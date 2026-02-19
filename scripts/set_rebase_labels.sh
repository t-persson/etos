#!/bin/bash
set -x

REPOSITORY="t-persson/etos"
LABEL="automated-rebase"
TIMEOUT=60  # Timeout in seconds (1 minute)

if [ ! -x "$(command -v gh)" ]; then
  echo "Error: gh CLI is not installed." >&2
  exit 1
fi

wait_for_status() {
  PR_NUM=$1
  START=$EPOCHSECONDS
  while true; do
    MERGEABLE=$(gh pr view -R $REPOSITORY $PR_NUM --json mergeable --jq ".mergeable")
    if [[ "$MERGEABLE" == "UNKNOWN" ]]; then
      sleep 5
      if (( EPOCHSECONDS - START > TIMEOUT )); then
        exit 1
      fi
      continue
    else
      echo "$MERGEABLE"
      break
    fi
  done
}

for NUM in `gh pr list -R $REPOSITORY 2>/dev/null | awk '{print $1}'`; do
  echo "Processing PR #$NUM..."
  MERGEABLE=$(wait_for_status $NUM)
  if [ $? -gt 0 ]; then
    echo "Failed to get mergeable status for PR #$NUM after $TIMEOUT seconds. Skipping."
    continue
  fi
  echo "PR #$NUM mergeable status: $MERGEABLE"

  has_label=$(gh pr view -R $REPOSITORY $NUM --json labels --jq ".labels | map(select(.name == \"$LABEL\")) | length > 0")
  if [ "$MERGEABLE" = "CONFLICTING" ]; then
    if [ "$has_label" = "false" ]; then
      echo "PR #$NUM is conflicting, enable automated rebase"
      gh pr edit -R $REPOSITORY --add-label "$LABEL" $NUM
    else
      echo "PR #$NUM is already marked for rebase. No action needed."
    fi
  else
    if [ "$has_label" = "true" ]; then
      echo "PR #$NUM does not have conflicts, disable automated rebase"
      gh pr edit -R $REPOSITORY --remove-label "$LABEL" $NUM
    else
      echo "PR #$NUM is not conflicting and does not have the rebase label. No action needed."
    fi
  fi
done
