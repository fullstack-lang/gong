const fs = require('fs');
const path = require('path');

// Target the specific test4 Angular output directory
const BUILD_DIR = path.join(__dirname, 'dist', 'ng-github.com-fullstack-lang-gong-test-test4', 'browser'); 
const OUTPUT_FILE = path.join(__dirname, '..', 'test4-portable-app.html');
console.log("Packaging Test4 Gong application into a single HTML file...");

// 1. Read the compiled WASM and convert to Base64
const wasmPath = path.join(BUILD_DIR, 'main.wasm');
if (!fs.existsSync(wasmPath)) {
    console.error("❌ Error: main.wasm not found. Did you compile Go to the public folder?");
    process.exit(1);
}
const wasmBuffer = fs.readFileSync(wasmPath);
const wasmBase64 = wasmBuffer.toString('base64');

// 2. Read the built index.html and wasm_exec.js
let html = fs.readFileSync(path.join(BUILD_DIR, 'index.html'), 'utf8');
const wasmExecJs = fs.readFileSync(path.join(BUILD_DIR, 'wasm_exec.js'), 'utf8');

// 3. Create the Base64 Bootloader
const base64Bootloader = `
<script>${wasmExecJs}</script>

<script>
  console.log("Initializing Test4 WASM Backend from Base64...");
  const base64String = "${wasmBase64}";
  const binaryString = window.atob(base64String);
  const bytes = new Uint8Array(binaryString.length);
  for (let i = 0; i < binaryString.length; i++) {
      bytes[i] = binaryString.charCodeAt(i);
  }
  
  const go = new Go();
  WebAssembly.instantiate(bytes.buffer, go.importObject)
    .then((result) => {
        go.run(result.instance);
        console.log("✅ Test4 WASM backend loaded successfully from memory!");
    })
    .catch(err => console.error("WASM Boot Error:", err));
</script>
`;

// 4. Inject the bootloader right before the closing </body> tag
html = html.replace('</body>', base64Bootloader + '\n</body>');

// 5. Inline Angular CSS styles (Using a robust Regex to catch all attributes)
html = html.replace(/<link[^>]*rel="stylesheet"[^>]*href="([^"]+)"[^>]*>/gi, (match, src) => {
    const cssPath = path.join(BUILD_DIR, src);
    if (fs.existsSync(cssPath)) {
        console.log(`css: Inlining ${src}`);
        return `<style>\n${fs.readFileSync(cssPath, 'utf8')}\n</style>`;
    }
    return match; 
});

// Optional: Angular sometimes wraps a fallback CSS link in a <noscript> tag. 
// We can safely delete it since we just embedded the CSS.
html = html.replace(/<noscript>[\s\S]*?<\/noscript>/gi, '');

// 5.5 Inline Favicon as a Base64 String
html = html.replace(/<link[^>]*rel="icon"[^>]*href="([^"]+)"[^>]*>/gi, (match, src) => {
    const faviconPath = path.join(BUILD_DIR, src);
    if (fs.existsSync(faviconPath)) {
        console.log(`favicon: Inlining ${src} as Base64`);
        const faviconBuffer = fs.readFileSync(faviconPath);
        const base64Favicon = faviconBuffer.toString('base64');
        return `<link rel="icon" type="image/x-icon" href="data:image/x-icon;base64,${base64Favicon}">`;
    }
    return match;
});

// 6. Inline Angular JS Modules (main.js, polyfills.js, etc.)
html = html.replace(/<script src="([^"]+)" type="module"><\/script>/g, (match, src) => {
    const jsPath = path.join(BUILD_DIR, src);
    if (fs.existsSync(jsPath)) {
        return `<script type="module">\n${fs.readFileSync(jsPath, 'utf8')}\n</script>`;
    }
    return match;
});

// Write the massive file to the root of the workspace
fs.writeFileSync(OUTPUT_FILE, html);
console.log("✅ Success! You can now double-click: " + OUTPUT_FILE);