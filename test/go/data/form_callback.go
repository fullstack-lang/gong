package data

import (
	"log"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	table "github.com/fullstack-lang/gongtable/go/models"
)

func NewAstructFormCallback(
	stageOfInterest *models.StageStruct,
) (astructFormCallback *AstructFormCallback) {
	astructFormCallback = new(AstructFormCallback)
	astructFormCallback.stageOfInterest = stageOfInterest
	return
}

type AstructFormCallback struct {
	stageOfInterest *models.StageStruct
}

func (astructFormCallback *AstructFormCallback) BeforeCommit(formStage *table.StageStruct) {

	log.Println("AstructFormCallback, BeforeCommit")

	astruct := new(models.Astruct).Stage(astructFormCallback.stageOfInterest)

	// get the formGroup
	formGroup := formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			astruct.Name = formDiv.FormFields[0].FormFieldString.Value
		case "Date":
			date := formDiv.FormFields[0].FormFieldDate.Value

			time := formDiv.FormFields[1].FormFieldTime.Value
			astruct.Date = addTimeComponents(date, time)
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
