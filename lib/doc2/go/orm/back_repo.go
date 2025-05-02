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

	BackRepoGongEnumValueEntry BackRepoGongEnumValueEntryStruct

	BackRepoGongStructShape BackRepoGongStructShapeStruct

	BackRepoLinkShape BackRepoLinkShapeStruct

	BackRepoNoteShape BackRepoNoteShapeStruct

	BackRepoNoteShapeLink BackRepoNoteShapeLinkStruct

	BackRepoPosition BackRepoPositionStruct

	BackRepoUmlState BackRepoUmlStateStruct

	BackRepoUmlsc BackRepoUmlscStruct

	BackRepoVertice BackRepoVerticeStruct

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
		&GongEnumValueEntryDB{},
		&GongStructShapeDB{},
		&LinkShapeDB{},
		&NoteShapeDB{},
		&NoteShapeLinkDB{},
		&PositionDB{},
		&UmlStateDB{},
		&UmlscDB{},
		&VerticeDB{},
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
	backRepo.BackRepoGongEnumValueEntry = BackRepoGongEnumValueEntryStruct{
		Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr: make(map[uint]*models.GongEnumValueEntry, 0),
		Map_GongEnumValueEntryDBID_GongEnumValueEntryDB:  make(map[uint]*GongEnumValueEntryDB, 0),
		Map_GongEnumValueEntryPtr_GongEnumValueEntryDBID: make(map[*models.GongEnumValueEntry]uint, 0),

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
	backRepo.BackRepoNoteShape = BackRepoNoteShapeStruct{
		Map_NoteShapeDBID_NoteShapePtr: make(map[uint]*models.NoteShape, 0),
		Map_NoteShapeDBID_NoteShapeDB:  make(map[uint]*NoteShapeDB, 0),
		Map_NoteShapePtr_NoteShapeDBID: make(map[*models.NoteShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNoteShapeLink = BackRepoNoteShapeLinkStruct{
		Map_NoteShapeLinkDBID_NoteShapeLinkPtr: make(map[uint]*models.NoteShapeLink, 0),
		Map_NoteShapeLinkDBID_NoteShapeLinkDB:  make(map[uint]*NoteShapeLinkDB, 0),
		Map_NoteShapeLinkPtr_NoteShapeLinkDBID: make(map[*models.NoteShapeLink]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPosition = BackRepoPositionStruct{
		Map_PositionDBID_PositionPtr: make(map[uint]*models.Position, 0),
		Map_PositionDBID_PositionDB:  make(map[uint]*PositionDB, 0),
		Map_PositionPtr_PositionDBID: make(map[*models.Position]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUmlState = BackRepoUmlStateStruct{
		Map_UmlStateDBID_UmlStatePtr: make(map[uint]*models.UmlState, 0),
		Map_UmlStateDBID_UmlStateDB:  make(map[uint]*UmlStateDB, 0),
		Map_UmlStatePtr_UmlStateDBID: make(map[*models.UmlState]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUmlsc = BackRepoUmlscStruct{
		Map_UmlscDBID_UmlscPtr: make(map[uint]*models.Umlsc, 0),
		Map_UmlscDBID_UmlscDB:  make(map[uint]*UmlscDB, 0),
		Map_UmlscPtr_UmlscDBID: make(map[*models.Umlsc]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVertice = BackRepoVerticeStruct{
		Map_VerticeDBID_VerticePtr: make(map[uint]*models.Vertice, 0),
		Map_VerticeDBID_VerticeDB:  make(map[uint]*VerticeDB, 0),
		Map_VerticePtr_VerticeDBID: make(map[*models.Vertice]uint, 0),

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
	backRepo.BackRepoGongEnumValueEntry.CommitPhaseOne(stage)
	backRepo.BackRepoGongStructShape.CommitPhaseOne(stage)
	backRepo.BackRepoLinkShape.CommitPhaseOne(stage)
	backRepo.BackRepoNoteShape.CommitPhaseOne(stage)
	backRepo.BackRepoNoteShapeLink.CommitPhaseOne(stage)
	backRepo.BackRepoPosition.CommitPhaseOne(stage)
	backRepo.BackRepoUmlState.CommitPhaseOne(stage)
	backRepo.BackRepoUmlsc.CommitPhaseOne(stage)
	backRepo.BackRepoVertice.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAttributeShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoClassdiagram.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDiagramPackage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValueEntry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongStructShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLinkShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNoteShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNoteShapeLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPosition.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUmlState.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUmlsc.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVertice.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoGongEnumValueEntry.CheckoutPhaseOne()
	backRepo.BackRepoGongStructShape.CheckoutPhaseOne()
	backRepo.BackRepoLinkShape.CheckoutPhaseOne()
	backRepo.BackRepoNoteShape.CheckoutPhaseOne()
	backRepo.BackRepoNoteShapeLink.CheckoutPhaseOne()
	backRepo.BackRepoPosition.CheckoutPhaseOne()
	backRepo.BackRepoUmlState.CheckoutPhaseOne()
	backRepo.BackRepoUmlsc.CheckoutPhaseOne()
	backRepo.BackRepoVertice.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAttributeShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoClassdiagram.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDiagramPackage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValueEntry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongStructShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLinkShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNoteShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNoteShapeLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPosition.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUmlState.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUmlsc.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVertice.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAttributeShape.Backup(dirPath)
	backRepo.BackRepoClassdiagram.Backup(dirPath)
	backRepo.BackRepoDiagramPackage.Backup(dirPath)
	backRepo.BackRepoGongEnumShape.Backup(dirPath)
	backRepo.BackRepoGongEnumValueEntry.Backup(dirPath)
	backRepo.BackRepoGongStructShape.Backup(dirPath)
	backRepo.BackRepoLinkShape.Backup(dirPath)
	backRepo.BackRepoNoteShape.Backup(dirPath)
	backRepo.BackRepoNoteShapeLink.Backup(dirPath)
	backRepo.BackRepoPosition.Backup(dirPath)
	backRepo.BackRepoUmlState.Backup(dirPath)
	backRepo.BackRepoUmlsc.Backup(dirPath)
	backRepo.BackRepoVertice.Backup(dirPath)
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
	backRepo.BackRepoGongEnumValueEntry.BackupXL(file)
	backRepo.BackRepoGongStructShape.BackupXL(file)
	backRepo.BackRepoLinkShape.BackupXL(file)
	backRepo.BackRepoNoteShape.BackupXL(file)
	backRepo.BackRepoNoteShapeLink.BackupXL(file)
	backRepo.BackRepoPosition.BackupXL(file)
	backRepo.BackRepoUmlState.BackupXL(file)
	backRepo.BackRepoUmlsc.BackupXL(file)
	backRepo.BackRepoVertice.BackupXL(file)

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
	backRepo.BackRepoGongEnumValueEntry.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongStructShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoLinkShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoNoteShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoNoteShapeLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoPosition.RestorePhaseOne(dirPath)
	backRepo.BackRepoUmlState.RestorePhaseOne(dirPath)
	backRepo.BackRepoUmlsc.RestorePhaseOne(dirPath)
	backRepo.BackRepoVertice.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAttributeShape.RestorePhaseTwo()
	backRepo.BackRepoClassdiagram.RestorePhaseTwo()
	backRepo.BackRepoDiagramPackage.RestorePhaseTwo()
	backRepo.BackRepoGongEnumShape.RestorePhaseTwo()
	backRepo.BackRepoGongEnumValueEntry.RestorePhaseTwo()
	backRepo.BackRepoGongStructShape.RestorePhaseTwo()
	backRepo.BackRepoLinkShape.RestorePhaseTwo()
	backRepo.BackRepoNoteShape.RestorePhaseTwo()
	backRepo.BackRepoNoteShapeLink.RestorePhaseTwo()
	backRepo.BackRepoPosition.RestorePhaseTwo()
	backRepo.BackRepoUmlState.RestorePhaseTwo()
	backRepo.BackRepoUmlsc.RestorePhaseTwo()
	backRepo.BackRepoVertice.RestorePhaseTwo()

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
	backRepo.BackRepoGongEnumValueEntry.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongStructShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoLinkShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoNoteShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoNoteShapeLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoPosition.RestoreXLPhaseOne(file)
	backRepo.BackRepoUmlState.RestoreXLPhaseOne(file)
	backRepo.BackRepoUmlsc.RestoreXLPhaseOne(file)
	backRepo.BackRepoVertice.RestoreXLPhaseOne(file)

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
