// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

// declaration in order to justify use of the models import
var __Xlsx__dummysDeclaration__ models.Xlsx
var _ = __Xlsx__dummysDeclaration__
var __Xlsx_time__dummyDeclaration time.Duration
var _ = __Xlsx_time__dummyDeclaration

var mutexXlsx sync.Mutex

// An XlsxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateXlsx
type XlsxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XlsxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateXlsx
type XlsxInput struct {
	// The Xlsx to submit or modify
	// in: body
	Xlsx *orm.XlsxAPI
}

// UpdateXlsx
//
// swagger:route PATCH /xlsxs/{ID} xlsxs updateXlsx
//
// # Update a xlsx
//
// Responses:
// default: genericError
//
//	200: xlsxDBResponse
func (controller *Controller) UpdateXlsx(w http.ResponseWriter, r *http.Request) {

	mutexXlsx.Lock()
	defer mutexXlsx.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	// Validate input
	var input orm.XlsxAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlsxDB orm.XlsxDB

	// fetch the xlsx
	_, err := db.First(&xlsxDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlsxDB.CopyBasicFieldsFromXlsx_WOP(&input.Xlsx_WOP)
	xlsxDB.XlsxPointersEncoding = input.XlsxPointersEncoding

	db, _ = db.Model(&xlsxDB)
	_, err = db.Updates(&xlsxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlsxNew := new(models.Xlsx)
	xlsxDB.CopyBasicFieldsToXlsx(xlsxNew)

	// redeem pointers
	xlsxDB.DecodePointers(backRepo, xlsxNew)

	// get stage instance from DB instance, and call callback function
	xlsxOld := backRepo.BackRepoXlsx.Map_XlsxDBID_XlsxPtr[xlsxDB.ID]
	if xlsxOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(xlsxOld, xlsxNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlsxDB
	writeJSON(w, http.StatusOK, xlsxDB)
}
