// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ButtonFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Button", true)
			} else {
				FillUpFormFromGongstruct(onSave.button, probe)
			}
		case *CellFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Cell", true)
			} else {
				FillUpFormFromGongstruct(onSave.cell, probe)
			}
		case *CellBooleanFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CellBoolean", true)
			} else {
				FillUpFormFromGongstruct(onSave.cellboolean, probe)
			}
		case *CellFloat64FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CellFloat64", true)
			} else {
				FillUpFormFromGongstruct(onSave.cellfloat64, probe)
			}
		case *CellIconFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CellIcon", true)
			} else {
				FillUpFormFromGongstruct(onSave.cellicon, probe)
			}
		case *CellIntFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CellInt", true)
			} else {
				FillUpFormFromGongstruct(onSave.cellint, probe)
			}
		case *CellStringFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CellString", true)
			} else {
				FillUpFormFromGongstruct(onSave.cellstring, probe)
			}
		case *DisplayedColumnFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DisplayedColumn", true)
			} else {
				FillUpFormFromGongstruct(onSave.displayedcolumn, probe)
			}
		case *RowFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Row", true)
			} else {
				FillUpFormFromGongstruct(onSave.row, probe)
			}
		case *SVGIconFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SVGIcon", true)
			} else {
				FillUpFormFromGongstruct(onSave.svgicon, probe)
			}
		case *TableFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Table", true)
			} else {
				FillUpFormFromGongstruct(onSave.table, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Button":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		button := new(models.Button)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(button, formGroup, probe)
	case "Cell":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Cell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFormCallback(
			nil,
			probe,
			formGroup,
		)
		cell := new(models.Cell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cell, formGroup, probe)
	case "CellBoolean":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CellBoolean Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellBooleanFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellboolean := new(models.CellBoolean)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellboolean, formGroup, probe)
	case "CellFloat64":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CellFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFloat64FormCallback(
			nil,
			probe,
			formGroup,
		)
		cellfloat64 := new(models.CellFloat64)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellfloat64, formGroup, probe)
	case "CellIcon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CellIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellicon := new(models.CellIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellicon, formGroup, probe)
	case "CellInt":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CellInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellIntFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellint := new(models.CellInt)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellint, formGroup, probe)
	case "CellString":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CellString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellStringFormCallback(
			nil,
			probe,
			formGroup,
		)
		cellstring := new(models.CellString)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cellstring, formGroup, probe)
	case "DisplayedColumn":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DisplayedColumn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplayedColumnFormCallback(
			nil,
			probe,
			formGroup,
		)
		displayedcolumn := new(models.DisplayedColumn)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(displayedcolumn, formGroup, probe)
	case "Row":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Row Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RowFormCallback(
			nil,
			probe,
			formGroup,
		)
		row := new(models.Row)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(row, formGroup, probe)
	case "SVGIcon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SVGIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		svgicon := new(models.SVGIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svgicon, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			probe,
			formGroup,
		)
		table := new(models.Table)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(table, formGroup, probe)
	}
	formStage.Commit()
}
