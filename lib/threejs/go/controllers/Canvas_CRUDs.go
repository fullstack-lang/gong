// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Canvas__dummysDeclaration__ models.Canvas
var _ = __Canvas__dummysDeclaration__
var __Canvas_time__dummyDeclaration time.Duration
var _ = __Canvas_time__dummyDeclaration

var mutexCanvas sync.Mutex

// An CanvasID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCanvas updateCanvas deleteCanvas
type CanvasID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CanvasInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCanvas updateCanvas
type CanvasInput struct {
	// The Canvas to submit or modify
	// in: body
	Canvas *orm.CanvasAPI
}

// GetCanvass
//
// swagger:route GET /canvass canvass getCanvass
//
// # Get all canvass
//
// Responses:
// default: genericError
//
//	200: canvasDBResponse
func (controller *Controller) GetCanvass(c *gin.Context) {

	// source slice
	var canvasDBs []orm.CanvasDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCanvass", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCanvas.GetDB()

	_, err := db.Find(&canvasDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	canvasAPIs := make([]orm.CanvasAPI, 0)

	// for each canvas, update fields from the database nullable fields
	for idx := range canvasDBs {
		canvasDB := &canvasDBs[idx]
		_ = canvasDB
		var canvasAPI orm.CanvasAPI

		// insertion point for updating fields
		canvasAPI.ID = canvasDB.ID
		canvasDB.CopyBasicFieldsToCanvas_WOP(&canvasAPI.Canvas_WOP)
		canvasAPI.CanvasPointersEncoding = canvasDB.CanvasPointersEncoding
		canvasAPIs = append(canvasAPIs, canvasAPI)
	}

	c.JSON(http.StatusOK, canvasAPIs)
}

// PostCanvas
//
// swagger:route POST /canvass canvass postCanvas
//
// Creates a canvas
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCanvas(c *gin.Context) {

	mutexCanvas.Lock()
	defer mutexCanvas.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCanvass", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCanvas.GetDB()

	// Validate input
	var input orm.CanvasAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create canvas
	canvasDB := orm.CanvasDB{}
	canvasDB.CanvasPointersEncoding = input.CanvasPointersEncoding
	canvasDB.CopyBasicFieldsFromCanvas_WOP(&input.Canvas_WOP)

	_, err = db.Create(&canvasDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCanvas.CheckoutPhaseOneInstance(&canvasDB)
	canvas := backRepo.BackRepoCanvas.Map_CanvasDBID_CanvasPtr[canvasDB.ID]

	if canvas != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), canvas)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, canvasDB)
}

// GetCanvas
//
// swagger:route GET /canvass/{ID} canvass getCanvas
//
// Gets the details for a canvas.
//
// Responses:
// default: genericError
//
//	200: canvasDBResponse
func (controller *Controller) GetCanvas(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCanvas", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCanvas.GetDB()

	// Get canvasDB in DB
	var canvasDB orm.CanvasDB
	if _, err := db.First(&canvasDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var canvasAPI orm.CanvasAPI
	canvasAPI.ID = canvasDB.ID
	canvasAPI.CanvasPointersEncoding = canvasDB.CanvasPointersEncoding
	canvasDB.CopyBasicFieldsToCanvas_WOP(&canvasAPI.Canvas_WOP)

	c.JSON(http.StatusOK, canvasAPI)
}

// UpdateCanvas
//
// swagger:route PATCH /canvass/{ID} canvass updateCanvas
//
// # Update a canvas
//
// Responses:
// default: genericError
//
//	200: canvasDBResponse
func (controller *Controller) UpdateCanvas(c *gin.Context) {

	mutexCanvas.Lock()
	defer mutexCanvas.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCanvas.GetDB()

	// Validate input
	var input orm.CanvasAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var canvasDB orm.CanvasDB

	// fetch the canvas
	_, err := db.First(&canvasDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	canvasDB.CopyBasicFieldsFromCanvas_WOP(&input.Canvas_WOP)
	canvasDB.CanvasPointersEncoding = input.CanvasPointersEncoding

	db, _ = db.Model(&canvasDB)
	_, err = db.Updates(&canvasDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	canvasNew := new(models.Canvas)
	canvasDB.CopyBasicFieldsToCanvas(canvasNew)

	// redeem pointers
	canvasDB.DecodePointers(backRepo, canvasNew)

	// get stage instance from DB instance, and call callback function
	canvasOld := backRepo.BackRepoCanvas.Map_CanvasDBID_CanvasPtr[canvasDB.ID]
	if canvasOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), canvasOld, canvasNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the canvasDB
	c.JSON(http.StatusOK, canvasDB)
}

// DeleteCanvas
//
// swagger:route DELETE /canvass/{ID} canvass deleteCanvas
//
// # Delete a canvas
//
// default: genericError
//
//	200: canvasDBResponse
func (controller *Controller) DeleteCanvas(c *gin.Context) {

	mutexCanvas.Lock()
	defer mutexCanvas.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCanvas", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoCanvas.GetDB()

	// Get model if exist
	var canvasDB orm.CanvasDB
	if _, err := db.First(&canvasDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&canvasDB)

	// get an instance (not staged) from DB instance, and call callback function
	canvasDeleted := new(models.Canvas)
	canvasDB.CopyBasicFieldsToCanvas(canvasDeleted)

	// get stage instance from DB instance, and call callback function
	canvasStaged := backRepo.BackRepoCanvas.Map_CanvasDBID_CanvasPtr[canvasDB.ID]
	if canvasStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), canvasStaged, canvasDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
