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

	"github.com/fullstack-lang/gongtree/go/db"
	"github.com/fullstack-lang/gongtree/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongtree/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoButton BackRepoButtonStruct

	BackRepoNode BackRepoNodeStruct

	BackRepoSVGIcon BackRepoSVGIconStruct

	BackRepoTree BackRepoTreeStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongtree_go",
		&ButtonDB{},
		&NodeDB{},
		&SVGIconDB{},
		&TreeDB{},
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
	backRepo.BackRepoNode = BackRepoNodeStruct{
		Map_NodeDBID_NodePtr: make(map[uint]*models.Node, 0),
		Map_NodeDBID_NodeDB:  make(map[uint]*NodeDB, 0),
		Map_NodePtr_NodeDBID: make(map[*models.Node]uint, 0),

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
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoButton.CommitPhaseOne(stage)
	backRepo.BackRepoNode.CommitPhaseOne(stage)
	backRepo.BackRepoSVGIcon.CommitPhaseOne(stage)
	backRepo.BackRepoTree.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNode.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSVGIcon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTree.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoButton.CheckoutPhaseOne()
	backRepo.BackRepoNode.CheckoutPhaseOne()
	backRepo.BackRepoSVGIcon.CheckoutPhaseOne()
	backRepo.BackRepoTree.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoButton.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNode.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSVGIcon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTree.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoButton.Backup(dirPath)
	backRepo.BackRepoNode.Backup(dirPath)
	backRepo.BackRepoSVGIcon.Backup(dirPath)
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
	backRepo.BackRepoSVGIcon.BackupXL(file)
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
	backRepo.BackRepoSVGIcon.RestorePhaseOne(dirPath)
	backRepo.BackRepoTree.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoButton.RestorePhaseTwo()
	backRepo.BackRepoNode.RestorePhaseTwo()
	backRepo.BackRepoSVGIcon.RestorePhaseTwo()
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
	backRepo.BackRepoSVGIcon.RestoreXLPhaseOne(file)
	backRepo.BackRepoTree.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.rwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.rwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()
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
	backRepoStruct.rwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.rwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}