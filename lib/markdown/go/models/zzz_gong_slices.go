// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Content
	// insertion point per field

	// Compute reverse map for named struct JpgImage
	// insertion point per field

	// Compute reverse map for named struct PngImage
	// insertion point per field

	// Compute reverse map for named struct SvgImage
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.JpgImages {
		res = append(res, instance)
	}

	for instance := range stage.PngImages {
		res = append(res, instance)
	}

	for instance := range stage.SvgImages {
		res = append(res, instance)
	}


	return
}
