// generated code by xsd2gong - do not edit
package barline

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
)

var _ xml.Attr

// EnumBackwardForward: The backward-forward type is used to specify repeat directions. The
// start of the repeat has a forward direction while the end of the repeat has a
// backward direction.
type EnumBackwardForward string

const (
	EnumBackwardForwardBackward EnumBackwardForward = "backward"
	EnumBackwardForwardForward EnumBackwardForward = "forward"
)

// EnumBarStyle: The bar-style type represents barline style information. Choices are
// regular, dotted, dashed, heavy, light-light, light-heavy, heavy-light, heavy-heavy,
// tick (a short stroke through the top line), short (a partial barline between the 2nd
// and 4th lines), and none.
type EnumBarStyle string

const (
	EnumBarStyleRegular EnumBarStyle = "regular"
	EnumBarStyleDotted EnumBarStyle = "dotted"
	EnumBarStyleDashed EnumBarStyle = "dashed"
	EnumBarStyleHeavy EnumBarStyle = "heavy"
	EnumBarStyleLightLight EnumBarStyle = "light-light"
	EnumBarStyleLightHeavy EnumBarStyle = "light-heavy"
	EnumBarStyleHeavyLight EnumBarStyle = "heavy-light"
	EnumBarStyleHeavyHeavy EnumBarStyle = "heavy-heavy"
	EnumBarStyleTick EnumBarStyle = "tick"
	EnumBarStyleShort EnumBarStyle = "short"
	EnumBarStyleNone EnumBarStyle = "none"
)

// EnumRightLeftMiddle: The right-left-middle type is used to specify barline location.
type EnumRightLeftMiddle string

const (
	EnumRightLeftMiddleRight EnumRightLeftMiddle = "right"
	EnumRightLeftMiddleLeft EnumRightLeftMiddle = "left"
	EnumRightLeftMiddleMiddle EnumRightLeftMiddle = "middle"
)

// EnumStartStopDiscontinue: The start-stop-discontinue type is used to specify ending types.
// Typically, the start type is associated with the left barline of the first measure
// in an ending. The stop and discontinue types are associated with the right barline
// of the last measure in an ending. Stop is used when the ending mark concludes with a
// downward jog, as is typical for first endings. Discontinue is used when there is no
// downward jog, as is typical for second endings that do not conclude a piece.
type EnumStartStopDiscontinue string

const (
	EnumStartStopDiscontinueStart EnumStartStopDiscontinue = "start"
	EnumStartStopDiscontinueStop EnumStartStopDiscontinue = "stop"
	EnumStartStopDiscontinueDiscontinue EnumStartStopDiscontinue = "discontinue"
)

// EnumWinged: The winged attribute indicates whether the repeat has winged
// extensions that appear above and below the barline. The straight and curved values
// represent single wings, while the double-straight and double-curved values represent
// double wings. The none value indicates no wings and is the default.
type EnumWinged string

const (
	EnumWingedNone EnumWinged = "none"
	EnumWingedStraight EnumWinged = "straight"
	EnumWingedCurved EnumWinged = "curved"
	EnumWingedDoubleStraight EnumWinged = "double-straight"
	EnumWingedDoubleCurved EnumWinged = "double-curved"
)

// EnumSystemRelationNumber: The system-relation-number type distinguishes measure numbers that are
// associated with a system rather than the particular part where the element appears.
// A value of only-top or only-bottom indicates that the number should appear only on
// the top or bottom part of the current system, respectively. A value of also-top or
// also-bottom indicates that the number should appear on both the current part and the
// top or bottom part of the current system, respectively. If these values appear in a
// score, when parts are created the number should only appear once in this part, not
// twice. A value of none indicates that the number is associated only with the current
// part, not with the system.
type EnumSystemRelationNumber string

const (
	EnumSystemRelationNumberOnlyTop EnumSystemRelationNumber = "only-top"
	EnumSystemRelationNumberOnlyBottom EnumSystemRelationNumber = "only-bottom"
	EnumSystemRelationNumberAlsoTop EnumSystemRelationNumber = "also-top"
	EnumSystemRelationNumberAlsoBottom EnumSystemRelationNumber = "also-bottom"
	EnumSystemRelationNumberNone EnumSystemRelationNumber = "none"
)

// EnumSystemRelation: The system-relation type distinguishes elements that are associated
// with a system rather than the particular part where the element appears. A value of
// only-top indicates that the element should appear only on the top part of the
// current system. A value of also-top indicates that the element should appear on both
// the current part and the top part of the current system. If this value appears in a
// score, when parts are created the element should only appear once in this part, not
// twice. A value of none indicates that the element is associated only with the
// current part, not with the system.
type EnumSystemRelation string

const (
	EnumSystemRelationOnlyTop EnumSystemRelation = "only-top"
	EnumSystemRelationAlsoTop EnumSystemRelation = "also-top"
	EnumSystemRelationNone EnumSystemRelation = "none"
)

// BarStyleColor: The bar-style-color type contains barline style and color information.
type BarStyleColor struct {
	Name string `xml:"-"`
	Color string `xml:"color,attr,omitempty"`
	Value EnumBarStyle `xml:",chardata"`
}

// Barline: If a barline is other than a normal single barline, it should be
// represented by a barline type that describes it. This includes information about
// repeats and multiple endings, as well as line style. Barline data is on the same
// level as the other musical data in a score - a child of a measure in a partwise
// score, or a part in a timewise score. This allows for barlines within measures, as
// in dotted barlines that subdivide measures in complex meters. The two fermata
// elements allow for fermatas on both sides of the barline (the lower one inverted).
// Barlines have a location attribute to make it easier to process barlines
// independently of the other musical data in a score. It is often easier to set up
// measures separately from entering notes. The location attribute must match where the
// barline element occurs within the rest of the musical data in the score. If location
// is left, it should be the first element in the measure, aside from the print,
// bookmark, and link elements. If location is right, it should be the last element,
// again with the possible exception of the print, bookmark, and link elements. If no
// location is specified, the right barline is the default. The segno, coda, and
// divisions attributes work the same way as in the sound element. They are used for
// playback when barline elements contain segno or coda child elements.
type Barline struct {
	Name string `xml:"-"`
	Location EnumRightLeftMiddle `xml:"location,attr,omitempty"`
	Segno string `xml:"segno,attr,omitempty"`
	Coda string `xml:"coda,attr,omitempty"`
	Divisions string `xml:"divisions,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	BarStyle *BarStyleColor `xml:"bar-style,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	WavyLine *common.WavyLine `xml:"wavy-line,omitempty"`
	Segno_1 *common.Segno `xml:"segno,omitempty"`
	Coda_1 *common.Coda `xml:"coda,omitempty"`
	Fermata []*common.Fermata `xml:"fermata,omitempty"`
	Ending *Ending `xml:"ending,omitempty"`
	Repeat *Repeat `xml:"repeat,omitempty"`
}

// Ending: The ending type represents multiple (e.g. first and second) endings.
// Typically, the start type is associated with the left barline of the first measure
// in an ending. The stop and discontinue types are associated with the right barline
// of the last measure in an ending. Stop is used when the ending mark concludes with a
// downward jog, as is typical for first endings. Discontinue is used when there is no
// downward jog, as is typical for second endings that do not conclude a piece. The
// length of the jog can be specified using the end-length attribute. The text-x and
// text-y attributes are offsets that specify where the baseline of the start of the
// ending text appears, relative to the start of the ending line. The number attribute
// indicates which times the ending is played, similar to the time-only attribute used
// by other elements. While this often represents the numeric values for what is under
// the ending line, it can also indicate whether an ending is played during a larger
// dal segno or da capo repeat. Single endings such as "1" or comma-separated multiple
// endings such as "1,2" may be used. The ending element text is used when the text
// displayed in the ending is different than what appears in the number attribute. The
// print-object attribute is used to indicate when an ending is present but not
// printed, as is often the case for many parts in a full score.
type Ending struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	Type EnumStartStopDiscontinue `xml:"type,attr,omitempty"`
	EndLength string `xml:"end-length,attr,omitempty"`
	TextX string `xml:"text-x,attr,omitempty"`
	TextY string `xml:"text-y,attr,omitempty"`
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
	System EnumSystemRelation `xml:"system,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Repeat: The repeat type represents repeat marks. The start of the repeat has a
// forward direction while the end of the repeat has a backward direction. The times
// and after-jump attributes are only used with backward repeats that are not part of
// an ending. The times attribute indicates the number of times the repeated section is
// played. The after-jump attribute indicates if the repeats are played after a jump
// due to a da capo or dal segno.
type Repeat struct {
	Name string `xml:"-"`
	Direction EnumBackwardForward `xml:"direction,attr,omitempty"`
	Times int `xml:"times,attr,omitempty"`
	AfterJump common.EnumYesNo `xml:"after-jump,attr,omitempty"`
	Winged EnumWinged `xml:"winged,attr,omitempty"`
}

