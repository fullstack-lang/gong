// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/orm"
)

// declaration in order to justify use of the models import
var __A__dummysDeclaration__ models.A
var _ = __A__dummysDeclaration__
var __A_time__dummyDeclaration time.Duration
var _ = __A_time__dummyDeclaration

var mutexA sync.Mutex

// An AID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA updateA deleteA
type AID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA updateA
type AInput struct {
	// The A to submit or modify
	// in: body
	A *orm.AAPI
}

// GetAs
//
// swagger:route GET /as as getAs
//
// # Get all as
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) GetAs(w http.ResponseWriter, r *http.Request) {

	// source slice
	var aDBs []orm.ADB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA.GetDB()

	_, err := db.Find(&aDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	aAPIs := make([]orm.AAPI, 0)

	// for each a, update fields from the database nullable fields
	for idx := range aDBs {
		aDB := &aDBs[idx]
		_ = aDB
		var aAPI orm.AAPI

		// insertion point for updating fields
		aAPI.ID = aDB.ID
		aDB.CopyBasicFieldsToA_WOP(&aAPI.A_WOP)
		aAPI.APointersEncoding = aDB.APointersEncoding
		aAPIs = append(aAPIs, aAPI)
	}

	writeJSON(w, http.StatusOK, aAPIs)
}

// PostA
//
// swagger:route POST /as as postA
//
// Creates a a
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA(w http.ResponseWriter, r *http.Request) {

	mutexA.Lock()
	defer mutexA.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA.GetDB()

	// Validate input
	var input orm.AAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a
	aDB := orm.ADB{}
	aDB.APointersEncoding = input.APointersEncoding
	aDB.CopyBasicFieldsFromA_WOP(&input.A_WOP)

	_, err = db.Create(&aDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA.CheckoutPhaseOneInstance(&aDB)
	a := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]

	if a != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, aDB)
}

// GetA
//
// swagger:route GET /as/{ID} as getA
//
// Gets the details for a a.
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) GetA(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA.GetDB()

	// Get aDB in DB
	var aDB orm.ADB
	if _, err := db.First(&aDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var aAPI orm.AAPI
	aAPI.ID = aDB.ID
	aAPI.APointersEncoding = aDB.APointersEncoding
	aDB.CopyBasicFieldsToA_WOP(&aAPI.A_WOP)

	writeJSON(w, http.StatusOK, aAPI)
}

// UpdateA
//
// swagger:route PATCH /as/{ID} as updateA
//
// # Update a a
//
// Responses:
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) UpdateA(w http.ResponseWriter, r *http.Request) {

	mutexA.Lock()
	defer mutexA.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA.GetDB()

	// Validate input
	var input orm.AAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var aDB orm.ADB

	// fetch the a
	_, err := db.First(&aDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	aDB.CopyBasicFieldsFromA_WOP(&input.A_WOP)
	aDB.APointersEncoding = input.APointersEncoding

	db, _ = db.Model(&aDB)
	_, err = db.Updates(&aDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	aNew := new(models.A)
	aDB.CopyBasicFieldsToA(aNew)

	// redeem pointers
	aDB.DecodePointers(backRepo, aNew)

	// get stage instance from DB instance, and call callback function
	aOld := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]
	if aOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), aOld, aNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the aDB
	writeJSON(w, http.StatusOK, aDB)
}

// DeleteA
//
// swagger:route DELETE /as/{ID} as deleteA
//
// # Delete a a
//
// default: genericError
//
//	200: aDBResponse
func (controller *Controller) DeleteA(w http.ResponseWriter, r *http.Request) {

	mutexA.Lock()
	defer mutexA.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA.GetDB()

	// Get model if exist
	var aDB orm.ADB
	if _, err := db.First(&aDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&aDB)

	// get an instance (not staged) from DB instance, and call callback function
	aDeleted := new(models.A)
	aDB.CopyBasicFieldsToA(aDeleted)

	// get stage instance from DB instance, and call callback function
	aStaged := backRepo.BackRepoA.Map_ADBID_APtr[aDB.ID]
	if aStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), aStaged, aDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
