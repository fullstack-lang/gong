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
var __FormFieldString__dummysDeclaration__ models.FormFieldString
var __FormFieldString_time__dummyDeclaration time.Duration

var mutexFormFieldString sync.Mutex

// An FormFieldStringID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldString updateFormFieldString deleteFormFieldString
type FormFieldStringID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldStringInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldString updateFormFieldString
type FormFieldStringInput struct {
	// The FormFieldString to submit or modify
	// in: body
	FormFieldString *orm.FormFieldStringAPI
}

// GetFormFieldStrings
//
// swagger:route GET /formfieldstrings formfieldstrings getFormFieldStrings
//
// # Get all formfieldstrings
//
// Responses:
// default: genericError
//
//	200: formfieldstringDBResponse
func (controller *Controller) GetFormFieldStrings(c *gin.Context) {

	// source slice
	var formfieldstringDBs []orm.FormFieldStringDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldStrings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldString.GetDB()

	_, err := db.Find(&formfieldstringDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldstringAPIs := make([]orm.FormFieldStringAPI, 0)

	// for each formfieldstring, update fields from the database nullable fields
	for idx := range formfieldstringDBs {
		formfieldstringDB := &formfieldstringDBs[idx]
		_ = formfieldstringDB
		var formfieldstringAPI orm.FormFieldStringAPI

		// insertion point for updating fields
		formfieldstringAPI.ID = formfieldstringDB.ID
		formfieldstringDB.CopyBasicFieldsToFormFieldString_WOP(&formfieldstringAPI.FormFieldString_WOP)
		formfieldstringAPI.FormFieldStringPointersEncoding = formfieldstringDB.FormFieldStringPointersEncoding
		formfieldstringAPIs = append(formfieldstringAPIs, formfieldstringAPI)
	}

	c.JSON(http.StatusOK, formfieldstringAPIs)
}

// PostFormFieldString
//
// swagger:route POST /formfieldstrings formfieldstrings postFormFieldString
//
// Creates a formfieldstring
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldString(c *gin.Context) {

	mutexFormFieldString.Lock()
	defer mutexFormFieldString.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldStrings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldString.GetDB()

	// Validate input
	var input orm.FormFieldStringAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldstring
	formfieldstringDB := orm.FormFieldStringDB{}
	formfieldstringDB.FormFieldStringPointersEncoding = input.FormFieldStringPointersEncoding
	formfieldstringDB.CopyBasicFieldsFromFormFieldString_WOP(&input.FormFieldString_WOP)

	_, err = db.Create(&formfieldstringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldString.CheckoutPhaseOneInstance(&formfieldstringDB)
	formfieldstring := backRepo.BackRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]

	if formfieldstring != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfieldstring)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldstringDB)
}

// GetFormFieldString
//
// swagger:route GET /formfieldstrings/{ID} formfieldstrings getFormFieldString
//
// Gets the details for a formfieldstring.
//
// Responses:
// default: genericError
//
//	200: formfieldstringDBResponse
func (controller *Controller) GetFormFieldString(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldString.GetDB()

	// Get formfieldstringDB in DB
	var formfieldstringDB orm.FormFieldStringDB
	if _, err := db.First(&formfieldstringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldstringAPI orm.FormFieldStringAPI
	formfieldstringAPI.ID = formfieldstringDB.ID
	formfieldstringAPI.FormFieldStringPointersEncoding = formfieldstringDB.FormFieldStringPointersEncoding
	formfieldstringDB.CopyBasicFieldsToFormFieldString_WOP(&formfieldstringAPI.FormFieldString_WOP)

	c.JSON(http.StatusOK, formfieldstringAPI)
}

// UpdateFormFieldString
//
// swagger:route PATCH /formfieldstrings/{ID} formfieldstrings updateFormFieldString
//
// # Update a formfieldstring
//
// Responses:
// default: genericError
//
//	200: formfieldstringDBResponse
func (controller *Controller) UpdateFormFieldString(c *gin.Context) {

	mutexFormFieldString.Lock()
	defer mutexFormFieldString.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldString.GetDB()

	// Validate input
	var input orm.FormFieldStringAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldstringDB orm.FormFieldStringDB

	// fetch the formfieldstring
	_, err := db.First(&formfieldstringDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldstringDB.CopyBasicFieldsFromFormFieldString_WOP(&input.FormFieldString_WOP)
	formfieldstringDB.FormFieldStringPointersEncoding = input.FormFieldStringPointersEncoding

	db, _ = db.Model(&formfieldstringDB)
	_, err = db.Updates(&formfieldstringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldstringNew := new(models.FormFieldString)
	formfieldstringDB.CopyBasicFieldsToFormFieldString(formfieldstringNew)

	// redeem pointers
	formfieldstringDB.DecodePointers(backRepo, formfieldstringNew)

	// get stage instance from DB instance, and call callback function
	formfieldstringOld := backRepo.BackRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]
	if formfieldstringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldstringOld, formfieldstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldstringDB
	c.JSON(http.StatusOK, formfieldstringDB)
}

// DeleteFormFieldString
//
// swagger:route DELETE /formfieldstrings/{ID} formfieldstrings deleteFormFieldString
//
// # Delete a formfieldstring
//
// default: genericError
//
//	200: formfieldstringDBResponse
func (controller *Controller) DeleteFormFieldString(c *gin.Context) {

	mutexFormFieldString.Lock()
	defer mutexFormFieldString.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldString", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormFieldString.GetDB()

	// Get model if exist
	var formfieldstringDB orm.FormFieldStringDB
	if _, err := db.First(&formfieldstringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formfieldstringDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldstringDeleted := new(models.FormFieldString)
	formfieldstringDB.CopyBasicFieldsToFormFieldString(formfieldstringDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldstringStaged := backRepo.BackRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]
	if formfieldstringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldstringStaged, formfieldstringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
