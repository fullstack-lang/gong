#!/bin/bash
set -e

echo "🚀 Starting Test4 Portable App Build Process..."

GO_DIR="go"
NG_DIR="ng-github.com-fullstack-lang-gong-test-test4"
DIST_BROWSER_DIR="$NG_DIR/dist/ng-github.com-fullstack-lang-gong-test-test4/browser"

# 0. CLEANUP (Crucial for Library updates!)
echo "🧹 0/5: Cleaning up previous build artifacts..."
rm -rf "$GO_DIR/cmd/test4/test4"
rm -rf "$NG_DIR/dist/"
rm -rf "$NG_DIR/.angular/cache/"
rm -f "test4-portable-app.html"

echo "✅ Success! Cleanup performed"