// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gong/lib/split/go/db"
	"github.com/fullstack-lang/gong/lib/split/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/split/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAsSplit BackRepoAsSplitStruct

	BackRepoAsSplitArea BackRepoAsSplitAreaStruct

	BackRepoDoc BackRepoDocStruct

	BackRepoForm BackRepoFormStruct

	BackRepoSvg BackRepoSvgStruct

	BackRepoTable BackRepoTableStruct

	BackRepoTree BackRepoTreeStruct

	BackRepoView BackRepoViewStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_split_go",
		&AsSplitDB{},
		&AsSplitAreaDB{},
		&DocDB{},
		&FormDB{},
		&SvgDB{},
		&TableDB{},
		&TreeDB{},
		&ViewDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAsSplit = BackRepoAsSplitStruct{
		Map_AsSplitDBID_AsSplitPtr: make(map[uint]*models.AsSplit, 0),
		Map_AsSplitDBID_AsSplitDB:  make(map[uint]*AsSplitDB, 0),
		Map_AsSplitPtr_AsSplitDBID: make(map[*models.AsSplit]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAsSplitArea = BackRepoAsSplitAreaStruct{
		Map_AsSplitAreaDBID_AsSplitAreaPtr: make(map[uint]*models.AsSplitArea, 0),
		Map_AsSplitAreaDBID_AsSplitAreaDB:  make(map[uint]*AsSplitAreaDB, 0),
		Map_AsSplitAreaPtr_AsSplitAreaDBID: make(map[*models.AsSplitArea]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDoc = BackRepoDocStruct{
		Map_DocDBID_DocPtr: make(map[uint]*models.Doc, 0),
		Map_DocDBID_DocDB:  make(map[uint]*DocDB, 0),
		Map_DocPtr_DocDBID: make(map[*models.Doc]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoForm = BackRepoFormStruct{
		Map_FormDBID_FormPtr: make(map[uint]*models.Form, 0),
		Map_FormDBID_FormDB:  make(map[uint]*FormDB, 0),
		Map_FormPtr_FormDBID: make(map[*models.Form]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSvg = BackRepoSvgStruct{
		Map_SvgDBID_SvgPtr: make(map[uint]*models.Svg, 0),
		Map_SvgDBID_SvgDB:  make(map[uint]*SvgDB, 0),
		Map_SvgPtr_SvgDBID: make(map[*models.Svg]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTable = BackRepoTableStruct{
		Map_TableDBID_TablePtr: make(map[uint]*models.Table, 0),
		Map_TableDBID_TableDB:  make(map[uint]*TableDB, 0),
		Map_TablePtr_TableDBID: make(map[*models.Table]uint, 0),

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
	backRepo.BackRepoView = BackRepoViewStruct{
		Map_ViewDBID_ViewPtr: make(map[uint]*models.View, 0),
		Map_ViewDBID_ViewDB:  make(map[uint]*ViewDB, 0),
		Map_ViewPtr_ViewDBID: make(map[*models.View]uint, 0),

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

	backRepo.broadcastNbCommitToBack()

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

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAsSplit.CommitPhaseOne(stage)
	backRepo.BackRepoAsSplitArea.CommitPhaseOne(stage)
	backRepo.BackRepoDoc.CommitPhaseOne(stage)
	backRepo.BackRepoForm.CommitPhaseOne(stage)
	backRepo.BackRepoSvg.CommitPhaseOne(stage)
	backRepo.BackRepoTable.CommitPhaseOne(stage)
	backRepo.BackRepoTree.CommitPhaseOne(stage)
	backRepo.BackRepoView.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAsSplit.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAsSplitArea.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDoc.CommitPhaseTwo(backRepo)
	backRepo.BackRepoForm.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSvg.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTable.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTree.CommitPhaseTwo(backRepo)
	backRepo.BackRepoView.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAsSplit.CheckoutPhaseOne()
	backRepo.BackRepoAsSplitArea.CheckoutPhaseOne()
	backRepo.BackRepoDoc.CheckoutPhaseOne()
	backRepo.BackRepoForm.CheckoutPhaseOne()
	backRepo.BackRepoSvg.CheckoutPhaseOne()
	backRepo.BackRepoTable.CheckoutPhaseOne()
	backRepo.BackRepoTree.CheckoutPhaseOne()
	backRepo.BackRepoView.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAsSplit.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAsSplitArea.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDoc.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoForm.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSvg.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTable.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTree.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoView.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAsSplit.Backup(dirPath)
	backRepo.BackRepoAsSplitArea.Backup(dirPath)
	backRepo.BackRepoDoc.Backup(dirPath)
	backRepo.BackRepoForm.Backup(dirPath)
	backRepo.BackRepoSvg.Backup(dirPath)
	backRepo.BackRepoTable.Backup(dirPath)
	backRepo.BackRepoTree.Backup(dirPath)
	backRepo.BackRepoView.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAsSplit.BackupXL(file)
	backRepo.BackRepoAsSplitArea.BackupXL(file)
	backRepo.BackRepoDoc.BackupXL(file)
	backRepo.BackRepoForm.BackupXL(file)
	backRepo.BackRepoSvg.BackupXL(file)
	backRepo.BackRepoTable.BackupXL(file)
	backRepo.BackRepoTree.BackupXL(file)
	backRepo.BackRepoView.BackupXL(file)

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
	backRepo.BackRepoAsSplit.RestorePhaseOne(dirPath)
	backRepo.BackRepoAsSplitArea.RestorePhaseOne(dirPath)
	backRepo.BackRepoDoc.RestorePhaseOne(dirPath)
	backRepo.BackRepoForm.RestorePhaseOne(dirPath)
	backRepo.BackRepoSvg.RestorePhaseOne(dirPath)
	backRepo.BackRepoTable.RestorePhaseOne(dirPath)
	backRepo.BackRepoTree.RestorePhaseOne(dirPath)
	backRepo.BackRepoView.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAsSplit.RestorePhaseTwo()
	backRepo.BackRepoAsSplitArea.RestorePhaseTwo()
	backRepo.BackRepoDoc.RestorePhaseTwo()
	backRepo.BackRepoForm.RestorePhaseTwo()
	backRepo.BackRepoSvg.RestorePhaseTwo()
	backRepo.BackRepoTable.RestorePhaseTwo()
	backRepo.BackRepoTree.RestorePhaseTwo()
	backRepo.BackRepoView.RestorePhaseTwo()

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
	backRepo.BackRepoAsSplit.RestoreXLPhaseOne(file)
	backRepo.BackRepoAsSplitArea.RestoreXLPhaseOne(file)
	backRepo.BackRepoDoc.RestoreXLPhaseOne(file)
	backRepo.BackRepoForm.RestoreXLPhaseOne(file)
	backRepo.BackRepoSvg.RestoreXLPhaseOne(file)
	backRepo.BackRepoTable.RestoreXLPhaseOne(file)
	backRepo.BackRepoTree.RestoreXLPhaseOne(file)
	backRepo.BackRepoView.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
