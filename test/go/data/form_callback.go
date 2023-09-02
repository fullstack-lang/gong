package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func NewAstructFormCallback(
	stageOfInterest *models.StageStruct,
	formStage *table.StageStruct,
	astruct *models.Astruct,
) (astructFormCallback *AstructFormCallback) {
	astructFormCallback = new(AstructFormCallback)
	astructFormCallback.stageOfInterest = stageOfInterest
	astructFormCallback.formStage = formStage
	astructFormCallback.astruct = astruct
	return
}

type AstructFormCallback struct {
	stageOfInterest *models.StageStruct
	formStage       *table.StageStruct
	astruct         *models.Astruct
}

func (astructFormCallback *AstructFormCallback) OnSave() {

	log.Println("AstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructFormCallback.formStage.Checkout()

	if astructFormCallback.astruct == nil {
		astructFormCallback.astruct = new(models.Astruct).Stage(astructFormCallback.stageOfInterest)
	}
	astruct := astructFormCallback.astruct

	// get the formGroup
	formGroup := astructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&(astruct.Name), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(astruct.Date), formDiv)
		case "Booleanfield":
			FormDivBasicFieldToField(&(astruct.Booleanfield), formDiv)
		case "Aenum":
			if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
				if err := (&astruct.Aenum).FromString(value.GetName()); err != nil {
					log.Println("Unkwnown enum value", value.GetName())
				}
			}
		case "Cenum":
			if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
				if err := (&astruct.CEnum).FromCodeString(value.GetName()); err != nil {
					log.Println("Unkwnown enum value", value.GetName())
				}

			}
		case "Intfield":
			FormDivBasicFieldToField(&(astruct.Intfield), formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&(astruct.Duration1), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(astruct.Floatfield), formDiv)
		}
	}

	astructFormCallback.stageOfInterest.Commit()
}

func addTimeComponents(x, y time.Time) time.Time {
	h, m, s := y.Clock()
	x = x.Add(time.Duration(h) * time.Hour)
	x = x.Add(time.Duration(m) * time.Minute)
	x = x.Add(time.Duration(s) * time.Second)
	x = x.Add(time.Duration(y.Nanosecond()) * time.Nanosecond)
	return x
}
