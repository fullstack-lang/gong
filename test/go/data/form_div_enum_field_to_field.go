package data

import (
	"log"

	"github.com/fullstack-lang/gong/test/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
)

func FormDivEnumStringFieldToField[TF models.PointerToGongstructEnumStringField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			log.Println("Unkwnown enum value", value.GetName())
		}
	}
}

func FormDivEnumIntFieldToField[TF models.PointerToGongstructEnumIntField](field TF, formDiv *form.FormDiv) {
	if value := formDiv.FormFields[0].FormFieldSelect.Value; value != nil {
		if err := (field).FromCodeString(value.GetName()); err != nil {
			log.Println("Unkwnown enum value", value.GetName())
		}
	}
}
