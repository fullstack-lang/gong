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
var __Polyline__dummysDeclaration__ models.Polyline
var __Polyline_time__dummyDeclaration time.Duration

var mutexPolyline sync.Mutex

// An PolylineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPolyline updatePolyline deletePolyline
type PolylineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolylineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPolyline updatePolyline
type PolylineInput struct {
	// The Polyline to submit or modify
	// in: body
	Polyline *orm.PolylineAPI
}

// GetPolylines
//
// swagger:route GET /polylines polylines getPolylines
//
// # Get all polylines
//
// Responses:
// default: genericError
//
//	200: polylineDBResponse
func (controller *Controller) GetPolylines(c *gin.Context) {

	// source slice
	var polylineDBs []orm.PolylineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPolylines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolyline.GetDB()

	_, err := db.Find(&polylineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	polylineAPIs := make([]orm.PolylineAPI, 0)

	// for each polyline, update fields from the database nullable fields
	for idx := range polylineDBs {
		polylineDB := &polylineDBs[idx]
		_ = polylineDB
		var polylineAPI orm.PolylineAPI

		// insertion point for updating fields
		polylineAPI.ID = polylineDB.ID
		polylineDB.CopyBasicFieldsToPolyline_WOP(&polylineAPI.Polyline_WOP)
		polylineAPI.PolylinePointersEncoding = polylineDB.PolylinePointersEncoding
		polylineAPIs = append(polylineAPIs, polylineAPI)
	}

	c.JSON(http.StatusOK, polylineAPIs)
}

// PostPolyline
//
// swagger:route POST /polylines polylines postPolyline
//
// Creates a polyline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPolyline(c *gin.Context) {

	mutexPolyline.Lock()
	defer mutexPolyline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPolylines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolyline.GetDB()

	// Validate input
	var input orm.PolylineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create polyline
	polylineDB := orm.PolylineDB{}
	polylineDB.PolylinePointersEncoding = input.PolylinePointersEncoding
	polylineDB.CopyBasicFieldsFromPolyline_WOP(&input.Polyline_WOP)

	_, err = db.Create(&polylineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPolyline.CheckoutPhaseOneInstance(&polylineDB)
	polyline := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]

	if polyline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), polyline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, polylineDB)
}

// GetPolyline
//
// swagger:route GET /polylines/{ID} polylines getPolyline
//
// Gets the details for a polyline.
//
// Responses:
// default: genericError
//
//	200: polylineDBResponse
func (controller *Controller) GetPolyline(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPolyline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolyline.GetDB()

	// Get polylineDB in DB
	var polylineDB orm.PolylineDB
	if _, err := db.First(&polylineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var polylineAPI orm.PolylineAPI
	polylineAPI.ID = polylineDB.ID
	polylineAPI.PolylinePointersEncoding = polylineDB.PolylinePointersEncoding
	polylineDB.CopyBasicFieldsToPolyline_WOP(&polylineAPI.Polyline_WOP)

	c.JSON(http.StatusOK, polylineAPI)
}

// UpdatePolyline
//
// swagger:route PATCH /polylines/{ID} polylines updatePolyline
//
// # Update a polyline
//
// Responses:
// default: genericError
//
//	200: polylineDBResponse
func (controller *Controller) UpdatePolyline(c *gin.Context) {

	mutexPolyline.Lock()
	defer mutexPolyline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePolyline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolyline.GetDB()

	// Validate input
	var input orm.PolylineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var polylineDB orm.PolylineDB

	// fetch the polyline
	_, err := db.First(&polylineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	polylineDB.CopyBasicFieldsFromPolyline_WOP(&input.Polyline_WOP)
	polylineDB.PolylinePointersEncoding = input.PolylinePointersEncoding

	db, _ = db.Model(&polylineDB)
	_, err = db.Updates(&polylineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	polylineNew := new(models.Polyline)
	polylineDB.CopyBasicFieldsToPolyline(polylineNew)

	// redeem pointers
	polylineDB.DecodePointers(backRepo, polylineNew)

	// get stage instance from DB instance, and call callback function
	polylineOld := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]
	if polylineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), polylineOld, polylineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polylineDB
	c.JSON(http.StatusOK, polylineDB)
}

// DeletePolyline
//
// swagger:route DELETE /polylines/{ID} polylines deletePolyline
//
// # Delete a polyline
//
// default: genericError
//
//	200: polylineDBResponse
func (controller *Controller) DeletePolyline(c *gin.Context) {

	mutexPolyline.Lock()
	defer mutexPolyline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePolyline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPolyline.GetDB()

	// Get model if exist
	var polylineDB orm.PolylineDB
	if _, err := db.First(&polylineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&polylineDB)

	// get an instance (not staged) from DB instance, and call callback function
	polylineDeleted := new(models.Polyline)
	polylineDB.CopyBasicFieldsToPolyline(polylineDeleted)

	// get stage instance from DB instance, and call callback function
	polylineStaged := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[polylineDB.ID]
	if polylineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), polylineStaged, polylineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
