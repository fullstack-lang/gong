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

	"github.com/fullstack-lang/gong/lib/doc2/go/db"
	"github.com/fullstack-lang/gong/lib/doc2/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/doc2/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAttributeShape BackRepoAttributeShapeStruct

	BackRepoClassdiagram BackRepoClassdiagramStruct

	BackRepoDiagramPackage BackRepoDiagramPackageStruct

	BackRepoGongEnumShape BackRepoGongEnumShapeStruct

	BackRepoGongEnumValueShape BackRepoGongEnumValueShapeStruct

	BackRepoGongNoteLinkShape BackRepoGongNoteLinkShapeStruct

	BackRepoGongNoteShape BackRepoGongNoteShapeStruct

	BackRepoGongStructShape BackRepoGongStructShapeStruct

	BackRepoLinkShape BackRepoLinkShapeStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_doc2_go",
		&AttributeShapeDB{},
		&ClassdiagramDB{},
		&DiagramPackageDB{},
		&GongEnumShapeDB{},
		&GongEnumValueShapeDB{},
		&GongNoteLinkShapeDB{},
		&GongNoteShapeDB{},
		&GongStructShapeDB{},
		&LinkShapeDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAttributeShape = BackRepoAttributeShapeStruct{
		Map_AttributeShapeDBID_AttributeShapePtr: make(map[uint]*models.AttributeShape, 0),
		Map_AttributeShapeDBID_AttributeShapeDB:  make(map[uint]*AttributeShapeDB, 0),
		Map_AttributeShapePtr_AttributeShapeDBID: make(map[*models.AttributeShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoClassdiagram = BackRepoClassdiagramStruct{
		Map_ClassdiagramDBID_ClassdiagramPtr: make(map[uint]*models.Classdiagram, 0),
		Map_ClassdiagramDBID_ClassdiagramDB:  make(map[uint]*ClassdiagramDB, 0),
		Map_ClassdiagramPtr_ClassdiagramDBID: make(map[*models.Classdiagram]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDiagramPackage = BackRepoDiagramPackageStruct{
		Map_DiagramPackageDBID_DiagramPackagePtr: make(map[uint]*models.DiagramPackage, 0),
		Map_DiagramPackageDBID_DiagramPackageDB:  make(map[uint]*DiagramPackageDB, 0),
		Map_DiagramPackagePtr_DiagramPackageDBID: make(map[*models.DiagramPackage]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongEnumShape = BackRepoGongEnumShapeStruct{
		Map_GongEnumShapeDBID_GongEnumShapePtr: make(map[uint]*models.GongEnumShape, 0),
		Map_GongEnumShapeDBID_GongEnumShapeDB:  make(map[uint]*GongEnumShapeDB, 0),
		Map_GongEnumShapePtr_GongEnumShapeDBID: make(map[*models.GongEnumShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongEnumValueShape = BackRepoGongEnumValueShapeStruct{
		Map_GongEnumValueShapeDBID_GongEnumValueShapePtr: make(map[uint]*models.GongEnumValueShape, 0),
		Map_GongEnumValueShapeDBID_GongEnumValueShapeDB:  make(map[uint]*GongEnumValueShapeDB, 0),
		Map_GongEnumValueShapePtr_GongEnumValueShapeDBID: make(map[*models.GongEnumValueShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongNoteLinkShape = BackRepoGongNoteLinkShapeStruct{
		Map_GongNoteLinkShapeDBID_GongNoteLinkShapePtr: make(map[uint]*models.GongNoteLinkShape, 0),
		Map_GongNoteLinkShapeDBID_GongNoteLinkShapeDB:  make(map[uint]*GongNoteLinkShapeDB, 0),
		Map_GongNoteLinkShapePtr_GongNoteLinkShapeDBID: make(map[*models.GongNoteLinkShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongNoteShape = BackRepoGongNoteShapeStruct{
		Map_GongNoteShapeDBID_GongNoteShapePtr: make(map[uint]*models.GongNoteShape, 0),
		Map_GongNoteShapeDBID_GongNoteShapeDB:  make(map[uint]*GongNoteShapeDB, 0),
		Map_GongNoteShapePtr_GongNoteShapeDBID: make(map[*models.GongNoteShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongStructShape = BackRepoGongStructShapeStruct{
		Map_GongStructShapeDBID_GongStructShapePtr: make(map[uint]*models.GongStructShape, 0),
		Map_GongStructShapeDBID_GongStructShapeDB:  make(map[uint]*GongStructShapeDB, 0),
		Map_GongStructShapePtr_GongStructShapeDBID: make(map[*models.GongStructShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLinkShape = BackRepoLinkShapeStruct{
		Map_LinkShapeDBID_LinkShapePtr: make(map[uint]*models.LinkShape, 0),
		Map_LinkShapeDBID_LinkShapeDB:  make(map[uint]*LinkShapeDB, 0),
		Map_LinkShapePtr_LinkShapeDBID: make(map[*models.LinkShape]uint, 0),

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
	backRepo.BackRepoAttributeShape.CommitPhaseOne(stage)
	backRepo.BackRepoClassdiagram.CommitPhaseOne(stage)
	backRepo.BackRepoDiagramPackage.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnumShape.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnumValueShape.CommitPhaseOne(stage)
	backRepo.BackRepoGongNoteLinkShape.CommitPhaseOne(stage)
	backRepo.BackRepoGongNoteShape.CommitPhaseOne(stage)
	backRepo.BackRepoGongStructShape.CommitPhaseOne(stage)
	backRepo.BackRepoLinkShape.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAttributeShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoClassdiagram.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDiagramPackage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValueShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongNoteLinkShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongNoteShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongStructShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLinkShape.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoAttributeShape.CheckoutPhaseOne()
	backRepo.BackRepoClassdiagram.CheckoutPhaseOne()
	backRepo.BackRepoDiagramPackage.CheckoutPhaseOne()
	backRepo.BackRepoGongEnumShape.CheckoutPhaseOne()
	backRepo.BackRepoGongEnumValueShape.CheckoutPhaseOne()
	backRepo.BackRepoGongNoteLinkShape.CheckoutPhaseOne()
	backRepo.BackRepoGongNoteShape.CheckoutPhaseOne()
	backRepo.BackRepoGongStructShape.CheckoutPhaseOne()
	backRepo.BackRepoLinkShape.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAttributeShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoClassdiagram.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDiagramPackage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValueShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongNoteLinkShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongNoteShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongStructShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLinkShape.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAttributeShape.Backup(dirPath)
	backRepo.BackRepoClassdiagram.Backup(dirPath)
	backRepo.BackRepoDiagramPackage.Backup(dirPath)
	backRepo.BackRepoGongEnumShape.Backup(dirPath)
	backRepo.BackRepoGongEnumValueShape.Backup(dirPath)
	backRepo.BackRepoGongNoteLinkShape.Backup(dirPath)
	backRepo.BackRepoGongNoteShape.Backup(dirPath)
	backRepo.BackRepoGongStructShape.Backup(dirPath)
	backRepo.BackRepoLinkShape.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAttributeShape.BackupXL(file)
	backRepo.BackRepoClassdiagram.BackupXL(file)
	backRepo.BackRepoDiagramPackage.BackupXL(file)
	backRepo.BackRepoGongEnumShape.BackupXL(file)
	backRepo.BackRepoGongEnumValueShape.BackupXL(file)
	backRepo.BackRepoGongNoteLinkShape.BackupXL(file)
	backRepo.BackRepoGongNoteShape.BackupXL(file)
	backRepo.BackRepoGongStructShape.BackupXL(file)
	backRepo.BackRepoLinkShape.BackupXL(file)

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
	backRepo.BackRepoAttributeShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoClassdiagram.RestorePhaseOne(dirPath)
	backRepo.BackRepoDiagramPackage.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnumShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnumValueShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongNoteLinkShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongNoteShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongStructShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoLinkShape.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAttributeShape.RestorePhaseTwo()
	backRepo.BackRepoClassdiagram.RestorePhaseTwo()
	backRepo.BackRepoDiagramPackage.RestorePhaseTwo()
	backRepo.BackRepoGongEnumShape.RestorePhaseTwo()
	backRepo.BackRepoGongEnumValueShape.RestorePhaseTwo()
	backRepo.BackRepoGongNoteLinkShape.RestorePhaseTwo()
	backRepo.BackRepoGongNoteShape.RestorePhaseTwo()
	backRepo.BackRepoGongStructShape.RestorePhaseTwo()
	backRepo.BackRepoLinkShape.RestorePhaseTwo()

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
	backRepo.BackRepoAttributeShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoClassdiagram.RestoreXLPhaseOne(file)
	backRepo.BackRepoDiagramPackage.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnumShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnumValueShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongNoteLinkShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongNoteShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongStructShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoLinkShape.RestoreXLPhaseOne(file)

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
