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

	"github.com/fullstack-lang/gong/test/statemachines/go/db"
	"github.com/fullstack-lang/gong/test/statemachines/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/test/statemachines/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoArchitecture BackRepoArchitectureStruct

	BackRepoDiagram BackRepoDiagramStruct

	BackRepoKill BackRepoKillStruct

	BackRepoMessage BackRepoMessageStruct

	BackRepoMessageType BackRepoMessageTypeStruct

	BackRepoObject BackRepoObjectStruct

	BackRepoRole BackRepoRoleStruct

	BackRepoState BackRepoStateStruct

	BackRepoStateMachine BackRepoStateMachineStruct

	BackRepoStateShape BackRepoStateShapeStruct

	BackRepoTransition BackRepoTransitionStruct

	BackRepoTransition_Shape BackRepoTransition_ShapeStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_test_statemachines_go",
		&ArchitectureDB{},
		&DiagramDB{},
		&KillDB{},
		&MessageDB{},
		&MessageTypeDB{},
		&ObjectDB{},
		&RoleDB{},
		&StateDB{},
		&StateMachineDB{},
		&StateShapeDB{},
		&TransitionDB{},
		&Transition_ShapeDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoArchitecture = BackRepoArchitectureStruct{
		Map_ArchitectureDBID_ArchitecturePtr: make(map[uint]*models.Architecture, 0),
		Map_ArchitectureDBID_ArchitectureDB:  make(map[uint]*ArchitectureDB, 0),
		Map_ArchitecturePtr_ArchitectureDBID: make(map[*models.Architecture]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDiagram = BackRepoDiagramStruct{
		Map_DiagramDBID_DiagramPtr: make(map[uint]*models.Diagram, 0),
		Map_DiagramDBID_DiagramDB:  make(map[uint]*DiagramDB, 0),
		Map_DiagramPtr_DiagramDBID: make(map[*models.Diagram]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKill = BackRepoKillStruct{
		Map_KillDBID_KillPtr: make(map[uint]*models.Kill, 0),
		Map_KillDBID_KillDB:  make(map[uint]*KillDB, 0),
		Map_KillPtr_KillDBID: make(map[*models.Kill]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMessage = BackRepoMessageStruct{
		Map_MessageDBID_MessagePtr: make(map[uint]*models.Message, 0),
		Map_MessageDBID_MessageDB:  make(map[uint]*MessageDB, 0),
		Map_MessagePtr_MessageDBID: make(map[*models.Message]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMessageType = BackRepoMessageTypeStruct{
		Map_MessageTypeDBID_MessageTypePtr: make(map[uint]*models.MessageType, 0),
		Map_MessageTypeDBID_MessageTypeDB:  make(map[uint]*MessageTypeDB, 0),
		Map_MessageTypePtr_MessageTypeDBID: make(map[*models.MessageType]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoObject = BackRepoObjectStruct{
		Map_ObjectDBID_ObjectPtr: make(map[uint]*models.Object, 0),
		Map_ObjectDBID_ObjectDB:  make(map[uint]*ObjectDB, 0),
		Map_ObjectPtr_ObjectDBID: make(map[*models.Object]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRole = BackRepoRoleStruct{
		Map_RoleDBID_RolePtr: make(map[uint]*models.Role, 0),
		Map_RoleDBID_RoleDB:  make(map[uint]*RoleDB, 0),
		Map_RolePtr_RoleDBID: make(map[*models.Role]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoState = BackRepoStateStruct{
		Map_StateDBID_StatePtr: make(map[uint]*models.State, 0),
		Map_StateDBID_StateDB:  make(map[uint]*StateDB, 0),
		Map_StatePtr_StateDBID: make(map[*models.State]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStateMachine = BackRepoStateMachineStruct{
		Map_StateMachineDBID_StateMachinePtr: make(map[uint]*models.StateMachine, 0),
		Map_StateMachineDBID_StateMachineDB:  make(map[uint]*StateMachineDB, 0),
		Map_StateMachinePtr_StateMachineDBID: make(map[*models.StateMachine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStateShape = BackRepoStateShapeStruct{
		Map_StateShapeDBID_StateShapePtr: make(map[uint]*models.StateShape, 0),
		Map_StateShapeDBID_StateShapeDB:  make(map[uint]*StateShapeDB, 0),
		Map_StateShapePtr_StateShapeDBID: make(map[*models.StateShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTransition = BackRepoTransitionStruct{
		Map_TransitionDBID_TransitionPtr: make(map[uint]*models.Transition, 0),
		Map_TransitionDBID_TransitionDB:  make(map[uint]*TransitionDB, 0),
		Map_TransitionPtr_TransitionDBID: make(map[*models.Transition]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTransition_Shape = BackRepoTransition_ShapeStruct{
		Map_Transition_ShapeDBID_Transition_ShapePtr: make(map[uint]*models.Transition_Shape, 0),
		Map_Transition_ShapeDBID_Transition_ShapeDB:  make(map[uint]*Transition_ShapeDB, 0),
		Map_Transition_ShapePtr_Transition_ShapeDBID: make(map[*models.Transition_Shape]uint, 0),

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
	backRepo.BackRepoArchitecture.CommitPhaseOne(stage)
	backRepo.BackRepoDiagram.CommitPhaseOne(stage)
	backRepo.BackRepoKill.CommitPhaseOne(stage)
	backRepo.BackRepoMessage.CommitPhaseOne(stage)
	backRepo.BackRepoMessageType.CommitPhaseOne(stage)
	backRepo.BackRepoObject.CommitPhaseOne(stage)
	backRepo.BackRepoRole.CommitPhaseOne(stage)
	backRepo.BackRepoState.CommitPhaseOne(stage)
	backRepo.BackRepoStateMachine.CommitPhaseOne(stage)
	backRepo.BackRepoStateShape.CommitPhaseOne(stage)
	backRepo.BackRepoTransition.CommitPhaseOne(stage)
	backRepo.BackRepoTransition_Shape.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoArchitecture.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDiagram.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKill.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMessage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMessageType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoObject.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRole.CommitPhaseTwo(backRepo)
	backRepo.BackRepoState.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStateMachine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStateShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTransition.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTransition_Shape.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoArchitecture.CheckoutPhaseOne()
	backRepo.BackRepoDiagram.CheckoutPhaseOne()
	backRepo.BackRepoKill.CheckoutPhaseOne()
	backRepo.BackRepoMessage.CheckoutPhaseOne()
	backRepo.BackRepoMessageType.CheckoutPhaseOne()
	backRepo.BackRepoObject.CheckoutPhaseOne()
	backRepo.BackRepoRole.CheckoutPhaseOne()
	backRepo.BackRepoState.CheckoutPhaseOne()
	backRepo.BackRepoStateMachine.CheckoutPhaseOne()
	backRepo.BackRepoStateShape.CheckoutPhaseOne()
	backRepo.BackRepoTransition.CheckoutPhaseOne()
	backRepo.BackRepoTransition_Shape.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoArchitecture.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDiagram.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKill.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMessage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMessageType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoObject.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRole.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoState.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStateMachine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStateShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTransition.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTransition_Shape.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoArchitecture.Backup(dirPath)
	backRepo.BackRepoDiagram.Backup(dirPath)
	backRepo.BackRepoKill.Backup(dirPath)
	backRepo.BackRepoMessage.Backup(dirPath)
	backRepo.BackRepoMessageType.Backup(dirPath)
	backRepo.BackRepoObject.Backup(dirPath)
	backRepo.BackRepoRole.Backup(dirPath)
	backRepo.BackRepoState.Backup(dirPath)
	backRepo.BackRepoStateMachine.Backup(dirPath)
	backRepo.BackRepoStateShape.Backup(dirPath)
	backRepo.BackRepoTransition.Backup(dirPath)
	backRepo.BackRepoTransition_Shape.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoArchitecture.BackupXL(file)
	backRepo.BackRepoDiagram.BackupXL(file)
	backRepo.BackRepoKill.BackupXL(file)
	backRepo.BackRepoMessage.BackupXL(file)
	backRepo.BackRepoMessageType.BackupXL(file)
	backRepo.BackRepoObject.BackupXL(file)
	backRepo.BackRepoRole.BackupXL(file)
	backRepo.BackRepoState.BackupXL(file)
	backRepo.BackRepoStateMachine.BackupXL(file)
	backRepo.BackRepoStateShape.BackupXL(file)
	backRepo.BackRepoTransition.BackupXL(file)
	backRepo.BackRepoTransition_Shape.BackupXL(file)

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
	backRepo.BackRepoArchitecture.RestorePhaseOne(dirPath)
	backRepo.BackRepoDiagram.RestorePhaseOne(dirPath)
	backRepo.BackRepoKill.RestorePhaseOne(dirPath)
	backRepo.BackRepoMessage.RestorePhaseOne(dirPath)
	backRepo.BackRepoMessageType.RestorePhaseOne(dirPath)
	backRepo.BackRepoObject.RestorePhaseOne(dirPath)
	backRepo.BackRepoRole.RestorePhaseOne(dirPath)
	backRepo.BackRepoState.RestorePhaseOne(dirPath)
	backRepo.BackRepoStateMachine.RestorePhaseOne(dirPath)
	backRepo.BackRepoStateShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoTransition.RestorePhaseOne(dirPath)
	backRepo.BackRepoTransition_Shape.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoArchitecture.RestorePhaseTwo()
	backRepo.BackRepoDiagram.RestorePhaseTwo()
	backRepo.BackRepoKill.RestorePhaseTwo()
	backRepo.BackRepoMessage.RestorePhaseTwo()
	backRepo.BackRepoMessageType.RestorePhaseTwo()
	backRepo.BackRepoObject.RestorePhaseTwo()
	backRepo.BackRepoRole.RestorePhaseTwo()
	backRepo.BackRepoState.RestorePhaseTwo()
	backRepo.BackRepoStateMachine.RestorePhaseTwo()
	backRepo.BackRepoStateShape.RestorePhaseTwo()
	backRepo.BackRepoTransition.RestorePhaseTwo()
	backRepo.BackRepoTransition_Shape.RestorePhaseTwo()

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
	backRepo.BackRepoArchitecture.RestoreXLPhaseOne(file)
	backRepo.BackRepoDiagram.RestoreXLPhaseOne(file)
	backRepo.BackRepoKill.RestoreXLPhaseOne(file)
	backRepo.BackRepoMessage.RestoreXLPhaseOne(file)
	backRepo.BackRepoMessageType.RestoreXLPhaseOne(file)
	backRepo.BackRepoObject.RestoreXLPhaseOne(file)
	backRepo.BackRepoRole.RestoreXLPhaseOne(file)
	backRepo.BackRepoState.RestoreXLPhaseOne(file)
	backRepo.BackRepoStateMachine.RestoreXLPhaseOne(file)
	backRepo.BackRepoStateShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoTransition.RestoreXLPhaseOne(file)
	backRepo.BackRepoTransition_Shape.RestoreXLPhaseOne(file)

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
