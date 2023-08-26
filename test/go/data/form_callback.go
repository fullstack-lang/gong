package data

import (
	"log"

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

	// get Name field
	{
		formDiv := formGroup.FormDivs[0]
		astruct.Name = formDiv.FormFields[0].FormFieldString.Value
	}

	astructFormCallback.stageOfInterest.Commit()
}
