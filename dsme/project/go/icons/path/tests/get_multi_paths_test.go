package tests

import (
	"fmt"
	"testing"

	"github.com/sociointel/weberv2specs/go/icons"
	"github.com/sociointel/weberv2specs/go/icons/path"
)

func TestMultiPathExtraction(t *testing.T) {

	// Replace 'yourfile.svg' with your SVG file path
	paths, err := path.ExtractSVGPaths(icons.Percent_000_Icon.SVG)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, path := range paths {
		fmt.Println(path)
	}
}
