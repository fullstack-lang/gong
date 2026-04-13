// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	"github.com/fullstack-lang/gong/lib/form/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Form2__dummysDeclaration__ models.Form2
var _ = __Form2__dummysDeclaration__
var __Form2_time__dummyDeclaration time.Duration
var _ = __Form2_time__dummyDeclaration

var mutexForm2 sync.Mutex

// An Form2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getForm2 updateForm2 deleteForm2
type Form2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Form2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postForm2 updateForm2
type Form2Input struct {
	// The Form2 to submit or modify
	// in: body
	Form2 *orm.Form2API
}

// GetForm2s
//
// swagger:route GET /form2s form2s getForm2s
//
// # Get all form2s
//
// Responses:
// default: genericError
//
//	200: form2DBResponse
func (controller *Controller) GetForm2s(c *gin.Context) {

	// source slice
	var form2DBs []orm.Form2DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForm2s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm2.GetDB()

	_, err := db.Find(&form2DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	form2APIs := make([]orm.Form2API, 0)

	// for each form2, update fields from the database nullable fields
	for idx := range form2DBs {
		form2DB := &form2DBs[idx]
		_ = form2DB
		var form2API orm.Form2API

		// insertion point for updating fields
		form2API.ID = form2DB.ID
		form2DB.CopyBasicFieldsToForm2_WOP(&form2API.Form2_WOP)
		form2API.Form2PointersEncoding = form2DB.Form2PointersEncoding
		form2APIs = append(form2APIs, form2API)
	}

	c.JSON(http.StatusOK, form2APIs)
}

// PostForm2
//
// swagger:route POST /form2s form2s postForm2
//
// Creates a form2
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostForm2(c *gin.Context) {

	mutexForm2.Lock()
	defer mutexForm2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostForm2s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm2.GetDB()

	// Validate input
	var input orm.Form2API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create form2
	form2DB := orm.Form2DB{}
	form2DB.Form2PointersEncoding = input.Form2PointersEncoding
	form2DB.CopyBasicFieldsFromForm2_WOP(&input.Form2_WOP)

	_, err = db.Create(&form2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoForm2.CheckoutPhaseOneInstance(&form2DB)
	form2 := backRepo.BackRepoForm2.Map_Form2DBID_Form2Ptr[form2DB.ID]

	if form2 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), form2)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, form2DB)
}

// GetForm2
//
// swagger:route GET /form2s/{ID} form2s getForm2
//
// Gets the details for a form2.
//
// Responses:
// default: genericError
//
//	200: form2DBResponse
func (controller *Controller) GetForm2(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForm2", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm2.GetDB()

	// Get form2DB in DB
	var form2DB orm.Form2DB
	if _, err := db.First(&form2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var form2API orm.Form2API
	form2API.ID = form2DB.ID
	form2API.Form2PointersEncoding = form2DB.Form2PointersEncoding
	form2DB.CopyBasicFieldsToForm2_WOP(&form2API.Form2_WOP)

	c.JSON(http.StatusOK, form2API)
}

// UpdateForm2
//
// swagger:route PATCH /form2s/{ID} form2s updateForm2
//
// # Update a form2
//
// Responses:
// default: genericError
//
//	200: form2DBResponse
func (controller *Controller) UpdateForm2(c *gin.Context) {

	mutexForm2.Lock()
	defer mutexForm2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm2.GetDB()

	// Validate input
	var input orm.Form2API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var form2DB orm.Form2DB

	// fetch the form2
	_, err := db.First(&form2DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	form2DB.CopyBasicFieldsFromForm2_WOP(&input.Form2_WOP)
	form2DB.Form2PointersEncoding = input.Form2PointersEncoding

	db, _ = db.Model(&form2DB)
	_, err = db.Updates(&form2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	form2New := new(models.Form2)
	form2DB.CopyBasicFieldsToForm2(form2New)

	// redeem pointers
	form2DB.DecodePointers(backRepo, form2New)

	// get stage instance from DB instance, and call callback function
	form2Old := backRepo.BackRepoForm2.Map_Form2DBID_Form2Ptr[form2DB.ID]
	if form2Old != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), form2Old, form2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the form2DB
	c.JSON(http.StatusOK, form2DB)
}

// DeleteForm2
//
// swagger:route DELETE /form2s/{ID} form2s deleteForm2
//
// # Delete a form2
//
// default: genericError
//
//	200: form2DBResponse
func (controller *Controller) DeleteForm2(c *gin.Context) {

	mutexForm2.Lock()
	defer mutexForm2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteForm2", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoForm2.GetDB()

	// Get model if exist
	var form2DB orm.Form2DB
	if _, err := db.First(&form2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&form2DB)

	// get an instance (not staged) from DB instance, and call callback function
	form2Deleted := new(models.Form2)
	form2DB.CopyBasicFieldsToForm2(form2Deleted)

	// get stage instance from DB instance, and call callback function
	form2Staged := backRepo.BackRepoForm2.Map_Form2DBID_Form2Ptr[form2DB.ID]
	if form2Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), form2Staged, form2Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
