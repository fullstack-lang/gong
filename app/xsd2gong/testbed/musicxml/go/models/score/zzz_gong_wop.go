// generated code - do not edit
package score

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Credit_WOP struct {
	// insertion point

	Name string

	Page int

	Id string

	CreditType string
}

func (from *Credit) GongCopyBasicFields(to *Credit) {
	// insertion point
	*to = *from
}

type Defaults_WOP struct {
	// insertion point

	Name string
}

func (from *Defaults) GongCopyBasicFields(to *Defaults) {
	// insertion point
	to.Name = from.Name
}

type EmptyFont_WOP struct {
	// insertion point

	Name string

	FontFamily string

	FontSize string
}

func (from *EmptyFont) GongCopyBasicFields(to *EmptyFont) {
	// insertion point
	*to = *from
}

type GroupBarline_WOP struct {
	// insertion point

	Name string

	Color string

	Value EnumGroupBarlineValue
}

func (from *GroupBarline) GongCopyBasicFields(to *GroupBarline) {
	// insertion point
	*to = *from
}

type GroupName_WOP struct {
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

func (from *GroupName) GongCopyBasicFields(to *GroupName) {
	// insertion point
	*to = *from
}

type GroupSymbol_WOP struct {
	// insertion point

	Name string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	Color string
}

func (from *GroupSymbol) GongCopyBasicFields(to *GroupSymbol) {
	// insertion point
	*to = *from
}

type InstrumentLink_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *InstrumentLink) GongCopyBasicFields(to *InstrumentLink) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type LyricFont_WOP struct {
	// insertion point

	Name string

	Number string

	NameAttr string

	FontFamily string

	FontSize string
}

func (from *LyricFont) GongCopyBasicFields(to *LyricFont) {
	// insertion point
	*to = *from
}

type LyricLanguage_WOP struct {
	// insertion point

	Name string

	Number string

	NameAttr string

	Lang string
}

func (from *LyricLanguage) GongCopyBasicFields(to *LyricLanguage) {
	// insertion point
	*to = *from
}

type Measure_WOP struct {
	// insertion point

	Name string

	Number string

	Text string

	Width string

	Id string
}

func (from *Measure) GongCopyBasicFields(to *Measure) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Text = from.Text
	to.Width = from.Width
	to.Id = from.Id
}

type MeasureMusicDataChoice_WOP struct {
	// insertion point

	Name string
}

func (from *MeasureMusicDataChoice) GongCopyBasicFields(to *MeasureMusicDataChoice) {
	// insertion point
	*to = *from
}

type Opus_WOP struct {
	// insertion point

	Name string

	Href string

	Role string

	Title string
}

func (from *Opus) GongCopyBasicFields(to *Opus) {
	// insertion point
	*to = *from
}

type Part_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Part) GongCopyBasicFields(to *Part) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type PartGroup_WOP struct {
	// insertion point

	Name string

	Number string
}

func (from *PartGroup) GongCopyBasicFields(to *PartGroup) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
}

type PartLink_WOP struct {
	// insertion point

	Name string

	Href string

	Role string

	Title string

	GroupLink string
}

func (from *PartLink) GongCopyBasicFields(to *PartLink) {
	// insertion point
	to.Name = from.Name
	to.Href = from.Href
	to.Role = from.Role
	to.Title = from.Title
	to.GroupLink = from.GroupLink
}

type PartList_WOP struct {
	// insertion point

	Name string
}

func (from *PartList) GongCopyBasicFields(to *PartList) {
	// insertion point
	to.Name = from.Name
}

type PartName_WOP struct {
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

func (from *PartName) GongCopyBasicFields(to *PartName) {
	// insertion point
	*to = *from
}

type Player_WOP struct {
	// insertion point

	Name string

	Id string

	PlayerName string
}

func (from *Player) GongCopyBasicFields(to *Player) {
	// insertion point
	*to = *from
}

type ScoreInstrument_WOP struct {
	// insertion point

	Name string

	Id string

	InstrumentName string

	InstrumentAbbreviation string

	InstrumentSound string

	Ensemble string
}

func (from *ScoreInstrument) GongCopyBasicFields(to *ScoreInstrument) {
	// insertion point
	*to = *from
}

type ScorePart_WOP struct {
	// insertion point

	Name string

	Id string

	Group string
}

func (from *ScorePart) GongCopyBasicFields(to *ScorePart) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Group = from.Group
}

type ScorePartwise_WOP struct {
	// insertion point

	Name string

	Version string

	MovementNumber string

	MovementTitle string
}

func (from *ScorePartwise) GongCopyBasicFields(to *ScorePartwise) {
	// insertion point
	to.Name = from.Name
	to.Version = from.Version
	to.MovementNumber = from.MovementNumber
	to.MovementTitle = from.MovementTitle
}

type ScoreTimewise_WOP struct {
	// insertion point

	Name string

	Version string

	MovementNumber string

	MovementTitle string
}

func (from *ScoreTimewise) GongCopyBasicFields(to *ScoreTimewise) {
	// insertion point
	to.Name = from.Name
	to.Version = from.Version
	to.MovementNumber = from.MovementNumber
	to.MovementTitle = from.MovementTitle
}

type ScoreTimewiseMeasure_WOP struct {
	// insertion point

	Name string

	Number string

	Text string

	Width string

	Id string
}

func (from *ScoreTimewiseMeasure) GongCopyBasicFields(to *ScoreTimewiseMeasure) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Text = from.Text
	to.Width = from.Width
	to.Id = from.Id
}

type ScoreTimewiseMeasurePart_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *ScoreTimewiseMeasurePart) GongCopyBasicFields(to *ScoreTimewiseMeasurePart) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Work_WOP struct {
	// insertion point

	Name string

	WorkNumber string

	WorkTitle string
}

func (from *Work) GongCopyBasicFields(to *Work) {
	// insertion point
	to.Name = from.Name
	to.WorkNumber = from.WorkNumber
	to.WorkTitle = from.WorkTitle
}

// end of insertion point
