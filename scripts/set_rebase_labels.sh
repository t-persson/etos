#!/bin/bash
set -x

REPOSITORY="t-persson/etos"
LABEL="automated-rebase"

if [ ! -x "$(command -v gh)" ]; then
  echo "Error: gh CLI is not installed." >&2
  exit 1
fi

JSON=$(gh pr status -R "$REPOSITORY" -c --json mergeable --json number --jq ".createdBy")

echo "$JSON" | jq -c '.[]' | while read i; do
  mergeable=$(echo $i | jq -r '.mergeable')
  number=$(echo $i | jq -r '.number')
  # Check if the "$LABEL" label is already present
  has_label=$(gh pr view -R "$REPOSITORY" "$number" --json labels --jq '.labels | map(select(.name == "$LABEL")) | length > 0')
  if [ "$mergeable" = "CONFLICTING" ]; then
    if [ "$has_label" = "true" ]; then
      echo "PR #$number is not mergeable but has the '$LABEL' label."
      success=$(gh pr edit -R "$REPOSITORY" --remove-label "$LABEL" "$number")
      if [ $? -eq 0 ]; then
        echo "Removed '$LABEL' label to PR #$number."
      fi
    fi
  else
    if [ ! "$has_label" = "true" ]; then
      echo "PR #$number is mergeable but does not have the '$LABEL' label."
      success=$(gh pr edit -R "$REPOSITORY" --add-label "$LABEL" "$number")
      if [ $? -eq 0 ]; then
        echo "Added '$LABEL' label to PR #$number."
      fi
    fi
  fi
done
