# Copyright Axis Communications AB.
#
# For a full list of individual contributors, please see the commit history.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Label automated integration PRs in source-commit order using the GitHub CLI."""

import argparse
import base64
import json
import re
import subprocess
from dataclasses import dataclass
from typing import Any
from urllib.parse import quote

NEXT = "integration-next"
BLOCKED = "integration-blocked"
SOURCE_OWNER = "eiffel-community"
SOURCES = {
    "etos-api": ("api", "sse", "log_area"),
    "etos-suite-runner": ("suite_runner", "log_listener"),
    "etos-suite-starter": ("suite_starter",),
    "etos-test-runner": ("test_runner",),
}
LABELS = {
    NEXT: ("0e8a16", "Next source update in order; normal merge requirements still apply"),
    BLOCKED: ("d93f0b", "An earlier source update is missing or source order is unresolved"),
}


class GitHub:
    """Access GitHub through gh, inheriting authentication from GH_TOKEN."""

    @staticmethod
    def get(endpoint: str, paginate: bool = False) -> Any:
        """Read a REST resource or every page of a REST collection.

        :param endpoint: Relative GitHub API endpoint.
        :param paginate: Return a list of all response pages.
        :return: Decoded JSON response.
        :raises subprocess.CalledProcessError: If GitHub rejects the request.
        """
        command = ["gh", "api", endpoint]
        if paginate:
            command.extend(["--paginate", "--jq", "tojson"])
        output = subprocess.run(command, capture_output=True, text=True, check=True).stdout
        return (
            [json.loads(page) for page in output.splitlines()] if paginate else json.loads(output)
        )

    @staticmethod
    def write(endpoint: str, method: str, data: dict[str, Any] | None = None) -> None:
        """Write a label resource without evaluating any PR content as shell code.

        :param endpoint: Relative GitHub API endpoint.
        :param method: HTTP mutation method.
        :param data: Optional JSON request body.
        :raises subprocess.CalledProcessError: If GitHub rejects the request.
        """
        command = ["gh", "api", "--method", method, endpoint]
        if data is not None:
            command.extend(["--input", "-"])
        subprocess.run(
            command,
            input=json.dumps(data) if data is not None else None,
            capture_output=True,
            text=True,
            check=True,
        )

    def collection(self, endpoint: str) -> list[dict[str, Any]]:
        """Read every entry of an array-valued REST collection.

        :param endpoint: Relative GitHub API collection endpoint.
        :return: Entries from all pages.
        """
        return [entry for page in self.get(endpoint, paginate=True) for entry in page]


@dataclass
class Decision:
    """Describe the label and explanation for one integration PR."""

    pull: dict[str, Any]
    label: str
    reason: str


def _resolve(github: GitHub, source: str, reference: str) -> str:
    """Resolve a source reference to its full commit SHA.

    :param github: GitHub REST client.
    :param source: Child repository in OWNER/REPO format.
    :param reference: Commit SHA, abbreviated SHA, branch, or tag to resolve.
    :return: Full source commit SHA.
    :raises subprocess.CalledProcessError: If GitHub cannot resolve the reference.
    """
    return github.get(f"repos/{source}/commits/{quote(reference, safe='')}")["sha"]


def _baseline(github: GitHub, repository: str, base: str, source: str) -> str:
    """Resolve the source commit already integrated into the target defaults.

    :param github: GitHub REST client.
    :param repository: Target ETOS repository in OWNER/REPO format.
    :param base: Target main branch snapshot SHA.
    :param source: Child repository name, without its owner.
    :return: Source commit shared by all component defaults.
    :raises ValueError: If a version is missing or component versions disagree.
    :raises subprocess.CalledProcessError: If a default or source reference cannot be read.
    """
    commits = set()
    for name in SOURCES[source]:
        content = github.get(f"repos/{repository}/contents/defaults/{name}.yaml?ref={base}")
        text = base64.b64decode(content["content"]).decode()
        version = re.search(r"""^version:\s*['"]?([a-zA-Z0-9_.-]+)['"]?\s*$""", text, re.MULTILINE)
        if not version:
            raise ValueError(f"Cannot read version from defaults/{name}.yaml")
        commits.add(_resolve(github, f"{SOURCE_OWNER}/{source}", version[1]))
    if len(commits) != 1:
        raise ValueError(f"Defaults for {source} do not reference the same source commit")
    return commits.pop()


def _history(github: GitHub, source: str, baseline: str) -> list[str]:
    """Walk source main's first-parent commits after the integrated baseline.

    :param github: GitHub REST client.
    :param source: Child repository in OWNER/REPO format.
    :param baseline: Previously integrated source commit SHA.
    :return: Unintegrated mainline commit SHAs, oldest first.
    :raises ValueError: If the baseline is not on source main's first-parent history.
    :raises subprocess.CalledProcessError: If source history cannot be read.
    """
    head = _resolve(github, source, "main")
    pages = github.get(f"repos/{source}/compare/{baseline}...{head}?per_page=100", paginate=True)
    if pages[0]["status"] == "identical":
        return []
    if pages[0]["status"] != "ahead":
        raise ValueError(f"Integrated version is not an ancestor of {source}/main")
    commits = {commit["sha"]: commit for page in pages for commit in page["commits"]}
    history = []
    current = head
    # Side-branch commits must not become separate integration prerequisites.
    while current != baseline:
        if current not in commits or not commits[current]["parents"]:
            raise ValueError(f"Integrated version is not on {source}/main's first-parent history")
        history.append(current)
        current = commits[current]["parents"][0]["sha"]
    return list(reversed(history))


def _source_commit(github: GitHub, source: str, pull: dict[str, Any]) -> str:
    """Resolve the unique source commit or release link from an integration PR.

    :param github: GitHub REST client.
    :param source: Expected child repository in OWNER/REPO format.
    :param pull: Integration PR containing publisher metadata in its body.
    :return: Full commit SHA referenced by the publisher.
    :raises ValueError: If exactly one matching source link cannot be found.
    :raises subprocess.CalledProcessError: If the linked reference cannot be resolved.
    """
    prefix = rf"https://github\.com/{re.escape(source)}"
    if source.endswith("/etos-test-runner"):
        pattern = rf"^Release: {prefix}/releases/tag/([a-zA-Z0-9_.-]+)\s*$"
    else:
        pattern = rf"^Commit: {prefix}/commit/([a-fA-F0-9]{{40}})\s*$"
    matches = re.findall(pattern, pull["body"] or "", re.MULTILINE)
    if len(matches) != 1:
        raise ValueError(f"PR #{pull['number']} needs exactly one source commit or release link")
    return _resolve(github, source, matches[0])


def _prerequisites(github: GitHub, source: str, history: list[str]) -> set[str]:
    """Find source merges or published releases that require integration.

    :param github: GitHub REST client.
    :param source: Child repository in OWNER/REPO format.
    :param history: Unintegrated first-parent source commit SHAs.
    :return: Source commit SHAs that must be integrated in history order.
    :raises ValueError: If a published release tag cannot be resolved.
    :raises subprocess.CalledProcessError: If source PRs, releases, or tags cannot be read.
    """
    if source.endswith("/etos-test-runner"):
        releases = github.collection(f"repos/{source}/releases?per_page=100")
        tags = {
            tag["name"]: tag["commit"]["sha"]
            for tag in github.collection(f"repos/{source}/tags?per_page=100")
        }
        commits = set()
        for release in releases:
            if not release["draft"] and release["published_at"]:
                tag = release["tag_name"]
                if tag not in tags:
                    raise ValueError(f"Published release tag {source}/{tag} cannot be resolved")
                commits.add(tags[tag])
        return commits.intersection(history)
    commits = set()
    for commit in history:
        pulls = github.collection(f"repos/{source}/commits/{commit}/pulls?per_page=100")
        if any(
            pull["merged_at"]
            and pull["base"]["ref"] == "main"
            and pull["base"]["repo"]["full_name"] == source
            and pull["merge_commit_sha"] == commit
            for pull in pulls
        ):
            commits.add(commit)
    return commits


def source_decisions(
    github: GitHub, repository: str, base: str, source: str, pulls: list[dict[str, Any]]
) -> list[Decision]:
    """Find the next unintegrated source update, including missing integration PRs.

    :param github: GitHub REST client.
    :param repository: Target ETOS repository.
    :param base: Snapshot of the target main branch.
    :param source: Child repository name, without its owner.
    :param pulls: Open automated PRs for this child repository.
    :return: A label decision for every supplied PR.
    :raises ValueError: If source metadata or ancestry is ambiguous.
    """
    baseline = _baseline(github, repository, base, source)
    source = f"{SOURCE_OWNER}/{source}"
    history = _history(github, source, baseline)
    commits = {pull["number"]: _source_commit(github, source, pull) for pull in pulls}
    required = _prerequisites(github, source, history)
    # Explicit manual builds on main are also ordered, even without a merged source PR.
    required.update(commit for commit in commits.values() if commit in history)
    expected = next((commit for commit in history if commit in required), None)
    candidates = [number for number, commit in commits.items() if commit == expected]
    decisions = []
    for pull in pulls:
        commit = commits[pull["number"]]
        if commit not in history:
            reason = "Source commit is already integrated or is not on source main"
        elif commit != expected:
            reason = f"Waiting for source commit {expected}"
        elif len(candidates) != 1:
            reason = "Multiple open integration PRs reference the next source commit"
        elif pull["draft"]:
            reason = "Next integration PR is a draft"
        else:
            decisions.append(Decision(pull, NEXT, f"Next source commit is {commit}"))
            continue
        decisions.append(Decision(pull, BLOCKED, reason))
    return decisions


def _apply(github: GitHub, repository: str, decisions: list[Decision]) -> None:
    """Apply queue labels, revoking stale eligibility before promoting successors.

    :param github: GitHub REST client.
    :param repository: Target ETOS repository in OWNER/REPO format.
    :param decisions: Desired labels and current PR label snapshots.
    :raises subprocess.CalledProcessError: If labels cannot be read or changed.
    """
    existing = {
        label["name"] for label in github.collection(f"repos/{repository}/labels?per_page=100")
    }
    for label, (color, description) in LABELS.items():
        if label not in existing:
            github.write(
                f"repos/{repository}/labels",
                "POST",
                {"name": label, "color": color, "description": description},
            )
    # Revoke stale eligibility everywhere before promoting any successor.
    for decision in decisions:
        labels = {label["name"] for label in decision.pull["labels"]}
        endpoint = f"repos/{repository}/issues/{decision.pull['number']}/labels"
        if decision.label != NEXT and NEXT in labels:
            github.write(f"{endpoint}/{NEXT}", "DELETE")
        if decision.label == BLOCKED and BLOCKED not in labels:
            github.write(endpoint, "POST", {"labels": [BLOCKED]})
        if not decision.label and BLOCKED in labels:
            github.write(f"{endpoint}/{BLOCKED}", "DELETE")
    for decision in decisions:
        if decision.label == NEXT:
            labels = {label["name"] for label in decision.pull["labels"]}
            endpoint = f"repos/{repository}/issues/{decision.pull['number']}/labels"
            if BLOCKED in labels:
                github.write(f"{endpoint}/{BLOCKED}", "DELETE")
            if NEXT not in labels:
                github.write(endpoint, "POST", {"labels": [NEXT]})


def _report_error(message: str, error: ValueError | subprocess.CalledProcessError) -> None:
    """Print an evaluation error and any GitHub CLI diagnostic output.

    :param message: Context describing the failed operation.
    :param error: Metadata validation or GitHub CLI error to report.
    """
    print(f"{message}: {error}")
    if isinstance(error, subprocess.CalledProcessError):
        print(error.stderr)


def reconcile(github: GitHub, repository: str, dry_run: bool = False) -> bool:
    """Reconcile every source queue, blocking a queue on lookup or metadata errors.

    :param github: GitHub REST client.
    :param repository: Target ETOS repository.
    :param dry_run: Report decisions without writing labels.
    :return: Whether every source queue could be evaluated successfully.
    """
    base_endpoint = f"repos/{repository}/git/ref/heads/main"
    base = github.get(base_endpoint)["object"]["sha"]
    pulls = github.collection(f"repos/{repository}/pulls?state=open&base=main&per_page=100")
    groups: dict[str, list[dict[str, Any]]] = {}
    decisions = []
    success = True
    for pull in pulls:
        labels = {label["name"] for label in pull["labels"]}
        sources = labels.intersection(SOURCES)
        if "automated-pr" in labels and sources:
            if len(sources) == 1:
                groups.setdefault(sources.pop(), []).append(pull)
            else:
                decisions.append(Decision(pull, BLOCKED, "Multiple source repository labels"))
                success = False
        elif labels.intersection(LABELS):
            decisions.append(Decision(pull, "", "No longer a managed integration PR"))
    for source, pulls in groups.items():
        try:
            decisions.extend(source_decisions(github, repository, base, source, pulls))
        except (ValueError, subprocess.CalledProcessError) as error:
            _report_error(f"Could not evaluate {source}", error)
            decisions.extend(
                Decision(pull, BLOCKED, "Source order could not be established") for pull in pulls
            )
            success = False
    try:
        if github.get(base_endpoint)["object"]["sha"] != base:
            raise ValueError("Target main changed during reconciliation")
    except (ValueError, subprocess.CalledProcessError) as error:
        _report_error("Could not confirm target main", error)
        decisions = [
            (
                Decision(decision.pull, BLOCKED, "Target main could not be confirmed unchanged")
                if decision.label
                else decision
            )
            for decision in decisions
        ]
        success = False
    for decision in decisions:
        print(f"PR #{decision.pull['number']}: {decision.label or 'unmanaged'}: {decision.reason}")
    if not dry_run:
        _apply(github, repository, decisions)
    return success


def main() -> None:
    """Run label reconciliation for the selected repository."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("repository", help="Target repository in OWNER/REPO format")
    parser.add_argument("--dry-run", action="store_true", help="Report without changing GitHub")
    args = parser.parse_args()
    if not re.fullmatch(r"[a-zA-Z0-9_.-]+/[a-zA-Z0-9_.-]+", args.repository):
        parser.error("repository must be in OWNER/REPO format")
    raise SystemExit(0 if reconcile(GitHub(), args.repository, args.dry_run) else 1)


if __name__ == "__main__":
    main()
