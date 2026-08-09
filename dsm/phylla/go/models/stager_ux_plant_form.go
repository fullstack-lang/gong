package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

type plantFormOnSave struct {
	stager    *Stager
	plant     *PlantAbstract
	formGroup *form.FormGroup
}

func (onSave *plantFormOnSave) OnSave() {
	// 1. Checkout plantFormStage so front changes are synchronized into back repo structs
	onSave.stager.plantFormStage.Checkout()

	if onSave.plant == nil {
		return
	}

	modified := false
	for _, formDiv := range onSave.formGroup.FormDivs {
		switch formDiv.Name {
		case "PlantType":
			if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
				var newType PlantType
				if err := (&newType).FromCodeString(value.Name); err == nil {
					if onSave.plant.PlantType != newType {
						onSave.plant.PlantType = newType
						if newType != Vase {
							onSave.plant.CurrentView = VIEW_PLANT_2D
						}
						modified = true
					}
				}
			}
		}
	}

	if modified {
		onSave.stager.stage.Commit()
	}
}

func (stager *Stager) ux_plant_form() {
	stager.plantFormStage.Reset()

	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}
	if plant == nil {
		stager.plantFormStage.Commit()
		return
	}

	formGroup := (&form.FormGroup{
		Name:  "Plant Form",
		Label: "Plant Configuration",
	}).Stage(stager.plantFormStage)

	formDiv := (&form.FormDiv{
		Name: "PlantType",
	}).Stage(stager.plantFormStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	formField := (&form.FormField{
		Name:  "PlantType",
		Label: "Plant Type",
	}).Stage(stager.plantFormStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name: "enum",
	}).Stage(stager.plantFormStage)
	formField.FormFieldSelect = formFieldSelect

	formFieldSelect.Options = make([]*form.Option, 0)
	for idx, optionCode := range plant.PlantType.Codes() {
		optionValue := plant.PlantType.CodeValues()[idx]

		option := (&form.Option{
			Name: optionCode,
		}).Stage(stager.plantFormStage)

		if plant.PlantType.ToString() == optionValue {
			formFieldSelect.Value = option
		}

		formFieldSelect.Options = append(formFieldSelect.Options, option)
	}

	formGroup.OnSave = &plantFormOnSave{
		stager:    stager,
		plant:     plant,
		formGroup: formGroup,
	}

	stager.plantFormStage.Commit()
}
