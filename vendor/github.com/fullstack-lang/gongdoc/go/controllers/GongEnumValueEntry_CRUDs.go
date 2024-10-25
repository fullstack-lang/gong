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
var __GongEnumValueEntry__dummysDeclaration__ models.GongEnumValueEntry
var __GongEnumValueEntry_time__dummyDeclaration time.Duration

var mutexGongEnumValueEntry sync.Mutex

// An GongEnumValueEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnumValueEntry updateGongEnumValueEntry deleteGongEnumValueEntry
type GongEnumValueEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumValueEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnumValueEntry updateGongEnumValueEntry
type GongEnumValueEntryInput struct {
	// The GongEnumValueEntry to submit or modify
	// in: body
	GongEnumValueEntry *orm.GongEnumValueEntryAPI
}

// GetGongEnumValueEntrys
//
// swagger:route GET /gongenumvalueentrys gongenumvalueentrys getGongEnumValueEntrys
//
// # Get all gongenumvalueentrys
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func (controller *Controller) GetGongEnumValueEntrys(c *gin.Context) {

	// source slice
	var gongenumvalueentryDBs []orm.GongEnumValueEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValueEntrys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValueEntry.GetDB()

	_, err := db.Find(&gongenumvalueentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumvalueentryAPIs := make([]orm.GongEnumValueEntryAPI, 0)

	// for each gongenumvalueentry, update fields from the database nullable fields
	for idx := range gongenumvalueentryDBs {
		gongenumvalueentryDB := &gongenumvalueentryDBs[idx]
		_ = gongenumvalueentryDB
		var gongenumvalueentryAPI orm.GongEnumValueEntryAPI

		// insertion point for updating fields
		gongenumvalueentryAPI.ID = gongenumvalueentryDB.ID
		gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry_WOP(&gongenumvalueentryAPI.GongEnumValueEntry_WOP)
		gongenumvalueentryAPI.GongEnumValueEntryPointersEncoding = gongenumvalueentryDB.GongEnumValueEntryPointersEncoding
		gongenumvalueentryAPIs = append(gongenumvalueentryAPIs, gongenumvalueentryAPI)
	}

	c.JSON(http.StatusOK, gongenumvalueentryAPIs)
}

// PostGongEnumValueEntry
//
// swagger:route POST /gongenumvalueentrys gongenumvalueentrys postGongEnumValueEntry
//
// Creates a gongenumvalueentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongEnumValueEntry(c *gin.Context) {

	mutexGongEnumValueEntry.Lock()
	defer mutexGongEnumValueEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongEnumValueEntrys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValueEntry.GetDB()

	// Validate input
	var input orm.GongEnumValueEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenumvalueentry
	gongenumvalueentryDB := orm.GongEnumValueEntryDB{}
	gongenumvalueentryDB.GongEnumValueEntryPointersEncoding = input.GongEnumValueEntryPointersEncoding
	gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry_WOP(&input.GongEnumValueEntry_WOP)

	_, err = db.Create(&gongenumvalueentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongEnumValueEntry.CheckoutPhaseOneInstance(&gongenumvalueentryDB)
	gongenumvalueentry := backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr[gongenumvalueentryDB.ID]

	if gongenumvalueentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongenumvalueentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumvalueentryDB)
}

// GetGongEnumValueEntry
//
// swagger:route GET /gongenumvalueentrys/{ID} gongenumvalueentrys getGongEnumValueEntry
//
// Gets the details for a gongenumvalueentry.
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func (controller *Controller) GetGongEnumValueEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValueEntry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValueEntry.GetDB()

	// Get gongenumvalueentryDB in DB
	var gongenumvalueentryDB orm.GongEnumValueEntryDB
	if _, err := db.First(&gongenumvalueentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumvalueentryAPI orm.GongEnumValueEntryAPI
	gongenumvalueentryAPI.ID = gongenumvalueentryDB.ID
	gongenumvalueentryAPI.GongEnumValueEntryPointersEncoding = gongenumvalueentryDB.GongEnumValueEntryPointersEncoding
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry_WOP(&gongenumvalueentryAPI.GongEnumValueEntry_WOP)

	c.JSON(http.StatusOK, gongenumvalueentryAPI)
}

// UpdateGongEnumValueEntry
//
// swagger:route PATCH /gongenumvalueentrys/{ID} gongenumvalueentrys updateGongEnumValueEntry
//
// # Update a gongenumvalueentry
//
// Responses:
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func (controller *Controller) UpdateGongEnumValueEntry(c *gin.Context) {

	mutexGongEnumValueEntry.Lock()
	defer mutexGongEnumValueEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongEnumValueEntry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValueEntry.GetDB()

	// Validate input
	var input orm.GongEnumValueEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongenumvalueentryDB orm.GongEnumValueEntryDB

	// fetch the gongenumvalueentry
	_, err := db.First(&gongenumvalueentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumvalueentryDB.CopyBasicFieldsFromGongEnumValueEntry_WOP(&input.GongEnumValueEntry_WOP)
	gongenumvalueentryDB.GongEnumValueEntryPointersEncoding = input.GongEnumValueEntryPointersEncoding

	db, _ = db.Model(&gongenumvalueentryDB)
	_, err = db.Updates(&gongenumvalueentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueentryNew := new(models.GongEnumValueEntry)
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentryNew)

	// redeem pointers
	gongenumvalueentryDB.DecodePointers(backRepo, gongenumvalueentryNew)

	// get stage instance from DB instance, and call callback function
	gongenumvalueentryOld := backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr[gongenumvalueentryDB.ID]
	if gongenumvalueentryOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongenumvalueentryOld, gongenumvalueentryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumvalueentryDB
	c.JSON(http.StatusOK, gongenumvalueentryDB)
}

// DeleteGongEnumValueEntry
//
// swagger:route DELETE /gongenumvalueentrys/{ID} gongenumvalueentrys deleteGongEnumValueEntry
//
// # Delete a gongenumvalueentry
//
// default: genericError
//
//	200: gongenumvalueentryDBResponse
func (controller *Controller) DeleteGongEnumValueEntry(c *gin.Context) {

	mutexGongEnumValueEntry.Lock()
	defer mutexGongEnumValueEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongEnumValueEntry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValueEntry.GetDB()

	// Get model if exist
	var gongenumvalueentryDB orm.GongEnumValueEntryDB
	if _, err := db.First(&gongenumvalueentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongenumvalueentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueentryDeleted := new(models.GongEnumValueEntry)
	gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry(gongenumvalueentryDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumvalueentryStaged := backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryPtr[gongenumvalueentryDB.ID]
	if gongenumvalueentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongenumvalueentryStaged, gongenumvalueentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
