// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	"github.com/fullstack-lang/gong/lib/doc2/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FieldShape__dummysDeclaration__ models.AttributeShape
var __FieldShape_time__dummyDeclaration time.Duration

var mutexFieldShape sync.Mutex

// An FieldShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFieldShape updateFieldShape deleteFieldShape
type FieldShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FieldShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFieldShape updateFieldShape
type FieldShapeInput struct {
	// The FieldShape to submit or modify
	// in: body
	FieldShape *orm.FieldShapeAPI
}

// GetFieldShapes
//
// swagger:route GET /fieldshapes fieldshapes getFieldShapes
//
// # Get all fieldshapes
//
// Responses:
// default: genericError
//
//	200: fieldshapeDBResponse
func (controller *Controller) GetFieldShapes(c *gin.Context) {

	// source slice
	var fieldshapeDBs []orm.FieldShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFieldShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFieldShape.GetDB()

	_, err := db.Find(&fieldshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fieldshapeAPIs := make([]orm.FieldShapeAPI, 0)

	// for each fieldshape, update fields from the database nullable fields
	for idx := range fieldshapeDBs {
		fieldshapeDB := &fieldshapeDBs[idx]
		_ = fieldshapeDB
		var fieldshapeAPI orm.FieldShapeAPI

		// insertion point for updating fields
		fieldshapeAPI.ID = fieldshapeDB.ID
		fieldshapeDB.CopyBasicFieldsToFieldShape_WOP(&fieldshapeAPI.FieldShape_WOP)
		fieldshapeAPI.FieldShapePointersEncoding = fieldshapeDB.FieldShapePointersEncoding
		fieldshapeAPIs = append(fieldshapeAPIs, fieldshapeAPI)
	}

	c.JSON(http.StatusOK, fieldshapeAPIs)
}

// PostFieldShape
//
// swagger:route POST /fieldshapes fieldshapes postFieldShape
//
// Creates a fieldshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFieldShape(c *gin.Context) {

	mutexFieldShape.Lock()
	defer mutexFieldShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFieldShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFieldShape.GetDB()

	// Validate input
	var input orm.FieldShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fieldshape
	fieldshapeDB := orm.FieldShapeDB{}
	fieldshapeDB.FieldShapePointersEncoding = input.FieldShapePointersEncoding
	fieldshapeDB.CopyBasicFieldsFromFieldShape_WOP(&input.FieldShape_WOP)

	_, err = db.Create(&fieldshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFieldShape.CheckoutPhaseOneInstance(&fieldshapeDB)
	fieldshape := backRepo.BackRepoFieldShape.Map_FieldShapeDBID_FieldShapePtr[fieldshapeDB.ID]

	if fieldshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fieldshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fieldshapeDB)
}

// GetFieldShape
//
// swagger:route GET /fieldshapes/{ID} fieldshapes getFieldShape
//
// Gets the details for a fieldshape.
//
// Responses:
// default: genericError
//
//	200: fieldshapeDBResponse
func (controller *Controller) GetFieldShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFieldShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFieldShape.GetDB()

	// Get fieldshapeDB in DB
	var fieldshapeDB orm.FieldShapeDB
	if _, err := db.First(&fieldshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fieldshapeAPI orm.FieldShapeAPI
	fieldshapeAPI.ID = fieldshapeDB.ID
	fieldshapeAPI.FieldShapePointersEncoding = fieldshapeDB.FieldShapePointersEncoding
	fieldshapeDB.CopyBasicFieldsToFieldShape_WOP(&fieldshapeAPI.FieldShape_WOP)

	c.JSON(http.StatusOK, fieldshapeAPI)
}

// UpdateFieldShape
//
// swagger:route PATCH /fieldshapes/{ID} fieldshapes updateFieldShape
//
// # Update a fieldshape
//
// Responses:
// default: genericError
//
//	200: fieldshapeDBResponse
func (controller *Controller) UpdateFieldShape(c *gin.Context) {

	mutexFieldShape.Lock()
	defer mutexFieldShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFieldShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFieldShape.GetDB()

	// Validate input
	var input orm.FieldShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fieldshapeDB orm.FieldShapeDB

	// fetch the fieldshape
	_, err := db.First(&fieldshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fieldshapeDB.CopyBasicFieldsFromFieldShape_WOP(&input.FieldShape_WOP)
	fieldshapeDB.FieldShapePointersEncoding = input.FieldShapePointersEncoding

	db, _ = db.Model(&fieldshapeDB)
	_, err = db.Updates(&fieldshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fieldshapeNew := new(models.AttributeShape)
	fieldshapeDB.CopyBasicFieldsToFieldShape(fieldshapeNew)

	// redeem pointers
	fieldshapeDB.DecodePointers(backRepo, fieldshapeNew)

	// get stage instance from DB instance, and call callback function
	fieldshapeOld := backRepo.BackRepoFieldShape.Map_FieldShapeDBID_FieldShapePtr[fieldshapeDB.ID]
	if fieldshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fieldshapeOld, fieldshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fieldshapeDB
	c.JSON(http.StatusOK, fieldshapeDB)
}

// DeleteFieldShape
//
// swagger:route DELETE /fieldshapes/{ID} fieldshapes deleteFieldShape
//
// # Delete a fieldshape
//
// default: genericError
//
//	200: fieldshapeDBResponse
func (controller *Controller) DeleteFieldShape(c *gin.Context) {

	mutexFieldShape.Lock()
	defer mutexFieldShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFieldShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFieldShape.GetDB()

	// Get model if exist
	var fieldshapeDB orm.FieldShapeDB
	if _, err := db.First(&fieldshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&fieldshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	fieldshapeDeleted := new(models.AttributeShape)
	fieldshapeDB.CopyBasicFieldsToFieldShape(fieldshapeDeleted)

	// get stage instance from DB instance, and call callback function
	fieldshapeStaged := backRepo.BackRepoFieldShape.Map_FieldShapeDBID_FieldShapePtr[fieldshapeDB.ID]
	if fieldshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fieldshapeStaged, fieldshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
