// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xmlxsd"
)

type FormCallbackIF = gongprobe.FormCallbackIF

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func saveStageSet_A_directive_Stage(
	inst *models.A_directive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_A_measure_Stage(
	inst *models.A_measure,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Implicit":
			StageSetFormDivEnumStringFieldToField(&inst.Implicit, formDiv)
		case "Non_controlling":
			StageSetFormDivEnumStringFieldToField(&inst.Non_controlling, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Note":
			FormDivSliceOfPointersToField(inst, "Note", &inst.Note, formDiv, probe)
		case "Backup":
			FormDivSliceOfPointersToField(inst, "Backup", &inst.Backup, formDiv, probe)
		case "Forward":
			FormDivSliceOfPointersToField(inst, "Forward", &inst.Forward, formDiv, probe)
		case "Direction":
			FormDivSliceOfPointersToField(inst, "Direction", &inst.Direction, formDiv, probe)
		case "Attributes":
			FormDivSliceOfPointersToField(inst, "Attributes", &inst.Attributes, formDiv, probe)
		case "Harmony":
			FormDivSliceOfPointersToField(inst, "Harmony", &inst.Harmony, formDiv, probe)
		case "Figured_bass":
			FormDivSliceOfPointersToField(inst, "Figured_bass", &inst.Figured_bass, formDiv, probe)
		case "Print_XSD":
			FormDivSliceOfPointersToField(inst, "Print_XSD", &inst.Print_XSD, formDiv, probe)
		case "Sound":
			FormDivSliceOfPointersToField(inst, "Sound", &inst.Sound, formDiv, probe)
		case "Listening":
			FormDivSliceOfPointersToField(inst, "Listening", &inst.Listening, formDiv, probe)
		case "Barline":
			FormDivSliceOfPointersToField(inst, "Barline", &inst.Barline, formDiv, probe)
		case "Grouping":
			FormDivSliceOfPointersToField(inst, "Grouping", &inst.Grouping, formDiv, probe)
		case "Link":
			FormDivSliceOfPointersToField(inst, "Link", &inst.Link, formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(inst, "Bookmark", &inst.Bookmark, formDiv, probe)
		}
	}
}

func saveStageSet_A_measure_1_Stage(
	inst *models.A_measure_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Implicit":
			StageSetFormDivEnumStringFieldToField(&inst.Implicit, formDiv)
		case "Non_controlling":
			StageSetFormDivEnumStringFieldToField(&inst.Non_controlling, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Part":
			FormDivSliceOfPointersToField(inst, "Part", &inst.Part, formDiv, probe)
		}
	}
}

func saveStageSet_A_part_Stage(
	inst *models.A_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Measure":
			FormDivSliceOfPointersToField(inst, "Measure", &inst.Measure, formDiv, probe)
		}
	}
}

func saveStageSet_A_part_1_Stage(
	inst *models.A_part_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Note":
			FormDivSliceOfPointersToField(inst, "Note", &inst.Note, formDiv, probe)
		case "Backup":
			FormDivSliceOfPointersToField(inst, "Backup", &inst.Backup, formDiv, probe)
		case "Forward":
			FormDivSliceOfPointersToField(inst, "Forward", &inst.Forward, formDiv, probe)
		case "Direction":
			FormDivSliceOfPointersToField(inst, "Direction", &inst.Direction, formDiv, probe)
		case "Attributes":
			FormDivSliceOfPointersToField(inst, "Attributes", &inst.Attributes, formDiv, probe)
		case "Harmony":
			FormDivSliceOfPointersToField(inst, "Harmony", &inst.Harmony, formDiv, probe)
		case "Figured_bass":
			FormDivSliceOfPointersToField(inst, "Figured_bass", &inst.Figured_bass, formDiv, probe)
		case "Print_XSD":
			FormDivSliceOfPointersToField(inst, "Print_XSD", &inst.Print_XSD, formDiv, probe)
		case "Sound":
			FormDivSliceOfPointersToField(inst, "Sound", &inst.Sound, formDiv, probe)
		case "Listening":
			FormDivSliceOfPointersToField(inst, "Listening", &inst.Listening, formDiv, probe)
		case "Barline":
			FormDivSliceOfPointersToField(inst, "Barline", &inst.Barline, formDiv, probe)
		case "Grouping":
			FormDivSliceOfPointersToField(inst, "Grouping", &inst.Grouping, formDiv, probe)
		case "Link":
			FormDivSliceOfPointersToField(inst, "Link", &inst.Link, formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(inst, "Bookmark", &inst.Bookmark, formDiv, probe)
		}
	}
}

func saveStageSet_Accidental_Stage(
	inst *models.Accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Cautionary":
			StageSetFormDivEnumStringFieldToField(&inst.Cautionary, formDiv)
		case "Editorial":
			StageSetFormDivEnumStringFieldToField(&inst.Editorial, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Accidental_mark_Stage(
	inst *models.Accidental_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Accidental_text_Stage(
	inst *models.Accidental_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&inst.Line_height, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Accord_Stage(
	inst *models.Accord,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "String_XSD":
			FormDivBasicFieldToField(&inst.String_XSD, formDiv)
		case "Tuning_step":
			StageSetFormDivEnumStringFieldToField(&inst.Tuning_step, formDiv)
		case "Tuning_alter":
			FormDivBasicFieldToField(&inst.Tuning_alter, formDiv)
		case "Tuning_octave":
			FormDivBasicFieldToField(&inst.Tuning_octave, formDiv)
		}
	}
}

func saveStageSet_Accordion_registration_Stage(
	inst *models.Accordion_registration,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Accordion_high":
			FormDivBasicFieldToField(&inst.Accordion_high, formDiv)
		case "Accordion_middle":
			FormDivBasicFieldToField(&inst.Accordion_middle, formDiv)
		case "Accordion_low":
			FormDivBasicFieldToField(&inst.Accordion_low, formDiv)
		}
	}
}

func saveStageSet_Appearance_Stage(
	inst *models.Appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line_width":
			FormDivSliceOfPointersToField(inst, "Line_width", &inst.Line_width, formDiv, probe)
		case "Note_size":
			FormDivSliceOfPointersToField(inst, "Note_size", &inst.Note_size, formDiv, probe)
		case "Distance":
			FormDivSliceOfPointersToField(inst, "Distance", &inst.Distance, formDiv, probe)
		case "Glyph":
			FormDivSliceOfPointersToField(inst, "Glyph", &inst.Glyph, formDiv, probe)
		case "Other_appearance":
			FormDivSliceOfPointersToField(inst, "Other_appearance", &inst.Other_appearance, formDiv, probe)
		}
	}
}

func saveStageSet_Arpeggiate_Stage(
	inst *models.Arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Direction":
			StageSetFormDivEnumStringFieldToField(&inst.Direction, formDiv)
		case "Unbroken":
			StageSetFormDivEnumStringFieldToField(&inst.Unbroken, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Arrow_Stage(
	inst *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Arrow_direction":
			StageSetFormDivEnumStringFieldToField(&inst.Arrow_direction, formDiv)
		case "Arrow_style":
			StageSetFormDivEnumStringFieldToField(&inst.Arrow_style, formDiv)
		case "Arrowhead":
			FormDivBasicFieldToField(&inst.Arrowhead, formDiv)
		case "Circular_arrow":
			StageSetFormDivEnumStringFieldToField(&inst.Circular_arrow, formDiv)
		}
	}
}

func saveStageSet_Articulations_Stage(
	inst *models.Articulations,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Accent":
			FormDivSliceOfPointersToField(inst, "Accent", &inst.Accent, formDiv, probe)
		case "Strong_accent":
			FormDivSliceOfPointersToField(inst, "Strong_accent", &inst.Strong_accent, formDiv, probe)
		case "Staccato":
			FormDivSliceOfPointersToField(inst, "Staccato", &inst.Staccato, formDiv, probe)
		case "Tenuto":
			FormDivSliceOfPointersToField(inst, "Tenuto", &inst.Tenuto, formDiv, probe)
		case "Detached_legato":
			FormDivSliceOfPointersToField(inst, "Detached_legato", &inst.Detached_legato, formDiv, probe)
		case "Staccatissimo":
			FormDivSliceOfPointersToField(inst, "Staccatissimo", &inst.Staccatissimo, formDiv, probe)
		case "Spiccato":
			FormDivSliceOfPointersToField(inst, "Spiccato", &inst.Spiccato, formDiv, probe)
		case "Scoop":
			FormDivSliceOfPointersToField(inst, "Scoop", &inst.Scoop, formDiv, probe)
		case "Plop":
			FormDivSliceOfPointersToField(inst, "Plop", &inst.Plop, formDiv, probe)
		case "Doit":
			FormDivSliceOfPointersToField(inst, "Doit", &inst.Doit, formDiv, probe)
		case "Falloff":
			FormDivSliceOfPointersToField(inst, "Falloff", &inst.Falloff, formDiv, probe)
		case "Breath_mark":
			FormDivSliceOfPointersToField(inst, "Breath_mark", &inst.Breath_mark, formDiv, probe)
		case "Caesura":
			FormDivSliceOfPointersToField(inst, "Caesura", &inst.Caesura, formDiv, probe)
		case "Stress":
			FormDivSliceOfPointersToField(inst, "Stress", &inst.Stress, formDiv, probe)
		case "Unstress":
			FormDivSliceOfPointersToField(inst, "Unstress", &inst.Unstress, formDiv, probe)
		case "Soft_accent":
			FormDivSliceOfPointersToField(inst, "Soft_accent", &inst.Soft_accent, formDiv, probe)
		case "Other_articulation":
			FormDivSliceOfPointersToField(inst, "Other_articulation", &inst.Other_articulation, formDiv, probe)
		}
	}
}

func saveStageSet_Assess_Stage(
	inst *models.Assess,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		}
	}
}

func saveStageSet_Attributes_Stage(
	inst *models.Attributes,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "Key":
			FormDivSliceOfPointersToField(inst, "Key", &inst.Key, formDiv, probe)
		case "Time":
			FormDivBasicFieldToField(&inst.Time, formDiv)
		case "Staves":
			FormDivBasicFieldToField(&inst.Staves, formDiv)
		case "Part_symbol":
			StageSetFormDivSelectFieldToField(&inst.Part_symbol, probe.stageOfInterest.GetInstancesSet[*models.Part_symbol](), formDiv)
		case "Instruments":
			FormDivBasicFieldToField(&inst.Instruments, formDiv)
		case "Clef":
			FormDivSliceOfPointersToField(inst, "Clef", &inst.Clef, formDiv, probe)
		case "Staff_details":
			FormDivSliceOfPointersToField(inst, "Staff_details", &inst.Staff_details, formDiv, probe)
		case "Transpose":
			FormDivSliceOfPointersToField(inst, "Transpose", &inst.Transpose, formDiv, probe)
		case "For_part":
			FormDivSliceOfPointersToField(inst, "For_part", &inst.For_part, formDiv, probe)
		case "Directive":
			FormDivSliceOfPointersToField(inst, "Directive", &inst.Directive, formDiv, probe)
		case "Measure_style":
			FormDivSliceOfPointersToField(inst, "Measure_style", &inst.Measure_style, formDiv, probe)
		}
	}
}

func saveStageSet_Backup_Stage(
	inst *models.Backup,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		}
	}
}

func saveStageSet_Bar_style_color_Stage(
	inst *models.Bar_style_color,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Barline_Stage(
	inst *models.Barline,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&inst.Segno, formDiv)
		case "Coda":
			FormDivBasicFieldToField(&inst.Coda, formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Bar_style":
			StageSetFormDivSelectFieldToField(&inst.Bar_style, probe.stageOfInterest.GetInstancesSet[*models.Bar_style_color](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Wavy_line":
			StageSetFormDivSelectFieldToField(&inst.Wavy_line, probe.stageOfInterest.GetInstancesSet[*models.Wavy_line](), formDiv)
		case "Segno_1":
			StageSetFormDivSelectFieldToField(&inst.Segno_1, probe.stageOfInterest.GetInstancesSet[*models.Segno](), formDiv)
		case "Coda_1":
			StageSetFormDivSelectFieldToField(&inst.Coda_1, probe.stageOfInterest.GetInstancesSet[*models.Coda](), formDiv)
		case "Fermata":
			StageSetFormDivSelectFieldToField(&inst.Fermata, probe.stageOfInterest.GetInstancesSet[*models.Fermata](), formDiv)
		case "Ending":
			StageSetFormDivSelectFieldToField(&inst.Ending, probe.stageOfInterest.GetInstancesSet[*models.Ending](), formDiv)
		case "Repeat":
			StageSetFormDivSelectFieldToField(&inst.Repeat, probe.stageOfInterest.GetInstancesSet[*models.Repeat](), formDiv)
		}
	}
}

func saveStageSet_Barre_Stage(
	inst *models.Barre,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Bass_Stage(
	inst *models.Bass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Arrangement":
			StageSetFormDivEnumStringFieldToField(&inst.Arrangement, formDiv)
		case "Bass_separator":
			StageSetFormDivSelectFieldToField(&inst.Bass_separator, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), formDiv)
		case "Bass_step":
			StageSetFormDivSelectFieldToField(&inst.Bass_step, probe.stageOfInterest.GetInstancesSet[*models.Bass_step](), formDiv)
		case "Bass_alter":
			StageSetFormDivSelectFieldToField(&inst.Bass_alter, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), formDiv)
		}
	}
}

func saveStageSet_Bass_step_Stage(
	inst *models.Bass_step,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Beam_Stage(
	inst *models.Beam,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Repeater":
			StageSetFormDivEnumStringFieldToField(&inst.Repeater, formDiv)
		case "Fan":
			StageSetFormDivEnumStringFieldToField(&inst.Fan, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Beat_repeat_Stage(
	inst *models.Beat_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&inst.Slashes, formDiv)
		case "Use_dots":
			StageSetFormDivEnumStringFieldToField(&inst.Use_dots, formDiv)
		case "Slash_type":
			StageSetFormDivEnumStringFieldToField(&inst.Slash_type, formDiv)
		case "Slash_dot":
			FormDivBasicFieldToField(&inst.Slash_dot, formDiv)
		case "Except_voice":
			FormDivBasicFieldToField(&inst.Except_voice, formDiv)
		}
	}
}

func saveStageSet_Beat_unit_tied_Stage(
	inst *models.Beat_unit_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Beat_unit":
			StageSetFormDivEnumStringFieldToField(&inst.Beat_unit, formDiv)
		case "Beat_unit_dot":
			FormDivBasicFieldToField(&inst.Beat_unit_dot, formDiv)
		}
	}
}

func saveStageSet_Beater_Stage(
	inst *models.Beater,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tip":
			StageSetFormDivEnumStringFieldToField(&inst.Tip, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Bend_Stage(
	inst *models.Bend,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Shape":
			StageSetFormDivEnumStringFieldToField(&inst.Shape, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "First_beat":
			FormDivBasicFieldToField(&inst.First_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		case "Bend_alter":
			FormDivBasicFieldToField(&inst.Bend_alter, formDiv)
		case "Pre_bend":
			FormDivBasicFieldToField(&inst.Pre_bend, formDiv)
		case "Release":
			StageSetFormDivSelectFieldToField(&inst.Release, probe.stageOfInterest.GetInstancesSet[*models.Release](), formDiv)
		case "With_bar":
			StageSetFormDivSelectFieldToField(&inst.With_bar, probe.stageOfInterest.GetInstancesSet[*models.Placement_text](), formDiv)
		}
	}
}

func saveStageSet_Bookmark_Stage(
	inst *models.Bookmark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Position":
			FormDivBasicFieldToField(&inst.Position, formDiv)
		}
	}
}

func saveStageSet_Bracket_Stage(
	inst *models.Bracket,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line_end":
			StageSetFormDivEnumStringFieldToField(&inst.Line_end, formDiv)
		case "End_length":
			FormDivBasicFieldToField(&inst.End_length, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Breath_mark_Stage(
	inst *models.Breath_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Caesura_Stage(
	inst *models.Caesura,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Cancel_Stage(
	inst *models.Cancel,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Clef_Stage(
	inst *models.Clef,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Additional":
			StageSetFormDivEnumStringFieldToField(&inst.Additional, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "After_barline":
			StageSetFormDivEnumStringFieldToField(&inst.After_barline, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "Clef_octave_change":
			FormDivBasicFieldToField(&inst.Clef_octave_change, formDiv)
		}
	}
}

func saveStageSet_Coda_Stage(
	inst *models.Coda,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Credit_Stage(
	inst *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Page":
			FormDivBasicFieldToField(&inst.Page, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Credit_type":
			FormDivBasicFieldToField(&inst.Credit_type, formDiv)
		case "Link":
			FormDivSliceOfPointersToField(inst, "Link", &inst.Link, formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(inst, "Bookmark", &inst.Bookmark, formDiv, probe)
		case "Credit_image":
			StageSetFormDivSelectFieldToField(&inst.Credit_image, probe.stageOfInterest.GetInstancesSet[*models.Image](), formDiv)
		case "Credit_words":
			StageSetFormDivSelectFieldToField(&inst.Credit_words, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text_id](), formDiv)
		case "Credit_symbol":
			StageSetFormDivSelectFieldToField(&inst.Credit_symbol, probe.stageOfInterest.GetInstancesSet[*models.Formatted_symbol_id](), formDiv)
		}
	}
}

func saveStageSet_Dashes_Stage(
	inst *models.Dashes,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Defaults_Stage(
	inst *models.Defaults,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Scaling":
			StageSetFormDivSelectFieldToField(&inst.Scaling, probe.stageOfInterest.GetInstancesSet[*models.Scaling](), formDiv)
		case "Concert_score":
			FormDivBasicFieldToField(&inst.Concert_score, formDiv)
		case "Page_layout":
			StageSetFormDivSelectFieldToField(&inst.Page_layout, probe.stageOfInterest.GetInstancesSet[*models.Page_layout](), formDiv)
		case "System_layout":
			StageSetFormDivSelectFieldToField(&inst.System_layout, probe.stageOfInterest.GetInstancesSet[*models.System_layout](), formDiv)
		case "Staff_layout":
			FormDivSliceOfPointersToField(inst, "Staff_layout", &inst.Staff_layout, formDiv, probe)
		case "Appearance":
			StageSetFormDivSelectFieldToField(&inst.Appearance, probe.stageOfInterest.GetInstancesSet[*models.Appearance](), formDiv)
		case "Music_font":
			StageSetFormDivSelectFieldToField(&inst.Music_font, probe.stageOfInterest.GetInstancesSet[*models.Empty_font](), formDiv)
		case "Word_font":
			StageSetFormDivSelectFieldToField(&inst.Word_font, probe.stageOfInterest.GetInstancesSet[*models.Empty_font](), formDiv)
		case "Lyric_font":
			FormDivSliceOfPointersToField(inst, "Lyric_font", &inst.Lyric_font, formDiv, probe)
		case "Lyric_language":
			FormDivSliceOfPointersToField(inst, "Lyric_language", &inst.Lyric_language, formDiv, probe)
		}
	}
}

func saveStageSet_Degree_Stage(
	inst *models.Degree,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Degree_value":
			StageSetFormDivSelectFieldToField(&inst.Degree_value, probe.stageOfInterest.GetInstancesSet[*models.Degree_value](), formDiv)
		case "Degree_alter":
			StageSetFormDivSelectFieldToField(&inst.Degree_alter, probe.stageOfInterest.GetInstancesSet[*models.Degree_alter](), formDiv)
		case "Degree_type":
			StageSetFormDivSelectFieldToField(&inst.Degree_type, probe.stageOfInterest.GetInstancesSet[*models.Degree_type](), formDiv)
		}
	}
}

func saveStageSet_Degree_alter_Stage(
	inst *models.Degree_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Plus_minus":
			StageSetFormDivEnumStringFieldToField(&inst.Plus_minus, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Degree_type_Stage(
	inst *models.Degree_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Degree_value_Stage(
	inst *models.Degree_value,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Direction_Stage(
	inst *models.Direction,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Directive":
			StageSetFormDivEnumStringFieldToField(&inst.Directive, formDiv)
		case "System":
			FormDivBasicFieldToField(&inst.System, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Direction_type":
			FormDivSliceOfPointersToField(inst, "Direction_type", &inst.Direction_type, formDiv, probe)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageOfInterest.GetInstancesSet[*models.Offset](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "Sound":
			StageSetFormDivSelectFieldToField(&inst.Sound, probe.stageOfInterest.GetInstancesSet[*models.Sound](), formDiv)
		case "Listening":
			StageSetFormDivSelectFieldToField(&inst.Listening, probe.stageOfInterest.GetInstancesSet[*models.Listening](), formDiv)
		}
	}
}

func saveStageSet_Direction_type_Stage(
	inst *models.Direction_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Rehearsal":
			FormDivSliceOfPointersToField(inst, "Rehearsal", &inst.Rehearsal, formDiv, probe)
		case "Segno":
			FormDivSliceOfPointersToField(inst, "Segno", &inst.Segno, formDiv, probe)
		case "Coda":
			FormDivSliceOfPointersToField(inst, "Coda", &inst.Coda, formDiv, probe)
		case "Words":
			FormDivSliceOfPointersToField(inst, "Words", &inst.Words, formDiv, probe)
		case "Symbol":
			FormDivSliceOfPointersToField(inst, "Symbol", &inst.Symbol, formDiv, probe)
		case "Wedge":
			StageSetFormDivSelectFieldToField(&inst.Wedge, probe.stageOfInterest.GetInstancesSet[*models.Wedge](), formDiv)
		case "Dynamics":
			FormDivSliceOfPointersToField(inst, "Dynamics", &inst.Dynamics, formDiv, probe)
		case "Dashes":
			StageSetFormDivSelectFieldToField(&inst.Dashes, probe.stageOfInterest.GetInstancesSet[*models.Dashes](), formDiv)
		case "Bracket":
			StageSetFormDivSelectFieldToField(&inst.Bracket, probe.stageOfInterest.GetInstancesSet[*models.Bracket](), formDiv)
		case "Pedal":
			StageSetFormDivSelectFieldToField(&inst.Pedal, probe.stageOfInterest.GetInstancesSet[*models.Pedal](), formDiv)
		case "Metronome":
			StageSetFormDivSelectFieldToField(&inst.Metronome, probe.stageOfInterest.GetInstancesSet[*models.Metronome](), formDiv)
		case "Octave_shift":
			StageSetFormDivSelectFieldToField(&inst.Octave_shift, probe.stageOfInterest.GetInstancesSet[*models.Octave_shift](), formDiv)
		case "Harp_pedals":
			StageSetFormDivSelectFieldToField(&inst.Harp_pedals, probe.stageOfInterest.GetInstancesSet[*models.Harp_pedals](), formDiv)
		case "Damp":
			StageSetFormDivSelectFieldToField(&inst.Damp, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), formDiv)
		case "Damp_all":
			StageSetFormDivSelectFieldToField(&inst.Damp_all, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), formDiv)
		case "Eyeglasses":
			StageSetFormDivSelectFieldToField(&inst.Eyeglasses, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), formDiv)
		case "String_mute":
			StageSetFormDivSelectFieldToField(&inst.String_mute, probe.stageOfInterest.GetInstancesSet[*models.String_mute](), formDiv)
		case "Scordatura":
			StageSetFormDivSelectFieldToField(&inst.Scordatura, probe.stageOfInterest.GetInstancesSet[*models.Scordatura](), formDiv)
		case "Image":
			StageSetFormDivSelectFieldToField(&inst.Image, probe.stageOfInterest.GetInstancesSet[*models.Image](), formDiv)
		case "Principal_voice":
			StageSetFormDivSelectFieldToField(&inst.Principal_voice, probe.stageOfInterest.GetInstancesSet[*models.Principal_voice](), formDiv)
		case "Percussion":
			FormDivSliceOfPointersToField(inst, "Percussion", &inst.Percussion, formDiv, probe)
		case "Accordion_registration":
			StageSetFormDivSelectFieldToField(&inst.Accordion_registration, probe.stageOfInterest.GetInstancesSet[*models.Accordion_registration](), formDiv)
		case "Staff_divide":
			StageSetFormDivSelectFieldToField(&inst.Staff_divide, probe.stageOfInterest.GetInstancesSet[*models.Staff_divide](), formDiv)
		case "Other_direction":
			StageSetFormDivSelectFieldToField(&inst.Other_direction, probe.stageOfInterest.GetInstancesSet[*models.Other_direction](), formDiv)
		}
	}
}

func saveStageSet_Distance_Stage(
	inst *models.Distance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Double_Stage(
	inst *models.Double,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Above":
			StageSetFormDivEnumStringFieldToField(&inst.Above, formDiv)
		}
	}
}

func saveStageSet_Dynamics_Stage(
	inst *models.Dynamics,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "P":
			FormDivBasicFieldToField(&inst.P, formDiv)
		case "Pp":
			FormDivBasicFieldToField(&inst.Pp, formDiv)
		case "Ppp":
			FormDivBasicFieldToField(&inst.Ppp, formDiv)
		case "Pppp":
			FormDivBasicFieldToField(&inst.Pppp, formDiv)
		case "Ppppp":
			FormDivBasicFieldToField(&inst.Ppppp, formDiv)
		case "Pppppp":
			FormDivBasicFieldToField(&inst.Pppppp, formDiv)
		case "F":
			FormDivBasicFieldToField(&inst.F, formDiv)
		case "Ff":
			FormDivBasicFieldToField(&inst.Ff, formDiv)
		case "Fff":
			FormDivBasicFieldToField(&inst.Fff, formDiv)
		case "Ffff":
			FormDivBasicFieldToField(&inst.Ffff, formDiv)
		case "Fffff":
			FormDivBasicFieldToField(&inst.Fffff, formDiv)
		case "Ffffff":
			FormDivBasicFieldToField(&inst.Ffffff, formDiv)
		case "Mp":
			FormDivBasicFieldToField(&inst.Mp, formDiv)
		case "Mf":
			FormDivBasicFieldToField(&inst.Mf, formDiv)
		case "Sf":
			FormDivBasicFieldToField(&inst.Sf, formDiv)
		case "Sfp":
			FormDivBasicFieldToField(&inst.Sfp, formDiv)
		case "Sfpp":
			FormDivBasicFieldToField(&inst.Sfpp, formDiv)
		case "Fp":
			FormDivBasicFieldToField(&inst.Fp, formDiv)
		case "Rf":
			FormDivBasicFieldToField(&inst.Rf, formDiv)
		case "Rfz":
			FormDivBasicFieldToField(&inst.Rfz, formDiv)
		case "Sfz":
			FormDivBasicFieldToField(&inst.Sfz, formDiv)
		case "Sffz":
			FormDivBasicFieldToField(&inst.Sffz, formDiv)
		case "Fz":
			FormDivBasicFieldToField(&inst.Fz, formDiv)
		case "N":
			FormDivBasicFieldToField(&inst.N, formDiv)
		case "Pf":
			FormDivBasicFieldToField(&inst.Pf, formDiv)
		case "Sfzp":
			FormDivBasicFieldToField(&inst.Sfzp, formDiv)
		case "Other_dynamics":
			FormDivSliceOfPointersToField(inst, "Other_dynamics", &inst.Other_dynamics, formDiv, probe)
		}
	}
}

func saveStageSet_Effect_Stage(
	inst *models.Effect,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Elision_Stage(
	inst *models.Elision,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Empty_Stage(
	inst *models.Empty,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_Empty_font_Stage(
	inst *models.Empty_font,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		}
	}
}

func saveStageSet_Empty_line_Stage(
	inst *models.Empty_line,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line_shape":
			StageSetFormDivEnumStringFieldToField(&inst.Line_shape, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Line_length":
			StageSetFormDivEnumStringFieldToField(&inst.Line_length, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		}
	}
}

func saveStageSet_Empty_placement_Stage(
	inst *models.Empty_placement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		}
	}
}

func saveStageSet_Empty_placement_smufl_Stage(
	inst *models.Empty_placement_smufl,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		}
	}
}

func saveStageSet_Empty_print_object_style_align_Stage(
	inst *models.Empty_print_object_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		}
	}
}

func saveStageSet_Empty_print_style_Stage(
	inst *models.Empty_print_style,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Empty_print_style_align_Stage(
	inst *models.Empty_print_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		}
	}
}

func saveStageSet_Empty_print_style_align_id_Stage(
	inst *models.Empty_print_style_align_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Empty_trill_sound_Stage(
	inst *models.Empty_trill_sound,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Start_note":
			StageSetFormDivEnumStringFieldToField(&inst.Start_note, formDiv)
		case "Trill_step":
			StageSetFormDivEnumStringFieldToField(&inst.Trill_step, formDiv)
		case "Two_note_turn":
			StageSetFormDivEnumStringFieldToField(&inst.Two_note_turn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&inst.Second_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		}
	}
}

func saveStageSet_Encoding_Stage(
	inst *models.Encoding,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Encoding_date":
			FormDivBasicFieldToField(&inst.Encoding_date, formDiv)
		case "Encoder":
			FormDivSliceOfPointersToField(inst, "Encoder", &inst.Encoder, formDiv, probe)
		case "Software":
			FormDivBasicFieldToField(&inst.Software, formDiv)
		case "Encoding_description":
			FormDivBasicFieldToField(&inst.Encoding_description, formDiv)
		case "Supports":
			FormDivSliceOfPointersToField(inst, "Supports", &inst.Supports, formDiv, probe)
		}
	}
}

func saveStageSet_Ending_Stage(
	inst *models.Ending,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "End_length":
			FormDivBasicFieldToField(&inst.End_length, formDiv)
		case "Text_x":
			FormDivBasicFieldToField(&inst.Text_x, formDiv)
		case "Text_y":
			FormDivBasicFieldToField(&inst.Text_y, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "System":
			FormDivBasicFieldToField(&inst.System, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Extend_Stage(
	inst *models.Extend,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Feature_Stage(
	inst *models.Feature,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Fermata_Stage(
	inst *models.Fermata,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Figure_Stage(
	inst *models.Figure,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Prefix":
			StageSetFormDivSelectFieldToField(&inst.Prefix, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), formDiv)
		case "Figure_number":
			StageSetFormDivSelectFieldToField(&inst.Figure_number, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), formDiv)
		case "Suffix":
			StageSetFormDivSelectFieldToField(&inst.Suffix, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), formDiv)
		case "Extend":
			StageSetFormDivSelectFieldToField(&inst.Extend, probe.stageOfInterest.GetInstancesSet[*models.Extend](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		}
	}
}

func saveStageSet_Figured_bass_Stage(
	inst *models.Figured_bass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Print_dot":
			StageSetFormDivEnumStringFieldToField(&inst.Print_dot, formDiv)
		case "Print_lyric":
			StageSetFormDivEnumStringFieldToField(&inst.Print_lyric, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Print_spacing":
			StageSetFormDivEnumStringFieldToField(&inst.Print_spacing, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Figure":
			FormDivSliceOfPointersToField(inst, "Figure", &inst.Figure, formDiv, probe)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		}
	}
}

func saveStageSet_Fingering_Stage(
	inst *models.Fingering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Substitution":
			StageSetFormDivEnumStringFieldToField(&inst.Substitution, formDiv)
		case "Alternate":
			StageSetFormDivEnumStringFieldToField(&inst.Alternate, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_First_fret_Stage(
	inst *models.First_fret,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_For_part_Stage(
	inst *models.For_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Part_clef":
			StageSetFormDivSelectFieldToField(&inst.Part_clef, probe.stageOfInterest.GetInstancesSet[*models.Part_clef](), formDiv)
		case "Part_transpose":
			StageSetFormDivSelectFieldToField(&inst.Part_transpose, probe.stageOfInterest.GetInstancesSet[*models.Part_transpose](), formDiv)
		}
	}
}

func saveStageSet_Formatted_symbol_Stage(
	inst *models.Formatted_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&inst.Line_height, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Formatted_symbol_id_Stage(
	inst *models.Formatted_symbol_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&inst.Line_height, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Formatted_text_Stage(
	inst *models.Formatted_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&inst.Line_height, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Formatted_text_id_Stage(
	inst *models.Formatted_text_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&inst.Line_height, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Forward_Stage(
	inst *models.Forward,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		}
	}
}

func saveStageSet_Frame_Stage(
	inst *models.Frame,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Unplayed":
			FormDivBasicFieldToField(&inst.Unplayed, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Frame_strings":
			FormDivBasicFieldToField(&inst.Frame_strings, formDiv)
		case "Frame_frets":
			FormDivBasicFieldToField(&inst.Frame_frets, formDiv)
		case "First_fret":
			StageSetFormDivSelectFieldToField(&inst.First_fret, probe.stageOfInterest.GetInstancesSet[*models.First_fret](), formDiv)
		case "Frame_note":
			FormDivSliceOfPointersToField(inst, "Frame_note", &inst.Frame_note, formDiv, probe)
		}
	}
}

func saveStageSet_Frame_note_Stage(
	inst *models.Frame_note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "String_XSD":
			StageSetFormDivSelectFieldToField(&inst.String_XSD, probe.stageOfInterest.GetInstancesSet[*models.String_type](), formDiv)
		case "Fret":
			StageSetFormDivSelectFieldToField(&inst.Fret, probe.stageOfInterest.GetInstancesSet[*models.Fret](), formDiv)
		case "Fingering":
			StageSetFormDivSelectFieldToField(&inst.Fingering, probe.stageOfInterest.GetInstancesSet[*models.Fingering](), formDiv)
		case "Barre":
			StageSetFormDivSelectFieldToField(&inst.Barre, probe.stageOfInterest.GetInstancesSet[*models.Barre](), formDiv)
		}
	}
}

func saveStageSet_Fret_Stage(
	inst *models.Fret,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Glass_Stage(
	inst *models.Glass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Glissando_Stage(
	inst *models.Glissando,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Glyph_Stage(
	inst *models.Glyph,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Grace_Stage(
	inst *models.Grace,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Steal_time_previous":
			FormDivBasicFieldToField(&inst.Steal_time_previous, formDiv)
		case "Steal_time_following":
			FormDivBasicFieldToField(&inst.Steal_time_following, formDiv)
		case "Make_time":
			FormDivBasicFieldToField(&inst.Make_time, formDiv)
		case "Slash":
			StageSetFormDivEnumStringFieldToField(&inst.Slash, formDiv)
		}
	}
}

func saveStageSet_Group_barline_Stage(
	inst *models.Group_barline,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Group_name_Stage(
	inst *models.Group_name,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Group_symbol_Stage(
	inst *models.Group_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Grouping_Stage(
	inst *models.Grouping,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Member_of":
			FormDivBasicFieldToField(&inst.Member_of, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Feature":
			FormDivSliceOfPointersToField(inst, "Feature", &inst.Feature, formDiv, probe)
		}
	}
}

func saveStageSet_Hammer_on_pull_off_Stage(
	inst *models.Hammer_on_pull_off,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Handbell_Stage(
	inst *models.Handbell,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Harmon_closed_Stage(
	inst *models.Harmon_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Harmon_mute_Stage(
	inst *models.Harmon_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Harmon_closed":
			StageSetFormDivSelectFieldToField(&inst.Harmon_closed, probe.stageOfInterest.GetInstancesSet[*models.Harmon_closed](), formDiv)
		}
	}
}

func saveStageSet_Harmonic_Stage(
	inst *models.Harmonic,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Natural":
			FormDivBasicFieldToField(&inst.Natural, formDiv)
		case "Artificial":
			FormDivBasicFieldToField(&inst.Artificial, formDiv)
		case "Base_pitch":
			FormDivBasicFieldToField(&inst.Base_pitch, formDiv)
		case "Touching_pitch":
			FormDivBasicFieldToField(&inst.Touching_pitch, formDiv)
		case "Sounding_pitch":
			FormDivBasicFieldToField(&inst.Sounding_pitch, formDiv)
		}
	}
}

func saveStageSet_Harmony_Stage(
	inst *models.Harmony,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Print_frame":
			StageSetFormDivEnumStringFieldToField(&inst.Print_frame, formDiv)
		case "Arrangement":
			StageSetFormDivEnumStringFieldToField(&inst.Arrangement, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "System":
			FormDivBasicFieldToField(&inst.System, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Root":
			StageSetFormDivSelectFieldToField(&inst.Root, probe.stageOfInterest.GetInstancesSet[*models.Root](), formDiv)
		case "Numeral":
			StageSetFormDivSelectFieldToField(&inst.Numeral, probe.stageOfInterest.GetInstancesSet[*models.Numeral](), formDiv)
		case "Function":
			StageSetFormDivSelectFieldToField(&inst.Function, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), formDiv)
		case "Kind":
			StageSetFormDivSelectFieldToField(&inst.Kind, probe.stageOfInterest.GetInstancesSet[*models.Kind](), formDiv)
		case "Inversion":
			StageSetFormDivSelectFieldToField(&inst.Inversion, probe.stageOfInterest.GetInstancesSet[*models.Inversion](), formDiv)
		case "Bass":
			StageSetFormDivSelectFieldToField(&inst.Bass, probe.stageOfInterest.GetInstancesSet[*models.Bass](), formDiv)
		case "Degree":
			FormDivSliceOfPointersToField(inst, "Degree", &inst.Degree, formDiv, probe)
		case "Frame":
			StageSetFormDivSelectFieldToField(&inst.Frame, probe.stageOfInterest.GetInstancesSet[*models.Frame](), formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageOfInterest.GetInstancesSet[*models.Offset](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		}
	}
}

func saveStageSet_Harmony_alter_Stage(
	inst *models.Harmony_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Harp_pedals_Stage(
	inst *models.Harp_pedals,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Pedal_tuning":
			FormDivSliceOfPointersToField(inst, "Pedal_tuning", &inst.Pedal_tuning, formDiv, probe)
		}
	}
}

func saveStageSet_Heel_toe_Stage(
	inst *models.Heel_toe,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Substitution":
			StageSetFormDivEnumStringFieldToField(&inst.Substitution, formDiv)
		}
	}
}

func saveStageSet_Hole_Stage(
	inst *models.Hole,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Hole_type":
			FormDivBasicFieldToField(&inst.Hole_type, formDiv)
		case "Hole_closed":
			StageSetFormDivSelectFieldToField(&inst.Hole_closed, probe.stageOfInterest.GetInstancesSet[*models.Hole_closed](), formDiv)
		case "Hole_shape":
			FormDivBasicFieldToField(&inst.Hole_shape, formDiv)
		}
	}
}

func saveStageSet_Hole_closed_Stage(
	inst *models.Hole_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Horizontal_turn_Stage(
	inst *models.Horizontal_turn,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Slash":
			StageSetFormDivEnumStringFieldToField(&inst.Slash, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Start_note":
			StageSetFormDivEnumStringFieldToField(&inst.Start_note, formDiv)
		case "Trill_step":
			StageSetFormDivEnumStringFieldToField(&inst.Trill_step, formDiv)
		case "Two_note_turn":
			StageSetFormDivEnumStringFieldToField(&inst.Two_note_turn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&inst.Second_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		}
	}
}

func saveStageSet_Identification_Stage(
	inst *models.Identification,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Creator":
			FormDivSliceOfPointersToField(inst, "Creator", &inst.Creator, formDiv, probe)
		case "Rights":
			FormDivSliceOfPointersToField(inst, "Rights", &inst.Rights, formDiv, probe)
		case "Encoding":
			StageSetFormDivSelectFieldToField(&inst.Encoding, probe.stageOfInterest.GetInstancesSet[*models.Encoding](), formDiv)
		case "Source":
			FormDivBasicFieldToField(&inst.Source, formDiv)
		case "Relation":
			FormDivSliceOfPointersToField(inst, "Relation", &inst.Relation, formDiv, probe)
		case "Miscellaneous":
			StageSetFormDivSelectFieldToField(&inst.Miscellaneous, probe.stageOfInterest.GetInstancesSet[*models.Miscellaneous](), formDiv)
		}
	}
}

func saveStageSet_Image_Stage(
	inst *models.Image,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Source":
			FormDivBasicFieldToField(&inst.Source, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Instrument_Stage(
	inst *models.Instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Instrument_change_Stage(
	inst *models.Instrument_change,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Instrument_sound":
			FormDivBasicFieldToField(&inst.Instrument_sound, formDiv)
		case "Solo":
			FormDivBasicFieldToField(&inst.Solo, formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&inst.Ensemble, formDiv)
		case "Virtual_instrument":
			StageSetFormDivSelectFieldToField(&inst.Virtual_instrument, probe.stageOfInterest.GetInstancesSet[*models.Virtual_instrument](), formDiv)
		}
	}
}

func saveStageSet_Instrument_link_Stage(
	inst *models.Instrument_link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Interchangeable_Stage(
	inst *models.Interchangeable,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Separator":
			StageSetFormDivEnumStringFieldToField(&inst.Separator, formDiv)
		case "Time_relation":
			StageSetFormDivEnumStringFieldToField(&inst.Time_relation, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Beat_type":
			FormDivBasicFieldToField(&inst.Beat_type, formDiv)
		}
	}
}

func saveStageSet_Inversion_Stage(
	inst *models.Inversion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Key_Stage(
	inst *models.Key,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Cancel":
			StageSetFormDivSelectFieldToField(&inst.Cancel, probe.stageOfInterest.GetInstancesSet[*models.Cancel](), formDiv)
		case "Fifths":
			FormDivBasicFieldToField(&inst.Fifths, formDiv)
		case "Mode":
			FormDivBasicFieldToField(&inst.Mode, formDiv)
		case "Key_step":
			StageSetFormDivEnumStringFieldToField(&inst.Key_step, formDiv)
		case "Key_alter":
			FormDivBasicFieldToField(&inst.Key_alter, formDiv)
		case "Key_accidental":
			StageSetFormDivSelectFieldToField(&inst.Key_accidental, probe.stageOfInterest.GetInstancesSet[*models.Key_accidental](), formDiv)
		case "Key_octave":
			FormDivSliceOfPointersToField(inst, "Key_octave", &inst.Key_octave, formDiv, probe)
		}
	}
}

func saveStageSet_Key_accidental_Stage(
	inst *models.Key_accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Key_octave_Stage(
	inst *models.Key_octave,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Cancel":
			StageSetFormDivEnumStringFieldToField(&inst.Cancel, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Kind_Stage(
	inst *models.Kind,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Use_symbols":
			StageSetFormDivEnumStringFieldToField(&inst.Use_symbols, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Stack_degrees":
			StageSetFormDivEnumStringFieldToField(&inst.Stack_degrees, formDiv)
		case "Parentheses_degrees":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses_degrees, formDiv)
		case "Bracket_degrees":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket_degrees, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Level_Stage(
	inst *models.Level,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Reference_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Reference_XSD, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Line_detail_Stage(
	inst *models.Line_detail,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		}
	}
}

func saveStageSet_Line_width_Stage(
	inst *models.Line_width,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Link_Stage(
	inst *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Position":
			FormDivBasicFieldToField(&inst.Position, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		}
	}
}

func saveStageSet_Listen_Stage(
	inst *models.Listen,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Assess":
			FormDivSliceOfPointersToField(inst, "Assess", &inst.Assess, formDiv, probe)
		case "Wait":
			FormDivSliceOfPointersToField(inst, "Wait", &inst.Wait, formDiv, probe)
		case "Other_listen":
			FormDivSliceOfPointersToField(inst, "Other_listen", &inst.Other_listen, formDiv, probe)
		}
	}
}

func saveStageSet_Listening_Stage(
	inst *models.Listening,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Sync":
			FormDivSliceOfPointersToField(inst, "Sync", &inst.Sync, formDiv, probe)
		case "Other_listening":
			FormDivSliceOfPointersToField(inst, "Other_listening", &inst.Other_listening, formDiv, probe)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageOfInterest.GetInstancesSet[*models.Offset](), formDiv)
		}
	}
}

func saveStageSet_Lyric_Stage(
	inst *models.Lyric,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Syllabic":
			StageSetFormDivEnumStringFieldToField(&inst.Syllabic, formDiv)
		case "Text":
			StageSetFormDivSelectFieldToField(&inst.Text, probe.stageOfInterest.GetInstancesSet[*models.Text_element_data](), formDiv)
		case "Elision":
			FormDivSliceOfPointersToField(inst, "Elision", &inst.Elision, formDiv, probe)
		case "Extend":
			StageSetFormDivSelectFieldToField(&inst.Extend, probe.stageOfInterest.GetInstancesSet[*models.Extend](), formDiv)
		case "Laughing":
			FormDivBasicFieldToField(&inst.Laughing, formDiv)
		case "Humming":
			FormDivBasicFieldToField(&inst.Humming, formDiv)
		case "End_line":
			FormDivBasicFieldToField(&inst.End_line, formDiv)
		case "End_paragraph":
			FormDivBasicFieldToField(&inst.End_paragraph, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		}
	}
}

func saveStageSet_Lyric_font_Stage(
	inst *models.Lyric_font,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		}
	}
}

func saveStageSet_Lyric_language_Stage(
	inst *models.Lyric_language,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		}
	}
}

func saveStageSet_Measure_layout_Stage(
	inst *models.Measure_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Measure_distance":
			FormDivBasicFieldToField(&inst.Measure_distance, formDiv)
		}
	}
}

func saveStageSet_Measure_numbering_Stage(
	inst *models.Measure_numbering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "System":
			StageSetFormDivEnumStringFieldToField(&inst.System, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "Multiple_rest_always":
			StageSetFormDivEnumStringFieldToField(&inst.Multiple_rest_always, formDiv)
		case "Multiple_rest_range":
			StageSetFormDivEnumStringFieldToField(&inst.Multiple_rest_range, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Measure_repeat_Stage(
	inst *models.Measure_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&inst.Slashes, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Measure_style_Stage(
	inst *models.Measure_style,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Multiple_rest":
			StageSetFormDivSelectFieldToField(&inst.Multiple_rest, probe.stageOfInterest.GetInstancesSet[*models.Multiple_rest](), formDiv)
		case "Measure_repeat":
			StageSetFormDivSelectFieldToField(&inst.Measure_repeat, probe.stageOfInterest.GetInstancesSet[*models.Measure_repeat](), formDiv)
		case "Beat_repeat":
			StageSetFormDivSelectFieldToField(&inst.Beat_repeat, probe.stageOfInterest.GetInstancesSet[*models.Beat_repeat](), formDiv)
		case "Slash":
			StageSetFormDivSelectFieldToField(&inst.Slash, probe.stageOfInterest.GetInstancesSet[*models.Slash](), formDiv)
		}
	}
}

func saveStageSet_Membrane_Stage(
	inst *models.Membrane,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Metal_Stage(
	inst *models.Metal,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Metronome_Stage(
	inst *models.Metronome,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Beat_unit":
			StageSetFormDivEnumStringFieldToField(&inst.Beat_unit, formDiv)
		case "Beat_unit_dot":
			FormDivBasicFieldToField(&inst.Beat_unit_dot, formDiv)
		case "Beat_unit_tied":
			FormDivSliceOfPointersToField(inst, "Beat_unit_tied", &inst.Beat_unit_tied, formDiv, probe)
		case "Per_minute":
			StageSetFormDivSelectFieldToField(&inst.Per_minute, probe.stageOfInterest.GetInstancesSet[*models.Per_minute](), formDiv)
		case "Metronome_arrows":
			FormDivBasicFieldToField(&inst.Metronome_arrows, formDiv)
		case "Metronome_note":
			FormDivSliceOfPointersToField(inst, "Metronome_note", &inst.Metronome_note, formDiv, probe)
		case "Metronome_relation":
			FormDivBasicFieldToField(&inst.Metronome_relation, formDiv)
		}
	}
}

func saveStageSet_Metronome_beam_Stage(
	inst *models.Metronome_beam,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Metronome_note_Stage(
	inst *models.Metronome_note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Metronome_type":
			StageSetFormDivEnumStringFieldToField(&inst.Metronome_type, formDiv)
		case "Metronome_dot":
			FormDivBasicFieldToField(&inst.Metronome_dot, formDiv)
		case "Metronome_beam":
			FormDivSliceOfPointersToField(inst, "Metronome_beam", &inst.Metronome_beam, formDiv, probe)
		case "Metronome_tied":
			StageSetFormDivSelectFieldToField(&inst.Metronome_tied, probe.stageOfInterest.GetInstancesSet[*models.Metronome_tied](), formDiv)
		case "Metronome_tuplet":
			StageSetFormDivSelectFieldToField(&inst.Metronome_tuplet, probe.stageOfInterest.GetInstancesSet[*models.Metronome_tuplet](), formDiv)
		}
	}
}

func saveStageSet_Metronome_tied_Stage(
	inst *models.Metronome_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		}
	}
}

func saveStageSet_Metronome_tuplet_Stage(
	inst *models.Metronome_tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Actual_notes":
			FormDivBasicFieldToField(&inst.Actual_notes, formDiv)
		case "Normal_notes":
			FormDivBasicFieldToField(&inst.Normal_notes, formDiv)
		case "Normal_type":
			StageSetFormDivEnumStringFieldToField(&inst.Normal_type, formDiv)
		case "Normal_dot":
			FormDivBasicFieldToField(&inst.Normal_dot, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Show_number":
			StageSetFormDivEnumStringFieldToField(&inst.Show_number, formDiv)
		}
	}
}

func saveStageSet_Midi_device_Stage(
	inst *models.Midi_device,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Port":
			FormDivBasicFieldToField(&inst.Port, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Midi_instrument_Stage(
	inst *models.Midi_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Midi_channel":
			FormDivBasicFieldToField(&inst.Midi_channel, formDiv)
		case "Midi_name":
			FormDivBasicFieldToField(&inst.Midi_name, formDiv)
		case "Midi_bank":
			FormDivBasicFieldToField(&inst.Midi_bank, formDiv)
		case "Midi_program":
			FormDivBasicFieldToField(&inst.Midi_program, formDiv)
		case "Midi_unpitched":
			FormDivBasicFieldToField(&inst.Midi_unpitched, formDiv)
		case "Volume":
			FormDivBasicFieldToField(&inst.Volume, formDiv)
		case "Pan":
			FormDivBasicFieldToField(&inst.Pan, formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&inst.Elevation, formDiv)
		}
	}
}

func saveStageSet_Miscellaneous_Stage(
	inst *models.Miscellaneous,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Miscellaneous_field":
			FormDivSliceOfPointersToField(inst, "Miscellaneous_field", &inst.Miscellaneous_field, formDiv, probe)
		}
	}
}

func saveStageSet_Miscellaneous_field_Stage(
	inst *models.Miscellaneous_field,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&inst.NameXSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Mordent_Stage(
	inst *models.Mordent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Start_note":
			StageSetFormDivEnumStringFieldToField(&inst.Start_note, formDiv)
		case "Trill_step":
			StageSetFormDivEnumStringFieldToField(&inst.Trill_step, formDiv)
		case "Two_note_turn":
			StageSetFormDivEnumStringFieldToField(&inst.Two_note_turn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&inst.Second_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		case "Long":
			StageSetFormDivEnumStringFieldToField(&inst.Long, formDiv)
		case "Approach":
			StageSetFormDivEnumStringFieldToField(&inst.Approach, formDiv)
		case "Departure":
			StageSetFormDivEnumStringFieldToField(&inst.Departure, formDiv)
		}
	}
}

func saveStageSet_Multiple_rest_Stage(
	inst *models.Multiple_rest,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Use_symbols":
			StageSetFormDivEnumStringFieldToField(&inst.Use_symbols, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Name_display_Stage(
	inst *models.Name_display,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Display_text":
			FormDivSliceOfPointersToField(inst, "Display_text", &inst.Display_text, formDiv, probe)
		case "Accidental_text":
			FormDivSliceOfPointersToField(inst, "Accidental_text", &inst.Accidental_text, formDiv, probe)
		}
	}
}

func saveStageSet_Non_arpeggiate_Stage(
	inst *models.Non_arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Notations_Stage(
	inst *models.Notations,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Tied":
			FormDivSliceOfPointersToField(inst, "Tied", &inst.Tied, formDiv, probe)
		case "Slur":
			FormDivSliceOfPointersToField(inst, "Slur", &inst.Slur, formDiv, probe)
		case "Tuplet":
			FormDivSliceOfPointersToField(inst, "Tuplet", &inst.Tuplet, formDiv, probe)
		case "Glissando":
			FormDivSliceOfPointersToField(inst, "Glissando", &inst.Glissando, formDiv, probe)
		case "Slide":
			FormDivSliceOfPointersToField(inst, "Slide", &inst.Slide, formDiv, probe)
		case "Ornaments":
			FormDivSliceOfPointersToField(inst, "Ornaments", &inst.Ornaments, formDiv, probe)
		case "Technical":
			FormDivSliceOfPointersToField(inst, "Technical", &inst.Technical, formDiv, probe)
		case "Articulations":
			FormDivSliceOfPointersToField(inst, "Articulations", &inst.Articulations, formDiv, probe)
		case "Dynamics":
			FormDivSliceOfPointersToField(inst, "Dynamics", &inst.Dynamics, formDiv, probe)
		case "Fermata":
			FormDivSliceOfPointersToField(inst, "Fermata", &inst.Fermata, formDiv, probe)
		case "Arpeggiate":
			FormDivSliceOfPointersToField(inst, "Arpeggiate", &inst.Arpeggiate, formDiv, probe)
		case "Non_arpeggiate":
			FormDivSliceOfPointersToField(inst, "Non_arpeggiate", &inst.Non_arpeggiate, formDiv, probe)
		case "Accidental_mark":
			FormDivSliceOfPointersToField(inst, "Accidental_mark", &inst.Accidental_mark, formDiv, probe)
		case "Other_notation":
			FormDivSliceOfPointersToField(inst, "Other_notation", &inst.Other_notation, formDiv, probe)
		}
	}
}

func saveStageSet_Note_Stage(
	inst *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_leger":
			StageSetFormDivEnumStringFieldToField(&inst.Print_leger, formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&inst.Dynamics, formDiv)
		case "End_dynamics":
			FormDivBasicFieldToField(&inst.End_dynamics, formDiv)
		case "Attack":
			FormDivBasicFieldToField(&inst.Attack, formDiv)
		case "Release":
			FormDivBasicFieldToField(&inst.Release, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		case "Pizzicato":
			StageSetFormDivEnumStringFieldToField(&inst.Pizzicato, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Print_dot":
			StageSetFormDivEnumStringFieldToField(&inst.Print_dot, formDiv)
		case "Print_lyric":
			StageSetFormDivEnumStringFieldToField(&inst.Print_lyric, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Print_spacing":
			StageSetFormDivEnumStringFieldToField(&inst.Print_spacing, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Grace":
			StageSetFormDivSelectFieldToField(&inst.Grace, probe.stageOfInterest.GetInstancesSet[*models.Grace](), formDiv)
		case "Chord":
			FormDivBasicFieldToField(&inst.Chord, formDiv)
		case "Pitch":
			StageSetFormDivSelectFieldToField(&inst.Pitch, probe.stageOfInterest.GetInstancesSet[*models.Pitch](), formDiv)
		case "Unpitched":
			StageSetFormDivSelectFieldToField(&inst.Unpitched, probe.stageOfInterest.GetInstancesSet[*models.Unpitched](), formDiv)
		case "Rest":
			StageSetFormDivSelectFieldToField(&inst.Rest, probe.stageOfInterest.GetInstancesSet[*models.Rest](), formDiv)
		case "Tie":
			StageSetFormDivSelectFieldToField(&inst.Tie, probe.stageOfInterest.GetInstancesSet[*models.Tie](), formDiv)
		case "Cue":
			FormDivBasicFieldToField(&inst.Cue, formDiv)
		case "Instrument":
			FormDivSliceOfPointersToField(inst, "Instrument", &inst.Instrument, formDiv, probe)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.GetInstancesSet[*models.Note_type](), formDiv)
		case "Dot":
			FormDivSliceOfPointersToField(inst, "Dot", &inst.Dot, formDiv, probe)
		case "Accidental":
			StageSetFormDivSelectFieldToField(&inst.Accidental, probe.stageOfInterest.GetInstancesSet[*models.Accidental](), formDiv)
		case "Time_modification":
			StageSetFormDivSelectFieldToField(&inst.Time_modification, probe.stageOfInterest.GetInstancesSet[*models.Time_modification](), formDiv)
		case "Stem":
			StageSetFormDivSelectFieldToField(&inst.Stem, probe.stageOfInterest.GetInstancesSet[*models.Stem](), formDiv)
		case "Notehead":
			StageSetFormDivSelectFieldToField(&inst.Notehead, probe.stageOfInterest.GetInstancesSet[*models.Notehead](), formDiv)
		case "Notehead_text":
			StageSetFormDivSelectFieldToField(&inst.Notehead_text, probe.stageOfInterest.GetInstancesSet[*models.Notehead_text](), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "Beam":
			StageSetFormDivSelectFieldToField(&inst.Beam, probe.stageOfInterest.GetInstancesSet[*models.Beam](), formDiv)
		case "Notations":
			FormDivSliceOfPointersToField(inst, "Notations", &inst.Notations, formDiv, probe)
		case "Lyric":
			FormDivSliceOfPointersToField(inst, "Lyric", &inst.Lyric, formDiv, probe)
		case "Play":
			StageSetFormDivSelectFieldToField(&inst.Play, probe.stageOfInterest.GetInstancesSet[*models.Play](), formDiv)
		case "Listen":
			StageSetFormDivSelectFieldToField(&inst.Listen, probe.stageOfInterest.GetInstancesSet[*models.Listen](), formDiv)
		}
	}
}

func saveStageSet_Note_size_Stage(
	inst *models.Note_size,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Note_type_Stage(
	inst *models.Note_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Notehead_Stage(
	inst *models.Notehead,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Filled":
			StageSetFormDivEnumStringFieldToField(&inst.Filled, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Notehead_text_Stage(
	inst *models.Notehead_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Display_text":
			FormDivSliceOfPointersToField(inst, "Display_text", &inst.Display_text, formDiv, probe)
		case "Accidental_text":
			FormDivSliceOfPointersToField(inst, "Accidental_text", &inst.Accidental_text, formDiv, probe)
		}
	}
}

func saveStageSet_Numeral_Stage(
	inst *models.Numeral,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Numeral_root":
			StageSetFormDivSelectFieldToField(&inst.Numeral_root, probe.stageOfInterest.GetInstancesSet[*models.Numeral_root](), formDiv)
		case "Numeral_alter":
			StageSetFormDivSelectFieldToField(&inst.Numeral_alter, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), formDiv)
		case "Numeral_key":
			StageSetFormDivSelectFieldToField(&inst.Numeral_key, probe.stageOfInterest.GetInstancesSet[*models.Numeral_key](), formDiv)
		}
	}
}

func saveStageSet_Numeral_key_Stage(
	inst *models.Numeral_key,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Numeral_fifths":
			FormDivBasicFieldToField(&inst.Numeral_fifths, formDiv)
		case "Numeral_mode":
			StageSetFormDivEnumStringFieldToField(&inst.Numeral_mode, formDiv)
		}
	}
}

func saveStageSet_Numeral_root_Stage(
	inst *models.Numeral_root,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Octave_shift_Stage(
	inst *models.Octave_shift,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Size":
			FormDivBasicFieldToField(&inst.Size, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Offset_Stage(
	inst *models.Offset,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Sound":
			StageSetFormDivEnumStringFieldToField(&inst.Sound, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Opus_Stage(
	inst *models.Opus,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_Ornaments_Stage(
	inst *models.Ornaments,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Trill_mark":
			FormDivSliceOfPointersToField(inst, "Trill_mark", &inst.Trill_mark, formDiv, probe)
		case "Turn":
			FormDivSliceOfPointersToField(inst, "Turn", &inst.Turn, formDiv, probe)
		case "Delayed_turn":
			FormDivSliceOfPointersToField(inst, "Delayed_turn", &inst.Delayed_turn, formDiv, probe)
		case "Inverted_turn":
			FormDivSliceOfPointersToField(inst, "Inverted_turn", &inst.Inverted_turn, formDiv, probe)
		case "Delayed_inverted_turn":
			FormDivSliceOfPointersToField(inst, "Delayed_inverted_turn", &inst.Delayed_inverted_turn, formDiv, probe)
		case "Vertical_turn":
			FormDivSliceOfPointersToField(inst, "Vertical_turn", &inst.Vertical_turn, formDiv, probe)
		case "Inverted_vertical_turn":
			FormDivSliceOfPointersToField(inst, "Inverted_vertical_turn", &inst.Inverted_vertical_turn, formDiv, probe)
		case "Shake":
			FormDivSliceOfPointersToField(inst, "Shake", &inst.Shake, formDiv, probe)
		case "Wavy_line":
			FormDivSliceOfPointersToField(inst, "Wavy_line", &inst.Wavy_line, formDiv, probe)
		case "Mordent":
			FormDivSliceOfPointersToField(inst, "Mordent", &inst.Mordent, formDiv, probe)
		case "Inverted_mordent":
			FormDivSliceOfPointersToField(inst, "Inverted_mordent", &inst.Inverted_mordent, formDiv, probe)
		case "Schleifer":
			FormDivSliceOfPointersToField(inst, "Schleifer", &inst.Schleifer, formDiv, probe)
		case "Tremolo":
			FormDivSliceOfPointersToField(inst, "Tremolo", &inst.Tremolo, formDiv, probe)
		case "Haydn":
			FormDivSliceOfPointersToField(inst, "Haydn", &inst.Haydn, formDiv, probe)
		case "Other_ornament":
			FormDivSliceOfPointersToField(inst, "Other_ornament", &inst.Other_ornament, formDiv, probe)
		case "Accidental_mark":
			FormDivSliceOfPointersToField(inst, "Accidental_mark", &inst.Accidental_mark, formDiv, probe)
		}
	}
}

func saveStageSet_Other_appearance_Stage(
	inst *models.Other_appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_direction_Stage(
	inst *models.Other_direction,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_listening_Stage(
	inst *models.Other_listening,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_notation_Stage(
	inst *models.Other_notation,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_placement_text_Stage(
	inst *models.Other_placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_play_Stage(
	inst *models.Other_play,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Other_text_Stage(
	inst *models.Other_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Page_layout_Stage(
	inst *models.Page_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Page_height":
			FormDivBasicFieldToField(&inst.Page_height, formDiv)
		case "Page_width":
			FormDivBasicFieldToField(&inst.Page_width, formDiv)
		case "Page_margins":
			StageSetFormDivSelectFieldToField(&inst.Page_margins, probe.stageOfInterest.GetInstancesSet[*models.Page_margins](), formDiv)
		}
	}
}

func saveStageSet_Page_margins_Stage(
	inst *models.Page_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Left_margin":
			FormDivBasicFieldToField(&inst.Left_margin, formDiv)
		case "Right_margin":
			FormDivBasicFieldToField(&inst.Right_margin, formDiv)
		case "Top_margin":
			FormDivBasicFieldToField(&inst.Top_margin, formDiv)
		case "Bottom_margin":
			FormDivBasicFieldToField(&inst.Bottom_margin, formDiv)
		}
	}
}

func saveStageSet_Part_clef_Stage(
	inst *models.Part_clef,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "Clef_octave_change":
			FormDivBasicFieldToField(&inst.Clef_octave_change, formDiv)
		}
	}
}

func saveStageSet_Part_group_Stage(
	inst *models.Part_group,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Group_name":
			StageSetFormDivSelectFieldToField(&inst.Group_name, probe.stageOfInterest.GetInstancesSet[*models.Group_name](), formDiv)
		case "Group_name_display":
			StageSetFormDivSelectFieldToField(&inst.Group_name_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		case "Group_abbreviation":
			StageSetFormDivSelectFieldToField(&inst.Group_abbreviation, probe.stageOfInterest.GetInstancesSet[*models.Group_name](), formDiv)
		case "Group_abbreviation_display":
			StageSetFormDivSelectFieldToField(&inst.Group_abbreviation_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		case "Group_symbol":
			StageSetFormDivSelectFieldToField(&inst.Group_symbol, probe.stageOfInterest.GetInstancesSet[*models.Group_symbol](), formDiv)
		case "Group_barline":
			StageSetFormDivSelectFieldToField(&inst.Group_barline, probe.stageOfInterest.GetInstancesSet[*models.Group_barline](), formDiv)
		case "Group_time":
			FormDivBasicFieldToField(&inst.Group_time, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.GetInstancesSet[*models.Level](), formDiv)
		}
	}
}

func saveStageSet_Part_link_Stage(
	inst *models.Part_link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		case "Instrument_link":
			FormDivSliceOfPointersToField(inst, "Instrument_link", &inst.Instrument_link, formDiv, probe)
		case "Group_link":
			FormDivBasicFieldToField(&inst.Group_link, formDiv)
		}
	}
}

func saveStageSet_Part_list_Stage(
	inst *models.Part_list,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Part_group":
			StageSetFormDivSelectFieldToField(&inst.Part_group, probe.stageOfInterest.GetInstancesSet[*models.Part_group](), formDiv)
		case "Score_part":
			StageSetFormDivSelectFieldToField(&inst.Score_part, probe.stageOfInterest.GetInstancesSet[*models.Score_part](), formDiv)
		}
	}
}

func saveStageSet_Part_name_Stage(
	inst *models.Part_name,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Part_symbol_Stage(
	inst *models.Part_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Top_staff":
			FormDivBasicFieldToField(&inst.Top_staff, formDiv)
		case "Bottom_staff":
			FormDivBasicFieldToField(&inst.Bottom_staff, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Part_transpose_Stage(
	inst *models.Part_transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&inst.Diatonic, formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&inst.Chromatic, formDiv)
		case "Octave_change":
			FormDivBasicFieldToField(&inst.Octave_change, formDiv)
		case "Double":
			FormDivBasicFieldToField(&inst.Double, formDiv)
		}
	}
}

func saveStageSet_Pedal_Stage(
	inst *models.Pedal,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line":
			StageSetFormDivEnumStringFieldToField(&inst.Line, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Abbreviated":
			StageSetFormDivEnumStringFieldToField(&inst.Abbreviated, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Pedal_tuning_Stage(
	inst *models.Pedal_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Pedal_step":
			StageSetFormDivEnumStringFieldToField(&inst.Pedal_step, formDiv)
		case "Pedal_alter":
			FormDivBasicFieldToField(&inst.Pedal_alter, formDiv)
		}
	}
}

func saveStageSet_Per_minute_Stage(
	inst *models.Per_minute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Percussion_Stage(
	inst *models.Percussion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Glass":
			StageSetFormDivSelectFieldToField(&inst.Glass, probe.stageOfInterest.GetInstancesSet[*models.Glass](), formDiv)
		case "Metal":
			StageSetFormDivSelectFieldToField(&inst.Metal, probe.stageOfInterest.GetInstancesSet[*models.Metal](), formDiv)
		case "Wood":
			StageSetFormDivSelectFieldToField(&inst.Wood, probe.stageOfInterest.GetInstancesSet[*models.Wood](), formDiv)
		case "Pitched":
			StageSetFormDivSelectFieldToField(&inst.Pitched, probe.stageOfInterest.GetInstancesSet[*models.Pitched](), formDiv)
		case "Membrane":
			StageSetFormDivSelectFieldToField(&inst.Membrane, probe.stageOfInterest.GetInstancesSet[*models.Membrane](), formDiv)
		case "Effect":
			StageSetFormDivSelectFieldToField(&inst.Effect, probe.stageOfInterest.GetInstancesSet[*models.Effect](), formDiv)
		case "Timpani":
			StageSetFormDivSelectFieldToField(&inst.Timpani, probe.stageOfInterest.GetInstancesSet[*models.Timpani](), formDiv)
		case "Beater":
			StageSetFormDivSelectFieldToField(&inst.Beater, probe.stageOfInterest.GetInstancesSet[*models.Beater](), formDiv)
		case "Stick":
			StageSetFormDivSelectFieldToField(&inst.Stick, probe.stageOfInterest.GetInstancesSet[*models.Stick](), formDiv)
		case "Stick_location":
			StageSetFormDivEnumStringFieldToField(&inst.Stick_location, formDiv)
		case "Other_percussion":
			StageSetFormDivSelectFieldToField(&inst.Other_percussion, probe.stageOfInterest.GetInstancesSet[*models.Other_text](), formDiv)
		}
	}
}

func saveStageSet_Pitch_Stage(
	inst *models.Pitch,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Step":
			StageSetFormDivEnumStringFieldToField(&inst.Step, formDiv)
		case "Alter":
			FormDivBasicFieldToField(&inst.Alter, formDiv)
		case "Octave":
			FormDivBasicFieldToField(&inst.Octave, formDiv)
		}
	}
}

func saveStageSet_Pitched_Stage(
	inst *models.Pitched,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Placement_text_Stage(
	inst *models.Placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Play_Stage(
	inst *models.Play,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Ipa":
			FormDivBasicFieldToField(&inst.Ipa, formDiv)
		case "Mute":
			StageSetFormDivEnumStringFieldToField(&inst.Mute, formDiv)
		case "Semi_pitched":
			StageSetFormDivEnumStringFieldToField(&inst.Semi_pitched, formDiv)
		case "Other_play":
			FormDivSliceOfPointersToField(inst, "Other_play", &inst.Other_play, formDiv, probe)
		}
	}
}

func saveStageSet_Player_Stage(
	inst *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Player_name":
			FormDivBasicFieldToField(&inst.Player_name, formDiv)
		}
	}
}

func saveStageSet_Principal_voice_Stage(
	inst *models.Principal_voice,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Print_XSD_Stage(
	inst *models.Print_XSD,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Staff_spacing":
			FormDivBasicFieldToField(&inst.Staff_spacing, formDiv)
		case "New_system":
			StageSetFormDivEnumStringFieldToField(&inst.New_system, formDiv)
		case "New_page":
			StageSetFormDivEnumStringFieldToField(&inst.New_page, formDiv)
		case "Blank_page":
			FormDivBasicFieldToField(&inst.Blank_page, formDiv)
		case "Page_number":
			FormDivBasicFieldToField(&inst.Page_number, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Page_layout":
			StageSetFormDivSelectFieldToField(&inst.Page_layout, probe.stageOfInterest.GetInstancesSet[*models.Page_layout](), formDiv)
		case "System_layout":
			StageSetFormDivSelectFieldToField(&inst.System_layout, probe.stageOfInterest.GetInstancesSet[*models.System_layout](), formDiv)
		case "Staff_layout":
			FormDivSliceOfPointersToField(inst, "Staff_layout", &inst.Staff_layout, formDiv, probe)
		case "Measure_layout":
			StageSetFormDivSelectFieldToField(&inst.Measure_layout, probe.stageOfInterest.GetInstancesSet[*models.Measure_layout](), formDiv)
		case "Measure_numbering":
			StageSetFormDivSelectFieldToField(&inst.Measure_numbering, probe.stageOfInterest.GetInstancesSet[*models.Measure_numbering](), formDiv)
		case "Part_name_display":
			StageSetFormDivSelectFieldToField(&inst.Part_name_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		case "Part_abbreviation_display":
			StageSetFormDivSelectFieldToField(&inst.Part_abbreviation_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		}
	}
}

func saveStageSet_Release_Stage(
	inst *models.Release,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Offset":
			FormDivBasicFieldToField(&inst.Offset, formDiv)
		}
	}
}

func saveStageSet_Repeat_Stage(
	inst *models.Repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Direction":
			StageSetFormDivEnumStringFieldToField(&inst.Direction, formDiv)
		case "Times":
			FormDivBasicFieldToField(&inst.Times, formDiv)
		case "After_jump":
			StageSetFormDivEnumStringFieldToField(&inst.After_jump, formDiv)
		case "Winged":
			StageSetFormDivEnumStringFieldToField(&inst.Winged, formDiv)
		}
	}
}

func saveStageSet_Rest_Stage(
	inst *models.Rest,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Measure":
			StageSetFormDivEnumStringFieldToField(&inst.Measure, formDiv)
		case "Display_step":
			StageSetFormDivEnumStringFieldToField(&inst.Display_step, formDiv)
		case "Display_octave":
			FormDivBasicFieldToField(&inst.Display_octave, formDiv)
		}
	}
}

func saveStageSet_Root_Stage(
	inst *models.Root,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Root_step":
			StageSetFormDivSelectFieldToField(&inst.Root_step, probe.stageOfInterest.GetInstancesSet[*models.Root_step](), formDiv)
		case "Root_alter":
			StageSetFormDivSelectFieldToField(&inst.Root_alter, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), formDiv)
		}
	}
}

func saveStageSet_Root_step_Stage(
	inst *models.Root_step,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Scaling_Stage(
	inst *models.Scaling,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Millimeters":
			FormDivBasicFieldToField(&inst.Millimeters, formDiv)
		case "Tenths":
			FormDivBasicFieldToField(&inst.Tenths, formDiv)
		}
	}
}

func saveStageSet_Scordatura_Stage(
	inst *models.Scordatura,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Accord":
			FormDivSliceOfPointersToField(inst, "Accord", &inst.Accord, formDiv, probe)
		}
	}
}

func saveStageSet_Score_instrument_Stage(
	inst *models.Score_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Instrument_name":
			FormDivBasicFieldToField(&inst.Instrument_name, formDiv)
		case "Instrument_abbreviation":
			FormDivBasicFieldToField(&inst.Instrument_abbreviation, formDiv)
		case "Instrument_sound":
			FormDivBasicFieldToField(&inst.Instrument_sound, formDiv)
		case "Solo":
			FormDivBasicFieldToField(&inst.Solo, formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&inst.Ensemble, formDiv)
		case "Virtual_instrument":
			StageSetFormDivSelectFieldToField(&inst.Virtual_instrument, probe.stageOfInterest.GetInstancesSet[*models.Virtual_instrument](), formDiv)
		}
	}
}

func saveStageSet_Score_part_Stage(
	inst *models.Score_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageOfInterest.GetInstancesSet[*models.Identification](), formDiv)
		case "Part_link":
			FormDivSliceOfPointersToField(inst, "Part_link", &inst.Part_link, formDiv, probe)
		case "Part_name":
			StageSetFormDivSelectFieldToField(&inst.Part_name, probe.stageOfInterest.GetInstancesSet[*models.Part_name](), formDiv)
		case "Part_name_display":
			StageSetFormDivSelectFieldToField(&inst.Part_name_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		case "Part_abbreviation":
			StageSetFormDivSelectFieldToField(&inst.Part_abbreviation, probe.stageOfInterest.GetInstancesSet[*models.Part_name](), formDiv)
		case "Part_abbreviation_display":
			StageSetFormDivSelectFieldToField(&inst.Part_abbreviation_display, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), formDiv)
		case "Group":
			FormDivBasicFieldToField(&inst.Group, formDiv)
		case "Score_instrument":
			FormDivSliceOfPointersToField(inst, "Score_instrument", &inst.Score_instrument, formDiv, probe)
		case "Player":
			FormDivSliceOfPointersToField(inst, "Player", &inst.Player, formDiv, probe)
		case "Midi_device":
			FormDivSliceOfPointersToField(inst, "Midi_device", &inst.Midi_device, formDiv, probe)
		case "Midi_instrument":
			FormDivSliceOfPointersToField(inst, "Midi_instrument", &inst.Midi_instrument, formDiv, probe)
		}
	}
}

func saveStageSet_Score_partwise_Stage(
	inst *models.Score_partwise,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Version":
			FormDivBasicFieldToField(&inst.Version, formDiv)
		case "Work":
			StageSetFormDivSelectFieldToField(&inst.Work, probe.stageOfInterest.GetInstancesSet[*models.Work](), formDiv)
		case "Movement_number":
			FormDivBasicFieldToField(&inst.Movement_number, formDiv)
		case "Movement_title":
			FormDivBasicFieldToField(&inst.Movement_title, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageOfInterest.GetInstancesSet[*models.Identification](), formDiv)
		case "Defaults":
			StageSetFormDivSelectFieldToField(&inst.Defaults, probe.stageOfInterest.GetInstancesSet[*models.Defaults](), formDiv)
		case "Credit":
			FormDivSliceOfPointersToField(inst, "Credit", &inst.Credit, formDiv, probe)
		case "Part_list":
			StageSetFormDivSelectFieldToField(&inst.Part_list, probe.stageOfInterest.GetInstancesSet[*models.Part_list](), formDiv)
		case "Part":
			FormDivSliceOfPointersToField(inst, "Part", &inst.Part, formDiv, probe)
		}
	}
}

func saveStageSet_Score_timewise_Stage(
	inst *models.Score_timewise,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Version":
			FormDivBasicFieldToField(&inst.Version, formDiv)
		case "Work":
			StageSetFormDivSelectFieldToField(&inst.Work, probe.stageOfInterest.GetInstancesSet[*models.Work](), formDiv)
		case "Movement_number":
			FormDivBasicFieldToField(&inst.Movement_number, formDiv)
		case "Movement_title":
			FormDivBasicFieldToField(&inst.Movement_title, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageOfInterest.GetInstancesSet[*models.Identification](), formDiv)
		case "Defaults":
			StageSetFormDivSelectFieldToField(&inst.Defaults, probe.stageOfInterest.GetInstancesSet[*models.Defaults](), formDiv)
		case "Credit":
			FormDivSliceOfPointersToField(inst, "Credit", &inst.Credit, formDiv, probe)
		case "Part_list":
			StageSetFormDivSelectFieldToField(&inst.Part_list, probe.stageOfInterest.GetInstancesSet[*models.Part_list](), formDiv)
		case "Measure":
			FormDivSliceOfPointersToField(inst, "Measure", &inst.Measure, formDiv, probe)
		}
	}
}

func saveStageSet_Segno_Stage(
	inst *models.Segno,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Slash_Stage(
	inst *models.Slash,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Use_dots":
			StageSetFormDivEnumStringFieldToField(&inst.Use_dots, formDiv)
		case "Use_stems":
			StageSetFormDivEnumStringFieldToField(&inst.Use_stems, formDiv)
		case "Slash_type":
			StageSetFormDivEnumStringFieldToField(&inst.Slash_type, formDiv)
		case "Slash_dot":
			FormDivBasicFieldToField(&inst.Slash_dot, formDiv)
		case "Except_voice":
			FormDivBasicFieldToField(&inst.Except_voice, formDiv)
		}
	}
}

func saveStageSet_Slide_Stage(
	inst *models.Slide,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "First_beat":
			FormDivBasicFieldToField(&inst.First_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Slur_Stage(
	inst *models.Slur,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Orientation":
			StageSetFormDivEnumStringFieldToField(&inst.Orientation, formDiv)
		case "Bezier_x":
			FormDivBasicFieldToField(&inst.Bezier_x, formDiv)
		case "Bezier_y":
			FormDivBasicFieldToField(&inst.Bezier_y, formDiv)
		case "Bezier_x2":
			FormDivBasicFieldToField(&inst.Bezier_x2, formDiv)
		case "Bezier_y2":
			FormDivBasicFieldToField(&inst.Bezier_y2, formDiv)
		case "Bezier_offset":
			FormDivBasicFieldToField(&inst.Bezier_offset, formDiv)
		case "Bezier_offset2":
			FormDivBasicFieldToField(&inst.Bezier_offset2, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Sound_Stage(
	inst *models.Sound,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tempo":
			FormDivBasicFieldToField(&inst.Tempo, formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&inst.Dynamics, formDiv)
		case "Dacapo":
			StageSetFormDivEnumStringFieldToField(&inst.Dacapo, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&inst.Segno, formDiv)
		case "Dalsegno":
			FormDivBasicFieldToField(&inst.Dalsegno, formDiv)
		case "Coda":
			FormDivBasicFieldToField(&inst.Coda, formDiv)
		case "Tocoda":
			FormDivBasicFieldToField(&inst.Tocoda, formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "Forward_repeat":
			StageSetFormDivEnumStringFieldToField(&inst.Forward_repeat, formDiv)
		case "Fine":
			FormDivBasicFieldToField(&inst.Fine, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		case "Pizzicato":
			StageSetFormDivEnumStringFieldToField(&inst.Pizzicato, formDiv)
		case "Pan":
			FormDivBasicFieldToField(&inst.Pan, formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&inst.Elevation, formDiv)
		case "Damper_pedal":
			FormDivBasicFieldToField(&inst.Damper_pedal, formDiv)
		case "Soft_pedal":
			FormDivBasicFieldToField(&inst.Soft_pedal, formDiv)
		case "Sostenuto_pedal":
			FormDivBasicFieldToField(&inst.Sostenuto_pedal, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Instrument_change":
			FormDivSliceOfPointersToField(inst, "Instrument_change", &inst.Instrument_change, formDiv, probe)
		case "Midi_device":
			FormDivSliceOfPointersToField(inst, "Midi_device", &inst.Midi_device, formDiv, probe)
		case "Midi_instrument":
			FormDivSliceOfPointersToField(inst, "Midi_instrument", &inst.Midi_instrument, formDiv, probe)
		case "Play":
			FormDivSliceOfPointersToField(inst, "Play", &inst.Play, formDiv, probe)
		case "Swing":
			StageSetFormDivSelectFieldToField(&inst.Swing, probe.stageOfInterest.GetInstancesSet[*models.Swing](), formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageOfInterest.GetInstancesSet[*models.Offset](), formDiv)
		}
	}
}

func saveStageSet_Staff_details_Stage(
	inst *models.Staff_details,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Show_frets":
			StageSetFormDivEnumStringFieldToField(&inst.Show_frets, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Print_spacing":
			StageSetFormDivEnumStringFieldToField(&inst.Print_spacing, formDiv)
		case "Staff_type":
			StageSetFormDivEnumStringFieldToField(&inst.Staff_type, formDiv)
		case "Staff_lines":
			FormDivBasicFieldToField(&inst.Staff_lines, formDiv)
		case "Line_detail":
			FormDivSliceOfPointersToField(inst, "Line_detail", &inst.Line_detail, formDiv, probe)
		case "Staff_tuning":
			FormDivSliceOfPointersToField(inst, "Staff_tuning", &inst.Staff_tuning, formDiv, probe)
		case "Capo":
			FormDivBasicFieldToField(&inst.Capo, formDiv)
		case "Staff_size":
			StageSetFormDivSelectFieldToField(&inst.Staff_size, probe.stageOfInterest.GetInstancesSet[*models.Staff_size](), formDiv)
		}
	}
}

func saveStageSet_Staff_divide_Stage(
	inst *models.Staff_divide,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Staff_layout_Stage(
	inst *models.Staff_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Staff_distance":
			FormDivBasicFieldToField(&inst.Staff_distance, formDiv)
		}
	}
}

func saveStageSet_Staff_size_Stage(
	inst *models.Staff_size,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Scaling":
			FormDivBasicFieldToField(&inst.Scaling, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Staff_tuning_Stage(
	inst *models.Staff_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "Tuning_step":
			StageSetFormDivEnumStringFieldToField(&inst.Tuning_step, formDiv)
		case "Tuning_alter":
			FormDivBasicFieldToField(&inst.Tuning_alter, formDiv)
		case "Tuning_octave":
			FormDivBasicFieldToField(&inst.Tuning_octave, formDiv)
		}
	}
}

func saveStageSet_Stem_Stage(
	inst *models.Stem,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Stick_Stage(
	inst *models.Stick,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tip":
			StageSetFormDivEnumStringFieldToField(&inst.Tip, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Dashed_circle":
			StageSetFormDivEnumStringFieldToField(&inst.Dashed_circle, formDiv)
		case "Stick_type":
			StageSetFormDivEnumStringFieldToField(&inst.Stick_type, formDiv)
		case "Stick_material":
			StageSetFormDivEnumStringFieldToField(&inst.Stick_material, formDiv)
		}
	}
}

func saveStageSet_String_mute_Stage(
	inst *models.String_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_String_type_Stage(
	inst *models.String_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Strong_accent_Stage(
	inst *models.Strong_accent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		}
	}
}

func saveStageSet_Style_text_Stage(
	inst *models.Style_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Supports_Stage(
	inst *models.Supports,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Attribute":
			FormDivBasicFieldToField(&inst.Attribute, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Swing_Stage(
	inst *models.Swing,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Straight":
			FormDivBasicFieldToField(&inst.Straight, formDiv)
		case "First":
			FormDivBasicFieldToField(&inst.First, formDiv)
		case "Second":
			FormDivBasicFieldToField(&inst.Second, formDiv)
		case "Swing_type":
			FormDivBasicFieldToField(&inst.Swing_type, formDiv)
		case "Swing_style":
			FormDivBasicFieldToField(&inst.Swing_style, formDiv)
		}
	}
}

func saveStageSet_Sync_Stage(
	inst *models.Sync,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Latency":
			FormDivBasicFieldToField(&inst.Latency, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		}
	}
}

func saveStageSet_System_dividers_Stage(
	inst *models.System_dividers,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Left_divider":
			StageSetFormDivSelectFieldToField(&inst.Left_divider, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_object_style_align](), formDiv)
		case "Right_divider":
			StageSetFormDivSelectFieldToField(&inst.Right_divider, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_object_style_align](), formDiv)
		}
	}
}

func saveStageSet_System_layout_Stage(
	inst *models.System_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "System_margins":
			StageSetFormDivSelectFieldToField(&inst.System_margins, probe.stageOfInterest.GetInstancesSet[*models.System_margins](), formDiv)
		case "System_distance":
			FormDivBasicFieldToField(&inst.System_distance, formDiv)
		case "Top_system_distance":
			FormDivBasicFieldToField(&inst.Top_system_distance, formDiv)
		case "System_dividers":
			StageSetFormDivSelectFieldToField(&inst.System_dividers, probe.stageOfInterest.GetInstancesSet[*models.System_dividers](), formDiv)
		}
	}
}

func saveStageSet_System_margins_Stage(
	inst *models.System_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Left_margin":
			FormDivBasicFieldToField(&inst.Left_margin, formDiv)
		case "Right_margin":
			FormDivBasicFieldToField(&inst.Right_margin, formDiv)
		}
	}
}

func saveStageSet_Tap_Stage(
	inst *models.Tap,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Hand":
			StageSetFormDivEnumStringFieldToField(&inst.Hand, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Technical_Stage(
	inst *models.Technical,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Up_bow":
			FormDivSliceOfPointersToField(inst, "Up_bow", &inst.Up_bow, formDiv, probe)
		case "Down_bow":
			FormDivSliceOfPointersToField(inst, "Down_bow", &inst.Down_bow, formDiv, probe)
		case "Harmonic":
			FormDivSliceOfPointersToField(inst, "Harmonic", &inst.Harmonic, formDiv, probe)
		case "Open_string":
			FormDivSliceOfPointersToField(inst, "Open_string", &inst.Open_string, formDiv, probe)
		case "Thumb_position":
			FormDivSliceOfPointersToField(inst, "Thumb_position", &inst.Thumb_position, formDiv, probe)
		case "Fingering":
			FormDivSliceOfPointersToField(inst, "Fingering", &inst.Fingering, formDiv, probe)
		case "Pluck":
			FormDivSliceOfPointersToField(inst, "Pluck", &inst.Pluck, formDiv, probe)
		case "Double_tongue":
			FormDivSliceOfPointersToField(inst, "Double_tongue", &inst.Double_tongue, formDiv, probe)
		case "Triple_tongue":
			FormDivSliceOfPointersToField(inst, "Triple_tongue", &inst.Triple_tongue, formDiv, probe)
		case "Stopped":
			FormDivSliceOfPointersToField(inst, "Stopped", &inst.Stopped, formDiv, probe)
		case "Snap_pizzicato":
			FormDivSliceOfPointersToField(inst, "Snap_pizzicato", &inst.Snap_pizzicato, formDiv, probe)
		case "Fret":
			FormDivSliceOfPointersToField(inst, "Fret", &inst.Fret, formDiv, probe)
		case "String_XSD":
			FormDivSliceOfPointersToField(inst, "String_XSD", &inst.String_XSD, formDiv, probe)
		case "Hammer_on":
			FormDivSliceOfPointersToField(inst, "Hammer_on", &inst.Hammer_on, formDiv, probe)
		case "Pull_off":
			FormDivSliceOfPointersToField(inst, "Pull_off", &inst.Pull_off, formDiv, probe)
		case "Bend":
			FormDivSliceOfPointersToField(inst, "Bend", &inst.Bend, formDiv, probe)
		case "Tap":
			FormDivSliceOfPointersToField(inst, "Tap", &inst.Tap, formDiv, probe)
		case "Heel":
			FormDivSliceOfPointersToField(inst, "Heel", &inst.Heel, formDiv, probe)
		case "Toe":
			FormDivSliceOfPointersToField(inst, "Toe", &inst.Toe, formDiv, probe)
		case "Fingernails":
			FormDivSliceOfPointersToField(inst, "Fingernails", &inst.Fingernails, formDiv, probe)
		case "Hole":
			FormDivSliceOfPointersToField(inst, "Hole", &inst.Hole, formDiv, probe)
		case "Arrow":
			FormDivSliceOfPointersToField(inst, "Arrow", &inst.Arrow, formDiv, probe)
		case "Handbell":
			FormDivSliceOfPointersToField(inst, "Handbell", &inst.Handbell, formDiv, probe)
		case "Brass_bend":
			FormDivSliceOfPointersToField(inst, "Brass_bend", &inst.Brass_bend, formDiv, probe)
		case "Flip":
			FormDivSliceOfPointersToField(inst, "Flip", &inst.Flip, formDiv, probe)
		case "Smear":
			FormDivSliceOfPointersToField(inst, "Smear", &inst.Smear, formDiv, probe)
		case "Open":
			FormDivSliceOfPointersToField(inst, "Open", &inst.Open, formDiv, probe)
		case "Half_muted":
			FormDivSliceOfPointersToField(inst, "Half_muted", &inst.Half_muted, formDiv, probe)
		case "Harmon_mute":
			FormDivSliceOfPointersToField(inst, "Harmon_mute", &inst.Harmon_mute, formDiv, probe)
		case "Golpe":
			FormDivSliceOfPointersToField(inst, "Golpe", &inst.Golpe, formDiv, probe)
		case "Other_technical":
			FormDivSliceOfPointersToField(inst, "Other_technical", &inst.Other_technical, formDiv, probe)
		}
	}
}

func saveStageSet_Text_element_data_Stage(
	inst *models.Text_element_data,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&inst.Line_through, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&inst.Letter_spacing, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Tie_Stage(
	inst *models.Tie,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		}
	}
}

func saveStageSet_Tied_Stage(
	inst *models.Tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Orientation":
			StageSetFormDivEnumStringFieldToField(&inst.Orientation, formDiv)
		case "Bezier_x":
			FormDivBasicFieldToField(&inst.Bezier_x, formDiv)
		case "Bezier_y":
			FormDivBasicFieldToField(&inst.Bezier_y, formDiv)
		case "Bezier_x2":
			FormDivBasicFieldToField(&inst.Bezier_x2, formDiv)
		case "Bezier_y2":
			FormDivBasicFieldToField(&inst.Bezier_y2, formDiv)
		case "Bezier_offset":
			FormDivBasicFieldToField(&inst.Bezier_offset, formDiv)
		case "Bezier_offset2":
			FormDivBasicFieldToField(&inst.Bezier_offset2, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Time_Stage(
	inst *models.Time,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Separator":
			StageSetFormDivEnumStringFieldToField(&inst.Separator, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Print_object":
			StageSetFormDivEnumStringFieldToField(&inst.Print_object, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Beat_type":
			FormDivBasicFieldToField(&inst.Beat_type, formDiv)
		case "Interchangeable":
			StageSetFormDivSelectFieldToField(&inst.Interchangeable, probe.stageOfInterest.GetInstancesSet[*models.Interchangeable](), formDiv)
		case "Senza_misura":
			FormDivBasicFieldToField(&inst.Senza_misura, formDiv)
		}
	}
}

func saveStageSet_Time_modification_Stage(
	inst *models.Time_modification,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Actual_notes":
			FormDivBasicFieldToField(&inst.Actual_notes, formDiv)
		case "Normal_notes":
			FormDivBasicFieldToField(&inst.Normal_notes, formDiv)
		case "Normal_type":
			StageSetFormDivEnumStringFieldToField(&inst.Normal_type, formDiv)
		case "Normal_dot":
			FormDivBasicFieldToField(&inst.Normal_dot, formDiv)
		}
	}
}

func saveStageSet_Timpani_Stage(
	inst *models.Timpani,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		}
	}
}

func saveStageSet_Transpose_Stage(
	inst *models.Transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&inst.Diatonic, formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&inst.Chromatic, formDiv)
		case "Octave_change":
			FormDivBasicFieldToField(&inst.Octave_change, formDiv)
		case "Double":
			FormDivBasicFieldToField(&inst.Double, formDiv)
		}
	}
}

func saveStageSet_Tremolo_Stage(
	inst *models.Tremolo,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Tuplet_Stage(
	inst *models.Tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Show_number":
			StageSetFormDivEnumStringFieldToField(&inst.Show_number, formDiv)
		case "Show_type":
			StageSetFormDivEnumStringFieldToField(&inst.Show_type, formDiv)
		case "Line_shape":
			StageSetFormDivEnumStringFieldToField(&inst.Line_shape, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Tuplet_actual":
			StageSetFormDivSelectFieldToField(&inst.Tuplet_actual, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_portion](), formDiv)
		case "Tuplet_normal":
			StageSetFormDivSelectFieldToField(&inst.Tuplet_normal, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_portion](), formDiv)
		}
	}
}

func saveStageSet_Tuplet_dot_Stage(
	inst *models.Tuplet_dot,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Tuplet_number_Stage(
	inst *models.Tuplet_number,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Tuplet_portion_Stage(
	inst *models.Tuplet_portion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tuplet_number":
			StageSetFormDivSelectFieldToField(&inst.Tuplet_number, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_number](), formDiv)
		case "Tuplet_type":
			StageSetFormDivSelectFieldToField(&inst.Tuplet_type, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_type](), formDiv)
		case "Tuplet_dot":
			FormDivSliceOfPointersToField(inst, "Tuplet_dot", &inst.Tuplet_dot, formDiv, probe)
		}
	}
}

func saveStageSet_Tuplet_type_Stage(
	inst *models.Tuplet_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&inst.Font_family, formDiv)
		case "Font_style":
			StageSetFormDivEnumStringFieldToField(&inst.Font_style, formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&inst.Font_size, formDiv)
		case "Font_weight":
			StageSetFormDivEnumStringFieldToField(&inst.Font_weight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Typed_text_Stage(
	inst *models.Typed_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Unpitched_Stage(
	inst *models.Unpitched,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Display_step":
			StageSetFormDivEnumStringFieldToField(&inst.Display_step, formDiv)
		case "Display_octave":
			FormDivBasicFieldToField(&inst.Display_octave, formDiv)
		}
	}
}

func saveStageSet_Virtual_instrument_Stage(
	inst *models.Virtual_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Virtual_library":
			FormDivBasicFieldToField(&inst.Virtual_library, formDiv)
		case "Virtual_name":
			FormDivBasicFieldToField(&inst.Virtual_name, formDiv)
		}
	}
}

func saveStageSet_Wait_Stage(
	inst *models.Wait,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&inst.Time_only, formDiv)
		}
	}
}

func saveStageSet_Wavy_line_Stage(
	inst *models.Wavy_line,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Start_note":
			StageSetFormDivEnumStringFieldToField(&inst.Start_note, formDiv)
		case "Trill_step":
			StageSetFormDivEnumStringFieldToField(&inst.Trill_step, formDiv)
		case "Two_note_turn":
			StageSetFormDivEnumStringFieldToField(&inst.Two_note_turn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&inst.Second_beat, formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&inst.Last_beat, formDiv)
		}
	}
}

func saveStageSet_Wedge_Stage(
	inst *models.Wedge,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			StageSetFormDivEnumStringFieldToField(&inst.Type_XSD, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Spread":
			FormDivBasicFieldToField(&inst.Spread, formDiv)
		case "Niente":
			StageSetFormDivEnumStringFieldToField(&inst.Niente, formDiv)
		case "Line_type":
			StageSetFormDivEnumStringFieldToField(&inst.Line_type, formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&inst.Dash_length, formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&inst.Space_length, formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&inst.Default_x, formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&inst.Default_y, formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&inst.Relative_x, formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&inst.Relative_y, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Wood_Stage(
	inst *models.Wood,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "EnclosedText":
			StageSetFormDivEnumStringFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Work_Stage(
	inst *models.Work,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Work_number":
			FormDivBasicFieldToField(&inst.Work_number, formDiv)
		case "Work_title":
			FormDivBasicFieldToField(&inst.Work_title, formDiv)
		case "Opus":
			StageSetFormDivSelectFieldToField(&inst.Opus, probe.stageOfInterest.GetInstancesSet[*models.Opus](), formDiv)
		}
	}
}

func saveStageSet_ArcType_XlinkStage(
	inst *xlink.ArcType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		case "From":
			FormDivBasicFieldToField(&inst.From, formDiv)
		case "To":
			FormDivBasicFieldToField(&inst.To, formDiv)
		}
	}
}

func saveStageSet_Extended_XlinkStage(
	inst *xlink.Extended,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Title_1":
			StageSetFormDivSelectFieldToField(&inst.Title_1, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.TitleEltType](), formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.ResourceType](), formDiv)
		case "Locator":
			StageSetFormDivSelectFieldToField(&inst.Locator, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.LocatorType](), formDiv)
		case "Arc":
			StageSetFormDivSelectFieldToField(&inst.Arc, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.ArcType](), formDiv)
		}
	}
}

func saveStageSet_LocatorType_XlinkStage(
	inst *xlink.LocatorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Label":
			FormDivBasicFieldToField(&inst.Label, formDiv)
		}
	}
}

func saveStageSet_ResourceType_XlinkStage(
	inst *xlink.ResourceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Label":
			FormDivBasicFieldToField(&inst.Label, formDiv)
		}
	}
}

func saveStageSet_Simple_XlinkStage(
	inst *xlink.Simple,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_TitleEltType_XlinkStage(
	inst *xlink.TitleEltType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		}
	}
}

func saveStageSet_Xmlxsd_SchemaElement_XmlxsdStage(
	inst *xmlxsd.Xmlxsd_SchemaElement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_A_directive_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_directive",
	}).Stage(probe.formStage)
	inst := new(models.A_directive)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_directive_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_directive_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_measure_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_measure",
	}).Stage(probe.formStage)
	inst := new(models.A_measure)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_measure_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_measure_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_measure_1_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_measure_1",
	}).Stage(probe.formStage)
	inst := new(models.A_measure_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_measure_1_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_measure_1_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_part_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_part",
	}).Stage(probe.formStage)
	inst := new(models.A_part)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_part_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_part_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_part_1_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_part_1",
	}).Stage(probe.formStage)
	inst := new(models.A_part_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_part_1_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_part_1_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accidental_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Accidental",
	}).Stage(probe.formStage)
	inst := new(models.Accidental)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Accidental_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Accidental_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accidental_mark_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Accidental_mark",
	}).Stage(probe.formStage)
	inst := new(models.Accidental_mark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Accidental_mark_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Accidental_mark_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accidental_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Accidental_text",
	}).Stage(probe.formStage)
	inst := new(models.Accidental_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Accidental_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Accidental_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accord_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Accord",
	}).Stage(probe.formStage)
	inst := new(models.Accord)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Accord_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Accord_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accordion_registration_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Accordion_registration",
	}).Stage(probe.formStage)
	inst := new(models.Accordion_registration)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Accordion_registration_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Accordion_registration_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Appearance_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Appearance",
	}).Stage(probe.formStage)
	inst := new(models.Appearance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Appearance_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Appearance_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Arpeggiate_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Arpeggiate",
	}).Stage(probe.formStage)
	inst := new(models.Arpeggiate)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Arpeggiate_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Arpeggiate_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Arrow_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Arrow",
	}).Stage(probe.formStage)
	inst := new(models.Arrow)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Arrow_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Arrow_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Articulations_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Articulations",
	}).Stage(probe.formStage)
	inst := new(models.Articulations)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Articulations_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Articulations_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Assess_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Assess",
	}).Stage(probe.formStage)
	inst := new(models.Assess)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Assess_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Assess_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Attributes_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Attributes",
	}).Stage(probe.formStage)
	inst := new(models.Attributes)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Attributes_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Attributes_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Backup_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Backup",
	}).Stage(probe.formStage)
	inst := new(models.Backup)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Backup_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Backup_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bar_style_color_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bar_style_color",
	}).Stage(probe.formStage)
	inst := new(models.Bar_style_color)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bar_style_color_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bar_style_color_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Barline_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Barline",
	}).Stage(probe.formStage)
	inst := new(models.Barline)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Barline_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Barline_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Barre_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Barre",
	}).Stage(probe.formStage)
	inst := new(models.Barre)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Barre_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Barre_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bass_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bass",
	}).Stage(probe.formStage)
	inst := new(models.Bass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bass_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bass_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bass_step_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bass_step",
	}).Stage(probe.formStage)
	inst := new(models.Bass_step)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bass_step_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bass_step_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beam_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Beam",
	}).Stage(probe.formStage)
	inst := new(models.Beam)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Beam_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Beam_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beat_repeat_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Beat_repeat",
	}).Stage(probe.formStage)
	inst := new(models.Beat_repeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Beat_repeat_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Beat_repeat_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beat_unit_tied_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Beat_unit_tied",
	}).Stage(probe.formStage)
	inst := new(models.Beat_unit_tied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Beat_unit_tied_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Beat_unit_tied_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beater_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Beater",
	}).Stage(probe.formStage)
	inst := new(models.Beater)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Beater_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Beater_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bend_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bend",
	}).Stage(probe.formStage)
	inst := new(models.Bend)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bend_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bend_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bookmark_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bookmark",
	}).Stage(probe.formStage)
	inst := new(models.Bookmark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bookmark_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bookmark_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bracket_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Bracket",
	}).Stage(probe.formStage)
	inst := new(models.Bracket)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Bracket_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Bracket_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Breath_mark_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Breath_mark",
	}).Stage(probe.formStage)
	inst := new(models.Breath_mark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Breath_mark_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Breath_mark_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Caesura_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Caesura",
	}).Stage(probe.formStage)
	inst := new(models.Caesura)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Caesura_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Caesura_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Cancel_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Cancel",
	}).Stage(probe.formStage)
	inst := new(models.Cancel)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Cancel_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Cancel_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Clef_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Clef",
	}).Stage(probe.formStage)
	inst := new(models.Clef)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Clef_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Clef_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Coda_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Coda",
	}).Stage(probe.formStage)
	inst := new(models.Coda)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Coda_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Coda_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Credit_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Credit",
	}).Stage(probe.formStage)
	inst := new(models.Credit)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Credit_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Credit_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Dashes_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Dashes",
	}).Stage(probe.formStage)
	inst := new(models.Dashes)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Dashes_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Dashes_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Defaults_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Defaults",
	}).Stage(probe.formStage)
	inst := new(models.Defaults)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Defaults_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Defaults_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Degree_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Degree",
	}).Stage(probe.formStage)
	inst := new(models.Degree)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Degree_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Degree_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Degree_alter_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Degree_alter",
	}).Stage(probe.formStage)
	inst := new(models.Degree_alter)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Degree_alter_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Degree_alter_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Degree_type_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Degree_type",
	}).Stage(probe.formStage)
	inst := new(models.Degree_type)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Degree_type_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Degree_type_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Degree_value_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Degree_value",
	}).Stage(probe.formStage)
	inst := new(models.Degree_value)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Degree_value_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Degree_value_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Direction_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Direction",
	}).Stage(probe.formStage)
	inst := new(models.Direction)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Direction_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Direction_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Direction_type_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Direction_type",
	}).Stage(probe.formStage)
	inst := new(models.Direction_type)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Direction_type_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Direction_type_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Distance_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Distance",
	}).Stage(probe.formStage)
	inst := new(models.Distance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Distance_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Distance_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Double_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Double",
	}).Stage(probe.formStage)
	inst := new(models.Double)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Double_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Double_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Dynamics_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Dynamics",
	}).Stage(probe.formStage)
	inst := new(models.Dynamics)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Dynamics_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Dynamics_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Effect_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Effect",
	}).Stage(probe.formStage)
	inst := new(models.Effect)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Effect_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Effect_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Elision_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Elision",
	}).Stage(probe.formStage)
	inst := new(models.Elision)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Elision_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Elision_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty",
	}).Stage(probe.formStage)
	inst := new(models.Empty)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_font_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_font",
	}).Stage(probe.formStage)
	inst := new(models.Empty_font)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_font_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_font_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_line_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_line",
	}).Stage(probe.formStage)
	inst := new(models.Empty_line)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_line_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_line_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_placement_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_placement",
	}).Stage(probe.formStage)
	inst := new(models.Empty_placement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_placement_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_placement_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_placement_smufl_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_placement_smufl",
	}).Stage(probe.formStage)
	inst := new(models.Empty_placement_smufl)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_placement_smufl_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_placement_smufl_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_print_object_style_align_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_print_object_style_align",
	}).Stage(probe.formStage)
	inst := new(models.Empty_print_object_style_align)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_print_object_style_align_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_print_object_style_align_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_print_style_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_print_style",
	}).Stage(probe.formStage)
	inst := new(models.Empty_print_style)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_print_style_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_print_style_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_print_style_align_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_print_style_align",
	}).Stage(probe.formStage)
	inst := new(models.Empty_print_style_align)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_print_style_align_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_print_style_align_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_print_style_align_id_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_print_style_align_id",
	}).Stage(probe.formStage)
	inst := new(models.Empty_print_style_align_id)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_print_style_align_id_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_print_style_align_id_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_trill_sound_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Empty_trill_sound",
	}).Stage(probe.formStage)
	inst := new(models.Empty_trill_sound)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Empty_trill_sound_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Empty_trill_sound_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Encoding_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Encoding",
	}).Stage(probe.formStage)
	inst := new(models.Encoding)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Encoding_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Encoding_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Ending_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Ending",
	}).Stage(probe.formStage)
	inst := new(models.Ending)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Ending_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Ending_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Extend_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Extend",
	}).Stage(probe.formStage)
	inst := new(models.Extend)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Extend_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Extend_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Feature_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Feature",
	}).Stage(probe.formStage)
	inst := new(models.Feature)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Feature_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Feature_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fermata_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Fermata",
	}).Stage(probe.formStage)
	inst := new(models.Fermata)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Fermata_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Fermata_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Figure_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Figure",
	}).Stage(probe.formStage)
	inst := new(models.Figure)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Figure_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Figure_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Figured_bass_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Figured_bass",
	}).Stage(probe.formStage)
	inst := new(models.Figured_bass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Figured_bass_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Figured_bass_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fingering_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Fingering",
	}).Stage(probe.formStage)
	inst := new(models.Fingering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Fingering_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Fingering_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_First_fret_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New First_fret",
	}).Stage(probe.formStage)
	inst := new(models.First_fret)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_First_fret_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_First_fret_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_For_part_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New For_part",
	}).Stage(probe.formStage)
	inst := new(models.For_part)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_For_part_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_For_part_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Formatted_symbol_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Formatted_symbol",
	}).Stage(probe.formStage)
	inst := new(models.Formatted_symbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Formatted_symbol_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Formatted_symbol_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Formatted_symbol_id_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Formatted_symbol_id",
	}).Stage(probe.formStage)
	inst := new(models.Formatted_symbol_id)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Formatted_symbol_id_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Formatted_symbol_id_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Formatted_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Formatted_text",
	}).Stage(probe.formStage)
	inst := new(models.Formatted_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Formatted_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Formatted_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Formatted_text_id_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Formatted_text_id",
	}).Stage(probe.formStage)
	inst := new(models.Formatted_text_id)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Formatted_text_id_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Formatted_text_id_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Forward_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Forward",
	}).Stage(probe.formStage)
	inst := new(models.Forward)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Forward_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Forward_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Frame_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Frame",
	}).Stage(probe.formStage)
	inst := new(models.Frame)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Frame_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Frame_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Frame_note_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Frame_note",
	}).Stage(probe.formStage)
	inst := new(models.Frame_note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Frame_note_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Frame_note_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fret_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Fret",
	}).Stage(probe.formStage)
	inst := new(models.Fret)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Fret_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Fret_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glass_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Glass",
	}).Stage(probe.formStage)
	inst := new(models.Glass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Glass_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Glass_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glissando_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Glissando",
	}).Stage(probe.formStage)
	inst := new(models.Glissando)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Glissando_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Glissando_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glyph_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Glyph",
	}).Stage(probe.formStage)
	inst := new(models.Glyph)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Glyph_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Glyph_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Grace_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Grace",
	}).Stage(probe.formStage)
	inst := new(models.Grace)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Grace_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Grace_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Group_barline_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Group_barline",
	}).Stage(probe.formStage)
	inst := new(models.Group_barline)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Group_barline_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Group_barline_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Group_name_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Group_name",
	}).Stage(probe.formStage)
	inst := new(models.Group_name)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Group_name_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Group_name_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Group_symbol_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Group_symbol",
	}).Stage(probe.formStage)
	inst := new(models.Group_symbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Group_symbol_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Group_symbol_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Grouping_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Grouping",
	}).Stage(probe.formStage)
	inst := new(models.Grouping)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Grouping_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Grouping_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Hammer_on_pull_off_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Hammer_on_pull_off",
	}).Stage(probe.formStage)
	inst := new(models.Hammer_on_pull_off)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Hammer_on_pull_off_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Hammer_on_pull_off_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Handbell_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Handbell",
	}).Stage(probe.formStage)
	inst := new(models.Handbell)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Handbell_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Handbell_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmon_closed_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harmon_closed",
	}).Stage(probe.formStage)
	inst := new(models.Harmon_closed)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harmon_closed_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harmon_closed_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmon_mute_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harmon_mute",
	}).Stage(probe.formStage)
	inst := new(models.Harmon_mute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harmon_mute_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harmon_mute_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmonic_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harmonic",
	}).Stage(probe.formStage)
	inst := new(models.Harmonic)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harmonic_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harmonic_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmony_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harmony",
	}).Stage(probe.formStage)
	inst := new(models.Harmony)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harmony_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harmony_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmony_alter_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harmony_alter",
	}).Stage(probe.formStage)
	inst := new(models.Harmony_alter)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harmony_alter_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harmony_alter_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harp_pedals_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Harp_pedals",
	}).Stage(probe.formStage)
	inst := new(models.Harp_pedals)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Harp_pedals_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Harp_pedals_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Heel_toe_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Heel_toe",
	}).Stage(probe.formStage)
	inst := new(models.Heel_toe)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Heel_toe_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Heel_toe_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Hole_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Hole",
	}).Stage(probe.formStage)
	inst := new(models.Hole)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Hole_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Hole_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Hole_closed_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Hole_closed",
	}).Stage(probe.formStage)
	inst := new(models.Hole_closed)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Hole_closed_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Hole_closed_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Horizontal_turn_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Horizontal_turn",
	}).Stage(probe.formStage)
	inst := new(models.Horizontal_turn)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Horizontal_turn_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Horizontal_turn_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Identification_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Identification",
	}).Stage(probe.formStage)
	inst := new(models.Identification)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Identification_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Identification_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Image_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Image",
	}).Stage(probe.formStage)
	inst := new(models.Image)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Image_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Image_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Instrument_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Instrument",
	}).Stage(probe.formStage)
	inst := new(models.Instrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Instrument_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Instrument_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Instrument_change_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Instrument_change",
	}).Stage(probe.formStage)
	inst := new(models.Instrument_change)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Instrument_change_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Instrument_change_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Instrument_link_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Instrument_link",
	}).Stage(probe.formStage)
	inst := new(models.Instrument_link)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Instrument_link_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Instrument_link_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Interchangeable_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Interchangeable",
	}).Stage(probe.formStage)
	inst := new(models.Interchangeable)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Interchangeable_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Interchangeable_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Inversion_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Inversion",
	}).Stage(probe.formStage)
	inst := new(models.Inversion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Inversion_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Inversion_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Key_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Key",
	}).Stage(probe.formStage)
	inst := new(models.Key)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Key_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Key_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Key_accidental_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Key_accidental",
	}).Stage(probe.formStage)
	inst := new(models.Key_accidental)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Key_accidental_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Key_accidental_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Key_octave_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Key_octave",
	}).Stage(probe.formStage)
	inst := new(models.Key_octave)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Key_octave_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Key_octave_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Kind_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Kind",
	}).Stage(probe.formStage)
	inst := new(models.Kind)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Kind_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Kind_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Level_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Level",
	}).Stage(probe.formStage)
	inst := new(models.Level)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Level_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Level_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Line_detail_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Line_detail",
	}).Stage(probe.formStage)
	inst := new(models.Line_detail)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Line_detail_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Line_detail_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Line_width_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Line_width",
	}).Stage(probe.formStage)
	inst := new(models.Line_width)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Line_width_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Line_width_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Link_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Link",
	}).Stage(probe.formStage)
	inst := new(models.Link)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Link_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Link_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Listen_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Listen",
	}).Stage(probe.formStage)
	inst := new(models.Listen)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Listen_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Listen_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Listening_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Listening",
	}).Stage(probe.formStage)
	inst := new(models.Listening)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Listening_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Listening_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Lyric_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Lyric",
	}).Stage(probe.formStage)
	inst := new(models.Lyric)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Lyric_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Lyric_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Lyric_font_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Lyric_font",
	}).Stage(probe.formStage)
	inst := new(models.Lyric_font)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Lyric_font_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Lyric_font_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Lyric_language_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Lyric_language",
	}).Stage(probe.formStage)
	inst := new(models.Lyric_language)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Lyric_language_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Lyric_language_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Measure_layout_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Measure_layout",
	}).Stage(probe.formStage)
	inst := new(models.Measure_layout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Measure_layout_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Measure_layout_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Measure_numbering_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Measure_numbering",
	}).Stage(probe.formStage)
	inst := new(models.Measure_numbering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Measure_numbering_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Measure_numbering_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Measure_repeat_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Measure_repeat",
	}).Stage(probe.formStage)
	inst := new(models.Measure_repeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Measure_repeat_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Measure_repeat_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Measure_style_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Measure_style",
	}).Stage(probe.formStage)
	inst := new(models.Measure_style)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Measure_style_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Measure_style_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Membrane_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Membrane",
	}).Stage(probe.formStage)
	inst := new(models.Membrane)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Membrane_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Membrane_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metal_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metal",
	}).Stage(probe.formStage)
	inst := new(models.Metal)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metal_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metal_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metronome",
	}).Stage(probe.formStage)
	inst := new(models.Metronome)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metronome_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metronome_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_beam_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metronome_beam",
	}).Stage(probe.formStage)
	inst := new(models.Metronome_beam)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metronome_beam_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metronome_beam_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_note_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metronome_note",
	}).Stage(probe.formStage)
	inst := new(models.Metronome_note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metronome_note_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metronome_note_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_tied_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metronome_tied",
	}).Stage(probe.formStage)
	inst := new(models.Metronome_tied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metronome_tied_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metronome_tied_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_tuplet_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Metronome_tuplet",
	}).Stage(probe.formStage)
	inst := new(models.Metronome_tuplet)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Metronome_tuplet_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Metronome_tuplet_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Midi_device_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Midi_device",
	}).Stage(probe.formStage)
	inst := new(models.Midi_device)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Midi_device_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Midi_device_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Midi_instrument_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Midi_instrument",
	}).Stage(probe.formStage)
	inst := new(models.Midi_instrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Midi_instrument_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Midi_instrument_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Miscellaneous_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Miscellaneous",
	}).Stage(probe.formStage)
	inst := new(models.Miscellaneous)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Miscellaneous_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Miscellaneous_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Miscellaneous_field_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Miscellaneous_field",
	}).Stage(probe.formStage)
	inst := new(models.Miscellaneous_field)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Miscellaneous_field_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Miscellaneous_field_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Mordent_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Mordent",
	}).Stage(probe.formStage)
	inst := new(models.Mordent)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Mordent_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Mordent_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Multiple_rest_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Multiple_rest",
	}).Stage(probe.formStage)
	inst := new(models.Multiple_rest)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Multiple_rest_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Multiple_rest_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Name_display_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Name_display",
	}).Stage(probe.formStage)
	inst := new(models.Name_display)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Name_display_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Name_display_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Non_arpeggiate_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Non_arpeggiate",
	}).Stage(probe.formStage)
	inst := new(models.Non_arpeggiate)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Non_arpeggiate_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Non_arpeggiate_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Notations_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Notations",
	}).Stage(probe.formStage)
	inst := new(models.Notations)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Notations_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Notations_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Note_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Note",
	}).Stage(probe.formStage)
	inst := new(models.Note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Note_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Note_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Note_size_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Note_size",
	}).Stage(probe.formStage)
	inst := new(models.Note_size)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Note_size_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Note_size_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Note_type_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Note_type",
	}).Stage(probe.formStage)
	inst := new(models.Note_type)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Note_type_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Note_type_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Notehead_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Notehead",
	}).Stage(probe.formStage)
	inst := new(models.Notehead)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Notehead_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Notehead_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Notehead_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Notehead_text",
	}).Stage(probe.formStage)
	inst := new(models.Notehead_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Notehead_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Notehead_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Numeral_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Numeral",
	}).Stage(probe.formStage)
	inst := new(models.Numeral)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Numeral_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Numeral_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Numeral_key_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Numeral_key",
	}).Stage(probe.formStage)
	inst := new(models.Numeral_key)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Numeral_key_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Numeral_key_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Numeral_root_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Numeral_root",
	}).Stage(probe.formStage)
	inst := new(models.Numeral_root)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Numeral_root_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Numeral_root_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Octave_shift_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Octave_shift",
	}).Stage(probe.formStage)
	inst := new(models.Octave_shift)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Octave_shift_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Octave_shift_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Offset_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Offset",
	}).Stage(probe.formStage)
	inst := new(models.Offset)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Offset_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Offset_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Opus_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Opus",
	}).Stage(probe.formStage)
	inst := new(models.Opus)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Opus_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Opus_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Ornaments_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Ornaments",
	}).Stage(probe.formStage)
	inst := new(models.Ornaments)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Ornaments_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Ornaments_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_appearance_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_appearance",
	}).Stage(probe.formStage)
	inst := new(models.Other_appearance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_appearance_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_appearance_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_direction_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_direction",
	}).Stage(probe.formStage)
	inst := new(models.Other_direction)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_direction_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_direction_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_listening_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_listening",
	}).Stage(probe.formStage)
	inst := new(models.Other_listening)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_listening_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_listening_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_notation_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_notation",
	}).Stage(probe.formStage)
	inst := new(models.Other_notation)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_notation_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_notation_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_placement_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_placement_text",
	}).Stage(probe.formStage)
	inst := new(models.Other_placement_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_placement_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_placement_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_play_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_play",
	}).Stage(probe.formStage)
	inst := new(models.Other_play)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_play_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_play_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Other_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Other_text",
	}).Stage(probe.formStage)
	inst := new(models.Other_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Other_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Other_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Page_layout_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Page_layout",
	}).Stage(probe.formStage)
	inst := new(models.Page_layout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Page_layout_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Page_layout_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Page_margins_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Page_margins",
	}).Stage(probe.formStage)
	inst := new(models.Page_margins)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Page_margins_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Page_margins_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_clef_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_clef",
	}).Stage(probe.formStage)
	inst := new(models.Part_clef)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_clef_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_clef_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_group_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_group",
	}).Stage(probe.formStage)
	inst := new(models.Part_group)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_group_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_group_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_link_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_link",
	}).Stage(probe.formStage)
	inst := new(models.Part_link)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_link_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_link_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_list_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_list",
	}).Stage(probe.formStage)
	inst := new(models.Part_list)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_list_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_list_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_name_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_name",
	}).Stage(probe.formStage)
	inst := new(models.Part_name)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_name_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_name_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_symbol_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_symbol",
	}).Stage(probe.formStage)
	inst := new(models.Part_symbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_symbol_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_symbol_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_transpose_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Part_transpose",
	}).Stage(probe.formStage)
	inst := new(models.Part_transpose)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Part_transpose_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Part_transpose_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pedal_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Pedal",
	}).Stage(probe.formStage)
	inst := new(models.Pedal)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Pedal_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Pedal_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pedal_tuning_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Pedal_tuning",
	}).Stage(probe.formStage)
	inst := new(models.Pedal_tuning)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Pedal_tuning_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Pedal_tuning_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Per_minute_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Per_minute",
	}).Stage(probe.formStage)
	inst := new(models.Per_minute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Per_minute_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Per_minute_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Percussion_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Percussion",
	}).Stage(probe.formStage)
	inst := new(models.Percussion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Percussion_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Percussion_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pitch_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Pitch",
	}).Stage(probe.formStage)
	inst := new(models.Pitch)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Pitch_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Pitch_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pitched_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Pitched",
	}).Stage(probe.formStage)
	inst := new(models.Pitched)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Pitched_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Pitched_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Placement_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Placement_text",
	}).Stage(probe.formStage)
	inst := new(models.Placement_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Placement_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Placement_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Play_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Play",
	}).Stage(probe.formStage)
	inst := new(models.Play)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Play_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Play_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Player_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Player",
	}).Stage(probe.formStage)
	inst := new(models.Player)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Player_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Player_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Principal_voice_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Principal_voice",
	}).Stage(probe.formStage)
	inst := new(models.Principal_voice)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Principal_voice_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Principal_voice_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Print_XSD_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Print_XSD",
	}).Stage(probe.formStage)
	inst := new(models.Print_XSD)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Print_XSD_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Print_XSD_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Release_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Release",
	}).Stage(probe.formStage)
	inst := new(models.Release)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Release_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Release_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Repeat_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Repeat",
	}).Stage(probe.formStage)
	inst := new(models.Repeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Repeat_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Repeat_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Rest_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Rest",
	}).Stage(probe.formStage)
	inst := new(models.Rest)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Rest_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Rest_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Root_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Root",
	}).Stage(probe.formStage)
	inst := new(models.Root)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Root_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Root_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Root_step_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Root_step",
	}).Stage(probe.formStage)
	inst := new(models.Root_step)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Root_step_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Root_step_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Scaling_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Scaling",
	}).Stage(probe.formStage)
	inst := new(models.Scaling)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Scaling_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Scaling_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Scordatura_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Scordatura",
	}).Stage(probe.formStage)
	inst := new(models.Scordatura)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Scordatura_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Scordatura_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Score_instrument_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Score_instrument",
	}).Stage(probe.formStage)
	inst := new(models.Score_instrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Score_instrument_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Score_instrument_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Score_part_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Score_part",
	}).Stage(probe.formStage)
	inst := new(models.Score_part)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Score_part_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Score_part_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Score_partwise_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Score_partwise",
	}).Stage(probe.formStage)
	inst := new(models.Score_partwise)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Score_partwise_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Score_partwise_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Score_timewise_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Score_timewise",
	}).Stage(probe.formStage)
	inst := new(models.Score_timewise)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Score_timewise_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Score_timewise_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Segno_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Segno",
	}).Stage(probe.formStage)
	inst := new(models.Segno)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Segno_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Segno_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slash_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Slash",
	}).Stage(probe.formStage)
	inst := new(models.Slash)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Slash_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Slash_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slide_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Slide",
	}).Stage(probe.formStage)
	inst := new(models.Slide)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Slide_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Slide_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slur_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Slur",
	}).Stage(probe.formStage)
	inst := new(models.Slur)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Slur_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Slur_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Sound_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Sound",
	}).Stage(probe.formStage)
	inst := new(models.Sound)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Sound_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Sound_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Staff_details_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Staff_details",
	}).Stage(probe.formStage)
	inst := new(models.Staff_details)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Staff_details_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Staff_details_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Staff_divide_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Staff_divide",
	}).Stage(probe.formStage)
	inst := new(models.Staff_divide)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Staff_divide_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Staff_divide_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Staff_layout_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Staff_layout",
	}).Stage(probe.formStage)
	inst := new(models.Staff_layout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Staff_layout_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Staff_layout_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Staff_size_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Staff_size",
	}).Stage(probe.formStage)
	inst := new(models.Staff_size)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Staff_size_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Staff_size_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Staff_tuning_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Staff_tuning",
	}).Stage(probe.formStage)
	inst := new(models.Staff_tuning)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Staff_tuning_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Staff_tuning_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Stem_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Stem",
	}).Stage(probe.formStage)
	inst := new(models.Stem)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Stem_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Stem_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Stick_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Stick",
	}).Stage(probe.formStage)
	inst := new(models.Stick)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Stick_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Stick_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_String_mute_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New String_mute",
	}).Stage(probe.formStage)
	inst := new(models.String_mute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_String_mute_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_String_mute_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_String_type_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New String_type",
	}).Stage(probe.formStage)
	inst := new(models.String_type)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_String_type_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_String_type_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Strong_accent_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Strong_accent",
	}).Stage(probe.formStage)
	inst := new(models.Strong_accent)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Strong_accent_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Strong_accent_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Style_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Style_text",
	}).Stage(probe.formStage)
	inst := new(models.Style_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Style_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Style_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Supports_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Supports",
	}).Stage(probe.formStage)
	inst := new(models.Supports)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Supports_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Supports_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Swing_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Swing",
	}).Stage(probe.formStage)
	inst := new(models.Swing)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Swing_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Swing_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Sync_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Sync",
	}).Stage(probe.formStage)
	inst := new(models.Sync)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Sync_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Sync_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_System_dividers_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New System_dividers",
	}).Stage(probe.formStage)
	inst := new(models.System_dividers)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_System_dividers_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_System_dividers_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_System_layout_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New System_layout",
	}).Stage(probe.formStage)
	inst := new(models.System_layout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_System_layout_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_System_layout_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_System_margins_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New System_margins",
	}).Stage(probe.formStage)
	inst := new(models.System_margins)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_System_margins_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_System_margins_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tap_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tap",
	}).Stage(probe.formStage)
	inst := new(models.Tap)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tap_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tap_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Technical_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Technical",
	}).Stage(probe.formStage)
	inst := new(models.Technical)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Technical_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Technical_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Text_element_data_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Text_element_data",
	}).Stage(probe.formStage)
	inst := new(models.Text_element_data)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Text_element_data_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Text_element_data_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tie_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tie",
	}).Stage(probe.formStage)
	inst := new(models.Tie)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tie_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tie_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tied_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tied",
	}).Stage(probe.formStage)
	inst := new(models.Tied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tied_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tied_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Time_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Time",
	}).Stage(probe.formStage)
	inst := new(models.Time)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Time_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Time_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Time_modification_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Time_modification",
	}).Stage(probe.formStage)
	inst := new(models.Time_modification)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Time_modification_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Time_modification_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Timpani_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Timpani",
	}).Stage(probe.formStage)
	inst := new(models.Timpani)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Timpani_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Timpani_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Transpose_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Transpose",
	}).Stage(probe.formStage)
	inst := new(models.Transpose)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Transpose_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Transpose_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tremolo_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tremolo",
	}).Stage(probe.formStage)
	inst := new(models.Tremolo)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tremolo_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tremolo_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tuplet",
	}).Stage(probe.formStage)
	inst := new(models.Tuplet)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tuplet_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tuplet_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_dot_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tuplet_dot",
	}).Stage(probe.formStage)
	inst := new(models.Tuplet_dot)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tuplet_dot_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tuplet_dot_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_number_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tuplet_number",
	}).Stage(probe.formStage)
	inst := new(models.Tuplet_number)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tuplet_number_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tuplet_number_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_portion_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tuplet_portion",
	}).Stage(probe.formStage)
	inst := new(models.Tuplet_portion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tuplet_portion_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tuplet_portion_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_type_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Tuplet_type",
	}).Stage(probe.formStage)
	inst := new(models.Tuplet_type)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Tuplet_type_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Tuplet_type_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Typed_text_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Typed_text",
	}).Stage(probe.formStage)
	inst := new(models.Typed_text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Typed_text_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Typed_text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Unpitched_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Unpitched",
	}).Stage(probe.formStage)
	inst := new(models.Unpitched)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Unpitched_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Unpitched_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Virtual_instrument_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Virtual_instrument",
	}).Stage(probe.formStage)
	inst := new(models.Virtual_instrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Virtual_instrument_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Virtual_instrument_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wait_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Wait",
	}).Stage(probe.formStage)
	inst := new(models.Wait)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Wait_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Wait_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wavy_line_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Wavy_line",
	}).Stage(probe.formStage)
	inst := new(models.Wavy_line)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Wavy_line_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Wavy_line_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wedge_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Wedge",
	}).Stage(probe.formStage)
	inst := new(models.Wedge)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Wedge_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Wedge_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wood_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Wood",
	}).Stage(probe.formStage)
	inst := new(models.Wood)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Wood_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Wood_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Work_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Work",
	}).Stage(probe.formStage)
	inst := new(models.Work)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Work_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Work_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArcType_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArcType",
	}).Stage(probe.formStage)
	inst := new(xlink.ArcType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_ArcType_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_ArcType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Extended_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Extended",
	}).Stage(probe.formStage)
	inst := new(xlink.Extended)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_Extended_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_Extended_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LocatorType_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LocatorType",
	}).Stage(probe.formStage)
	inst := new(xlink.LocatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_LocatorType_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_LocatorType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ResourceType_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ResourceType",
	}).Stage(probe.formStage)
	inst := new(xlink.ResourceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_ResourceType_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_ResourceType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Simple_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Simple",
	}).Stage(probe.formStage)
	inst := new(xlink.Simple)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_Simple_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_Simple_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TitleEltType_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TitleEltType",
	}).Stage(probe.formStage)
	inst := new(xlink.TitleEltType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_TitleEltType_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_TitleEltType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Xmlxsd_SchemaElement_XmlxsdStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Xmlxsd_SchemaElement",
	}).Stage(probe.formStage)
	inst := new(xmlxsd.Xmlxsd_SchemaElement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XmlxsdStage.Lock()
			defer probe.stageOfInterest.XmlxsdStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XmlxsdStage)
			saveStageSet_Xmlxsd_SchemaElement_XmlxsdStage(inst, probe, formGroup)
			probe.stageOfInterest.XmlxsdStage.Commit()
			updateStageSetTable_Xmlxsd_SchemaElement_XmlxsdStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

