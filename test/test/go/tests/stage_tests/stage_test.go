package stage_test

import (
	"testing"

	models "github.com/fullstack-lang/gong/test/test/go/models"
	test_stack "github.com/fullstack-lang/gong/test/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/test/go/static"
)

// TestStageCount
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
func TestStageCount(t *testing.T) {

	r := test_static.ServeStaticFiles(false)
	stack := test_stack.NewStack(r, "test", "", "", "", true, true)
	stage := stack.Stage

	b1 := (&models.Bstruct{Name: "B1"}).Stage(stage)
	b2 := (&models.Bstruct{Name: "B2"}).Stage(stage)

	a1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      b2,
		Anarrayofb: []*models.Bstruct{
			b1,
			b2,
		},
	})
	_ = a1

	// empty stage
	want := 0
	got := len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage one instance
	a1.Stage(stage)
	a2 := (&models.Astruct{
		Name: "A2",
	}).Stage(stage)
	_ = a2

	want = 2
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	a1.Stage(stage)

	want = 2
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// unstage it
	a1.Unstage(stage)

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	a1.Stage(stage)

	want = 2
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	models.SerializeStage(stage, "test.xlsx")
	models.SerializeStage2(stage, "test_withIDs.xlsx", true)
}
