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
var __GongNoteShape__dummysDeclaration__ models.GongNoteShape
var __GongNoteShape_time__dummyDeclaration time.Duration

var mutexGongNoteShape sync.Mutex

// An GongNoteShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongNoteShape updateGongNoteShape deleteGongNoteShape
type GongNoteShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongNoteShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongNoteShape updateGongNoteShape
type GongNoteShapeInput struct {
	// The GongNoteShape to submit or modify
	// in: body
	GongNoteShape *orm.GongNoteShapeAPI
}

// GetGongNoteShapes
//
// swagger:route GET /gongnoteshapes gongnoteshapes getGongNoteShapes
//
// # Get all gongnoteshapes
//
// Responses:
// default: genericError
//
//	200: gongnoteshapeDBResponse
func (controller *Controller) GetGongNoteShapes(c *gin.Context) {

	// source slice
	var gongnoteshapeDBs []orm.GongNoteShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNoteShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteShape.GetDB()

	_, err := db.Find(&gongnoteshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongnoteshapeAPIs := make([]orm.GongNoteShapeAPI, 0)

	// for each gongnoteshape, update fields from the database nullable fields
	for idx := range gongnoteshapeDBs {
		gongnoteshapeDB := &gongnoteshapeDBs[idx]
		_ = gongnoteshapeDB
		var gongnoteshapeAPI orm.GongNoteShapeAPI

		// insertion point for updating fields
		gongnoteshapeAPI.ID = gongnoteshapeDB.ID
		gongnoteshapeDB.CopyBasicFieldsToGongNoteShape_WOP(&gongnoteshapeAPI.GongNoteShape_WOP)
		gongnoteshapeAPI.GongNoteShapePointersEncoding = gongnoteshapeDB.GongNoteShapePointersEncoding
		gongnoteshapeAPIs = append(gongnoteshapeAPIs, gongnoteshapeAPI)
	}

	c.JSON(http.StatusOK, gongnoteshapeAPIs)
}

// PostGongNoteShape
//
// swagger:route POST /gongnoteshapes gongnoteshapes postGongNoteShape
//
// Creates a gongnoteshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongNoteShape(c *gin.Context) {

	mutexGongNoteShape.Lock()
	defer mutexGongNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongNoteShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteShape.GetDB()

	// Validate input
	var input orm.GongNoteShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongnoteshape
	gongnoteshapeDB := orm.GongNoteShapeDB{}
	gongnoteshapeDB.GongNoteShapePointersEncoding = input.GongNoteShapePointersEncoding
	gongnoteshapeDB.CopyBasicFieldsFromGongNoteShape_WOP(&input.GongNoteShape_WOP)

	_, err = db.Create(&gongnoteshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongNoteShape.CheckoutPhaseOneInstance(&gongnoteshapeDB)
	gongnoteshape := backRepo.BackRepoGongNoteShape.Map_GongNoteShapeDBID_GongNoteShapePtr[gongnoteshapeDB.ID]

	if gongnoteshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongnoteshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongnoteshapeDB)
}

// GetGongNoteShape
//
// swagger:route GET /gongnoteshapes/{ID} gongnoteshapes getGongNoteShape
//
// Gets the details for a gongnoteshape.
//
// Responses:
// default: genericError
//
//	200: gongnoteshapeDBResponse
func (controller *Controller) GetGongNoteShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNoteShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteShape.GetDB()

	// Get gongnoteshapeDB in DB
	var gongnoteshapeDB orm.GongNoteShapeDB
	if _, err := db.First(&gongnoteshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongnoteshapeAPI orm.GongNoteShapeAPI
	gongnoteshapeAPI.ID = gongnoteshapeDB.ID
	gongnoteshapeAPI.GongNoteShapePointersEncoding = gongnoteshapeDB.GongNoteShapePointersEncoding
	gongnoteshapeDB.CopyBasicFieldsToGongNoteShape_WOP(&gongnoteshapeAPI.GongNoteShape_WOP)

	c.JSON(http.StatusOK, gongnoteshapeAPI)
}

// UpdateGongNoteShape
//
// swagger:route PATCH /gongnoteshapes/{ID} gongnoteshapes updateGongNoteShape
//
// # Update a gongnoteshape
//
// Responses:
// default: genericError
//
//	200: gongnoteshapeDBResponse
func (controller *Controller) UpdateGongNoteShape(c *gin.Context) {

	mutexGongNoteShape.Lock()
	defer mutexGongNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongNoteShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteShape.GetDB()

	// Validate input
	var input orm.GongNoteShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongnoteshapeDB orm.GongNoteShapeDB

	// fetch the gongnoteshape
	_, err := db.First(&gongnoteshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongnoteshapeDB.CopyBasicFieldsFromGongNoteShape_WOP(&input.GongNoteShape_WOP)
	gongnoteshapeDB.GongNoteShapePointersEncoding = input.GongNoteShapePointersEncoding

	db, _ = db.Model(&gongnoteshapeDB)
	_, err = db.Updates(&gongnoteshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongnoteshapeNew := new(models.GongNoteShape)
	gongnoteshapeDB.CopyBasicFieldsToGongNoteShape(gongnoteshapeNew)

	// redeem pointers
	gongnoteshapeDB.DecodePointers(backRepo, gongnoteshapeNew)

	// get stage instance from DB instance, and call callback function
	gongnoteshapeOld := backRepo.BackRepoGongNoteShape.Map_GongNoteShapeDBID_GongNoteShapePtr[gongnoteshapeDB.ID]
	if gongnoteshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongnoteshapeOld, gongnoteshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongnoteshapeDB
	c.JSON(http.StatusOK, gongnoteshapeDB)
}

// DeleteGongNoteShape
//
// swagger:route DELETE /gongnoteshapes/{ID} gongnoteshapes deleteGongNoteShape
//
// # Delete a gongnoteshape
//
// default: genericError
//
//	200: gongnoteshapeDBResponse
func (controller *Controller) DeleteGongNoteShape(c *gin.Context) {

	mutexGongNoteShape.Lock()
	defer mutexGongNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongNoteShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteShape.GetDB()

	// Get model if exist
	var gongnoteshapeDB orm.GongNoteShapeDB
	if _, err := db.First(&gongnoteshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongnoteshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongnoteshapeDeleted := new(models.GongNoteShape)
	gongnoteshapeDB.CopyBasicFieldsToGongNoteShape(gongnoteshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongnoteshapeStaged := backRepo.BackRepoGongNoteShape.Map_GongNoteShapeDBID_GongNoteShapePtr[gongnoteshapeDB.ID]
	if gongnoteshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongnoteshapeStaged, gongnoteshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
