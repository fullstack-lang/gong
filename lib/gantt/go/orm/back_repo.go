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

	"github.com/fullstack-lang/gong/lib/gantt/go/db"
	"github.com/fullstack-lang/gong/lib/gantt/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/gantt/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoArrow BackRepoArrowStruct

	BackRepoBar BackRepoBarStruct

	BackRepoGantt BackRepoGanttStruct

	BackRepoGroup BackRepoGroupStruct

	BackRepoLane BackRepoLaneStruct

	BackRepoLaneUse BackRepoLaneUseStruct

	BackRepoMilestone BackRepoMilestoneStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_gantt_go",
		&ArrowDB{},
		&BarDB{},
		&GanttDB{},
		&GroupDB{},
		&LaneDB{},
		&LaneUseDB{},
		&MilestoneDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoArrow = BackRepoArrowStruct{
		Map_ArrowDBID_ArrowPtr: make(map[uint]*models.Arrow, 0),
		Map_ArrowDBID_ArrowDB:  make(map[uint]*ArrowDB, 0),
		Map_ArrowPtr_ArrowDBID: make(map[*models.Arrow]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBar = BackRepoBarStruct{
		Map_BarDBID_BarPtr: make(map[uint]*models.Bar, 0),
		Map_BarDBID_BarDB:  make(map[uint]*BarDB, 0),
		Map_BarPtr_BarDBID: make(map[*models.Bar]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGantt = BackRepoGanttStruct{
		Map_GanttDBID_GanttPtr: make(map[uint]*models.Gantt, 0),
		Map_GanttDBID_GanttDB:  make(map[uint]*GanttDB, 0),
		Map_GanttPtr_GanttDBID: make(map[*models.Gantt]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroup = BackRepoGroupStruct{
		Map_GroupDBID_GroupPtr: make(map[uint]*models.Group, 0),
		Map_GroupDBID_GroupDB:  make(map[uint]*GroupDB, 0),
		Map_GroupPtr_GroupDBID: make(map[*models.Group]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLane = BackRepoLaneStruct{
		Map_LaneDBID_LanePtr: make(map[uint]*models.Lane, 0),
		Map_LaneDBID_LaneDB:  make(map[uint]*LaneDB, 0),
		Map_LanePtr_LaneDBID: make(map[*models.Lane]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLaneUse = BackRepoLaneUseStruct{
		Map_LaneUseDBID_LaneUsePtr: make(map[uint]*models.LaneUse, 0),
		Map_LaneUseDBID_LaneUseDB:  make(map[uint]*LaneUseDB, 0),
		Map_LaneUsePtr_LaneUseDBID: make(map[*models.LaneUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMilestone = BackRepoMilestoneStruct{
		Map_MilestoneDBID_MilestonePtr: make(map[uint]*models.Milestone, 0),
		Map_MilestoneDBID_MilestoneDB:  make(map[uint]*MilestoneDB, 0),
		Map_MilestonePtr_MilestoneDBID: make(map[*models.Milestone]uint, 0),

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
	backRepo.BackRepoArrow.CommitPhaseOne(stage)
	backRepo.BackRepoBar.CommitPhaseOne(stage)
	backRepo.BackRepoGantt.CommitPhaseOne(stage)
	backRepo.BackRepoGroup.CommitPhaseOne(stage)
	backRepo.BackRepoLane.CommitPhaseOne(stage)
	backRepo.BackRepoLaneUse.CommitPhaseOne(stage)
	backRepo.BackRepoMilestone.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoArrow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBar.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGantt.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLane.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLaneUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMilestone.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoArrow.CheckoutPhaseOne()
	backRepo.BackRepoBar.CheckoutPhaseOne()
	backRepo.BackRepoGantt.CheckoutPhaseOne()
	backRepo.BackRepoGroup.CheckoutPhaseOne()
	backRepo.BackRepoLane.CheckoutPhaseOne()
	backRepo.BackRepoLaneUse.CheckoutPhaseOne()
	backRepo.BackRepoMilestone.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoArrow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBar.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGantt.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLane.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLaneUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMilestone.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoArrow.Backup(dirPath)
	backRepo.BackRepoBar.Backup(dirPath)
	backRepo.BackRepoGantt.Backup(dirPath)
	backRepo.BackRepoGroup.Backup(dirPath)
	backRepo.BackRepoLane.Backup(dirPath)
	backRepo.BackRepoLaneUse.Backup(dirPath)
	backRepo.BackRepoMilestone.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoArrow.BackupXL(file)
	backRepo.BackRepoBar.BackupXL(file)
	backRepo.BackRepoGantt.BackupXL(file)
	backRepo.BackRepoGroup.BackupXL(file)
	backRepo.BackRepoLane.BackupXL(file)
	backRepo.BackRepoLaneUse.BackupXL(file)
	backRepo.BackRepoMilestone.BackupXL(file)

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
	backRepo.BackRepoArrow.RestorePhaseOne(dirPath)
	backRepo.BackRepoBar.RestorePhaseOne(dirPath)
	backRepo.BackRepoGantt.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoLane.RestorePhaseOne(dirPath)
	backRepo.BackRepoLaneUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoMilestone.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoArrow.RestorePhaseTwo()
	backRepo.BackRepoBar.RestorePhaseTwo()
	backRepo.BackRepoGantt.RestorePhaseTwo()
	backRepo.BackRepoGroup.RestorePhaseTwo()
	backRepo.BackRepoLane.RestorePhaseTwo()
	backRepo.BackRepoLaneUse.RestorePhaseTwo()
	backRepo.BackRepoMilestone.RestorePhaseTwo()

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
	backRepo.BackRepoArrow.RestoreXLPhaseOne(file)
	backRepo.BackRepoBar.RestoreXLPhaseOne(file)
	backRepo.BackRepoGantt.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoLane.RestoreXLPhaseOne(file)
	backRepo.BackRepoLaneUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoMilestone.RestoreXLPhaseOne(file)

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
