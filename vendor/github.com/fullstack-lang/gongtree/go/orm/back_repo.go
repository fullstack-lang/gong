package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongtree/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoButton BackRepoButtonStruct

	BackRepoNode BackRepoNodeStruct

	BackRepoTree BackRepoTreeStruct

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
		&ButtonDB{},
		&NodeDB{},
		&TreeDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoButton = BackRepoButtonStruct{
		Map_ButtonDBID_ButtonPtr: make(map[uint]*models.Button, 0),
		Map_ButtonDBID_ButtonDB:  make(map[uint]*ButtonDB, 0),
		Map_ButtonPtr_ButtonDBID: make(map[*models.Button]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNode = BackRepoNodeStruct{
		Map_NodeDBID_NodePtr: make(map[uint]*models.Node, 0),
		Map_NodeDBID_NodeDB:  make(map[uint]*NodeDB, 0),
		Map_NodePtr_NodeDBID: make(map[*models.Node]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTree = BackRepoTreeStruct{
		Map_TreeDBID_TreePtr: make(map[uint]*models.Tree, 0),
		Map_TreeDBID_TreeDB:  make(map[uint]*TreeDB, 0),
		Map_TreePtr_TreeDBID: make(map[*models.Tree]uint, 0),

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
	backRepo.BackRepoButton.CommitPhaseOne(stage)
	backRepo.BackRepoNode.CommitPhaseOne(stage)
	backRepo.BackRepoTree.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNode.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTree.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoButton.CheckoutPhaseOne()
	backRepo.BackRepoNode.CheckoutPhaseOne()
	backRepo.BackRepoTree.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNode.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTree.CheckoutPhaseTwo(backRepo)
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
	backRepo.BackRepoButton.Backup(dirPath)
	backRepo.BackRepoNode.Backup(dirPath)
	backRepo.BackRepoTree.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoButton.BackupXL(file)
	backRepo.BackRepoNode.BackupXL(file)
	backRepo.BackRepoTree.BackupXL(file)

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
	backRepo.BackRepoButton.RestorePhaseOne(dirPath)
	backRepo.BackRepoNode.RestorePhaseOne(dirPath)
	backRepo.BackRepoTree.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoButton.RestorePhaseTwo()
	backRepo.BackRepoNode.RestorePhaseTwo()
	backRepo.BackRepoTree.RestorePhaseTwo()

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
	backRepo.BackRepoButton.RestoreXLPhaseOne(file)
	backRepo.BackRepoNode.RestoreXLPhaseOne(file)
	backRepo.BackRepoTree.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}
