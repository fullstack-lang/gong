// generated code - do not edit
package music

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type MusicAbstract_WOP struct {
	// insertion point

	Name string

	IsChecked bool

	PitchHeight float64

	NbOfBeatsInTheme int

	BeatsPerSecond float64

	FirstVoiceShiftX float64

	FirstVoiceShiftY float64

	PitchDifference int

	Level float64

	ActualBeatsTemporalShift int

	IsMinor bool

	ThemeBinaryEncoding int

	BezierControlLengthRatio float64

	NbPitchLines int

	NbBeatLines int

	OriginX float64

	OriginY float64

	ScoreScale float64

	ShowFirstVoice bool

	ShowFirstVoiceShiftRight bool

	ShowSecondVoice bool

	ShowSecondVoiceShiftRight bool

	ShowFirstVoiceNotes bool

	ShowFirstVoiceNotesShiftRight bool

	ShowSecondVoiceNotes bool

	ShowSecondVoiceNotesShiftRight bool

	IsComposerNodeExpanded bool
}

func (from *MusicAbstract) GongCopyBasicFields(to *MusicAbstract) {
	// insertion point
	to.Name = from.Name
	to.IsChecked = from.IsChecked
	to.PitchHeight = from.PitchHeight
	to.NbOfBeatsInTheme = from.NbOfBeatsInTheme
	to.BeatsPerSecond = from.BeatsPerSecond
	to.FirstVoiceShiftX = from.FirstVoiceShiftX
	to.FirstVoiceShiftY = from.FirstVoiceShiftY
	to.PitchDifference = from.PitchDifference
	to.Level = from.Level
	to.ActualBeatsTemporalShift = from.ActualBeatsTemporalShift
	to.IsMinor = from.IsMinor
	to.ThemeBinaryEncoding = from.ThemeBinaryEncoding
	to.BezierControlLengthRatio = from.BezierControlLengthRatio
	to.NbPitchLines = from.NbPitchLines
	to.NbBeatLines = from.NbBeatLines
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.ScoreScale = from.ScoreScale
	to.ShowFirstVoice = from.ShowFirstVoice
	to.ShowFirstVoiceShiftRight = from.ShowFirstVoiceShiftRight
	to.ShowSecondVoice = from.ShowSecondVoice
	to.ShowSecondVoiceShiftRight = from.ShowSecondVoiceShiftRight
	to.ShowFirstVoiceNotes = from.ShowFirstVoiceNotes
	to.ShowFirstVoiceNotesShiftRight = from.ShowFirstVoiceNotesShiftRight
	to.ShowSecondVoiceNotes = from.ShowSecondVoiceNotes
	to.ShowSecondVoiceNotesShiftRight = from.ShowSecondVoiceNotesShiftRight
	to.IsComposerNodeExpanded = from.IsComposerNodeExpanded
}

// end of insertion point
