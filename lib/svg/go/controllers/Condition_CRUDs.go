// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Condition__dummysDeclaration__ models.Condition
var _ = __Condition__dummysDeclaration__
var __Condition_time__dummyDeclaration time.Duration
var _ = __Condition_time__dummyDeclaration

var mutexCondition sync.Mutex

// An ConditionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCondition updateCondition deleteCondition
type ConditionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ConditionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCondition updateCondition
type ConditionInput struct {
	// The Condition to submit or modify
	// in: body
	Condition *orm.ConditionAPI
}

// GetConditions
//
// swagger:route GET /conditions conditions getConditions
//
// # Get all conditions
//
// Responses:
// default: genericError
//
//	200: conditionDBResponse
func (controller *Controller) GetConditions(c *gin.Context) {

	// source slice
	var conditionDBs []orm.ConditionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetConditions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	_, err := db.Find(&conditionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	conditionAPIs := make([]orm.ConditionAPI, 0)

	// for each condition, update fields from the database nullable fields
	for idx := range conditionDBs {
		conditionDB := &conditionDBs[idx]
		_ = conditionDB
		var conditionAPI orm.ConditionAPI

		// insertion point for updating fields
		conditionAPI.ID = conditionDB.ID
		conditionDB.CopyBasicFieldsToCondition_WOP(&conditionAPI.Condition_WOP)
		conditionAPI.ConditionPointersEncoding = conditionDB.ConditionPointersEncoding
		conditionAPIs = append(conditionAPIs, conditionAPI)
	}

	c.JSON(http.StatusOK, conditionAPIs)
}

// PostCondition
//
// swagger:route POST /conditions conditions postCondition
//
// Creates a condition
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCondition(c *gin.Context) {

	mutexCondition.Lock()
	defer mutexCondition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostConditions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	// Validate input
	var input orm.ConditionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create condition
	conditionDB := orm.ConditionDB{}
	conditionDB.ConditionPointersEncoding = input.ConditionPointersEncoding
	conditionDB.CopyBasicFieldsFromCondition_WOP(&input.Condition_WOP)

	_, err = db.Create(&conditionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCondition.CheckoutPhaseOneInstance(&conditionDB)
	condition := backRepo.BackRepoCondition.Map_ConditionDBID_ConditionPtr[conditionDB.ID]

	if condition != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), condition)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, conditionDB)
}

// GetCondition
//
// swagger:route GET /conditions/{ID} conditions getCondition
//
// Gets the details for a condition.
//
// Responses:
// default: genericError
//
//	200: conditionDBResponse
func (controller *Controller) GetCondition(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCondition", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	// Get conditionDB in DB
	var conditionDB orm.ConditionDB
	if _, err := db.First(&conditionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var conditionAPI orm.ConditionAPI
	conditionAPI.ID = conditionDB.ID
	conditionAPI.ConditionPointersEncoding = conditionDB.ConditionPointersEncoding
	conditionDB.CopyBasicFieldsToCondition_WOP(&conditionAPI.Condition_WOP)

	c.JSON(http.StatusOK, conditionAPI)
}

// UpdateCondition
//
// swagger:route PATCH /conditions/{ID} conditions updateCondition
//
// # Update a condition
//
// Responses:
// default: genericError
//
//	200: conditionDBResponse
func (controller *Controller) UpdateCondition(c *gin.Context) {

	mutexCondition.Lock()
	defer mutexCondition.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	// Validate input
	var input orm.ConditionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var conditionDB orm.ConditionDB

	// fetch the condition
	_, err := db.First(&conditionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	conditionDB.CopyBasicFieldsFromCondition_WOP(&input.Condition_WOP)
	conditionDB.ConditionPointersEncoding = input.ConditionPointersEncoding

	db, _ = db.Model(&conditionDB)
	_, err = db.Updates(&conditionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	conditionNew := new(models.Condition)
	conditionDB.CopyBasicFieldsToCondition(conditionNew)

	// redeem pointers
	conditionDB.DecodePointers(backRepo, conditionNew)

	// get stage instance from DB instance, and call callback function
	conditionOld := backRepo.BackRepoCondition.Map_ConditionDBID_ConditionPtr[conditionDB.ID]
	if conditionOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), conditionOld, conditionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the conditionDB
	c.JSON(http.StatusOK, conditionDB)
}

// DeleteCondition
//
// swagger:route DELETE /conditions/{ID} conditions deleteCondition
//
// # Delete a condition
//
// default: genericError
//
//	200: conditionDBResponse
func (controller *Controller) DeleteCondition(c *gin.Context) {

	mutexCondition.Lock()
	defer mutexCondition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCondition", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCondition.GetDB()

	// Get model if exist
	var conditionDB orm.ConditionDB
	if _, err := db.First(&conditionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&conditionDB)

	// get an instance (not staged) from DB instance, and call callback function
	conditionDeleted := new(models.Condition)
	conditionDB.CopyBasicFieldsToCondition(conditionDeleted)

	// get stage instance from DB instance, and call callback function
	conditionStaged := backRepo.BackRepoCondition.Map_ConditionDBID_ConditionPtr[conditionDB.ID]
	if conditionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), conditionStaged, conditionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
