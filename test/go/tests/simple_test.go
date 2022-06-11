package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

// TestStageCallBack
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
//
//
func TestStageCallBack(t *testing.T) {

	// setup GORM
	db := orm.SetupModels(false, ":memory:")

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
	}).Stage()

	aclass2 := (&models.Astruct{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
		AnAstruct:           aclass1,
	})
	aclass2.Stage()

	bclass2 := (&models.Bstruct{Name: "B2"}).Stage()
	_ = bclass2

	models.Stage.Commit()

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
	got := len(models.Stage.Astructs)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	log.Printf("After models.Stage reset, ng of aclass instance %d", len(models.Stage.Astructs))

	aclass2.Unstage()

	want = 1
	got = len(models.Stage.Astructs)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	var aclasss []orm.AstructDB
	{
		query := db.Find(&aclasss)
		if query.Error != nil {
			t.Errorf("Find error %s", query.Error.Error())
			return
		}
	}

	want = 2
	got = len(aclasss)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	// after the commmit, there should only one aclassDB left
	models.Stage.Commit()
	{
		query := db.Find(&aclasss)
		if query.Error != nil {
			t.Errorf("Find error %s", query.Error.Error())
			return
		}
	}

	want = 1
	got = len(aclasss)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	// get the AstructDB
	aclassDB := &(orm.AstructDB{})
	if err := db.First(aclassDB).Error; err != nil {
		t.Errorf("Bad Query")
	}

	log.Printf("aclass Associationtob " + aclass1.Associationtob.Name)
	log.Printf("aclass float field %f ", aclass1.Floatfield)

	// more direct
	aclassDB.Floatfield_Data.Float64 = 10.5

	models.Stage.Checkout()

	// resets stage
	models.Stage.Reset()
	log.Printf("After models.Stage reset, ng of aclass instance %d", len(models.Stage.Astructs))

	models.Stage.Checkout()
	log.Printf("After models.Stage checkout, ng of aclass instance %d", len(models.Stage.Astructs))
	log.Printf("After models.Stage checkout, ng of bclass instance %d", len(models.Stage.Bstructs))

	for aclass := range models.Stage.Astructs {
		log.Printf("After models.Stage checkout, aclass Floatfield value %f", aclass.Floatfield)
	}

	// modify association to b on the stage
	aclass1.Associationtob = bclass2

	want = 0
	got = 0

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}
