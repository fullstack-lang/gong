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
var __Position__dummysDeclaration__ models.Position
var __Position_time__dummyDeclaration time.Duration

var mutexPosition sync.Mutex

// An PositionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPosition updatePosition deletePosition
type PositionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PositionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPosition updatePosition
type PositionInput struct {
	// The Position to submit or modify
	// in: body
	Position *orm.PositionAPI
}

// GetPositions
//
// swagger:route GET /positions positions getPositions
//
// # Get all positions
//
// Responses:
// default: genericError
//
//	200: positionDBResponse
func (controller *Controller) GetPositions(c *gin.Context) {

	// source slice
	var positionDBs []orm.PositionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPositions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPosition.GetDB()

	_, err := db.Find(&positionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	positionAPIs := make([]orm.PositionAPI, 0)

	// for each position, update fields from the database nullable fields
	for idx := range positionDBs {
		positionDB := &positionDBs[idx]
		_ = positionDB
		var positionAPI orm.PositionAPI

		// insertion point for updating fields
		positionAPI.ID = positionDB.ID
		positionDB.CopyBasicFieldsToPosition_WOP(&positionAPI.Position_WOP)
		positionAPI.PositionPointersEncoding = positionDB.PositionPointersEncoding
		positionAPIs = append(positionAPIs, positionAPI)
	}

	c.JSON(http.StatusOK, positionAPIs)
}

// PostPosition
//
// swagger:route POST /positions positions postPosition
//
// Creates a position
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPosition(c *gin.Context) {

	mutexPosition.Lock()
	defer mutexPosition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPositions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPosition.GetDB()

	// Validate input
	var input orm.PositionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create position
	positionDB := orm.PositionDB{}
	positionDB.PositionPointersEncoding = input.PositionPointersEncoding
	positionDB.CopyBasicFieldsFromPosition_WOP(&input.Position_WOP)

	_, err = db.Create(&positionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPosition.CheckoutPhaseOneInstance(&positionDB)
	position := backRepo.BackRepoPosition.Map_PositionDBID_PositionPtr[positionDB.ID]

	if position != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), position)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, positionDB)
}

// GetPosition
//
// swagger:route GET /positions/{ID} positions getPosition
//
// Gets the details for a position.
//
// Responses:
// default: genericError
//
//	200: positionDBResponse
func (controller *Controller) GetPosition(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPosition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPosition.GetDB()

	// Get positionDB in DB
	var positionDB orm.PositionDB
	if _, err := db.First(&positionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var positionAPI orm.PositionAPI
	positionAPI.ID = positionDB.ID
	positionAPI.PositionPointersEncoding = positionDB.PositionPointersEncoding
	positionDB.CopyBasicFieldsToPosition_WOP(&positionAPI.Position_WOP)

	c.JSON(http.StatusOK, positionAPI)
}

// UpdatePosition
//
// swagger:route PATCH /positions/{ID} positions updatePosition
//
// # Update a position
//
// Responses:
// default: genericError
//
//	200: positionDBResponse
func (controller *Controller) UpdatePosition(c *gin.Context) {

	mutexPosition.Lock()
	defer mutexPosition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePosition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPosition.GetDB()

	// Validate input
	var input orm.PositionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var positionDB orm.PositionDB

	// fetch the position
	_, err := db.First(&positionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	positionDB.CopyBasicFieldsFromPosition_WOP(&input.Position_WOP)
	positionDB.PositionPointersEncoding = input.PositionPointersEncoding

	db, _ = db.Model(&positionDB)
	_, err = db.Updates(&positionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	positionNew := new(models.Position)
	positionDB.CopyBasicFieldsToPosition(positionNew)

	// redeem pointers
	positionDB.DecodePointers(backRepo, positionNew)

	// get stage instance from DB instance, and call callback function
	positionOld := backRepo.BackRepoPosition.Map_PositionDBID_PositionPtr[positionDB.ID]
	if positionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), positionOld, positionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the positionDB
	c.JSON(http.StatusOK, positionDB)
}

// DeletePosition
//
// swagger:route DELETE /positions/{ID} positions deletePosition
//
// # Delete a position
//
// default: genericError
//
//	200: positionDBResponse
func (controller *Controller) DeletePosition(c *gin.Context) {

	mutexPosition.Lock()
	defer mutexPosition.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePosition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPosition.GetDB()

	// Get model if exist
	var positionDB orm.PositionDB
	if _, err := db.First(&positionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&positionDB)

	// get an instance (not staged) from DB instance, and call callback function
	positionDeleted := new(models.Position)
	positionDB.CopyBasicFieldsToPosition(positionDeleted)

	// get stage instance from DB instance, and call callback function
	positionStaged := backRepo.BackRepoPosition.Map_PositionDBID_PositionPtr[positionDB.ID]
	if positionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), positionStaged, positionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
