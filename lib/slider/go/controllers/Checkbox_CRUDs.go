// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Checkbox__dummysDeclaration__ models.Checkbox
var __Checkbox_time__dummyDeclaration time.Duration

var mutexCheckbox sync.Mutex

// An CheckboxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCheckbox updateCheckbox deleteCheckbox
type CheckboxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CheckboxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCheckbox updateCheckbox
type CheckboxInput struct {
	// The Checkbox to submit or modify
	// in: body
	Checkbox *orm.CheckboxAPI
}

// GetCheckboxs
//
// swagger:route GET /checkboxs checkboxs getCheckboxs
//
// # Get all checkboxs
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) GetCheckboxs(c *gin.Context) {

	// source slice
	var checkboxDBs []orm.CheckboxDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCheckboxs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	_, err := db.Find(&checkboxDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	checkboxAPIs := make([]orm.CheckboxAPI, 0)

	// for each checkbox, update fields from the database nullable fields
	for idx := range checkboxDBs {
		checkboxDB := &checkboxDBs[idx]
		_ = checkboxDB
		var checkboxAPI orm.CheckboxAPI

		// insertion point for updating fields
		checkboxAPI.ID = checkboxDB.ID
		checkboxDB.CopyBasicFieldsToCheckbox_WOP(&checkboxAPI.Checkbox_WOP)
		checkboxAPI.CheckboxPointersEncoding = checkboxDB.CheckboxPointersEncoding
		checkboxAPIs = append(checkboxAPIs, checkboxAPI)
	}

	c.JSON(http.StatusOK, checkboxAPIs)
}

// PostCheckbox
//
// swagger:route POST /checkboxs checkboxs postCheckbox
//
// Creates a checkbox
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCheckbox(c *gin.Context) {

	mutexCheckbox.Lock()
	defer mutexCheckbox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCheckboxs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	// Validate input
	var input orm.CheckboxAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create checkbox
	checkboxDB := orm.CheckboxDB{}
	checkboxDB.CheckboxPointersEncoding = input.CheckboxPointersEncoding
	checkboxDB.CopyBasicFieldsFromCheckbox_WOP(&input.Checkbox_WOP)

	_, err = db.Create(&checkboxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCheckbox.CheckoutPhaseOneInstance(&checkboxDB)
	checkbox := backRepo.BackRepoCheckbox.Map_CheckboxDBID_CheckboxPtr[checkboxDB.ID]

	if checkbox != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), checkbox)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, checkboxDB)
}

// GetCheckbox
//
// swagger:route GET /checkboxs/{ID} checkboxs getCheckbox
//
// Gets the details for a checkbox.
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) GetCheckbox(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCheckbox", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	// Get checkboxDB in DB
	var checkboxDB orm.CheckboxDB
	if _, err := db.First(&checkboxDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var checkboxAPI orm.CheckboxAPI
	checkboxAPI.ID = checkboxDB.ID
	checkboxAPI.CheckboxPointersEncoding = checkboxDB.CheckboxPointersEncoding
	checkboxDB.CopyBasicFieldsToCheckbox_WOP(&checkboxAPI.Checkbox_WOP)

	c.JSON(http.StatusOK, checkboxAPI)
}

// UpdateCheckbox
//
// swagger:route PATCH /checkboxs/{ID} checkboxs updateCheckbox
//
// # Update a checkbox
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) UpdateCheckbox(c *gin.Context) {

	mutexCheckbox.Lock()
	defer mutexCheckbox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCheckbox", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	// Validate input
	var input orm.CheckboxAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var checkboxDB orm.CheckboxDB

	// fetch the checkbox
	_, err := db.First(&checkboxDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	checkboxDB.CopyBasicFieldsFromCheckbox_WOP(&input.Checkbox_WOP)
	checkboxDB.CheckboxPointersEncoding = input.CheckboxPointersEncoding

	db, _ = db.Model(&checkboxDB)
	_, err = db.Updates(&checkboxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	checkboxNew := new(models.Checkbox)
	checkboxDB.CopyBasicFieldsToCheckbox(checkboxNew)

	// redeem pointers
	checkboxDB.DecodePointers(backRepo, checkboxNew)

	// get stage instance from DB instance, and call callback function
	checkboxOld := backRepo.BackRepoCheckbox.Map_CheckboxDBID_CheckboxPtr[checkboxDB.ID]
	if checkboxOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), checkboxOld, checkboxNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the checkboxDB
	c.JSON(http.StatusOK, checkboxDB)
}

// DeleteCheckbox
//
// swagger:route DELETE /checkboxs/{ID} checkboxs deleteCheckbox
//
// # Delete a checkbox
//
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) DeleteCheckbox(c *gin.Context) {

	mutexCheckbox.Lock()
	defer mutexCheckbox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCheckbox", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoCheckbox.GetDB()

	// Get model if exist
	var checkboxDB orm.CheckboxDB
	if _, err := db.First(&checkboxDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&checkboxDB)

	// get an instance (not staged) from DB instance, and call callback function
	checkboxDeleted := new(models.Checkbox)
	checkboxDB.CopyBasicFieldsToCheckbox(checkboxDeleted)

	// get stage instance from DB instance, and call callback function
	checkboxStaged := backRepo.BackRepoCheckbox.Map_CheckboxDBID_CheckboxPtr[checkboxDB.ID]
	if checkboxStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), checkboxStaged, checkboxDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
