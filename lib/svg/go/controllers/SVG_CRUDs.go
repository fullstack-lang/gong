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
var __SVG__dummysDeclaration__ models.SVG
var _ = __SVG__dummysDeclaration__
var __SVG_time__dummyDeclaration time.Duration
var _ = __SVG_time__dummyDeclaration

var mutexSVG sync.Mutex

// An SVGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSVG
type SVGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SVGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSVG
type SVGInput struct {
	// The SVG to submit or modify
	// in: body
	SVG *orm.SVGAPI
}

// UpdateSVG
//
// swagger:route PATCH /svgs/{ID} svgs updateSVG
//
// # Update a svg
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) UpdateSVG(w http.ResponseWriter, r *http.Request) {

	mutexSVG.Lock()
	defer mutexSVG.Unlock()

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
	db := backRepo.BackRepoSVG.GetDB()

	// Validate input
	var input orm.SVGAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgDB orm.SVGDB

	// fetch the svg
	_, err := db.First(&svgDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgDB.CopyBasicFieldsFromSVG_WOP(&input.SVG_WOP)
	svgDB.SVGPointersEncoding = input.SVGPointersEncoding

	db, _ = db.Model(&svgDB)
	_, err = db.Updates(&svgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgNew := new(models.SVG)
	svgDB.CopyBasicFieldsToSVG(svgNew)

	// redeem pointers
	svgDB.DecodePointers(backRepo, svgNew)

	// get stage instance from DB instance, and call callback function
	svgOld := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	if svgOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(svgOld, svgNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgDB
	writeJSON(w, http.StatusOK, svgDB)
}
