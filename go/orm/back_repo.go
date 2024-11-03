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

	"github.com/fullstack-lang/gong/go/db"
	"github.com/fullstack-lang/gong/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoGongBasicField BackRepoGongBasicFieldStruct

	BackRepoGongEnum BackRepoGongEnumStruct

	BackRepoGongEnumValue BackRepoGongEnumValueStruct

	BackRepoGongLink BackRepoGongLinkStruct

	BackRepoGongNote BackRepoGongNoteStruct

	BackRepoGongStruct BackRepoGongStructStruct

	BackRepoGongTimeField BackRepoGongTimeFieldStruct

	BackRepoMeta BackRepoMetaStruct

	BackRepoMetaReference BackRepoMetaReferenceStruct

	BackRepoModelPkg BackRepoModelPkgStruct

	BackRepoPointerToGongStructField BackRepoPointerToGongStructFieldStruct

	BackRepoSliceOfPointerToGongStructField BackRepoSliceOfPointerToGongStructFieldStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_go",
		&GongBasicFieldDB{},
		&GongEnumDB{},
		&GongEnumValueDB{},
		&GongLinkDB{},
		&GongNoteDB{},
		&GongStructDB{},
		&GongTimeFieldDB{},
		&MetaDB{},
		&MetaReferenceDB{},
		&ModelPkgDB{},
		&PointerToGongStructFieldDB{},
		&SliceOfPointerToGongStructFieldDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoGongBasicField = BackRepoGongBasicFieldStruct{
		Map_GongBasicFieldDBID_GongBasicFieldPtr: make(map[uint]*models.GongBasicField, 0),
		Map_GongBasicFieldDBID_GongBasicFieldDB:  make(map[uint]*GongBasicFieldDB, 0),
		Map_GongBasicFieldPtr_GongBasicFieldDBID: make(map[*models.GongBasicField]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongEnum = BackRepoGongEnumStruct{
		Map_GongEnumDBID_GongEnumPtr: make(map[uint]*models.GongEnum, 0),
		Map_GongEnumDBID_GongEnumDB:  make(map[uint]*GongEnumDB, 0),
		Map_GongEnumPtr_GongEnumDBID: make(map[*models.GongEnum]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongEnumValue = BackRepoGongEnumValueStruct{
		Map_GongEnumValueDBID_GongEnumValuePtr: make(map[uint]*models.GongEnumValue, 0),
		Map_GongEnumValueDBID_GongEnumValueDB:  make(map[uint]*GongEnumValueDB, 0),
		Map_GongEnumValuePtr_GongEnumValueDBID: make(map[*models.GongEnumValue]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongLink = BackRepoGongLinkStruct{
		Map_GongLinkDBID_GongLinkPtr: make(map[uint]*models.GongLink, 0),
		Map_GongLinkDBID_GongLinkDB:  make(map[uint]*GongLinkDB, 0),
		Map_GongLinkPtr_GongLinkDBID: make(map[*models.GongLink]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongNote = BackRepoGongNoteStruct{
		Map_GongNoteDBID_GongNotePtr: make(map[uint]*models.GongNote, 0),
		Map_GongNoteDBID_GongNoteDB:  make(map[uint]*GongNoteDB, 0),
		Map_GongNotePtr_GongNoteDBID: make(map[*models.GongNote]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongStruct = BackRepoGongStructStruct{
		Map_GongStructDBID_GongStructPtr: make(map[uint]*models.GongStruct, 0),
		Map_GongStructDBID_GongStructDB:  make(map[uint]*GongStructDB, 0),
		Map_GongStructPtr_GongStructDBID: make(map[*models.GongStruct]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGongTimeField = BackRepoGongTimeFieldStruct{
		Map_GongTimeFieldDBID_GongTimeFieldPtr: make(map[uint]*models.GongTimeField, 0),
		Map_GongTimeFieldDBID_GongTimeFieldDB:  make(map[uint]*GongTimeFieldDB, 0),
		Map_GongTimeFieldPtr_GongTimeFieldDBID: make(map[*models.GongTimeField]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeta = BackRepoMetaStruct{
		Map_MetaDBID_MetaPtr: make(map[uint]*models.Meta, 0),
		Map_MetaDBID_MetaDB:  make(map[uint]*MetaDB, 0),
		Map_MetaPtr_MetaDBID: make(map[*models.Meta]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetaReference = BackRepoMetaReferenceStruct{
		Map_MetaReferenceDBID_MetaReferencePtr: make(map[uint]*models.MetaReference, 0),
		Map_MetaReferenceDBID_MetaReferenceDB:  make(map[uint]*MetaReferenceDB, 0),
		Map_MetaReferencePtr_MetaReferenceDBID: make(map[*models.MetaReference]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoModelPkg = BackRepoModelPkgStruct{
		Map_ModelPkgDBID_ModelPkgPtr: make(map[uint]*models.ModelPkg, 0),
		Map_ModelPkgDBID_ModelPkgDB:  make(map[uint]*ModelPkgDB, 0),
		Map_ModelPkgPtr_ModelPkgDBID: make(map[*models.ModelPkg]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPointerToGongStructField = BackRepoPointerToGongStructFieldStruct{
		Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr: make(map[uint]*models.PointerToGongStructField, 0),
		Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB:  make(map[uint]*PointerToGongStructFieldDB, 0),
		Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID: make(map[*models.PointerToGongStructField]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSliceOfPointerToGongStructField = BackRepoSliceOfPointerToGongStructFieldStruct{
		Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr: make(map[uint]*models.SliceOfPointerToGongStructField, 0),
		Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB:  make(map[uint]*SliceOfPointerToGongStructFieldDB, 0),
		Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID: make(map[*models.SliceOfPointerToGongStructField]uint, 0),

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
	backRepo.BackRepoGongBasicField.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnum.CommitPhaseOne(stage)
	backRepo.BackRepoGongEnumValue.CommitPhaseOne(stage)
	backRepo.BackRepoGongLink.CommitPhaseOne(stage)
	backRepo.BackRepoGongNote.CommitPhaseOne(stage)
	backRepo.BackRepoGongStruct.CommitPhaseOne(stage)
	backRepo.BackRepoGongTimeField.CommitPhaseOne(stage)
	backRepo.BackRepoMeta.CommitPhaseOne(stage)
	backRepo.BackRepoMetaReference.CommitPhaseOne(stage)
	backRepo.BackRepoModelPkg.CommitPhaseOne(stage)
	backRepo.BackRepoPointerToGongStructField.CommitPhaseOne(stage)
	backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoGongBasicField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnum.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValue.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongNote.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongStruct.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGongTimeField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeta.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetaReference.CommitPhaseTwo(backRepo)
	backRepo.BackRepoModelPkg.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPointerToGongStructField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoGongBasicField.CheckoutPhaseOne()
	backRepo.BackRepoGongEnum.CheckoutPhaseOne()
	backRepo.BackRepoGongEnumValue.CheckoutPhaseOne()
	backRepo.BackRepoGongLink.CheckoutPhaseOne()
	backRepo.BackRepoGongNote.CheckoutPhaseOne()
	backRepo.BackRepoGongStruct.CheckoutPhaseOne()
	backRepo.BackRepoGongTimeField.CheckoutPhaseOne()
	backRepo.BackRepoMeta.CheckoutPhaseOne()
	backRepo.BackRepoMetaReference.CheckoutPhaseOne()
	backRepo.BackRepoModelPkg.CheckoutPhaseOne()
	backRepo.BackRepoPointerToGongStructField.CheckoutPhaseOne()
	backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoGongBasicField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnum.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongEnumValue.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongNote.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongStruct.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGongTimeField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeta.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetaReference.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoModelPkg.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPointerToGongStructField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.Backup(dirPath)
	backRepo.BackRepoGongEnum.Backup(dirPath)
	backRepo.BackRepoGongEnumValue.Backup(dirPath)
	backRepo.BackRepoGongLink.Backup(dirPath)
	backRepo.BackRepoGongNote.Backup(dirPath)
	backRepo.BackRepoGongStruct.Backup(dirPath)
	backRepo.BackRepoGongTimeField.Backup(dirPath)
	backRepo.BackRepoMeta.Backup(dirPath)
	backRepo.BackRepoMetaReference.Backup(dirPath)
	backRepo.BackRepoModelPkg.Backup(dirPath)
	backRepo.BackRepoPointerToGongStructField.Backup(dirPath)
	backRepo.BackRepoSliceOfPointerToGongStructField.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.BackupXL(file)
	backRepo.BackRepoGongEnum.BackupXL(file)
	backRepo.BackRepoGongEnumValue.BackupXL(file)
	backRepo.BackRepoGongLink.BackupXL(file)
	backRepo.BackRepoGongNote.BackupXL(file)
	backRepo.BackRepoGongStruct.BackupXL(file)
	backRepo.BackRepoGongTimeField.BackupXL(file)
	backRepo.BackRepoMeta.BackupXL(file)
	backRepo.BackRepoMetaReference.BackupXL(file)
	backRepo.BackRepoModelPkg.BackupXL(file)
	backRepo.BackRepoPointerToGongStructField.BackupXL(file)
	backRepo.BackRepoSliceOfPointerToGongStructField.BackupXL(file)

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
	backRepo.BackRepoGongBasicField.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnum.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongEnumValue.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongNote.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongStruct.RestorePhaseOne(dirPath)
	backRepo.BackRepoGongTimeField.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeta.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetaReference.RestorePhaseOne(dirPath)
	backRepo.BackRepoModelPkg.RestorePhaseOne(dirPath)
	backRepo.BackRepoPointerToGongStructField.RestorePhaseOne(dirPath)
	backRepo.BackRepoSliceOfPointerToGongStructField.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoGongBasicField.RestorePhaseTwo()
	backRepo.BackRepoGongEnum.RestorePhaseTwo()
	backRepo.BackRepoGongEnumValue.RestorePhaseTwo()
	backRepo.BackRepoGongLink.RestorePhaseTwo()
	backRepo.BackRepoGongNote.RestorePhaseTwo()
	backRepo.BackRepoGongStruct.RestorePhaseTwo()
	backRepo.BackRepoGongTimeField.RestorePhaseTwo()
	backRepo.BackRepoMeta.RestorePhaseTwo()
	backRepo.BackRepoMetaReference.RestorePhaseTwo()
	backRepo.BackRepoModelPkg.RestorePhaseTwo()
	backRepo.BackRepoPointerToGongStructField.RestorePhaseTwo()
	backRepo.BackRepoSliceOfPointerToGongStructField.RestorePhaseTwo()

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
	backRepo.BackRepoGongBasicField.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnum.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongEnumValue.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongNote.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongStruct.RestoreXLPhaseOne(file)
	backRepo.BackRepoGongTimeField.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeta.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetaReference.RestoreXLPhaseOne(file)
	backRepo.BackRepoModelPkg.RestoreXLPhaseOne(file)
	backRepo.BackRepoPointerToGongStructField.RestoreXLPhaseOne(file)
	backRepo.BackRepoSliceOfPointerToGongStructField.RestoreXLPhaseOne(file)

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
