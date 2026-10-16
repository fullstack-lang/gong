// generated code by xsd2gong - do not edit
package score

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

var _ xml.Attr

// EnumGroupBarlineValue: The group-barline-value type indicates if the group should have common
// barlines.
type EnumGroupBarlineValue string

const (
	EnumGroupBarlineValueYes EnumGroupBarlineValue = "yes"
	EnumGroupBarlineValueNo EnumGroupBarlineValue = "no"
	EnumGroupBarlineValueMensurstrich EnumGroupBarlineValue = "Mensurstrich"
)

// Credit: The credit type represents the appearance of the title, composer,
// arranger, lyricist, copyright, dedication, and other text, symbols, and graphics
// that commonly appear on the first page of a score. The credit-words, credit-symbol,
// and credit-image elements are similar to the words, symbol, and image elements for
// directions. However, since the credit is not part of a measure, the default-x and
// default-y attributes adjust the origin relative to the bottom left-hand corner of
// the page. The enclosure for credit-words and credit-symbol is none by default. By
// default, a series of credit-words and credit-symbol elements within a single credit
// element follow one another in sequence visually. Non-positional formatting
// attributes are carried over from the previous element by default. The page attribute
// for the credit element specifies the page number where the credit should appear.
// This is an integer value that starts with 1 for the first page. Its value is 1 by
// default. Since credits occur before the music, these page numbers do not refer to
// the page numbering specified by the print element's page-number attribute. The
// credit-type element indicates the purpose behind a credit. Multiple types of data
// may be combined in a single credit, so multiple elements may be used. Standard
// values include page number, title, subtitle, composer, arranger, lyricist, rights,
// and part name.
type Credit struct {
	Name string `xml:"-"`
	Page int `xml:"page,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	CreditType string `xml:"credit-type,omitempty"`
	Link []*common.Link `xml:"link,omitempty"`
	Bookmark []*common.Bookmark `xml:"bookmark,omitempty"`
	CreditImage *direction.Image `xml:"credit-image,omitempty"`
	CreditWords []*common.FormattedTextId `xml:"credit-words,omitempty"`
	CreditSymbol []*common.FormattedSymbolId `xml:"credit-symbol,omitempty"`
}

// Defaults: The defaults type specifies score-wide defaults for scaling; whether
// or not the file is a concert score; layout; and default values for the music font,
// word font, lyric font, and lyric language. Except for the concert-score element, if
// any defaults are missing, the choice of what to use is determined by the
// application.
type Defaults struct {
	Name string `xml:"-"`
	Scaling *common.Scaling `xml:"scaling,omitempty"`
	ConcertScore *common.Empty `xml:"concert-score,omitempty"`
	PageLayout *common.PageLayout `xml:"page-layout,omitempty"`
	SystemLayout *common.SystemLayout `xml:"system-layout,omitempty"`
	StaffLayout []*common.StaffLayout `xml:"staff-layout,omitempty"`
	Appearance *common.Appearance `xml:"appearance,omitempty"`
	MusicFont *EmptyFont `xml:"music-font,omitempty"`
	WordFont *EmptyFont `xml:"word-font,omitempty"`
	LyricFont []*LyricFont `xml:"lyric-font,omitempty"`
	LyricLanguage []*LyricLanguage `xml:"lyric-language,omitempty"`
}

// EmptyFont: The empty-font type represents an empty element with font attributes.
type EmptyFont struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
}

// GroupBarline: The group-barline type indicates if the group should have common
// barlines.
type GroupBarline struct {
	Name string `xml:"-"`
	Color string `xml:"color,attr,omitempty"`
	Value EnumGroupBarlineValue `xml:",chardata"`
}

// GroupName: The group-name type describes the name or abbreviation of a part-group
// element. Formatting attributes in the group-name type are deprecated in Version 2.0
// in favor of the new group-name-display and group-abbreviation-display elements.
type GroupName struct {
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
	Justify common.EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	Value string `xml:",chardata"`
}

// GroupSymbol: The group-symbol type indicates how the symbol for a group is
// indicated in the score. It is none if not specified.
type GroupSymbol struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value attributes.EnumGroupSymbolValue `xml:",chardata"`
}

// InstrumentLink: Multiple part-link elements can link a condensed part within a score
// file to multiple MusicXML parts files. For example, a "Clarinet 1 and 2" part in a
// score file could link to separate "Clarinet 1" and "Clarinet 2" part files. The
// instrument-link type distinguish which of the score-instruments within a score-part
// are in which part file. The instrument-link id attribute refers to a
// score-instrument id attribute.
type InstrumentLink struct {
	Name string `xml:"-"`
	RefScoreInstrument *ScoreInstrument `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
}

// LyricFont: The lyric-font type specifies the default font for a particular name
// and number of lyric.
type LyricFont struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	NameAttr string `xml:"name,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
}

// LyricLanguage: The lyric-language type specifies the default language for a
// particular name and number of lyric.
type LyricLanguage struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	NameAttr string `xml:"name,attr,omitempty"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Opus: The opus type represents a link to a MusicXML opus document that
// composes multiple MusicXML scores into a collection.
type Opus struct {
	Name string `xml:"-"`
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Type xlink.EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show xlink.EnumShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate xlink.EnumActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

// PartGroup: The part-group element indicates groupings of parts in the score,
// usually indicated by braces and brackets. Braces that are used for multi-staff parts
// should be defined in the attributes element for that part. The part-group start
// element appears before the first score-part in the group. The part-group stop
// element appears after the last score-part in the group. The number attribute is used
// to distinguish overlapping and nested part-groups, not the sequence of groups. As
// with parts, groups can have a name and abbreviation. Values for the child elements
// are ignored at the stop of a group. A part-group element is not needed for a single
// multi-staff part. By default, multi-staff parts include a brace symbol and (if
// appropriate given the bar-style) common barlines. The symbol formatting for a
// multi-staff part can be more fully specified using the part-symbol element.
type PartGroup struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Number string `xml:"number,attr,omitempty"`
	GroupName *GroupName `xml:"group-name,omitempty"`
	GroupNameDisplay *common.NameDisplay `xml:"group-name-display,omitempty"`
	GroupAbbreviation *GroupName `xml:"group-abbreviation,omitempty"`
	GroupAbbreviationDisplay *common.NameDisplay `xml:"group-abbreviation-display,omitempty"`
	GroupSymbol *GroupSymbol `xml:"group-symbol,omitempty"`
	GroupBarline *GroupBarline `xml:"group-barline,omitempty"`
	GroupTime *common.Empty `xml:"group-time,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
}

// PartLink: The part-link type allows MusicXML data for both score and parts to be
// contained within a single compressed MusicXML file. It links a score-part from a
// score document to MusicXML documents that contain parts data. In the case of a
// single compressed MusicXML file, the link href values are paths that are relative to
// the root folder of the zip file.
type PartLink struct {
	Name string `xml:"-"`
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Type xlink.EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show xlink.EnumShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate xlink.EnumActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	InstrumentLink []*InstrumentLink `xml:"instrument-link,omitempty"`
	GroupLink string `xml:"group-link,omitempty"`
}

// PartList: The part-list identifies the different musical parts in this document.
// Each part has an ID that is used later within the musical data. Since parts may be
// encoded separately and combined later, identification elements are present at both
// the score and score-part levels. There must be at least one score-part, combined as
// desired with part-group elements that indicate braces and brackets. Parts are
// ordered from top to bottom in a score based on the order in which they appear in the
// part-list.
type PartList struct {
	Name string `xml:"-"`
	PartGroup []*PartGroup `xml:"part-group,omitempty"`
	ScorePart []*ScorePart `xml:"score-part,omitempty"`
}

// PartName: The part-name type describes the name or abbreviation of a score-part
// element. Formatting attributes for the part-name element are deprecated in Version
// 2.0 in favor of the new part-name-display and part-abbreviation-display elements.
type PartName struct {
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
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Justify common.EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Player: The player type allows for multiple players per score-part for use in
// listening applications. One player may play multiple instruments, while a single
// instrument may include multiple players in divisi sections.
type Player struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	PlayerName string `xml:"player-name,omitempty"`
}

// ScoreInstrument: The score-instrument type represents a single instrument within a
// score-part. As with the score-part type, each score-instrument has a required ID
// attribute, a name, and an optional abbreviation. A score-instrument type is also
// required if the score specifies MIDI 1.0 channels, banks, or programs. An initial
// midi-instrument assignment can also be made here. MusicXML software should be able
// to automatically assign reasonable channels and instruments without these elements
// in simple cases, such as where part names match General MIDI instrument names. The
// score-instrument element can also distinguish multiple instruments of the same type
// that are on the same part, such as Clarinet 1 and Clarinet 2 instruments within a
// Clarinets 1 and 2 part.
type ScoreInstrument struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	InstrumentName string `xml:"instrument-name,omitempty"`
	InstrumentAbbreviation string `xml:"instrument-abbreviation,omitempty"`
	InstrumentSound string `xml:"instrument-sound,omitempty"`
	Solo *common.Empty `xml:"solo,omitempty"`
	Ensemble string `xml:"ensemble,omitempty"`
	VirtualInstrument *direction.VirtualInstrument `xml:"virtual-instrument,omitempty"`
}

// ScorePart: The score-part type collects part-wide information for each part in a
// score. Often, each MusicXML part corresponds to a track in a Standard MIDI Format 1
// file. In this case, the midi-device element is used to make a MIDI device or port
// assignment for the given track or specific MIDI instruments. Initial midi-instrument
// assignments may be made here as well. The score-instrument elements are used when
// there are multiple instruments per track.
type ScorePart struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Identification *common.Identification `xml:"identification,omitempty"`
	PartLink []*PartLink `xml:"part-link,omitempty"`
	PartName *PartName `xml:"part-name,omitempty"`
	PartNameDisplay *common.NameDisplay `xml:"part-name-display,omitempty"`
	PartAbbreviation *PartName `xml:"part-abbreviation,omitempty"`
	PartAbbreviationDisplay *common.NameDisplay `xml:"part-abbreviation-display,omitempty"`
	Group string `xml:"group,omitempty"`
	ScoreInstrument []*ScoreInstrument `xml:"score-instrument,omitempty"`
	Player []*Player `xml:"player,omitempty"`
	MidiDevice []*common.MidiDevice `xml:"midi-device,omitempty"`
	MidiInstrument []*common.MidiInstrument `xml:"midi-instrument,omitempty"`
}

// Work: Works are optionally identified by number and title. The work type
// also may indicate a link to the opus document that composes multiple scores into a
// collection.
type Work struct {
	Name string `xml:"-"`
	WorkNumber string `xml:"work-number,omitempty"`
	WorkTitle string `xml:"work-title,omitempty"`
	Opus *Opus `xml:"opus,omitempty"`
}

// ScorePartwise: The score-partwise element is the root element for a partwise MusicXML
// score. It includes a score-header group followed by a series of parts with measures
// inside. The document-attributes attribute group includes the version attribute.
type ScorePartwise struct {
	Name string `xml:"-"`
	XMLName xml.Name `xml:"score-partwise"`
	Version string `xml:"version,attr,omitempty"`
	Work *Work `xml:"work,omitempty"`
	MovementNumber string `xml:"movement-number,omitempty"`
	MovementTitle string `xml:"movement-title,omitempty"`
	Identification *common.Identification `xml:"identification,omitempty"`
	Defaults *Defaults `xml:"defaults,omitempty"`
	Credit []*Credit `xml:"credit,omitempty"`
	PartList *PartList `xml:"part-list,omitempty"`
	Part []*Part `xml:"part,omitempty"`
}

// ScoreTimewise: The score-timewise element is the root element for a timewise MusicXML
// score. It includes a score-header group followed by a series of measures with parts
// inside. The document-attributes attribute group includes the version attribute.
type ScoreTimewise struct {
	Name string `xml:"-"`
	XMLName xml.Name `xml:"score-timewise"`
	Version string `xml:"version,attr,omitempty"`
	Work *Work `xml:"work,omitempty"`
	MovementNumber string `xml:"movement-number,omitempty"`
	MovementTitle string `xml:"movement-title,omitempty"`
	Identification *common.Identification `xml:"identification,omitempty"`
	Defaults *Defaults `xml:"defaults,omitempty"`
	Credit []*Credit `xml:"credit,omitempty"`
	PartList *PartList `xml:"part-list,omitempty"`
	Measure []*ScoreTimewiseMeasure `xml:"measure,omitempty"`
}

type Part struct {
	Name string `xml:"-"`
	RefScorePart *ScorePart `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Measure []*Measure `xml:"measure,omitempty"`
}

type Measure struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	Text string `xml:"text,attr,omitempty"`
	Implicit common.EnumYesNo `xml:"implicit,attr,omitempty"`
	NonControlling common.EnumYesNo `xml:"non-controlling,attr,omitempty"`
	Width string `xml:"width,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	MusicData []*MeasureMusicDataChoice `xml:"-"`
}

// MeasureMusicDataChoice: Order-preserving xs:choice wrapper for Measure
type MeasureMusicDataChoice struct {
	Name string `xml:"-"`
	Note *note.Note `xml:"note,omitempty"`
	Backup *note.Backup `xml:"backup,omitempty"`
	Forward *note.Forward `xml:"forward,omitempty"`
	Direction *direction.Direction `xml:"direction,omitempty"`
	Attributes *attributes.Attributes `xml:"attributes,omitempty"`
	Harmony *direction.Harmony `xml:"harmony,omitempty"`
	FiguredBass *note.FiguredBass `xml:"figured-bass,omitempty"`
	Print *direction.Print `xml:"print,omitempty"`
	Sound *direction.Sound `xml:"sound,omitempty"`
	Listening *direction.Listening `xml:"listening,omitempty"`
	Barline *barline.Barline `xml:"barline,omitempty"`
	Grouping *direction.Grouping `xml:"grouping,omitempty"`
	Link *common.Link `xml:"link,omitempty"`
	Bookmark *common.Bookmark `xml:"bookmark,omitempty"`
}

type ScoreTimewiseMeasure struct {
	Name string `xml:"-"`
	Number string `xml:"number,attr,omitempty"`
	Text string `xml:"text,attr,omitempty"`
	Implicit common.EnumYesNo `xml:"implicit,attr,omitempty"`
	NonControlling common.EnumYesNo `xml:"non-controlling,attr,omitempty"`
	Width string `xml:"width,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Part []*ScoreTimewiseMeasurePart `xml:"part,omitempty"`
}

type ScoreTimewiseMeasurePart struct {
	Name string `xml:"-"`
	RefScorePart *ScorePart `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Note []*note.Note `xml:"note,omitempty"`
	Backup []*note.Backup `xml:"backup,omitempty"`
	Forward []*note.Forward `xml:"forward,omitempty"`
	Direction []*direction.Direction `xml:"direction,omitempty"`
	Attributes []*attributes.Attributes `xml:"attributes,omitempty"`
	Harmony []*direction.Harmony `xml:"harmony,omitempty"`
	FiguredBass []*note.FiguredBass `xml:"figured-bass,omitempty"`
	Print []*direction.Print `xml:"print,omitempty"`
	Sound []*direction.Sound `xml:"sound,omitempty"`
	Listening []*direction.Listening `xml:"listening,omitempty"`
	Barline []*barline.Barline `xml:"barline,omitempty"`
	Grouping []*direction.Grouping `xml:"grouping,omitempty"`
	Link []*common.Link `xml:"link,omitempty"`
	Bookmark []*common.Bookmark `xml:"bookmark,omitempty"`
}

