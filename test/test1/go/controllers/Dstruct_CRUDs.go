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
var __Dstruct__dummysDeclaration__ models.Dstruct
var _ = __Dstruct__dummysDeclaration__
var __Dstruct_time__dummyDeclaration time.Duration
var _ = __Dstruct_time__dummyDeclaration

var mutexDstruct sync.Mutex

// An DstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDstruct updateDstruct deleteDstruct
type DstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDstruct updateDstruct
type DstructInput struct {
	// The Dstruct to submit or modify
	// in: body
	Dstruct *orm.DstructAPI
}

// GetDstructs
//
// swagger:route GET /dstructs dstructs getDstructs
//
// # Get all dstructs
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) GetDstructs(w http.ResponseWriter, r *http.Request) {

	// source slice
	var dstructDBs []orm.DstructDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDstructs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	_, err := db.Find(&dstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dstructAPIs := make([]orm.DstructAPI, 0)

	// for each dstruct, update fields from the database nullable fields
	for idx := range dstructDBs {
		dstructDB := &dstructDBs[idx]
		_ = dstructDB
		var dstructAPI orm.DstructAPI

		// insertion point for updating fields
		dstructAPI.ID = dstructDB.ID
		dstructDB.CopyBasicFieldsToDstruct_WOP(&dstructAPI.Dstruct_WOP)
		dstructAPI.DstructPointersEncoding = dstructDB.DstructPointersEncoding
		dstructAPIs = append(dstructAPIs, dstructAPI)
	}

	writeJSON(w, http.StatusOK, dstructAPIs)
}

// PostDstruct
//
// swagger:route POST /dstructs dstructs postDstruct
//
// Creates a dstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDstruct(w http.ResponseWriter, r *http.Request) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDstructs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Validate input
	var input orm.DstructAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dstruct
	dstructDB := orm.DstructDB{}
	dstructDB.DstructPointersEncoding = input.DstructPointersEncoding
	dstructDB.CopyBasicFieldsFromDstruct_WOP(&input.Dstruct_WOP)

	_, err = db.Create(&dstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDstruct.CheckoutPhaseOneInstance(&dstructDB)
	dstruct := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]

	if dstruct != nil {
		backRepo.GetStage().AfterCreateFromFront(dstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, dstructDB)
}

// GetDstruct
//
// swagger:route GET /dstructs/{ID} dstructs getDstruct
//
// Gets the details for a dstruct.
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) GetDstruct(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDstruct", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Get dstructDB in DB
	var dstructDB orm.DstructDB
	if _, err := db.First(&dstructDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var dstructAPI orm.DstructAPI
	dstructAPI.ID = dstructDB.ID
	dstructAPI.DstructPointersEncoding = dstructDB.DstructPointersEncoding
	dstructDB.CopyBasicFieldsToDstruct_WOP(&dstructAPI.Dstruct_WOP)

	writeJSON(w, http.StatusOK, dstructAPI)
}

// UpdateDstruct
//
// swagger:route PATCH /dstructs/{ID} dstructs updateDstruct
//
// # Update a dstruct
//
// Responses:
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) UpdateDstruct(w http.ResponseWriter, r *http.Request) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

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
	db := backRepo.BackRepoDstruct.GetDB()

	// Validate input
	var input orm.DstructAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dstructDB orm.DstructDB

	// fetch the dstruct
	_, err := db.First(&dstructDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dstructDB.CopyBasicFieldsFromDstruct_WOP(&input.Dstruct_WOP)
	dstructDB.DstructPointersEncoding = input.DstructPointersEncoding

	db, _ = db.Model(&dstructDB)
	_, err = db.Updates(&dstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dstructNew := new(models.Dstruct)
	dstructDB.CopyBasicFieldsToDstruct(dstructNew)

	// redeem pointers
	dstructDB.DecodePointers(backRepo, dstructNew)

	// get stage instance from DB instance, and call callback function
	dstructOld := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]
	if dstructOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(dstructOld, dstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dstructDB
	writeJSON(w, http.StatusOK, dstructDB)
}

// DeleteDstruct
//
// swagger:route DELETE /dstructs/{ID} dstructs deleteDstruct
//
// # Delete a dstruct
//
// default: genericError
//
//	200: dstructDBResponse
func (controller *Controller) DeleteDstruct(w http.ResponseWriter, r *http.Request) {

	mutexDstruct.Lock()
	defer mutexDstruct.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDstruct", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/test1/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDstruct.GetDB()

	// Get model if exist
	var dstructDB orm.DstructDB
	if _, err := db.First(&dstructDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&dstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	dstructDeleted := new(models.Dstruct)
	dstructDB.CopyBasicFieldsToDstruct(dstructDeleted)

	// get stage instance from DB instance, and call callback function
	dstructStaged := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[dstructDB.ID]
	if dstructStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(dstructStaged, dstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
