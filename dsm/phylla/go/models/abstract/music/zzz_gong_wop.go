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
	*to = *from
}

// end of insertion point
