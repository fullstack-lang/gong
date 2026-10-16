// generated code - do not edit
package common

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AccidentalText_WOP struct {
	// insertion point

	Name string

	Smufl string

	Lang string

	Space string

	Justify EnumLeftCenterRight

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	LineHeight string

	Dir EnumTextDirection

	Enclosure EnumEnclosureShape

	Value EnumAccidentalValue
}

func (from *AccidentalText) GongCopyBasicFields(to *AccidentalText) {
	// insertion point
	*to = *from
}

type Appearance_WOP struct {
	// insertion point

	Name string
}

func (from *Appearance) GongCopyBasicFields(to *Appearance) {
	// insertion point
	to.Name = from.Name
}

type Bookmark_WOP struct {
	// insertion point

	Name string

	Id string

	NameAttr string

	Element string

	Position int
}

func (from *Bookmark) GongCopyBasicFields(to *Bookmark) {
	// insertion point
	*to = *from
}

type Coda_WOP struct {
	// insertion point

	Name string

	Smufl string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Id string
}

func (from *Coda) GongCopyBasicFields(to *Coda) {
	// insertion point
	*to = *from
}

type Distance_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *Distance) GongCopyBasicFields(to *Distance) {
	// insertion point
	*to = *from
}

type Dynamics_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Placement EnumAboveBelow

	Underline int

	Overline int

	LineThrough int

	Enclosure EnumEnclosureShape

	Id string
}

func (from *Dynamics) GongCopyBasicFields(to *Dynamics) {
	// insertion point
	to.Name = from.Name
	to.DefaultX = from.DefaultX
	to.DefaultY = from.DefaultY
	to.RelativeX = from.RelativeX
	to.RelativeY = from.RelativeY
	to.FontFamily = from.FontFamily
	to.FontStyle = from.FontStyle
	to.FontSize = from.FontSize
	to.FontWeight = from.FontWeight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Placement = from.Placement
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.LineThrough = from.LineThrough
	to.Enclosure = from.Enclosure
	to.Id = from.Id
}

type Empty_WOP struct {
	// insertion point

	Name string
}

func (from *Empty) GongCopyBasicFields(to *Empty) {
	// insertion point
	*to = *from
}

type EmptyPlacement_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow
}

func (from *EmptyPlacement) GongCopyBasicFields(to *EmptyPlacement) {
	// insertion point
	*to = *from
}

type EmptyPlacementSmufl_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow

	Smufl string
}

func (from *EmptyPlacementSmufl) GongCopyBasicFields(to *EmptyPlacementSmufl) {
	// insertion point
	*to = *from
}

type EmptyPrintObjectStyleAlign_WOP struct {
	// insertion point

	Name string

	PrintObject EnumYesNo

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign
}

func (from *EmptyPrintObjectStyleAlign) GongCopyBasicFields(to *EmptyPrintObjectStyleAlign) {
	// insertion point
	*to = *from
}

type EmptyPrintStyle_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string
}

func (from *EmptyPrintStyle) GongCopyBasicFields(to *EmptyPrintStyle) {
	// insertion point
	*to = *from
}

type EmptyPrintStyleAlign_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign
}

func (from *EmptyPrintStyleAlign) GongCopyBasicFields(to *EmptyPrintStyleAlign) {
	// insertion point
	*to = *from
}

type EmptyPrintStyleAlignId_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Id string
}

func (from *EmptyPrintStyleAlignId) GongCopyBasicFields(to *EmptyPrintStyleAlignId) {
	// insertion point
	*to = *from
}

type EmptyTrillSound_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow

	StartNote EnumStartNote

	TrillStep EnumTrillStep

	TwoNoteTurn EnumTwoNoteTurn

	Accelerate EnumYesNo

	Beats string

	SecondBeat string

	LastBeat string
}

func (from *EmptyTrillSound) GongCopyBasicFields(to *EmptyTrillSound) {
	// insertion point
	*to = *from
}

type Encoding_WOP struct {
	// insertion point

	Name string

	EncodingDate string

	Software string

	EncodingDescription string
}

func (from *Encoding) GongCopyBasicFields(to *Encoding) {
	// insertion point
	to.Name = from.Name
	to.EncodingDate = from.EncodingDate
	to.Software = from.Software
	to.EncodingDescription = from.EncodingDescription
}

type Fermata_WOP struct {
	// insertion point

	Name string

	Type EnumUprightInverted

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Id string

	Value EnumFermataShape
}

func (from *Fermata) GongCopyBasicFields(to *Fermata) {
	// insertion point
	*to = *from
}

type Fingering_WOP struct {
	// insertion point

	Name string

	Substitution EnumYesNo

	Alternate EnumYesNo

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow

	Value string
}

func (from *Fingering) GongCopyBasicFields(to *Fingering) {
	// insertion point
	*to = *from
}

type FormattedSymbol_WOP struct {
	// insertion point

	Name string

	Justify EnumLeftCenterRight

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	LineHeight string

	Dir EnumTextDirection

	Enclosure EnumEnclosureShape

	Value string
}

func (from *FormattedSymbol) GongCopyBasicFields(to *FormattedSymbol) {
	// insertion point
	*to = *from
}

type FormattedSymbolId_WOP struct {
	// insertion point

	Name string

	Justify EnumLeftCenterRight

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	LineHeight string

	Dir EnumTextDirection

	Enclosure EnumEnclosureShape

	Id string

	Value string
}

func (from *FormattedSymbolId) GongCopyBasicFields(to *FormattedSymbolId) {
	// insertion point
	*to = *from
}

type FormattedText_WOP struct {
	// insertion point

	Name string

	Lang string

	Space string

	Justify EnumLeftCenterRight

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	LineHeight string

	Dir EnumTextDirection

	Enclosure EnumEnclosureShape

	Value string
}

func (from *FormattedText) GongCopyBasicFields(to *FormattedText) {
	// insertion point
	*to = *from
}

type FormattedTextId_WOP struct {
	// insertion point

	Name string

	Lang string

	Space string

	Justify EnumLeftCenterRight

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Underline int

	Overline int

	LineThrough int

	Rotation string

	LetterSpacing string

	LineHeight string

	Dir EnumTextDirection

	Enclosure EnumEnclosureShape

	Id string

	Value string
}

func (from *FormattedTextId) GongCopyBasicFields(to *FormattedTextId) {
	// insertion point
	*to = *from
}

type Fret_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Value int
}

func (from *Fret) GongCopyBasicFields(to *Fret) {
	// insertion point
	*to = *from
}

type Glyph_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *Glyph) GongCopyBasicFields(to *Glyph) {
	// insertion point
	*to = *from
}

type HorizontalTurn_WOP struct {
	// insertion point

	Name string

	Slash EnumYesNo

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow

	StartNote EnumStartNote

	TrillStep EnumTrillStep

	TwoNoteTurn EnumTwoNoteTurn

	Accelerate EnumYesNo

	Beats string

	SecondBeat string

	LastBeat string
}

func (from *HorizontalTurn) GongCopyBasicFields(to *HorizontalTurn) {
	// insertion point
	*to = *from
}

type Identification_WOP struct {
	// insertion point

	Name string

	Source string
}

func (from *Identification) GongCopyBasicFields(to *Identification) {
	// insertion point
	to.Name = from.Name
	to.Source = from.Source
}

type Level_WOP struct {
	// insertion point

	Name string

	Reference EnumYesNo

	Type EnumStartStopSingle

	Parentheses EnumYesNo

	Bracket EnumYesNo

	Size EnumSymbolSize

	Value string
}

func (from *Level) GongCopyBasicFields(to *Level) {
	// insertion point
	*to = *from
}

type LineWidth_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *LineWidth) GongCopyBasicFields(to *LineWidth) {
	// insertion point
	*to = *from
}

type Link_WOP struct {
	// insertion point

	Name string

	NameAttr string

	Href string

	Role string

	Title string

	Element string

	Position int

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string
}

func (from *Link) GongCopyBasicFields(to *Link) {
	// insertion point
	*to = *from
}

type MeasureLayout_WOP struct {
	// insertion point

	Name string

	MeasureDistance string
}

func (from *MeasureLayout) GongCopyBasicFields(to *MeasureLayout) {
	// insertion point
	*to = *from
}

type MidiDevice_WOP struct {
	// insertion point

	Name string

	Port int

	Id string

	Value string
}

func (from *MidiDevice) GongCopyBasicFields(to *MidiDevice) {
	// insertion point
	*to = *from
}

type MidiInstrument_WOP struct {
	// insertion point

	Name string

	Id string

	MidiChannel int

	MidiName string

	MidiBank int

	MidiProgram int

	MidiUnpitched int

	Volume string

	Pan string

	Elevation string
}

func (from *MidiInstrument) GongCopyBasicFields(to *MidiInstrument) {
	// insertion point
	*to = *from
}

type Miscellaneous_WOP struct {
	// insertion point

	Name string
}

func (from *Miscellaneous) GongCopyBasicFields(to *Miscellaneous) {
	// insertion point
	to.Name = from.Name
}

type MiscellaneousField_WOP struct {
	// insertion point

	Name string

	NameAttr string

	Value string
}

func (from *MiscellaneousField) GongCopyBasicFields(to *MiscellaneousField) {
	// insertion point
	*to = *from
}

type NameDisplay_WOP struct {
	// insertion point

	Name string

	PrintObject EnumYesNo
}

func (from *NameDisplay) GongCopyBasicFields(to *NameDisplay) {
	// insertion point
	to.Name = from.Name
	to.PrintObject = from.PrintObject
}

type NoteSize_WOP struct {
	// insertion point

	Name string

	Type EnumNoteSizeType

	Value string
}

func (from *NoteSize) GongCopyBasicFields(to *NoteSize) {
	// insertion point
	*to = *from
}

type OtherAppearance_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *OtherAppearance) GongCopyBasicFields(to *OtherAppearance) {
	// insertion point
	*to = *from
}

type OtherPlay_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *OtherPlay) GongCopyBasicFields(to *OtherPlay) {
	// insertion point
	*to = *from
}

type OtherText_WOP struct {
	// insertion point

	Name string

	Smufl string

	Value string
}

func (from *OtherText) GongCopyBasicFields(to *OtherText) {
	// insertion point
	*to = *from
}

type PageLayout_WOP struct {
	// insertion point

	Name string

	PageHeight string

	PageWidth string
}

func (from *PageLayout) GongCopyBasicFields(to *PageLayout) {
	// insertion point
	to.Name = from.Name
	to.PageHeight = from.PageHeight
	to.PageWidth = from.PageWidth
}

type PageMargins_WOP struct {
	// insertion point

	Name string

	Type EnumMarginType

	LeftMargin string

	RightMargin string

	TopMargin string

	BottomMargin string
}

func (from *PageMargins) GongCopyBasicFields(to *PageMargins) {
	// insertion point
	*to = *from
}

type Play_WOP struct {
	// insertion point

	Name string

	Id string

	Ipa string

	Mute EnumMute

	SemiPitched EnumSemiPitched
}

func (from *Play) GongCopyBasicFields(to *Play) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Ipa = from.Ipa
	to.Mute = from.Mute
	to.SemiPitched = from.SemiPitched
}

type Scaling_WOP struct {
	// insertion point

	Name string

	Millimeters string

	Tenths string
}

func (from *Scaling) GongCopyBasicFields(to *Scaling) {
	// insertion point
	*to = *from
}

type Segno_WOP struct {
	// insertion point

	Name string

	Smufl string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Halign EnumLeftCenterRight

	Valign EnumValign

	Id string
}

func (from *Segno) GongCopyBasicFields(to *Segno) {
	// insertion point
	*to = *from
}

type StaffLayout_WOP struct {
	// insertion point

	Name string

	Number int

	StaffDistance string
}

func (from *StaffLayout) GongCopyBasicFields(to *StaffLayout) {
	// insertion point
	*to = *from
}

type StringType_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontStyle EnumFontStyle

	FontSize string

	FontWeight EnumFontWeight

	Color string

	Placement EnumAboveBelow

	Value int
}

func (from *StringType) GongCopyBasicFields(to *StringType) {
	// insertion point
	*to = *from
}

type Supports_WOP struct {
	// insertion point

	Name string

	Type EnumYesNo

	Element string

	Attribute string

	Value string
}

func (from *Supports) GongCopyBasicFields(to *Supports) {
	// insertion point
	*to = *from
}

type SystemDividers_WOP struct {
	// insertion point

	Name string
}

func (from *SystemDividers) GongCopyBasicFields(to *SystemDividers) {
	// insertion point
	to.Name = from.Name
}

type SystemLayout_WOP struct {
	// insertion point

	Name string

	SystemDistance string

	TopSystemDistance string
}

func (from *SystemLayout) GongCopyBasicFields(to *SystemLayout) {
	// insertion point
	to.Name = from.Name
	to.SystemDistance = from.SystemDistance
	to.TopSystemDistance = from.TopSystemDistance
}

type SystemMargins_WOP struct {
	// insertion point

	Name string

	LeftMargin string

	RightMargin string
}

func (from *SystemMargins) GongCopyBasicFields(to *SystemMargins) {
	// insertion point
	*to = *from
}

type TypedText_WOP struct {
	// insertion point

	Name string

	Type string

	Value string
}

func (from *TypedText) GongCopyBasicFields(to *TypedText) {
	// insertion point
	*to = *from
}

type WavyLine_WOP struct {
	// insertion point

	Name string

	Type EnumStartStopContinue

	Number int

	Smufl string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Placement EnumAboveBelow

	Color string

	StartNote EnumStartNote

	TrillStep EnumTrillStep

	TwoNoteTurn EnumTwoNoteTurn

	Accelerate EnumYesNo

	Beats string

	SecondBeat string

	LastBeat string
}

func (from *WavyLine) GongCopyBasicFields(to *WavyLine) {
	// insertion point
	*to = *from
}

// end of insertion point
