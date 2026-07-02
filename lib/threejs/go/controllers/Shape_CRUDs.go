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
var __Shape__dummysDeclaration__ models.Shape
var _ = __Shape__dummysDeclaration__
var __Shape_time__dummyDeclaration time.Duration
var _ = __Shape_time__dummyDeclaration

var mutexShape sync.Mutex

// An ShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getShape updateShape deleteShape
type ShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postShape updateShape
type ShapeInput struct {
	// The Shape to submit or modify
	// in: body
	Shape *orm.ShapeAPI
}

// GetShapes
//
// swagger:route GET /shapes shapes getShapes
//
// # Get all shapes
//
// Responses:
// default: genericError
//
//	200: shapeDBResponse
func (controller *Controller) GetShapes(c *gin.Context) {

	// source slice
	var shapeDBs []orm.ShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetShapes", "Name", stackPath)
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
	db := backRepo.BackRepoShape.GetDB()

	_, err := db.Find(&shapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	shapeAPIs := make([]orm.ShapeAPI, 0)

	// for each shape, update fields from the database nullable fields
	for idx := range shapeDBs {
		shapeDB := &shapeDBs[idx]
		_ = shapeDB
		var shapeAPI orm.ShapeAPI

		// insertion point for updating fields
		shapeAPI.ID = shapeDB.ID
		shapeDB.CopyBasicFieldsToShape_WOP(&shapeAPI.Shape_WOP)
		shapeAPI.ShapePointersEncoding = shapeDB.ShapePointersEncoding
		shapeAPIs = append(shapeAPIs, shapeAPI)
	}

	c.JSON(http.StatusOK, shapeAPIs)
}

// PostShape
//
// swagger:route POST /shapes shapes postShape
//
// Creates a shape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostShape(c *gin.Context) {

	mutexShape.Lock()
	defer mutexShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostShapes", "Name", stackPath)
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
	db := backRepo.BackRepoShape.GetDB()

	// Validate input
	var input orm.ShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create shape
	shapeDB := orm.ShapeDB{}
	shapeDB.ShapePointersEncoding = input.ShapePointersEncoding
	shapeDB.CopyBasicFieldsFromShape_WOP(&input.Shape_WOP)

	_, err = db.Create(&shapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoShape.CheckoutPhaseOneInstance(&shapeDB)
	shape := backRepo.BackRepoShape.Map_ShapeDBID_ShapePtr[shapeDB.ID]

	if shape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), shape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, shapeDB)
}

// GetShape
//
// swagger:route GET /shapes/{ID} shapes getShape
//
// Gets the details for a shape.
//
// Responses:
// default: genericError
//
//	200: shapeDBResponse
func (controller *Controller) GetShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetShape", "Name", stackPath)
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
	db := backRepo.BackRepoShape.GetDB()

	// Get shapeDB in DB
	var shapeDB orm.ShapeDB
	if _, err := db.First(&shapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var shapeAPI orm.ShapeAPI
	shapeAPI.ID = shapeDB.ID
	shapeAPI.ShapePointersEncoding = shapeDB.ShapePointersEncoding
	shapeDB.CopyBasicFieldsToShape_WOP(&shapeAPI.Shape_WOP)

	c.JSON(http.StatusOK, shapeAPI)
}

// UpdateShape
//
// swagger:route PATCH /shapes/{ID} shapes updateShape
//
// # Update a shape
//
// Responses:
// default: genericError
//
//	200: shapeDBResponse
func (controller *Controller) UpdateShape(c *gin.Context) {

	mutexShape.Lock()
	defer mutexShape.Unlock()

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
	db := backRepo.BackRepoShape.GetDB()

	// Validate input
	var input orm.ShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var shapeDB orm.ShapeDB

	// fetch the shape
	_, err := db.First(&shapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	shapeDB.CopyBasicFieldsFromShape_WOP(&input.Shape_WOP)
	shapeDB.ShapePointersEncoding = input.ShapePointersEncoding

	db, _ = db.Model(&shapeDB)
	_, err = db.Updates(&shapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	shapeNew := new(models.Shape)
	shapeDB.CopyBasicFieldsToShape(shapeNew)

	// redeem pointers
	shapeDB.DecodePointers(backRepo, shapeNew)

	// get stage instance from DB instance, and call callback function
	shapeOld := backRepo.BackRepoShape.Map_ShapeDBID_ShapePtr[shapeDB.ID]
	if shapeOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), shapeOld, shapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the shapeDB
	c.JSON(http.StatusOK, shapeDB)
}

// DeleteShape
//
// swagger:route DELETE /shapes/{ID} shapes deleteShape
//
// # Delete a shape
//
// default: genericError
//
//	200: shapeDBResponse
func (controller *Controller) DeleteShape(c *gin.Context) {

	mutexShape.Lock()
	defer mutexShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteShape", "Name", stackPath)
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
	db := backRepo.BackRepoShape.GetDB()

	// Get model if exist
	var shapeDB orm.ShapeDB
	if _, err := db.First(&shapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&shapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	shapeDeleted := new(models.Shape)
	shapeDB.CopyBasicFieldsToShape(shapeDeleted)

	// get stage instance from DB instance, and call callback function
	shapeStaged := backRepo.BackRepoShape.Map_ShapeDBID_ShapePtr[shapeDB.ID]
	if shapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), shapeStaged, shapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
