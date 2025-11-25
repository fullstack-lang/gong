// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Transition__dummysDeclaration__ models.Transition
var __Transition_time__dummyDeclaration time.Duration

var mutexTransition sync.Mutex

// An TransitionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTransition updateTransition deleteTransition
type TransitionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TransitionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTransition updateTransition
type TransitionInput struct {
	// The Transition to submit or modify
	// in: body
	Transition *orm.TransitionAPI
}

// GetTransitions
//
// swagger:route GET /transitions transitions getTransitions
//
// # Get all transitions
//
// Responses:
// default: genericError
//
//	200: transitionDBResponse
func (controller *Controller) GetTransitions(c *gin.Context) {

	// source slice
	var transitionDBs []orm.TransitionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTransitions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition.GetDB()

	_, err := db.Find(&transitionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	transitionAPIs := make([]orm.TransitionAPI, 0)

	// for each transition, update fields from the database nullable fields
	for idx := range transitionDBs {
		transitionDB := &transitionDBs[idx]
		_ = transitionDB
		var transitionAPI orm.TransitionAPI

		// insertion point for updating fields
		transitionAPI.ID = transitionDB.ID
		transitionDB.CopyBasicFieldsToTransition_WOP(&transitionAPI.Transition_WOP)
		transitionAPI.TransitionPointersEncoding = transitionDB.TransitionPointersEncoding
		transitionAPIs = append(transitionAPIs, transitionAPI)
	}

	c.JSON(http.StatusOK, transitionAPIs)
}

// PostTransition
//
// swagger:route POST /transitions transitions postTransition
//
// Creates a transition
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTransition(c *gin.Context) {

	mutexTransition.Lock()
	defer mutexTransition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTransitions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition.GetDB()

	// Validate input
	var input orm.TransitionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create transition
	transitionDB := orm.TransitionDB{}
	transitionDB.TransitionPointersEncoding = input.TransitionPointersEncoding
	transitionDB.CopyBasicFieldsFromTransition_WOP(&input.Transition_WOP)

	_, err = db.Create(&transitionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTransition.CheckoutPhaseOneInstance(&transitionDB)
	transition := backRepo.BackRepoTransition.Map_TransitionDBID_TransitionPtr[transitionDB.ID]

	if transition != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), transition)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, transitionDB)
}

// GetTransition
//
// swagger:route GET /transitions/{ID} transitions getTransition
//
// Gets the details for a transition.
//
// Responses:
// default: genericError
//
//	200: transitionDBResponse
func (controller *Controller) GetTransition(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTransition", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition.GetDB()

	// Get transitionDB in DB
	var transitionDB orm.TransitionDB
	if _, err := db.First(&transitionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var transitionAPI orm.TransitionAPI
	transitionAPI.ID = transitionDB.ID
	transitionAPI.TransitionPointersEncoding = transitionDB.TransitionPointersEncoding
	transitionDB.CopyBasicFieldsToTransition_WOP(&transitionAPI.Transition_WOP)

	c.JSON(http.StatusOK, transitionAPI)
}

// UpdateTransition
//
// swagger:route PATCH /transitions/{ID} transitions updateTransition
//
// # Update a transition
//
// Responses:
// default: genericError
//
//	200: transitionDBResponse
func (controller *Controller) UpdateTransition(c *gin.Context) {

	mutexTransition.Lock()
	defer mutexTransition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition.GetDB()

	// Validate input
	var input orm.TransitionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var transitionDB orm.TransitionDB

	// fetch the transition
	_, err := db.First(&transitionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	transitionDB.CopyBasicFieldsFromTransition_WOP(&input.Transition_WOP)
	transitionDB.TransitionPointersEncoding = input.TransitionPointersEncoding

	db, _ = db.Model(&transitionDB)
	_, err = db.Updates(&transitionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	transitionNew := new(models.Transition)
	transitionDB.CopyBasicFieldsToTransition(transitionNew)

	// redeem pointers
	transitionDB.DecodePointers(backRepo, transitionNew)

	// get stage instance from DB instance, and call callback function
	transitionOld := backRepo.BackRepoTransition.Map_TransitionDBID_TransitionPtr[transitionDB.ID]
	if transitionOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), transitionOld, transitionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the transitionDB
	c.JSON(http.StatusOK, transitionDB)
}

// DeleteTransition
//
// swagger:route DELETE /transitions/{ID} transitions deleteTransition
//
// # Delete a transition
//
// default: genericError
//
//	200: transitionDBResponse
func (controller *Controller) DeleteTransition(c *gin.Context) {

	mutexTransition.Lock()
	defer mutexTransition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTransition", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition.GetDB()

	// Get model if exist
	var transitionDB orm.TransitionDB
	if _, err := db.First(&transitionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&transitionDB)

	// get an instance (not staged) from DB instance, and call callback function
	transitionDeleted := new(models.Transition)
	transitionDB.CopyBasicFieldsToTransition(transitionDeleted)

	// get stage instance from DB instance, and call callback function
	transitionStaged := backRepo.BackRepoTransition.Map_TransitionDBID_TransitionPtr[transitionDB.ID]
	if transitionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), transitionStaged, transitionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
