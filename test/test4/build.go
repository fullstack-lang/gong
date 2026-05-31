package main

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// =========================================================================
	// 1. Bash Script Equivalent (Compilation)
	// =========================================================================
	fmt.Println("🚀 Step 1: Building project...")

	// 0. Cleanup previous builds and recreate .tmp_build
	os.RemoveAll(".tmp_build")
	os.RemoveAll("go/cmd/test4-wasm/test4-wasm")
	os.Remove("test4-portable-app.html")
	os.Remove("test4-portable-app.zip")
	os.MkdirAll(".tmp_build", 0755)

	// 1. Copy Frontend Assets from the lib/split directory
	splitNgDir := filepath.Join("..", "..", "lib", "split", "ng-github.com-fullstack-lang-gong-lib-split")
	distBrowserDir := filepath.Join(splitNgDir, "dist", "ng-github.com-fullstack-lang-gong-lib-split", "browser")
	if _, err := os.Stat(distBrowserDir); os.IsNotExist(err) {
		fmt.Printf("❌ Error: Angular dist directory not found at %s. Please build the Angular app first.\n", distBrowserDir)
		os.Exit(1)
	}
	copyDir(distBrowserDir, ".tmp_build")

	// Copy wasm_exec.js (Handles both older and newer Go versions)
	goRootBytes, _ := exec.Command("go", "env", "GOROOT").Output()
	goRoot := strings.TrimSpace(string(goRootBytes))
	wasmExecPath := filepath.Join(goRoot, "misc", "wasm", "wasm_exec.js")
	if _, err := os.Stat(wasmExecPath); os.IsNotExist(err) {
		wasmExecPath = filepath.Join(goRoot, "lib", "wasm", "wasm_exec.js")
	}
	copyFile(wasmExecPath, filepath.Join(".tmp_build", "wasm_exec.js"))

	// Build the Go Wasm binary:
	runWasmBuild("go", "build", "-a", "-o", ".tmp_build/main.wasm", "./go/cmd/test4-wasm")

	// =========================================================================
	// 2. bundle.js Equivalent (Packaging)
	// =========================================================================
	fmt.Println("\n📦 Step 2: Packaging Test4 Gong application into a single HTML file...")

	buildDir := ".tmp_build"
	outputFile := "test4-portable-app.html"

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
<script>%s</script>

<script>
  console.log("Initializing Test4 WASM Backend from Base64...");
  const base64String = "%s";
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
`, wasmExecJs, wasmBase64)

	// Inject the bootloader right before the closing </body> tag
	html = strings.Replace(html, "</body>", bootloader+"\n</body>", 1)

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
	err = os.WriteFile(outputFile, []byte(html), 0644)
	if err != nil {
		fmt.Printf("❌ Error writing output file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n✅ Success! You can now double-click: %s\n", outputFile)

	// Zip the HTML file
	zipFile := "test4-portable-app.zip"
	if err := createZip(zipFile, outputFile); err != nil {
		fmt.Printf("❌ Error zipping output file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("🗜️ Zipped output to: %s\n", zipFile)
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
