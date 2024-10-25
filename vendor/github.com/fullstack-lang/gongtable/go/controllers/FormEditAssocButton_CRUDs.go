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
var __FormEditAssocButton__dummysDeclaration__ models.FormEditAssocButton
var __FormEditAssocButton_time__dummyDeclaration time.Duration

var mutexFormEditAssocButton sync.Mutex

// An FormEditAssocButtonID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormEditAssocButton updateFormEditAssocButton deleteFormEditAssocButton
type FormEditAssocButtonID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormEditAssocButtonInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormEditAssocButton updateFormEditAssocButton
type FormEditAssocButtonInput struct {
	// The FormEditAssocButton to submit or modify
	// in: body
	FormEditAssocButton *orm.FormEditAssocButtonAPI
}

// GetFormEditAssocButtons
//
// swagger:route GET /formeditassocbuttons formeditassocbuttons getFormEditAssocButtons
//
// # Get all formeditassocbuttons
//
// Responses:
// default: genericError
//
//	200: formeditassocbuttonDBResponse
func (controller *Controller) GetFormEditAssocButtons(c *gin.Context) {

	// source slice
	var formeditassocbuttonDBs []orm.FormEditAssocButtonDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormEditAssocButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	_, err := db.Find(&formeditassocbuttonDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formeditassocbuttonAPIs := make([]orm.FormEditAssocButtonAPI, 0)

	// for each formeditassocbutton, update fields from the database nullable fields
	for idx := range formeditassocbuttonDBs {
		formeditassocbuttonDB := &formeditassocbuttonDBs[idx]
		_ = formeditassocbuttonDB
		var formeditassocbuttonAPI orm.FormEditAssocButtonAPI

		// insertion point for updating fields
		formeditassocbuttonAPI.ID = formeditassocbuttonDB.ID
		formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton_WOP(&formeditassocbuttonAPI.FormEditAssocButton_WOP)
		formeditassocbuttonAPI.FormEditAssocButtonPointersEncoding = formeditassocbuttonDB.FormEditAssocButtonPointersEncoding
		formeditassocbuttonAPIs = append(formeditassocbuttonAPIs, formeditassocbuttonAPI)
	}

	c.JSON(http.StatusOK, formeditassocbuttonAPIs)
}

// PostFormEditAssocButton
//
// swagger:route POST /formeditassocbuttons formeditassocbuttons postFormEditAssocButton
//
// Creates a formeditassocbutton
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormEditAssocButton(c *gin.Context) {

	mutexFormEditAssocButton.Lock()
	defer mutexFormEditAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormEditAssocButtons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	// Validate input
	var input orm.FormEditAssocButtonAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formeditassocbutton
	formeditassocbuttonDB := orm.FormEditAssocButtonDB{}
	formeditassocbuttonDB.FormEditAssocButtonPointersEncoding = input.FormEditAssocButtonPointersEncoding
	formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButton_WOP(&input.FormEditAssocButton_WOP)

	_, err = db.Create(&formeditassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormEditAssocButton.CheckoutPhaseOneInstance(&formeditassocbuttonDB)
	formeditassocbutton := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]

	if formeditassocbutton != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formeditassocbutton)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formeditassocbuttonDB)
}

// GetFormEditAssocButton
//
// swagger:route GET /formeditassocbuttons/{ID} formeditassocbuttons getFormEditAssocButton
//
// Gets the details for a formeditassocbutton.
//
// Responses:
// default: genericError
//
//	200: formeditassocbuttonDBResponse
func (controller *Controller) GetFormEditAssocButton(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormEditAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	// Get formeditassocbuttonDB in DB
	var formeditassocbuttonDB orm.FormEditAssocButtonDB
	if _, err := db.First(&formeditassocbuttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formeditassocbuttonAPI orm.FormEditAssocButtonAPI
	formeditassocbuttonAPI.ID = formeditassocbuttonDB.ID
	formeditassocbuttonAPI.FormEditAssocButtonPointersEncoding = formeditassocbuttonDB.FormEditAssocButtonPointersEncoding
	formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton_WOP(&formeditassocbuttonAPI.FormEditAssocButton_WOP)

	c.JSON(http.StatusOK, formeditassocbuttonAPI)
}

// UpdateFormEditAssocButton
//
// swagger:route PATCH /formeditassocbuttons/{ID} formeditassocbuttons updateFormEditAssocButton
//
// # Update a formeditassocbutton
//
// Responses:
// default: genericError
//
//	200: formeditassocbuttonDBResponse
func (controller *Controller) UpdateFormEditAssocButton(c *gin.Context) {

	mutexFormEditAssocButton.Lock()
	defer mutexFormEditAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormEditAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	// Validate input
	var input orm.FormEditAssocButtonAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formeditassocbuttonDB orm.FormEditAssocButtonDB

	// fetch the formeditassocbutton
	_, err := db.First(&formeditassocbuttonDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButton_WOP(&input.FormEditAssocButton_WOP)
	formeditassocbuttonDB.FormEditAssocButtonPointersEncoding = input.FormEditAssocButtonPointersEncoding

	db, _ = db.Model(&formeditassocbuttonDB)
	_, err = db.Updates(&formeditassocbuttonDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formeditassocbuttonNew := new(models.FormEditAssocButton)
	formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton(formeditassocbuttonNew)

	// redeem pointers
	formeditassocbuttonDB.DecodePointers(backRepo, formeditassocbuttonNew)

	// get stage instance from DB instance, and call callback function
	formeditassocbuttonOld := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
	if formeditassocbuttonOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formeditassocbuttonOld, formeditassocbuttonNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formeditassocbuttonDB
	c.JSON(http.StatusOK, formeditassocbuttonDB)
}

// DeleteFormEditAssocButton
//
// swagger:route DELETE /formeditassocbuttons/{ID} formeditassocbuttons deleteFormEditAssocButton
//
// # Delete a formeditassocbutton
//
// default: genericError
//
//	200: formeditassocbuttonDBResponse
func (controller *Controller) DeleteFormEditAssocButton(c *gin.Context) {

	mutexFormEditAssocButton.Lock()
	defer mutexFormEditAssocButton.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormEditAssocButton", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormEditAssocButton.GetDB()

	// Get model if exist
	var formeditassocbuttonDB orm.FormEditAssocButtonDB
	if _, err := db.First(&formeditassocbuttonDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formeditassocbuttonDB)

	// get an instance (not staged) from DB instance, and call callback function
	formeditassocbuttonDeleted := new(models.FormEditAssocButton)
	formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton(formeditassocbuttonDeleted)

	// get stage instance from DB instance, and call callback function
	formeditassocbuttonStaged := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
	if formeditassocbuttonStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formeditassocbuttonStaged, formeditassocbuttonDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
