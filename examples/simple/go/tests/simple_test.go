package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/examples/simple/go/models"
	"github.com/fullstack-lang/gong/examples/simple/go/orm"
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
	db.DB().SetMaxOpenConns(1)

	// initiate back repo a callback functions
	orm.BackRepo.Init(db)

	bclass1 := (&models.Bclass{Name: "B1"}).Stage()

	aclass1 := (&models.Aclass{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage()

	aclass2 := (&models.Aclass{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	})
	aclass2.Stage()

	bclass2 := (&models.Bclass{Name: "B2"}).Stage()
	_ = bclass2

	models.Stage.Commit()

	want := 2
	got := len(models.Stage.Aclasss)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	log.Printf("After models.Stage reset, ng of aclass instance %d", len(models.Stage.Aclasss))

	aclass2.Unstage()

	want = 1
	got = len(models.Stage.Aclasss)

	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}

	var aclasss []orm.AclassDB
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

	// get the AclassDB
	aclassDB := &(orm.AclassDB{})
	if err := db.First(aclassDB).Error; err != nil {
		t.Errorf("Bad Query")
	}

	log.Printf("aclass Associationtob " + aclass1.Associationtob.Name)
	log.Printf("aclass float field %f ", aclass1.Floatfield)

	// more direct
	aclassDB.Floatfield = 10.5

	models.Stage.Checkout()

	// resets stage
	models.Stage.Reset()
	log.Printf("After models.Stage reset, ng of aclass instance %d", len(models.Stage.Aclasss))

	models.Stage.Checkout()
	log.Printf("After models.Stage checkout, ng of aclass instance %d", len(models.Stage.Aclasss))
	log.Printf("After models.Stage checkout, ng of bclass instance %d", len(models.Stage.Bclasss))

	for aclass := range models.Stage.Aclasss {
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