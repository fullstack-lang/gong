// generated code - do not edit
package content

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
// Clean garbage collect unstaged instances that are referenced by MD_Band_PropertyType
func (md_band_propertytype *MD_Band_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_band_propertytype.MD_Band) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ContentInformation_PropertyType
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_contentinformation_propertytype.MD_FeatureCatalogueDescription) || modified
	modified = stage.CleanPointer(&md_contentinformation_propertytype.MD_CoverageDescription) || modified
	modified = stage.CleanPointer(&md_contentinformation_propertytype.MD_ImageDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_CoverageDescription_PropertyType
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_coveragedescription_propertytype.MD_CoverageDescription) || modified
	modified = stage.CleanPointer(&md_coveragedescription_propertytype.MD_ImageDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_CoverageDescription_Type
func (md_coveragedescription_type *MD_CoverageDescription_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_coveragedescription_type.Dimension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_coveragedescription_type.ContentType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_FeatureCatalogueDescription_PropertyType
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ImageDescription_PropertyType
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_imagedescription_propertytype.MD_ImageDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ImageDescription_Type
func (md_imagedescription_type *MD_ImageDescription_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_imagedescription_type.Dimension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_imagedescription_type.ContentType) || modified
	modified = stage.CleanPointer(&md_imagedescription_type.ImagingCondition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_RangeDimension_PropertyType
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_rangedimension_propertytype.MD_RangeDimension) || modified
	modified = stage.CleanPointer(&md_rangedimension_propertytype.MD_Band) || modified
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
