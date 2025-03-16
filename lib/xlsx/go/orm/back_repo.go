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

	"github.com/fullstack-lang/gong/lib/xlsx/go/db"
	"github.com/fullstack-lang/gong/lib/xlsx/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/xlsx/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoDisplaySelection BackRepoDisplaySelectionStruct

	BackRepoXLCell BackRepoXLCellStruct

	BackRepoXLFile BackRepoXLFileStruct

	BackRepoXLRow BackRepoXLRowStruct

	BackRepoXLSheet BackRepoXLSheetStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_xlsx_go",
		&DisplaySelectionDB{},
		&XLCellDB{},
		&XLFileDB{},
		&XLRowDB{},
		&XLSheetDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoDisplaySelection = BackRepoDisplaySelectionStruct{
		Map_DisplaySelectionDBID_DisplaySelectionPtr: make(map[uint]*models.DisplaySelection, 0),
		Map_DisplaySelectionDBID_DisplaySelectionDB:  make(map[uint]*DisplaySelectionDB, 0),
		Map_DisplaySelectionPtr_DisplaySelectionDBID: make(map[*models.DisplaySelection]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXLCell = BackRepoXLCellStruct{
		Map_XLCellDBID_XLCellPtr: make(map[uint]*models.XLCell, 0),
		Map_XLCellDBID_XLCellDB:  make(map[uint]*XLCellDB, 0),
		Map_XLCellPtr_XLCellDBID: make(map[*models.XLCell]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXLFile = BackRepoXLFileStruct{
		Map_XLFileDBID_XLFilePtr: make(map[uint]*models.XLFile, 0),
		Map_XLFileDBID_XLFileDB:  make(map[uint]*XLFileDB, 0),
		Map_XLFilePtr_XLFileDBID: make(map[*models.XLFile]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXLRow = BackRepoXLRowStruct{
		Map_XLRowDBID_XLRowPtr: make(map[uint]*models.XLRow, 0),
		Map_XLRowDBID_XLRowDB:  make(map[uint]*XLRowDB, 0),
		Map_XLRowPtr_XLRowDBID: make(map[*models.XLRow]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXLSheet = BackRepoXLSheetStruct{
		Map_XLSheetDBID_XLSheetPtr: make(map[uint]*models.XLSheet, 0),
		Map_XLSheetDBID_XLSheetDB:  make(map[uint]*XLSheetDB, 0),
		Map_XLSheetPtr_XLSheetDBID: make(map[*models.XLSheet]uint, 0),

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

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDisplaySelection.CommitPhaseOne(stage)
	backRepo.BackRepoXLCell.CommitPhaseOne(stage)
	backRepo.BackRepoXLFile.CommitPhaseOne(stage)
	backRepo.BackRepoXLRow.CommitPhaseOne(stage)
	backRepo.BackRepoXLSheet.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDisplaySelection.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXLCell.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXLFile.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXLRow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXLSheet.CommitPhaseTwo(backRepo)

	// important to release the mutex before calls to IncrementCommitFromBackNb
	// because it will block otherwise
	backRepo.rwMutex.Unlock()

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {

	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoDisplaySelection.CheckoutPhaseOne()
	backRepo.BackRepoXLCell.CheckoutPhaseOne()
	backRepo.BackRepoXLFile.CheckoutPhaseOne()
	backRepo.BackRepoXLRow.CheckoutPhaseOne()
	backRepo.BackRepoXLSheet.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoDisplaySelection.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXLCell.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXLFile.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXLRow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXLSheet.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoDisplaySelection.Backup(dirPath)
	backRepo.BackRepoXLCell.Backup(dirPath)
	backRepo.BackRepoXLFile.Backup(dirPath)
	backRepo.BackRepoXLRow.Backup(dirPath)
	backRepo.BackRepoXLSheet.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoDisplaySelection.BackupXL(file)
	backRepo.BackRepoXLCell.BackupXL(file)
	backRepo.BackRepoXLFile.BackupXL(file)
	backRepo.BackRepoXLRow.BackupXL(file)
	backRepo.BackRepoXLSheet.BackupXL(file)

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
	backRepo.BackRepoDisplaySelection.RestorePhaseOne(dirPath)
	backRepo.BackRepoXLCell.RestorePhaseOne(dirPath)
	backRepo.BackRepoXLFile.RestorePhaseOne(dirPath)
	backRepo.BackRepoXLRow.RestorePhaseOne(dirPath)
	backRepo.BackRepoXLSheet.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoDisplaySelection.RestorePhaseTwo()
	backRepo.BackRepoXLCell.RestorePhaseTwo()
	backRepo.BackRepoXLFile.RestorePhaseTwo()
	backRepo.BackRepoXLRow.RestorePhaseTwo()
	backRepo.BackRepoXLSheet.RestorePhaseTwo()

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
	backRepo.BackRepoDisplaySelection.RestoreXLPhaseOne(file)
	backRepo.BackRepoXLCell.RestoreXLPhaseOne(file)
	backRepo.BackRepoXLFile.RestoreXLPhaseOne(file)
	backRepo.BackRepoXLRow.RestoreXLPhaseOne(file)
	backRepo.BackRepoXLSheet.RestoreXLPhaseOne(file)

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
