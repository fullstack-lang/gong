package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/fullstack"
	"github.com/fullstack-lang/gong/test/go/models"
)

// TestStageCallBack
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
func TestStageCallBack(t *testing.T) {

	stage := fullstack.NewStackInstance(nil, "")

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
	}).Stage(stage)

	aclass2 := (&models.Astruct{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
		AnAstruct:           aclass1,
	})
	aclass2.Stage(stage)

	bclass2 := (&models.Bstruct{Name: "B2"}).Stage(stage)
	_ = bclass2

	stage.Commit()

	mapOfAstruct := models.GongGetMap[map[string]*models.Astruct]()
	for id, astruct := range *mapOfAstruct {
		log.Println(id, " ", astruct.Name)
	}

	setOfAstruct := models.GongGetSet[map[*models.Astruct]any]()
	for astruct := range *setOfAstruct {
		log.Println(" ", astruct.Name)
	}

	mapOfAstruct = models.GetGongstructInstancesMap[models.Astruct]()
	for id, astruct := range *mapOfAstruct {
		log.Println(id, " ", astruct.Name)
	}

	setOfAstruct = models.GetGongstructInstancesSet[models.Astruct]()
	for astruct := range *setOfAstruct {
		log.Println(" ", astruct.Name)
	}

	log.Println(models.GetAssociationName[models.Astruct]().Associationtob.Name)
	log.Println(models.GetAssociationName[models.Astruct]().Bstruct.Name)
	log.Println(models.GetAssociationName[models.Astruct]().Anotherassociationtob_2.Name)
	log.Println(models.GetAssociationName[models.Astruct]().Anarrayofb[0].Name)
	log.Println(models.GetAssociationName[models.Astruct]().Anarrayofa[0].Name)

	reverseMapAstruct_AssocationTob :=
		models.GetPointerReverseMap[
			models.Astruct, models.Bstruct](
			models.GetAssociationName[models.Astruct]().Associationtob.Name)
	for _, astruct := range reverseMapAstruct_AssocationTob[bclass1] {
		log.Println("astruct ", astruct.Name)
	}
	reverseMapAstruct_AnAstruct :=
		models.GetPointerReverseMap[
			models.Astruct, models.Astruct](
			models.GetAssociationName[models.Astruct]().AnAstruct.Name)
	for _, astruct := range reverseMapAstruct_AnAstruct[aclass2] {
		log.Println("astruct from aclass2", astruct.Name)
	}
	reverseMapAstruct_AnAstruct2 :=
		models.GetPointerReverseMap[
			models.Astruct, models.Astruct](
			models.GetAssociationName[models.Astruct]().AnAstruct.Name)
	for _, astruct := range reverseMapAstruct_AnAstruct2[aclass1] {
		log.Println("astruct from aclass1", astruct.Name)
	}
	reverseMapAstruct_AnarrayOfB :=
		models.GetSliceOfPointersReverseMap[
			models.Astruct, models.Bstruct](
			models.GetAssociationName[models.Astruct]().Anarrayofb[0].Name)

	log.Println("astruct from an array of b", reverseMapAstruct_AnarrayOfB[bclass1].Name)

	log.Println()

	want := 2
	got := len(stage.Astructs)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	log.Printf("After stage reset, ng of aclass instance %d", len(stage.Astructs))

	aclass2.Unstage(stage)

	want = 1
	got = len(stage.Astructs)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	stage.Checkout()

	// resets stage
	stage.Reset()
	log.Printf("After stage reset, ng of aclass instance %d", len(stage.Astructs))

	stage.Checkout()
	log.Printf("After stage checkout, ng of aclass instance %d", len(stage.Astructs))
	log.Printf("After stage checkout, ng of bclass instance %d", len(stage.Bstructs))

	for aclass := range stage.Astructs {
		log.Printf("After stage checkout, aclass Floatfield value %f", aclass.Floatfield)
	}

	// modify association to b on the stage
	aclass1.Associationtob = bclass2

	want = 0
	got = 0

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}
