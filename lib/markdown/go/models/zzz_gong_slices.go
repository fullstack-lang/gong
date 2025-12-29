// generated code - do not edit
package models

import "time"
var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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


func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenDeletedInstances int
	
	// insertion point per named struct
	var contents_newInstances []*Content
	var contents_deletedInstances []*Content

	// parse all staged instances and check if they have a reference
	for content := range stage.Contents {
		if _, ok := stage.Contents_reference[content]; !ok {
			contents_newInstances = append(contents_newInstances, content)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of Content "+content.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for content := range stage.Contents_reference {
		if _, ok := stage.Contents[content]; !ok {
			contents_deletedInstances = append(contents_deletedInstances, content)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of Content "+content.Name,
				)
			}
		}
	}

	lenNewInstances += len(contents_newInstances)
	lenDeletedInstances += len(contents_deletedInstances)
	var jpgimages_newInstances []*JpgImage
	var jpgimages_deletedInstances []*JpgImage

	// parse all staged instances and check if they have a reference
	for jpgimage := range stage.JpgImages {
		if _, ok := stage.JpgImages_reference[jpgimage]; !ok {
			jpgimages_newInstances = append(jpgimages_newInstances, jpgimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of JpgImage "+jpgimage.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for jpgimage := range stage.JpgImages_reference {
		if _, ok := stage.JpgImages[jpgimage]; !ok {
			jpgimages_deletedInstances = append(jpgimages_deletedInstances, jpgimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of JpgImage "+jpgimage.Name,
				)
			}
		}
	}

	lenNewInstances += len(jpgimages_newInstances)
	lenDeletedInstances += len(jpgimages_deletedInstances)
	var pngimages_newInstances []*PngImage
	var pngimages_deletedInstances []*PngImage

	// parse all staged instances and check if they have a reference
	for pngimage := range stage.PngImages {
		if _, ok := stage.PngImages_reference[pngimage]; !ok {
			pngimages_newInstances = append(pngimages_newInstances, pngimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of PngImage "+pngimage.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for pngimage := range stage.PngImages_reference {
		if _, ok := stage.PngImages[pngimage]; !ok {
			pngimages_deletedInstances = append(pngimages_deletedInstances, pngimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of PngImage "+pngimage.Name,
				)
			}
		}
	}

	lenNewInstances += len(pngimages_newInstances)
	lenDeletedInstances += len(pngimages_deletedInstances)
	var svgimages_newInstances []*SvgImage
	var svgimages_deletedInstances []*SvgImage

	// parse all staged instances and check if they have a reference
	for svgimage := range stage.SvgImages {
		if _, ok := stage.SvgImages_reference[svgimage]; !ok {
			svgimages_newInstances = append(svgimages_newInstances, svgimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of SvgImage "+svgimage.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svgimage := range stage.SvgImages_reference {
		if _, ok := stage.SvgImages[svgimage]; !ok {
			svgimages_deletedInstances = append(svgimages_deletedInstances, svgimage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of SvgImage "+svgimage.Name,
				)
			}
		}
	}

	lenNewInstances += len(svgimages_newInstances)
	lenDeletedInstances += len(svgimages_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Contents_reference = make(map[*Content]*Content)
	for instance := range stage.Contents {
		stage.Contents_reference[instance] = instance
	}

	stage.JpgImages_reference = make(map[*JpgImage]*JpgImage)
	for instance := range stage.JpgImages {
		stage.JpgImages_reference[instance] = instance
	}

	stage.PngImages_reference = make(map[*PngImage]*PngImage)
	for instance := range stage.PngImages {
		stage.PngImages_reference[instance] = instance
	}

	stage.SvgImages_reference = make(map[*SvgImage]*SvgImage)
	for instance := range stage.SvgImages {
		stage.SvgImages_reference[instance] = instance
	}

}
