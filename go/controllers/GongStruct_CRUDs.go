// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongStruct__dummysDeclaration__ models.GongStruct
var __GongStruct_time__dummyDeclaration time.Duration

var mutexGongStruct sync.Mutex

// An GongStructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongStruct updateGongStruct deleteGongStruct
type GongStructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongStructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongStruct updateGongStruct
type GongStructInput struct {
	// The GongStruct to submit or modify
	// in: body
	GongStruct *orm.GongStructAPI
}

// GetGongStructs
//
// swagger:route GET /gongstructs gongstructs getGongStructs
//
// # Get all gongstructs
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func (controller *Controller) GetGongStructs(c *gin.Context) {

	// source slice
	var gongstructDBs []orm.GongStructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongStructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongStruct.GetDB()

	_, err := db.Find(&gongstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongstructAPIs := make([]orm.GongStructAPI, 0)

	// for each gongstruct, update fields from the database nullable fields
	for idx := range gongstructDBs {
		gongstructDB := &gongstructDBs[idx]
		_ = gongstructDB
		var gongstructAPI orm.GongStructAPI

		// insertion point for updating fields
		gongstructAPI.ID = gongstructDB.ID
		gongstructDB.CopyBasicFieldsToGongStruct_WOP(&gongstructAPI.GongStruct_WOP)
		gongstructAPI.GongStructPointersEncoding = gongstructDB.GongStructPointersEncoding
		gongstructAPIs = append(gongstructAPIs, gongstructAPI)
	}

	c.JSON(http.StatusOK, gongstructAPIs)
}

// PostGongStruct
//
// swagger:route POST /gongstructs gongstructs postGongStruct
//
// Creates a gongstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongStruct(c *gin.Context) {

	mutexGongStruct.Lock()
	defer mutexGongStruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongStructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongStruct.GetDB()

	// Validate input
	var input orm.GongStructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongstruct
	gongstructDB := orm.GongStructDB{}
	gongstructDB.GongStructPointersEncoding = input.GongStructPointersEncoding
	gongstructDB.CopyBasicFieldsFromGongStruct_WOP(&input.GongStruct_WOP)

	_, err = db.Create(&gongstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongStruct.CheckoutPhaseOneInstance(&gongstructDB)
	gongstruct := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]

	if gongstruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongstructDB)
}

// GetGongStruct
//
// swagger:route GET /gongstructs/{ID} gongstructs getGongStruct
//
// Gets the details for a gongstruct.
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func (controller *Controller) GetGongStruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongStruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongStruct.GetDB()

	// Get gongstructDB in DB
	var gongstructDB orm.GongStructDB
	if _, err := db.First(&gongstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongstructAPI orm.GongStructAPI
	gongstructAPI.ID = gongstructDB.ID
	gongstructAPI.GongStructPointersEncoding = gongstructDB.GongStructPointersEncoding
	gongstructDB.CopyBasicFieldsToGongStruct_WOP(&gongstructAPI.GongStruct_WOP)

	c.JSON(http.StatusOK, gongstructAPI)
}

// UpdateGongStruct
//
// swagger:route PATCH /gongstructs/{ID} gongstructs updateGongStruct
//
// # Update a gongstruct
//
// Responses:
// default: genericError
//
//	200: gongstructDBResponse
func (controller *Controller) UpdateGongStruct(c *gin.Context) {

	mutexGongStruct.Lock()
	defer mutexGongStruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongStruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongStruct.GetDB()

	// Validate input
	var input orm.GongStructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongstructDB orm.GongStructDB

	// fetch the gongstruct
	_, err := db.First(&gongstructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongstructDB.CopyBasicFieldsFromGongStruct_WOP(&input.GongStruct_WOP)
	gongstructDB.GongStructPointersEncoding = input.GongStructPointersEncoding

	db, _ = db.Model(&gongstructDB)
	_, err = db.Updates(&gongstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongstructNew := new(models.GongStruct)
	gongstructDB.CopyBasicFieldsToGongStruct(gongstructNew)

	// redeem pointers
	gongstructDB.DecodePointers(backRepo, gongstructNew)

	// get stage instance from DB instance, and call callback function
	gongstructOld := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]
	if gongstructOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongstructOld, gongstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongstructDB
	c.JSON(http.StatusOK, gongstructDB)
}

// DeleteGongStruct
//
// swagger:route DELETE /gongstructs/{ID} gongstructs deleteGongStruct
//
// # Delete a gongstruct
//
// default: genericError
//
//	200: gongstructDBResponse
func (controller *Controller) DeleteGongStruct(c *gin.Context) {

	mutexGongStruct.Lock()
	defer mutexGongStruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongStruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongStruct.GetDB()

	// Get model if exist
	var gongstructDB orm.GongStructDB
	if _, err := db.First(&gongstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongstructDeleted := new(models.GongStruct)
	gongstructDB.CopyBasicFieldsToGongStruct(gongstructDeleted)

	// get stage instance from DB instance, and call callback function
	gongstructStaged := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]
	if gongstructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongstructStaged, gongstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
