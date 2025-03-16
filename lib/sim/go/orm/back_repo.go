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

	"github.com/fullstack-lang/gong/lib/sim/go/db"
	"github.com/fullstack-lang/gong/lib/sim/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/sim/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCommand BackRepoCommandStruct

	BackRepoDummyAgent BackRepoDummyAgentStruct

	BackRepoEngine BackRepoEngineStruct

	BackRepoEvent BackRepoEventStruct

	BackRepoStatus BackRepoStatusStruct

	BackRepoUpdateState BackRepoUpdateStateStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_sim_go",
		&CommandDB{},
		&DummyAgentDB{},
		&EngineDB{},
		&EventDB{},
		&StatusDB{},
		&UpdateStateDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCommand = BackRepoCommandStruct{
		Map_CommandDBID_CommandPtr: make(map[uint]*models.Command, 0),
		Map_CommandDBID_CommandDB:  make(map[uint]*CommandDB, 0),
		Map_CommandPtr_CommandDBID: make(map[*models.Command]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDummyAgent = BackRepoDummyAgentStruct{
		Map_DummyAgentDBID_DummyAgentPtr: make(map[uint]*models.DummyAgent, 0),
		Map_DummyAgentDBID_DummyAgentDB:  make(map[uint]*DummyAgentDB, 0),
		Map_DummyAgentPtr_DummyAgentDBID: make(map[*models.DummyAgent]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEngine = BackRepoEngineStruct{
		Map_EngineDBID_EnginePtr: make(map[uint]*models.Engine, 0),
		Map_EngineDBID_EngineDB:  make(map[uint]*EngineDB, 0),
		Map_EnginePtr_EngineDBID: make(map[*models.Engine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEvent = BackRepoEventStruct{
		Map_EventDBID_EventPtr: make(map[uint]*models.Event, 0),
		Map_EventDBID_EventDB:  make(map[uint]*EventDB, 0),
		Map_EventPtr_EventDBID: make(map[*models.Event]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStatus = BackRepoStatusStruct{
		Map_StatusDBID_StatusPtr: make(map[uint]*models.Status, 0),
		Map_StatusDBID_StatusDB:  make(map[uint]*StatusDB, 0),
		Map_StatusPtr_StatusDBID: make(map[*models.Status]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUpdateState = BackRepoUpdateStateStruct{
		Map_UpdateStateDBID_UpdateStatePtr: make(map[uint]*models.UpdateState, 0),
		Map_UpdateStateDBID_UpdateStateDB:  make(map[uint]*UpdateStateDB, 0),
		Map_UpdateStatePtr_UpdateStateDBID: make(map[*models.UpdateState]uint, 0),

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
	backRepo.BackRepoCommand.CommitPhaseOne(stage)
	backRepo.BackRepoDummyAgent.CommitPhaseOne(stage)
	backRepo.BackRepoEngine.CommitPhaseOne(stage)
	backRepo.BackRepoEvent.CommitPhaseOne(stage)
	backRepo.BackRepoStatus.CommitPhaseOne(stage)
	backRepo.BackRepoUpdateState.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCommand.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDummyAgent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEngine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEvent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStatus.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUpdateState.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoCommand.CheckoutPhaseOne()
	backRepo.BackRepoDummyAgent.CheckoutPhaseOne()
	backRepo.BackRepoEngine.CheckoutPhaseOne()
	backRepo.BackRepoEvent.CheckoutPhaseOne()
	backRepo.BackRepoStatus.CheckoutPhaseOne()
	backRepo.BackRepoUpdateState.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCommand.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDummyAgent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEngine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEvent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStatus.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUpdateState.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCommand.Backup(dirPath)
	backRepo.BackRepoDummyAgent.Backup(dirPath)
	backRepo.BackRepoEngine.Backup(dirPath)
	backRepo.BackRepoEvent.Backup(dirPath)
	backRepo.BackRepoStatus.Backup(dirPath)
	backRepo.BackRepoUpdateState.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCommand.BackupXL(file)
	backRepo.BackRepoDummyAgent.BackupXL(file)
	backRepo.BackRepoEngine.BackupXL(file)
	backRepo.BackRepoEvent.BackupXL(file)
	backRepo.BackRepoStatus.BackupXL(file)
	backRepo.BackRepoUpdateState.BackupXL(file)

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
	backRepo.BackRepoCommand.RestorePhaseOne(dirPath)
	backRepo.BackRepoDummyAgent.RestorePhaseOne(dirPath)
	backRepo.BackRepoEngine.RestorePhaseOne(dirPath)
	backRepo.BackRepoEvent.RestorePhaseOne(dirPath)
	backRepo.BackRepoStatus.RestorePhaseOne(dirPath)
	backRepo.BackRepoUpdateState.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCommand.RestorePhaseTwo()
	backRepo.BackRepoDummyAgent.RestorePhaseTwo()
	backRepo.BackRepoEngine.RestorePhaseTwo()
	backRepo.BackRepoEvent.RestorePhaseTwo()
	backRepo.BackRepoStatus.RestorePhaseTwo()
	backRepo.BackRepoUpdateState.RestorePhaseTwo()

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
	backRepo.BackRepoCommand.RestoreXLPhaseOne(file)
	backRepo.BackRepoDummyAgent.RestoreXLPhaseOne(file)
	backRepo.BackRepoEngine.RestoreXLPhaseOne(file)
	backRepo.BackRepoEvent.RestoreXLPhaseOne(file)
	backRepo.BackRepoStatus.RestoreXLPhaseOne(file)
	backRepo.BackRepoUpdateState.RestoreXLPhaseOne(file)

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
