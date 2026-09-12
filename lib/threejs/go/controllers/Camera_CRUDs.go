// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"
)

// declaration in order to justify use of the models import
var __Camera__dummysDeclaration__ models.Camera
var _ = __Camera__dummysDeclaration__
var __Camera_time__dummyDeclaration time.Duration
var _ = __Camera_time__dummyDeclaration

var mutexCamera sync.Mutex

// An CameraID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCamera updateCamera deleteCamera
type CameraID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CameraInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCamera updateCamera
type CameraInput struct {
	// The Camera to submit or modify
	// in: body
	Camera *orm.CameraAPI
}

// GetCameras
//
// swagger:route GET /cameras cameras getCameras
//
// # Get all cameras
//
// Responses:
// default: genericError
//
//	200: cameraDBResponse
func (controller *Controller) GetCameras(w http.ResponseWriter, r *http.Request) {

	// source slice
	var cameraDBs []orm.CameraDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCameras", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCamera.GetDB()

	_, err := db.Find(&cameraDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cameraAPIs := make([]orm.CameraAPI, 0)

	// for each camera, update fields from the database nullable fields
	for idx := range cameraDBs {
		cameraDB := &cameraDBs[idx]
		_ = cameraDB
		var cameraAPI orm.CameraAPI

		// insertion point for updating fields
		cameraAPI.ID = cameraDB.ID
		cameraDB.CopyBasicFieldsToCamera_WOP(&cameraAPI.Camera_WOP)
		cameraAPI.CameraPointersEncoding = cameraDB.CameraPointersEncoding
		cameraAPIs = append(cameraAPIs, cameraAPI)
	}

	writeJSON(w, http.StatusOK, cameraAPIs)
}

// PostCamera
//
// swagger:route POST /cameras cameras postCamera
//
// Creates a camera
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCamera(w http.ResponseWriter, r *http.Request) {

	mutexCamera.Lock()
	defer mutexCamera.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCameras", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCamera.GetDB()

	// Validate input
	var input orm.CameraAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create camera
	cameraDB := orm.CameraDB{}
	cameraDB.CameraPointersEncoding = input.CameraPointersEncoding
	cameraDB.CopyBasicFieldsFromCamera_WOP(&input.Camera_WOP)

	_, err = db.Create(&cameraDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCamera.CheckoutPhaseOneInstance(&cameraDB)
	camera := backRepo.BackRepoCamera.Map_CameraDBID_CameraPtr[cameraDB.ID]

	if camera != nil {
		backRepo.GetStage().AfterCreateFromFront(camera)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, cameraDB)
}

// GetCamera
//
// swagger:route GET /cameras/{ID} cameras getCamera
//
// Gets the details for a camera.
//
// Responses:
// default: genericError
//
//	200: cameraDBResponse
func (controller *Controller) GetCamera(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCamera", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCamera.GetDB()

	// Get cameraDB in DB
	var cameraDB orm.CameraDB
	if _, err := db.First(&cameraDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var cameraAPI orm.CameraAPI
	cameraAPI.ID = cameraDB.ID
	cameraAPI.CameraPointersEncoding = cameraDB.CameraPointersEncoding
	cameraDB.CopyBasicFieldsToCamera_WOP(&cameraAPI.Camera_WOP)

	writeJSON(w, http.StatusOK, cameraAPI)
}

// UpdateCamera
//
// swagger:route PATCH /cameras/{ID} cameras updateCamera
//
// # Update a camera
//
// Responses:
// default: genericError
//
//	200: cameraDBResponse
func (controller *Controller) UpdateCamera(w http.ResponseWriter, r *http.Request) {

	mutexCamera.Lock()
	defer mutexCamera.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCamera.GetDB()

	// Validate input
	var input orm.CameraAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cameraDB orm.CameraDB

	// fetch the camera
	_, err := db.First(&cameraDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cameraDB.CopyBasicFieldsFromCamera_WOP(&input.Camera_WOP)
	cameraDB.CameraPointersEncoding = input.CameraPointersEncoding

	db, _ = db.Model(&cameraDB)
	_, err = db.Updates(&cameraDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cameraNew := new(models.Camera)
	cameraDB.CopyBasicFieldsToCamera(cameraNew)

	// redeem pointers
	cameraDB.DecodePointers(backRepo, cameraNew)

	// get stage instance from DB instance, and call callback function
	cameraOld := backRepo.BackRepoCamera.Map_CameraDBID_CameraPtr[cameraDB.ID]
	if cameraOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(cameraOld, cameraNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cameraDB
	writeJSON(w, http.StatusOK, cameraDB)
}

// DeleteCamera
//
// swagger:route DELETE /cameras/{ID} cameras deleteCamera
//
// # Delete a camera
//
// default: genericError
//
//	200: cameraDBResponse
func (controller *Controller) DeleteCamera(w http.ResponseWriter, r *http.Request) {

	mutexCamera.Lock()
	defer mutexCamera.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCamera", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCamera.GetDB()

	// Get model if exist
	var cameraDB orm.CameraDB
	if _, err := db.First(&cameraDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cameraDB)

	// get an instance (not staged) from DB instance, and call callback function
	cameraDeleted := new(models.Camera)
	cameraDB.CopyBasicFieldsToCamera(cameraDeleted)

	// get stage instance from DB instance, and call callback function
	cameraStaged := backRepo.BackRepoCamera.Map_CameraDBID_CameraPtr[cameraDB.ID]
	if cameraStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(cameraStaged, cameraDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
