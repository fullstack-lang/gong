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
var __Triangle__dummysDeclaration__ models.Triangle
var _ = __Triangle__dummysDeclaration__
var __Triangle_time__dummyDeclaration time.Duration
var _ = __Triangle_time__dummyDeclaration

var mutexTriangle sync.Mutex

// An TriangleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTriangle
type TriangleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TriangleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTriangle
type TriangleInput struct {
	// The Triangle to submit or modify
	// in: body
	Triangle *orm.TriangleAPI
}

// UpdateTriangle
//
// swagger:route PATCH /triangles/{ID} triangles updateTriangle
//
// # Update a triangle
//
// Responses:
// default: genericError
//
//	200: triangleDBResponse
func (controller *Controller) UpdateTriangle(w http.ResponseWriter, r *http.Request) {

	mutexTriangle.Lock()
	defer mutexTriangle.Unlock()

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
	db := backRepo.BackRepoTriangle.GetDB()

	// Validate input
	var input orm.TriangleAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var triangleDB orm.TriangleDB

	// fetch the triangle
	_, err := db.First(&triangleDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	triangleDB.CopyBasicFieldsFromTriangle_WOP(&input.Triangle_WOP)
	triangleDB.TrianglePointersEncoding = input.TrianglePointersEncoding

	db, _ = db.Model(&triangleDB)
	_, err = db.Updates(&triangleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	triangleNew := new(models.Triangle)
	triangleDB.CopyBasicFieldsToTriangle(triangleNew)

	// redeem pointers
	triangleDB.DecodePointers(backRepo, triangleNew)

	// get stage instance from DB instance, and call callback function
	triangleOld := backRepo.BackRepoTriangle.Map_TriangleDBID_TrianglePtr[triangleDB.ID]
	if triangleOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(triangleOld, triangleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the triangleDB
	writeJSON(w, http.StatusOK, triangleDB)
}
