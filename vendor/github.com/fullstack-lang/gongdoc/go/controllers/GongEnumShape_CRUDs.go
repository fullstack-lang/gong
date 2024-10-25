// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongEnumShape__dummysDeclaration__ models.GongEnumShape
var __GongEnumShape_time__dummyDeclaration time.Duration

var mutexGongEnumShape sync.Mutex

// An GongEnumShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnumShape updateGongEnumShape deleteGongEnumShape
type GongEnumShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnumShape updateGongEnumShape
type GongEnumShapeInput struct {
	// The GongEnumShape to submit or modify
	// in: body
	GongEnumShape *orm.GongEnumShapeAPI
}

// GetGongEnumShapes
//
// swagger:route GET /gongenumshapes gongenumshapes getGongEnumShapes
//
// # Get all gongenumshapes
//
// Responses:
// default: genericError
//
//	200: gongenumshapeDBResponse
func (controller *Controller) GetGongEnumShapes(c *gin.Context) {

	// source slice
	var gongenumshapeDBs []orm.GongEnumShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumShape.GetDB()

	_, err := db.Find(&gongenumshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumshapeAPIs := make([]orm.GongEnumShapeAPI, 0)

	// for each gongenumshape, update fields from the database nullable fields
	for idx := range gongenumshapeDBs {
		gongenumshapeDB := &gongenumshapeDBs[idx]
		_ = gongenumshapeDB
		var gongenumshapeAPI orm.GongEnumShapeAPI

		// insertion point for updating fields
		gongenumshapeAPI.ID = gongenumshapeDB.ID
		gongenumshapeDB.CopyBasicFieldsToGongEnumShape_WOP(&gongenumshapeAPI.GongEnumShape_WOP)
		gongenumshapeAPI.GongEnumShapePointersEncoding = gongenumshapeDB.GongEnumShapePointersEncoding
		gongenumshapeAPIs = append(gongenumshapeAPIs, gongenumshapeAPI)
	}

	c.JSON(http.StatusOK, gongenumshapeAPIs)
}

// PostGongEnumShape
//
// swagger:route POST /gongenumshapes gongenumshapes postGongEnumShape
//
// Creates a gongenumshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongEnumShape(c *gin.Context) {

	mutexGongEnumShape.Lock()
	defer mutexGongEnumShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongEnumShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumShape.GetDB()

	// Validate input
	var input orm.GongEnumShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenumshape
	gongenumshapeDB := orm.GongEnumShapeDB{}
	gongenumshapeDB.GongEnumShapePointersEncoding = input.GongEnumShapePointersEncoding
	gongenumshapeDB.CopyBasicFieldsFromGongEnumShape_WOP(&input.GongEnumShape_WOP)

	_, err = db.Create(&gongenumshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongEnumShape.CheckoutPhaseOneInstance(&gongenumshapeDB)
	gongenumshape := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]

	if gongenumshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongenumshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumshapeDB)
}

// GetGongEnumShape
//
// swagger:route GET /gongenumshapes/{ID} gongenumshapes getGongEnumShape
//
// Gets the details for a gongenumshape.
//
// Responses:
// default: genericError
//
//	200: gongenumshapeDBResponse
func (controller *Controller) GetGongEnumShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnumShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumShape.GetDB()

	// Get gongenumshapeDB in DB
	var gongenumshapeDB orm.GongEnumShapeDB
	if _, err := db.First(&gongenumshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumshapeAPI orm.GongEnumShapeAPI
	gongenumshapeAPI.ID = gongenumshapeDB.ID
	gongenumshapeAPI.GongEnumShapePointersEncoding = gongenumshapeDB.GongEnumShapePointersEncoding
	gongenumshapeDB.CopyBasicFieldsToGongEnumShape_WOP(&gongenumshapeAPI.GongEnumShape_WOP)

	c.JSON(http.StatusOK, gongenumshapeAPI)
}

// UpdateGongEnumShape
//
// swagger:route PATCH /gongenumshapes/{ID} gongenumshapes updateGongEnumShape
//
// # Update a gongenumshape
//
// Responses:
// default: genericError
//
//	200: gongenumshapeDBResponse
func (controller *Controller) UpdateGongEnumShape(c *gin.Context) {

	mutexGongEnumShape.Lock()
	defer mutexGongEnumShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongEnumShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumShape.GetDB()

	// Validate input
	var input orm.GongEnumShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongenumshapeDB orm.GongEnumShapeDB

	// fetch the gongenumshape
	_, err := db.First(&gongenumshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumshapeDB.CopyBasicFieldsFromGongEnumShape_WOP(&input.GongEnumShape_WOP)
	gongenumshapeDB.GongEnumShapePointersEncoding = input.GongEnumShapePointersEncoding

	db, _ = db.Model(&gongenumshapeDB)
	_, err = db.Updates(&gongenumshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumshapeNew := new(models.GongEnumShape)
	gongenumshapeDB.CopyBasicFieldsToGongEnumShape(gongenumshapeNew)

	// redeem pointers
	gongenumshapeDB.DecodePointers(backRepo, gongenumshapeNew)

	// get stage instance from DB instance, and call callback function
	gongenumshapeOld := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]
	if gongenumshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongenumshapeOld, gongenumshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumshapeDB
	c.JSON(http.StatusOK, gongenumshapeDB)
}

// DeleteGongEnumShape
//
// swagger:route DELETE /gongenumshapes/{ID} gongenumshapes deleteGongEnumShape
//
// # Delete a gongenumshape
//
// default: genericError
//
//	200: gongenumshapeDBResponse
func (controller *Controller) DeleteGongEnumShape(c *gin.Context) {

	mutexGongEnumShape.Lock()
	defer mutexGongEnumShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongEnumShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongEnumShape.GetDB()

	// Get model if exist
	var gongenumshapeDB orm.GongEnumShapeDB
	if _, err := db.First(&gongenumshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongenumshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumshapeDeleted := new(models.GongEnumShape)
	gongenumshapeDB.CopyBasicFieldsToGongEnumShape(gongenumshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumshapeStaged := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]
	if gongenumshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongenumshapeStaged, gongenumshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
