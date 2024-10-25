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
var __ModelPkg__dummysDeclaration__ models.ModelPkg
var __ModelPkg_time__dummyDeclaration time.Duration

var mutexModelPkg sync.Mutex

// An ModelPkgID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getModelPkg updateModelPkg deleteModelPkg
type ModelPkgID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ModelPkgInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postModelPkg updateModelPkg
type ModelPkgInput struct {
	// The ModelPkg to submit or modify
	// in: body
	ModelPkg *orm.ModelPkgAPI
}

// GetModelPkgs
//
// swagger:route GET /modelpkgs modelpkgs getModelPkgs
//
// # Get all modelpkgs
//
// Responses:
// default: genericError
//
//	200: modelpkgDBResponse
func (controller *Controller) GetModelPkgs(c *gin.Context) {

	// source slice
	var modelpkgDBs []orm.ModelPkgDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetModelPkgs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelPkg.GetDB()

	_, err := db.Find(&modelpkgDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	modelpkgAPIs := make([]orm.ModelPkgAPI, 0)

	// for each modelpkg, update fields from the database nullable fields
	for idx := range modelpkgDBs {
		modelpkgDB := &modelpkgDBs[idx]
		_ = modelpkgDB
		var modelpkgAPI orm.ModelPkgAPI

		// insertion point for updating fields
		modelpkgAPI.ID = modelpkgDB.ID
		modelpkgDB.CopyBasicFieldsToModelPkg_WOP(&modelpkgAPI.ModelPkg_WOP)
		modelpkgAPI.ModelPkgPointersEncoding = modelpkgDB.ModelPkgPointersEncoding
		modelpkgAPIs = append(modelpkgAPIs, modelpkgAPI)
	}

	c.JSON(http.StatusOK, modelpkgAPIs)
}

// PostModelPkg
//
// swagger:route POST /modelpkgs modelpkgs postModelPkg
//
// Creates a modelpkg
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostModelPkg(c *gin.Context) {

	mutexModelPkg.Lock()
	defer mutexModelPkg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostModelPkgs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelPkg.GetDB()

	// Validate input
	var input orm.ModelPkgAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create modelpkg
	modelpkgDB := orm.ModelPkgDB{}
	modelpkgDB.ModelPkgPointersEncoding = input.ModelPkgPointersEncoding
	modelpkgDB.CopyBasicFieldsFromModelPkg_WOP(&input.ModelPkg_WOP)

	_, err = db.Create(&modelpkgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoModelPkg.CheckoutPhaseOneInstance(&modelpkgDB)
	modelpkg := backRepo.BackRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr[modelpkgDB.ID]

	if modelpkg != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), modelpkg)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, modelpkgDB)
}

// GetModelPkg
//
// swagger:route GET /modelpkgs/{ID} modelpkgs getModelPkg
//
// Gets the details for a modelpkg.
//
// Responses:
// default: genericError
//
//	200: modelpkgDBResponse
func (controller *Controller) GetModelPkg(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetModelPkg", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelPkg.GetDB()

	// Get modelpkgDB in DB
	var modelpkgDB orm.ModelPkgDB
	if _, err := db.First(&modelpkgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var modelpkgAPI orm.ModelPkgAPI
	modelpkgAPI.ID = modelpkgDB.ID
	modelpkgAPI.ModelPkgPointersEncoding = modelpkgDB.ModelPkgPointersEncoding
	modelpkgDB.CopyBasicFieldsToModelPkg_WOP(&modelpkgAPI.ModelPkg_WOP)

	c.JSON(http.StatusOK, modelpkgAPI)
}

// UpdateModelPkg
//
// swagger:route PATCH /modelpkgs/{ID} modelpkgs updateModelPkg
//
// # Update a modelpkg
//
// Responses:
// default: genericError
//
//	200: modelpkgDBResponse
func (controller *Controller) UpdateModelPkg(c *gin.Context) {

	mutexModelPkg.Lock()
	defer mutexModelPkg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateModelPkg", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelPkg.GetDB()

	// Validate input
	var input orm.ModelPkgAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var modelpkgDB orm.ModelPkgDB

	// fetch the modelpkg
	_, err := db.First(&modelpkgDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	modelpkgDB.CopyBasicFieldsFromModelPkg_WOP(&input.ModelPkg_WOP)
	modelpkgDB.ModelPkgPointersEncoding = input.ModelPkgPointersEncoding

	db, _ = db.Model(&modelpkgDB)
	_, err = db.Updates(&modelpkgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	modelpkgNew := new(models.ModelPkg)
	modelpkgDB.CopyBasicFieldsToModelPkg(modelpkgNew)

	// redeem pointers
	modelpkgDB.DecodePointers(backRepo, modelpkgNew)

	// get stage instance from DB instance, and call callback function
	modelpkgOld := backRepo.BackRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr[modelpkgDB.ID]
	if modelpkgOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), modelpkgOld, modelpkgNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the modelpkgDB
	c.JSON(http.StatusOK, modelpkgDB)
}

// DeleteModelPkg
//
// swagger:route DELETE /modelpkgs/{ID} modelpkgs deleteModelPkg
//
// # Delete a modelpkg
//
// default: genericError
//
//	200: modelpkgDBResponse
func (controller *Controller) DeleteModelPkg(c *gin.Context) {

	mutexModelPkg.Lock()
	defer mutexModelPkg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteModelPkg", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelPkg.GetDB()

	// Get model if exist
	var modelpkgDB orm.ModelPkgDB
	if _, err := db.First(&modelpkgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&modelpkgDB)

	// get an instance (not staged) from DB instance, and call callback function
	modelpkgDeleted := new(models.ModelPkg)
	modelpkgDB.CopyBasicFieldsToModelPkg(modelpkgDeleted)

	// get stage instance from DB instance, and call callback function
	modelpkgStaged := backRepo.BackRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr[modelpkgDB.ID]
	if modelpkgStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), modelpkgStaged, modelpkgDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
