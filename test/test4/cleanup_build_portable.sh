#!/bin/bash
set -e

echo "🚀 Starting Test4 Portable App Build Process..."

GO_DIR="go"
TMP_BUILD_DIR=".tmp_build"

# 0. CLEANUP (Crucial for Library updates!)
echo "🧹 Cleaning up previous build artifacts..."
rm -rf "$GO_DIR/cmd/test4/test4"
rm -f "test4-portable-app.html"
rm -f "test4-portable-app.zip"
rm -rf "$TMP_BUILD_DIR"

echo "✅ Success! Cleanup performed"