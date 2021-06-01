package tests

import (
	"log"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

func TestBackup(t *testing.T) {

	// setup GORM
	db := orm.SetupModels(false, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// initiate back repo a callback functions
	orm.BackRepo.Init(db)

	bclass1 := (&models.Bclass{Name: "B1"}).Stage().Commit()
	bclass2 := (&models.Bclass{Name: "B2"}).Stage().Commit()

	aclass1 := (&models.Aclass{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()

	// test renumbering
	aclass1_bis := (&models.Aclass{
		Name:                "A1_bis",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()
	_ = aclass1_bis
	aclass1_bis.Unstage()

	aclass2 := (&models.Aclass{
		Name:                "A2",
		Floatfield:          10.77,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	})
	aclass2.Stage().Commit()

	// order is important, since it is stored in the index
	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass2)
	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass1)

	aclass1.Anotherarrayofb = append(aclass1.Anotherarrayofb, bclass1)
	aclass1.Anotherarrayofb = append(aclass1.Anotherarrayofb, bclass2)

	aclass1.Anarrayofa = append(aclass1.Anarrayofa, aclass1)
	aclass1.Anarrayofa = append(aclass1.Anarrayofa, aclass2)

	models.Stage.Commit()

	for aclass := range models.Stage.Aclasss {
		aclassDB := orm.BackRepo.BackRepoAclass.GetAclassDBFromAclassPtr(aclass)
		aclassDB.CreatedAt = time.Time{}
		aclassDB.UpdatedAt = time.Time{}
	}
	for bclass := range models.Stage.Bclasss {
		bclassDB := orm.BackRepo.BackRepoBclass.GetBclassDBFromBclassPtr(bclass)
		bclassDB.CreatedAt = time.Time{}
		bclassDB.UpdatedAt = time.Time{}
	}

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
		if aclass.Name == "A1" {
			log.Print("aclass.Anarrayofb[0].Name of b : ", aclass.Anarrayofb[0].Name)
			log.Print("aclass.Anarrayofb[1].Name of b : ", aclass.Anarrayofb[1].Name)
			log.Print("aclass.Anarrayofa[0].Name of a : ", aclass.Anarrayofa[0].Name)
			log.Print("aclass.Anarrayofa[1].Name of a : ", aclass.Anarrayofa[1].Name)
		}
	}

	models.Stage.Commit()

	for aclass := range models.Stage.Aclasss {
		aclassDB := orm.BackRepo.BackRepoAclass.GetAclassDBFromAclassPtr(aclass)
		aclassDB.CreatedAt = time.Time{}
		aclassDB.UpdatedAt = time.Time{}
	}
	for bclass := range models.Stage.Bclasss {
		bclassDB := orm.BackRepo.BackRepoBclass.GetBclassDBFromBclassPtr(bclass)
		bclassDB.CreatedAt = time.Time{}
		bclassDB.UpdatedAt = time.Time{}
	}

	models.Stage.Backup("bckp-after-restore")
}
