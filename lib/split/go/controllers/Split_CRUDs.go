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
var __Split__dummysDeclaration__ models.Split
var __Split_time__dummyDeclaration time.Duration

var mutexSplit sync.Mutex

// An SplitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSplit updateSplit deleteSplit
type SplitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SplitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSplit updateSplit
type SplitInput struct {
	// The Split to submit or modify
	// in: body
	Split *orm.SplitAPI
}

// GetSplits
//
// swagger:route GET /splits splits getSplits
//
// # Get all splits
//
// Responses:
// default: genericError
//
//	200: splitDBResponse
func (controller *Controller) GetSplits(c *gin.Context) {

	// source slice
	var splitDBs []orm.SplitDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSplits", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSplit.GetDB()

	_, err := db.Find(&splitDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	splitAPIs := make([]orm.SplitAPI, 0)

	// for each split, update fields from the database nullable fields
	for idx := range splitDBs {
		splitDB := &splitDBs[idx]
		_ = splitDB
		var splitAPI orm.SplitAPI

		// insertion point for updating fields
		splitAPI.ID = splitDB.ID
		splitDB.CopyBasicFieldsToSplit_WOP(&splitAPI.Split_WOP)
		splitAPI.SplitPointersEncoding = splitDB.SplitPointersEncoding
		splitAPIs = append(splitAPIs, splitAPI)
	}

	c.JSON(http.StatusOK, splitAPIs)
}

// PostSplit
//
// swagger:route POST /splits splits postSplit
//
// Creates a split
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSplit(c *gin.Context) {

	mutexSplit.Lock()
	defer mutexSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSplits", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSplit.GetDB()

	// Validate input
	var input orm.SplitAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create split
	splitDB := orm.SplitDB{}
	splitDB.SplitPointersEncoding = input.SplitPointersEncoding
	splitDB.CopyBasicFieldsFromSplit_WOP(&input.Split_WOP)

	_, err = db.Create(&splitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSplit.CheckoutPhaseOneInstance(&splitDB)
	split := backRepo.BackRepoSplit.Map_SplitDBID_SplitPtr[splitDB.ID]

	if split != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), split)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, splitDB)
}

// GetSplit
//
// swagger:route GET /splits/{ID} splits getSplit
//
// Gets the details for a split.
//
// Responses:
// default: genericError
//
//	200: splitDBResponse
func (controller *Controller) GetSplit(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSplit.GetDB()

	// Get splitDB in DB
	var splitDB orm.SplitDB
	if _, err := db.First(&splitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var splitAPI orm.SplitAPI
	splitAPI.ID = splitDB.ID
	splitAPI.SplitPointersEncoding = splitDB.SplitPointersEncoding
	splitDB.CopyBasicFieldsToSplit_WOP(&splitAPI.Split_WOP)

	c.JSON(http.StatusOK, splitAPI)
}

// UpdateSplit
//
// swagger:route PATCH /splits/{ID} splits updateSplit
//
// # Update a split
//
// Responses:
// default: genericError
//
//	200: splitDBResponse
func (controller *Controller) UpdateSplit(c *gin.Context) {

	mutexSplit.Lock()
	defer mutexSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSplit.GetDB()

	// Validate input
	var input orm.SplitAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var splitDB orm.SplitDB

	// fetch the split
	_, err := db.First(&splitDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	splitDB.CopyBasicFieldsFromSplit_WOP(&input.Split_WOP)
	splitDB.SplitPointersEncoding = input.SplitPointersEncoding

	db, _ = db.Model(&splitDB)
	_, err = db.Updates(&splitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	splitNew := new(models.Split)
	splitDB.CopyBasicFieldsToSplit(splitNew)

	// redeem pointers
	splitDB.DecodePointers(backRepo, splitNew)

	// get stage instance from DB instance, and call callback function
	splitOld := backRepo.BackRepoSplit.Map_SplitDBID_SplitPtr[splitDB.ID]
	if splitOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), splitOld, splitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the splitDB
	c.JSON(http.StatusOK, splitDB)
}

// DeleteSplit
//
// swagger:route DELETE /splits/{ID} splits deleteSplit
//
// # Delete a split
//
// default: genericError
//
//	200: splitDBResponse
func (controller *Controller) DeleteSplit(c *gin.Context) {

	mutexSplit.Lock()
	defer mutexSplit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSplit", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoSplit.GetDB()

	// Get model if exist
	var splitDB orm.SplitDB
	if _, err := db.First(&splitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&splitDB)

	// get an instance (not staged) from DB instance, and call callback function
	splitDeleted := new(models.Split)
	splitDB.CopyBasicFieldsToSplit(splitDeleted)

	// get stage instance from DB instance, and call callback function
	splitStaged := backRepo.BackRepoSplit.Map_SplitDBID_SplitPtr[splitDB.ID]
	if splitStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), splitStaged, splitDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
