// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
	case *models.DisplaySelection:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("XLFile", inst.XLFile, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.XLFile](), probe.formStage)
		StageSetAssociationFieldToForm("XLSheet", inst.XLSheet, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.XLSheet](), probe.formStage)
	case *models.XLCell:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.XLRows {
				for _, target := range src.Cells {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.XLRow", "Cells", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.XLSheets {
				for _, target := range src.SheetCells {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.XLSheet", "SheetCells", refNames, formGroup, probe.formStage)
		}
	case *models.XLFile:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbSheets", inst.NbSheets, probe.formStage, formGroup)

		{
			// Slice of pointers: Sheets
			div := (&form.FormDiv{Name: "Sheets"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sheets {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sheets",
				Label: "Sheets",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DisplaySelections {
				if src.XLFile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DisplaySelection", "XLFile", refNames, formGroup, probe.formStage)
		}
	case *models.XLRow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RowIndex", inst.RowIndex, probe.formStage, formGroup)

		{
			// Slice of pointers: Cells
			div := (&form.FormDiv{Name: "Cells"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Cells {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Cells",
				Label: "Cells",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.XLSheets {
				for _, target := range src.Rows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.XLSheet", "Rows", refNames, formGroup, probe.formStage)
		}
	case *models.XLSheet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxRow", inst.MaxRow, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxCol", inst.MaxCol, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbRows", inst.NbRows, probe.formStage, formGroup)

		{
			// Slice of pointers: Rows
			div := (&form.FormDiv{Name: "Rows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rows",
				Label: "Rows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SheetCells
			div := (&form.FormDiv{Name: "SheetCells"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SheetCells {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SheetCells",
				Label: "SheetCells",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DisplaySelections {
				if src.XLSheet == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DisplaySelection", "XLSheet", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.XLFiles {
				for _, target := range src.Sheets {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.XLFile", "Sheets", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
