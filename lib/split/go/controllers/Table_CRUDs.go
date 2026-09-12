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
var __Table__dummysDeclaration__ models.Table
var _ = __Table__dummysDeclaration__
var __Table_time__dummyDeclaration time.Duration
var _ = __Table_time__dummyDeclaration

var mutexTable sync.Mutex

// An TableID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTable updateTable deleteTable
type TableID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTable updateTable
type TableInput struct {
	// The Table to submit or modify
	// in: body
	Table *orm.TableAPI
}

// GetTables
//
// swagger:route GET /tables tables getTables
//
// # Get all tables
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) GetTables(w http.ResponseWriter, r *http.Request) {

	// source slice
	var tableDBs []orm.TableDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTables", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTable.GetDB()

	_, err := db.Find(&tableDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tableAPIs := make([]orm.TableAPI, 0)

	// for each table, update fields from the database nullable fields
	for idx := range tableDBs {
		tableDB := &tableDBs[idx]
		_ = tableDB
		var tableAPI orm.TableAPI

		// insertion point for updating fields
		tableAPI.ID = tableDB.ID
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableAPIs = append(tableAPIs, tableAPI)
	}

	writeJSON(w, http.StatusOK, tableAPIs)
}

// PostTable
//
// swagger:route POST /tables tables postTable
//
// Creates a table
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTable(w http.ResponseWriter, r *http.Request) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTables", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Validate input
	var input orm.TableAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create table
	tableDB := orm.TableDB{}
	tableDB.TablePointersEncoding = input.TablePointersEncoding
	tableDB.CopyBasicFieldsFromTable_WOP(&input.Table_WOP)

	_, err = db.Create(&tableDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTable.CheckoutPhaseOneInstance(&tableDB)
	table := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]

	if table != nil {
		backRepo.GetStage().AfterCreateFromFront(table)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, tableDB)
}

// GetTable
//
// swagger:route GET /tables/{ID} tables getTable
//
// Gets the details for a table.
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) GetTable(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTable", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Get tableDB in DB
	var tableDB orm.TableDB
	if _, err := db.First(&tableDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var tableAPI orm.TableAPI
	tableAPI.ID = tableDB.ID
	tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
	tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

	writeJSON(w, http.StatusOK, tableAPI)
}

// UpdateTable
//
// swagger:route PATCH /tables/{ID} tables updateTable
//
// # Update a table
//
// Responses:
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) UpdateTable(w http.ResponseWriter, r *http.Request) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

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
	db := backRepo.BackRepoTable.GetDB()

	// Validate input
	var input orm.TableAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tableDB orm.TableDB

	// fetch the table
	_, err := db.First(&tableDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tableDB.CopyBasicFieldsFromTable_WOP(&input.Table_WOP)
	tableDB.TablePointersEncoding = input.TablePointersEncoding

	db, _ = db.Model(&tableDB)
	_, err = db.Updates(&tableDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tableNew := new(models.Table)
	tableDB.CopyBasicFieldsToTable(tableNew)

	// redeem pointers
	tableDB.DecodePointers(backRepo, tableNew)

	// get stage instance from DB instance, and call callback function
	tableOld := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]
	if tableOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(tableOld, tableNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tableDB
	writeJSON(w, http.StatusOK, tableDB)
}

// DeleteTable
//
// swagger:route DELETE /tables/{ID} tables deleteTable
//
// # Delete a table
//
// default: genericError
//
//	200: tableDBResponse
func (controller *Controller) DeleteTable(w http.ResponseWriter, r *http.Request) {

	mutexTable.Lock()
	defer mutexTable.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTable", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTable.GetDB()

	// Get model if exist
	var tableDB orm.TableDB
	if _, err := db.First(&tableDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tableDB)

	// get an instance (not staged) from DB instance, and call callback function
	tableDeleted := new(models.Table)
	tableDB.CopyBasicFieldsToTable(tableDeleted)

	// get stage instance from DB instance, and call callback function
	tableStaged := backRepo.BackRepoTable.Map_TableDBID_TablePtr[tableDB.ID]
	if tableStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(tableStaged, tableDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
