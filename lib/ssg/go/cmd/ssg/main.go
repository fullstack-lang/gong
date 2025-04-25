package main

import (
	"strconv"

	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time" // Import the time package

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta" // Use alias 'meta' to avoid conflict
	"github.com/yuin/goldmark/extension" // <-- Import the standard extension package
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html" // Import the html renderer package

	// YAML parsing is handled by goldmark-meta

	// insertion point for models import
	ssg_models "github.com/fullstack-lang/gong/lib/ssg/go/models"
	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/stack"
	ssg_static "github.com/fullstack-lang/gong/lib/ssg/go/static"
)

// --- Structs (Page, SiteInfo) - Unchanged ---
// Page holds information about a single content file
type Page struct {
	SourcePath    string                 // Original path in content/
	OutputPath    string                 // Target path in the specified output directory
	URL           string                 // Absolute URL path (e.g., "/chapter1/page1/") used for internal logic
	RelativeURL   string                 // Path relative to output root (e.g., "chapter1/page1/index.html") used by templates
	RelPathToRoot string                 // Path from page to root (e.g., "../../" or "/") used by templates
	Title         string                 // Page title from front matter
	Weight        int                    // Weight for ordering from front matter
	ContentHTML   template.HTML          // Processed HTML content from Markdown
	IsHome        bool                   // Is this the root _index.md?
	IsSection     bool                   // Is this an _index.md file (a section index)?
	FrontMatter   map[string]interface{} // Store all front matter
	Pages         []*Page                // Child pages (for sections), sorted by weight
	Section       *Page                  // Parent section page (nil for top-level sections/home)
	Site          *SiteInfo              // Link back to global site info
}

// SiteInfo holds global information about the site
type SiteInfo struct {
	Pages     map[string]*Page   // Map Absolute URLs to Pages for easy lookup
	Sections  []*Page            // Top-level sections (chapters) sorted by weight
	Templates *template.Template // Parsed HTML templates
}

// Initialize Goldmark with extensions
var md = goldmark.New(
	goldmark.WithExtensions(
		meta.New(meta.WithStoresInDocument()), // For front matter
		extension.Table,                       // Use the built-in Table extension
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(), // Allows raw HTML in Markdown (needed for things like <img> tags)
	),
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	contentDir    = "content"
	staticDir     = "static"
	layoutDir     = "layouts"
	serverPort    = "8080"
	serveFlag     = flag.Bool("serve", false, "Serve the site on localhost after building (or without building if -no-build is used)")
	targetFlag    = flag.String("target", "file", "Build target: 'web' for server or 'file' for local filesystem (ignored if -no-build is used)")
	outputDirFlag = flag.String("output", "public", "Output directory for the generated site")
	noBuildFlag   = flag.Bool("no-build", false, "Skip the build process (requires -serve and existing output directory)")
)

func main() {

	log.SetPrefix("ssg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := ssg_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := ssg_stack.NewStack(r, "ssg", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	ssg_models.NewStager(r, stack.Stage)

	buildTarget := *targetFlag
	outputDir := *outputDirFlag // Use the directory specified by the flag

	// --- Handle -no-build -serve combination first ---
	if *noBuildFlag {
		if !*serveFlag {
			log.Fatalf("Error: -no-build flag requires the -serve flag.")
		}

		// Check if output directory exists
		info, err := os.Stat(outputDir)
		if os.IsNotExist(err) {
			log.Fatalf("Error: Output directory '%s' does not exist. Cannot serve without building.", outputDir)
		} else if err != nil {
			log.Fatalf("Error checking output directory '%s': %v", outputDir, err)
		} else if !info.IsDir() {
			log.Fatalf("Error: Output path '%s' exists but is not a directory. Cannot serve.", outputDir)
		}

		log.Printf("Skipping build process due to -no-build flag.")
		serveSite(outputDir) // Directly serve the existing directory
		return               // Exit after starting server
	}

	// --- Proceed with build if -no-build is not set ---

	if buildTarget != "web" && buildTarget != "file" {
		log.Fatalf("Invalid target '%s'. Must be 'web' or 'file'.", buildTarget)
	}
	// Warning if -serve is used with -target file (build still happens)
	if *serveFlag && buildTarget != "web" {
		log.Printf("Warning: -serve flag provided but target is 'file'. Site will be built but not served.")
		*serveFlag = false // Disable serve if target is not web
	}

	log.Printf("Starting book generator build (Target: %s, Output: %s)...", buildTarget, outputDir)

	// --- Build Steps ---
	if err := cleanOutputDir(outputDir); err != nil {
		log.Fatalf("Error cleaning output directory '%s': %v", outputDir, err)
	}
	log.Printf("Cleaned output directory '%s'.\n", outputDir)

	templates, err := loadTemplates()
	if err != nil {
		log.Fatalf("Error loading templates from '%s': %v", layoutDir, err)
	}
	log.Println("Loaded HTML templates.")

	site := &SiteInfo{
		Pages:     make(map[string]*Page),
		Templates: templates,
	}
	// Pass build target and output dir to parseContent
	err = parseContent(site, buildTarget, outputDir)
	if err != nil {
		log.Fatalf("Error parsing content from '%s': %v", contentDir, err)
	}
	log.Printf("Parsed %d content files.\n", len(site.Pages))

	buildSiteStructure(site) // <-- Call the updated function
	log.Println("Built site structure (sections and pages).")

	// Pass output dir and build target to renderPages
	err = renderPages(site, outputDir, buildTarget)
	if err != nil {
		log.Fatalf("Error rendering HTML pages: %v", err)
	}
	log.Println("Rendered HTML pages.")

	// Pass output dir to copyStaticFiles
	if err := copyStaticFiles(outputDir); err != nil {
		log.Fatalf("Error copying static files from '%s' to '%s': %v", staticDir, outputDir, err)
	}
	log.Println("Copied static files.")

	log.Println("Build finished successfully!")

	// // --- Serve Site (if flag is set and target is web, and -no-build was NOT set) ---
	// if *serveFlag {
	// 	serveSite(outputDir) // Serve from the specified output directory
	// }

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err = r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}

// --- cleanOutputDir - Unchanged ---
func cleanOutputDir(dir string) error {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		log.Printf("Removing existing output directory: %s\n", dir)
		return os.RemoveAll(dir)
	}
	return nil
}

// --- loadTemplates - Unchanged ---
func loadTemplates() (*template.Template, error) {
	tmpl := template.New("base").Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	})

	globPattern := filepath.Join(layoutDir, "**/*.html")
	log.Printf("Parsing templates with glob pattern: %s\n", globPattern)

	tmpl, err := tmpl.ParseGlob(globPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates with glob '%s': %w", globPattern, err)
	}

	return tmpl, nil
}

// --- parseContent - Unchanged ---
// Uses the globally configured `md` variable which now includes the Table extension.
func parseContent(site *SiteInfo, buildTarget string, outputDir string) error {
	log.Printf("Parsing content from directory: %s\n", contentDir)
	return filepath.WalkDir(contentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if d.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v\n", path, err)
			return nil // Continue with other files
		}

		var buf bytes.Buffer
		context := parser.NewContext()
		// Use the globally configured md parser (now with Table support)
		if err := md.Convert(contentBytes, &buf, parser.WithContext(context)); err != nil {
			log.Printf("Error converting markdown file %s: %v\n", path, err)
			return nil // Continue with other files
		}

		metaData := meta.Get(context)
		pageMeta := make(map[string]interface{})
		if metaData != nil {
			pageMeta = metaData
		} else {
			log.Printf("Warning: No front matter found in %s\n", path)
		}

		// --- Determine Paths and URLs ---
		relSourcePath, _ := filepath.Rel(contentDir, path)
		outputPath := ""
		url := ""           // Absolute URL path for internal logic (remains consistent)
		relativeURL := ""   // Path relative to output root (consistent for both targets)
		relPathToRoot := "" // Path from page to root (depends on target)

		isSection := strings.HasSuffix(filepath.Base(path), "_index.md")
		isHome := (relSourcePath == "_index.md")

		// Calculate OutputPath, internal URL, and root-relative URL (used by templates)
		if isHome {
			outputPath = filepath.Join(outputDir, "index.html")
			url = "/"
			relativeURL = "index.html" // Path relative to output root
		} else {
			dir := filepath.Dir(relSourcePath)
			base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

			if isSection {
				// Section index page (_index.md)
				outputPath = filepath.Join(outputDir, dir, "index.html")
				url = "/" + filepath.ToSlash(dir) + "/"                          // URL ends with /
				relativeURL = filepath.ToSlash(filepath.Join(dir, "index.html")) // Path relative to output root
			} else {
				// Regular page (e.g., page.md)
				outputPath = filepath.Join(outputDir, dir, base, "index.html")         // Output into its own folder
				url = "/" + filepath.ToSlash(filepath.Join(dir, base)) + "/"           // URL ends with /
				relativeURL = filepath.ToSlash(filepath.Join(dir, base, "index.html")) // Path relative to output root
			}
		}

		// Calculate RelPathToRoot based on target
		if buildTarget == "web" {
			// Web target ALWAYS uses root-relative paths for assets by setting RelPathToRoot to "/"
			relPathToRoot = "/"
		} else { // target == "file"
			// File target calculates relative path back to root based on depth
			if isHome {
				relPathToRoot = ""
			} else {
				// Calculate depth based on the *output* directory structure
				outputRel, _ := filepath.Rel(outputDir, outputPath)
				depth := len(strings.Split(filepath.ToSlash(filepath.Dir(outputRel)), "/"))
				relPathToRoot = strings.Repeat("../", depth)
			}
		}

		// --- Create Page Struct ---
		page := &Page{
			SourcePath:    path,
			OutputPath:    outputPath,
			URL:           url,           // Internal absolute URL
			RelativeURL:   relativeURL,   // Path relative to output root (used by template)
			RelPathToRoot: relPathToRoot, // Path from page to root (depends on target, used by template)
			ContentHTML:   template.HTML(buf.String()),
			IsHome:        isHome,
			IsSection:     isSection,
			FrontMatter:   pageMeta,
			Site:          site,
			Title:         getString(pageMeta, "title", "Untitled"),
			Weight:        getInt(pageMeta, "weight", 0),
		}

		if _, exists := site.Pages[page.URL]; exists {
			log.Printf("Warning: Duplicate URL detected '%s' from file %s. Overwriting.", page.URL, path)
		}
		site.Pages[page.URL] = page

		return nil
	})
}

// --- buildSiteStructure (MODIFIED) ---
// This function now walks up the directory tree to find the nearest parent section.
func buildSiteStructure(site *SiteInfo) {
	// First pass: Link children to their nearest ancestor section
	for url, page := range site.Pages {
		if page.IsHome {
			continue // Skip the home page itself
		}

		// Start searching for parent from the page's directory URL
		currentPath := url
		foundParent := false

		// Loop upwards until a parent section is found or root is reached
		for currentPath != "/" {
			// Calculate the parent directory's URL path
			trimmedPath := strings.TrimSuffix(currentPath, "/")
			lastSlash := strings.LastIndex(trimmedPath, "/")
			if lastSlash < 0 { // Should not happen if currentPath != "/"
				break
			}
			parentURL := "/" // Default to root if parent is top level
			if lastSlash > 0 {
				parentURL = trimmedPath[:lastSlash] + "/"
			}

			// Check if a page exists at the parent URL and if it's a section or home
			if parentSection, exists := site.Pages[parentURL]; exists {
				if parentSection.IsHome || parentSection.IsSection {
					// Found the nearest ancestor section
					page.Section = parentSection
					parentSection.Pages = append(parentSection.Pages, page)
					foundParent = true
					// log.Printf("Linked page '%s' to section '%s'", page.URL, parentSection.URL) // Debug logging
					break // Stop searching upwards
				}
			}

			// Move up one level
			currentPath = parentURL
		}

		if !foundParent && !page.IsSection { // Only warn for non-section pages that couldn't be linked
			log.Printf("Warning: Could not find ancestor section for page '%s'. It might not appear in navigation correctly.", page.URL)
		}
	}

	// Second pass: Populate top-level sections and sort everything
	homePage := site.Pages["/"] // Assume home page exists
	for _, page := range site.Pages {
		// Identify top-level sections (direct children of the home page)
		if page.IsSection && !page.IsHome && page.Section == homePage {
			site.Sections = append(site.Sections, page)
		}

		// Sort child pages within each section (and the home page)
		if page.IsHome || page.IsSection {
			sort.Slice(page.Pages, func(i, j int) bool {
				// Primary sort by Weight
				if page.Pages[i].Weight != page.Pages[j].Weight {
					return page.Pages[i].Weight < page.Pages[j].Weight
				}
				// Secondary sort by Title if weights are equal
				return page.Pages[i].Title < page.Pages[j].Title
			})
		}
	}

	// Sort the top-level sections themselves
	sort.Slice(site.Sections, func(i, j int) bool {
		// Primary sort by Weight
		if site.Sections[i].Weight != site.Sections[j].Weight {
			return site.Sections[i].Weight < site.Sections[j].Weight
		}
		// Secondary sort by Title if weights are equal
		return site.Sections[i].Title < site.Sections[j].Title
	})

	log.Printf("Identified %d top-level sections.\n", len(site.Sections))
}

// --- renderPages - Unchanged ---
// It now accepts the build target and passes it to the template context.
func renderPages(site *SiteInfo, outputDir string, buildTarget string) error {
	log.Println("Rendering pages...")
	var renderErrors []error

	for _, page := range site.Pages {
		// Note: page.OutputPath already includes the base outputDir from parseContent
		outputPageDir := filepath.Dir(page.OutputPath)
		if err := os.MkdirAll(outputPageDir, 0755); err != nil {
			log.Printf("Error creating directory %s for page %s: %v\n", outputPageDir, page.SourcePath, err)
			renderErrors = append(renderErrors, fmt.Errorf("mkdir failed for %s: %w", page.OutputPath, err))
			continue
		}

		outFile, err := os.Create(page.OutputPath)
		if err != nil {
			log.Printf("Error creating output file %s for page %s: %v\n", page.OutputPath, page.SourcePath, err)
			renderErrors = append(renderErrors, fmt.Errorf("create failed for %s: %w", page.OutputPath, err))
			continue
		}

		// Determine which layout to use: list.html for sections/home, single.html otherwise
		layoutName := "single.html"
		if page.IsHome || page.IsSection {
			layoutName = "list.html"
		}

		// Define the data to pass to the template
		templateData := map[string]interface{}{
			"Page":        page,
			"Site":        site,
			"BuildTarget": buildTarget, // Pass buildTarget to template data
		}

		// Execute the base template, which will in turn call the correct block (list or single)
		err = site.Templates.ExecuteTemplate(outFile, "baseof.html", templateData)
		closeErr := outFile.Close() // Close file after execution attempt

		if err != nil {
			log.Printf("Error executing template (layout: %s) for %s (URL: %s): %v\n", layoutName, page.SourcePath, page.URL, err)
			renderErrors = append(renderErrors, fmt.Errorf("template failed for %s: %w", page.SourcePath, err))
		} else if closeErr != nil {
			log.Printf("Error closing output file %s: %v\n", page.OutputPath, closeErr)
			// Decide if this is fatal or just a warning. For now, treat as an error.
			renderErrors = append(renderErrors, fmt.Errorf("close failed for %s: %w", page.OutputPath, closeErr))
		}
	}

	if len(renderErrors) > 0 {
		log.Printf("Encountered %d errors during page rendering.\n", len(renderErrors))
		// Return the first error encountered for simplicity
		return fmt.Errorf("encountered %d errors during rendering, first error: %w", len(renderErrors), renderErrors[0])
	}
	return nil
}

// --- copyStaticFiles - Unchanged ---
func copyStaticFiles(outputDir string) error {
	log.Printf("Copying static files from '%s' to '%s'\n", staticDir, outputDir)
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Printf("Static directory '%s' does not exist, skipping copy.\n", staticDir)
		return nil
	}

	return filepath.WalkDir(staticDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(staticDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}
		if relPath == "." {
			return nil // Skip the root directory itself
		}

		// Target path is relative to the specified output directory
		targetPath := filepath.Join(outputDir, relPath)

		if d.IsDir() {
			// Create directory structure in the output directory
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				log.Printf("Error creating static directory %s: %v\n", targetPath, err)
				return err // Stop copying if directory creation fails
			}
			return nil // Directory processed, continue walking
		}

		// It's a file, copy it
		sourceFile, err := os.Open(path)
		if err != nil {
			log.Printf("Error opening static source file %s: %v\n", path, err)
			return err // Stop if source file cannot be opened
		}
		defer sourceFile.Close() // Ensure source file is closed

		// Ensure the target directory exists (needed if copying single files into new dirs)
		targetFileDir := filepath.Dir(targetPath)
		if err := os.MkdirAll(targetFileDir, 0755); err != nil {
			log.Printf("Error creating directory %s for static file %s: %v\n", targetFileDir, targetPath, err)
			return err
		}

		targetFile, err := os.Create(targetPath)
		if err != nil {
			log.Printf("Error creating static target file %s: %v\n", targetPath, err)
			return err // Stop if target file cannot be created
		}
		defer targetFile.Close() // Ensure target file is closed

		_, err = io.Copy(targetFile, sourceFile)
		if err != nil {
			log.Printf("Error copying static file from %s to %s: %v\n", path, targetPath, err)
			return err // Stop if copying fails
		}
		// log.Printf("Copied static file: %s -> %s\n", path, targetPath) // Optional: log successful copies
		return nil
	})
}

// --- Logging Middleware Helpers - Unchanged ---
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r) // Serve the request first
		duration := time.Since(start)
		// Log details after the request is served
		log.Printf(
			"%s %d \"%s %s %s\" %s \"%s\" \"%s\"\n",
			r.RemoteAddr, lrw.statusCode, r.Method, r.URL.RequestURI(), r.Proto,
			duration, r.Referer(), r.UserAgent(),
		)
	})
}

// --- serveSite - Unchanged ---
// serveSite starts a web server with explicit 404 handling for non-existent files.
func serveSite(outputDir string) {
	addr := ":" + serverPort
	log.Printf("Serving site from '%s' on http://localhost%s\n", outputDir, addr)

	// Base file server
	fs := http.FileServer(http.Dir(outputDir))

	// Wrap the base file server with our custom handler and logging
	handler := loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Construct the full path on the filesystem
		// Use filepath.Join for OS-independent path construction
		// Use filepath.Clean to prevent directory traversal attacks (e.g., /../..)
		requestedPath := filepath.Join(outputDir, filepath.Clean(r.URL.Path))

		// If the path ends with '/', check for index.html in that directory
		if strings.HasSuffix(r.URL.Path, "/") {
			requestedPath = filepath.Join(requestedPath, "index.html")
		} else {
			// Check if the requested path corresponds to a directory in the output
			// If so, check if an index.html exists within it.
			info, err := os.Stat(requestedPath)
			if err == nil && info.IsDir() {
				indexPath := filepath.Join(requestedPath, "index.html")
				if _, err := os.Stat(indexPath); err == nil {
					// If index.html exists, serve it by redirecting
					// Ensure the redirect URL has a trailing slash
					redirectURL := r.URL.Path
					if !strings.HasSuffix(redirectURL, "/") {
						redirectURL += "/"
					}
					// log.Printf("Redirecting directory request: %s -> %s", r.URL.Path, redirectURL) // Debug log
					http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
					return
				}
			}
		}

		// Check if the file exists (after potential index.html adjustment)
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			// File does not exist, return 404
			log.Printf("File not found: %s (Request URI: %s)", requestedPath, r.URL.RequestURI())
			// Check if it's a directory without index.html - could serve a directory listing or 404
			// For now, just 404
			http.NotFound(w, r)
			return // Stop processing here
		}

		// File exists, let the standard file server handle it
		fs.ServeHTTP(w, r)
	}))

	// Register the wrapped handler
	http.Handle("/", handler)

	// Start the server
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server on address %s: %v\n", addr, err)
	}
}

// --- Helper functions (getString, getInt) - Unchanged ---
func getString(m map[string]interface{}, key string, defaultValue string) string {
	if val, ok := m[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return defaultValue
}

func getInt(m map[string]interface{}, key string, defaultValue int) int {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case int:
			return v
		case float64: // Handle potential number types from YAML/JSON
			return int(v)
		case int64:
			return int(v)
		case int32:
			return int(v)
		}
		log.Printf("Warning: Front matter key '%s' has unexpected type %T, expected int or number.\n", key, val)
	}
	return defaultValue
}
