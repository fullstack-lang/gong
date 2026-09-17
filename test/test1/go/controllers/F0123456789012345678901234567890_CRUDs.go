// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test1/go/models"
	"github.com/fullstack-lang/gong/test/test1/go/orm"
)

// declaration in order to justify use of the models import
var __F0123456789012345678901234567890__dummysDeclaration__ models.F0123456789012345678901234567890
var _ = __F0123456789012345678901234567890__dummysDeclaration__
var __F0123456789012345678901234567890_time__dummyDeclaration time.Duration
var _ = __F0123456789012345678901234567890_time__dummyDeclaration

var mutexF0123456789012345678901234567890 sync.Mutex

// An F0123456789012345678901234567890ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateF0123456789012345678901234567890
type F0123456789012345678901234567890ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// F0123456789012345678901234567890Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateF0123456789012345678901234567890
type F0123456789012345678901234567890Input struct {
	// The F0123456789012345678901234567890 to submit or modify
	// in: body
	F0123456789012345678901234567890 *orm.F0123456789012345678901234567890API
}

// UpdateF0123456789012345678901234567890
//
// swagger:route PATCH /f0123456789012345678901234567890s/{ID} f0123456789012345678901234567890s updateF0123456789012345678901234567890
//
// # Update a f0123456789012345678901234567890
//
// Responses:
// default: genericError
//
//	200: f0123456789012345678901234567890DBResponse
func (controller *Controller) UpdateF0123456789012345678901234567890(w http.ResponseWriter, r *http.Request) {

	mutexF0123456789012345678901234567890.Lock()
	defer mutexF0123456789012345678901234567890.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	// Validate input
	var input orm.F0123456789012345678901234567890API
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var f0123456789012345678901234567890DB orm.F0123456789012345678901234567890DB

	// fetch the f0123456789012345678901234567890
	_, err := db.First(&f0123456789012345678901234567890DB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	f0123456789012345678901234567890DB.CopyBasicFieldsFromF0123456789012345678901234567890_WOP(&input.F0123456789012345678901234567890_WOP)
	f0123456789012345678901234567890DB.F0123456789012345678901234567890PointersEncoding = input.F0123456789012345678901234567890PointersEncoding

	db, _ = db.Model(&f0123456789012345678901234567890DB)
	_, err = db.Updates(&f0123456789012345678901234567890DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	f0123456789012345678901234567890New := new(models.F0123456789012345678901234567890)
	f0123456789012345678901234567890DB.CopyBasicFieldsToF0123456789012345678901234567890(f0123456789012345678901234567890New)

	// redeem pointers
	f0123456789012345678901234567890DB.DecodePointers(backRepo, f0123456789012345678901234567890New)

	// get stage instance from DB instance, and call callback function
	f0123456789012345678901234567890Old := backRepo.BackRepoF0123456789012345678901234567890.Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890Ptr[f0123456789012345678901234567890DB.ID]
	if f0123456789012345678901234567890Old != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(f0123456789012345678901234567890Old, f0123456789012345678901234567890New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the f0123456789012345678901234567890DB
	writeJSON(w, http.StatusOK, f0123456789012345678901234567890DB)
}
