package models

type MusicAbstract struct {
	Name string

	IsChecked bool

	// PitchHeight is the vertical spacing between chromatic pitches relative to SideLength
	PitchHeight float64

	// NbOfBeatsInTheme is the number of beats per theme/measure (default 16)
	NbOfBeatsInTheme int

	// BeatsPerSecond is the tempo in beats per second
	BeatsPerSecond float64

	// FirstVoiceShiftX is the relative horizontal shift of the first voice
	FirstVoiceShiftX float64

	// FirstVoiceShiftY is the relative vertical shift of the first voice
	FirstVoiceShiftY float64

	// PitchDifference is the transposition in semitones between the 1st and 2nd voice
	PitchDifference int

	// Level is the velocity/volume level
	Level float64

	// ActualBeatsTemporalShift is the temporal shift in beats between voices
	ActualBeatsTemporalShift int

	// IsMinor selects between minor and major scales
	IsMinor bool

	// ThemeBinaryEncoding is the 64-bit integer representing which beats have notes played
	ThemeBinaryEncoding int

	// Ratio of the length of the cubic Bezier control vector to the side length
	BezierControlLengthRatio float64

	// Number of pitch lines (horizontal grid lines)
	NbPitchLines int

	// Number of beat lines (vertical grid lines)
	NbBeatLines int

	// Origin of the score in SVG canvas space
	OriginX float64
	OriginY float64

	// Tree visibility checkboxes ("Composer")
	ShowFirstVoice                 bool
	ShowFirstVoiceShiftRight       bool
	ShowSecondVoice                bool
	ShowSecondVoiceShiftRight      bool
	ShowFirstVoiceNotes            bool
	ShowFirstVoiceNotesShiftRight  bool
	ShowSecondVoiceNotes           bool
	ShowSecondVoiceNotesShiftRight bool

	IsComposerNodeExpanded bool
}

// IsNotePlayed checks whether the note at the specified beat index is played.
func (musicAbstract *MusicAbstract) IsNotePlayed(beatNb int) bool {
	if beatNb < 0 || beatNb > 63 {
		return false
	}
	return musicAbstract.ThemeBinaryEncoding&(1<<beatNb) != 0
}

// ToggleNotePlayed toggles whether the note at the specified beat index is played.
func (musicAbstract *MusicAbstract) ToggleNotePlayed(beatNb int) {
	if beatNb < 0 || beatNb > 63 {
		return
	}
	musicAbstract.ThemeBinaryEncoding ^= 1 << beatNb
}
