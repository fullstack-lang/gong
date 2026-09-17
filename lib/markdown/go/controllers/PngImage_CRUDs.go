// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
	"github.com/fullstack-lang/gong/lib/markdown/go/orm"
)

// declaration in order to justify use of the models import
var __PngImage__dummysDeclaration__ models.PngImage
var _ = __PngImage__dummysDeclaration__
var __PngImage_time__dummyDeclaration time.Duration
var _ = __PngImage_time__dummyDeclaration

var mutexPngImage sync.Mutex

// An PngImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePngImage
type PngImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PngImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePngImage
type PngImageInput struct {
	// The PngImage to submit or modify
	// in: body
	PngImage *orm.PngImageAPI
}

// UpdatePngImage
//
// swagger:route PATCH /pngimages/{ID} pngimages updatePngImage
//
// # Update a pngimage
//
// Responses:
// default: genericError
//
//	200: pngimageDBResponse
func (controller *Controller) UpdatePngImage(w http.ResponseWriter, r *http.Request) {

	mutexPngImage.Lock()
	defer mutexPngImage.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoPngImage.GetDB()

	// Validate input
	var input orm.PngImageAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pngimageDB orm.PngImageDB

	// fetch the pngimage
	_, err := db.First(&pngimageDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pngimageDB.CopyBasicFieldsFromPngImage_WOP(&input.PngImage_WOP)
	pngimageDB.PngImagePointersEncoding = input.PngImagePointersEncoding

	db, _ = db.Model(&pngimageDB)
	_, err = db.Updates(&pngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pngimageNew := new(models.PngImage)
	pngimageDB.CopyBasicFieldsToPngImage(pngimageNew)

	// redeem pointers
	pngimageDB.DecodePointers(backRepo, pngimageNew)

	// get stage instance from DB instance, and call callback function
	pngimageOld := backRepo.BackRepoPngImage.Map_PngImageDBID_PngImagePtr[pngimageDB.ID]
	if pngimageOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(pngimageOld, pngimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pngimageDB
	writeJSON(w, http.StatusOK, pngimageDB)
}
