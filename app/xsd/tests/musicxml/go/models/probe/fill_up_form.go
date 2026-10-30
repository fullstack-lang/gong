// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xmlxsd"
)

const FormName = "Form"

var (
	_ = sort.Strings
	_ = strings.Join
)

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {
	switch inst := instance.(type) {
	case *models.A_directive:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Directive {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Directive", refNames, formGroup, probe.formStage)
		}
	case *models.A_measure:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Implicit", inst.Implicit, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Non_controlling", inst.Non_controlling, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Note", inst, &inst.Note, formGroup, probe)
		AssociationSliceToForm("Backup", inst, &inst.Backup, formGroup, probe)
		AssociationSliceToForm("Forward", inst, &inst.Forward, formGroup, probe)
		AssociationSliceToForm("Direction", inst, &inst.Direction, formGroup, probe)
		AssociationSliceToForm("Attributes", inst, &inst.Attributes, formGroup, probe)
		AssociationSliceToForm("Harmony", inst, &inst.Harmony, formGroup, probe)
		AssociationSliceToForm("Figured_bass", inst, &inst.Figured_bass, formGroup, probe)
		AssociationSliceToForm("Print_XSD", inst, &inst.Print_XSD, formGroup, probe)
		AssociationSliceToForm("Sound", inst, &inst.Sound, formGroup, probe)
		AssociationSliceToForm("Listening", inst, &inst.Listening, formGroup, probe)
		AssociationSliceToForm("Barline", inst, &inst.Barline, formGroup, probe)
		AssociationSliceToForm("Grouping", inst, &inst.Grouping, formGroup, probe)
		AssociationSliceToForm("Link", inst, &inst.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", inst, &inst.Bookmark, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_parts {
				for _, target := range src.Measure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part", "Measure", refNames, formGroup, probe.formStage)
		}
	case *models.A_measure_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Implicit", inst.Implicit, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Non_controlling", inst.Non_controlling, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Part", inst, &inst.Part, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				for _, target := range src.Measure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Measure", refNames, formGroup, probe.formStage)
		}
	case *models.A_part:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Measure", inst, &inst.Measure, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				for _, target := range src.Part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Part", refNames, formGroup, probe.formStage)
		}
	case *models.A_part_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Note", inst, &inst.Note, formGroup, probe)
		AssociationSliceToForm("Backup", inst, &inst.Backup, formGroup, probe)
		AssociationSliceToForm("Forward", inst, &inst.Forward, formGroup, probe)
		AssociationSliceToForm("Direction", inst, &inst.Direction, formGroup, probe)
		AssociationSliceToForm("Attributes", inst, &inst.Attributes, formGroup, probe)
		AssociationSliceToForm("Harmony", inst, &inst.Harmony, formGroup, probe)
		AssociationSliceToForm("Figured_bass", inst, &inst.Figured_bass, formGroup, probe)
		AssociationSliceToForm("Print_XSD", inst, &inst.Print_XSD, formGroup, probe)
		AssociationSliceToForm("Sound", inst, &inst.Sound, formGroup, probe)
		AssociationSliceToForm("Listening", inst, &inst.Listening, formGroup, probe)
		AssociationSliceToForm("Barline", inst, &inst.Barline, formGroup, probe)
		AssociationSliceToForm("Grouping", inst, &inst.Grouping, formGroup, probe)
		AssociationSliceToForm("Link", inst, &inst.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", inst, &inst.Bookmark, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measure_1s {
				for _, target := range src.Part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure_1", "Part", refNames, formGroup, probe.formStage)
		}
	case *models.Accidental:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Cautionary", inst.Cautionary, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Editorial", inst.Editorial, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Accidental == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Accidental", refNames, formGroup, probe.formStage)
		}
	case *models.Accidental_mark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Accidental_mark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Accidental_mark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Accidental_mark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Accidental_mark", refNames, formGroup, probe.formStage)
		}
	case *models.Accidental_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Name_displays {
				for _, target := range src.Accidental_text {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Name_display", "Accidental_text", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notehead_texts {
				for _, target := range src.Accidental_text {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notehead_text", "Accidental_text", refNames, formGroup, probe.formStage)
		}
	case *models.Accord:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("String_XSD", inst.String_XSD, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tuning_step", inst.Tuning_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Tuning_alter", inst.Tuning_alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tuning_octave", inst.Tuning_octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Scordaturas {
				for _, target := range src.Accord {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scordatura", "Accord", refNames, formGroup, probe.formStage)
		}
	case *models.Accordion_registration:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_high", inst.Accordion_high, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_middle", inst.Accordion_middle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_low", inst.Accordion_low, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Accordion_registration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Accordion_registration", refNames, formGroup, probe.formStage)
		}
	case *models.Appearance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Line_width", inst, &inst.Line_width, formGroup, probe)
		AssociationSliceToForm("Note_size", inst, &inst.Note_size, formGroup, probe)
		AssociationSliceToForm("Distance", inst, &inst.Distance, formGroup, probe)
		AssociationSliceToForm("Glyph", inst, &inst.Glyph, formGroup, probe)
		AssociationSliceToForm("Other_appearance", inst, &inst.Other_appearance, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.Appearance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Appearance", refNames, formGroup, probe.formStage)
		}
	case *models.Arpeggiate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Unbroken", inst.Unbroken, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Arpeggiate {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Arpeggiate", refNames, formGroup, probe.formStage)
		}
	case *models.Arrow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Arrow_direction", inst.Arrow_direction, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Arrow_style", inst.Arrow_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Arrowhead", inst.Arrowhead, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Circular_arrow", inst.Circular_arrow, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Arrow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Arrow", refNames, formGroup, probe.formStage)
		}
	case *models.Articulations:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Accent", inst, &inst.Accent, formGroup, probe)
		AssociationSliceToForm("Strong_accent", inst, &inst.Strong_accent, formGroup, probe)
		AssociationSliceToForm("Staccato", inst, &inst.Staccato, formGroup, probe)
		AssociationSliceToForm("Tenuto", inst, &inst.Tenuto, formGroup, probe)
		AssociationSliceToForm("Detached_legato", inst, &inst.Detached_legato, formGroup, probe)
		AssociationSliceToForm("Staccatissimo", inst, &inst.Staccatissimo, formGroup, probe)
		AssociationSliceToForm("Spiccato", inst, &inst.Spiccato, formGroup, probe)
		AssociationSliceToForm("Scoop", inst, &inst.Scoop, formGroup, probe)
		AssociationSliceToForm("Plop", inst, &inst.Plop, formGroup, probe)
		AssociationSliceToForm("Doit", inst, &inst.Doit, formGroup, probe)
		AssociationSliceToForm("Falloff", inst, &inst.Falloff, formGroup, probe)
		AssociationSliceToForm("Breath_mark", inst, &inst.Breath_mark, formGroup, probe)
		AssociationSliceToForm("Caesura", inst, &inst.Caesura, formGroup, probe)
		AssociationSliceToForm("Stress", inst, &inst.Stress, formGroup, probe)
		AssociationSliceToForm("Unstress", inst, &inst.Unstress, formGroup, probe)
		AssociationSliceToForm("Soft_accent", inst, &inst.Soft_accent, formGroup, probe)
		AssociationSliceToForm("Other_articulation", inst, &inst.Other_articulation, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Articulations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Articulations", refNames, formGroup, probe.formStage)
		}
	case *models.Assess:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listens {
				for _, target := range src.Assess {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listen", "Assess", refNames, formGroup, probe.formStage)
		}
	case *models.Attributes:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		AssociationSliceToForm("Key", inst, &inst.Key, formGroup, probe)
		StageSetBasicFieldtoForm("Time", inst.Time, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staves", inst.Staves, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Part_symbol", inst.Part_symbol, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_symbol](), probe.formStage)
		StageSetBasicFieldtoForm("Instruments", inst.Instruments, probe.formStage, formGroup)
		AssociationSliceToForm("Clef", inst, &inst.Clef, formGroup, probe)
		AssociationSliceToForm("Staff_details", inst, &inst.Staff_details, formGroup, probe)
		AssociationSliceToForm("Transpose", inst, &inst.Transpose, formGroup, probe)
		AssociationSliceToForm("For_part", inst, &inst.For_part, formGroup, probe)
		AssociationSliceToForm("Directive", inst, &inst.Directive, formGroup, probe)
		AssociationSliceToForm("Measure_style", inst, &inst.Measure_style, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Attributes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Attributes", refNames, formGroup, probe.formStage)
		}
	case *models.Backup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Backup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Backup", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Backup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Backup", refNames, formGroup, probe.formStage)
		}
	case *models.Bar_style_color:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Bar_style == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Bar_style", refNames, formGroup, probe.formStage)
		}
	case *models.Barline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Segno", inst.Segno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Coda", inst.Coda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bar_style", inst.Bar_style, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Bar_style_color](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetAssociationFieldToForm("Wavy_line", inst.Wavy_line, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Wavy_line](), probe.formStage)
		StageSetAssociationFieldToForm("Segno_1", inst.Segno_1, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Segno](), probe.formStage)
		StageSetAssociationFieldToForm("Coda_1", inst.Coda_1, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Coda](), probe.formStage)
		StageSetAssociationFieldToForm("Fermata", inst.Fermata, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Fermata](), probe.formStage)
		StageSetAssociationFieldToForm("Ending", inst.Ending, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Ending](), probe.formStage)
		StageSetAssociationFieldToForm("Repeat", inst.Repeat, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Repeat](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Barline {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Barline", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Barline {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Barline", refNames, formGroup, probe.formStage)
		}
	case *models.Barre:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frame_notes {
				if src.Barre == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Barre", refNames, formGroup, probe.formStage)
		}
	case *models.Bass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Arrangement", inst.Arrangement, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Bass_separator", inst.Bass_separator, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Bass_step", inst.Bass_step, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Bass_step](), probe.formStage)
		StageSetAssociationFieldToForm("Bass_alter", inst.Bass_alter, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Bass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Bass", refNames, formGroup, probe.formStage)
		}
	case *models.Bass_step:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Basss {
				if src.Bass_step == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bass", "Bass_step", refNames, formGroup, probe.formStage)
		}
	case *models.Beam:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Repeater", inst.Repeater, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Fan", inst.Fan, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Beam == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Beam", refNames, formGroup, probe.formStage)
		}
	case *models.Beat_repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Use_dots", inst.Use_dots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Slash_type", inst.Slash_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slash_dot", inst.Slash_dot, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Except_voice", inst.Except_voice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Measure_styles {
				if src.Beat_repeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Measure_style", "Beat_repeat", refNames, formGroup, probe.formStage)
		}
	case *models.Beat_unit_tied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Beat_unit", inst.Beat_unit, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beat_unit_dot", inst.Beat_unit_dot, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronomes {
				for _, target := range src.Beat_unit_tied {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome", "Beat_unit_tied", refNames, formGroup, probe.formStage)
		}
	case *models.Beater:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tip", inst.Tip, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Beater == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Beater", refNames, formGroup, probe.formStage)
		}
	case *models.Bend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Shape", inst.Shape, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("First_beat", inst.First_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bend_alter", inst.Bend_alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pre_bend", inst.Pre_bend, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Release", inst.Release, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Release](), probe.formStage)
		StageSetAssociationFieldToForm("With_bar", inst.With_bar, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Placement_text](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Bend {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Bend", refNames, formGroup, probe.formStage)
		}
	case *models.Bookmark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Position", inst.Position, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Bookmark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Bookmark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Bookmark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Bookmark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Credits {
				for _, target := range src.Bookmark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Bookmark", refNames, formGroup, probe.formStage)
		}
	case *models.Bracket:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_end", inst.Line_end, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("End_length", inst.End_length, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Bracket == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Bracket", refNames, formGroup, probe.formStage)
		}
	case *models.Breath_mark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Breath_mark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Breath_mark", refNames, formGroup, probe.formStage)
		}
	case *models.Caesura:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Caesura {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Caesura", refNames, formGroup, probe.formStage)
		}
	case *models.Cancel:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Keys {
				if src.Cancel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Key", "Cancel", refNames, formGroup, probe.formStage)
		}
	case *models.Clef:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Additional", inst.Additional, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("After_barline", inst.After_barline, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Clef_octave_change", inst.Clef_octave_change, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Clef {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Clef", refNames, formGroup, probe.formStage)
		}
	case *models.Coda:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Coda_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Coda_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Coda {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Coda", refNames, formGroup, probe.formStage)
		}
	case *models.Credit:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page", inst.Page, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Credit_type", inst.Credit_type, probe.formStage, formGroup)
		AssociationSliceToForm("Link", inst, &inst.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", inst, &inst.Bookmark, formGroup, probe)
		StageSetAssociationFieldToForm("Credit_image", inst.Credit_image, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Image](), probe.formStage)
		StageSetAssociationFieldToForm("Credit_words", inst.Credit_words, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text_id](), probe.formStage)
		StageSetAssociationFieldToForm("Credit_symbol", inst.Credit_symbol, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_symbol_id](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				for _, target := range src.Credit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Credit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				for _, target := range src.Credit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Credit", refNames, formGroup, probe.formStage)
		}
	case *models.Dashes:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Dashes == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Dashes", refNames, formGroup, probe.formStage)
		}
	case *models.Defaults:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Scaling", inst.Scaling, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Scaling](), probe.formStage)
		StageSetBasicFieldtoForm("Concert_score", inst.Concert_score, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Page_layout", inst.Page_layout, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Page_layout](), probe.formStage)
		StageSetAssociationFieldToForm("System_layout", inst.System_layout, formGroup, probe.stageOfInterest.GetInstancesSet[*models.System_layout](), probe.formStage)
		AssociationSliceToForm("Staff_layout", inst, &inst.Staff_layout, formGroup, probe)
		StageSetAssociationFieldToForm("Appearance", inst.Appearance, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Appearance](), probe.formStage)
		StageSetAssociationFieldToForm("Music_font", inst.Music_font, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_font](), probe.formStage)
		StageSetAssociationFieldToForm("Word_font", inst.Word_font, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_font](), probe.formStage)
		AssociationSliceToForm("Lyric_font", inst, &inst.Lyric_font, formGroup, probe)
		AssociationSliceToForm("Lyric_language", inst, &inst.Lyric_language, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				if src.Defaults == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Defaults", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				if src.Defaults == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Defaults", refNames, formGroup, probe.formStage)
		}
	case *models.Degree:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Degree_value", inst.Degree_value, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Degree_value](), probe.formStage)
		StageSetAssociationFieldToForm("Degree_alter", inst.Degree_alter, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Degree_alter](), probe.formStage)
		StageSetAssociationFieldToForm("Degree_type", inst.Degree_type, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Degree_type](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				for _, target := range src.Degree {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Degree", refNames, formGroup, probe.formStage)
		}
	case *models.Degree_alter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Plus_minus", inst.Plus_minus, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Degrees {
				if src.Degree_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Degree", "Degree_alter", refNames, formGroup, probe.formStage)
		}
	case *models.Degree_type:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Degrees {
				if src.Degree_type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Degree", "Degree_type", refNames, formGroup, probe.formStage)
		}
	case *models.Degree_value:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Degrees {
				if src.Degree_value == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Degree", "Degree_value", refNames, formGroup, probe.formStage)
		}
	case *models.Direction:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Directive", inst.Directive, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("System", inst.System, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Direction_type", inst, &inst.Direction_type, formGroup, probe)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Sound", inst.Sound, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Sound](), probe.formStage)
		StageSetAssociationFieldToForm("Listening", inst.Listening, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Listening](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Direction {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Direction", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Direction {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Direction", refNames, formGroup, probe.formStage)
		}
	case *models.Direction_type:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Rehearsal", inst, &inst.Rehearsal, formGroup, probe)
		AssociationSliceToForm("Segno", inst, &inst.Segno, formGroup, probe)
		AssociationSliceToForm("Coda", inst, &inst.Coda, formGroup, probe)
		AssociationSliceToForm("Words", inst, &inst.Words, formGroup, probe)
		AssociationSliceToForm("Symbol", inst, &inst.Symbol, formGroup, probe)
		StageSetAssociationFieldToForm("Wedge", inst.Wedge, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Wedge](), probe.formStage)
		AssociationSliceToForm("Dynamics", inst, &inst.Dynamics, formGroup, probe)
		StageSetAssociationFieldToForm("Dashes", inst.Dashes, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Dashes](), probe.formStage)
		StageSetAssociationFieldToForm("Bracket", inst.Bracket, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Bracket](), probe.formStage)
		StageSetAssociationFieldToForm("Pedal", inst.Pedal, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Pedal](), probe.formStage)
		StageSetAssociationFieldToForm("Metronome", inst.Metronome, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Metronome](), probe.formStage)
		StageSetAssociationFieldToForm("Octave_shift", inst.Octave_shift, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Octave_shift](), probe.formStage)
		StageSetAssociationFieldToForm("Harp_pedals", inst.Harp_pedals, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Harp_pedals](), probe.formStage)
		StageSetAssociationFieldToForm("Damp", inst.Damp, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("Damp_all", inst.Damp_all, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("Eyeglasses", inst.Eyeglasses, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("String_mute", inst.String_mute, formGroup, probe.stageOfInterest.GetInstancesSet[*models.String_mute](), probe.formStage)
		StageSetAssociationFieldToForm("Scordatura", inst.Scordatura, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Scordatura](), probe.formStage)
		StageSetAssociationFieldToForm("Image", inst.Image, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Image](), probe.formStage)
		StageSetAssociationFieldToForm("Principal_voice", inst.Principal_voice, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Principal_voice](), probe.formStage)
		AssociationSliceToForm("Percussion", inst, &inst.Percussion, formGroup, probe)
		StageSetAssociationFieldToForm("Accordion_registration", inst.Accordion_registration, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Accordion_registration](), probe.formStage)
		StageSetAssociationFieldToForm("Staff_divide", inst.Staff_divide, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Staff_divide](), probe.formStage)
		StageSetAssociationFieldToForm("Other_direction", inst.Other_direction, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Other_direction](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				for _, target := range src.Direction_type {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Direction_type", refNames, formGroup, probe.formStage)
		}
	case *models.Distance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Appearances {
				for _, target := range src.Distance {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Appearance", "Distance", refNames, formGroup, probe.formStage)
		}
	case *models.Double:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Above", inst.Above, formGroup, probe.formStage)
	case *models.Dynamics:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("P", inst.P, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pp", inst.Pp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ppp", inst.Ppp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pppp", inst.Pppp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ppppp", inst.Ppppp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pppppp", inst.Pppppp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("F", inst.F, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ff", inst.Ff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fff", inst.Fff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ffff", inst.Ffff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fffff", inst.Fffff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ffffff", inst.Ffffff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Mp", inst.Mp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Mf", inst.Mf, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sf", inst.Sf, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sfp", inst.Sfp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sfpp", inst.Sfpp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fp", inst.Fp, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rf", inst.Rf, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rfz", inst.Rfz, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sfz", inst.Sfz, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sffz", inst.Sffz, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fz", inst.Fz, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("N", inst.N, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pf", inst.Pf, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sfzp", inst.Sfzp, probe.formStage, formGroup)
		AssociationSliceToForm("Other_dynamics", inst, &inst.Other_dynamics, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Dynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Dynamics", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Dynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Dynamics", refNames, formGroup, probe.formStage)
		}
	case *models.Effect:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Effect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Effect", refNames, formGroup, probe.formStage)
		}
	case *models.Elision:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Lyrics {
				for _, target := range src.Elision {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Elision", refNames, formGroup, probe.formStage)
		}
	case *models.Empty:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.Empty_font:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.Music_font == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Music_font", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.Word_font == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Word_font", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_line:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_shape", inst.Line_shape, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Line_length", inst.Line_length, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Scoop {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Scoop", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Plop {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Plop", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Doit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Doit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Falloff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Falloff", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_placement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Accent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Accent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Staccato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Staccato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Tenuto {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Tenuto", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Detached_legato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Detached_legato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Staccatissimo {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Staccatissimo", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Spiccato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Spiccato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Stress {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Stress", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Unstress {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Unstress", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Soft_accent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Soft_accent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				for _, target := range src.Dot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Dot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Schleifer {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Schleifer", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Up_bow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Up_bow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Down_bow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Down_bow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Open_string {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Open_string", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Thumb_position {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Thumb_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Double_tongue {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Double_tongue", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Triple_tongue {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Triple_tongue", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Snap_pizzicato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Snap_pizzicato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Fingernails {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Fingernails", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Brass_bend {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Brass_bend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Flip {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Flip", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Smear {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Smear", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Golpe {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Golpe", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_placement_smufl:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Stopped {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Stopped", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Open {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Open", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Half_muted {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Half_muted", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_print_object_style_align:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.System_dividerss {
				if src.Left_divider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_dividers", "Left_divider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.System_dividerss {
				if src.Right_divider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_dividers", "Right_divider", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_print_style:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
	case *models.Empty_print_style_align:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
	case *models.Empty_print_style_align_id:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Damp == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Damp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Damp_all == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Damp_all", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Eyeglasses == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Eyeglasses", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_trill_sound:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Start_note", inst.Start_note, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Trill_step", inst.Trill_step, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Two_note_turn", inst.Two_note_turn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Trill_mark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Trill_mark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Vertical_turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Vertical_turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Inverted_vertical_turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Inverted_vertical_turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Shake {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Shake", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Haydn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Haydn", refNames, formGroup, probe.formStage)
		}
	case *models.Encoding:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Encoding_date", inst.Encoding_date, probe.formStage, formGroup)
		AssociationSliceToForm("Encoder", inst, &inst.Encoder, formGroup, probe)
		StageSetBasicFieldtoForm("Software", inst.Software, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Encoding_description", inst.Encoding_description, probe.formStage, formGroup)
		AssociationSliceToForm("Supports", inst, &inst.Supports, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Identifications {
				if src.Encoding == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Identification", "Encoding", refNames, formGroup, probe.formStage)
		}
	case *models.Ending:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("End_length", inst.End_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text_x", inst.Text_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text_y", inst.Text_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("System", inst.System, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Ending == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Ending", refNames, formGroup, probe.formStage)
		}
	case *models.Extend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Extend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Lyrics {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Extend", refNames, formGroup, probe.formStage)
		}
	case *models.Feature:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Groupings {
				for _, target := range src.Feature {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Grouping", "Feature", refNames, formGroup, probe.formStage)
		}
	case *models.Fermata:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Fermata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Fermata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Fermata {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Fermata", refNames, formGroup, probe.formStage)
		}
	case *models.Figure:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Prefix", inst.Prefix, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Figure_number", inst.Figure_number, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Suffix", inst.Suffix, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Extend](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figured_basss {
				for _, target := range src.Figure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figured_bass", "Figure", refNames, formGroup, probe.formStage)
		}
	case *models.Figured_bass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_dot", inst.Print_dot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_lyric", inst.Print_lyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Figure", inst, &inst.Figure, formGroup, probe)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Figured_bass {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Figured_bass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Figured_bass {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Figured_bass", refNames, formGroup, probe.formStage)
		}
	case *models.Fingering:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Substitution", inst.Substitution, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Alternate", inst.Alternate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frame_notes {
				if src.Fingering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Fingering", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Fingering {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Fingering", refNames, formGroup, probe.formStage)
		}
	case *models.First_fret:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frames {
				if src.First_fret == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame", "First_fret", refNames, formGroup, probe.formStage)
		}
	case *models.For_part:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Part_clef", inst.Part_clef, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_clef](), probe.formStage)
		StageSetAssociationFieldToForm("Part_transpose", inst.Part_transpose, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_transpose](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.For_part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "For_part", refNames, formGroup, probe.formStage)
		}
	case *models.Formatted_symbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)
	case *models.Formatted_symbol_id:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Credits {
				if src.Credit_symbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_symbol", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Symbol {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Symbol", refNames, formGroup, probe.formStage)
		}
	case *models.Formatted_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Backups {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Backup", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figured_basss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figured_bass", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Forwards {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Forward", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Lyrics {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Name_displays {
				for _, target := range src.Display_text {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Name_display", "Display_text", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notehead_texts {
				for _, target := range src.Display_text {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notehead_text", "Display_text", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Footnote", refNames, formGroup, probe.formStage)
		}
	case *models.Formatted_text_id:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Credits {
				if src.Credit_words == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_words", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Rehearsal {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Rehearsal", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Words {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Words", refNames, formGroup, probe.formStage)
		}
	case *models.Forward:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Forward {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Forward", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Forward {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Forward", refNames, formGroup, probe.formStage)
		}
	case *models.Frame:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Unplayed", inst.Unplayed, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame_strings", inst.Frame_strings, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame_frets", inst.Frame_frets, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("First_fret", inst.First_fret, formGroup, probe.stageOfInterest.GetInstancesSet[*models.First_fret](), probe.formStage)
		AssociationSliceToForm("Frame_note", inst, &inst.Frame_note, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Frame == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Frame", refNames, formGroup, probe.formStage)
		}
	case *models.Frame_note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("String_XSD", inst.String_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.String_type](), probe.formStage)
		StageSetAssociationFieldToForm("Fret", inst.Fret, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Fret](), probe.formStage)
		StageSetAssociationFieldToForm("Fingering", inst.Fingering, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Fingering](), probe.formStage)
		StageSetAssociationFieldToForm("Barre", inst.Barre, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Barre](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frames {
				for _, target := range src.Frame_note {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame", "Frame_note", refNames, formGroup, probe.formStage)
		}
	case *models.Fret:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frame_notes {
				if src.Fret == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Fret", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Fret {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Fret", refNames, formGroup, probe.formStage)
		}
	case *models.Glass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Glass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Glass", refNames, formGroup, probe.formStage)
		}
	case *models.Glissando:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Glissando {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Glissando", refNames, formGroup, probe.formStage)
		}
	case *models.Glyph:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Appearances {
				for _, target := range src.Glyph {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Appearance", "Glyph", refNames, formGroup, probe.formStage)
		}
	case *models.Grace:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Steal_time_previous", inst.Steal_time_previous, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Steal_time_following", inst.Steal_time_following, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Make_time", inst.Make_time, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Slash", inst.Slash, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Grace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Grace", refNames, formGroup, probe.formStage)
		}
	case *models.Group_barline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_barline == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_barline", refNames, formGroup, probe.formStage)
		}
	case *models.Group_name:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_name == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_name", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_abbreviation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_abbreviation", refNames, formGroup, probe.formStage)
		}
	case *models.Group_symbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_symbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_symbol", refNames, formGroup, probe.formStage)
		}
	case *models.Grouping:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Member_of", inst.Member_of, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Feature", inst, &inst.Feature, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Grouping {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Grouping", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Grouping {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Grouping", refNames, formGroup, probe.formStage)
		}
	case *models.Hammer_on_pull_off:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Hammer_on {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Hammer_on", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Pull_off {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Pull_off", refNames, formGroup, probe.formStage)
		}
	case *models.Handbell:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Handbell {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Handbell", refNames, formGroup, probe.formStage)
		}
	case *models.Harmon_closed:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmon_mutes {
				if src.Harmon_closed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmon_mute", "Harmon_closed", refNames, formGroup, probe.formStage)
		}
	case *models.Harmon_mute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Harmon_closed", inst.Harmon_closed, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Harmon_closed](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Harmon_mute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Harmon_mute", refNames, formGroup, probe.formStage)
		}
	case *models.Harmonic:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Natural", inst.Natural, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Artificial", inst.Artificial, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base_pitch", inst.Base_pitch, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Touching_pitch", inst.Touching_pitch, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sounding_pitch", inst.Sounding_pitch, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Harmonic {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Harmonic", refNames, formGroup, probe.formStage)
		}
	case *models.Harmony:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_frame", inst.Print_frame, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Arrangement", inst.Arrangement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("System", inst.System, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Root", inst.Root, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Root](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral", inst.Numeral, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Numeral](), probe.formStage)
		StageSetAssociationFieldToForm("Function", inst.Function, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Kind", inst.Kind, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Kind](), probe.formStage)
		StageSetAssociationFieldToForm("Inversion", inst.Inversion, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Inversion](), probe.formStage)
		StageSetAssociationFieldToForm("Bass", inst.Bass, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Bass](), probe.formStage)
		AssociationSliceToForm("Degree", inst, &inst.Degree, formGroup, probe)
		StageSetAssociationFieldToForm("Frame", inst.Frame, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Frame](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Harmony {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Harmony", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Harmony {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Harmony", refNames, formGroup, probe.formStage)
		}
	case *models.Harmony_alter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Basss {
				if src.Bass_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bass", "Bass_alter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Numerals {
				if src.Numeral_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Numeral", "Numeral_alter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Roots {
				if src.Root_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Root", "Root_alter", refNames, formGroup, probe.formStage)
		}
	case *models.Harp_pedals:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Pedal_tuning", inst, &inst.Pedal_tuning, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Harp_pedals == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Harp_pedals", refNames, formGroup, probe.formStage)
		}
	case *models.Heel_toe:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Substitution", inst.Substitution, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Heel {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Heel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Toe {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Toe", refNames, formGroup, probe.formStage)
		}
	case *models.Hole:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Hole_type", inst.Hole_type, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Hole_closed", inst.Hole_closed, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Hole_closed](), probe.formStage)
		StageSetBasicFieldtoForm("Hole_shape", inst.Hole_shape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Hole {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Hole", refNames, formGroup, probe.formStage)
		}
	case *models.Hole_closed:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Holes {
				if src.Hole_closed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Hole", "Hole_closed", refNames, formGroup, probe.formStage)
		}
	case *models.Horizontal_turn:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Slash", inst.Slash, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Start_note", inst.Start_note, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Trill_step", inst.Trill_step, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Two_note_turn", inst.Two_note_turn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Delayed_turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Delayed_turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Inverted_turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Inverted_turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Delayed_inverted_turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Delayed_inverted_turn", refNames, formGroup, probe.formStage)
		}
	case *models.Identification:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Creator", inst, &inst.Creator, formGroup, probe)
		AssociationSliceToForm("Rights", inst, &inst.Rights, formGroup, probe)
		StageSetAssociationFieldToForm("Encoding", inst.Encoding, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Encoding](), probe.formStage)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)
		AssociationSliceToForm("Relation", inst, &inst.Relation, formGroup, probe)
		StageSetAssociationFieldToForm("Miscellaneous", inst.Miscellaneous, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Miscellaneous](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Identification", refNames, formGroup, probe.formStage)
		}
	case *models.Image:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Credits {
				if src.Credit_image == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_image", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Image == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Image", refNames, formGroup, probe.formStage)
		}
	case *models.Instrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				for _, target := range src.Instrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Instrument", refNames, formGroup, probe.formStage)
		}
	case *models.Instrument_change:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Instrument_sound", inst.Instrument_sound, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Solo", inst.Solo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ensemble", inst.Ensemble, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Virtual_instrument", inst.Virtual_instrument, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Virtual_instrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				for _, target := range src.Instrument_change {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Instrument_change", refNames, formGroup, probe.formStage)
		}
	case *models.Instrument_link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_links {
				for _, target := range src.Instrument_link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_link", "Instrument_link", refNames, formGroup, probe.formStage)
		}
	case *models.Interchangeable:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Separator", inst.Separator, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Time_relation", inst.Time_relation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beat_type", inst.Beat_type, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Times {
				if src.Interchangeable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Time", "Interchangeable", refNames, formGroup, probe.formStage)
		}
	case *models.Inversion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Inversion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Inversion", refNames, formGroup, probe.formStage)
		}
	case *models.Key:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Cancel", inst.Cancel, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Cancel](), probe.formStage)
		StageSetBasicFieldtoForm("Fifths", inst.Fifths, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Mode", inst.Mode, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Key_step", inst.Key_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Key_alter", inst.Key_alter, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Key_accidental", inst.Key_accidental, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Key_accidental](), probe.formStage)
		AssociationSliceToForm("Key_octave", inst, &inst.Key_octave, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Key {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Key", refNames, formGroup, probe.formStage)
		}
	case *models.Key_accidental:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Keys {
				if src.Key_accidental == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Key", "Key_accidental", refNames, formGroup, probe.formStage)
		}
	case *models.Key_octave:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Cancel", inst.Cancel, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Keys {
				for _, target := range src.Key_octave {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Key", "Key_octave", refNames, formGroup, probe.formStage)
		}
	case *models.Kind:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Use_symbols", inst.Use_symbols, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Stack_degrees", inst.Stack_degrees, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses_degrees", inst.Parentheses_degrees, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket_degrees", inst.Bracket_degrees, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Kind == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Kind", refNames, formGroup, probe.formStage)
		}
	case *models.Level:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Reference_XSD", inst.Reference_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Backups {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Backup", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figured_basss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figured_bass", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Forwards {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Forward", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Lyrics {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Level", refNames, formGroup, probe.formStage)
		}
	case *models.Line_detail:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Staff_detailss {
				for _, target := range src.Line_detail {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Staff_details", "Line_detail", refNames, formGroup, probe.formStage)
		}
	case *models.Line_width:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Appearances {
				for _, target := range src.Line_width {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Appearance", "Line_width", refNames, formGroup, probe.formStage)
		}
	case *models.Link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Position", inst.Position, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Link", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Link", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Credits {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Link", refNames, formGroup, probe.formStage)
		}
	case *models.Listen:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Assess", inst, &inst.Assess, formGroup, probe)
		AssociationSliceToForm("Wait", inst, &inst.Wait, formGroup, probe)
		AssociationSliceToForm("Other_listen", inst, &inst.Other_listen, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Listen == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Listen", refNames, formGroup, probe.formStage)
		}
	case *models.Listening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Sync", inst, &inst.Sync, formGroup, probe)
		AssociationSliceToForm("Other_listening", inst, &inst.Other_listening, formGroup, probe)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Listening {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Listening", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Listening {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Listening", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				if src.Listening == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Listening", refNames, formGroup, probe.formStage)
		}
	case *models.Lyric:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Syllabic", inst.Syllabic, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Text", inst.Text, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Text_element_data](), probe.formStage)
		AssociationSliceToForm("Elision", inst, &inst.Elision, formGroup, probe)
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Extend](), probe.formStage)
		StageSetBasicFieldtoForm("Laughing", inst.Laughing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Humming", inst.Humming, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_line", inst.End_line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_paragraph", inst.End_paragraph, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				for _, target := range src.Lyric {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Lyric", refNames, formGroup, probe.formStage)
		}
	case *models.Lyric_font:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				for _, target := range src.Lyric_font {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Lyric_font", refNames, formGroup, probe.formStage)
		}
	case *models.Lyric_language:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				for _, target := range src.Lyric_language {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Lyric_language", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Measure_distance", inst.Measure_distance, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.Measure_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Measure_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_numbering:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Multiple_rest_always", inst.Multiple_rest_always, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Multiple_rest_range", inst.Multiple_rest_range, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.Measure_numbering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Measure_numbering", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Measure_styles {
				if src.Measure_repeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Measure_style", "Measure_repeat", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_style:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Multiple_rest", inst.Multiple_rest, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Multiple_rest](), probe.formStage)
		StageSetAssociationFieldToForm("Measure_repeat", inst.Measure_repeat, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Measure_repeat](), probe.formStage)
		StageSetAssociationFieldToForm("Beat_repeat", inst.Beat_repeat, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Beat_repeat](), probe.formStage)
		StageSetAssociationFieldToForm("Slash", inst.Slash, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Slash](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Measure_style {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Measure_style", refNames, formGroup, probe.formStage)
		}
	case *models.Membrane:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Membrane == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Membrane", refNames, formGroup, probe.formStage)
		}
	case *models.Metal:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Metal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Metal", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Beat_unit", inst.Beat_unit, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beat_unit_dot", inst.Beat_unit_dot, probe.formStage, formGroup)
		AssociationSliceToForm("Beat_unit_tied", inst, &inst.Beat_unit_tied, formGroup, probe)
		StageSetAssociationFieldToForm("Per_minute", inst.Per_minute, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Per_minute](), probe.formStage)
		StageSetBasicFieldtoForm("Metronome_arrows", inst.Metronome_arrows, probe.formStage, formGroup)
		AssociationSliceToForm("Metronome_note", inst, &inst.Metronome_note, formGroup, probe)
		StageSetBasicFieldtoForm("Metronome_relation", inst.Metronome_relation, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Metronome == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Metronome", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome_beam:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronome_notes {
				for _, target := range src.Metronome_beam {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome_note", "Metronome_beam", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome_note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Metronome_type", inst.Metronome_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Metronome_dot", inst.Metronome_dot, probe.formStage, formGroup)
		AssociationSliceToForm("Metronome_beam", inst, &inst.Metronome_beam, formGroup, probe)
		StageSetAssociationFieldToForm("Metronome_tied", inst.Metronome_tied, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Metronome_tied](), probe.formStage)
		StageSetAssociationFieldToForm("Metronome_tuplet", inst.Metronome_tuplet, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Metronome_tuplet](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronomes {
				for _, target := range src.Metronome_note {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome", "Metronome_note", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome_tied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronome_notes {
				if src.Metronome_tied == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome_note", "Metronome_tied", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome_tuplet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actual_notes", inst.Actual_notes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Normal_notes", inst.Normal_notes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Normal_type", inst.Normal_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Normal_dot", inst.Normal_dot, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Show_number", inst.Show_number, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronome_notes {
				if src.Metronome_tuplet == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome_note", "Metronome_tuplet", refNames, formGroup, probe.formStage)
		}
	case *models.Midi_device:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Port", inst.Port, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				for _, target := range src.Midi_device {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Midi_device", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				for _, target := range src.Midi_device {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Midi_device", refNames, formGroup, probe.formStage)
		}
	case *models.Midi_instrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Midi_channel", inst.Midi_channel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Midi_name", inst.Midi_name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Midi_bank", inst.Midi_bank, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Midi_program", inst.Midi_program, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Midi_unpitched", inst.Midi_unpitched, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Volume", inst.Volume, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pan", inst.Pan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Elevation", inst.Elevation, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				for _, target := range src.Midi_instrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Midi_instrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				for _, target := range src.Midi_instrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Midi_instrument", refNames, formGroup, probe.formStage)
		}
	case *models.Miscellaneous:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Miscellaneous_field", inst, &inst.Miscellaneous_field, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Identifications {
				if src.Miscellaneous == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Identification", "Miscellaneous", refNames, formGroup, probe.formStage)
		}
	case *models.Miscellaneous_field:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Miscellaneouss {
				for _, target := range src.Miscellaneous_field {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Miscellaneous", "Miscellaneous_field", refNames, formGroup, probe.formStage)
		}
	case *models.Mordent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Start_note", inst.Start_note, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Trill_step", inst.Trill_step, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Two_note_turn", inst.Two_note_turn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Long", inst.Long, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Approach", inst.Approach, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Departure", inst.Departure, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Mordent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Mordent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Inverted_mordent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Inverted_mordent", refNames, formGroup, probe.formStage)
		}
	case *models.Multiple_rest:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Use_symbols", inst.Use_symbols, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Measure_styles {
				if src.Multiple_rest == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Measure_style", "Multiple_rest", refNames, formGroup, probe.formStage)
		}
	case *models.Name_display:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		AssociationSliceToForm("Display_text", inst, &inst.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", inst, &inst.Accidental_text, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_groups {
				if src.Group_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_abbreviation_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.Part_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Part_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.Part_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Part_abbreviation_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				if src.Part_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				if src.Part_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_abbreviation_display", refNames, formGroup, probe.formStage)
		}
	case *models.Non_arpeggiate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Non_arpeggiate {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Non_arpeggiate", refNames, formGroup, probe.formStage)
		}
	case *models.Notations:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		AssociationSliceToForm("Tied", inst, &inst.Tied, formGroup, probe)
		AssociationSliceToForm("Slur", inst, &inst.Slur, formGroup, probe)
		AssociationSliceToForm("Tuplet", inst, &inst.Tuplet, formGroup, probe)
		AssociationSliceToForm("Glissando", inst, &inst.Glissando, formGroup, probe)
		AssociationSliceToForm("Slide", inst, &inst.Slide, formGroup, probe)
		AssociationSliceToForm("Ornaments", inst, &inst.Ornaments, formGroup, probe)
		AssociationSliceToForm("Technical", inst, &inst.Technical, formGroup, probe)
		AssociationSliceToForm("Articulations", inst, &inst.Articulations, formGroup, probe)
		AssociationSliceToForm("Dynamics", inst, &inst.Dynamics, formGroup, probe)
		AssociationSliceToForm("Fermata", inst, &inst.Fermata, formGroup, probe)
		AssociationSliceToForm("Arpeggiate", inst, &inst.Arpeggiate, formGroup, probe)
		AssociationSliceToForm("Non_arpeggiate", inst, &inst.Non_arpeggiate, formGroup, probe)
		AssociationSliceToForm("Accidental_mark", inst, &inst.Accidental_mark, formGroup, probe)
		AssociationSliceToForm("Other_notation", inst, &inst.Other_notation, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				for _, target := range src.Notations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Notations", refNames, formGroup, probe.formStage)
		}
	case *models.Note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_leger", inst.Print_leger, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dynamics", inst.Dynamics, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_dynamics", inst.End_dynamics, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Attack", inst.Attack, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Release", inst.Release, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_dot", inst.Print_dot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_lyric", inst.Print_lyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Grace", inst.Grace, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Grace](), probe.formStage)
		StageSetBasicFieldtoForm("Chord", inst.Chord, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pitch", inst.Pitch, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Pitch](), probe.formStage)
		StageSetAssociationFieldToForm("Unpitched", inst.Unpitched, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Unpitched](), probe.formStage)
		StageSetAssociationFieldToForm("Rest", inst.Rest, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Rest](), probe.formStage)
		StageSetAssociationFieldToForm("Tie", inst.Tie, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Tie](), probe.formStage)
		StageSetBasicFieldtoForm("Cue", inst.Cue, probe.formStage, formGroup)
		AssociationSliceToForm("Instrument", inst, &inst.Instrument, formGroup, probe)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Note_type](), probe.formStage)
		AssociationSliceToForm("Dot", inst, &inst.Dot, formGroup, probe)
		StageSetAssociationFieldToForm("Accidental", inst.Accidental, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Accidental](), probe.formStage)
		StageSetAssociationFieldToForm("Time_modification", inst.Time_modification, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Time_modification](), probe.formStage)
		StageSetAssociationFieldToForm("Stem", inst.Stem, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Stem](), probe.formStage)
		StageSetAssociationFieldToForm("Notehead", inst.Notehead, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Notehead](), probe.formStage)
		StageSetAssociationFieldToForm("Notehead_text", inst.Notehead_text, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Notehead_text](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Beam", inst.Beam, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Beam](), probe.formStage)
		AssociationSliceToForm("Notations", inst, &inst.Notations, formGroup, probe)
		AssociationSliceToForm("Lyric", inst, &inst.Lyric, formGroup, probe)
		StageSetAssociationFieldToForm("Play", inst.Play, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Play](), probe.formStage)
		StageSetAssociationFieldToForm("Listen", inst.Listen, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Listen](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Note {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Note {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Note", refNames, formGroup, probe.formStage)
		}
	case *models.Note_size:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Appearances {
				for _, target := range src.Note_size {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Appearance", "Note_size", refNames, formGroup, probe.formStage)
		}
	case *models.Note_type:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.Notehead:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Filled", inst.Filled, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Notehead == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Notehead", refNames, formGroup, probe.formStage)
		}
	case *models.Notehead_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("Display_text", inst, &inst.Display_text, formGroup, probe)
		AssociationSliceToForm("Accidental_text", inst, &inst.Accidental_text, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Notehead_text == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Notehead_text", refNames, formGroup, probe.formStage)
		}
	case *models.Numeral:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Numeral_root", inst.Numeral_root, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Numeral_root](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral_alter", inst.Numeral_alter, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral_key", inst.Numeral_key, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Numeral_key](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Numeral == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Numeral", refNames, formGroup, probe.formStage)
		}
	case *models.Numeral_key:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Numeral_fifths", inst.Numeral_fifths, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Numeral_mode", inst.Numeral_mode, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Numerals {
				if src.Numeral_key == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Numeral", "Numeral_key", refNames, formGroup, probe.formStage)
		}
	case *models.Numeral_root:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Numerals {
				if src.Numeral_root == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Numeral", "Numeral_root", refNames, formGroup, probe.formStage)
		}
	case *models.Octave_shift:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Size", inst.Size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Octave_shift == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Octave_shift", refNames, formGroup, probe.formStage)
		}
	case *models.Offset:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sound", inst.Sound, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listenings {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listening", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Offset", refNames, formGroup, probe.formStage)
		}
	case *models.Opus:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Works {
				if src.Opus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Work", "Opus", refNames, formGroup, probe.formStage)
		}
	case *models.Ornaments:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Trill_mark", inst, &inst.Trill_mark, formGroup, probe)
		AssociationSliceToForm("Turn", inst, &inst.Turn, formGroup, probe)
		AssociationSliceToForm("Delayed_turn", inst, &inst.Delayed_turn, formGroup, probe)
		AssociationSliceToForm("Inverted_turn", inst, &inst.Inverted_turn, formGroup, probe)
		AssociationSliceToForm("Delayed_inverted_turn", inst, &inst.Delayed_inverted_turn, formGroup, probe)
		AssociationSliceToForm("Vertical_turn", inst, &inst.Vertical_turn, formGroup, probe)
		AssociationSliceToForm("Inverted_vertical_turn", inst, &inst.Inverted_vertical_turn, formGroup, probe)
		AssociationSliceToForm("Shake", inst, &inst.Shake, formGroup, probe)
		AssociationSliceToForm("Wavy_line", inst, &inst.Wavy_line, formGroup, probe)
		AssociationSliceToForm("Mordent", inst, &inst.Mordent, formGroup, probe)
		AssociationSliceToForm("Inverted_mordent", inst, &inst.Inverted_mordent, formGroup, probe)
		AssociationSliceToForm("Schleifer", inst, &inst.Schleifer, formGroup, probe)
		AssociationSliceToForm("Tremolo", inst, &inst.Tremolo, formGroup, probe)
		AssociationSliceToForm("Haydn", inst, &inst.Haydn, formGroup, probe)
		AssociationSliceToForm("Other_ornament", inst, &inst.Other_ornament, formGroup, probe)
		AssociationSliceToForm("Accidental_mark", inst, &inst.Accidental_mark, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Ornaments {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Ornaments", refNames, formGroup, probe.formStage)
		}
	case *models.Other_appearance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Appearances {
				for _, target := range src.Other_appearance {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Appearance", "Other_appearance", refNames, formGroup, probe.formStage)
		}
	case *models.Other_direction:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Other_direction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Other_direction", refNames, formGroup, probe.formStage)
		}
	case *models.Other_listening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listens {
				for _, target := range src.Other_listen {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listen", "Other_listen", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listenings {
				for _, target := range src.Other_listening {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listening", "Other_listening", refNames, formGroup, probe.formStage)
		}
	case *models.Other_notation:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Other_notation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Other_notation", refNames, formGroup, probe.formStage)
		}
	case *models.Other_placement_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Other_articulation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Other_articulation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Other_ornament {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Other_ornament", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Other_technical {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Other_technical", refNames, formGroup, probe.formStage)
		}
	case *models.Other_play:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Plays {
				for _, target := range src.Other_play {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Play", "Other_play", refNames, formGroup, probe.formStage)
		}
	case *models.Other_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Dynamicss {
				for _, target := range src.Other_dynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Dynamics", "Other_dynamics", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Other_percussion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Other_percussion", refNames, formGroup, probe.formStage)
		}
	case *models.Page_layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page_height", inst.Page_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page_width", inst.Page_width, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Page_margins", inst.Page_margins, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Page_margins](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.Page_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Page_layout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.Page_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Page_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Page_margins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Left_margin", inst.Left_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Right_margin", inst.Right_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Top_margin", inst.Top_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bottom_margin", inst.Bottom_margin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Page_layouts {
				if src.Page_margins == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Page_layout", "Page_margins", refNames, formGroup, probe.formStage)
		}
	case *models.Part_clef:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Clef_octave_change", inst.Clef_octave_change, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.For_parts {
				if src.Part_clef == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.For_part", "Part_clef", refNames, formGroup, probe.formStage)
		}
	case *models.Part_group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Group_name", inst.Group_name, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Group_name](), probe.formStage)
		StageSetAssociationFieldToForm("Group_name_display", inst.Group_name_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Group_abbreviation", inst.Group_abbreviation, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Group_name](), probe.formStage)
		StageSetAssociationFieldToForm("Group_abbreviation_display", inst.Group_abbreviation_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Group_symbol", inst.Group_symbol, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Group_symbol](), probe.formStage)
		StageSetAssociationFieldToForm("Group_barline", inst.Group_barline, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Group_barline](), probe.formStage)
		StageSetBasicFieldtoForm("Group_time", inst.Group_time, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_lists {
				if src.Part_group == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_list", "Part_group", refNames, formGroup, probe.formStage)
		}
	case *models.Part_link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		AssociationSliceToForm("Instrument_link", inst, &inst.Instrument_link, formGroup, probe)
		StageSetBasicFieldtoForm("Group_link", inst.Group_link, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				for _, target := range src.Part_link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_link", refNames, formGroup, probe.formStage)
		}
	case *models.Part_list:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Part_group", inst.Part_group, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_group](), probe.formStage)
		StageSetAssociationFieldToForm("Score_part", inst.Score_part, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Score_part](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				if src.Part_list == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Part_list", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				if src.Part_list == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Part_list", refNames, formGroup, probe.formStage)
		}
	case *models.Part_name:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				if src.Part_name == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_name", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				if src.Part_abbreviation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_abbreviation", refNames, formGroup, probe.formStage)
		}
	case *models.Part_symbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Top_staff", inst.Top_staff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bottom_staff", inst.Bottom_staff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				if src.Part_symbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Part_symbol", refNames, formGroup, probe.formStage)
		}
	case *models.Part_transpose:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Diatonic", inst.Diatonic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Chromatic", inst.Chromatic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Octave_change", inst.Octave_change, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Double", inst.Double, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.For_parts {
				if src.Part_transpose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.For_part", "Part_transpose", refNames, formGroup, probe.formStage)
		}
	case *models.Pedal:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line", inst.Line, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Abbreviated", inst.Abbreviated, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Pedal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Pedal", refNames, formGroup, probe.formStage)
		}
	case *models.Pedal_tuning:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Pedal_step", inst.Pedal_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Pedal_alter", inst.Pedal_alter, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harp_pedalss {
				for _, target := range src.Pedal_tuning {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harp_pedals", "Pedal_tuning", refNames, formGroup, probe.formStage)
		}
	case *models.Per_minute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Metronomes {
				if src.Per_minute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome", "Per_minute", refNames, formGroup, probe.formStage)
		}
	case *models.Percussion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Glass", inst.Glass, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Glass](), probe.formStage)
		StageSetAssociationFieldToForm("Metal", inst.Metal, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Metal](), probe.formStage)
		StageSetAssociationFieldToForm("Wood", inst.Wood, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Wood](), probe.formStage)
		StageSetAssociationFieldToForm("Pitched", inst.Pitched, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Pitched](), probe.formStage)
		StageSetAssociationFieldToForm("Membrane", inst.Membrane, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Membrane](), probe.formStage)
		StageSetAssociationFieldToForm("Effect", inst.Effect, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Effect](), probe.formStage)
		StageSetAssociationFieldToForm("Timpani", inst.Timpani, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Timpani](), probe.formStage)
		StageSetAssociationFieldToForm("Beater", inst.Beater, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Beater](), probe.formStage)
		StageSetAssociationFieldToForm("Stick", inst.Stick, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Stick](), probe.formStage)
		StageSetEnumStringFieldToForm("Stick_location", inst.Stick_location, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Other_percussion", inst.Other_percussion, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Other_text](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Percussion {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Percussion", refNames, formGroup, probe.formStage)
		}
	case *models.Pitch:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Step", inst.Step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Alter", inst.Alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Octave", inst.Octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Pitch == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Pitch", refNames, formGroup, probe.formStage)
		}
	case *models.Pitched:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Pitched == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Pitched", refNames, formGroup, probe.formStage)
		}
	case *models.Placement_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Bends {
				if src.With_bar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bend", "With_bar", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Pluck {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Pluck", refNames, formGroup, probe.formStage)
		}
	case *models.Play:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ipa", inst.Ipa, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Mute", inst.Mute, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Semi_pitched", inst.Semi_pitched, formGroup, probe.formStage)
		AssociationSliceToForm("Other_play", inst, &inst.Other_play, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Play == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Play", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				for _, target := range src.Play {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Play", refNames, formGroup, probe.formStage)
		}
	case *models.Player:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player_name", inst.Player_name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				for _, target := range src.Player {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Player", refNames, formGroup, probe.formStage)
		}
	case *models.Principal_voice:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Principal_voice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Principal_voice", refNames, formGroup, probe.formStage)
		}
	case *models.Print_XSD:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff_spacing", inst.Staff_spacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("New_system", inst.New_system, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("New_page", inst.New_page, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Blank_page", inst.Blank_page, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page_number", inst.Page_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Page_layout", inst.Page_layout, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Page_layout](), probe.formStage)
		StageSetAssociationFieldToForm("System_layout", inst.System_layout, formGroup, probe.stageOfInterest.GetInstancesSet[*models.System_layout](), probe.formStage)
		AssociationSliceToForm("Staff_layout", inst, &inst.Staff_layout, formGroup, probe)
		StageSetAssociationFieldToForm("Measure_layout", inst.Measure_layout, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Measure_layout](), probe.formStage)
		StageSetAssociationFieldToForm("Measure_numbering", inst.Measure_numbering, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Measure_numbering](), probe.formStage)
		StageSetAssociationFieldToForm("Part_name_display", inst.Part_name_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation_display", inst.Part_abbreviation_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Print_XSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Print_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Print_XSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Print_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.Release:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Offset", inst.Offset, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Bends {
				if src.Release == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bend", "Release", refNames, formGroup, probe.formStage)
		}
	case *models.Repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Times", inst.Times, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("After_jump", inst.After_jump, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Winged", inst.Winged, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Repeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Repeat", refNames, formGroup, probe.formStage)
		}
	case *models.Rest:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Measure", inst.Measure, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Display_step", inst.Display_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Display_octave", inst.Display_octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Rest == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Rest", refNames, formGroup, probe.formStage)
		}
	case *models.Root:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Root_step", inst.Root_step, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Root_step](), probe.formStage)
		StageSetAssociationFieldToForm("Root_alter", inst.Root_alter, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Harmony_alter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Root == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Root", refNames, formGroup, probe.formStage)
		}
	case *models.Root_step:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Roots {
				if src.Root_step == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Root", "Root_step", refNames, formGroup, probe.formStage)
		}
	case *models.Scaling:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Millimeters", inst.Millimeters, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tenths", inst.Tenths, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.Scaling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Scaling", refNames, formGroup, probe.formStage)
		}
	case *models.Scordatura:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Accord", inst, &inst.Accord, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Scordatura == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Scordatura", refNames, formGroup, probe.formStage)
		}
	case *models.Score_instrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Instrument_name", inst.Instrument_name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Instrument_abbreviation", inst.Instrument_abbreviation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Instrument_sound", inst.Instrument_sound, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Solo", inst.Solo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ensemble", inst.Ensemble, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Virtual_instrument", inst.Virtual_instrument, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Virtual_instrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_parts {
				for _, target := range src.Score_instrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Score_instrument", refNames, formGroup, probe.formStage)
		}
	case *models.Score_part:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Identification](), probe.formStage)
		AssociationSliceToForm("Part_link", inst, &inst.Part_link, formGroup, probe)
		StageSetAssociationFieldToForm("Part_name", inst.Part_name, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_name](), probe.formStage)
		StageSetAssociationFieldToForm("Part_name_display", inst.Part_name_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation", inst.Part_abbreviation, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_name](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation_display", inst.Part_abbreviation_display, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetBasicFieldtoForm("Group", inst.Group, probe.formStage, formGroup)
		AssociationSliceToForm("Score_instrument", inst, &inst.Score_instrument, formGroup, probe)
		AssociationSliceToForm("Player", inst, &inst.Player, formGroup, probe)
		AssociationSliceToForm("Midi_device", inst, &inst.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", inst, &inst.Midi_instrument, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Part_lists {
				if src.Score_part == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_list", "Score_part", refNames, formGroup, probe.formStage)
		}
	case *models.Score_partwise:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Version", inst.Version, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Work](), probe.formStage)
		StageSetBasicFieldtoForm("Movement_number", inst.Movement_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Movement_title", inst.Movement_title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Defaults](), probe.formStage)
		AssociationSliceToForm("Credit", inst, &inst.Credit, formGroup, probe)
		StageSetAssociationFieldToForm("Part_list", inst.Part_list, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_list](), probe.formStage)
		AssociationSliceToForm("Part", inst, &inst.Part, formGroup, probe)
	case *models.Score_timewise:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Version", inst.Version, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Work](), probe.formStage)
		StageSetBasicFieldtoForm("Movement_number", inst.Movement_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Movement_title", inst.Movement_title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Defaults](), probe.formStage)
		AssociationSliceToForm("Credit", inst, &inst.Credit, formGroup, probe)
		StageSetAssociationFieldToForm("Part_list", inst.Part_list, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Part_list](), probe.formStage)
		AssociationSliceToForm("Measure", inst, &inst.Measure, formGroup, probe)
	case *models.Segno:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Segno_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Segno_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				for _, target := range src.Segno {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Segno", refNames, formGroup, probe.formStage)
		}
	case *models.Slash:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Use_dots", inst.Use_dots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Use_stems", inst.Use_stems, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Slash_type", inst.Slash_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slash_dot", inst.Slash_dot, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Except_voice", inst.Except_voice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Measure_styles {
				if src.Slash == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Measure_style", "Slash", refNames, formGroup, probe.formStage)
		}
	case *models.Slide:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("First_beat", inst.First_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Slide {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Slide", refNames, formGroup, probe.formStage)
		}
	case *models.Slur:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Orientation", inst.Orientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Bezier_x", inst.Bezier_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_y", inst.Bezier_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_x2", inst.Bezier_x2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_y2", inst.Bezier_y2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_offset", inst.Bezier_offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_offset2", inst.Bezier_offset2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Slur {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Slur", refNames, formGroup, probe.formStage)
		}
	case *models.Sound:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tempo", inst.Tempo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dynamics", inst.Dynamics, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dacapo", inst.Dacapo, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Segno", inst.Segno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dalsegno", inst.Dalsegno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Coda", inst.Coda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tocoda", inst.Tocoda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Forward_repeat", inst.Forward_repeat, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Fine", inst.Fine, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Pan", inst.Pan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Elevation", inst.Elevation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Damper_pedal", inst.Damper_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Soft_pedal", inst.Soft_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sostenuto_pedal", inst.Sostenuto_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Instrument_change", inst, &inst.Instrument_change, formGroup, probe)
		AssociationSliceToForm("Midi_device", inst, &inst.Midi_device, formGroup, probe)
		AssociationSliceToForm("Midi_instrument", inst, &inst.Midi_instrument, formGroup, probe)
		AssociationSliceToForm("Play", inst, &inst.Play, formGroup, probe)
		StageSetAssociationFieldToForm("Swing", inst.Swing, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Swing](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_measures {
				for _, target := range src.Sound {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Sound", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_part_1s {
				for _, target := range src.Sound {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Sound", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Directions {
				if src.Sound == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Sound", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_details:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show_frets", inst.Show_frets, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Staff_type", inst.Staff_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Staff_lines", inst.Staff_lines, probe.formStage, formGroup)
		AssociationSliceToForm("Line_detail", inst, &inst.Line_detail, formGroup, probe)
		AssociationSliceToForm("Staff_tuning", inst, &inst.Staff_tuning, formGroup, probe)
		StageSetBasicFieldtoForm("Capo", inst.Capo, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Staff_size", inst.Staff_size, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Staff_size](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Staff_details {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Staff_details", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_divide:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Staff_divide == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Staff_divide", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff_distance", inst.Staff_distance, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				for _, target := range src.Staff_layout {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Staff_layout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				for _, target := range src.Staff_layout {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "Staff_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_size:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Scaling", inst.Scaling, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Staff_detailss {
				if src.Staff_size == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Staff_details", "Staff_size", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_tuning:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tuning_step", inst.Tuning_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Tuning_alter", inst.Tuning_alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tuning_octave", inst.Tuning_octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Staff_detailss {
				for _, target := range src.Staff_tuning {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Staff_details", "Staff_tuning", refNames, formGroup, probe.formStage)
		}
	case *models.Stem:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Stem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Stem", refNames, formGroup, probe.formStage)
		}
	case *models.Stick:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tip", inst.Tip, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Dashed_circle", inst.Dashed_circle, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Stick_type", inst.Stick_type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Stick_material", inst.Stick_material, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Stick == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Stick", refNames, formGroup, probe.formStage)
		}
	case *models.String_mute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.String_mute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "String_mute", refNames, formGroup, probe.formStage)
		}
	case *models.String_type:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Frame_notes {
				if src.String_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "String_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.String_XSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "String_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.Strong_accent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Articulationss {
				for _, target := range src.Strong_accent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Articulations", "Strong_accent", refNames, formGroup, probe.formStage)
		}
	case *models.Style_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Basss {
				if src.Bass_separator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bass", "Bass_separator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Prefix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Prefix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Figure_number == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Figure_number", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Figures {
				if src.Suffix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Suffix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Harmonys {
				if src.Function == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Function", refNames, formGroup, probe.formStage)
		}
	case *models.Supports:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Attribute", inst.Attribute, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Encodings {
				for _, target := range src.Supports {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Encoding", "Supports", refNames, formGroup, probe.formStage)
		}
	case *models.Swing:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Straight", inst.Straight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("First", inst.First, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second", inst.Second, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Swing_type", inst.Swing_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Swing_style", inst.Swing_style, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Sounds {
				if src.Swing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Swing", refNames, formGroup, probe.formStage)
		}
	case *models.Sync:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Latency", inst.Latency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listenings {
				for _, target := range src.Sync {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listening", "Sync", refNames, formGroup, probe.formStage)
		}
	case *models.System_dividers:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Left_divider", inst.Left_divider, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_object_style_align](), probe.formStage)
		StageSetAssociationFieldToForm("Right_divider", inst.Right_divider, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Empty_print_object_style_align](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.System_layouts {
				if src.System_dividers == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_layout", "System_dividers", refNames, formGroup, probe.formStage)
		}
	case *models.System_layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("System_margins", inst.System_margins, formGroup, probe.stageOfInterest.GetInstancesSet[*models.System_margins](), probe.formStage)
		StageSetBasicFieldtoForm("System_distance", inst.System_distance, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Top_system_distance", inst.Top_system_distance, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("System_dividers", inst.System_dividers, formGroup, probe.stageOfInterest.GetInstancesSet[*models.System_dividers](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Defaultss {
				if src.System_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "System_layout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Print_XSDs {
				if src.System_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print_XSD", "System_layout", refNames, formGroup, probe.formStage)
		}
	case *models.System_margins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Left_margin", inst.Left_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Right_margin", inst.Right_margin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.System_layouts {
				if src.System_margins == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_layout", "System_margins", refNames, formGroup, probe.formStage)
		}
	case *models.Tap:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Hand", inst.Hand, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Technicals {
				for _, target := range src.Tap {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "Tap", refNames, formGroup, probe.formStage)
		}
	case *models.Technical:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		AssociationSliceToForm("Up_bow", inst, &inst.Up_bow, formGroup, probe)
		AssociationSliceToForm("Down_bow", inst, &inst.Down_bow, formGroup, probe)
		AssociationSliceToForm("Harmonic", inst, &inst.Harmonic, formGroup, probe)
		AssociationSliceToForm("Open_string", inst, &inst.Open_string, formGroup, probe)
		AssociationSliceToForm("Thumb_position", inst, &inst.Thumb_position, formGroup, probe)
		AssociationSliceToForm("Fingering", inst, &inst.Fingering, formGroup, probe)
		AssociationSliceToForm("Pluck", inst, &inst.Pluck, formGroup, probe)
		AssociationSliceToForm("Double_tongue", inst, &inst.Double_tongue, formGroup, probe)
		AssociationSliceToForm("Triple_tongue", inst, &inst.Triple_tongue, formGroup, probe)
		AssociationSliceToForm("Stopped", inst, &inst.Stopped, formGroup, probe)
		AssociationSliceToForm("Snap_pizzicato", inst, &inst.Snap_pizzicato, formGroup, probe)
		AssociationSliceToForm("Fret", inst, &inst.Fret, formGroup, probe)
		AssociationSliceToForm("String_XSD", inst, &inst.String_XSD, formGroup, probe)
		AssociationSliceToForm("Hammer_on", inst, &inst.Hammer_on, formGroup, probe)
		AssociationSliceToForm("Pull_off", inst, &inst.Pull_off, formGroup, probe)
		AssociationSliceToForm("Bend", inst, &inst.Bend, formGroup, probe)
		AssociationSliceToForm("Tap", inst, &inst.Tap, formGroup, probe)
		AssociationSliceToForm("Heel", inst, &inst.Heel, formGroup, probe)
		AssociationSliceToForm("Toe", inst, &inst.Toe, formGroup, probe)
		AssociationSliceToForm("Fingernails", inst, &inst.Fingernails, formGroup, probe)
		AssociationSliceToForm("Hole", inst, &inst.Hole, formGroup, probe)
		AssociationSliceToForm("Arrow", inst, &inst.Arrow, formGroup, probe)
		AssociationSliceToForm("Handbell", inst, &inst.Handbell, formGroup, probe)
		AssociationSliceToForm("Brass_bend", inst, &inst.Brass_bend, formGroup, probe)
		AssociationSliceToForm("Flip", inst, &inst.Flip, formGroup, probe)
		AssociationSliceToForm("Smear", inst, &inst.Smear, formGroup, probe)
		AssociationSliceToForm("Open", inst, &inst.Open, formGroup, probe)
		AssociationSliceToForm("Half_muted", inst, &inst.Half_muted, formGroup, probe)
		AssociationSliceToForm("Harmon_mute", inst, &inst.Harmon_mute, formGroup, probe)
		AssociationSliceToForm("Golpe", inst, &inst.Golpe, formGroup, probe)
		AssociationSliceToForm("Other_technical", inst, &inst.Other_technical, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Technical {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Technical", refNames, formGroup, probe.formStage)
		}
	case *models.Text_element_data:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Lyrics {
				if src.Text == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Text", refNames, formGroup, probe.formStage)
		}
	case *models.Tie:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Tie == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Tie", refNames, formGroup, probe.formStage)
		}
	case *models.Tied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Orientation", inst.Orientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Bezier_x", inst.Bezier_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_y", inst.Bezier_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_x2", inst.Bezier_x2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_y2", inst.Bezier_y2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_offset", inst.Bezier_offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bezier_offset2", inst.Bezier_offset2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Tied {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Tied", refNames, formGroup, probe.formStage)
		}
	case *models.Time:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Separator", inst.Separator, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beat_type", inst.Beat_type, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Interchangeable", inst.Interchangeable, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Interchangeable](), probe.formStage)
		StageSetBasicFieldtoForm("Senza_misura", inst.Senza_misura, probe.formStage, formGroup)
	case *models.Time_modification:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actual_notes", inst.Actual_notes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Normal_notes", inst.Normal_notes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Normal_type", inst.Normal_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Normal_dot", inst.Normal_dot, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Time_modification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Time_modification", refNames, formGroup, probe.formStage)
		}
	case *models.Timpani:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Timpani == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Timpani", refNames, formGroup, probe.formStage)
		}
	case *models.Transpose:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Diatonic", inst.Diatonic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Chromatic", inst.Chromatic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Octave_change", inst.Octave_change, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Double", inst.Double, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Attributess {
				for _, target := range src.Transpose {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Transpose", refNames, formGroup, probe.formStage)
		}
	case *models.Tremolo:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Tremolo {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Tremolo", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Show_number", inst.Show_number, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Show_type", inst.Show_type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Line_shape", inst.Line_shape, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Tuplet_actual", inst.Tuplet_actual, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_portion](), probe.formStage)
		StageSetAssociationFieldToForm("Tuplet_normal", inst.Tuplet_normal, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_portion](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notationss {
				for _, target := range src.Tuplet {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Tuplet", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet_dot:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Tuplet_portions {
				for _, target := range src.Tuplet_dot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet_portion", "Tuplet_dot", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet_number:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Tuplet_portions {
				if src.Tuplet_number == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet_portion", "Tuplet_number", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet_portion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Tuplet_number", inst.Tuplet_number, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_number](), probe.formStage)
		StageSetAssociationFieldToForm("Tuplet_type", inst.Tuplet_type, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Tuplet_type](), probe.formStage)
		AssociationSliceToForm("Tuplet_dot", inst, &inst.Tuplet_dot, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Tuplets {
				if src.Tuplet_actual == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet", "Tuplet_actual", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Tuplets {
				if src.Tuplet_normal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet", "Tuplet_normal", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet_type:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_style", inst.Font_style, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Font_weight", inst.Font_weight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Tuplet_portions {
				if src.Tuplet_type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet_portion", "Tuplet_type", refNames, formGroup, probe.formStage)
		}
	case *models.Typed_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Encodings {
				for _, target := range src.Encoder {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Encoding", "Encoder", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Identifications {
				for _, target := range src.Creator {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Identification", "Creator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Identifications {
				for _, target := range src.Rights {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Identification", "Rights", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Identifications {
				for _, target := range src.Relation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Identification", "Relation", refNames, formGroup, probe.formStage)
		}
	case *models.Unpitched:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Display_step", inst.Display_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Display_octave", inst.Display_octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Notes {
				if src.Unpitched == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Unpitched", refNames, formGroup, probe.formStage)
		}
	case *models.Virtual_instrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Virtual_library", inst.Virtual_library, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Virtual_name", inst.Virtual_name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Instrument_changes {
				if src.Virtual_instrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Instrument_change", "Virtual_instrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_instruments {
				if src.Virtual_instrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_instrument", "Virtual_instrument", refNames, formGroup, probe.formStage)
		}
	case *models.Wait:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Listens {
				for _, target := range src.Wait {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listen", "Wait", refNames, formGroup, probe.formStage)
		}
	case *models.Wavy_line:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Start_note", inst.Start_note, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Trill_step", inst.Trill_step, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Two_note_turn", inst.Two_note_turn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Barlines {
				if src.Wavy_line == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Wavy_line", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Ornamentss {
				for _, target := range src.Wavy_line {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ornaments", "Wavy_line", refNames, formGroup, probe.formStage)
		}
	case *models.Wedge:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Spread", inst.Spread, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Niente", inst.Niente, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Line_type", inst.Line_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Direction_types {
				if src.Wedge == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Wedge", refNames, formGroup, probe.formStage)
		}
	case *models.Wood:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Percussions {
				if src.Wood == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Wood", refNames, formGroup, probe.formStage)
		}
	case *models.Work:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Work_number", inst.Work_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Work_title", inst.Work_title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Opus", inst.Opus, formGroup, probe.stageOfInterest.GetInstancesSet[*models.Opus](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_partwises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Work", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.Score_timewises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Work", refNames, formGroup, probe.formStage)
		}
	case *xlink.ArcType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("From", inst.From, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("To", inst.To, probe.formStage, formGroup)

		{
			// Slice of pointers: Title_1
			div := (&form.FormDiv{Name: "Title_1"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Title_1 {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Title_1",
				Label: "Title_1",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.Extendeds {
				if src.Arc == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Arc", refNames, formGroup, probe.formStage)
		}
	case *xlink.Extended:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Title_1", inst.Title_1, formGroup, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.TitleEltType](), probe.formStage)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.ResourceType](), probe.formStage)
		StageSetAssociationFieldToForm("Locator", inst.Locator, formGroup, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.LocatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Arc", inst.Arc, formGroup, probe.stageOfInterest.XlinkStage.GetInstancesSet[*xlink.ArcType](), probe.formStage)
	case *xlink.LocatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)

		{
			// Slice of pointers: Title_1
			div := (&form.FormDiv{Name: "Title_1"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Title_1 {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Title_1",
				Label: "Title_1",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.Extendeds {
				if src.Locator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Locator", refNames, formGroup, probe.formStage)
		}
	case *xlink.ResourceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.Extendeds {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Resource", refNames, formGroup, probe.formStage)
		}
	case *xlink.Simple:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
	case *xlink.TitleEltType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.ArcTypes {
				for _, target := range src.Title_1 {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.ArcType", "Title_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.Extendeds {
				if src.Title_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Title_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.XlinkStage.LocatorTypes {
				for _, target := range src.Title_1 {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.LocatorType", "Title_1", refNames, formGroup, probe.formStage)
		}
	case *xmlxsd.Xmlxsd_SchemaElement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
