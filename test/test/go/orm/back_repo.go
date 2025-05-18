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

	"github.com/fullstack-lang/gong/test/test/go/db"
	"github.com/fullstack-lang/gong/test/test/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/test/test/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

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

	BackRepoF0123456789012345678901234567890 BackRepoF0123456789012345678901234567890Struct

	BackRepoGstruct BackRepoGstructStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.Stage

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.Stage, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_test_test_go",
		&AstructDB{},
		&AstructBstruct2UseDB{},
		&AstructBstructUseDB{},
		&BstructDB{},
		&DstructDB{},
		&F0123456789012345678901234567890DB{},
		&GstructDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

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
	backRepo.BackRepoF0123456789012345678901234567890 = BackRepoF0123456789012345678901234567890Struct{
		Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890Ptr: make(map[uint]*models.F0123456789012345678901234567890, 0),
		Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890DB:  make(map[uint]*F0123456789012345678901234567890DB, 0),
		Map_F0123456789012345678901234567890Ptr_F0123456789012345678901234567890DBID: make(map[*models.F0123456789012345678901234567890]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGstruct = BackRepoGstructStruct{
		Map_GstructDBID_GstructPtr: make(map[uint]*models.Gstruct, 0),
		Map_GstructDBID_GstructDB:  make(map[uint]*GstructDB, 0),
		Map_GstructPtr_GstructDBID: make(map[*models.Gstruct]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.Stage) {
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
func (backRepo *BackRepoStruct) Commit(stage *models.Stage) {

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAstruct.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstruct2Use.CommitPhaseOne(stage)
	backRepo.BackRepoAstructBstructUse.CommitPhaseOne(stage)
	backRepo.BackRepoBstruct.CommitPhaseOne(stage)
	backRepo.BackRepoDstruct.CommitPhaseOne(stage)
	backRepo.BackRepoF0123456789012345678901234567890.CommitPhaseOne(stage)
	backRepo.BackRepoGstruct.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAstruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstruct2Use.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstructUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBstruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDstruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoF0123456789012345678901234567890.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGstruct.CommitPhaseTwo(backRepo)

	// important to release the mutex before calls to IncrementCommitFromBackNb
	// because it will block otherwise
	backRepo.rwMutex.Unlock()

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.Stage) {

	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAstruct.CheckoutPhaseOne()
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseOne()
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseOne()
	backRepo.BackRepoBstruct.CheckoutPhaseOne()
	backRepo.BackRepoDstruct.CheckoutPhaseOne()
	backRepo.BackRepoF0123456789012345678901234567890.CheckoutPhaseOne()
	backRepo.BackRepoGstruct.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDstruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoF0123456789012345678901234567890.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGstruct.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.Backup(dirPath)
	backRepo.BackRepoAstructBstruct2Use.Backup(dirPath)
	backRepo.BackRepoAstructBstructUse.Backup(dirPath)
	backRepo.BackRepoBstruct.Backup(dirPath)
	backRepo.BackRepoDstruct.Backup(dirPath)
	backRepo.BackRepoF0123456789012345678901234567890.Backup(dirPath)
	backRepo.BackRepoGstruct.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.BackupXL(file)
	backRepo.BackRepoAstructBstruct2Use.BackupXL(file)
	backRepo.BackRepoAstructBstructUse.BackupXL(file)
	backRepo.BackRepoBstruct.BackupXL(file)
	backRepo.BackRepoDstruct.BackupXL(file)
	backRepo.BackRepoF0123456789012345678901234567890.BackupXL(file)
	backRepo.BackRepoGstruct.BackupXL(file)

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
func (backRepo *BackRepoStruct) Restore(stage *models.Stage, dirPath string) {
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
	backRepo.BackRepoF0123456789012345678901234567890.RestorePhaseOne(dirPath)
	backRepo.BackRepoGstruct.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAstruct.RestorePhaseTwo()
	backRepo.BackRepoAstructBstruct2Use.RestorePhaseTwo()
	backRepo.BackRepoAstructBstructUse.RestorePhaseTwo()
	backRepo.BackRepoBstruct.RestorePhaseTwo()
	backRepo.BackRepoDstruct.RestorePhaseTwo()
	backRepo.BackRepoF0123456789012345678901234567890.RestorePhaseTwo()
	backRepo.BackRepoGstruct.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.Stage, dirPath string) {

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
	backRepo.BackRepoF0123456789012345678901234567890.RestoreXLPhaseOne(file)
	backRepo.BackRepoGstruct.RestoreXLPhaseOne(file)

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
