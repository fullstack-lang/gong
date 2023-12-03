package tests

import (
	"testing"

	test_stack "github.com/fullstack-lang/gong/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/go/static"

	"github.com/fullstack-lang/gong/test/go/models"
)

func TestSerialize(t *testing.T) {

	r := test_static.ServeStaticFiles(false)

	// setup stack
	stack := test_stack.NewTranscientStack(r, "test", false)
	stage := stack.Stage

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
