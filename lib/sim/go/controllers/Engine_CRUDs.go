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
var __Engine__dummysDeclaration__ models.Engine
var __Engine_time__dummyDeclaration time.Duration

var mutexEngine sync.Mutex

// An EngineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEngine updateEngine deleteEngine
type EngineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EngineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEngine updateEngine
type EngineInput struct {
	// The Engine to submit or modify
	// in: body
	Engine *orm.EngineAPI
}

// GetEngines
//
// swagger:route GET /engines engines getEngines
//
// # Get all engines
//
// Responses:
// default: genericError
//
//	200: engineDBResponse
func (controller *Controller) GetEngines(c *gin.Context) {

	// source slice
	var engineDBs []orm.EngineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEngines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEngine.GetDB()

	_, err := db.Find(&engineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	engineAPIs := make([]orm.EngineAPI, 0)

	// for each engine, update fields from the database nullable fields
	for idx := range engineDBs {
		engineDB := &engineDBs[idx]
		_ = engineDB
		var engineAPI orm.EngineAPI

		// insertion point for updating fields
		engineAPI.ID = engineDB.ID
		engineDB.CopyBasicFieldsToEngine_WOP(&engineAPI.Engine_WOP)
		engineAPI.EnginePointersEncoding = engineDB.EnginePointersEncoding
		engineAPIs = append(engineAPIs, engineAPI)
	}

	c.JSON(http.StatusOK, engineAPIs)
}

// PostEngine
//
// swagger:route POST /engines engines postEngine
//
// Creates a engine
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEngine(c *gin.Context) {

	mutexEngine.Lock()
	defer mutexEngine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEngines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEngine.GetDB()

	// Validate input
	var input orm.EngineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create engine
	engineDB := orm.EngineDB{}
	engineDB.EnginePointersEncoding = input.EnginePointersEncoding
	engineDB.CopyBasicFieldsFromEngine_WOP(&input.Engine_WOP)

	_, err = db.Create(&engineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEngine.CheckoutPhaseOneInstance(&engineDB)
	engine := backRepo.BackRepoEngine.Map_EngineDBID_EnginePtr[engineDB.ID]

	if engine != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), engine)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, engineDB)
}

// GetEngine
//
// swagger:route GET /engines/{ID} engines getEngine
//
// Gets the details for a engine.
//
// Responses:
// default: genericError
//
//	200: engineDBResponse
func (controller *Controller) GetEngine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEngine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEngine.GetDB()

	// Get engineDB in DB
	var engineDB orm.EngineDB
	if _, err := db.First(&engineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var engineAPI orm.EngineAPI
	engineAPI.ID = engineDB.ID
	engineAPI.EnginePointersEncoding = engineDB.EnginePointersEncoding
	engineDB.CopyBasicFieldsToEngine_WOP(&engineAPI.Engine_WOP)

	c.JSON(http.StatusOK, engineAPI)
}

// UpdateEngine
//
// swagger:route PATCH /engines/{ID} engines updateEngine
//
// # Update a engine
//
// Responses:
// default: genericError
//
//	200: engineDBResponse
func (controller *Controller) UpdateEngine(c *gin.Context) {

	mutexEngine.Lock()
	defer mutexEngine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEngine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEngine.GetDB()

	// Validate input
	var input orm.EngineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var engineDB orm.EngineDB

	// fetch the engine
	_, err := db.First(&engineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	engineDB.CopyBasicFieldsFromEngine_WOP(&input.Engine_WOP)
	engineDB.EnginePointersEncoding = input.EnginePointersEncoding

	db, _ = db.Model(&engineDB)
	_, err = db.Updates(&engineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	engineNew := new(models.Engine)
	engineDB.CopyBasicFieldsToEngine(engineNew)

	// redeem pointers
	engineDB.DecodePointers(backRepo, engineNew)

	// get stage instance from DB instance, and call callback function
	engineOld := backRepo.BackRepoEngine.Map_EngineDBID_EnginePtr[engineDB.ID]
	if engineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), engineOld, engineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the engineDB
	c.JSON(http.StatusOK, engineDB)
}

// DeleteEngine
//
// swagger:route DELETE /engines/{ID} engines deleteEngine
//
// # Delete a engine
//
// default: genericError
//
//	200: engineDBResponse
func (controller *Controller) DeleteEngine(c *gin.Context) {

	mutexEngine.Lock()
	defer mutexEngine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEngine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEngine.GetDB()

	// Get model if exist
	var engineDB orm.EngineDB
	if _, err := db.First(&engineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&engineDB)

	// get an instance (not staged) from DB instance, and call callback function
	engineDeleted := new(models.Engine)
	engineDB.CopyBasicFieldsToEngine(engineDeleted)

	// get stage instance from DB instance, and call callback function
	engineStaged := backRepo.BackRepoEngine.Map_EngineDBID_EnginePtr[engineDB.ID]
	if engineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), engineStaged, engineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
