// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Polygone__dummysDeclaration__ models.Polygone
var __Polygone_time__dummyDeclaration time.Duration

var mutexPolygone sync.Mutex

// An PolygoneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPolygone updatePolygone deletePolygone
type PolygoneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolygoneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPolygone updatePolygone
type PolygoneInput struct {
	// The Polygone to submit or modify
	// in: body
	Polygone *orm.PolygoneAPI
}

// GetPolygones
//
// swagger:route GET /polygones polygones getPolygones
//
// # Get all polygones
//
// Responses:
// default: genericError
//
//	200: polygoneDBResponse
func (controller *Controller) GetPolygones(c *gin.Context) {

	// source slice
	var polygoneDBs []orm.PolygoneDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPolygones", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	_, err := db.Find(&polygoneDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	polygoneAPIs := make([]orm.PolygoneAPI, 0)

	// for each polygone, update fields from the database nullable fields
	for idx := range polygoneDBs {
		polygoneDB := &polygoneDBs[idx]
		_ = polygoneDB
		var polygoneAPI orm.PolygoneAPI

		// insertion point for updating fields
		polygoneAPI.ID = polygoneDB.ID
		polygoneDB.CopyBasicFieldsToPolygone_WOP(&polygoneAPI.Polygone_WOP)
		polygoneAPI.PolygonePointersEncoding = polygoneDB.PolygonePointersEncoding
		polygoneAPIs = append(polygoneAPIs, polygoneAPI)
	}

	c.JSON(http.StatusOK, polygoneAPIs)
}

// PostPolygone
//
// swagger:route POST /polygones polygones postPolygone
//
// Creates a polygone
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPolygone(c *gin.Context) {

	mutexPolygone.Lock()
	defer mutexPolygone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPolygones", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	// Validate input
	var input orm.PolygoneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create polygone
	polygoneDB := orm.PolygoneDB{}
	polygoneDB.PolygonePointersEncoding = input.PolygonePointersEncoding
	polygoneDB.CopyBasicFieldsFromPolygone_WOP(&input.Polygone_WOP)

	_, err = db.Create(&polygoneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPolygone.CheckoutPhaseOneInstance(&polygoneDB)
	polygone := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]

	if polygone != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), polygone)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, polygoneDB)
}

// GetPolygone
//
// swagger:route GET /polygones/{ID} polygones getPolygone
//
// Gets the details for a polygone.
//
// Responses:
// default: genericError
//
//	200: polygoneDBResponse
func (controller *Controller) GetPolygone(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPolygone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	// Get polygoneDB in DB
	var polygoneDB orm.PolygoneDB
	if _, err := db.First(&polygoneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var polygoneAPI orm.PolygoneAPI
	polygoneAPI.ID = polygoneDB.ID
	polygoneAPI.PolygonePointersEncoding = polygoneDB.PolygonePointersEncoding
	polygoneDB.CopyBasicFieldsToPolygone_WOP(&polygoneAPI.Polygone_WOP)

	c.JSON(http.StatusOK, polygoneAPI)
}

// UpdatePolygone
//
// swagger:route PATCH /polygones/{ID} polygones updatePolygone
//
// # Update a polygone
//
// Responses:
// default: genericError
//
//	200: polygoneDBResponse
func (controller *Controller) UpdatePolygone(c *gin.Context) {

	mutexPolygone.Lock()
	defer mutexPolygone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePolygone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	// Validate input
	var input orm.PolygoneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var polygoneDB orm.PolygoneDB

	// fetch the polygone
	_, err := db.First(&polygoneDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	polygoneDB.CopyBasicFieldsFromPolygone_WOP(&input.Polygone_WOP)
	polygoneDB.PolygonePointersEncoding = input.PolygonePointersEncoding

	db, _ = db.Model(&polygoneDB)
	_, err = db.Updates(&polygoneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	polygoneNew := new(models.Polygone)
	polygoneDB.CopyBasicFieldsToPolygone(polygoneNew)

	// redeem pointers
	polygoneDB.DecodePointers(backRepo, polygoneNew)

	// get stage instance from DB instance, and call callback function
	polygoneOld := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
	if polygoneOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), polygoneOld, polygoneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polygoneDB
	c.JSON(http.StatusOK, polygoneDB)
}

// DeletePolygone
//
// swagger:route DELETE /polygones/{ID} polygones deletePolygone
//
// # Delete a polygone
//
// default: genericError
//
//	200: polygoneDBResponse
func (controller *Controller) DeletePolygone(c *gin.Context) {

	mutexPolygone.Lock()
	defer mutexPolygone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePolygone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolygone.GetDB()

	// Get model if exist
	var polygoneDB orm.PolygoneDB
	if _, err := db.First(&polygoneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&polygoneDB)

	// get an instance (not staged) from DB instance, and call callback function
	polygoneDeleted := new(models.Polygone)
	polygoneDB.CopyBasicFieldsToPolygone(polygoneDeleted)

	// get stage instance from DB instance, and call callback function
	polygoneStaged := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[polygoneDB.ID]
	if polygoneStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), polygoneStaged, polygoneDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
