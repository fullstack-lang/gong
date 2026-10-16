// generated code - do not edit
package attributes

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
// Clean garbage collect unstaged instances that are referenced by Attributes
func (attributes *Attributes) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&attributes.Key) || modified
	modified = stage.CleanSlice(&attributes.Time) || modified
	modified = stage.CleanSlice(&attributes.Clef) || modified
	modified = stage.CleanSlice(&attributes.StaffDetails) || modified
	modified = stage.CleanSlice(&attributes.Transpose) || modified
	modified = stage.CleanSlice(&attributes.ForPart) || modified
	modified = stage.CleanSlice(&attributes.Directive) || modified
	modified = stage.CleanSlice(&attributes.MeasureStyle) || modified
	// insertion point per field
	modified = stage.CleanPointer(&attributes.PartSymbol) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ForPart
func (forpart *ForPart) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&forpart.PartClef) || modified
	modified = stage.CleanPointer(&forpart.PartTranspose) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Key
func (key *Key) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&key.KeyAccidental) || modified
	modified = stage.CleanSlice(&key.KeyOctave) || modified
	// insertion point per field
	modified = stage.CleanPointer(&key.Cancel) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MeasureStyle
func (measurestyle *MeasureStyle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&measurestyle.MultipleRest) || modified
	modified = stage.CleanPointer(&measurestyle.MeasureRepeat) || modified
	modified = stage.CleanPointer(&measurestyle.BeatRepeat) || modified
	modified = stage.CleanPointer(&measurestyle.Slash) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PartTranspose
func (parttranspose *PartTranspose) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&parttranspose.Double) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StaffDetails
func (staffdetails *StaffDetails) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&staffdetails.LineDetail) || modified
	modified = stage.CleanSlice(&staffdetails.StaffTuning) || modified
	// insertion point per field
	modified = stage.CleanPointer(&staffdetails.StaffSize) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Time
func (time *Time) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&time.Interchangeable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Transpose
func (transpose *Transpose) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&transpose.Double) || modified
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
