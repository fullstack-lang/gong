package stage_test

import (
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
)

// TestStageCount
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
func TestStageCount(t *testing.T) {

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage()

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

	stage := models.Stage
	_ = stage

	// empty stage
	want := 0
	got := len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage one instance
	aclass1.Stage()

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	aclass1.Stage()

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// unstage it
	aclass1.Unstage()

	want = 0
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

	// stage it again
	aclass1.Stage()

	want = 1
	got = len(stage.Astructs)
	if got != want {
		t.Fatal("Wanted ", want, "got", got)
	}

}
