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
var __SplitArea__dummysDeclaration__ models.SplitArea
var __SplitArea_time__dummyDeclaration time.Duration

var mutexSplitArea sync.Mutex

// An SplitAreaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSplitArea updateSplitArea deleteSplitArea
type SplitAreaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SplitAreaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSplitArea updateSplitArea
type SplitAreaInput struct {
	// The SplitArea to submit or modify
	// in: body
	SplitArea *orm.SplitAreaAPI
}

// GetSplitAreas
//
// swagger:route GET /splitareas splitareas getSplitAreas
//
// # Get all splitareas
//
// Responses:
// default: genericError
//
//	200: splitareaDBResponse
func (controller *Controller) GetSplitAreas(c *gin.Context) {

	// source slice
	var splitareaDBs []orm.SplitAreaDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSplitAreas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSplitArea.GetDB()

	_, err := db.Find(&splitareaDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	splitareaAPIs := make([]orm.SplitAreaAPI, 0)

	// for each splitarea, update fields from the database nullable fields
	for idx := range splitareaDBs {
		splitareaDB := &splitareaDBs[idx]
		_ = splitareaDB
		var splitareaAPI orm.SplitAreaAPI

		// insertion point for updating fields
		splitareaAPI.ID = splitareaDB.ID
		splitareaDB.CopyBasicFieldsToSplitArea_WOP(&splitareaAPI.SplitArea_WOP)
		splitareaAPI.SplitAreaPointersEncoding = splitareaDB.SplitAreaPointersEncoding
		splitareaAPIs = append(splitareaAPIs, splitareaAPI)
	}

	c.JSON(http.StatusOK, splitareaAPIs)
}

// PostSplitArea
//
// swagger:route POST /splitareas splitareas postSplitArea
//
// Creates a splitarea
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSplitArea(c *gin.Context) {

	mutexSplitArea.Lock()
	defer mutexSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSplitAreas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSplitArea.GetDB()

	// Validate input
	var input orm.SplitAreaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create splitarea
	splitareaDB := orm.SplitAreaDB{}
	splitareaDB.SplitAreaPointersEncoding = input.SplitAreaPointersEncoding
	splitareaDB.CopyBasicFieldsFromSplitArea_WOP(&input.SplitArea_WOP)

	_, err = db.Create(&splitareaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSplitArea.CheckoutPhaseOneInstance(&splitareaDB)
	splitarea := backRepo.BackRepoSplitArea.Map_SplitAreaDBID_SplitAreaPtr[splitareaDB.ID]

	if splitarea != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), splitarea)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, splitareaDB)
}

// GetSplitArea
//
// swagger:route GET /splitareas/{ID} splitareas getSplitArea
//
// Gets the details for a splitarea.
//
// Responses:
// default: genericError
//
//	200: splitareaDBResponse
func (controller *Controller) GetSplitArea(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSplitArea", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSplitArea.GetDB()

	// Get splitareaDB in DB
	var splitareaDB orm.SplitAreaDB
	if _, err := db.First(&splitareaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var splitareaAPI orm.SplitAreaAPI
	splitareaAPI.ID = splitareaDB.ID
	splitareaAPI.SplitAreaPointersEncoding = splitareaDB.SplitAreaPointersEncoding
	splitareaDB.CopyBasicFieldsToSplitArea_WOP(&splitareaAPI.SplitArea_WOP)

	c.JSON(http.StatusOK, splitareaAPI)
}

// UpdateSplitArea
//
// swagger:route PATCH /splitareas/{ID} splitareas updateSplitArea
//
// # Update a splitarea
//
// Responses:
// default: genericError
//
//	200: splitareaDBResponse
func (controller *Controller) UpdateSplitArea(c *gin.Context) {

	mutexSplitArea.Lock()
	defer mutexSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSplitArea", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSplitArea.GetDB()

	// Validate input
	var input orm.SplitAreaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var splitareaDB orm.SplitAreaDB

	// fetch the splitarea
	_, err := db.First(&splitareaDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	splitareaDB.CopyBasicFieldsFromSplitArea_WOP(&input.SplitArea_WOP)
	splitareaDB.SplitAreaPointersEncoding = input.SplitAreaPointersEncoding

	db, _ = db.Model(&splitareaDB)
	_, err = db.Updates(&splitareaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	splitareaNew := new(models.SplitArea)
	splitareaDB.CopyBasicFieldsToSplitArea(splitareaNew)

	// redeem pointers
	splitareaDB.DecodePointers(backRepo, splitareaNew)

	// get stage instance from DB instance, and call callback function
	splitareaOld := backRepo.BackRepoSplitArea.Map_SplitAreaDBID_SplitAreaPtr[splitareaDB.ID]
	if splitareaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), splitareaOld, splitareaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the splitareaDB
	c.JSON(http.StatusOK, splitareaDB)
}

// DeleteSplitArea
//
// swagger:route DELETE /splitareas/{ID} splitareas deleteSplitArea
//
// # Delete a splitarea
//
// default: genericError
//
//	200: splitareaDBResponse
func (controller *Controller) DeleteSplitArea(c *gin.Context) {

	mutexSplitArea.Lock()
	defer mutexSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSplitArea", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSplitArea.GetDB()

	// Get model if exist
	var splitareaDB orm.SplitAreaDB
	if _, err := db.First(&splitareaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&splitareaDB)

	// get an instance (not staged) from DB instance, and call callback function
	splitareaDeleted := new(models.SplitArea)
	splitareaDB.CopyBasicFieldsToSplitArea(splitareaDeleted)

	// get stage instance from DB instance, and call callback function
	splitareaStaged := backRepo.BackRepoSplitArea.Map_SplitAreaDBID_SplitAreaPtr[splitareaDB.ID]
	if splitareaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), splitareaStaged, splitareaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
