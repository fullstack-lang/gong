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

	"github.com/fullstack-lang/gong/lib/maplibre/go/db"
	"github.com/fullstack-lang/gong/lib/maplibre/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/maplibre/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCoordinate BackRepoCoordinateStruct

	BackRepoLayer BackRepoLayerStruct

	BackRepoMapView BackRepoMapViewStruct

	BackRepoMarker BackRepoMarkerStruct

	BackRepoPolygon BackRepoPolygonStruct

	BackRepoPolyline BackRepoPolylineStruct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_maplibre_go",
		&CoordinateDB{},
		&LayerDB{},
		&MapViewDB{},
		&MarkerDB{},
		&PolygonDB{},
		&PolylineDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCoordinate = BackRepoCoordinateStruct{
		Map_CoordinateDBID_CoordinatePtr: make(map[uint]*models.Coordinate, 0),
		Map_CoordinateDBID_CoordinateDB:  make(map[uint]*CoordinateDB, 0),
		Map_CoordinatePtr_CoordinateDBID: make(map[*models.Coordinate]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLayer = BackRepoLayerStruct{
		Map_LayerDBID_LayerPtr: make(map[uint]*models.Layer, 0),
		Map_LayerDBID_LayerDB:  make(map[uint]*LayerDB, 0),
		Map_LayerPtr_LayerDBID: make(map[*models.Layer]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMapView = BackRepoMapViewStruct{
		Map_MapViewDBID_MapViewPtr: make(map[uint]*models.MapView, 0),
		Map_MapViewDBID_MapViewDB:  make(map[uint]*MapViewDB, 0),
		Map_MapViewPtr_MapViewDBID: make(map[*models.MapView]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMarker = BackRepoMarkerStruct{
		Map_MarkerDBID_MarkerPtr: make(map[uint]*models.Marker, 0),
		Map_MarkerDBID_MarkerDB:  make(map[uint]*MarkerDB, 0),
		Map_MarkerPtr_MarkerDBID: make(map[*models.Marker]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPolygon = BackRepoPolygonStruct{
		Map_PolygonDBID_PolygonPtr: make(map[uint]*models.Polygon, 0),
		Map_PolygonDBID_PolygonDB:  make(map[uint]*PolygonDB, 0),
		Map_PolygonPtr_PolygonDBID: make(map[*models.Polygon]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPolyline = BackRepoPolylineStruct{
		Map_PolylineDBID_PolylinePtr: make(map[uint]*models.Polyline, 0),
		Map_PolylineDBID_PolylineDB:  make(map[uint]*PolylineDB, 0),
		Map_PolylinePtr_PolylineDBID: make(map[*models.Polyline]uint, 0),

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
	backRepo.BackRepoCoordinate.CommitPhaseOne(stage)
	backRepo.BackRepoLayer.CommitPhaseOne(stage)
	backRepo.BackRepoMapView.CommitPhaseOne(stage)
	backRepo.BackRepoMarker.CommitPhaseOne(stage)
	backRepo.BackRepoPolygon.CommitPhaseOne(stage)
	backRepo.BackRepoPolyline.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCoordinate.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLayer.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMapView.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMarker.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPolygon.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPolyline.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoCoordinate.CheckoutPhaseOne()
	backRepo.BackRepoLayer.CheckoutPhaseOne()
	backRepo.BackRepoMapView.CheckoutPhaseOne()
	backRepo.BackRepoMarker.CheckoutPhaseOne()
	backRepo.BackRepoPolygon.CheckoutPhaseOne()
	backRepo.BackRepoPolyline.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCoordinate.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLayer.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMapView.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMarker.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPolygon.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPolyline.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCoordinate.Backup(dirPath)
	backRepo.BackRepoLayer.Backup(dirPath)
	backRepo.BackRepoMapView.Backup(dirPath)
	backRepo.BackRepoMarker.Backup(dirPath)
	backRepo.BackRepoPolygon.Backup(dirPath)
	backRepo.BackRepoPolyline.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCoordinate.BackupXL(file)
	backRepo.BackRepoLayer.BackupXL(file)
	backRepo.BackRepoMapView.BackupXL(file)
	backRepo.BackRepoMarker.BackupXL(file)
	backRepo.BackRepoPolygon.BackupXL(file)
	backRepo.BackRepoPolyline.BackupXL(file)

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
	backRepo.BackRepoCoordinate.RestorePhaseOne(dirPath)
	backRepo.BackRepoLayer.RestorePhaseOne(dirPath)
	backRepo.BackRepoMapView.RestorePhaseOne(dirPath)
	backRepo.BackRepoMarker.RestorePhaseOne(dirPath)
	backRepo.BackRepoPolygon.RestorePhaseOne(dirPath)
	backRepo.BackRepoPolyline.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCoordinate.RestorePhaseTwo()
	backRepo.BackRepoLayer.RestorePhaseTwo()
	backRepo.BackRepoMapView.RestorePhaseTwo()
	backRepo.BackRepoMarker.RestorePhaseTwo()
	backRepo.BackRepoPolygon.RestorePhaseTwo()
	backRepo.BackRepoPolyline.RestorePhaseTwo()

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
	backRepo.BackRepoCoordinate.RestoreXLPhaseOne(file)
	backRepo.BackRepoLayer.RestoreXLPhaseOne(file)
	backRepo.BackRepoMapView.RestoreXLPhaseOne(file)
	backRepo.BackRepoMarker.RestoreXLPhaseOne(file)
	backRepo.BackRepoPolygon.RestoreXLPhaseOne(file)
	backRepo.BackRepoPolyline.RestoreXLPhaseOne(file)

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
