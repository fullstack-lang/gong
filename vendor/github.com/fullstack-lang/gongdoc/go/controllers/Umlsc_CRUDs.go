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
var __Umlsc__dummysDeclaration__ models.Umlsc
var __Umlsc_time__dummyDeclaration time.Duration

var mutexUmlsc sync.Mutex

// An UmlscID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUmlsc updateUmlsc deleteUmlsc
type UmlscID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UmlscInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUmlsc updateUmlsc
type UmlscInput struct {
	// The Umlsc to submit or modify
	// in: body
	Umlsc *orm.UmlscAPI
}

// GetUmlscs
//
// swagger:route GET /umlscs umlscs getUmlscs
//
// # Get all umlscs
//
// Responses:
// default: genericError
//
//	200: umlscDBResponse
func (controller *Controller) GetUmlscs(c *gin.Context) {

	// source slice
	var umlscDBs []orm.UmlscDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUmlscs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlsc.GetDB()

	_, err := db.Find(&umlscDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	umlscAPIs := make([]orm.UmlscAPI, 0)

	// for each umlsc, update fields from the database nullable fields
	for idx := range umlscDBs {
		umlscDB := &umlscDBs[idx]
		_ = umlscDB
		var umlscAPI orm.UmlscAPI

		// insertion point for updating fields
		umlscAPI.ID = umlscDB.ID
		umlscDB.CopyBasicFieldsToUmlsc_WOP(&umlscAPI.Umlsc_WOP)
		umlscAPI.UmlscPointersEncoding = umlscDB.UmlscPointersEncoding
		umlscAPIs = append(umlscAPIs, umlscAPI)
	}

	c.JSON(http.StatusOK, umlscAPIs)
}

// PostUmlsc
//
// swagger:route POST /umlscs umlscs postUmlsc
//
// Creates a umlsc
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUmlsc(c *gin.Context) {

	mutexUmlsc.Lock()
	defer mutexUmlsc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUmlscs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlsc.GetDB()

	// Validate input
	var input orm.UmlscAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create umlsc
	umlscDB := orm.UmlscDB{}
	umlscDB.UmlscPointersEncoding = input.UmlscPointersEncoding
	umlscDB.CopyBasicFieldsFromUmlsc_WOP(&input.Umlsc_WOP)

	_, err = db.Create(&umlscDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUmlsc.CheckoutPhaseOneInstance(&umlscDB)
	umlsc := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[umlscDB.ID]

	if umlsc != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), umlsc)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, umlscDB)
}

// GetUmlsc
//
// swagger:route GET /umlscs/{ID} umlscs getUmlsc
//
// Gets the details for a umlsc.
//
// Responses:
// default: genericError
//
//	200: umlscDBResponse
func (controller *Controller) GetUmlsc(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUmlsc", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlsc.GetDB()

	// Get umlscDB in DB
	var umlscDB orm.UmlscDB
	if _, err := db.First(&umlscDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var umlscAPI orm.UmlscAPI
	umlscAPI.ID = umlscDB.ID
	umlscAPI.UmlscPointersEncoding = umlscDB.UmlscPointersEncoding
	umlscDB.CopyBasicFieldsToUmlsc_WOP(&umlscAPI.Umlsc_WOP)

	c.JSON(http.StatusOK, umlscAPI)
}

// UpdateUmlsc
//
// swagger:route PATCH /umlscs/{ID} umlscs updateUmlsc
//
// # Update a umlsc
//
// Responses:
// default: genericError
//
//	200: umlscDBResponse
func (controller *Controller) UpdateUmlsc(c *gin.Context) {

	mutexUmlsc.Lock()
	defer mutexUmlsc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUmlsc", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlsc.GetDB()

	// Validate input
	var input orm.UmlscAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var umlscDB orm.UmlscDB

	// fetch the umlsc
	_, err := db.First(&umlscDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	umlscDB.CopyBasicFieldsFromUmlsc_WOP(&input.Umlsc_WOP)
	umlscDB.UmlscPointersEncoding = input.UmlscPointersEncoding

	db, _ = db.Model(&umlscDB)
	_, err = db.Updates(&umlscDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	umlscNew := new(models.Umlsc)
	umlscDB.CopyBasicFieldsToUmlsc(umlscNew)

	// redeem pointers
	umlscDB.DecodePointers(backRepo, umlscNew)

	// get stage instance from DB instance, and call callback function
	umlscOld := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[umlscDB.ID]
	if umlscOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), umlscOld, umlscNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the umlscDB
	c.JSON(http.StatusOK, umlscDB)
}

// DeleteUmlsc
//
// swagger:route DELETE /umlscs/{ID} umlscs deleteUmlsc
//
// # Delete a umlsc
//
// default: genericError
//
//	200: umlscDBResponse
func (controller *Controller) DeleteUmlsc(c *gin.Context) {

	mutexUmlsc.Lock()
	defer mutexUmlsc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUmlsc", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUmlsc.GetDB()

	// Get model if exist
	var umlscDB orm.UmlscDB
	if _, err := db.First(&umlscDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&umlscDB)

	// get an instance (not staged) from DB instance, and call callback function
	umlscDeleted := new(models.Umlsc)
	umlscDB.CopyBasicFieldsToUmlsc(umlscDeleted)

	// get stage instance from DB instance, and call callback function
	umlscStaged := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[umlscDB.ID]
	if umlscStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), umlscStaged, umlscDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
