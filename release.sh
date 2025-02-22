#!/bin/bash

set -e

VERSION_FILE="VERSION"

if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "Uncommitted changes detected. Please commit or stash them before releasing."
    exit 1
fi

if [[ ! -f "$VERSION_FILE" ]]; then 
    echo "0.1.0" > "$VERSION_FILE"
fi

VERSION=$(cat "$VERSION_FILE")
echo "Current version: $VERSION"

NEW_VERSION=$(echo "$VERSION" | awk -F. -v OFS=. '{$NF += 1 ; print}')
echo "Next version: $NEW_VERSION"

echo "$NEW_VERSION" > "$VERSION_FILE"

echo "Git tagging..."
git add "$VERSION_FILE"
git commit -m "Release v$NEW_VERSION"

git tag "v$NEW_VERSION"

git push origin main --tags

echo "Version bumped to $NEW_VERSION, commited, tagged and pushed to origin"
