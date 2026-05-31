#!/bin/bash
set -e

echo "🚀 Starting Test4 Portable App Build Process..."

GO_DIR="go"
SPLIT_NG_DIR="../../lib/split/ng-github.com-fullstack-lang-gong-lib-split"
DIST_BROWSER_DIR="$SPLIT_NG_DIR/dist/ng-github.com-fullstack-lang-gong-lib-split/browser"
TMP_BUILD_DIR=".tmp_build"

# 0. CLEANUP (Crucial for Library updates!)
echo "🧹 0/5: Cleaning up previous build artifacts..."
rm -rf "$GO_DIR/cmd/test4/test4"
rm -rf "$GO_DIR/cmd/test4-wasm/test4-wasm"
rm -f "test4-portable-app.html"
rm -f "test4-portable-app.zip"
rm -rf "$TMP_BUILD_DIR"

# 1. Prepare temporary build directory
echo "📦 1/5: Preparing temporary build directory and copying frontend assets..."
mkdir -p "$TMP_BUILD_DIR"

if [ ! -d "$DIST_BROWSER_DIR" ]; then
    echo "❌ Error: Angular dist directory not found. Please build the Angular app first."
    exit 1
fi

cp -r "$DIST_BROWSER_DIR/"* "$TMP_BUILD_DIR/"

if [ -f "$(go env GOROOT)/lib/wasm/wasm_exec.js" ]; then
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$TMP_BUILD_DIR/"
else
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$TMP_BUILD_DIR/"
fi

# 2. Compile Go to WebAssembly
echo "🐹 2/5: Compiling Go backend to WebAssembly..."
cd "$GO_DIR"
# Use -a flag to force rebuild of all packages, preventing stale cache usage
GOOS=js GOARCH=wasm go build -a -o "../$TMP_BUILD_DIR/main.wasm" ./cmd/test4-wasm
cd ..

# 3. Run the Node Bundler
echo "⚙️ 3/5: Bundling everything into a single HTML file..."
node bundle.js

# 4. Zip the HTML file
echo "🗜️ 4/5: Zipping the portable app..."
zip test4-portable-app.zip test4-portable-app.html

# 5. Cleanup temporary build directory
echo "🧹 5/5: Cleaning up temporary build directory..."
# rm -rf "$TMP_BUILD_DIR"

echo "✅ Success! Your offline app is ready at: test4-portable-app.html and test4-portable-app.zip"