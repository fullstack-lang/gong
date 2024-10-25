// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FormDiv__dummysDeclaration__ models.FormDiv
var __FormDiv_time__dummyDeclaration time.Duration

var mutexFormDiv sync.Mutex

// An FormDivID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormDiv updateFormDiv deleteFormDiv
type FormDivID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormDivInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormDiv updateFormDiv
type FormDivInput struct {
	// The FormDiv to submit or modify
	// in: body
	FormDiv *orm.FormDivAPI
}

// GetFormDivs
//
// swagger:route GET /formdivs formdivs getFormDivs
//
// # Get all formdivs
//
// Responses:
// default: genericError
//
//	200: formdivDBResponse
func (controller *Controller) GetFormDivs(c *gin.Context) {

	// source slice
	var formdivDBs []orm.FormDivDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormDivs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	_, err := db.Find(&formdivDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formdivAPIs := make([]orm.FormDivAPI, 0)

	// for each formdiv, update fields from the database nullable fields
	for idx := range formdivDBs {
		formdivDB := &formdivDBs[idx]
		_ = formdivDB
		var formdivAPI orm.FormDivAPI

		// insertion point for updating fields
		formdivAPI.ID = formdivDB.ID
		formdivDB.CopyBasicFieldsToFormDiv_WOP(&formdivAPI.FormDiv_WOP)
		formdivAPI.FormDivPointersEncoding = formdivDB.FormDivPointersEncoding
		formdivAPIs = append(formdivAPIs, formdivAPI)
	}

	c.JSON(http.StatusOK, formdivAPIs)
}

// PostFormDiv
//
// swagger:route POST /formdivs formdivs postFormDiv
//
// Creates a formdiv
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormDiv(c *gin.Context) {

	mutexFormDiv.Lock()
	defer mutexFormDiv.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormDivs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	// Validate input
	var input orm.FormDivAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formdiv
	formdivDB := orm.FormDivDB{}
	formdivDB.FormDivPointersEncoding = input.FormDivPointersEncoding
	formdivDB.CopyBasicFieldsFromFormDiv_WOP(&input.FormDiv_WOP)

	_, err = db.Create(&formdivDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormDiv.CheckoutPhaseOneInstance(&formdivDB)
	formdiv := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]

	if formdiv != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formdiv)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formdivDB)
}

// GetFormDiv
//
// swagger:route GET /formdivs/{ID} formdivs getFormDiv
//
// Gets the details for a formdiv.
//
// Responses:
// default: genericError
//
//	200: formdivDBResponse
func (controller *Controller) GetFormDiv(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormDiv", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	// Get formdivDB in DB
	var formdivDB orm.FormDivDB
	if _, err := db.First(&formdivDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formdivAPI orm.FormDivAPI
	formdivAPI.ID = formdivDB.ID
	formdivAPI.FormDivPointersEncoding = formdivDB.FormDivPointersEncoding
	formdivDB.CopyBasicFieldsToFormDiv_WOP(&formdivAPI.FormDiv_WOP)

	c.JSON(http.StatusOK, formdivAPI)
}

// UpdateFormDiv
//
// swagger:route PATCH /formdivs/{ID} formdivs updateFormDiv
//
// # Update a formdiv
//
// Responses:
// default: genericError
//
//	200: formdivDBResponse
func (controller *Controller) UpdateFormDiv(c *gin.Context) {

	mutexFormDiv.Lock()
	defer mutexFormDiv.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormDiv", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	// Validate input
	var input orm.FormDivAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formdivDB orm.FormDivDB

	// fetch the formdiv
	_, err := db.First(&formdivDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formdivDB.CopyBasicFieldsFromFormDiv_WOP(&input.FormDiv_WOP)
	formdivDB.FormDivPointersEncoding = input.FormDivPointersEncoding

	db, _ = db.Model(&formdivDB)
	_, err = db.Updates(&formdivDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formdivNew := new(models.FormDiv)
	formdivDB.CopyBasicFieldsToFormDiv(formdivNew)

	// redeem pointers
	formdivDB.DecodePointers(backRepo, formdivNew)

	// get stage instance from DB instance, and call callback function
	formdivOld := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
	if formdivOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formdivOld, formdivNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formdivDB
	c.JSON(http.StatusOK, formdivDB)
}

// DeleteFormDiv
//
// swagger:route DELETE /formdivs/{ID} formdivs deleteFormDiv
//
// # Delete a formdiv
//
// default: genericError
//
//	200: formdivDBResponse
func (controller *Controller) DeleteFormDiv(c *gin.Context) {

	mutexFormDiv.Lock()
	defer mutexFormDiv.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormDiv", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormDiv.GetDB()

	// Get model if exist
	var formdivDB orm.FormDivDB
	if _, err := db.First(&formdivDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formdivDB)

	// get an instance (not staged) from DB instance, and call callback function
	formdivDeleted := new(models.FormDiv)
	formdivDB.CopyBasicFieldsToFormDiv(formdivDeleted)

	// get stage instance from DB instance, and call callback function
	formdivStaged := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
	if formdivStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formdivStaged, formdivDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
