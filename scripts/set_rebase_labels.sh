#!/bin/bash
set -x

REPOSITORY="t-persson/etos"
LABEL="automated-rebase"

if [ ! -x "$(command -v gh)" ]; then
  echo "Error: gh CLI is not installed." >&2
  exit 1
fi

for num in `gh pr list -R $REPOSITORY 2>/dev/null | awk '{print $1}'`; do
  MERGEABLE=$(gh pr view -R $REPOSITORY $num --json mergeable --jq ".mergeable")
  echo "PR #$num mergeable status: $MERGEABLE"
  has_label=$(gh pr view -R $REPOSITORY $num --json labels --jq ".labels | map(select(.name == \"$LABEL\")) | length > 0")
  if [ "$MERGEABLE" = "CONFLICTING" ]; then
    echo "PR #$num is not mergeable. Checking for '$LABEL' label..."
    if [ "$has_label" = "true" ]; then
      echo "PR #$num has the '$LABEL' label but is not mergeable. Removing label..."
      gh pr edit -R $REPOSITORY --remove-label "$LABEL" $num
    else
      echo "PR #$num does not have the '$LABEL' label and is not mergeable. No action needed."
    fi
  else
    echo "PR #$num is mergeable. Checking for '$LABEL' label..."
    if [ "$has_label" = "false" ]; then
      echo "PR #$num does not have the '$LABEL' label but is mergeable. Adding label..."
      gh pr edit -R $REPOSITORY --add-label "$LABEL" $num
    else
      echo "PR #$num has the '$LABEL' label and is mergeable. No action needed."
    fi
  fi
done
