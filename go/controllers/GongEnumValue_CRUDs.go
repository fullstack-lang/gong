// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongEnumValue__dummysDeclaration__ models.GongEnumValue
var __GongEnumValue_time__dummyDeclaration time.Duration

var mutexGongEnumValue sync.Mutex

// An GongEnumValueID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnumValue updateGongEnumValue deleteGongEnumValue
type GongEnumValueID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumValueInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnumValue updateGongEnumValue
type GongEnumValueInput struct {
	// The GongEnumValue to submit or modify
	// in: body
	GongEnumValue *orm.GongEnumValueAPI
}

// GetGongEnumValues
//
// swagger:route GET /gongenumvalues gongenumvalues getGongEnumValues
//
// # Get all gongenumvalues
//
// Responses:
// default: genericError
//
//	200: gongenumvalueDBResponse
func (controller *Controller) GetGongEnumValues(c *gin.Context) {

	// source slice
	var gongenumvalueDBs []orm.GongEnumValueDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValues", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValue.GetDB()

	_, err := db.Find(&gongenumvalueDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumvalueAPIs := make([]orm.GongEnumValueAPI, 0)

	// for each gongenumvalue, update fields from the database nullable fields
	for idx := range gongenumvalueDBs {
		gongenumvalueDB := &gongenumvalueDBs[idx]
		_ = gongenumvalueDB
		var gongenumvalueAPI orm.GongEnumValueAPI

		// insertion point for updating fields
		gongenumvalueAPI.ID = gongenumvalueDB.ID
		gongenumvalueDB.CopyBasicFieldsToGongEnumValue_WOP(&gongenumvalueAPI.GongEnumValue_WOP)
		gongenumvalueAPI.GongEnumValuePointersEncoding = gongenumvalueDB.GongEnumValuePointersEncoding
		gongenumvalueAPIs = append(gongenumvalueAPIs, gongenumvalueAPI)
	}

	c.JSON(http.StatusOK, gongenumvalueAPIs)
}

// PostGongEnumValue
//
// swagger:route POST /gongenumvalues gongenumvalues postGongEnumValue
//
// Creates a gongenumvalue
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongEnumValue(c *gin.Context) {

	mutexGongEnumValue.Lock()
	defer mutexGongEnumValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongEnumValues", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValue.GetDB()

	// Validate input
	var input orm.GongEnumValueAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenumvalue
	gongenumvalueDB := orm.GongEnumValueDB{}
	gongenumvalueDB.GongEnumValuePointersEncoding = input.GongEnumValuePointersEncoding
	gongenumvalueDB.CopyBasicFieldsFromGongEnumValue_WOP(&input.GongEnumValue_WOP)

	_, err = db.Create(&gongenumvalueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongEnumValue.CheckoutPhaseOneInstance(&gongenumvalueDB)
	gongenumvalue := backRepo.BackRepoGongEnumValue.Map_GongEnumValueDBID_GongEnumValuePtr[gongenumvalueDB.ID]

	if gongenumvalue != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongenumvalue)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumvalueDB)
}

// GetGongEnumValue
//
// swagger:route GET /gongenumvalues/{ID} gongenumvalues getGongEnumValue
//
// Gets the details for a gongenumvalue.
//
// Responses:
// default: genericError
//
//	200: gongenumvalueDBResponse
func (controller *Controller) GetGongEnumValue(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValue.GetDB()

	// Get gongenumvalueDB in DB
	var gongenumvalueDB orm.GongEnumValueDB
	if _, err := db.First(&gongenumvalueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumvalueAPI orm.GongEnumValueAPI
	gongenumvalueAPI.ID = gongenumvalueDB.ID
	gongenumvalueAPI.GongEnumValuePointersEncoding = gongenumvalueDB.GongEnumValuePointersEncoding
	gongenumvalueDB.CopyBasicFieldsToGongEnumValue_WOP(&gongenumvalueAPI.GongEnumValue_WOP)

	c.JSON(http.StatusOK, gongenumvalueAPI)
}

// UpdateGongEnumValue
//
// swagger:route PATCH /gongenumvalues/{ID} gongenumvalues updateGongEnumValue
//
// # Update a gongenumvalue
//
// Responses:
// default: genericError
//
//	200: gongenumvalueDBResponse
func (controller *Controller) UpdateGongEnumValue(c *gin.Context) {

	mutexGongEnumValue.Lock()
	defer mutexGongEnumValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongEnumValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValue.GetDB()

	// Validate input
	var input orm.GongEnumValueAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongenumvalueDB orm.GongEnumValueDB

	// fetch the gongenumvalue
	_, err := db.First(&gongenumvalueDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumvalueDB.CopyBasicFieldsFromGongEnumValue_WOP(&input.GongEnumValue_WOP)
	gongenumvalueDB.GongEnumValuePointersEncoding = input.GongEnumValuePointersEncoding

	db, _ = db.Model(&gongenumvalueDB)
	_, err = db.Updates(&gongenumvalueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueNew := new(models.GongEnumValue)
	gongenumvalueDB.CopyBasicFieldsToGongEnumValue(gongenumvalueNew)

	// redeem pointers
	gongenumvalueDB.DecodePointers(backRepo, gongenumvalueNew)

	// get stage instance from DB instance, and call callback function
	gongenumvalueOld := backRepo.BackRepoGongEnumValue.Map_GongEnumValueDBID_GongEnumValuePtr[gongenumvalueDB.ID]
	if gongenumvalueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongenumvalueOld, gongenumvalueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumvalueDB
	c.JSON(http.StatusOK, gongenumvalueDB)
}

// DeleteGongEnumValue
//
// swagger:route DELETE /gongenumvalues/{ID} gongenumvalues deleteGongEnumValue
//
// # Delete a gongenumvalue
//
// default: genericError
//
//	200: gongenumvalueDBResponse
func (controller *Controller) DeleteGongEnumValue(c *gin.Context) {

	mutexGongEnumValue.Lock()
	defer mutexGongEnumValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongEnumValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumValue.GetDB()

	// Get model if exist
	var gongenumvalueDB orm.GongEnumValueDB
	if _, err := db.First(&gongenumvalueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongenumvalueDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueDeleted := new(models.GongEnumValue)
	gongenumvalueDB.CopyBasicFieldsToGongEnumValue(gongenumvalueDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumvalueStaged := backRepo.BackRepoGongEnumValue.Map_GongEnumValueDBID_GongEnumValuePtr[gongenumvalueDB.ID]
	if gongenumvalueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongenumvalueStaged, gongenumvalueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
