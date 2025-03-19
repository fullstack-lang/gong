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
var __AsSplit__dummysDeclaration__ models.AsSplit
var __AsSplit_time__dummyDeclaration time.Duration

var mutexAsSplit sync.Mutex

// An AsSplitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAsSplit updateAsSplit deleteAsSplit
type AsSplitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AsSplitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAsSplit updateAsSplit
type AsSplitInput struct {
	// The AsSplit to submit or modify
	// in: body
	AsSplit *orm.AsSplitAPI
}

// GetAsSplits
//
// swagger:route GET /assplits assplits getAsSplits
//
// # Get all assplits
//
// Responses:
// default: genericError
//
//	200: assplitDBResponse
func (controller *Controller) GetAsSplits(c *gin.Context) {

	// source slice
	var assplitDBs []orm.AsSplitDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAsSplits", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	_, err := db.Find(&assplitDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	assplitAPIs := make([]orm.AsSplitAPI, 0)

	// for each assplit, update fields from the database nullable fields
	for idx := range assplitDBs {
		assplitDB := &assplitDBs[idx]
		_ = assplitDB
		var assplitAPI orm.AsSplitAPI

		// insertion point for updating fields
		assplitAPI.ID = assplitDB.ID
		assplitDB.CopyBasicFieldsToAsSplit_WOP(&assplitAPI.AsSplit_WOP)
		assplitAPI.AsSplitPointersEncoding = assplitDB.AsSplitPointersEncoding
		assplitAPIs = append(assplitAPIs, assplitAPI)
	}

	c.JSON(http.StatusOK, assplitAPIs)
}

// PostAsSplit
//
// swagger:route POST /assplits assplits postAsSplit
//
// Creates a assplit
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAsSplit(c *gin.Context) {

	mutexAsSplit.Lock()
	defer mutexAsSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAsSplits", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	// Validate input
	var input orm.AsSplitAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create assplit
	assplitDB := orm.AsSplitDB{}
	assplitDB.AsSplitPointersEncoding = input.AsSplitPointersEncoding
	assplitDB.CopyBasicFieldsFromAsSplit_WOP(&input.AsSplit_WOP)

	_, err = db.Create(&assplitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAsSplit.CheckoutPhaseOneInstance(&assplitDB)
	assplit := backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitPtr[assplitDB.ID]

	if assplit != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), assplit)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, assplitDB)
}

// GetAsSplit
//
// swagger:route GET /assplits/{ID} assplits getAsSplit
//
// Gets the details for a assplit.
//
// Responses:
// default: genericError
//
//	200: assplitDBResponse
func (controller *Controller) GetAsSplit(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAsSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	// Get assplitDB in DB
	var assplitDB orm.AsSplitDB
	if _, err := db.First(&assplitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var assplitAPI orm.AsSplitAPI
	assplitAPI.ID = assplitDB.ID
	assplitAPI.AsSplitPointersEncoding = assplitDB.AsSplitPointersEncoding
	assplitDB.CopyBasicFieldsToAsSplit_WOP(&assplitAPI.AsSplit_WOP)

	c.JSON(http.StatusOK, assplitAPI)
}

// UpdateAsSplit
//
// swagger:route PATCH /assplits/{ID} assplits updateAsSplit
//
// # Update a assplit
//
// Responses:
// default: genericError
//
//	200: assplitDBResponse
func (controller *Controller) UpdateAsSplit(c *gin.Context) {

	mutexAsSplit.Lock()
	defer mutexAsSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAsSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	// Validate input
	var input orm.AsSplitAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var assplitDB orm.AsSplitDB

	// fetch the assplit
	_, err := db.First(&assplitDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	assplitDB.CopyBasicFieldsFromAsSplit_WOP(&input.AsSplit_WOP)
	assplitDB.AsSplitPointersEncoding = input.AsSplitPointersEncoding

	db, _ = db.Model(&assplitDB)
	_, err = db.Updates(&assplitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	assplitNew := new(models.AsSplit)
	assplitDB.CopyBasicFieldsToAsSplit(assplitNew)

	// redeem pointers
	assplitDB.DecodePointers(backRepo, assplitNew)

	// get stage instance from DB instance, and call callback function
	assplitOld := backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitPtr[assplitDB.ID]
	if assplitOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), assplitOld, assplitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the assplitDB
	c.JSON(http.StatusOK, assplitDB)
}

// DeleteAsSplit
//
// swagger:route DELETE /assplits/{ID} assplits deleteAsSplit
//
// # Delete a assplit
//
// default: genericError
//
//	200: assplitDBResponse
func (controller *Controller) DeleteAsSplit(c *gin.Context) {

	mutexAsSplit.Lock()
	defer mutexAsSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAsSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoAsSplit.GetDB()

	// Get model if exist
	var assplitDB orm.AsSplitDB
	if _, err := db.First(&assplitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&assplitDB)

	// get an instance (not staged) from DB instance, and call callback function
	assplitDeleted := new(models.AsSplit)
	assplitDB.CopyBasicFieldsToAsSplit(assplitDeleted)

	// get stage instance from DB instance, and call callback function
	assplitStaged := backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitPtr[assplitDB.ID]
	if assplitStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), assplitStaged, assplitDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
