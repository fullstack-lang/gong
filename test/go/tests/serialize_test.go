package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

func TestSerialize(t *testing.T) {

	orm.SetupModels(false, ":memory:")

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage().Commit()
	bclass2 := (&models.Bstruct{Name: "B2"}).Stage().Commit()

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()

	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass2)
	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass1)

	models.SerializeStage("serialize_test.xlsx")
}
