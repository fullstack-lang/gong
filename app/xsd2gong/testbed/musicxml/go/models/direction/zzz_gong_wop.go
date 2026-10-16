// generated code - do not edit
package direction

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Accord_WOP struct {
	// insertion point

	Name string

	String int

	TuningAlter string

	TuningOctave int
}

func (from *Accord) GongCopyBasicFields(to *Accord) {
	// insertion point
	*to = *from
}

type AccordionRegistration_WOP struct {
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

	AccordionMiddle int
}

func (from *AccordionRegistration) GongCopyBasicFields(to *AccordionRegistration) {
	// insertion point
	*to = *from
}

type Barre_WOP struct {
	// insertion point

	Name string

	Color string
}

func (from *Barre) GongCopyBasicFields(to *Barre) {
	// insertion point
	*to = *from
}

type Bass_WOP struct {
	// insertion point

	Name string

	Arrangement EnumHarmonyArrangement
}

func (from *Bass) GongCopyBasicFields(to *Bass) {
	// insertion point
	to.Name = from.Name
	to.Arrangement = from.Arrangement
}

type BassStep_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *BassStep) GongCopyBasicFields(to *BassStep) {
	// insertion point
	*to = *from
}

type BeatUnitTied_WOP struct {
	// insertion point

	Name string
}

func (from *BeatUnitTied) GongCopyBasicFields(to *BeatUnitTied) {
	// insertion point
	*to = *from
}

type Beater_WOP struct {
	// insertion point

	Name string

	Tip EnumTipDirection

	Value EnumBeaterValue
}

func (from *Beater) GongCopyBasicFields(to *Beater) {
	// insertion point
	*to = *from
}

type Bracket_WOP struct {
	// insertion point

	Name string

	Number int

	LineEnd EnumLineEnd

	EndLength string

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string
}

func (from *Bracket) GongCopyBasicFields(to *Bracket) {
	// insertion point
	*to = *from
}

type Dashes_WOP struct {
	// insertion point

	Name string

	Number int

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string
}

func (from *Dashes) GongCopyBasicFields(to *Dashes) {
	// insertion point
	*to = *from
}

type Degree_WOP struct {
	// insertion point

	Name string
}

func (from *Degree) GongCopyBasicFields(to *Degree) {
	// insertion point
	to.Name = from.Name
}

type DegreeAlter_WOP struct {
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

func (from *DegreeAlter) GongCopyBasicFields(to *DegreeAlter) {
	// insertion point
	*to = *from
}

type DegreeType_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumDegreeTypeValue
}

func (from *DegreeType) GongCopyBasicFields(to *DegreeType) {
	// insertion point
	*to = *from
}

type DegreeValue_WOP struct {
	// insertion point

	Name string

	Symbol EnumDegreeSymbolValue

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value int
}

func (from *DegreeValue) GongCopyBasicFields(to *DegreeValue) {
	// insertion point
	*to = *from
}

type Direction_WOP struct {
	// insertion point

	Name string

	Id string

	Voice string

	Staff int
}

func (from *Direction) GongCopyBasicFields(to *Direction) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Voice = from.Voice
	to.Staff = from.Staff
}

type DirectionType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DirectionType) GongCopyBasicFields(to *DirectionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Effect_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumEffectValue
}

func (from *Effect) GongCopyBasicFields(to *Effect) {
	// insertion point
	*to = *from
}

type Feature_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *Feature) GongCopyBasicFields(to *Feature) {
	// insertion point
	*to = *from
}

type FirstFret_WOP struct {
	// insertion point

	Name string

	Text string

	Value int
}

func (from *FirstFret) GongCopyBasicFields(to *FirstFret) {
	// insertion point
	*to = *from
}

type Frame_WOP struct {
	// insertion point

	Name string

	Height string

	Width string

	Unplayed string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string

	FrameStrings int

	FrameFrets int
}

func (from *Frame) GongCopyBasicFields(to *Frame) {
	// insertion point
	to.Name = from.Name
	to.Height = from.Height
	to.Width = from.Width
	to.Unplayed = from.Unplayed
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.Color = from.Color
	to.Id = from.Id
	to.FrameStrings = from.FrameStrings
	to.FrameFrets = from.FrameFrets
}

type FrameNote_WOP struct {
	// insertion point

	Name string
}

func (from *FrameNote) GongCopyBasicFields(to *FrameNote) {
	// insertion point
	to.Name = from.Name
}

type Glass_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumGlassValue
}

func (from *Glass) GongCopyBasicFields(to *Glass) {
	// insertion point
	*to = *from
}

type Grouping_WOP struct {
	// insertion point

	Name string

	Number string

	MemberOf string

	Id string
}

func (from *Grouping) GongCopyBasicFields(to *Grouping) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.MemberOf = from.MemberOf
	to.Id = from.Id
}

type Harmony_WOP struct {
	// insertion point

	Name string

	Type EnumHarmonyType

	Arrangement EnumHarmonyArrangement

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string

	Staff int
}

func (from *Harmony) GongCopyBasicFields(to *Harmony) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Arrangement = from.Arrangement
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
	to.Staff = from.Staff
}

type HarmonyAlter_WOP struct {
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

func (from *HarmonyAlter) GongCopyBasicFields(to *HarmonyAlter) {
	// insertion point
	*to = *from
}

type HarpPedals_WOP struct {
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
}

func (from *HarpPedals) GongCopyBasicFields(to *HarpPedals) {
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
}

type Image_WOP struct {
	// insertion point

	Name string

	Source string

	Type string

	Height string

	Width string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Id string
}

func (from *Image) GongCopyBasicFields(to *Image) {
	// insertion point
	*to = *from
}

type InstrumentChange_WOP struct {
	// insertion point

	Name string

	Id string

	InstrumentSound string

	Ensemble string
}

func (from *InstrumentChange) GongCopyBasicFields(to *InstrumentChange) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.InstrumentSound = from.InstrumentSound
	to.Ensemble = from.Ensemble
}

type Inversion_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value int
}

func (from *Inversion) GongCopyBasicFields(to *Inversion) {
	// insertion point
	*to = *from
}

type Kind_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumKindValue
}

func (from *Kind) GongCopyBasicFields(to *Kind) {
	// insertion point
	*to = *from
}

type Listening_WOP struct {
	// insertion point

	Name string
}

func (from *Listening) GongCopyBasicFields(to *Listening) {
	// insertion point
	to.Name = from.Name
}

type MeasureNumbering_WOP struct {
	// insertion point

	Name string

	Staff int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value EnumMeasureNumberingValue
}

func (from *MeasureNumbering) GongCopyBasicFields(to *MeasureNumbering) {
	// insertion point
	*to = *from
}

type Membrane_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumMembraneValue
}

func (from *Membrane) GongCopyBasicFields(to *Membrane) {
	// insertion point
	*to = *from
}

type Metal_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumMetalValue
}

func (from *Metal) GongCopyBasicFields(to *Metal) {
	// insertion point
	*to = *from
}

type Metronome_WOP struct {
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

	MetronomeRelation string
}

func (from *Metronome) GongCopyBasicFields(to *Metronome) {
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
	to.MetronomeRelation = from.MetronomeRelation
}

type MetronomeBeam_WOP struct {
	// insertion point

	Name string

	Number int

	Value EnumBeamValue
}

func (from *MetronomeBeam) GongCopyBasicFields(to *MetronomeBeam) {
	// insertion point
	*to = *from
}

type MetronomeNote_WOP struct {
	// insertion point

	Name string
}

func (from *MetronomeNote) GongCopyBasicFields(to *MetronomeNote) {
	// insertion point
	to.Name = from.Name
}

type MetronomeTied_WOP struct {
	// insertion point

	Name string
}

func (from *MetronomeTied) GongCopyBasicFields(to *MetronomeTied) {
	// insertion point
	*to = *from
}

type MetronomeTuplet_WOP struct {
	// insertion point

	Name string

	ShowNumber EnumShowTuplet

	ActualNotes int

	NormalNotes int
}

func (from *MetronomeTuplet) GongCopyBasicFields(to *MetronomeTuplet) {
	// insertion point
	*to = *from
}

type Numeral_WOP struct {
	// insertion point

	Name string
}

func (from *Numeral) GongCopyBasicFields(to *Numeral) {
	// insertion point
	to.Name = from.Name
}

type NumeralKey_WOP struct {
	// insertion point

	Name string

	NumeralFifths int

	NumeralMode EnumNumeralMode
}

func (from *NumeralKey) GongCopyBasicFields(to *NumeralKey) {
	// insertion point
	*to = *from
}

type NumeralRoot_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value int
}

func (from *NumeralRoot) GongCopyBasicFields(to *NumeralRoot) {
	// insertion point
	*to = *from
}

type OctaveShift_WOP struct {
	// insertion point

	Name string

	Type EnumUpDownStopContinue

	Number int

	Size int

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
}

func (from *OctaveShift) GongCopyBasicFields(to *OctaveShift) {
	// insertion point
	*to = *from
}

type Offset_WOP struct {
	// insertion point

	Name string

	Value string
}

func (from *Offset) GongCopyBasicFields(to *Offset) {
	// insertion point
	*to = *from
}

type OtherDirection_WOP struct {
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

	Id string

	Value string
}

func (from *OtherDirection) GongCopyBasicFields(to *OtherDirection) {
	// insertion point
	*to = *from
}

type OtherListening_WOP struct {
	// insertion point

	Name string

	Type string

	Player string

	TimeOnly string

	Value string
}

func (from *OtherListening) GongCopyBasicFields(to *OtherListening) {
	// insertion point
	*to = *from
}

type Pedal_WOP struct {
	// insertion point

	Name string

	Type EnumPedalType

	Number int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string
}

func (from *Pedal) GongCopyBasicFields(to *Pedal) {
	// insertion point
	*to = *from
}

type PedalTuning_WOP struct {
	// insertion point

	Name string

	PedalAlter string
}

func (from *PedalTuning) GongCopyBasicFields(to *PedalTuning) {
	// insertion point
	*to = *from
}

type PerMinute_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string

	Value string
}

func (from *PerMinute) GongCopyBasicFields(to *PerMinute) {
	// insertion point
	*to = *from
}

type Percussion_WOP struct {
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

	StickLocation EnumStickLocation
}

func (from *Percussion) GongCopyBasicFields(to *Percussion) {
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
	to.StickLocation = from.StickLocation
}

type Pitched_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumPitchedValue
}

func (from *Pitched) GongCopyBasicFields(to *Pitched) {
	// insertion point
	*to = *from
}

type PrincipalVoice_WOP struct {
	// insertion point

	Name string

	Symbol EnumPrincipalVoiceSymbol

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

func (from *PrincipalVoice) GongCopyBasicFields(to *PrincipalVoice) {
	// insertion point
	*to = *from
}

type Print_WOP struct {
	// insertion point

	Name string

	StaffSpacing string

	BlankPage int

	PageNumber string

	Id string
}

func (from *Print) GongCopyBasicFields(to *Print) {
	// insertion point
	to.Name = from.Name
	to.StaffSpacing = from.StaffSpacing
	to.BlankPage = from.BlankPage
	to.PageNumber = from.PageNumber
	to.Id = from.Id
}

type Root_WOP struct {
	// insertion point

	Name string
}

func (from *Root) GongCopyBasicFields(to *Root) {
	// insertion point
	to.Name = from.Name
}

type RootStep_WOP struct {
	// insertion point

	Name string

	Text string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string
}

func (from *RootStep) GongCopyBasicFields(to *RootStep) {
	// insertion point
	*to = *from
}

type Scordatura_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Scordatura) GongCopyBasicFields(to *Scordatura) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Sound_WOP struct {
	// insertion point

	Name string

	Tempo string

	Dynamics string

	Segno string

	Dalsegno string

	Coda string

	Tocoda string

	Divisions string

	Fine string

	TimeOnly string

	Pan string

	Elevation string

	DamperPedal string

	SoftPedal string

	SostenutoPedal string

	Id string
}

func (from *Sound) GongCopyBasicFields(to *Sound) {
	// insertion point
	to.Name = from.Name
	to.Tempo = from.Tempo
	to.Dynamics = from.Dynamics
	to.Segno = from.Segno
	to.Dalsegno = from.Dalsegno
	to.Coda = from.Coda
	to.Tocoda = from.Tocoda
	to.Divisions = from.Divisions
	to.Fine = from.Fine
	to.TimeOnly = from.TimeOnly
	to.Pan = from.Pan
	to.Elevation = from.Elevation
	to.DamperPedal = from.DamperPedal
	to.SoftPedal = from.SoftPedal
	to.SostenutoPedal = from.SostenutoPedal
	to.Id = from.Id
}

type StaffDivide_WOP struct {
	// insertion point

	Name string

	Type EnumStaffDivideSymbol

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string
}

func (from *StaffDivide) GongCopyBasicFields(to *StaffDivide) {
	// insertion point
	*to = *from
}

type Stick_WOP struct {
	// insertion point

	Name string

	Tip EnumTipDirection

	StickType EnumStickType

	StickMaterial EnumStickMaterial
}

func (from *Stick) GongCopyBasicFields(to *Stick) {
	// insertion point
	*to = *from
}

type StringMute_WOP struct {
	// insertion point

	Name string

	Type EnumOnOff

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string
}

func (from *StringMute) GongCopyBasicFields(to *StringMute) {
	// insertion point
	*to = *from
}

type StyleText_WOP struct {
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

func (from *StyleText) GongCopyBasicFields(to *StyleText) {
	// insertion point
	*to = *from
}

type Swing_WOP struct {
	// insertion point

	Name string

	First int

	Second int

	SwingStyle string
}

func (from *Swing) GongCopyBasicFields(to *Swing) {
	// insertion point
	*to = *from
}

type Sync_WOP struct {
	// insertion point

	Name string

	Type EnumSyncType

	Latency int

	Player string

	TimeOnly string
}

func (from *Sync) GongCopyBasicFields(to *Sync) {
	// insertion point
	*to = *from
}

type TimeModification_WOP struct {
	// insertion point

	Name string

	ActualNotes int

	NormalNotes int
}

func (from *TimeModification) GongCopyBasicFields(to *TimeModification) {
	// insertion point
	*to = *from
}

type Timpani_WOP struct {
	// insertion point

	Name string

	Smufl string
}

func (from *Timpani) GongCopyBasicFields(to *Timpani) {
	// insertion point
	*to = *from
}

type VirtualInstrument_WOP struct {
	// insertion point

	Name string

	VirtualLibrary string

	VirtualName string
}

func (from *VirtualInstrument) GongCopyBasicFields(to *VirtualInstrument) {
	// insertion point
	*to = *from
}

type Wedge_WOP struct {
	// insertion point

	Name string

	Type EnumWedgeType

	Number int

	Spread string

	DashLength string

	SpaceLength string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Id string
}

func (from *Wedge) GongCopyBasicFields(to *Wedge) {
	// insertion point
	*to = *from
}

type Wood_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value EnumWoodValue
}

func (from *Wood) GongCopyBasicFields(to *Wood) {
	// insertion point
	*to = *from
}

// end of insertion point
