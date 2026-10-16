// generated code by xsd2gong - do not edit
package attributes

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
)

var _ xml.Attr

// EnumCancelLocation: The cancel-location type is used to indicate where a key signature
// cancellation appears relative to a new key signature: to the left, to the right, or
// before the barline and to the left. It is left by default. For mid-measure key
// elements, a cancel-location of before-barline should be treated like a
// cancel-location of left.
type EnumCancelLocation string

const (
	EnumCancelLocationLeft EnumCancelLocation = "left"
	EnumCancelLocationRight EnumCancelLocation = "right"
	EnumCancelLocationBeforeBarline EnumCancelLocation = "before-barline"
)

// EnumClefSign: The clef-sign type represents the different clef symbols. The jianpu
// sign indicates that the music that follows should be in jianpu numbered notation,
// just as the TAB sign indicates that the music that follows should be in tablature
// notation. Unlike TAB, a jianpu sign does not correspond to a visual clef notation.
// The none sign is deprecated as of MusicXML 4.0. Use the clef element's print-object
// attribute instead. When the none sign is used, notes should be displayed as if in
// treble clef.
type EnumClefSign string

const (
	EnumClefSignG EnumClefSign = "G"
	EnumClefSignF EnumClefSign = "F"
	EnumClefSignC EnumClefSign = "C"
	EnumClefSignPercussion EnumClefSign = "percussion"
	EnumClefSignTAB EnumClefSign = "TAB"
	EnumClefSignJianpu EnumClefSign = "jianpu"
	EnumClefSignNone EnumClefSign = "none"
)

// EnumShowFrets: The show-frets type indicates whether to show tablature frets as
// numbers (0, 1, 2) or letters (a, b, c). The default choice is numbers.
type EnumShowFrets string

const (
	EnumShowFretsNumbers EnumShowFrets = "numbers"
	EnumShowFretsLetters EnumShowFrets = "letters"
)

// EnumStaffType: The staff-type value can be ossia, editorial, cue, alternate, or
// regular. An ossia staff represents music that can be played instead of what appears
// on the regular staff. An editorial staff also represents musical alternatives, but
// is created by an editor rather than the composer. It can be used for suggested
// interpretations or alternatives from other sources. A cue staff represents music
// from another part. An alternate staff shares the same music as the prior staff, but
// displayed differently (e.g., treble and bass clef, standard notation and tablature).
// It is not included in playback. An alternate staff provides more information to an
// application reading a file than encoding the same music in separate parts, so its
// use is preferred in this situation if feasible. A regular staff is the standard
// default staff-type.
type EnumStaffType string

const (
	EnumStaffTypeOssia EnumStaffType = "ossia"
	EnumStaffTypeEditorial EnumStaffType = "editorial"
	EnumStaffTypeCue EnumStaffType = "cue"
	EnumStaffTypeAlternate EnumStaffType = "alternate"
	EnumStaffTypeRegular EnumStaffType = "regular"
)

// EnumTimeRelation: The time-relation type indicates the symbol used to represent the
// interchangeable aspect of dual time signatures.
type EnumTimeRelation string

const (
	EnumTimeRelationParentheses EnumTimeRelation = "parentheses"
	EnumTimeRelationBracket EnumTimeRelation = "bracket"
	EnumTimeRelationEquals EnumTimeRelation = "equals"
	EnumTimeRelationSlash EnumTimeRelation = "slash"
	EnumTimeRelationSpace EnumTimeRelation = "space"
	EnumTimeRelationHyphen EnumTimeRelation = "hyphen"
)

// EnumTimeSeparator: The time-separator type indicates how to display the arrangement
// between the beats and beat-type values in a time signature. The default value is
// none. The horizontal, diagonal, and vertical values represent horizontal, diagonal
// lower-left to upper-right, and vertical lines respectively. For these values, the
// beats and beat-type values are arranged on either side of the separator line. The
// none value represents no separator with the beats and beat-type arranged vertically.
// The adjacent value represents no separator with the beats and beat-type arranged
// horizontally.
type EnumTimeSeparator string

const (
	EnumTimeSeparatorNone EnumTimeSeparator = "none"
	EnumTimeSeparatorHorizontal EnumTimeSeparator = "horizontal"
	EnumTimeSeparatorDiagonal EnumTimeSeparator = "diagonal"
	EnumTimeSeparatorVertical EnumTimeSeparator = "vertical"
	EnumTimeSeparatorAdjacent EnumTimeSeparator = "adjacent"
)

// EnumTimeSymbol: The time-symbol type indicates how to display a time signature. The
// normal value is the usual fractional display, and is the implied symbol type if none
// is specified. Other options are the common and cut time symbols, as well as a single
// number with an implied denominator. The note symbol indicates that the beat-type
// should be represented with the corresponding downstem note rather than a number. The
// dotted-note symbol indicates that the beat-type should be represented with a dotted
// downstem note that corresponds to three times the beat-type value, and a numerator
// that is one third the beats value.
type EnumTimeSymbol string

const (
	EnumTimeSymbolCommon EnumTimeSymbol = "common"
	EnumTimeSymbolCut EnumTimeSymbol = "cut"
	EnumTimeSymbolSingleNumber EnumTimeSymbol = "single-number"
	EnumTimeSymbolNote EnumTimeSymbol = "note"
	EnumTimeSymbolDottedNote EnumTimeSymbol = "dotted-note"
	EnumTimeSymbolNormal EnumTimeSymbol = "normal"
)

// EnumStep: The step type represents a step of the diatonic scale, represented
// using the English letters A through G.
type EnumStep string

const (
	EnumStepA EnumStep = "A"
	EnumStepB EnumStep = "B"
	EnumStepC EnumStep = "C"
	EnumStepD EnumStep = "D"
	EnumStepE EnumStep = "E"
	EnumStepF EnumStep = "F"
	EnumStepG EnumStep = "G"
)

// EnumGroupSymbolValue: The group-symbol-value type indicates how the symbol for a group or
// multi-staff part is indicated in the score.
type EnumGroupSymbolValue string

const (
	EnumGroupSymbolValueNone EnumGroupSymbolValue = "none"
	EnumGroupSymbolValueBrace EnumGroupSymbolValue = "brace"
	EnumGroupSymbolValueLine EnumGroupSymbolValue = "line"
	EnumGroupSymbolValueBracket EnumGroupSymbolValue = "bracket"
	EnumGroupSymbolValueSquare EnumGroupSymbolValue = "square"
)

// Attributes: The attributes element contains musical information that typically
// changes on measure boundaries. This includes key and time signatures, clefs,
// transpositions, and staving. When attributes are changed mid-measure, it affects the
// music in score order, not in MusicXML document order.
type Attributes struct {
	Name string `xml:"-"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Divisions string `xml:"divisions,omitempty"`
	Key []*Key `xml:"key,omitempty"`
	Time []*Time `xml:"time,omitempty"`
	Staves int `xml:"staves,omitempty"`
	PartSymbol *PartSymbol `xml:"part-symbol,omitempty"`
	Instruments int `xml:"instruments,omitempty"`
	Clef []*Clef `xml:"clef,omitempty"`
	StaffDetails []*StaffDetails `xml:"staff-details,omitempty"`
	Transpose []*Transpose `xml:"transpose,omitempty"`
	ForPart []*ForPart `xml:"for-part,omitempty"`
	Directive []*AttributesDirective `xml:"directive,omitempty"`
	MeasureStyle []*MeasureStyle `xml:"measure-style,omitempty"`
}

// BeatRepeat: The beat-repeat type is used to indicate that a single beat (but
// possibly many notes) is repeated. The slashes attribute specifies the number of
// slashes to use in the symbol. The use-dots attribute indicates whether or not to use
// dots as well (for instance, with mixed rhythm patterns). The value for slashes is 1
// and the value for use-dots is no if not specified. The stop type indicates the first
// beat where the repeats are no longer displayed. Both the start and stop of the beat
// being repeated should be specified unless the repeats are displayed through the end
// of the part. The beat-repeat element specifies a notation style for repetitions. The
// actual music being repeated needs to be repeated within the MusicXML file. This
// element specifies the notation that indicates the repeat.
type BeatRepeat struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Slashes int `xml:"slashes,attr,omitempty"`
	UseDots common.EnumYesNo `xml:"use-dots,attr,omitempty"`
	SlashType common.EnumNoteTypeValue `xml:"slash-type,omitempty"`
	SlashDot []*common.Empty `xml:"slash-dot,omitempty"`
	ExceptVoice string `xml:"except-voice,omitempty"`
}

// Cancel: A cancel element indicates that the old key signature should be
// cancelled before the new one appears. This will always happen when changing to C
// major or A minor and need not be specified then. The cancel value matches the fifths
// value of the cancelled key signature (e.g., a cancel of -2 will provide an explicit
// cancellation for changing from B flat major to F major). The optional location
// attribute indicates where the cancellation appears relative to the new key
// signature.
type Cancel struct {
	Name string `xml:"-"`
	Location EnumCancelLocation `xml:"location,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Clef: Clefs are represented by a combination of sign, line, and
// clef-octave-change elements. The optional number attribute refers to staff numbers
// within the part. A value of 1 is assumed if not present. Sometimes clefs are added
// to the staff in non-standard line positions, either to indicate cue passages, or
// when there are multiple clefs present simultaneously on one staff. In this
// situation, the additional attribute is set to "yes" and the line value is ignored.
// The size attribute is used for clefs where the additional attribute is "yes". It is
// typically used to indicate cue clefs. Sometimes clefs at the start of a measure need
// to appear after the barline rather than before, as for cues or for use after a
// repeated section. The after-barline attribute is set to "yes" in this situation. The
// attribute is ignored for mid-measure clefs. Clefs appear at the start of each system
// unless the print-object attribute has been set to "no" or the additional attribute
// has been set to "yes".
type Clef struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Additional common.EnumYesNo `xml:"additional,attr,omitempty"`
	Size common.EnumSymbolSize `xml:"size,attr,omitempty"`
	AfterBarline common.EnumYesNo `xml:"after-barline,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Sign EnumClefSign `xml:"sign,omitempty"`
	Line int `xml:"line,omitempty"`
	ClefOctaveChange int `xml:"clef-octave-change,omitempty"`
}

// Double: The double type indicates that the music is doubled one octave from
// what is currently written. If the above attribute is set to yes, the doubling is one
// octave above what is written, as for mixed flute / piccolo parts in band literature.
// Otherwise the doubling is one octave below what is written, as for mixed cello /
// bass parts in orchestral literature.
type Double struct {
	Name string `xml:"-"`
	Above common.EnumYesNo `xml:"above,attr,omitempty"`
}

// ForPart: The for-part type is used in a concert score to indicate the
// transposition for a transposed part created from that score. It is only used in
// score files that contain a concert-score element in the defaults. This allows
// concert scores with transposed parts to be represented in a single uncompressed
// MusicXML file. The optional number attribute refers to staff numbers, from top to
// bottom on the system. If absent, the child elements apply to all staves in the
// created part.
type ForPart struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	PartClef *PartClef `xml:"part-clef,omitempty"`
	PartTranspose *PartTranspose `xml:"part-transpose,omitempty"`
}

// Interchangeable: The interchangeable type is used to represent the second in a pair of
// interchangeable dual time signatures, such as the 6/8 in 3/4 (6/8). A separate
// symbol attribute value is available compared to the time element's symbol attribute,
// which applies to the first of the dual time signatures.
type Interchangeable struct {
	Name string `xml:"-"`
	Symbol EnumTimeSymbol `xml:"symbol,attr,omitempty"`
	Separator EnumTimeSeparator `xml:"separator,attr,omitempty"`
	TimeRelation EnumTimeRelation `xml:"time-relation,omitempty"`
	Beats string `xml:"beats,omitempty"`
	BeatType string `xml:"beat-type,omitempty"`
}

// Key: The key type represents a key signature. Both traditional and
// non-traditional key signatures are supported. The optional number attribute refers
// to staff numbers. If absent, the key signature applies to all staves in the part.
// Key signatures appear at the start of each system unless the print-object attribute
// has been set to "no".
type Key struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Cancel *Cancel `xml:"cancel,omitempty"`
	Fifths int `xml:"fifths,omitempty"`
	Mode string `xml:"mode,omitempty"`
	KeyStep EnumStep `xml:"key-step,omitempty"`
	KeyAlter string `xml:"key-alter,omitempty"`
	KeyAccidental []*KeyAccidental `xml:"key-accidental,omitempty"`
	KeyOctave []*KeyOctave `xml:"key-octave,omitempty"`
}

// KeyAccidental: The key-accidental type indicates the accidental to be displayed in a
// non-traditional key signature, represented in the same manner as the accidental type
// without the formatting attributes.
type KeyAccidental struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value common.EnumAccidentalValue `xml:",chardata"`
}

// KeyOctave: The key-octave type specifies in which octave an element of a key
// signature appears. The content specifies the octave value using the same values as
// the display-octave element. The number attribute is a positive integer that refers
// to the key signature element in left-to-right order. If the cancel attribute is set
// to yes, then this number refers to the canceling key signature specified by the
// cancel element in the parent key element. The cancel attribute cannot be set to yes
// if there is no corresponding cancel element within the parent key element. It is no
// by default.
type KeyOctave struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Cancel common.EnumYesNo `xml:"cancel,attr,omitempty"`
	Value int `xml:",chardata"`
}

// LineDetail: If the staff-lines element is present, the appearance of each line may
// be individually specified with a line-detail type. Staff lines are numbered from
// bottom to top. The print-object attribute allows lines to be hidden within a staff.
// This is used in special situations such as a widely-spaced percussion staff where a
// note placed below the higher line is distinct from a note placed above the lower
// line. Hidden staff lines are included when specifying clef lines and determining
// display-step / display-octave values, but are not counted as lines for the purposes
// of the system-layout and staff-layout elements.
type LineDetail struct {
	Name string `xml:"-"`
	Line int `xml:"line,attr,omitempty"`
	Width string `xml:"width,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
}

// MeasureRepeat: The measure-repeat type is used for both single and multiple measure
// repeats. The text of the element indicates the number of measures to be repeated in
// a single pattern. The slashes attribute specifies the number of slashes to use in
// the repeat sign. It is 1 if not specified. The text of the element is ignored when
// the type is stop. The stop type indicates the first measure where the repeats are no
// longer displayed. Both the start and the stop of the measure-repeat should be
// specified unless the repeats are displayed through the end of the part. The
// measure-repeat element specifies a notation style for repetitions. The actual music
// being repeated needs to be repeated within each measure of the MusicXML file. This
// element specifies the notation that indicates the repeat.
type MeasureRepeat struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Slashes int `xml:"slashes,attr,omitempty"`
	Value string `xml:",chardata"`
}

// MeasureStyle: A measure-style indicates a special way to print partial to multiple
// measures within a part. This includes multiple rests over several measures, repeats
// of beats, single, or multiple measures, and use of slash notation. The multiple-rest
// and measure-repeat elements indicate the number of measures covered in the element
// content. The beat-repeat and slash elements can cover partial measures. All but the
// multiple-rest element use a type attribute to indicate starting and stopping the use
// of the style. The optional number attribute specifies the staff number from top to
// bottom on the system, as with clef.
type MeasureStyle struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	MultipleRest *MultipleRest `xml:"multiple-rest,omitempty"`
	MeasureRepeat *MeasureRepeat `xml:"measure-repeat,omitempty"`
	BeatRepeat *BeatRepeat `xml:"beat-repeat,omitempty"`
	Slash *Slash `xml:"slash,omitempty"`
}

// MultipleRest: The text of the multiple-rest type indicates the number of measures in
// the multiple rest. Multiple rests may use the 1-bar / 2-bar / 4-bar rest symbols, or
// a single shape. The use-symbols attribute indicates which to use; it is no if not
// specified.
type MultipleRest struct {
	Name string `xml:"-"`
	UseSymbols common.EnumYesNo `xml:"use-symbols,attr,omitempty"`
	Value int `xml:",chardata"`
}

// PartClef: The child elements of the part-clef type have the same meaning as for
// the clef type. However that meaning applies to a transposed part created from the
// existing score file.
type PartClef struct {
	Name string `xml:"-"`
	Sign EnumClefSign `xml:"sign,omitempty"`
	Line int `xml:"line,omitempty"`
	ClefOctaveChange int `xml:"clef-octave-change,omitempty"`
}

// PartSymbol: The part-symbol type indicates how a symbol for a multi-staff part is
// indicated in the score; brace is the default value. The top-staff and bottom-staff
// attributes are used when the brace does not extend across the entire part. For
// example, in a 3-staff organ part, the top-staff will typically be 1 for the right
// hand, while the bottom-staff will typically be 2 for the left hand. Staff 3 for the
// pedals is usually outside the brace. By default, the presence of a part-symbol
// element that does not extend across the entire part also indicates a corresponding
// change in the common barlines within a part.
type PartSymbol struct {
	Name string `xml:"-"`
	TopStaff int `xml:"top-staff,attr,omitempty"`
	BottomStaff int `xml:"bottom-staff,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value EnumGroupSymbolValue `xml:",chardata"`
}

// PartTranspose: The child elements of the part-transpose type have the same meaning as
// for the transpose type. However that meaning applies to a transposed part created
// from the existing score file.
type PartTranspose struct {
	Name string `xml:"-"`
	Diatonic int `xml:"diatonic,omitempty"`
	Chromatic string `xml:"chromatic,omitempty"`
	OctaveChange int `xml:"octave-change,omitempty"`
	Double *Double `xml:"double,omitempty"`
}

// Slash: The slash type is used to indicate that slash notation is to be used.
// If the slash is on every beat, use-stems is no (the default). To indicate rhythms
// but not pitches, use-stems is set to yes. The type attribute indicates whether this
// is the start or stop of a slash notation style. The use-dots attribute works as for
// the beat-repeat element, and only has effect if use-stems is no.
type Slash struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	UseDots common.EnumYesNo `xml:"use-dots,attr,omitempty"`
	UseStems common.EnumYesNo `xml:"use-stems,attr,omitempty"`
	SlashType common.EnumNoteTypeValue `xml:"slash-type,omitempty"`
	SlashDot []*common.Empty `xml:"slash-dot,omitempty"`
	ExceptVoice string `xml:"except-voice,omitempty"`
}

// StaffDetails: The staff-details element is used to indicate different types of
// staves. The optional number attribute specifies the staff number from top to bottom
// on the system, as with clef. The print-object attribute is used to indicate when a
// staff is not printed in a part, usually in large scores where empty parts are
// omitted. It is yes by default. If print-spacing is yes while print-object is no, the
// score is printed in cutaway format where vertical space is left for the empty part.
type StaffDetails struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	ShowFrets EnumShowFrets `xml:"show-frets,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	PrintSpacing common.EnumYesNo `xml:"print-spacing,attr,omitempty"`
	StaffType EnumStaffType `xml:"staff-type,omitempty"`
	StaffLines int `xml:"staff-lines,omitempty"`
	LineDetail []*LineDetail `xml:"line-detail,omitempty"`
	StaffTuning []*StaffTuning `xml:"staff-tuning,omitempty"`
	Capo int `xml:"capo,omitempty"`
	StaffSize *StaffSize `xml:"staff-size,omitempty"`
}

// StaffSize: The staff-size element indicates how large a staff space is on this
// staff, expressed as a percentage of the work's default scaling. Values less than 100
// make the staff space smaller while values over 100 make the staff space larger. A
// staff-type of cue, ossia, or editorial implies a staff-size of less than 100, but
// the exact value is implementation-dependent unless specified here. Staff size
// affects staff height only, not the relationship of the staff to the left and right
// margins. In some cases, a staff-size different than 100 also scales the notation on
// the staff, such as with a cue staff. In other cases, such as percussion staves, the
// lines may be more widely spaced without scaling the notation on the staff. The
// scaling attribute allows these two cases to be distinguished. It specifies the
// percentage scaling that applies to the notation. Values less that 100 make the
// notation smaller while values over 100 make the notation larger. The staff-size
// content and scaling attribute are both non-negative decimal values.
type StaffSize struct {
	Name string `xml:"-"`
	Scaling string `xml:"scaling,attr,omitempty"`
	Value string `xml:",chardata"`
}

// StaffTuning: The staff-tuning type specifies the open, non-capo tuning of the lines
// on a tablature staff.
type StaffTuning struct {
	Name string `xml:"-"`
	Line int `xml:"line,attr,omitempty"`
	TuningStep EnumStep `xml:"tuning-step,omitempty"`
	TuningAlter string `xml:"tuning-alter,omitempty"`
	TuningOctave int `xml:"tuning-octave,omitempty"`
}

// Time: Time signatures are represented by the beats element for the numerator
// and the beat-type element for the denominator. The symbol attribute is used to
// indicate common and cut time symbols as well as a single number display. Multiple
// pairs of beat and beat-type elements are used for composite time signatures with
// multiple denominators, such as 2/4 + 3/8. A composite such as 3+2/8 requires only
// one beat/beat-type pair. The print-object attribute allows a time signature to be
// specified but not printed, as is the case for excerpts from the middle of a score.
// The value is "yes" if not present. The optional number attribute refers to staff
// numbers within the part. If absent, the time signature applies to all staves in the
// part.
type Time struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Symbol EnumTimeSymbol `xml:"symbol,attr,omitempty"`
	Separator EnumTimeSeparator `xml:"separator,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Beats string `xml:"beats,omitempty"`
	BeatType string `xml:"beat-type,omitempty"`
	Interchangeable *Interchangeable `xml:"interchangeable,omitempty"`
	SenzaMisura string `xml:"senza-misura,omitempty"`
}

// Transpose: The transpose type represents what must be added to a written pitch to
// get a correct sounding pitch. The optional number attribute refers to staff numbers,
// from top to bottom on the system. If absent, the transposition applies to all staves
// in the part. Per-staff transposition is most often used in parts that represent
// multiple instruments.
type Transpose struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Diatonic int `xml:"diatonic,omitempty"`
	Chromatic string `xml:"chromatic,omitempty"`
	OctaveChange int `xml:"octave-change,omitempty"`
	Double *Double `xml:"double,omitempty"`
}

// AttributesDirective: Directives are like directions, but can be grouped together
// with attributes for convenience. This is typically used for tempo markings
// at the beginning of a piece of music. This element was deprecated in Version
// 2.0 in favor of the direction element's directive attribute. Language names
// come from ISO 639, with optional country subcodes from ISO 3166.
type AttributesDirective struct {
	Name string `xml:"-"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

