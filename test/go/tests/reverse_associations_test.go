package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
)

func TestReverseAssociation(t *testing.T) {

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage()

	bclass2 := (&models.Bstruct{Name: "B2"}).Stage()
	_ = bclass2

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
		Anarrayofb: []*models.Bstruct{
			bclass1,
			bclass2,
		},
	}).Stage()
	_ = aclass1

	aclass2 := (&models.Astruct{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	})
	aclass2.Stage()

	want, got := 0, 0

	reverse_N_01_Map := models.Stage.CreateReverseMap_Astruct_Associationtob()
	_ = reverse_N_01_Map

	reverse_01_N_Map := models.Stage.CreateReverseMap_Astruct_Anarrayofb()
	_ = reverse_01_N_Map

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

}
