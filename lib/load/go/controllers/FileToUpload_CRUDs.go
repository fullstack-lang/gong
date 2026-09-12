// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	"github.com/fullstack-lang/gong/lib/load/go/orm"
)

// declaration in order to justify use of the models import
var __FileToUpload__dummysDeclaration__ models.FileToUpload
var _ = __FileToUpload__dummysDeclaration__
var __FileToUpload_time__dummyDeclaration time.Duration
var _ = __FileToUpload_time__dummyDeclaration

var mutexFileToUpload sync.Mutex

// An FileToUploadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFileToUpload updateFileToUpload deleteFileToUpload
type FileToUploadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FileToUploadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFileToUpload updateFileToUpload
type FileToUploadInput struct {
	// The FileToUpload to submit or modify
	// in: body
	FileToUpload *orm.FileToUploadAPI
}

// GetFileToUploads
//
// swagger:route GET /filetouploads filetouploads getFileToUploads
//
// # Get all filetouploads
//
// Responses:
// default: genericError
//
//	200: filetouploadDBResponse
func (controller *Controller) GetFileToUploads(w http.ResponseWriter, r *http.Request) {

	// source slice
	var filetouploadDBs []orm.FileToUploadDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFileToUploads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToUpload.GetDB()

	_, err := db.Find(&filetouploadDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	filetouploadAPIs := make([]orm.FileToUploadAPI, 0)

	// for each filetoupload, update fields from the database nullable fields
	for idx := range filetouploadDBs {
		filetouploadDB := &filetouploadDBs[idx]
		_ = filetouploadDB
		var filetouploadAPI orm.FileToUploadAPI

		// insertion point for updating fields
		filetouploadAPI.ID = filetouploadDB.ID
		filetouploadDB.CopyBasicFieldsToFileToUpload_WOP(&filetouploadAPI.FileToUpload_WOP)
		filetouploadAPI.FileToUploadPointersEncoding = filetouploadDB.FileToUploadPointersEncoding
		filetouploadAPIs = append(filetouploadAPIs, filetouploadAPI)
	}

	writeJSON(w, http.StatusOK, filetouploadAPIs)
}

// PostFileToUpload
//
// swagger:route POST /filetouploads filetouploads postFileToUpload
//
// Creates a filetoupload
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFileToUpload(w http.ResponseWriter, r *http.Request) {

	mutexFileToUpload.Lock()
	defer mutexFileToUpload.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFileToUploads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToUpload.GetDB()

	// Validate input
	var input orm.FileToUploadAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create filetoupload
	filetouploadDB := orm.FileToUploadDB{}
	filetouploadDB.FileToUploadPointersEncoding = input.FileToUploadPointersEncoding
	filetouploadDB.CopyBasicFieldsFromFileToUpload_WOP(&input.FileToUpload_WOP)

	_, err = db.Create(&filetouploadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFileToUpload.CheckoutPhaseOneInstance(&filetouploadDB)
	filetoupload := backRepo.BackRepoFileToUpload.Map_FileToUploadDBID_FileToUploadPtr[filetouploadDB.ID]

	if filetoupload != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), filetoupload)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, filetouploadDB)
}

// GetFileToUpload
//
// swagger:route GET /filetouploads/{ID} filetouploads getFileToUpload
//
// Gets the details for a filetoupload.
//
// Responses:
// default: genericError
//
//	200: filetouploadDBResponse
func (controller *Controller) GetFileToUpload(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFileToUpload", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToUpload.GetDB()

	// Get filetouploadDB in DB
	var filetouploadDB orm.FileToUploadDB
	if _, err := db.First(&filetouploadDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var filetouploadAPI orm.FileToUploadAPI
	filetouploadAPI.ID = filetouploadDB.ID
	filetouploadAPI.FileToUploadPointersEncoding = filetouploadDB.FileToUploadPointersEncoding
	filetouploadDB.CopyBasicFieldsToFileToUpload_WOP(&filetouploadAPI.FileToUpload_WOP)

	writeJSON(w, http.StatusOK, filetouploadAPI)
}

// UpdateFileToUpload
//
// swagger:route PATCH /filetouploads/{ID} filetouploads updateFileToUpload
//
// # Update a filetoupload
//
// Responses:
// default: genericError
//
//	200: filetouploadDBResponse
func (controller *Controller) UpdateFileToUpload(w http.ResponseWriter, r *http.Request) {

	mutexFileToUpload.Lock()
	defer mutexFileToUpload.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToUpload.GetDB()

	// Validate input
	var input orm.FileToUploadAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var filetouploadDB orm.FileToUploadDB

	// fetch the filetoupload
	_, err := db.First(&filetouploadDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	filetouploadDB.CopyBasicFieldsFromFileToUpload_WOP(&input.FileToUpload_WOP)
	filetouploadDB.FileToUploadPointersEncoding = input.FileToUploadPointersEncoding

	db, _ = db.Model(&filetouploadDB)
	_, err = db.Updates(&filetouploadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	filetouploadNew := new(models.FileToUpload)
	filetouploadDB.CopyBasicFieldsToFileToUpload(filetouploadNew)

	// redeem pointers
	filetouploadDB.DecodePointers(backRepo, filetouploadNew)

	// get stage instance from DB instance, and call callback function
	filetouploadOld := backRepo.BackRepoFileToUpload.Map_FileToUploadDBID_FileToUploadPtr[filetouploadDB.ID]
	if filetouploadOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), filetouploadOld, filetouploadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the filetouploadDB
	writeJSON(w, http.StatusOK, filetouploadDB)
}

// DeleteFileToUpload
//
// swagger:route DELETE /filetouploads/{ID} filetouploads deleteFileToUpload
//
// # Delete a filetoupload
//
// default: genericError
//
//	200: filetouploadDBResponse
func (controller *Controller) DeleteFileToUpload(w http.ResponseWriter, r *http.Request) {

	mutexFileToUpload.Lock()
	defer mutexFileToUpload.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFileToUpload", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToUpload.GetDB()

	// Get model if exist
	var filetouploadDB orm.FileToUploadDB
	if _, err := db.First(&filetouploadDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&filetouploadDB)

	// get an instance (not staged) from DB instance, and call callback function
	filetouploadDeleted := new(models.FileToUpload)
	filetouploadDB.CopyBasicFieldsToFileToUpload(filetouploadDeleted)

	// get stage instance from DB instance, and call callback function
	filetouploadStaged := backRepo.BackRepoFileToUpload.Map_FileToUploadDBID_FileToUploadPtr[filetouploadDB.ID]
	if filetouploadStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), filetouploadStaged, filetouploadDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
