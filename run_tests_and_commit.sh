#!/bin/bash
set -e

echo "Formatting codebase..."
go fmt ./...

echo "Running tests..."
go test ./...

echo "Staging files..."
git add .

echo "Committing..."
git commit --no-gpg-sign -m "Add advanced data structures, concurrency, cache, and optional types"

echo "Done!"
