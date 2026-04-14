// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.A_directive:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Directive"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Directive",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Directive",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_measure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Implicit", instanceWithInferedType.Implicit, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Non_controlling", instanceWithInferedType.Non_controlling, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Note", instanceWithInferedType, &instanceWithInferedType.Note, formGroup, probe)
		AssociationSliceToForm("Backup", instanceWithInferedType, &instanceWithInferedType.Backup, formGroup, probe)
		AssociationSliceToForm("Forward", instanceWithInferedType, &instanceWithInferedType.Forward, formGroup, probe)
		AssociationSliceToForm("Direction", instanceWithInferedType, &instanceWithInferedType.Direction, formGroup, probe)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		AssociationSliceToForm("Harmony", instanceWithInferedType, &instanceWithInferedType.Harmony, formGroup, probe)
		AssociationSliceToForm("Figured_bass", instanceWithInferedType, &instanceWithInferedType.Figured_bass, formGroup, probe)
		AssociationSliceToForm("Print", instanceWithInferedType, &instanceWithInferedType.Print, formGroup, probe)
		AssociationSliceToForm("Sound", instanceWithInferedType, &instanceWithInferedType.Sound, formGroup, probe)
		AssociationSliceToForm("Listening", instanceWithInferedType, &instanceWithInferedType.Listening, formGroup, probe)
		AssociationSliceToForm("Barline", instanceWithInferedType, &instanceWithInferedType.Barline, formGroup, probe)
		AssociationSliceToForm("Grouping", instanceWithInferedType, &instanceWithInferedType.Grouping, formGroup, probe)
		AssociationSliceToForm("Link", instanceWithInferedType, &instanceWithInferedType.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", instanceWithInferedType, &instanceWithInferedType.Bookmark, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part"
			rf.Fieldname = "Measure"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part),
					"Measure",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part](
					nil,
					"Measure",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_measure_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Implicit", instanceWithInferedType.Implicit, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Non_controlling", instanceWithInferedType.Non_controlling, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Part", instanceWithInferedType, &instanceWithInferedType.Part, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_timewise"
			rf.Fieldname = "Measure"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_timewise),
					"Measure",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_timewise](
					nil,
					"Measure",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Measure", instanceWithInferedType, &instanceWithInferedType.Measure, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_partwise"
			rf.Fieldname = "Part"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_partwise),
					"Part",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_partwise](
					nil,
					"Part",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_part_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Note", instanceWithInferedType, &instanceWithInferedType.Note, formGroup, probe)
		AssociationSliceToForm("Backup", instanceWithInferedType, &instanceWithInferedType.Backup, formGroup, probe)
		AssociationSliceToForm("Forward", instanceWithInferedType, &instanceWithInferedType.Forward, formGroup, probe)
		AssociationSliceToForm("Direction", instanceWithInferedType, &instanceWithInferedType.Direction, formGroup, probe)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		AssociationSliceToForm("Harmony", instanceWithInferedType, &instanceWithInferedType.Harmony, formGroup, probe)
		AssociationSliceToForm("Figured_bass", instanceWithInferedType, &instanceWithInferedType.Figured_bass, formGroup, probe)
		AssociationSliceToForm("Print", instanceWithInferedType, &instanceWithInferedType.Print, formGroup, probe)
		AssociationSliceToForm("Sound", instanceWithInferedType, &instanceWithInferedType.Sound, formGroup, probe)
		AssociationSliceToForm("Listening", instanceWithInferedType, &instanceWithInferedType.Listening, formGroup, probe)
		AssociationSliceToForm("Barline", instanceWithInferedType, &instanceWithInferedType.Barline, formGroup, probe)
		AssociationSliceToForm("Grouping", instanceWithInferedType, &instanceWithInferedType.Grouping, formGroup, probe)
		AssociationSliceToForm("Link", instanceWithInferedType, &instanceWithInferedType.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", instanceWithInferedType, &instanceWithInferedType.Bookmark, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure_1"
			rf.Fieldname = "Part"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure_1),
					"Part",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure_1](
					nil,
					"Part",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Cautionary", instanceWithInferedType.Cautionary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Editorial", instanceWithInferedType.Editorial, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Accidental_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Accidental_mark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Accidental_mark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accidental_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Name_display"
			rf.Fieldname = "Accidental_text"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Name_display),
					"Accidental_text",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Name_display](
					nil,
					"Accidental_text",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notehead_text"
			rf.Fieldname = "Accidental_text"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notehead_text),
					"Accidental_text",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notehead_text](
					nil,
					"Accidental_text",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accord:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("String", instanceWithInferedType.String, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Tuning_step", instanceWithInferedType.Tuning_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Tuning_alter", instanceWithInferedType.Tuning_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tuning_octave", instanceWithInferedType.Tuning_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scordatura"
			rf.Fieldname = "Accord"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Scordatura),
					"Accord",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Scordatura](
					nil,
					"Accord",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accordion_registration:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Accordion_high", instanceWithInferedType.Accordion_high, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Accordion_middle", instanceWithInferedType.Accordion_middle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Accordion_low", instanceWithInferedType.Accordion_low, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Line_width", instanceWithInferedType, &instanceWithInferedType.Line_width, formGroup, probe)
		AssociationSliceToForm("Note_size", instanceWithInferedType, &instanceWithInferedType.Note_size, formGroup, probe)
		AssociationSliceToForm("Distance", instanceWithInferedType, &instanceWithInferedType.Distance, formGroup, probe)
		AssociationSliceToForm("Glyph", instanceWithInferedType, &instanceWithInferedType.Glyph, formGroup, probe)
		AssociationSliceToForm("Other_appearance", instanceWithInferedType, &instanceWithInferedType.Other_appearance, formGroup, probe)

	case *models.Arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Unbroken", instanceWithInferedType.Unbroken, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Arpeggiate"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Arpeggiate",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Arpeggiate",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Arrow_direction", instanceWithInferedType.Arrow_direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Arrow_style", instanceWithInferedType.Arrow_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Arrowhead", instanceWithInferedType.Arrowhead, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Circular_arrow", instanceWithInferedType.Circular_arrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Arrow"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Arrow",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Arrow",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Articulations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Accent", instanceWithInferedType, &instanceWithInferedType.Accent, formGroup, probe)
		AssociationSliceToForm("Strong_accent", instanceWithInferedType, &instanceWithInferedType.Strong_accent, formGroup, probe)
		AssociationSliceToForm("Staccato", instanceWithInferedType, &instanceWithInferedType.Staccato, formGroup, probe)
		AssociationSliceToForm("Tenuto", instanceWithInferedType, &instanceWithInferedType.Tenuto, formGroup, probe)
		AssociationSliceToForm("Detached_legato", instanceWithInferedType, &instanceWithInferedType.Detached_legato, formGroup, probe)
		AssociationSliceToForm("Staccatissimo", instanceWithInferedType, &instanceWithInferedType.Staccatissimo, formGroup, probe)
		AssociationSliceToForm("Spiccato", instanceWithInferedType, &instanceWithInferedType.Spiccato, formGroup, probe)
		AssociationSliceToForm("Scoop", instanceWithInferedType, &instanceWithInferedType.Scoop, formGroup, probe)
		AssociationSliceToForm("Plop", instanceWithInferedType, &instanceWithInferedType.Plop, formGroup, probe)
		AssociationSliceToForm("Doit", instanceWithInferedType, &instanceWithInferedType.Doit, formGroup, probe)
		AssociationSliceToForm("Falloff", instanceWithInferedType, &instanceWithInferedType.Falloff, formGroup, probe)
		AssociationSliceToForm("Breath_mark", instanceWithInferedType, &instanceWithInferedType.Breath_mark, formGroup, probe)
		AssociationSliceToForm("Caesura", instanceWithInferedType, &instanceWithInferedType.Caesura, formGroup, probe)
		AssociationSliceToForm("Stress", instanceWithInferedType, &instanceWithInferedType.Stress, formGroup, probe)
		AssociationSliceToForm("Unstress", instanceWithInferedType, &instanceWithInferedType.Unstress, formGroup, probe)
		AssociationSliceToForm("Soft_accent", instanceWithInferedType, &instanceWithInferedType.Soft_accent, formGroup, probe)
		AssociationSliceToForm("Other_articulation", instanceWithInferedType, &instanceWithInferedType.Other_articulation, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Articulations"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Articulations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Articulations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Assess:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Listen"
			rf.Fieldname = "Assess"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Listen),
					"Assess",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Listen](
					nil,
					"Assess",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Attributes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Key", instanceWithInferedType, &instanceWithInferedType.Key, formGroup, probe)
		AssociationSliceToForm("Time", instanceWithInferedType, &instanceWithInferedType.Time, formGroup, probe)
		BasicFieldtoForm("Staves", instanceWithInferedType.Staves, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Part_symbol", instanceWithInferedType.Part_symbol, formGroup, probe)
		BasicFieldtoForm("Instruments", instanceWithInferedType.Instruments, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Clef", instanceWithInferedType, &instanceWithInferedType.Clef, formGroup, probe)
		AssociationSliceToForm("Staff_details", instanceWithInferedType, &instanceWithInferedType.Staff_details, formGroup, probe)
		AssociationSliceToForm("Transpose", instanceWithInferedType, &instanceWithInferedType.Transpose, formGroup, probe)
		AssociationSliceToForm("For_part", instanceWithInferedType, &instanceWithInferedType.For_part, formGroup, probe)
		AssociationSliceToForm("Directive", instanceWithInferedType, &instanceWithInferedType.Directive, formGroup, probe)
		AssociationSliceToForm("Measure_style", instanceWithInferedType, &instanceWithInferedType.Measure_style, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Backup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Backup"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Backup",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Backup",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Backup"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Backup",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Backup",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Bar_style_color:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bar_style", instanceWithInferedType.Bar_style, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		AssociationFieldToForm("Wavy_line", instanceWithInferedType.Wavy_line, formGroup, probe)
		AssociationFieldToForm("Segno_1", instanceWithInferedType.Segno_1, formGroup, probe)
		AssociationFieldToForm("Coda_1", instanceWithInferedType.Coda_1, formGroup, probe)
		AssociationFieldToForm("Fermata", instanceWithInferedType.Fermata, formGroup, probe)
		AssociationFieldToForm("Ending", instanceWithInferedType.Ending, formGroup, probe)
		AssociationFieldToForm("Repeat", instanceWithInferedType.Repeat, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Barline"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Barline",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Barline",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Barline"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Barline",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Barline",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Barre:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Arrangement", instanceWithInferedType.Arrangement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bass_separator", instanceWithInferedType.Bass_separator, formGroup, probe)
		AssociationFieldToForm("Bass_step", instanceWithInferedType.Bass_step, formGroup, probe)
		AssociationFieldToForm("Bass_alter", instanceWithInferedType.Bass_alter, formGroup, probe)

	case *models.Bass_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Repeater", instanceWithInferedType.Repeater, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Fan", instanceWithInferedType.Fan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beat_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slashes", instanceWithInferedType.Slashes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Use_dots", instanceWithInferedType.Use_dots, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Slash_type", instanceWithInferedType.Slash_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slash_dot", instanceWithInferedType.Slash_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Except_voice", instanceWithInferedType.Except_voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beat_unit_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Beat_unit", instanceWithInferedType.Beat_unit, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beat_unit_dot", instanceWithInferedType.Beat_unit_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome"
			rf.Fieldname = "Beat_unit_tied"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Metronome),
					"Beat_unit_tied",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Metronome](
					nil,
					"Beat_unit_tied",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Beater:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tip", instanceWithInferedType.Tip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Bend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Shape", instanceWithInferedType.Shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("First_beat", instanceWithInferedType.First_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bend_alter", instanceWithInferedType.Bend_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pre_bend", instanceWithInferedType.Pre_bend, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Release", instanceWithInferedType.Release, formGroup, probe)
		AssociationFieldToForm("With_bar", instanceWithInferedType.With_bar, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Bend"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Bend",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Bend",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Bookmark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Position", instanceWithInferedType.Position, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Bookmark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Bookmark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Bookmark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit](
					nil,
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Bracket:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_end", instanceWithInferedType.Line_end, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End_length", instanceWithInferedType.End_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Breath_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Breath_mark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Breath_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Breath_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Caesura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Caesura"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Caesura",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Caesura",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Cancel:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Additional", instanceWithInferedType.Additional, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("After_barline", instanceWithInferedType.After_barline, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Clef_octave_change", instanceWithInferedType.Clef_octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Clef"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Clef",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Clef",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Coda:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Coda"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Coda",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Coda",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Credit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Page", instanceWithInferedType.Page, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Credit_type", instanceWithInferedType.Credit_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Credit_image", instanceWithInferedType.Credit_image, formGroup, probe)
		AssociationSliceToForm("Link", instanceWithInferedType, &instanceWithInferedType.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", instanceWithInferedType, &instanceWithInferedType.Bookmark, formGroup, probe)
		AssociationSliceToForm("Credit_words", instanceWithInferedType, &instanceWithInferedType.Credit_words, formGroup, probe)
		AssociationSliceToForm("Credit_symbol", instanceWithInferedType, &instanceWithInferedType.Credit_symbol, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_partwise"
			rf.Fieldname = "Credit"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_partwise),
					"Credit",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_partwise](
					nil,
					"Credit",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_timewise"
			rf.Fieldname = "Credit"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_timewise),
					"Credit",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_timewise](
					nil,
					"Credit",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Dashes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Defaults:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Scaling", instanceWithInferedType.Scaling, formGroup, probe)
		BasicFieldtoForm("Concert_score", instanceWithInferedType.Concert_score, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Page_layout", instanceWithInferedType.Page_layout, formGroup, probe)
		AssociationFieldToForm("System_layout", instanceWithInferedType.System_layout, formGroup, probe)
		AssociationSliceToForm("Staff_layout", instanceWithInferedType, &instanceWithInferedType.Staff_layout, formGroup, probe)
		AssociationFieldToForm("Appearance", instanceWithInferedType.Appearance, formGroup, probe)
		AssociationFieldToForm("Music_font", instanceWithInferedType.Music_font, formGroup, probe)
		AssociationFieldToForm("Word_font", instanceWithInferedType.Word_font, formGroup, probe)
		AssociationSliceToForm("Lyric_font", instanceWithInferedType, &instanceWithInferedType.Lyric_font, formGroup, probe)
		AssociationSliceToForm("Lyric_language", instanceWithInferedType, &instanceWithInferedType.Lyric_language, formGroup, probe)

	case *models.Degree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("Degree_value", instanceWithInferedType.Degree_value, formGroup, probe)
		AssociationFieldToForm("Degree_alter", instanceWithInferedType.Degree_alter, formGroup, probe)
		AssociationFieldToForm("Degree_type", instanceWithInferedType.Degree_type, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Harmony"
			rf.Fieldname = "Degree"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Harmony),
					"Degree",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Harmony](
					nil,
					"Degree",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Degree_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Plus_minus", instanceWithInferedType.Plus_minus, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Degree_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Degree_value:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Direction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Directive", instanceWithInferedType.Directive, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Direction_type", instanceWithInferedType, &instanceWithInferedType.Direction_type, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Sound", instanceWithInferedType.Sound, formGroup, probe)
		AssociationFieldToForm("Listening", instanceWithInferedType.Listening, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Direction"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Direction",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Direction",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Direction"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Direction",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Direction",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Direction_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Rehearsal", instanceWithInferedType, &instanceWithInferedType.Rehearsal, formGroup, probe)
		AssociationSliceToForm("Segno", instanceWithInferedType, &instanceWithInferedType.Segno, formGroup, probe)
		AssociationSliceToForm("Coda", instanceWithInferedType, &instanceWithInferedType.Coda, formGroup, probe)
		AssociationSliceToForm("Words", instanceWithInferedType, &instanceWithInferedType.Words, formGroup, probe)
		AssociationSliceToForm("Symbol", instanceWithInferedType, &instanceWithInferedType.Symbol, formGroup, probe)
		AssociationFieldToForm("Wedge", instanceWithInferedType.Wedge, formGroup, probe)
		AssociationSliceToForm("Dynamics", instanceWithInferedType, &instanceWithInferedType.Dynamics, formGroup, probe)
		AssociationFieldToForm("Dashes", instanceWithInferedType.Dashes, formGroup, probe)
		AssociationFieldToForm("Bracket", instanceWithInferedType.Bracket, formGroup, probe)
		AssociationFieldToForm("Pedal", instanceWithInferedType.Pedal, formGroup, probe)
		AssociationFieldToForm("Metronome", instanceWithInferedType.Metronome, formGroup, probe)
		AssociationFieldToForm("Octave_shift", instanceWithInferedType.Octave_shift, formGroup, probe)
		AssociationFieldToForm("Harp_pedals", instanceWithInferedType.Harp_pedals, formGroup, probe)
		AssociationFieldToForm("Damp", instanceWithInferedType.Damp, formGroup, probe)
		AssociationFieldToForm("Damp_all", instanceWithInferedType.Damp_all, formGroup, probe)
		AssociationFieldToForm("Eyeglasses", instanceWithInferedType.Eyeglasses, formGroup, probe)
		AssociationFieldToForm("String_mute", instanceWithInferedType.String_mute, formGroup, probe)
		AssociationFieldToForm("Scordatura", instanceWithInferedType.Scordatura, formGroup, probe)
		AssociationFieldToForm("Image", instanceWithInferedType.Image, formGroup, probe)
		AssociationFieldToForm("Principal_voice", instanceWithInferedType.Principal_voice, formGroup, probe)
		AssociationSliceToForm("Percussion", instanceWithInferedType, &instanceWithInferedType.Percussion, formGroup, probe)
		AssociationFieldToForm("Accordion_registration", instanceWithInferedType.Accordion_registration, formGroup, probe)
		AssociationFieldToForm("Staff_divide", instanceWithInferedType.Staff_divide, formGroup, probe)
		AssociationFieldToForm("Other_direction", instanceWithInferedType.Other_direction, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction"
			rf.Fieldname = "Direction_type"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction),
					"Direction_type",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction](
					nil,
					"Direction_type",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Distance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Distance"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Distance",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance](
					nil,
					"Distance",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Double:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Above", instanceWithInferedType.Above, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Dynamics:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("P", instanceWithInferedType.P, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pp", instanceWithInferedType.Pp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ppp", instanceWithInferedType.Ppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pppp", instanceWithInferedType.Pppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ppppp", instanceWithInferedType.Ppppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pppppp", instanceWithInferedType.Pppppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("F", instanceWithInferedType.F, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ff", instanceWithInferedType.Ff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fff", instanceWithInferedType.Fff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ffff", instanceWithInferedType.Ffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fffff", instanceWithInferedType.Fffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ffffff", instanceWithInferedType.Ffffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Mp", instanceWithInferedType.Mp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Mf", instanceWithInferedType.Mf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sf", instanceWithInferedType.Sf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sfp", instanceWithInferedType.Sfp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sfpp", instanceWithInferedType.Sfpp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fp", instanceWithInferedType.Fp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rf", instanceWithInferedType.Rf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rfz", instanceWithInferedType.Rfz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sfz", instanceWithInferedType.Sfz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sffz", instanceWithInferedType.Sffz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fz", instanceWithInferedType.Fz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("N", instanceWithInferedType.N, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pf", instanceWithInferedType.Pf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sfzp", instanceWithInferedType.Sfzp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Other_dynamics", instanceWithInferedType, &instanceWithInferedType.Other_dynamics, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Dynamics"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Dynamics"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Effect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Elision:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lyric"
			rf.Fieldname = "Elision"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lyric),
					"Elision",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lyric](
					nil,
					"Elision",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_shape", instanceWithInferedType.Line_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_length", instanceWithInferedType.Line_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Scoop"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Scoop",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Scoop",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Plop"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Plop",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Plop",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Doit"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Doit",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Doit",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Falloff"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Falloff",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Falloff",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty_placement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Accent"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Accent",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Accent",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Staccato"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Staccato",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Staccato",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Tenuto"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Tenuto",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Tenuto",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Detached_legato"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Detached_legato",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Detached_legato",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Staccatissimo"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Staccatissimo",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Staccatissimo",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Spiccato"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Spiccato",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Spiccato",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Stress"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Stress",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Stress",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Unstress"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Unstress",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Unstress",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Soft_accent"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Soft_accent",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Soft_accent",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Dot"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Dot",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note](
					nil,
					"Dot",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Schleifer"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Schleifer",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Schleifer",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Up_bow"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Up_bow",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Up_bow",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Down_bow"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Down_bow",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Down_bow",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Open_string"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Open_string",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Open_string",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Thumb_position"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Thumb_position",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Thumb_position",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Double_tongue"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Double_tongue",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Double_tongue",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Triple_tongue"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Triple_tongue",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Triple_tongue",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Snap_pizzicato"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Snap_pizzicato",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Snap_pizzicato",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Fingernails"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Fingernails",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Fingernails",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Brass_bend"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Brass_bend",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Brass_bend",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Flip"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Flip",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Flip",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Smear"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Smear",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Smear",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Golpe"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Golpe",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Golpe",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty_placement_smufl:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Stopped"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Stopped",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Stopped",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Open"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Open",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Open",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Half_muted"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Half_muted",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Half_muted",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty_print_object_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style_align_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_trill_sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Trill_mark"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Trill_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Trill_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Vertical_turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Vertical_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Vertical_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Inverted_vertical_turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Inverted_vertical_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Inverted_vertical_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Shake"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Shake",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Shake",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Haydn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Haydn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Haydn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Encoding:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Encoder", instanceWithInferedType, &instanceWithInferedType.Encoder, formGroup, probe)
		BasicFieldtoForm("Software", instanceWithInferedType.Software, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Encoding_description", instanceWithInferedType.Encoding_description, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Supports", instanceWithInferedType, &instanceWithInferedType.Supports, formGroup, probe)

	case *models.Ending:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End_length", instanceWithInferedType.End_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text_x", instanceWithInferedType.Text_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text_y", instanceWithInferedType.Text_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Extend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Feature:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Grouping"
			rf.Fieldname = "Feature"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Grouping),
					"Feature",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Grouping](
					nil,
					"Feature",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Fermata:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Fermata"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Fermata",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Fermata",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Figure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Prefix", instanceWithInferedType.Prefix, formGroup, probe)
		AssociationFieldToForm("Figure_number", instanceWithInferedType.Figure_number, formGroup, probe)
		AssociationFieldToForm("Suffix", instanceWithInferedType.Suffix, formGroup, probe)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Figured_bass"
			rf.Fieldname = "Figure"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Figured_bass),
					"Figure",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Figured_bass](
					nil,
					"Figure",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Figured_bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_dot", instanceWithInferedType.Print_dot, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_lyric", instanceWithInferedType.Print_lyric, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Figure", instanceWithInferedType, &instanceWithInferedType.Figure, formGroup, probe)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Figured_bass"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Figured_bass",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Figured_bass",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Figured_bass"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Figured_bass",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Figured_bass",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Fingering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Substitution", instanceWithInferedType.Substitution, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Alternate", instanceWithInferedType.Alternate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Fingering"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Fingering",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Fingering",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.First_fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.For_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Part_clef", instanceWithInferedType.Part_clef, formGroup, probe)
		AssociationFieldToForm("Part_transpose", instanceWithInferedType.Part_transpose, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "For_part"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"For_part",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"For_part",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Formatted_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Formatted_symbol_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Credit_symbol"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Credit_symbol",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit](
					nil,
					"Credit_symbol",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Symbol"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Symbol",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Symbol",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Formatted_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Name_display"
			rf.Fieldname = "Display_text"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Name_display),
					"Display_text",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Name_display](
					nil,
					"Display_text",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notehead_text"
			rf.Fieldname = "Display_text"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notehead_text),
					"Display_text",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notehead_text](
					nil,
					"Display_text",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Formatted_text_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Credit_words"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Credit_words",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit](
					nil,
					"Credit_words",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Rehearsal"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Rehearsal",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Rehearsal",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Words"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Words",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Words",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Forward:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Forward"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Forward",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Forward",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Forward"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Forward",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Forward",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Frame:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Unplayed", instanceWithInferedType.Unplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Frame_strings", instanceWithInferedType.Frame_strings, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Frame_frets", instanceWithInferedType.Frame_frets, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("First_fret", instanceWithInferedType.First_fret, formGroup, probe)
		AssociationSliceToForm("Frame_note", instanceWithInferedType, &instanceWithInferedType.Frame_note, formGroup, probe)

	case *models.Frame_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("String", instanceWithInferedType.String, formGroup, probe)
		AssociationFieldToForm("Fret", instanceWithInferedType.Fret, formGroup, probe)
		AssociationFieldToForm("Fingering", instanceWithInferedType.Fingering, formGroup, probe)
		AssociationFieldToForm("Barre", instanceWithInferedType.Barre, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Frame"
			rf.Fieldname = "Frame_note"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Frame),
					"Frame_note",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Frame](
					nil,
					"Frame_note",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Fret"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Fret",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Fret",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Glass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Glissando:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Glissando"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Glissando",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Glissando",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Glyph:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Glyph"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Glyph",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance](
					nil,
					"Glyph",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Grace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Steal_time_previous", instanceWithInferedType.Steal_time_previous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Steal_time_following", instanceWithInferedType.Steal_time_following, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Make_time", instanceWithInferedType.Make_time, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Slash", instanceWithInferedType.Slash, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Group_barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Group_name:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Group_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Grouping:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Member_of", instanceWithInferedType.Member_of, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Feature", instanceWithInferedType, &instanceWithInferedType.Feature, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Grouping"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Grouping",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Grouping",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Grouping"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Grouping",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Grouping",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Hammer_on_pull_off:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Hammer_on"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Hammer_on",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Hammer_on",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Pull_off"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Pull_off",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Pull_off",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Handbell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Handbell"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Handbell",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Handbell",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Harmon_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Harmon_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Harmon_closed", instanceWithInferedType.Harmon_closed, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Harmon_mute"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Harmon_mute",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Harmon_mute",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Harmonic:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Natural", instanceWithInferedType.Natural, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Artificial", instanceWithInferedType.Artificial, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base_pitch", instanceWithInferedType.Base_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Touching_pitch", instanceWithInferedType.Touching_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sounding_pitch", instanceWithInferedType.Sounding_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Harmonic"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Harmonic",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Harmonic",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Harmony:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_frame", instanceWithInferedType.Print_frame, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Arrangement", instanceWithInferedType.Arrangement, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Root", instanceWithInferedType.Root, formGroup, probe)
		AssociationFieldToForm("Numeral", instanceWithInferedType.Numeral, formGroup, probe)
		AssociationFieldToForm("Function", instanceWithInferedType.Function, formGroup, probe)
		AssociationFieldToForm("Kind", instanceWithInferedType.Kind, formGroup, probe)
		AssociationFieldToForm("Inversion", instanceWithInferedType.Inversion, formGroup, probe)
		AssociationFieldToForm("Bass", instanceWithInferedType.Bass, formGroup, probe)
		AssociationSliceToForm("Degree", instanceWithInferedType, &instanceWithInferedType.Degree, formGroup, probe)
		AssociationFieldToForm("Frame", instanceWithInferedType.Frame, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Harmony"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Harmony",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Harmony",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Harmony"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Harmony",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Harmony",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Harmony_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Harp_pedals:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Pedal_tuning", instanceWithInferedType, &instanceWithInferedType.Pedal_tuning, formGroup, probe)

	case *models.Heel_toe:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Heel"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Heel",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Heel",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Toe"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Toe",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Toe",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Hole:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Hole_type", instanceWithInferedType.Hole_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Hole_closed", instanceWithInferedType.Hole_closed, formGroup, probe)
		BasicFieldtoForm("Hole_shape", instanceWithInferedType.Hole_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Hole"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Hole",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Hole",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Hole_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Horizontal_turn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Slash", instanceWithInferedType.Slash, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Delayed_turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Delayed_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Delayed_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Inverted_turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Inverted_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Inverted_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Delayed_inverted_turn"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Delayed_inverted_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Delayed_inverted_turn",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Identification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Creator", instanceWithInferedType, &instanceWithInferedType.Creator, formGroup, probe)
		AssociationSliceToForm("Rights", instanceWithInferedType, &instanceWithInferedType.Rights, formGroup, probe)
		AssociationFieldToForm("Encoding", instanceWithInferedType.Encoding, formGroup, probe)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Relation", instanceWithInferedType, &instanceWithInferedType.Relation, formGroup, probe)
		AssociationFieldToForm("Miscellaneous", instanceWithInferedType.Miscellaneous, formGroup, probe)

	case *models.Image:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Instrument"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note](
					nil,
					"Instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Instrument_change:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_sound", instanceWithInferedType.Instrument_sound, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Solo", instanceWithInferedType.Solo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ensemble", instanceWithInferedType.Ensemble, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Virtual_instrument", instanceWithInferedType.Virtual_instrument, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sound"
			rf.Fieldname = "Instrument_change"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sound),
					"Instrument_change",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sound](
					nil,
					"Instrument_change",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Instrument_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Part_link"
			rf.Fieldname = "Instrument_link"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Part_link),
					"Instrument_link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Part_link](
					nil,
					"Instrument_link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Interchangeable:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Separator", instanceWithInferedType.Separator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Time_relation", instanceWithInferedType.Time_relation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Beat_type", instanceWithInferedType.Beat_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Inversion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Cancel", instanceWithInferedType.Cancel, formGroup, probe)
		BasicFieldtoForm("Fifths", instanceWithInferedType.Fifths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Mode", instanceWithInferedType.Mode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Key_step", instanceWithInferedType.Key_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Key_alter", instanceWithInferedType.Key_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Key_accidental", instanceWithInferedType.Key_accidental, formGroup, probe)
		AssociationSliceToForm("Key_octave", instanceWithInferedType, &instanceWithInferedType.Key_octave, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Key"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Key",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Key",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Key_accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Key_octave:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Cancel", instanceWithInferedType.Cancel, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Key"
			rf.Fieldname = "Key_octave"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Key),
					"Key_octave",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Key](
					nil,
					"Key_octave",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Kind:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Use_symbols", instanceWithInferedType.Use_symbols, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Stack_degrees", instanceWithInferedType.Stack_degrees, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses_degrees", instanceWithInferedType.Parentheses_degrees, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket_degrees", instanceWithInferedType.Bracket_degrees, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Level:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Reference", instanceWithInferedType.Reference, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Line_detail:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Staff_details"
			rf.Fieldname = "Line_detail"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Staff_details),
					"Line_detail",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Staff_details](
					nil,
					"Line_detail",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Line_width:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Line_width"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Line_width",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance](
					nil,
					"Line_width",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Position", instanceWithInferedType.Position, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Link"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Link"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Link"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit](
					nil,
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Listen:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Assess", instanceWithInferedType, &instanceWithInferedType.Assess, formGroup, probe)
		AssociationSliceToForm("Wait", instanceWithInferedType, &instanceWithInferedType.Wait, formGroup, probe)
		AssociationSliceToForm("Other_listen", instanceWithInferedType, &instanceWithInferedType.Other_listen, formGroup, probe)

	case *models.Listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Sync", instanceWithInferedType, &instanceWithInferedType.Sync, formGroup, probe)
		AssociationSliceToForm("Other_listening", instanceWithInferedType, &instanceWithInferedType.Other_listening, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Listening"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Listening",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Listening",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Listening"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Listening",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Listening",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lyric:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Elision", instanceWithInferedType, &instanceWithInferedType.Elision, formGroup, probe)
		BasicFieldtoForm("Syllabic", instanceWithInferedType.Syllabic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Text", instanceWithInferedType, &instanceWithInferedType.Text, formGroup, probe)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		BasicFieldtoForm("Laughing", instanceWithInferedType.Laughing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Humming", instanceWithInferedType.Humming, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End_line", instanceWithInferedType.End_line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End_paragraph", instanceWithInferedType.End_paragraph, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Lyric"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Lyric",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note](
					nil,
					"Lyric",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lyric_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_font"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Defaults),
					"Lyric_font",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Defaults](
					nil,
					"Lyric_font",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lyric_language:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_language"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Defaults),
					"Lyric_language",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Defaults](
					nil,
					"Lyric_language",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Measure_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Measure_distance", instanceWithInferedType.Measure_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_numbering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Multiple_rest_always", instanceWithInferedType.Multiple_rest_always, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Multiple_rest_range", instanceWithInferedType.Multiple_rest_range, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slashes", instanceWithInferedType.Slashes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Multiple_rest", instanceWithInferedType.Multiple_rest, formGroup, probe)
		AssociationFieldToForm("Measure_repeat", instanceWithInferedType.Measure_repeat, formGroup, probe)
		AssociationFieldToForm("Beat_repeat", instanceWithInferedType.Beat_repeat, formGroup, probe)
		AssociationFieldToForm("Slash", instanceWithInferedType.Slash, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Measure_style"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Measure_style",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Measure_style",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Membrane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metronome:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Beat_unit", instanceWithInferedType.Beat_unit, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beat_unit_dot", instanceWithInferedType.Beat_unit_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Per_minute", instanceWithInferedType.Per_minute, formGroup, probe)
		AssociationSliceToForm("Beat_unit_tied", instanceWithInferedType, &instanceWithInferedType.Beat_unit_tied, formGroup, probe)
		BasicFieldtoForm("Metronome_arrows", instanceWithInferedType.Metronome_arrows, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Metronome_relation", instanceWithInferedType.Metronome_relation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Metronome_note", instanceWithInferedType, &instanceWithInferedType.Metronome_note, formGroup, probe)

	case *models.Metronome_beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome_note"
			rf.Fieldname = "Metronome_beam"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Metronome_note),
					"Metronome_beam",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Metronome_note](
					nil,
					"Metronome_beam",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Metronome_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Metronome_type", instanceWithInferedType.Metronome_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Metronome_dot", instanceWithInferedType.Metronome_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Metronome_beam", instanceWithInferedType, &instanceWithInferedType.Metronome_beam, formGroup, probe)
		AssociationFieldToForm("Metronome_tied", instanceWithInferedType.Metronome_tied, formGroup, probe)
		AssociationFieldToForm("Metronome_tuplet", instanceWithInferedType.Metronome_tuplet, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome"
			rf.Fieldname = "Metronome_note"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Metronome),
					"Metronome_note",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Metronome](
					nil,
					"Metronome_note",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Metronome_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Metronome_tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Midi_device:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Port", instanceWithInferedType.Port, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Midi_device"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Midi_device",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part](
					nil,
					"Midi_device",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sound"
			rf.Fieldname = "Midi_device"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sound),
					"Midi_device",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sound](
					nil,
					"Midi_device",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Midi_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_channel", instanceWithInferedType.Midi_channel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_name", instanceWithInferedType.Midi_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_bank", instanceWithInferedType.Midi_bank, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_program", instanceWithInferedType.Midi_program, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_unpitched", instanceWithInferedType.Midi_unpitched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Volume", instanceWithInferedType.Volume, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Pan", instanceWithInferedType.Pan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Elevation", instanceWithInferedType.Elevation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Midi_instrument"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Midi_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part](
					nil,
					"Midi_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sound"
			rf.Fieldname = "Midi_instrument"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sound),
					"Midi_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sound](
					nil,
					"Midi_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Miscellaneous:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Miscellaneous_field", instanceWithInferedType, &instanceWithInferedType.Miscellaneous_field, formGroup, probe)

	case *models.Miscellaneous_field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Miscellaneous"
			rf.Fieldname = "Miscellaneous_field"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Miscellaneous),
					"Miscellaneous_field",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Miscellaneous](
					nil,
					"Miscellaneous_field",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Mordent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Mordent"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Mordent",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Mordent",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Inverted_mordent"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Inverted_mordent",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Inverted_mordent",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Multiple_rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Use_symbols", instanceWithInferedType.Use_symbols, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Name_display:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Display_text", instanceWithInferedType, &instanceWithInferedType.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", instanceWithInferedType, &instanceWithInferedType.Accidental_text, formGroup, probe)

	case *models.Non_arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Non_arpeggiate"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Non_arpeggiate",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Non_arpeggiate",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Notations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		AssociationSliceToForm("Tied", instanceWithInferedType, &instanceWithInferedType.Tied, formGroup, probe)
		AssociationSliceToForm("Slur", instanceWithInferedType, &instanceWithInferedType.Slur, formGroup, probe)
		AssociationSliceToForm("Tuplet", instanceWithInferedType, &instanceWithInferedType.Tuplet, formGroup, probe)
		AssociationSliceToForm("Glissando", instanceWithInferedType, &instanceWithInferedType.Glissando, formGroup, probe)
		AssociationSliceToForm("Slide", instanceWithInferedType, &instanceWithInferedType.Slide, formGroup, probe)
		AssociationSliceToForm("Ornaments", instanceWithInferedType, &instanceWithInferedType.Ornaments, formGroup, probe)
		AssociationSliceToForm("Technical", instanceWithInferedType, &instanceWithInferedType.Technical, formGroup, probe)
		AssociationSliceToForm("Articulations", instanceWithInferedType, &instanceWithInferedType.Articulations, formGroup, probe)
		AssociationSliceToForm("Dynamics", instanceWithInferedType, &instanceWithInferedType.Dynamics, formGroup, probe)
		AssociationSliceToForm("Fermata", instanceWithInferedType, &instanceWithInferedType.Fermata, formGroup, probe)
		AssociationSliceToForm("Arpeggiate", instanceWithInferedType, &instanceWithInferedType.Arpeggiate, formGroup, probe)
		AssociationSliceToForm("Non_arpeggiate", instanceWithInferedType, &instanceWithInferedType.Non_arpeggiate, formGroup, probe)
		AssociationSliceToForm("Accidental_mark", instanceWithInferedType, &instanceWithInferedType.Accidental_mark, formGroup, probe)
		AssociationSliceToForm("Other_notation", instanceWithInferedType, &instanceWithInferedType.Other_notation, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Notations"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Notations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note](
					nil,
					"Notations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_leger", instanceWithInferedType.Print_leger, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Dynamics", instanceWithInferedType.Dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End_dynamics", instanceWithInferedType.End_dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Attack", instanceWithInferedType.Attack, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Release", instanceWithInferedType.Release, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Pizzicato", instanceWithInferedType.Pizzicato, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_dot", instanceWithInferedType.Print_dot, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_lyric", instanceWithInferedType.Print_lyric, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Grace", instanceWithInferedType.Grace, formGroup, probe)
		BasicFieldtoForm("Chord", instanceWithInferedType.Chord, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Pitch", instanceWithInferedType.Pitch, formGroup, probe)
		AssociationFieldToForm("Unpitched", instanceWithInferedType.Unpitched, formGroup, probe)
		AssociationFieldToForm("Rest", instanceWithInferedType.Rest, formGroup, probe)
		AssociationFieldToForm("Tie", instanceWithInferedType.Tie, formGroup, probe)
		BasicFieldtoForm("Cue", instanceWithInferedType.Cue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Instrument", instanceWithInferedType, &instanceWithInferedType.Instrument, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Type", instanceWithInferedType.Type, formGroup, probe)
		AssociationSliceToForm("Dot", instanceWithInferedType, &instanceWithInferedType.Dot, formGroup, probe)
		AssociationFieldToForm("Accidental", instanceWithInferedType.Accidental, formGroup, probe)
		AssociationFieldToForm("Time_modification", instanceWithInferedType.Time_modification, formGroup, probe)
		AssociationFieldToForm("Stem", instanceWithInferedType.Stem, formGroup, probe)
		AssociationFieldToForm("Notehead", instanceWithInferedType.Notehead, formGroup, probe)
		AssociationFieldToForm("Notehead_text", instanceWithInferedType.Notehead_text, formGroup, probe)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Beam", instanceWithInferedType.Beam, formGroup, probe)
		AssociationSliceToForm("Notations", instanceWithInferedType, &instanceWithInferedType.Notations, formGroup, probe)
		AssociationSliceToForm("Lyric", instanceWithInferedType, &instanceWithInferedType.Lyric, formGroup, probe)
		AssociationFieldToForm("Play", instanceWithInferedType.Play, formGroup, probe)
		AssociationFieldToForm("Listen", instanceWithInferedType.Listen, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Note"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Note",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Note",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Note"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Note",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Note",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Note_size"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Note_size",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance](
					nil,
					"Note_size",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Notehead:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Filled", instanceWithInferedType.Filled, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Notehead_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Display_text", instanceWithInferedType, &instanceWithInferedType.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", instanceWithInferedType, &instanceWithInferedType.Accidental_text, formGroup, probe)

	case *models.Numeral:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Numeral_root", instanceWithInferedType.Numeral_root, formGroup, probe)
		AssociationFieldToForm("Numeral_alter", instanceWithInferedType.Numeral_alter, formGroup, probe)
		AssociationFieldToForm("Numeral_key", instanceWithInferedType.Numeral_key, formGroup, probe)

	case *models.Numeral_key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Numeral_fifths", instanceWithInferedType.Numeral_fifths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Numeral_mode", instanceWithInferedType.Numeral_mode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Numeral_root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Octave_shift:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Offset:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Sound", instanceWithInferedType.Sound, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Opus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Ornaments:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Trill_mark", instanceWithInferedType, &instanceWithInferedType.Trill_mark, formGroup, probe)
		AssociationSliceToForm("Turn", instanceWithInferedType, &instanceWithInferedType.Turn, formGroup, probe)
		AssociationSliceToForm("Delayed_turn", instanceWithInferedType, &instanceWithInferedType.Delayed_turn, formGroup, probe)
		AssociationSliceToForm("Inverted_turn", instanceWithInferedType, &instanceWithInferedType.Inverted_turn, formGroup, probe)
		AssociationSliceToForm("Delayed_inverted_turn", instanceWithInferedType, &instanceWithInferedType.Delayed_inverted_turn, formGroup, probe)
		AssociationSliceToForm("Vertical_turn", instanceWithInferedType, &instanceWithInferedType.Vertical_turn, formGroup, probe)
		AssociationSliceToForm("Inverted_vertical_turn", instanceWithInferedType, &instanceWithInferedType.Inverted_vertical_turn, formGroup, probe)
		AssociationSliceToForm("Shake", instanceWithInferedType, &instanceWithInferedType.Shake, formGroup, probe)
		AssociationSliceToForm("Wavy_line", instanceWithInferedType, &instanceWithInferedType.Wavy_line, formGroup, probe)
		AssociationSliceToForm("Mordent", instanceWithInferedType, &instanceWithInferedType.Mordent, formGroup, probe)
		AssociationSliceToForm("Inverted_mordent", instanceWithInferedType, &instanceWithInferedType.Inverted_mordent, formGroup, probe)
		AssociationSliceToForm("Schleifer", instanceWithInferedType, &instanceWithInferedType.Schleifer, formGroup, probe)
		AssociationSliceToForm("Tremolo", instanceWithInferedType, &instanceWithInferedType.Tremolo, formGroup, probe)
		AssociationSliceToForm("Haydn", instanceWithInferedType, &instanceWithInferedType.Haydn, formGroup, probe)
		AssociationSliceToForm("Other_ornament", instanceWithInferedType, &instanceWithInferedType.Other_ornament, formGroup, probe)
		AssociationSliceToForm("Accidental_mark", instanceWithInferedType, &instanceWithInferedType.Accidental_mark, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Ornaments"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Ornaments",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Ornaments",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Other_appearance"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Other_appearance",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance](
					nil,
					"Other_appearance",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_direction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Other_listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Listen"
			rf.Fieldname = "Other_listen"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Listen),
					"Other_listen",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Listen](
					nil,
					"Other_listen",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Listening"
			rf.Fieldname = "Other_listening"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Listening),
					"Other_listening",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Listening](
					nil,
					"Other_listening",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_notation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Other_notation"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Other_notation",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Other_notation",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_placement_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Other_articulation"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Other_articulation",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Other_articulation",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Other_ornament"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Other_ornament",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Other_ornament",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Other_technical"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Other_technical",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Other_technical",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Play"
			rf.Fieldname = "Other_play"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Play),
					"Other_play",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Play](
					nil,
					"Other_play",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Dynamics"
			rf.Fieldname = "Other_dynamics"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Dynamics),
					"Other_dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Dynamics](
					nil,
					"Other_dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Page_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Page_height", instanceWithInferedType.Page_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Page_width", instanceWithInferedType.Page_width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Page_margins", instanceWithInferedType.Page_margins, formGroup, probe)

	case *models.Page_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Left_margin", instanceWithInferedType.Left_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Right_margin", instanceWithInferedType.Right_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Top_margin", instanceWithInferedType.Top_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bottom_margin", instanceWithInferedType.Bottom_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Clef_octave_change", instanceWithInferedType.Clef_octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Group_name", instanceWithInferedType.Group_name, formGroup, probe)
		AssociationFieldToForm("Group_name_display", instanceWithInferedType.Group_name_display, formGroup, probe)
		AssociationFieldToForm("Group_abbreviation", instanceWithInferedType.Group_abbreviation, formGroup, probe)
		AssociationFieldToForm("Group_abbreviation_display", instanceWithInferedType.Group_abbreviation_display, formGroup, probe)
		AssociationFieldToForm("Group_symbol", instanceWithInferedType.Group_symbol, formGroup, probe)
		AssociationFieldToForm("Group_barline", instanceWithInferedType.Group_barline, formGroup, probe)
		BasicFieldtoForm("Group_time", instanceWithInferedType.Group_time, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)

	case *models.Part_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Instrument_link", instanceWithInferedType, &instanceWithInferedType.Instrument_link, formGroup, probe)
		BasicFieldtoForm("Group_link", instanceWithInferedType.Group_link, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Part_link"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Part_link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part](
					nil,
					"Part_link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Part_list:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Part_group", instanceWithInferedType.Part_group, formGroup, probe)
		AssociationFieldToForm("Score_part", instanceWithInferedType.Score_part, formGroup, probe)

	case *models.Part_name:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Top_staff", instanceWithInferedType.Top_staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bottom_staff", instanceWithInferedType.Bottom_staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Part_transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Diatonic", instanceWithInferedType.Diatonic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Chromatic", instanceWithInferedType.Chromatic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Octave_change", instanceWithInferedType.Octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Double", instanceWithInferedType.Double, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pedal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Abbreviated", instanceWithInferedType.Abbreviated, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pedal_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Pedal_step", instanceWithInferedType.Pedal_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Pedal_alter", instanceWithInferedType.Pedal_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Harp_pedals"
			rf.Fieldname = "Pedal_tuning"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Harp_pedals),
					"Pedal_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Harp_pedals](
					nil,
					"Pedal_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Per_minute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Percussion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Glass", instanceWithInferedType.Glass, formGroup, probe)
		AssociationFieldToForm("Metal", instanceWithInferedType.Metal, formGroup, probe)
		AssociationFieldToForm("Wood", instanceWithInferedType.Wood, formGroup, probe)
		AssociationFieldToForm("Pitched", instanceWithInferedType.Pitched, formGroup, probe)
		AssociationFieldToForm("Membrane", instanceWithInferedType.Membrane, formGroup, probe)
		AssociationFieldToForm("Effect", instanceWithInferedType.Effect, formGroup, probe)
		AssociationFieldToForm("Timpani", instanceWithInferedType.Timpani, formGroup, probe)
		AssociationFieldToForm("Beater", instanceWithInferedType.Beater, formGroup, probe)
		AssociationFieldToForm("Stick", instanceWithInferedType.Stick, formGroup, probe)
		BasicFieldtoForm("Stick_location", instanceWithInferedType.Stick_location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Other_percussion", instanceWithInferedType.Other_percussion, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Percussion"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Percussion",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Percussion",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Pitch:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Step", instanceWithInferedType.Step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Alter", instanceWithInferedType.Alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Octave", instanceWithInferedType.Octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Placement_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Pluck"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Pluck",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Pluck",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ipa", instanceWithInferedType.Ipa, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Mute", instanceWithInferedType.Mute, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Semi_pitched", instanceWithInferedType.Semi_pitched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Other_play", instanceWithInferedType, &instanceWithInferedType.Other_play, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sound"
			rf.Fieldname = "Play"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sound),
					"Play",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sound](
					nil,
					"Play",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Player:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Player_name", instanceWithInferedType.Player_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Player"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Player",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part](
					nil,
					"Player",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Principal_voice:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Print:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff_spacing", instanceWithInferedType.Staff_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("New_system", instanceWithInferedType.New_system, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("New_page", instanceWithInferedType.New_page, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Blank_page", instanceWithInferedType.Blank_page, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Page_number", instanceWithInferedType.Page_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Page_layout", instanceWithInferedType.Page_layout, formGroup, probe)
		AssociationFieldToForm("System_layout", instanceWithInferedType.System_layout, formGroup, probe)
		AssociationSliceToForm("Staff_layout", instanceWithInferedType, &instanceWithInferedType.Staff_layout, formGroup, probe)
		AssociationFieldToForm("Measure_layout", instanceWithInferedType.Measure_layout, formGroup, probe)
		AssociationFieldToForm("Measure_numbering", instanceWithInferedType.Measure_numbering, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Print"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Print",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Print",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Print"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Print",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Print",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Release:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Times", instanceWithInferedType.Times, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("After_jump", instanceWithInferedType.After_jump, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Winged", instanceWithInferedType.Winged, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Measure", instanceWithInferedType.Measure, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Display_step", instanceWithInferedType.Display_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Display_octave", instanceWithInferedType.Display_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Root_step", instanceWithInferedType.Root_step, formGroup, probe)
		AssociationFieldToForm("Root_alter", instanceWithInferedType.Root_alter, formGroup, probe)

	case *models.Root_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Scaling:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Millimeters", instanceWithInferedType.Millimeters, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tenths", instanceWithInferedType.Tenths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Scordatura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Accord", instanceWithInferedType, &instanceWithInferedType.Accord, formGroup, probe)

	case *models.Score_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_name", instanceWithInferedType.Instrument_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_abbreviation", instanceWithInferedType.Instrument_abbreviation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_sound", instanceWithInferedType.Instrument_sound, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Solo", instanceWithInferedType.Solo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ensemble", instanceWithInferedType.Ensemble, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Virtual_instrument", instanceWithInferedType.Virtual_instrument, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Score_instrument"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Score_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part](
					nil,
					"Score_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Score_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationSliceToForm("Part_link", instanceWithInferedType, &instanceWithInferedType.Part_link, formGroup, probe)
		AssociationFieldToForm("Part_name", instanceWithInferedType.Part_name, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation", instanceWithInferedType.Part_abbreviation, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)
		BasicFieldtoForm("Group", instanceWithInferedType.Group, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Score_instrument", instanceWithInferedType, &instanceWithInferedType.Score_instrument, formGroup, probe)
		AssociationSliceToForm("Player", instanceWithInferedType, &instanceWithInferedType.Player, formGroup, probe)
		AssociationSliceToForm("Midi_device", instanceWithInferedType, &instanceWithInferedType.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", instanceWithInferedType, &instanceWithInferedType.Midi_instrument, formGroup, probe)

	case *models.Score_partwise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Version", instanceWithInferedType.Version, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Work", instanceWithInferedType.Work, formGroup, probe)
		BasicFieldtoForm("Movement_number", instanceWithInferedType.Movement_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Movement_title", instanceWithInferedType.Movement_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationFieldToForm("Defaults", instanceWithInferedType.Defaults, formGroup, probe)
		AssociationSliceToForm("Credit", instanceWithInferedType, &instanceWithInferedType.Credit, formGroup, probe)
		AssociationFieldToForm("Part_list", instanceWithInferedType.Part_list, formGroup, probe)
		AssociationSliceToForm("Part", instanceWithInferedType, &instanceWithInferedType.Part, formGroup, probe)

	case *models.Score_timewise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Version", instanceWithInferedType.Version, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Work", instanceWithInferedType.Work, formGroup, probe)
		BasicFieldtoForm("Movement_number", instanceWithInferedType.Movement_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Movement_title", instanceWithInferedType.Movement_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationFieldToForm("Defaults", instanceWithInferedType.Defaults, formGroup, probe)
		AssociationSliceToForm("Credit", instanceWithInferedType, &instanceWithInferedType.Credit, formGroup, probe)
		AssociationFieldToForm("Part_list", instanceWithInferedType.Part_list, formGroup, probe)
		AssociationSliceToForm("Measure", instanceWithInferedType, &instanceWithInferedType.Measure, formGroup, probe)

	case *models.Segno:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Segno"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Segno",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type](
					nil,
					"Segno",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Slash:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Use_dots", instanceWithInferedType.Use_dots, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Use_stems", instanceWithInferedType.Use_stems, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Slash_type", instanceWithInferedType.Slash_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slash_dot", instanceWithInferedType.Slash_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Except_voice", instanceWithInferedType.Except_voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Slide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("First_beat", instanceWithInferedType.First_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Slide"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Slide",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Slide",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Slur:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Orientation", instanceWithInferedType.Orientation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_x", instanceWithInferedType.Bezier_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_y", instanceWithInferedType.Bezier_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_x2", instanceWithInferedType.Bezier_x2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_y2", instanceWithInferedType.Bezier_y2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_offset", instanceWithInferedType.Bezier_offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_offset2", instanceWithInferedType.Bezier_offset2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Slur"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Slur",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Slur",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tempo", instanceWithInferedType.Tempo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dynamics", instanceWithInferedType.Dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Dacapo", instanceWithInferedType.Dacapo, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dalsegno", instanceWithInferedType.Dalsegno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tocoda", instanceWithInferedType.Tocoda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Forward_repeat", instanceWithInferedType.Forward_repeat, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Fine", instanceWithInferedType.Fine, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Pizzicato", instanceWithInferedType.Pizzicato, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Pan", instanceWithInferedType.Pan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Elevation", instanceWithInferedType.Elevation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Damper_pedal", instanceWithInferedType.Damper_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Soft_pedal", instanceWithInferedType.Soft_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Sostenuto_pedal", instanceWithInferedType.Sostenuto_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Instrument_change", instanceWithInferedType, &instanceWithInferedType.Instrument_change, formGroup, probe)
		AssociationSliceToForm("Midi_device", instanceWithInferedType, &instanceWithInferedType.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", instanceWithInferedType, &instanceWithInferedType.Midi_instrument, formGroup, probe)
		AssociationSliceToForm("Play", instanceWithInferedType, &instanceWithInferedType.Play, formGroup, probe)
		AssociationFieldToForm("Swing", instanceWithInferedType.Swing, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_measure"
			rf.Fieldname = "Sound"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_measure),
					"Sound",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_measure](
					nil,
					"Sound",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_part_1"
			rf.Fieldname = "Sound"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_part_1),
					"Sound",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_part_1](
					nil,
					"Sound",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Staff_details:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Show_frets", instanceWithInferedType.Show_frets, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Staff_type", instanceWithInferedType.Staff_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff_lines", instanceWithInferedType.Staff_lines, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Line_detail", instanceWithInferedType, &instanceWithInferedType.Line_detail, formGroup, probe)
		AssociationSliceToForm("Staff_tuning", instanceWithInferedType, &instanceWithInferedType.Staff_tuning, formGroup, probe)
		BasicFieldtoForm("Capo", instanceWithInferedType.Capo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Staff_size", instanceWithInferedType.Staff_size, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Staff_details"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Staff_details",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Staff_details",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Staff_divide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Staff_distance", instanceWithInferedType.Staff_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Staff_layout"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Defaults),
					"Staff_layout",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Defaults](
					nil,
					"Staff_layout",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Print"
			rf.Fieldname = "Staff_layout"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Print),
					"Staff_layout",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Print](
					nil,
					"Staff_layout",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Staff_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Scaling", instanceWithInferedType.Scaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Tuning_step", instanceWithInferedType.Tuning_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Tuning_alter", instanceWithInferedType.Tuning_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tuning_octave", instanceWithInferedType.Tuning_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Staff_details"
			rf.Fieldname = "Staff_tuning"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Staff_details),
					"Staff_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Staff_details](
					nil,
					"Staff_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Stem:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Stick:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Tip", instanceWithInferedType.Tip, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Dashed_circle", instanceWithInferedType.Dashed_circle, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Stick_type", instanceWithInferedType.Stick_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stick_material", instanceWithInferedType.Stick_material, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.String_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.String_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "String"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"String",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"String",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Strong_accent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Articulations"
			rf.Fieldname = "Strong_accent"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Articulations),
					"Strong_accent",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Articulations](
					nil,
					"Strong_accent",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Style_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Supports:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Attribute", instanceWithInferedType.Attribute, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Encoding"
			rf.Fieldname = "Supports"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Encoding),
					"Supports",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Encoding](
					nil,
					"Supports",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Swing:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Straight", instanceWithInferedType.Straight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("First", instanceWithInferedType.First, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Second", instanceWithInferedType.Second, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Swing_type", instanceWithInferedType.Swing_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Swing_style", instanceWithInferedType.Swing_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Sync:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Latency", instanceWithInferedType.Latency, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Listening"
			rf.Fieldname = "Sync"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Listening),
					"Sync",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Listening](
					nil,
					"Sync",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.System_dividers:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Left_divider", instanceWithInferedType.Left_divider, formGroup, probe)
		AssociationFieldToForm("Right_divider", instanceWithInferedType.Right_divider, formGroup, probe)

	case *models.System_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("System_margins", instanceWithInferedType.System_margins, formGroup, probe)
		BasicFieldtoForm("System_distance", instanceWithInferedType.System_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Top_system_distance", instanceWithInferedType.Top_system_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("System_dividers", instanceWithInferedType.System_dividers, formGroup, probe)

	case *models.System_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Left_margin", instanceWithInferedType.Left_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Right_margin", instanceWithInferedType.Right_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tap:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Hand", instanceWithInferedType.Hand, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Technical"
			rf.Fieldname = "Tap"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Technical),
					"Tap",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Technical](
					nil,
					"Tap",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Technical:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Up_bow", instanceWithInferedType, &instanceWithInferedType.Up_bow, formGroup, probe)
		AssociationSliceToForm("Down_bow", instanceWithInferedType, &instanceWithInferedType.Down_bow, formGroup, probe)
		AssociationSliceToForm("Harmonic", instanceWithInferedType, &instanceWithInferedType.Harmonic, formGroup, probe)
		AssociationSliceToForm("Open_string", instanceWithInferedType, &instanceWithInferedType.Open_string, formGroup, probe)
		AssociationSliceToForm("Thumb_position", instanceWithInferedType, &instanceWithInferedType.Thumb_position, formGroup, probe)
		AssociationSliceToForm("Fingering", instanceWithInferedType, &instanceWithInferedType.Fingering, formGroup, probe)
		AssociationSliceToForm("Pluck", instanceWithInferedType, &instanceWithInferedType.Pluck, formGroup, probe)
		AssociationSliceToForm("Double_tongue", instanceWithInferedType, &instanceWithInferedType.Double_tongue, formGroup, probe)
		AssociationSliceToForm("Triple_tongue", instanceWithInferedType, &instanceWithInferedType.Triple_tongue, formGroup, probe)
		AssociationSliceToForm("Stopped", instanceWithInferedType, &instanceWithInferedType.Stopped, formGroup, probe)
		AssociationSliceToForm("Snap_pizzicato", instanceWithInferedType, &instanceWithInferedType.Snap_pizzicato, formGroup, probe)
		AssociationSliceToForm("Fret", instanceWithInferedType, &instanceWithInferedType.Fret, formGroup, probe)
		AssociationSliceToForm("String", instanceWithInferedType, &instanceWithInferedType.String, formGroup, probe)
		AssociationSliceToForm("Hammer_on", instanceWithInferedType, &instanceWithInferedType.Hammer_on, formGroup, probe)
		AssociationSliceToForm("Pull_off", instanceWithInferedType, &instanceWithInferedType.Pull_off, formGroup, probe)
		AssociationSliceToForm("Bend", instanceWithInferedType, &instanceWithInferedType.Bend, formGroup, probe)
		AssociationSliceToForm("Tap", instanceWithInferedType, &instanceWithInferedType.Tap, formGroup, probe)
		AssociationSliceToForm("Heel", instanceWithInferedType, &instanceWithInferedType.Heel, formGroup, probe)
		AssociationSliceToForm("Toe", instanceWithInferedType, &instanceWithInferedType.Toe, formGroup, probe)
		AssociationSliceToForm("Fingernails", instanceWithInferedType, &instanceWithInferedType.Fingernails, formGroup, probe)
		AssociationSliceToForm("Hole", instanceWithInferedType, &instanceWithInferedType.Hole, formGroup, probe)
		AssociationSliceToForm("Arrow", instanceWithInferedType, &instanceWithInferedType.Arrow, formGroup, probe)
		AssociationSliceToForm("Handbell", instanceWithInferedType, &instanceWithInferedType.Handbell, formGroup, probe)
		AssociationSliceToForm("Brass_bend", instanceWithInferedType, &instanceWithInferedType.Brass_bend, formGroup, probe)
		AssociationSliceToForm("Flip", instanceWithInferedType, &instanceWithInferedType.Flip, formGroup, probe)
		AssociationSliceToForm("Smear", instanceWithInferedType, &instanceWithInferedType.Smear, formGroup, probe)
		AssociationSliceToForm("Open", instanceWithInferedType, &instanceWithInferedType.Open, formGroup, probe)
		AssociationSliceToForm("Half_muted", instanceWithInferedType, &instanceWithInferedType.Half_muted, formGroup, probe)
		AssociationSliceToForm("Harmon_mute", instanceWithInferedType, &instanceWithInferedType.Harmon_mute, formGroup, probe)
		AssociationSliceToForm("Golpe", instanceWithInferedType, &instanceWithInferedType.Golpe, formGroup, probe)
		AssociationSliceToForm("Other_technical", instanceWithInferedType, &instanceWithInferedType.Other_technical, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Technical"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Technical",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Technical",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Text_element_data:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lyric"
			rf.Fieldname = "Text"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lyric),
					"Text",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lyric](
					nil,
					"Text",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Orientation", instanceWithInferedType.Orientation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_x", instanceWithInferedType.Bezier_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_y", instanceWithInferedType.Bezier_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_x2", instanceWithInferedType.Bezier_x2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_y2", instanceWithInferedType.Bezier_y2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_offset", instanceWithInferedType.Bezier_offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Bezier_offset2", instanceWithInferedType.Bezier_offset2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Tied"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Tied",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Tied",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Time:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Separator", instanceWithInferedType.Separator, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Beat_type", instanceWithInferedType.Beat_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Interchangeable", instanceWithInferedType.Interchangeable, formGroup, probe)
		BasicFieldtoForm("Senza_misura", instanceWithInferedType.Senza_misura, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Time"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Time",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Time",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Time_modification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Actual_notes", instanceWithInferedType.Actual_notes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Normal_notes", instanceWithInferedType.Normal_notes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Normal_type", instanceWithInferedType.Normal_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Normal_dot", instanceWithInferedType.Normal_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Timpani:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Diatonic", instanceWithInferedType.Diatonic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Chromatic", instanceWithInferedType.Chromatic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Octave_change", instanceWithInferedType.Octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Double", instanceWithInferedType.Double, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Transpose"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Transpose",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes](
					nil,
					"Transpose",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tremolo:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Tremolo"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Tremolo",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Tremolo",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Show_number", instanceWithInferedType.Show_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Show_type", instanceWithInferedType.Show_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Line_shape", instanceWithInferedType.Line_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tuplet_actual", instanceWithInferedType.Tuplet_actual, formGroup, probe)
		AssociationFieldToForm("Tuplet_normal", instanceWithInferedType.Tuplet_normal, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Notations"
			rf.Fieldname = "Tuplet"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Notations),
					"Tuplet",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Notations](
					nil,
					"Tuplet",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tuplet_dot:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tuplet_portion"
			rf.Fieldname = "Tuplet_dot"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Tuplet_portion),
					"Tuplet_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Tuplet_portion](
					nil,
					"Tuplet_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tuplet_number:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tuplet_portion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tuplet_number", instanceWithInferedType.Tuplet_number, formGroup, probe)
		AssociationFieldToForm("Tuplet_type", instanceWithInferedType.Tuplet_type, formGroup, probe)
		AssociationSliceToForm("Tuplet_dot", instanceWithInferedType, &instanceWithInferedType.Tuplet_dot, formGroup, probe)

	case *models.Tuplet_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)

	case *models.Typed_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Encoding"
			rf.Fieldname = "Encoder"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Encoding),
					"Encoder",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Encoding](
					nil,
					"Encoder",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Creator"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Creator",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification](
					nil,
					"Creator",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Rights"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Rights",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification](
					nil,
					"Rights",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Relation"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Relation",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification](
					nil,
					"Relation",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Unpitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Display_step", instanceWithInferedType.Display_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Display_octave", instanceWithInferedType.Display_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Virtual_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Virtual_library", instanceWithInferedType.Virtual_library, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Virtual_name", instanceWithInferedType.Virtual_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wait:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Listen"
			rf.Fieldname = "Wait"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Listen),
					"Wait",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Listen](
					nil,
					"Wait",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Wavy_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Wavy_line"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Wavy_line",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments](
					nil,
					"Wavy_line",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Wedge:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Spread", instanceWithInferedType.Spread, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Niente", instanceWithInferedType.Niente, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wood:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Work:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Work_number", instanceWithInferedType.Work_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Work_title", instanceWithInferedType.Work_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Opus", instanceWithInferedType.Opus, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
