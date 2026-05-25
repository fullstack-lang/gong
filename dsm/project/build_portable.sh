#!/bin/bash
set -e

echo "🚀 Starting DSM Project Portable App Build Process..."

GO_DIR="go"
SPLIT_NG_DIR="../../lib/split/ng-github.com-fullstack-lang-gong-lib-split"
LOCAL_DIST_DIR="dist"

# 0. CLEANUP
echo "🧹 0/5: Cleaning up previous build artifacts..."
rm -rf dist/
rm -f "project-portable-app.html"

# 1. Build the Angular Application in lib/split
echo "🅰️ 1/5: Building Angular frontend in lib/split..."
cd "$SPLIT_NG_DIR"
npm run build
cd - > /dev/null

# 2. Copy the built frontend locally
echo "📦 2/5: Copying frontend to local dist folder..."
mkdir -p "$LOCAL_DIST_DIR"
cp -r "$SPLIT_NG_DIR/dist/ng-github.com-fullstack-lang-gong-lib-split/browser/"* "$LOCAL_DIST_DIR/"

# 3. Copy the Go WebAssembly glue code
echo "📦 3/5: Copying wasm_exec.js directly to output folder..."
if [ -f "$(go env GOROOT)/lib/wasm/wasm_exec.js" ]; then
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$LOCAL_DIST_DIR/"
else
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$LOCAL_DIST_DIR/"
fi

# 4. Compile Go to WebAssembly
echo "🐹 4/5: Compiling Go backend to WebAssembly..."
cd "$GO_DIR"
GOOS=js GOARCH=wasm go build -o "../$LOCAL_DIST_DIR/main.wasm" ./cmd/project
cd ..

# 5. Run the Node Bundler
echo "⚙️ 5/5: Bundling everything into a single HTML file..."
node go/cmd/project/bundle.js

echo "✅ Success! Your offline app is ready at: project-portable-app.html"