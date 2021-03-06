// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/fullstack-lang/gong/test/go/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAstruct BackRepoAstructStruct

	BackRepoAstructBstruct2Use BackRepoAstructBstruct2UseStruct

	BackRepoAstructBstructUse BackRepoAstructBstructUseStruct

	BackRepoBstruct BackRepoBstructStruct

	BackRepoDstruct BackRepoDstructStruct

	CommitFromBackNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromBackCallback != nil {
		models.Stage.OnInitCommitFromBackCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromFrontCallback != nil {
		models.Stage.OnInitCommitFromFrontCallback.BeforeCommit(&models.Stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoAstruct.Init(db)
	backRepo.BackRepoAstructBstruct2Use.Init(db)
	backRepo.BackRepoAstructBstructUse.Init(db)
	backRepo.BackRepoBstruct.Init(db)
	backRepo.BackRepoDstruct.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAstruct.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstruct2Use.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstructUse.CommitPhaseOne(stage)
	backRepo.BackRepoBstruct.CommitPhaseOne(stage)
	backRepo.BackRepoDstruct.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAstruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstruct2Use.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstructUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBstruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDstruct.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAstruct.CheckoutPhaseOne()
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseOne()
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseOne()
	backRepo.BackRepoBstruct.CheckoutPhaseOne()
	backRepo.BackRepoDstruct.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDstruct.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitFromBackNb() uint {
	return BackRepo.GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.Backup(dirPath)
	backRepo.BackRepoAstructBstruct2Use.Backup(dirPath)
	backRepo.BackRepoAstructBstructUse.Backup(dirPath)
	backRepo.BackRepoBstruct.Backup(dirPath)
	backRepo.BackRepoDstruct.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.BackupXL(file)
	backRepo.BackRepoAstructBstruct2Use.BackupXL(file)
	backRepo.BackRepoAstructBstructUse.BackupXL(file)
	backRepo.BackRepoBstruct.BackupXL(file)
	backRepo.BackRepoDstruct.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.RestorePhaseOne(dirPath)
	backRepo.BackRepoAstructBstruct2Use.RestorePhaseOne(dirPath)
	backRepo.BackRepoAstructBstructUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoBstruct.RestorePhaseOne(dirPath)
	backRepo.BackRepoDstruct.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.RestorePhaseTwo()
	backRepo.BackRepoAstructBstruct2Use.RestorePhaseTwo()
	backRepo.BackRepoAstructBstructUse.RestorePhaseTwo()
	backRepo.BackRepoBstruct.RestorePhaseTwo()
	backRepo.BackRepoDstruct.RestorePhaseTwo()

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	models.Stage.Reset()

	// commit the cleaned stage
	models.Stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.RestoreXLPhaseOne(file)
	backRepo.BackRepoAstructBstruct2Use.RestoreXLPhaseOne(file)
	backRepo.BackRepoAstructBstructUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoBstruct.RestoreXLPhaseOne(file)
	backRepo.BackRepoDstruct.RestoreXLPhaseOne(file)

	// commit the restored stage
	models.Stage.Commit()
}
