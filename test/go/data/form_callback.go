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
) (astructFormCallback *AstructFormCallback) {
	astructFormCallback = new(AstructFormCallback)
	astructFormCallback.stageOfInterest = stageOfInterest
	astructFormCallback.formStage = formStage
	return
}

type AstructFormCallback struct {
	stageOfInterest *models.StageStruct
	formStage       *table.StageStruct
}

func (astructFormCallback *AstructFormCallback) OnSave() {

	log.Println("AstructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	astructFormCallback.formStage.Checkout()

	astruct := new(models.Astruct).Stage(astructFormCallback.stageOfInterest)

	// get the formGroup
	formGroup := astructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			newValue := formDiv.FormFields[0].FormFieldString.Value
			astruct.Name = newValue
		case "Date":
			date := formDiv.FormFields[0].FormFieldDate.Value

			time := formDiv.FormFields[1].FormFieldTime.Value
			astruct.Date = addTimeComponents(date, time)
		case "Booleanfield":
			value := formDiv.CheckBoxs[0].Value
			astruct.Booleanfield = value
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
			newValue := formDiv.FormFields[0].FormFieldInt.Value
			astruct.Intfield = newValue
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
