#!/bin/bash

# Simple Release Script
# Usage: ./scripts/simple-release.sh 1.2.3

set -e

if [ $# -eq 0 ]; then
    echo "Usage: ./scripts/simple-release.sh <version>"
    echo "Example: ./scripts/simple-release.sh 1.2.3"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "❌ Invalid version format. Use semantic versioning (e.g., 1.2.3)"
    exit 1
fi

echo "🚀 Releasing version $VERSION"

# Check if tag already exists
if git tag -l | grep -q "^v$VERSION$"; then
    echo "❌ Version $VERSION already exists as a git tag"
    exit 1
fi

# Update root package.json
echo "📝 Updating root package.json..."
npm version $VERSION --no-git-tag-version

# Update TypeScript package.json
echo "📝 Updating packages/ts/package.json..."
sed -i.bak "s/\"version\": \"[^\"]*\"/\"version\": \"$VERSION\"/" packages/ts/package.json
rm -f packages/ts/package.json.bak

# Verify versions match
ROOT_VERSION=$(node -p "require('./package.json').version")
TS_VERSION=$(node -p "require('./packages/ts/package.json').version")

if [ "$ROOT_VERSION" != "$VERSION" ] || [ "$TS_VERSION" != "$VERSION" ]; then
    echo "❌ Version mismatch after update!"
    echo "Root: $ROOT_VERSION, TS: $TS_VERSION, Expected: $VERSION"
    exit 1
fi

echo "✅ Both package.json files updated to $VERSION"

# Commit and tag
echo "📦 Committing and tagging..."
git add package.json packages/ts/package.json
git commit -m "Release version $VERSION

🤖 Generated with simple-release script"

git tag "v$VERSION"

# Push
echo "🚀 Pushing to GitHub..."
git push origin main
git push origin "v$VERSION"

# Create and push Go module tag
echo "📦 Creating Go module tag..."
git tag "codegen/go/v$VERSION" "v$VERSION"
git push origin "codegen/go/v$VERSION"

echo ""
echo "✅ Release $VERSION completed!"
echo "📋 Monitor at: https://github.com/rakeyshgidwani/sunday-schemas/actions"
echo "📦 Packages will be available at:"
echo "   npm: @rakeyshgidwani/sunday-schemas@$VERSION"
echo "   Go:  github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"