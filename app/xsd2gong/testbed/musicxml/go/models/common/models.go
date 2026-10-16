// generated code by xsd2gong - do not edit
package common

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

var _ xml.Attr

// EnumAboveBelow: The above-below type is used to indicate whether one element appears
// above or below another element.
type EnumAboveBelow string

const (
	EnumAboveBelowAbove EnumAboveBelow = "above"
	EnumAboveBelowBelow EnumAboveBelow = "below"
)

// EnumCssFontSize: The css-font-size type includes the CSS font sizes used as an
// alternative to a numeric point size.
type EnumCssFontSize string

const (
	EnumCssFontSizeXxSmall EnumCssFontSize = "xx-small"
	EnumCssFontSizeXSmall EnumCssFontSize = "x-small"
	EnumCssFontSizeSmall EnumCssFontSize = "small"
	EnumCssFontSizeMedium EnumCssFontSize = "medium"
	EnumCssFontSizeLarge EnumCssFontSize = "large"
	EnumCssFontSizeXLarge EnumCssFontSize = "x-large"
	EnumCssFontSizeXxLarge EnumCssFontSize = "xx-large"
)

// EnumEnclosureShape: The enclosure-shape type describes the shape and presence / absence of
// an enclosure around text or symbols. A bracket enclosure is similar to a rectangle
// with the bottom line missing, as is common in jazz notation. An inverted-bracket
// enclosure is similar to a rectangle with the top line missing.
type EnumEnclosureShape string

const (
	EnumEnclosureShapeRectangle EnumEnclosureShape = "rectangle"
	EnumEnclosureShapeSquare EnumEnclosureShape = "square"
	EnumEnclosureShapeOval EnumEnclosureShape = "oval"
	EnumEnclosureShapeCircle EnumEnclosureShape = "circle"
	EnumEnclosureShapeBracket EnumEnclosureShape = "bracket"
	EnumEnclosureShapeInvertedBracket EnumEnclosureShape = "inverted-bracket"
	EnumEnclosureShapeTriangle EnumEnclosureShape = "triangle"
	EnumEnclosureShapeDiamond EnumEnclosureShape = "diamond"
	EnumEnclosureShapePentagon EnumEnclosureShape = "pentagon"
	EnumEnclosureShapeHexagon EnumEnclosureShape = "hexagon"
	EnumEnclosureShapeHeptagon EnumEnclosureShape = "heptagon"
	EnumEnclosureShapeOctagon EnumEnclosureShape = "octagon"
	EnumEnclosureShapeNonagon EnumEnclosureShape = "nonagon"
	EnumEnclosureShapeDecagon EnumEnclosureShape = "decagon"
	EnumEnclosureShapeNone EnumEnclosureShape = "none"
)

// EnumFermataShape: The fermata-shape type represents the shape of the fermata sign. The
// empty value is equivalent to the normal value.
type EnumFermataShape string

const (
	EnumFermataShapeNormal EnumFermataShape = "normal"
	EnumFermataShapeAngled EnumFermataShape = "angled"
	EnumFermataShapeSquare EnumFermataShape = "square"
	EnumFermataShapeDoubleAngled EnumFermataShape = "double-angled"
	EnumFermataShapeDoubleSquare EnumFermataShape = "double-square"
	EnumFermataShapeDoubleDot EnumFermataShape = "double-dot"
	EnumFermataShapeHalfCurve EnumFermataShape = "half-curve"
	EnumFermataShapeCurlew EnumFermataShape = "curlew"
	EnumFermataShapeEmpty EnumFermataShape = ""
)

// EnumFontStyle: The font-style type represents a simplified version of the CSS
// font-style property.
type EnumFontStyle string

const (
	EnumFontStyleNormal EnumFontStyle = "normal"
	EnumFontStyleItalic EnumFontStyle = "italic"
)

// EnumFontWeight: The font-weight type represents a simplified version of the CSS
// font-weight property.
type EnumFontWeight string

const (
	EnumFontWeightNormal EnumFontWeight = "normal"
	EnumFontWeightBold EnumFontWeight = "bold"
)

// EnumLeftCenterRight: The left-center-right type is used to define horizontal alignment and
// text justification.
type EnumLeftCenterRight string

const (
	EnumLeftCenterRightLeft EnumLeftCenterRight = "left"
	EnumLeftCenterRightCenter EnumLeftCenterRight = "center"
	EnumLeftCenterRightRight EnumLeftCenterRight = "right"
)

// EnumLeftRight: The left-right type is used to indicate whether one element appears to
// the left or the right of another element.
type EnumLeftRight string

const (
	EnumLeftRightLeft EnumLeftRight = "left"
	EnumLeftRightRight EnumLeftRight = "right"
)

// EnumLineLength: The line-length type distinguishes between different line lengths for
// doit, falloff, plop, and scoop articulations.
type EnumLineLength string

const (
	EnumLineLengthShort EnumLineLength = "short"
	EnumLineLengthMedium EnumLineLength = "medium"
	EnumLineLengthLong EnumLineLength = "long"
)

// EnumLineShape: The line-shape type distinguishes between straight and curved lines.
type EnumLineShape string

const (
	EnumLineShapeStraight EnumLineShape = "straight"
	EnumLineShapeCurved EnumLineShape = "curved"
)

// EnumLineType: The line-type type distinguishes between solid, dashed, dotted, and
// wavy lines.
type EnumLineType string

const (
	EnumLineTypeSolid EnumLineType = "solid"
	EnumLineTypeDashed EnumLineType = "dashed"
	EnumLineTypeDotted EnumLineType = "dotted"
	EnumLineTypeWavy EnumLineType = "wavy"
)

// EnumMute: The mute type represents muting for different instruments, including
// brass, winds, and strings. The on and off values are used for undifferentiated
// mutes. The remaining values represent specific mutes.
type EnumMute string

const (
	EnumMuteOn EnumMute = "on"
	EnumMuteOff EnumMute = "off"
	EnumMuteStraight EnumMute = "straight"
	EnumMuteCup EnumMute = "cup"
	EnumMuteHarmonNoStem EnumMute = "harmon-no-stem"
	EnumMuteHarmonStem EnumMute = "harmon-stem"
	EnumMuteBucket EnumMute = "bucket"
	EnumMutePlunger EnumMute = "plunger"
	EnumMuteHat EnumMute = "hat"
	EnumMuteSolotone EnumMute = "solotone"
	EnumMutePractice EnumMute = "practice"
	EnumMuteStopMute EnumMute = "stop-mute"
	EnumMuteStopHand EnumMute = "stop-hand"
	EnumMuteEcho EnumMute = "echo"
	EnumMutePalm EnumMute = "palm"
)

// EnumOverUnder: The over-under type is used to indicate whether the tips of curved
// lines such as slurs and ties are overhand (tips down) or underhand (tips up).
type EnumOverUnder string

const (
	EnumOverUnderOver EnumOverUnder = "over"
	EnumOverUnderUnder EnumOverUnder = "under"
)

// EnumSemiPitched: The semi-pitched type represents categories of indefinite pitch for
// percussion instruments.
type EnumSemiPitched string

const (
	EnumSemiPitchedHigh EnumSemiPitched = "high"
	EnumSemiPitchedMediumHigh EnumSemiPitched = "medium-high"
	EnumSemiPitchedMedium EnumSemiPitched = "medium"
	EnumSemiPitchedMediumLow EnumSemiPitched = "medium-low"
	EnumSemiPitchedLow EnumSemiPitched = "low"
	EnumSemiPitchedVeryLow EnumSemiPitched = "very-low"
)

// EnumStartNote: The start-note type describes the starting note of trills and mordents
// for playback, relative to the current note.
type EnumStartNote string

const (
	EnumStartNoteUpper EnumStartNote = "upper"
	EnumStartNoteMain EnumStartNote = "main"
	EnumStartNoteBelow EnumStartNote = "below"
)

// EnumStartStop: The start-stop type is used for an attribute of musical elements that
// can either start or stop, such as tuplets. The values of start and stop refer to how
// an element appears in musical score order, not in MusicXML document order. An
// element with a stop attribute may precede the corresponding element with a start
// attribute within a MusicXML document. This is particularly common in multi-staff
// music. For example, the stopping point for a tuplet may appear in staff 1 before the
// starting point for the tuplet appears in staff 2 later in the document. When
// multiple elements with the same tag are used within the same note, their order
// within the MusicXML document should match the musical score order.
type EnumStartStop string

const (
	EnumStartStopStart EnumStartStop = "start"
	EnumStartStopStop EnumStartStop = "stop"
)

// EnumStartStopContinue: The start-stop-continue type is used for an attribute of musical
// elements that can either start or stop, but also need to refer to an intermediate
// point in the symbol, as for complex slurs or for formatting of symbols across system
// breaks. The values of start, stop, and continue refer to how an element appears in
// musical score order, not in MusicXML document order. An element with a stop
// attribute may precede the corresponding element with a start attribute within a
// MusicXML document. This is particularly common in multi-staff music. For example,
// the stopping point for a slur may appear in staff 1 before the starting point for
// the slur appears in staff 2 later in the document. When multiple elements with the
// same tag are used within the same note, their order within the MusicXML document
// should match the musical score order. For example, a note that marks both the end of
// one slur and the start of a new slur should have the incoming slur element with a
// type of stop precede the outgoing slur element with a type of start.
type EnumStartStopContinue string

const (
	EnumStartStopContinueStart EnumStartStopContinue = "start"
	EnumStartStopContinueStop EnumStartStopContinue = "stop"
	EnumStartStopContinueContinue EnumStartStopContinue = "continue"
)

// EnumStartStopSingle: The start-stop-single type is used for an attribute of musical
// elements that can be used for either multi-note or single-note musical elements, as
// for groupings. When multiple elements with the same tag are used within the same
// note, their order within the MusicXML document should match the musical score order.
type EnumStartStopSingle string

const (
	EnumStartStopSingleStart EnumStartStopSingle = "start"
	EnumStartStopSingleStop EnumStartStopSingle = "stop"
	EnumStartStopSingleSingle EnumStartStopSingle = "single"
)

// EnumSymbolSize: The symbol-size type is used to distinguish between full, cue sized,
// grace cue sized, and oversized symbols.
type EnumSymbolSize string

const (
	EnumSymbolSizeFull EnumSymbolSize = "full"
	EnumSymbolSizeCue EnumSymbolSize = "cue"
	EnumSymbolSizeGraceCue EnumSymbolSize = "grace-cue"
	EnumSymbolSizeLarge EnumSymbolSize = "large"
)

// EnumTextDirection: The text-direction type is used to adjust and override the Unicode
// bidirectional text algorithm, similar to the Directionality data category in the W3C
// Internationalization Tag Set recommendation. Values are ltr (left-to-right embed),
// rtl (right-to-left embed), lro (left-to-right bidi-override), and rlo (right-to-left
// bidi-override). The default value is ltr. This type is typically used by
// applications that store text in left-to-right visual order rather than logical
// order. Such applications can use the lro value to better communicate with other
// applications that more fully support bidirectional text.
type EnumTextDirection string

const (
	EnumTextDirectionLtr EnumTextDirection = "ltr"
	EnumTextDirectionRtl EnumTextDirection = "rtl"
	EnumTextDirectionLro EnumTextDirection = "lro"
	EnumTextDirectionRlo EnumTextDirection = "rlo"
)

// EnumTiedType: The tied-type type is used as an attribute of the tied element to
// specify where the visual representation of a tie begins and ends. A tied element
// which joins two notes of the same pitch can be specified with tied-type start on the
// first note and tied-type stop on the second note. To indicate a note should be
// undamped, use a single tied element with tied-type let-ring. For other ties that are
// visually attached to a single note, such as a tie leading into or out of a repeated
// section or coda, use two tied elements on the same note, one start and one stop. In
// start-stop cases, ties can add more elements using a continue type. This is
// typically used to specify the formatting of cross-system ties. When multiple
// elements with the same tag are used within the same note, their order within the
// MusicXML document should match the musical score order. For example, a note with a
// tie at the end of a first ending should have the tied element with a type of start
// precede the tied element with a type of stop.
type EnumTiedType string

const (
	EnumTiedTypeStart EnumTiedType = "start"
	EnumTiedTypeStop EnumTiedType = "stop"
	EnumTiedTypeContinue EnumTiedType = "continue"
	EnumTiedTypeLetRing EnumTiedType = "let-ring"
)

// EnumTopBottom: The top-bottom type is used to indicate the top or bottom part of a
// vertical shape like non-arpeggiate.
type EnumTopBottom string

const (
	EnumTopBottomTop EnumTopBottom = "top"
	EnumTopBottomBottom EnumTopBottom = "bottom"
)

// EnumTremoloType: The tremolo-type is used to distinguish double-note, single-note, and
// unmeasured tremolos.
type EnumTremoloType string

const (
	EnumTremoloTypeStart EnumTremoloType = "start"
	EnumTremoloTypeStop EnumTremoloType = "stop"
	EnumTremoloTypeSingle EnumTremoloType = "single"
	EnumTremoloTypeUnmeasured EnumTremoloType = "unmeasured"
)

// EnumTrillStep: The trill-step type describes the alternating note of trills and
// mordents for playback, relative to the current note.
type EnumTrillStep string

const (
	EnumTrillStepWhole EnumTrillStep = "whole"
	EnumTrillStepHalf EnumTrillStep = "half"
	EnumTrillStepUnison EnumTrillStep = "unison"
)

// EnumTwoNoteTurn: The two-note-turn type describes the ending notes of trills and
// mordents for playback, relative to the current note.
type EnumTwoNoteTurn string

const (
	EnumTwoNoteTurnWhole EnumTwoNoteTurn = "whole"
	EnumTwoNoteTurnHalf EnumTwoNoteTurn = "half"
	EnumTwoNoteTurnNone EnumTwoNoteTurn = "none"
)

// EnumUpDown: The up-down type is used for the direction of arrows and other pointed
// symbols like vertical accents, indicating which way the tip is pointing.
type EnumUpDown string

const (
	EnumUpDownUp EnumUpDown = "up"
	EnumUpDownDown EnumUpDown = "down"
)

// EnumUprightInverted: The upright-inverted type describes the appearance of a fermata
// element. The value is upright if not specified.
type EnumUprightInverted string

const (
	EnumUprightInvertedUpright EnumUprightInverted = "upright"
	EnumUprightInvertedInverted EnumUprightInverted = "inverted"
)

// EnumValign: The valign type is used to indicate vertical alignment to the top,
// middle, bottom, or baseline of the text. If the text is on multiple lines, baseline
// alignment refers to the baseline of the lowest line of text. Defaults are
// implementation-dependent.
type EnumValign string

const (
	EnumValignTop EnumValign = "top"
	EnumValignMiddle EnumValign = "middle"
	EnumValignBottom EnumValign = "bottom"
	EnumValignBaseline EnumValign = "baseline"
)

// EnumValignImage: The valign-image type is used to indicate vertical alignment for
// images and graphics, so it does not include a baseline value. Defaults are
// implementation-dependent.
type EnumValignImage string

const (
	EnumValignImageTop EnumValignImage = "top"
	EnumValignImageMiddle EnumValignImage = "middle"
	EnumValignImageBottom EnumValignImage = "bottom"
)

// EnumYesNo: The yes-no type is used for boolean-like attributes. We cannot use W3C
// XML Schema booleans due to their restrictions on expression of boolean values.
type EnumYesNo string

const (
	EnumYesNoYes EnumYesNo = "yes"
	EnumYesNoNo EnumYesNo = "no"
)

// EnumMarginType: The margin-type type specifies whether margins apply to even page, odd
// pages, or both.
type EnumMarginType string

const (
	EnumMarginTypeOdd EnumMarginType = "odd"
	EnumMarginTypeEven EnumMarginType = "even"
	EnumMarginTypeBoth EnumMarginType = "both"
)

// EnumNoteSizeType: The note-size-type type indicates the type of note being defined by a
// note-size element. The grace-cue type is used for notes of grace-cue size. The grace
// type is used for notes of cue size that include a grace element. The cue type is
// used for all other notes with cue size, whether defined explicitly or implicitly via
// a cue element. The large type is used for notes of large size.
type EnumNoteSizeType string

const (
	EnumNoteSizeTypeCue EnumNoteSizeType = "cue"
	EnumNoteSizeTypeGrace EnumNoteSizeType = "grace"
	EnumNoteSizeTypeGraceCue EnumNoteSizeType = "grace-cue"
	EnumNoteSizeTypeLarge EnumNoteSizeType = "large"
)

// EnumAccidentalValue: The accidental-value type represents notated accidentals supported by
// MusicXML. In the MusicXML 2.0 DTD this was a string with values that could be
// included. The XSD strengthens the data typing to an enumerated list. The quarter-
// and three-quarters- accidentals are Tartini-style quarter-tone accidentals. The
// -down and -up accidentals are quarter-tone accidentals that include arrows pointing
// down or up. The slash- accidentals are used in Turkish classical music. The numbered
// sharp and flat accidentals are superscripted versions of the accidental signs, used
// in Turkish folk music. The sori and koron accidentals are microtonal sharp and flat
// accidentals used in Iranian and Persian music. The other accidental covers
// accidentals other than those listed here. It is usually used in combination with the
// smufl attribute to specify a particular SMuFL accidental. The smufl attribute may be
// used with any accidental value to help specify the appearance of symbols that share
// the same MusicXML semantics.
type EnumAccidentalValue string

const (
	EnumAccidentalValueSharp EnumAccidentalValue = "sharp"
	EnumAccidentalValueNatural EnumAccidentalValue = "natural"
	EnumAccidentalValueFlat EnumAccidentalValue = "flat"
	EnumAccidentalValueDoubleSharp EnumAccidentalValue = "double-sharp"
	EnumAccidentalValueSharpSharp EnumAccidentalValue = "sharp-sharp"
	EnumAccidentalValueFlatFlat EnumAccidentalValue = "flat-flat"
	EnumAccidentalValueNaturalSharp EnumAccidentalValue = "natural-sharp"
	EnumAccidentalValueNaturalFlat EnumAccidentalValue = "natural-flat"
	EnumAccidentalValueQuarterFlat EnumAccidentalValue = "quarter-flat"
	EnumAccidentalValueQuarterSharp EnumAccidentalValue = "quarter-sharp"
	EnumAccidentalValueThreeQuartersFlat EnumAccidentalValue = "three-quarters-flat"
	EnumAccidentalValueThreeQuartersSharp EnumAccidentalValue = "three-quarters-sharp"
	EnumAccidentalValueSharpDown EnumAccidentalValue = "sharp-down"
	EnumAccidentalValueSharpUp EnumAccidentalValue = "sharp-up"
	EnumAccidentalValueNaturalDown EnumAccidentalValue = "natural-down"
	EnumAccidentalValueNaturalUp EnumAccidentalValue = "natural-up"
	EnumAccidentalValueFlatDown EnumAccidentalValue = "flat-down"
	EnumAccidentalValueFlatUp EnumAccidentalValue = "flat-up"
	EnumAccidentalValueDoubleSharpDown EnumAccidentalValue = "double-sharp-down"
	EnumAccidentalValueDoubleSharpUp EnumAccidentalValue = "double-sharp-up"
	EnumAccidentalValueFlatFlatDown EnumAccidentalValue = "flat-flat-down"
	EnumAccidentalValueFlatFlatUp EnumAccidentalValue = "flat-flat-up"
	EnumAccidentalValueArrowDown EnumAccidentalValue = "arrow-down"
	EnumAccidentalValueArrowUp EnumAccidentalValue = "arrow-up"
	EnumAccidentalValueTripleSharp EnumAccidentalValue = "triple-sharp"
	EnumAccidentalValueTripleFlat EnumAccidentalValue = "triple-flat"
	EnumAccidentalValueSlashQuarterSharp EnumAccidentalValue = "slash-quarter-sharp"
	EnumAccidentalValueSlashSharp EnumAccidentalValue = "slash-sharp"
	EnumAccidentalValueSlashFlat EnumAccidentalValue = "slash-flat"
	EnumAccidentalValueDoubleSlashFlat EnumAccidentalValue = "double-slash-flat"
	EnumAccidentalValueSharp1 EnumAccidentalValue = "sharp-1"
	EnumAccidentalValueSharp2 EnumAccidentalValue = "sharp-2"
	EnumAccidentalValueSharp3 EnumAccidentalValue = "sharp-3"
	EnumAccidentalValueSharp5 EnumAccidentalValue = "sharp-5"
	EnumAccidentalValueFlat1 EnumAccidentalValue = "flat-1"
	EnumAccidentalValueFlat2 EnumAccidentalValue = "flat-2"
	EnumAccidentalValueFlat3 EnumAccidentalValue = "flat-3"
	EnumAccidentalValueFlat4 EnumAccidentalValue = "flat-4"
	EnumAccidentalValueSori EnumAccidentalValue = "sori"
	EnumAccidentalValueKoron EnumAccidentalValue = "koron"
	EnumAccidentalValueOther EnumAccidentalValue = "other"
)

// EnumNoteTypeValue: The note-type-value type is used for the MusicXML type element and
// represents the graphic note type, from 1024th (shortest) to maxima (longest).
type EnumNoteTypeValue string

const (
	EnumNoteTypeValueN1024th EnumNoteTypeValue = "1024th"
	EnumNoteTypeValueN512th EnumNoteTypeValue = "512th"
	EnumNoteTypeValueN256th EnumNoteTypeValue = "256th"
	EnumNoteTypeValueN128th EnumNoteTypeValue = "128th"
	EnumNoteTypeValueN64th EnumNoteTypeValue = "64th"
	EnumNoteTypeValueN32nd EnumNoteTypeValue = "32nd"
	EnumNoteTypeValueN16th EnumNoteTypeValue = "16th"
	EnumNoteTypeValueEighth EnumNoteTypeValue = "eighth"
	EnumNoteTypeValueQuarter EnumNoteTypeValue = "quarter"
	EnumNoteTypeValueHalf EnumNoteTypeValue = "half"
	EnumNoteTypeValueWhole EnumNoteTypeValue = "whole"
	EnumNoteTypeValueBreve EnumNoteTypeValue = "breve"
	EnumNoteTypeValueLong EnumNoteTypeValue = "long"
	EnumNoteTypeValueMaxima EnumNoteTypeValue = "maxima"
)

// EnumSwingTypeValue: The swing-type-value type specifies the note type, either eighth or
// 16th, to which the ratio defined in the swing element is applied.
type EnumSwingTypeValue string

const (
	EnumSwingTypeValueN16th EnumSwingTypeValue = "16th"
	EnumSwingTypeValueEighth EnumSwingTypeValue = "eighth"
)

// AccidentalText: The accidental-text type represents an element with an accidental
// value and text-formatting attributes.
type AccidentalText struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`
	Justify EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	LineHeight string `xml:"line-height,attr,omitempty"`
	Dir EnumTextDirection `xml:"dir,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Value EnumAccidentalValue `xml:",chardata"`
}

// Coda: The coda type is the visual indicator of a coda sign. The exact glyph
// can be specified with the smufl attribute. A sound element is also needed to guide
// playback applications reliably.
type Coda struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Dynamics: Dynamics can be associated either with a note or a general musical
// direction. To avoid inconsistencies between and amongst the letter abbreviations for
// dynamics (what is sf vs. sfz, standing alone or with a trailing dynamic that is not
// always piano), we use the actual letters as the names of these dynamic elements. The
// other-dynamics element allows other dynamic marks that are not covered here.
// Dynamics elements may also be combined to create marks not covered by a single
// element, such as sfmp. These letter dynamic symbols are separated from crescendo,
// decrescendo, and wedge indications. Dynamic representation is inconsistent in
// scores. Many things are assumed by the composer and left out, such as returns to
// original dynamics. The MusicXML format captures what is in the score, but does not
// try to be optimal for analysis or synthesis of dynamics. The placement attribute is
// used when the dynamics are associated with a note. It is ignored when the dynamics
// are associated with a direction. In that case the direction element's placement
// attribute is used instead.
type Dynamics struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	P []*Empty `xml:"p,omitempty"`
	Pp []*Empty `xml:"pp,omitempty"`
	Ppp []*Empty `xml:"ppp,omitempty"`
	Pppp []*Empty `xml:"pppp,omitempty"`
	Ppppp []*Empty `xml:"ppppp,omitempty"`
	Pppppp []*Empty `xml:"pppppp,omitempty"`
	F []*Empty `xml:"f,omitempty"`
	Ff []*Empty `xml:"ff,omitempty"`
	Fff []*Empty `xml:"fff,omitempty"`
	Ffff []*Empty `xml:"ffff,omitempty"`
	Fffff []*Empty `xml:"fffff,omitempty"`
	Ffffff []*Empty `xml:"ffffff,omitempty"`
	Mp []*Empty `xml:"mp,omitempty"`
	Mf []*Empty `xml:"mf,omitempty"`
	Sf []*Empty `xml:"sf,omitempty"`
	Sfp []*Empty `xml:"sfp,omitempty"`
	Sfpp []*Empty `xml:"sfpp,omitempty"`
	Fp []*Empty `xml:"fp,omitempty"`
	Rf []*Empty `xml:"rf,omitempty"`
	Rfz []*Empty `xml:"rfz,omitempty"`
	Sfz []*Empty `xml:"sfz,omitempty"`
	Sffz []*Empty `xml:"sffz,omitempty"`
	Fz []*Empty `xml:"fz,omitempty"`
	N []*Empty `xml:"n,omitempty"`
	Pf []*Empty `xml:"pf,omitempty"`
	Sfzp []*Empty `xml:"sfzp,omitempty"`
	OtherDynamics []*OtherText `xml:"other-dynamics,omitempty"`
}

// Empty: The empty type represents an empty element with no attributes.
type Empty struct {
	Name string `xml:"-"`
}

// EmptyPlacement: The empty-placement type represents an empty element with print-style
// and placement attributes.
type EmptyPlacement struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
}

// EmptyPlacementSmufl: The empty-placement-smufl type represents an empty element with
// print-style, placement, and smufl attributes.
type EmptyPlacementSmufl struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
}

// EmptyPrintStyle: The empty-print-style type represents an empty element with
// print-style attribute group.
type EmptyPrintStyle struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

// EmptyPrintStyleAlign: The empty-print-style-align type represents an empty element with
// print-style-align attribute group.
type EmptyPrintStyleAlign struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
}

// EmptyPrintStyleAlignId: The empty-print-style-align-id type represents an empty element with
// print-style-align and optional-unique-id attribute groups.
type EmptyPrintStyleAlignId struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// EmptyPrintObjectStyleAlign: The empty-print-style-align-object type represents an empty element
// with print-object and print-style-align attribute groups.
type EmptyPrintObjectStyleAlign struct {
	Name string `xml:"-"`
	PrintObject EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
}

// EmptyTrillSound: The empty-trill-sound type represents an empty element with
// print-style, placement, and trill-sound attributes.
type EmptyTrillSound struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	StartNote EnumStartNote `xml:"start-note,attr,omitempty"`
	TrillStep EnumTrillStep `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn EnumTwoNoteTurn `xml:"two-note-turn,attr,omitempty"`
	Accelerate EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	SecondBeat string `xml:"second-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
}

// HorizontalTurn: The horizontal-turn type represents turn elements that are horizontal
// rather than vertical. These are empty elements with print-style, placement,
// trill-sound, and slash attributes. If the slash attribute is yes, then a vertical
// line is used to slash the turn. It is no if not specified.
type HorizontalTurn struct {
	Name string `xml:"-"`
	Slash EnumYesNo `xml:"slash,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	StartNote EnumStartNote `xml:"start-note,attr,omitempty"`
	TrillStep EnumTrillStep `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn EnumTwoNoteTurn `xml:"two-note-turn,attr,omitempty"`
	Accelerate EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	SecondBeat string `xml:"second-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
}

// Fermata: The fermata text content represents the shape of the fermata sign. An
// empty fermata element represents a normal fermata. The fermata type is upright if
// not specified.
type Fermata struct {
	Name string `xml:"-"`
	Type EnumUprightInverted `xml:"type,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value EnumFermataShape `xml:",chardata"`
}

// Fingering: Fingering is typically indicated 1,2,3,4,5. Multiple fingerings may be
// given, typically to substitute fingerings in the middle of a note. The substitution
// and alternate values are "no" if the attribute is not present. For guitar and other
// fretted instruments, the fingering element represents the fretting finger; the pluck
// element represents the plucking finger.
type Fingering struct {
	Name string `xml:"-"`
	Substitution EnumYesNo `xml:"substitution,attr,omitempty"`
	Alternate EnumYesNo `xml:"alternate,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value string `xml:",chardata"`
}

// FormattedSymbol: The formatted-symbol type represents a SMuFL musical symbol element
// with formatting attributes.
type FormattedSymbol struct {
	Name string `xml:"-"`
	Justify EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	LineHeight string `xml:"line-height,attr,omitempty"`
	Dir EnumTextDirection `xml:"dir,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Value string `xml:",chardata"`
}

// FormattedSymbolId: The formatted-symbol-id type represents a SMuFL musical symbol element
// with formatting and id attributes.
type FormattedSymbolId struct {
	Name string `xml:"-"`
	Justify EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	LineHeight string `xml:"line-height,attr,omitempty"`
	Dir EnumTextDirection `xml:"dir,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// FormattedText: The formatted-text type represents a text element with text-formatting
// attributes.
type FormattedText struct {
	Name string `xml:"-"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`
	Justify EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	LineHeight string `xml:"line-height,attr,omitempty"`
	Dir EnumTextDirection `xml:"dir,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Value string `xml:",chardata"`
}

// FormattedTextId: The formatted-text-id type represents a text element with
// text-formatting and id attributes.
type FormattedTextId struct {
	Name string `xml:"-"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`
	Justify EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Underline int `xml:"underline,attr,omitempty"`
	Overline int `xml:"overline,attr,omitempty"`
	LineThrough int `xml:"line-through,attr,omitempty"`
	Rotation string `xml:"rotation,attr,omitempty"`
	LetterSpacing string `xml:"letter-spacing,attr,omitempty"`
	LineHeight string `xml:"line-height,attr,omitempty"`
	Dir EnumTextDirection `xml:"dir,attr,omitempty"`
	Enclosure EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Fret: The fret element is used with tablature notation and chord diagrams.
// Fret numbers start with 0 for an open string and 1 for the first fret.
type Fret struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Level: The level type is used to specify editorial information for different
// MusicXML elements. The content contains identifying and/or descriptive text about
// the editorial status of the parent element. If the reference attribute is yes, this
// indicates editorial information that is for display only and should not affect
// playback. For instance, a modern edition of older music may set reference="yes" on
// the attributes containing the music's original clef, key, and time signature. It is
// no if not specified. The type attribute indicates whether the editorial information
// applies to the start of a series of symbols, the end of a series of symbols, or a
// single symbol. It is single if not specified for compatibility with earlier MusicXML
// versions.
type Level struct {
	Name string `xml:"-"`
	Reference EnumYesNo `xml:"reference,attr,omitempty"`
	Type EnumStartStopSingle `xml:"type,attr,omitempty"`
	Parentheses EnumYesNo `xml:"parentheses,attr,omitempty"`
	Bracket EnumYesNo `xml:"bracket,attr,omitempty"`
	Size EnumSymbolSize `xml:"size,attr,omitempty"`
	Value string `xml:",chardata"`
}

// MidiDevice: The midi-device type corresponds to the DeviceName meta event in
// Standard MIDI Files. The optional port attribute is a number from 1 to 16 that can
// be used with the unofficial MIDI 1.0 port (or cable) meta event. Unlike the
// DeviceName meta event, there can be multiple midi-device elements per MusicXML part.
// The optional id attribute refers to the score-instrument assigned to this device. If
// missing, the device assignment affects all score-instrument elements in the
// score-part.
type MidiDevice struct {
	Name string `xml:"-"`
	Port int `xml:"port,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// MidiInstrument: The midi-instrument type defines MIDI 1.0 instrument playback. The
// midi-instrument element can be a part of either the score-instrument element at the
// start of a part, or the sound element within a part. The id attribute refers to the
// score-instrument affected by the change.
type MidiInstrument struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	MidiChannel int `xml:"midi-channel,omitempty"`
	MidiName string `xml:"midi-name,omitempty"`
	MidiBank int `xml:"midi-bank,omitempty"`
	MidiProgram int `xml:"midi-program,omitempty"`
	MidiUnpitched int `xml:"midi-unpitched,omitempty"`
	Volume string `xml:"volume,omitempty"`
	Pan string `xml:"pan,omitempty"`
	Elevation string `xml:"elevation,omitempty"`
}

// NameDisplay: The name-display type is used for exact formatting of multi-font text
// in part and group names to the left of the system. The print-object attribute can be
// used to determine what, if anything, is printed at the start of each system.
// Enclosure for the display-text element is none by default. Language for the
// display-text element is Italian ("it") by default.
type NameDisplay struct {
	Name string `xml:"-"`
	PrintObject EnumYesNo `xml:"print-object,attr,omitempty"`
	DisplayText []*FormattedText `xml:"display-text,omitempty"`
	AccidentalText []*AccidentalText `xml:"accidental-text,omitempty"`
}

// OtherPlay: The other-play element represents other types of playback. The
// required type attribute indicates the type of playback to which the element content
// applies.
type OtherPlay struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Play: The play type specifies playback techniques to be used in conjunction
// with the instrument-sound element. When used as part of a sound element, it applies
// to all notes going forward in score order. In multi-instrument parts, the affected
// instrument should be specified using the id attribute. When used as part of a note
// element, it applies to the current note only.
type Play struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Ipa string `xml:"ipa,omitempty"`
	Mute EnumMute `xml:"mute,omitempty"`
	SemiPitched EnumSemiPitched `xml:"semi-pitched,omitempty"`
	OtherPlay []*OtherPlay `xml:"other-play,omitempty"`
}

// Segno: The segno type is the visual indicator of a segno sign. The exact
// glyph can be specified with the smufl attribute. A sound element is also needed to
// guide playback applications reliably.
type Segno struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// StringType: The string type is used with tablature notation, regular notation
// (where it is often circled), and chord diagrams. String numbers start with 1 for the
// highest pitched full-length string.
type StringType struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	Value int `xml:",chardata"`
}

// TypedText: The typed-text type represents a text element with a type attribute.
type TypedText struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// WavyLine: Wavy lines are one way to indicate trills and vibrato. When used with
// a barline element, they should always have type="continue" set. The smufl attribute
// specifies a particular wavy line glyph from the SMuFL Multi-segment lines range.
type WavyLine struct {
	Name string `xml:"-"`
	Type EnumStartStopContinue `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Placement EnumAboveBelow `xml:"placement,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	StartNote EnumStartNote `xml:"start-note,attr,omitempty"`
	TrillStep EnumTrillStep `xml:"trill-step,attr,omitempty"`
	TwoNoteTurn EnumTwoNoteTurn `xml:"two-note-turn,attr,omitempty"`
	Accelerate EnumYesNo `xml:"accelerate,attr,omitempty"`
	Beats string `xml:"beats,attr,omitempty"`
	SecondBeat string `xml:"second-beat,attr,omitempty"`
	LastBeat string `xml:"last-beat,attr,omitempty"`
}

// Encoding: The encoding element contains information about who did the digital
// encoding, when, with what software, and in what aspects. Standard type values for
// the encoder element are music, words, and arrangement, but other types may be used.
// The type attribute is only needed when there are multiple encoder elements.
type Encoding struct {
	Name string `xml:"-"`
	EncodingDate string `xml:"encoding-date,omitempty"`
	Encoder []*TypedText `xml:"encoder,omitempty"`
	Software string `xml:"software,omitempty"`
	EncodingDescription string `xml:"encoding-description,omitempty"`
	Supports []*Supports `xml:"supports,omitempty"`
}

// Identification: Identification contains basic metadata about the score. It includes
// information that may apply at a score-wide, movement-wide, or part-wide level. The
// creator, rights, source, and relation elements are based on Dublin Core.
type Identification struct {
	Name string `xml:"-"`
	Creator []*TypedText `xml:"creator,omitempty"`
	Rights []*TypedText `xml:"rights,omitempty"`
	Encoding *Encoding `xml:"encoding,omitempty"`
	Source string `xml:"source,omitempty"`
	Relation []*TypedText `xml:"relation,omitempty"`
	Miscellaneous *Miscellaneous `xml:"miscellaneous,omitempty"`
}

// Miscellaneous: If a program has other metadata not yet supported in the MusicXML
// format, it can go in the miscellaneous element. The miscellaneous type puts each
// separate part of metadata into its own miscellaneous-field type.
type Miscellaneous struct {
	Name string `xml:"-"`
	MiscellaneousField []*MiscellaneousField `xml:"miscellaneous-field,omitempty"`
}

// MiscellaneousField: If a program has other metadata not yet supported in the MusicXML
// format, each type of metadata can go in a miscellaneous-field element. The required
// name attribute indicates the type of metadata the element content represents.
type MiscellaneousField struct {
	Name string `xml:"-"`
	NameAttr string `xml:"name,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Supports: The supports type indicates if a MusicXML encoding supports a
// particular MusicXML element. This is recommended for elements like beam, stem, and
// accidental, where the absence of an element is ambiguous if you do not know if the
// encoding supports that element. For Version 2.0, the supports element is expanded to
// allow programs to indicate support for particular attributes or particular values.
// This lets applications communicate, for example, that all system and/or page breaks
// are contained in the MusicXML file.
type Supports struct {
	Name string `xml:"-"`
	Type EnumYesNo `xml:"type,attr,omitempty"`
	Element string `xml:"element,attr,omitempty"`
	Attribute string `xml:"attribute,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

// Appearance: The appearance type controls general graphical settings for the
// music's final form appearance on a printed page of display. This includes support
// for line widths, definitions for note sizes, and standard distances between notation
// elements, plus an extension element for other aspects of appearance.
type Appearance struct {
	Name string `xml:"-"`
	LineWidth []*LineWidth `xml:"line-width,omitempty"`
	NoteSize []*NoteSize `xml:"note-size,omitempty"`
	Distance []*Distance `xml:"distance,omitempty"`
	Glyph []*Glyph `xml:"glyph,omitempty"`
	OtherAppearance []*OtherAppearance `xml:"other-appearance,omitempty"`
}

// Distance: The distance element represents standard distances between notation
// elements in tenths. The type attribute defines what type of distance is being
// defined. Valid values include hyphen (for hyphens in lyrics) and beam.
type Distance struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Glyph: The glyph element represents what SMuFL glyph should be used for
// different variations of symbols that are semantically identical. The type attribute
// specifies what type of glyph is being defined. The element value specifies what
// SMuFL glyph to use, including recommended stylistic alternates. The SMuFL glyph name
// should match the type. For instance, a type of quarter-rest would use values
// restQuarter, restQuarterOld, or restQuarterZ. A type of g-clef-ottava-bassa would
// use values gClef8vb, gClef8vbOld, or gClef8vbCClef. A type of octave-shift-up-8
// would use values ottava, ottavaBassa, ottavaBassaBa, ottavaBassaVb, or octaveBassa.
type Glyph struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// LineWidth: The line-width type indicates the width of a line type in tenths. The
// type attribute defines what type of line is being defined. Values include beam,
// bracket, dashes, enclosure, ending, extend, heavy barline, leger, light barline,
// octave shift, pedal, slur middle, slur tip, staff, stem, tie middle, tie tip, tuplet
// bracket, and wedge. The text content is expressed in tenths.
type LineWidth struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// MeasureLayout: The measure-layout type includes the horizontal distance from the
// previous measure. It applies to the current measure only.
type MeasureLayout struct {
	Name string `xml:"-"`
	MeasureDistance string `xml:"measure-distance,omitempty"`
}

// NoteSize: The note-size type indicates the percentage of the regular note size
// to use for notes with a cue and large size as defined in the type element. The grace
// type is used for notes of cue size that that include a grace element. The cue type
// is used for all other notes with cue size, whether defined explicitly or implicitly
// via a cue element. The large type is used for notes of large size. The text content
// represent the numeric percentage. A value of 100 would be identical to the size of a
// regular note as defined by the music font.
type NoteSize struct {
	Name string `xml:"-"`
	Type EnumNoteSizeType `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// OtherAppearance: The other-appearance type is used to define any graphical settings not
// yet in the current version of the MusicXML format. This allows extended
// representation, though without application interoperability.
type OtherAppearance struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// PageLayout: Page layout can be defined both in score-wide defaults and in the
// print element. Page margins are specified either for both even and odd pages, or via
// separate odd and even page number values. The type is not needed when used as part
// of a print element. If omitted when used in the defaults element, "both" is the
// default. If no page-layout element is present in the defaults element, default page
// layout values are chosen by the application. When used in the print element, the
// page-layout element affects the appearance of the current page only. All other pages
// use the default values as determined by the defaults element. If any child elements
// are missing from the page-layout element in a print element, the values determined
// by the defaults element are used there as well.
type PageLayout struct {
	Name string `xml:"-"`
	PageHeight string `xml:"page-height,omitempty"`
	PageWidth string `xml:"page-width,omitempty"`
	PageMargins []*PageMargins `xml:"page-margins,omitempty"`
}

// PageMargins: Page margins are specified either for both even and odd pages, or via
// separate odd and even page number values. The type attribute is not needed when used
// as part of a print element. If omitted when the page-margins type is used in the
// defaults element, "both" is the default value.
type PageMargins struct {
	Name string `xml:"-"`
	Type EnumMarginType `xml:"type,attr,omitempty"`
	LeftMargin string `xml:"left-margin,omitempty"`
	RightMargin string `xml:"right-margin,omitempty"`
	TopMargin string `xml:"top-margin,omitempty"`
	BottomMargin string `xml:"bottom-margin,omitempty"`
}

// Scaling: Margins, page sizes, and distances are all measured in tenths to keep
// MusicXML data in a consistent coordinate system as much as possible. The translation
// to absolute units is done with the scaling type, which specifies how many
// millimeters are equal to how many tenths. For a staff height of 7 mm, millimeters
// would be set to 7 while tenths is set to 40. The ability to set a formula rather
// than a single scaling factor helps avoid roundoff errors.
type Scaling struct {
	Name string `xml:"-"`
	Millimeters string `xml:"millimeters,omitempty"`
	Tenths string `xml:"tenths,omitempty"`
}

// StaffLayout: Staff layout includes the vertical distance from the bottom line of
// the previous staff in this system to the top line of the staff specified by the
// number attribute. The optional number attribute refers to staff numbers within the
// part, from top to bottom on the system. A value of 1 is used if not present. When
// used in the defaults element, the values apply to all systems in all parts. When
// used in the print element, the values apply to the current system only. This value
// is ignored for the first staff in a system.
type StaffLayout struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	StaffDistance string `xml:"staff-distance,omitempty"`
}

// SystemDividers: The system-dividers element indicates the presence or absence of
// system dividers (also known as system separation marks) between systems displayed on
// the same page. Dividers on the left and right side of the page are controlled by the
// left-divider and right-divider elements respectively. The default vertical position
// is half the system-distance value from the top of the system that is below the
// divider. The default horizontal position is the left and right system margin,
// respectively. When used in the print element, the system-dividers element affects
// the dividers that would appear between the current system and the previous system.
type SystemDividers struct {
	Name string `xml:"-"`
	LeftDivider *EmptyPrintObjectStyleAlign `xml:"left-divider,omitempty"`
	RightDivider *EmptyPrintObjectStyleAlign `xml:"right-divider,omitempty"`
}

// SystemLayout: A system is a group of staves that are read and played simultaneously.
// System layout includes left and right margins and the vertical distance from the
// previous system. The system distance is measured from the bottom line of the
// previous system to the top line of the current system. It is ignored for the first
// system on a page. The top system distance is measured from the page's top margin to
// the top line of the first system. It is ignored for all but the first system on a
// page. Sometimes the sum of measure widths in a system may not equal the system width
// specified by the layout elements due to roundoff or other errors. The behavior when
// reading MusicXML files in these cases is application-dependent. For instance,
// applications may find that the system layout data is more reliable than the sum of
// the measure widths, and adjust the measure widths accordingly. When used in the
// defaults element, the system-layout element defines a default appearance for all
// systems in the score. If no system-layout element is present in the defaults
// element, default system layout values are chosen by the application. When used in
// the print element, the system-layout element affects the appearance of the current
// system only. All other systems use the default values as determined by the defaults
// element. If any child elements are missing from the system-layout element in a print
// element, the values determined by the defaults element are used there as well. This
// type of system-layout element need only be read from or written to the first visible
// part in the score.
type SystemLayout struct {
	Name string `xml:"-"`
	SystemMargins *SystemMargins `xml:"system-margins,omitempty"`
	SystemDistance string `xml:"system-distance,omitempty"`
	TopSystemDistance string `xml:"top-system-distance,omitempty"`
	SystemDividers *SystemDividers `xml:"system-dividers,omitempty"`
}

// SystemMargins: System margins are relative to the page margins. Positive values
// indent and negative values reduce the margin size.
type SystemMargins struct {
	Name string `xml:"-"`
	LeftMargin string `xml:"left-margin,omitempty"`
	RightMargin string `xml:"right-margin,omitempty"`
}

// Bookmark: The bookmark type serves as a well-defined target for an incoming
// simple XLink.
type Bookmark struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	NameAttr string `xml:"name,attr,omitempty"`
	Element string `xml:"element,attr,omitempty"`
	Position int `xml:"position,attr,omitempty"`
}

// Link: The link type serves as an outgoing simple XLink. If a relative link
// is used within a document that is part of a compressed MusicXML file, the link is
// relative to the root folder of the zip file.
type Link struct {
	Name string `xml:"-"`
	NameAttr string `xml:"name,attr,omitempty"`
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Type xlink.EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show xlink.EnumShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate xlink.EnumActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	Element string `xml:"element,attr,omitempty"`
	Position int `xml:"position,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
}

// OtherText: The other-text type represents a text element with a smufl attribute
// group. This type is used by MusicXML direction extension elements to allow
// specification of specific SMuFL glyphs without needed to add every glyph as a
// MusicXML element.
type OtherText struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value string `xml:",chardata"`
}

