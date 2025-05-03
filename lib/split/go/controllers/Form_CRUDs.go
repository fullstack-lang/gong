// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Form__dummysDeclaration__ models.Form
var __Form_time__dummyDeclaration time.Duration

var mutexForm sync.Mutex

// An FormID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getForm updateForm deleteForm
type FormID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postForm updateForm
type FormInput struct {
	// The Form to submit or modify
	// in: body
	Form *orm.FormAPI
}

// GetForms
//
// swagger:route GET /forms forms getForms
//
// # Get all forms
//
// Responses:
// default: genericError
//
//	200: formDBResponse
func (controller *Controller) GetForms(c *gin.Context) {

	// source slice
	var formDBs []orm.FormDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForms", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	_, err := db.Find(&formDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formAPIs := make([]orm.FormAPI, 0)

	// for each form, update fields from the database nullable fields
	for idx := range formDBs {
		formDB := &formDBs[idx]
		_ = formDB
		var formAPI orm.FormAPI

		// insertion point for updating fields
		formAPI.ID = formDB.ID
		formDB.CopyBasicFieldsToForm_WOP(&formAPI.Form_WOP)
		formAPI.FormPointersEncoding = formDB.FormPointersEncoding
		formAPIs = append(formAPIs, formAPI)
	}

	c.JSON(http.StatusOK, formAPIs)
}

// PostForm
//
// swagger:route POST /forms forms postForm
//
// Creates a form
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostForm(c *gin.Context) {

	mutexForm.Lock()
	defer mutexForm.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostForms", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	// Validate input
	var input orm.FormAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create form
	formDB := orm.FormDB{}
	formDB.FormPointersEncoding = input.FormPointersEncoding
	formDB.CopyBasicFieldsFromForm_WOP(&input.Form_WOP)

	_, err = db.Create(&formDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoForm.CheckoutPhaseOneInstance(&formDB)
	form := backRepo.BackRepoForm.Map_FormDBID_FormPtr[formDB.ID]

	if form != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), form)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formDB)
}

// GetForm
//
// swagger:route GET /forms/{ID} forms getForm
//
// Gets the details for a form.
//
// Responses:
// default: genericError
//
//	200: formDBResponse
func (controller *Controller) GetForm(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForm", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	// Get formDB in DB
	var formDB orm.FormDB
	if _, err := db.First(&formDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formAPI orm.FormAPI
	formAPI.ID = formDB.ID
	formAPI.FormPointersEncoding = formDB.FormPointersEncoding
	formDB.CopyBasicFieldsToForm_WOP(&formAPI.Form_WOP)

	c.JSON(http.StatusOK, formAPI)
}

// UpdateForm
//
// swagger:route PATCH /forms/{ID} forms updateForm
//
// # Update a form
//
// Responses:
// default: genericError
//
//	200: formDBResponse
func (controller *Controller) UpdateForm(c *gin.Context) {

	mutexForm.Lock()
	defer mutexForm.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateForm", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	// Validate input
	var input orm.FormAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formDB orm.FormDB

	// fetch the form
	_, err := db.First(&formDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formDB.CopyBasicFieldsFromForm_WOP(&input.Form_WOP)
	formDB.FormPointersEncoding = input.FormPointersEncoding

	db, _ = db.Model(&formDB)
	_, err = db.Updates(&formDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formNew := new(models.Form)
	formDB.CopyBasicFieldsToForm(formNew)

	// redeem pointers
	formDB.DecodePointers(backRepo, formNew)

	// get stage instance from DB instance, and call callback function
	formOld := backRepo.BackRepoForm.Map_FormDBID_FormPtr[formDB.ID]
	if formOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formOld, formNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formDB
	c.JSON(http.StatusOK, formDB)
}

// DeleteForm
//
// swagger:route DELETE /forms/{ID} forms deleteForm
//
// # Delete a form
//
// default: genericError
//
//	200: formDBResponse
func (controller *Controller) DeleteForm(c *gin.Context) {

	mutexForm.Lock()
	defer mutexForm.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteForm", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm.GetDB()

	// Get model if exist
	var formDB orm.FormDB
	if _, err := db.First(&formDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&formDB)

	// get an instance (not staged) from DB instance, and call callback function
	formDeleted := new(models.Form)
	formDB.CopyBasicFieldsToForm(formDeleted)

	// get stage instance from DB instance, and call callback function
	formStaged := backRepo.BackRepoForm.Map_FormDBID_FormPtr[formDB.ID]
	if formStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formStaged, formDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
