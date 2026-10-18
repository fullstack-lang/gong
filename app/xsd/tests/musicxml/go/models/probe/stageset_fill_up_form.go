// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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

		{
			// Slice of pointers: Note
			div := (&form.FormDiv{Name: "Note"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note",
				Label: "Note",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Backup
			div := (&form.FormDiv{Name: "Backup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Backup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Backup",
				Label: "Backup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Forward
			div := (&form.FormDiv{Name: "Forward"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Forward {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Forward",
				Label: "Forward",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Direction
			div := (&form.FormDiv{Name: "Direction"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Direction {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Direction",
				Label: "Direction",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmony
			div := (&form.FormDiv{Name: "Harmony"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmony {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmony",
				Label: "Harmony",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Figured_bass
			div := (&form.FormDiv{Name: "Figured_bass"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Figured_bass {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Figured_bass",
				Label: "Figured_bass",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Print
			div := (&form.FormDiv{Name: "Print"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Print {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Print",
				Label: "Print",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sound
			div := (&form.FormDiv{Name: "Sound"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sound {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sound",
				Label: "Sound",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Listening
			div := (&form.FormDiv{Name: "Listening"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Listening {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Listening",
				Label: "Listening",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Barline
			div := (&form.FormDiv{Name: "Barline"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Barline {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Barline",
				Label: "Barline",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Grouping
			div := (&form.FormDiv{Name: "Grouping"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Grouping {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Grouping",
				Label: "Grouping",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bookmark
			div := (&form.FormDiv{Name: "Bookmark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bookmark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bookmark",
				Label: "Bookmark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_parts {
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

		{
			// Slice of pointers: Part
			div := (&form.FormDiv{Name: "Part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part",
				Label: "Part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_timewises {
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

		{
			// Slice of pointers: Measure
			div := (&form.FormDiv{Name: "Measure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Measure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Measure",
				Label: "Measure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
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

		{
			// Slice of pointers: Note
			div := (&form.FormDiv{Name: "Note"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note",
				Label: "Note",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Backup
			div := (&form.FormDiv{Name: "Backup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Backup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Backup",
				Label: "Backup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Forward
			div := (&form.FormDiv{Name: "Forward"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Forward {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Forward",
				Label: "Forward",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Direction
			div := (&form.FormDiv{Name: "Direction"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Direction {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Direction",
				Label: "Direction",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmony
			div := (&form.FormDiv{Name: "Harmony"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmony {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmony",
				Label: "Harmony",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Figured_bass
			div := (&form.FormDiv{Name: "Figured_bass"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Figured_bass {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Figured_bass",
				Label: "Figured_bass",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Print
			div := (&form.FormDiv{Name: "Print"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Print {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Print",
				Label: "Print",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sound
			div := (&form.FormDiv{Name: "Sound"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sound {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sound",
				Label: "Sound",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Listening
			div := (&form.FormDiv{Name: "Listening"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Listening {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Listening",
				Label: "Listening",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Barline
			div := (&form.FormDiv{Name: "Barline"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Barline {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Barline",
				Label: "Barline",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Grouping
			div := (&form.FormDiv{Name: "Grouping"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Grouping {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Grouping",
				Label: "Grouping",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bookmark
			div := (&form.FormDiv{Name: "Bookmark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bookmark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bookmark",
				Label: "Bookmark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measure_1s {
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
		StageSetBasicFieldtoForm("Cautionary", inst.Cautionary, probe.formStage, formGroup)
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Name_displays {
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
			for src := range probe.stageSet.Stage.Notehead_texts {
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
		StageSetBasicFieldtoForm("String", inst.String, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tuning_step", inst.Tuning_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Tuning_alter", inst.Tuning_alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tuning_octave", inst.Tuning_octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scordaturas {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_high", inst.Accordion_high, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_middle", inst.Accordion_middle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Accordion_low", inst.Accordion_low, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Accordion_registration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Accordion_registration", refNames, formGroup, probe.formStage)
		}
	case *models.Appearance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Line_width
			div := (&form.FormDiv{Name: "Line_width"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Line_width {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Line_width",
				Label: "Line_width",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Note_size
			div := (&form.FormDiv{Name: "Note_size"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note_size {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note_size",
				Label: "Note_size",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Distance
			div := (&form.FormDiv{Name: "Distance"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Distance {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Distance",
				Label: "Distance",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Glyph
			div := (&form.FormDiv{Name: "Glyph"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Glyph {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Glyph",
				Label: "Glyph",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_appearance
			div := (&form.FormDiv{Name: "Other_appearance"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_appearance {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_appearance",
				Label: "Other_appearance",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
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
		StageSetBasicFieldtoForm("Direction", inst.Direction, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Unbroken", inst.Unbroken, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arrow_direction", inst.Arrow_direction, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arrow_style", inst.Arrow_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arrowhead", inst.Arrowhead, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Circular_arrow", inst.Circular_arrow, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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

		{
			// Slice of pointers: Accent
			div := (&form.FormDiv{Name: "Accent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accent",
				Label: "Accent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Strong_accent
			div := (&form.FormDiv{Name: "Strong_accent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Strong_accent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Strong_accent",
				Label: "Strong_accent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staccato
			div := (&form.FormDiv{Name: "Staccato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staccato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staccato",
				Label: "Staccato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tenuto
			div := (&form.FormDiv{Name: "Tenuto"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tenuto {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tenuto",
				Label: "Tenuto",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Detached_legato
			div := (&form.FormDiv{Name: "Detached_legato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Detached_legato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Detached_legato",
				Label: "Detached_legato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staccatissimo
			div := (&form.FormDiv{Name: "Staccatissimo"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staccatissimo {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staccatissimo",
				Label: "Staccatissimo",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Spiccato
			div := (&form.FormDiv{Name: "Spiccato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Spiccato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Spiccato",
				Label: "Spiccato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Scoop
			div := (&form.FormDiv{Name: "Scoop"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Scoop {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Scoop",
				Label: "Scoop",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Plop
			div := (&form.FormDiv{Name: "Plop"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Plop {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Plop",
				Label: "Plop",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Doit
			div := (&form.FormDiv{Name: "Doit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Doit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Doit",
				Label: "Doit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Falloff
			div := (&form.FormDiv{Name: "Falloff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Falloff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Falloff",
				Label: "Falloff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Breath_mark
			div := (&form.FormDiv{Name: "Breath_mark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Breath_mark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Breath_mark",
				Label: "Breath_mark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Caesura
			div := (&form.FormDiv{Name: "Caesura"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Caesura {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Caesura",
				Label: "Caesura",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Stress
			div := (&form.FormDiv{Name: "Stress"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stress {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stress",
				Label: "Stress",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Unstress
			div := (&form.FormDiv{Name: "Unstress"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Unstress {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Unstress",
				Label: "Unstress",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Soft_accent
			div := (&form.FormDiv{Name: "Soft_accent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Soft_accent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Soft_accent",
				Label: "Soft_accent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_articulation
			div := (&form.FormDiv{Name: "Other_articulation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_articulation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_articulation",
				Label: "Other_articulation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_only", inst.Time_only, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Listens {
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
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)

		{
			// Slice of pointers: Key
			div := (&form.FormDiv{Name: "Key"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Key {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Key",
				Label: "Key",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Time
			div := (&form.FormDiv{Name: "Time"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Time {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Time",
				Label: "Time",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Staves", inst.Staves, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Part_symbol", inst.Part_symbol, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_symbol](), probe.formStage)
		StageSetBasicFieldtoForm("Instruments", inst.Instruments, probe.formStage, formGroup)

		{
			// Slice of pointers: Clef
			div := (&form.FormDiv{Name: "Clef"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Clef {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Clef",
				Label: "Clef",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staff_details
			div := (&form.FormDiv{Name: "Staff_details"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staff_details {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staff_details",
				Label: "Staff_details",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Transpose
			div := (&form.FormDiv{Name: "Transpose"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Transpose {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Transpose",
				Label: "Transpose",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: For_part
			div := (&form.FormDiv{Name: "For_part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.For_part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "For_part",
				Label: "For_part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Directive
			div := (&form.FormDiv{Name: "Directive"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Directive {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Directive",
				Label: "Directive",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Measure_style
			div := (&form.FormDiv{Name: "Measure_style"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Measure_style {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Measure_style",
				Label: "Measure_style",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Bar_style == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Bar_style", refNames, formGroup, probe.formStage)
		}
	case *models.Barline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Location", inst.Location, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Segno", inst.Segno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Coda", inst.Coda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bar_style", inst.Bar_style, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bar_style_color](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetAssociationFieldToForm("Wavy_line", inst.Wavy_line, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Wavy_line](), probe.formStage)
		StageSetAssociationFieldToForm("Segno_1", inst.Segno_1, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Segno](), probe.formStage)
		StageSetAssociationFieldToForm("Coda_1", inst.Coda_1, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Coda](), probe.formStage)
		StageSetAssociationFieldToForm("Fermata", inst.Fermata, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Fermata](), probe.formStage)
		StageSetAssociationFieldToForm("Ending", inst.Ending, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Ending](), probe.formStage)
		StageSetAssociationFieldToForm("Repeat", inst.Repeat, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Repeat](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frame_notes {
				if src.Barre == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Barre", refNames, formGroup, probe.formStage)
		}
	case *models.Bass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arrangement", inst.Arrangement, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bass_separator", inst.Bass_separator, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Bass_step", inst.Bass_step, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bass_step](), probe.formStage)
		StageSetAssociationFieldToForm("Bass_alter", inst.Bass_alter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Harmony_alter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Basss {
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
		StageSetBasicFieldtoForm("Fan", inst.Fan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Beam == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Beam", refNames, formGroup, probe.formStage)
		}
	case *models.Beat_repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Use_dots", inst.Use_dots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Slash_type", inst.Slash_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slash_dot", inst.Slash_dot, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Except_voice", inst.Except_voice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Measure_styles {
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
			for src := range probe.stageSet.Stage.Metronomes {
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
		StageSetBasicFieldtoForm("Tip", inst.Tip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
				if src.Beater == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Beater", refNames, formGroup, probe.formStage)
		}
	case *models.Bend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Shape", inst.Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("First_beat", inst.First_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bend_alter", inst.Bend_alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pre_bend", inst.Pre_bend, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Release", inst.Release, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Release](), probe.formStage)
		StageSetAssociationFieldToForm("With_bar", inst.With_bar, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Placement_text](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
			for src := range probe.stageSet.Stage.Credits {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_end", inst.Line_end, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_length", inst.End_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
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
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
		StageSetBasicFieldtoForm("Location", inst.Location, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Keys {
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
		StageSetBasicFieldtoForm("Size", inst.Size, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("After_barline", inst.After_barline, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sign", inst.Sign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Clef_octave_change", inst.Clef_octave_change, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Coda_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Coda_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetAssociationFieldToForm("Credit_image", inst.Credit_image, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Image](), probe.formStage)

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bookmark
			div := (&form.FormDiv{Name: "Bookmark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bookmark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bookmark",
				Label: "Bookmark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Credit_words
			div := (&form.FormDiv{Name: "Credit_words"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit_words {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit_words",
				Label: "Credit_words",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Credit_symbol
			div := (&form.FormDiv{Name: "Credit_symbol"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit_symbol {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit_symbol",
				Label: "Credit_symbol",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
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
			for src := range probe.stageSet.Stage.Score_timewises {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
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
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Dashes == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Dashes", refNames, formGroup, probe.formStage)
		}
	case *models.Defaults:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Scaling", inst.Scaling, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Scaling](), probe.formStage)
		StageSetBasicFieldtoForm("Concert_score", inst.Concert_score, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Page_layout", inst.Page_layout, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Page_layout](), probe.formStage)
		StageSetAssociationFieldToForm("System_layout", inst.System_layout, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.System_layout](), probe.formStage)

		{
			// Slice of pointers: Staff_layout
			div := (&form.FormDiv{Name: "Staff_layout"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staff_layout {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staff_layout",
				Label: "Staff_layout",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Appearance", inst.Appearance, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Appearance](), probe.formStage)
		StageSetAssociationFieldToForm("Music_font", inst.Music_font, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_font](), probe.formStage)
		StageSetAssociationFieldToForm("Word_font", inst.Word_font, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_font](), probe.formStage)

		{
			// Slice of pointers: Lyric_font
			div := (&form.FormDiv{Name: "Lyric_font"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lyric_font {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lyric_font",
				Label: "Lyric_font",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Lyric_language
			div := (&form.FormDiv{Name: "Lyric_language"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lyric_language {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lyric_language",
				Label: "Lyric_language",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
				if src.Defaults == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Defaults", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_timewises {
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
		StageSetAssociationFieldToForm("Degree_value", inst.Degree_value, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Degree_value](), probe.formStage)
		StageSetAssociationFieldToForm("Degree_alter", inst.Degree_alter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Degree_alter](), probe.formStage)
		StageSetAssociationFieldToForm("Degree_type", inst.Degree_type, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Degree_type](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Degrees {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Degrees {
				if src.Degree_type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Degree", "Degree_type", refNames, formGroup, probe.formStage)
		}
	case *models.Degree_value:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Symbol", inst.Symbol, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Degrees {
				if src.Degree_value == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Degree", "Degree_value", refNames, formGroup, probe.formStage)
		}
	case *models.Direction:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Directive", inst.Directive, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Direction_type
			div := (&form.FormDiv{Name: "Direction_type"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Direction_type {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Direction_type",
				Label: "Direction_type",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Sound", inst.Sound, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Sound](), probe.formStage)
		StageSetAssociationFieldToForm("Listening", inst.Listening, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Listening](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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

		{
			// Slice of pointers: Rehearsal
			div := (&form.FormDiv{Name: "Rehearsal"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rehearsal {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rehearsal",
				Label: "Rehearsal",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Segno
			div := (&form.FormDiv{Name: "Segno"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Segno {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Segno",
				Label: "Segno",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Coda
			div := (&form.FormDiv{Name: "Coda"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Coda {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Coda",
				Label: "Coda",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Words
			div := (&form.FormDiv{Name: "Words"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Words {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Words",
				Label: "Words",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Symbol
			div := (&form.FormDiv{Name: "Symbol"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Symbol {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Symbol",
				Label: "Symbol",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Wedge", inst.Wedge, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Wedge](), probe.formStage)

		{
			// Slice of pointers: Dynamics
			div := (&form.FormDiv{Name: "Dynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dynamics",
				Label: "Dynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Dashes", inst.Dashes, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Dashes](), probe.formStage)
		StageSetAssociationFieldToForm("Bracket", inst.Bracket, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bracket](), probe.formStage)
		StageSetAssociationFieldToForm("Pedal", inst.Pedal, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Pedal](), probe.formStage)
		StageSetAssociationFieldToForm("Metronome", inst.Metronome, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Metronome](), probe.formStage)
		StageSetAssociationFieldToForm("Octave_shift", inst.Octave_shift, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Octave_shift](), probe.formStage)
		StageSetAssociationFieldToForm("Harp_pedals", inst.Harp_pedals, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Harp_pedals](), probe.formStage)
		StageSetAssociationFieldToForm("Damp", inst.Damp, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("Damp_all", inst.Damp_all, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("Eyeglasses", inst.Eyeglasses, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_print_style_align_id](), probe.formStage)
		StageSetAssociationFieldToForm("String_mute", inst.String_mute, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.String_mute](), probe.formStage)
		StageSetAssociationFieldToForm("Scordatura", inst.Scordatura, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Scordatura](), probe.formStage)
		StageSetAssociationFieldToForm("Image", inst.Image, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Image](), probe.formStage)
		StageSetAssociationFieldToForm("Principal_voice", inst.Principal_voice, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Principal_voice](), probe.formStage)

		{
			// Slice of pointers: Percussion
			div := (&form.FormDiv{Name: "Percussion"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Percussion {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Percussion",
				Label: "Percussion",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Accordion_registration", inst.Accordion_registration, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Accordion_registration](), probe.formStage)
		StageSetAssociationFieldToForm("Staff_divide", inst.Staff_divide, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Staff_divide](), probe.formStage)
		StageSetAssociationFieldToForm("Other_direction", inst.Other_direction, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Other_direction](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Directions {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Appearances {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
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

		{
			// Slice of pointers: Other_dynamics
			div := (&form.FormDiv{Name: "Other_dynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_dynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_dynamics",
				Label: "Other_dynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lyrics {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
				if src.Music_font == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Music_font", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
				if src.Word_font == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Word_font", refNames, formGroup, probe.formStage)
		}
	case *models.Empty_line:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_shape", inst.Line_shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_length", inst.Line_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Notes {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.System_dividerss {
				if src.Left_divider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_dividers", "Left_divider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.System_dividerss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
	case *models.Empty_print_style_align:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
	case *models.Empty_print_style_align_id:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Damp == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Damp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Damp_all == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Damp_all", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start_note", inst.Start_note, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Trill_step", inst.Trill_step, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Two_note_turn", inst.Two_note_turn, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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

		{
			// Slice of pointers: Encoder
			div := (&form.FormDiv{Name: "Encoder"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Encoder {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Encoder",
				Label: "Encoder",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Software", inst.Software, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Encoding_description", inst.Encoding_description, probe.formStage, formGroup)

		{
			// Slice of pointers: Supports
			div := (&form.FormDiv{Name: "Supports"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Supports {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Supports",
				Label: "Supports",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Identifications {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_length", inst.End_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text_x", inst.Text_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text_y", inst.Text_y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Ending == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Ending", refNames, formGroup, probe.formStage)
		}
	case *models.Extend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Extend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lyrics {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Extend", refNames, formGroup, probe.formStage)
		}
	case *models.Feature:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groupings {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Fermata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Fermata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetAssociationFieldToForm("Prefix", inst.Prefix, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Figure_number", inst.Figure_number, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Suffix", inst.Suffix, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Extend](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figured_basss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_dot", inst.Print_dot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_lyric", inst.Print_lyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Figure
			div := (&form.FormDiv{Name: "Figure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Figure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Figure",
				Label: "Figure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frame_notes {
				if src.Fingering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Fingering", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Location", inst.Location, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frames {
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
		StageSetAssociationFieldToForm("Part_clef", inst.Part_clef, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_clef](), probe.formStage)
		StageSetAssociationFieldToForm("Part_transpose", inst.Part_transpose, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_transpose](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)
	case *models.Formatted_symbol_id:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Credits {
				for _, target := range src.Credit_symbol {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_symbol", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Backups {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Backup", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Directions {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figured_basss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figured_bass", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Forwards {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Forward", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lyrics {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Name_displays {
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
			for src := range probe.stageSet.Stage.Notationss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notehead_texts {
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
			for src := range probe.stageSet.Stage.Part_groups {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_height", inst.Line_height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Credits {
				for _, target := range src.Credit_words {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_words", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame_strings", inst.Frame_strings, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame_frets", inst.Frame_frets, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("First_fret", inst.First_fret, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.First_fret](), probe.formStage)

		{
			// Slice of pointers: Frame_note
			div := (&form.FormDiv{Name: "Frame_note"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Frame_note {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Frame_note",
				Label: "Frame_note",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Frame == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Frame", refNames, formGroup, probe.formStage)
		}
	case *models.Frame_note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("String", inst.String, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.String_type](), probe.formStage)
		StageSetAssociationFieldToForm("Fret", inst.Fret, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Fret](), probe.formStage)
		StageSetAssociationFieldToForm("Fingering", inst.Fingering, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Fingering](), probe.formStage)
		StageSetAssociationFieldToForm("Barre", inst.Barre, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Barre](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frames {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frame_notes {
				if src.Fret == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "Fret", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
				if src.Glass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Glass", refNames, formGroup, probe.formStage)
		}
	case *models.Glissando:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Appearances {
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
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
				if src.Group_name == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_name", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
				if src.Group_symbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_symbol", refNames, formGroup, probe.formStage)
		}
	case *models.Grouping:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Member_of", inst.Member_of, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Feature
			div := (&form.FormDiv{Name: "Feature"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Feature {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Feature",
				Label: "Feature",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Location", inst.Location, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmon_mutes {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Harmon_closed", inst.Harmon_closed, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Harmon_closed](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Natural", inst.Natural, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Artificial", inst.Artificial, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base_pitch", inst.Base_pitch, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Touching_pitch", inst.Touching_pitch, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sounding_pitch", inst.Sounding_pitch, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_frame", inst.Print_frame, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Arrangement", inst.Arrangement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Root", inst.Root, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Root](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral", inst.Numeral, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Numeral](), probe.formStage)
		StageSetAssociationFieldToForm("Function", inst.Function, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Style_text](), probe.formStage)
		StageSetAssociationFieldToForm("Kind", inst.Kind, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Kind](), probe.formStage)
		StageSetAssociationFieldToForm("Inversion", inst.Inversion, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Inversion](), probe.formStage)
		StageSetAssociationFieldToForm("Bass", inst.Bass, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bass](), probe.formStage)

		{
			// Slice of pointers: Degree
			div := (&form.FormDiv{Name: "Degree"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Degree {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Degree",
				Label: "Degree",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Frame", inst.Frame, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Frame](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Basss {
				if src.Bass_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bass", "Bass_alter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Numerals {
				if src.Numeral_alter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Numeral", "Numeral_alter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Roots {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Pedal_tuning
			div := (&form.FormDiv{Name: "Pedal_tuning"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pedal_tuning {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pedal_tuning",
				Label: "Pedal_tuning",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Harp_pedals == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Harp_pedals", refNames, formGroup, probe.formStage)
		}
	case *models.Heel_toe:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Hole_type", inst.Hole_type, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Hole_closed", inst.Hole_closed, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Hole_closed](), probe.formStage)
		StageSetBasicFieldtoForm("Hole_shape", inst.Hole_shape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Location", inst.Location, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Holes {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start_note", inst.Start_note, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Trill_step", inst.Trill_step, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Two_note_turn", inst.Two_note_turn, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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

		{
			// Slice of pointers: Creator
			div := (&form.FormDiv{Name: "Creator"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Creator {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Creator",
				Label: "Creator",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Rights
			div := (&form.FormDiv{Name: "Rights"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rights {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rights",
				Label: "Rights",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Encoding", inst.Encoding, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Encoding](), probe.formStage)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)

		{
			// Slice of pointers: Relation
			div := (&form.FormDiv{Name: "Relation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Relation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Relation",
				Label: "Relation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Miscellaneous", inst.Miscellaneous, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Miscellaneous](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_timewises {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Credits {
				if src.Credit_image == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Credit_image", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetAssociationFieldToForm("Virtual_instrument", inst.Virtual_instrument, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Virtual_instrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sounds {
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
			for src := range probe.stageSet.Stage.Part_links {
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
		StageSetBasicFieldtoForm("Symbol", inst.Symbol, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Separator", inst.Separator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Time_relation", inst.Time_relation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beat_type", inst.Beat_type, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Times {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Cancel", inst.Cancel, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Cancel](), probe.formStage)
		StageSetBasicFieldtoForm("Fifths", inst.Fifths, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Mode", inst.Mode, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Key_step", inst.Key_step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Key_alter", inst.Key_alter, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Key_accidental", inst.Key_accidental, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Key_accidental](), probe.formStage)

		{
			// Slice of pointers: Key_octave
			div := (&form.FormDiv{Name: "Key_octave"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Key_octave {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Key_octave",
				Label: "Key_octave",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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
			for src := range probe.stageSet.Stage.Keys {
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
			for src := range probe.stageSet.Stage.Keys {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Kind == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Kind", refNames, formGroup, probe.formStage)
		}
	case *models.Level:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Reference", inst.Reference, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Backups {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Backup", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Directions {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figured_basss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figured_bass", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Forwards {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Forward", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lyrics {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Notations", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
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
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Staff_detailss {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Appearances {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
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
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
			for src := range probe.stageSet.Stage.Credits {
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

		{
			// Slice of pointers: Assess
			div := (&form.FormDiv{Name: "Assess"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Assess {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Assess",
				Label: "Assess",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Wait
			div := (&form.FormDiv{Name: "Wait"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Wait {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Wait",
				Label: "Wait",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_listen
			div := (&form.FormDiv{Name: "Other_listen"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_listen {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_listen",
				Label: "Other_listen",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Listen == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Listen", refNames, formGroup, probe.formStage)
		}
	case *models.Listening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Sync
			div := (&form.FormDiv{Name: "Sync"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sync {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sync",
				Label: "Sync",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_listening
			div := (&form.FormDiv{Name: "Other_listening"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_listening {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_listening",
				Label: "Other_listening",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
			for src := range probe.stageSet.Stage.Directions {
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
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Elision
			div := (&form.FormDiv{Name: "Elision"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elision {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elision",
				Label: "Elision",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Syllabic", inst.Syllabic, probe.formStage, formGroup)

		{
			// Slice of pointers: Text
			div := (&form.FormDiv{Name: "Text"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Text {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Text",
				Label: "Text",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Extend](), probe.formStage)
		StageSetBasicFieldtoForm("Laughing", inst.Laughing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Humming", inst.Humming, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_line", inst.End_line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End_paragraph", inst.End_paragraph, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
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
			for src := range probe.stageSet.Stage.Defaultss {
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
			for src := range probe.stageSet.Stage.Prints {
				if src.Measure_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Measure_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_numbering:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("System", inst.System, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Multiple_rest_always", inst.Multiple_rest_always, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Multiple_rest_range", inst.Multiple_rest_range, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Prints {
				if src.Measure_numbering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Measure_numbering", refNames, formGroup, probe.formStage)
		}
	case *models.Measure_repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Measure_styles {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Multiple_rest", inst.Multiple_rest, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Multiple_rest](), probe.formStage)
		StageSetAssociationFieldToForm("Measure_repeat", inst.Measure_repeat, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Measure_repeat](), probe.formStage)
		StageSetAssociationFieldToForm("Beat_repeat", inst.Beat_repeat, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Beat_repeat](), probe.formStage)
		StageSetAssociationFieldToForm("Slash", inst.Slash, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Slash](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Beat_unit", inst.Beat_unit, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beat_unit_dot", inst.Beat_unit_dot, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Per_minute", inst.Per_minute, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Per_minute](), probe.formStage)

		{
			// Slice of pointers: Beat_unit_tied
			div := (&form.FormDiv{Name: "Beat_unit_tied"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Beat_unit_tied {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Beat_unit_tied",
				Label: "Beat_unit_tied",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Metronome_arrows", inst.Metronome_arrows, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Metronome_relation", inst.Metronome_relation, probe.formStage, formGroup)

		{
			// Slice of pointers: Metronome_note
			div := (&form.FormDiv{Name: "Metronome_note"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Metronome_note {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Metronome_note",
				Label: "Metronome_note",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Metronome_notes {
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
		StageSetBasicFieldtoForm("Metronome_type", inst.Metronome_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Metronome_dot", inst.Metronome_dot, probe.formStage, formGroup)

		{
			// Slice of pointers: Metronome_beam
			div := (&form.FormDiv{Name: "Metronome_beam"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Metronome_beam {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Metronome_beam",
				Label: "Metronome_beam",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Metronome_tied", inst.Metronome_tied, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Metronome_tied](), probe.formStage)
		StageSetAssociationFieldToForm("Metronome_tuplet", inst.Metronome_tuplet, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Metronome_tuplet](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Metronomes {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Metronome_notes {
				if src.Metronome_tied == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Metronome_note", "Metronome_tied", refNames, formGroup, probe.formStage)
		}
	case *models.Metronome_tuplet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Metronome_notes {
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
			for src := range probe.stageSet.Stage.Score_parts {
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
			for src := range probe.stageSet.Stage.Sounds {
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
			for src := range probe.stageSet.Stage.Score_parts {
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
			for src := range probe.stageSet.Stage.Sounds {
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

		{
			// Slice of pointers: Miscellaneous_field
			div := (&form.FormDiv{Name: "Miscellaneous_field"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Miscellaneous_field {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Miscellaneous_field",
				Label: "Miscellaneous_field",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Identifications {
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
			for src := range probe.stageSet.Stage.Miscellaneouss {
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

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Measure_styles {
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

		{
			// Slice of pointers: Display_text
			div := (&form.FormDiv{Name: "Display_text"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Display_text {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Display_text",
				Label: "Display_text",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Accidental_text
			div := (&form.FormDiv{Name: "Accidental_text"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accidental_text {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accidental_text",
				Label: "Accidental_text",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
				if src.Group_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_groups {
				if src.Group_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Part_group", "Group_abbreviation_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Prints {
				if src.Part_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Part_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Prints {
				if src.Part_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Part_abbreviation_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
				if src.Part_name_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_name_display", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
				if src.Part_abbreviation_display == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_abbreviation_display", refNames, formGroup, probe.formStage)
		}
	case *models.Non_arpeggiate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			// Slice of pointers: Tied
			div := (&form.FormDiv{Name: "Tied"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tied {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tied",
				Label: "Tied",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Slur
			div := (&form.FormDiv{Name: "Slur"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Slur {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Slur",
				Label: "Slur",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tuplet
			div := (&form.FormDiv{Name: "Tuplet"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tuplet {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tuplet",
				Label: "Tuplet",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Glissando
			div := (&form.FormDiv{Name: "Glissando"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Glissando {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Glissando",
				Label: "Glissando",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Slide
			div := (&form.FormDiv{Name: "Slide"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Slide {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Slide",
				Label: "Slide",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ornaments
			div := (&form.FormDiv{Name: "Ornaments"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ornaments {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ornaments",
				Label: "Ornaments",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Technical
			div := (&form.FormDiv{Name: "Technical"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Technical {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Technical",
				Label: "Technical",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Articulations
			div := (&form.FormDiv{Name: "Articulations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Articulations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Articulations",
				Label: "Articulations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Dynamics
			div := (&form.FormDiv{Name: "Dynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dynamics",
				Label: "Dynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fermata
			div := (&form.FormDiv{Name: "Fermata"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fermata {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fermata",
				Label: "Fermata",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arpeggiate
			div := (&form.FormDiv{Name: "Arpeggiate"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arpeggiate {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arpeggiate",
				Label: "Arpeggiate",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Non_arpeggiate
			div := (&form.FormDiv{Name: "Non_arpeggiate"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Non_arpeggiate {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Non_arpeggiate",
				Label: "Non_arpeggiate",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Accidental_mark
			div := (&form.FormDiv{Name: "Accidental_mark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accidental_mark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accidental_mark",
				Label: "Accidental_mark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_notation
			div := (&form.FormDiv{Name: "Other_notation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_notation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_notation",
				Label: "Other_notation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_dot", inst.Print_dot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_lyric", inst.Print_lyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Grace", inst.Grace, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Grace](), probe.formStage)
		StageSetBasicFieldtoForm("Chord", inst.Chord, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pitch", inst.Pitch, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Pitch](), probe.formStage)
		StageSetAssociationFieldToForm("Unpitched", inst.Unpitched, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Unpitched](), probe.formStage)
		StageSetAssociationFieldToForm("Rest", inst.Rest, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rest](), probe.formStage)
		StageSetAssociationFieldToForm("Tie", inst.Tie, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tie](), probe.formStage)
		StageSetBasicFieldtoForm("Cue", inst.Cue, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)

		{
			// Slice of pointers: Instrument
			div := (&form.FormDiv{Name: "Instrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Instrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Instrument",
				Label: "Instrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Type", inst.Type, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note_type](), probe.formStage)

		{
			// Slice of pointers: Dot
			div := (&form.FormDiv{Name: "Dot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dot",
				Label: "Dot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Accidental", inst.Accidental, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Accidental](), probe.formStage)
		StageSetAssociationFieldToForm("Time_modification", inst.Time_modification, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Time_modification](), probe.formStage)
		StageSetAssociationFieldToForm("Stem", inst.Stem, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stem](), probe.formStage)
		StageSetAssociationFieldToForm("Notehead", inst.Notehead, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Notehead](), probe.formStage)
		StageSetAssociationFieldToForm("Notehead_text", inst.Notehead_text, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Notehead_text](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Beam", inst.Beam, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Beam](), probe.formStage)

		{
			// Slice of pointers: Notations
			div := (&form.FormDiv{Name: "Notations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Notations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Notations",
				Label: "Notations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Lyric
			div := (&form.FormDiv{Name: "Lyric"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lyric {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lyric",
				Label: "Lyric",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Play", inst.Play, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Play](), probe.formStage)
		StageSetAssociationFieldToForm("Listen", inst.Listen, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Listen](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Appearances {
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
			for src := range probe.stageSet.Stage.Notes {
				if src.Type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Type", refNames, formGroup, probe.formStage)
		}
	case *models.Notehead:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Filled", inst.Filled, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Notehead == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Notehead", refNames, formGroup, probe.formStage)
		}
	case *models.Notehead_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Display_text
			div := (&form.FormDiv{Name: "Display_text"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Display_text {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Display_text",
				Label: "Display_text",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Accidental_text
			div := (&form.FormDiv{Name: "Accidental_text"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accidental_text {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accidental_text",
				Label: "Accidental_text",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Notehead_text == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Notehead_text", refNames, formGroup, probe.formStage)
		}
	case *models.Numeral:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Numeral_root", inst.Numeral_root, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Numeral_root](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral_alter", inst.Numeral_alter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Harmony_alter](), probe.formStage)
		StageSetAssociationFieldToForm("Numeral_key", inst.Numeral_key, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Numeral_key](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
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
		StageSetBasicFieldtoForm("Numeral_mode", inst.Numeral_mode, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Numerals {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Numerals {
				if src.Numeral_root == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Numeral", "Numeral_root", refNames, formGroup, probe.formStage)
		}
	case *models.Octave_shift:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Size", inst.Size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Directions {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Listenings {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Listening", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sounds {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Works {
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

		{
			// Slice of pointers: Trill_mark
			div := (&form.FormDiv{Name: "Trill_mark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Trill_mark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Trill_mark",
				Label: "Trill_mark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Turn
			div := (&form.FormDiv{Name: "Turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Turn",
				Label: "Turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Delayed_turn
			div := (&form.FormDiv{Name: "Delayed_turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Delayed_turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Delayed_turn",
				Label: "Delayed_turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Inverted_turn
			div := (&form.FormDiv{Name: "Inverted_turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Inverted_turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Inverted_turn",
				Label: "Inverted_turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Delayed_inverted_turn
			div := (&form.FormDiv{Name: "Delayed_inverted_turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Delayed_inverted_turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Delayed_inverted_turn",
				Label: "Delayed_inverted_turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Vertical_turn
			div := (&form.FormDiv{Name: "Vertical_turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Vertical_turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Vertical_turn",
				Label: "Vertical_turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Inverted_vertical_turn
			div := (&form.FormDiv{Name: "Inverted_vertical_turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Inverted_vertical_turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Inverted_vertical_turn",
				Label: "Inverted_vertical_turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Shake
			div := (&form.FormDiv{Name: "Shake"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Shake {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Shake",
				Label: "Shake",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Wavy_line
			div := (&form.FormDiv{Name: "Wavy_line"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Wavy_line {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Wavy_line",
				Label: "Wavy_line",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Mordent
			div := (&form.FormDiv{Name: "Mordent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Mordent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Mordent",
				Label: "Mordent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Inverted_mordent
			div := (&form.FormDiv{Name: "Inverted_mordent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Inverted_mordent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Inverted_mordent",
				Label: "Inverted_mordent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Schleifer
			div := (&form.FormDiv{Name: "Schleifer"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Schleifer {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Schleifer",
				Label: "Schleifer",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tremolo
			div := (&form.FormDiv{Name: "Tremolo"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tremolo {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tremolo",
				Label: "Tremolo",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Haydn
			div := (&form.FormDiv{Name: "Haydn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Haydn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Haydn",
				Label: "Haydn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_ornament
			div := (&form.FormDiv{Name: "Other_ornament"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_ornament {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_ornament",
				Label: "Other_ornament",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Accidental_mark
			div := (&form.FormDiv{Name: "Accidental_mark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accidental_mark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accidental_mark",
				Label: "Accidental_mark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Appearances {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Other_direction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Other_direction", refNames, formGroup, probe.formStage)
		}
	case *models.Other_listening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Listens {
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
			for src := range probe.stageSet.Stage.Listenings {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
			for src := range probe.stageSet.Stage.Ornamentss {
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
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plays {
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
			for src := range probe.stageSet.Stage.Dynamicss {
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
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetAssociationFieldToForm("Page_margins", inst.Page_margins, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Page_margins](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
				if src.Page_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "Page_layout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Prints {
				if src.Page_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Page_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Page_margins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Left_margin", inst.Left_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Right_margin", inst.Right_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Top_margin", inst.Top_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bottom_margin", inst.Bottom_margin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Page_layouts {
				if src.Page_margins == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Page_layout", "Page_margins", refNames, formGroup, probe.formStage)
		}
	case *models.Part_clef:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sign", inst.Sign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Clef_octave_change", inst.Clef_octave_change, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.For_parts {
				if src.Part_clef == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.For_part", "Part_clef", refNames, formGroup, probe.formStage)
		}
	case *models.Part_group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Group_name", inst.Group_name, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Group_name](), probe.formStage)
		StageSetAssociationFieldToForm("Group_name_display", inst.Group_name_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Group_abbreviation", inst.Group_abbreviation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Group_name](), probe.formStage)
		StageSetAssociationFieldToForm("Group_abbreviation_display", inst.Group_abbreviation_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Group_symbol", inst.Group_symbol, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Group_symbol](), probe.formStage)
		StageSetAssociationFieldToForm("Group_barline", inst.Group_barline, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Group_barline](), probe.formStage)
		StageSetBasicFieldtoForm("Group_time", inst.Group_time, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Formatted_text](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_lists {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			// Slice of pointers: Instrument_link
			div := (&form.FormDiv{Name: "Instrument_link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Instrument_link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Instrument_link",
				Label: "Instrument_link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Group_link", inst.Group_link, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
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
		StageSetAssociationFieldToForm("Part_group", inst.Part_group, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_group](), probe.formStage)
		StageSetAssociationFieldToForm("Score_part", inst.Score_part, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Score_part](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
				if src.Part_list == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Part_list", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_timewises {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
				if src.Part_name == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_part", "Part_name", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
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
			for src := range probe.stageSet.Stage.Attributess {
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
			for src := range probe.stageSet.Stage.For_parts {
				if src.Part_transpose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.For_part", "Part_transpose", refNames, formGroup, probe.formStage)
		}
	case *models.Pedal:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line", inst.Line, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Abbreviated", inst.Abbreviated, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Harp_pedalss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Metronomes {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Enclosure", inst.Enclosure, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Glass", inst.Glass, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Glass](), probe.formStage)
		StageSetAssociationFieldToForm("Metal", inst.Metal, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Metal](), probe.formStage)
		StageSetAssociationFieldToForm("Wood", inst.Wood, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Wood](), probe.formStage)
		StageSetAssociationFieldToForm("Pitched", inst.Pitched, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Pitched](), probe.formStage)
		StageSetAssociationFieldToForm("Membrane", inst.Membrane, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Membrane](), probe.formStage)
		StageSetAssociationFieldToForm("Effect", inst.Effect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Effect](), probe.formStage)
		StageSetAssociationFieldToForm("Timpani", inst.Timpani, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Timpani](), probe.formStage)
		StageSetAssociationFieldToForm("Beater", inst.Beater, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Beater](), probe.formStage)
		StageSetAssociationFieldToForm("Stick", inst.Stick, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stick](), probe.formStage)
		StageSetBasicFieldtoForm("Stick_location", inst.Stick_location, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Other_percussion", inst.Other_percussion, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Other_text](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bends {
				if src.With_bar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bend", "With_bar", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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
		StageSetBasicFieldtoForm("Mute", inst.Mute, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Semi_pitched", inst.Semi_pitched, probe.formStage, formGroup)

		{
			// Slice of pointers: Other_play
			div := (&form.FormDiv{Name: "Other_play"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_play {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_play",
				Label: "Other_play",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Play == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Play", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sounds {
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
			for src := range probe.stageSet.Stage.Score_parts {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Symbol", inst.Symbol, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
				if src.Principal_voice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Direction_type", "Principal_voice", refNames, formGroup, probe.formStage)
		}
	case *models.Print:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff_spacing", inst.Staff_spacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("New_system", inst.New_system, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("New_page", inst.New_page, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Blank_page", inst.Blank_page, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page_number", inst.Page_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Page_layout", inst.Page_layout, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Page_layout](), probe.formStage)
		StageSetAssociationFieldToForm("System_layout", inst.System_layout, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.System_layout](), probe.formStage)

		{
			// Slice of pointers: Staff_layout
			div := (&form.FormDiv{Name: "Staff_layout"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staff_layout {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staff_layout",
				Label: "Staff_layout",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Measure_layout", inst.Measure_layout, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Measure_layout](), probe.formStage)
		StageSetAssociationFieldToForm("Measure_numbering", inst.Measure_numbering, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Measure_numbering](), probe.formStage)
		StageSetAssociationFieldToForm("Part_name_display", inst.Part_name_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation_display", inst.Part_abbreviation_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
				for _, target := range src.Print {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_measure", "Print", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_part_1s {
				for _, target := range src.Print {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_part_1", "Print", refNames, formGroup, probe.formStage)
		}
	case *models.Release:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bends {
				if src.Release == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bend", "Release", refNames, formGroup, probe.formStage)
		}
	case *models.Repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Direction", inst.Direction, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Times", inst.Times, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("After_jump", inst.After_jump, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Winged", inst.Winged, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
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
			for src := range probe.stageSet.Stage.Notes {
				if src.Rest == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Rest", refNames, formGroup, probe.formStage)
		}
	case *models.Root:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Root_step", inst.Root_step, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Root_step](), probe.formStage)
		StageSetAssociationFieldToForm("Root_alter", inst.Root_alter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Harmony_alter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Roots {
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
			for src := range probe.stageSet.Stage.Defaultss {
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

		{
			// Slice of pointers: Accord
			div := (&form.FormDiv{Name: "Accord"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accord {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accord",
				Label: "Accord",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetAssociationFieldToForm("Virtual_instrument", inst.Virtual_instrument, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Virtual_instrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_parts {
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
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Identification](), probe.formStage)

		{
			// Slice of pointers: Part_link
			div := (&form.FormDiv{Name: "Part_link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part_link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part_link",
				Label: "Part_link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Part_name", inst.Part_name, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_name](), probe.formStage)
		StageSetAssociationFieldToForm("Part_name_display", inst.Part_name_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation", inst.Part_abbreviation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_name](), probe.formStage)
		StageSetAssociationFieldToForm("Part_abbreviation_display", inst.Part_abbreviation_display, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Name_display](), probe.formStage)
		StageSetBasicFieldtoForm("Group", inst.Group, probe.formStage, formGroup)

		{
			// Slice of pointers: Score_instrument
			div := (&form.FormDiv{Name: "Score_instrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Score_instrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Score_instrument",
				Label: "Score_instrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Player
			div := (&form.FormDiv{Name: "Player"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Player {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Player",
				Label: "Player",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Midi_device
			div := (&form.FormDiv{Name: "Midi_device"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Midi_device {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Midi_device",
				Label: "Midi_device",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Midi_instrument
			div := (&form.FormDiv{Name: "Midi_instrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Midi_instrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Midi_instrument",
				Label: "Midi_instrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Part_lists {
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
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Work](), probe.formStage)
		StageSetBasicFieldtoForm("Movement_number", inst.Movement_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Movement_title", inst.Movement_title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Defaults](), probe.formStage)

		{
			// Slice of pointers: Credit
			div := (&form.FormDiv{Name: "Credit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit",
				Label: "Credit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Part_list", inst.Part_list, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_list](), probe.formStage)

		{
			// Slice of pointers: Part
			div := (&form.FormDiv{Name: "Part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part",
				Label: "Part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.Score_timewise:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Version", inst.Version, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Work](), probe.formStage)
		StageSetBasicFieldtoForm("Movement_number", inst.Movement_number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Movement_title", inst.Movement_title, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Defaults](), probe.formStage)

		{
			// Slice of pointers: Credit
			div := (&form.FormDiv{Name: "Credit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit",
				Label: "Credit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Part_list", inst.Part_list, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Part_list](), probe.formStage)

		{
			// Slice of pointers: Measure
			div := (&form.FormDiv{Name: "Measure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Measure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Measure",
				Label: "Measure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.Segno:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Segno_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Segno_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Use_dots", inst.Use_dots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Use_stems", inst.Use_stems, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Slash_type", inst.Slash_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slash_dot", inst.Slash_dot, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Except_voice", inst.Except_voice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Measure_styles {
				if src.Slash == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Measure_style", "Slash", refNames, formGroup, probe.formStage)
		}
	case *models.Slide:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("First_beat", inst.First_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Orientation", inst.Orientation, probe.formStage, formGroup)
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
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Pan", inst.Pan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Elevation", inst.Elevation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Damper_pedal", inst.Damper_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Soft_pedal", inst.Soft_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Sostenuto_pedal", inst.Sostenuto_pedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Instrument_change
			div := (&form.FormDiv{Name: "Instrument_change"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Instrument_change {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Instrument_change",
				Label: "Instrument_change",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Midi_device
			div := (&form.FormDiv{Name: "Midi_device"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Midi_device {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Midi_device",
				Label: "Midi_device",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Midi_instrument
			div := (&form.FormDiv{Name: "Midi_instrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Midi_instrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Midi_instrument",
				Label: "Midi_instrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Play
			div := (&form.FormDiv{Name: "Play"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Play {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Play",
				Label: "Play",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Swing", inst.Swing, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Swing](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.A_measures {
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
			for src := range probe.stageSet.Stage.A_part_1s {
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
			for src := range probe.stageSet.Stage.Directions {
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
		StageSetBasicFieldtoForm("Show_frets", inst.Show_frets, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Print_spacing", inst.Print_spacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Staff_type", inst.Staff_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff_lines", inst.Staff_lines, probe.formStage, formGroup)

		{
			// Slice of pointers: Line_detail
			div := (&form.FormDiv{Name: "Line_detail"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Line_detail {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Line_detail",
				Label: "Line_detail",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staff_tuning
			div := (&form.FormDiv{Name: "Staff_tuning"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staff_tuning {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staff_tuning",
				Label: "Staff_tuning",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Capo", inst.Capo, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Staff_size", inst.Staff_size, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Staff_size](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
			for src := range probe.stageSet.Stage.Defaultss {
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
			for src := range probe.stageSet.Stage.Prints {
				for _, target := range src.Staff_layout {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "Staff_layout", refNames, formGroup, probe.formStage)
		}
	case *models.Staff_size:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Scaling", inst.Scaling, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Staff_detailss {
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
			for src := range probe.stageSet.Stage.Staff_detailss {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
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
		StageSetBasicFieldtoForm("Stick_type", inst.Stick_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stick_material", inst.Stick_material, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
				if src.Stick == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Percussion", "Stick", refNames, formGroup, probe.formStage)
		}
	case *models.String_mute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Frame_notes {
				if src.String == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Frame_note", "String", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
				for _, target := range src.String {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Technical", "String", refNames, formGroup, probe.formStage)
		}
	case *models.Strong_accent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Articulationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Basss {
				if src.Bass_separator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Bass", "Bass_separator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Prefix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Prefix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Figure_number == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Figure_number", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Figures {
				if src.Suffix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Figure", "Suffix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Harmonys {
				if src.Function == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Harmony", "Function", refNames, formGroup, probe.formStage)
		}
	case *models.Supports:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Attribute", inst.Attribute, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Encodings {
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
		StageSetEnumStringFieldToForm("Swing_type", inst.Swing_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Swing_style", inst.Swing_style, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sounds {
				if src.Swing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sound", "Swing", refNames, formGroup, probe.formStage)
		}
	case *models.Sync:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Latency", inst.Latency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Listenings {
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
		StageSetAssociationFieldToForm("Left_divider", inst.Left_divider, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_print_object_style_align](), probe.formStage)
		StageSetAssociationFieldToForm("Right_divider", inst.Right_divider, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Empty_print_object_style_align](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.System_layouts {
				if src.System_dividers == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_layout", "System_dividers", refNames, formGroup, probe.formStage)
		}
	case *models.System_layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("System_margins", inst.System_margins, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.System_margins](), probe.formStage)
		StageSetBasicFieldtoForm("System_distance", inst.System_distance, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Top_system_distance", inst.Top_system_distance, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("System_dividers", inst.System_dividers, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.System_dividers](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Defaultss {
				if src.System_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Defaults", "System_layout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Prints {
				if src.System_layout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Print", "System_layout", refNames, formGroup, probe.formStage)
		}
	case *models.System_margins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Left_margin", inst.Left_margin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Right_margin", inst.Right_margin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.System_layouts {
				if src.System_margins == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.System_layout", "System_margins", refNames, formGroup, probe.formStage)
		}
	case *models.Tap:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Hand", inst.Hand, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Technicals {
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

		{
			// Slice of pointers: Up_bow
			div := (&form.FormDiv{Name: "Up_bow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Up_bow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Up_bow",
				Label: "Up_bow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Down_bow
			div := (&form.FormDiv{Name: "Down_bow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Down_bow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Down_bow",
				Label: "Down_bow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmonic
			div := (&form.FormDiv{Name: "Harmonic"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmonic {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmonic",
				Label: "Harmonic",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Open_string
			div := (&form.FormDiv{Name: "Open_string"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Open_string {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Open_string",
				Label: "Open_string",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Thumb_position
			div := (&form.FormDiv{Name: "Thumb_position"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Thumb_position {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Thumb_position",
				Label: "Thumb_position",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fingering
			div := (&form.FormDiv{Name: "Fingering"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fingering {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fingering",
				Label: "Fingering",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pluck
			div := (&form.FormDiv{Name: "Pluck"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pluck {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pluck",
				Label: "Pluck",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Double_tongue
			div := (&form.FormDiv{Name: "Double_tongue"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Double_tongue {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Double_tongue",
				Label: "Double_tongue",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Triple_tongue
			div := (&form.FormDiv{Name: "Triple_tongue"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Triple_tongue {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Triple_tongue",
				Label: "Triple_tongue",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Stopped
			div := (&form.FormDiv{Name: "Stopped"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stopped {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stopped",
				Label: "Stopped",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Snap_pizzicato
			div := (&form.FormDiv{Name: "Snap_pizzicato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Snap_pizzicato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Snap_pizzicato",
				Label: "Snap_pizzicato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fret
			div := (&form.FormDiv{Name: "Fret"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fret {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fret",
				Label: "Fret",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: String
			div := (&form.FormDiv{Name: "String"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.String {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "String",
				Label: "String",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Hammer_on
			div := (&form.FormDiv{Name: "Hammer_on"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Hammer_on {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Hammer_on",
				Label: "Hammer_on",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pull_off
			div := (&form.FormDiv{Name: "Pull_off"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pull_off {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pull_off",
				Label: "Pull_off",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bend
			div := (&form.FormDiv{Name: "Bend"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bend {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bend",
				Label: "Bend",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tap
			div := (&form.FormDiv{Name: "Tap"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tap {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tap",
				Label: "Tap",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Heel
			div := (&form.FormDiv{Name: "Heel"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Heel {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Heel",
				Label: "Heel",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Toe
			div := (&form.FormDiv{Name: "Toe"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Toe {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Toe",
				Label: "Toe",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fingernails
			div := (&form.FormDiv{Name: "Fingernails"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fingernails {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fingernails",
				Label: "Fingernails",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Hole
			div := (&form.FormDiv{Name: "Hole"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Hole {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Hole",
				Label: "Hole",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arrow
			div := (&form.FormDiv{Name: "Arrow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arrow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arrow",
				Label: "Arrow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Handbell
			div := (&form.FormDiv{Name: "Handbell"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Handbell {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Handbell",
				Label: "Handbell",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Brass_bend
			div := (&form.FormDiv{Name: "Brass_bend"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Brass_bend {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Brass_bend",
				Label: "Brass_bend",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Flip
			div := (&form.FormDiv{Name: "Flip"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Flip {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Flip",
				Label: "Flip",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Smear
			div := (&form.FormDiv{Name: "Smear"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Smear {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Smear",
				Label: "Smear",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Open
			div := (&form.FormDiv{Name: "Open"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Open {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Open",
				Label: "Open",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Half_muted
			div := (&form.FormDiv{Name: "Half_muted"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Half_muted {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Half_muted",
				Label: "Half_muted",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmon_mute
			div := (&form.FormDiv{Name: "Harmon_mute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmon_mute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmon_mute",
				Label: "Harmon_mute",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Golpe
			div := (&form.FormDiv{Name: "Golpe"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Golpe {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Golpe",
				Label: "Golpe",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Other_technical
			div := (&form.FormDiv{Name: "Other_technical"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Other_technical {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Other_technical",
				Label: "Other_technical",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_through", inst.Line_through, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Letter_spacing", inst.Letter_spacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dir", inst.Dir, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lyrics {
				for _, target := range src.Text {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lyric", "Text", refNames, formGroup, probe.formStage)
		}
	case *models.Tie:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				if src.Tie == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Tie", refNames, formGroup, probe.formStage)
		}
	case *models.Tied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dash_length", inst.Dash_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space_length", inst.Space_length, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Orientation", inst.Orientation, probe.formStage, formGroup)
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
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Halign", inst.Halign, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Valign", inst.Valign, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Print_object", inst.Print_object, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beat_type", inst.Beat_type, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Interchangeable", inst.Interchangeable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Interchangeable](), probe.formStage)
		StageSetBasicFieldtoForm("Senza_misura", inst.Senza_misura, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributess {
				for _, target := range src.Time {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attributes", "Time", refNames, formGroup, probe.formStage)
		}
	case *models.Time_modification:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actual_notes", inst.Actual_notes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Normal_notes", inst.Normal_notes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Normal_type", inst.Normal_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Normal_dot", inst.Normal_dot, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
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
			for src := range probe.stageSet.Stage.Percussions {
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
			for src := range probe.stageSet.Stage.Attributess {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_family", inst.Font_family, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ornamentss {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Show_number", inst.Show_number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show_type", inst.Show_type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line_shape", inst.Line_shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Tuplet_actual", inst.Tuplet_actual, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tuplet_portion](), probe.formStage)
		StageSetAssociationFieldToForm("Tuplet_normal", inst.Tuplet_normal, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tuplet_portion](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notationss {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tuplet_portions {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tuplet_portions {
				if src.Tuplet_number == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet_portion", "Tuplet_number", refNames, formGroup, probe.formStage)
		}
	case *models.Tuplet_portion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Tuplet_number", inst.Tuplet_number, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tuplet_number](), probe.formStage)
		StageSetAssociationFieldToForm("Tuplet_type", inst.Tuplet_type, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tuplet_type](), probe.formStage)

		{
			// Slice of pointers: Tuplet_dot
			div := (&form.FormDiv{Name: "Tuplet_dot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tuplet_dot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tuplet_dot",
				Label: "Tuplet_dot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tuplets {
				if src.Tuplet_actual == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet", "Tuplet_actual", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tuplets {
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
		StageSetBasicFieldtoForm("Font_style", inst.Font_style, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_size", inst.Font_size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Font_weight", inst.Font_weight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnclosedText", inst.EnclosedText, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tuplet_portions {
				if src.Tuplet_type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tuplet_portion", "Tuplet_type", refNames, formGroup, probe.formStage)
		}
	case *models.Typed_text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Encodings {
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
			for src := range probe.stageSet.Stage.Identifications {
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
			for src := range probe.stageSet.Stage.Identifications {
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
			for src := range probe.stageSet.Stage.Identifications {
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
			for src := range probe.stageSet.Stage.Notes {
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
			for src := range probe.stageSet.Stage.Instrument_changes {
				if src.Virtual_instrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Instrument_change", "Virtual_instrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_instruments {
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
		StageSetEnumStringFieldToForm("Time_only", inst.Time_only, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Listens {
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
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_x", inst.Default_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default_y", inst.Default_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_x", inst.Relative_x, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Relative_y", inst.Relative_y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placement", inst.Placement, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start_note", inst.Start_note, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Trill_step", inst.Trill_step, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Two_note_turn", inst.Two_note_turn, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second_beat", inst.Second_beat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Last_beat", inst.Last_beat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Barlines {
				if src.Wavy_line == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Barline", "Wavy_line", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ornamentss {
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
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Spread", inst.Spread, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Niente", inst.Niente, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line_type", inst.Line_type, probe.formStage, formGroup)
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
			for src := range probe.stageSet.Stage.Direction_types {
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
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Percussions {
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
		StageSetAssociationFieldToForm("Opus", inst.Opus, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Opus](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_partwises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_partwise", "Work", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Score_timewises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Score_timewise", "Work", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
