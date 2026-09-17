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
var __MeshPhysicalMaterial__dummysDeclaration__ models.MeshPhysicalMaterial
var _ = __MeshPhysicalMaterial__dummysDeclaration__
var __MeshPhysicalMaterial_time__dummyDeclaration time.Duration
var _ = __MeshPhysicalMaterial_time__dummyDeclaration

var mutexMeshPhysicalMaterial sync.Mutex

// An MeshPhysicalMaterialID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMeshPhysicalMaterial
type MeshPhysicalMaterialID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshPhysicalMaterialInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMeshPhysicalMaterial
type MeshPhysicalMaterialInput struct {
	// The MeshPhysicalMaterial to submit or modify
	// in: body
	MeshPhysicalMaterial *orm.MeshPhysicalMaterialAPI
}

// UpdateMeshPhysicalMaterial
//
// swagger:route PATCH /meshphysicalmaterials/{ID} meshphysicalmaterials updateMeshPhysicalMaterial
//
// # Update a meshphysicalmaterial
//
// Responses:
// default: genericError
//
//	200: meshphysicalmaterialDBResponse
func (controller *Controller) UpdateMeshPhysicalMaterial(w http.ResponseWriter, r *http.Request) {

	mutexMeshPhysicalMaterial.Lock()
	defer mutexMeshPhysicalMaterial.Unlock()

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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	// Validate input
	var input orm.MeshPhysicalMaterialAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshphysicalmaterialDB orm.MeshPhysicalMaterialDB

	// fetch the meshphysicalmaterial
	_, err := db.First(&meshphysicalmaterialDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshphysicalmaterialDB.CopyBasicFieldsFromMeshPhysicalMaterial_WOP(&input.MeshPhysicalMaterial_WOP)
	meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding = input.MeshPhysicalMaterialPointersEncoding

	db, _ = db.Model(&meshphysicalmaterialDB)
	_, err = db.Updates(&meshphysicalmaterialDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshphysicalmaterialNew := new(models.MeshPhysicalMaterial)
	meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial(meshphysicalmaterialNew)

	// redeem pointers
	meshphysicalmaterialDB.DecodePointers(backRepo, meshphysicalmaterialNew)

	// get stage instance from DB instance, and call callback function
	meshphysicalmaterialOld := backRepo.BackRepoMeshPhysicalMaterial.Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialPtr[meshphysicalmaterialDB.ID]
	if meshphysicalmaterialOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(meshphysicalmaterialOld, meshphysicalmaterialNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshphysicalmaterialDB
	writeJSON(w, http.StatusOK, meshphysicalmaterialDB)
}
