package tests

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"

	"github.com/tealeg/xlsx/v3"
)

func TestBackupTest(t *testing.T) {

	// setup GORM
	orm.SetupModels(false, "../../../test/test.db")

	models.Stage.Checkout()

	models.Stage.Backup("bckp-test")
}

func TestRestoreTest(t *testing.T) {

	// setup GORM
	orm.SetupModels(false, "../../../test/test.db")

	models.Stage.Restore("bckp-test")

	models.Stage.Commit()
}

func CreateTestStage() {

	// setup GORM
	orm.SetupModels(false, ":memory:")

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage().Commit()
	bclass2 := (&models.Bstruct{Name: "B2"}).Stage().Commit()

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage().Commit()

	// test renumbering
	aclass1_bis := (&models.Astruct{
		Name:                "A1_bis",
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

}

func TestBackup(t *testing.T) {

	CreateTestStage()

	models.Stage.Backup("bckp")
	models.Stage.BackupXL("bckp")
}

func TestRestore(t *testing.T) {

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

func TestNewXLBackup(t *testing.T) {

	CreateTestStage()

	bckpFile := xlsx.NewFile()

	sheet, _ := bckpFile.AddSheet("Astruct")

	row := sheet.AddRow()

	set := models.GetGongstructInstancesSet[models.Astruct]()
	for astruct := range *set {
		cell := row.AddCell()
		cell.Value = astruct.Name
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	bckpFile.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(".", "new_bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}

}
