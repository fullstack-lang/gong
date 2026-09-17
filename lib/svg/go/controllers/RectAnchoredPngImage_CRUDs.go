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
var __RectAnchoredPngImage__dummysDeclaration__ models.RectAnchoredPngImage
var _ = __RectAnchoredPngImage__dummysDeclaration__
var __RectAnchoredPngImage_time__dummyDeclaration time.Duration
var _ = __RectAnchoredPngImage_time__dummyDeclaration

var mutexRectAnchoredPngImage sync.Mutex

// An RectAnchoredPngImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateRectAnchoredPngImage
type RectAnchoredPngImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredPngImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateRectAnchoredPngImage
type RectAnchoredPngImageInput struct {
	// The RectAnchoredPngImage to submit or modify
	// in: body
	RectAnchoredPngImage *orm.RectAnchoredPngImageAPI
}

// UpdateRectAnchoredPngImage
//
// swagger:route PATCH /rectanchoredpngimages/{ID} rectanchoredpngimages updateRectAnchoredPngImage
//
// # Update a rectanchoredpngimage
//
// Responses:
// default: genericError
//
//	200: rectanchoredpngimageDBResponse
func (controller *Controller) UpdateRectAnchoredPngImage(w http.ResponseWriter, r *http.Request) {

	mutexRectAnchoredPngImage.Lock()
	defer mutexRectAnchoredPngImage.Unlock()

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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	// Validate input
	var input orm.RectAnchoredPngImageAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredpngimageDB orm.RectAnchoredPngImageDB

	// fetch the rectanchoredpngimage
	_, err := db.First(&rectanchoredpngimageDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredpngimageDB.CopyBasicFieldsFromRectAnchoredPngImage_WOP(&input.RectAnchoredPngImage_WOP)
	rectanchoredpngimageDB.RectAnchoredPngImagePointersEncoding = input.RectAnchoredPngImagePointersEncoding

	db, _ = db.Model(&rectanchoredpngimageDB)
	_, err = db.Updates(&rectanchoredpngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpngimageNew := new(models.RectAnchoredPngImage)
	rectanchoredpngimageDB.CopyBasicFieldsToRectAnchoredPngImage(rectanchoredpngimageNew)

	// redeem pointers
	rectanchoredpngimageDB.DecodePointers(backRepo, rectanchoredpngimageNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredpngimageOld := backRepo.BackRepoRectAnchoredPngImage.Map_RectAnchoredPngImageDBID_RectAnchoredPngImagePtr[rectanchoredpngimageDB.ID]
	if rectanchoredpngimageOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(rectanchoredpngimageOld, rectanchoredpngimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredpngimageDB
	writeJSON(w, http.StatusOK, rectanchoredpngimageDB)
}
