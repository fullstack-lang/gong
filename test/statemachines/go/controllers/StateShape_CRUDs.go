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
var __StateShape__dummysDeclaration__ models.StateShape
var __StateShape_time__dummyDeclaration time.Duration

var mutexStateShape sync.Mutex

// An StateShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStateShape updateStateShape deleteStateShape
type StateShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StateShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStateShape updateStateShape
type StateShapeInput struct {
	// The StateShape to submit or modify
	// in: body
	StateShape *orm.StateShapeAPI
}

// GetStateShapes
//
// swagger:route GET /stateshapes stateshapes getStateShapes
//
// # Get all stateshapes
//
// Responses:
// default: genericError
//
//	200: stateshapeDBResponse
func (controller *Controller) GetStateShapes(c *gin.Context) {

	// source slice
	var stateshapeDBs []orm.StateShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStateShapes", "Name", stackPath)
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
	db := backRepo.BackRepoStateShape.GetDB()

	_, err := db.Find(&stateshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	stateshapeAPIs := make([]orm.StateShapeAPI, 0)

	// for each stateshape, update fields from the database nullable fields
	for idx := range stateshapeDBs {
		stateshapeDB := &stateshapeDBs[idx]
		_ = stateshapeDB
		var stateshapeAPI orm.StateShapeAPI

		// insertion point for updating fields
		stateshapeAPI.ID = stateshapeDB.ID
		stateshapeDB.CopyBasicFieldsToStateShape_WOP(&stateshapeAPI.StateShape_WOP)
		stateshapeAPI.StateShapePointersEncoding = stateshapeDB.StateShapePointersEncoding
		stateshapeAPIs = append(stateshapeAPIs, stateshapeAPI)
	}

	c.JSON(http.StatusOK, stateshapeAPIs)
}

// PostStateShape
//
// swagger:route POST /stateshapes stateshapes postStateShape
//
// Creates a stateshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStateShape(c *gin.Context) {

	mutexStateShape.Lock()
	defer mutexStateShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStateShapes", "Name", stackPath)
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
	db := backRepo.BackRepoStateShape.GetDB()

	// Validate input
	var input orm.StateShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create stateshape
	stateshapeDB := orm.StateShapeDB{}
	stateshapeDB.StateShapePointersEncoding = input.StateShapePointersEncoding
	stateshapeDB.CopyBasicFieldsFromStateShape_WOP(&input.StateShape_WOP)

	_, err = db.Create(&stateshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStateShape.CheckoutPhaseOneInstance(&stateshapeDB)
	stateshape := backRepo.BackRepoStateShape.Map_StateShapeDBID_StateShapePtr[stateshapeDB.ID]

	if stateshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), stateshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, stateshapeDB)
}

// GetStateShape
//
// swagger:route GET /stateshapes/{ID} stateshapes getStateShape
//
// Gets the details for a stateshape.
//
// Responses:
// default: genericError
//
//	200: stateshapeDBResponse
func (controller *Controller) GetStateShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStateShape", "Name", stackPath)
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
	db := backRepo.BackRepoStateShape.GetDB()

	// Get stateshapeDB in DB
	var stateshapeDB orm.StateShapeDB
	if _, err := db.First(&stateshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var stateshapeAPI orm.StateShapeAPI
	stateshapeAPI.ID = stateshapeDB.ID
	stateshapeAPI.StateShapePointersEncoding = stateshapeDB.StateShapePointersEncoding
	stateshapeDB.CopyBasicFieldsToStateShape_WOP(&stateshapeAPI.StateShape_WOP)

	c.JSON(http.StatusOK, stateshapeAPI)
}

// UpdateStateShape
//
// swagger:route PATCH /stateshapes/{ID} stateshapes updateStateShape
//
// # Update a stateshape
//
// Responses:
// default: genericError
//
//	200: stateshapeDBResponse
func (controller *Controller) UpdateStateShape(c *gin.Context) {

	mutexStateShape.Lock()
	defer mutexStateShape.Unlock()

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
	db := backRepo.BackRepoStateShape.GetDB()

	// Validate input
	var input orm.StateShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var stateshapeDB orm.StateShapeDB

	// fetch the stateshape
	_, err := db.First(&stateshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	stateshapeDB.CopyBasicFieldsFromStateShape_WOP(&input.StateShape_WOP)
	stateshapeDB.StateShapePointersEncoding = input.StateShapePointersEncoding

	db, _ = db.Model(&stateshapeDB)
	_, err = db.Updates(&stateshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	stateshapeNew := new(models.StateShape)
	stateshapeDB.CopyBasicFieldsToStateShape(stateshapeNew)

	// redeem pointers
	stateshapeDB.DecodePointers(backRepo, stateshapeNew)

	// get stage instance from DB instance, and call callback function
	stateshapeOld := backRepo.BackRepoStateShape.Map_StateShapeDBID_StateShapePtr[stateshapeDB.ID]
	if stateshapeOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), stateshapeOld, stateshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the stateshapeDB
	c.JSON(http.StatusOK, stateshapeDB)
}

// DeleteStateShape
//
// swagger:route DELETE /stateshapes/{ID} stateshapes deleteStateShape
//
// # Delete a stateshape
//
// default: genericError
//
//	200: stateshapeDBResponse
func (controller *Controller) DeleteStateShape(c *gin.Context) {

	mutexStateShape.Lock()
	defer mutexStateShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStateShape", "Name", stackPath)
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
	db := backRepo.BackRepoStateShape.GetDB()

	// Get model if exist
	var stateshapeDB orm.StateShapeDB
	if _, err := db.First(&stateshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&stateshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	stateshapeDeleted := new(models.StateShape)
	stateshapeDB.CopyBasicFieldsToStateShape(stateshapeDeleted)

	// get stage instance from DB instance, and call callback function
	stateshapeStaged := backRepo.BackRepoStateShape.Map_StateShapeDBID_StateShapePtr[stateshapeDB.ID]
	if stateshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), stateshapeStaged, stateshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
