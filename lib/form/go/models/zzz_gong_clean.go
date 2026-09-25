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
// Clean garbage collect unstaged instances that are referenced by FormDiv
func (formdiv *FormDiv) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&formdiv.FormFields) || modified
	modified = stage.CleanSlice(&formdiv.CheckBoxs) || modified
	// insertion point per field
	modified = stage.CleanPointer(&formdiv.FormEditAssocButton) || modified
	modified = stage.CleanPointer(&formdiv.FormSortAssocButton) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormField
func (formfield *FormField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&formfield.FormFieldString) || modified
	modified = stage.CleanPointer(&formfield.FormFieldFloat64) || modified
	modified = stage.CleanPointer(&formfield.FormFieldInt) || modified
	modified = stage.CleanPointer(&formfield.FormFieldDate) || modified
	modified = stage.CleanPointer(&formfield.FormFieldTime) || modified
	modified = stage.CleanPointer(&formfield.FormFieldDateTime) || modified
	modified = stage.CleanPointer(&formfield.FormFieldSelect) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldSelect
func (formfieldselect *FormFieldSelect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&formfieldselect.Options) || modified
	// insertion point per field
	modified = stage.CleanPointer(&formfieldselect.Value) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormGroup
func (formgroup *FormGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&formgroup.FormDivs) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormSortAssocButton
func (formsortassocbutton *FormSortAssocButton) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&formsortassocbutton.FormEditAssocButton) || modified
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
