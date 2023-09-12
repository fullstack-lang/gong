// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gong/test/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAstruct BackRepoAstructStruct

	BackRepoAstructBstruct2Use BackRepoAstructBstruct2UseStruct

	BackRepoAstructBstructUse BackRepoAstructBstructUseStruct

	BackRepoBstruct BackRepoBstructStruct

	BackRepoDstruct BackRepoDstructStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&AstructDB{},
		&AstructBstruct2UseDB{},
		&AstructBstructUseDB{},
		&BstructDB{},
		&DstructDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAstruct = BackRepoAstructStruct{
		Map_AstructDBID_AstructPtr: make(map[uint]*models.Astruct, 0),
		Map_AstructDBID_AstructDB:  make(map[uint]*AstructDB, 0),
		Map_AstructPtr_AstructDBID: make(map[*models.Astruct]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAstructBstruct2Use = BackRepoAstructBstruct2UseStruct{
		Map_AstructBstruct2UseDBID_AstructBstruct2UsePtr: make(map[uint]*models.AstructBstruct2Use, 0),
		Map_AstructBstruct2UseDBID_AstructBstruct2UseDB:  make(map[uint]*AstructBstruct2UseDB, 0),
		Map_AstructBstruct2UsePtr_AstructBstruct2UseDBID: make(map[*models.AstructBstruct2Use]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAstructBstructUse = BackRepoAstructBstructUseStruct{
		Map_AstructBstructUseDBID_AstructBstructUsePtr: make(map[uint]*models.AstructBstructUse, 0),
		Map_AstructBstructUseDBID_AstructBstructUseDB:  make(map[uint]*AstructBstructUseDB, 0),
		Map_AstructBstructUsePtr_AstructBstructUseDBID: make(map[*models.AstructBstructUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBstruct = BackRepoBstructStruct{
		Map_BstructDBID_BstructPtr: make(map[uint]*models.Bstruct, 0),
		Map_BstructDBID_BstructDB:  make(map[uint]*BstructDB, 0),
		Map_BstructPtr_BstructDBID: make(map[*models.Bstruct]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDstruct = BackRepoDstructStruct{
		Map_DstructDBID_DstructPtr: make(map[uint]*models.Dstruct, 0),
		Map_DstructDBID_DstructDB:  make(map[uint]*DstructDB, 0),
		Map_DstructPtr_DstructDBID: make(map[*models.Dstruct]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAstruct.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstruct2Use.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstructUse.CommitPhaseOne(stage)
	backRepo.BackRepoBstruct.CommitPhaseOne(stage)
	backRepo.BackRepoDstruct.CommitPhaseOne(stage)

	// backRepo.BackRepoAstruct.ResetReversePointers(backRepo)
	// backRepo.BackRepoBstruct.ResetReversePointers(backRepo)

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

	// backRepo.BackRepoBstruct.ResetReversePointers(backRepo)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDstruct.CheckoutPhaseTwo(backRepo)
}

var _backRepo *BackRepoStruct

var once sync.Once

func GetDefaultBackRepo() *BackRepoStruct {
	once.Do(func() {
		_backRepo = NewBackRepo(models.GetDefaultStage(), "")
	})
	return _backRepo
}

func GetLastCommitFromBackNb() uint {
	return GetDefaultBackRepo().GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return GetDefaultBackRepo().GetLastPushFromFrontNb()
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
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

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

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

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
	backRepo.stage.Commit()
}
