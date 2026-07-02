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

	"github.com/fullstack-lang/gong/lib/threejs/go/db"
	"github.com/fullstack-lang/gong/lib/threejs/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/threejs/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAmbiantLight BackRepoAmbiantLightStruct

	BackRepoBoxGeometry BackRepoBoxGeometryStruct

	BackRepoCanvas BackRepoCanvasStruct

	BackRepoCurve BackRepoCurveStruct

	BackRepoCylinderGeometry BackRepoCylinderGeometryStruct

	BackRepoDirectionalLight BackRepoDirectionalLightStruct

	BackRepoExtrudeGeometry BackRepoExtrudeGeometryStruct

	BackRepoMesh BackRepoMeshStruct

	BackRepoMeshMaterialBasic BackRepoMeshMaterialBasicStruct

	BackRepoMeshPhysicalMaterial BackRepoMeshPhysicalMaterialStruct

	BackRepoPlaneGeometry BackRepoPlaneGeometryStruct

	BackRepoShape BackRepoShapeStruct

	BackRepoSphereGeometry BackRepoSphereGeometryStruct

	BackRepoTorusGeometry BackRepoTorusGeometryStruct

	BackRepoTubeGeometry BackRepoTubeGeometryStruct

	BackRepoVector2 BackRepoVector2Struct

	BackRepoVector3 BackRepoVector3Struct

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
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_threejs_go",
		&AmbiantLightDB{},
		&BoxGeometryDB{},
		&CanvasDB{},
		&CurveDB{},
		&CylinderGeometryDB{},
		&DirectionalLightDB{},
		&ExtrudeGeometryDB{},
		&MeshDB{},
		&MeshMaterialBasicDB{},
		&MeshPhysicalMaterialDB{},
		&PlaneGeometryDB{},
		&ShapeDB{},
		&SphereGeometryDB{},
		&TorusGeometryDB{},
		&TubeGeometryDB{},
		&Vector2DB{},
		&Vector3DB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAmbiantLight = BackRepoAmbiantLightStruct{
		Map_AmbiantLightDBID_AmbiantLightPtr: make(map[uint]*models.AmbiantLight, 0),
		Map_AmbiantLightDBID_AmbiantLightDB:  make(map[uint]*AmbiantLightDB, 0),
		Map_AmbiantLightPtr_AmbiantLightDBID: make(map[*models.AmbiantLight]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBoxGeometry = BackRepoBoxGeometryStruct{
		Map_BoxGeometryDBID_BoxGeometryPtr: make(map[uint]*models.BoxGeometry, 0),
		Map_BoxGeometryDBID_BoxGeometryDB:  make(map[uint]*BoxGeometryDB, 0),
		Map_BoxGeometryPtr_BoxGeometryDBID: make(map[*models.BoxGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCanvas = BackRepoCanvasStruct{
		Map_CanvasDBID_CanvasPtr: make(map[uint]*models.Canvas, 0),
		Map_CanvasDBID_CanvasDB:  make(map[uint]*CanvasDB, 0),
		Map_CanvasPtr_CanvasDBID: make(map[*models.Canvas]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCurve = BackRepoCurveStruct{
		Map_CurveDBID_CurvePtr: make(map[uint]*models.Curve, 0),
		Map_CurveDBID_CurveDB:  make(map[uint]*CurveDB, 0),
		Map_CurvePtr_CurveDBID: make(map[*models.Curve]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCylinderGeometry = BackRepoCylinderGeometryStruct{
		Map_CylinderGeometryDBID_CylinderGeometryPtr: make(map[uint]*models.CylinderGeometry, 0),
		Map_CylinderGeometryDBID_CylinderGeometryDB:  make(map[uint]*CylinderGeometryDB, 0),
		Map_CylinderGeometryPtr_CylinderGeometryDBID: make(map[*models.CylinderGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDirectionalLight = BackRepoDirectionalLightStruct{
		Map_DirectionalLightDBID_DirectionalLightPtr: make(map[uint]*models.DirectionalLight, 0),
		Map_DirectionalLightDBID_DirectionalLightDB:  make(map[uint]*DirectionalLightDB, 0),
		Map_DirectionalLightPtr_DirectionalLightDBID: make(map[*models.DirectionalLight]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoExtrudeGeometry = BackRepoExtrudeGeometryStruct{
		Map_ExtrudeGeometryDBID_ExtrudeGeometryPtr: make(map[uint]*models.ExtrudeGeometry, 0),
		Map_ExtrudeGeometryDBID_ExtrudeGeometryDB:  make(map[uint]*ExtrudeGeometryDB, 0),
		Map_ExtrudeGeometryPtr_ExtrudeGeometryDBID: make(map[*models.ExtrudeGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMesh = BackRepoMeshStruct{
		Map_MeshDBID_MeshPtr: make(map[uint]*models.Mesh, 0),
		Map_MeshDBID_MeshDB:  make(map[uint]*MeshDB, 0),
		Map_MeshPtr_MeshDBID: make(map[*models.Mesh]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeshMaterialBasic = BackRepoMeshMaterialBasicStruct{
		Map_MeshMaterialBasicDBID_MeshMaterialBasicPtr: make(map[uint]*models.MeshMaterialBasic, 0),
		Map_MeshMaterialBasicDBID_MeshMaterialBasicDB:  make(map[uint]*MeshMaterialBasicDB, 0),
		Map_MeshMaterialBasicPtr_MeshMaterialBasicDBID: make(map[*models.MeshMaterialBasic]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeshPhysicalMaterial = BackRepoMeshPhysicalMaterialStruct{
		Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialPtr: make(map[uint]*models.MeshPhysicalMaterial, 0),
		Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialDB:  make(map[uint]*MeshPhysicalMaterialDB, 0),
		Map_MeshPhysicalMaterialPtr_MeshPhysicalMaterialDBID: make(map[*models.MeshPhysicalMaterial]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPlaneGeometry = BackRepoPlaneGeometryStruct{
		Map_PlaneGeometryDBID_PlaneGeometryPtr: make(map[uint]*models.PlaneGeometry, 0),
		Map_PlaneGeometryDBID_PlaneGeometryDB:  make(map[uint]*PlaneGeometryDB, 0),
		Map_PlaneGeometryPtr_PlaneGeometryDBID: make(map[*models.PlaneGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoShape = BackRepoShapeStruct{
		Map_ShapeDBID_ShapePtr: make(map[uint]*models.Shape, 0),
		Map_ShapeDBID_ShapeDB:  make(map[uint]*ShapeDB, 0),
		Map_ShapePtr_ShapeDBID: make(map[*models.Shape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSphereGeometry = BackRepoSphereGeometryStruct{
		Map_SphereGeometryDBID_SphereGeometryPtr: make(map[uint]*models.SphereGeometry, 0),
		Map_SphereGeometryDBID_SphereGeometryDB:  make(map[uint]*SphereGeometryDB, 0),
		Map_SphereGeometryPtr_SphereGeometryDBID: make(map[*models.SphereGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTorusGeometry = BackRepoTorusGeometryStruct{
		Map_TorusGeometryDBID_TorusGeometryPtr: make(map[uint]*models.TorusGeometry, 0),
		Map_TorusGeometryDBID_TorusGeometryDB:  make(map[uint]*TorusGeometryDB, 0),
		Map_TorusGeometryPtr_TorusGeometryDBID: make(map[*models.TorusGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTubeGeometry = BackRepoTubeGeometryStruct{
		Map_TubeGeometryDBID_TubeGeometryPtr: make(map[uint]*models.TubeGeometry, 0),
		Map_TubeGeometryDBID_TubeGeometryDB:  make(map[uint]*TubeGeometryDB, 0),
		Map_TubeGeometryPtr_TubeGeometryDBID: make(map[*models.TubeGeometry]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVector2 = BackRepoVector2Struct{
		Map_Vector2DBID_Vector2Ptr: make(map[uint]*models.Vector2, 0),
		Map_Vector2DBID_Vector2DB:  make(map[uint]*Vector2DB, 0),
		Map_Vector2Ptr_Vector2DBID: make(map[*models.Vector2]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVector3 = BackRepoVector3Struct{
		Map_Vector3DBID_Vector3Ptr: make(map[uint]*models.Vector3, 0),
		Map_Vector3DBID_Vector3DB:  make(map[uint]*Vector3DB, 0),
		Map_Vector3Ptr_Vector3DBID: make(map[*models.Vector3]uint, 0),

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
	backRepo.BackRepoAmbiantLight.CommitPhaseOne(stage)
	backRepo.BackRepoBoxGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoCanvas.CommitPhaseOne(stage)
	backRepo.BackRepoCurve.CommitPhaseOne(stage)
	backRepo.BackRepoCylinderGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoDirectionalLight.CommitPhaseOne(stage)
	backRepo.BackRepoExtrudeGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoMesh.CommitPhaseOne(stage)
	backRepo.BackRepoMeshMaterialBasic.CommitPhaseOne(stage)
	backRepo.BackRepoMeshPhysicalMaterial.CommitPhaseOne(stage)
	backRepo.BackRepoPlaneGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoShape.CommitPhaseOne(stage)
	backRepo.BackRepoSphereGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoTorusGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoTubeGeometry.CommitPhaseOne(stage)
	backRepo.BackRepoVector2.CommitPhaseOne(stage)
	backRepo.BackRepoVector3.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAmbiantLight.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBoxGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCanvas.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCurve.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCylinderGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDirectionalLight.CommitPhaseTwo(backRepo)
	backRepo.BackRepoExtrudeGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMesh.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeshMaterialBasic.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeshPhysicalMaterial.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPlaneGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSphereGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTorusGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTubeGeometry.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVector2.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVector3.CommitPhaseTwo(backRepo)

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
	backRepo.BackRepoAmbiantLight.CheckoutPhaseOne()
	backRepo.BackRepoBoxGeometry.CheckoutPhaseOne()
	backRepo.BackRepoCanvas.CheckoutPhaseOne()
	backRepo.BackRepoCurve.CheckoutPhaseOne()
	backRepo.BackRepoCylinderGeometry.CheckoutPhaseOne()
	backRepo.BackRepoDirectionalLight.CheckoutPhaseOne()
	backRepo.BackRepoExtrudeGeometry.CheckoutPhaseOne()
	backRepo.BackRepoMesh.CheckoutPhaseOne()
	backRepo.BackRepoMeshMaterialBasic.CheckoutPhaseOne()
	backRepo.BackRepoMeshPhysicalMaterial.CheckoutPhaseOne()
	backRepo.BackRepoPlaneGeometry.CheckoutPhaseOne()
	backRepo.BackRepoShape.CheckoutPhaseOne()
	backRepo.BackRepoSphereGeometry.CheckoutPhaseOne()
	backRepo.BackRepoTorusGeometry.CheckoutPhaseOne()
	backRepo.BackRepoTubeGeometry.CheckoutPhaseOne()
	backRepo.BackRepoVector2.CheckoutPhaseOne()
	backRepo.BackRepoVector3.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAmbiantLight.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBoxGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCanvas.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCurve.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCylinderGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDirectionalLight.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoExtrudeGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMesh.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeshMaterialBasic.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeshPhysicalMaterial.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPlaneGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSphereGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTorusGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTubeGeometry.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVector2.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVector3.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAmbiantLight.Backup(dirPath)
	backRepo.BackRepoBoxGeometry.Backup(dirPath)
	backRepo.BackRepoCanvas.Backup(dirPath)
	backRepo.BackRepoCurve.Backup(dirPath)
	backRepo.BackRepoCylinderGeometry.Backup(dirPath)
	backRepo.BackRepoDirectionalLight.Backup(dirPath)
	backRepo.BackRepoExtrudeGeometry.Backup(dirPath)
	backRepo.BackRepoMesh.Backup(dirPath)
	backRepo.BackRepoMeshMaterialBasic.Backup(dirPath)
	backRepo.BackRepoMeshPhysicalMaterial.Backup(dirPath)
	backRepo.BackRepoPlaneGeometry.Backup(dirPath)
	backRepo.BackRepoShape.Backup(dirPath)
	backRepo.BackRepoSphereGeometry.Backup(dirPath)
	backRepo.BackRepoTorusGeometry.Backup(dirPath)
	backRepo.BackRepoTubeGeometry.Backup(dirPath)
	backRepo.BackRepoVector2.Backup(dirPath)
	backRepo.BackRepoVector3.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAmbiantLight.BackupXL(file)
	backRepo.BackRepoBoxGeometry.BackupXL(file)
	backRepo.BackRepoCanvas.BackupXL(file)
	backRepo.BackRepoCurve.BackupXL(file)
	backRepo.BackRepoCylinderGeometry.BackupXL(file)
	backRepo.BackRepoDirectionalLight.BackupXL(file)
	backRepo.BackRepoExtrudeGeometry.BackupXL(file)
	backRepo.BackRepoMesh.BackupXL(file)
	backRepo.BackRepoMeshMaterialBasic.BackupXL(file)
	backRepo.BackRepoMeshPhysicalMaterial.BackupXL(file)
	backRepo.BackRepoPlaneGeometry.BackupXL(file)
	backRepo.BackRepoShape.BackupXL(file)
	backRepo.BackRepoSphereGeometry.BackupXL(file)
	backRepo.BackRepoTorusGeometry.BackupXL(file)
	backRepo.BackRepoTubeGeometry.BackupXL(file)
	backRepo.BackRepoVector2.BackupXL(file)
	backRepo.BackRepoVector3.BackupXL(file)

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
	backRepo.BackRepoAmbiantLight.RestorePhaseOne(dirPath)
	backRepo.BackRepoBoxGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoCanvas.RestorePhaseOne(dirPath)
	backRepo.BackRepoCurve.RestorePhaseOne(dirPath)
	backRepo.BackRepoCylinderGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoDirectionalLight.RestorePhaseOne(dirPath)
	backRepo.BackRepoExtrudeGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoMesh.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeshMaterialBasic.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeshPhysicalMaterial.RestorePhaseOne(dirPath)
	backRepo.BackRepoPlaneGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoSphereGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoTorusGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoTubeGeometry.RestorePhaseOne(dirPath)
	backRepo.BackRepoVector2.RestorePhaseOne(dirPath)
	backRepo.BackRepoVector3.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAmbiantLight.RestorePhaseTwo()
	backRepo.BackRepoBoxGeometry.RestorePhaseTwo()
	backRepo.BackRepoCanvas.RestorePhaseTwo()
	backRepo.BackRepoCurve.RestorePhaseTwo()
	backRepo.BackRepoCylinderGeometry.RestorePhaseTwo()
	backRepo.BackRepoDirectionalLight.RestorePhaseTwo()
	backRepo.BackRepoExtrudeGeometry.RestorePhaseTwo()
	backRepo.BackRepoMesh.RestorePhaseTwo()
	backRepo.BackRepoMeshMaterialBasic.RestorePhaseTwo()
	backRepo.BackRepoMeshPhysicalMaterial.RestorePhaseTwo()
	backRepo.BackRepoPlaneGeometry.RestorePhaseTwo()
	backRepo.BackRepoShape.RestorePhaseTwo()
	backRepo.BackRepoSphereGeometry.RestorePhaseTwo()
	backRepo.BackRepoTorusGeometry.RestorePhaseTwo()
	backRepo.BackRepoTubeGeometry.RestorePhaseTwo()
	backRepo.BackRepoVector2.RestorePhaseTwo()
	backRepo.BackRepoVector3.RestorePhaseTwo()

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
	backRepo.BackRepoAmbiantLight.RestoreXLPhaseOne(file)
	backRepo.BackRepoBoxGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoCanvas.RestoreXLPhaseOne(file)
	backRepo.BackRepoCurve.RestoreXLPhaseOne(file)
	backRepo.BackRepoCylinderGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoDirectionalLight.RestoreXLPhaseOne(file)
	backRepo.BackRepoExtrudeGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoMesh.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeshMaterialBasic.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeshPhysicalMaterial.RestoreXLPhaseOne(file)
	backRepo.BackRepoPlaneGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoSphereGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoTorusGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoTubeGeometry.RestoreXLPhaseOne(file)
	backRepo.BackRepoVector2.RestoreXLPhaseOne(file)
	backRepo.BackRepoVector3.RestoreXLPhaseOne(file)

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
