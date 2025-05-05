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
var __GongNoteLinkShape__dummysDeclaration__ models.GongNoteLinkShape
var __GongNoteLinkShape_time__dummyDeclaration time.Duration

var mutexGongNoteLinkShape sync.Mutex

// An GongNoteLinkShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongNoteLinkShape updateGongNoteLinkShape deleteGongNoteLinkShape
type GongNoteLinkShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongNoteLinkShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongNoteLinkShape updateGongNoteLinkShape
type GongNoteLinkShapeInput struct {
	// The GongNoteLinkShape to submit or modify
	// in: body
	GongNoteLinkShape *orm.GongNoteLinkShapeAPI
}

// GetGongNoteLinkShapes
//
// swagger:route GET /gongnotelinkshapes gongnotelinkshapes getGongNoteLinkShapes
//
// # Get all gongnotelinkshapes
//
// Responses:
// default: genericError
//
//	200: gongnotelinkshapeDBResponse
func (controller *Controller) GetGongNoteLinkShapes(c *gin.Context) {

	// source slice
	var gongnotelinkshapeDBs []orm.GongNoteLinkShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNoteLinkShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteLinkShape.GetDB()

	_, err := db.Find(&gongnotelinkshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongnotelinkshapeAPIs := make([]orm.GongNoteLinkShapeAPI, 0)

	// for each gongnotelinkshape, update fields from the database nullable fields
	for idx := range gongnotelinkshapeDBs {
		gongnotelinkshapeDB := &gongnotelinkshapeDBs[idx]
		_ = gongnotelinkshapeDB
		var gongnotelinkshapeAPI orm.GongNoteLinkShapeAPI

		// insertion point for updating fields
		gongnotelinkshapeAPI.ID = gongnotelinkshapeDB.ID
		gongnotelinkshapeDB.CopyBasicFieldsToGongNoteLinkShape_WOP(&gongnotelinkshapeAPI.GongNoteLinkShape_WOP)
		gongnotelinkshapeAPI.GongNoteLinkShapePointersEncoding = gongnotelinkshapeDB.GongNoteLinkShapePointersEncoding
		gongnotelinkshapeAPIs = append(gongnotelinkshapeAPIs, gongnotelinkshapeAPI)
	}

	c.JSON(http.StatusOK, gongnotelinkshapeAPIs)
}

// PostGongNoteLinkShape
//
// swagger:route POST /gongnotelinkshapes gongnotelinkshapes postGongNoteLinkShape
//
// Creates a gongnotelinkshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongNoteLinkShape(c *gin.Context) {

	mutexGongNoteLinkShape.Lock()
	defer mutexGongNoteLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongNoteLinkShapes", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteLinkShape.GetDB()

	// Validate input
	var input orm.GongNoteLinkShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongnotelinkshape
	gongnotelinkshapeDB := orm.GongNoteLinkShapeDB{}
	gongnotelinkshapeDB.GongNoteLinkShapePointersEncoding = input.GongNoteLinkShapePointersEncoding
	gongnotelinkshapeDB.CopyBasicFieldsFromGongNoteLinkShape_WOP(&input.GongNoteLinkShape_WOP)

	_, err = db.Create(&gongnotelinkshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongNoteLinkShape.CheckoutPhaseOneInstance(&gongnotelinkshapeDB)
	gongnotelinkshape := backRepo.BackRepoGongNoteLinkShape.Map_GongNoteLinkShapeDBID_GongNoteLinkShapePtr[gongnotelinkshapeDB.ID]

	if gongnotelinkshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongnotelinkshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongnotelinkshapeDB)
}

// GetGongNoteLinkShape
//
// swagger:route GET /gongnotelinkshapes/{ID} gongnotelinkshapes getGongNoteLinkShape
//
// Gets the details for a gongnotelinkshape.
//
// Responses:
// default: genericError
//
//	200: gongnotelinkshapeDBResponse
func (controller *Controller) GetGongNoteLinkShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNoteLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteLinkShape.GetDB()

	// Get gongnotelinkshapeDB in DB
	var gongnotelinkshapeDB orm.GongNoteLinkShapeDB
	if _, err := db.First(&gongnotelinkshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongnotelinkshapeAPI orm.GongNoteLinkShapeAPI
	gongnotelinkshapeAPI.ID = gongnotelinkshapeDB.ID
	gongnotelinkshapeAPI.GongNoteLinkShapePointersEncoding = gongnotelinkshapeDB.GongNoteLinkShapePointersEncoding
	gongnotelinkshapeDB.CopyBasicFieldsToGongNoteLinkShape_WOP(&gongnotelinkshapeAPI.GongNoteLinkShape_WOP)

	c.JSON(http.StatusOK, gongnotelinkshapeAPI)
}

// UpdateGongNoteLinkShape
//
// swagger:route PATCH /gongnotelinkshapes/{ID} gongnotelinkshapes updateGongNoteLinkShape
//
// # Update a gongnotelinkshape
//
// Responses:
// default: genericError
//
//	200: gongnotelinkshapeDBResponse
func (controller *Controller) UpdateGongNoteLinkShape(c *gin.Context) {

	mutexGongNoteLinkShape.Lock()
	defer mutexGongNoteLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongNoteLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteLinkShape.GetDB()

	// Validate input
	var input orm.GongNoteLinkShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongnotelinkshapeDB orm.GongNoteLinkShapeDB

	// fetch the gongnotelinkshape
	_, err := db.First(&gongnotelinkshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongnotelinkshapeDB.CopyBasicFieldsFromGongNoteLinkShape_WOP(&input.GongNoteLinkShape_WOP)
	gongnotelinkshapeDB.GongNoteLinkShapePointersEncoding = input.GongNoteLinkShapePointersEncoding

	db, _ = db.Model(&gongnotelinkshapeDB)
	_, err = db.Updates(&gongnotelinkshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongnotelinkshapeNew := new(models.GongNoteLinkShape)
	gongnotelinkshapeDB.CopyBasicFieldsToGongNoteLinkShape(gongnotelinkshapeNew)

	// redeem pointers
	gongnotelinkshapeDB.DecodePointers(backRepo, gongnotelinkshapeNew)

	// get stage instance from DB instance, and call callback function
	gongnotelinkshapeOld := backRepo.BackRepoGongNoteLinkShape.Map_GongNoteLinkShapeDBID_GongNoteLinkShapePtr[gongnotelinkshapeDB.ID]
	if gongnotelinkshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongnotelinkshapeOld, gongnotelinkshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongnotelinkshapeDB
	c.JSON(http.StatusOK, gongnotelinkshapeDB)
}

// DeleteGongNoteLinkShape
//
// swagger:route DELETE /gongnotelinkshapes/{ID} gongnotelinkshapes deleteGongNoteLinkShape
//
// # Delete a gongnotelinkshape
//
// default: genericError
//
//	200: gongnotelinkshapeDBResponse
func (controller *Controller) DeleteGongNoteLinkShape(c *gin.Context) {

	mutexGongNoteLinkShape.Lock()
	defer mutexGongNoteLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongNoteLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoGongNoteLinkShape.GetDB()

	// Get model if exist
	var gongnotelinkshapeDB orm.GongNoteLinkShapeDB
	if _, err := db.First(&gongnotelinkshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongnotelinkshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongnotelinkshapeDeleted := new(models.GongNoteLinkShape)
	gongnotelinkshapeDB.CopyBasicFieldsToGongNoteLinkShape(gongnotelinkshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongnotelinkshapeStaged := backRepo.BackRepoGongNoteLinkShape.Map_GongNoteLinkShapeDBID_GongNoteLinkShapePtr[gongnotelinkshapeDB.ID]
	if gongnotelinkshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongnotelinkshapeStaged, gongnotelinkshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
