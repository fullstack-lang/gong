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
var __SvgImage__dummysDeclaration__ models.SvgImage
var _ = __SvgImage__dummysDeclaration__
var __SvgImage_time__dummyDeclaration time.Duration
var _ = __SvgImage_time__dummyDeclaration

var mutexSvgImage sync.Mutex

// An SvgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSvgImage
type SvgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SvgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSvgImage
type SvgImageInput struct {
	// The SvgImage to submit or modify
	// in: body
	SvgImage *orm.SvgImageAPI
}

// UpdateSvgImage
//
// swagger:route PATCH /svgimages/{ID} svgimages updateSvgImage
//
// # Update a svgimage
//
// Responses:
// default: genericError
//
//	200: svgimageDBResponse
func (controller *Controller) UpdateSvgImage(w http.ResponseWriter, r *http.Request) {

	mutexSvgImage.Lock()
	defer mutexSvgImage.Unlock()

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
	db := backRepo.BackRepoSvgImage.GetDB()

	// Validate input
	var input orm.SvgImageAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgimageDB orm.SvgImageDB

	// fetch the svgimage
	_, err := db.First(&svgimageDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgimageDB.CopyBasicFieldsFromSvgImage_WOP(&input.SvgImage_WOP)
	svgimageDB.SvgImagePointersEncoding = input.SvgImagePointersEncoding

	db, _ = db.Model(&svgimageDB)
	_, err = db.Updates(&svgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgimageNew := new(models.SvgImage)
	svgimageDB.CopyBasicFieldsToSvgImage(svgimageNew)

	// redeem pointers
	svgimageDB.DecodePointers(backRepo, svgimageNew)

	// get stage instance from DB instance, and call callback function
	svgimageOld := backRepo.BackRepoSvgImage.Map_SvgImageDBID_SvgImagePtr[svgimageDB.ID]
	if svgimageOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(svgimageOld, svgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgimageDB
	writeJSON(w, http.StatusOK, svgimageDB)
}
