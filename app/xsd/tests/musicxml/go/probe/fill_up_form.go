// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
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
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.A_directive](
				"Attributes",
				"Directive",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.A_directive {
					return owner.Directive
				})
		}

	case *models.A_measure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Implicit", instanceWithInferedType.Implicit, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Non_controlling", instanceWithInferedType.Non_controlling, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_part, *models.A_measure](
				"A_part",
				"Measure",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part) []*models.A_measure {
					return owner.Measure
				})
		}

	case *models.A_measure_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Implicit", instanceWithInferedType.Implicit, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Non_controlling", instanceWithInferedType.Non_controlling, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Part", instanceWithInferedType, &instanceWithInferedType.Part, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_timewise, *models.A_measure_1](
				"Score_timewise",
				"Measure",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_timewise) []*models.A_measure_1 {
					return owner.Measure
				})
		}

	case *models.A_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Measure", instanceWithInferedType, &instanceWithInferedType.Measure, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_partwise, *models.A_part](
				"Score_partwise",
				"Part",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_partwise) []*models.A_part {
					return owner.Part
				})
		}

	case *models.A_part_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure_1, *models.A_part_1](
				"A_measure_1",
				"Part",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure_1) []*models.A_part_1 {
					return owner.Part
				})
		}

	case *models.Accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Cautionary", instanceWithInferedType.Cautionary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Editorial", instanceWithInferedType.Editorial, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Accidental_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Accidental_mark](
				"Notations",
				"Accidental_mark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Accidental_mark {
					return owner.Accidental_mark
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Accidental_mark](
				"Ornaments",
				"Accidental_mark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Accidental_mark {
					return owner.Accidental_mark
				})
		}

	case *models.Accidental_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Name_display, *models.Accidental_text](
				"Name_display",
				"Accidental_text",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Name_display) []*models.Accidental_text {
					return owner.Accidental_text
				})
		}
		{
			AssociationReverseSliceToForm[*models.Notehead_text, *models.Accidental_text](
				"Notehead_text",
				"Accidental_text",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notehead_text) []*models.Accidental_text {
					return owner.Accidental_text
				})
		}

	case *models.Accord:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("String", instanceWithInferedType.String, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Tuning_step", instanceWithInferedType.Tuning_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Tuning_alter", instanceWithInferedType.Tuning_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tuning_octave", instanceWithInferedType.Tuning_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Scordatura, *models.Accord](
				"Scordatura",
				"Accord",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scordatura) []*models.Accord {
					return owner.Accord
				})
		}

	case *models.Accordion_registration:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Accordion_high", instanceWithInferedType.Accordion_high, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Accordion_middle", instanceWithInferedType.Accordion_middle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Accordion_low", instanceWithInferedType.Accordion_low, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Line_width", instanceWithInferedType, &instanceWithInferedType.Line_width, formGroup, probe)
		AssociationSliceToForm("Note_size", instanceWithInferedType, &instanceWithInferedType.Note_size, formGroup, probe)
		AssociationSliceToForm("Distance", instanceWithInferedType, &instanceWithInferedType.Distance, formGroup, probe)
		AssociationSliceToForm("Glyph", instanceWithInferedType, &instanceWithInferedType.Glyph, formGroup, probe)
		AssociationSliceToForm("Other_appearance", instanceWithInferedType, &instanceWithInferedType.Other_appearance, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Unbroken", instanceWithInferedType.Unbroken, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Arpeggiate](
				"Notations",
				"Arpeggiate",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Arpeggiate {
					return owner.Arpeggiate
				})
		}

	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Arrow_direction", instanceWithInferedType.Arrow_direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Arrow_style", instanceWithInferedType.Arrow_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Arrowhead", instanceWithInferedType.Arrowhead, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Circular_arrow", instanceWithInferedType.Circular_arrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Arrow](
				"Technical",
				"Arrow",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Arrow {
					return owner.Arrow
				})
		}

	case *models.Articulations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Articulations](
				"Notations",
				"Articulations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Articulations {
					return owner.Articulations
				})
		}

	case *models.Assess:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Listen, *models.Assess](
				"Listen",
				"Assess",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Listen) []*models.Assess {
					return owner.Assess
				})
		}

	case *models.Attributes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Key", instanceWithInferedType, &instanceWithInferedType.Key, formGroup, probe)
		AssociationSliceToForm("Time", instanceWithInferedType, &instanceWithInferedType.Time, formGroup, probe)
		BasicFieldtoForm("Staves", instanceWithInferedType.Staves, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part_symbol", instanceWithInferedType.Part_symbol, formGroup, probe)
		BasicFieldtoForm("Instruments", instanceWithInferedType.Instruments, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Clef", instanceWithInferedType, &instanceWithInferedType.Clef, formGroup, probe)
		AssociationSliceToForm("Staff_details", instanceWithInferedType, &instanceWithInferedType.Staff_details, formGroup, probe)
		AssociationSliceToForm("Transpose", instanceWithInferedType, &instanceWithInferedType.Transpose, formGroup, probe)
		AssociationSliceToForm("For_part", instanceWithInferedType, &instanceWithInferedType.For_part, formGroup, probe)
		AssociationSliceToForm("Directive", instanceWithInferedType, &instanceWithInferedType.Directive, formGroup, probe)
		AssociationSliceToForm("Measure_style", instanceWithInferedType, &instanceWithInferedType.Measure_style, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Attributes](
				"A_measure",
				"Attributes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Attributes {
					return owner.Attributes
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Attributes](
				"A_part_1",
				"Attributes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Attributes {
					return owner.Attributes
				})
		}

	case *models.Backup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Backup](
				"A_measure",
				"Backup",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Backup {
					return owner.Backup
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Backup](
				"A_part_1",
				"Backup",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Backup {
					return owner.Backup
				})
		}

	case *models.Bar_style_color:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Bar_style", instanceWithInferedType.Bar_style, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		AssociationFieldToForm("Wavy_line", instanceWithInferedType.Wavy_line, formGroup, probe)
		AssociationFieldToForm("Segno_1", instanceWithInferedType.Segno_1, formGroup, probe)
		AssociationFieldToForm("Coda_1", instanceWithInferedType.Coda_1, formGroup, probe)
		AssociationFieldToForm("Fermata", instanceWithInferedType.Fermata, formGroup, probe)
		AssociationFieldToForm("Ending", instanceWithInferedType.Ending, formGroup, probe)
		AssociationFieldToForm("Repeat", instanceWithInferedType.Repeat, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Barline](
				"A_measure",
				"Barline",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Barline {
					return owner.Barline
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Barline](
				"A_part_1",
				"Barline",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Barline {
					return owner.Barline
				})
		}

	case *models.Barre:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Arrangement", instanceWithInferedType.Arrangement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Bass_separator", instanceWithInferedType.Bass_separator, formGroup, probe)
		AssociationFieldToForm("Bass_step", instanceWithInferedType.Bass_step, formGroup, probe)
		AssociationFieldToForm("Bass_alter", instanceWithInferedType.Bass_alter, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Bass_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Repeater", instanceWithInferedType.Repeater, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Fan", instanceWithInferedType.Fan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Beat_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slashes", instanceWithInferedType.Slashes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Use_dots", instanceWithInferedType.Use_dots, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Slash_type", instanceWithInferedType.Slash_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slash_dot", instanceWithInferedType.Slash_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Except_voice", instanceWithInferedType.Except_voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Beat_unit_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Beat_unit", instanceWithInferedType.Beat_unit, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beat_unit_dot", instanceWithInferedType.Beat_unit_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Metronome, *models.Beat_unit_tied](
				"Metronome",
				"Beat_unit_tied",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Metronome) []*models.Beat_unit_tied {
					return owner.Beat_unit_tied
				})
		}

	case *models.Beater:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tip", instanceWithInferedType.Tip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Bend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Shape", instanceWithInferedType.Shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("First_beat", instanceWithInferedType.First_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bend_alter", instanceWithInferedType.Bend_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pre_bend", instanceWithInferedType.Pre_bend, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Release", instanceWithInferedType.Release, formGroup, probe)
		AssociationFieldToForm("With_bar", instanceWithInferedType.With_bar, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Bend](
				"Technical",
				"Bend",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Bend {
					return owner.Bend
				})
		}

	case *models.Bookmark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Position", instanceWithInferedType.Position, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Bookmark](
				"A_measure",
				"Bookmark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Bookmark {
					return owner.Bookmark
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Bookmark](
				"A_part_1",
				"Bookmark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Bookmark {
					return owner.Bookmark
				})
		}
		{
			AssociationReverseSliceToForm[*models.Credit, *models.Bookmark](
				"Credit",
				"Bookmark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Credit) []*models.Bookmark {
					return owner.Bookmark
				})
		}

	case *models.Bracket:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_end", instanceWithInferedType.Line_end, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("End_length", instanceWithInferedType.End_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Breath_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Breath_mark](
				"Articulations",
				"Breath_mark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Breath_mark {
					return owner.Breath_mark
				})
		}

	case *models.Caesura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Caesura](
				"Articulations",
				"Caesura",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Caesura {
					return owner.Caesura
				})
		}

	case *models.Cancel:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Additional", instanceWithInferedType.Additional, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("After_barline", instanceWithInferedType.After_barline, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Clef_octave_change", instanceWithInferedType.Clef_octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Clef](
				"Attributes",
				"Clef",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Clef {
					return owner.Clef
				})
		}

	case *models.Coda:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Coda](
				"Direction_type",
				"Coda",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Coda {
					return owner.Coda
				})
		}

	case *models.Credit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Page", instanceWithInferedType.Page, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Credit_type", instanceWithInferedType.Credit_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Credit_image", instanceWithInferedType.Credit_image, formGroup, probe)
		AssociationSliceToForm("Link", instanceWithInferedType, &instanceWithInferedType.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", instanceWithInferedType, &instanceWithInferedType.Bookmark, formGroup, probe)
		AssociationSliceToForm("Credit_words", instanceWithInferedType, &instanceWithInferedType.Credit_words, formGroup, probe)
		AssociationSliceToForm("Credit_symbol", instanceWithInferedType, &instanceWithInferedType.Credit_symbol, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_partwise, *models.Credit](
				"Score_partwise",
				"Credit",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_partwise) []*models.Credit {
					return owner.Credit
				})
		}
		{
			AssociationReverseSliceToForm[*models.Score_timewise, *models.Credit](
				"Score_timewise",
				"Credit",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_timewise) []*models.Credit {
					return owner.Credit
				})
		}

	case *models.Dashes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Defaults:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Scaling", instanceWithInferedType.Scaling, formGroup, probe)
		BasicFieldtoForm("Concert_score", instanceWithInferedType.Concert_score, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Page_layout", instanceWithInferedType.Page_layout, formGroup, probe)
		AssociationFieldToForm("System_layout", instanceWithInferedType.System_layout, formGroup, probe)
		AssociationSliceToForm("Staff_layout", instanceWithInferedType, &instanceWithInferedType.Staff_layout, formGroup, probe)
		AssociationFieldToForm("Appearance", instanceWithInferedType.Appearance, formGroup, probe)
		AssociationFieldToForm("Music_font", instanceWithInferedType.Music_font, formGroup, probe)
		AssociationFieldToForm("Word_font", instanceWithInferedType.Word_font, formGroup, probe)
		AssociationSliceToForm("Lyric_font", instanceWithInferedType, &instanceWithInferedType.Lyric_font, formGroup, probe)
		AssociationSliceToForm("Lyric_language", instanceWithInferedType, &instanceWithInferedType.Lyric_language, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Degree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("Degree_value", instanceWithInferedType.Degree_value, formGroup, probe)
		AssociationFieldToForm("Degree_alter", instanceWithInferedType.Degree_alter, formGroup, probe)
		AssociationFieldToForm("Degree_type", instanceWithInferedType.Degree_type, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Harmony, *models.Degree](
				"Harmony",
				"Degree",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Harmony) []*models.Degree {
					return owner.Degree
				})
		}

	case *models.Degree_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Plus_minus", instanceWithInferedType.Plus_minus, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Degree_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Degree_value:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Direction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Directive", instanceWithInferedType.Directive, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Direction_type", instanceWithInferedType, &instanceWithInferedType.Direction_type, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Sound", instanceWithInferedType.Sound, formGroup, probe)
		AssociationFieldToForm("Listening", instanceWithInferedType.Listening, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Direction](
				"A_measure",
				"Direction",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Direction {
					return owner.Direction
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Direction](
				"A_part_1",
				"Direction",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Direction {
					return owner.Direction
				})
		}

	case *models.Direction_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Direction, *models.Direction_type](
				"Direction",
				"Direction_type",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction) []*models.Direction_type {
					return owner.Direction_type
				})
		}

	case *models.Distance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Appearance, *models.Distance](
				"Appearance",
				"Distance",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Appearance) []*models.Distance {
					return owner.Distance
				})
		}

	case *models.Double:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Above", instanceWithInferedType.Above, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Dynamics:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("P", instanceWithInferedType.P, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pp", instanceWithInferedType.Pp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ppp", instanceWithInferedType.Ppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pppp", instanceWithInferedType.Pppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ppppp", instanceWithInferedType.Ppppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pppppp", instanceWithInferedType.Pppppp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("F", instanceWithInferedType.F, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ff", instanceWithInferedType.Ff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fff", instanceWithInferedType.Fff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ffff", instanceWithInferedType.Ffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fffff", instanceWithInferedType.Fffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ffffff", instanceWithInferedType.Ffffff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Mp", instanceWithInferedType.Mp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Mf", instanceWithInferedType.Mf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sf", instanceWithInferedType.Sf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sfp", instanceWithInferedType.Sfp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sfpp", instanceWithInferedType.Sfpp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fp", instanceWithInferedType.Fp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rf", instanceWithInferedType.Rf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rfz", instanceWithInferedType.Rfz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sfz", instanceWithInferedType.Sfz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sffz", instanceWithInferedType.Sffz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fz", instanceWithInferedType.Fz, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("N", instanceWithInferedType.N, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pf", instanceWithInferedType.Pf, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sfzp", instanceWithInferedType.Sfzp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Other_dynamics", instanceWithInferedType, &instanceWithInferedType.Other_dynamics, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Dynamics](
				"Direction_type",
				"Dynamics",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Dynamics {
					return owner.Dynamics
				})
		}
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Dynamics](
				"Notations",
				"Dynamics",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Dynamics {
					return owner.Dynamics
				})
		}

	case *models.Effect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Elision:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Lyric, *models.Elision](
				"Lyric",
				"Elision",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Lyric) []*models.Elision {
					return owner.Elision
				})
		}

	case *models.Empty:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_shape", instanceWithInferedType.Line_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_length", instanceWithInferedType.Line_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_line](
				"Articulations",
				"Scoop",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_line {
					return owner.Scoop
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_line](
				"Articulations",
				"Plop",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_line {
					return owner.Plop
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_line](
				"Articulations",
				"Doit",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_line {
					return owner.Doit
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_line](
				"Articulations",
				"Falloff",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_line {
					return owner.Falloff
				})
		}

	case *models.Empty_placement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Accent",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Accent
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Staccato",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Staccato
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Tenuto",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Tenuto
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Detached_legato",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Detached_legato
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Staccatissimo",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Staccatissimo
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Spiccato",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Spiccato
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Stress",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Stress
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Unstress",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Unstress
				})
		}
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Empty_placement](
				"Articulations",
				"Soft_accent",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Empty_placement {
					return owner.Soft_accent
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Empty_placement](
				"Note",
				"Dot",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Empty_placement {
					return owner.Dot
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_placement](
				"Ornaments",
				"Schleifer",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_placement {
					return owner.Schleifer
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Up_bow",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Up_bow
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Down_bow",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Down_bow
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Open_string",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Open_string
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Thumb_position",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Thumb_position
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Double_tongue",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Double_tongue
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Triple_tongue",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Triple_tongue
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Snap_pizzicato",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Snap_pizzicato
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Fingernails",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Fingernails
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Brass_bend",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Brass_bend
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Flip",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Flip
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Smear",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Smear
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement](
				"Technical",
				"Golpe",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement {
					return owner.Golpe
				})
		}

	case *models.Empty_placement_smufl:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement_smufl](
				"Technical",
				"Stopped",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement_smufl {
					return owner.Stopped
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement_smufl](
				"Technical",
				"Open",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement_smufl {
					return owner.Open
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Empty_placement_smufl](
				"Technical",
				"Half_muted",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Empty_placement_smufl {
					return owner.Half_muted
				})
		}

	case *models.Empty_print_object_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_print_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_print_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_print_style_align_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Empty_trill_sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_trill_sound](
				"Ornaments",
				"Trill_mark",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_trill_sound {
					return owner.Trill_mark
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_trill_sound](
				"Ornaments",
				"Vertical_turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_trill_sound {
					return owner.Vertical_turn
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_trill_sound](
				"Ornaments",
				"Inverted_vertical_turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_trill_sound {
					return owner.Inverted_vertical_turn
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_trill_sound](
				"Ornaments",
				"Shake",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_trill_sound {
					return owner.Shake
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Empty_trill_sound](
				"Ornaments",
				"Haydn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Empty_trill_sound {
					return owner.Haydn
				})
		}

	case *models.Encoding:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Encoder", instanceWithInferedType, &instanceWithInferedType.Encoder, formGroup, probe)
		BasicFieldtoForm("Software", instanceWithInferedType.Software, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Encoding_description", instanceWithInferedType.Encoding_description, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Supports", instanceWithInferedType, &instanceWithInferedType.Supports, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Ending:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("End_length", instanceWithInferedType.End_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text_x", instanceWithInferedType.Text_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text_y", instanceWithInferedType.Text_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Extend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Feature:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Grouping, *models.Feature](
				"Grouping",
				"Feature",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Grouping) []*models.Feature {
					return owner.Feature
				})
		}

	case *models.Fermata:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Fermata](
				"Notations",
				"Fermata",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Fermata {
					return owner.Fermata
				})
		}

	case *models.Figure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Prefix", instanceWithInferedType.Prefix, formGroup, probe)
		AssociationFieldToForm("Figure_number", instanceWithInferedType.Figure_number, formGroup, probe)
		AssociationFieldToForm("Suffix", instanceWithInferedType.Suffix, formGroup, probe)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Figured_bass, *models.Figure](
				"Figured_bass",
				"Figure",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Figured_bass) []*models.Figure {
					return owner.Figure
				})
		}

	case *models.Figured_bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_dot", instanceWithInferedType.Print_dot, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_lyric", instanceWithInferedType.Print_lyric, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Figure", instanceWithInferedType, &instanceWithInferedType.Figure, formGroup, probe)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Figured_bass](
				"A_measure",
				"Figured_bass",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Figured_bass {
					return owner.Figured_bass
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Figured_bass](
				"A_part_1",
				"Figured_bass",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Figured_bass {
					return owner.Figured_bass
				})
		}

	case *models.Fingering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Substitution", instanceWithInferedType.Substitution, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Alternate", instanceWithInferedType.Alternate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Fingering](
				"Technical",
				"Fingering",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Fingering {
					return owner.Fingering
				})
		}

	case *models.First_fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.For_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part_clef", instanceWithInferedType.Part_clef, formGroup, probe)
		AssociationFieldToForm("Part_transpose", instanceWithInferedType.Part_transpose, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.For_part](
				"Attributes",
				"For_part",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.For_part {
					return owner.For_part
				})
		}

	case *models.Formatted_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Formatted_symbol_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Credit, *models.Formatted_symbol_id](
				"Credit",
				"Credit_symbol",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Credit) []*models.Formatted_symbol_id {
					return owner.Credit_symbol
				})
		}
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Formatted_symbol_id](
				"Direction_type",
				"Symbol",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Formatted_symbol_id {
					return owner.Symbol
				})
		}

	case *models.Formatted_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Name_display, *models.Formatted_text](
				"Name_display",
				"Display_text",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Name_display) []*models.Formatted_text {
					return owner.Display_text
				})
		}
		{
			AssociationReverseSliceToForm[*models.Notehead_text, *models.Formatted_text](
				"Notehead_text",
				"Display_text",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notehead_text) []*models.Formatted_text {
					return owner.Display_text
				})
		}

	case *models.Formatted_text_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space", instanceWithInferedType.Space, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_height", instanceWithInferedType.Line_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Credit, *models.Formatted_text_id](
				"Credit",
				"Credit_words",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Credit) []*models.Formatted_text_id {
					return owner.Credit_words
				})
		}
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Formatted_text_id](
				"Direction_type",
				"Rehearsal",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Formatted_text_id {
					return owner.Rehearsal
				})
		}
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Formatted_text_id](
				"Direction_type",
				"Words",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Formatted_text_id {
					return owner.Words
				})
		}

	case *models.Forward:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Forward](
				"A_measure",
				"Forward",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Forward {
					return owner.Forward
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Forward](
				"A_part_1",
				"Forward",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Forward {
					return owner.Forward
				})
		}

	case *models.Frame:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Unplayed", instanceWithInferedType.Unplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Frame_strings", instanceWithInferedType.Frame_strings, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Frame_frets", instanceWithInferedType.Frame_frets, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("First_fret", instanceWithInferedType.First_fret, formGroup, probe)
		AssociationSliceToForm("Frame_note", instanceWithInferedType, &instanceWithInferedType.Frame_note, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Frame_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("String", instanceWithInferedType.String, formGroup, probe)
		AssociationFieldToForm("Fret", instanceWithInferedType.Fret, formGroup, probe)
		AssociationFieldToForm("Fingering", instanceWithInferedType.Fingering, formGroup, probe)
		AssociationFieldToForm("Barre", instanceWithInferedType.Barre, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Frame, *models.Frame_note](
				"Frame",
				"Frame_note",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Frame) []*models.Frame_note {
					return owner.Frame_note
				})
		}

	case *models.Fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Fret](
				"Technical",
				"Fret",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Fret {
					return owner.Fret
				})
		}

	case *models.Glass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Glissando:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Glissando](
				"Notations",
				"Glissando",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Glissando {
					return owner.Glissando
				})
		}

	case *models.Glyph:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Appearance, *models.Glyph](
				"Appearance",
				"Glyph",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Appearance) []*models.Glyph {
					return owner.Glyph
				})
		}

	case *models.Grace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Steal_time_previous", instanceWithInferedType.Steal_time_previous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Steal_time_following", instanceWithInferedType.Steal_time_following, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Make_time", instanceWithInferedType.Make_time, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Slash", instanceWithInferedType.Slash, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Group_barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Group_name:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Group_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Grouping:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Member_of", instanceWithInferedType.Member_of, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Feature", instanceWithInferedType, &instanceWithInferedType.Feature, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Grouping](
				"A_measure",
				"Grouping",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Grouping {
					return owner.Grouping
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Grouping](
				"A_part_1",
				"Grouping",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Grouping {
					return owner.Grouping
				})
		}

	case *models.Hammer_on_pull_off:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Hammer_on_pull_off](
				"Technical",
				"Hammer_on",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Hammer_on_pull_off {
					return owner.Hammer_on
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Hammer_on_pull_off](
				"Technical",
				"Pull_off",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Hammer_on_pull_off {
					return owner.Pull_off
				})
		}

	case *models.Handbell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Handbell](
				"Technical",
				"Handbell",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Handbell {
					return owner.Handbell
				})
		}

	case *models.Harmon_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Harmon_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Harmon_closed", instanceWithInferedType.Harmon_closed, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Harmon_mute](
				"Technical",
				"Harmon_mute",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Harmon_mute {
					return owner.Harmon_mute
				})
		}

	case *models.Harmonic:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Natural", instanceWithInferedType.Natural, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Artificial", instanceWithInferedType.Artificial, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Base_pitch", instanceWithInferedType.Base_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Touching_pitch", instanceWithInferedType.Touching_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sounding_pitch", instanceWithInferedType.Sounding_pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Harmonic](
				"Technical",
				"Harmonic",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Harmonic {
					return owner.Harmonic
				})
		}

	case *models.Harmony:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_frame", instanceWithInferedType.Print_frame, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Arrangement", instanceWithInferedType.Arrangement, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Harmony](
				"A_measure",
				"Harmony",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Harmony {
					return owner.Harmony
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Harmony](
				"A_part_1",
				"Harmony",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Harmony {
					return owner.Harmony
				})
		}

	case *models.Harmony_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Harp_pedals:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Pedal_tuning", instanceWithInferedType, &instanceWithInferedType.Pedal_tuning, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Heel_toe:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Heel_toe](
				"Technical",
				"Heel",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Heel_toe {
					return owner.Heel
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Heel_toe](
				"Technical",
				"Toe",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Heel_toe {
					return owner.Toe
				})
		}

	case *models.Hole:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Hole_type", instanceWithInferedType.Hole_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Hole_closed", instanceWithInferedType.Hole_closed, formGroup, probe)
		BasicFieldtoForm("Hole_shape", instanceWithInferedType.Hole_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Hole](
				"Technical",
				"Hole",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Hole {
					return owner.Hole
				})
		}

	case *models.Hole_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Location", instanceWithInferedType.Location, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Horizontal_turn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Slash", instanceWithInferedType.Slash, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Horizontal_turn](
				"Ornaments",
				"Turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Horizontal_turn {
					return owner.Turn
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Horizontal_turn](
				"Ornaments",
				"Delayed_turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Horizontal_turn {
					return owner.Delayed_turn
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Horizontal_turn](
				"Ornaments",
				"Inverted_turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Horizontal_turn {
					return owner.Inverted_turn
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Horizontal_turn](
				"Ornaments",
				"Delayed_inverted_turn",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Horizontal_turn {
					return owner.Delayed_inverted_turn
				})
		}

	case *models.Identification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Creator", instanceWithInferedType, &instanceWithInferedType.Creator, formGroup, probe)
		AssociationSliceToForm("Rights", instanceWithInferedType, &instanceWithInferedType.Rights, formGroup, probe)
		AssociationFieldToForm("Encoding", instanceWithInferedType.Encoding, formGroup, probe)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Relation", instanceWithInferedType, &instanceWithInferedType.Relation, formGroup, probe)
		AssociationFieldToForm("Miscellaneous", instanceWithInferedType.Miscellaneous, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Image:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Note, *models.Instrument](
				"Note",
				"Instrument",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Instrument {
					return owner.Instrument
				})
		}

	case *models.Instrument_change:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Instrument_sound", instanceWithInferedType.Instrument_sound, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Solo", instanceWithInferedType.Solo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ensemble", instanceWithInferedType.Ensemble, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Virtual_instrument", instanceWithInferedType.Virtual_instrument, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Sound, *models.Instrument_change](
				"Sound",
				"Instrument_change",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sound) []*models.Instrument_change {
					return owner.Instrument_change
				})
		}

	case *models.Instrument_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Part_link, *models.Instrument_link](
				"Part_link",
				"Instrument_link",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Part_link) []*models.Instrument_link {
					return owner.Instrument_link
				})
		}

	case *models.Interchangeable:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Separator", instanceWithInferedType.Separator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Time_relation", instanceWithInferedType.Time_relation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Beat_type", instanceWithInferedType.Beat_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Inversion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Cancel", instanceWithInferedType.Cancel, formGroup, probe)
		BasicFieldtoForm("Fifths", instanceWithInferedType.Fifths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Mode", instanceWithInferedType.Mode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Key_step", instanceWithInferedType.Key_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Key_alter", instanceWithInferedType.Key_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Key_accidental", instanceWithInferedType.Key_accidental, formGroup, probe)
		AssociationSliceToForm("Key_octave", instanceWithInferedType, &instanceWithInferedType.Key_octave, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Key](
				"Attributes",
				"Key",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Key {
					return owner.Key
				})
		}

	case *models.Key_accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Key_octave:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Cancel", instanceWithInferedType.Cancel, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Key, *models.Key_octave](
				"Key",
				"Key_octave",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Key) []*models.Key_octave {
					return owner.Key_octave
				})
		}

	case *models.Kind:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Use_symbols", instanceWithInferedType.Use_symbols, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Stack_degrees", instanceWithInferedType.Stack_degrees, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses_degrees", instanceWithInferedType.Parentheses_degrees, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket_degrees", instanceWithInferedType.Bracket_degrees, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Level:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Reference", instanceWithInferedType.Reference, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Line_detail:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Staff_details, *models.Line_detail](
				"Staff_details",
				"Line_detail",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Staff_details) []*models.Line_detail {
					return owner.Line_detail
				})
		}

	case *models.Line_width:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Appearance, *models.Line_width](
				"Appearance",
				"Line_width",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Appearance) []*models.Line_width {
					return owner.Line_width
				})
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Position", instanceWithInferedType.Position, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Link](
				"A_measure",
				"Link",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Link {
					return owner.Link
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Link](
				"A_part_1",
				"Link",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Link {
					return owner.Link
				})
		}
		{
			AssociationReverseSliceToForm[*models.Credit, *models.Link](
				"Credit",
				"Link",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Credit) []*models.Link {
					return owner.Link
				})
		}

	case *models.Listen:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Assess", instanceWithInferedType, &instanceWithInferedType.Assess, formGroup, probe)
		AssociationSliceToForm("Wait", instanceWithInferedType, &instanceWithInferedType.Wait, formGroup, probe)
		AssociationSliceToForm("Other_listen", instanceWithInferedType, &instanceWithInferedType.Other_listen, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sync", instanceWithInferedType, &instanceWithInferedType.Sync, formGroup, probe)
		AssociationSliceToForm("Other_listening", instanceWithInferedType, &instanceWithInferedType.Other_listening, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Listening](
				"A_measure",
				"Listening",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Listening {
					return owner.Listening
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Listening](
				"A_part_1",
				"Listening",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Listening {
					return owner.Listening
				})
		}

	case *models.Lyric:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Elision", instanceWithInferedType, &instanceWithInferedType.Elision, formGroup, probe)
		BasicFieldtoForm("Syllabic", instanceWithInferedType.Syllabic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Text", instanceWithInferedType, &instanceWithInferedType.Text, formGroup, probe)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		BasicFieldtoForm("Laughing", instanceWithInferedType.Laughing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Humming", instanceWithInferedType.Humming, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("End_line", instanceWithInferedType.End_line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("End_paragraph", instanceWithInferedType.End_paragraph, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Note, *models.Lyric](
				"Note",
				"Lyric",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Lyric {
					return owner.Lyric
				})
		}

	case *models.Lyric_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Defaults, *models.Lyric_font](
				"Defaults",
				"Lyric_font",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Defaults) []*models.Lyric_font {
					return owner.Lyric_font
				})
		}

	case *models.Lyric_language:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Defaults, *models.Lyric_language](
				"Defaults",
				"Lyric_language",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Defaults) []*models.Lyric_language {
					return owner.Lyric_language
				})
		}

	case *models.Measure_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Measure_distance", instanceWithInferedType.Measure_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Measure_numbering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("System", instanceWithInferedType.System, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Multiple_rest_always", instanceWithInferedType.Multiple_rest_always, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Multiple_rest_range", instanceWithInferedType.Multiple_rest_range, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Measure_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slashes", instanceWithInferedType.Slashes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Measure_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Multiple_rest", instanceWithInferedType.Multiple_rest, formGroup, probe)
		AssociationFieldToForm("Measure_repeat", instanceWithInferedType.Measure_repeat, formGroup, probe)
		AssociationFieldToForm("Beat_repeat", instanceWithInferedType.Beat_repeat, formGroup, probe)
		AssociationFieldToForm("Slash", instanceWithInferedType.Slash, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Measure_style](
				"Attributes",
				"Measure_style",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Measure_style {
					return owner.Measure_style
				})
		}

	case *models.Membrane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Metal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Metronome:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Beat_unit", instanceWithInferedType.Beat_unit, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beat_unit_dot", instanceWithInferedType.Beat_unit_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Per_minute", instanceWithInferedType.Per_minute, formGroup, probe)
		AssociationSliceToForm("Beat_unit_tied", instanceWithInferedType, &instanceWithInferedType.Beat_unit_tied, formGroup, probe)
		BasicFieldtoForm("Metronome_arrows", instanceWithInferedType.Metronome_arrows, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Metronome_relation", instanceWithInferedType.Metronome_relation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Metronome_note", instanceWithInferedType, &instanceWithInferedType.Metronome_note, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Metronome_beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Metronome_note, *models.Metronome_beam](
				"Metronome_note",
				"Metronome_beam",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Metronome_note) []*models.Metronome_beam {
					return owner.Metronome_beam
				})
		}

	case *models.Metronome_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Metronome_type", instanceWithInferedType.Metronome_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Metronome_dot", instanceWithInferedType.Metronome_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Metronome_beam", instanceWithInferedType, &instanceWithInferedType.Metronome_beam, formGroup, probe)
		AssociationFieldToForm("Metronome_tied", instanceWithInferedType.Metronome_tied, formGroup, probe)
		AssociationFieldToForm("Metronome_tuplet", instanceWithInferedType.Metronome_tuplet, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Metronome, *models.Metronome_note](
				"Metronome",
				"Metronome_note",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Metronome) []*models.Metronome_note {
					return owner.Metronome_note
				})
		}

	case *models.Metronome_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Metronome_tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Midi_device:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Port", instanceWithInferedType.Port, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_part, *models.Midi_device](
				"Score_part",
				"Midi_device",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_part) []*models.Midi_device {
					return owner.Midi_device
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sound, *models.Midi_device](
				"Sound",
				"Midi_device",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sound) []*models.Midi_device {
					return owner.Midi_device
				})
		}

	case *models.Midi_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Midi_channel", instanceWithInferedType.Midi_channel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Midi_name", instanceWithInferedType.Midi_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Midi_bank", instanceWithInferedType.Midi_bank, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Midi_program", instanceWithInferedType.Midi_program, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Midi_unpitched", instanceWithInferedType.Midi_unpitched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Volume", instanceWithInferedType.Volume, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pan", instanceWithInferedType.Pan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Elevation", instanceWithInferedType.Elevation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_part, *models.Midi_instrument](
				"Score_part",
				"Midi_instrument",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_part) []*models.Midi_instrument {
					return owner.Midi_instrument
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sound, *models.Midi_instrument](
				"Sound",
				"Midi_instrument",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sound) []*models.Midi_instrument {
					return owner.Midi_instrument
				})
		}

	case *models.Miscellaneous:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Miscellaneous_field", instanceWithInferedType, &instanceWithInferedType.Miscellaneous_field, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Miscellaneous_field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Miscellaneous, *models.Miscellaneous_field](
				"Miscellaneous",
				"Miscellaneous_field",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Miscellaneous) []*models.Miscellaneous_field {
					return owner.Miscellaneous_field
				})
		}

	case *models.Mordent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Mordent](
				"Ornaments",
				"Mordent",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Mordent {
					return owner.Mordent
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Mordent](
				"Ornaments",
				"Inverted_mordent",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Mordent {
					return owner.Inverted_mordent
				})
		}

	case *models.Multiple_rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Use_symbols", instanceWithInferedType.Use_symbols, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Name_display:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Display_text", instanceWithInferedType, &instanceWithInferedType.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", instanceWithInferedType, &instanceWithInferedType.Accidental_text, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Non_arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Non_arpeggiate](
				"Notations",
				"Non_arpeggiate",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Non_arpeggiate {
					return owner.Non_arpeggiate
				})
		}

	case *models.Notations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Note, *models.Notations](
				"Note",
				"Notations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Notations {
					return owner.Notations
				})
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_leger", instanceWithInferedType.Print_leger, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Dynamics", instanceWithInferedType.Dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("End_dynamics", instanceWithInferedType.End_dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Attack", instanceWithInferedType.Attack, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Release", instanceWithInferedType.Release, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Pizzicato", instanceWithInferedType.Pizzicato, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_dot", instanceWithInferedType.Print_dot, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_lyric", instanceWithInferedType.Print_lyric, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Grace", instanceWithInferedType.Grace, formGroup, probe)
		BasicFieldtoForm("Chord", instanceWithInferedType.Chord, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Pitch", instanceWithInferedType.Pitch, formGroup, probe)
		AssociationFieldToForm("Unpitched", instanceWithInferedType.Unpitched, formGroup, probe)
		AssociationFieldToForm("Rest", instanceWithInferedType.Rest, formGroup, probe)
		AssociationFieldToForm("Tie", instanceWithInferedType.Tie, formGroup, probe)
		BasicFieldtoForm("Cue", instanceWithInferedType.Cue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Instrument", instanceWithInferedType, &instanceWithInferedType.Instrument, formGroup, probe)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		BasicFieldtoForm("Voice", instanceWithInferedType.Voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Type", instanceWithInferedType.Type, formGroup, probe)
		AssociationSliceToForm("Dot", instanceWithInferedType, &instanceWithInferedType.Dot, formGroup, probe)
		AssociationFieldToForm("Accidental", instanceWithInferedType.Accidental, formGroup, probe)
		AssociationFieldToForm("Time_modification", instanceWithInferedType.Time_modification, formGroup, probe)
		AssociationFieldToForm("Stem", instanceWithInferedType.Stem, formGroup, probe)
		AssociationFieldToForm("Notehead", instanceWithInferedType.Notehead, formGroup, probe)
		AssociationFieldToForm("Notehead_text", instanceWithInferedType.Notehead_text, formGroup, probe)
		BasicFieldtoForm("Staff", instanceWithInferedType.Staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Beam", instanceWithInferedType.Beam, formGroup, probe)
		AssociationSliceToForm("Notations", instanceWithInferedType, &instanceWithInferedType.Notations, formGroup, probe)
		AssociationSliceToForm("Lyric", instanceWithInferedType, &instanceWithInferedType.Lyric, formGroup, probe)
		AssociationFieldToForm("Play", instanceWithInferedType.Play, formGroup, probe)
		AssociationFieldToForm("Listen", instanceWithInferedType.Listen, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Note](
				"A_measure",
				"Note",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Note {
					return owner.Note
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Note](
				"A_part_1",
				"Note",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Note {
					return owner.Note
				})
		}

	case *models.Note_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Appearance, *models.Note_size](
				"Appearance",
				"Note_size",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Appearance) []*models.Note_size {
					return owner.Note_size
				})
		}

	case *models.Note_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Notehead:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Filled", instanceWithInferedType.Filled, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Notehead_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Display_text", instanceWithInferedType, &instanceWithInferedType.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", instanceWithInferedType, &instanceWithInferedType.Accidental_text, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Numeral:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Numeral_root", instanceWithInferedType.Numeral_root, formGroup, probe)
		AssociationFieldToForm("Numeral_alter", instanceWithInferedType.Numeral_alter, formGroup, probe)
		AssociationFieldToForm("Numeral_key", instanceWithInferedType.Numeral_key, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Numeral_key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Numeral_fifths", instanceWithInferedType.Numeral_fifths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Numeral_mode", instanceWithInferedType.Numeral_mode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Numeral_root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Octave_shift:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Offset:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Sound", instanceWithInferedType.Sound, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Opus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Ornaments:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Ornaments](
				"Notations",
				"Ornaments",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Ornaments {
					return owner.Ornaments
				})
		}

	case *models.Other_appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Appearance, *models.Other_appearance](
				"Appearance",
				"Other_appearance",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Appearance) []*models.Other_appearance {
					return owner.Other_appearance
				})
		}

	case *models.Other_direction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Other_listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Listen, *models.Other_listening](
				"Listen",
				"Other_listen",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Listen) []*models.Other_listening {
					return owner.Other_listen
				})
		}
		{
			AssociationReverseSliceToForm[*models.Listening, *models.Other_listening](
				"Listening",
				"Other_listening",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Listening) []*models.Other_listening {
					return owner.Other_listening
				})
		}

	case *models.Other_notation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Other_notation](
				"Notations",
				"Other_notation",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Other_notation {
					return owner.Other_notation
				})
		}

	case *models.Other_placement_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Other_placement_text](
				"Articulations",
				"Other_articulation",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Other_placement_text {
					return owner.Other_articulation
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Other_placement_text](
				"Ornaments",
				"Other_ornament",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Other_placement_text {
					return owner.Other_ornament
				})
		}
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Other_placement_text](
				"Technical",
				"Other_technical",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Other_placement_text {
					return owner.Other_technical
				})
		}

	case *models.Other_play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Play, *models.Other_play](
				"Play",
				"Other_play",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Play) []*models.Other_play {
					return owner.Other_play
				})
		}

	case *models.Other_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Dynamics, *models.Other_text](
				"Dynamics",
				"Other_dynamics",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Dynamics) []*models.Other_text {
					return owner.Other_dynamics
				})
		}

	case *models.Page_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Page_height", instanceWithInferedType.Page_height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Page_width", instanceWithInferedType.Page_width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Page_margins", instanceWithInferedType.Page_margins, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Page_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Left_margin", instanceWithInferedType.Left_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Right_margin", instanceWithInferedType.Right_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Top_margin", instanceWithInferedType.Top_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bottom_margin", instanceWithInferedType.Bottom_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Clef_octave_change", instanceWithInferedType.Clef_octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Group_name", instanceWithInferedType.Group_name, formGroup, probe)
		AssociationFieldToForm("Group_name_display", instanceWithInferedType.Group_name_display, formGroup, probe)
		AssociationFieldToForm("Group_abbreviation", instanceWithInferedType.Group_abbreviation, formGroup, probe)
		AssociationFieldToForm("Group_abbreviation_display", instanceWithInferedType.Group_abbreviation_display, formGroup, probe)
		AssociationFieldToForm("Group_symbol", instanceWithInferedType.Group_symbol, formGroup, probe)
		AssociationFieldToForm("Group_barline", instanceWithInferedType.Group_barline, formGroup, probe)
		BasicFieldtoForm("Group_time", instanceWithInferedType.Group_time, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Footnote", instanceWithInferedType.Footnote, formGroup, probe)
		AssociationFieldToForm("Level", instanceWithInferedType.Level, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Href", instanceWithInferedType.Href, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Role", instanceWithInferedType.Role, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Title", instanceWithInferedType.Title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Show", instanceWithInferedType.Show, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Actuate", instanceWithInferedType.Actuate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Instrument_link", instanceWithInferedType, &instanceWithInferedType.Instrument_link, formGroup, probe)
		BasicFieldtoForm("Group_link", instanceWithInferedType.Group_link, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_part, *models.Part_link](
				"Score_part",
				"Part_link",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_part) []*models.Part_link {
					return owner.Part_link
				})
		}

	case *models.Part_list:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part_group", instanceWithInferedType.Part_group, formGroup, probe)
		AssociationFieldToForm("Score_part", instanceWithInferedType.Score_part, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_name:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Justify", instanceWithInferedType.Justify, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Top_staff", instanceWithInferedType.Top_staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bottom_staff", instanceWithInferedType.Bottom_staff, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Part_transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Diatonic", instanceWithInferedType.Diatonic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Chromatic", instanceWithInferedType.Chromatic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Octave_change", instanceWithInferedType.Octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Double", instanceWithInferedType.Double, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Pedal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Sign", instanceWithInferedType.Sign, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Abbreviated", instanceWithInferedType.Abbreviated, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Pedal_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Pedal_step", instanceWithInferedType.Pedal_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Pedal_alter", instanceWithInferedType.Pedal_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Harp_pedals, *models.Pedal_tuning](
				"Harp_pedals",
				"Pedal_tuning",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Harp_pedals) []*models.Pedal_tuning {
					return owner.Pedal_tuning
				})
		}

	case *models.Per_minute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Percussion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Enclosure", instanceWithInferedType.Enclosure, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Other_percussion", instanceWithInferedType.Other_percussion, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Percussion](
				"Direction_type",
				"Percussion",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Percussion {
					return owner.Percussion
				})
		}

	case *models.Pitch:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Step", instanceWithInferedType.Step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Alter", instanceWithInferedType.Alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Octave", instanceWithInferedType.Octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Pitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Placement_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Placement_text](
				"Technical",
				"Pluck",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Placement_text {
					return owner.Pluck
				})
		}

	case *models.Play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ipa", instanceWithInferedType.Ipa, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Mute", instanceWithInferedType.Mute, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Semi_pitched", instanceWithInferedType.Semi_pitched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Other_play", instanceWithInferedType, &instanceWithInferedType.Other_play, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Sound, *models.Play](
				"Sound",
				"Play",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sound) []*models.Play {
					return owner.Play
				})
		}

	case *models.Player:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Player_name", instanceWithInferedType.Player_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_part, *models.Player](
				"Score_part",
				"Player",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_part) []*models.Player {
					return owner.Player
				})
		}

	case *models.Principal_voice:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Print:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff_spacing", instanceWithInferedType.Staff_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("New_system", instanceWithInferedType.New_system, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("New_page", instanceWithInferedType.New_page, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Blank_page", instanceWithInferedType.Blank_page, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Page_number", instanceWithInferedType.Page_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Page_layout", instanceWithInferedType.Page_layout, formGroup, probe)
		AssociationFieldToForm("System_layout", instanceWithInferedType.System_layout, formGroup, probe)
		AssociationSliceToForm("Staff_layout", instanceWithInferedType, &instanceWithInferedType.Staff_layout, formGroup, probe)
		AssociationFieldToForm("Measure_layout", instanceWithInferedType.Measure_layout, formGroup, probe)
		AssociationFieldToForm("Measure_numbering", instanceWithInferedType.Measure_numbering, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Print](
				"A_measure",
				"Print",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Print {
					return owner.Print
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Print](
				"A_part_1",
				"Print",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Print {
					return owner.Print
				})
		}

	case *models.Release:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Times", instanceWithInferedType.Times, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("After_jump", instanceWithInferedType.After_jump, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Winged", instanceWithInferedType.Winged, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Measure", instanceWithInferedType.Measure, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Display_step", instanceWithInferedType.Display_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Display_octave", instanceWithInferedType.Display_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Root_step", instanceWithInferedType.Root_step, formGroup, probe)
		AssociationFieldToForm("Root_alter", instanceWithInferedType.Root_alter, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Root_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Scaling:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Millimeters", instanceWithInferedType.Millimeters, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tenths", instanceWithInferedType.Tenths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Scordatura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Accord", instanceWithInferedType, &instanceWithInferedType.Accord, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Score_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Instrument_name", instanceWithInferedType.Instrument_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Instrument_abbreviation", instanceWithInferedType.Instrument_abbreviation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Instrument_sound", instanceWithInferedType.Instrument_sound, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Solo", instanceWithInferedType.Solo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ensemble", instanceWithInferedType.Ensemble, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Virtual_instrument", instanceWithInferedType.Virtual_instrument, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Score_part, *models.Score_instrument](
				"Score_part",
				"Score_instrument",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Score_part) []*models.Score_instrument {
					return owner.Score_instrument
				})
		}

	case *models.Score_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationSliceToForm("Part_link", instanceWithInferedType, &instanceWithInferedType.Part_link, formGroup, probe)
		AssociationFieldToForm("Part_name", instanceWithInferedType.Part_name, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation", instanceWithInferedType.Part_abbreviation, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)
		BasicFieldtoForm("Group", instanceWithInferedType.Group, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Score_instrument", instanceWithInferedType, &instanceWithInferedType.Score_instrument, formGroup, probe)
		AssociationSliceToForm("Player", instanceWithInferedType, &instanceWithInferedType.Player, formGroup, probe)
		AssociationSliceToForm("Midi_device", instanceWithInferedType, &instanceWithInferedType.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", instanceWithInferedType, &instanceWithInferedType.Midi_instrument, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Score_partwise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Version", instanceWithInferedType.Version, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Work", instanceWithInferedType.Work, formGroup, probe)
		BasicFieldtoForm("Movement_number", instanceWithInferedType.Movement_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Movement_title", instanceWithInferedType.Movement_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationFieldToForm("Defaults", instanceWithInferedType.Defaults, formGroup, probe)
		AssociationSliceToForm("Credit", instanceWithInferedType, &instanceWithInferedType.Credit, formGroup, probe)
		AssociationFieldToForm("Part_list", instanceWithInferedType.Part_list, formGroup, probe)
		AssociationSliceToForm("Part", instanceWithInferedType, &instanceWithInferedType.Part, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Score_timewise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Version", instanceWithInferedType.Version, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Work", instanceWithInferedType.Work, formGroup, probe)
		BasicFieldtoForm("Movement_number", instanceWithInferedType.Movement_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Movement_title", instanceWithInferedType.Movement_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationFieldToForm("Defaults", instanceWithInferedType.Defaults, formGroup, probe)
		AssociationSliceToForm("Credit", instanceWithInferedType, &instanceWithInferedType.Credit, formGroup, probe)
		AssociationFieldToForm("Part_list", instanceWithInferedType.Part_list, formGroup, probe)
		AssociationSliceToForm("Measure", instanceWithInferedType, &instanceWithInferedType.Measure, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Segno:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Direction_type, *models.Segno](
				"Direction_type",
				"Segno",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Direction_type) []*models.Segno {
					return owner.Segno
				})
		}

	case *models.Slash:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Use_dots", instanceWithInferedType.Use_dots, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Use_stems", instanceWithInferedType.Use_stems, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Slash_type", instanceWithInferedType.Slash_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Slash_dot", instanceWithInferedType.Slash_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Except_voice", instanceWithInferedType.Except_voice, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Slide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("First_beat", instanceWithInferedType.First_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Slide](
				"Notations",
				"Slide",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Slide {
					return owner.Slide
				})
		}

	case *models.Slur:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Orientation", instanceWithInferedType.Orientation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_x", instanceWithInferedType.Bezier_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_y", instanceWithInferedType.Bezier_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_x2", instanceWithInferedType.Bezier_x2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_y2", instanceWithInferedType.Bezier_y2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_offset", instanceWithInferedType.Bezier_offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_offset2", instanceWithInferedType.Bezier_offset2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Slur](
				"Notations",
				"Slur",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Slur {
					return owner.Slur
				})
		}

	case *models.Sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tempo", instanceWithInferedType.Tempo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dynamics", instanceWithInferedType.Dynamics, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Dacapo", instanceWithInferedType.Dacapo, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dalsegno", instanceWithInferedType.Dalsegno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tocoda", instanceWithInferedType.Tocoda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Divisions", instanceWithInferedType.Divisions, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Forward_repeat", instanceWithInferedType.Forward_repeat, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Fine", instanceWithInferedType.Fine, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Pizzicato", instanceWithInferedType.Pizzicato, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Pan", instanceWithInferedType.Pan, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Elevation", instanceWithInferedType.Elevation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Damper_pedal", instanceWithInferedType.Damper_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Soft_pedal", instanceWithInferedType.Soft_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Sostenuto_pedal", instanceWithInferedType.Sostenuto_pedal, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Instrument_change", instanceWithInferedType, &instanceWithInferedType.Instrument_change, formGroup, probe)
		AssociationSliceToForm("Midi_device", instanceWithInferedType, &instanceWithInferedType.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", instanceWithInferedType, &instanceWithInferedType.Midi_instrument, formGroup, probe)
		AssociationSliceToForm("Play", instanceWithInferedType, &instanceWithInferedType.Play, formGroup, probe)
		AssociationFieldToForm("Swing", instanceWithInferedType.Swing, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_measure, *models.Sound](
				"A_measure",
				"Sound",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_measure) []*models.Sound {
					return owner.Sound
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_part_1, *models.Sound](
				"A_part_1",
				"Sound",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_part_1) []*models.Sound {
					return owner.Sound
				})
		}

	case *models.Staff_details:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Show_frets", instanceWithInferedType.Show_frets, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Print_spacing", instanceWithInferedType.Print_spacing, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Staff_type", instanceWithInferedType.Staff_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff_lines", instanceWithInferedType.Staff_lines, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Line_detail", instanceWithInferedType, &instanceWithInferedType.Line_detail, formGroup, probe)
		AssociationSliceToForm("Staff_tuning", instanceWithInferedType, &instanceWithInferedType.Staff_tuning, formGroup, probe)
		BasicFieldtoForm("Capo", instanceWithInferedType.Capo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Staff_size", instanceWithInferedType.Staff_size, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Staff_details](
				"Attributes",
				"Staff_details",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Staff_details {
					return owner.Staff_details
				})
		}

	case *models.Staff_divide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Staff_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Staff_distance", instanceWithInferedType.Staff_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Defaults, *models.Staff_layout](
				"Defaults",
				"Staff_layout",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Defaults) []*models.Staff_layout {
					return owner.Staff_layout
				})
		}
		{
			AssociationReverseSliceToForm[*models.Print, *models.Staff_layout](
				"Print",
				"Staff_layout",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Print) []*models.Staff_layout {
					return owner.Staff_layout
				})
		}

	case *models.Staff_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Scaling", instanceWithInferedType.Scaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Staff_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line", instanceWithInferedType.Line, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Tuning_step", instanceWithInferedType.Tuning_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Tuning_alter", instanceWithInferedType.Tuning_alter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Tuning_octave", instanceWithInferedType.Tuning_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Staff_details, *models.Staff_tuning](
				"Staff_details",
				"Staff_tuning",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Staff_details) []*models.Staff_tuning {
					return owner.Staff_tuning
				})
		}

	case *models.Stem:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Stick:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Tip", instanceWithInferedType.Tip, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Parentheses", instanceWithInferedType.Parentheses, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Dashed_circle", instanceWithInferedType.Dashed_circle, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Stick_type", instanceWithInferedType.Stick_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Stick_material", instanceWithInferedType.Stick_material, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.String_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.String_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.String_type](
				"Technical",
				"String",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.String_type {
					return owner.String
				})
		}

	case *models.Strong_accent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Articulations, *models.Strong_accent](
				"Articulations",
				"Strong_accent",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Articulations) []*models.Strong_accent {
					return owner.Strong_accent
				})
		}

	case *models.Style_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Supports:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Element", instanceWithInferedType.Element, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Attribute", instanceWithInferedType.Attribute, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Encoding, *models.Supports](
				"Encoding",
				"Supports",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Encoding) []*models.Supports {
					return owner.Supports
				})
		}

	case *models.Swing:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Straight", instanceWithInferedType.Straight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("First", instanceWithInferedType.First, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Second", instanceWithInferedType.Second, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Swing_type", instanceWithInferedType.Swing_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Swing_style", instanceWithInferedType.Swing_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Sync:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Latency", instanceWithInferedType.Latency, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Listening, *models.Sync](
				"Listening",
				"Sync",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Listening) []*models.Sync {
					return owner.Sync
				})
		}

	case *models.System_dividers:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Left_divider", instanceWithInferedType.Left_divider, formGroup, probe)
		AssociationFieldToForm("Right_divider", instanceWithInferedType.Right_divider, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.System_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("System_margins", instanceWithInferedType.System_margins, formGroup, probe)
		BasicFieldtoForm("System_distance", instanceWithInferedType.System_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Top_system_distance", instanceWithInferedType.Top_system_distance, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("System_dividers", instanceWithInferedType.System_dividers, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.System_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Left_margin", instanceWithInferedType.Left_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Right_margin", instanceWithInferedType.Right_margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Tap:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Hand", instanceWithInferedType.Hand, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Technical, *models.Tap](
				"Technical",
				"Tap",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Technical) []*models.Tap {
					return owner.Tap
				})
		}

	case *models.Technical:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Technical](
				"Notations",
				"Technical",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Technical {
					return owner.Technical
				})
		}

	case *models.Text_element_data:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Underline", instanceWithInferedType.Underline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Overline", instanceWithInferedType.Overline, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_through", instanceWithInferedType.Line_through, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Letter_spacing", instanceWithInferedType.Letter_spacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dir", instanceWithInferedType.Dir, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Lyric, *models.Text_element_data](
				"Lyric",
				"Text",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Lyric) []*models.Text_element_data {
					return owner.Text
				})
		}

	case *models.Tie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Orientation", instanceWithInferedType.Orientation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_x", instanceWithInferedType.Bezier_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_y", instanceWithInferedType.Bezier_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_x2", instanceWithInferedType.Bezier_x2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_y2", instanceWithInferedType.Bezier_y2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_offset", instanceWithInferedType.Bezier_offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bezier_offset2", instanceWithInferedType.Bezier_offset2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Tied](
				"Notations",
				"Tied",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Tied {
					return owner.Tied
				})
		}

	case *models.Time:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Symbol", instanceWithInferedType.Symbol, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Separator", instanceWithInferedType.Separator, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Halign", instanceWithInferedType.Halign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Valign", instanceWithInferedType.Valign, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Print_object", instanceWithInferedType.Print_object, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Beat_type", instanceWithInferedType.Beat_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Interchangeable", instanceWithInferedType.Interchangeable, formGroup, probe)
		BasicFieldtoForm("Senza_misura", instanceWithInferedType.Senza_misura, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Time](
				"Attributes",
				"Time",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Time {
					return owner.Time
				})
		}

	case *models.Time_modification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Actual_notes", instanceWithInferedType.Actual_notes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Normal_notes", instanceWithInferedType.Normal_notes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Normal_type", instanceWithInferedType.Normal_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Normal_dot", instanceWithInferedType.Normal_dot, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Timpani:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Diatonic", instanceWithInferedType.Diatonic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Chromatic", instanceWithInferedType.Chromatic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Octave_change", instanceWithInferedType.Octave_change, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Double", instanceWithInferedType.Double, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Attributes, *models.Transpose](
				"Attributes",
				"Transpose",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Attributes) []*models.Transpose {
					return owner.Transpose
				})
		}

	case *models.Tremolo:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Tremolo](
				"Ornaments",
				"Tremolo",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Tremolo {
					return owner.Tremolo
				})
		}

	case *models.Tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Bracket", instanceWithInferedType.Bracket, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Show_number", instanceWithInferedType.Show_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Show_type", instanceWithInferedType.Show_type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Line_shape", instanceWithInferedType.Line_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Tuplet_actual", instanceWithInferedType.Tuplet_actual, formGroup, probe)
		AssociationFieldToForm("Tuplet_normal", instanceWithInferedType.Tuplet_normal, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Notations, *models.Tuplet](
				"Notations",
				"Tuplet",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Notations) []*models.Tuplet {
					return owner.Tuplet
				})
		}

	case *models.Tuplet_dot:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Tuplet_portion, *models.Tuplet_dot](
				"Tuplet_portion",
				"Tuplet_dot",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Tuplet_portion) []*models.Tuplet_dot {
					return owner.Tuplet_dot
				})
		}

	case *models.Tuplet_number:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Tuplet_portion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Tuplet_number", instanceWithInferedType.Tuplet_number, formGroup, probe)
		AssociationFieldToForm("Tuplet_type", instanceWithInferedType.Tuplet_type, formGroup, probe)
		AssociationSliceToForm("Tuplet_dot", instanceWithInferedType, &instanceWithInferedType.Tuplet_dot, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Tuplet_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_family", instanceWithInferedType.Font_family, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_style", instanceWithInferedType.Font_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_size", instanceWithInferedType.Font_size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Font_weight", instanceWithInferedType.Font_weight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Typed_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Encoding, *models.Typed_text](
				"Encoding",
				"Encoder",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Encoding) []*models.Typed_text {
					return owner.Encoder
				})
		}
		{
			AssociationReverseSliceToForm[*models.Identification, *models.Typed_text](
				"Identification",
				"Creator",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Identification) []*models.Typed_text {
					return owner.Creator
				})
		}
		{
			AssociationReverseSliceToForm[*models.Identification, *models.Typed_text](
				"Identification",
				"Rights",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Identification) []*models.Typed_text {
					return owner.Rights
				})
		}
		{
			AssociationReverseSliceToForm[*models.Identification, *models.Typed_text](
				"Identification",
				"Relation",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Identification) []*models.Typed_text {
					return owner.Relation
				})
		}

	case *models.Unpitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Display_step", instanceWithInferedType.Display_step, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Display_octave", instanceWithInferedType.Display_octave, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Virtual_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Virtual_library", instanceWithInferedType.Virtual_library, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Virtual_name", instanceWithInferedType.Virtual_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Wait:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Player", instanceWithInferedType.Player, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Time_only", instanceWithInferedType.Time_only, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Listen, *models.Wait](
				"Listen",
				"Wait",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Listen) []*models.Wait {
					return owner.Wait
				})
		}

	case *models.Wavy_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Placement", instanceWithInferedType.Placement, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Start_note", instanceWithInferedType.Start_note, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Trill_step", instanceWithInferedType.Trill_step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Two_note_turn", instanceWithInferedType.Two_note_turn, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Accelerate", instanceWithInferedType.Accelerate, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Beats", instanceWithInferedType.Beats, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Second_beat", instanceWithInferedType.Second_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Last_beat", instanceWithInferedType.Last_beat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Ornaments, *models.Wavy_line](
				"Ornaments",
				"Wavy_line",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ornaments) []*models.Wavy_line {
					return owner.Wavy_line
				})
		}

	case *models.Wedge:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Spread", instanceWithInferedType.Spread, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Niente", instanceWithInferedType.Niente, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Line_type", instanceWithInferedType.Line_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Dash_length", instanceWithInferedType.Dash_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Space_length", instanceWithInferedType.Space_length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_x", instanceWithInferedType.Default_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default_y", instanceWithInferedType.Default_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_x", instanceWithInferedType.Relative_x, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Relative_y", instanceWithInferedType.Relative_y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Wood:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Smufl", instanceWithInferedType.Smufl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Work:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Work_number", instanceWithInferedType.Work_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Work_title", instanceWithInferedType.Work_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Opus", instanceWithInferedType.Opus, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
