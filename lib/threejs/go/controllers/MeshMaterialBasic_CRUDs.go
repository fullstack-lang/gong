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
var __MeshMaterialBasic__dummysDeclaration__ models.MeshMaterialBasic
var _ = __MeshMaterialBasic__dummysDeclaration__
var __MeshMaterialBasic_time__dummyDeclaration time.Duration
var _ = __MeshMaterialBasic_time__dummyDeclaration

var mutexMeshMaterialBasic sync.Mutex

// An MeshMaterialBasicID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateMeshMaterialBasic
type MeshMaterialBasicID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshMaterialBasicInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMeshMaterialBasic
type MeshMaterialBasicInput struct {
	// The MeshMaterialBasic to submit or modify
	// in: body
	MeshMaterialBasic *orm.MeshMaterialBasicAPI
}

// UpdateMeshMaterialBasic
//
// swagger:route PATCH /meshmaterialbasics/{ID} meshmaterialbasics updateMeshMaterialBasic
//
// # Update a meshmaterialbasic
//
// Responses:
// default: genericError
//
//	200: meshmaterialbasicDBResponse
func (controller *Controller) UpdateMeshMaterialBasic(w http.ResponseWriter, r *http.Request) {

	mutexMeshMaterialBasic.Lock()
	defer mutexMeshMaterialBasic.Unlock()

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
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	// Validate input
	var input orm.MeshMaterialBasicAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshmaterialbasicDB orm.MeshMaterialBasicDB

	// fetch the meshmaterialbasic
	_, err := db.First(&meshmaterialbasicDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshmaterialbasicDB.CopyBasicFieldsFromMeshMaterialBasic_WOP(&input.MeshMaterialBasic_WOP)
	meshmaterialbasicDB.MeshMaterialBasicPointersEncoding = input.MeshMaterialBasicPointersEncoding

	db, _ = db.Model(&meshmaterialbasicDB)
	_, err = db.Updates(&meshmaterialbasicDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshmaterialbasicNew := new(models.MeshMaterialBasic)
	meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic(meshmaterialbasicNew)

	// redeem pointers
	meshmaterialbasicDB.DecodePointers(backRepo, meshmaterialbasicNew)

	// get stage instance from DB instance, and call callback function
	meshmaterialbasicOld := backRepo.BackRepoMeshMaterialBasic.Map_MeshMaterialBasicDBID_MeshMaterialBasicPtr[meshmaterialbasicDB.ID]
	if meshmaterialbasicOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(meshmaterialbasicOld, meshmaterialbasicNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshmaterialbasicDB
	writeJSON(w, http.StatusOK, meshmaterialbasicDB)
}
