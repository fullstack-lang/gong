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
var __Vector2__dummysDeclaration__ models.Vector2
var _ = __Vector2__dummysDeclaration__
var __Vector2_time__dummyDeclaration time.Duration
var _ = __Vector2_time__dummyDeclaration

var mutexVector2 sync.Mutex

// An Vector2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateVector2
type Vector2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Vector2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateVector2
type Vector2Input struct {
	// The Vector2 to submit or modify
	// in: body
	Vector2 *orm.Vector2API
}

// UpdateVector2
//
// swagger:route PATCH /vector2s/{ID} vector2s updateVector2
//
// # Update a vector2
//
// Responses:
// default: genericError
//
//	200: vector2DBResponse
func (controller *Controller) UpdateVector2(w http.ResponseWriter, r *http.Request) {

	mutexVector2.Lock()
	defer mutexVector2.Unlock()

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
	db := backRepo.BackRepoVector2.GetDB()

	// Validate input
	var input orm.Vector2API
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var vector2DB orm.Vector2DB

	// fetch the vector2
	_, err := db.First(&vector2DB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	vector2DB.CopyBasicFieldsFromVector2_WOP(&input.Vector2_WOP)
	vector2DB.Vector2PointersEncoding = input.Vector2PointersEncoding

	db, _ = db.Model(&vector2DB)
	_, err = db.Updates(&vector2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	vector2New := new(models.Vector2)
	vector2DB.CopyBasicFieldsToVector2(vector2New)

	// redeem pointers
	vector2DB.DecodePointers(backRepo, vector2New)

	// get stage instance from DB instance, and call callback function
	vector2Old := backRepo.BackRepoVector2.Map_Vector2DBID_Vector2Ptr[vector2DB.ID]
	if vector2Old != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(vector2Old, vector2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the vector2DB
	writeJSON(w, http.StatusOK, vector2DB)
}
