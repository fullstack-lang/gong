// generated code by xsd2gong - do not edit
package note

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
)

var _ xml.Attr

// EnumArrowDirection: The arrow-direction type represents the direction in which an arrow
// points, using Unicode arrow terminology.
type EnumArrowDirection string

const (
	EnumArrowDirectionLeft EnumArrowDirection = "left"
	EnumArrowDirectionUp EnumArrowDirection = "up"
	EnumArrowDirectionRight EnumArrowDirection = "right"
	EnumArrowDirectionDown EnumArrowDirection = "down"
	EnumArrowDirectionNorthwest EnumArrowDirection = "northwest"
	EnumArrowDirectionNortheast EnumArrowDirection = "northeast"
	EnumArrowDirectionSoutheast EnumArrowDirection = "southeast"
	EnumArrowDirectionSouthwest EnumArrowDirection = "southwest"
	EnumArrowDirectionLeftRight EnumArrowDirection = "left right"
	EnumArrowDirectionUpDown EnumArrowDirection = "up down"
	EnumArrowDirectionNorthwestSoutheast EnumArrowDirection = "northwest southeast"
	EnumArrowDirectionNortheastSouthwest EnumArrowDirection = "northeast southwest"
	EnumArrowDirectionOther EnumArrowDirection = "other"
)

// EnumArrowStyle: The arrow-style type represents the style of an arrow, using Unicode
// arrow terminology. Filled and hollow arrows indicate polygonal single arrows. Paired
// arrows are duplicate single arrows in the same direction. Combined arrows apply to
// double direction arrows like left right, indicating that an arrow in one direction
// should be combined with an arrow in the other direction.
type EnumArrowStyle string

const (
	EnumArrowStyleSingle EnumArrowStyle = "single"
	EnumArrowStyleDouble EnumArrowStyle = "double"
	EnumArrowStyleFilled EnumArrowStyle = "filled"
	EnumArrowStyleHollow EnumArrowStyle = "hollow"
	EnumArrowStylePaired EnumArrowStyle = "paired"
	EnumArrowStyleCombined EnumArrowStyle = "combined"
	EnumArrowStyleOther EnumArrowStyle = "other"
)

// EnumBendShape: The bend-shape type distinguishes between the angled bend symbols
// commonly used in standard notation and the curved bend symbols commonly used in both
// tablature and standard notation.
type EnumBendShape string

const (
	EnumBendShapeAngled EnumBendShape = "angled"
	EnumBendShapeCurved EnumBendShape = "curved"
)

// EnumBreathMarkValue: The breath-mark-value type represents the symbol used for a breath
// mark.
type EnumBreathMarkValue string

const (
	EnumBreathMarkValueEmpty EnumBreathMarkValue = ""
	EnumBreathMarkValueComma EnumBreathMarkValue = "comma"
	EnumBreathMarkValueTick EnumBreathMarkValue = "tick"
	EnumBreathMarkValueUpbow EnumBreathMarkValue = "upbow"
	EnumBreathMarkValueSalzedo EnumBreathMarkValue = "salzedo"
)

// EnumCaesuraValue: The caesura-value type represents the shape of the caesura sign.
type EnumCaesuraValue string

const (
	EnumCaesuraValueNormal EnumCaesuraValue = "normal"
	EnumCaesuraValueThick EnumCaesuraValue = "thick"
	EnumCaesuraValueShort EnumCaesuraValue = "short"
	EnumCaesuraValueCurved EnumCaesuraValue = "curved"
	EnumCaesuraValueSingle EnumCaesuraValue = "single"
	EnumCaesuraValueEmpty EnumCaesuraValue = ""
)

// EnumCircularArrow: The circular-arrow type represents the direction in which a circular
// arrow points, using Unicode arrow terminology.
type EnumCircularArrow string

const (
	EnumCircularArrowClockwise EnumCircularArrow = "clockwise"
	EnumCircularArrowAnticlockwise EnumCircularArrow = "anticlockwise"
)

// EnumFan: The fan type represents the type of beam fanning present on a note,
// used to represent accelerandos and ritardandos.
type EnumFan string

const (
	EnumFanAccel EnumFan = "accel"
	EnumFanRit EnumFan = "rit"
	EnumFanNone EnumFan = "none"
)

// EnumHandbellValue: The handbell-value type represents the type of handbell technique
// being notated.
type EnumHandbellValue string

const (
	EnumHandbellValueBelltree EnumHandbellValue = "belltree"
	EnumHandbellValueDamp EnumHandbellValue = "damp"
	EnumHandbellValueEcho EnumHandbellValue = "echo"
	EnumHandbellValueGyro EnumHandbellValue = "gyro"
	EnumHandbellValueHandMartellato EnumHandbellValue = "hand martellato"
	EnumHandbellValueMalletLift EnumHandbellValue = "mallet lift"
	EnumHandbellValueMalletTable EnumHandbellValue = "mallet table"
	EnumHandbellValueMartellato EnumHandbellValue = "martellato"
	EnumHandbellValueMartellatoLift EnumHandbellValue = "martellato lift"
	EnumHandbellValueMutedMartellato EnumHandbellValue = "muted martellato"
	EnumHandbellValuePluckLift EnumHandbellValue = "pluck lift"
	EnumHandbellValueSwing EnumHandbellValue = "swing"
)

// EnumHarmonClosedLocation: The harmon-closed-location type indicates which portion of the symbol
// is filled in when the corresponding harmon-closed-value is half.
type EnumHarmonClosedLocation string

const (
	EnumHarmonClosedLocationRight EnumHarmonClosedLocation = "right"
	EnumHarmonClosedLocationBottom EnumHarmonClosedLocation = "bottom"
	EnumHarmonClosedLocationLeft EnumHarmonClosedLocation = "left"
	EnumHarmonClosedLocationTop EnumHarmonClosedLocation = "top"
)

// EnumHarmonClosedValue: The harmon-closed-value type represents whether the harmon mute is
// closed, open, or half-open.
type EnumHarmonClosedValue string

const (
	EnumHarmonClosedValueYes EnumHarmonClosedValue = "yes"
	EnumHarmonClosedValueNo EnumHarmonClosedValue = "no"
	EnumHarmonClosedValueHalf EnumHarmonClosedValue = "half"
)

// EnumHoleClosedLocation: The hole-closed-location type indicates which portion of the hole is
// filled in when the corresponding hole-closed-value is half.
type EnumHoleClosedLocation string

const (
	EnumHoleClosedLocationRight EnumHoleClosedLocation = "right"
	EnumHoleClosedLocationBottom EnumHoleClosedLocation = "bottom"
	EnumHoleClosedLocationLeft EnumHoleClosedLocation = "left"
	EnumHoleClosedLocationTop EnumHoleClosedLocation = "top"
)

// EnumHoleClosedValue: The hole-closed-value type represents whether the hole is closed,
// open, or half-open.
type EnumHoleClosedValue string

const (
	EnumHoleClosedValueYes EnumHoleClosedValue = "yes"
	EnumHoleClosedValueNo EnumHoleClosedValue = "no"
	EnumHoleClosedValueHalf EnumHoleClosedValue = "half"
)

// EnumNoteheadValue: The notehead-value type indicates shapes other than the open and
// closed ovals associated with note durations. The values do, re, mi, fa, fa up, so,
// la, and ti correspond to Aikin's 7-shape system. The fa up shape is typically used
// with upstems; the fa shape is typically used with downstems or no stems. The arrow
// shapes differ from triangle and inverted triangle by being centered on the stem.
// Slashed and back slashed notes include both the normal notehead and a slash. The
// triangle shape has the tip of the triangle pointing up; the inverted triangle shape
// has the tip of the triangle pointing down. The left triangle shape is a right
// triangle with the hypotenuse facing up and to the left. The other notehead covers
// noteheads other than those listed here. It is usually used in combination with the
// smufl attribute to specify a particular SMuFL notehead. The smufl attribute may be
// used with any notehead value to help specify the appearance of symbols that share
// the same MusicXML semantics. Noteheads in the SMuFL Note name noteheads and Note
// name noteheads supplement ranges (U+E150–U+E1AF and U+EEE0–U+EEFF) should not use
// the smufl attribute or the "other" value, but instead use the notehead-text element.
type EnumNoteheadValue string

const (
	EnumNoteheadValueSlash EnumNoteheadValue = "slash"
	EnumNoteheadValueTriangle EnumNoteheadValue = "triangle"
	EnumNoteheadValueDiamond EnumNoteheadValue = "diamond"
	EnumNoteheadValueSquare EnumNoteheadValue = "square"
	EnumNoteheadValueCross EnumNoteheadValue = "cross"
	EnumNoteheadValueX EnumNoteheadValue = "x"
	EnumNoteheadValueCircleX EnumNoteheadValue = "circle-x"
	EnumNoteheadValueInvertedTriangle EnumNoteheadValue = "inverted triangle"
	EnumNoteheadValueArrowDown EnumNoteheadValue = "arrow down"
	EnumNoteheadValueArrowUp EnumNoteheadValue = "arrow up"
	EnumNoteheadValueCircled EnumNoteheadValue = "circled"
	EnumNoteheadValueSlashed EnumNoteheadValue = "slashed"
	EnumNoteheadValueBackSlashed EnumNoteheadValue = "back slashed"
	EnumNoteheadValueNormal EnumNoteheadValue = "normal"
	EnumNoteheadValueCluster EnumNoteheadValue = "cluster"
	EnumNoteheadValueCircleDot EnumNoteheadValue = "circle dot"
	EnumNoteheadValueLeftTriangle EnumNoteheadValue = "left triangle"
	EnumNoteheadValueRectangle EnumNoteheadValue = "rectangle"
	EnumNoteheadValueNone EnumNoteheadValue = "none"
	EnumNoteheadValueDo EnumNoteheadValue = "do"
	EnumNoteheadValueRe EnumNoteheadValue = "re"
	EnumNoteheadValueMi EnumNoteheadValue = "mi"
	EnumNoteheadValueFa EnumNoteheadValue = "fa"
	EnumNoteheadValueFaUp EnumNoteheadValue = "fa up"
	EnumNoteheadValueSo EnumNoteheadValue = "so"
	EnumNoteheadValueLa EnumNoteheadValue = "la"
	EnumNoteheadValueTi EnumNoteheadValue = "ti"
	EnumNoteheadValueOther EnumNoteheadValue = "other"
)

// EnumStemValue: The stem-value type represents the notated stem direction.
type EnumStemValue string

const (
	EnumStemValueDown EnumStemValue = "down"
	EnumStemValueUp EnumStemValue = "up"
	EnumStemValueDouble EnumStemValue = "double"
	EnumStemValueNone EnumStemValue = "none"
)

// EnumSyllabic: Lyric hyphenation is indicated by the syllabic type. The single,
// begin, end, and middle values represent single-syllable words, word-beginning
// syllables, word-ending syllables, and mid-word syllables, respectively.
type EnumSyllabic string

const (
	EnumSyllabicSingle EnumSyllabic = "single"
	EnumSyllabicBegin EnumSyllabic = "begin"
	EnumSyllabicEnd EnumSyllabic = "end"
	EnumSyllabicMiddle EnumSyllabic = "middle"
)

// EnumTapHand: The tap-hand type represents the symbol to use for a tap element. The
// left and right values refer to the SMuFL guitarLeftHandTapping and
// guitarRightHandTapping glyphs respectively.
type EnumTapHand string

const (
	EnumTapHandLeft EnumTapHand = "left"
	EnumTapHandRight EnumTapHand = "right"
)

// Accidental: The accidental type represents actual notated accidentals. Editorial
// and cautionary indications are indicated by attributes. Values for these attributes
// are "no" if not present. Specific graphic display such as parentheses, brackets, and
// size are controlled by the level-display attribute group.
type Accidental struct {
	Name string `xml:"-"`
	Cautionary common.EnumYesNo `xml:"cautionary,attr,omitempty"`
	Editorial common.EnumYesNo `xml:"editorial,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
	Bracket common.EnumYesNo `xml:"bracket,attr,omitempty"`
	Size common.EnumSymbolSize `xml:"size,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value common.EnumAccidentalValue `xml:",chardata"`
}

// AccidentalMark: An accidental-mark can be used as a separate notation or as part of an
// ornament. When used in an ornament, position and placement are relative to the
// ornament, not relative to the note.
type AccidentalMark struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
	Bracket common.EnumYesNo `xml:"bracket,attr,omitempty"`
	Size common.EnumSymbolSize `xml:"size,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value common.EnumAccidentalValue `xml:",chardata"`
}

// Arpeggiate: The arpeggiate type indicates that this note is part of an arpeggiated
// chord. The number attribute can be used to distinguish between two simultaneous
// chords arpeggiated separately (different numbers) or together (same number). The
// direction attribute is used if there is an arrow on the arpeggio sign. By default,
// arpeggios go from the lowest to highest note. The length of the sign can be
// determined from the position attributes for the arpeggiate elements used with the
// top and bottom notes of the arpeggiated chord. If the unbroken attribute is set to
// yes, it indicates that the arpeggio continues onto another staff within the part.
// This serves as a hint to applications and is not required for cross-staff arpeggios.
type Arpeggiate struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Direction common.EnumUpDown `xml:"direction,attr,omitempty"`
	Unbroken common.EnumYesNo `xml:"unbroken,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Articulations: Articulations and accents are grouped together here.
type Articulations struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Accent []*common.EmptyPlacement `xml:"accent,omitempty"`
	StrongAccent []*StrongAccent `xml:"strong-accent,omitempty"`
	Staccato []*common.EmptyPlacement `xml:"staccato,omitempty"`
	Tenuto []*common.EmptyPlacement `xml:"tenuto,omitempty"`
	DetachedLegato []*common.EmptyPlacement `xml:"detached-legato,omitempty"`
	Staccatissimo []*common.EmptyPlacement `xml:"staccatissimo,omitempty"`
	Spiccato []*common.EmptyPlacement `xml:"spiccato,omitempty"`
	Scoop []*EmptyLine `xml:"scoop,omitempty"`
	Plop []*EmptyLine `xml:"plop,omitempty"`
	Doit []*EmptyLine `xml:"doit,omitempty"`
	Falloff []*EmptyLine `xml:"falloff,omitempty"`
	BreathMark []*BreathMark `xml:"breath-mark,omitempty"`
	Caesura []*Caesura `xml:"caesura,omitempty"`
	Stress []*common.EmptyPlacement `xml:"stress,omitempty"`
	Unstress []*common.EmptyPlacement `xml:"unstress,omitempty"`
	SoftAccent []*common.EmptyPlacement `xml:"soft-accent,omitempty"`
	OtherArticulation []*OtherPlacementText `xml:"other-articulation,omitempty"`
}

// Arrow: The arrow element represents an arrow used for a musical technical
// indication. It can represent both Unicode and SMuFL arrows. The presence of an
// arrowhead element indicates that only the arrowhead is displayed, not the arrow
// stem. The smufl attribute distinguishes different SMuFL glyphs that have an arrow
// appearance such as arrowBlackUp, guitarStrumUp, or handbellsSwingUp. The specified
// glyph should match the descriptive representation.
type Arrow struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	ArrowDirection EnumArrowDirection `xml:"arrow-direction,omitempty"`
	ArrowStyle EnumArrowStyle `xml:"arrow-style,omitempty"`
	Arrowhead *common.Empty `xml:"arrowhead,omitempty"`
	CircularArrow EnumCircularArrow `xml:"circular-arrow,omitempty"`
}

// Assess: By default, an assessment application should assess all notes without
// a cue child element, and not assess any note with a cue child element. The assess
// type allows this default assessment to be overridden for individual notes. The
// optional player and time-only attributes restrict the type to apply to a single
// player or set of times through a repeated section, respectively. If missing, the
// type applies to all players or all times through the repeated section, respectively.
// The player attribute references the id attribute of a player element defined within
// the matching score-part.
type Assess struct {
	Name string `xml:"-"`
	Type common.EnumYesNo `xml:"type,attr,omitempty"`
	Player string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
}

// Backup: The backup and forward elements are required to coordinate multiple
// voices in one part, including music on multiple staves. The backup type is generally
// used to move between voices and staves. Thus the backup element does not include
// voice or staff elements. Duration values should always be positive, and should not
// cross measure boundaries or mid-measure changes in the divisions value.
type Backup struct {
	Name string `xml:"-"`
	Duration string `xml:"duration,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
}

// Beam: Beam values include begin, continue, end, forward hook, and backward
// hook. Up to eight concurrent beams are available to cover up to 1024th notes. Each
// beam in a note is represented with a separate beam element, starting with the eighth
// note beam using a number attribute of 1. Note that the beam number does not
// distinguish sets of beams that overlap, as it does for slur and other elements.
// Beaming groups are distinguished by being in different voices and/or the presence or
// absence of grace and cue elements. Beams that have a begin value can also have a fan
// attribute to indicate accelerandos and ritardandos using fanned beams. The fan
// attribute may also be used with a continue value if the fanning direction changes on
// that note. The value is "none" if not specified. The repeater attribute has been
// deprecated in MusicXML 3.0. Formerly used for tremolos, it needs to be specified
// with a "yes" value for each beam using it.
type Beam struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Repeater common.EnumYesNo `xml:"repeater,attr,omitempty"`
	Fan EnumFan `xml:"fan,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value direction.EnumBeamValue `xml:",chardata"`
}

// Bend: The bend type is used in guitar notation and tablature. A single note
// with a bend and release will contain two bend elements: the first to represent the
// bend and the second to represent the release. The shape attribute distinguishes
// between the angled bend symbols commonly used in standard notation and the curved
// bend symbols commonly used in both tablature and standard notation.
type Bend struct {
	Name string `xml:"-"`
	Shape EnumBendShape `xml:"shape,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Accelerate common.EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	FirstBeat string `xml:"first-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
	BendAlter string `xml:"bend-alter,omitempty"`
	PreBend *common.Empty `xml:"pre-bend,omitempty"`
	Release *Release `xml:"release,omitempty"`
	WithBar *PlacementText `xml:"with-bar,omitempty"`
}

// BreathMark: The breath-mark element indicates a place to take a breath.
type BreathMark struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value EnumBreathMarkValue `xml:",chardata"`
}

// Caesura: The caesura element indicates a slight pause. It is notated using a
// "railroad tracks" symbol or other variations specified in the element content.
type Caesura struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value EnumCaesuraValue `xml:",chardata"`
}

// Elision: The elision type represents an elision between lyric syllables. The
// text content specifies the symbol used to display the elision. Common values are a
// no-break space (Unicode 00A0), an underscore (Unicode 005F), or an undertie (Unicode
// 203F). If the text content is empty, the smufl attribute is used to specify the
// symbol to use. Its value is a SMuFL canonical glyph name that starts with lyrics.
// The SMuFL attribute is ignored if the elision glyph is already specified by the text
// content. If neither text content nor a smufl attribute are present, the elision
// glyph is application-specific.
type Elision struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

// EmptyLine: The empty-line type represents an empty element with line-shape,
// line-type, line-length, dashed-formatting, print-style and placement attributes.
type EmptyLine struct {
	Name string `xml:"-"`
	LineShape common.EnumLineShape `xml:"line-shape,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	LineLength common.EnumLineLength `xml:"line-length,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
}

// Extend: The extend type represents lyric word extension / melisma lines as
// well as figured bass extensions. The optional type and position attributes are added
// in Version 3.0 to provide better formatting control.
type Extend struct {
	Name string `xml:"-"`
	Type common.EnumStartStopContinue `xml:"type,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

// Figure: The figure type represents a single figure within a figured-bass
// element.
type Figure struct {
	Name string `xml:"-"`
	Prefix *direction.StyleText `xml:"prefix,omitempty"`
	FigureNumber *direction.StyleText `xml:"figure-number,omitempty"`
	Suffix *direction.StyleText `xml:"suffix,omitempty"`
	Extend *Extend `xml:"extend,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
}

// FiguredBass: The figured-bass element represents figured bass notation. Figured
// bass elements take their position from the first regular note (not a grace note or
// chord note) that follows in score order. The optional duration element is used to
// indicate changes of figures under a note. Figures are ordered from top to bottom.
// The value of parentheses is "no" if not present.
type FiguredBass struct {
	Name string `xml:"-"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
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
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	PrintDot common.EnumYesNo `xml:"print-dot,attr,omitempty"`
	PrintLyric common.EnumYesNo `xml:"print-lyric,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	PrintSpacing common.EnumYesNo `xml:"print-spacing,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Figure []*Figure `xml:"figure,omitempty"`
	Duration string `xml:"duration,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
}

// Forward: The backup and forward elements are required to coordinate multiple
// voices in one part, including music on multiple staves. The forward element is
// generally used within voices and staves. Duration values should always be positive,
// and should not cross measure boundaries or mid-measure changes in the divisions
// value.
type Forward struct {
	Name string `xml:"-"`
	Duration string `xml:"duration,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Voice string `xml:"voice,omitempty"`
	Staff int `xml:"staff,omitempty"`
}

// Glissando: Glissando and slide types both indicate rapidly moving from one pitch
// to the other so that individual notes are not discerned. A glissando sounds the
// distinct notes in between the two pitches and defaults to a wavy line. The optional
// text is printed alongside the line.
type Glissando struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Grace: The grace type indicates the presence of a grace note. The slash
// attribute for a grace note is yes for slashed grace notes. The steal-time-previous
// attribute indicates the percentage of time to steal from the previous note for the
// grace note. The steal-time-following attribute indicates the percentage of time to
// steal from the following note for the grace note, as for appoggiaturas. The
// make-time attribute indicates to make time, not steal time; the units are in
// real-time divisions for the grace note.
type Grace struct {
	Name string `xml:"-"`
	StealTimePrevious string `xml:"steal-time-previous,attr,omitempty"`
	StealTimeFollowing string `xml:"steal-time-following,attr,omitempty"`
	MakeTime string `xml:"make-time,attr,omitempty"`
	Slash common.EnumYesNo `xml:"slash,attr,omitempty"`
}

// HammerOnPullOff: The hammer-on and pull-off elements are used in guitar and fretted
// instrument notation. Since a single slur can be marked over many notes, the
// hammer-on and pull-off elements are separate so the individual pair of notes can be
// specified. The element content can be used to specify how the hammer-on or pull-off
// should be notated. An empty element leaves this choice up to the application.
type HammerOnPullOff struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
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
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Handbell: The handbell element represents notation for various techniques used
// in handbell and handchime music.
type Handbell struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value EnumHandbellValue `xml:",chardata"`
}

// HarmonClosed: The harmon-closed type represents whether the harmon mute is closed,
// open, or half-open. The optional location attribute indicates which portion of the
// symbol is filled in when the element value is half.
type HarmonClosed struct {
	Name string `xml:"-"`
	Location EnumHarmonClosedLocation `xml:"location,attr,omitempty"`
	Value EnumHarmonClosedValue `xml:",chardata"`
}

// HarmonMute: The harmon-mute type represents the symbols used for harmon mutes in
// brass notation.
type HarmonMute struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	HarmonClosed *HarmonClosed `xml:"harmon-closed,omitempty"`
}

// Harmonic: The harmonic type indicates natural and artificial harmonics. Allowing
// the type of pitch to be specified, combined with controls for appearance/playback
// differences, allows both the notation and the sound to be represented. Artificial
// harmonics can add a notated touching pitch; artificial pinch harmonics will usually
// not notate a touching pitch. The attributes for the harmonic element refer to the
// use of the circular harmonic symbol, typically but not always used with natural
// harmonics.
type Harmonic struct {
	Name string `xml:"-"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Natural *common.Empty `xml:"natural,omitempty"`
	Artificial *common.Empty `xml:"artificial,omitempty"`
	BasePitch *common.Empty `xml:"base-pitch,omitempty"`
	TouchingPitch *common.Empty `xml:"touching-pitch,omitempty"`
	SoundingPitch *common.Empty `xml:"sounding-pitch,omitempty"`
}

// HeelToe: The heel and toe elements are used with organ pedals. The substitution
// value is "no" if the attribute is not present.
type HeelToe struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Substitution common.EnumYesNo `xml:"substitution,attr,omitempty"`
}

// Hole: The hole type represents the symbols used for woodwind and brass
// fingerings as well as other notations.
type Hole struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	HoleType string `xml:"hole-type,omitempty"`
	HoleClosed *HoleClosed `xml:"hole-closed,omitempty"`
	HoleShape string `xml:"hole-shape,omitempty"`
}

// HoleClosed: The hole-closed type represents whether the hole is closed, open, or
// half-open. The optional location attribute indicates which portion of the hole is
// filled in when the element value is half.
type HoleClosed struct {
	Name string `xml:"-"`
	Location EnumHoleClosedLocation `xml:"location,attr,omitempty"`
	Value EnumHoleClosedValue `xml:",chardata"`
}

// Instrument: The instrument type distinguishes between score-instrument elements in
// a score-part. The id attribute is an IDREF back to the score-instrument ID. If
// multiple score-instruments are specified in a score-part, there should be an
// instrument element for each note in the part. Notes that are shared between multiple
// score-instruments can have more than one instrument element.
type Instrument struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
}

// Listen: The listen and listening types, new in Version 4.0, specify different
// ways that a score following or machine listening application can interact with a
// performer. The listen type handles interactions that are specific to a note. If
// multiple child elements of the same type are present, they should have distinct
// player and/or time-only attributes.
type Listen struct {
	Name string `xml:"-"`
	Assess []*Assess `xml:"assess,omitempty"`
	Wait []*Wait `xml:"wait,omitempty"`
	OtherListen []*direction.OtherListening `xml:"other-listen,omitempty"`
}

// Lyric: The lyric type represents text underlays for lyrics. Two text elements
// that are not separated by an elision element are part of the same syllable, but may
// have different text formatting. The MusicXML XSD is more strict than the DTD in
// enforcing this by disallowing a second syllabic element unless preceded by an
// elision element. The lyric number indicates multiple lines, though a name can be
// used as well. Common name examples are verse and chorus. Justification is center by
// default; placement is below by default. Vertical alignment is to the baseline of the
// text and horizontal alignment matches justification. The print-object attribute can
// override a note's print-lyric attribute in cases where only some lyrics on a note
// are printed, as when lyrics for later verses are printed in a block of text rather
// than with each note. The time-only attribute precisely specifies which lyrics are to
// be sung which time through a repeated section.
type Lyric struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	NameAttr string `xml:"name,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
	Justify common.EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Syllabic EnumSyllabic `xml:"syllabic,omitempty"`
	Text *TextElementData `xml:"text,omitempty"`
	Elision []*Elision `xml:"elision,omitempty"`
	Extend *Extend `xml:"extend,omitempty"`
	Laughing *common.Empty `xml:"laughing,omitempty"`
	Humming *common.Empty `xml:"humming,omitempty"`
	EndLine *common.Empty `xml:"end-line,omitempty"`
	EndParagraph *common.Empty `xml:"end-paragraph,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
}

// Mordent: The mordent type is used for both represents the mordent sign with the
// vertical line and the inverted-mordent sign without the line. The long attribute is
// "no" by default. The approach and departure attributes are used for compound
// ornaments, indicating how the beginning and ending of the ornament look relative to
// the main part of the mordent.
type Mordent struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	StartNote common.EnumStartNote `xml:"start-note,attr,omitempty"`
	TrillStep common.EnumTrillStep `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn common.EnumTwoNoteTurn `xml:"two-note-turn,attr,omitempty"`
	Accelerate common.EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	SecondBeat string `xml:"second-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
	Long common.EnumYesNo `xml:"long,attr,omitempty"`
	Approach common.EnumAboveBelow `xml:"approach,attr,omitempty"`
	Departure common.EnumAboveBelow `xml:"departure,attr,omitempty"`
}

// NonArpeggiate: The non-arpeggiate type indicates that this note is at the top or
// bottom of a bracket indicating to not arpeggiate these notes. Since this does not
// involve playback, it is only used on the top or bottom notes, not on each note as
// for the arpeggiate type.
type NonArpeggiate struct {
	Name string `xml:"-"`
	Type common.EnumTopBottom `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Notations: Notations refer to musical notations, not XML notations. Multiple
// notations are allowed in order to represent multiple editorial levels. The
// print-object attribute, added in Version 3.0, allows notations to represent details
// of performance technique, such as fingerings, without having them appear in the
// score.
type Notations struct {
	Name string `xml:"-"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Tied []*Tied `xml:"tied,omitempty"`
	Slur []*Slur `xml:"slur,omitempty"`
	Tuplet []*Tuplet `xml:"tuplet,omitempty"`
	Glissando []*Glissando `xml:"glissando,omitempty"`
	Slide []*Slide `xml:"slide,omitempty"`
	Ornaments []*Ornaments `xml:"ornaments,omitempty"`
	Technical []*Technical `xml:"technical,omitempty"`
	Articulations []*Articulations `xml:"articulations,omitempty"`
	Dynamics []*common.Dynamics `xml:"dynamics,omitempty"`
	Fermata []*common.Fermata `xml:"fermata,omitempty"`
	Arpeggiate []*Arpeggiate `xml:"arpeggiate,omitempty"`
	NonArpeggiate []*NonArpeggiate `xml:"non-arpeggiate,omitempty"`
	AccidentalMark []*AccidentalMark `xml:"accidental-mark,omitempty"`
	OtherNotation []*OtherNotation `xml:"other-notation,omitempty"`
}

// Note: Notes are the most common type of MusicXML data. The MusicXML format
// distinguishes between elements used for sound information and elements used for
// notation information (e.g., tie is used for sound, tied for notation). Thus grace
// notes do not have a duration element. Cue notes have a duration element, as do
// forward elements, but no tie elements. Having these two types of information
// available can make interchange easier, as some programs handle one type of
// information more readily than the other. The print-leger attribute is used to
// indicate whether leger lines are printed. Notes without leger lines are used to
// indicate indeterminate high and low notes. By default, it is set to yes. If
// print-object is set to no, print-leger is interpreted to also be set to no if not
// present. This attribute is ignored for rests. The dynamics and end-dynamics
// attributes correspond to MIDI 1.0's Note On and Note Off velocities, respectively.
// They are expressed in terms of percentages of the default forte value (90 for MIDI
// 1.0). The attack and release attributes are used to alter the starting and stopping
// time of the note from when it would otherwise occur based on the flow of durations -
// information that is specific to a performance. They are expressed in terms of
// divisions, either positive or negative. A note that starts a tie should not have a
// release attribute, and a note that stops a tie should not have an attack attribute.
// The attack and release attributes are independent of each other. The attack
// attribute only changes the starting time of a note, and the release attribute only
// changes the stopping time of a note. If a note is played only particular times
// through a repeat, the time-only attribute shows which times to play the note. The
// pizzicato attribute is used when just this note is sounded pizzicato, vs. the
// pizzicato element which changes overall playback between pizzicato and arco.
type Note struct {
	Name string `xml:"-"`
	PrintLeger common.EnumYesNo `xml:"print-leger,attr,omitempty"`
	Dynamics string `xml:"dynamics,attr,omitempty"`
	EndDynamics string `xml:"end-dynamics,attr,omitempty"`
	Attack string `xml:"attack,attr,omitempty"`
	Release string `xml:"release,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
	Pizzicato common.EnumYesNo `xml:"pizzicato,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	PrintDot common.EnumYesNo `xml:"print-dot,attr,omitempty"`
	PrintLyric common.EnumYesNo `xml:"print-lyric,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	PrintSpacing common.EnumYesNo `xml:"print-spacing,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Grace *Grace `xml:"grace,omitempty"`
	Chord *common.Empty `xml:"chord,omitempty"`
	Pitch *Pitch `xml:"pitch,omitempty"`
	Unpitched *Unpitched `xml:"unpitched,omitempty"`
	Rest *Rest `xml:"rest,omitempty"`
	Tie []*Tie `xml:"tie,omitempty"`
	Cue *common.Empty `xml:"cue,omitempty"`
	Duration string `xml:"duration,omitempty"`
	Instrument []*Instrument `xml:"instrument,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Voice string `xml:"voice,omitempty"`
	Type *NoteType `xml:"type,omitempty"`
	Dot []*common.EmptyPlacement `xml:"dot,omitempty"`
	Accidental *Accidental `xml:"accidental,omitempty"`
	TimeModification *direction.TimeModification `xml:"time-modification,omitempty"`
	Stem *Stem `xml:"stem,omitempty"`
	Notehead *Notehead `xml:"notehead,omitempty"`
	NoteheadText *NoteheadText `xml:"notehead-text,omitempty"`
	Staff int `xml:"staff,omitempty"`
	Beam []*Beam `xml:"beam,omitempty"`
	Notations []*Notations `xml:"notations,omitempty"`
	Lyric []*Lyric `xml:"lyric,omitempty"`
	Play *common.Play `xml:"play,omitempty"`
	Listen *Listen `xml:"listen,omitempty"`
}

// NoteType: The note-type type indicates the graphic note type. Values range from
// 1024th to maxima. The size attribute indicates full, cue, grace-cue, or large size.
// The default is full for regular notes, grace-cue for notes that contain both grace
// and cue elements, and cue for notes that contain either a cue or a grace element,
// but not both.
type NoteType struct {
	Name string `xml:"-"`
	Size common.EnumSymbolSize `xml:"size,attr,omitempty"`
	Value common.EnumNoteTypeValue `xml:",chardata"`
}

// Notehead: The notehead type indicates shapes other than the open and closed
// ovals associated with note durations. The smufl attribute can be used to specify a
// particular notehead, allowing application interoperability without requiring every
// SMuFL glyph to have a MusicXML element equivalent. This attribute can be used either
// with the "other" value, or to refine a specific notehead value such as "cluster".
// Noteheads in the SMuFL Note name noteheads and Note name noteheads supplement ranges
// (U+E150–U+E1AF and U+EEE0–U+EEFF) should not use the smufl attribute or the "other"
// value, but instead use the notehead-text element. For the enclosed shapes, the
// default is to be hollow for half notes and longer, and filled otherwise. The filled
// attribute can be set to change this if needed. If the parentheses attribute is set
// to yes, the notehead is parenthesized. It is no by default.
type Notehead struct {
	Name string `xml:"-"`
	Filled common.EnumYesNo `xml:"filled,attr,omitempty"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumNoteheadValue `xml:",chardata"`
}

// NoteheadText: The notehead-text type represents text that is displayed inside a
// notehead, as is done in some educational music. It is not needed for the numbers
// used in tablature or jianpu notation. The presence of a TAB or jianpu clefs is
// sufficient to indicate that numbers are used. The display-text and accidental-text
// elements allow display of fully formatted text and accidentals.
type NoteheadText struct {
	Name string `xml:"-"`
	DisplayText []*common.FormattedText `xml:"display-text,omitempty"`
	AccidentalText []*common.AccidentalText `xml:"accidental-text,omitempty"`
}

// Ornaments: Ornaments can be any of several types, followed optionally by
// accidentals. The accidental-mark element's content is represented the same as an
// accidental element, but with a different name to reflect the different musical
// meaning.
type Ornaments struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	TrillMark []*common.EmptyTrillSound `xml:"trill-mark,omitempty"`
	Turn []*common.HorizontalTurn `xml:"turn,omitempty"`
	DelayedTurn []*common.HorizontalTurn `xml:"delayed-turn,omitempty"`
	InvertedTurn []*common.HorizontalTurn `xml:"inverted-turn,omitempty"`
	DelayedInvertedTurn []*common.HorizontalTurn `xml:"delayed-inverted-turn,omitempty"`
	VerticalTurn []*common.EmptyTrillSound `xml:"vertical-turn,omitempty"`
	InvertedVerticalTurn []*common.EmptyTrillSound `xml:"inverted-vertical-turn,omitempty"`
	Shake []*common.EmptyTrillSound `xml:"shake,omitempty"`
	WavyLine []*common.WavyLine `xml:"wavy-line,omitempty"`
	Mordent []*Mordent `xml:"mordent,omitempty"`
	InvertedMordent []*Mordent `xml:"inverted-mordent,omitempty"`
	Schleifer []*common.EmptyPlacement `xml:"schleifer,omitempty"`
	Tremolo []*Tremolo `xml:"tremolo,omitempty"`
	Haydn []*common.EmptyTrillSound `xml:"haydn,omitempty"`
	OtherOrnament []*OtherPlacementText `xml:"other-ornament,omitempty"`
	AccidentalMark []*AccidentalMark `xml:"accidental-mark,omitempty"`
}

// OtherNotation: The other-notation type is used to define any notations not yet in the
// MusicXML format. It handles notations where more specific extension elements such as
// other-dynamics and other-technical are not appropriate. The smufl attribute can be
// used to specify a particular notation, allowing application interoperability without
// requiring every SMuFL glyph to have a MusicXML element equivalent. Using the
// other-notation type without the smufl attribute allows for extended representation,
// though without application interoperability.
type OtherNotation struct {
	Name string `xml:"-"`
	Type common.EnumStartStopSingle `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// OtherPlacementText: The other-placement-text type represents a text element with
// print-style, placement, and smufl attribute groups. This type is used by MusicXML
// notation extension elements to allow specification of specific SMuFL glyphs without
// needed to add every glyph as a MusicXML element.
type OtherPlacementText struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Pitch: Pitch is represented as a combination of the step of the diatonic
// scale, the chromatic alteration, and the octave.
type Pitch struct {
	Name string `xml:"-"`
	Step attributes.EnumStep `xml:"step,omitempty"`
	Alter string `xml:"alter,omitempty"`
	Octave int `xml:"octave,omitempty"`
}

// PlacementText: The placement-text type represents a text element with print-style and
// placement attribute groups.
type PlacementText struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Release: The release type indicates that a bend is a release rather than a
// normal bend or pre-bend. The offset attribute specifies where the release starts in
// terms of divisions relative to the current note. The first-beat and last-beat
// attributes of the parent bend element are relative to the original note position,
// not this offset value.
type Release struct {
	Name string `xml:"-"`
	Offset string `xml:"offset,attr,omitempty"`
}

// Rest: The rest element indicates notated rests or silences. Rest elements
// are usually empty, but placement on the staff can be specified using display-step
// and display-octave elements. If the measure attribute is set to yes, this indicates
// this is a complete measure rest.
type Rest struct {
	Name string `xml:"-"`
	Measure common.EnumYesNo `xml:"measure,attr,omitempty"`
	DisplayStep attributes.EnumStep `xml:"display-step,omitempty"`
	DisplayOctave int `xml:"display-octave,omitempty"`
}

// Slide: Glissando and slide types both indicate rapidly moving from one pitch
// to the other so that individual notes are not discerned. A slide is continuous
// between the two pitches and defaults to a solid line. The optional text for a is
// printed alongside the line.
type Slide struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Accelerate common.EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	FirstBeat string `xml:"first-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Slur: Slur types are empty. Most slurs are represented with two elements:
// one with a start type, and one with a stop type. Slurs can add more elements using a
// continue type. This is typically used to specify the formatting of cross-system
// slurs, or to specify the shape of very complex slurs.
type Slur struct {
	Name string `xml:"-"`
	Type common.EnumStartStopContinue `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Orientation common.EnumOverUnder `xml:"orientation,attr,omitempty"`
	BezierX string `xml:"bezier-x,attr,omitempty"`
	BezierY string `xml:"bezier-y,attr,omitempty"`
	BezierX2 string `xml:"bezier-x2,attr,omitempty"`
	BezierY2 string `xml:"bezier-y2,attr,omitempty"`
	BezierOffset string `xml:"bezier-offset,attr,omitempty"`
	BezierOffset2 string `xml:"bezier-offset2,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Stem: Stems can be down, up, none, or double. For down and up stems, the
// position attributes can be used to specify stem length. The relative values specify
// the end of the stem relative to the program default. Default values specify an
// absolute end stem position. Negative values of relative-y that would flip a stem
// instead of shortening it are ignored. A stem element associated with a rest refers
// to a stemlet.
type Stem struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value EnumStemValue `xml:",chardata"`
}

// StrongAccent: The strong-accent type indicates a vertical accent mark. The type
// attribute indicates if the point of the accent is down or up.
type StrongAccent struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Type common.EnumUpDown `xml:"type,attr,omitempty"`
}

// Tap: The tap type indicates a tap on the fretboard. The text content allows
// specification of the notation; + and T are common choices. If the element is empty,
// the hand attribute is used to specify the symbol to use. The hand attribute is
// ignored if the tap glyph is already specified by the text content. If neither text
// content nor the hand attribute are present, the display is application-specific.
type Tap struct {
	Name string `xml:"-"`
	Hand EnumTapHand `xml:"hand,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Technical: Technical indications give performance information for individual
// instruments.
type Technical struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	UpBow []*common.EmptyPlacement `xml:"up-bow,omitempty"`
	DownBow []*common.EmptyPlacement `xml:"down-bow,omitempty"`
	Harmonic []*Harmonic `xml:"harmonic,omitempty"`
	OpenString []*common.EmptyPlacement `xml:"open-string,omitempty"`
	ThumbPosition []*common.EmptyPlacement `xml:"thumb-position,omitempty"`
	Fingering []*common.Fingering `xml:"fingering,omitempty"`
	Pluck []*PlacementText `xml:"pluck,omitempty"`
	DoubleTongue []*common.EmptyPlacement `xml:"double-tongue,omitempty"`
	TripleTongue []*common.EmptyPlacement `xml:"triple-tongue,omitempty"`
	Stopped []*common.EmptyPlacementSmufl `xml:"stopped,omitempty"`
	SnapPizzicato []*common.EmptyPlacement `xml:"snap-pizzicato,omitempty"`
	Fret []*common.Fret `xml:"fret,omitempty"`
	String []*common.StringType `xml:"string,omitempty"`
	HammerOn []*HammerOnPullOff `xml:"hammer-on,omitempty"`
	PullOff []*HammerOnPullOff `xml:"pull-off,omitempty"`
	Bend []*Bend `xml:"bend,omitempty"`
	Tap []*Tap `xml:"tap,omitempty"`
	Heel []*HeelToe `xml:"heel,omitempty"`
	Toe []*HeelToe `xml:"toe,omitempty"`
	Fingernails []*common.EmptyPlacement `xml:"fingernails,omitempty"`
	Hole []*Hole `xml:"hole,omitempty"`
	Arrow []*Arrow `xml:"arrow,omitempty"`
	Handbell []*Handbell `xml:"handbell,omitempty"`
	BrassBend []*common.EmptyPlacement `xml:"brass-bend,omitempty"`
	Flip []*common.EmptyPlacement `xml:"flip,omitempty"`
	Smear []*common.EmptyPlacement `xml:"smear,omitempty"`
	Open []*common.EmptyPlacementSmufl `xml:"open,omitempty"`
	HalfMuted []*common.EmptyPlacementSmufl `xml:"half-muted,omitempty"`
	HarmonMute []*HarmonMute `xml:"harmon-mute,omitempty"`
	Golpe []*common.EmptyPlacement `xml:"golpe,omitempty"`
	OtherTechnical []*OtherPlacementText `xml:"other-technical,omitempty"`
}

// TextElementData: The text-element-data type represents a syllable or portion of a
// syllable for lyric text underlay. A hyphen in the string content should only be used
// for an actual hyphenated word. Language names for text elements come from ISO 639,
// with optional country subcodes from ISO 3166.
type TextElementData struct {
	Name string `xml:"-"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	Dir common.EnumTextDirection `xml:"dir,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Tie: The tie element indicates that a tie begins or ends with this note. If
// the tie element applies only particular times through a repeat, the time-only
// attribute indicates which times to apply it. The tie element indicates sound; the
// tied element indicates notation.
type Tie struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
}

// Tied: The tied element represents the notated tie. The tie element
// represents the tie sound. The number attribute is rarely needed to disambiguate
// ties, since note pitches will usually suffice. The attribute is implied rather than
// defaulting to 1 as with most elements. It is available for use in more complex tied
// notation situations. Ties that join two notes of the same pitch together should be
// represented with a tied element on the first note with type="start" and a tied
// element on the second note with type="stop". This can also be done if the two notes
// being tied are enharmonically equivalent, but have different step values. It is not
// recommended to use tied elements to join two notes with enharmonically inequivalent
// pitches. Ties that indicate that an instrument should be undamped are specified with
// a single tied element with type="let-ring". Ties that are visually attached to only
// one note, other than undamped ties, should be specified with two tied elements on
// the same note, first type="start" then type="stop". This can be used to represent
// ties into or out of repeated sections or codas.
type Tied struct {
	Name string `xml:"-"`
	Type common.EnumTiedType `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Orientation common.EnumOverUnder `xml:"orientation,attr,omitempty"`
	BezierX string `xml:"bezier-x,attr,omitempty"`
	BezierY string `xml:"bezier-y,attr,omitempty"`
	BezierX2 string `xml:"bezier-x2,attr,omitempty"`
	BezierY2 string `xml:"bezier-y2,attr,omitempty"`
	BezierOffset string `xml:"bezier-offset,attr,omitempty"`
	BezierOffset2 string `xml:"bezier-offset2,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Tremolo: The tremolo ornament can be used to indicate single-note, double-note,
// or unmeasured tremolos. Single-note tremolos use the single type, double-note
// tremolos use the start and stop types, and unmeasured tremolos use the unmeasured
// type. The default is "single" for compatibility with Version 1.1. The text of the
// element indicates the number of tremolo marks and is an integer from 0 to 8. Note
// that the number of attached beams is not included in this value, but is represented
// separately using the beam element. The value should be 0 for unmeasured tremolos.
// When using double-note tremolos, the duration of each note in the tremolo should
// correspond to half of the notated type value. A time-modification element should
// also be added with an actual-notes value of 2 and a normal-notes value of 1. If used
// within a tuplet, this 2/1 ratio should be multiplied by the existing tuplet ratio.
// The smufl attribute specifies the glyph to use from the SMuFL Tremolos range for an
// unmeasured tremolo. It is ignored for other tremolo types. The SMuFL buzzRoll glyph
// is used by default if the attribute is missing. Using repeater beams for indicating
// tremolos is deprecated as of MusicXML 3.0.
type Tremolo struct {
	Name string `xml:"-"`
	Type common.EnumTremoloType `xml:"type,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Tuplet: A tuplet element is present when a tuplet is to be displayed
// graphically, in addition to the sound data provided by the time-modification
// elements. The number attribute is used to distinguish nested tuplets. The bracket
// attribute is used to indicate the presence of a bracket. If unspecified, the results
// are implementation-dependent. The line-shape attribute is used to specify whether
// the bracket is straight or in the older curved or slurred style. It is straight by
// default. Whereas a time-modification element shows how the cumulative, sounding
// effect of tuplets and double-note tremolos compare to the written note type, the
// tuplet element describes how this is displayed. The tuplet element also provides
// more detailed representation information than the time-modification element, and is
// needed to represent nested tuplets and other complex tuplets accurately. The
// show-number attribute is used to display either the number of actual notes, the
// number of both actual and normal notes, or neither. It is actual by default. The
// show-type attribute is used to display either the actual type, both the actual and
// normal types, or neither. It is none by default.
type Tuplet struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	Bracket common.EnumYesNo `xml:"bracket,attr,omitempty"`
	ShowNumber direction.EnumShowTuplet `xml:"show-number,attr,omitempty"`
	ShowType direction.EnumShowTuplet `xml:"show-type,attr,omitempty"`
	LineShape common.EnumLineShape `xml:"line-shape,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	TupletActual *TupletPortion `xml:"tuplet-actual,omitempty"`
	TupletNormal *TupletPortion `xml:"tuplet-normal,omitempty"`
}

// TupletDot: The tuplet-dot type is used to specify dotted tuplet types.
type TupletDot struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

// TupletNumber: The tuplet-number type indicates the number of notes for this portion
// of the tuplet.
type TupletNumber struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value int `xml:",chardata"`
}

// TupletPortion: The tuplet-portion type provides optional full control over tuplet
// specifications. It allows the number and note type (including dots) to be set for
// the actual and normal portions of a single tuplet. If any of these elements are
// absent, their values are based on the time-modification element.
type TupletPortion struct {
	Name string `xml:"-"`
	TupletNumber *TupletNumber `xml:"tuplet-number,omitempty"`
	TupletType *TupletType `xml:"tuplet-type,omitempty"`
	TupletDot []*TupletDot `xml:"tuplet-dot,omitempty"`
}

// TupletType: The tuplet-type type indicates the graphical note type of the notes
// for this portion of the tuplet.
type TupletType struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value common.EnumNoteTypeValue `xml:",chardata"`
}

// Unpitched: The unpitched type represents musical elements that are notated on the
// staff but lack definite pitch, such as unpitched percussion and speaking voice. If
// the child elements are not present, the note is placed on the middle line of the
// staff. This is generally used with a one-line staff. Notes in percussion clef should
// always use an unpitched element rather than a pitch element.
type Unpitched struct {
	Name string `xml:"-"`
	DisplayStep attributes.EnumStep `xml:"display-step,omitempty"`
	DisplayOctave int `xml:"display-octave,omitempty"`
}

// Wait: The wait type specifies a point where the accompaniment should wait
// for a performer event before continuing. This typically happens at the start of new
// sections or after a held note or indeterminate music. These waiting points cannot
// always be inferred reliably from the contents of the displayed score. The optional
// player and time-only attributes restrict the type to apply to a single player or set
// of times through a repeated section, respectively.
type Wait struct {
	Name string `xml:"-"`
	Player string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
}

