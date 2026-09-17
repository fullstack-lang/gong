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
var __Vector3__dummysDeclaration__ models.Vector3
var _ = __Vector3__dummysDeclaration__
var __Vector3_time__dummyDeclaration time.Duration
var _ = __Vector3_time__dummyDeclaration

var mutexVector3 sync.Mutex

// An Vector3ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateVector3
type Vector3ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Vector3Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateVector3
type Vector3Input struct {
	// The Vector3 to submit or modify
	// in: body
	Vector3 *orm.Vector3API
}

// UpdateVector3
//
// swagger:route PATCH /vector3s/{ID} vector3s updateVector3
//
// # Update a vector3
//
// Responses:
// default: genericError
//
//	200: vector3DBResponse
func (controller *Controller) UpdateVector3(w http.ResponseWriter, r *http.Request) {

	mutexVector3.Lock()
	defer mutexVector3.Unlock()

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
	db := backRepo.BackRepoVector3.GetDB()

	// Validate input
	var input orm.Vector3API
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var vector3DB orm.Vector3DB

	// fetch the vector3
	_, err := db.First(&vector3DB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	vector3DB.CopyBasicFieldsFromVector3_WOP(&input.Vector3_WOP)
	vector3DB.Vector3PointersEncoding = input.Vector3PointersEncoding

	db, _ = db.Model(&vector3DB)
	_, err = db.Updates(&vector3DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	vector3New := new(models.Vector3)
	vector3DB.CopyBasicFieldsToVector3(vector3New)

	// redeem pointers
	vector3DB.DecodePointers(backRepo, vector3New)

	// get stage instance from DB instance, and call callback function
	vector3Old := backRepo.BackRepoVector3.Map_Vector3DBID_Vector3Ptr[vector3DB.ID]
	if vector3Old != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(vector3Old, vector3New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the vector3DB
	writeJSON(w, http.StatusOK, vector3DB)
}
