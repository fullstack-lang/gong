package tests

import (
	"log"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

func TestXLBackup(t *testing.T) {

	// setup GORM
	orm.SetupModels(false, ":memory:")

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage().Commit()
	bclass2 := (&models.Bstruct{Name: "B2"}).Stage().Commit()

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Date:                time.Date(2021, time.February, 20, 20, 12, 4, 0, time.UTC),
		Duration1:           time.Hour + 3*time.Minute + time.Millisecond,
		Aenum:               models.ENUM_VAL1,
		Intfield:            4,
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()

	// test renumbering
	aclass1_bis := (&models.Astruct{
		Name:                "A1_bis",
		Date:                time.Date(2021, time.February, 20, 20, 12, 4, 0, time.UTC),
		Duration1:           time.Hour + 3*time.Minute + time.Millisecond,
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()
	_ = aclass1_bis
	aclass1_bis.Unstage()

	aclass2 := (&models.Astruct{
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

	for aclass := range models.Stage.Astructs {
		aclassDB := orm.BackRepo.BackRepoAstruct.GetAstructDBFromAstructPtr(aclass)
		aclassDB.CreatedAt = time.Time{}
		aclassDB.UpdatedAt = time.Time{}
	}
	for bclass := range models.Stage.Bstructs {
		bclassDB := orm.BackRepo.BackRepoBstruct.GetBstructDBFromBstructPtr(bclass)
		bclassDB.CreatedAt = time.Time{}
		bclassDB.UpdatedAt = time.Time{}
	}

	models.Stage.BackupXL("bckp-xl")
}

func TestRestoreXL(t *testing.T) {

	// setup GORM
	orm.SetupModels(false, ":memory:")

	models.Stage.Restore("bckp")

	for aclass := range models.Stage.Astructs {
		if aclass.Name == "A1" {
			log.Print("aclass.Anarrayofb[0].Name of b : ", aclass.Anarrayofb[0].Name)
			log.Print("aclass.Anarrayofb[1].Name of b : ", aclass.Anarrayofb[1].Name)
			log.Print("aclass.Anarrayofa[0].Name of a : ", aclass.Anarrayofa[0].Name)
			log.Print("aclass.Anarrayofa[1].Name of a : ", aclass.Anarrayofa[1].Name)
		}
	}

	models.Stage.Commit()

	for aclass := range models.Stage.Astructs {
		aclassDB := orm.BackRepo.BackRepoAstruct.GetAstructDBFromAstructPtr(aclass)
		aclassDB.CreatedAt = time.Time{}
		aclassDB.UpdatedAt = time.Time{}
	}
	for bclass := range models.Stage.Bstructs {
		bclassDB := orm.BackRepo.BackRepoBstruct.GetBstructDBFromBstructPtr(bclass)
		bclassDB.CreatedAt = time.Time{}
		bclassDB.UpdatedAt = time.Time{}
	}

	models.Stage.Backup("bckp-after-restore")
}
