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
var __AttributeShape__dummysDeclaration__ models.AttributeShape
var __AttributeShape_time__dummyDeclaration time.Duration

var mutexAttributeShape sync.Mutex

// An AttributeShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAttributeShape updateAttributeShape deleteAttributeShape
type AttributeShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AttributeShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAttributeShape updateAttributeShape
type AttributeShapeInput struct {
	// The AttributeShape to submit or modify
	// in: body
	AttributeShape *orm.AttributeShapeAPI
}

// GetAttributeShapes
//
// swagger:route GET /attributeshapes attributeshapes getAttributeShapes
//
// # Get all attributeshapes
//
// Responses:
// default: genericError
//
//	200: attributeshapeDBResponse
func (controller *Controller) GetAttributeShapes(c *gin.Context) {

	// source slice
	var attributeshapeDBs []orm.AttributeShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributeShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAttributeShape.GetDB()

	_, err := db.Find(&attributeshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributeshapeAPIs := make([]orm.AttributeShapeAPI, 0)

	// for each attributeshape, update fields from the database nullable fields
	for idx := range attributeshapeDBs {
		attributeshapeDB := &attributeshapeDBs[idx]
		_ = attributeshapeDB
		var attributeshapeAPI orm.AttributeShapeAPI

		// insertion point for updating fields
		attributeshapeAPI.ID = attributeshapeDB.ID
		attributeshapeDB.CopyBasicFieldsToAttributeShape_WOP(&attributeshapeAPI.AttributeShape_WOP)
		attributeshapeAPI.AttributeShapePointersEncoding = attributeshapeDB.AttributeShapePointersEncoding
		attributeshapeAPIs = append(attributeshapeAPIs, attributeshapeAPI)
	}

	c.JSON(http.StatusOK, attributeshapeAPIs)
}

// PostAttributeShape
//
// swagger:route POST /attributeshapes attributeshapes postAttributeShape
//
// Creates a attributeshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAttributeShape(c *gin.Context) {

	mutexAttributeShape.Lock()
	defer mutexAttributeShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAttributeShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAttributeShape.GetDB()

	// Validate input
	var input orm.AttributeShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributeshape
	attributeshapeDB := orm.AttributeShapeDB{}
	attributeshapeDB.AttributeShapePointersEncoding = input.AttributeShapePointersEncoding
	attributeshapeDB.CopyBasicFieldsFromAttributeShape_WOP(&input.AttributeShape_WOP)

	_, err = db.Create(&attributeshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAttributeShape.CheckoutPhaseOneInstance(&attributeshapeDB)
	attributeshape := backRepo.BackRepoAttributeShape.Map_AttributeShapeDBID_AttributeShapePtr[attributeshapeDB.ID]

	if attributeshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributeshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributeshapeDB)
}

// GetAttributeShape
//
// swagger:route GET /attributeshapes/{ID} attributeshapes getAttributeShape
//
// Gets the details for a attributeshape.
//
// Responses:
// default: genericError
//
//	200: attributeshapeDBResponse
func (controller *Controller) GetAttributeShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributeShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAttributeShape.GetDB()

	// Get attributeshapeDB in DB
	var attributeshapeDB orm.AttributeShapeDB
	if _, err := db.First(&attributeshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributeshapeAPI orm.AttributeShapeAPI
	attributeshapeAPI.ID = attributeshapeDB.ID
	attributeshapeAPI.AttributeShapePointersEncoding = attributeshapeDB.AttributeShapePointersEncoding
	attributeshapeDB.CopyBasicFieldsToAttributeShape_WOP(&attributeshapeAPI.AttributeShape_WOP)

	c.JSON(http.StatusOK, attributeshapeAPI)
}

// UpdateAttributeShape
//
// swagger:route PATCH /attributeshapes/{ID} attributeshapes updateAttributeShape
//
// # Update a attributeshape
//
// Responses:
// default: genericError
//
//	200: attributeshapeDBResponse
func (controller *Controller) UpdateAttributeShape(c *gin.Context) {

	mutexAttributeShape.Lock()
	defer mutexAttributeShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAttributeShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAttributeShape.GetDB()

	// Validate input
	var input orm.AttributeShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributeshapeDB orm.AttributeShapeDB

	// fetch the attributeshape
	_, err := db.First(&attributeshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributeshapeDB.CopyBasicFieldsFromAttributeShape_WOP(&input.AttributeShape_WOP)
	attributeshapeDB.AttributeShapePointersEncoding = input.AttributeShapePointersEncoding

	db, _ = db.Model(&attributeshapeDB)
	_, err = db.Updates(&attributeshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributeshapeNew := new(models.AttributeShape)
	attributeshapeDB.CopyBasicFieldsToAttributeShape(attributeshapeNew)

	// redeem pointers
	attributeshapeDB.DecodePointers(backRepo, attributeshapeNew)

	// get stage instance from DB instance, and call callback function
	attributeshapeOld := backRepo.BackRepoAttributeShape.Map_AttributeShapeDBID_AttributeShapePtr[attributeshapeDB.ID]
	if attributeshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributeshapeOld, attributeshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributeshapeDB
	c.JSON(http.StatusOK, attributeshapeDB)
}

// DeleteAttributeShape
//
// swagger:route DELETE /attributeshapes/{ID} attributeshapes deleteAttributeShape
//
// # Delete a attributeshape
//
// default: genericError
//
//	200: attributeshapeDBResponse
func (controller *Controller) DeleteAttributeShape(c *gin.Context) {

	mutexAttributeShape.Lock()
	defer mutexAttributeShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAttributeShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAttributeShape.GetDB()

	// Get model if exist
	var attributeshapeDB orm.AttributeShapeDB
	if _, err := db.First(&attributeshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attributeshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributeshapeDeleted := new(models.AttributeShape)
	attributeshapeDB.CopyBasicFieldsToAttributeShape(attributeshapeDeleted)

	// get stage instance from DB instance, and call callback function
	attributeshapeStaged := backRepo.BackRepoAttributeShape.Map_AttributeShapeDBID_AttributeShapePtr[attributeshapeDB.ID]
	if attributeshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributeshapeStaged, attributeshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
