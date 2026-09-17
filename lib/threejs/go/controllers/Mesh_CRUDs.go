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
var __Mesh__dummysDeclaration__ models.Mesh
var _ = __Mesh__dummysDeclaration__
var __Mesh_time__dummyDeclaration time.Duration
var _ = __Mesh_time__dummyDeclaration

var mutexMesh sync.Mutex

// An MeshID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMesh
type MeshID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMesh
type MeshInput struct {
	// The Mesh to submit or modify
	// in: body
	Mesh *orm.MeshAPI
}

// UpdateMesh
//
// swagger:route PATCH /meshs/{ID} meshs updateMesh
//
// # Update a mesh
//
// Responses:
// default: genericError
//
//	200: meshDBResponse
func (controller *Controller) UpdateMesh(w http.ResponseWriter, r *http.Request) {

	mutexMesh.Lock()
	defer mutexMesh.Unlock()

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
	db := backRepo.BackRepoMesh.GetDB()

	// Validate input
	var input orm.MeshAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshDB orm.MeshDB

	// fetch the mesh
	_, err := db.First(&meshDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshDB.CopyBasicFieldsFromMesh_WOP(&input.Mesh_WOP)
	meshDB.MeshPointersEncoding = input.MeshPointersEncoding

	db, _ = db.Model(&meshDB)
	_, err = db.Updates(&meshDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshNew := new(models.Mesh)
	meshDB.CopyBasicFieldsToMesh(meshNew)

	// redeem pointers
	meshDB.DecodePointers(backRepo, meshNew)

	// get stage instance from DB instance, and call callback function
	meshOld := backRepo.BackRepoMesh.Map_MeshDBID_MeshPtr[meshDB.ID]
	if meshOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(meshOld, meshNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshDB
	writeJSON(w, http.StatusOK, meshDB)
}
