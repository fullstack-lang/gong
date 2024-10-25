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
var __Point__dummysDeclaration__ models.Point
var __Point_time__dummyDeclaration time.Duration

var mutexPoint sync.Mutex

// An PointID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPoint updatePoint deletePoint
type PointID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PointInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPoint updatePoint
type PointInput struct {
	// The Point to submit or modify
	// in: body
	Point *orm.PointAPI
}

// GetPoints
//
// swagger:route GET /points points getPoints
//
// # Get all points
//
// Responses:
// default: genericError
//
//	200: pointDBResponse
func (controller *Controller) GetPoints(c *gin.Context) {

	// source slice
	var pointDBs []orm.PointDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPoints", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPoint.GetDB()

	_, err := db.Find(&pointDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pointAPIs := make([]orm.PointAPI, 0)

	// for each point, update fields from the database nullable fields
	for idx := range pointDBs {
		pointDB := &pointDBs[idx]
		_ = pointDB
		var pointAPI orm.PointAPI

		// insertion point for updating fields
		pointAPI.ID = pointDB.ID
		pointDB.CopyBasicFieldsToPoint_WOP(&pointAPI.Point_WOP)
		pointAPI.PointPointersEncoding = pointDB.PointPointersEncoding
		pointAPIs = append(pointAPIs, pointAPI)
	}

	c.JSON(http.StatusOK, pointAPIs)
}

// PostPoint
//
// swagger:route POST /points points postPoint
//
// Creates a point
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPoint(c *gin.Context) {

	mutexPoint.Lock()
	defer mutexPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPoints", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPoint.GetDB()

	// Validate input
	var input orm.PointAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create point
	pointDB := orm.PointDB{}
	pointDB.PointPointersEncoding = input.PointPointersEncoding
	pointDB.CopyBasicFieldsFromPoint_WOP(&input.Point_WOP)

	_, err = db.Create(&pointDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPoint.CheckoutPhaseOneInstance(&pointDB)
	point := backRepo.BackRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]

	if point != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), point)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pointDB)
}

// GetPoint
//
// swagger:route GET /points/{ID} points getPoint
//
// Gets the details for a point.
//
// Responses:
// default: genericError
//
//	200: pointDBResponse
func (controller *Controller) GetPoint(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPoint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPoint.GetDB()

	// Get pointDB in DB
	var pointDB orm.PointDB
	if _, err := db.First(&pointDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pointAPI orm.PointAPI
	pointAPI.ID = pointDB.ID
	pointAPI.PointPointersEncoding = pointDB.PointPointersEncoding
	pointDB.CopyBasicFieldsToPoint_WOP(&pointAPI.Point_WOP)

	c.JSON(http.StatusOK, pointAPI)
}

// UpdatePoint
//
// swagger:route PATCH /points/{ID} points updatePoint
//
// # Update a point
//
// Responses:
// default: genericError
//
//	200: pointDBResponse
func (controller *Controller) UpdatePoint(c *gin.Context) {

	mutexPoint.Lock()
	defer mutexPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePoint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPoint.GetDB()

	// Validate input
	var input orm.PointAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pointDB orm.PointDB

	// fetch the point
	_, err := db.First(&pointDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pointDB.CopyBasicFieldsFromPoint_WOP(&input.Point_WOP)
	pointDB.PointPointersEncoding = input.PointPointersEncoding

	db, _ = db.Model(&pointDB)
	_, err = db.Updates(&pointDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pointNew := new(models.Point)
	pointDB.CopyBasicFieldsToPoint(pointNew)

	// redeem pointers
	pointDB.DecodePointers(backRepo, pointNew)

	// get stage instance from DB instance, and call callback function
	pointOld := backRepo.BackRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
	if pointOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pointOld, pointNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pointDB
	c.JSON(http.StatusOK, pointDB)
}

// DeletePoint
//
// swagger:route DELETE /points/{ID} points deletePoint
//
// # Delete a point
//
// default: genericError
//
//	200: pointDBResponse
func (controller *Controller) DeletePoint(c *gin.Context) {

	mutexPoint.Lock()
	defer mutexPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePoint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPoint.GetDB()

	// Get model if exist
	var pointDB orm.PointDB
	if _, err := db.First(&pointDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&pointDB)

	// get an instance (not staged) from DB instance, and call callback function
	pointDeleted := new(models.Point)
	pointDB.CopyBasicFieldsToPoint(pointDeleted)

	// get stage instance from DB instance, and call callback function
	pointStaged := backRepo.BackRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
	if pointStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pointStaged, pointDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
