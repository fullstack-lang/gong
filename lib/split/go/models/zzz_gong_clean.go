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
// Clean garbage collect unstaged instances that are referenced by AsSplit
func (assplit *AsSplit) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&assplit.AsSplitAreas) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AsSplitArea
func (assplitarea *AsSplitArea) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&assplitarea.AsSplit) || modified
	modified = stage.CleanPointer(&assplitarea.Button) || modified
	modified = stage.CleanPointer(&assplitarea.Cursor) || modified
	modified = stage.CleanPointer(&assplitarea.Form) || modified
	modified = stage.CleanPointer(&assplitarea.Load) || modified
	modified = stage.CleanPointer(&assplitarea.Markdown) || modified
	modified = stage.CleanPointer(&assplitarea.Slider) || modified
	modified = stage.CleanPointer(&assplitarea.Split) || modified
	modified = stage.CleanPointer(&assplitarea.Svg) || modified
	modified = stage.CleanPointer(&assplitarea.Table) || modified
	modified = stage.CleanPointer(&assplitarea.Tone) || modified
	modified = stage.CleanPointer(&assplitarea.Tree) || modified
	modified = stage.CleanPointer(&assplitarea.Threejs) || modified
	modified = stage.CleanPointer(&assplitarea.Xlsx) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by View
func (view *View) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&view.RootAsSplitAreas) || modified
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
