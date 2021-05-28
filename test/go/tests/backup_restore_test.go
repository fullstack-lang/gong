package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

func TestBackup(t *testing.T) {

	// setup GORM
	db := orm.SetupModels(false, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// initiate back repo a callback functions
	orm.BackRepo.Init(db)

	bclass1 := (&models.Bclass{Name: "B1"}).Stage()
	bclass2 := (&models.Bclass{Name: "B2"}).Stage()
	_ = bclass2

	aclass1 := (&models.Aclass{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage()
	_ = aclass1

	aclass2 := (&models.Aclass{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	})
	aclass2.Stage()

	aclass1.Anotherarrayofb = append(aclass1.Anotherarrayofb, bclass1)
	aclass1.Anotherarrayofb = append(aclass1.Anotherarrayofb, bclass2)
	aclass1.Anarrayofa = append(aclass1.Anarrayofa, aclass1)
	aclass1.Anarrayofa = append(aclass1.Anarrayofa, aclass2)

	models.Stage.Commit()

	models.Stage.Backup("bckp")
}

func TestRestore(t *testing.T) {

	// setup GORM
	db := orm.SetupModels(false, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// initiate back repo a callback functions
	orm.BackRepo.Init(db)

	models.Stage.Restore("bckp")

	for aclass := range models.Stage.Aclasss {
		log.Print(aclass)
	}

	var a1 *models.Aclass
	for a := range models.Stage.Aclasss {
		a1 = a
	}
	a1.Floatfield = 17.0

	models.Stage.Commit()

	models.Stage.Backup("bckp-after-restore")
}
