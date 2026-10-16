// generated code - do not edit
package attributes

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Attributes_WOP struct {
	// insertion point

	Name string

	Divisions string

	Staves int

	Instruments int
}

func (from *Attributes) GongCopyBasicFields(to *Attributes) {
	// insertion point
	to.Name = from.Name
	to.Divisions = from.Divisions
	to.Staves = from.Staves
	to.Instruments = from.Instruments
}

type AttributesDirective_WOP struct {
	// insertion point

	Name string

	Lang string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Value string
}

func (from *AttributesDirective) GongCopyBasicFields(to *AttributesDirective) {
	// insertion point
	*to = *from
}

type BeatRepeat_WOP struct {
	// insertion point

	Name string

	Slashes int

	ExceptVoice string
}

func (from *BeatRepeat) GongCopyBasicFields(to *BeatRepeat) {
	// insertion point
	*to = *from
}

type Cancel_WOP struct {
	// insertion point

	Name string

	Location EnumCancelLocation

	Value int
}

func (from *Cancel) GongCopyBasicFields(to *Cancel) {
	// insertion point
	*to = *from
}

type Clef_WOP struct {
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

	Id string

	Sign EnumClefSign

	Line int

	ClefOctaveChange int
}

func (from *Clef) GongCopyBasicFields(to *Clef) {
	// insertion point
	*to = *from
}

type Double_WOP struct {
	// insertion point

	Name string
}

func (from *Double) GongCopyBasicFields(to *Double) {
	// insertion point
	*to = *from
}

type ForPart_WOP struct {
	// insertion point

	Name string

	Number int

	Id string
}

func (from *ForPart) GongCopyBasicFields(to *ForPart) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Id = from.Id
}

type Interchangeable_WOP struct {
	// insertion point

	Name string

	Symbol EnumTimeSymbol

	Separator EnumTimeSeparator

	TimeRelation EnumTimeRelation

	Beats string

	BeatType string
}

func (from *Interchangeable) GongCopyBasicFields(to *Interchangeable) {
	// insertion point
	*to = *from
}

type Key_WOP struct {
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

	Id string

	Fifths int

	Mode string

	KeyStep EnumStep

	KeyAlter string
}

func (from *Key) GongCopyBasicFields(to *Key) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
	to.Fifths = from.Fifths
	to.Mode = from.Mode
	to.KeyStep = from.KeyStep
	to.KeyAlter = from.KeyAlter
}

type KeyAccidental_WOP struct {
	// insertion point

	Name string

	Smufl string
}

func (from *KeyAccidental) GongCopyBasicFields(to *KeyAccidental) {
	// insertion point
	*to = *from
}

type KeyOctave_WOP struct {
	// insertion point

	Name string

	Number int

	Value int
}

func (from *KeyOctave) GongCopyBasicFields(to *KeyOctave) {
	// insertion point
	*to = *from
}

type LineDetail_WOP struct {
	// insertion point

	Name string

	Line int

	Width string

	Color string
}

func (from *LineDetail) GongCopyBasicFields(to *LineDetail) {
	// insertion point
	*to = *from
}

type MeasureRepeat_WOP struct {
	// insertion point

	Name string

	Slashes int

	Value string
}

func (from *MeasureRepeat) GongCopyBasicFields(to *MeasureRepeat) {
	// insertion point
	*to = *from
}

type MeasureStyle_WOP struct {
	// insertion point

	Name string

	Number int

	FontFamily string

	FontSize string

	Color string

	Id string
}

func (from *MeasureStyle) GongCopyBasicFields(to *MeasureStyle) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
}

type MultipleRest_WOP struct {
	// insertion point

	Name string

	Value int
}

func (from *MultipleRest) GongCopyBasicFields(to *MultipleRest) {
	// insertion point
	*to = *from
}

type PartClef_WOP struct {
	// insertion point

	Name string

	Sign EnumClefSign

	Line int

	ClefOctaveChange int
}

func (from *PartClef) GongCopyBasicFields(to *PartClef) {
	// insertion point
	*to = *from
}

type PartSymbol_WOP struct {
	// insertion point

	Name string

	TopStaff int

	BottomStaff int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string

	Value EnumGroupSymbolValue
}

func (from *PartSymbol) GongCopyBasicFields(to *PartSymbol) {
	// insertion point
	*to = *from
}

type PartTranspose_WOP struct {
	// insertion point

	Name string

	Diatonic int

	Chromatic string

	OctaveChange int
}

func (from *PartTranspose) GongCopyBasicFields(to *PartTranspose) {
	// insertion point
	to.Name = from.Name
	to.Diatonic = from.Diatonic
	to.Chromatic = from.Chromatic
	to.OctaveChange = from.OctaveChange
}

type Slash_WOP struct {
	// insertion point

	Name string

	ExceptVoice string
}

func (from *Slash) GongCopyBasicFields(to *Slash) {
	// insertion point
	*to = *from
}

type StaffDetails_WOP struct {
	// insertion point

	Name string

	Number int

	ShowFrets EnumShowFrets

	StaffType EnumStaffType

	StaffLines int

	Capo int
}

func (from *StaffDetails) GongCopyBasicFields(to *StaffDetails) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.ShowFrets = from.ShowFrets
	to.StaffType = from.StaffType
	to.StaffLines = from.StaffLines
	to.Capo = from.Capo
}

type StaffSize_WOP struct {
	// insertion point

	Name string

	Scaling string

	Value string
}

func (from *StaffSize) GongCopyBasicFields(to *StaffSize) {
	// insertion point
	*to = *from
}

type StaffTuning_WOP struct {
	// insertion point

	Name string

	Line int

	TuningStep EnumStep

	TuningAlter string

	TuningOctave int
}

func (from *StaffTuning) GongCopyBasicFields(to *StaffTuning) {
	// insertion point
	*to = *from
}

type Time_WOP struct {
	// insertion point

	Name string

	Number int

	Symbol EnumTimeSymbol

	Separator EnumTimeSeparator

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	Id string

	Beats string

	BeatType string

	SenzaMisura string
}

func (from *Time) GongCopyBasicFields(to *Time) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Symbol = from.Symbol
	to.Separator = from.Separator
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontSize = from.FontSize
	to.Color = from.Color
	to.Id = from.Id
	to.Beats = from.Beats
	to.BeatType = from.BeatType
	to.SenzaMisura = from.SenzaMisura
}

type Transpose_WOP struct {
	// insertion point

	Name string

	Number int

	Id string

	Diatonic int

	Chromatic string

	OctaveChange int
}

func (from *Transpose) GongCopyBasicFields(to *Transpose) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Id = from.Id
	to.Diatonic = from.Diatonic
	to.Chromatic = from.Chromatic
	to.OctaveChange = from.OctaveChange
}

// end of insertion point
