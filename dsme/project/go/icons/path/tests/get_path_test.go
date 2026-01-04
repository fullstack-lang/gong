package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/sociointel/weberv2specs/go/icons/path"
)

func TestPathExtration(t *testing.T) {
	// Define the test data
	svgContent := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<svg
	   height="24"
	   viewBox="0 -960 960 960"
	   width="24"
	   version="1.1"
	   id="svg1"
	   sodipodi:docname="arrow_up_sided_to_right.svg"
	   inkscape:version="1.3.2 (091e20e, 2023-11-25)"
	   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
	   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
	   xmlns="http://www.w3.org/2000/svg"
	   xmlns:svg="http://www.w3.org/2000/svg">
	  <defs
		 id="defs1" />
	  <sodipodi:namedview
		 id="namedview1"
		 pagecolor="#ffffff"
		 bordercolor="#666666"
		 borderopacity="1.0"
		 inkscape:showpageshadow="2"
		 inkscape:pageopacity="0.0"
		 inkscape:pagecheckerboard="0"
		 inkscape:deskcolor="#d1d1d1"
		 inkscape:zoom="39.333333"
		 inkscape:cx="11.961864"
		 inkscape:cy="13.817797"
		 inkscape:window-width="1792"
		 inkscape:window-height="1120"
		 inkscape:window-x="0"
		 inkscape:window-y="0"
		 inkscape:window-maximized="0"
		 inkscape:current-layer="svg1" />
	  <path
		 d="m 680,-120 v -567 l -64,63 -56,-56 160,-160 160,160 -56,56 -64,-63 v 567 z"
		 id="path1"
		 sodipodi:nodetypes="cccccccccc" />
	</svg>
`

	// Create a temporary file to store the gitignore content
	tmpFile, err := os.CreateTemp("", "test.svg")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the gitignore content to the temporary file
	if _, err := tmpFile.WriteString(svgContent); err != nil {
		t.Fatalf("Failed to write svg content: %v", err)
	}

	// Replace 'yourfile.svg' with your SVG file path
	paths, err := path.ExtractSVGPaths(svgContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, path := range paths {
		fmt.Println(path)
	}
}
