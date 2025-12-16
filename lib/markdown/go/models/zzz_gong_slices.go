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

// insertion point per named struct
func (content *Content) GongCopy() GongstructIF {
	newInstance := *content
	return &newInstance
}

func (jpgimage *JpgImage) GongCopy() GongstructIF {
	newInstance := *jpgimage
	return &newInstance
}

func (pngimage *PngImage) GongCopy() GongstructIF {
	newInstance := *pngimage
	return &newInstance
}

func (svgimage *SvgImage) GongCopy() GongstructIF {
	newInstance := *svgimage
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
}
