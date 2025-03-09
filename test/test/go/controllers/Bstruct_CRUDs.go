// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	"github.com/fullstack-lang/gong/test/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Bstruct__dummysDeclaration__ models.Bstruct
var __Bstruct_time__dummyDeclaration time.Duration

var mutexBstruct sync.Mutex

// An BstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBstruct updateBstruct deleteBstruct
type BstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBstruct updateBstruct
type BstructInput struct {
	// The Bstruct to submit or modify
	// in: body
	Bstruct *orm.BstructAPI
}

// GetBstructs
//
// swagger:route GET /bstructs bstructs getBstructs
//
// # Get all bstructs
//
// Responses:
// default: genericError
//
//	200: bstructDBResponse
func (controller *Controller) GetBstructs(c *gin.Context) {

	// source slice
	var bstructDBs []orm.BstructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	_, err := db.Find(&bstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bstructAPIs := make([]orm.BstructAPI, 0)

	// for each bstruct, update fields from the database nullable fields
	for idx := range bstructDBs {
		bstructDB := &bstructDBs[idx]
		_ = bstructDB
		var bstructAPI orm.BstructAPI

		// insertion point for updating fields
		bstructAPI.ID = bstructDB.ID
		bstructDB.CopyBasicFieldsToBstruct_WOP(&bstructAPI.Bstruct_WOP)
		bstructAPI.BstructPointersEncoding = bstructDB.BstructPointersEncoding
		bstructAPIs = append(bstructAPIs, bstructAPI)
	}

	c.JSON(http.StatusOK, bstructAPIs)
}

// PostBstruct
//
// swagger:route POST /bstructs bstructs postBstruct
//
// Creates a bstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBstruct(c *gin.Context) {

	mutexBstruct.Lock()
	defer mutexBstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBstructs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	// Validate input
	var input orm.BstructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bstruct
	bstructDB := orm.BstructDB{}
	bstructDB.BstructPointersEncoding = input.BstructPointersEncoding
	bstructDB.CopyBasicFieldsFromBstruct_WOP(&input.Bstruct_WOP)

	_, err = db.Create(&bstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBstruct.CheckoutPhaseOneInstance(&bstructDB)
	bstruct := backRepo.BackRepoBstruct.Map_BstructDBID_BstructPtr[bstructDB.ID]

	if bstruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bstructDB)
}

// GetBstruct
//
// swagger:route GET /bstructs/{ID} bstructs getBstruct
//
// Gets the details for a bstruct.
//
// Responses:
// default: genericError
//
//	200: bstructDBResponse
func (controller *Controller) GetBstruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	// Get bstructDB in DB
	var bstructDB orm.BstructDB
	if _, err := db.First(&bstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bstructAPI orm.BstructAPI
	bstructAPI.ID = bstructDB.ID
	bstructAPI.BstructPointersEncoding = bstructDB.BstructPointersEncoding
	bstructDB.CopyBasicFieldsToBstruct_WOP(&bstructAPI.Bstruct_WOP)

	c.JSON(http.StatusOK, bstructAPI)
}

// UpdateBstruct
//
// swagger:route PATCH /bstructs/{ID} bstructs updateBstruct
//
// # Update a bstruct
//
// Responses:
// default: genericError
//
//	200: bstructDBResponse
func (controller *Controller) UpdateBstruct(c *gin.Context) {

	mutexBstruct.Lock()
	defer mutexBstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	// Validate input
	var input orm.BstructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bstructDB orm.BstructDB

	// fetch the bstruct
	_, err := db.First(&bstructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bstructDB.CopyBasicFieldsFromBstruct_WOP(&input.Bstruct_WOP)
	bstructDB.BstructPointersEncoding = input.BstructPointersEncoding

	db, _ = db.Model(&bstructDB)
	_, err = db.Updates(&bstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bstructNew := new(models.Bstruct)
	bstructDB.CopyBasicFieldsToBstruct(bstructNew)

	// redeem pointers
	bstructDB.DecodePointers(backRepo, bstructNew)

	// get stage instance from DB instance, and call callback function
	bstructOld := backRepo.BackRepoBstruct.Map_BstructDBID_BstructPtr[bstructDB.ID]
	if bstructOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bstructOld, bstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bstructDB
	c.JSON(http.StatusOK, bstructDB)
}

// DeleteBstruct
//
// swagger:route DELETE /bstructs/{ID} bstructs deleteBstruct
//
// # Delete a bstruct
//
// default: genericError
//
//	200: bstructDBResponse
func (controller *Controller) DeleteBstruct(c *gin.Context) {

	mutexBstruct.Lock()
	defer mutexBstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBstruct", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoBstruct.GetDB()

	// Get model if exist
	var bstructDB orm.BstructDB
	if _, err := db.First(&bstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&bstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	bstructDeleted := new(models.Bstruct)
	bstructDB.CopyBasicFieldsToBstruct(bstructDeleted)

	// get stage instance from DB instance, and call callback function
	bstructStaged := backRepo.BackRepoBstruct.Map_BstructDBID_BstructPtr[bstructDB.ID]
	if bstructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bstructStaged, bstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
