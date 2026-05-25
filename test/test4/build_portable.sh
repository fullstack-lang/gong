#!/bin/bash
set -e

echo "🚀 Starting Test4 Portable App Build Process..."

GO_DIR="go"
NG_DIR="ng-github.com-fullstack-lang-gong-test-test4"
DIST_BROWSER_DIR="$NG_DIR/dist/ng-github.com-fullstack-lang-gong-test-test4/browser"

# 0. CLEANUP (Crucial for Library updates!)
echo "🧹 0/5: Cleaning up previous build artifacts..."
rm -rf "$NG_DIR/dist/"
rm -rf "$NG_DIR/.angular/cache/"
rm -f "test4-portable-app.html"
rm -f "test4-portable-app.zip"

# 1. Copy the Go WebAssembly glue code
echo "📦 1/5: Copying wasm_exec.js to Angular public folder..."
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$NG_DIR/public/"

# 2. Build the Angular Application
echo "🅰️ 2/5: Building Angular frontend for production..."
cd "$NG_DIR"

# ---> NEW CRITICAL STEP: Build the library first to apply your WebSocket changes
npx ng build test4 

# ---> Build the main app (which will now pull the fresh library)
npm run build
cd ..

# 3. Compile Go to WebAssembly
echo "🐹 3/5: Compiling Go backend to WebAssembly..."
cd "$GO_DIR"
GOOS=js GOARCH=wasm go build -o "../$DIST_BROWSER_DIR/main.wasm" ./cmd/test4
cd ..

# 4. Run the Node Bundler
echo "⚙️ 4/5: Bundling everything into a single HTML file..."
node bundle.js

# 5. Zip the HTML file
echo "🗜️ 5/5: Zipping the portable app..."
zip test4-portable-app.zip test4-portable-app.html

echo "✅ Success! Your offline app is ready at: test4-portable-app.html and test4-portable-app.zip"