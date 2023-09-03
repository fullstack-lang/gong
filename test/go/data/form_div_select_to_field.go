package data

import (
	"github.com/fullstack-lang/gong/test/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
)

func FormDivSelectFieldToField[TF models.PointerToGongstruct](field *TF, stageOfInterest *models.StageStruct, formDiv *form.FormDiv) {

	if formDiv.FormFields[0].FormFieldSelect.Value == nil {
		*field = nil
	} else {
		for _instance := range *models.GetGongstructInstancesSetFromPointerType[TF](stageOfInterest) {
			if any(_instance).(TF).GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
				*field = any(_instance).(TF)
			}
		}
	}
}
