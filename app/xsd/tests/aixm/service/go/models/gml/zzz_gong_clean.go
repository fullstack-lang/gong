// generated code - do not edit
package gml

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
// Clean garbage collect unstaged instances that are referenced by AbstractFeatureType
func (abstractfeaturetype *AbstractFeatureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractfeaturetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractfeaturetype.Identifier) || modified
	modified = stage.CleanPointer(&abstractfeaturetype.BoundedBy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGMLType
func (abstractgmltype *AbstractGMLType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgmltype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgmltype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeometricPrimitiveType
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeometricprimitivetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeometricprimitivetype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeometryType
func (abstractgeometrytype *AbstractGeometryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeometrytype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeometrytype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeGeometricPrimitiveType
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimegeometricprimitivetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimegeometricprimitivetype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeObjectType
func (abstracttimeobjecttype *AbstractTimeObjectType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeobjecttype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeobjecttype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimePrimitiveType
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeprimitivetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeprimitivetype.Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeSliceType
func (abstracttimeslicetype *AbstractTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeslicetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeslicetype.Identifier) || modified
	modified = stage.CleanPointer(&abstracttimeslicetype.ValidTime) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BoundingShapeType
func (boundingshapetype *BoundingShapeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&boundingshapetype.Envelope) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DynamicFeatureType
func (dynamicfeaturetype *DynamicFeatureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dynamicfeaturetype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dynamicfeaturetype.Identifier) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.BoundedBy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EnvelopeType
func (envelopetype *EnvelopeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&envelopetype.LowerCorner) || modified
	modified = stage.CleanPointer(&envelopetype.UpperCorner) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PointType
func (pointtype *PointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointtype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&pointtype.Identifier) || modified
	modified = stage.CleanPointer(&pointtype.Pos) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeInstantType
func (timeinstanttype *TimeInstantType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeinstanttype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeinstanttype.Identifier) || modified
	modified = stage.CleanPointer(&timeinstanttype.TimePosition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimePeriodType
func (timeperiodtype *TimePeriodType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeperiodtype.NameXSD) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeperiodtype.Identifier) || modified
	modified = stage.CleanPointer(&timeperiodtype.BeginPosition) || modified
	modified = stage.CleanPointer(&timeperiodtype.EndPosition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimePrimitivePropertyType
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimeInstant) || modified
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimePeriod) || modified
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
