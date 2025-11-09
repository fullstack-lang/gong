// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __ControlPoint__dummysDeclaration__ models.ControlPoint
var __ControlPoint_time__dummyDeclaration time.Duration

var mutexControlPoint sync.Mutex

// An ControlPointID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getControlPoint updateControlPoint deleteControlPoint
type ControlPointID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ControlPointInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postControlPoint updateControlPoint
type ControlPointInput struct {
	// The ControlPoint to submit or modify
	// in: body
	ControlPoint *orm.ControlPointAPI
}

// GetControlPoints
//
// swagger:route GET /controlpoints controlpoints getControlPoints
//
// # Get all controlpoints
//
// Responses:
// default: genericError
//
//	200: controlpointDBResponse
func (controller *Controller) GetControlPoints(c *gin.Context) {

	// source slice
	var controlpointDBs []orm.ControlPointDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetControlPoints", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoControlPoint.GetDB()

	_, err := db.Find(&controlpointDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	controlpointAPIs := make([]orm.ControlPointAPI, 0)

	// for each controlpoint, update fields from the database nullable fields
	for idx := range controlpointDBs {
		controlpointDB := &controlpointDBs[idx]
		_ = controlpointDB
		var controlpointAPI orm.ControlPointAPI

		// insertion point for updating fields
		controlpointAPI.ID = controlpointDB.ID
		controlpointDB.CopyBasicFieldsToControlPoint_WOP(&controlpointAPI.ControlPoint_WOP)
		controlpointAPI.ControlPointPointersEncoding = controlpointDB.ControlPointPointersEncoding
		controlpointAPIs = append(controlpointAPIs, controlpointAPI)
	}

	c.JSON(http.StatusOK, controlpointAPIs)
}

// PostControlPoint
//
// swagger:route POST /controlpoints controlpoints postControlPoint
//
// Creates a controlpoint
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostControlPoint(c *gin.Context) {

	mutexControlPoint.Lock()
	defer mutexControlPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostControlPoints", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoControlPoint.GetDB()

	// Validate input
	var input orm.ControlPointAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create controlpoint
	controlpointDB := orm.ControlPointDB{}
	controlpointDB.ControlPointPointersEncoding = input.ControlPointPointersEncoding
	controlpointDB.CopyBasicFieldsFromControlPoint_WOP(&input.ControlPoint_WOP)

	_, err = db.Create(&controlpointDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoControlPoint.CheckoutPhaseOneInstance(&controlpointDB)
	controlpoint := backRepo.BackRepoControlPoint.Map_ControlPointDBID_ControlPointPtr[controlpointDB.ID]

	if controlpoint != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), controlpoint)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, controlpointDB)
}

// GetControlPoint
//
// swagger:route GET /controlpoints/{ID} controlpoints getControlPoint
//
// Gets the details for a controlpoint.
//
// Responses:
// default: genericError
//
//	200: controlpointDBResponse
func (controller *Controller) GetControlPoint(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetControlPoint", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoControlPoint.GetDB()

	// Get controlpointDB in DB
	var controlpointDB orm.ControlPointDB
	if _, err := db.First(&controlpointDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var controlpointAPI orm.ControlPointAPI
	controlpointAPI.ID = controlpointDB.ID
	controlpointAPI.ControlPointPointersEncoding = controlpointDB.ControlPointPointersEncoding
	controlpointDB.CopyBasicFieldsToControlPoint_WOP(&controlpointAPI.ControlPoint_WOP)

	c.JSON(http.StatusOK, controlpointAPI)
}

// UpdateControlPoint
//
// swagger:route PATCH /controlpoints/{ID} controlpoints updateControlPoint
//
// # Update a controlpoint
//
// Responses:
// default: genericError
//
//	200: controlpointDBResponse
func (controller *Controller) UpdateControlPoint(c *gin.Context) {

	mutexControlPoint.Lock()
	defer mutexControlPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoControlPoint.GetDB()

	// Validate input
	var input orm.ControlPointAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var controlpointDB orm.ControlPointDB

	// fetch the controlpoint
	_, err := db.First(&controlpointDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	controlpointDB.CopyBasicFieldsFromControlPoint_WOP(&input.ControlPoint_WOP)
	controlpointDB.ControlPointPointersEncoding = input.ControlPointPointersEncoding

	db, _ = db.Model(&controlpointDB)
	_, err = db.Updates(&controlpointDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	controlpointNew := new(models.ControlPoint)
	controlpointDB.CopyBasicFieldsToControlPoint(controlpointNew)

	// redeem pointers
	controlpointDB.DecodePointers(backRepo, controlpointNew)

	// get stage instance from DB instance, and call callback function
	controlpointOld := backRepo.BackRepoControlPoint.Map_ControlPointDBID_ControlPointPtr[controlpointDB.ID]
	if controlpointOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), controlpointOld, controlpointNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), controlpointOld, controlpointNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the controlpointDB
	c.JSON(http.StatusOK, controlpointDB)
}

// DeleteControlPoint
//
// swagger:route DELETE /controlpoints/{ID} controlpoints deleteControlPoint
//
// # Delete a controlpoint
//
// default: genericError
//
//	200: controlpointDBResponse
func (controller *Controller) DeleteControlPoint(c *gin.Context) {

	mutexControlPoint.Lock()
	defer mutexControlPoint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteControlPoint", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoControlPoint.GetDB()

	// Get model if exist
	var controlpointDB orm.ControlPointDB
	if _, err := db.First(&controlpointDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&controlpointDB)

	// get an instance (not staged) from DB instance, and call callback function
	controlpointDeleted := new(models.ControlPoint)
	controlpointDB.CopyBasicFieldsToControlPoint(controlpointDeleted)

	// get stage instance from DB instance, and call callback function
	controlpointStaged := backRepo.BackRepoControlPoint.Map_ControlPointDBID_ControlPointPtr[controlpointDB.ID]
	if controlpointStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), controlpointStaged, controlpointDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
