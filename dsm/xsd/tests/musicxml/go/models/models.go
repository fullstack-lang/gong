// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "above-below"
type Enum_Above_below string

const (
	Enum_Above_below_Above Enum_Above_below = "above"
	Enum_Above_below_Below Enum_Above_below = "below"
)

// From xsd simple type with enumerate restriction "accidental-value"
type Enum_Accidental_value string

const (
	Enum_Accidental_value_Sharp                Enum_Accidental_value = "sharp"
	Enum_Accidental_value_Natural              Enum_Accidental_value = "natural"
	Enum_Accidental_value_Flat                 Enum_Accidental_value = "flat"
	Enum_Accidental_value_Double_sharp         Enum_Accidental_value = "double-sharp"
	Enum_Accidental_value_Sharp_sharp          Enum_Accidental_value = "sharp-sharp"
	Enum_Accidental_value_Flat_flat            Enum_Accidental_value = "flat-flat"
	Enum_Accidental_value_Natural_sharp        Enum_Accidental_value = "natural-sharp"
	Enum_Accidental_value_Natural_flat         Enum_Accidental_value = "natural-flat"
	Enum_Accidental_value_Quarter_flat         Enum_Accidental_value = "quarter-flat"
	Enum_Accidental_value_Quarter_sharp        Enum_Accidental_value = "quarter-sharp"
	Enum_Accidental_value_Three_quarters_flat  Enum_Accidental_value = "three-quarters-flat"
	Enum_Accidental_value_Three_quarters_sharp Enum_Accidental_value = "three-quarters-sharp"
	Enum_Accidental_value_Sharp_down           Enum_Accidental_value = "sharp-down"
	Enum_Accidental_value_Sharp_up             Enum_Accidental_value = "sharp-up"
	Enum_Accidental_value_Natural_down         Enum_Accidental_value = "natural-down"
	Enum_Accidental_value_Natural_up           Enum_Accidental_value = "natural-up"
	Enum_Accidental_value_Flat_down            Enum_Accidental_value = "flat-down"
	Enum_Accidental_value_Flat_up              Enum_Accidental_value = "flat-up"
	Enum_Accidental_value_Double_sharp_down    Enum_Accidental_value = "double-sharp-down"
	Enum_Accidental_value_Double_sharp_up      Enum_Accidental_value = "double-sharp-up"
	Enum_Accidental_value_Flat_flat_down       Enum_Accidental_value = "flat-flat-down"
	Enum_Accidental_value_Flat_flat_up         Enum_Accidental_value = "flat-flat-up"
	Enum_Accidental_value_Arrow_down           Enum_Accidental_value = "arrow-down"
	Enum_Accidental_value_Arrow_up             Enum_Accidental_value = "arrow-up"
	Enum_Accidental_value_Triple_sharp         Enum_Accidental_value = "triple-sharp"
	Enum_Accidental_value_Triple_flat          Enum_Accidental_value = "triple-flat"
	Enum_Accidental_value_Slash_quarter_sharp  Enum_Accidental_value = "slash-quarter-sharp"
	Enum_Accidental_value_Slash_sharp          Enum_Accidental_value = "slash-sharp"
	Enum_Accidental_value_Slash_flat           Enum_Accidental_value = "slash-flat"
	Enum_Accidental_value_Double_slash_flat    Enum_Accidental_value = "double-slash-flat"
	Enum_Accidental_value_Sharp_1              Enum_Accidental_value = "sharp-1"
	Enum_Accidental_value_Sharp_2              Enum_Accidental_value = "sharp-2"
	Enum_Accidental_value_Sharp_3              Enum_Accidental_value = "sharp-3"
	Enum_Accidental_value_Sharp_5              Enum_Accidental_value = "sharp-5"
	Enum_Accidental_value_Flat_1               Enum_Accidental_value = "flat-1"
	Enum_Accidental_value_Flat_2               Enum_Accidental_value = "flat-2"
	Enum_Accidental_value_Flat_3               Enum_Accidental_value = "flat-3"
	Enum_Accidental_value_Flat_4               Enum_Accidental_value = "flat-4"
	Enum_Accidental_value_Sori                 Enum_Accidental_value = "sori"
	Enum_Accidental_value_Koron                Enum_Accidental_value = "koron"
	Enum_Accidental_value_Other                Enum_Accidental_value = "other"
)

// From xsd simple type with enumerate restriction "arrow-direction"
type Enum_Arrow_direction string

const (
	Enum_Arrow_direction_Left                Enum_Arrow_direction = "left"
	Enum_Arrow_direction_Up                  Enum_Arrow_direction = "up"
	Enum_Arrow_direction_Right               Enum_Arrow_direction = "right"
	Enum_Arrow_direction_Down                Enum_Arrow_direction = "down"
	Enum_Arrow_direction_Northwest           Enum_Arrow_direction = "northwest"
	Enum_Arrow_direction_Northeast           Enum_Arrow_direction = "northeast"
	Enum_Arrow_direction_Southeast           Enum_Arrow_direction = "southeast"
	Enum_Arrow_direction_Southwest           Enum_Arrow_direction = "southwest"
	Enum_Arrow_direction_Left_right          Enum_Arrow_direction = "left right"
	Enum_Arrow_direction_Up_down             Enum_Arrow_direction = "up down"
	Enum_Arrow_direction_Northwest_southeast Enum_Arrow_direction = "northwest southeast"
	Enum_Arrow_direction_Northeast_southwest Enum_Arrow_direction = "northeast southwest"
	Enum_Arrow_direction_Other               Enum_Arrow_direction = "other"
)

// From xsd simple type with enumerate restriction "arrow-style"
type Enum_Arrow_style string

const (
	Enum_Arrow_style_Single   Enum_Arrow_style = "single"
	Enum_Arrow_style_Double   Enum_Arrow_style = "double"
	Enum_Arrow_style_Filled   Enum_Arrow_style = "filled"
	Enum_Arrow_style_Hollow   Enum_Arrow_style = "hollow"
	Enum_Arrow_style_Paired   Enum_Arrow_style = "paired"
	Enum_Arrow_style_Combined Enum_Arrow_style = "combined"
	Enum_Arrow_style_Other    Enum_Arrow_style = "other"
)

// From xsd simple type with enumerate restriction "backward-forward"
type Enum_Backward_forward string

const (
	Enum_Backward_forward_Backward Enum_Backward_forward = "backward"
	Enum_Backward_forward_Forward  Enum_Backward_forward = "forward"
)

// From xsd simple type with enumerate restriction "bar-style"
type Enum_Bar_style string

const (
	Enum_Bar_style_Regular     Enum_Bar_style = "regular"
	Enum_Bar_style_Dotted      Enum_Bar_style = "dotted"
	Enum_Bar_style_Dashed      Enum_Bar_style = "dashed"
	Enum_Bar_style_Heavy       Enum_Bar_style = "heavy"
	Enum_Bar_style_Light_light Enum_Bar_style = "light-light"
	Enum_Bar_style_Light_heavy Enum_Bar_style = "light-heavy"
	Enum_Bar_style_Heavy_light Enum_Bar_style = "heavy-light"
	Enum_Bar_style_Heavy_heavy Enum_Bar_style = "heavy-heavy"
	Enum_Bar_style_Tick        Enum_Bar_style = "tick"
	Enum_Bar_style_Short       Enum_Bar_style = "short"
	Enum_Bar_style_None        Enum_Bar_style = "none"
)

// From xsd simple type with enumerate restriction "beam-value"
type Enum_Beam_value string

const (
	Enum_Beam_value_Begin         Enum_Beam_value = "begin"
	Enum_Beam_value_Continue      Enum_Beam_value = "continue"
	Enum_Beam_value_End           Enum_Beam_value = "end"
	Enum_Beam_value_Forward_hook  Enum_Beam_value = "forward hook"
	Enum_Beam_value_Backward_hook Enum_Beam_value = "backward hook"
)

// From xsd simple type with enumerate restriction "beater-value"
type Enum_Beater_value string

const (
	Enum_Beater_value_Bow                   Enum_Beater_value = "bow"
	Enum_Beater_value_Chime_hammer          Enum_Beater_value = "chime hammer"
	Enum_Beater_value_Coin                  Enum_Beater_value = "coin"
	Enum_Beater_value_Drum_stick            Enum_Beater_value = "drum stick"
	Enum_Beater_value_Finger                Enum_Beater_value = "finger"
	Enum_Beater_value_Fingernail            Enum_Beater_value = "fingernail"
	Enum_Beater_value_Fist                  Enum_Beater_value = "fist"
	Enum_Beater_value_Guiro_scraper         Enum_Beater_value = "guiro scraper"
	Enum_Beater_value_Hammer                Enum_Beater_value = "hammer"
	Enum_Beater_value_Hand                  Enum_Beater_value = "hand"
	Enum_Beater_value_Jazz_stick            Enum_Beater_value = "jazz stick"
	Enum_Beater_value_Knitting_needle       Enum_Beater_value = "knitting needle"
	Enum_Beater_value_Metal_hammer          Enum_Beater_value = "metal hammer"
	Enum_Beater_value_Slide_brush_on_gong   Enum_Beater_value = "slide brush on gong"
	Enum_Beater_value_Snare_stick           Enum_Beater_value = "snare stick"
	Enum_Beater_value_Spoon_mallet          Enum_Beater_value = "spoon mallet"
	Enum_Beater_value_Superball             Enum_Beater_value = "superball"
	Enum_Beater_value_Triangle_beater       Enum_Beater_value = "triangle beater"
	Enum_Beater_value_Triangle_beater_plain Enum_Beater_value = "triangle beater plain"
	Enum_Beater_value_Wire_brush            Enum_Beater_value = "wire brush"
)

// From xsd simple type with enumerate restriction "bend-shape"
type Enum_Bend_shape string

const (
	Enum_Bend_shape_Angled Enum_Bend_shape = "angled"
	Enum_Bend_shape_Curved Enum_Bend_shape = "curved"
)

// From xsd simple type with enumerate restriction "breath-mark-value"
type Enum_Breath_mark_value string

const (
	Enum_Breath_mark_value_        Enum_Breath_mark_value = ""
	Enum_Breath_mark_value_Comma   Enum_Breath_mark_value = "comma"
	Enum_Breath_mark_value_Tick    Enum_Breath_mark_value = "tick"
	Enum_Breath_mark_value_Upbow   Enum_Breath_mark_value = "upbow"
	Enum_Breath_mark_value_Salzedo Enum_Breath_mark_value = "salzedo"
)

// From xsd simple type with enumerate restriction "caesura-value"
type Enum_Caesura_value string

const (
	Enum_Caesura_value_Normal Enum_Caesura_value = "normal"
	Enum_Caesura_value_Thick  Enum_Caesura_value = "thick"
	Enum_Caesura_value_Short  Enum_Caesura_value = "short"
	Enum_Caesura_value_Curved Enum_Caesura_value = "curved"
	Enum_Caesura_value_Single Enum_Caesura_value = "single"
	Enum_Caesura_value_       Enum_Caesura_value = ""
)

// From xsd simple type with enumerate restriction "cancel-location"
type Enum_Cancel_location string

const (
	Enum_Cancel_location_Left           Enum_Cancel_location = "left"
	Enum_Cancel_location_Right          Enum_Cancel_location = "right"
	Enum_Cancel_location_Before_barline Enum_Cancel_location = "before-barline"
)

// From xsd simple type with enumerate restriction "circular-arrow"
type Enum_Circular_arrow string

const (
	Enum_Circular_arrow_Clockwise     Enum_Circular_arrow = "clockwise"
	Enum_Circular_arrow_Anticlockwise Enum_Circular_arrow = "anticlockwise"
)

// From xsd simple type with enumerate restriction "clef-sign"
type Enum_Clef_sign string

const (
	Enum_Clef_sign_G          Enum_Clef_sign = "G"
	Enum_Clef_sign_F          Enum_Clef_sign = "F"
	Enum_Clef_sign_C          Enum_Clef_sign = "C"
	Enum_Clef_sign_Percussion Enum_Clef_sign = "percussion"
	Enum_Clef_sign_TAB        Enum_Clef_sign = "TAB"
	Enum_Clef_sign_Jianpu     Enum_Clef_sign = "jianpu"
	Enum_Clef_sign_None       Enum_Clef_sign = "none"
)

// From xsd simple type with enumerate restriction "color"
type Enum_Color string

const ()

// From xsd simple type with enumerate restriction "comma-separated-text"
type Enum_Comma_separated_text string

const ()

// From xsd simple type with enumerate restriction "degree-symbol-value"
type Enum_Degree_symbol_value string

const (
	Enum_Degree_symbol_value_Major           Enum_Degree_symbol_value = "major"
	Enum_Degree_symbol_value_Minor           Enum_Degree_symbol_value = "minor"
	Enum_Degree_symbol_value_Augmented       Enum_Degree_symbol_value = "augmented"
	Enum_Degree_symbol_value_Diminished      Enum_Degree_symbol_value = "diminished"
	Enum_Degree_symbol_value_Half_diminished Enum_Degree_symbol_value = "half-diminished"
)

// From xsd simple type with enumerate restriction "degree-type-value"
type Enum_Degree_type_value string

const (
	Enum_Degree_type_value_Add      Enum_Degree_type_value = "add"
	Enum_Degree_type_value_Alter    Enum_Degree_type_value = "alter"
	Enum_Degree_type_value_Subtract Enum_Degree_type_value = "subtract"
)

// From xsd simple type with enumerate restriction "distance-type"
type Enum_Distance_type string

const ()

// From xsd simple type with enumerate restriction "effect-value"
type Enum_Effect_value string

const (
	Enum_Effect_value_Anvil          Enum_Effect_value = "anvil"
	Enum_Effect_value_Auto_horn      Enum_Effect_value = "auto horn"
	Enum_Effect_value_Bird_whistle   Enum_Effect_value = "bird whistle"
	Enum_Effect_value_Cannon         Enum_Effect_value = "cannon"
	Enum_Effect_value_Duck_call      Enum_Effect_value = "duck call"
	Enum_Effect_value_Gun_shot       Enum_Effect_value = "gun shot"
	Enum_Effect_value_Klaxon_horn    Enum_Effect_value = "klaxon horn"
	Enum_Effect_value_Lions_roar     Enum_Effect_value = "lions roar"
	Enum_Effect_value_Lotus_flute    Enum_Effect_value = "lotus flute"
	Enum_Effect_value_Megaphone      Enum_Effect_value = "megaphone"
	Enum_Effect_value_Police_whistle Enum_Effect_value = "police whistle"
	Enum_Effect_value_Siren          Enum_Effect_value = "siren"
	Enum_Effect_value_Slide_whistle  Enum_Effect_value = "slide whistle"
	Enum_Effect_value_Thunder_sheet  Enum_Effect_value = "thunder sheet"
	Enum_Effect_value_Wind_machine   Enum_Effect_value = "wind machine"
	Enum_Effect_value_Wind_whistle   Enum_Effect_value = "wind whistle"
)

// From xsd simple type with enumerate restriction "enclosure-shape"
type Enum_Enclosure_shape string

const (
	Enum_Enclosure_shape_Rectangle        Enum_Enclosure_shape = "rectangle"
	Enum_Enclosure_shape_Square           Enum_Enclosure_shape = "square"
	Enum_Enclosure_shape_Oval             Enum_Enclosure_shape = "oval"
	Enum_Enclosure_shape_Circle           Enum_Enclosure_shape = "circle"
	Enum_Enclosure_shape_Bracket          Enum_Enclosure_shape = "bracket"
	Enum_Enclosure_shape_Inverted_bracket Enum_Enclosure_shape = "inverted-bracket"
	Enum_Enclosure_shape_Triangle         Enum_Enclosure_shape = "triangle"
	Enum_Enclosure_shape_Diamond          Enum_Enclosure_shape = "diamond"
	Enum_Enclosure_shape_Pentagon         Enum_Enclosure_shape = "pentagon"
	Enum_Enclosure_shape_Hexagon          Enum_Enclosure_shape = "hexagon"
	Enum_Enclosure_shape_Heptagon         Enum_Enclosure_shape = "heptagon"
	Enum_Enclosure_shape_Octagon          Enum_Enclosure_shape = "octagon"
	Enum_Enclosure_shape_Nonagon          Enum_Enclosure_shape = "nonagon"
	Enum_Enclosure_shape_Decagon          Enum_Enclosure_shape = "decagon"
	Enum_Enclosure_shape_None             Enum_Enclosure_shape = "none"
)

// From xsd simple type with enumerate restriction "ending-number"
type Enum_Ending_number string

const ()

// From xsd simple type with enumerate restriction "fan"
type Enum_Fan string

const (
	Enum_Fan_Accel Enum_Fan = "accel"
	Enum_Fan_Rit   Enum_Fan = "rit"
	Enum_Fan_None  Enum_Fan = "none"
)

// From xsd simple type with enumerate restriction "fermata-shape"
type Enum_Fermata_shape string

const (
	Enum_Fermata_shape_Normal        Enum_Fermata_shape = "normal"
	Enum_Fermata_shape_Angled        Enum_Fermata_shape = "angled"
	Enum_Fermata_shape_Square        Enum_Fermata_shape = "square"
	Enum_Fermata_shape_Double_angled Enum_Fermata_shape = "double-angled"
	Enum_Fermata_shape_Double_square Enum_Fermata_shape = "double-square"
	Enum_Fermata_shape_Double_dot    Enum_Fermata_shape = "double-dot"
	Enum_Fermata_shape_Half_curve    Enum_Fermata_shape = "half-curve"
	Enum_Fermata_shape_Curlew        Enum_Fermata_shape = "curlew"
	Enum_Fermata_shape_              Enum_Fermata_shape = ""
)

// From xsd simple type with enumerate restriction "font-style"
type Enum_Font_style string

const (
	Enum_Font_style_Normal Enum_Font_style = "normal"
	Enum_Font_style_Italic Enum_Font_style = "italic"
)

// From xsd simple type with enumerate restriction "font-weight"
type Enum_Font_weight string

const (
	Enum_Font_weight_Normal Enum_Font_weight = "normal"
	Enum_Font_weight_Bold   Enum_Font_weight = "bold"
)

// From xsd simple type with enumerate restriction "glass-value"
type Enum_Glass_value string

const (
	Enum_Glass_value_Glass_harmonica Enum_Glass_value = "glass harmonica"
	Enum_Glass_value_Glass_harp      Enum_Glass_value = "glass harp"
	Enum_Glass_value_Wind_chimes     Enum_Glass_value = "wind chimes"
)

// From xsd simple type with enumerate restriction "glyph-type"
type Enum_Glyph_type string

const ()

// From xsd simple type with enumerate restriction "group-barline-value"
type Enum_Group_barline_value string

const (
	Enum_Group_barline_value_Yes          Enum_Group_barline_value = "yes"
	Enum_Group_barline_value_No           Enum_Group_barline_value = "no"
	Enum_Group_barline_value_Mensurstrich Enum_Group_barline_value = "Mensurstrich"
)

// From xsd simple type with enumerate restriction "group-symbol-value"
type Enum_Group_symbol_value string

const (
	Enum_Group_symbol_value_None    Enum_Group_symbol_value = "none"
	Enum_Group_symbol_value_Brace   Enum_Group_symbol_value = "brace"
	Enum_Group_symbol_value_Line    Enum_Group_symbol_value = "line"
	Enum_Group_symbol_value_Bracket Enum_Group_symbol_value = "bracket"
	Enum_Group_symbol_value_Square  Enum_Group_symbol_value = "square"
)

// From xsd simple type with enumerate restriction "handbell-value"
type Enum_Handbell_value string

const (
	Enum_Handbell_value_Belltree         Enum_Handbell_value = "belltree"
	Enum_Handbell_value_Damp             Enum_Handbell_value = "damp"
	Enum_Handbell_value_Echo             Enum_Handbell_value = "echo"
	Enum_Handbell_value_Gyro             Enum_Handbell_value = "gyro"
	Enum_Handbell_value_Hand_martellato  Enum_Handbell_value = "hand martellato"
	Enum_Handbell_value_Mallet_lift      Enum_Handbell_value = "mallet lift"
	Enum_Handbell_value_Mallet_table     Enum_Handbell_value = "mallet table"
	Enum_Handbell_value_Martellato       Enum_Handbell_value = "martellato"
	Enum_Handbell_value_Martellato_lift  Enum_Handbell_value = "martellato lift"
	Enum_Handbell_value_Muted_martellato Enum_Handbell_value = "muted martellato"
	Enum_Handbell_value_Pluck_lift       Enum_Handbell_value = "pluck lift"
	Enum_Handbell_value_Swing            Enum_Handbell_value = "swing"
)

// From xsd simple type with enumerate restriction "harmon-closed-location"
type Enum_Harmon_closed_location string

const (
	Enum_Harmon_closed_location_Right  Enum_Harmon_closed_location = "right"
	Enum_Harmon_closed_location_Bottom Enum_Harmon_closed_location = "bottom"
	Enum_Harmon_closed_location_Left   Enum_Harmon_closed_location = "left"
	Enum_Harmon_closed_location_Top    Enum_Harmon_closed_location = "top"
)

// From xsd simple type with enumerate restriction "harmon-closed-value"
type Enum_Harmon_closed_value string

const (
	Enum_Harmon_closed_value_Yes  Enum_Harmon_closed_value = "yes"
	Enum_Harmon_closed_value_No   Enum_Harmon_closed_value = "no"
	Enum_Harmon_closed_value_Half Enum_Harmon_closed_value = "half"
)

// From xsd simple type with enumerate restriction "harmony-arrangement"
type Enum_Harmony_arrangement string

const (
	Enum_Harmony_arrangement_Vertical   Enum_Harmony_arrangement = "vertical"
	Enum_Harmony_arrangement_Horizontal Enum_Harmony_arrangement = "horizontal"
	Enum_Harmony_arrangement_Diagonal   Enum_Harmony_arrangement = "diagonal"
)

// From xsd simple type with enumerate restriction "harmony-type"
type Enum_Harmony_type string

const (
	Enum_Harmony_type_Explicit  Enum_Harmony_type = "explicit"
	Enum_Harmony_type_Implied   Enum_Harmony_type = "implied"
	Enum_Harmony_type_Alternate Enum_Harmony_type = "alternate"
)

// From xsd simple type with enumerate restriction "hole-closed-location"
type Enum_Hole_closed_location string

const (
	Enum_Hole_closed_location_Right  Enum_Hole_closed_location = "right"
	Enum_Hole_closed_location_Bottom Enum_Hole_closed_location = "bottom"
	Enum_Hole_closed_location_Left   Enum_Hole_closed_location = "left"
	Enum_Hole_closed_location_Top    Enum_Hole_closed_location = "top"
)

// From xsd simple type with enumerate restriction "hole-closed-value"
type Enum_Hole_closed_value string

const (
	Enum_Hole_closed_value_Yes  Enum_Hole_closed_value = "yes"
	Enum_Hole_closed_value_No   Enum_Hole_closed_value = "no"
	Enum_Hole_closed_value_Half Enum_Hole_closed_value = "half"
)

// From xsd simple type with enumerate restriction "kind-value"
type Enum_Kind_value string

const (
	Enum_Kind_value_Major              Enum_Kind_value = "major"
	Enum_Kind_value_Minor              Enum_Kind_value = "minor"
	Enum_Kind_value_Augmented          Enum_Kind_value = "augmented"
	Enum_Kind_value_Diminished         Enum_Kind_value = "diminished"
	Enum_Kind_value_Dominant           Enum_Kind_value = "dominant"
	Enum_Kind_value_Major_seventh      Enum_Kind_value = "major-seventh"
	Enum_Kind_value_Minor_seventh      Enum_Kind_value = "minor-seventh"
	Enum_Kind_value_Diminished_seventh Enum_Kind_value = "diminished-seventh"
	Enum_Kind_value_Augmented_seventh  Enum_Kind_value = "augmented-seventh"
	Enum_Kind_value_Half_diminished    Enum_Kind_value = "half-diminished"
	Enum_Kind_value_Major_minor        Enum_Kind_value = "major-minor"
	Enum_Kind_value_Major_sixth        Enum_Kind_value = "major-sixth"
	Enum_Kind_value_Minor_sixth        Enum_Kind_value = "minor-sixth"
	Enum_Kind_value_Dominant_ninth     Enum_Kind_value = "dominant-ninth"
	Enum_Kind_value_Major_ninth        Enum_Kind_value = "major-ninth"
	Enum_Kind_value_Minor_ninth        Enum_Kind_value = "minor-ninth"
	Enum_Kind_value_Dominant_11th      Enum_Kind_value = "dominant-11th"
	Enum_Kind_value_Major_11th         Enum_Kind_value = "major-11th"
	Enum_Kind_value_Minor_11th         Enum_Kind_value = "minor-11th"
	Enum_Kind_value_Dominant_13th      Enum_Kind_value = "dominant-13th"
	Enum_Kind_value_Major_13th         Enum_Kind_value = "major-13th"
	Enum_Kind_value_Minor_13th         Enum_Kind_value = "minor-13th"
	Enum_Kind_value_Suspended_second   Enum_Kind_value = "suspended-second"
	Enum_Kind_value_Suspended_fourth   Enum_Kind_value = "suspended-fourth"
	Enum_Kind_value_Neapolitan         Enum_Kind_value = "Neapolitan"
	Enum_Kind_value_Italian            Enum_Kind_value = "Italian"
	Enum_Kind_value_French             Enum_Kind_value = "French"
	Enum_Kind_value_German             Enum_Kind_value = "German"
	Enum_Kind_value_Pedal              Enum_Kind_value = "pedal"
	Enum_Kind_value_Power              Enum_Kind_value = "power"
	Enum_Kind_value_Tristan            Enum_Kind_value = "Tristan"
	Enum_Kind_value_Other              Enum_Kind_value = "other"
	Enum_Kind_value_None               Enum_Kind_value = "none"
)

// From xsd simple type with enumerate restriction "left-center-right"
type Enum_Left_center_right string

const (
	Enum_Left_center_right_Left   Enum_Left_center_right = "left"
	Enum_Left_center_right_Center Enum_Left_center_right = "center"
	Enum_Left_center_right_Right  Enum_Left_center_right = "right"
)

// From xsd simple type with enumerate restriction "left-right"
type Enum_Left_right string

const (
	Enum_Left_right_Left  Enum_Left_right = "left"
	Enum_Left_right_Right Enum_Left_right = "right"
)

// From xsd simple type with enumerate restriction "line-end"
type Enum_Line_end string

const (
	Enum_Line_end_Up    Enum_Line_end = "up"
	Enum_Line_end_Down  Enum_Line_end = "down"
	Enum_Line_end_Both  Enum_Line_end = "both"
	Enum_Line_end_Arrow Enum_Line_end = "arrow"
	Enum_Line_end_None  Enum_Line_end = "none"
)

// From xsd simple type with enumerate restriction "line-length"
type Enum_Line_length string

const (
	Enum_Line_length_Short  Enum_Line_length = "short"
	Enum_Line_length_Medium Enum_Line_length = "medium"
	Enum_Line_length_Long   Enum_Line_length = "long"
)

// From xsd simple type with enumerate restriction "line-shape"
type Enum_Line_shape string

const (
	Enum_Line_shape_Straight Enum_Line_shape = "straight"
	Enum_Line_shape_Curved   Enum_Line_shape = "curved"
)

// From xsd simple type with enumerate restriction "line-type"
type Enum_Line_type string

const (
	Enum_Line_type_Solid  Enum_Line_type = "solid"
	Enum_Line_type_Dashed Enum_Line_type = "dashed"
	Enum_Line_type_Dotted Enum_Line_type = "dotted"
	Enum_Line_type_Wavy   Enum_Line_type = "wavy"
)

// From xsd simple type with enumerate restriction "line-width-type"
type Enum_Line_width_type string

const ()

// From xsd simple type with enumerate restriction "margin-type"
type Enum_Margin_type string

const (
	Enum_Margin_type_Odd  Enum_Margin_type = "odd"
	Enum_Margin_type_Even Enum_Margin_type = "even"
	Enum_Margin_type_Both Enum_Margin_type = "both"
)

// From xsd simple type with enumerate restriction "measure-numbering-value"
type Enum_Measure_numbering_value string

const (
	Enum_Measure_numbering_value_None    Enum_Measure_numbering_value = "none"
	Enum_Measure_numbering_value_Measure Enum_Measure_numbering_value = "measure"
	Enum_Measure_numbering_value_System  Enum_Measure_numbering_value = "system"
)

// From xsd simple type with enumerate restriction "measure-text"
type Enum_Measure_text string

const ()

// From xsd simple type with enumerate restriction "membrane-value"
type Enum_Membrane_value string

const (
	Enum_Membrane_value_Bass_drum             Enum_Membrane_value = "bass drum"
	Enum_Membrane_value_Bass_drum_on_side     Enum_Membrane_value = "bass drum on side"
	Enum_Membrane_value_Bongos                Enum_Membrane_value = "bongos"
	Enum_Membrane_value_Chinese_tomtom        Enum_Membrane_value = "Chinese tomtom"
	Enum_Membrane_value_Conga_drum            Enum_Membrane_value = "conga drum"
	Enum_Membrane_value_Cuica                 Enum_Membrane_value = "cuica"
	Enum_Membrane_value_Goblet_drum           Enum_Membrane_value = "goblet drum"
	Enum_Membrane_value_Indo_American_tomtom  Enum_Membrane_value = "Indo-American tomtom"
	Enum_Membrane_value_Japanese_tomtom       Enum_Membrane_value = "Japanese tomtom"
	Enum_Membrane_value_Military_drum         Enum_Membrane_value = "military drum"
	Enum_Membrane_value_Snare_drum            Enum_Membrane_value = "snare drum"
	Enum_Membrane_value_Snare_drum_snares_off Enum_Membrane_value = "snare drum snares off"
	Enum_Membrane_value_Tabla                 Enum_Membrane_value = "tabla"
	Enum_Membrane_value_Tambourine            Enum_Membrane_value = "tambourine"
	Enum_Membrane_value_Tenor_drum            Enum_Membrane_value = "tenor drum"
	Enum_Membrane_value_Timbales              Enum_Membrane_value = "timbales"
	Enum_Membrane_value_Tomtom                Enum_Membrane_value = "tomtom"
)

// From xsd simple type with enumerate restriction "metal-value"
type Enum_Metal_value string

const (
	Enum_Metal_value_Agogo               Enum_Metal_value = "agogo"
	Enum_Metal_value_Almglocken          Enum_Metal_value = "almglocken"
	Enum_Metal_value_Bell                Enum_Metal_value = "bell"
	Enum_Metal_value_Bell_plate          Enum_Metal_value = "bell plate"
	Enum_Metal_value_Bell_tree           Enum_Metal_value = "bell tree"
	Enum_Metal_value_Brake_drum          Enum_Metal_value = "brake drum"
	Enum_Metal_value_Cencerro            Enum_Metal_value = "cencerro"
	Enum_Metal_value_Chain_rattle        Enum_Metal_value = "chain rattle"
	Enum_Metal_value_Chinese_cymbal      Enum_Metal_value = "Chinese cymbal"
	Enum_Metal_value_Cowbell             Enum_Metal_value = "cowbell"
	Enum_Metal_value_Crash_cymbals       Enum_Metal_value = "crash cymbals"
	Enum_Metal_value_Crotale             Enum_Metal_value = "crotale"
	Enum_Metal_value_Cymbal_tongs        Enum_Metal_value = "cymbal tongs"
	Enum_Metal_value_Domed_gong          Enum_Metal_value = "domed gong"
	Enum_Metal_value_Finger_cymbals      Enum_Metal_value = "finger cymbals"
	Enum_Metal_value_Flexatone           Enum_Metal_value = "flexatone"
	Enum_Metal_value_Gong                Enum_Metal_value = "gong"
	Enum_Metal_value_Hi_hat              Enum_Metal_value = "hi-hat"
	Enum_Metal_value_High_hat_cymbals    Enum_Metal_value = "high-hat cymbals"
	Enum_Metal_value_Handbell            Enum_Metal_value = "handbell"
	Enum_Metal_value_Jaw_harp            Enum_Metal_value = "jaw harp"
	Enum_Metal_value_Jingle_bells        Enum_Metal_value = "jingle bells"
	Enum_Metal_value_Musical_saw         Enum_Metal_value = "musical saw"
	Enum_Metal_value_Shell_bells         Enum_Metal_value = "shell bells"
	Enum_Metal_value_Sistrum             Enum_Metal_value = "sistrum"
	Enum_Metal_value_Sizzle_cymbal       Enum_Metal_value = "sizzle cymbal"
	Enum_Metal_value_Sleigh_bells        Enum_Metal_value = "sleigh bells"
	Enum_Metal_value_Suspended_cymbal    Enum_Metal_value = "suspended cymbal"
	Enum_Metal_value_Tam_tam             Enum_Metal_value = "tam tam"
	Enum_Metal_value_Tam_tam_with_beater Enum_Metal_value = "tam tam with beater"
	Enum_Metal_value_Triangle            Enum_Metal_value = "triangle"
	Enum_Metal_value_Vietnamese_hat      Enum_Metal_value = "Vietnamese hat"
)

// From xsd simple type with enumerate restriction "mode"
type Enum_Mode string

const ()

// From xsd simple type with enumerate restriction "mute"
type Enum_Mute string

const (
	Enum_Mute_On             Enum_Mute = "on"
	Enum_Mute_Off            Enum_Mute = "off"
	Enum_Mute_Straight       Enum_Mute = "straight"
	Enum_Mute_Cup            Enum_Mute = "cup"
	Enum_Mute_Harmon_no_stem Enum_Mute = "harmon-no-stem"
	Enum_Mute_Harmon_stem    Enum_Mute = "harmon-stem"
	Enum_Mute_Bucket         Enum_Mute = "bucket"
	Enum_Mute_Plunger        Enum_Mute = "plunger"
	Enum_Mute_Hat            Enum_Mute = "hat"
	Enum_Mute_Solotone       Enum_Mute = "solotone"
	Enum_Mute_Practice       Enum_Mute = "practice"
	Enum_Mute_Stop_mute      Enum_Mute = "stop-mute"
	Enum_Mute_Stop_hand      Enum_Mute = "stop-hand"
	Enum_Mute_Echo           Enum_Mute = "echo"
	Enum_Mute_Palm           Enum_Mute = "palm"
)

// From xsd simple type with enumerate restriction "note-size-type"
type Enum_Note_size_type string

const (
	Enum_Note_size_type_Cue       Enum_Note_size_type = "cue"
	Enum_Note_size_type_Grace     Enum_Note_size_type = "grace"
	Enum_Note_size_type_Grace_cue Enum_Note_size_type = "grace-cue"
	Enum_Note_size_type_Large     Enum_Note_size_type = "large"
)

// From xsd simple type with enumerate restriction "note-type-value"
type Enum_Note_type_value string

const (
	Enum_Note_type_value_1024th  Enum_Note_type_value = "1024th"
	Enum_Note_type_value_512th   Enum_Note_type_value = "512th"
	Enum_Note_type_value_256th   Enum_Note_type_value = "256th"
	Enum_Note_type_value_128th   Enum_Note_type_value = "128th"
	Enum_Note_type_value_64th    Enum_Note_type_value = "64th"
	Enum_Note_type_value_32nd    Enum_Note_type_value = "32nd"
	Enum_Note_type_value_16th    Enum_Note_type_value = "16th"
	Enum_Note_type_value_Eighth  Enum_Note_type_value = "eighth"
	Enum_Note_type_value_Quarter Enum_Note_type_value = "quarter"
	Enum_Note_type_value_Half    Enum_Note_type_value = "half"
	Enum_Note_type_value_Whole   Enum_Note_type_value = "whole"
	Enum_Note_type_value_Breve   Enum_Note_type_value = "breve"
	Enum_Note_type_value_Long    Enum_Note_type_value = "long"
	Enum_Note_type_value_Maxima  Enum_Note_type_value = "maxima"
)

// From xsd simple type with enumerate restriction "notehead-value"
type Enum_Notehead_value string

const (
	Enum_Notehead_value_Slash             Enum_Notehead_value = "slash"
	Enum_Notehead_value_Triangle          Enum_Notehead_value = "triangle"
	Enum_Notehead_value_Diamond           Enum_Notehead_value = "diamond"
	Enum_Notehead_value_Square            Enum_Notehead_value = "square"
	Enum_Notehead_value_Cross             Enum_Notehead_value = "cross"
	Enum_Notehead_value_X                 Enum_Notehead_value = "x"
	Enum_Notehead_value_Circle_x          Enum_Notehead_value = "circle-x"
	Enum_Notehead_value_Inverted_triangle Enum_Notehead_value = "inverted triangle"
	Enum_Notehead_value_Arrow_down        Enum_Notehead_value = "arrow down"
	Enum_Notehead_value_Arrow_up          Enum_Notehead_value = "arrow up"
	Enum_Notehead_value_Circled           Enum_Notehead_value = "circled"
	Enum_Notehead_value_Slashed           Enum_Notehead_value = "slashed"
	Enum_Notehead_value_Back_slashed      Enum_Notehead_value = "back slashed"
	Enum_Notehead_value_Normal            Enum_Notehead_value = "normal"
	Enum_Notehead_value_Cluster           Enum_Notehead_value = "cluster"
	Enum_Notehead_value_Circle_dot        Enum_Notehead_value = "circle dot"
	Enum_Notehead_value_Left_triangle     Enum_Notehead_value = "left triangle"
	Enum_Notehead_value_Rectangle         Enum_Notehead_value = "rectangle"
	Enum_Notehead_value_None              Enum_Notehead_value = "none"
	Enum_Notehead_value_Do                Enum_Notehead_value = "do"
	Enum_Notehead_value_Re                Enum_Notehead_value = "re"
	Enum_Notehead_value_Mi                Enum_Notehead_value = "mi"
	Enum_Notehead_value_Fa                Enum_Notehead_value = "fa"
	Enum_Notehead_value_Fa_up             Enum_Notehead_value = "fa up"
	Enum_Notehead_value_So                Enum_Notehead_value = "so"
	Enum_Notehead_value_La                Enum_Notehead_value = "la"
	Enum_Notehead_value_Ti                Enum_Notehead_value = "ti"
	Enum_Notehead_value_Other             Enum_Notehead_value = "other"
)

// From xsd simple type with enumerate restriction "numeral-mode"
type Enum_Numeral_mode string

const (
	Enum_Numeral_mode_Major          Enum_Numeral_mode = "major"
	Enum_Numeral_mode_Minor          Enum_Numeral_mode = "minor"
	Enum_Numeral_mode_Natural_minor  Enum_Numeral_mode = "natural minor"
	Enum_Numeral_mode_Melodic_minor  Enum_Numeral_mode = "melodic minor"
	Enum_Numeral_mode_Harmonic_minor Enum_Numeral_mode = "harmonic minor"
)

// From xsd simple type with enumerate restriction "on-off"
type Enum_On_off string

const (
	Enum_On_off_On  Enum_On_off = "on"
	Enum_On_off_Off Enum_On_off = "off"
)

// From xsd simple type with enumerate restriction "over-under"
type Enum_Over_under string

const (
	Enum_Over_under_Over  Enum_Over_under = "over"
	Enum_Over_under_Under Enum_Over_under = "under"
)

// From xsd simple type with enumerate restriction "pedal-type"
type Enum_Pedal_type string

const (
	Enum_Pedal_type_Start       Enum_Pedal_type = "start"
	Enum_Pedal_type_Stop        Enum_Pedal_type = "stop"
	Enum_Pedal_type_Sostenuto   Enum_Pedal_type = "sostenuto"
	Enum_Pedal_type_Change      Enum_Pedal_type = "change"
	Enum_Pedal_type_Continue    Enum_Pedal_type = "continue"
	Enum_Pedal_type_Discontinue Enum_Pedal_type = "discontinue"
	Enum_Pedal_type_Resume      Enum_Pedal_type = "resume"
)

// From xsd simple type with enumerate restriction "pitched-value"
type Enum_Pitched_value string

const (
	Enum_Pitched_value_Celesta        Enum_Pitched_value = "celesta"
	Enum_Pitched_value_Chimes         Enum_Pitched_value = "chimes"
	Enum_Pitched_value_Glockenspiel   Enum_Pitched_value = "glockenspiel"
	Enum_Pitched_value_Lithophone     Enum_Pitched_value = "lithophone"
	Enum_Pitched_value_Mallet         Enum_Pitched_value = "mallet"
	Enum_Pitched_value_Marimba        Enum_Pitched_value = "marimba"
	Enum_Pitched_value_Steel_drums    Enum_Pitched_value = "steel drums"
	Enum_Pitched_value_Tubaphone      Enum_Pitched_value = "tubaphone"
	Enum_Pitched_value_Tubular_chimes Enum_Pitched_value = "tubular chimes"
	Enum_Pitched_value_Vibraphone     Enum_Pitched_value = "vibraphone"
	Enum_Pitched_value_Xylophone      Enum_Pitched_value = "xylophone"
)

// From xsd simple type with enumerate restriction "principal-voice-symbol"
type Enum_Principal_voice_symbol string

const (
	Enum_Principal_voice_symbol_Hauptstimme Enum_Principal_voice_symbol = "Hauptstimme"
	Enum_Principal_voice_symbol_Nebenstimme Enum_Principal_voice_symbol = "Nebenstimme"
	Enum_Principal_voice_symbol_Plain       Enum_Principal_voice_symbol = "plain"
	Enum_Principal_voice_symbol_None        Enum_Principal_voice_symbol = "none"
)

// From xsd simple type with enumerate restriction "right-left-middle"
type Enum_Right_left_middle string

const (
	Enum_Right_left_middle_Right  Enum_Right_left_middle = "right"
	Enum_Right_left_middle_Left   Enum_Right_left_middle = "left"
	Enum_Right_left_middle_Middle Enum_Right_left_middle = "middle"
)

// From xsd simple type with enumerate restriction "semi-pitched"
type Enum_Semi_pitched string

const (
	Enum_Semi_pitched_High        Enum_Semi_pitched = "high"
	Enum_Semi_pitched_Medium_high Enum_Semi_pitched = "medium-high"
	Enum_Semi_pitched_Medium      Enum_Semi_pitched = "medium"
	Enum_Semi_pitched_Medium_low  Enum_Semi_pitched = "medium-low"
	Enum_Semi_pitched_Low         Enum_Semi_pitched = "low"
	Enum_Semi_pitched_Very_low    Enum_Semi_pitched = "very-low"
)

// From xsd simple type with enumerate restriction "show-frets"
type Enum_Show_frets string

const (
	Enum_Show_frets_Numbers Enum_Show_frets = "numbers"
	Enum_Show_frets_Letters Enum_Show_frets = "letters"
)

// From xsd simple type with enumerate restriction "show-tuplet"
type Enum_Show_tuplet string

const (
	Enum_Show_tuplet_Actual Enum_Show_tuplet = "actual"
	Enum_Show_tuplet_Both   Enum_Show_tuplet = "both"
	Enum_Show_tuplet_None   Enum_Show_tuplet = "none"
)

// From xsd simple type with enumerate restriction "staff-divide-symbol"
type Enum_Staff_divide_symbol string

const (
	Enum_Staff_divide_symbol_Down    Enum_Staff_divide_symbol = "down"
	Enum_Staff_divide_symbol_Up      Enum_Staff_divide_symbol = "up"
	Enum_Staff_divide_symbol_Up_down Enum_Staff_divide_symbol = "up-down"
)

// From xsd simple type with enumerate restriction "staff-type"
type Enum_Staff_type string

const (
	Enum_Staff_type_Ossia     Enum_Staff_type = "ossia"
	Enum_Staff_type_Editorial Enum_Staff_type = "editorial"
	Enum_Staff_type_Cue       Enum_Staff_type = "cue"
	Enum_Staff_type_Alternate Enum_Staff_type = "alternate"
	Enum_Staff_type_Regular   Enum_Staff_type = "regular"
)

// From xsd simple type with enumerate restriction "start-note"
type Enum_Start_note string

const (
	Enum_Start_note_Upper Enum_Start_note = "upper"
	Enum_Start_note_Main  Enum_Start_note = "main"
	Enum_Start_note_Below Enum_Start_note = "below"
)

// From xsd simple type with enumerate restriction "start-stop"
type Enum_Start_stop string

const (
	Enum_Start_stop_Start Enum_Start_stop = "start"
	Enum_Start_stop_Stop  Enum_Start_stop = "stop"
)

// From xsd simple type with enumerate restriction "start-stop-continue"
type Enum_Start_stop_continue string

const (
	Enum_Start_stop_continue_Start    Enum_Start_stop_continue = "start"
	Enum_Start_stop_continue_Stop     Enum_Start_stop_continue = "stop"
	Enum_Start_stop_continue_Continue Enum_Start_stop_continue = "continue"
)

// From xsd simple type with enumerate restriction "start-stop-discontinue"
type Enum_Start_stop_discontinue string

const (
	Enum_Start_stop_discontinue_Start       Enum_Start_stop_discontinue = "start"
	Enum_Start_stop_discontinue_Stop        Enum_Start_stop_discontinue = "stop"
	Enum_Start_stop_discontinue_Discontinue Enum_Start_stop_discontinue = "discontinue"
)

// From xsd simple type with enumerate restriction "start-stop-single"
type Enum_Start_stop_single string

const (
	Enum_Start_stop_single_Start  Enum_Start_stop_single = "start"
	Enum_Start_stop_single_Stop   Enum_Start_stop_single = "stop"
	Enum_Start_stop_single_Single Enum_Start_stop_single = "single"
)

// From xsd simple type with enumerate restriction "stem-value"
type Enum_Stem_value string

const (
	Enum_Stem_value_Down   Enum_Stem_value = "down"
	Enum_Stem_value_Up     Enum_Stem_value = "up"
	Enum_Stem_value_Double Enum_Stem_value = "double"
	Enum_Stem_value_None   Enum_Stem_value = "none"
)

// From xsd simple type with enumerate restriction "step"
type Enum_Step string

const (
	Enum_Step_A Enum_Step = "A"
	Enum_Step_B Enum_Step = "B"
	Enum_Step_C Enum_Step = "C"
	Enum_Step_D Enum_Step = "D"
	Enum_Step_E Enum_Step = "E"
	Enum_Step_F Enum_Step = "F"
	Enum_Step_G Enum_Step = "G"
)

// From xsd simple type with enumerate restriction "stick-location"
type Enum_Stick_location string

const (
	Enum_Stick_location_Center      Enum_Stick_location = "center"
	Enum_Stick_location_Rim         Enum_Stick_location = "rim"
	Enum_Stick_location_Cymbal_bell Enum_Stick_location = "cymbal bell"
	Enum_Stick_location_Cymbal_edge Enum_Stick_location = "cymbal edge"
)

// From xsd simple type with enumerate restriction "stick-material"
type Enum_Stick_material string

const (
	Enum_Stick_material_Soft   Enum_Stick_material = "soft"
	Enum_Stick_material_Medium Enum_Stick_material = "medium"
	Enum_Stick_material_Hard   Enum_Stick_material = "hard"
	Enum_Stick_material_Shaded Enum_Stick_material = "shaded"
	Enum_Stick_material_X      Enum_Stick_material = "x"
)

// From xsd simple type with enumerate restriction "stick-type"
type Enum_Stick_type string

const (
	Enum_Stick_type_Bass_drum        Enum_Stick_type = "bass drum"
	Enum_Stick_type_Double_bass_drum Enum_Stick_type = "double bass drum"
	Enum_Stick_type_Glockenspiel     Enum_Stick_type = "glockenspiel"
	Enum_Stick_type_Gum              Enum_Stick_type = "gum"
	Enum_Stick_type_Hammer           Enum_Stick_type = "hammer"
	Enum_Stick_type_Superball        Enum_Stick_type = "superball"
	Enum_Stick_type_Timpani          Enum_Stick_type = "timpani"
	Enum_Stick_type_Wound            Enum_Stick_type = "wound"
	Enum_Stick_type_Xylophone        Enum_Stick_type = "xylophone"
	Enum_Stick_type_Yarn             Enum_Stick_type = "yarn"
)

// From xsd simple type with enumerate restriction "syllabic"
type Enum_Syllabic string

const (
	Enum_Syllabic_Single Enum_Syllabic = "single"
	Enum_Syllabic_Begin  Enum_Syllabic = "begin"
	Enum_Syllabic_End    Enum_Syllabic = "end"
	Enum_Syllabic_Middle Enum_Syllabic = "middle"
)

// From xsd simple type with enumerate restriction "symbol-size"
type Enum_Symbol_size string

const (
	Enum_Symbol_size_Full      Enum_Symbol_size = "full"
	Enum_Symbol_size_Cue       Enum_Symbol_size = "cue"
	Enum_Symbol_size_Grace_cue Enum_Symbol_size = "grace-cue"
	Enum_Symbol_size_Large     Enum_Symbol_size = "large"
)

// From xsd simple type with enumerate restriction "sync-type"
type Enum_Sync_type string

const (
	Enum_Sync_type_None         Enum_Sync_type = "none"
	Enum_Sync_type_Tempo        Enum_Sync_type = "tempo"
	Enum_Sync_type_Mostly_tempo Enum_Sync_type = "mostly-tempo"
	Enum_Sync_type_Mostly_event Enum_Sync_type = "mostly-event"
	Enum_Sync_type_Event        Enum_Sync_type = "event"
	Enum_Sync_type_Always_event Enum_Sync_type = "always-event"
)

// From xsd simple type with enumerate restriction "system-relation-number"
type Enum_System_relation_number string

const (
	Enum_System_relation_number_Only_top    Enum_System_relation_number = "only-top"
	Enum_System_relation_number_Only_bottom Enum_System_relation_number = "only-bottom"
	Enum_System_relation_number_Also_top    Enum_System_relation_number = "also-top"
	Enum_System_relation_number_Also_bottom Enum_System_relation_number = "also-bottom"
	Enum_System_relation_number_None        Enum_System_relation_number = "none"
)

// From xsd simple type with enumerate restriction "tap-hand"
type Enum_Tap_hand string

const (
	Enum_Tap_hand_Left  Enum_Tap_hand = "left"
	Enum_Tap_hand_Right Enum_Tap_hand = "right"
)

// From xsd simple type with enumerate restriction "text-direction"
type Enum_Text_direction string

const (
	Enum_Text_direction_Ltr Enum_Text_direction = "ltr"
	Enum_Text_direction_Rtl Enum_Text_direction = "rtl"
	Enum_Text_direction_Lro Enum_Text_direction = "lro"
	Enum_Text_direction_Rlo Enum_Text_direction = "rlo"
)

// From xsd simple type with enumerate restriction "tied-type"
type Enum_Tied_type string

const (
	Enum_Tied_type_Start    Enum_Tied_type = "start"
	Enum_Tied_type_Stop     Enum_Tied_type = "stop"
	Enum_Tied_type_Continue Enum_Tied_type = "continue"
	Enum_Tied_type_Let_ring Enum_Tied_type = "let-ring"
)

// From xsd simple type with enumerate restriction "time-only"
type Enum_Time_only string

const ()

// From xsd simple type with enumerate restriction "time-relation"
type Enum_Time_relation string

const (
	Enum_Time_relation_Parentheses Enum_Time_relation = "parentheses"
	Enum_Time_relation_Bracket     Enum_Time_relation = "bracket"
	Enum_Time_relation_Equals      Enum_Time_relation = "equals"
	Enum_Time_relation_Slash       Enum_Time_relation = "slash"
	Enum_Time_relation_Space       Enum_Time_relation = "space"
	Enum_Time_relation_Hyphen      Enum_Time_relation = "hyphen"
)

// From xsd simple type with enumerate restriction "time-separator"
type Enum_Time_separator string

const (
	Enum_Time_separator_None       Enum_Time_separator = "none"
	Enum_Time_separator_Horizontal Enum_Time_separator = "horizontal"
	Enum_Time_separator_Diagonal   Enum_Time_separator = "diagonal"
	Enum_Time_separator_Vertical   Enum_Time_separator = "vertical"
	Enum_Time_separator_Adjacent   Enum_Time_separator = "adjacent"
)

// From xsd simple type with enumerate restriction "time-symbol"
type Enum_Time_symbol string

const (
	Enum_Time_symbol_Common        Enum_Time_symbol = "common"
	Enum_Time_symbol_Cut           Enum_Time_symbol = "cut"
	Enum_Time_symbol_Single_number Enum_Time_symbol = "single-number"
	Enum_Time_symbol_Note          Enum_Time_symbol = "note"
	Enum_Time_symbol_Dotted_note   Enum_Time_symbol = "dotted-note"
	Enum_Time_symbol_Normal        Enum_Time_symbol = "normal"
)

// From xsd simple type with enumerate restriction "tip-direction"
type Enum_Tip_direction string

const (
	Enum_Tip_direction_Up        Enum_Tip_direction = "up"
	Enum_Tip_direction_Down      Enum_Tip_direction = "down"
	Enum_Tip_direction_Left      Enum_Tip_direction = "left"
	Enum_Tip_direction_Right     Enum_Tip_direction = "right"
	Enum_Tip_direction_Northwest Enum_Tip_direction = "northwest"
	Enum_Tip_direction_Northeast Enum_Tip_direction = "northeast"
	Enum_Tip_direction_Southeast Enum_Tip_direction = "southeast"
	Enum_Tip_direction_Southwest Enum_Tip_direction = "southwest"
)

// From xsd simple type with enumerate restriction "top-bottom"
type Enum_Top_bottom string

const (
	Enum_Top_bottom_Top    Enum_Top_bottom = "top"
	Enum_Top_bottom_Bottom Enum_Top_bottom = "bottom"
)

// From xsd simple type with enumerate restriction "tremolo-type"
type Enum_Tremolo_type string

const (
	Enum_Tremolo_type_Start      Enum_Tremolo_type = "start"
	Enum_Tremolo_type_Stop       Enum_Tremolo_type = "stop"
	Enum_Tremolo_type_Single     Enum_Tremolo_type = "single"
	Enum_Tremolo_type_Unmeasured Enum_Tremolo_type = "unmeasured"
)

// From xsd simple type with enumerate restriction "trill-step"
type Enum_Trill_step string

const (
	Enum_Trill_step_Whole  Enum_Trill_step = "whole"
	Enum_Trill_step_Half   Enum_Trill_step = "half"
	Enum_Trill_step_Unison Enum_Trill_step = "unison"
)

// From xsd simple type with enumerate restriction "two-note-turn"
type Enum_Two_note_turn string

const (
	Enum_Two_note_turn_Whole Enum_Two_note_turn = "whole"
	Enum_Two_note_turn_Half  Enum_Two_note_turn = "half"
	Enum_Two_note_turn_None  Enum_Two_note_turn = "none"
)

// From xsd simple type with enumerate restriction "up-down"
type Enum_Up_down string

const (
	Enum_Up_down_Up   Enum_Up_down = "up"
	Enum_Up_down_Down Enum_Up_down = "down"
)

// From xsd simple type with enumerate restriction "up-down-stop-continue"
type Enum_Up_down_stop_continue string

const (
	Enum_Up_down_stop_continue_Up       Enum_Up_down_stop_continue = "up"
	Enum_Up_down_stop_continue_Down     Enum_Up_down_stop_continue = "down"
	Enum_Up_down_stop_continue_Stop     Enum_Up_down_stop_continue = "stop"
	Enum_Up_down_stop_continue_Continue Enum_Up_down_stop_continue = "continue"
)

// From xsd simple type with enumerate restriction "upright-inverted"
type Enum_Upright_inverted string

const (
	Enum_Upright_inverted_Upright  Enum_Upright_inverted = "upright"
	Enum_Upright_inverted_Inverted Enum_Upright_inverted = "inverted"
)

// From xsd simple type with enumerate restriction "valign"
type Enum_Valign string

const (
	Enum_Valign_Top      Enum_Valign = "top"
	Enum_Valign_Middle   Enum_Valign = "middle"
	Enum_Valign_Bottom   Enum_Valign = "bottom"
	Enum_Valign_Baseline Enum_Valign = "baseline"
)

// From xsd simple type with enumerate restriction "valign-image"
type Enum_Valign_image string

const (
	Enum_Valign_image_Top    Enum_Valign_image = "top"
	Enum_Valign_image_Middle Enum_Valign_image = "middle"
	Enum_Valign_image_Bottom Enum_Valign_image = "bottom"
)

// From xsd simple type with enumerate restriction "wedge-type"
type Enum_Wedge_type string

const (
	Enum_Wedge_type_Crescendo  Enum_Wedge_type = "crescendo"
	Enum_Wedge_type_Diminuendo Enum_Wedge_type = "diminuendo"
	Enum_Wedge_type_Stop       Enum_Wedge_type = "stop"
	Enum_Wedge_type_Continue   Enum_Wedge_type = "continue"
)

// From xsd simple type with enumerate restriction "winged"
type Enum_Winged string

const (
	Enum_Winged_None            Enum_Winged = "none"
	Enum_Winged_Straight        Enum_Winged = "straight"
	Enum_Winged_Curved          Enum_Winged = "curved"
	Enum_Winged_Double_straight Enum_Winged = "double-straight"
	Enum_Winged_Double_curved   Enum_Winged = "double-curved"
)

// From xsd simple type with enumerate restriction "wood-value"
type Enum_Wood_value string

const (
	Enum_Wood_value_Bamboo_scraper        Enum_Wood_value = "bamboo scraper"
	Enum_Wood_value_Board_clapper         Enum_Wood_value = "board clapper"
	Enum_Wood_value_Cabasa                Enum_Wood_value = "cabasa"
	Enum_Wood_value_Castanets             Enum_Wood_value = "castanets"
	Enum_Wood_value_Castanets_with_handle Enum_Wood_value = "castanets with handle"
	Enum_Wood_value_Claves                Enum_Wood_value = "claves"
	Enum_Wood_value_Football_rattle       Enum_Wood_value = "football rattle"
	Enum_Wood_value_Guiro                 Enum_Wood_value = "guiro"
	Enum_Wood_value_Log_drum              Enum_Wood_value = "log drum"
	Enum_Wood_value_Maraca                Enum_Wood_value = "maraca"
	Enum_Wood_value_Maracas               Enum_Wood_value = "maracas"
	Enum_Wood_value_Quijada               Enum_Wood_value = "quijada"
	Enum_Wood_value_Rainstick             Enum_Wood_value = "rainstick"
	Enum_Wood_value_Ratchet               Enum_Wood_value = "ratchet"
	Enum_Wood_value_Reco_reco             Enum_Wood_value = "reco-reco"
	Enum_Wood_value_Sandpaper_blocks      Enum_Wood_value = "sandpaper blocks"
	Enum_Wood_value_Slit_drum             Enum_Wood_value = "slit drum"
	Enum_Wood_value_Temple_block          Enum_Wood_value = "temple block"
	Enum_Wood_value_Vibraslap             Enum_Wood_value = "vibraslap"
	Enum_Wood_value_Whip                  Enum_Wood_value = "whip"
	Enum_Wood_value_Wood_block            Enum_Wood_value = "wood block"
)

// From xsd simple type with enumerate restriction "yes-no"
type Enum_Yes_no string

const (
	Enum_Yes_no_Yes Enum_Yes_no = "yes"
	Enum_Yes_no_No  Enum_Yes_no = "no"
)

// insertion point for gongstructs declarations

// AttributeGroup_bend_sound UnNamed source named attribute group "bend-sound"
type AttributeGroup_bend_sound struct {

	// insertion point for fields

	// generated from attribute "accelerate
	Accelerate Enum_Yes_no `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "first-beat
	First_beat string `xml:"first-beat,attr,omitempty"`

	// generated from attribute "last-beat
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_bezier UnNamed source named attribute group "bezier"
type AttributeGroup_bezier struct {

	// insertion point for fields

	// generated from attribute "bezier-x
	Bezier_x string `xml:"bezier-x,attr,omitempty"`

	// generated from attribute "bezier-y
	Bezier_y string `xml:"bezier-y,attr,omitempty"`

	// generated from attribute "bezier-x2
	Bezier_x2 string `xml:"bezier-x2,attr,omitempty"`

	// generated from attribute "bezier-y2
	Bezier_y2 string `xml:"bezier-y2,attr,omitempty"`

	// generated from attribute "bezier-offset
	Bezier_offset string `xml:"bezier-offset,attr,omitempty"`

	// generated from attribute "bezier-offset2
	Bezier_offset2 string `xml:"bezier-offset2,attr,omitempty"`
}

// AttributeGroup_color UnNamed source named attribute group "color"
type AttributeGroup_color struct {

	// insertion point for fields

	// generated from attribute "color
	Color string `xml:"color,attr,omitempty"`
}

// AttributeGroup_dashed_formatting UnNamed source named attribute group "dashed-formatting"
type AttributeGroup_dashed_formatting struct {

	// insertion point for fields

	// generated from attribute "dash-length
	Dash_length string `xml:"dash-length,attr,omitempty"`

	// generated from attribute "space-length
	Space_length string `xml:"space-length,attr,omitempty"`
}

// AttributeGroup_directive UnNamed source named attribute group "directive"
type AttributeGroup_directive struct {

	// insertion point for fields

	// generated from attribute "directive
	Directive Enum_Yes_no `xml:"directive,attr,omitempty"`
}

// AttributeGroup_document_attributes UnNamed source named attribute group "document-attributes"
type AttributeGroup_document_attributes struct {

	// insertion point for fields

	// generated from attribute "version
	Version string `xml:"version,attr,omitempty"`
}

// AttributeGroup_enclosure UnNamed source named attribute group "enclosure"
type AttributeGroup_enclosure struct {

	// insertion point for fields

	// generated from attribute "enclosure
	Enclosure string `xml:"enclosure,attr,omitempty"`
}

// AttributeGroup_font UnNamed source named attribute group "font"
type AttributeGroup_font struct {

	// insertion point for fields

	// generated from attribute "font-family
	Font_family string `xml:"font-family,attr,omitempty"`

	// generated from attribute "font-style
	Font_style string `xml:"font-style,attr,omitempty"`

	// generated from attribute "font-size
	Font_size string `xml:"font-size,attr,omitempty"`

	// generated from attribute "font-weight
	Font_weight string `xml:"font-weight,attr,omitempty"`
}

// AttributeGroup_halign UnNamed source named attribute group "halign"
type AttributeGroup_halign struct {

	// insertion point for fields

	// generated from attribute "halign
	Halign string `xml:"halign,attr,omitempty"`
}

// AttributeGroup_justify UnNamed source named attribute group "justify"
type AttributeGroup_justify struct {

	// insertion point for fields

	// generated from attribute "justify
	Justify Enum_Left_center_right `xml:"justify,attr,omitempty"`
}

// AttributeGroup_letter_spacing UnNamed source named attribute group "letter-spacing"
type AttributeGroup_letter_spacing struct {

	// insertion point for fields

	// generated from attribute "letter-spacing
	Letter_spacing string `xml:"letter-spacing,attr,omitempty"`
}

// AttributeGroup_level_display UnNamed source named attribute group "level-display"
type AttributeGroup_level_display struct {

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses Enum_Yes_no `xml:"parentheses,attr,omitempty"`

	// generated from attribute "bracket
	Bracket Enum_Yes_no `xml:"bracket,attr,omitempty"`

	// generated from attribute "size
	Size Enum_Symbol_size `xml:"size,attr,omitempty"`
}

// AttributeGroup_line_height UnNamed source named attribute group "line-height"
type AttributeGroup_line_height struct {

	// insertion point for fields

	// generated from attribute "line-height
	Line_height string `xml:"line-height,attr,omitempty"`
}

// AttributeGroup_line_length UnNamed source named attribute group "line-length"
type AttributeGroup_line_length struct {

	// insertion point for fields

	// generated from attribute "line-length
	Line_length string `xml:"line-length,attr,omitempty"`
}

// AttributeGroup_line_shape UnNamed source named attribute group "line-shape"
type AttributeGroup_line_shape struct {

	// insertion point for fields

	// generated from attribute "line-shape
	Line_shape string `xml:"line-shape,attr,omitempty"`
}

// AttributeGroup_line_type UnNamed source named attribute group "line-type"
type AttributeGroup_line_type struct {

	// insertion point for fields

	// generated from attribute "line-type
	Line_type string `xml:"line-type,attr,omitempty"`
}

// AttributeGroup_optional_unique_id UnNamed source named attribute group "optional-unique-id"
type AttributeGroup_optional_unique_id struct {

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// AttributeGroup_orientation UnNamed source named attribute group "orientation"
type AttributeGroup_orientation struct {

	// insertion point for fields

	// generated from attribute "orientation
	Orientation string `xml:"orientation,attr,omitempty"`
}

// AttributeGroup_placement UnNamed source named attribute group "placement"
type AttributeGroup_placement struct {

	// insertion point for fields

	// generated from attribute "placement
	Placement string `xml:"placement,attr,omitempty"`
}

// AttributeGroup_position UnNamed source named attribute group "position"
type AttributeGroup_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_print_object UnNamed source named attribute group "print-object"
type AttributeGroup_print_object struct {

	// insertion point for fields

	// generated from attribute "print-object
	Print_object Enum_Yes_no `xml:"print-object,attr,omitempty"`
}

// AttributeGroup_print_spacing UnNamed source named attribute group "print-spacing"
type AttributeGroup_print_spacing struct {

	// insertion point for fields

	// generated from attribute "print-spacing
	Print_spacing Enum_Yes_no `xml:"print-spacing,attr,omitempty"`
}

// AttributeGroup_print_style UnNamed source named attribute group "print-style"
type AttributeGroup_print_style struct {

	// insertion point for fields

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color
}

// AttributeGroup_print_style_align UnNamed source named attribute group "print-style-align"
type AttributeGroup_print_style_align struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign
	AttributeGroup_valign
}

// AttributeGroup_printout UnNamed source named attribute group "printout"
type AttributeGroup_printout struct {

	// insertion point for fields

	// generated from attribute "print-dot
	Print_dot Enum_Yes_no `xml:"print-dot,attr,omitempty"`

	// generated from attribute "print-lyric
	Print_lyric Enum_Yes_no `xml:"print-lyric,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing
}

// AttributeGroup_smufl UnNamed source named attribute group "smufl"
type AttributeGroup_smufl struct {

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`
}

// AttributeGroup_system_relation UnNamed source named attribute group "system-relation"
type AttributeGroup_system_relation struct {

	// insertion point for fields

	// generated from attribute "system
	System Enum_System_relation_number `xml:"system,attr,omitempty"`
}

// AttributeGroup_symbol_formatting UnNamed source named attribute group "symbol-formatting"
type AttributeGroup_symbol_formatting struct {

	// insertion point for fields

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "text-rotation
	AttributeGroup_text_rotation

	// generated from attribute group "letter-spacing
	AttributeGroup_letter_spacing

	// generated from attribute group "line-height
	AttributeGroup_line_height

	// generated from attribute group "text-direction
	AttributeGroup_text_direction

	// generated from attribute group "enclosure
	AttributeGroup_enclosure
}

// AttributeGroup_text_decoration UnNamed source named attribute group "text-decoration"
type AttributeGroup_text_decoration struct {

	// insertion point for fields

	// generated from attribute "underline
	Underline int `xml:"underline,attr,omitempty"`

	// generated from attribute "overline
	Overline int `xml:"overline,attr,omitempty"`

	// generated from attribute "line-through
	Line_through int `xml:"line-through,attr,omitempty"`
}

// AttributeGroup_text_direction UnNamed source named attribute group "text-direction"
type AttributeGroup_text_direction struct {

	// insertion point for fields

	// generated from attribute "dir
	Dir string `xml:"dir,attr,omitempty"`
}

// AttributeGroup_text_formatting UnNamed source named attribute group "text-formatting"
type AttributeGroup_text_formatting struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace space
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "text-rotation
	AttributeGroup_text_rotation

	// generated from attribute group "letter-spacing
	AttributeGroup_letter_spacing

	// generated from attribute group "line-height
	AttributeGroup_line_height

	// generated from attribute group "text-direction
	AttributeGroup_text_direction

	// generated from attribute group "enclosure
	AttributeGroup_enclosure
}

// AttributeGroup_text_rotation UnNamed source named attribute group "text-rotation"
type AttributeGroup_text_rotation struct {

	// insertion point for fields

	// generated from attribute "rotation
	Rotation string `xml:"rotation,attr,omitempty"`
}

// AttributeGroup_trill_sound UnNamed source named attribute group "trill-sound"
type AttributeGroup_trill_sound struct {

	// insertion point for fields

	// generated from attribute "start-note
	Start_note string `xml:"start-note,attr,omitempty"`

	// generated from attribute "trill-step
	Trill_step string `xml:"trill-step,attr,omitempty"`

	// generated from attribute "two-note-turn
	Two_note_turn string `xml:"two-note-turn,attr,omitempty"`

	// generated from attribute "accelerate
	Accelerate Enum_Yes_no `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "second-beat
	Second_beat string `xml:"second-beat,attr,omitempty"`

	// generated from attribute "last-beat
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_valign UnNamed source named attribute group "valign"
type AttributeGroup_valign struct {

	// insertion point for fields

	// generated from attribute "valign
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_valign_image UnNamed source named attribute group "valign-image"
type AttributeGroup_valign_image struct {

	// insertion point for fields

	// generated from attribute "valign
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_x_position UnNamed source named attribute group "x-position"
type AttributeGroup_x_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_y_position UnNamed source named attribute group "y-position"
type AttributeGroup_y_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_image_attributes UnNamed source named attribute group "image-attributes"
type AttributeGroup_image_attributes struct {

	// insertion point for fields

	// generated from attribute "source
	Source string `xml:"source,attr,omitempty"`

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "height
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign-image
	AttributeGroup_valign_image
}

// AttributeGroup_print_attributes UnNamed source named attribute group "print-attributes"
type AttributeGroup_print_attributes struct {

	// insertion point for fields

	// generated from attribute "staff-spacing
	Staff_spacing string `xml:"staff-spacing,attr,omitempty"`

	// generated from attribute "new-system
	New_system Enum_Yes_no `xml:"new-system,attr,omitempty"`

	// generated from attribute "new-page
	New_page Enum_Yes_no `xml:"new-page,attr,omitempty"`

	// generated from attribute "blank-page
	Blank_page int `xml:"blank-page,attr,omitempty"`

	// generated from attribute "page-number
	Page_number string `xml:"page-number,attr,omitempty"`
}

// AttributeGroup_element_position UnNamed source named attribute group "element-position"
type AttributeGroup_element_position struct {

	// insertion point for fields

	// generated from attribute "element
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "position
	Position int `xml:"position,attr,omitempty"`
}

// AttributeGroup_link_attributes UnNamed source named attribute group "link-attributes"
type AttributeGroup_link_attributes struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink href
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink type
	Type string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink role
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink title
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink show
	Show string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink actuate
	Actuate string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

// AttributeGroup_group_name_text UnNamed source named attribute group "group-name-text"
type AttributeGroup_group_name_text struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "justify
	AttributeGroup_justify
}

// AttributeGroup_measure_attributes UnNamed source named attribute group "measure-attributes"
type AttributeGroup_measure_attributes struct {

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "implicit
	Implicit Enum_Yes_no `xml:"implicit,attr,omitempty"`

	// generated from attribute "non-controlling
	Non_controlling Enum_Yes_no `xml:"non-controlling,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// AttributeGroup_part_attributes UnNamed source named attribute group "part-attributes"
type AttributeGroup_part_attributes struct {

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// AttributeGroup_part_name_text UnNamed source named attribute group "part-name-text"
type AttributeGroup_part_name_text struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify
}

// Accidental_text Named source named complex type "accidental-text"
// The accidental-text type represents an element with an accidental
// value and text-formatting attributes.
type Accidental_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "text-formatting
	AttributeGroup_text_formatting

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Accidental_value `xml:",chardata"`
}

// Coda Named source named complex type "coda"
// The coda type is the visual indicator of a coda sign. The exact glyph
// can be specified with the smufl attribute. A sound element is also needed to guide
// playback applications reliably.
type Coda struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Dynamics Named source named complex type "dynamics"
// Dynamics can be associated either with a note or a general musical
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

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "enclosure
	AttributeGroup_enclosure

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "p" of type empty order 230 depth 1
	P string `xml:"p,omitempty"`

	// generated from element "pp" of type empty order 231 depth 1
	Pp string `xml:"pp,omitempty"`

	// generated from element "ppp" of type empty order 232 depth 1
	Ppp string `xml:"ppp,omitempty"`

	// generated from element "pppp" of type empty order 233 depth 1
	Pppp string `xml:"pppp,omitempty"`

	// generated from element "ppppp" of type empty order 234 depth 1
	Ppppp string `xml:"ppppp,omitempty"`

	// generated from element "pppppp" of type empty order 235 depth 1
	Pppppp string `xml:"pppppp,omitempty"`

	// generated from element "f" of type empty order 236 depth 1
	F string `xml:"f,omitempty"`

	// generated from element "ff" of type empty order 237 depth 1
	Ff string `xml:"ff,omitempty"`

	// generated from element "fff" of type empty order 238 depth 1
	Fff string `xml:"fff,omitempty"`

	// generated from element "ffff" of type empty order 239 depth 1
	Ffff string `xml:"ffff,omitempty"`

	// generated from element "fffff" of type empty order 240 depth 1
	Fffff string `xml:"fffff,omitempty"`

	// generated from element "ffffff" of type empty order 241 depth 1
	Ffffff string `xml:"ffffff,omitempty"`

	// generated from element "mp" of type empty order 242 depth 1
	Mp string `xml:"mp,omitempty"`

	// generated from element "mf" of type empty order 243 depth 1
	Mf string `xml:"mf,omitempty"`

	// generated from element "sf" of type empty order 244 depth 1
	Sf string `xml:"sf,omitempty"`

	// generated from element "sfp" of type empty order 245 depth 1
	Sfp string `xml:"sfp,omitempty"`

	// generated from element "sfpp" of type empty order 246 depth 1
	Sfpp string `xml:"sfpp,omitempty"`

	// generated from element "fp" of type empty order 247 depth 1
	Fp string `xml:"fp,omitempty"`

	// generated from element "rf" of type empty order 248 depth 1
	Rf string `xml:"rf,omitempty"`

	// generated from element "rfz" of type empty order 249 depth 1
	Rfz string `xml:"rfz,omitempty"`

	// generated from element "sfz" of type empty order 250 depth 1
	Sfz string `xml:"sfz,omitempty"`

	// generated from element "sffz" of type empty order 251 depth 1
	Sffz string `xml:"sffz,omitempty"`

	// generated from element "fz" of type empty order 252 depth 1
	Fz string `xml:"fz,omitempty"`

	// generated from element "n" of type empty order 253 depth 1
	N string `xml:"n,omitempty"`

	// generated from element "pf" of type empty order 254 depth 1
	Pf string `xml:"pf,omitempty"`

	// generated from element "sfzp" of type empty order 255 depth 1
	Sfzp string `xml:"sfzp,omitempty"`

	// generated from element "other-dynamics" of type other-text order 256 depth 1
	Other_dynamics []*Other_text `xml:"other-dynamics,omitempty"`
}

// Empty Named source named complex type "empty"
// The empty type represents an empty element with no attributes.
type Empty struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Empty_placement Named source named complex type "empty-placement"
// The empty-placement type represents an empty element with print-style
// and placement attributes.
type Empty_placement struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement
}

// Empty_placement_smufl Named source named complex type "empty-placement-smufl"
// The empty-placement-smufl type represents an empty element with
// print-style, placement, and smufl attributes.
type Empty_placement_smufl struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl
}

// Empty_print_style Named source named complex type "empty-print-style"
// The empty-print-style type represents an empty element with
// print-style attribute group.
type Empty_print_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style
}

// Empty_print_style_align Named source named complex type "empty-print-style-align"
// The empty-print-style-align type represents an empty element with
// print-style-align attribute group.
type Empty_print_style_align struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align
}

// Empty_print_style_align_id Named source named complex type "empty-print-style-align-id"
// The empty-print-style-align-id type represents an empty element with
// print-style-align and optional-unique-id attribute groups.
type Empty_print_style_align_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Empty_print_object_style_align Named source named complex type "empty-print-object-style-align"
// The empty-print-style-align-object type represents an empty element
// with print-object and print-style-align attribute groups.
type Empty_print_object_style_align struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align
}

// Empty_trill_sound Named source named complex type "empty-trill-sound"
// The empty-trill-sound type represents an empty element with
// print-style, placement, and trill-sound attributes.
type Empty_trill_sound struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Horizontal_turn Named source named complex type "horizontal-turn"
// The horizontal-turn type represents turn elements that are horizontal
// rather than vertical. These are empty elements with print-style, placement,
// trill-sound, and slash attributes. If the slash attribute is yes, then a vertical
// line is used to slash the turn. It is no if not specified.
type Horizontal_turn struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "slash
	Slash Enum_Yes_no `xml:"slash,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Fermata Named source named complex type "fermata"
// The fermata text content represents the shape of the fermata sign. An
// empty fermata element represents a normal fermata. The fermata type is upright if
// not specified.
type Fermata struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Fingering Named source named complex type "fingering"
// Fingering is typically indicated 1,2,3,4,5. Multiple fingerings may be
// given, typically to substitute fingerings in the middle of a note. The substitution
// and alternate values are "no" if the attribute is not present. For guitar and other
// fretted instruments, the fingering element represents the fretting finger; the pluck
// element represents the plucking finger.
type Fingering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "substitution
	Substitution Enum_Yes_no `xml:"substitution,attr,omitempty"`

	// generated from attribute "alternate
	Alternate Enum_Yes_no `xml:"alternate,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_symbol Named source named complex type "formatted-symbol"
// The formatted-symbol type represents a SMuFL musical symbol element
// with formatting attributes.
type Formatted_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "symbol-formatting
	AttributeGroup_symbol_formatting

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_symbol_id Named source named complex type "formatted-symbol-id"
// The formatted-symbol-id type represents a SMuFL musical symbol element
// with formatting and id attributes.
type Formatted_symbol_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "symbol-formatting
	AttributeGroup_symbol_formatting

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_text Named source named complex type "formatted-text"
// The formatted-text type represents a text element with text-formatting
// attributes.
type Formatted_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "text-formatting
	AttributeGroup_text_formatting

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_text_id Named source named complex type "formatted-text-id"
// The formatted-text-id type represents a text element with
// text-formatting and id attributes.
type Formatted_text_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "text-formatting
	AttributeGroup_text_formatting

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Fret Named source named complex type "fret"
// The fret element is used with tablature notation and chord diagrams.
// Fret numbers start with 0 for an open string and 1 for the first fret.
type Fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Level Named source named complex type "level"
// The level type is used to specify editorial information for different
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

	// insertion point for fields

	// generated from attribute "reference
	Reference Enum_Yes_no `xml:"reference,attr,omitempty"`

	// generated from attribute "type
	Type Enum_Start_stop_single `xml:"type,attr,omitempty"`

	// generated from attribute group "level-display
	AttributeGroup_level_display

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Midi_device Named source named complex type "midi-device"
// The midi-device type corresponds to the DeviceName meta event in
// Standard MIDI Files. The optional port attribute is a number from 1 to 16 that can
// be used with the unofficial MIDI 1.0 port (or cable) meta event. Unlike the
// DeviceName meta event, there can be multiple midi-device elements per MusicXML part.
// The optional id attribute refers to the score-instrument assigned to this device. If
// missing, the device assignment affects all score-instrument elements in the
// score-part.
type Midi_device struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "port
	Port int `xml:"port,attr,omitempty"`

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Midi_instrument Named source named complex type "midi-instrument"
// The midi-instrument type defines MIDI 1.0 instrument playback. The
// midi-instrument element can be a part of either the score-instrument element at the
// start of a part, or the sound element within a part. The id attribute refers to the
// score-instrument affected by the change.
type Midi_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "midi-channel" of type midi-16 order 311 depth 1
	Midi_channel int `xml:"midi-channel,omitempty"`

	// generated from element "midi-name" of type string order 312 depth 1
	Midi_name string `xml:"midi-name,omitempty"`

	// generated from element "midi-bank" of type midi-16384 order 313 depth 1
	Midi_bank int `xml:"midi-bank,omitempty"`

	// generated from element "midi-program" of type midi-128 order 314 depth 1
	Midi_program int `xml:"midi-program,omitempty"`

	// generated from element "midi-unpitched" of type midi-128 order 315 depth 1
	Midi_unpitched int `xml:"midi-unpitched,omitempty"`

	// generated from element "volume" of type percent order 316 depth 1
	Volume string `xml:"volume,omitempty"`

	// generated from element "pan" of type rotation-degrees order 317 depth 1
	Pan string `xml:"pan,omitempty"`

	// generated from element "elevation" of type rotation-degrees order 318 depth 1
	Elevation string `xml:"elevation,omitempty"`
}

// Name_display Named source named complex type "name-display"
// The name-display type is used for exact formatting of multi-font text
// in part and group names to the left of the system. The print-object attribute can be
// used to determine what, if anything, is printed at the start of each system.
// Enclosure for the display-text element is none by default. Language for the
// display-text element is Italian ("it") by default.
type Name_display struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "display-text" of type formatted-text order 320 depth 1
	Display_text []*Formatted_text `xml:"display-text,omitempty"`

	// generated from element "accidental-text" of type accidental-text order 321 depth 1
	Accidental_text []*Accidental_text `xml:"accidental-text,omitempty"`
}

// Other_play Named source named complex type "other-play"
// The other-play element represents other types of playback. The
// required type attribute indicates the type of playback to which the element content
// applies.
type Other_play struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Play Named source named complex type "play"
// The play type specifies playback techniques to be used in conjunction
// with the instrument-sound element. When used as part of a sound element, it applies
// to all notes going forward in score order. In multi-instrument parts, the affected
// instrument should be specified using the id attribute. When used as part of a note
// element, it applies to the current note only.
type Play struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "ipa" of type string order 325 depth 1
	Ipa string `xml:"ipa,omitempty"`

	// generated from element "mute" of type mute order 326 depth 1
	Mute string `xml:"mute,omitempty"`

	// generated from element "semi-pitched" of type semi-pitched order 327 depth 1
	Semi_pitched string `xml:"semi-pitched,omitempty"`

	// generated from element "other-play" of type other-play order 328 depth 1
	Other_play []*Other_play `xml:"other-play,omitempty"`
}

// Segno Named source named complex type "segno"
// The segno type is the visual indicator of a segno sign. The exact
// glyph can be specified with the smufl attribute. A sound element is also needed to
// guide playback applications reliably.
type Segno struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// String_type Named source named complex type "string-type"
// The string type is used with tablature notation, regular notation
// (where it is often circled), and chord diagrams. String numbers start with 1 for the
// highest pitched full-length string.
type String_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Typed_text Named source named complex type "typed-text"
// The typed-text type represents a text element with a type attribute.
type Typed_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Wavy_line Named source named complex type "wavy-line"
// Wavy lines are one way to indicate trills and vibrato. When used with
// a barline element, they should always have type="continue" set. The smufl attribute
// specifies a particular wavy line glyph from the SMuFL Multi-segment lines range.
type Wavy_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop_continue `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Attributes Named source named complex type "attributes"
// The attributes element contains musical information that typically
// changes on measure boundaries. This includes key and time signatures, clefs,
// transpositions, and staving. When attributes are changed mid-measure, it affects the
// music in score order, not in MusicXML document order.
type Attributes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 342 depth 1
	Group_editorial

	// generated from element "divisions" of type positive-divisions order 343 depth 1
	Divisions string `xml:"divisions,omitempty"`

	// generated from element "key" of type key order 344 depth 1
	Key []*Key `xml:"key,omitempty"`

	// generated from element "time" of type time order 345 depth 1
	Time []*Time `xml:"time,omitempty"`

	// generated from element "staves" of type nonNegativeInteger order 346 depth 1
	Staves int `xml:"staves,omitempty"`

	// generated from element "part-symbol" of type part-symbol order 347 depth 1
	Part_symbol *Part_symbol `xml:"part-symbol,omitempty"`

	// generated from element "instruments" of type nonNegativeInteger order 348 depth 1
	Instruments int `xml:"instruments,omitempty"`

	// generated from element "clef" of type clef order 349 depth 1
	Clef []*Clef `xml:"clef,omitempty"`

	// generated from element "staff-details" of type staff-details order 350 depth 1
	Staff_details []*Staff_details `xml:"staff-details,omitempty"`

	// generated from element "transpose" of type transpose order 351 depth 1
	Transpose []*Transpose `xml:"transpose,omitempty"`

	// generated from element "for-part" of type for-part order 352 depth 1
	For_part []*For_part `xml:"for-part,omitempty"`

	// generated from anonymous type within outer element "directive" of type A_directive.
	Directive []*A_directive `xml:"directive,omitempty"`

	// generated from element "measure-style" of type measure-style order 356 depth 1
	Measure_style []*Measure_style `xml:"measure-style,omitempty"`
}

// A_directive Named source within outer element "directive"
type A_directive struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beat_repeat Named source named complex type "beat-repeat"
// The beat-repeat type is used to indicate that a single beat (but
// possibly many notes) is repeated. The slashes attribute specifies the number of
// slashes to use in the symbol. The use-dots attribute indicates whether or not to use
// dots as well (for instance, with mixed rhythm patterns). The value for slashes is 1
// and the value for use-dots is no if not specified. The stop type indicates the first
// beat where the repeats are no longer displayed. Both the start and stop of the beat
// being repeated should be specified unless the repeats are displayed through the end
// of the part. The beat-repeat element specifies a notation style for repetitions. The
// actual music being repeated needs to be repeated within the MusicXML file. This
// element specifies the notation that indicates the repeat.
type Beat_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "slashes
	Slashes int `xml:"slashes,attr,omitempty"`

	// generated from attribute "use-dots
	Use_dots Enum_Yes_no `xml:"use-dots,attr,omitempty"`

	// generated from group with order 358 depth 1
	Group_slash
}

// Cancel Named source named complex type "cancel"
// A cancel element indicates that the old key signature should be
// cancelled before the new one appears. This will always happen when changing to C
// major or A minor and need not be specified then. The cancel value matches the fifths
// value of the cancelled key signature (e.g., a cancel of -2 will provide an explicit
// cancellation for changing from B flat major to F major). The optional location
// attribute indicates where the cancellation appears relative to the new key
// signature.
type Cancel struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Clef Named source named complex type "clef"
// Clefs are represented by a combination of sign, line, and
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

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "additional
	Additional Enum_Yes_no `xml:"additional,attr,omitempty"`

	// generated from attribute "size
	Size string `xml:"size,attr,omitempty"`

	// generated from attribute "after-barline
	After_barline Enum_Yes_no `xml:"after-barline,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 361 depth 1
	Group_clef
}

// Double Named source named complex type "double"
// The double type indicates that the music is doubled one octave from
// what is currently written. If the above attribute is set to yes, the doubling is one
// octave above what is written, as for mixed flute / piccolo parts in band literature.
// Otherwise the doubling is one octave below what is written, as for mixed cello /
// bass parts in orchestral literature.
type Double struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "above
	Above Enum_Yes_no `xml:"above,attr,omitempty"`
}

// For_part Named source named complex type "for-part"
// The for-part type is used in a concert score to indicate the
// transposition for a transposed part created from that score. It is only used in
// score files that contain a concert-score element in the defaults. This allows
// concert scores with transposed parts to be represented in a single uncompressed
// MusicXML file. The optional number attribute refers to staff numbers, from top to
// bottom on the system. If absent, the child elements apply to all staves in the
// created part.
type For_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "part-clef" of type part-clef order 367 depth 1
	Part_clef *Part_clef `xml:"part-clef,omitempty"`

	// generated from element "part-transpose" of type part-transpose order 368 depth 1
	Part_transpose *Part_transpose `xml:"part-transpose,omitempty"`
}

// Interchangeable Named source named complex type "interchangeable"
// The interchangeable type is used to represent the second in a pair of
// interchangeable dual time signatures, such as the 6/8 in 3/4 (6/8). A separate
// symbol attribute value is available compared to the time element's symbol attribute,
// which applies to the first of the dual time signatures.
type Interchangeable struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator
	Separator string `xml:"separator,attr,omitempty"`

	// generated from element "time-relation" of type time-relation order 371 depth 1
	Time_relation string `xml:"time-relation,omitempty"`

	// generated from group with order 372 depth 1
	Group_time_signature
}

// Key Named source named complex type "key"
// The key type represents a key signature. Both traditional and
// non-traditional key signatures are supported. The optional number attribute refers
// to staff numbers. If absent, the key signature applies to all staves in the part.
// Key signatures appear at the start of each system unless the print-object attribute
// has been set to "no".
type Key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 374 depth 1
	Group_traditional_key

	// generated from group with order 375 depth 1
	Group_non_traditional_key

	// generated from element "key-octave" of type key-octave order 376 depth 1
	Key_octave []*Key_octave `xml:"key-octave,omitempty"`
}

// Key_accidental Named source named complex type "key-accidental"
// The key-accidental type indicates the accidental to be displayed in a
// non-traditional key signature, represented in the same manner as the accidental type
// without the formatting attributes.
type Key_accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Accidental_value `xml:",chardata"`
}

// Key_octave Named source named complex type "key-octave"
// The key-octave type specifies in which octave an element of a key
// signature appears. The content specifies the octave value using the same values as
// the display-octave element. The number attribute is a positive integer that refers
// to the key signature element in left-to-right order. If the cancel attribute is set
// to yes, then this number refers to the canceling key signature specified by the
// cancel element in the parent key element. The cancel attribute cannot be set to yes
// if there is no corresponding cancel element within the parent key element. It is no
// by default.
type Key_octave struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "cancel
	Cancel Enum_Yes_no `xml:"cancel,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Line_detail Named source named complex type "line-detail"
// If the staff-lines element is present, the appearance of each line may
// be individually specified with a line-detail type. Staff lines are numbered from
// bottom to top. The print-object attribute allows lines to be hidden within a staff.
// This is used in special situations such as a widely-spaced percussion staff where a
// note placed below the higher line is distinct from a note placed above the lower
// line. Hidden staff lines are included when specifying clef lines and determining
// display-step / display-octave values, but are not counted as lines for the purposes
// of the system-layout and staff-layout elements.
type Line_detail struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line
	Line int `xml:"line,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "print-object
	AttributeGroup_print_object
}

// Measure_repeat Named source named complex type "measure-repeat"
// The measure-repeat type is used for both single and multiple measure
// repeats. The text of the element indicates the number of measures to be repeated in
// a single pattern. The slashes attribute specifies the number of slashes to use in
// the repeat sign. It is 1 if not specified. The text of the element is ignored when
// the type is stop. The stop type indicates the first measure where the repeats are no
// longer displayed. Both the start and the stop of the measure-repeat should be
// specified unless the repeats are displayed through the end of the part. The
// measure-repeat element specifies a notation style for repetitions. The actual music
// being repeated needs to be repeated within each measure of the MusicXML file. This
// element specifies the notation that indicates the repeat.
type Measure_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "slashes
	Slashes int `xml:"slashes,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_style Named source named complex type "measure-style"
// A measure-style indicates a special way to print partial to multiple
// measures within a part. This includes multiple rests over several measures, repeats
// of beats, single, or multiple measures, and use of slash notation. The multiple-rest
// and measure-repeat elements indicate the number of measures covered in the element
// content. The beat-repeat and slash elements can cover partial measures. All but the
// multiple-rest element use a type attribute to indicate starting and stopping the use
// of the style. The optional number attribute specifies the staff number from top to
// bottom on the system, as with clef.
type Measure_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "multiple-rest" of type multiple-rest order 388 depth 1
	Multiple_rest *Multiple_rest `xml:"multiple-rest,omitempty"`

	// generated from element "measure-repeat" of type measure-repeat order 389 depth 1
	Measure_repeat *Measure_repeat `xml:"measure-repeat,omitempty"`

	// generated from element "beat-repeat" of type beat-repeat order 390 depth 1
	Beat_repeat *Beat_repeat `xml:"beat-repeat,omitempty"`

	// generated from element "slash" of type slash order 391 depth 1
	Slash *Slash `xml:"slash,omitempty"`
}

// Multiple_rest Named source named complex type "multiple-rest"
// The text of the multiple-rest type indicates the number of measures in
// the multiple rest. Multiple rests may use the 1-bar / 2-bar / 4-bar rest symbols, or
// a single shape. The use-symbols attribute indicates which to use; it is no if not
// specified.
type Multiple_rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols
	Use_symbols Enum_Yes_no `xml:"use-symbols,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Part_clef Named source named complex type "part-clef"
// The child elements of the part-clef type have the same meaning as for
// the clef type. However that meaning applies to a transposed part created from the
// existing score file.
type Part_clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 397 depth 1
	Group_clef
}

// Part_symbol Named source named complex type "part-symbol"
// The part-symbol type indicates how a symbol for a multi-staff part is
// indicated in the score; brace is the default value. The top-staff and bottom-staff
// attributes are used when the brace does not extend across the entire part. For
// example, in a 3-staff organ part, the top-staff will typically be 1 for the right
// hand, while the bottom-staff will typically be 2 for the left hand. Staff 3 for the
// pedals is usually outside the brace. By default, the presence of a part-symbol
// element that does not extend across the entire part also indicates a corresponding
// change in the common barlines within a part.
type Part_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "top-staff
	Top_staff int `xml:"top-staff,attr,omitempty"`

	// generated from attribute "bottom-staff
	Bottom_staff int `xml:"bottom-staff,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Group_symbol_value `xml:",chardata"`
}

// Part_transpose Named source named complex type "part-transpose"
// The child elements of the part-transpose type have the same meaning as
// for the transpose type. However that meaning applies to a transposed part created
// from the existing score file.
type Part_transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 402 depth 1
	Group_transpose
}

// Slash Named source named complex type "slash"
// The slash type is used to indicate that slash notation is to be used.
// If the slash is on every beat, use-stems is no (the default). To indicate rhythms
// but not pitches, use-stems is set to yes. The type attribute indicates whether this
// is the start or stop of a slash notation style. The use-dots attribute works as for
// the beat-repeat element, and only has effect if use-stems is no.
type Slash struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "use-dots
	Use_dots Enum_Yes_no `xml:"use-dots,attr,omitempty"`

	// generated from attribute "use-stems
	Use_stems Enum_Yes_no `xml:"use-stems,attr,omitempty"`

	// generated from group with order 404 depth 1
	Group_slash
}

// Staff_details Named source named complex type "staff-details"
// The staff-details element is used to indicate different types of
// staves. The optional number attribute specifies the staff number from top to bottom
// on the system, as with clef. The print-object attribute is used to indicate when a
// staff is not printed in a part, usually in large scores where empty parts are
// omitted. It is yes by default. If print-spacing is yes while print-object is no, the
// score is printed in cutaway format where vertical space is left for the empty part.
type Staff_details struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "show-frets
	Show_frets string `xml:"show-frets,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing

	// generated from element "staff-type" of type staff-type order 406 depth 1
	Staff_type string `xml:"staff-type,omitempty"`

	// generated from element "staff-lines" of type nonNegativeInteger order 407 depth 1
	Staff_lines int `xml:"staff-lines,omitempty"`

	// generated from element "line-detail" of type line-detail order 408 depth 1
	Line_detail []*Line_detail `xml:"line-detail,omitempty"`

	// generated from element "staff-tuning" of type staff-tuning order 409 depth 1
	Staff_tuning []*Staff_tuning `xml:"staff-tuning,omitempty"`

	// generated from element "capo" of type nonNegativeInteger order 410 depth 1
	Capo int `xml:"capo,omitempty"`

	// generated from element "staff-size" of type staff-size order 411 depth 1
	Staff_size *Staff_size `xml:"staff-size,omitempty"`
}

// Staff_size Named source named complex type "staff-size"
// The staff-size element indicates how large a staff space is on this
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
type Staff_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "scaling
	Scaling string `xml:"scaling,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Staff_tuning Named source named complex type "staff-tuning"
// The staff-tuning type specifies the open, non-capo tuning of the lines
// on a tablature staff.
type Staff_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line
	Line int `xml:"line,attr,omitempty"`

	// generated from group with order 416 depth 1
	Group_tuning
}

// Time Named source named complex type "time"
// Time signatures are represented by the beats element for the numerator
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

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "symbol
	Symbol Enum_Time_symbol `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator
	Separator Enum_Time_separator `xml:"separator,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 418 depth 1
	Group_time_signature

	// generated from element "interchangeable" of type interchangeable order 419 depth 1
	Interchangeable *Interchangeable `xml:"interchangeable,omitempty"`

	// generated from element "senza-misura" of type string order 420 depth 1
	Senza_misura string `xml:"senza-misura,omitempty"`
}

// Transpose Named source named complex type "transpose"
// The transpose type represents what must be added to a written pitch to
// get a correct sounding pitch. The optional number attribute refers to staff numbers,
// from top to bottom on the system. If absent, the transposition applies to all staves
// in the part. Per-staff transposition is most often used in parts that represent
// multiple instruments.
type Transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 425 depth 1
	Group_transpose
}

// Bar_style_color Named source named complex type "bar-style-color"
// The bar-style-color type contains barline style and color information.
type Bar_style_color struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Barline Named source named complex type "barline"
// If a barline is other than a normal single barline, it should be
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

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// generated from attribute "segno
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "coda
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "bar-style" of type bar-style-color order 430 depth 1
	Bar_style *Bar_style_color `xml:"bar-style,omitempty"`

	// generated from group with order 431 depth 1
	Group_editorial

	// generated from element "wavy-line" of type wavy-line order 432 depth 1
	Wavy_line *Wavy_line `xml:"wavy-line,omitempty"`

	// generated from element "segno" of type segno order 433 depth 1
	Segno_1 *Segno `xml:"segno,omitempty"`

	// generated from element "coda" of type coda order 434 depth 1
	Coda_1 *Coda `xml:"coda,omitempty"`

	// generated from element "fermata" of type fermata order 435 depth 1
	Fermata *Fermata `xml:"fermata,omitempty"`

	// generated from element "ending" of type ending order 436 depth 1
	Ending *Ending `xml:"ending,omitempty"`

	// generated from element "repeat" of type repeat order 437 depth 1
	Repeat *Repeat `xml:"repeat,omitempty"`
}

// Ending Named source named complex type "ending"
// The ending type represents multiple (e.g. first and second) endings.
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

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "end-length
	End_length string `xml:"end-length,attr,omitempty"`

	// generated from attribute "text-x
	Text_x string `xml:"text-x,attr,omitempty"`

	// generated from attribute "text-y
	Text_y string `xml:"text-y,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "system-relation
	AttributeGroup_system_relation

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Repeat Named source named complex type "repeat"
// The repeat type represents repeat marks. The start of the repeat has a
// forward direction while the end of the repeat has a backward direction. The times
// and after-jump attributes are only used with backward repeats that are not part of
// an ending. The times attribute indicates the number of times the repeated section is
// played. The after-jump attribute indicates if the repeats are played after a jump
// due to a da capo or dal segno.
type Repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "direction
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "times
	Times int `xml:"times,attr,omitempty"`

	// generated from attribute "after-jump
	After_jump Enum_Yes_no `xml:"after-jump,attr,omitempty"`

	// generated from attribute "winged
	Winged string `xml:"winged,attr,omitempty"`
}

// Accord Named source named complex type "accord"
// The accord type represents the tuning of a single string in the
// scordatura element. It uses the same group of elements as the staff-tuning element.
// Strings are numbered from high to low.
type Accord struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "string
	String int `xml:"string,attr,omitempty"`

	// generated from group with order 445 depth 1
	Group_tuning
}

// Accordion_registration Named source named complex type "accordion-registration"
// The accordion-registration type is used for accordion registration
// symbols. These are circular symbols divided horizontally into high, middle, and low
// sections that correspond to 4', 8', and 16' pipes. Each accordion-high,
// accordion-middle, and accordion-low element represents the presence of one or more
// dots in the registration diagram. An accordion-registration element needs to have at
// least one of the child elements present.
type Accordion_registration struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accordion-high" of type empty order 447 depth 1
	Accordion_high string `xml:"accordion-high,omitempty"`

	// generated from element "accordion-middle" of type accordion-middle order 448 depth 1
	Accordion_middle int `xml:"accordion-middle,omitempty"`

	// generated from element "accordion-low" of type empty order 449 depth 1
	Accordion_low string `xml:"accordion-low,omitempty"`
}

// Barre Named source named complex type "barre"
// The barre element indicates placing a finger over multiple strings on
// a single fret. The type is "start" for the lowest pitched string (e.g., the string
// with the highest MusicXML number) and is "stop" for the highest pitched string.
type Barre struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color
}

// Bass Named source named complex type "bass"
// The bass type is used to indicate a bass note in popular music chord
// symbols, e.g. G/C. It is generally not used in functional harmony, as inversion is
// generally not used in pop chord symbols. As with root, it is divided into step and
// alter elements, similar to pitches. The arrangement attribute specifies where the
// bass is displayed relative to what precedes it.
type Bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "arrangement
	Arrangement string `xml:"arrangement,attr,omitempty"`

	// generated from element "bass-separator" of type style-text order 455 depth 1
	Bass_separator *Style_text `xml:"bass-separator,omitempty"`

	// generated from element "bass-step" of type bass-step order 456 depth 1
	Bass_step *Bass_step `xml:"bass-step,omitempty"`

	// generated from element "bass-alter" of type harmony-alter order 457 depth 1
	Bass_alter *Harmony_alter `xml:"bass-alter,omitempty"`
}

// Harmony_alter Named source named complex type "harmony-alter"
// The harmony-alter type represents the chromatic alteration of the
// root, numeral, or bass of the current harmony-chord group within the harmony
// element. In some chord styles, the text of the preceding element may include
// alteration information. In that case, the print-object attribute of this type can be
// set to no. The location attribute indicates whether the alteration should appear to
// the left or the right of the preceding element. Its default value varies by element.
type Harmony_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location Enum_Left_right `xml:"location,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Bass_step Named source named complex type "bass-step"
// The bass-step type represents the pitch step of the bass of the
// current chord within the harmony element. The text attribute indicates how the bass
// should appear in a score if not using the element contents.
type Bass_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beater Named source named complex type "beater"
// The beater type represents pictograms for beaters, mallets, and sticks
// that do not have different materials represented in the pictogram.
type Beater struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip
	Tip string `xml:"tip,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beat_unit_tied Named source named complex type "beat-unit-tied"
// The beat-unit-tied type indicates a beat-unit within a metronome mark
// that is tied to the preceding beat-unit. This allows two or more tied notes to be
// associated with a per-minute value in a metronome mark, whereas the metronome-tied
// element is restricted to metric relationship marks.
type Beat_unit_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 465 depth 1
	Group_beat_unit
}

// Bracket Named source named complex type "bracket"
// Brackets are combined with words in a variety of modern directions.
// The line-end attribute specifies if there is a jog up or down (or both), an arrow,
// or nothing at the start or end of the bracket. If the line-end is up or down, the
// length of the jog can be specified using the end-length attribute. The line-type is
// solid if not specified.
type Bracket struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line-end
	Line_end string `xml:"line-end,attr,omitempty"`

	// generated from attribute "end-length
	End_length string `xml:"end-length,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Dashes Named source named complex type "dashes"
// The dashes type represents dashes, used for instance with cresc. and
// dim. marks.
type Dashes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop_continue `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Degree Named source named complex type "degree"
// The degree type is used to add, alter, or subtract individual notes in
// the chord. The print-object attribute can be used to keep the degree from printing
// separately when it has already taken into account in the text attribute of the kind
// element. The degree-value and degree-type text attributes specify how the value and
// type of the degree should be displayed. A harmony of kind "other" can be spelled
// explicitly by using a series of degree elements together with a root.
type Degree struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "degree-value" of type degree-value order 478 depth 1
	Degree_value *Degree_value `xml:"degree-value,omitempty"`

	// generated from element "degree-alter" of type degree-alter order 479 depth 1
	Degree_alter *Degree_alter `xml:"degree-alter,omitempty"`

	// generated from element "degree-type" of type degree-type order 480 depth 1
	Degree_type *Degree_type `xml:"degree-type,omitempty"`
}

// Degree_alter Named source named complex type "degree-alter"
// The degree-alter type represents the chromatic alteration for the
// current degree. If the degree-type value is alter or subtract, the degree-alter
// value is relative to the degree already in the chord based on its kind element. If
// the degree-type value is add, the degree-alter is relative to a dominant chord
// (major and perfect intervals except for a minor seventh). The plus-minus attribute
// is used to indicate if plus and minus symbols should be used instead of sharp and
// flat symbols to display the degree alteration. It is no if not specified.
type Degree_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "plus-minus
	Plus_minus Enum_Yes_no `xml:"plus-minus,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_type Named source named complex type "degree-type"
// The degree-type type indicates if this degree is an addition,
// alteration, or subtraction relative to the kind of the current chord. The value of
// the degree-type element affects the interpretation of the value of the degree-alter
// element. The text attribute specifies how the type of the degree should be
// displayed.
type Degree_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_value Named source named complex type "degree-value"
// The content of the degree-value type is a number indicating the degree
// of the chord (1 for the root, 3 for third, etc). The text attribute specifies how
// the value of the degree should be displayed. The symbol attribute indicates that a
// symbol should be used in specifying the degree. If the symbol attribute is present,
// the value of the text attribute follows the symbol.
type Degree_value struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Direction Named source named complex type "direction"
// A direction is a musical indication that is not necessarily attached
// to a specific note. Two or more may be combined to indicate words followed by the
// start of a dashed line, the end of a wedge followed by dynamics, etc. For
// applications where a specific direction is indeed attached to a specific note, the
// direction element can be associated with the first note element that follows it in
// score order that is not in a different voice. By default, a series of direction-type
// elements and a series of child elements of a direction-type within a single
// direction element follow one another in sequence visually. For a series of
// direction-type children, non-positional formatting attributes are carried over from
// the previous element by default.
type Direction struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "directive
	AttributeGroup_directive

	// generated from attribute group "system-relation
	AttributeGroup_system_relation

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "direction-type" of type direction-type order 489 depth 1
	Direction_type []*Direction_type `xml:"direction-type,omitempty"`

	// generated from element "offset" of type offset order 490 depth 1
	Offset *Offset `xml:"offset,omitempty"`

	// generated from group with order 491 depth 1
	Group_editorial_voice_direction

	// generated from group with order 492 depth 1
	Group_staff

	// generated from element "sound" of type sound order 493 depth 1
	Sound *Sound `xml:"sound,omitempty"`

	// generated from element "listening" of type listening order 494 depth 1
	Listening *Listening `xml:"listening,omitempty"`
}

// Direction_type Named source named complex type "direction-type"
// Textual direction types may have more than 1 component due to multiple
// fonts. The dynamics element may also be used in the notations element. Attribute
// groups related to print suggestions apply to the individual direction-type, not to
// the overall direction.
type Direction_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "rehearsal" of type formatted-text-id order 500 depth 1
	Rehearsal []*Formatted_text_id `xml:"rehearsal,omitempty"`

	// generated from element "segno" of type segno order 501 depth 1
	Segno []*Segno `xml:"segno,omitempty"`

	// generated from element "coda" of type coda order 502 depth 1
	Coda []*Coda `xml:"coda,omitempty"`

	// generated from element "words" of type formatted-text-id order 503 depth 1
	Words []*Formatted_text_id `xml:"words,omitempty"`

	// generated from element "symbol" of type formatted-symbol-id order 504 depth 1
	Symbol []*Formatted_symbol_id `xml:"symbol,omitempty"`

	// generated from element "wedge" of type wedge order 505 depth 1
	Wedge *Wedge `xml:"wedge,omitempty"`

	// generated from element "dynamics" of type dynamics order 506 depth 1
	Dynamics []*Dynamics `xml:"dynamics,omitempty"`

	// generated from element "dashes" of type dashes order 507 depth 1
	Dashes *Dashes `xml:"dashes,omitempty"`

	// generated from element "bracket" of type bracket order 508 depth 1
	Bracket *Bracket `xml:"bracket,omitempty"`

	// generated from element "pedal" of type pedal order 509 depth 1
	Pedal *Pedal `xml:"pedal,omitempty"`

	// generated from element "metronome" of type metronome order 510 depth 1
	Metronome *Metronome `xml:"metronome,omitempty"`

	// generated from element "octave-shift" of type octave-shift order 511 depth 1
	Octave_shift *Octave_shift `xml:"octave-shift,omitempty"`

	// generated from element "harp-pedals" of type harp-pedals order 512 depth 1
	Harp_pedals *Harp_pedals `xml:"harp-pedals,omitempty"`

	// generated from element "damp" of type empty-print-style-align-id order 513 depth 1
	Damp *Empty_print_style_align_id `xml:"damp,omitempty"`

	// generated from element "damp-all" of type empty-print-style-align-id order 514 depth 1
	Damp_all *Empty_print_style_align_id `xml:"damp-all,omitempty"`

	// generated from element "eyeglasses" of type empty-print-style-align-id order 515 depth 1
	Eyeglasses *Empty_print_style_align_id `xml:"eyeglasses,omitempty"`

	// generated from element "string-mute" of type string-mute order 516 depth 1
	String_mute *String_mute `xml:"string-mute,omitempty"`

	// generated from element "scordatura" of type scordatura order 517 depth 1
	Scordatura *Scordatura `xml:"scordatura,omitempty"`

	// generated from element "image" of type image order 518 depth 1
	Image *Image `xml:"image,omitempty"`

	// generated from element "principal-voice" of type principal-voice order 519 depth 1
	Principal_voice *Principal_voice `xml:"principal-voice,omitempty"`

	// generated from element "percussion" of type percussion order 520 depth 1
	Percussion []*Percussion `xml:"percussion,omitempty"`

	// generated from element "accordion-registration" of type accordion-registration order 521 depth 1
	Accordion_registration *Accordion_registration `xml:"accordion-registration,omitempty"`

	// generated from element "staff-divide" of type staff-divide order 522 depth 1
	Staff_divide *Staff_divide `xml:"staff-divide,omitempty"`

	// generated from element "other-direction" of type other-direction order 523 depth 1
	Other_direction *Other_direction `xml:"other-direction,omitempty"`
}

// Effect Named source named complex type "effect"
// The effect type represents pictograms for sound effect percussion
// instruments. The smufl attribute is used to distinguish different SMuFL stylistic
// alternates.
type Effect struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Feature Named source named complex type "feature"
// The feature type is a part of the grouping element used for musical
// analysis. The type attribute represents the type of the feature and the element
// content represents its value. This type is flexible to allow for different analyses.
type Feature struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// First_fret Named source named complex type "first-fret"
// The first-fret type indicates which fret is shown in the top space of
// the frame; it is fret 1 if the element is not present. The optional text attribute
// indicates how this is represented in the fret diagram, while the location attribute
// indicates whether the text appears to the left or right of the frame.
type First_fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Frame Named source named complex type "frame"
// The frame type represents a frame or fretboard diagram used together
// with a chord symbol. The representation is based on the NIFF guitar grid with
// additional information. The frame type's unplayed attribute indicates what to
// display above a string that has no associated frame-note element. Typical values are
// x and the empty string. If the attribute is not present, the display of the unplayed
// string is application-defined.
type Frame struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "height
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute "unplayed
	Unplayed string `xml:"unplayed,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign-image
	AttributeGroup_valign_image

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "frame-strings" of type positiveInteger order 529 depth 1
	Frame_strings int `xml:"frame-strings,omitempty"`

	// generated from element "frame-frets" of type positiveInteger order 530 depth 1
	Frame_frets int `xml:"frame-frets,omitempty"`

	// generated from element "first-fret" of type first-fret order 531 depth 1
	First_fret *First_fret `xml:"first-fret,omitempty"`

	// generated from element "frame-note" of type frame-note order 532 depth 1
	Frame_note []*Frame_note `xml:"frame-note,omitempty"`
}

// Frame_note Named source named complex type "frame-note"
// The frame-note type represents each note included in the frame. An
// open string will have a fret value of 0, while a muted string will not be associated
// with a frame-note element.
type Frame_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "string" of type string-type order 539 depth 1
	String *String_type `xml:"string,omitempty"`

	// generated from element "fret" of type fret order 540 depth 1
	Fret *Fret `xml:"fret,omitempty"`

	// generated from element "fingering" of type fingering order 541 depth 1
	Fingering *Fingering `xml:"fingering,omitempty"`

	// generated from element "barre" of type barre order 542 depth 1
	Barre *Barre `xml:"barre,omitempty"`
}

// Glass Named source named complex type "glass"
// The glass type represents pictograms for glass percussion instruments.
// The smufl attribute is used to distinguish different SMuFL glyphs for wind chimes in
// the Chimes pictograms range, including those made of materials other than glass.
type Glass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grouping Named source named complex type "grouping"
// The grouping type is used for musical analysis. When the type
// attribute is "start" or "single", it usually contains one or more feature elements.
// The number attribute is used for distinguishing between overlapping and hierarchical
// groupings. The member-of attribute allows for easy distinguishing of what grouping
// elements are in what hierarchy. Feature elements contained within a "stop" type of
// grouping may be ignored. This element is flexible to allow for different types of
// analyses. Future versions of the MusicXML format may add elements that can represent
// more standardized categories of analysis data, allowing for easier data sharing.
type Grouping struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "member-of
	Member_of string `xml:"member-of,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "feature" of type feature order 545 depth 1
	Feature []*Feature `xml:"feature,omitempty"`
}

// Harmony Named source named complex type "harmony"
// The harmony type represents harmony analysis, including chord symbols
// in popular music as well as functional harmony analysis in classical music. If there
// are alternate harmonies possible, this can be specified using multiple harmony
// elements differentiated by type. Explicit harmonies have all note present in the
// music; implied have some notes missing but implied; alternate represents alternate
// analyses. The print-object attribute controls whether or not anything is printed due
// to the harmony element. The print-frame attribute controls printing of a frame or
// fretboard diagram. The print-style attribute group sets the default for the harmony,
// but individual elements can override this with their own print-style values. The
// arrangement attribute specifies how multiple harmony-chord groups are arranged
// relative to each other. Harmony-chords with vertical arrangement are separated by
// horizontal lines. Harmony-chords with diagonal or horizontal arrangement are
// separated by diagonal lines or slashes.
type Harmony struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "print-frame
	Print_frame Enum_Yes_no `xml:"print-frame,attr,omitempty"`

	// generated from attribute "arrangement
	Arrangement Enum_Harmony_arrangement `xml:"arrangement,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "system-relation
	AttributeGroup_system_relation

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 548 depth 1
	Group_harmony_chord

	// generated from element "frame" of type frame order 549 depth 1
	Frame *Frame `xml:"frame,omitempty"`

	// generated from element "offset" of type offset order 550 depth 1
	Offset *Offset `xml:"offset,omitempty"`

	// generated from group with order 551 depth 1
	Group_editorial

	// generated from group with order 552 depth 1
	Group_staff
}

// Harp_pedals Named source named complex type "harp-pedals"
// The harp-pedals type is used to create harp pedal diagrams. The
// pedal-step and pedal-alter elements use the same values as the step and alter
// elements. For easiest reading, the pedal-tuning elements should follow standard harp
// pedal order, with pedal-step values of D, C, B, E, F, G, and A.
type Harp_pedals struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "pedal-tuning" of type pedal-tuning order 559 depth 1
	Pedal_tuning []*Pedal_tuning `xml:"pedal-tuning,omitempty"`
}

// Image Named source named complex type "image"
// The image type is used to include graphical images in a score.
type Image struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "image-attributes
	AttributeGroup_image_attributes

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Instrument_change Named source named complex type "instrument-change"
// The instrument-change element type represents a change to the virtual
// instrument sound for a given score-instrument. The id attribute refers to the
// score-instrument affected by the change. All instrument-change child elements can
// also be initially specified within the score-instrument element.
type Instrument_change struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from group with order 566 depth 1
	Group_virtual_instrument_data
}

// Inversion Named source named complex type "inversion"
// The inversion type represents harmony inversions. The value is a
// number indicating which inversion is used: 0 for root position, 1 for first
// inversion, etc. The text attribute indicates how the inversion should be displayed
// in a score.
type Inversion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Kind Named source named complex type "kind"
// Kind indicates the type of chord. Degree elements can then add,
// subtract, or alter from these starting points The attributes are used to indicate
// the formatting of the symbol. Since the kind element is the constant in all the
// harmony-chord groups that can make up a polychord, many formatting attributes are
// here. The use-symbols attribute is yes if the kind should be represented when
// possible with harmony symbols rather than letters and numbers. These symbols
// include: major: a triangle, like Unicode 25B3 minor: -, like Unicode 002D augmented:
// +, like Unicode 002B diminished: °, like Unicode 00B0 half-diminished: ø, like
// Unicode 00F8 For the major-minor kind, only the minor symbol is used when
// use-symbols is yes. The major symbol is set using the symbol attribute in the
// degree-value element. The corresponding degree-alter value will usually be 0 in this
// case. The text attribute describes how the kind should be spelled in a score. If
// use-symbols is yes, the value of the text attribute follows the symbol. The
// stack-degrees attribute is yes if the degree elements should be stacked above each
// other. The parentheses-degrees attribute is yes if all the degrees should be in
// parentheses. The bracket-degrees attribute is yes if all the degrees should be in a
// bracket. If not specified, these values are implementation-specific. The alignment
// attributes are for the entire harmony-chord group of which this kind element is a
// part. The text attribute may use strings such as "13sus" that refer to both the kind
// and one or more degree elements. In this case, the corresponding degree elements
// should have the print-object attribute set to "no" to keep redundant alterations
// from being displayed.
type Kind struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols
	Use_symbols Enum_Yes_no `xml:"use-symbols,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "stack-degrees
	Stack_degrees Enum_Yes_no `xml:"stack-degrees,attr,omitempty"`

	// generated from attribute "parentheses-degrees
	Parentheses_degrees Enum_Yes_no `xml:"parentheses-degrees,attr,omitempty"`

	// generated from attribute "bracket-degrees
	Bracket_degrees Enum_Yes_no `xml:"bracket-degrees,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign
	AttributeGroup_valign

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Listening Named source named complex type "listening"
// The listen and listening types, new in Version 4.0, specify different
// ways that a score following or machine listening application can interact with a
// performer. The listening type handles interactions that change the state of the
// listening application from the specified point in the performance onward. If
// multiple child elements of the same type are present, they should have distinct
// player and/or time-only attributes. The offset element is used to indicate that the
// listening change takes place offset from the current score position. If the
// listening element is a child of a direction element, the listening offset element
// overrides the direction offset element if both elements are present. Note that the
// offset reflects the intended musical position for the change in state. It should not
// be used to compensate for latency issues in particular hardware configurations.
type Listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sync" of type sync order 574 depth 1
	Sync []*Sync `xml:"sync,omitempty"`

	// generated from element "other-listening" of type other-listening order 575 depth 1
	Other_listening []*Other_listening `xml:"other-listening,omitempty"`

	// generated from element "offset" of type offset order 576 depth 1
	Offset *Offset `xml:"offset,omitempty"`
}

// Measure_numbering Named source named complex type "measure-numbering"
// The measure-numbering type describes how frequently measure numbers
// are displayed on this part. The text attribute from the measure element is used for
// display, or the number attribute if the text attribute is not present. Measures with
// an implicit attribute set to "yes" never display a measure number, regardless of the
// measure-numbering setting. The optional staff attribute refers to staff numbers
// within the part, from top to bottom on the system. It indicates which staff is used
// as the reference point for vertical positioning. A value of 1 is assumed if not
// present. The optional multiple-rest-always and multiple-rest-range attributes
// describe how measure numbers are shown on multiple rests when the measure-numbering
// value is not set to none. The multiple-rest-always attribute is set to yes when the
// measure number should always be shown, even if the multiple rest starts midway
// through a system when measure numbering is set to system level. The
// multiple-rest-range attribute is set to yes when measure numbers on multiple rests
// display the range of numbers for the first and last measure, rather than just the
// number of the first measure.
type Measure_numbering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "system
	System string `xml:"system,attr,omitempty"`

	// generated from attribute "staff
	Staff int `xml:"staff,attr,omitempty"`

	// generated from attribute "multiple-rest-always
	Multiple_rest_always Enum_Yes_no `xml:"multiple-rest-always,attr,omitempty"`

	// generated from attribute "multiple-rest-range
	Multiple_rest_range Enum_Yes_no `xml:"multiple-rest-range,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Membrane Named source named complex type "membrane"
// The membrane type represents pictograms for membrane percussion
// instruments. The smufl attribute is used to distinguish different SMuFL stylistic
// alternates.
type Membrane struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metal Named source named complex type "metal"
// The metal type represents pictograms for metal percussion instruments.
// The smufl attribute is used to distinguish different SMuFL stylistic alternates.
type Metal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome Named source named complex type "metronome"
// The metronome type represents metronome marks and other metric
// relationships. The beat-unit group and per-minute element specify regular metronome
// marks. The metronome-note and metronome-relation elements allow for the
// specification of metric modulations and other metric relationships, such as swing
// tempo marks where two eighths are equated to a quarter note / eighth note triplet.
// Tied notes can be represented in both types of metronome marks by using the
// beat-unit-tied and metronome-tied elements. The parentheses attribute indicates
// whether or not to put the metronome mark in parentheses; its value is no if not
// specified. The print-object attribute is set to no in cases where the metronome
// element represents a relationship or range that is not displayed in the music
// notation.
type Metronome struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses Enum_Yes_no `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 582 depth 1
	Group_beat_unit

	// generated from element "per-minute" of type per-minute order 584 depth 1
	Per_minute *Per_minute `xml:"per-minute,omitempty"`

	// generated from element "beat-unit-tied" of type beat-unit-tied order 586 depth 1
	Beat_unit_tied []*Beat_unit_tied `xml:"beat-unit-tied,omitempty"`

	// generated from element "metronome-arrows" of type empty order 587 depth 1
	Metronome_arrows string `xml:"metronome-arrows,omitempty"`

	// generated from element "metronome-relation" of type string order 589 depth 1
	Metronome_relation string `xml:"metronome-relation,omitempty"`

	// generated from element "metronome-note" of type metronome-note order 590 depth 1
	Metronome_note []*Metronome_note `xml:"metronome-note,omitempty"`
}

// Metronome_beam Named source named complex type "metronome-beam"
// The metronome-beam type works like the beam type in defining metric
// relationships, but does not include all the attributes available in the beam type.
type Metronome_beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Beam_value `xml:",chardata"`
}

// Metronome_note Named source named complex type "metronome-note"
// The metronome-note type defines the appearance of a note within a
// metric relationship mark.
type Metronome_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "metronome-type" of type note-type-value order 597 depth 1
	Metronome_type string `xml:"metronome-type,omitempty"`

	// generated from element "metronome-dot" of type empty order 598 depth 1
	Metronome_dot string `xml:"metronome-dot,omitempty"`

	// generated from element "metronome-beam" of type metronome-beam order 599 depth 1
	Metronome_beam []*Metronome_beam `xml:"metronome-beam,omitempty"`

	// generated from element "metronome-tied" of type metronome-tied order 600 depth 1
	Metronome_tied *Metronome_tied `xml:"metronome-tied,omitempty"`

	// generated from element "metronome-tuplet" of type metronome-tuplet order 601 depth 1
	Metronome_tuplet *Metronome_tuplet `xml:"metronome-tuplet,omitempty"`
}

// Metronome_tied Named source named complex type "metronome-tied"
// The metronome-tied indicates the presence of a tie within a metric
// relationship mark. As with the tied element, both the start and stop of the tie
// should be specified, in this case within separate metronome-note elements.
type Metronome_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`
}

// Metronome_tuplet Named source named complex type "metronome-tuplet"
// The metronome-tuplet type uses the same element structure as the
// time-modification element along with some attributes from the tuplet element.
type Metronome_tuplet struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Numeral Named source named complex type "numeral"
// The numeral type represents the Roman numeral or Nashville number part
// of a harmony. It requires that the key be specified in the encoding, either with a
// key or numeral-key element.
type Numeral struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "numeral-root" of type numeral-root order 605 depth 1
	Numeral_root *Numeral_root `xml:"numeral-root,omitempty"`

	// generated from element "numeral-alter" of type harmony-alter order 606 depth 1
	Numeral_alter *Harmony_alter `xml:"numeral-alter,omitempty"`

	// generated from element "numeral-key" of type numeral-key order 607 depth 1
	Numeral_key *Numeral_key `xml:"numeral-key,omitempty"`
}

// Numeral_key Named source named complex type "numeral-key"
// The numeral-key type is used when the key for the numeral is different
// than the key specified by the key signature. The numeral-fifths element specifies
// the key in the same way as the fifths element. The numeral-mode element specifies
// the mode similar to the mode element, but with a restricted set of values
type Numeral_key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "numeral-fifths" of type fifths order 609 depth 1
	Numeral_fifths int `xml:"numeral-fifths,omitempty"`

	// generated from element "numeral-mode" of type numeral-mode order 610 depth 1
	Numeral_mode string `xml:"numeral-mode,omitempty"`
}

// Numeral_root Named source named complex type "numeral-root"
// The numeral-root type represents the Roman numeral or Nashville number
// as a positive integer from 1 to 7. The text attribute indicates how the numeral
// should appear in the score. A numeral-root value of 5 with a kind of major would
// have a text attribute of "V" if displayed as a Roman numeral, and "5" if displayed
// as a Nashville number. If the text attribute is not specified, the display is
// application-dependent.
type Numeral_root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Octave_shift Named source named complex type "octave-shift"
// The octave shift type indicates where notes are shifted up or down
// from their true pitched values because of printing difficulty. Thus a treble clef
// line noted with 8va will be indicated with an octave-shift down from the pitch data
// indicated in the notes. A size of 8 indicates one octave; a size of 15 indicates two
// octaves.
type Octave_shift struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "size
	Size int `xml:"size,attr,omitempty"`

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Offset Named source named complex type "offset"
// An offset is represented in terms of divisions, and indicates where
// the direction will appear relative to the current musical location. The current
// musical location is always within the current measure, even at the end of a measure.
// The offset affects the visual appearance of the direction. If the sound attribute is
// "yes", then the offset affects playback and listening too. If the sound attribute is
// "no", then any sound or listening associated with the direction takes effect at the
// current location. The sound attribute is "no" by default for compatibility with
// earlier versions of the MusicXML format. If an element within a direction includes a
// default-x attribute, the offset value will be ignored when determining the
// appearance of that element.
type Offset struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "sound
	Sound Enum_Yes_no `xml:"sound,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_direction Named source named complex type "other-direction"
// The other-direction type is used to define any direction symbols not
// yet in the MusicXML format. The smufl attribute can be used to specify a particular
// direction symbol, allowing application interoperability without requiring every
// SMuFL glyph to have a MusicXML element equivalent. Using the other-direction type
// without the smufl attribute allows for extended representation, though without
// application interoperability.
type Other_direction struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_listening Named source named complex type "other-listening"
// The other-listening type represents other types of listening control
// and interaction. The required type attribute indicates the type of listening to
// which the element content applies. The optional player and time-only attributes
// restrict the element to apply to a single player or set of times through a repeated
// section, respectively.
type Other_listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Pedal Named source named complex type "pedal"
// The pedal type represents piano pedal marks, including damper and
// sostenuto pedal marks. The line attribute is yes if pedal lines are used. The sign
// attribute is yes if Ped, Sost, and * signs are used. For compatibility with older
// versions, the sign attribute is yes by default if the line attribute is no, and is
// no by default if the line attribute is yes. If the sign attribute is set to yes and
// the type is start or sostenuto, the abbreviated attribute is yes if the short P and
// S signs are used, and no if the full Ped and Sost signs are used. It is no by
// default. Otherwise the abbreviated attribute is ignored. The alignment attributes
// are ignored if the sign attribute is no.
type Pedal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line
	Line Enum_Yes_no `xml:"line,attr,omitempty"`

	// generated from attribute "sign
	Sign Enum_Yes_no `xml:"sign,attr,omitempty"`

	// generated from attribute "abbreviated
	Abbreviated Enum_Yes_no `xml:"abbreviated,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Pedal_tuning Named source named complex type "pedal-tuning"
// The pedal-tuning type specifies the tuning of a single harp pedal.
type Pedal_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "pedal-step" of type step order 629 depth 1
	Pedal_step Enum_Step `xml:"pedal-step,omitempty"`

	// generated from element "pedal-alter" of type semitones order 630 depth 1
	Pedal_alter string `xml:"pedal-alter,omitempty"`
}

// Per_minute Named source named complex type "per-minute"
// The per-minute type can be a number, or a text description including
// numbers. If a font is specified, it overrides the font specified for the overall
// metronome element. This allows separate specification of a music font for the
// beat-unit and a text font for the numeric value, in cases where a single metronome
// font is not used.
type Per_minute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Percussion Named source named complex type "percussion"
// The percussion element is used to define percussion pictogram symbols.
// Definitions for these symbols can be found in Kurt Stone's "Music Notation in the
// Twentieth Century" on pages 206-212 and 223. Some values are added to these based on
// how usage has evolved in the 30 years since Stone's book was published.
type Percussion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "enclosure
	AttributeGroup_enclosure

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "glass" of type glass order 634 depth 1
	Glass *Glass `xml:"glass,omitempty"`

	// generated from element "metal" of type metal order 635 depth 1
	Metal *Metal `xml:"metal,omitempty"`

	// generated from element "wood" of type wood order 636 depth 1
	Wood *Wood `xml:"wood,omitempty"`

	// generated from element "pitched" of type pitched order 637 depth 1
	Pitched *Pitched `xml:"pitched,omitempty"`

	// generated from element "membrane" of type membrane order 638 depth 1
	Membrane *Membrane `xml:"membrane,omitempty"`

	// generated from element "effect" of type effect order 639 depth 1
	Effect *Effect `xml:"effect,omitempty"`

	// generated from element "timpani" of type timpani order 640 depth 1
	Timpani *Timpani `xml:"timpani,omitempty"`

	// generated from element "beater" of type beater order 641 depth 1
	Beater *Beater `xml:"beater,omitempty"`

	// generated from element "stick" of type stick order 642 depth 1
	Stick *Stick `xml:"stick,omitempty"`

	// generated from element "stick-location" of type stick-location order 643 depth 1
	Stick_location string `xml:"stick-location,omitempty"`

	// generated from element "other-percussion" of type other-text order 644 depth 1
	Other_percussion *Other_text `xml:"other-percussion,omitempty"`
}

// Pitched Named source named complex type "pitched"
// The pitched-value type represents pictograms for pitched percussion
// instruments. The smufl attribute is used to distinguish different SMuFL glyphs for a
// particular pictogram within the Tuned mallet percussion pictograms range.
type Pitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Principal_voice Named source named complex type "principal-voice"
// The principal-voice type represents principal and secondary voices in
// a score, either for analysis or for square bracket symbols that appear in a score.
// The element content is used for analysis and may be any text value. The symbol
// attribute indicates the type of symbol used. When used for analysis separate from
// any printed score markings, it should be set to none. Otherwise if the type is stop
// it should be set to plain.
type Principal_voice struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Print Named source named complex type "print"
// The print type contains general printing parameters, including layout
// elements. The part-name-display and part-abbreviation-display elements may also be
// used here to change how a part name or abbreviation is displayed over the course of
// a piece. They take effect when the current measure or a succeeding measure starts a
// new system. Layout group elements in a print element only apply to the current page,
// system, or staff. Music that follows continues to take the default values from the
// layout determined by the defaults element.
type Print struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-attributes
	AttributeGroup_print_attributes

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 653 depth 1
	Group_layout

	// generated from element "measure-layout" of type measure-layout order 654 depth 1
	Measure_layout *Measure_layout `xml:"measure-layout,omitempty"`

	// generated from element "measure-numbering" of type measure-numbering order 655 depth 1
	Measure_numbering *Measure_numbering `xml:"measure-numbering,omitempty"`

	// generated from element "part-name-display" of type name-display order 656 depth 1
	Part_name_display *Name_display `xml:"part-name-display,omitempty"`

	// generated from element "part-abbreviation-display" of type name-display order 657 depth 1
	Part_abbreviation_display *Name_display `xml:"part-abbreviation-display,omitempty"`
}

// Root Named source named complex type "root"
// The root type indicates a pitch like C, D, E vs. a scale degree like
// 1, 2, 3. It is used with chord symbols in popular music. The root element has a
// root-step and optional root-alter element similar to the step and alter elements,
// but renamed to distinguish the different musical meanings.
type Root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "root-step" of type root-step order 661 depth 1
	Root_step *Root_step `xml:"root-step,omitempty"`

	// generated from element "root-alter" of type harmony-alter order 662 depth 1
	Root_alter *Harmony_alter `xml:"root-alter,omitempty"`
}

// Root_step Named source named complex type "root-step"
// The root-step type represents the pitch step of the root of the
// current chord within the harmony element. The text attribute indicates how the root
// should appear in a score if not using the element contents.
type Root_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Step `xml:",chardata"`
}

// Scordatura Named source named complex type "scordatura"
// Scordatura string tunings are represented by a series of accord
// elements, similar to the staff-tuning elements. Strings are numbered from high to
// low.
type Scordatura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accord" of type accord order 666 depth 1
	Accord []*Accord `xml:"accord,omitempty"`
}

// Sound Named source named complex type "sound"
// The sound element contains general playback parameters. They can stand
// alone within a part/measure, or be a component element within a direction. Tempo is
// expressed in quarter notes per minute. If 0, the sound-generating program should
// prompt the user at the time of compiling a sound (MIDI) file. Dynamics (or MIDI
// velocity) are expressed as a percentage of the default forte value (90 for MIDI
// 1.0). Dacapo indicates to go back to the beginning of the movement. When used it
// always has the value "yes". Segno and dalsegno are used for backwards jumps to a
// segno sign; coda and tocoda are used for forward jumps to a coda sign. If there are
// multiple jumps, the value of these parameters can be used to name and distinguish
// them. If segno or coda is used, the divisions attribute can also be used to indicate
// the number of divisions per quarter note. Otherwise sound and MIDI generating
// programs may have to recompute this. By default, a dalsegno or dacapo attribute
// indicates that the jump should occur the first time through, while a tocoda
// attribute indicates the jump should occur the second time through. The time that
// jumps occur can be changed by using the time-only attribute. The forward-repeat
// attribute indicates that a forward repeat sign is implied but not displayed. It is
// used for example in two-part forms with repeats, such as a minuet and trio where no
// repeat is displayed at the start of the trio. This usually occurs after a barline.
// When used it always has the value of "yes". The fine attribute follows the final
// note or rest in a movement with a da capo or dal segno direction. If numeric, the
// value represents the actual duration of the final note or rest, which can be
// ambiguous in written notation and different among parts and voices. The value may
// also be "yes" to indicate no change to the final duration. If the sound element
// applies only particular times through a repeat, the time-only attribute indicates
// which times to apply the sound element. Pizzicato in a sound element effects all
// following notes. Yes indicates pizzicato, no indicates arco. The pan and elevation
// attributes are deprecated in Version 2.0. The pan and elevation elements in the
// midi-instrument element should be used instead. The meaning of the pan and elevation
// attributes is the same as for the pan and elevation elements. If both are present,
// the mid-instrument elements take priority. The damper-pedal, soft-pedal, and
// sostenuto-pedal attributes effect playback of the three common piano pedals and
// their MIDI controller equivalents. The yes value indicates the pedal is depressed;
// no indicates the pedal is released. A numeric value from 0 to 100 may also be used
// for half pedaling. This value is the percentage that the pedal is depressed. A value
// of 0 is equivalent to no, and a value of 100 is equivalent to yes. Instrument
// changes, MIDI devices, MIDI instruments, and playback techniques are changed using
// the instrument-change, midi-device, midi-instrument, and play elements. When there
// are multiple instances of these elements, they should be grouped together by
// instrument using the id attribute values. The offset element is used to indicate
// that the sound takes place offset from the current score position. If the sound
// element is a child of a direction element, the sound offset element overrides the
// direction offset element if both elements are present. Note that the offset reflects
// the intended musical position for the change in sound. It should not be used to
// compensate for latency issues in particular hardware configurations.
type Sound struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tempo
	Tempo string `xml:"tempo,attr,omitempty"`

	// generated from attribute "dynamics
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "dacapo
	Dacapo Enum_Yes_no `xml:"dacapo,attr,omitempty"`

	// generated from attribute "segno
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "dalsegno
	Dalsegno string `xml:"dalsegno,attr,omitempty"`

	// generated from attribute "coda
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "tocoda
	Tocoda string `xml:"tocoda,attr,omitempty"`

	// generated from attribute "divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute "forward-repeat
	Forward_repeat Enum_Yes_no `xml:"forward-repeat,attr,omitempty"`

	// generated from attribute "fine
	Fine string `xml:"fine,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato
	Pizzicato Enum_Yes_no `xml:"pizzicato,attr,omitempty"`

	// generated from attribute "pan
	Pan string `xml:"pan,attr,omitempty"`

	// generated from attribute "elevation
	Elevation string `xml:"elevation,attr,omitempty"`

	// generated from attribute "damper-pedal
	Damper_pedal string `xml:"damper-pedal,attr,omitempty"`

	// generated from attribute "soft-pedal
	Soft_pedal string `xml:"soft-pedal,attr,omitempty"`

	// generated from attribute "sostenuto-pedal
	Sostenuto_pedal string `xml:"sostenuto-pedal,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "instrument-change" of type instrument-change order 669 depth 1
	Instrument_change []*Instrument_change `xml:"instrument-change,omitempty"`

	// generated from element "midi-device" of type midi-device order 670 depth 1
	Midi_device []*Midi_device `xml:"midi-device,omitempty"`

	// generated from element "midi-instrument" of type midi-instrument order 671 depth 1
	Midi_instrument []*Midi_instrument `xml:"midi-instrument,omitempty"`

	// generated from element "play" of type play order 672 depth 1
	Play []*Play `xml:"play,omitempty"`

	// generated from element "swing" of type swing order 673 depth 1
	Swing *Swing `xml:"swing,omitempty"`

	// generated from element "offset" of type offset order 674 depth 1
	Offset *Offset `xml:"offset,omitempty"`
}

// Staff_divide Named source named complex type "staff-divide"
// The staff-divide element represents the staff division arrow symbols
// found at SMuFL code points U+E00B, U+E00C, and U+E00D.
type Staff_divide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Stick Named source named complex type "stick"
// The stick type represents pictograms where the material of the stick,
// mallet, or beater is included.The parentheses and dashed-circle attributes indicate
// the presence of these marks around the round beater part of a pictogram. Values for
// these attributes are "no" if not present.
type Stick struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip
	Tip Enum_Tip_direction `xml:"tip,attr,omitempty"`

	// generated from attribute "parentheses
	Parentheses Enum_Yes_no `xml:"parentheses,attr,omitempty"`

	// generated from attribute "dashed-circle
	Dashed_circle Enum_Yes_no `xml:"dashed-circle,attr,omitempty"`

	// generated from element "stick-type" of type stick-type order 680 depth 1
	Stick_type string `xml:"stick-type,omitempty"`

	// generated from element "stick-material" of type stick-material order 681 depth 1
	Stick_material string `xml:"stick-material,omitempty"`
}

// String_mute Named source named complex type "string-mute"
// The string-mute type represents string mute on and mute off symbols.
type String_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Swing Named source named complex type "swing"
// The swing element specifies whether or not to use swing playback,
// where consecutive on-beat / off-beat eighth or 16th notes are played with unequal
// nominal durations. The straight element specifies that no swing is present, so
// consecutive notes have equal durations. The first and second elements are positive
// integers that specify the ratio between durations of consecutive notes. For example,
// a first element with a value of 2 and a second element with a value of 1 applied to
// eighth notes specifies a quarter note / eighth note tuplet playback, where the first
// note is twice as long as the second note. Ratios should be specified with the
// smallest integers possible. For example, a ratio of 6 to 4 should be specified as 3
// to 2 instead. The optional swing-type element specifies the note type, either eighth
// or 16th, to which the ratio is applied. The value is eighth if this element is not
// present. The optional swing-style element is a string describing the style of swing
// used. The swing element has no effect for playback of grace notes, notes where a
// type element is not present, and notes where the specified duration is different
// than the nominal value associated with the specified type. If a swung note has
// attack and release attributes, those values modify the swung playback.
type Swing struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "straight" of type empty order 686 depth 1
	Straight string `xml:"straight,omitempty"`

	// generated from element "first" of type positiveInteger order 687 depth 1
	First int `xml:"first,omitempty"`

	// generated from element "second" of type positiveInteger order 688 depth 1
	Second int `xml:"second,omitempty"`

	// generated from element "swing-type" of type swing-type-value order 689 depth 1
	Swing_type Enum_Note_type_value `xml:"swing-type,omitempty"`

	// generated from element "swing-style" of type string order 690 depth 1
	Swing_style string `xml:"swing-style,omitempty"`
}

// Sync Named source named complex type "sync"
// The sync type specifies the style that a score following application
// should use the synchronize an accompaniment with a performer. If this type is not
// included in a score, default synchronization depends on the application. The
// optional latency attribute specifies a time in milliseconds that the listening
// application should expect from the performer. The optional player and time-only
// attributes restrict the element to apply to a single player or set of times through
// a repeated section, respectively.
type Sync struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "latency
	Latency int `xml:"latency,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`
}

// Timpani Named source named complex type "timpani"
// The timpani type represents the timpani pictogram. The smufl attribute
// is used to distinguish different SMuFL stylistic alternates.
type Timpani struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`
}

// Wedge Named source named complex type "wedge"
// The wedge type represents crescendo and diminuendo wedge symbols. The
// type attribute is crescendo for the start of a wedge that is closed at the left
// side, and diminuendo for the start of a wedge that is closed on the right side.
// Spread values are measured in tenths; those at the start of a crescendo wedge or end
// of a diminuendo wedge are ignored. The niente attribute is yes if a circle appears
// at the point of the wedge, indicating a crescendo from nothing or diminuendo to
// nothing. It is no by default, and used only when the type is crescendo, or the type
// is stop for a wedge that began with a diminuendo type. The line-type is solid if not
// specified.
type Wedge struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "spread
	Spread string `xml:"spread,attr,omitempty"`

	// generated from attribute "niente
	Niente Enum_Yes_no `xml:"niente,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Wood Named source named complex type "wood"
// The wood type represents pictograms for wood percussion instruments.
// The smufl attribute is used to distinguish different SMuFL stylistic alternates.
type Wood struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Encoding Named source named complex type "encoding"
// The encoding element contains information about who did the digital
// encoding, when, with what software, and in what aspects. Standard type values for
// the encoder element are music, words, and arrangement, but other types may be used.
// The type attribute is only needed when there are multiple encoder elements.
type Encoding struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "encoder" of type typed-text order 702 depth 1
	Encoder []*Typed_text `xml:"encoder,omitempty"`

	// generated from element "software" of type string order 703 depth 1
	Software string `xml:"software,omitempty"`

	// generated from element "encoding-description" of type string order 704 depth 1
	Encoding_description string `xml:"encoding-description,omitempty"`

	// generated from element "supports" of type supports order 705 depth 1
	Supports []*Supports `xml:"supports,omitempty"`
}

// Identification Named source named complex type "identification"
// Identification contains basic metadata about the score. It includes
// information that may apply at a score-wide, movement-wide, or part-wide level. The
// creator, rights, source, and relation elements are based on Dublin Core.
type Identification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "creator" of type typed-text order 707 depth 1
	Creator []*Typed_text `xml:"creator,omitempty"`

	// generated from element "rights" of type typed-text order 708 depth 1
	Rights []*Typed_text `xml:"rights,omitempty"`

	// generated from element "encoding" of type encoding order 709 depth 1
	Encoding *Encoding `xml:"encoding,omitempty"`

	// generated from element "source" of type string order 710 depth 1
	Source string `xml:"source,omitempty"`

	// generated from element "relation" of type typed-text order 711 depth 1
	Relation []*Typed_text `xml:"relation,omitempty"`

	// generated from element "miscellaneous" of type miscellaneous order 712 depth 1
	Miscellaneous *Miscellaneous `xml:"miscellaneous,omitempty"`
}

// Miscellaneous Named source named complex type "miscellaneous"
// If a program has other metadata not yet supported in the MusicXML
// format, it can go in the miscellaneous element. The miscellaneous type puts each
// separate part of metadata into its own miscellaneous-field type.
type Miscellaneous struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "miscellaneous-field" of type miscellaneous-field order 714 depth 1
	Miscellaneous_field []*Miscellaneous_field `xml:"miscellaneous-field,omitempty"`
}

// Miscellaneous_field Named source named complex type "miscellaneous-field"
// If a program has other metadata not yet supported in the MusicXML
// format, each type of metadata can go in a miscellaneous-field element. The required
// name attribute indicates the type of metadata the element content represents.
type Miscellaneous_field struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Supports Named source named complex type "supports"
// The supports type indicates if a MusicXML encoding supports a
// particular MusicXML element. This is recommended for elements like beam, stem, and
// accidental, where the absence of an element is ambiguous if you do not know if the
// encoding supports that element. For Version 2.0, the supports element is expanded to
// allow programs to indicate support for particular attributes or particular values.
// This lets applications communicate, for example, that all system and/or page breaks
// are contained in the MusicXML file.
type Supports struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Yes_no `xml:"type,attr,omitempty"`

	// generated from attribute "element
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "attribute
	Attribute string `xml:"attribute,attr,omitempty"`

	// generated from attribute "value
	Value string `xml:"value,attr,omitempty"`
}

// Appearance Named source named complex type "appearance"
// The appearance type controls general graphical settings for the
// music's final form appearance on a printed page of display. This includes support
// for line widths, definitions for note sizes, and standard distances between notation
// elements, plus an extension element for other aspects of appearance.
type Appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "line-width" of type line-width order 718 depth 1
	Line_width []*Line_width `xml:"line-width,omitempty"`

	// generated from element "note-size" of type note-size order 719 depth 1
	Note_size []*Note_size `xml:"note-size,omitempty"`

	// generated from element "distance" of type distance order 720 depth 1
	Distance []*Distance `xml:"distance,omitempty"`

	// generated from element "glyph" of type glyph order 721 depth 1
	Glyph []*Glyph `xml:"glyph,omitempty"`

	// generated from element "other-appearance" of type other-appearance order 722 depth 1
	Other_appearance []*Other_appearance `xml:"other-appearance,omitempty"`
}

// Distance Named source named complex type "distance"
// The distance element represents standard distances between notation
// elements in tenths. The type attribute defines what type of distance is being
// defined. Valid values include hyphen (for hyphens in lyrics) and beam.
type Distance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glyph Named source named complex type "glyph"
// The glyph element represents what SMuFL glyph should be used for
// different variations of symbols that are semantically identical. The type attribute
// specifies what type of glyph is being defined. The element value specifies what
// SMuFL glyph to use, including recommended stylistic alternates. The SMuFL glyph name
// should match the type. For instance, a type of quarter-rest would use values
// restQuarter, restQuarterOld, or restQuarterZ. A type of g-clef-ottava-bassa would
// use values gClef8vb, gClef8vbOld, or gClef8vbCClef. A type of octave-shift-up-8
// would use values ottava, ottavaBassa, ottavaBassaBa, ottavaBassaVb, or octaveBassa.
type Glyph struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Line_width Named source named complex type "line-width"
// The line-width type indicates the width of a line type in tenths. The
// type attribute defines what type of line is being defined. Values include beam,
// bracket, dashes, enclosure, ending, extend, heavy barline, leger, light barline,
// octave shift, pedal, slur middle, slur tip, staff, stem, tie middle, tie tip, tuplet
// bracket, and wedge. The text content is expressed in tenths.
type Line_width struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_layout Named source named complex type "measure-layout"
// The measure-layout type includes the horizontal distance from the
// previous measure. It applies to the current measure only.
type Measure_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "measure-distance" of type tenths order 727 depth 1
	Measure_distance string `xml:"measure-distance,omitempty"`
}

// Note_size Named source named complex type "note-size"
// The note-size type indicates the percentage of the regular note size
// to use for notes with a cue and large size as defined in the type element. The grace
// type is used for notes of cue size that that include a grace element. The cue type
// is used for all other notes with cue size, whether defined explicitly or implicitly
// via a cue element. The large type is used for notes of large size. The text content
// represent the numeric percentage. A value of 100 would be identical to the size of a
// regular note as defined by the music font.
type Note_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_appearance Named source named complex type "other-appearance"
// The other-appearance type is used to define any graphical settings not
// yet in the current version of the MusicXML format. This allows extended
// representation, though without application interoperability.
type Other_appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Page_layout Named source named complex type "page-layout"
// Page layout can be defined both in score-wide defaults and in the
// print element. Page margins are specified either for both even and odd pages, or via
// separate odd and even page number values. The type is not needed when used as part
// of a print element. If omitted when used in the defaults element, "both" is the
// default. If no page-layout element is present in the defaults element, default page
// layout values are chosen by the application. When used in the print element, the
// page-layout element affects the appearance of the current page only. All other pages
// use the default values as determined by the defaults element. If any child elements
// are missing from the page-layout element in a print element, the values determined
// by the defaults element are used there as well.
type Page_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "page-height" of type tenths order 731 depth 1
	Page_height string `xml:"page-height,omitempty"`

	// generated from element "page-width" of type tenths order 732 depth 1
	Page_width string `xml:"page-width,omitempty"`

	// generated from element "page-margins" of type page-margins order 733 depth 1
	Page_margins *Page_margins `xml:"page-margins,omitempty"`
}

// Page_margins Named source named complex type "page-margins"
// Page margins are specified either for both even and odd pages, or via
// separate odd and even page number values. The type attribute is not needed when used
// as part of a print element. If omitted when the page-margins type is used in the
// defaults element, "both" is the default value.
type Page_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from group with order 735 depth 1
	Group_all_margins
}

// Scaling Named source named complex type "scaling"
// Margins, page sizes, and distances are all measured in tenths to keep
// MusicXML data in a consistent coordinate system as much as possible. The translation
// to absolute units is done with the scaling type, which specifies how many
// millimeters are equal to how many tenths. For a staff height of 7 mm, millimeters
// would be set to 7 while tenths is set to 40. The ability to set a formula rather
// than a single scaling factor helps avoid roundoff errors.
type Scaling struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "millimeters" of type millimeters order 737 depth 1
	Millimeters string `xml:"millimeters,omitempty"`

	// generated from element "tenths" of type tenths order 738 depth 1
	Tenths string `xml:"tenths,omitempty"`
}

// Staff_layout Named source named complex type "staff-layout"
// Staff layout includes the vertical distance from the bottom line of
// the previous staff in this system to the top line of the staff specified by the
// number attribute. The optional number attribute refers to staff numbers within the
// part, from top to bottom on the system. A value of 1 is used if not present. When
// used in the defaults element, the values apply to all systems in all parts. When
// used in the print element, the values apply to the current system only. This value
// is ignored for the first staff in a system.
type Staff_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from element "staff-distance" of type tenths order 740 depth 1
	Staff_distance string `xml:"staff-distance,omitempty"`
}

// System_dividers Named source named complex type "system-dividers"
// The system-dividers element indicates the presence or absence of
// system dividers (also known as system separation marks) between systems displayed on
// the same page. Dividers on the left and right side of the page are controlled by the
// left-divider and right-divider elements respectively. The default vertical position
// is half the system-distance value from the top of the system that is below the
// divider. The default horizontal position is the left and right system margin,
// respectively. When used in the print element, the system-dividers element affects
// the dividers that would appear between the current system and the previous system.
type System_dividers struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-divider" of type empty-print-object-style-align order 742 depth 1
	Left_divider *Empty_print_object_style_align `xml:"left-divider,omitempty"`

	// generated from element "right-divider" of type empty-print-object-style-align order 743 depth 1
	Right_divider *Empty_print_object_style_align `xml:"right-divider,omitempty"`
}

// System_layout Named source named complex type "system-layout"
// A system is a group of staves that are read and played simultaneously.
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
type System_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "system-margins" of type system-margins order 745 depth 1
	System_margins *System_margins `xml:"system-margins,omitempty"`

	// generated from element "system-distance" of type tenths order 746 depth 1
	System_distance string `xml:"system-distance,omitempty"`

	// generated from element "top-system-distance" of type tenths order 747 depth 1
	Top_system_distance string `xml:"top-system-distance,omitempty"`

	// generated from element "system-dividers" of type system-dividers order 748 depth 1
	System_dividers *System_dividers `xml:"system-dividers,omitempty"`
}

// System_margins Named source named complex type "system-margins"
// System margins are relative to the page margins. Positive values
// indent and negative values reduce the margin size.
type System_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 750 depth 1
	Group_left_right_margins
}

// Bookmark Named source named complex type "bookmark"
// The bookmark type serves as a well-defined target for an incoming
// simple XLink.
type Bookmark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "element-position
	AttributeGroup_element_position
}

// Link Named source named complex type "link"
// The link type serves as an outgoing simple XLink. If a relative link
// is used within a document that is part of a compressed MusicXML file, the link is
// relative to the root folder of the zip file.
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from attribute group "element-position
	AttributeGroup_element_position

	// generated from attribute group "position
	AttributeGroup_position
}

// Accidental Named source named complex type "accidental"
// The accidental type represents actual notated accidentals. Editorial
// and cautionary indications are indicated by attributes. Values for these attributes
// are "no" if not present. Specific graphic display such as parentheses, brackets, and
// size are controlled by the level-display attribute group.
type Accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "cautionary
	Cautionary string `xml:"cautionary,attr,omitempty"`

	// generated from attribute "editorial
	Editorial Enum_Yes_no `xml:"editorial,attr,omitempty"`

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "level-display
	AttributeGroup_level_display

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_mark Named source named complex type "accidental-mark"
// An accidental-mark can be used as a separate notation or as part of an
// ornament. When used in an ornament, position and placement are relative to the
// ornament, not relative to the note.
type Accidental_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "level-display
	AttributeGroup_level_display

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Accidental_value `xml:",chardata"`
}

// Arpeggiate Named source named complex type "arpeggiate"
// The arpeggiate type indicates that this note is part of an arpeggiated
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

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "direction
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "unbroken
	Unbroken Enum_Yes_no `xml:"unbroken,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Articulations Named source named complex type "articulations"
// Articulations and accents are grouped together here.
type Articulations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accent" of type empty-placement order 771 depth 1
	Accent []*Empty_placement `xml:"accent,omitempty"`

	// generated from element "strong-accent" of type strong-accent order 772 depth 1
	Strong_accent []*Strong_accent `xml:"strong-accent,omitempty"`

	// generated from element "staccato" of type empty-placement order 773 depth 1
	Staccato []*Empty_placement `xml:"staccato,omitempty"`

	// generated from element "tenuto" of type empty-placement order 774 depth 1
	Tenuto []*Empty_placement `xml:"tenuto,omitempty"`

	// generated from element "detached-legato" of type empty-placement order 775 depth 1
	Detached_legato []*Empty_placement `xml:"detached-legato,omitempty"`

	// generated from element "staccatissimo" of type empty-placement order 776 depth 1
	Staccatissimo []*Empty_placement `xml:"staccatissimo,omitempty"`

	// generated from element "spiccato" of type empty-placement order 777 depth 1
	Spiccato []*Empty_placement `xml:"spiccato,omitempty"`

	// generated from element "scoop" of type empty-line order 778 depth 1
	Scoop []*Empty_line `xml:"scoop,omitempty"`

	// generated from element "plop" of type empty-line order 779 depth 1
	Plop []*Empty_line `xml:"plop,omitempty"`

	// generated from element "doit" of type empty-line order 780 depth 1
	Doit []*Empty_line `xml:"doit,omitempty"`

	// generated from element "falloff" of type empty-line order 781 depth 1
	Falloff []*Empty_line `xml:"falloff,omitempty"`

	// generated from element "breath-mark" of type breath-mark order 782 depth 1
	Breath_mark []*Breath_mark `xml:"breath-mark,omitempty"`

	// generated from element "caesura" of type caesura order 783 depth 1
	Caesura []*Caesura `xml:"caesura,omitempty"`

	// generated from element "stress" of type empty-placement order 784 depth 1
	Stress []*Empty_placement `xml:"stress,omitempty"`

	// generated from element "unstress" of type empty-placement order 785 depth 1
	Unstress []*Empty_placement `xml:"unstress,omitempty"`

	// generated from element "soft-accent" of type empty-placement order 786 depth 1
	Soft_accent []*Empty_placement `xml:"soft-accent,omitempty"`

	// generated from element "other-articulation" of type other-placement-text order 787 depth 1
	Other_articulation []*Other_placement_text `xml:"other-articulation,omitempty"`
}

// Arrow Named source named complex type "arrow"
// The arrow element represents an arrow used for a musical technical
// indication. It can represent both Unicode and SMuFL arrows. The presence of an
// arrowhead element indicates that only the arrowhead is displayed, not the arrow
// stem. The smufl attribute distinguishes different SMuFL glyphs that have an arrow
// appearance such as arrowBlackUp, guitarStrumUp, or handbellsSwingUp. The specified
// glyph should match the descriptive representation.
type Arrow struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// generated from element "arrow-direction" of type arrow-direction order 790 depth 1
	Arrow_direction string `xml:"arrow-direction,omitempty"`

	// generated from element "arrow-style" of type arrow-style order 791 depth 1
	Arrow_style string `xml:"arrow-style,omitempty"`

	// generated from element "arrowhead" of type empty order 792 depth 1
	Arrowhead string `xml:"arrowhead,omitempty"`

	// generated from element "circular-arrow" of type circular-arrow order 793 depth 1
	Circular_arrow string `xml:"circular-arrow,omitempty"`
}

// Assess Named source named complex type "assess"
// By default, an assessment application should assess all notes without
// a cue child element, and not assess any note with a cue child element. The assess
// type allows this default assessment to be overridden for individual notes. The
// optional player and time-only attributes restrict the type to apply to a single
// player or set of times through a repeated section, respectively. If missing, the
// type applies to all players or all times through the repeated section, respectively.
// The player attribute references the id attribute of a player element defined within
// the matching score-part.
type Assess struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Yes_no `xml:"type,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Backup Named source named complex type "backup"
// The backup and forward elements are required to coordinate multiple
// voices in one part, including music on multiple staves. The backup type is generally
// used to move between voices and staves. Thus the backup element does not include
// voice or staff elements. Duration values should always be positive, and should not
// cross measure boundaries or mid-measure changes in the divisions value.
type Backup struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 799 depth 1
	Group_duration

	// generated from group with order 800 depth 1
	Group_editorial
}

// Beam Named source named complex type "beam"
// Beam values include begin, continue, end, forward hook, and backward
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

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "repeater
	Repeater Enum_Yes_no `xml:"repeater,attr,omitempty"`

	// generated from attribute "fan
	Fan string `xml:"fan,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Bend Named source named complex type "bend"
// The bend type is used in guitar notation and tablature. A single note
// with a bend and release will contain two bend elements: the first to represent the
// bend and the second to represent the release. The shape attribute distinguishes
// between the angled bend symbols commonly used in standard notation and the curved
// bend symbols commonly used in both tablature and standard notation.
type Bend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "shape
	Shape string `xml:"shape,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "bend-sound
	AttributeGroup_bend_sound

	// generated from element "bend-alter" of type semitones order 805 depth 1
	Bend_alter string `xml:"bend-alter,omitempty"`

	// generated from element "pre-bend" of type empty order 806 depth 1
	Pre_bend string `xml:"pre-bend,omitempty"`

	// generated from element "release" of type release order 807 depth 1
	Release *Release `xml:"release,omitempty"`

	// generated from element "with-bar" of type placement-text order 808 depth 1
	With_bar *Placement_text `xml:"with-bar,omitempty"`
}

// Breath_mark Named source named complex type "breath-mark"
// The breath-mark element indicates a place to take a breath.
type Breath_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Caesura Named source named complex type "caesura"
// The caesura element indicates a slight pause. It is notated using a
// "railroad tracks" symbol or other variations specified in the element content.
type Caesura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Elision Named source named complex type "elision"
// The elision type represents an elision between lyric syllables. The
// text content specifies the symbol used to display the elision. Common values are a
// no-break space (Unicode 00A0), an underscore (Unicode 005F), or an undertie (Unicode
// 203F). If the text content is empty, the smufl attribute is used to specify the
// symbol to use. Its value is a SMuFL canonical glyph name that starts with lyrics.
// The SMuFL attribute is ignored if the elision glyph is already specified by the text
// content. If neither text content nor a smufl attribute are present, the elision
// glyph is application-specific.
type Elision struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Empty_line Named source named complex type "empty-line"
// The empty-line type represents an empty element with line-shape,
// line-type, line-length, dashed-formatting, print-style and placement attributes.
type Empty_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "line-length
	AttributeGroup_line_length

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement
}

// Extend Named source named complex type "extend"
// The extend type represents lyric word extension / melisma lines as
// well as figured bass extensions. The optional type and position attributes are added
// in Version 3.0 to provide better formatting control.
type Extend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop_continue `xml:"type,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color
}

// Figure Named source named complex type "figure"
// The figure type represents a single figure within a figured-bass
// element.
type Figure struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "prefix" of type style-text order 831 depth 1
	Prefix *Style_text `xml:"prefix,omitempty"`

	// generated from element "figure-number" of type style-text order 832 depth 1
	Figure_number *Style_text `xml:"figure-number,omitempty"`

	// generated from element "suffix" of type style-text order 833 depth 1
	Suffix *Style_text `xml:"suffix,omitempty"`

	// generated from element "extend" of type extend order 834 depth 1
	Extend *Extend `xml:"extend,omitempty"`

	// generated from group with order 835 depth 1
	Group_editorial
}

// Figured_bass Named source named complex type "figured-bass"
// The figured-bass element represents figured bass notation. Figured
// bass elements take their position from the first regular note (not a grace note or
// chord note) that follows in score order. The optional duration element is used to
// indicate changes of figures under a note. Figures are ordered from top to bottom.
// The value of parentheses is "no" if not present.
type Figured_bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses Enum_Yes_no `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "figure" of type figure order 837 depth 1
	Figure []*Figure `xml:"figure,omitempty"`

	// generated from group with order 838 depth 1
	Group_duration

	// generated from group with order 839 depth 1
	Group_editorial
}

// Forward Named source named complex type "forward"
// The backup and forward elements are required to coordinate multiple
// voices in one part, including music on multiple staves. The forward element is
// generally used within voices and staves. Duration values should always be positive,
// and should not cross measure boundaries or mid-measure changes in the divisions
// value.
type Forward struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 845 depth 1
	Group_duration

	// generated from group with order 846 depth 1
	Group_editorial_voice

	// generated from group with order 847 depth 1
	Group_staff
}

// Glissando Named source named complex type "glissando"
// Glissando and slide types both indicate rapidly moving from one pitch
// to the other so that individual notes are not discerned. A glissando sounds the
// distinct notes in between the two pitches and defaults to a wavy line. The optional
// text is printed alongside the line.
type Glissando struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grace Named source named complex type "grace"
// The grace type indicates the presence of a grace note. The slash
// attribute for a grace note is yes for slashed grace notes. The steal-time-previous
// attribute indicates the percentage of time to steal from the previous note for the
// grace note. The steal-time-following attribute indicates the percentage of time to
// steal from the following note for the grace note, as for appoggiaturas. The
// make-time attribute indicates to make time, not steal time; the units are in
// real-time divisions for the grace note.
type Grace struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "steal-time-previous
	Steal_time_previous string `xml:"steal-time-previous,attr,omitempty"`

	// generated from attribute "steal-time-following
	Steal_time_following string `xml:"steal-time-following,attr,omitempty"`

	// generated from attribute "make-time
	Make_time string `xml:"make-time,attr,omitempty"`

	// generated from attribute "slash
	Slash Enum_Yes_no `xml:"slash,attr,omitempty"`
}

// Hammer_on_pull_off Named source named complex type "hammer-on-pull-off"
// The hammer-on and pull-off elements are used in guitar and fretted
// instrument notation. Since a single slur can be marked over many notes, the
// hammer-on and pull-off elements are separate so the individual pair of notes can be
// specified. The element content can be used to specify how the hammer-on or pull-off
// should be notated. An empty element leaves this choice up to the application.
type Hammer_on_pull_off struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Handbell Named source named complex type "handbell"
// The handbell element represents notation for various techniques used
// in handbell and handchime music.
type Handbell struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Harmon_closed Named source named complex type "harmon-closed"
// The harmon-closed type represents whether the harmon mute is closed,
// open, or half-open. The optional location attribute indicates which portion of the
// symbol is filled in when the element value is half.
type Harmon_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Harmon_mute Named source named complex type "harmon-mute"
// The harmon-mute type represents the symbols used for harmon mutes in
// brass notation.
type Harmon_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "harmon-closed" of type harmon-closed order 862 depth 1
	Harmon_closed *Harmon_closed `xml:"harmon-closed,omitempty"`
}

// Harmonic Named source named complex type "harmonic"
// The harmonic type indicates natural and artificial harmonics. Allowing
// the type of pitch to be specified, combined with controls for appearance/playback
// differences, allows both the notation and the sound to be represented. Artificial
// harmonics can add a notated touching pitch; artificial pinch harmonics will usually
// not notate a touching pitch. The attributes for the harmonic element refer to the
// use of the circular harmonic symbol, typically but not always used with natural
// harmonics.
type Harmonic struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "natural" of type empty order 866 depth 1
	Natural string `xml:"natural,omitempty"`

	// generated from element "artificial" of type empty order 867 depth 1
	Artificial string `xml:"artificial,omitempty"`

	// generated from element "base-pitch" of type empty order 868 depth 1
	Base_pitch string `xml:"base-pitch,omitempty"`

	// generated from element "touching-pitch" of type empty order 869 depth 1
	Touching_pitch string `xml:"touching-pitch,omitempty"`

	// generated from element "sounding-pitch" of type empty order 870 depth 1
	Sounding_pitch string `xml:"sounding-pitch,omitempty"`
}

// Heel_toe Named source named complex type "heel-toe"
// The heel and toe elements are used with organ pedals. The substitution
// value is "no" if the attribute is not present.
type Heel_toe struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Hole Named source named complex type "hole"
// The hole type represents the symbols used for woodwind and brass
// fingerings as well as other notations.
type Hole struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "hole-type" of type string order 876 depth 1
	Hole_type string `xml:"hole-type,omitempty"`

	// generated from element "hole-closed" of type hole-closed order 877 depth 1
	Hole_closed *Hole_closed `xml:"hole-closed,omitempty"`

	// generated from element "hole-shape" of type string order 878 depth 1
	Hole_shape string `xml:"hole-shape,omitempty"`
}

// Hole_closed Named source named complex type "hole-closed"
// The hole-closed type represents whether the hole is closed, open, or
// half-open. The optional location attribute indicates which portion of the hole is
// filled in when the element value is half.
type Hole_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Instrument Named source named complex type "instrument"
// The instrument type distinguishes between score-instrument elements in
// a score-part. The id attribute is an IDREF back to the score-instrument ID. If
// multiple score-instruments are specified in a score-part, there should be an
// instrument element for each note in the part. Notes that are shared between multiple
// score-instruments can have more than one instrument element.
type Instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// Listen Named source named complex type "listen"
// The listen and listening types, new in Version 4.0, specify different
// ways that a score following or machine listening application can interact with a
// performer. The listen type handles interactions that are specific to a note. If
// multiple child elements of the same type are present, they should have distinct
// player and/or time-only attributes.
type Listen struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "assess" of type assess order 884 depth 1
	Assess []*Assess `xml:"assess,omitempty"`

	// generated from element "wait" of type wait order 885 depth 1
	Wait []*Wait `xml:"wait,omitempty"`

	// generated from element "other-listen" of type other-listening order 886 depth 1
	Other_listen []*Other_listening `xml:"other-listen,omitempty"`
}

// Lyric Named source named complex type "lyric"
// The lyric type represents text underlays for lyrics. Two text elements
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

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "elision" of type elision order 890 depth 1
	Elision []*Elision `xml:"elision,omitempty"`

	// generated from element "syllabic" of type syllabic order 891 depth 1
	Syllabic string `xml:"syllabic,omitempty"`

	// generated from element "text" of type text-element-data order 892 depth 1
	Text []*Text_element_data `xml:"text,omitempty"`

	// generated from element "extend" of type extend order 893 depth 1
	Extend *Extend `xml:"extend,omitempty"`

	// generated from element "laughing" of type empty order 895 depth 1
	Laughing string `xml:"laughing,omitempty"`

	// generated from element "humming" of type empty order 896 depth 1
	Humming string `xml:"humming,omitempty"`

	// generated from element "end-line" of type empty order 897 depth 1
	End_line string `xml:"end-line,omitempty"`

	// generated from element "end-paragraph" of type empty order 898 depth 1
	End_paragraph string `xml:"end-paragraph,omitempty"`

	// generated from group with order 899 depth 1
	Group_editorial
}

// Mordent Named source named complex type "mordent"
// The mordent type is used for both represents the mordent sign with the
// vertical line and the inverted-mordent sign without the line. The long attribute is
// "no" by default. The approach and departure attributes are used for compound
// ornaments, indicating how the beginning and ending of the ornament look relative to
// the main part of the mordent.
type Mordent struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Non_arpeggiate Named source named complex type "non-arpeggiate"
// The non-arpeggiate type indicates that this note is at the top or
// bottom of a bracket indicating to not arpeggiate these notes. Since this does not
// involve playback, it is only used on the top or bottom notes, not on each note as
// for the arpeggiate type.
type Non_arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Notations Named source named complex type "notations"
// Notations refer to musical notations, not XML notations. Multiple
// notations are allowed in order to represent multiple editorial levels. The
// print-object attribute, added in Version 3.0, allows notations to represent details
// of performance technique, such as fingerings, without having them appear in the
// score.
type Notations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 913 depth 1
	Group_editorial

	// generated from element "tied" of type tied order 914 depth 1
	Tied []*Tied `xml:"tied,omitempty"`

	// generated from element "slur" of type slur order 915 depth 1
	Slur []*Slur `xml:"slur,omitempty"`

	// generated from element "tuplet" of type tuplet order 916 depth 1
	Tuplet []*Tuplet `xml:"tuplet,omitempty"`

	// generated from element "glissando" of type glissando order 917 depth 1
	Glissando []*Glissando `xml:"glissando,omitempty"`

	// generated from element "slide" of type slide order 918 depth 1
	Slide []*Slide `xml:"slide,omitempty"`

	// generated from element "ornaments" of type ornaments order 919 depth 1
	Ornaments []*Ornaments `xml:"ornaments,omitempty"`

	// generated from element "technical" of type technical order 920 depth 1
	Technical []*Technical `xml:"technical,omitempty"`

	// generated from element "articulations" of type articulations order 921 depth 1
	Articulations []*Articulations `xml:"articulations,omitempty"`

	// generated from element "dynamics" of type dynamics order 922 depth 1
	Dynamics []*Dynamics `xml:"dynamics,omitempty"`

	// generated from element "fermata" of type fermata order 923 depth 1
	Fermata []*Fermata `xml:"fermata,omitempty"`

	// generated from element "arpeggiate" of type arpeggiate order 924 depth 1
	Arpeggiate []*Arpeggiate `xml:"arpeggiate,omitempty"`

	// generated from element "non-arpeggiate" of type non-arpeggiate order 925 depth 1
	Non_arpeggiate []*Non_arpeggiate `xml:"non-arpeggiate,omitempty"`

	// generated from element "accidental-mark" of type accidental-mark order 926 depth 1
	Accidental_mark []*Accidental_mark `xml:"accidental-mark,omitempty"`

	// generated from element "other-notation" of type other-notation order 927 depth 1
	Other_notation []*Other_notation `xml:"other-notation,omitempty"`
}

// Note Named source named complex type "note"
// Notes are the most common type of MusicXML data. The MusicXML format
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

	// insertion point for fields

	// generated from attribute "print-leger
	Print_leger Enum_Yes_no `xml:"print-leger,attr,omitempty"`

	// generated from attribute "dynamics
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "end-dynamics
	End_dynamics string `xml:"end-dynamics,attr,omitempty"`

	// generated from attribute "attack
	Attack string `xml:"attack,attr,omitempty"`

	// generated from attribute "release
	Release string `xml:"release,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato
	Pizzicato Enum_Yes_no `xml:"pizzicato,attr,omitempty"`

	// generated from attribute group "x-position
	AttributeGroup_x_position

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "grace" of type grace order 931 depth 1
	Grace *Grace `xml:"grace,omitempty"`

	// generated from group with order 932 depth 1
	Group_full_note

	// generated from element "tie" of type tie order 933 depth 1
	Tie *Tie `xml:"tie,omitempty"`

	// generated from element "cue" of type empty order 934 depth 1
	Cue string `xml:"cue,omitempty"`

	// generated from group with order 938 depth 1
	Group_duration

	// generated from element "instrument" of type instrument order 942 depth 1
	Instrument []*Instrument `xml:"instrument,omitempty"`

	// generated from group with order 943 depth 1
	Group_editorial_voice

	// generated from element "type" of type note-type order 944 depth 1
	Type *Note_type `xml:"type,omitempty"`

	// generated from element "dot" of type empty-placement order 945 depth 1
	Dot []*Empty_placement `xml:"dot,omitempty"`

	// generated from element "accidental" of type accidental order 946 depth 1
	Accidental *Accidental `xml:"accidental,omitempty"`

	// generated from element "time-modification" of type time-modification order 947 depth 1
	Time_modification *Time_modification `xml:"time-modification,omitempty"`

	// generated from element "stem" of type stem order 948 depth 1
	Stem *Stem `xml:"stem,omitempty"`

	// generated from element "notehead" of type notehead order 949 depth 1
	Notehead *Notehead `xml:"notehead,omitempty"`

	// generated from element "notehead-text" of type notehead-text order 950 depth 1
	Notehead_text *Notehead_text `xml:"notehead-text,omitempty"`

	// generated from group with order 951 depth 1
	Group_staff

	// generated from element "beam" of type beam order 952 depth 1
	Beam *Beam `xml:"beam,omitempty"`

	// generated from element "notations" of type notations order 953 depth 1
	Notations []*Notations `xml:"notations,omitempty"`

	// generated from element "lyric" of type lyric order 954 depth 1
	Lyric []*Lyric `xml:"lyric,omitempty"`

	// generated from element "play" of type play order 955 depth 1
	Play *Play `xml:"play,omitempty"`

	// generated from element "listen" of type listen order 956 depth 1
	Listen *Listen `xml:"listen,omitempty"`
}

// Note_type Named source named complex type "note-type"
// The note-type type indicates the graphic note type. Values range from
// 1024th to maxima. The size attribute indicates full, cue, grace-cue, or large size.
// The default is full for regular notes, grace-cue for notes that contain both grace
// and cue elements, and cue for notes that contain either a cue or a grace element,
// but not both.
type Note_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "size
	Size Enum_Symbol_size `xml:"size,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Note_type_value `xml:",chardata"`
}

// Notehead Named source named complex type "notehead"
// The notehead type indicates shapes other than the open and closed
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

	// insertion point for fields

	// generated from attribute "filled
	Filled Enum_Yes_no `xml:"filled,attr,omitempty"`

	// generated from attribute "parentheses
	Parentheses Enum_Yes_no `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead_text Named source named complex type "notehead-text"
// The notehead-text type represents text that is displayed inside a
// notehead, as is done in some educational music. It is not needed for the numbers
// used in tablature or jianpu notation. The presence of a TAB or jianpu clefs is
// sufficient to indicate that numbers are used. The display-text and accidental-text
// elements allow display of fully formatted text and accidentals.
type Notehead_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-text" of type formatted-text order 968 depth 1
	Display_text []*Formatted_text `xml:"display-text,omitempty"`

	// generated from element "accidental-text" of type accidental-text order 969 depth 1
	Accidental_text []*Accidental_text `xml:"accidental-text,omitempty"`
}

// Ornaments Named source named complex type "ornaments"
// Ornaments can be any of several types, followed optionally by
// accidentals. The accidental-mark element's content is represented the same as an
// accidental element, but with a different name to reflect the different musical
// meaning.
type Ornaments struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "trill-mark" of type empty-trill-sound order 971 depth 1
	Trill_mark []*Empty_trill_sound `xml:"trill-mark,omitempty"`

	// generated from element "turn" of type horizontal-turn order 972 depth 1
	Turn []*Horizontal_turn `xml:"turn,omitempty"`

	// generated from element "delayed-turn" of type horizontal-turn order 973 depth 1
	Delayed_turn []*Horizontal_turn `xml:"delayed-turn,omitempty"`

	// generated from element "inverted-turn" of type horizontal-turn order 974 depth 1
	Inverted_turn []*Horizontal_turn `xml:"inverted-turn,omitempty"`

	// generated from element "delayed-inverted-turn" of type horizontal-turn order 975 depth 1
	Delayed_inverted_turn []*Horizontal_turn `xml:"delayed-inverted-turn,omitempty"`

	// generated from element "vertical-turn" of type empty-trill-sound order 976 depth 1
	Vertical_turn []*Empty_trill_sound `xml:"vertical-turn,omitempty"`

	// generated from element "inverted-vertical-turn" of type empty-trill-sound order 977 depth 1
	Inverted_vertical_turn []*Empty_trill_sound `xml:"inverted-vertical-turn,omitempty"`

	// generated from element "shake" of type empty-trill-sound order 978 depth 1
	Shake []*Empty_trill_sound `xml:"shake,omitempty"`

	// generated from element "wavy-line" of type wavy-line order 979 depth 1
	Wavy_line []*Wavy_line `xml:"wavy-line,omitempty"`

	// generated from element "mordent" of type mordent order 980 depth 1
	Mordent []*Mordent `xml:"mordent,omitempty"`

	// generated from element "inverted-mordent" of type mordent order 981 depth 1
	Inverted_mordent []*Mordent `xml:"inverted-mordent,omitempty"`

	// generated from element "schleifer" of type empty-placement order 982 depth 1
	Schleifer []*Empty_placement `xml:"schleifer,omitempty"`

	// generated from element "tremolo" of type tremolo order 983 depth 1
	Tremolo []*Tremolo `xml:"tremolo,omitempty"`

	// generated from element "haydn" of type empty-trill-sound order 984 depth 1
	Haydn []*Empty_trill_sound `xml:"haydn,omitempty"`

	// generated from element "other-ornament" of type other-placement-text order 985 depth 1
	Other_ornament []*Other_placement_text `xml:"other-ornament,omitempty"`

	// generated from element "accidental-mark" of type accidental-mark order 986 depth 1
	Accidental_mark []*Accidental_mark `xml:"accidental-mark,omitempty"`
}

// Other_notation Named source named complex type "other-notation"
// The other-notation type is used to define any notations not yet in the
// MusicXML format. It handles notations where more specific extension elements such as
// other-dynamics and other-technical are not appropriate. The smufl attribute can be
// used to specify a particular notation, allowing application interoperability without
// requiring every SMuFL glyph to have a MusicXML element equivalent. Using the
// other-notation type without the smufl attribute allows for extended representation,
// though without application interoperability.
type Other_notation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop_single `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_placement_text Named source named complex type "other-placement-text"
// The other-placement-text type represents a text element with
// print-style, placement, and smufl attribute groups. This type is used by MusicXML
// notation extension elements to allow specification of specific SMuFL glyphs without
// needed to add every glyph as a MusicXML element.
type Other_placement_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_text Named source named complex type "other-text"
// The other-text type represents a text element with a smufl attribute
// group. This type is used by MusicXML direction extension elements to allow
// specification of specific SMuFL glyphs without needed to add every glyph as a
// MusicXML element.
type Other_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Pitch Named source named complex type "pitch"
// Pitch is represented as a combination of the step of the diatonic
// scale, the chromatic alteration, and the octave.
type Pitch struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "step" of type step order 1001 depth 1
	Step Enum_Step `xml:"step,omitempty"`

	// generated from element "alter" of type semitones order 1002 depth 1
	Alter string `xml:"alter,omitempty"`

	// generated from element "octave" of type octave order 1003 depth 1
	Octave int `xml:"octave,omitempty"`
}

// Placement_text Named source named complex type "placement-text"
// The placement-text type represents a text element with print-style and
// placement attribute groups.
type Placement_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Release Named source named complex type "release"
// The release type indicates that a bend is a release rather than a
// normal bend or pre-bend. The offset attribute specifies where the release starts in
// terms of divisions relative to the current note. The first-beat and last-beat
// attributes of the parent bend element are relative to the original note position,
// not this offset value.
type Release struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Rest Named source named complex type "rest"
// The rest element indicates notated rests or silences. Rest elements
// are usually empty, but placement on the staff can be specified using display-step
// and display-octave elements. If the measure attribute is set to yes, this indicates
// this is a complete measure rest.
type Rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "measure
	Measure Enum_Yes_no `xml:"measure,attr,omitempty"`

	// generated from group with order 1009 depth 1
	Group_display_step_octave
}

// Slide Named source named complex type "slide"
// Glissando and slide types both indicate rapidly moving from one pitch
// to the other so that individual notes are not discerned. A slide is continuous
// between the two pitches and defaults to a solid line. The optional text for a is
// printed alongside the line.
type Slide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "bend-sound
	AttributeGroup_bend_sound

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Slur Named source named complex type "slur"
// Slur types are empty. Most slurs are represented with two elements:
// one with a start type, and one with a stop type. Slurs can add more elements using a
// continue type. This is typically used to specify the formatting of cross-system
// slurs, or to specify the shape of very complex slurs.
type Slur struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop_continue `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "orientation
	AttributeGroup_orientation

	// generated from attribute group "bezier
	AttributeGroup_bezier

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Stem Named source named complex type "stem"
// Stems can be down, up, none, or double. For down and up stems, the
// position attributes can be used to specify stem length. The relative values specify
// the end of the stem relative to the program default. Default values specify an
// absolute end stem position. Negative values of relative-y that would flip a stem
// instead of shortening it are ignored. A stem element associated with a rest refers
// to a stemlet.
type Stem struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "y-position
	AttributeGroup_y_position

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Strong_accent Named source named complex type "strong-accent"
// The strong-accent type indicates a vertical accent mark. The type
// attribute indicates if the point of the accent is down or up.
type Strong_accent struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Style_text Named source named complex type "style-text"
// The style-text type represents a text element with a print-style
// attribute group.
type Style_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Tap Named source named complex type "tap"
// The tap type indicates a tap on the fretboard. The text content allows
// specification of the notation; + and T are common choices. If the element is empty,
// the hand attribute is used to specify the symbol to use. The hand attribute is
// ignored if the tap glyph is already specified by the text content. If neither text
// content nor the hand attribute are present, the display is application-specific.
type Tap struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "hand
	Hand string `xml:"hand,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Technical Named source named complex type "technical"
// Technical indications give performance information for individual
// instruments.
type Technical struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "up-bow" of type empty-placement order 1035 depth 1
	Up_bow []*Empty_placement `xml:"up-bow,omitempty"`

	// generated from element "down-bow" of type empty-placement order 1036 depth 1
	Down_bow []*Empty_placement `xml:"down-bow,omitempty"`

	// generated from element "harmonic" of type harmonic order 1037 depth 1
	Harmonic []*Harmonic `xml:"harmonic,omitempty"`

	// generated from element "open-string" of type empty-placement order 1038 depth 1
	Open_string []*Empty_placement `xml:"open-string,omitempty"`

	// generated from element "thumb-position" of type empty-placement order 1039 depth 1
	Thumb_position []*Empty_placement `xml:"thumb-position,omitempty"`

	// generated from element "fingering" of type fingering order 1040 depth 1
	Fingering []*Fingering `xml:"fingering,omitempty"`

	// generated from element "pluck" of type placement-text order 1041 depth 1
	Pluck []*Placement_text `xml:"pluck,omitempty"`

	// generated from element "double-tongue" of type empty-placement order 1042 depth 1
	Double_tongue []*Empty_placement `xml:"double-tongue,omitempty"`

	// generated from element "triple-tongue" of type empty-placement order 1043 depth 1
	Triple_tongue []*Empty_placement `xml:"triple-tongue,omitempty"`

	// generated from element "stopped" of type empty-placement-smufl order 1044 depth 1
	Stopped []*Empty_placement_smufl `xml:"stopped,omitempty"`

	// generated from element "snap-pizzicato" of type empty-placement order 1045 depth 1
	Snap_pizzicato []*Empty_placement `xml:"snap-pizzicato,omitempty"`

	// generated from element "fret" of type fret order 1046 depth 1
	Fret []*Fret `xml:"fret,omitempty"`

	// generated from element "string" of type string-type order 1047 depth 1
	String []*String_type `xml:"string,omitempty"`

	// generated from element "hammer-on" of type hammer-on-pull-off order 1048 depth 1
	Hammer_on []*Hammer_on_pull_off `xml:"hammer-on,omitempty"`

	// generated from element "pull-off" of type hammer-on-pull-off order 1049 depth 1
	Pull_off []*Hammer_on_pull_off `xml:"pull-off,omitempty"`

	// generated from element "bend" of type bend order 1050 depth 1
	Bend []*Bend `xml:"bend,omitempty"`

	// generated from element "tap" of type tap order 1051 depth 1
	Tap []*Tap `xml:"tap,omitempty"`

	// generated from element "heel" of type heel-toe order 1052 depth 1
	Heel []*Heel_toe `xml:"heel,omitempty"`

	// generated from element "toe" of type heel-toe order 1053 depth 1
	Toe []*Heel_toe `xml:"toe,omitempty"`

	// generated from element "fingernails" of type empty-placement order 1054 depth 1
	Fingernails []*Empty_placement `xml:"fingernails,omitempty"`

	// generated from element "hole" of type hole order 1055 depth 1
	Hole []*Hole `xml:"hole,omitempty"`

	// generated from element "arrow" of type arrow order 1056 depth 1
	Arrow []*Arrow `xml:"arrow,omitempty"`

	// generated from element "handbell" of type handbell order 1057 depth 1
	Handbell []*Handbell `xml:"handbell,omitempty"`

	// generated from element "brass-bend" of type empty-placement order 1058 depth 1
	Brass_bend []*Empty_placement `xml:"brass-bend,omitempty"`

	// generated from element "flip" of type empty-placement order 1059 depth 1
	Flip []*Empty_placement `xml:"flip,omitempty"`

	// generated from element "smear" of type empty-placement order 1060 depth 1
	Smear []*Empty_placement `xml:"smear,omitempty"`

	// generated from element "open" of type empty-placement-smufl order 1061 depth 1
	Open []*Empty_placement_smufl `xml:"open,omitempty"`

	// generated from element "half-muted" of type empty-placement-smufl order 1062 depth 1
	Half_muted []*Empty_placement_smufl `xml:"half-muted,omitempty"`

	// generated from element "harmon-mute" of type harmon-mute order 1063 depth 1
	Harmon_mute []*Harmon_mute `xml:"harmon-mute,omitempty"`

	// generated from element "golpe" of type empty-placement order 1064 depth 1
	Golpe []*Empty_placement `xml:"golpe,omitempty"`

	// generated from element "other-technical" of type other-placement-text order 1065 depth 1
	Other_technical []*Other_placement_text `xml:"other-technical,omitempty"`
}

// Text_element_data Named source named complex type "text-element-data"
// The text-element-data type represents a syllable or portion of a
// syllable for lyric text underlay. A hyphen in the string content should only be used
// for an actual hyphenated word. Language names for text elements come from ISO 639,
// with optional country subcodes from ISO 3166.
type Text_element_data struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "text-rotation
	AttributeGroup_text_rotation

	// generated from attribute group "letter-spacing
	AttributeGroup_letter_spacing

	// generated from attribute group "text-direction
	AttributeGroup_text_direction

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Tie Named source named complex type "tie"
// The tie element indicates that a tie begins or ends with this note. If
// the tie element applies only particular times through a repeat, the time-only
// attribute indicates which times to apply it. The tie element indicates sound; the
// tied element indicates notation.
type Tie struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`
}

// Tied Named source named complex type "tied"
// The tied element represents the notated tie. The tie element
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

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "orientation
	AttributeGroup_orientation

	// generated from attribute group "bezier
	AttributeGroup_bezier

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Time_modification Named source named complex type "time-modification"
// Time modification indicates tuplets, double-note tremolos, and other
// durational changes. A time-modification element shows how the cumulative, sounding
// effect of tuplets and double-note tremolos compare to the written note type
// represented by the type and dot elements. Nested tuplets and other notations that
// use more detailed information need both the time-modification and tuplet elements to
// be represented accurately.
type Time_modification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "actual-notes" of type nonNegativeInteger order 1085 depth 1
	Actual_notes int `xml:"actual-notes,omitempty"`

	// generated from element "normal-notes" of type nonNegativeInteger order 1086 depth 1
	Normal_notes int `xml:"normal-notes,omitempty"`

	// generated from element "normal-type" of type note-type-value order 1087 depth 1
	Normal_type Enum_Note_type_value `xml:"normal-type,omitempty"`

	// generated from element "normal-dot" of type empty order 1088 depth 1
	Normal_dot string `xml:"normal-dot,omitempty"`
}

// Tremolo Named source named complex type "tremolo"
// The tremolo ornament can be used to indicate single-note, double-note,
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

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet Named source named complex type "tuplet"
// A tuplet element is present when a tuplet is to be displayed
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

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "bracket
	Bracket Enum_Yes_no `xml:"bracket,attr,omitempty"`

	// generated from attribute "show-number
	Show_number string `xml:"show-number,attr,omitempty"`

	// generated from attribute "show-type
	Show_type Enum_Show_tuplet `xml:"show-type,attr,omitempty"`

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "tuplet-actual" of type tuplet-portion order 1094 depth 1
	Tuplet_actual *Tuplet_portion `xml:"tuplet-actual,omitempty"`

	// generated from element "tuplet-normal" of type tuplet-portion order 1095 depth 1
	Tuplet_normal *Tuplet_portion `xml:"tuplet-normal,omitempty"`
}

// Tuplet_dot Named source named complex type "tuplet-dot"
// The tuplet-dot type is used to specify dotted tuplet types.
type Tuplet_dot struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color
}

// Tuplet_number Named source named complex type "tuplet-number"
// The tuplet-number type indicates the number of notes for this portion
// of the tuplet.
type Tuplet_number struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet_portion Named source named complex type "tuplet-portion"
// The tuplet-portion type provides optional full control over tuplet
// specifications. It allows the number and note type (including dots) to be set for
// the actual and normal portions of a single tuplet. If any of these elements are
// absent, their values are based on the time-modification element.
type Tuplet_portion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "tuplet-number" of type tuplet-number order 1107 depth 1
	Tuplet_number *Tuplet_number `xml:"tuplet-number,omitempty"`

	// generated from element "tuplet-type" of type tuplet-type order 1108 depth 1
	Tuplet_type *Tuplet_type `xml:"tuplet-type,omitempty"`

	// generated from element "tuplet-dot" of type tuplet-dot order 1109 depth 1
	Tuplet_dot []*Tuplet_dot `xml:"tuplet-dot,omitempty"`
}

// Tuplet_type Named source named complex type "tuplet-type"
// The tuplet-type type indicates the graphical note type of the notes
// for this portion of the tuplet.
type Tuplet_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText Enum_Note_type_value `xml:",chardata"`
}

// Unpitched Named source named complex type "unpitched"
// The unpitched type represents musical elements that are notated on the
// staff but lack definite pitch, such as unpitched percussion and speaking voice. If
// the child elements are not present, the note is placed on the middle line of the
// staff. This is generally used with a one-line staff. Notes in percussion clef should
// always use an unpitched element rather than a pitch element.
type Unpitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 1114 depth 1
	Group_display_step_octave
}

// Wait Named source named complex type "wait"
// The wait type specifies a point where the accompaniment should wait
// for a performer event before continuing. This typically happens at the start of new
// sections or after a held note or indeterminate music. These waiting points cannot
// always be inferred reliably from the contents of the displayed score. The optional
// player and time-only attributes restrict the type to apply to a single player or set
// of times through a repeated section, respectively.
type Wait struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only Enum_Time_only `xml:"time-only,attr,omitempty"`
}

// Credit Named source named complex type "credit"
// The credit type represents the appearance of the title, composer,
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

	// insertion point for fields

	// generated from attribute "page
	Page int `xml:"page,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "credit-type" of type string order 1117 depth 1
	Credit_type string `xml:"credit-type,omitempty"`

	// generated from element "credit-image" of type image order 1120 depth 1
	Credit_image *Image `xml:"credit-image,omitempty"`

	// generated from element "link" of type link order 1123 depth 1
	Link []*Link `xml:"link,omitempty"`

	// generated from element "bookmark" of type bookmark order 1124 depth 1
	Bookmark []*Bookmark `xml:"bookmark,omitempty"`

	// generated from element "credit-words" of type formatted-text-id order 1125 depth 1
	Credit_words []*Formatted_text_id `xml:"credit-words,omitempty"`

	// generated from element "credit-symbol" of type formatted-symbol-id order 1126 depth 1
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol,omitempty"`
}

// Defaults Named source named complex type "defaults"
// The defaults type specifies score-wide defaults for scaling; whether
// or not the file is a concert score; layout; and default values for the music font,
// word font, lyric font, and lyric language. Except for the concert-score element, if
// any defaults are missing, the choice of what to use is determined by the
// application.
type Defaults struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "scaling" of type scaling order 1129 depth 1
	Scaling *Scaling `xml:"scaling,omitempty"`

	// generated from element "concert-score" of type empty order 1130 depth 1
	Concert_score string `xml:"concert-score,omitempty"`

	// generated from group with order 1131 depth 1
	Group_layout

	// generated from element "appearance" of type appearance order 1132 depth 1
	Appearance *Appearance `xml:"appearance,omitempty"`

	// generated from element "music-font" of type empty-font order 1133 depth 1
	Music_font *Empty_font `xml:"music-font,omitempty"`

	// generated from element "word-font" of type empty-font order 1134 depth 1
	Word_font *Empty_font `xml:"word-font,omitempty"`

	// generated from element "lyric-font" of type lyric-font order 1135 depth 1
	Lyric_font []*Lyric_font `xml:"lyric-font,omitempty"`

	// generated from element "lyric-language" of type lyric-language order 1136 depth 1
	Lyric_language []*Lyric_language `xml:"lyric-language,omitempty"`
}

// Empty_font Named source named complex type "empty-font"
// The empty-font type represents an empty element with font attributes.
type Empty_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font
}

// Group_barline Named source named complex type "group-barline"
// The group-barline type indicates if the group should have common
// barlines.
type Group_barline struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_name Named source named complex type "group-name"
// The group-name type describes the name or abbreviation of a part-group
// element. Formatting attributes in the group-name type are deprecated in Version 2.0
// in favor of the new group-name-display and group-abbreviation-display elements.
type Group_name struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "group-name-text
	AttributeGroup_group_name_text

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_symbol Named source named complex type "group-symbol"
// The group-symbol type indicates how the symbol for a group is
// indicated in the score. It is none if not specified.
type Group_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Instrument_link Named source named complex type "instrument-link"
// Multiple part-link elements can link a condensed part within a score
// file to multiple MusicXML parts files. For example, a "Clarinet 1 and 2" part in a
// score file could link to separate "Clarinet 1" and "Clarinet 2" part files. The
// instrument-link type distinguish which of the score-instruments within a score-part
// are in which part file. The instrument-link id attribute refers to a
// score-instrument id attribute.
type Instrument_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// Lyric_font Named source named complex type "lyric-font"
// The lyric-font type specifies the default font for a particular name
// and number of lyric.
type Lyric_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font
}

// Lyric_language Named source named complex type "lyric-language"
// The lyric-language type specifies the default language for a
// particular name and number of lyric.
type Lyric_language struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Opus Named source named complex type "opus"
// The opus type represents a link to a MusicXML opus document that
// composes multiple MusicXML scores into a collection.
type Opus struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes
}

// Part_group Named source named complex type "part-group"
// The part-group element indicates groupings of parts in the score,
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
type Part_group struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type Enum_Start_stop `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from element "group-name" of type group-name order 1153 depth 1
	Group_name *Group_name `xml:"group-name,omitempty"`

	// generated from element "group-name-display" of type name-display order 1154 depth 1
	Group_name_display *Name_display `xml:"group-name-display,omitempty"`

	// generated from element "group-abbreviation" of type group-name order 1155 depth 1
	Group_abbreviation *Group_name `xml:"group-abbreviation,omitempty"`

	// generated from element "group-abbreviation-display" of type name-display order 1156 depth 1
	Group_abbreviation_display *Name_display `xml:"group-abbreviation-display,omitempty"`

	// generated from element "group-symbol" of type group-symbol order 1157 depth 1
	Group_symbol *Group_symbol `xml:"group-symbol,omitempty"`

	// generated from element "group-barline" of type group-barline order 1158 depth 1
	Group_barline *Group_barline `xml:"group-barline,omitempty"`

	// generated from element "group-time" of type empty order 1159 depth 1
	Group_time string `xml:"group-time,omitempty"`

	// generated from group with order 1160 depth 1
	Group_editorial
}

// Part_link Named source named complex type "part-link"
// The part-link type allows MusicXML data for both score and parts to be
// contained within a single compressed MusicXML file. It links a score-part from a
// score document to MusicXML documents that contain parts data. In the case of a
// single compressed MusicXML file, the link href values are paths that are relative to
// the root folder of the zip file.
type Part_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from element "instrument-link" of type instrument-link order 1162 depth 1
	Instrument_link []*Instrument_link `xml:"instrument-link,omitempty"`

	// generated from element "group-link" of type string order 1163 depth 1
	Group_link string `xml:"group-link,omitempty"`
}

// Part_list Named source named complex type "part-list"
// The part-list identifies the different musical parts in this document.
// Each part has an ID that is used later within the musical data. Since parts may be
// encoded separately and combined later, identification elements are present at both
// the score and score-part levels. There must be at least one score-part, combined as
// desired with part-group elements that indicate braces and brackets. Parts are
// ordered from top to bottom in a score based on the order in which they appear in the
// part-list.
type Part_list struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 1166 depth 1
	Group_part_group

	// generated from group with order 1167 depth 1
	Group_score_part
}

// Part_name Named source named complex type "part-name"
// The part-name type describes the name or abbreviation of a score-part
// element. Formatting attributes for the part-name element are deprecated in Version
// 2.0 in favor of the new part-name-display and part-abbreviation-display elements.
type Part_name struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "part-name-text
	AttributeGroup_part_name_text

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Player Named source named complex type "player"
// The player type allows for multiple players per score-part for use in
// listening applications. One player may play multiple instruments, while a single
// instrument may include multiple players in divisi sections.
type Player struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "player-name" of type string order 1173 depth 1
	Player_name string `xml:"player-name,omitempty"`
}

// Score_instrument Named source named complex type "score-instrument"
// The score-instrument type represents a single instrument within a
// score-part. As with the score-part type, each score-instrument has a required ID
// attribute, a name, and an optional abbreviation. A score-instrument type is also
// required if the score specifies MIDI 1.0 channels, banks, or programs. An initial
// midi-instrument assignment can also be made here. MusicXML software should be able
// to automatically assign reasonable channels and instruments without these elements
// in simple cases, such as where part names match General MIDI instrument names. The
// score-instrument element can also distinguish multiple instruments of the same type
// that are on the same part, such as Clarinet 1 and Clarinet 2 instruments within a
// Clarinets 1 and 2 part.
type Score_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "instrument-name" of type string order 1175 depth 1
	Instrument_name string `xml:"instrument-name,omitempty"`

	// generated from element "instrument-abbreviation" of type string order 1176 depth 1
	Instrument_abbreviation string `xml:"instrument-abbreviation,omitempty"`

	// generated from group with order 1177 depth 1
	Group_virtual_instrument_data
}

// Score_part Named source named complex type "score-part"
// The score-part type collects part-wide information for each part in a
// score. Often, each MusicXML part corresponds to a track in a Standard MIDI Format 1
// file. In this case, the midi-device element is used to make a MIDI device or port
// assignment for the given track or specific MIDI instruments. Initial midi-instrument
// assignments may be made here as well. The score-instrument elements are used when
// there are multiple instruments per track.
type Score_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "identification" of type identification order 1179 depth 1
	Identification *Identification `xml:"identification,omitempty"`

	// generated from element "part-link" of type part-link order 1180 depth 1
	Part_link []*Part_link `xml:"part-link,omitempty"`

	// generated from element "part-name" of type part-name order 1181 depth 1
	Part_name *Part_name `xml:"part-name,omitempty"`

	// generated from element "part-name-display" of type name-display order 1182 depth 1
	Part_name_display *Name_display `xml:"part-name-display,omitempty"`

	// generated from element "part-abbreviation" of type part-name order 1183 depth 1
	Part_abbreviation *Part_name `xml:"part-abbreviation,omitempty"`

	// generated from element "part-abbreviation-display" of type name-display order 1184 depth 1
	Part_abbreviation_display *Name_display `xml:"part-abbreviation-display,omitempty"`

	// generated from element "group" of type string order 1185 depth 1
	Group string `xml:"group,omitempty"`

	// generated from element "score-instrument" of type score-instrument order 1186 depth 1
	Score_instrument []*Score_instrument `xml:"score-instrument,omitempty"`

	// generated from element "player" of type player order 1187 depth 1
	Player []*Player `xml:"player,omitempty"`

	// generated from element "midi-device" of type midi-device order 1188 depth 1
	Midi_device []*Midi_device `xml:"midi-device,omitempty"`

	// generated from element "midi-instrument" of type midi-instrument order 1189 depth 1
	Midi_instrument []*Midi_instrument `xml:"midi-instrument,omitempty"`
}

// Virtual_instrument Named source named complex type "virtual-instrument"
// The virtual-instrument element defines a specific virtual instrument
// used for an instrument sound.
type Virtual_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "virtual-library" of type string order 1191 depth 1
	Virtual_library string `xml:"virtual-library,omitempty"`

	// generated from element "virtual-name" of type string order 1192 depth 1
	Virtual_name string `xml:"virtual-name,omitempty"`
}

// Work Named source named complex type "work"
// Works are optionally identified by number and title. The work type
// also may indicate a link to the opus document that composes multiple scores into a
// collection.
type Work struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "work-number" of type string order 1194 depth 1
	Work_number string `xml:"work-number,omitempty"`

	// generated from element "work-title" of type string order 1195 depth 1
	Work_title string `xml:"work-title,omitempty"`

	// generated from element "opus" of type opus order 1196 depth 1
	Opus *Opus `xml:"opus,omitempty"`
}

// Group_editorial UnNamed source named group "editorial"
type Group_editorial struct {

	// insertion point for fields

	// generated from group with order 1198 depth 1
	Group_footnote

	// generated from group with order 1199 depth 1
	Group_level
}

// Group_editorial_voice UnNamed source named group "editorial-voice"
type Group_editorial_voice struct {

	// insertion point for fields

	// generated from group with order 1201 depth 1
	Group_footnote

	// generated from group with order 1202 depth 1
	Group_level

	// generated from group with order 1203 depth 1
	Group_voice
}

// Group_editorial_voice_direction UnNamed source named group "editorial-voice-direction"
type Group_editorial_voice_direction struct {

	// insertion point for fields

	// generated from group with order 1205 depth 1
	Group_footnote

	// generated from group with order 1206 depth 1
	Group_level

	// generated from group with order 1207 depth 1
	Group_voice
}

// Group_footnote UnNamed source named group "footnote"
type Group_footnote struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 1209 depth 1
	Footnote *Formatted_text `xml:"footnote,omitempty"`
}

// Group_level UnNamed source named group "level"
type Group_level struct {

	// insertion point for fields

	// generated from element "level" of type level order 1211 depth 1
	Level *Level `xml:"level,omitempty"`
}

// Group_staff UnNamed source named group "staff"
type Group_staff struct {

	// insertion point for fields

	// generated from element "staff" of type positiveInteger order 1213 depth 1
	Staff int `xml:"staff,omitempty"`
}

// Group_tuning UnNamed source named group "tuning"
type Group_tuning struct {

	// insertion point for fields

	// generated from element "tuning-step" of type step order 1215 depth 1
	Tuning_step Enum_Step `xml:"tuning-step,omitempty"`

	// generated from element "tuning-alter" of type semitones order 1216 depth 1
	Tuning_alter string `xml:"tuning-alter,omitempty"`

	// generated from element "tuning-octave" of type octave order 1217 depth 1
	Tuning_octave int `xml:"tuning-octave,omitempty"`
}

// Group_virtual_instrument_data UnNamed source named group "virtual-instrument-data"
type Group_virtual_instrument_data struct {

	// insertion point for fields

	// generated from element "instrument-sound" of type string order 1219 depth 1
	Instrument_sound string `xml:"instrument-sound,omitempty"`

	// generated from element "solo" of type empty order 1220 depth 1
	Solo string `xml:"solo,omitempty"`

	// generated from element "ensemble" of type positive-integer-or-empty order 1221 depth 1
	Ensemble string `xml:"ensemble,omitempty"`

	// generated from element "virtual-instrument" of type virtual-instrument order 1222 depth 1
	Virtual_instrument *Virtual_instrument `xml:"virtual-instrument,omitempty"`
}

// Group_voice UnNamed source named group "voice"
type Group_voice struct {

	// insertion point for fields

	// generated from element "voice" of type string order 1224 depth 1
	Voice string `xml:"voice,omitempty"`
}

// Group_clef UnNamed source named group "clef"
type Group_clef struct {

	// insertion point for fields

	// generated from element "sign" of type clef-sign order 1226 depth 1
	Sign string `xml:"sign,omitempty"`

	// generated from element "line" of type staff-line-position order 1227 depth 1
	Line int `xml:"line,omitempty"`

	// generated from element "clef-octave-change" of type integer order 1228 depth 1
	Clef_octave_change int `xml:"clef-octave-change,omitempty"`
}

// Group_non_traditional_key UnNamed source named group "non-traditional-key"
type Group_non_traditional_key struct {

	// insertion point for fields

	// generated from element "key-step" of type step order 1230 depth 1
	Key_step Enum_Step `xml:"key-step,omitempty"`

	// generated from element "key-alter" of type semitones order 1231 depth 1
	Key_alter string `xml:"key-alter,omitempty"`

	// generated from element "key-accidental" of type key-accidental order 1232 depth 1
	Key_accidental *Key_accidental `xml:"key-accidental,omitempty"`
}

// Group_slash UnNamed source named group "slash"
type Group_slash struct {

	// insertion point for fields

	// generated from element "slash-type" of type note-type-value order 1234 depth 1
	Slash_type Enum_Note_type_value `xml:"slash-type,omitempty"`

	// generated from element "slash-dot" of type empty order 1235 depth 1
	Slash_dot string `xml:"slash-dot,omitempty"`

	// generated from element "except-voice" of type string order 1236 depth 1
	Except_voice string `xml:"except-voice,omitempty"`
}

// Group_time_signature UnNamed source named group "time-signature"
type Group_time_signature struct {

	// insertion point for fields

	// generated from element "beats" of type string order 1238 depth 1
	Beats string `xml:"beats,omitempty"`

	// generated from element "beat-type" of type string order 1239 depth 1
	Beat_type string `xml:"beat-type,omitempty"`
}

// Group_traditional_key UnNamed source named group "traditional-key"
type Group_traditional_key struct {

	// insertion point for fields

	// generated from element "cancel" of type cancel order 1241 depth 1
	Cancel *Cancel `xml:"cancel,omitempty"`

	// generated from element "fifths" of type fifths order 1242 depth 1
	Fifths int `xml:"fifths,omitempty"`

	// generated from element "mode" of type mode order 1243 depth 1
	Mode string `xml:"mode,omitempty"`
}

// Group_transpose UnNamed source named group "transpose"
type Group_transpose struct {

	// insertion point for fields

	// generated from element "diatonic" of type integer order 1245 depth 1
	Diatonic int `xml:"diatonic,omitempty"`

	// generated from element "chromatic" of type semitones order 1246 depth 1
	Chromatic string `xml:"chromatic,omitempty"`

	// generated from element "octave-change" of type integer order 1247 depth 1
	Octave_change int `xml:"octave-change,omitempty"`

	// generated from element "double" of type double order 1248 depth 1
	Double float64 `xml:"double,omitempty"`
}

// Group_beat_unit UnNamed source named group "beat-unit"
type Group_beat_unit struct {

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value order 1250 depth 1
	Beat_unit Enum_Note_type_value `xml:"beat-unit,omitempty"`

	// generated from element "beat-unit-dot" of type empty order 1251 depth 1
	Beat_unit_dot string `xml:"beat-unit-dot,omitempty"`
}

// Group_harmony_chord UnNamed source named group "harmony-chord"
type Group_harmony_chord struct {

	// insertion point for fields

	// generated from element "root" of type root order 1253 depth 1
	Root *Root `xml:"root,omitempty"`

	// generated from element "numeral" of type numeral order 1254 depth 1
	Numeral *Numeral `xml:"numeral,omitempty"`

	// generated from element "function" of type style-text order 1255 depth 1
	Function *Style_text `xml:"function,omitempty"`

	// generated from element "kind" of type kind order 1256 depth 1
	Kind *Kind `xml:"kind,omitempty"`

	// generated from element "inversion" of type inversion order 1257 depth 1
	Inversion *Inversion `xml:"inversion,omitempty"`

	// generated from element "bass" of type bass order 1258 depth 1
	Bass *Bass `xml:"bass,omitempty"`

	// generated from element "degree" of type degree order 1259 depth 1
	Degree []*Degree `xml:"degree,omitempty"`
}

// Group_all_margins UnNamed source named group "all-margins"
type Group_all_margins struct {

	// insertion point for fields

	// generated from group with order 1261 depth 1
	Group_left_right_margins

	// generated from element "top-margin" of type tenths order 1262 depth 1
	Top_margin string `xml:"top-margin,omitempty"`

	// generated from element "bottom-margin" of type tenths order 1263 depth 1
	Bottom_margin string `xml:"bottom-margin,omitempty"`
}

// Group_layout UnNamed source named group "layout"
type Group_layout struct {

	// insertion point for fields

	// generated from element "page-layout" of type page-layout order 1265 depth 1
	Page_layout *Page_layout `xml:"page-layout,omitempty"`

	// generated from element "system-layout" of type system-layout order 1266 depth 1
	System_layout *System_layout `xml:"system-layout,omitempty"`

	// generated from element "staff-layout" of type staff-layout order 1267 depth 1
	Staff_layout []*Staff_layout `xml:"staff-layout,omitempty"`
}

// Group_left_right_margins UnNamed source named group "left-right-margins"
type Group_left_right_margins struct {

	// insertion point for fields

	// generated from element "left-margin" of type tenths order 1269 depth 1
	Left_margin string `xml:"left-margin,omitempty"`

	// generated from element "right-margin" of type tenths order 1270 depth 1
	Right_margin string `xml:"right-margin,omitempty"`
}

// Group_duration UnNamed source named group "duration"
type Group_duration struct {

	// insertion point for fields

	// generated from element "duration" of type positive-divisions order 1272 depth 1
	Duration string `xml:"duration,omitempty"`
}

// Group_display_step_octave UnNamed source named group "display-step-octave"
type Group_display_step_octave struct {

	// insertion point for fields

	// generated from element "display-step" of type step order 1274 depth 1
	Display_step Enum_Step `xml:"display-step,omitempty"`

	// generated from element "display-octave" of type octave order 1275 depth 1
	Display_octave int `xml:"display-octave,omitempty"`
}

// Group_full_note UnNamed source named group "full-note"
type Group_full_note struct {

	// insertion point for fields

	// generated from element "chord" of type empty order 1277 depth 1
	Chord string `xml:"chord,omitempty"`

	// generated from element "pitch" of type pitch order 1278 depth 1
	Pitch *Pitch `xml:"pitch,omitempty"`

	// generated from element "unpitched" of type unpitched order 1279 depth 1
	Unpitched *Unpitched `xml:"unpitched,omitempty"`

	// generated from element "rest" of type rest order 1280 depth 1
	Rest *Rest `xml:"rest,omitempty"`
}

// Group_music_data UnNamed source named group "music-data"
type Group_music_data struct {

	// insertion point for fields

	// generated from element "note" of type note order 1282 depth 1
	Note []*Note `xml:"note,omitempty"`

	// generated from element "backup" of type backup order 1283 depth 1
	Backup []*Backup `xml:"backup,omitempty"`

	// generated from element "forward" of type forward order 1284 depth 1
	Forward []*Forward `xml:"forward,omitempty"`

	// generated from element "direction" of type direction order 1285 depth 1
	Direction []*Direction `xml:"direction,omitempty"`

	// generated from element "attributes" of type attributes order 1286 depth 1
	Attributes []*Attributes `xml:"attributes,omitempty"`

	// generated from element "harmony" of type harmony order 1287 depth 1
	Harmony []*Harmony `xml:"harmony,omitempty"`

	// generated from element "figured-bass" of type figured-bass order 1288 depth 1
	Figured_bass []*Figured_bass `xml:"figured-bass,omitempty"`

	// generated from element "print" of type print order 1289 depth 1
	Print []*Print `xml:"print,omitempty"`

	// generated from element "sound" of type sound order 1290 depth 1
	Sound []*Sound `xml:"sound,omitempty"`

	// generated from element "listening" of type listening order 1291 depth 1
	Listening []*Listening `xml:"listening,omitempty"`

	// generated from element "barline" of type barline order 1292 depth 1
	Barline []*Barline `xml:"barline,omitempty"`

	// generated from element "grouping" of type grouping order 1293 depth 1
	Grouping []*Grouping `xml:"grouping,omitempty"`

	// generated from element "link" of type link order 1294 depth 1
	Link []*Link `xml:"link,omitempty"`

	// generated from element "bookmark" of type bookmark order 1295 depth 1
	Bookmark []*Bookmark `xml:"bookmark,omitempty"`
}

// Group_part_group UnNamed source named group "part-group"
type Group_part_group struct {

	// insertion point for fields

	// generated from element "part-group" of type part-group order 1297 depth 1
	Part_group *Part_group `xml:"part-group,omitempty"`
}

// Group_score_header UnNamed source named group "score-header"
type Group_score_header struct {

	// insertion point for fields

	// generated from element "work" of type work order 1299 depth 1
	Work *Work `xml:"work,omitempty"`

	// generated from element "movement-number" of type string order 1300 depth 1
	Movement_number string `xml:"movement-number,omitempty"`

	// generated from element "movement-title" of type string order 1301 depth 1
	Movement_title string `xml:"movement-title,omitempty"`

	// generated from element "identification" of type identification order 1302 depth 1
	Identification *Identification `xml:"identification,omitempty"`

	// generated from element "defaults" of type defaults order 1303 depth 1
	Defaults *Defaults `xml:"defaults,omitempty"`

	// generated from element "credit" of type credit order 1304 depth 1
	Credit []*Credit `xml:"credit,omitempty"`

	// generated from element "part-list" of type part-list order 1305 depth 1
	Part_list *Part_list `xml:"part-list,omitempty"`
}

// Group_score_part UnNamed source named group "score-part"
type Group_score_part struct {

	// insertion point for fields

	// generated from element "score-part" of type score-part order 1307 depth 1
	Score_part *Score_part `xml:"score-part,omitempty"`
}

// Score_partwise Named source element score-partwise within root schema
// The score-partwise element is the root element for a partwise MusicXML
// score. It includes a score-header group followed by a series of parts with measures
// inside. The document-attributes attribute group includes the version attribute.
type Score_partwise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"score-partwise"`

	// generated from inline complex type
	A_score_partwise
}

// A_score_partwise Named source within outer element "score-partwise"
type A_score_partwise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from group with order 1310 depth 2
	Group_score_header

	// generated from anonymous type within outer element "part" of type A_part.
	Part []*A_part `xml:"part,omitempty"`
}

// A_part Named source within outer element "part"
type A_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from anonymous type within outer element "measure" of type A_measure.
	Measure []*A_measure `xml:"measure,omitempty"`
}

// A_measure Named source within outer element "measure"
type A_measure struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from group with order 1315 depth 6
	Group_music_data
}

// Score_timewise Named source element score-timewise within root schema
// The score-timewise element is the root element for a timewise MusicXML
// score. It includes a score-header group followed by a series of measures with parts
// inside. The document-attributes attribute group includes the version attribute.
type Score_timewise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"score-timewise"`

	// generated from inline complex type
	A_score_timewise
}

// A_score_timewise Named source within outer element "score-timewise"
type A_score_timewise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from group with order 1321 depth 2
	Group_score_header

	// generated from anonymous type within outer element "measure" of type A_measure.
	Measure []*A_measure_1 `xml:"measure,omitempty"`
}

// A_measure_1 Named source within outer element "measure"
type A_measure_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from anonymous type within outer element "part" of type A_part.
	Part []*A_part_1 `xml:"part,omitempty"`
}

// A_part_1 Named source within outer element "part"
type A_part_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from group with order 1326 depth 6
	Group_music_data
}
