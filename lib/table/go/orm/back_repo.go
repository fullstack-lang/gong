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

	"github.com/fullstack-lang/gong/lib/table/go/db"
	"github.com/fullstack-lang/gong/lib/table/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/table/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoButton BackRepoButtonStruct

	BackRepoCell BackRepoCellStruct

	BackRepoCellBoolean BackRepoCellBooleanStruct

	BackRepoCellFloat64 BackRepoCellFloat64Struct

	BackRepoCellIcon BackRepoCellIconStruct

	BackRepoCellInt BackRepoCellIntStruct

	BackRepoCellString BackRepoCellStringStruct

	BackRepoDisplayedColumn BackRepoDisplayedColumnStruct

	BackRepoRow BackRepoRowStruct

	BackRepoSVGIcon BackRepoSVGIconStruct

	BackRepoTable BackRepoTableStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.Stage

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers        []chan int
}

func NewBackRepo(stage *models.Stage, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_table_go",
		&ButtonDB{},
		&CellDB{},
		&CellBooleanDB{},
		&CellFloat64DB{},
		&CellIconDB{},
		&CellIntDB{},
		&CellStringDB{},
		&DisplayedColumnDB{},
		&RowDB{},
		&SVGIconDB{},
		&TableDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoButton = BackRepoButtonStruct{
		Map_ButtonDBID_ButtonPtr: make(map[uint]*models.Button, 0),
		Map_ButtonDBID_ButtonDB:  make(map[uint]*ButtonDB, 0),
		Map_ButtonPtr_ButtonDBID: make(map[*models.Button]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCell = BackRepoCellStruct{
		Map_CellDBID_CellPtr: make(map[uint]*models.Cell, 0),
		Map_CellDBID_CellDB:  make(map[uint]*CellDB, 0),
		Map_CellPtr_CellDBID: make(map[*models.Cell]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellBoolean = BackRepoCellBooleanStruct{
		Map_CellBooleanDBID_CellBooleanPtr: make(map[uint]*models.CellBoolean, 0),
		Map_CellBooleanDBID_CellBooleanDB:  make(map[uint]*CellBooleanDB, 0),
		Map_CellBooleanPtr_CellBooleanDBID: make(map[*models.CellBoolean]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellFloat64 = BackRepoCellFloat64Struct{
		Map_CellFloat64DBID_CellFloat64Ptr: make(map[uint]*models.CellFloat64, 0),
		Map_CellFloat64DBID_CellFloat64DB:  make(map[uint]*CellFloat64DB, 0),
		Map_CellFloat64Ptr_CellFloat64DBID: make(map[*models.CellFloat64]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellIcon = BackRepoCellIconStruct{
		Map_CellIconDBID_CellIconPtr: make(map[uint]*models.CellIcon, 0),
		Map_CellIconDBID_CellIconDB:  make(map[uint]*CellIconDB, 0),
		Map_CellIconPtr_CellIconDBID: make(map[*models.CellIcon]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellInt = BackRepoCellIntStruct{
		Map_CellIntDBID_CellIntPtr: make(map[uint]*models.CellInt, 0),
		Map_CellIntDBID_CellIntDB:  make(map[uint]*CellIntDB, 0),
		Map_CellIntPtr_CellIntDBID: make(map[*models.CellInt]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCellString = BackRepoCellStringStruct{
		Map_CellStringDBID_CellStringPtr: make(map[uint]*models.CellString, 0),
		Map_CellStringDBID_CellStringDB:  make(map[uint]*CellStringDB, 0),
		Map_CellStringPtr_CellStringDBID: make(map[*models.CellString]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDisplayedColumn = BackRepoDisplayedColumnStruct{
		Map_DisplayedColumnDBID_DisplayedColumnPtr: make(map[uint]*models.DisplayedColumn, 0),
		Map_DisplayedColumnDBID_DisplayedColumnDB:  make(map[uint]*DisplayedColumnDB, 0),
		Map_DisplayedColumnPtr_DisplayedColumnDBID: make(map[*models.DisplayedColumn]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRow = BackRepoRowStruct{
		Map_RowDBID_RowPtr: make(map[uint]*models.Row, 0),
		Map_RowDBID_RowDB:  make(map[uint]*RowDB, 0),
		Map_RowPtr_RowDBID: make(map[*models.Row]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSVGIcon = BackRepoSVGIconStruct{
		Map_SVGIconDBID_SVGIconPtr: make(map[uint]*models.SVGIcon, 0),
		Map_SVGIconDBID_SVGIconDB:  make(map[uint]*SVGIconDB, 0),
		Map_SVGIconPtr_SVGIconDBID: make(map[*models.SVGIcon]uint, 0),

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
	backRepo.BackRepoButton.CommitPhaseOne(stage)
	backRepo.BackRepoCell.CommitPhaseOne(stage)
	backRepo.BackRepoCellBoolean.CommitPhaseOne(stage)
	backRepo.BackRepoCellFloat64.CommitPhaseOne(stage)
	backRepo.BackRepoCellIcon.CommitPhaseOne(stage)
	backRepo.BackRepoCellInt.CommitPhaseOne(stage)
	backRepo.BackRepoCellString.CommitPhaseOne(stage)
	backRepo.BackRepoDisplayedColumn.CommitPhaseOne(stage)
	backRepo.BackRepoRow.CommitPhaseOne(stage)
	backRepo.BackRepoSVGIcon.CommitPhaseOne(stage)
	backRepo.BackRepoTable.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCell.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellBoolean.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellFloat64.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellIcon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellInt.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCellString.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDisplayedColumn.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSVGIcon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTable.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoButton.CheckoutPhaseOne()
	backRepo.BackRepoCell.CheckoutPhaseOne()
	backRepo.BackRepoCellBoolean.CheckoutPhaseOne()
	backRepo.BackRepoCellFloat64.CheckoutPhaseOne()
	backRepo.BackRepoCellIcon.CheckoutPhaseOne()
	backRepo.BackRepoCellInt.CheckoutPhaseOne()
	backRepo.BackRepoCellString.CheckoutPhaseOne()
	backRepo.BackRepoDisplayedColumn.CheckoutPhaseOne()
	backRepo.BackRepoRow.CheckoutPhaseOne()
	backRepo.BackRepoSVGIcon.CheckoutPhaseOne()
	backRepo.BackRepoTable.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCell.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellBoolean.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellFloat64.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellIcon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellInt.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCellString.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDisplayedColumn.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSVGIcon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTable.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoButton.Backup(dirPath)
	backRepo.BackRepoCell.Backup(dirPath)
	backRepo.BackRepoCellBoolean.Backup(dirPath)
	backRepo.BackRepoCellFloat64.Backup(dirPath)
	backRepo.BackRepoCellIcon.Backup(dirPath)
	backRepo.BackRepoCellInt.Backup(dirPath)
	backRepo.BackRepoCellString.Backup(dirPath)
	backRepo.BackRepoDisplayedColumn.Backup(dirPath)
	backRepo.BackRepoRow.Backup(dirPath)
	backRepo.BackRepoSVGIcon.Backup(dirPath)
	backRepo.BackRepoTable.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoButton.BackupXL(file)
	backRepo.BackRepoCell.BackupXL(file)
	backRepo.BackRepoCellBoolean.BackupXL(file)
	backRepo.BackRepoCellFloat64.BackupXL(file)
	backRepo.BackRepoCellIcon.BackupXL(file)
	backRepo.BackRepoCellInt.BackupXL(file)
	backRepo.BackRepoCellString.BackupXL(file)
	backRepo.BackRepoDisplayedColumn.BackupXL(file)
	backRepo.BackRepoRow.BackupXL(file)
	backRepo.BackRepoSVGIcon.BackupXL(file)
	backRepo.BackRepoTable.BackupXL(file)

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
	backRepo.BackRepoButton.RestorePhaseOne(dirPath)
	backRepo.BackRepoCell.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellBoolean.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellFloat64.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellIcon.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellInt.RestorePhaseOne(dirPath)
	backRepo.BackRepoCellString.RestorePhaseOne(dirPath)
	backRepo.BackRepoDisplayedColumn.RestorePhaseOne(dirPath)
	backRepo.BackRepoRow.RestorePhaseOne(dirPath)
	backRepo.BackRepoSVGIcon.RestorePhaseOne(dirPath)
	backRepo.BackRepoTable.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoButton.RestorePhaseTwo()
	backRepo.BackRepoCell.RestorePhaseTwo()
	backRepo.BackRepoCellBoolean.RestorePhaseTwo()
	backRepo.BackRepoCellFloat64.RestorePhaseTwo()
	backRepo.BackRepoCellIcon.RestorePhaseTwo()
	backRepo.BackRepoCellInt.RestorePhaseTwo()
	backRepo.BackRepoCellString.RestorePhaseTwo()
	backRepo.BackRepoDisplayedColumn.RestorePhaseTwo()
	backRepo.BackRepoRow.RestorePhaseTwo()
	backRepo.BackRepoSVGIcon.RestorePhaseTwo()
	backRepo.BackRepoTable.RestorePhaseTwo()

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
	backRepo.BackRepoButton.RestoreXLPhaseOne(file)
	backRepo.BackRepoCell.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellBoolean.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellFloat64.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellIcon.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellInt.RestoreXLPhaseOne(file)
	backRepo.BackRepoCellString.RestoreXLPhaseOne(file)
	backRepo.BackRepoDisplayedColumn.RestoreXLPhaseOne(file)
	backRepo.BackRepoRow.RestoreXLPhaseOne(file)
	backRepo.BackRepoSVGIcon.RestoreXLPhaseOne(file)
	backRepo.BackRepoTable.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

// SubscribeToCommitNb is a function that is called by the front to be notified of new commits
// It returns a channel that will be closed when the connection is closed
//
// The channel is used to propagate the commit number to the front
// This commit number is used by the front to validate the tree after commit.
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

	// if len(subscribers) == 0 {
	// 	log.Println(backRepoStruct.stage.GetType(), backRepoStruct.stage.GetName(), "no subsribers to broadcast to")
	// }

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
