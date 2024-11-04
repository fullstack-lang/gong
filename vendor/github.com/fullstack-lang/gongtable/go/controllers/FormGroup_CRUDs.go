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
var __FormGroup__dummysDeclaration__ models.FormGroup
var __FormGroup_time__dummyDeclaration time.Duration

var mutexFormGroup sync.Mutex

// An FormGroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormGroup updateFormGroup deleteFormGroup
type FormGroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormGroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormGroup updateFormGroup
type FormGroupInput struct {
	// The FormGroup to submit or modify
	// in: body
	FormGroup *orm.FormGroupAPI
}

// GetFormGroups
//
// swagger:route GET /formgroups formgroups getFormGroups
//
// # Get all formgroups
//
// Responses:
// default: genericError
//
//	200: formgroupDBResponse
func (controller *Controller) GetFormGroups(c *gin.Context) {

	// source slice
	var formgroupDBs []orm.FormGroupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormGroup.GetDB()

	_, err := db.Find(&formgroupDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formgroupAPIs := make([]orm.FormGroupAPI, 0)

	// for each formgroup, update fields from the database nullable fields
	for idx := range formgroupDBs {
		formgroupDB := &formgroupDBs[idx]
		_ = formgroupDB
		var formgroupAPI orm.FormGroupAPI

		// insertion point for updating fields
		formgroupAPI.ID = formgroupDB.ID
		formgroupDB.CopyBasicFieldsToFormGroup_WOP(&formgroupAPI.FormGroup_WOP)
		formgroupAPI.FormGroupPointersEncoding = formgroupDB.FormGroupPointersEncoding
		formgroupAPIs = append(formgroupAPIs, formgroupAPI)
	}

	c.JSON(http.StatusOK, formgroupAPIs)
}

// PostFormGroup
//
// swagger:route POST /formgroups formgroups postFormGroup
//
// Creates a formgroup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormGroup(c *gin.Context) {

	mutexFormGroup.Lock()
	defer mutexFormGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormGroup.GetDB()

	// Validate input
	var input orm.FormGroupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formgroup
	formgroupDB := orm.FormGroupDB{}
	formgroupDB.FormGroupPointersEncoding = input.FormGroupPointersEncoding
	formgroupDB.CopyBasicFieldsFromFormGroup_WOP(&input.FormGroup_WOP)

	_, err = db.Create(&formgroupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormGroup.CheckoutPhaseOneInstance(&formgroupDB)
	formgroup := backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]

	if formgroup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formgroup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formgroupDB)
}

// GetFormGroup
//
// swagger:route GET /formgroups/{ID} formgroups getFormGroup
//
// Gets the details for a formgroup.
//
// Responses:
// default: genericError
//
//	200: formgroupDBResponse
func (controller *Controller) GetFormGroup(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormGroup.GetDB()

	// Get formgroupDB in DB
	var formgroupDB orm.FormGroupDB
	if _, err := db.First(&formgroupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formgroupAPI orm.FormGroupAPI
	formgroupAPI.ID = formgroupDB.ID
	formgroupAPI.FormGroupPointersEncoding = formgroupDB.FormGroupPointersEncoding
	formgroupDB.CopyBasicFieldsToFormGroup_WOP(&formgroupAPI.FormGroup_WOP)

	c.JSON(http.StatusOK, formgroupAPI)
}

// UpdateFormGroup
//
// swagger:route PATCH /formgroups/{ID} formgroups updateFormGroup
//
// # Update a formgroup
//
// Responses:
// default: genericError
//
//	200: formgroupDBResponse
func (controller *Controller) UpdateFormGroup(c *gin.Context) {

	mutexFormGroup.Lock()
	defer mutexFormGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormGroup.GetDB()

	// Validate input
	var input orm.FormGroupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formgroupDB orm.FormGroupDB

	// fetch the formgroup
	_, err := db.First(&formgroupDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formgroupDB.CopyBasicFieldsFromFormGroup_WOP(&input.FormGroup_WOP)
	formgroupDB.FormGroupPointersEncoding = input.FormGroupPointersEncoding

	db, _ = db.Model(&formgroupDB)
	_, err = db.Updates(&formgroupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formgroupNew := new(models.FormGroup)
	formgroupDB.CopyBasicFieldsToFormGroup(formgroupNew)

	// redeem pointers
	formgroupDB.DecodePointers(backRepo, formgroupNew)

	// get stage instance from DB instance, and call callback function
	formgroupOld := backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
	if formgroupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formgroupOld, formgroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formgroupDB
	c.JSON(http.StatusOK, formgroupDB)
}

// DeleteFormGroup
//
// swagger:route DELETE /formgroups/{ID} formgroups deleteFormGroup
//
// # Delete a formgroup
//
// default: genericError
//
//	200: formgroupDBResponse
func (controller *Controller) DeleteFormGroup(c *gin.Context) {

	mutexFormGroup.Lock()
	defer mutexFormGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormGroup.GetDB()

	// Get model if exist
	var formgroupDB orm.FormGroupDB
	if _, err := db.First(&formgroupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formgroupDB)

	// get an instance (not staged) from DB instance, and call callback function
	formgroupDeleted := new(models.FormGroup)
	formgroupDB.CopyBasicFieldsToFormGroup(formgroupDeleted)

	// get stage instance from DB instance, and call callback function
	formgroupStaged := backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
	if formgroupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formgroupStaged, formgroupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
