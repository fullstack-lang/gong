// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AIXMBasicMessage
func (aixmbasicmessage *AIXMBasicMessage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aixmbasicmessage.HasMember) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AIXMBasicMessageType
func (aixmbasicmessagetype *AIXMBasicMessageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aixmbasicmessagetype.HasMember) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	if stage.BasictypesStage != nil {
		modified = stage.BasictypesStage.Clean() || modified
	}
	if stage.XmlxsdStage != nil {
		modified = stage.XmlxsdStage.Clean() || modified
	}
	if stage.XlinkStage != nil {
		modified = stage.XlinkStage.Clean() || modified
	}
	if stage.GcobaseStage != nil {
		modified = stage.GcobaseStage.Clean() || modified
	}
	if stage.MeasuresStage != nil {
		modified = stage.MeasuresStage.Clean() || modified
	}
	if stage.GmlStage != nil {
		modified = stage.GmlStage.Clean() || modified
	}
	if stage.ApplicationschemaStage != nil {
		modified = stage.ApplicationschemaStage.Clean() || modified
	}
	if stage.ContentStage != nil {
		modified = stage.ContentStage.Clean() || modified
	}
	if stage.MetadataextensionStage != nil {
		modified = stage.MetadataextensionStage.Clean() || modified
	}
	if stage.PortrayalcatalogueStage != nil {
		modified = stage.PortrayalcatalogueStage.Clean() || modified
	}
	if stage.SpatialrepresentationStage != nil {
		modified = stage.SpatialrepresentationStage.Clean() || modified
	}
	if stage.GmdStage != nil {
		modified = stage.GmdStage.Clean() || modified
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		modified = stage.AixmabstractgmlobjecttypesStage.Clean() || modified
	}
	if stage.AixmdatatypesStage != nil {
		modified = stage.AixmdatatypesStage.Clean() || modified
	}
	if stage.AixmfeaturesStage != nil {
		modified = stage.AixmfeaturesStage.Clean() || modified
	}
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
