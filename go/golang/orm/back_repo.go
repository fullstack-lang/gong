package orm

const BackRepoTemplateCode = `// do not modify, generated file
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

	"{{PkgPathRoot}}/db"
	"{{PkgPathRoot}}/models"

	` + gormFirstLineToBeRemoved + `
	"{{PkgPathRoot}}/orm/dbgorm"
	` + gormLastLineToBeRemoved + `

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations{{` + string(rune(BackRepoPerStructDeclarations)) + `}}
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

	` + liteFirstLineToBeRemoved + `
	db = NewDBLite()
	` + liteLastLineToBeRemoved + `

	` + gormFirstLineToBeRemoved + `
	db = dbgorm.NewDBWrapper(filename, "{{PkgPathRootWithoutSlashes}}",{{` + string(rune(BackRepoPerStructRefToStructDB)) + `}}
	)
	` + gormLastLineToBeRemoved + `

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations{{` + string(rune(BackRepoPerStructInits)) + `}}

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

	// insertion point for per struct back repo phase one commit{{` + string(rune(BackRepoPerStructPhaseOneCommits)) + `}}

	// insertion point for per struct back repo phase two commit{{` + string(rune(BackRepoPerStructPhaseTwoCommits)) + `}}

	// important to release the mutex before calls to IncrementCommitFromBackNb
	// because it will block otherwise
	backRepo.rwMutex.Unlock()

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.Stage) {

	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()
	// insertion point for per struct back repo phase one commit{{` + string(rune(BackRepoPerStructPhaseOneCheckouts)) + `}}

	// insertion point for per struct back repo phase two commit{{` + string(rune(BackRepoPerStructPhaseTwoCheckouts)) + `}}
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup{{` + string(rune(BackRepoBackup)) + `}}
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup{{` + string(rune(BackRepoBackupXL)) + `}}

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

	// insertion point for per struct backup{{` + string(rune(BackRepoRestorePhaseOne)) + `}}

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup{{` + string(rune(BackRepoRestorePhaseTwo)) + `}}

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

	// insertion point for per struct backup{{` + string(rune(BackRepoRestoreXLPhaseOne)) + `}}

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

`

type BackRepoSubTemplateInsertion int

const (
	BackRepoPerStructDeclarations BackRepoSubTemplateInsertion = iota
	BackRepoPerStructInits
	BackRepoPerStructRefToStructDB
	BackRepoPerStructPhaseOneCommits
	BackRepoPerStructPhaseTwoCommits
	BackRepoPerStructPhaseOneCheckouts
	BackRepoPerStructPhaseTwoCheckouts
	BackRepoInitAndCommit
	BackRepoInitAndCheckout
	BackRepoCommit
	BackRepoCheckout
	BackRepoBackup
	BackRepoBackupXL
	BackRepoRestorePhaseOne
	BackRepoRestoreXLPhaseOne
	BackRepoRestorePhaseTwo
)

var BackRepoSubTemplate map[string]string = // new line
map[string]string{

	string(rune(BackRepoPerStructDeclarations)): `
	BackRepo{{Structname}} BackRepo{{Structname}}Struct
`,

	string(rune(BackRepoPerStructInits)): `
	backRepo.BackRepo{{Structname}} = BackRepo{{Structname}}Struct{
		Map_{{Structname}}DBID_{{Structname}}Ptr: make(map[uint]*models.{{Structname}}, 0),
		Map_{{Structname}}DBID_{{Structname}}DB:  make(map[uint]*{{Structname}}DB, 0),
		Map_{{Structname}}Ptr_{{Structname}}DBID: make(map[*models.{{Structname}}]uint, 0),

		db:    db,
		stage: stage,
	}`,

	string(rune(BackRepoPerStructRefToStructDB)): `
		&{{Structname}}DB{},`,

	string(rune(BackRepoPerStructPhaseOneCommits)): `
	backRepo.BackRepo{{Structname}}.CommitPhaseOne(stage)`,

	string(rune(BackRepoPerStructPhaseTwoCommits)): `
	backRepo.BackRepo{{Structname}}.CommitPhaseTwo(backRepo)`,

	string(rune(BackRepoPerStructPhaseOneCheckouts)): `
	backRepo.BackRepo{{Structname}}.CheckoutPhaseOne()`,

	string(rune(BackRepoPerStructPhaseTwoCheckouts)): `
	backRepo.BackRepo{{Structname}}.CheckoutPhaseTwo(backRepo)`,

	string(rune(BackRepoInitAndCommit)): `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := BackRepo{{Structname}}Init(
		CreateMode,
		db); err != nil {
		return err
	}
`,

	string(rune(BackRepoInitAndCheckout)): `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := BackRepo{{Structname}}Init(
		CreateMode,
		db); err != nil {
		err := errors.New("AllORMToModels, CreateMode Translation of {{Structname}} failed")
		return err
	}
`,

	string(rune(BackRepoCheckout)): `
	if err := BackRepo{{Structname}}Init(
		UpdateMode,
		db); err != nil {
		err := errors.New("AllORMToModels, UpdateMode Translation of {{Structname}} failed")
		return err
	}
`,

	string(rune(BackRepoCommit)): `
	if err := BackRepo{{Structname}}Init(
		UpdateMode,
		db); err != nil {
		return err
	}
`,

	string(rune(BackRepoBackup)): `
	backRepo.BackRepo{{Structname}}.Backup(dirPath)`,

	string(rune(BackRepoBackupXL)): `
	backRepo.BackRepo{{Structname}}.BackupXL(file)`,

	string(rune(BackRepoRestorePhaseOne)): `
	backRepo.BackRepo{{Structname}}.RestorePhaseOne(dirPath)`,

	string(rune(BackRepoRestoreXLPhaseOne)): `
	backRepo.BackRepo{{Structname}}.RestoreXLPhaseOne(file)`,

	string(rune(BackRepoRestorePhaseTwo)): `
	backRepo.BackRepo{{Structname}}.RestorePhaseTwo()`,
}
