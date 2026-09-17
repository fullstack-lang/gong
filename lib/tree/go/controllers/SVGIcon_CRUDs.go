// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/orm"
)

// declaration in order to justify use of the models import
var __SVGIcon__dummysDeclaration__ models.SVGIcon
var _ = __SVGIcon__dummysDeclaration__
var __SVGIcon_time__dummyDeclaration time.Duration
var _ = __SVGIcon_time__dummyDeclaration

var mutexSVGIcon sync.Mutex

// An SVGIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSVGIcon
type SVGIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SVGIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSVGIcon
type SVGIconInput struct {
	// The SVGIcon to submit or modify
	// in: body
	SVGIcon *orm.SVGIconAPI
}

// UpdateSVGIcon
//
// swagger:route PATCH /svgicons/{ID} svgicons updateSVGIcon
//
// # Update a svgicon
//
// Responses:
// default: genericError
//
//	200: svgiconDBResponse
func (controller *Controller) UpdateSVGIcon(w http.ResponseWriter, r *http.Request) {

	mutexSVGIcon.Lock()
	defer mutexSVGIcon.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	// Validate input
	var input orm.SVGIconAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgiconDB orm.SVGIconDB

	// fetch the svgicon
	_, err := db.First(&svgiconDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgiconDB.CopyBasicFieldsFromSVGIcon_WOP(&input.SVGIcon_WOP)
	svgiconDB.SVGIconPointersEncoding = input.SVGIconPointersEncoding

	db, _ = db.Model(&svgiconDB)
	_, err = db.Updates(&svgiconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgiconNew := new(models.SVGIcon)
	svgiconDB.CopyBasicFieldsToSVGIcon(svgiconNew)

	// redeem pointers
	svgiconDB.DecodePointers(backRepo, svgiconNew)

	// get stage instance from DB instance, and call callback function
	svgiconOld := backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]
	if svgiconOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(svgiconOld, svgiconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgiconDB
	writeJSON(w, http.StatusOK, svgiconDB)
}
