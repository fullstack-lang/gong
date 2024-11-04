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
var __CheckBox__dummysDeclaration__ models.CheckBox
var __CheckBox_time__dummyDeclaration time.Duration

var mutexCheckBox sync.Mutex

// An CheckBoxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCheckBox updateCheckBox deleteCheckBox
type CheckBoxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CheckBoxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCheckBox updateCheckBox
type CheckBoxInput struct {
	// The CheckBox to submit or modify
	// in: body
	CheckBox *orm.CheckBoxAPI
}

// GetCheckBoxs
//
// swagger:route GET /checkboxs checkboxs getCheckBoxs
//
// # Get all checkboxs
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) GetCheckBoxs(c *gin.Context) {

	// source slice
	var checkboxDBs []orm.CheckBoxDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCheckBoxs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

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
	checkboxAPIs := make([]orm.CheckBoxAPI, 0)

	// for each checkbox, update fields from the database nullable fields
	for idx := range checkboxDBs {
		checkboxDB := &checkboxDBs[idx]
		_ = checkboxDB
		var checkboxAPI orm.CheckBoxAPI

		// insertion point for updating fields
		checkboxAPI.ID = checkboxDB.ID
		checkboxDB.CopyBasicFieldsToCheckBox_WOP(&checkboxAPI.CheckBox_WOP)
		checkboxAPI.CheckBoxPointersEncoding = checkboxDB.CheckBoxPointersEncoding
		checkboxAPIs = append(checkboxAPIs, checkboxAPI)
	}

	c.JSON(http.StatusOK, checkboxAPIs)
}

// PostCheckBox
//
// swagger:route POST /checkboxs checkboxs postCheckBox
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
func (controller *Controller) PostCheckBox(c *gin.Context) {

	mutexCheckBox.Lock()
	defer mutexCheckBox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCheckBoxs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

	// Validate input
	var input orm.CheckBoxAPI

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
	checkboxDB := orm.CheckBoxDB{}
	checkboxDB.CheckBoxPointersEncoding = input.CheckBoxPointersEncoding
	checkboxDB.CopyBasicFieldsFromCheckBox_WOP(&input.CheckBox_WOP)

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
	backRepo.BackRepoCheckBox.CheckoutPhaseOneInstance(&checkboxDB)
	checkbox := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]

	if checkbox != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), checkbox)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, checkboxDB)
}

// GetCheckBox
//
// swagger:route GET /checkboxs/{ID} checkboxs getCheckBox
//
// Gets the details for a checkbox.
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) GetCheckBox(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCheckBox", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

	// Get checkboxDB in DB
	var checkboxDB orm.CheckBoxDB
	if _, err := db.First(&checkboxDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var checkboxAPI orm.CheckBoxAPI
	checkboxAPI.ID = checkboxDB.ID
	checkboxAPI.CheckBoxPointersEncoding = checkboxDB.CheckBoxPointersEncoding
	checkboxDB.CopyBasicFieldsToCheckBox_WOP(&checkboxAPI.CheckBox_WOP)

	c.JSON(http.StatusOK, checkboxAPI)
}

// UpdateCheckBox
//
// swagger:route PATCH /checkboxs/{ID} checkboxs updateCheckBox
//
// # Update a checkbox
//
// Responses:
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) UpdateCheckBox(c *gin.Context) {

	mutexCheckBox.Lock()
	defer mutexCheckBox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCheckBox", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

	// Validate input
	var input orm.CheckBoxAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var checkboxDB orm.CheckBoxDB

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
	checkboxDB.CopyBasicFieldsFromCheckBox_WOP(&input.CheckBox_WOP)
	checkboxDB.CheckBoxPointersEncoding = input.CheckBoxPointersEncoding

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
	checkboxNew := new(models.CheckBox)
	checkboxDB.CopyBasicFieldsToCheckBox(checkboxNew)

	// redeem pointers
	checkboxDB.DecodePointers(backRepo, checkboxNew)

	// get stage instance from DB instance, and call callback function
	checkboxOld := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]
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

// DeleteCheckBox
//
// swagger:route DELETE /checkboxs/{ID} checkboxs deleteCheckBox
//
// # Delete a checkbox
//
// default: genericError
//
//	200: checkboxDBResponse
func (controller *Controller) DeleteCheckBox(c *gin.Context) {

	mutexCheckBox.Lock()
	defer mutexCheckBox.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCheckBox", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCheckBox.GetDB()

	// Get model if exist
	var checkboxDB orm.CheckBoxDB
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
	checkboxDeleted := new(models.CheckBox)
	checkboxDB.CopyBasicFieldsToCheckBox(checkboxDeleted)

	// get stage instance from DB instance, and call callback function
	checkboxStaged := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]
	if checkboxStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), checkboxStaged, checkboxDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
