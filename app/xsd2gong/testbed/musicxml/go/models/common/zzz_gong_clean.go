// generated code - do not edit
package common

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
// Clean garbage collect unstaged instances that are referenced by Appearance
func (appearance *Appearance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&appearance.LineWidth) || modified
	modified = stage.CleanSlice(&appearance.NoteSize) || modified
	modified = stage.CleanSlice(&appearance.Distance) || modified
	modified = stage.CleanSlice(&appearance.Glyph) || modified
	modified = stage.CleanSlice(&appearance.OtherAppearance) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Dynamics
func (dynamics *Dynamics) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dynamics.P) || modified
	modified = stage.CleanSlice(&dynamics.Pp) || modified
	modified = stage.CleanSlice(&dynamics.Ppp) || modified
	modified = stage.CleanSlice(&dynamics.Pppp) || modified
	modified = stage.CleanSlice(&dynamics.Ppppp) || modified
	modified = stage.CleanSlice(&dynamics.Pppppp) || modified
	modified = stage.CleanSlice(&dynamics.F) || modified
	modified = stage.CleanSlice(&dynamics.Ff) || modified
	modified = stage.CleanSlice(&dynamics.Fff) || modified
	modified = stage.CleanSlice(&dynamics.Ffff) || modified
	modified = stage.CleanSlice(&dynamics.Fffff) || modified
	modified = stage.CleanSlice(&dynamics.Ffffff) || modified
	modified = stage.CleanSlice(&dynamics.Mp) || modified
	modified = stage.CleanSlice(&dynamics.Mf) || modified
	modified = stage.CleanSlice(&dynamics.Sf) || modified
	modified = stage.CleanSlice(&dynamics.Sfp) || modified
	modified = stage.CleanSlice(&dynamics.Sfpp) || modified
	modified = stage.CleanSlice(&dynamics.Fp) || modified
	modified = stage.CleanSlice(&dynamics.Rf) || modified
	modified = stage.CleanSlice(&dynamics.Rfz) || modified
	modified = stage.CleanSlice(&dynamics.Sfz) || modified
	modified = stage.CleanSlice(&dynamics.Sffz) || modified
	modified = stage.CleanSlice(&dynamics.Fz) || modified
	modified = stage.CleanSlice(&dynamics.N) || modified
	modified = stage.CleanSlice(&dynamics.Pf) || modified
	modified = stage.CleanSlice(&dynamics.Sfzp) || modified
	modified = stage.CleanSlice(&dynamics.OtherDynamics) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Encoding
func (encoding *Encoding) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&encoding.Encoder) || modified
	modified = stage.CleanSlice(&encoding.Supports) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Identification
func (identification *Identification) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&identification.Creator) || modified
	modified = stage.CleanSlice(&identification.Rights) || modified
	modified = stage.CleanSlice(&identification.Relation) || modified
	// insertion point per field
	modified = stage.CleanPointer(&identification.Encoding) || modified
	modified = stage.CleanPointer(&identification.Miscellaneous) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Miscellaneous
func (miscellaneous *Miscellaneous) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&miscellaneous.MiscellaneousField) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NameDisplay
func (namedisplay *NameDisplay) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&namedisplay.DisplayText) || modified
	modified = stage.CleanSlice(&namedisplay.AccidentalText) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PageLayout
func (pagelayout *PageLayout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pagelayout.PageMargins) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Play
func (play *Play) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&play.OtherPlay) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SystemDividers
func (systemdividers *SystemDividers) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&systemdividers.LeftDivider) || modified
	modified = stage.CleanPointer(&systemdividers.RightDivider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SystemLayout
func (systemlayout *SystemLayout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&systemlayout.SystemMargins) || modified
	modified = stage.CleanPointer(&systemlayout.SystemDividers) || modified
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
