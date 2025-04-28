// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Load__dummysDeclaration__ models.Load
var __Load_time__dummyDeclaration time.Duration

var mutexLoad sync.Mutex

// An LoadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLoad updateLoad deleteLoad
type LoadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LoadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLoad updateLoad
type LoadInput struct {
	// The Load to submit or modify
	// in: body
	Load *orm.LoadAPI
}

// GetLoads
//
// swagger:route GET /loads loads getLoads
//
// # Get all loads
//
// Responses:
// default: genericError
//
//	200: loadDBResponse
func (controller *Controller) GetLoads(c *gin.Context) {

	// source slice
	var loadDBs []orm.LoadDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLoads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLoad.GetDB()

	_, err := db.Find(&loadDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	loadAPIs := make([]orm.LoadAPI, 0)

	// for each load, update fields from the database nullable fields
	for idx := range loadDBs {
		loadDB := &loadDBs[idx]
		_ = loadDB
		var loadAPI orm.LoadAPI

		// insertion point for updating fields
		loadAPI.ID = loadDB.ID
		loadDB.CopyBasicFieldsToLoad_WOP(&loadAPI.Load_WOP)
		loadAPI.LoadPointersEncoding = loadDB.LoadPointersEncoding
		loadAPIs = append(loadAPIs, loadAPI)
	}

	c.JSON(http.StatusOK, loadAPIs)
}

// PostLoad
//
// swagger:route POST /loads loads postLoad
//
// Creates a load
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLoad(c *gin.Context) {

	mutexLoad.Lock()
	defer mutexLoad.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLoads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLoad.GetDB()

	// Validate input
	var input orm.LoadAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create load
	loadDB := orm.LoadDB{}
	loadDB.LoadPointersEncoding = input.LoadPointersEncoding
	loadDB.CopyBasicFieldsFromLoad_WOP(&input.Load_WOP)

	_, err = db.Create(&loadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLoad.CheckoutPhaseOneInstance(&loadDB)
	load := backRepo.BackRepoLoad.Map_LoadDBID_LoadPtr[loadDB.ID]

	if load != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), load)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, loadDB)
}

// GetLoad
//
// swagger:route GET /loads/{ID} loads getLoad
//
// Gets the details for a load.
//
// Responses:
// default: genericError
//
//	200: loadDBResponse
func (controller *Controller) GetLoad(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLoad", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLoad.GetDB()

	// Get loadDB in DB
	var loadDB orm.LoadDB
	if _, err := db.First(&loadDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var loadAPI orm.LoadAPI
	loadAPI.ID = loadDB.ID
	loadAPI.LoadPointersEncoding = loadDB.LoadPointersEncoding
	loadDB.CopyBasicFieldsToLoad_WOP(&loadAPI.Load_WOP)

	c.JSON(http.StatusOK, loadAPI)
}

// UpdateLoad
//
// swagger:route PATCH /loads/{ID} loads updateLoad
//
// # Update a load
//
// Responses:
// default: genericError
//
//	200: loadDBResponse
func (controller *Controller) UpdateLoad(c *gin.Context) {

	mutexLoad.Lock()
	defer mutexLoad.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLoad", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLoad.GetDB()

	// Validate input
	var input orm.LoadAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var loadDB orm.LoadDB

	// fetch the load
	_, err := db.First(&loadDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	loadDB.CopyBasicFieldsFromLoad_WOP(&input.Load_WOP)
	loadDB.LoadPointersEncoding = input.LoadPointersEncoding

	db, _ = db.Model(&loadDB)
	_, err = db.Updates(&loadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	loadNew := new(models.Load)
	loadDB.CopyBasicFieldsToLoad(loadNew)

	// redeem pointers
	loadDB.DecodePointers(backRepo, loadNew)

	// get stage instance from DB instance, and call callback function
	loadOld := backRepo.BackRepoLoad.Map_LoadDBID_LoadPtr[loadDB.ID]
	if loadOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), loadOld, loadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the loadDB
	c.JSON(http.StatusOK, loadDB)
}

// DeleteLoad
//
// swagger:route DELETE /loads/{ID} loads deleteLoad
//
// # Delete a load
//
// default: genericError
//
//	200: loadDBResponse
func (controller *Controller) DeleteLoad(c *gin.Context) {

	mutexLoad.Lock()
	defer mutexLoad.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLoad", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLoad.GetDB()

	// Get model if exist
	var loadDB orm.LoadDB
	if _, err := db.First(&loadDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&loadDB)

	// get an instance (not staged) from DB instance, and call callback function
	loadDeleted := new(models.Load)
	loadDB.CopyBasicFieldsToLoad(loadDeleted)

	// get stage instance from DB instance, and call callback function
	loadStaged := backRepo.BackRepoLoad.Map_LoadDBID_LoadPtr[loadDB.ID]
	if loadStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), loadStaged, loadDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
