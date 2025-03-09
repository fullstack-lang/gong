// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	"github.com/fullstack-lang/gong/test/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Gstruct__dummysDeclaration__ models.Gstruct
var __Gstruct_time__dummyDeclaration time.Duration

var mutexGstruct sync.Mutex

// An GstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGstruct updateGstruct deleteGstruct
type GstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGstruct updateGstruct
type GstructInput struct {
	// The Gstruct to submit or modify
	// in: body
	Gstruct *orm.GstructAPI
}

// GetGstructs
//
// swagger:route GET /gstructs gstructs getGstructs
//
// # Get all gstructs
//
// Responses:
// default: genericError
//
//	200: gstructDBResponse
func (controller *Controller) GetGstructs(c *gin.Context) {

	// source slice
	var gstructDBs []orm.GstructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	_, err := db.Find(&gstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gstructAPIs := make([]orm.GstructAPI, 0)

	// for each gstruct, update fields from the database nullable fields
	for idx := range gstructDBs {
		gstructDB := &gstructDBs[idx]
		_ = gstructDB
		var gstructAPI orm.GstructAPI

		// insertion point for updating fields
		gstructAPI.ID = gstructDB.ID
		gstructDB.CopyBasicFieldsToGstruct_WOP(&gstructAPI.Gstruct_WOP)
		gstructAPI.GstructPointersEncoding = gstructDB.GstructPointersEncoding
		gstructAPIs = append(gstructAPIs, gstructAPI)
	}

	c.JSON(http.StatusOK, gstructAPIs)
}

// PostGstruct
//
// swagger:route POST /gstructs gstructs postGstruct
//
// Creates a gstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGstruct(c *gin.Context) {

	mutexGstruct.Lock()
	defer mutexGstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	// Validate input
	var input orm.GstructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gstruct
	gstructDB := orm.GstructDB{}
	gstructDB.GstructPointersEncoding = input.GstructPointersEncoding
	gstructDB.CopyBasicFieldsFromGstruct_WOP(&input.Gstruct_WOP)

	_, err = db.Create(&gstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGstruct.CheckoutPhaseOneInstance(&gstructDB)
	gstruct := backRepo.BackRepoGstruct.Map_GstructDBID_GstructPtr[gstructDB.ID]

	if gstruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gstructDB)
}

// GetGstruct
//
// swagger:route GET /gstructs/{ID} gstructs getGstruct
//
// Gets the details for a gstruct.
//
// Responses:
// default: genericError
//
//	200: gstructDBResponse
func (controller *Controller) GetGstruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	// Get gstructDB in DB
	var gstructDB orm.GstructDB
	if _, err := db.First(&gstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gstructAPI orm.GstructAPI
	gstructAPI.ID = gstructDB.ID
	gstructAPI.GstructPointersEncoding = gstructDB.GstructPointersEncoding
	gstructDB.CopyBasicFieldsToGstruct_WOP(&gstructAPI.Gstruct_WOP)

	c.JSON(http.StatusOK, gstructAPI)
}

// UpdateGstruct
//
// swagger:route PATCH /gstructs/{ID} gstructs updateGstruct
//
// # Update a gstruct
//
// Responses:
// default: genericError
//
//	200: gstructDBResponse
func (controller *Controller) UpdateGstruct(c *gin.Context) {

	mutexGstruct.Lock()
	defer mutexGstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	// Validate input
	var input orm.GstructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gstructDB orm.GstructDB

	// fetch the gstruct
	_, err := db.First(&gstructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gstructDB.CopyBasicFieldsFromGstruct_WOP(&input.Gstruct_WOP)
	gstructDB.GstructPointersEncoding = input.GstructPointersEncoding

	db, _ = db.Model(&gstructDB)
	_, err = db.Updates(&gstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gstructNew := new(models.Gstruct)
	gstructDB.CopyBasicFieldsToGstruct(gstructNew)

	// redeem pointers
	gstructDB.DecodePointers(backRepo, gstructNew)

	// get stage instance from DB instance, and call callback function
	gstructOld := backRepo.BackRepoGstruct.Map_GstructDBID_GstructPtr[gstructDB.ID]
	if gstructOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gstructOld, gstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gstructDB
	c.JSON(http.StatusOK, gstructDB)
}

// DeleteGstruct
//
// swagger:route DELETE /gstructs/{ID} gstructs deleteGstruct
//
// # Delete a gstruct
//
// default: genericError
//
//	200: gstructDBResponse
func (controller *Controller) DeleteGstruct(c *gin.Context) {

	mutexGstruct.Lock()
	defer mutexGstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGstruct.GetDB()

	// Get model if exist
	var gstructDB orm.GstructDB
	if _, err := db.First(&gstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	gstructDeleted := new(models.Gstruct)
	gstructDB.CopyBasicFieldsToGstruct(gstructDeleted)

	// get stage instance from DB instance, and call callback function
	gstructStaged := backRepo.BackRepoGstruct.Map_GstructDBID_GstructPtr[gstructDB.ID]
	if gstructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gstructStaged, gstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
