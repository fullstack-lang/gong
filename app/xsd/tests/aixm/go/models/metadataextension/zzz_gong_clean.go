// generated code - do not edit
package metadataextension

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
// Clean garbage collect unstaged instances that are referenced by MD_ExtendedElementInformation_PropertyType
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_extendedelementinformation_propertytype.MD_ExtendedElementInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ExtendedElementInformation_Type
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_extendedelementinformation_type.Obligation) || modified
	modified = stage.CleanPointer(&md_extendedelementinformation_type.DataType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_MetadataExtensionInformation_PropertyType
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_MetadataExtensionInformation_Type
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_metadataextensioninformation_type.ExtendedElementInformation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
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
