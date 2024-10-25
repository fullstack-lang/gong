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
var __Circle__dummysDeclaration__ models.Circle
var __Circle_time__dummyDeclaration time.Duration

var mutexCircle sync.Mutex

// An CircleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCircle updateCircle deleteCircle
type CircleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CircleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCircle updateCircle
type CircleInput struct {
	// The Circle to submit or modify
	// in: body
	Circle *orm.CircleAPI
}

// GetCircles
//
// swagger:route GET /circles circles getCircles
//
// # Get all circles
//
// Responses:
// default: genericError
//
//	200: circleDBResponse
func (controller *Controller) GetCircles(c *gin.Context) {

	// source slice
	var circleDBs []orm.CircleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCircles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircle.GetDB()

	_, err := db.Find(&circleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	circleAPIs := make([]orm.CircleAPI, 0)

	// for each circle, update fields from the database nullable fields
	for idx := range circleDBs {
		circleDB := &circleDBs[idx]
		_ = circleDB
		var circleAPI orm.CircleAPI

		// insertion point for updating fields
		circleAPI.ID = circleDB.ID
		circleDB.CopyBasicFieldsToCircle_WOP(&circleAPI.Circle_WOP)
		circleAPI.CirclePointersEncoding = circleDB.CirclePointersEncoding
		circleAPIs = append(circleAPIs, circleAPI)
	}

	c.JSON(http.StatusOK, circleAPIs)
}

// PostCircle
//
// swagger:route POST /circles circles postCircle
//
// Creates a circle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCircle(c *gin.Context) {

	mutexCircle.Lock()
	defer mutexCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCircles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircle.GetDB()

	// Validate input
	var input orm.CircleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create circle
	circleDB := orm.CircleDB{}
	circleDB.CirclePointersEncoding = input.CirclePointersEncoding
	circleDB.CopyBasicFieldsFromCircle_WOP(&input.Circle_WOP)

	_, err = db.Create(&circleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCircle.CheckoutPhaseOneInstance(&circleDB)
	circle := backRepo.BackRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]

	if circle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), circle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, circleDB)
}

// GetCircle
//
// swagger:route GET /circles/{ID} circles getCircle
//
// Gets the details for a circle.
//
// Responses:
// default: genericError
//
//	200: circleDBResponse
func (controller *Controller) GetCircle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircle.GetDB()

	// Get circleDB in DB
	var circleDB orm.CircleDB
	if _, err := db.First(&circleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var circleAPI orm.CircleAPI
	circleAPI.ID = circleDB.ID
	circleAPI.CirclePointersEncoding = circleDB.CirclePointersEncoding
	circleDB.CopyBasicFieldsToCircle_WOP(&circleAPI.Circle_WOP)

	c.JSON(http.StatusOK, circleAPI)
}

// UpdateCircle
//
// swagger:route PATCH /circles/{ID} circles updateCircle
//
// # Update a circle
//
// Responses:
// default: genericError
//
//	200: circleDBResponse
func (controller *Controller) UpdateCircle(c *gin.Context) {

	mutexCircle.Lock()
	defer mutexCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircle.GetDB()

	// Validate input
	var input orm.CircleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var circleDB orm.CircleDB

	// fetch the circle
	_, err := db.First(&circleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	circleDB.CopyBasicFieldsFromCircle_WOP(&input.Circle_WOP)
	circleDB.CirclePointersEncoding = input.CirclePointersEncoding

	db, _ = db.Model(&circleDB)
	_, err = db.Updates(&circleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	circleNew := new(models.Circle)
	circleDB.CopyBasicFieldsToCircle(circleNew)

	// redeem pointers
	circleDB.DecodePointers(backRepo, circleNew)

	// get stage instance from DB instance, and call callback function
	circleOld := backRepo.BackRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]
	if circleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), circleOld, circleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the circleDB
	c.JSON(http.StatusOK, circleDB)
}

// DeleteCircle
//
// swagger:route DELETE /circles/{ID} circles deleteCircle
//
// # Delete a circle
//
// default: genericError
//
//	200: circleDBResponse
func (controller *Controller) DeleteCircle(c *gin.Context) {

	mutexCircle.Lock()
	defer mutexCircle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCircle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCircle.GetDB()

	// Get model if exist
	var circleDB orm.CircleDB
	if _, err := db.First(&circleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&circleDB)

	// get an instance (not staged) from DB instance, and call callback function
	circleDeleted := new(models.Circle)
	circleDB.CopyBasicFieldsToCircle(circleDeleted)

	// get stage instance from DB instance, and call callback function
	circleStaged := backRepo.BackRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]
	if circleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), circleStaged, circleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
