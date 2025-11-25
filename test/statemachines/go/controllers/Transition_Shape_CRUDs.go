// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Transition_Shape__dummysDeclaration__ models.Transition_Shape
var __Transition_Shape_time__dummyDeclaration time.Duration

var mutexTransition_Shape sync.Mutex

// An Transition_ShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTransition_Shape updateTransition_Shape deleteTransition_Shape
type Transition_ShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Transition_ShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTransition_Shape updateTransition_Shape
type Transition_ShapeInput struct {
	// The Transition_Shape to submit or modify
	// in: body
	Transition_Shape *orm.Transition_ShapeAPI
}

// GetTransition_Shapes
//
// swagger:route GET /transition_shapes transition_shapes getTransition_Shapes
//
// # Get all transition_shapes
//
// Responses:
// default: genericError
//
//	200: transition_shapeDBResponse
func (controller *Controller) GetTransition_Shapes(c *gin.Context) {

	// source slice
	var transition_shapeDBs []orm.Transition_ShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTransition_Shapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition_Shape.GetDB()

	_, err := db.Find(&transition_shapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	transition_shapeAPIs := make([]orm.Transition_ShapeAPI, 0)

	// for each transition_shape, update fields from the database nullable fields
	for idx := range transition_shapeDBs {
		transition_shapeDB := &transition_shapeDBs[idx]
		_ = transition_shapeDB
		var transition_shapeAPI orm.Transition_ShapeAPI

		// insertion point for updating fields
		transition_shapeAPI.ID = transition_shapeDB.ID
		transition_shapeDB.CopyBasicFieldsToTransition_Shape_WOP(&transition_shapeAPI.Transition_Shape_WOP)
		transition_shapeAPI.Transition_ShapePointersEncoding = transition_shapeDB.Transition_ShapePointersEncoding
		transition_shapeAPIs = append(transition_shapeAPIs, transition_shapeAPI)
	}

	c.JSON(http.StatusOK, transition_shapeAPIs)
}

// PostTransition_Shape
//
// swagger:route POST /transition_shapes transition_shapes postTransition_Shape
//
// Creates a transition_shape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTransition_Shape(c *gin.Context) {

	mutexTransition_Shape.Lock()
	defer mutexTransition_Shape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTransition_Shapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition_Shape.GetDB()

	// Validate input
	var input orm.Transition_ShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create transition_shape
	transition_shapeDB := orm.Transition_ShapeDB{}
	transition_shapeDB.Transition_ShapePointersEncoding = input.Transition_ShapePointersEncoding
	transition_shapeDB.CopyBasicFieldsFromTransition_Shape_WOP(&input.Transition_Shape_WOP)

	_, err = db.Create(&transition_shapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTransition_Shape.CheckoutPhaseOneInstance(&transition_shapeDB)
	transition_shape := backRepo.BackRepoTransition_Shape.Map_Transition_ShapeDBID_Transition_ShapePtr[transition_shapeDB.ID]

	if transition_shape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), transition_shape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, transition_shapeDB)
}

// GetTransition_Shape
//
// swagger:route GET /transition_shapes/{ID} transition_shapes getTransition_Shape
//
// Gets the details for a transition_shape.
//
// Responses:
// default: genericError
//
//	200: transition_shapeDBResponse
func (controller *Controller) GetTransition_Shape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTransition_Shape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition_Shape.GetDB()

	// Get transition_shapeDB in DB
	var transition_shapeDB orm.Transition_ShapeDB
	if _, err := db.First(&transition_shapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var transition_shapeAPI orm.Transition_ShapeAPI
	transition_shapeAPI.ID = transition_shapeDB.ID
	transition_shapeAPI.Transition_ShapePointersEncoding = transition_shapeDB.Transition_ShapePointersEncoding
	transition_shapeDB.CopyBasicFieldsToTransition_Shape_WOP(&transition_shapeAPI.Transition_Shape_WOP)

	c.JSON(http.StatusOK, transition_shapeAPI)
}

// UpdateTransition_Shape
//
// swagger:route PATCH /transition_shapes/{ID} transition_shapes updateTransition_Shape
//
// # Update a transition_shape
//
// Responses:
// default: genericError
//
//	200: transition_shapeDBResponse
func (controller *Controller) UpdateTransition_Shape(c *gin.Context) {

	mutexTransition_Shape.Lock()
	defer mutexTransition_Shape.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition_Shape.GetDB()

	// Validate input
	var input orm.Transition_ShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var transition_shapeDB orm.Transition_ShapeDB

	// fetch the transition_shape
	_, err := db.First(&transition_shapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	transition_shapeDB.CopyBasicFieldsFromTransition_Shape_WOP(&input.Transition_Shape_WOP)
	transition_shapeDB.Transition_ShapePointersEncoding = input.Transition_ShapePointersEncoding

	db, _ = db.Model(&transition_shapeDB)
	_, err = db.Updates(&transition_shapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	transition_shapeNew := new(models.Transition_Shape)
	transition_shapeDB.CopyBasicFieldsToTransition_Shape(transition_shapeNew)

	// redeem pointers
	transition_shapeDB.DecodePointers(backRepo, transition_shapeNew)

	// get stage instance from DB instance, and call callback function
	transition_shapeOld := backRepo.BackRepoTransition_Shape.Map_Transition_ShapeDBID_Transition_ShapePtr[transition_shapeDB.ID]
	if transition_shapeOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), transition_shapeOld, transition_shapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the transition_shapeDB
	c.JSON(http.StatusOK, transition_shapeDB)
}

// DeleteTransition_Shape
//
// swagger:route DELETE /transition_shapes/{ID} transition_shapes deleteTransition_Shape
//
// # Delete a transition_shape
//
// default: genericError
//
//	200: transition_shapeDBResponse
func (controller *Controller) DeleteTransition_Shape(c *gin.Context) {

	mutexTransition_Shape.Lock()
	defer mutexTransition_Shape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTransition_Shape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTransition_Shape.GetDB()

	// Get model if exist
	var transition_shapeDB orm.Transition_ShapeDB
	if _, err := db.First(&transition_shapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&transition_shapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	transition_shapeDeleted := new(models.Transition_Shape)
	transition_shapeDB.CopyBasicFieldsToTransition_Shape(transition_shapeDeleted)

	// get stage instance from DB instance, and call callback function
	transition_shapeStaged := backRepo.BackRepoTransition_Shape.Map_Transition_ShapeDBID_Transition_ShapePtr[transition_shapeDB.ID]
	if transition_shapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), transition_shapeStaged, transition_shapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
