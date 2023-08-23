package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/test/go/fullstack"
	"github.com/fullstack-lang/gong/test/go/models"
)

func TestSerialize(t *testing.T) {

	stage, _ := fullstack.NewStackInstance(nil, "")

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage(stage).Commit(stage)
	bclass2 := (&models.Bstruct{Name: "B2"}).Stage(stage).Commit(stage)

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage(stage).Commit(stage)

	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass2)
	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass1)

	models.SerializeStage(stage, "serialize_test.xlsx")
}
