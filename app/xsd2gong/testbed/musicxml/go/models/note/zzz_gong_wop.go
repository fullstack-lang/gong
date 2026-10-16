// generated code - do not edit
package note

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Accidental_WOP struct {
	// insertion point

	Name string

	Smufl string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *Accidental) GongCopyBasicFields(to *Accidental) {
	// insertion point
	*to = *from
}

type AccidentalMark_WOP struct {
	// insertion point

	Name string

	Smufl string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string
}

func (from *AccidentalMark) GongCopyBasicFields(to *AccidentalMark) {
	// insertion point
	*to = *from
}

type Arpeggiate_WOP struct {
	// insertion point

	Name string

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string
}

func (from *Arpeggiate) GongCopyBasicFields(to *Arpeggiate) {
	// insertion point
	*to = *from
}

type Arrow_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Smufl string

	ArrowDirection EnumArrowDirection

	ArrowStyle EnumArrowStyle

	CircularArrow EnumCircularArrow
}

func (from *Arrow) GongCopyBasicFields(to *Arrow) {
	// insertion point
	*to = *from
}

type Articulations_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Articulations) GongCopyBasicFields(to *Articulations) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Assess_WOP struct {
	// insertion point

	Name string

	Player string

	TimeOnly string
}

func (from *Assess) GongCopyBasicFields(to *Assess) {
	// insertion point
	*to = *from
}

type Backup_WOP struct {
	// insertion point

	Name string

	Duration string
}

func (from *Backup) GongCopyBasicFields(to *Backup) {
	// insertion point
	*to = *from
}

type Beam_WOP struct {
	// insertion point

	Name string

	Number int

	Fan EnumFan

	Color string

	Id string
}

func (from *Beam) GongCopyBasicFields(to *Beam) {
	// insertion point
	*to = *from
}

type Bend_WOP struct {
	// insertion point

	Name string

	Shape EnumBendShape

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Beats string

	FirstBeat string

	LastBeat string

	BendAlter string
}

func (from *Bend) GongCopyBasicFields(to *Bend) {
	// insertion point
	to.Name = from.Name
	to.Shape = from.Shape
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Beats = from.Beats
	to.FirstBeat = from.FirstBeat
	to.LastBeat = from.LastBeat
	to.BendAlter = from.BendAlter
}

type BreathMark_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumBreathMarkValue
}

func (from *BreathMark) GongCopyBasicFields(to *BreathMark) {
	// insertion point
	*to = *from
}

type Caesura_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumCaesuraValue
}

func (from *Caesura) GongCopyBasicFields(to *Caesura) {
	// insertion point
	*to = *from
}

type Elision_WOP struct {
	// insertion point

	Name string

	Smufl string

	FontFamily string

	FontSize string

	Color string

	Value string
}

func (from *Elision) GongCopyBasicFields(to *Elision) {
	// insertion point
	*to = *from
}

type EmptyLine_WOP struct {
	// insertion point

	Name string

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *EmptyLine) GongCopyBasicFields(to *EmptyLine) {
	// insertion point
	*to = *from
}

type Extend_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string
}

func (from *Extend) GongCopyBasicFields(to *Extend) {
	// insertion point
	*to = *from
}

type Figure_WOP struct {
	// insertion point

	Name string
}

func (from *Figure) GongCopyBasicFields(to *Figure) {
	// insertion point
	to.Name = from.Name
}

type FiguredBass_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string

	Duration string
}

func (from *FiguredBass) GongCopyBasicFields(to *FiguredBass) {
	// insertion point
	to.Name = from.Name
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
	to.Duration = from.Duration
}

type Forward_WOP struct {
	// insertion point

	Name string

	Duration string

	Voice string

	Staff int
}

func (from *Forward) GongCopyBasicFields(to *Forward) {
	// insertion point
	*to = *from
}

type Glissando_WOP struct {
	// insertion point

	Name string

	Number int

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string

	Value string
}

func (from *Glissando) GongCopyBasicFields(to *Glissando) {
	// insertion point
	*to = *from
}

type Grace_WOP struct {
	// insertion point

	Name string

	StealTimePrevious string

	StealTimeFollowing string

	MakeTime string
}

func (from *Grace) GongCopyBasicFields(to *Grace) {
	// insertion point
	*to = *from
}

type HammerOnPullOff_WOP struct {
	// insertion point

	Name string

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value string
}

func (from *HammerOnPullOff) GongCopyBasicFields(to *HammerOnPullOff) {
	// insertion point
	*to = *from
}

type Handbell_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumHandbellValue
}

func (from *Handbell) GongCopyBasicFields(to *Handbell) {
	// insertion point
	*to = *from
}

type HarmonClosed_WOP struct {
	// insertion point

	Name string

	Location EnumHarmonClosedLocation

	Value EnumHarmonClosedValue
}

func (from *HarmonClosed) GongCopyBasicFields(to *HarmonClosed) {
	// insertion point
	*to = *from
}

type HarmonMute_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *HarmonMute) GongCopyBasicFields(to *HarmonMute) {
	// insertion point
	to.Name = from.Name
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
}

type Harmonic_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *Harmonic) GongCopyBasicFields(to *Harmonic) {
	// insertion point
	*to = *from
}

type HeelToe_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *HeelToe) GongCopyBasicFields(to *HeelToe) {
	// insertion point
	*to = *from
}

type Hole_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	HoleType string

	HoleShape string
}

func (from *Hole) GongCopyBasicFields(to *Hole) {
	// insertion point
	to.Name = from.Name
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.HoleType = from.HoleType
	to.HoleShape = from.HoleShape
}

type HoleClosed_WOP struct {
	// insertion point

	Name string

	Location EnumHoleClosedLocation

	Value EnumHoleClosedValue
}

func (from *HoleClosed) GongCopyBasicFields(to *HoleClosed) {
	// insertion point
	*to = *from
}

type Instrument_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Instrument) GongCopyBasicFields(to *Instrument) {
	// insertion point
	*to = *from
}

type Listen_WOP struct {
	// insertion point

	Name string
}

func (from *Listen) GongCopyBasicFields(to *Listen) {
	// insertion point
	to.Name = from.Name
}

type Lyric_WOP struct {
	// insertion point

	Name string

	Number string

	NameAttr string

	TimeOnly string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string

	Syllabic EnumSyllabic
}

func (from *Lyric) GongCopyBasicFields(to *Lyric) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.NameAttr = from.NameAttr
	to.TimeOnly = from.TimeOnly
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.Color = from.Color
	to.Id = from.Id
	to.Syllabic = from.Syllabic
}

type Mordent_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Beats string

	SecondBeat string

	LastBeat string
}

func (from *Mordent) GongCopyBasicFields(to *Mordent) {
	// insertion point
	*to = *from
}

type NonArpeggiate_WOP struct {
	// insertion point

	Name string

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string
}

func (from *NonArpeggiate) GongCopyBasicFields(to *NonArpeggiate) {
	// insertion point
	*to = *from
}

type Notations_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Notations) GongCopyBasicFields(to *Notations) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Note_WOP struct {
	// insertion point

	Name string

	Dynamics string

	EndDynamics string

	Attack string

	Release string

	TimeOnly string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string

	Duration string

	Voice string

	Staff int
}

func (from *Note) GongCopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.Dynamics = from.Dynamics
	to.EndDynamics = from.EndDynamics
	to.Attack = from.Attack
	to.Release = from.Release
	to.TimeOnly = from.TimeOnly
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
	to.Duration = from.Duration
	to.Voice = from.Voice
	to.Staff = from.Staff
}

type NoteType_WOP struct {
	// insertion point

	Name string
}

func (from *NoteType) GongCopyBasicFields(to *NoteType) {
	// insertion point
	*to = *from
}

type Notehead_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string

	Color string

	Smufl string

	Value EnumNoteheadValue
}

func (from *Notehead) GongCopyBasicFields(to *Notehead) {
	// insertion point
	*to = *from
}

type NoteheadText_WOP struct {
	// insertion point

	Name string
}

func (from *NoteheadText) GongCopyBasicFields(to *NoteheadText) {
	// insertion point
	*to = *from
}

type Ornaments_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Ornaments) GongCopyBasicFields(to *Ornaments) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type OtherNotation_WOP struct {
	// insertion point

	Name string

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Smufl string

	Id string

	Value string
}

func (from *OtherNotation) GongCopyBasicFields(to *OtherNotation) {
	// insertion point
	*to = *from
}

type OtherPlacementText_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Smufl string

	Value string
}

func (from *OtherPlacementText) GongCopyBasicFields(to *OtherPlacementText) {
	// insertion point
	*to = *from
}

type Pitch_WOP struct {
	// insertion point

	Name string

	Alter string

	Octave int
}

func (from *Pitch) GongCopyBasicFields(to *Pitch) {
	// insertion point
	*to = *from
}

type PlacementText_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value string
}

func (from *PlacementText) GongCopyBasicFields(to *PlacementText) {
	// insertion point
	*to = *from
}

type Release_WOP struct {
	// insertion point

	Name string

	Offset string
}

func (from *Release) GongCopyBasicFields(to *Release) {
	// insertion point
	*to = *from
}

type Rest_WOP struct {
	// insertion point

	Name string

	DisplayOctave int
}

func (from *Rest) GongCopyBasicFields(to *Rest) {
	// insertion point
	*to = *from
}

type Slide_WOP struct {
	// insertion point

	Name string

	Number int

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Beats string

	FirstBeat string

	LastBeat string

	Id string

	Value string
}

func (from *Slide) GongCopyBasicFields(to *Slide) {
	// insertion point
	*to = *from
}

type Slur_WOP struct {
	// insertion point

	Name string

	Number int

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	BezierX string

	BezierY string

	BezierX2 string

	BezierY2 string

	BezierOffset string

	BezierOffset2 string

	Color string

	Id string
}

func (from *Slur) GongCopyBasicFields(to *Slur) {
	// insertion point
	*to = *from
}

type Stem_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Value EnumStemValue
}

func (from *Stem) GongCopyBasicFields(to *Stem) {
	// insertion point
	*to = *from
}

type StrongAccent_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *StrongAccent) GongCopyBasicFields(to *StrongAccent) {
	// insertion point
	*to = *from
}

type Tap_WOP struct {
	// insertion point

	Name string

	Hand EnumTapHand

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value string
}

func (from *Tap) GongCopyBasicFields(to *Tap) {
	// insertion point
	*to = *from
}

type Technical_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Technical) GongCopyBasicFields(to *Technical) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type TextElementData_WOP struct {
	// insertion point

	Name string

	Lang string

	FontFamily string

	FontSize string

	Color string

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	Value string
}

func (from *TextElementData) GongCopyBasicFields(to *TextElementData) {
	// insertion point
	*to = *from
}

type Tie_WOP struct {
	// insertion point

	Name string

	TimeOnly string
}

func (from *Tie) GongCopyBasicFields(to *Tie) {
	// insertion point
	*to = *from
}

type Tied_WOP struct {
	// insertion point

	Name string

	Number int

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	BezierX string

	BezierY string

	BezierX2 string

	BezierY2 string

	BezierOffset string

	BezierOffset2 string

	Color string

	Id string
}

func (from *Tied) GongCopyBasicFields(to *Tied) {
	// insertion point
	*to = *from
}

type Tremolo_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Smufl string

	Value int
}

func (from *Tremolo) GongCopyBasicFields(to *Tremolo) {
	// insertion point
	*to = *from
}

type Tuplet_WOP struct {
	// insertion point

	Name string

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Id string
}

func (from *Tuplet) GongCopyBasicFields(to *Tuplet) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.Id = from.Id
}

type TupletDot_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string

	Color string
}

func (from *TupletDot) GongCopyBasicFields(to *TupletDot) {
	// insertion point
	*to = *from
}

type TupletNumber_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string

	Color string

	Value int
}

func (from *TupletNumber) GongCopyBasicFields(to *TupletNumber) {
	// insertion point
	*to = *from
}

type TupletPortion_WOP struct {
	// insertion point

	Name string
}

func (from *TupletPortion) GongCopyBasicFields(to *TupletPortion) {
	// insertion point
	to.Name = from.Name
}

type TupletType_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string

	Color string
}

func (from *TupletType) GongCopyBasicFields(to *TupletType) {
	// insertion point
	*to = *from
}

type Unpitched_WOP struct {
	// insertion point

	Name string

	DisplayOctave int
}

func (from *Unpitched) GongCopyBasicFields(to *Unpitched) {
	// insertion point
	*to = *from
}

type Wait_WOP struct {
	// insertion point

	Name string

	Player string

	TimeOnly string
}

func (from *Wait) GongCopyBasicFields(to *Wait) {
	// insertion point
	*to = *from
}

// end of insertion point
