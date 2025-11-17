package stage_test

import (
	"log"
	"testing"

	models "github.com/fullstack-lang/gong/test/test/go/models"
	test_stack "github.com/fullstack-lang/gong/test/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/test/go/static"
	"github.com/xuri/excelize/v2"
)

// TestStageCount
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
func TestStageCount(t *testing.T) {

	r := test_static.ServeStaticFiles(false)
	stack := test_stack.NewStack(r, "test", "", "", "", true, true)
	stage := stack.Stage

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage(stage)

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
		Anarrayofb: []*models.Bstruct{
			bclass1,
		},
	})
	_ = aclass1

	// empty stage
	want := 0
	got := len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage one instance
	aclass1.Stage(stage)

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	aclass1.Stage(stage)

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// unstage it
	aclass1.Unstage(stage)

	want = 0
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	aclass1.Stage(stage)

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	models.SerializeExcelizePointerToGongstruct[*models.Astruct](stage, f)

	if err := f.SaveAs("Astructs.xlsx"); err != nil {
		log.Fatal(err)
	}
}
