package path

import (
	"encoding/xml"
)

// SVG structure to hold the XML data
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Paths   []Path   `xml:"path"`
}

// Path structure for the path elements in the SVG
type Path struct {
	D string `xml:"d,attr"`
}

// Function to parse SVG file and return path definitions
func ExtractSVGPaths(svgContent string) ([]string, error) {
	var svg SVG
	err := xml.Unmarshal([]byte(svgContent), &svg)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, path := range svg.Paths {
		paths = append(paths, path.D)
	}

	return paths, nil
}
