// generated code - do not edit
package score

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
// Clean garbage collect unstaged instances that are referenced by Defaults
func (defaults *Defaults) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&defaults.LyricFont) || modified
	modified = stage.CleanSlice(&defaults.LyricLanguage) || modified
	// insertion point per field
	modified = stage.CleanPointer(&defaults.MusicFont) || modified
	modified = stage.CleanPointer(&defaults.WordFont) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InstrumentLink
func (instrumentlink *InstrumentLink) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&instrumentlink.RefScoreInstrument) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Measure
func (measure *Measure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&measure.MusicData) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part
func (part *Part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&part.Measure) || modified
	// insertion point per field
	modified = stage.CleanPointer(&part.RefScorePart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PartGroup
func (partgroup *PartGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&partgroup.GroupName) || modified
	modified = stage.CleanPointer(&partgroup.GroupAbbreviation) || modified
	modified = stage.CleanPointer(&partgroup.GroupSymbol) || modified
	modified = stage.CleanPointer(&partgroup.GroupBarline) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PartLink
func (partlink *PartLink) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&partlink.InstrumentLink) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartList
func (partlist *PartList) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&partlist.PartGroup) || modified
	modified = stage.CleanSlice(&partlist.ScorePart) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ScorePart
func (scorepart *ScorePart) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scorepart.PartLink) || modified
	modified = stage.CleanSlice(&scorepart.ScoreInstrument) || modified
	modified = stage.CleanSlice(&scorepart.Player) || modified
	// insertion point per field
	modified = stage.CleanPointer(&scorepart.PartName) || modified
	modified = stage.CleanPointer(&scorepart.PartAbbreviation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ScorePartwise
func (scorepartwise *ScorePartwise) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scorepartwise.Credit) || modified
	modified = stage.CleanSlice(&scorepartwise.Part) || modified
	// insertion point per field
	modified = stage.CleanPointer(&scorepartwise.Work) || modified
	modified = stage.CleanPointer(&scorepartwise.Defaults) || modified
	modified = stage.CleanPointer(&scorepartwise.PartList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ScoreTimewise
func (scoretimewise *ScoreTimewise) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scoretimewise.Credit) || modified
	modified = stage.CleanSlice(&scoretimewise.Measure) || modified
	// insertion point per field
	modified = stage.CleanPointer(&scoretimewise.Work) || modified
	modified = stage.CleanPointer(&scoretimewise.Defaults) || modified
	modified = stage.CleanPointer(&scoretimewise.PartList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ScoreTimewiseMeasure
func (scoretimewisemeasure *ScoreTimewiseMeasure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scoretimewisemeasure.Part) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ScoreTimewiseMeasurePart
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&scoretimewisemeasurepart.RefScorePart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Work
func (work *Work) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&work.Opus) || modified
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
