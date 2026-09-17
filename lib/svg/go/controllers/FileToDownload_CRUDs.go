// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// declaration in order to justify use of the models import
var __FileToDownload__dummysDeclaration__ models.FileToDownload
var _ = __FileToDownload__dummysDeclaration__
var __FileToDownload_time__dummyDeclaration time.Duration
var _ = __FileToDownload_time__dummyDeclaration

var mutexFileToDownload sync.Mutex

// An FileToDownloadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFileToDownload
type FileToDownloadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FileToDownloadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFileToDownload
type FileToDownloadInput struct {
	// The FileToDownload to submit or modify
	// in: body
	FileToDownload *orm.FileToDownloadAPI
}

// UpdateFileToDownload
//
// swagger:route PATCH /filetodownloads/{ID} filetodownloads updateFileToDownload
//
// # Update a filetodownload
//
// Responses:
// default: genericError
//
//	200: filetodownloadDBResponse
func (controller *Controller) UpdateFileToDownload(w http.ResponseWriter, r *http.Request) {

	mutexFileToDownload.Lock()
	defer mutexFileToDownload.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	// Validate input
	var input orm.FileToDownloadAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var filetodownloadDB orm.FileToDownloadDB

	// fetch the filetodownload
	_, err := db.First(&filetodownloadDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	filetodownloadDB.CopyBasicFieldsFromFileToDownload_WOP(&input.FileToDownload_WOP)
	filetodownloadDB.FileToDownloadPointersEncoding = input.FileToDownloadPointersEncoding

	db, _ = db.Model(&filetodownloadDB)
	_, err = db.Updates(&filetodownloadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	filetodownloadNew := new(models.FileToDownload)
	filetodownloadDB.CopyBasicFieldsToFileToDownload(filetodownloadNew)

	// redeem pointers
	filetodownloadDB.DecodePointers(backRepo, filetodownloadNew)

	// get stage instance from DB instance, and call callback function
	filetodownloadOld := backRepo.BackRepoFileToDownload.Map_FileToDownloadDBID_FileToDownloadPtr[filetodownloadDB.ID]
	if filetodownloadOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(filetodownloadOld, filetodownloadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the filetodownloadDB
	writeJSON(w, http.StatusOK, filetodownloadDB)
}
