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
var __FormSortAssocButton__dummysDeclaration__ models.FormSortAssocButton
var __FormSortAssocButton_time__dummyDeclaration time.Duration

var mutexFormSortAssocButton sync.Mutex

// An FormSortAssocButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormSortAssocButton updateFormSortAssocButton deleteFormSortAssocButton
type FormSortAssocButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormSortAssocButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormSortAssocButton updateFormSortAssocButton
type FormSortAssocButtonInput struct {
	// The FormSortAssocButton to submit or modify
	// in: body
	FormSortAssocButton *orm.FormSortAssocButtonAPI
}

// GetFormSortAssocButtons
//
// swagger:route GET /formsortassocbuttons formsortassocbuttons getFormSortAssocButtons
//
// # Get all formsortassocbuttons
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) GetFormSortAssocButtons(c *gin.Context) {

	// source slice
	var formsortassocbuttonDBs []orm.FormSortAssocButtonDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormSortAssocButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	_, err := db.Find(&formsortassocbuttonDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formsortassocbuttonAPIs := make([]orm.FormSortAssocButtonAPI, 0)

	// for each formsortassocbutton, update fields from the database nullable fields
	for idx := range formsortassocbuttonDBs {
		formsortassocbuttonDB := &formsortassocbuttonDBs[idx]
		_ = formsortassocbuttonDB
		var formsortassocbuttonAPI orm.FormSortAssocButtonAPI

		// insertion point for updating fields
		formsortassocbuttonAPI.ID = formsortassocbuttonDB.ID
		formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton_WOP(&formsortassocbuttonAPI.FormSortAssocButton_WOP)
		formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding = formsortassocbuttonDB.FormSortAssocButtonPointersEncoding
		formsortassocbuttonAPIs = append(formsortassocbuttonAPIs, formsortassocbuttonAPI)
	}

	c.JSON(http.StatusOK, formsortassocbuttonAPIs)
}

// PostFormSortAssocButton
//
// swagger:route POST /formsortassocbuttons formsortassocbuttons postFormSortAssocButton
//
// Creates a formsortassocbutton
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormSortAssocButton(c *gin.Context) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormSortAssocButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Validate input
	var input orm.FormSortAssocButtonAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formsortassocbutton
	formsortassocbuttonDB := orm.FormSortAssocButtonDB{}
	formsortassocbuttonDB.FormSortAssocButtonPointersEncoding = input.FormSortAssocButtonPointersEncoding
	formsortassocbuttonDB.CopyBasicFieldsFromFormSortAssocButton_WOP(&input.FormSortAssocButton_WOP)

	_, err = db.Create(&formsortassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormSortAssocButton.CheckoutPhaseOneInstance(&formsortassocbuttonDB)
	formsortassocbutton := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]

	if formsortassocbutton != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formsortassocbutton)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formsortassocbuttonDB)
}

// GetFormSortAssocButton
//
// swagger:route GET /formsortassocbuttons/{ID} formsortassocbuttons getFormSortAssocButton
//
// Gets the details for a formsortassocbutton.
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) GetFormSortAssocButton(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormSortAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Get formsortassocbuttonDB in DB
	var formsortassocbuttonDB orm.FormSortAssocButtonDB
	if _, err := db.First(&formsortassocbuttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formsortassocbuttonAPI orm.FormSortAssocButtonAPI
	formsortassocbuttonAPI.ID = formsortassocbuttonDB.ID
	formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding = formsortassocbuttonDB.FormSortAssocButtonPointersEncoding
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton_WOP(&formsortassocbuttonAPI.FormSortAssocButton_WOP)

	c.JSON(http.StatusOK, formsortassocbuttonAPI)
}

// UpdateFormSortAssocButton
//
// swagger:route PATCH /formsortassocbuttons/{ID} formsortassocbuttons updateFormSortAssocButton
//
// # Update a formsortassocbutton
//
// Responses:
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) UpdateFormSortAssocButton(c *gin.Context) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormSortAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Validate input
	var input orm.FormSortAssocButtonAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formsortassocbuttonDB orm.FormSortAssocButtonDB

	// fetch the formsortassocbutton
	_, err := db.First(&formsortassocbuttonDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formsortassocbuttonDB.CopyBasicFieldsFromFormSortAssocButton_WOP(&input.FormSortAssocButton_WOP)
	formsortassocbuttonDB.FormSortAssocButtonPointersEncoding = input.FormSortAssocButtonPointersEncoding

	db, _ = db.Model(&formsortassocbuttonDB)
	_, err = db.Updates(&formsortassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formsortassocbuttonNew := new(models.FormSortAssocButton)
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton(formsortassocbuttonNew)

	// redeem pointers
	formsortassocbuttonDB.DecodePointers(backRepo, formsortassocbuttonNew)

	// get stage instance from DB instance, and call callback function
	formsortassocbuttonOld := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]
	if formsortassocbuttonOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formsortassocbuttonOld, formsortassocbuttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formsortassocbuttonDB
	c.JSON(http.StatusOK, formsortassocbuttonDB)
}

// DeleteFormSortAssocButton
//
// swagger:route DELETE /formsortassocbuttons/{ID} formsortassocbuttons deleteFormSortAssocButton
//
// # Delete a formsortassocbutton
//
// default: genericError
//
//	200: formsortassocbuttonDBResponse
func (controller *Controller) DeleteFormSortAssocButton(c *gin.Context) {

	mutexFormSortAssocButton.Lock()
	defer mutexFormSortAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormSortAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormSortAssocButton.GetDB()

	// Get model if exist
	var formsortassocbuttonDB orm.FormSortAssocButtonDB
	if _, err := db.First(&formsortassocbuttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formsortassocbuttonDB)

	// get an instance (not staged) from DB instance, and call callback function
	formsortassocbuttonDeleted := new(models.FormSortAssocButton)
	formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton(formsortassocbuttonDeleted)

	// get stage instance from DB instance, and call callback function
	formsortassocbuttonStaged := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[formsortassocbuttonDB.ID]
	if formsortassocbuttonStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formsortassocbuttonStaged, formsortassocbuttonDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
