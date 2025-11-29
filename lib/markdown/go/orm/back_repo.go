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

	"github.com/fullstack-lang/gong/lib/markdown/go/db"
	"github.com/fullstack-lang/gong/lib/markdown/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/markdown/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoContent BackRepoContentStruct

	BackRepoJpgImage BackRepoJpgImageStruct

	BackRepoPngImage BackRepoPngImageStruct

	BackRepoSvgImage BackRepoSvgImageStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_markdown_go",
		&ContentDB{},
		&JpgImageDB{},
		&PngImageDB{},
		&SvgImageDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoContent = BackRepoContentStruct{
		Map_ContentDBID_ContentPtr: make(map[uint]*models.Content, 0),
		Map_ContentDBID_ContentDB:  make(map[uint]*ContentDB, 0),
		Map_ContentPtr_ContentDBID: make(map[*models.Content]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoJpgImage = BackRepoJpgImageStruct{
		Map_JpgImageDBID_JpgImagePtr: make(map[uint]*models.JpgImage, 0),
		Map_JpgImageDBID_JpgImageDB:  make(map[uint]*JpgImageDB, 0),
		Map_JpgImagePtr_JpgImageDBID: make(map[*models.JpgImage]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPngImage = BackRepoPngImageStruct{
		Map_PngImageDBID_PngImagePtr: make(map[uint]*models.PngImage, 0),
		Map_PngImageDBID_PngImageDB:  make(map[uint]*PngImageDB, 0),
		Map_PngImagePtr_PngImageDBID: make(map[*models.PngImage]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSvgImage = BackRepoSvgImageStruct{
		Map_SvgImageDBID_SvgImagePtr: make(map[uint]*models.SvgImage, 0),
		Map_SvgImageDBID_SvgImageDB:  make(map[uint]*SvgImageDB, 0),
		Map_SvgImagePtr_SvgImageDBID: make(map[*models.SvgImage]uint, 0),

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
	backRepo.BackRepoContent.CommitPhaseOne(stage)
	backRepo.BackRepoJpgImage.CommitPhaseOne(stage)
	backRepo.BackRepoPngImage.CommitPhaseOne(stage)
	backRepo.BackRepoSvgImage.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoContent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoJpgImage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPngImage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSvgImage.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoContent.CheckoutPhaseOne()
	backRepo.BackRepoJpgImage.CheckoutPhaseOne()
	backRepo.BackRepoPngImage.CheckoutPhaseOne()
	backRepo.BackRepoSvgImage.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoContent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoJpgImage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPngImage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSvgImage.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoContent.Backup(dirPath)
	backRepo.BackRepoJpgImage.Backup(dirPath)
	backRepo.BackRepoPngImage.Backup(dirPath)
	backRepo.BackRepoSvgImage.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoContent.BackupXL(file)
	backRepo.BackRepoJpgImage.BackupXL(file)
	backRepo.BackRepoPngImage.BackupXL(file)
	backRepo.BackRepoSvgImage.BackupXL(file)

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
	backRepo.BackRepoContent.RestorePhaseOne(dirPath)
	backRepo.BackRepoJpgImage.RestorePhaseOne(dirPath)
	backRepo.BackRepoPngImage.RestorePhaseOne(dirPath)
	backRepo.BackRepoSvgImage.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoContent.RestorePhaseTwo()
	backRepo.BackRepoJpgImage.RestorePhaseTwo()
	backRepo.BackRepoPngImage.RestorePhaseTwo()
	backRepo.BackRepoSvgImage.RestorePhaseTwo()

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
	backRepo.BackRepoContent.RestoreXLPhaseOne(file)
	backRepo.BackRepoJpgImage.RestoreXLPhaseOne(file)
	backRepo.BackRepoPngImage.RestoreXLPhaseOne(file)
	backRepo.BackRepoSvgImage.RestoreXLPhaseOne(file)

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
