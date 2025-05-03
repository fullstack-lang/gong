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
var __GongEnumValueShape__dummysDeclaration__ models.GongEnumValueShape
var __GongEnumValueShape_time__dummyDeclaration time.Duration

var mutexGongEnumValueShape sync.Mutex

// An GongEnumValueShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnumValueShape updateGongEnumValueShape deleteGongEnumValueShape
type GongEnumValueShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumValueShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnumValueShape updateGongEnumValueShape
type GongEnumValueShapeInput struct {
	// The GongEnumValueShape to submit or modify
	// in: body
	GongEnumValueShape *orm.GongEnumValueShapeAPI
}

// GetGongEnumValueShapes
//
// swagger:route GET /gongenumvalueshapes gongenumvalueshapes getGongEnumValueShapes
//
// # Get all gongenumvalueshapes
//
// Responses:
// default: genericError
//
//	200: gongenumvalueshapeDBResponse
func (controller *Controller) GetGongEnumValueShapes(c *gin.Context) {

	// source slice
	var gongenumvalueshapeDBs []orm.GongEnumValueShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValueShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongEnumValueShape.GetDB()

	_, err := db.Find(&gongenumvalueshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumvalueshapeAPIs := make([]orm.GongEnumValueShapeAPI, 0)

	// for each gongenumvalueshape, update fields from the database nullable fields
	for idx := range gongenumvalueshapeDBs {
		gongenumvalueshapeDB := &gongenumvalueshapeDBs[idx]
		_ = gongenumvalueshapeDB
		var gongenumvalueshapeAPI orm.GongEnumValueShapeAPI

		// insertion point for updating fields
		gongenumvalueshapeAPI.ID = gongenumvalueshapeDB.ID
		gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape_WOP(&gongenumvalueshapeAPI.GongEnumValueShape_WOP)
		gongenumvalueshapeAPI.GongEnumValueShapePointersEncoding = gongenumvalueshapeDB.GongEnumValueShapePointersEncoding
		gongenumvalueshapeAPIs = append(gongenumvalueshapeAPIs, gongenumvalueshapeAPI)
	}

	c.JSON(http.StatusOK, gongenumvalueshapeAPIs)
}

// PostGongEnumValueShape
//
// swagger:route POST /gongenumvalueshapes gongenumvalueshapes postGongEnumValueShape
//
// Creates a gongenumvalueshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongEnumValueShape(c *gin.Context) {

	mutexGongEnumValueShape.Lock()
	defer mutexGongEnumValueShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongEnumValueShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongEnumValueShape.GetDB()

	// Validate input
	var input orm.GongEnumValueShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenumvalueshape
	gongenumvalueshapeDB := orm.GongEnumValueShapeDB{}
	gongenumvalueshapeDB.GongEnumValueShapePointersEncoding = input.GongEnumValueShapePointersEncoding
	gongenumvalueshapeDB.CopyBasicFieldsFromGongEnumValueShape_WOP(&input.GongEnumValueShape_WOP)

	_, err = db.Create(&gongenumvalueshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongEnumValueShape.CheckoutPhaseOneInstance(&gongenumvalueshapeDB)
	gongenumvalueshape := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]

	if gongenumvalueshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongenumvalueshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumvalueshapeDB)
}

// GetGongEnumValueShape
//
// swagger:route GET /gongenumvalueshapes/{ID} gongenumvalueshapes getGongEnumValueShape
//
// Gets the details for a gongenumvalueshape.
//
// Responses:
// default: genericError
//
//	200: gongenumvalueshapeDBResponse
func (controller *Controller) GetGongEnumValueShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumValueShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongEnumValueShape.GetDB()

	// Get gongenumvalueshapeDB in DB
	var gongenumvalueshapeDB orm.GongEnumValueShapeDB
	if _, err := db.First(&gongenumvalueshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumvalueshapeAPI orm.GongEnumValueShapeAPI
	gongenumvalueshapeAPI.ID = gongenumvalueshapeDB.ID
	gongenumvalueshapeAPI.GongEnumValueShapePointersEncoding = gongenumvalueshapeDB.GongEnumValueShapePointersEncoding
	gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape_WOP(&gongenumvalueshapeAPI.GongEnumValueShape_WOP)

	c.JSON(http.StatusOK, gongenumvalueshapeAPI)
}

// UpdateGongEnumValueShape
//
// swagger:route PATCH /gongenumvalueshapes/{ID} gongenumvalueshapes updateGongEnumValueShape
//
// # Update a gongenumvalueshape
//
// Responses:
// default: genericError
//
//	200: gongenumvalueshapeDBResponse
func (controller *Controller) UpdateGongEnumValueShape(c *gin.Context) {

	mutexGongEnumValueShape.Lock()
	defer mutexGongEnumValueShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongEnumValueShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongEnumValueShape.GetDB()

	// Validate input
	var input orm.GongEnumValueShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongenumvalueshapeDB orm.GongEnumValueShapeDB

	// fetch the gongenumvalueshape
	_, err := db.First(&gongenumvalueshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumvalueshapeDB.CopyBasicFieldsFromGongEnumValueShape_WOP(&input.GongEnumValueShape_WOP)
	gongenumvalueshapeDB.GongEnumValueShapePointersEncoding = input.GongEnumValueShapePointersEncoding

	db, _ = db.Model(&gongenumvalueshapeDB)
	_, err = db.Updates(&gongenumvalueshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueshapeNew := new(models.GongEnumValueShape)
	gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape(gongenumvalueshapeNew)

	// redeem pointers
	gongenumvalueshapeDB.DecodePointers(backRepo, gongenumvalueshapeNew)

	// get stage instance from DB instance, and call callback function
	gongenumvalueshapeOld := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]
	if gongenumvalueshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongenumvalueshapeOld, gongenumvalueshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumvalueshapeDB
	c.JSON(http.StatusOK, gongenumvalueshapeDB)
}

// DeleteGongEnumValueShape
//
// swagger:route DELETE /gongenumvalueshapes/{ID} gongenumvalueshapes deleteGongEnumValueShape
//
// # Delete a gongenumvalueshape
//
// default: genericError
//
//	200: gongenumvalueshapeDBResponse
func (controller *Controller) DeleteGongEnumValueShape(c *gin.Context) {

	mutexGongEnumValueShape.Lock()
	defer mutexGongEnumValueShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongEnumValueShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongEnumValueShape.GetDB()

	// Get model if exist
	var gongenumvalueshapeDB orm.GongEnumValueShapeDB
	if _, err := db.First(&gongenumvalueshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongenumvalueshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumvalueshapeDeleted := new(models.GongEnumValueShape)
	gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape(gongenumvalueshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumvalueshapeStaged := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]
	if gongenumvalueshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongenumvalueshapeStaged, gongenumvalueshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
