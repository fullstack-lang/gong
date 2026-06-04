package main

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fullstack-lang/gong/lib/split"
)

func main() {
	if len(os.Args) > 1 {
		targetDir := os.Args[1]
		if err := os.Chdir(targetDir); err != nil {
			fmt.Printf("❌ Error changing to directory %s: %v\n", targetDir, err)
			os.Exit(1)
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("❌ Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	pkgName := filepath.Base(cwd)
	outputFile := fmt.Sprintf("%s-app-portable.html", pkgName)
	zipFile := fmt.Sprintf("%s-app-portable.zip", pkgName)

	// =========================================================================
	// 1. Bash Script Equivalent (Compilation)
	// =========================================================================
	fmt.Printf("🚀 Step 1: Building project %s...\n", pkgName)

	// 0. Cleanup previous builds and recreate .tmp_build
	os.RemoveAll(".tmp_build")
	os.Remove(outputFile)
	os.Remove(zipFile)
	os.MkdirAll(".tmp_build", 0o755)

	// 1. Copy Frontend Assets from the embedded lib/split directory
	baseEmbedPath := "ng-github.com-fullstack-lang-gong-lib-split/dist/ng-github.com-fullstack-lang-gong-lib-split"
	embedPath := baseEmbedPath + "/browser"
	if err := copyFS(split.NgDistNg, embedPath, ".tmp_build"); err != nil {
		fmt.Printf("❌ Error extracting embedded Angular files: %v\n", err)
		os.Exit(1)
	}

	// Copy wasm_exec.js (Handles both older and newer Go versions)
	goRootBytes, _ := exec.Command("go", "env", "GOROOT").Output()
	goRoot := strings.TrimSpace(string(goRootBytes))
	wasmExecPath := filepath.Join(goRoot, "misc", "wasm", "wasm_exec.js")
	if _, err := os.Stat(wasmExecPath); os.IsNotExist(err) {
		wasmExecPath = filepath.Join(goRoot, "lib", "wasm", "wasm_exec.js")
	}
	copyFile(wasmExecPath, filepath.Join(".tmp_build", "wasm_exec.js"))

	// Build the Go Wasm binary:
	runWasmBuild("go", "build", "-a", "-o", ".tmp_build/main.wasm", ".")

	// =========================================================================
	// 2. bundle.js Equivalent (Packaging)
	// =========================================================================
	fmt.Printf("\n📦 Step 2: Packaging %s Gong application into a single HTML file...\n", pkgName)

	buildDir := ".tmp_build"

	// Read main.wasm and encode to Base64
	wasmPath := filepath.Join(buildDir, "main.wasm")
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		fmt.Printf("❌ Error: main.wasm not found at %s. Did you compile Go to the public folder?\n", wasmPath)
		os.Exit(1)
	}
	wasmBase64 := base64.StdEncoding.EncodeToString(wasmBytes)

	// Read index.html and wasm_exec.js
	htmlBytes, err := os.ReadFile(filepath.Join(buildDir, "index.html"))
	if err != nil {
		fmt.Printf("❌ Error reading index.html: %v\n", err)
		os.Exit(1)
	}
	html := string(htmlBytes)

	wasmExecBytes, err := os.ReadFile(filepath.Join(buildDir, "wasm_exec.js"))
	if err != nil {
		fmt.Printf("❌ Error reading wasm_exec.js: %v\n", err)
		os.Exit(1)
	}
	wasmExecJs := string(wasmExecBytes)

	// Create the Base64 Bootloader
	bootloader := fmt.Sprintf(`
<div id="wasm-progress-container" style="position: fixed; top: 50%%; left: 50%%; transform: translate(-50%%, -50%%); font-family: sans-serif; text-align: center; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 4px 10px rgba(0,0,0,0.2); z-index: 9999;">
    <h3 style="margin-top: 0;">Loading Application...</h3>
    <div style="margin-bottom: 10px;">Processing WASM: <span id="wasm-progress-text">0%%</span></div>
    <div style="width: 300px; height: 20px; background: #eee; border-radius: 10px; overflow: hidden; margin: 0 auto;">
        <div id="wasm-progress-bar" style="width: 0%%; height: 100%%; background: #007bff; transition: width 0.1s;"></div>
    </div>
</div>

<script>%s</script>

<script id="wasm-base64-data" type="text/plain">%s</script>

<script>
  console.log("Initializing %s WASM Backend from Base64...");
  const base64String = document.getElementById('wasm-base64-data').textContent.trim();
  
  async function processBase64InChunks(base64Str) {
      let padding = 0;
      if (base64Str.endsWith('==')) padding = 2;
      else if (base64Str.endsWith('=')) padding = 1;
      
      const totalLength = base64Str.length;
      const byteLength = (totalLength * 3) / 4 - padding;
      const bytes = new Uint8Array(byteLength);
      
      const chunkSize = 4 * 1024 * 256; 
      let byteOffset = 0;
      let stringOffset = 0;
      
      const progressText = document.getElementById('wasm-progress-text');
      const progressBar = document.getElementById('wasm-progress-bar');

      return new Promise((resolve) => {
          function processChunk() {
              const end = Math.min(stringOffset + chunkSize, totalLength);
              const chunk = base64Str.substring(stringOffset, end);
              const binaryString = window.atob(chunk);
              const len = binaryString.length;
              
              for (let i = 0; i < len; i++) {
                  bytes[byteOffset++] = binaryString.charCodeAt(i);
              }
              stringOffset = end;
              
              const percentage = Math.round((stringOffset / totalLength) * 100);
              if (progressText) progressText.innerText = percentage + '%%';
              if (progressBar) progressBar.style.width = percentage + '%%';
              
              if (stringOffset < totalLength) {
                  // 2. Yield to the browser's rendering engine so it actually paints the progress bar
                  requestAnimationFrame(() => {
                      setTimeout(processChunk, 0);
                  });
              } else {
                  resolve(bytes);
              }
          }
          
          processChunk();
      });
  }

  processBase64InChunks(base64String).then((bytes) => {
      const progressText = document.getElementById('wasm-progress-text');
      if (progressText) {
          progressText.innerText = "Compiling WASM (this may take a moment)...";
      }

      // 3. Small timeout to allow the browser to paint the "Compiling..." text before blocking the main thread
      setTimeout(() => {
          const go = new Go();
          WebAssembly.instantiate(bytes.buffer, go.importObject)
            .then((result) => {
                // 4. Hide the progress bar ONLY when the app is completely instantiated and ready
                const progressContainer = document.getElementById('wasm-progress-container');
                if (progressContainer) {
                    progressContainer.style.display = 'none';
                }
                go.run(result.instance);
                console.log("✅ %s WASM backend loaded successfully from memory!");
            })
            .catch(err => console.error("WASM Boot Error:", err));
      }, 50);
  });
</script>
`, wasmExecJs, wasmBase64, pkgName, pkgName)

	// Inject the bootloader right before the closing </body> tag
	html = strings.Replace(html, "</body>", bootloader+"\n</body>", 1)

	// Inject history API patch in <head> to prevent Angular SecurityError on file:///
	fileProtocolPatch := `
<script>
  if (window.location.protocol === 'file:') {
    const originalReplaceState = history.replaceState;
    history.replaceState = function() {
      try { originalReplaceState.apply(this, arguments); } catch (e) {}
    };
    const originalPushState = history.pushState;
    history.pushState = function() {
      try { originalPushState.apply(this, arguments); } catch (e) {}
    };
  }
</script>
`
	html = strings.Replace(html, "<head>", "<head>\n"+fileProtocolPatch, 1)

	// Inline Angular CSS styles AND embed all referenced fonts/assets as Base64
	cssLinkRe := regexp.MustCompile(`(?i)<link[^>]*rel="stylesheet"[^>]*href="([^"]+)"[^>]*>`)
	urlRe := regexp.MustCompile(`(?i)url\((?:['"]?)([^'"\)]+)(?:['"]?)\)`)

	html = cssLinkRe.ReplaceAllStringFunc(html, func(match string) string {
		src := cssLinkRe.FindStringSubmatch(match)[1]
		cssPath := filepath.Join(buildDir, src)

		cssBytes, err := os.ReadFile(cssPath)
		if err == nil {
			fmt.Printf("css: Inlining %s\n", src)
			cssContent := string(cssBytes)

			// Replace urls inside css
			cssContent = urlRe.ReplaceAllStringFunc(cssContent, func(urlMatch string) string {
				assetPath := urlRe.FindStringSubmatch(urlMatch)[1]

				if strings.HasPrefix(assetPath, "data:") || strings.HasPrefix(assetPath, "http") {
					return urlMatch
				}

				cleanAssetPath := strings.Split(strings.Split(assetPath, "?")[0], "#")[0]
				absAssetPath := filepath.Join(buildDir, cleanAssetPath)

				assetBytes, err := os.ReadFile(absAssetPath)
				if err == nil {
					fmt.Printf(" ↳ font/asset: Inlining %s as Base64\n", cleanAssetPath)
					ext := strings.ToLower(filepath.Ext(cleanAssetPath))
					mime := "application/octet-stream"
					switch ext {
					case ".woff2":
						mime = "font/woff2"
					case ".woff":
						mime = "font/woff"
					case ".ttf":
						mime = "font/ttf"
					case ".eot":
						mime = "application/vnd.ms-fontobject"
					case ".svg":
						mime = "image/svg+xml"
					case ".png":
						mime = "image/png"
					case ".jpg", ".jpeg":
						mime = "image/jpeg"
					}

					b64Str := base64.StdEncoding.EncodeToString(assetBytes)
					return fmt.Sprintf(`url("data:%s;base64,%s")`, mime, b64Str)
				}
				fmt.Printf(" ⚠️ Warning: Could not find asset to inline: %s\n", absAssetPath)
				return urlMatch
			})
			return fmt.Sprintf("<style>\n%s\n</style>", cssContent)
		}
		return match
	})

	// Remove <noscript> tags
	noscriptRe := regexp.MustCompile(`(?is)<noscript>.*?</noscript>`)
	html = noscriptRe.ReplaceAllString(html, "")

	// Inline Favicon as a Base64 String
	faviconRe := regexp.MustCompile(`(?i)<link[^>]*rel="icon"[^>]*href="([^"]+)"[^>]*>`)
	html = faviconRe.ReplaceAllStringFunc(html, func(match string) string {
		src := faviconRe.FindStringSubmatch(match)[1]
		faviconBytes, err := os.ReadFile(filepath.Join(buildDir, src))
		if err == nil {
			fmt.Printf("favicon: Inlining %s as Base64\n", src)
			return fmt.Sprintf(`<link rel="icon" type="image/x-icon" href="data:image/x-icon;base64,%s">`, base64.StdEncoding.EncodeToString(faviconBytes))
		}
		return match
	})

	// Inline Angular JS Modules
	jsModuleRe := regexp.MustCompile(`(?i)<script src="([^"]+)" type="module"><\/script>`)
	html = jsModuleRe.ReplaceAllStringFunc(html, func(match string) string {
		src := jsModuleRe.FindStringSubmatch(match)[1]
		jsBytes, err := os.ReadFile(filepath.Join(buildDir, src))
		if err == nil {
			fmt.Printf("js: Inlining %s\n", src)
			return fmt.Sprintf("<script type=\"module\">\n%s\n</script>", string(jsBytes))
		}
		return match
	})

	// Write the final bundled HTML
	err = os.WriteFile(outputFile, []byte(html), 0o644)
	if err != nil {
		fmt.Printf("❌ Error writing output file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n✅ Success! You can now double-click: %s\n", outputFile)

	// Zip the HTML file
	if err := createZip(zipFile, outputFile); err != nil {
		fmt.Printf("❌ Error zipping output file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("🗜️ Zipped output to: %s\n", zipFile)

	// Clean up temporary build directory
	fmt.Println("🧹 Cleaning up temporary build directory...")
	os.RemoveAll(".tmp_build")
}

// Helper: Run standard build commands
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Error running command '%s %s': %v\n", name, strings.Join(args, " "), err)
		os.Exit(1)
	}
}

// Helper: Create Zip
func createZip(zipName, fileName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToZip, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = fileName
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

// Helper: Copy a file
func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// Helper: Copy files from an embed.FS (or any fs.FS) to a local directory
func copyFS(srcFS fs.FS, srcDir string, dst string) error {
	return fs.WalkDir(srcFS, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(srcDir, path)
		dstPath := filepath.Join(dst, relPath)
		if d.IsDir() {
			return os.MkdirAll(dstPath, 0o755)
		}
		data, err := fs.ReadFile(srcFS, path)
		if err != nil {
			return err
		}
		return os.WriteFile(dstPath, data, 0o644)
	})
}

// Helper: Copy a directory
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}
		return copyFile(path, dstPath)
	})
}

// Helper: Run Wasm builds (forces GOOS=js and GOARCH=wasm into environment)
func runWasmBuild(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Error running command '%s %s': %v\n", name, strings.Join(args, " "), err)
		os.Exit(1)
	}
}
