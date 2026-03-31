// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type A_directive_WOP struct {
	// insertion point

	Name string

	Lang string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *A_directive) CopyBasicFields(to *A_directive) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type A_measure_WOP struct {
	// insertion point

	Name string

	Number string

	Text string

	Implicit Enum_Yes_no

	Non_controlling Enum_Yes_no

	Width string

	Id string
}

func (from *A_measure) CopyBasicFields(to *A_measure) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Text = from.Text
	to.Implicit = from.Implicit
	to.Non_controlling = from.Non_controlling
	to.Width = from.Width
	to.Id = from.Id
}

type A_measure_1_WOP struct {
	// insertion point

	Name string

	Number string

	Text string

	Implicit Enum_Yes_no

	Non_controlling Enum_Yes_no

	Width string

	Id string
}

func (from *A_measure_1) CopyBasicFields(to *A_measure_1) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Text = from.Text
	to.Implicit = from.Implicit
	to.Non_controlling = from.Non_controlling
	to.Width = from.Width
	to.Id = from.Id
}

type A_part_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *A_part) CopyBasicFields(to *A_part) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type A_part_1_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *A_part_1) CopyBasicFields(to *A_part_1) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Accidental_WOP struct {
	// insertion point

	Name string

	Cautionary string

	Editorial Enum_Yes_no

	Smufl string

	Parentheses Enum_Yes_no

	Bracket Enum_Yes_no

	Size Enum_Symbol_size

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Accidental) CopyBasicFields(to *Accidental) {
	// insertion point
	to.Name = from.Name
	to.Cautionary = from.Cautionary
	to.Editorial = from.Editorial
	to.Smufl = from.Smufl
	to.Parentheses = from.Parentheses
	to.Bracket = from.Bracket
	to.Size = from.Size
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Accidental_mark_WOP struct {
	// insertion point

	Name string

	Smufl string

	Parentheses Enum_Yes_no

	Bracket Enum_Yes_no

	Size Enum_Symbol_size

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Id string

	EnclosedText Enum_Accidental_value
}

func (from *Accidental_mark) CopyBasicFields(to *Accidental_mark) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.Parentheses = from.Parentheses
	to.Bracket = from.Bracket
	to.Size = from.Size
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Accidental_text_WOP struct {
	// insertion point

	Name string

	Smufl string

	Lang string

	Space string

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Line_height string

	Dir string

	Enclosure string

	EnclosedText Enum_Accidental_value
}

func (from *Accidental_text) CopyBasicFields(to *Accidental_text) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.Lang = from.Lang
	to.Space = from.Space
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Line_height = from.Line_height
	to.Dir = from.Dir
	to.Enclosure = from.Enclosure
	to.EnclosedText = from.EnclosedText
}

type Accord_WOP struct {
	// insertion point

	Name string

	String int

	Tuning_step Enum_Step

	Tuning_alter string

	Tuning_octave int
}

func (from *Accord) CopyBasicFields(to *Accord) {
	// insertion point
	to.Name = from.Name
	to.String = from.String
	to.Tuning_step = from.Tuning_step
	to.Tuning_alter = from.Tuning_alter
	to.Tuning_octave = from.Tuning_octave
}

type Accordion_registration_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string

	Accordion_high string

	Accordion_middle int

	Accordion_low string
}

func (from *Accordion_registration) CopyBasicFields(to *Accordion_registration) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
	to.Accordion_high = from.Accordion_high
	to.Accordion_middle = from.Accordion_middle
	to.Accordion_low = from.Accordion_low
}

type Appearance_WOP struct {
	// insertion point

	Name string
}

func (from *Appearance) CopyBasicFields(to *Appearance) {
	// insertion point
	to.Name = from.Name
}

type Arpeggiate_WOP struct {
	// insertion point

	Name string

	Number int

	Direction string

	Unbroken Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Color string

	Id string
}

func (from *Arpeggiate) CopyBasicFields(to *Arpeggiate) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Direction = from.Direction
	to.Unbroken = from.Unbroken
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Color = from.Color
	to.Id = from.Id
}

type Arrow_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Smufl string

	Arrow_direction string

	Arrow_style string

	Arrowhead string

	Circular_arrow string
}

func (from *Arrow) CopyBasicFields(to *Arrow) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Smufl = from.Smufl
	to.Arrow_direction = from.Arrow_direction
	to.Arrow_style = from.Arrow_style
	to.Arrowhead = from.Arrowhead
	to.Circular_arrow = from.Circular_arrow
}

type Articulations_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Articulations) CopyBasicFields(to *Articulations) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Assess_WOP struct {
	// insertion point

	Name string

	Type Enum_Yes_no

	Player string

	Time_only string
}

func (from *Assess) CopyBasicFields(to *Assess) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Player = from.Player
	to.Time_only = from.Time_only
}

type Attributes_WOP struct {
	// insertion point

	Name string

	Divisions string

	Staves int

	Instruments int
}

func (from *Attributes) CopyBasicFields(to *Attributes) {
	// insertion point
	to.Name = from.Name
	to.Divisions = from.Divisions
	to.Staves = from.Staves
	to.Instruments = from.Instruments
}

type Backup_WOP struct {
	// insertion point

	Name string

	Duration string
}

func (from *Backup) CopyBasicFields(to *Backup) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
}

type Bar_style_color_WOP struct {
	// insertion point

	Name string

	Color string

	EnclosedText string
}

func (from *Bar_style_color) CopyBasicFields(to *Bar_style_color) {
	// insertion point
	to.Name = from.Name
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Barline_WOP struct {
	// insertion point

	Name string

	Location string

	Segno string

	Coda string

	Divisions string

	Id string
}

func (from *Barline) CopyBasicFields(to *Barline) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.Segno = from.Segno
	to.Coda = from.Coda
	to.Divisions = from.Divisions
	to.Id = from.Id
}

type Barre_WOP struct {
	// insertion point

	Name string

	Type string

	Color string
}

func (from *Barre) CopyBasicFields(to *Barre) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Color = from.Color
}

type Bass_WOP struct {
	// insertion point

	Name string

	Arrangement string
}

func (from *Bass) CopyBasicFields(to *Bass) {
	// insertion point
	to.Name = from.Name
	to.Arrangement = from.Arrangement
}

type Bass_step_WOP struct {
	// insertion point

	Name string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Bass_step) CopyBasicFields(to *Bass_step) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Beam_WOP struct {
	// insertion point

	Name string

	Number int

	Repeater Enum_Yes_no

	Fan string

	Color string

	Id string

	EnclosedText string
}

func (from *Beam) CopyBasicFields(to *Beam) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Repeater = from.Repeater
	to.Fan = from.Fan
	to.Color = from.Color
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Beat_repeat_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Slashes int

	Use_dots Enum_Yes_no

	Slash_type Enum_Note_type_value

	Slash_dot string

	Except_voice string
}

func (from *Beat_repeat) CopyBasicFields(to *Beat_repeat) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Slashes = from.Slashes
	to.Use_dots = from.Use_dots
	to.Slash_type = from.Slash_type
	to.Slash_dot = from.Slash_dot
	to.Except_voice = from.Except_voice
}

type Beat_unit_tied_WOP struct {
	// insertion point

	Name string

	Beat_unit Enum_Note_type_value

	Beat_unit_dot string
}

func (from *Beat_unit_tied) CopyBasicFields(to *Beat_unit_tied) {
	// insertion point
	to.Name = from.Name
	to.Beat_unit = from.Beat_unit
	to.Beat_unit_dot = from.Beat_unit_dot
}

type Beater_WOP struct {
	// insertion point

	Name string

	Tip string

	EnclosedText string
}

func (from *Beater) CopyBasicFields(to *Beater) {
	// insertion point
	to.Name = from.Name
	to.Tip = from.Tip
	to.EnclosedText = from.EnclosedText
}

type Bend_WOP struct {
	// insertion point

	Name string

	Shape string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Accelerate Enum_Yes_no

	Beats string

	First_beat string

	Last_beat string

	Bend_alter string

	Pre_bend string
}

func (from *Bend) CopyBasicFields(to *Bend) {
	// insertion point
	to.Name = from.Name
	to.Shape = from.Shape
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Accelerate = from.Accelerate
	to.Beats = from.Beats
	to.First_beat = from.First_beat
	to.Last_beat = from.Last_beat
	to.Bend_alter = from.Bend_alter
	to.Pre_bend = from.Pre_bend
}

type Bookmark_WOP struct {
	// insertion point

	Name string

	Id string

	NameXSD string

	Element string

	Position int
}

func (from *Bookmark) CopyBasicFields(to *Bookmark) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.NameXSD = from.NameXSD
	to.Element = from.Element
	to.Position = from.Position
}

type Bracket_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Line_end string

	End_length string

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	Id string
}

func (from *Bracket) CopyBasicFields(to *Bracket) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line_end = from.Line_end
	to.End_length = from.End_length
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.Id = from.Id
}

type Breath_mark_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Breath_mark) CopyBasicFields(to *Breath_mark) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Caesura_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Caesura) CopyBasicFields(to *Caesura) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Cancel_WOP struct {
	// insertion point

	Name string

	Location string

	EnclosedText int
}

func (from *Cancel) CopyBasicFields(to *Cancel) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.EnclosedText = from.EnclosedText
}

type Clef_WOP struct {
	// insertion point

	Name string

	Number int

	Additional Enum_Yes_no

	Size string

	After_barline Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Print_object Enum_Yes_no

	Id string

	Sign string

	Line int

	Clef_octave_change int
}

func (from *Clef) CopyBasicFields(to *Clef) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Additional = from.Additional
	to.Size = from.Size
	to.After_barline = from.After_barline
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Print_object = from.Print_object
	to.Id = from.Id
	to.Sign = from.Sign
	to.Line = from.Line
	to.Clef_octave_change = from.Clef_octave_change
}

type Coda_WOP struct {
	// insertion point

	Name string

	Smufl string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Coda) CopyBasicFields(to *Coda) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Credit_WOP struct {
	// insertion point

	Name string

	Page int

	Id string

	Credit_type string
}

func (from *Credit) CopyBasicFields(to *Credit) {
	// insertion point
	to.Name = from.Name
	to.Page = from.Page
	to.Id = from.Id
	to.Credit_type = from.Credit_type
}

type Dashes_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop_continue

	Number int

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	Id string
}

func (from *Dashes) CopyBasicFields(to *Dashes) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.Id = from.Id
}

type Defaults_WOP struct {
	// insertion point

	Name string

	Concert_score string
}

func (from *Defaults) CopyBasicFields(to *Defaults) {
	// insertion point
	to.Name = from.Name
	to.Concert_score = from.Concert_score
}

type Degree_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no
}

func (from *Degree) CopyBasicFields(to *Degree) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
}

type Degree_alter_WOP struct {
	// insertion point

	Name string

	Plus_minus Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Degree_alter) CopyBasicFields(to *Degree_alter) {
	// insertion point
	to.Name = from.Name
	to.Plus_minus = from.Plus_minus
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Degree_type_WOP struct {
	// insertion point

	Name string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Degree_type) CopyBasicFields(to *Degree_type) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Degree_value_WOP struct {
	// insertion point

	Name string

	Symbol string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText int
}

func (from *Degree_value) CopyBasicFields(to *Degree_value) {
	// insertion point
	to.Name = from.Name
	to.Symbol = from.Symbol
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Direction_WOP struct {
	// insertion point

	Name string

	Placement string

	Directive Enum_Yes_no

	System Enum_System_relation_number

	Id string

	Voice string

	Staff int
}

func (from *Direction) CopyBasicFields(to *Direction) {
	// insertion point
	to.Name = from.Name
	to.Placement = from.Placement
	to.Directive = from.Directive
	to.System = from.System
	to.Id = from.Id
	to.Voice = from.Voice
	to.Staff = from.Staff
}

type Direction_type_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Direction_type) CopyBasicFields(to *Direction_type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Distance_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Distance) CopyBasicFields(to *Distance) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Double_WOP struct {
	// insertion point

	Name string

	Above Enum_Yes_no
}

func (from *Double) CopyBasicFields(to *Double) {
	// insertion point
	to.Name = from.Name
	to.Above = from.Above
}

type Dynamics_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Placement string

	Underline int

	Overline int

	Line_through int

	Enclosure string

	Id string

	P string

	Pp string

	Ppp string

	Pppp string

	Ppppp string

	Pppppp string

	F string

	Ff string

	Fff string

	Ffff string

	Fffff string

	Ffffff string

	Mp string

	Mf string

	Sf string

	Sfp string

	Sfpp string

	Fp string

	Rf string

	Rfz string

	Sfz string

	Sffz string

	Fz string

	N string

	Pf string

	Sfzp string
}

func (from *Dynamics) CopyBasicFields(to *Dynamics) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Placement = from.Placement
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Enclosure = from.Enclosure
	to.Id = from.Id
	to.P = from.P
	to.Pp = from.Pp
	to.Ppp = from.Ppp
	to.Pppp = from.Pppp
	to.Ppppp = from.Ppppp
	to.Pppppp = from.Pppppp
	to.F = from.F
	to.Ff = from.Ff
	to.Fff = from.Fff
	to.Ffff = from.Ffff
	to.Fffff = from.Fffff
	to.Ffffff = from.Ffffff
	to.Mp = from.Mp
	to.Mf = from.Mf
	to.Sf = from.Sf
	to.Sfp = from.Sfp
	to.Sfpp = from.Sfpp
	to.Fp = from.Fp
	to.Rf = from.Rf
	to.Rfz = from.Rfz
	to.Sfz = from.Sfz
	to.Sffz = from.Sffz
	to.Fz = from.Fz
	to.N = from.N
	to.Pf = from.Pf
	to.Sfzp = from.Sfzp
}

type Effect_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Effect) CopyBasicFields(to *Effect) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Elision_WOP struct {
	// insertion point

	Name string

	Smufl string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Elision) CopyBasicFields(to *Elision) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Empty_WOP struct {
	// insertion point

	Name string
}

func (from *Empty) CopyBasicFields(to *Empty) {
	// insertion point
	to.Name = from.Name
}

type Empty_font_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string
}

func (from *Empty_font) CopyBasicFields(to *Empty_font) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
}

type Empty_line_WOP struct {
	// insertion point

	Name string

	Line_shape string

	Line_type string

	Line_length string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string
}

func (from *Empty_line) CopyBasicFields(to *Empty_line) {
	// insertion point
	to.Name = from.Name
	to.Line_shape = from.Line_shape
	to.Line_type = from.Line_type
	to.Line_length = from.Line_length
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
}

type Empty_placement_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string
}

func (from *Empty_placement) CopyBasicFields(to *Empty_placement) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
}

type Empty_placement_smufl_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Smufl string
}

func (from *Empty_placement_smufl) CopyBasicFields(to *Empty_placement_smufl) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Smufl = from.Smufl
}

type Empty_print_object_style_align_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string
}

func (from *Empty_print_object_style_align) CopyBasicFields(to *Empty_print_object_style_align) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
}

type Empty_print_style_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string
}

func (from *Empty_print_style) CopyBasicFields(to *Empty_print_style) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
}

type Empty_print_style_align_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string
}

func (from *Empty_print_style_align) CopyBasicFields(to *Empty_print_style_align) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
}

type Empty_print_style_align_id_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Empty_print_style_align_id) CopyBasicFields(to *Empty_print_style_align_id) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Empty_trill_sound_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Start_note string

	Trill_step string

	Two_note_turn string

	Accelerate Enum_Yes_no

	Beats string

	Second_beat string

	Last_beat string
}

func (from *Empty_trill_sound) CopyBasicFields(to *Empty_trill_sound) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Start_note = from.Start_note
	to.Trill_step = from.Trill_step
	to.Two_note_turn = from.Two_note_turn
	to.Accelerate = from.Accelerate
	to.Beats = from.Beats
	to.Second_beat = from.Second_beat
	to.Last_beat = from.Last_beat
}

type Encoding_WOP struct {
	// insertion point

	Name string

	Software string

	Encoding_description string
}

func (from *Encoding) CopyBasicFields(to *Encoding) {
	// insertion point
	to.Name = from.Name
	to.Software = from.Software
	to.Encoding_description = from.Encoding_description
}

type Ending_WOP struct {
	// insertion point

	Name string

	Number string

	Type string

	End_length string

	Text_x string

	Text_y string

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	System Enum_System_relation_number

	EnclosedText string
}

func (from *Ending) CopyBasicFields(to *Ending) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Type = from.Type
	to.End_length = from.End_length
	to.Text_x = from.Text_x
	to.Text_y = from.Text_y
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.System = from.System
	to.EnclosedText = from.EnclosedText
}

type Extend_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop_continue

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string
}

func (from *Extend) CopyBasicFields(to *Extend) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
}

type Feature_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Feature) CopyBasicFields(to *Feature) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Fermata_WOP struct {
	// insertion point

	Name string

	Type string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Id string

	EnclosedText string
}

func (from *Fermata) CopyBasicFields(to *Fermata) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Figure_WOP struct {
	// insertion point

	Name string
}

func (from *Figure) CopyBasicFields(to *Figure) {
	// insertion point
	to.Name = from.Name
}

type Figured_bass_WOP struct {
	// insertion point

	Name string

	Parentheses Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Placement string

	Print_dot Enum_Yes_no

	Print_lyric Enum_Yes_no

	Print_object Enum_Yes_no

	Print_spacing Enum_Yes_no

	Id string

	Duration string
}

func (from *Figured_bass) CopyBasicFields(to *Figured_bass) {
	// insertion point
	to.Name = from.Name
	to.Parentheses = from.Parentheses
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Placement = from.Placement
	to.Print_dot = from.Print_dot
	to.Print_lyric = from.Print_lyric
	to.Print_object = from.Print_object
	to.Print_spacing = from.Print_spacing
	to.Id = from.Id
	to.Duration = from.Duration
}

type Fingering_WOP struct {
	// insertion point

	Name string

	Substitution Enum_Yes_no

	Alternate Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Fingering) CopyBasicFields(to *Fingering) {
	// insertion point
	to.Name = from.Name
	to.Substitution = from.Substitution
	to.Alternate = from.Alternate
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type First_fret_WOP struct {
	// insertion point

	Name string

	Text string

	Location string

	EnclosedText int
}

func (from *First_fret) CopyBasicFields(to *First_fret) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Location = from.Location
	to.EnclosedText = from.EnclosedText
}

type For_part_WOP struct {
	// insertion point

	Name string

	Number int

	Id string
}

func (from *For_part) CopyBasicFields(to *For_part) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Id = from.Id
}

type Formatted_symbol_WOP struct {
	// insertion point

	Name string

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Line_height string

	Dir string

	Enclosure string

	EnclosedText string
}

func (from *Formatted_symbol) CopyBasicFields(to *Formatted_symbol) {
	// insertion point
	to.Name = from.Name
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Line_height = from.Line_height
	to.Dir = from.Dir
	to.Enclosure = from.Enclosure
	to.EnclosedText = from.EnclosedText
}

type Formatted_symbol_id_WOP struct {
	// insertion point

	Name string

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Line_height string

	Dir string

	Enclosure string

	Id string

	EnclosedText string
}

func (from *Formatted_symbol_id) CopyBasicFields(to *Formatted_symbol_id) {
	// insertion point
	to.Name = from.Name
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Line_height = from.Line_height
	to.Dir = from.Dir
	to.Enclosure = from.Enclosure
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Formatted_text_WOP struct {
	// insertion point

	Name string

	Lang string

	Space string

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Line_height string

	Dir string

	Enclosure string

	EnclosedText string
}

func (from *Formatted_text) CopyBasicFields(to *Formatted_text) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
	to.Space = from.Space
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Line_height = from.Line_height
	to.Dir = from.Dir
	to.Enclosure = from.Enclosure
	to.EnclosedText = from.EnclosedText
}

type Formatted_text_id_WOP struct {
	// insertion point

	Name string

	Lang string

	Space string

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Line_height string

	Dir string

	Enclosure string

	Id string

	EnclosedText string
}

func (from *Formatted_text_id) CopyBasicFields(to *Formatted_text_id) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
	to.Space = from.Space
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Line_height = from.Line_height
	to.Dir = from.Dir
	to.Enclosure = from.Enclosure
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Forward_WOP struct {
	// insertion point

	Name string

	Duration string

	Voice string

	Staff int
}

func (from *Forward) CopyBasicFields(to *Forward) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
	to.Voice = from.Voice
	to.Staff = from.Staff
}

type Frame_WOP struct {
	// insertion point

	Name string

	Height string

	Width string

	Unplayed string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	Halign string

	Valign string

	Id string

	Frame_strings int

	Frame_frets int
}

func (from *Frame) CopyBasicFields(to *Frame) {
	// insertion point
	to.Name = from.Name
	to.Height = from.Height
	to.Width = from.Width
	to.Unplayed = from.Unplayed
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
	to.Frame_strings = from.Frame_strings
	to.Frame_frets = from.Frame_frets
}

type Frame_note_WOP struct {
	// insertion point

	Name string
}

func (from *Frame_note) CopyBasicFields(to *Frame_note) {
	// insertion point
	to.Name = from.Name
}

type Fret_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText int
}

func (from *Fret) CopyBasicFields(to *Fret) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Glass_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Glass) CopyBasicFields(to *Glass) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Glissando_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Number int

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Id string

	EnclosedText string
}

func (from *Glissando) CopyBasicFields(to *Glissando) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Glyph_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Glyph) CopyBasicFields(to *Glyph) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Grace_WOP struct {
	// insertion point

	Name string

	Steal_time_previous string

	Steal_time_following string

	Make_time string

	Slash Enum_Yes_no
}

func (from *Grace) CopyBasicFields(to *Grace) {
	// insertion point
	to.Name = from.Name
	to.Steal_time_previous = from.Steal_time_previous
	to.Steal_time_following = from.Steal_time_following
	to.Make_time = from.Make_time
	to.Slash = from.Slash
}

type Group_barline_WOP struct {
	// insertion point

	Name string

	Color string

	EnclosedText string
}

func (from *Group_barline) CopyBasicFields(to *Group_barline) {
	// insertion point
	to.Name = from.Name
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Group_name_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Justify Enum_Left_center_right

	EnclosedText string
}

func (from *Group_name) CopyBasicFields(to *Group_name) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Justify = from.Justify
	to.EnclosedText = from.EnclosedText
}

type Group_symbol_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	EnclosedText string
}

func (from *Group_symbol) CopyBasicFields(to *Group_symbol) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Grouping_WOP struct {
	// insertion point

	Name string

	Type string

	Number string

	Member_of string

	Id string
}

func (from *Grouping) CopyBasicFields(to *Grouping) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Member_of = from.Member_of
	to.Id = from.Id
}

type Hammer_on_pull_off_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Number int

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Hammer_on_pull_off) CopyBasicFields(to *Hammer_on_pull_off) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Handbell_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Handbell) CopyBasicFields(to *Handbell) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Harmon_closed_WOP struct {
	// insertion point

	Name string

	Location string

	EnclosedText string
}

func (from *Harmon_closed) CopyBasicFields(to *Harmon_closed) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.EnclosedText = from.EnclosedText
}

type Harmon_mute_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string
}

func (from *Harmon_mute) CopyBasicFields(to *Harmon_mute) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
}

type Harmonic_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Natural string

	Artificial string

	Base_pitch string

	Touching_pitch string

	Sounding_pitch string
}

func (from *Harmonic) CopyBasicFields(to *Harmonic) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Natural = from.Natural
	to.Artificial = from.Artificial
	to.Base_pitch = from.Base_pitch
	to.Touching_pitch = from.Touching_pitch
	to.Sounding_pitch = from.Sounding_pitch
}

type Harmony_WOP struct {
	// insertion point

	Name string

	Type string

	Print_frame Enum_Yes_no

	Arrangement Enum_Harmony_arrangement

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	System Enum_System_relation_number

	Id string

	Staff int
}

func (from *Harmony) CopyBasicFields(to *Harmony) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Print_frame = from.Print_frame
	to.Arrangement = from.Arrangement
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.System = from.System
	to.Id = from.Id
	to.Staff = from.Staff
}

type Harmony_alter_WOP struct {
	// insertion point

	Name string

	Location Enum_Left_right

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Harmony_alter) CopyBasicFields(to *Harmony_alter) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Harp_pedals_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Harp_pedals) CopyBasicFields(to *Harp_pedals) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Heel_toe_WOP struct {
	// insertion point

	Name string
}

func (from *Heel_toe) CopyBasicFields(to *Heel_toe) {
	// insertion point
	to.Name = from.Name
}

type Hole_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Hole_type string

	Hole_shape string
}

func (from *Hole) CopyBasicFields(to *Hole) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Hole_type = from.Hole_type
	to.Hole_shape = from.Hole_shape
}

type Hole_closed_WOP struct {
	// insertion point

	Name string

	Location string

	EnclosedText string
}

func (from *Hole_closed) CopyBasicFields(to *Hole_closed) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.EnclosedText = from.EnclosedText
}

type Horizontal_turn_WOP struct {
	// insertion point

	Name string

	Slash Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Start_note string

	Trill_step string

	Two_note_turn string

	Accelerate Enum_Yes_no

	Beats string

	Second_beat string

	Last_beat string
}

func (from *Horizontal_turn) CopyBasicFields(to *Horizontal_turn) {
	// insertion point
	to.Name = from.Name
	to.Slash = from.Slash
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Start_note = from.Start_note
	to.Trill_step = from.Trill_step
	to.Two_note_turn = from.Two_note_turn
	to.Accelerate = from.Accelerate
	to.Beats = from.Beats
	to.Second_beat = from.Second_beat
	to.Last_beat = from.Last_beat
}

type Identification_WOP struct {
	// insertion point

	Name string

	Source string
}

func (from *Identification) CopyBasicFields(to *Identification) {
	// insertion point
	to.Name = from.Name
	to.Source = from.Source
}

type Image_WOP struct {
	// insertion point

	Name string

	Source string

	Type string

	Height string

	Width string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Halign string

	Valign string

	Id string
}

func (from *Image) CopyBasicFields(to *Image) {
	// insertion point
	to.Name = from.Name
	to.Source = from.Source
	to.Type = from.Type
	to.Height = from.Height
	to.Width = from.Width
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Instrument_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Instrument) CopyBasicFields(to *Instrument) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Instrument_change_WOP struct {
	// insertion point

	Name string

	Id string

	Instrument_sound string

	Solo string

	Ensemble string
}

func (from *Instrument_change) CopyBasicFields(to *Instrument_change) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Instrument_sound = from.Instrument_sound
	to.Solo = from.Solo
	to.Ensemble = from.Ensemble
}

type Instrument_link_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Instrument_link) CopyBasicFields(to *Instrument_link) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Interchangeable_WOP struct {
	// insertion point

	Name string

	Symbol string

	Separator string

	Time_relation string

	Beats string

	Beat_type string
}

func (from *Interchangeable) CopyBasicFields(to *Interchangeable) {
	// insertion point
	to.Name = from.Name
	to.Symbol = from.Symbol
	to.Separator = from.Separator
	to.Time_relation = from.Time_relation
	to.Beats = from.Beats
	to.Beat_type = from.Beat_type
}

type Inversion_WOP struct {
	// insertion point

	Name string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText int
}

func (from *Inversion) CopyBasicFields(to *Inversion) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Key_WOP struct {
	// insertion point

	Name string

	Number int

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Print_object Enum_Yes_no

	Id string

	Fifths int

	Mode string

	Key_step Enum_Step

	Key_alter string
}

func (from *Key) CopyBasicFields(to *Key) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Print_object = from.Print_object
	to.Id = from.Id
	to.Fifths = from.Fifths
	to.Mode = from.Mode
	to.Key_step = from.Key_step
	to.Key_alter = from.Key_alter
}

type Key_accidental_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText Enum_Accidental_value
}

func (from *Key_accidental) CopyBasicFields(to *Key_accidental) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Key_octave_WOP struct {
	// insertion point

	Name string

	Number int

	Cancel Enum_Yes_no

	EnclosedText int
}

func (from *Key_octave) CopyBasicFields(to *Key_octave) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Cancel = from.Cancel
	to.EnclosedText = from.EnclosedText
}

type Kind_WOP struct {
	// insertion point

	Name string

	Use_symbols Enum_Yes_no

	Text string

	Stack_degrees Enum_Yes_no

	Parentheses_degrees Enum_Yes_no

	Bracket_degrees Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	EnclosedText string
}

func (from *Kind) CopyBasicFields(to *Kind) {
	// insertion point
	to.Name = from.Name
	to.Use_symbols = from.Use_symbols
	to.Text = from.Text
	to.Stack_degrees = from.Stack_degrees
	to.Parentheses_degrees = from.Parentheses_degrees
	to.Bracket_degrees = from.Bracket_degrees
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.EnclosedText = from.EnclosedText
}

type Level_WOP struct {
	// insertion point

	Name string

	Reference Enum_Yes_no

	Type Enum_Start_stop_single

	Parentheses Enum_Yes_no

	Bracket Enum_Yes_no

	Size Enum_Symbol_size

	EnclosedText string
}

func (from *Level) CopyBasicFields(to *Level) {
	// insertion point
	to.Name = from.Name
	to.Reference = from.Reference
	to.Type = from.Type
	to.Parentheses = from.Parentheses
	to.Bracket = from.Bracket
	to.Size = from.Size
	to.EnclosedText = from.EnclosedText
}

type Line_detail_WOP struct {
	// insertion point

	Name string

	Line int

	Width string

	Color string

	Line_type string

	Print_object Enum_Yes_no
}

func (from *Line_detail) CopyBasicFields(to *Line_detail) {
	// insertion point
	to.Name = from.Name
	to.Line = from.Line
	to.Width = from.Width
	to.Color = from.Color
	to.Line_type = from.Line_type
	to.Print_object = from.Print_object
}

type Line_width_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Line_width) CopyBasicFields(to *Line_width) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Link_WOP struct {
	// insertion point

	Name string

	NameXSD string

	Href string

	Type string

	Role string

	Title string

	Show string

	Actuate string

	Element string

	Position int

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string
}

func (from *Link) CopyBasicFields(to *Link) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.Href = from.Href
	to.Type = from.Type
	to.Role = from.Role
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Element = from.Element
	to.Position = from.Position
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
}

type Listen_WOP struct {
	// insertion point

	Name string
}

func (from *Listen) CopyBasicFields(to *Listen) {
	// insertion point
	to.Name = from.Name
}

type Listening_WOP struct {
	// insertion point

	Name string
}

func (from *Listening) CopyBasicFields(to *Listening) {
	// insertion point
	to.Name = from.Name
}

type Lyric_WOP struct {
	// insertion point

	Name string

	Number string

	NameXSD string

	Time_only Enum_Time_only

	Justify Enum_Left_center_right

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Color string

	Print_object Enum_Yes_no

	Id string

	Syllabic string

	Laughing string

	Humming string

	End_line string

	End_paragraph string
}

func (from *Lyric) CopyBasicFields(to *Lyric) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.NameXSD = from.NameXSD
	to.Time_only = from.Time_only
	to.Justify = from.Justify
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Color = from.Color
	to.Print_object = from.Print_object
	to.Id = from.Id
	to.Syllabic = from.Syllabic
	to.Laughing = from.Laughing
	to.Humming = from.Humming
	to.End_line = from.End_line
	to.End_paragraph = from.End_paragraph
}

type Lyric_font_WOP struct {
	// insertion point

	Name string

	Number string

	NameXSD string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string
}

func (from *Lyric_font) CopyBasicFields(to *Lyric_font) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.NameXSD = from.NameXSD
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
}

type Lyric_language_WOP struct {
	// insertion point

	Name string

	Number string

	NameXSD string

	Lang string
}

func (from *Lyric_language) CopyBasicFields(to *Lyric_language) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.NameXSD = from.NameXSD
	to.Lang = from.Lang
}

type Measure_layout_WOP struct {
	// insertion point

	Name string

	Measure_distance string
}

func (from *Measure_layout) CopyBasicFields(to *Measure_layout) {
	// insertion point
	to.Name = from.Name
	to.Measure_distance = from.Measure_distance
}

type Measure_numbering_WOP struct {
	// insertion point

	Name string

	System string

	Staff int

	Multiple_rest_always Enum_Yes_no

	Multiple_rest_range Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	EnclosedText string
}

func (from *Measure_numbering) CopyBasicFields(to *Measure_numbering) {
	// insertion point
	to.Name = from.Name
	to.System = from.System
	to.Staff = from.Staff
	to.Multiple_rest_always = from.Multiple_rest_always
	to.Multiple_rest_range = from.Multiple_rest_range
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.EnclosedText = from.EnclosedText
}

type Measure_repeat_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Slashes int

	EnclosedText string
}

func (from *Measure_repeat) CopyBasicFields(to *Measure_repeat) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Slashes = from.Slashes
	to.EnclosedText = from.EnclosedText
}

type Measure_style_WOP struct {
	// insertion point

	Name string

	Number int

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Id string
}

func (from *Measure_style) CopyBasicFields(to *Measure_style) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Id = from.Id
}

type Membrane_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Membrane) CopyBasicFields(to *Membrane) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Metal_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Metal) CopyBasicFields(to *Metal) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Metronome_WOP struct {
	// insertion point

	Name string

	Parentheses Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Print_object Enum_Yes_no

	Justify Enum_Left_center_right

	Id string

	Beat_unit Enum_Note_type_value

	Beat_unit_dot string

	Metronome_arrows string

	Metronome_relation string
}

func (from *Metronome) CopyBasicFields(to *Metronome) {
	// insertion point
	to.Name = from.Name
	to.Parentheses = from.Parentheses
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Print_object = from.Print_object
	to.Justify = from.Justify
	to.Id = from.Id
	to.Beat_unit = from.Beat_unit
	to.Beat_unit_dot = from.Beat_unit_dot
	to.Metronome_arrows = from.Metronome_arrows
	to.Metronome_relation = from.Metronome_relation
}

type Metronome_beam_WOP struct {
	// insertion point

	Name string

	Number int

	EnclosedText Enum_Beam_value
}

func (from *Metronome_beam) CopyBasicFields(to *Metronome_beam) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.EnclosedText = from.EnclosedText
}

type Metronome_note_WOP struct {
	// insertion point

	Name string

	Metronome_type string

	Metronome_dot string
}

func (from *Metronome_note) CopyBasicFields(to *Metronome_note) {
	// insertion point
	to.Name = from.Name
	to.Metronome_type = from.Metronome_type
	to.Metronome_dot = from.Metronome_dot
}

type Metronome_tied_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop
}

func (from *Metronome_tied) CopyBasicFields(to *Metronome_tied) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
}

type Metronome_tuplet_WOP struct {
	// insertion point

	Name string
}

func (from *Metronome_tuplet) CopyBasicFields(to *Metronome_tuplet) {
	// insertion point
	to.Name = from.Name
}

type Midi_device_WOP struct {
	// insertion point

	Name string

	Port int

	Id string

	EnclosedText string
}

func (from *Midi_device) CopyBasicFields(to *Midi_device) {
	// insertion point
	to.Name = from.Name
	to.Port = from.Port
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Midi_instrument_WOP struct {
	// insertion point

	Name string

	Id string

	Midi_channel int

	Midi_name string

	Midi_bank int

	Midi_program int

	Midi_unpitched int

	Volume string

	Pan string

	Elevation string
}

func (from *Midi_instrument) CopyBasicFields(to *Midi_instrument) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Midi_channel = from.Midi_channel
	to.Midi_name = from.Midi_name
	to.Midi_bank = from.Midi_bank
	to.Midi_program = from.Midi_program
	to.Midi_unpitched = from.Midi_unpitched
	to.Volume = from.Volume
	to.Pan = from.Pan
	to.Elevation = from.Elevation
}

type Miscellaneous_WOP struct {
	// insertion point

	Name string
}

func (from *Miscellaneous) CopyBasicFields(to *Miscellaneous) {
	// insertion point
	to.Name = from.Name
}

type Miscellaneous_field_WOP struct {
	// insertion point

	Name string

	NameXSD string

	EnclosedText string
}

func (from *Miscellaneous_field) CopyBasicFields(to *Miscellaneous_field) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.EnclosedText = from.EnclosedText
}

type Mordent_WOP struct {
	// insertion point

	Name string
}

func (from *Mordent) CopyBasicFields(to *Mordent) {
	// insertion point
	to.Name = from.Name
}

type Multiple_rest_WOP struct {
	// insertion point

	Name string

	Use_symbols Enum_Yes_no

	EnclosedText int
}

func (from *Multiple_rest) CopyBasicFields(to *Multiple_rest) {
	// insertion point
	to.Name = from.Name
	to.Use_symbols = from.Use_symbols
	to.EnclosedText = from.EnclosedText
}

type Name_display_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no
}

func (from *Name_display) CopyBasicFields(to *Name_display) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
}

type Non_arpeggiate_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Color string

	Id string
}

func (from *Non_arpeggiate) CopyBasicFields(to *Non_arpeggiate) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Color = from.Color
	to.Id = from.Id
}

type Notations_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no

	Id string
}

func (from *Notations) CopyBasicFields(to *Notations) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
	to.Id = from.Id
}

type Note_WOP struct {
	// insertion point

	Name string

	Print_leger Enum_Yes_no

	Dynamics string

	End_dynamics string

	Attack string

	Release string

	Time_only Enum_Time_only

	Pizzicato Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Print_dot Enum_Yes_no

	Print_lyric Enum_Yes_no

	Print_object Enum_Yes_no

	Print_spacing Enum_Yes_no

	Id string

	Chord string

	Cue string

	Duration string

	Voice string

	Staff int
}

func (from *Note) CopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.Print_leger = from.Print_leger
	to.Dynamics = from.Dynamics
	to.End_dynamics = from.End_dynamics
	to.Attack = from.Attack
	to.Release = from.Release
	to.Time_only = from.Time_only
	to.Pizzicato = from.Pizzicato
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Print_dot = from.Print_dot
	to.Print_lyric = from.Print_lyric
	to.Print_object = from.Print_object
	to.Print_spacing = from.Print_spacing
	to.Id = from.Id
	to.Chord = from.Chord
	to.Cue = from.Cue
	to.Duration = from.Duration
	to.Voice = from.Voice
	to.Staff = from.Staff
}

type Note_size_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Note_size) CopyBasicFields(to *Note_size) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Note_type_WOP struct {
	// insertion point

	Name string

	Size Enum_Symbol_size

	EnclosedText Enum_Note_type_value
}

func (from *Note_type) CopyBasicFields(to *Note_type) {
	// insertion point
	to.Name = from.Name
	to.Size = from.Size
	to.EnclosedText = from.EnclosedText
}

type Notehead_WOP struct {
	// insertion point

	Name string

	Filled Enum_Yes_no

	Parentheses Enum_Yes_no

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Smufl string

	EnclosedText string
}

func (from *Notehead) CopyBasicFields(to *Notehead) {
	// insertion point
	to.Name = from.Name
	to.Filled = from.Filled
	to.Parentheses = from.Parentheses
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Notehead_text_WOP struct {
	// insertion point

	Name string
}

func (from *Notehead_text) CopyBasicFields(to *Notehead_text) {
	// insertion point
	to.Name = from.Name
}

type Numeral_WOP struct {
	// insertion point

	Name string
}

func (from *Numeral) CopyBasicFields(to *Numeral) {
	// insertion point
	to.Name = from.Name
}

type Numeral_key_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no

	Numeral_fifths int

	Numeral_mode string
}

func (from *Numeral_key) CopyBasicFields(to *Numeral_key) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
	to.Numeral_fifths = from.Numeral_fifths
	to.Numeral_mode = from.Numeral_mode
}

type Numeral_root_WOP struct {
	// insertion point

	Name string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText int
}

func (from *Numeral_root) CopyBasicFields(to *Numeral_root) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Octave_shift_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Size int

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Id string
}

func (from *Octave_shift) CopyBasicFields(to *Octave_shift) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Size = from.Size
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Id = from.Id
}

type Offset_WOP struct {
	// insertion point

	Name string

	Sound Enum_Yes_no

	EnclosedText string
}

func (from *Offset) CopyBasicFields(to *Offset) {
	// insertion point
	to.Name = from.Name
	to.Sound = from.Sound
	to.EnclosedText = from.EnclosedText
}

type Opus_WOP struct {
	// insertion point

	Name string

	Href string

	Type string

	Role string

	Title string

	Show string

	Actuate string
}

func (from *Opus) CopyBasicFields(to *Opus) {
	// insertion point
	to.Name = from.Name
	to.Href = from.Href
	to.Type = from.Type
	to.Role = from.Role
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Ornaments_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Ornaments) CopyBasicFields(to *Ornaments) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Other_appearance_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Other_appearance) CopyBasicFields(to *Other_appearance) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Other_direction_WOP struct {
	// insertion point

	Name string

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Smufl string

	Id string

	EnclosedText string
}

func (from *Other_direction) CopyBasicFields(to *Other_direction) {
	// insertion point
	to.Name = from.Name
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Smufl = from.Smufl
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Other_listening_WOP struct {
	// insertion point

	Name string

	Type string

	Player string

	Time_only Enum_Time_only

	EnclosedText string
}

func (from *Other_listening) CopyBasicFields(to *Other_listening) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Player = from.Player
	to.Time_only = from.Time_only
	to.EnclosedText = from.EnclosedText
}

type Other_notation_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop_single

	Number int

	Print_object Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Smufl string

	Id string

	EnclosedText string
}

func (from *Other_notation) CopyBasicFields(to *Other_notation) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Print_object = from.Print_object
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Smufl = from.Smufl
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Other_placement_text_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Smufl string

	EnclosedText string
}

func (from *Other_placement_text) CopyBasicFields(to *Other_placement_text) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Other_play_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Other_play) CopyBasicFields(to *Other_play) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Other_text_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Other_text) CopyBasicFields(to *Other_text) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Page_layout_WOP struct {
	// insertion point

	Name string

	Page_height string

	Page_width string
}

func (from *Page_layout) CopyBasicFields(to *Page_layout) {
	// insertion point
	to.Name = from.Name
	to.Page_height = from.Page_height
	to.Page_width = from.Page_width
}

type Page_margins_WOP struct {
	// insertion point

	Name string

	Type string

	Left_margin string

	Right_margin string

	Top_margin string

	Bottom_margin string
}

func (from *Page_margins) CopyBasicFields(to *Page_margins) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Left_margin = from.Left_margin
	to.Right_margin = from.Right_margin
	to.Top_margin = from.Top_margin
	to.Bottom_margin = from.Bottom_margin
}

type Part_clef_WOP struct {
	// insertion point

	Name string

	Sign string

	Line int

	Clef_octave_change int
}

func (from *Part_clef) CopyBasicFields(to *Part_clef) {
	// insertion point
	to.Name = from.Name
	to.Sign = from.Sign
	to.Line = from.Line
	to.Clef_octave_change = from.Clef_octave_change
}

type Part_group_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Number string

	Group_time string
}

func (from *Part_group) CopyBasicFields(to *Part_group) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Group_time = from.Group_time
}

type Part_link_WOP struct {
	// insertion point

	Name string

	Href string

	Type string

	Role string

	Title string

	Show string

	Actuate string

	Group_link string
}

func (from *Part_link) CopyBasicFields(to *Part_link) {
	// insertion point
	to.Name = from.Name
	to.Href = from.Href
	to.Type = from.Type
	to.Role = from.Role
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Group_link = from.Group_link
}

type Part_list_WOP struct {
	// insertion point

	Name string
}

func (from *Part_list) CopyBasicFields(to *Part_list) {
	// insertion point
	to.Name = from.Name
}

type Part_name_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Print_object Enum_Yes_no

	Justify Enum_Left_center_right

	EnclosedText string
}

func (from *Part_name) CopyBasicFields(to *Part_name) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Print_object = from.Print_object
	to.Justify = from.Justify
	to.EnclosedText = from.EnclosedText
}

type Part_symbol_WOP struct {
	// insertion point

	Name string

	Top_staff int

	Bottom_staff int

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	EnclosedText Enum_Group_symbol_value
}

func (from *Part_symbol) CopyBasicFields(to *Part_symbol) {
	// insertion point
	to.Name = from.Name
	to.Top_staff = from.Top_staff
	to.Bottom_staff = from.Bottom_staff
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Part_transpose_WOP struct {
	// insertion point

	Name string

	Diatonic int

	Chromatic string

	Octave_change int

	Double float64
}

func (from *Part_transpose) CopyBasicFields(to *Part_transpose) {
	// insertion point
	to.Name = from.Name
	to.Diatonic = from.Diatonic
	to.Chromatic = from.Chromatic
	to.Octave_change = from.Octave_change
	to.Double = from.Double
}

type Pedal_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Line Enum_Yes_no

	Sign Enum_Yes_no

	Abbreviated Enum_Yes_no

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Pedal) CopyBasicFields(to *Pedal) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line = from.Line
	to.Sign = from.Sign
	to.Abbreviated = from.Abbreviated
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Pedal_tuning_WOP struct {
	// insertion point

	Name string

	Pedal_step Enum_Step

	Pedal_alter string
}

func (from *Pedal_tuning) CopyBasicFields(to *Pedal_tuning) {
	// insertion point
	to.Name = from.Name
	to.Pedal_step = from.Pedal_step
	to.Pedal_alter = from.Pedal_alter
}

type Per_minute_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	EnclosedText string
}

func (from *Per_minute) CopyBasicFields(to *Per_minute) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.EnclosedText = from.EnclosedText
}

type Percussion_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Enclosure string

	Id string

	Stick_location string
}

func (from *Percussion) CopyBasicFields(to *Percussion) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Enclosure = from.Enclosure
	to.Id = from.Id
	to.Stick_location = from.Stick_location
}

type Pitch_WOP struct {
	// insertion point

	Name string

	Step Enum_Step

	Alter string

	Octave int
}

func (from *Pitch) CopyBasicFields(to *Pitch) {
	// insertion point
	to.Name = from.Name
	to.Step = from.Step
	to.Alter = from.Alter
	to.Octave = from.Octave
}

type Pitched_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Pitched) CopyBasicFields(to *Pitched) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Placement_text_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Placement_text) CopyBasicFields(to *Placement_text) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Play_WOP struct {
	// insertion point

	Name string

	Id string

	Ipa string

	Mute string

	Semi_pitched string
}

func (from *Play) CopyBasicFields(to *Play) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Ipa = from.Ipa
	to.Mute = from.Mute
	to.Semi_pitched = from.Semi_pitched
}

type Player_WOP struct {
	// insertion point

	Name string

	Id string

	Player_name string
}

func (from *Player) CopyBasicFields(to *Player) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Player_name = from.Player_name
}

type Principal_voice_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Symbol string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string

	EnclosedText string
}

func (from *Principal_voice) CopyBasicFields(to *Principal_voice) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Symbol = from.Symbol
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Print_WOP struct {
	// insertion point

	Name string

	Staff_spacing string

	New_system Enum_Yes_no

	New_page Enum_Yes_no

	Blank_page int

	Page_number string

	Id string
}

func (from *Print) CopyBasicFields(to *Print) {
	// insertion point
	to.Name = from.Name
	to.Staff_spacing = from.Staff_spacing
	to.New_system = from.New_system
	to.New_page = from.New_page
	to.Blank_page = from.Blank_page
	to.Page_number = from.Page_number
	to.Id = from.Id
}

type Release_WOP struct {
	// insertion point

	Name string
}

func (from *Release) CopyBasicFields(to *Release) {
	// insertion point
	to.Name = from.Name
}

type Repeat_WOP struct {
	// insertion point

	Name string

	Direction string

	Times int

	After_jump Enum_Yes_no

	Winged string
}

func (from *Repeat) CopyBasicFields(to *Repeat) {
	// insertion point
	to.Name = from.Name
	to.Direction = from.Direction
	to.Times = from.Times
	to.After_jump = from.After_jump
	to.Winged = from.Winged
}

type Rest_WOP struct {
	// insertion point

	Name string

	Measure Enum_Yes_no

	Display_step Enum_Step

	Display_octave int
}

func (from *Rest) CopyBasicFields(to *Rest) {
	// insertion point
	to.Name = from.Name
	to.Measure = from.Measure
	to.Display_step = from.Display_step
	to.Display_octave = from.Display_octave
}

type Root_WOP struct {
	// insertion point

	Name string
}

func (from *Root) CopyBasicFields(to *Root) {
	// insertion point
	to.Name = from.Name
}

type Root_step_WOP struct {
	// insertion point

	Name string

	Text string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText Enum_Step
}

func (from *Root_step) CopyBasicFields(to *Root_step) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Scaling_WOP struct {
	// insertion point

	Name string

	Millimeters string

	Tenths string
}

func (from *Scaling) CopyBasicFields(to *Scaling) {
	// insertion point
	to.Name = from.Name
	to.Millimeters = from.Millimeters
	to.Tenths = from.Tenths
}

type Scordatura_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Scordatura) CopyBasicFields(to *Scordatura) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Score_instrument_WOP struct {
	// insertion point

	Name string

	Id string

	Instrument_name string

	Instrument_abbreviation string

	Instrument_sound string

	Solo string

	Ensemble string
}

func (from *Score_instrument) CopyBasicFields(to *Score_instrument) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Instrument_name = from.Instrument_name
	to.Instrument_abbreviation = from.Instrument_abbreviation
	to.Instrument_sound = from.Instrument_sound
	to.Solo = from.Solo
	to.Ensemble = from.Ensemble
}

type Score_part_WOP struct {
	// insertion point

	Name string

	Id string

	Group string
}

func (from *Score_part) CopyBasicFields(to *Score_part) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Group = from.Group
}

type Score_partwise_WOP struct {
	// insertion point

	Name string

	Version string

	Movement_number string

	Movement_title string
}

func (from *Score_partwise) CopyBasicFields(to *Score_partwise) {
	// insertion point
	to.Name = from.Name
	to.Version = from.Version
	to.Movement_number = from.Movement_number
	to.Movement_title = from.Movement_title
}

type Score_timewise_WOP struct {
	// insertion point

	Name string

	Version string

	Movement_number string

	Movement_title string
}

func (from *Score_timewise) CopyBasicFields(to *Score_timewise) {
	// insertion point
	to.Name = from.Name
	to.Version = from.Version
	to.Movement_number = from.Movement_number
	to.Movement_title = from.Movement_title
}

type Segno_WOP struct {
	// insertion point

	Name string

	Smufl string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Segno) CopyBasicFields(to *Segno) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Slash_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Use_dots Enum_Yes_no

	Use_stems Enum_Yes_no

	Slash_type Enum_Note_type_value

	Slash_dot string

	Except_voice string
}

func (from *Slash) CopyBasicFields(to *Slash) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Use_dots = from.Use_dots
	to.Use_stems = from.Use_stems
	to.Slash_type = from.Slash_type
	to.Slash_dot = from.Slash_dot
	to.Except_voice = from.Except_voice
}

type Slide_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Number int

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Accelerate Enum_Yes_no

	Beats string

	First_beat string

	Last_beat string

	Id string

	EnclosedText string
}

func (from *Slide) CopyBasicFields(to *Slide) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Accelerate = from.Accelerate
	to.Beats = from.Beats
	to.First_beat = from.First_beat
	to.Last_beat = from.Last_beat
	to.Id = from.Id
	to.EnclosedText = from.EnclosedText
}

type Slur_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop_continue

	Number int

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Orientation string

	Bezier_x string

	Bezier_y string

	Bezier_x2 string

	Bezier_y2 string

	Bezier_offset string

	Bezier_offset2 string

	Color string

	Id string
}

func (from *Slur) CopyBasicFields(to *Slur) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Orientation = from.Orientation
	to.Bezier_x = from.Bezier_x
	to.Bezier_y = from.Bezier_y
	to.Bezier_x2 = from.Bezier_x2
	to.Bezier_y2 = from.Bezier_y2
	to.Bezier_offset = from.Bezier_offset
	to.Bezier_offset2 = from.Bezier_offset2
	to.Color = from.Color
	to.Id = from.Id
}

type Sound_WOP struct {
	// insertion point

	Name string

	Tempo string

	Dynamics string

	Dacapo Enum_Yes_no

	Segno string

	Dalsegno string

	Coda string

	Tocoda string

	Divisions string

	Forward_repeat Enum_Yes_no

	Fine string

	Time_only Enum_Time_only

	Pizzicato Enum_Yes_no

	Pan string

	Elevation string

	Damper_pedal string

	Soft_pedal string

	Sostenuto_pedal string

	Id string
}

func (from *Sound) CopyBasicFields(to *Sound) {
	// insertion point
	to.Name = from.Name
	to.Tempo = from.Tempo
	to.Dynamics = from.Dynamics
	to.Dacapo = from.Dacapo
	to.Segno = from.Segno
	to.Dalsegno = from.Dalsegno
	to.Coda = from.Coda
	to.Tocoda = from.Tocoda
	to.Divisions = from.Divisions
	to.Forward_repeat = from.Forward_repeat
	to.Fine = from.Fine
	to.Time_only = from.Time_only
	to.Pizzicato = from.Pizzicato
	to.Pan = from.Pan
	to.Elevation = from.Elevation
	to.Damper_pedal = from.Damper_pedal
	to.Soft_pedal = from.Soft_pedal
	to.Sostenuto_pedal = from.Sostenuto_pedal
	to.Id = from.Id
}

type Staff_details_WOP struct {
	// insertion point

	Name string

	Number int

	Show_frets string

	Print_object Enum_Yes_no

	Print_spacing Enum_Yes_no

	Staff_type string

	Staff_lines int

	Capo int
}

func (from *Staff_details) CopyBasicFields(to *Staff_details) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Show_frets = from.Show_frets
	to.Print_object = from.Print_object
	to.Print_spacing = from.Print_spacing
	to.Staff_type = from.Staff_type
	to.Staff_lines = from.Staff_lines
	to.Capo = from.Capo
}

type Staff_divide_WOP struct {
	// insertion point

	Name string

	Type string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *Staff_divide) CopyBasicFields(to *Staff_divide) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type Staff_layout_WOP struct {
	// insertion point

	Name string

	Number int

	Staff_distance string
}

func (from *Staff_layout) CopyBasicFields(to *Staff_layout) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Staff_distance = from.Staff_distance
}

type Staff_size_WOP struct {
	// insertion point

	Name string

	Scaling string

	EnclosedText string
}

func (from *Staff_size) CopyBasicFields(to *Staff_size) {
	// insertion point
	to.Name = from.Name
	to.Scaling = from.Scaling
	to.EnclosedText = from.EnclosedText
}

type Staff_tuning_WOP struct {
	// insertion point

	Name string

	Line int

	Tuning_step Enum_Step

	Tuning_alter string

	Tuning_octave int
}

func (from *Staff_tuning) CopyBasicFields(to *Staff_tuning) {
	// insertion point
	to.Name = from.Name
	to.Line = from.Line
	to.Tuning_step = from.Tuning_step
	to.Tuning_alter = from.Tuning_alter
	to.Tuning_octave = from.Tuning_octave
}

type Stem_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	EnclosedText string
}

func (from *Stem) CopyBasicFields(to *Stem) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Stick_WOP struct {
	// insertion point

	Name string

	Tip Enum_Tip_direction

	Parentheses Enum_Yes_no

	Dashed_circle Enum_Yes_no

	Stick_type string

	Stick_material string
}

func (from *Stick) CopyBasicFields(to *Stick) {
	// insertion point
	to.Name = from.Name
	to.Tip = from.Tip
	to.Parentheses = from.Parentheses
	to.Dashed_circle = from.Dashed_circle
	to.Stick_type = from.Stick_type
	to.Stick_material = from.Stick_material
}

type String_mute_WOP struct {
	// insertion point

	Name string

	Type string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Id string
}

func (from *String_mute) CopyBasicFields(to *String_mute) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Id = from.Id
}

type String_type_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText int
}

func (from *String_type) CopyBasicFields(to *String_type) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Strong_accent_WOP struct {
	// insertion point

	Name string
}

func (from *Strong_accent) CopyBasicFields(to *Strong_accent) {
	// insertion point
	to.Name = from.Name
}

type Style_text_WOP struct {
	// insertion point

	Name string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText string
}

func (from *Style_text) CopyBasicFields(to *Style_text) {
	// insertion point
	to.Name = from.Name
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Supports_WOP struct {
	// insertion point

	Name string

	Type Enum_Yes_no

	Element string

	Attribute string

	Value string
}

func (from *Supports) CopyBasicFields(to *Supports) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Element = from.Element
	to.Attribute = from.Attribute
	to.Value = from.Value
}

type Swing_WOP struct {
	// insertion point

	Name string

	Straight string

	First int

	Second int

	Swing_type Enum_Note_type_value

	Swing_style string
}

func (from *Swing) CopyBasicFields(to *Swing) {
	// insertion point
	to.Name = from.Name
	to.Straight = from.Straight
	to.First = from.First
	to.Second = from.Second
	to.Swing_type = from.Swing_type
	to.Swing_style = from.Swing_style
}

type Sync_WOP struct {
	// insertion point

	Name string

	Type string

	Latency int

	Player string

	Time_only Enum_Time_only
}

func (from *Sync) CopyBasicFields(to *Sync) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Latency = from.Latency
	to.Player = from.Player
	to.Time_only = from.Time_only
}

type System_dividers_WOP struct {
	// insertion point

	Name string
}

func (from *System_dividers) CopyBasicFields(to *System_dividers) {
	// insertion point
	to.Name = from.Name
}

type System_layout_WOP struct {
	// insertion point

	Name string

	System_distance string

	Top_system_distance string
}

func (from *System_layout) CopyBasicFields(to *System_layout) {
	// insertion point
	to.Name = from.Name
	to.System_distance = from.System_distance
	to.Top_system_distance = from.Top_system_distance
}

type System_margins_WOP struct {
	// insertion point

	Name string

	Left_margin string

	Right_margin string
}

func (from *System_margins) CopyBasicFields(to *System_margins) {
	// insertion point
	to.Name = from.Name
	to.Left_margin = from.Left_margin
	to.Right_margin = from.Right_margin
}

type Tap_WOP struct {
	// insertion point

	Name string

	Hand string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	EnclosedText string
}

func (from *Tap) CopyBasicFields(to *Tap) {
	// insertion point
	to.Name = from.Name
	to.Hand = from.Hand
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.EnclosedText = from.EnclosedText
}

type Technical_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *Technical) CopyBasicFields(to *Technical) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Text_element_data_WOP struct {
	// insertion point

	Name string

	Lang string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Underline int

	Overline int

	Line_through int

	Rotation string

	Letter_spacing string

	Dir string

	EnclosedText string
}

func (from *Text_element_data) CopyBasicFields(to *Text_element_data) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Underline = from.Underline
	to.Overline = from.Overline
	to.Line_through = from.Line_through
	to.Rotation = from.Rotation
	to.Letter_spacing = from.Letter_spacing
	to.Dir = from.Dir
	to.EnclosedText = from.EnclosedText
}

type Tie_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Time_only Enum_Time_only
}

func (from *Tie) CopyBasicFields(to *Tie) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Time_only = from.Time_only
}

type Tied_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Orientation string

	Bezier_x string

	Bezier_y string

	Bezier_x2 string

	Bezier_y2 string

	Bezier_offset string

	Bezier_offset2 string

	Color string

	Id string
}

func (from *Tied) CopyBasicFields(to *Tied) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Orientation = from.Orientation
	to.Bezier_x = from.Bezier_x
	to.Bezier_y = from.Bezier_y
	to.Bezier_x2 = from.Bezier_x2
	to.Bezier_y2 = from.Bezier_y2
	to.Bezier_offset = from.Bezier_offset
	to.Bezier_offset2 = from.Bezier_offset2
	to.Color = from.Color
	to.Id = from.Id
}

type Time_WOP struct {
	// insertion point

	Name string

	Number int

	Symbol Enum_Time_symbol

	Separator Enum_Time_separator

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Halign string

	Valign string

	Print_object Enum_Yes_no

	Id string

	Beats string

	Beat_type string

	Senza_misura string
}

func (from *Time) CopyBasicFields(to *Time) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Symbol = from.Symbol
	to.Separator = from.Separator
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Halign = from.Halign
	to.Valign = from.Valign
	to.Print_object = from.Print_object
	to.Id = from.Id
	to.Beats = from.Beats
	to.Beat_type = from.Beat_type
	to.Senza_misura = from.Senza_misura
}

type Time_modification_WOP struct {
	// insertion point

	Name string

	Actual_notes int

	Normal_notes int

	Normal_type Enum_Note_type_value

	Normal_dot string
}

func (from *Time_modification) CopyBasicFields(to *Time_modification) {
	// insertion point
	to.Name = from.Name
	to.Actual_notes = from.Actual_notes
	to.Normal_notes = from.Normal_notes
	to.Normal_type = from.Normal_type
	to.Normal_dot = from.Normal_dot
}

type Timpani_WOP struct {
	// insertion point

	Name string

	Smufl string
}

func (from *Timpani) CopyBasicFields(to *Timpani) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
}

type Transpose_WOP struct {
	// insertion point

	Name string

	Number int

	Id string

	Diatonic int

	Chromatic string

	Octave_change int

	Double float64
}

func (from *Transpose) CopyBasicFields(to *Transpose) {
	// insertion point
	to.Name = from.Name
	to.Number = from.Number
	to.Id = from.Id
	to.Diatonic = from.Diatonic
	to.Chromatic = from.Chromatic
	to.Octave_change = from.Octave_change
	to.Double = from.Double
}

type Tremolo_WOP struct {
	// insertion point

	Name string

	Type string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	Placement string

	Smufl string

	EnclosedText int
}

func (from *Tremolo) CopyBasicFields(to *Tremolo) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.Placement = from.Placement
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Tuplet_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop

	Number int

	Bracket Enum_Yes_no

	Show_number string

	Show_type Enum_Show_tuplet

	Line_shape string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Id string
}

func (from *Tuplet) CopyBasicFields(to *Tuplet) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Bracket = from.Bracket
	to.Show_number = from.Show_number
	to.Show_type = from.Show_type
	to.Line_shape = from.Line_shape
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Id = from.Id
}

type Tuplet_dot_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string
}

func (from *Tuplet_dot) CopyBasicFields(to *Tuplet_dot) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
}

type Tuplet_number_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText int
}

func (from *Tuplet_number) CopyBasicFields(to *Tuplet_number) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Tuplet_portion_WOP struct {
	// insertion point

	Name string
}

func (from *Tuplet_portion) CopyBasicFields(to *Tuplet_portion) {
	// insertion point
	to.Name = from.Name
}

type Tuplet_type_WOP struct {
	// insertion point

	Name string

	Font_family string

	Font_style string

	Font_size string

	Font_weight string

	Color string

	EnclosedText Enum_Note_type_value
}

func (from *Tuplet_type) CopyBasicFields(to *Tuplet_type) {
	// insertion point
	to.Name = from.Name
	to.Font_family = from.Font_family
	to.Font_style = from.Font_style
	to.Font_size = from.Font_size
	to.Font_weight = from.Font_weight
	to.Color = from.Color
	to.EnclosedText = from.EnclosedText
}

type Typed_text_WOP struct {
	// insertion point

	Name string

	Type string

	EnclosedText string
}

func (from *Typed_text) CopyBasicFields(to *Typed_text) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.EnclosedText = from.EnclosedText
}

type Unpitched_WOP struct {
	// insertion point

	Name string

	Display_step Enum_Step

	Display_octave int
}

func (from *Unpitched) CopyBasicFields(to *Unpitched) {
	// insertion point
	to.Name = from.Name
	to.Display_step = from.Display_step
	to.Display_octave = from.Display_octave
}

type Virtual_instrument_WOP struct {
	// insertion point

	Name string

	Virtual_library string

	Virtual_name string
}

func (from *Virtual_instrument) CopyBasicFields(to *Virtual_instrument) {
	// insertion point
	to.Name = from.Name
	to.Virtual_library = from.Virtual_library
	to.Virtual_name = from.Virtual_name
}

type Wait_WOP struct {
	// insertion point

	Name string

	Player string

	Time_only Enum_Time_only
}

func (from *Wait) CopyBasicFields(to *Wait) {
	// insertion point
	to.Name = from.Name
	to.Player = from.Player
	to.Time_only = from.Time_only
}

type Wavy_line_WOP struct {
	// insertion point

	Name string

	Type Enum_Start_stop_continue

	Number int

	Smufl string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Placement string

	Color string

	Start_note string

	Trill_step string

	Two_note_turn string

	Accelerate Enum_Yes_no

	Beats string

	Second_beat string

	Last_beat string
}

func (from *Wavy_line) CopyBasicFields(to *Wavy_line) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Smufl = from.Smufl
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Placement = from.Placement
	to.Color = from.Color
	to.Start_note = from.Start_note
	to.Trill_step = from.Trill_step
	to.Two_note_turn = from.Two_note_turn
	to.Accelerate = from.Accelerate
	to.Beats = from.Beats
	to.Second_beat = from.Second_beat
	to.Last_beat = from.Last_beat
}

type Wedge_WOP struct {
	// insertion point

	Name string

	Type string

	Number int

	Spread string

	Niente Enum_Yes_no

	Line_type string

	Dash_length string

	Space_length string

	Default_x string

	Default_y string

	Relative_x string

	Relative_y string

	Color string

	Id string
}

func (from *Wedge) CopyBasicFields(to *Wedge) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Number = from.Number
	to.Spread = from.Spread
	to.Niente = from.Niente
	to.Line_type = from.Line_type
	to.Dash_length = from.Dash_length
	to.Space_length = from.Space_length
	to.Default_x = from.Default_x
	to.Default_y = from.Default_y
	to.Relative_x = from.Relative_x
	to.Relative_y = from.Relative_y
	to.Color = from.Color
	to.Id = from.Id
}

type Wood_WOP struct {
	// insertion point

	Name string

	Smufl string

	EnclosedText string
}

func (from *Wood) CopyBasicFields(to *Wood) {
	// insertion point
	to.Name = from.Name
	to.Smufl = from.Smufl
	to.EnclosedText = from.EnclosedText
}

type Work_WOP struct {
	// insertion point

	Name string

	Work_number string

	Work_title string
}

func (from *Work) CopyBasicFields(to *Work) {
	// insertion point
	to.Name = from.Name
	to.Work_number = from.Work_number
	to.Work_title = from.Work_title
}

