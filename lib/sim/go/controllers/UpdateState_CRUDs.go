// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __UpdateState__dummysDeclaration__ models.UpdateState
var __UpdateState_time__dummyDeclaration time.Duration

var mutexUpdateState sync.Mutex

// An UpdateStateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUpdateState updateUpdateState deleteUpdateState
type UpdateStateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UpdateStateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUpdateState updateUpdateState
type UpdateStateInput struct {
	// The UpdateState to submit or modify
	// in: body
	UpdateState *orm.UpdateStateAPI
}

// GetUpdateStates
//
// swagger:route GET /updatestates updatestates getUpdateStates
//
// # Get all updatestates
//
// Responses:
// default: genericError
//
//	200: updatestateDBResponse
func (controller *Controller) GetUpdateStates(c *gin.Context) {

	// source slice
	var updatestateDBs []orm.UpdateStateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUpdateStates", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	_, err := db.Find(&updatestateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	updatestateAPIs := make([]orm.UpdateStateAPI, 0)

	// for each updatestate, update fields from the database nullable fields
	for idx := range updatestateDBs {
		updatestateDB := &updatestateDBs[idx]
		_ = updatestateDB
		var updatestateAPI orm.UpdateStateAPI

		// insertion point for updating fields
		updatestateAPI.ID = updatestateDB.ID
		updatestateDB.CopyBasicFieldsToUpdateState_WOP(&updatestateAPI.UpdateState_WOP)
		updatestateAPI.UpdateStatePointersEncoding = updatestateDB.UpdateStatePointersEncoding
		updatestateAPIs = append(updatestateAPIs, updatestateAPI)
	}

	c.JSON(http.StatusOK, updatestateAPIs)
}

// PostUpdateState
//
// swagger:route POST /updatestates updatestates postUpdateState
//
// Creates a updatestate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUpdateState(c *gin.Context) {

	mutexUpdateState.Lock()
	defer mutexUpdateState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUpdateStates", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	// Validate input
	var input orm.UpdateStateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create updatestate
	updatestateDB := orm.UpdateStateDB{}
	updatestateDB.UpdateStatePointersEncoding = input.UpdateStatePointersEncoding
	updatestateDB.CopyBasicFieldsFromUpdateState_WOP(&input.UpdateState_WOP)

	_, err = db.Create(&updatestateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUpdateState.CheckoutPhaseOneInstance(&updatestateDB)
	updatestate := backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStatePtr[updatestateDB.ID]

	if updatestate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), updatestate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, updatestateDB)
}

// GetUpdateState
//
// swagger:route GET /updatestates/{ID} updatestates getUpdateState
//
// Gets the details for a updatestate.
//
// Responses:
// default: genericError
//
//	200: updatestateDBResponse
func (controller *Controller) GetUpdateState(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUpdateState", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	// Get updatestateDB in DB
	var updatestateDB orm.UpdateStateDB
	if _, err := db.First(&updatestateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var updatestateAPI orm.UpdateStateAPI
	updatestateAPI.ID = updatestateDB.ID
	updatestateAPI.UpdateStatePointersEncoding = updatestateDB.UpdateStatePointersEncoding
	updatestateDB.CopyBasicFieldsToUpdateState_WOP(&updatestateAPI.UpdateState_WOP)

	c.JSON(http.StatusOK, updatestateAPI)
}

// UpdateUpdateState
//
// swagger:route PATCH /updatestates/{ID} updatestates updateUpdateState
//
// # Update a updatestate
//
// Responses:
// default: genericError
//
//	200: updatestateDBResponse
func (controller *Controller) UpdateUpdateState(c *gin.Context) {

	mutexUpdateState.Lock()
	defer mutexUpdateState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUpdateState", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	// Validate input
	var input orm.UpdateStateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var updatestateDB orm.UpdateStateDB

	// fetch the updatestate
	_, err := db.First(&updatestateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	updatestateDB.CopyBasicFieldsFromUpdateState_WOP(&input.UpdateState_WOP)
	updatestateDB.UpdateStatePointersEncoding = input.UpdateStatePointersEncoding

	db, _ = db.Model(&updatestateDB)
	_, err = db.Updates(&updatestateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	updatestateNew := new(models.UpdateState)
	updatestateDB.CopyBasicFieldsToUpdateState(updatestateNew)

	// redeem pointers
	updatestateDB.DecodePointers(backRepo, updatestateNew)

	// get stage instance from DB instance, and call callback function
	updatestateOld := backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStatePtr[updatestateDB.ID]
	if updatestateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), updatestateOld, updatestateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the updatestateDB
	c.JSON(http.StatusOK, updatestateDB)
}

// DeleteUpdateState
//
// swagger:route DELETE /updatestates/{ID} updatestates deleteUpdateState
//
// # Delete a updatestate
//
// default: genericError
//
//	200: updatestateDBResponse
func (controller *Controller) DeleteUpdateState(c *gin.Context) {

	mutexUpdateState.Lock()
	defer mutexUpdateState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUpdateState", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoUpdateState.GetDB()

	// Get model if exist
	var updatestateDB orm.UpdateStateDB
	if _, err := db.First(&updatestateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&updatestateDB)

	// get an instance (not staged) from DB instance, and call callback function
	updatestateDeleted := new(models.UpdateState)
	updatestateDB.CopyBasicFieldsToUpdateState(updatestateDeleted)

	// get stage instance from DB instance, and call callback function
	updatestateStaged := backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStatePtr[updatestateDB.ID]
	if updatestateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), updatestateStaged, updatestateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
