// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __UmlState__dummysDeclaration__ models.UmlState
var __UmlState_time__dummyDeclaration time.Duration

var mutexUmlState sync.Mutex

// An UmlStateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUmlState updateUmlState deleteUmlState
type UmlStateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UmlStateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUmlState updateUmlState
type UmlStateInput struct {
	// The UmlState to submit or modify
	// in: body
	UmlState *orm.UmlStateAPI
}

// GetUmlStates
//
// swagger:route GET /umlstates umlstates getUmlStates
//
// # Get all umlstates
//
// Responses:
// default: genericError
//
//	200: umlstateDBResponse
func (controller *Controller) GetUmlStates(c *gin.Context) {

	// source slice
	var umlstateDBs []orm.UmlStateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUmlStates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlState.GetDB()

	_, err := db.Find(&umlstateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	umlstateAPIs := make([]orm.UmlStateAPI, 0)

	// for each umlstate, update fields from the database nullable fields
	for idx := range umlstateDBs {
		umlstateDB := &umlstateDBs[idx]
		_ = umlstateDB
		var umlstateAPI orm.UmlStateAPI

		// insertion point for updating fields
		umlstateAPI.ID = umlstateDB.ID
		umlstateDB.CopyBasicFieldsToUmlState_WOP(&umlstateAPI.UmlState_WOP)
		umlstateAPI.UmlStatePointersEncoding = umlstateDB.UmlStatePointersEncoding
		umlstateAPIs = append(umlstateAPIs, umlstateAPI)
	}

	c.JSON(http.StatusOK, umlstateAPIs)
}

// PostUmlState
//
// swagger:route POST /umlstates umlstates postUmlState
//
// Creates a umlstate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUmlState(c *gin.Context) {

	mutexUmlState.Lock()
	defer mutexUmlState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUmlStates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlState.GetDB()

	// Validate input
	var input orm.UmlStateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create umlstate
	umlstateDB := orm.UmlStateDB{}
	umlstateDB.UmlStatePointersEncoding = input.UmlStatePointersEncoding
	umlstateDB.CopyBasicFieldsFromUmlState_WOP(&input.UmlState_WOP)

	_, err = db.Create(&umlstateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUmlState.CheckoutPhaseOneInstance(&umlstateDB)
	umlstate := backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]

	if umlstate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), umlstate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, umlstateDB)
}

// GetUmlState
//
// swagger:route GET /umlstates/{ID} umlstates getUmlState
//
// Gets the details for a umlstate.
//
// Responses:
// default: genericError
//
//	200: umlstateDBResponse
func (controller *Controller) GetUmlState(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUmlState", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlState.GetDB()

	// Get umlstateDB in DB
	var umlstateDB orm.UmlStateDB
	if _, err := db.First(&umlstateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var umlstateAPI orm.UmlStateAPI
	umlstateAPI.ID = umlstateDB.ID
	umlstateAPI.UmlStatePointersEncoding = umlstateDB.UmlStatePointersEncoding
	umlstateDB.CopyBasicFieldsToUmlState_WOP(&umlstateAPI.UmlState_WOP)

	c.JSON(http.StatusOK, umlstateAPI)
}

// UpdateUmlState
//
// swagger:route PATCH /umlstates/{ID} umlstates updateUmlState
//
// # Update a umlstate
//
// Responses:
// default: genericError
//
//	200: umlstateDBResponse
func (controller *Controller) UpdateUmlState(c *gin.Context) {

	mutexUmlState.Lock()
	defer mutexUmlState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUmlState", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlState.GetDB()

	// Validate input
	var input orm.UmlStateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var umlstateDB orm.UmlStateDB

	// fetch the umlstate
	_, err := db.First(&umlstateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	umlstateDB.CopyBasicFieldsFromUmlState_WOP(&input.UmlState_WOP)
	umlstateDB.UmlStatePointersEncoding = input.UmlStatePointersEncoding

	db, _ = db.Model(&umlstateDB)
	_, err = db.Updates(&umlstateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	umlstateNew := new(models.UmlState)
	umlstateDB.CopyBasicFieldsToUmlState(umlstateNew)

	// redeem pointers
	umlstateDB.DecodePointers(backRepo, umlstateNew)

	// get stage instance from DB instance, and call callback function
	umlstateOld := backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]
	if umlstateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), umlstateOld, umlstateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the umlstateDB
	c.JSON(http.StatusOK, umlstateDB)
}

// DeleteUmlState
//
// swagger:route DELETE /umlstates/{ID} umlstates deleteUmlState
//
// # Delete a umlstate
//
// default: genericError
//
//	200: umlstateDBResponse
func (controller *Controller) DeleteUmlState(c *gin.Context) {

	mutexUmlState.Lock()
	defer mutexUmlState.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUmlState", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlState.GetDB()

	// Get model if exist
	var umlstateDB orm.UmlStateDB
	if _, err := db.First(&umlstateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&umlstateDB)

	// get an instance (not staged) from DB instance, and call callback function
	umlstateDeleted := new(models.UmlState)
	umlstateDB.CopyBasicFieldsToUmlState(umlstateDeleted)

	// get stage instance from DB instance, and call callback function
	umlstateStaged := backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]
	if umlstateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), umlstateStaged, umlstateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
