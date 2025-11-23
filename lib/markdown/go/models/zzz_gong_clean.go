// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Content
	for content := range stage.Contents {
		_ = content
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct JpgImage
	for jpgimage := range stage.JpgImages {
		_ = jpgimage
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct PngImage
	for pngimage := range stage.PngImages {
		_ = pngimage
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct SvgImage
	for svgimage := range stage.SvgImages {
		_ = svgimage
		// insertion point per field
		// insertion point per field
	}

}
