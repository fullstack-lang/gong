// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	"github.com/fullstack-lang/gong/test/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __F0123456789012345678901234567890__dummysDeclaration__ models.F0123456789012345678901234567890
var __F0123456789012345678901234567890_time__dummyDeclaration time.Duration

var mutexF0123456789012345678901234567890 sync.Mutex

// An F0123456789012345678901234567890ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getF0123456789012345678901234567890 updateF0123456789012345678901234567890 deleteF0123456789012345678901234567890
type F0123456789012345678901234567890ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// F0123456789012345678901234567890Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postF0123456789012345678901234567890 updateF0123456789012345678901234567890
type F0123456789012345678901234567890Input struct {
	// The F0123456789012345678901234567890 to submit or modify
	// in: body
	F0123456789012345678901234567890 *orm.F0123456789012345678901234567890API
}

// GetF0123456789012345678901234567890s
//
// swagger:route GET /f0123456789012345678901234567890s f0123456789012345678901234567890s getF0123456789012345678901234567890s
//
// # Get all f0123456789012345678901234567890s
//
// Responses:
// default: genericError
//
//	200: f0123456789012345678901234567890DBResponse
func (controller *Controller) GetF0123456789012345678901234567890s(c *gin.Context) {

	// source slice
	var f0123456789012345678901234567890DBs []orm.F0123456789012345678901234567890DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetF0123456789012345678901234567890s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	_, err := db.Find(&f0123456789012345678901234567890DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	f0123456789012345678901234567890APIs := make([]orm.F0123456789012345678901234567890API, 0)

	// for each f0123456789012345678901234567890, update fields from the database nullable fields
	for idx := range f0123456789012345678901234567890DBs {
		f0123456789012345678901234567890DB := &f0123456789012345678901234567890DBs[idx]
		_ = f0123456789012345678901234567890DB
		var f0123456789012345678901234567890API orm.F0123456789012345678901234567890API

		// insertion point for updating fields
		f0123456789012345678901234567890API.ID = f0123456789012345678901234567890DB.ID
		f0123456789012345678901234567890DB.CopyBasicFieldsToF0123456789012345678901234567890_WOP(&f0123456789012345678901234567890API.F0123456789012345678901234567890_WOP)
		f0123456789012345678901234567890API.F0123456789012345678901234567890PointersEncoding = f0123456789012345678901234567890DB.F0123456789012345678901234567890PointersEncoding
		f0123456789012345678901234567890APIs = append(f0123456789012345678901234567890APIs, f0123456789012345678901234567890API)
	}

	c.JSON(http.StatusOK, f0123456789012345678901234567890APIs)
}

// PostF0123456789012345678901234567890
//
// swagger:route POST /f0123456789012345678901234567890s f0123456789012345678901234567890s postF0123456789012345678901234567890
//
// Creates a f0123456789012345678901234567890
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostF0123456789012345678901234567890(c *gin.Context) {

	mutexF0123456789012345678901234567890.Lock()
	defer mutexF0123456789012345678901234567890.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostF0123456789012345678901234567890s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	// Validate input
	var input orm.F0123456789012345678901234567890API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create f0123456789012345678901234567890
	f0123456789012345678901234567890DB := orm.F0123456789012345678901234567890DB{}
	f0123456789012345678901234567890DB.F0123456789012345678901234567890PointersEncoding = input.F0123456789012345678901234567890PointersEncoding
	f0123456789012345678901234567890DB.CopyBasicFieldsFromF0123456789012345678901234567890_WOP(&input.F0123456789012345678901234567890_WOP)

	_, err = db.Create(&f0123456789012345678901234567890DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoF0123456789012345678901234567890.CheckoutPhaseOneInstance(&f0123456789012345678901234567890DB)
	f0123456789012345678901234567890 := backRepo.BackRepoF0123456789012345678901234567890.Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890Ptr[f0123456789012345678901234567890DB.ID]

	if f0123456789012345678901234567890 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), f0123456789012345678901234567890)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, f0123456789012345678901234567890DB)
}

// GetF0123456789012345678901234567890
//
// swagger:route GET /f0123456789012345678901234567890s/{ID} f0123456789012345678901234567890s getF0123456789012345678901234567890
//
// Gets the details for a f0123456789012345678901234567890.
//
// Responses:
// default: genericError
//
//	200: f0123456789012345678901234567890DBResponse
func (controller *Controller) GetF0123456789012345678901234567890(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetF0123456789012345678901234567890", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	// Get f0123456789012345678901234567890DB in DB
	var f0123456789012345678901234567890DB orm.F0123456789012345678901234567890DB
	if _, err := db.First(&f0123456789012345678901234567890DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var f0123456789012345678901234567890API orm.F0123456789012345678901234567890API
	f0123456789012345678901234567890API.ID = f0123456789012345678901234567890DB.ID
	f0123456789012345678901234567890API.F0123456789012345678901234567890PointersEncoding = f0123456789012345678901234567890DB.F0123456789012345678901234567890PointersEncoding
	f0123456789012345678901234567890DB.CopyBasicFieldsToF0123456789012345678901234567890_WOP(&f0123456789012345678901234567890API.F0123456789012345678901234567890_WOP)

	c.JSON(http.StatusOK, f0123456789012345678901234567890API)
}

// UpdateF0123456789012345678901234567890
//
// swagger:route PATCH /f0123456789012345678901234567890s/{ID} f0123456789012345678901234567890s updateF0123456789012345678901234567890
//
// # Update a f0123456789012345678901234567890
//
// Responses:
// default: genericError
//
//	200: f0123456789012345678901234567890DBResponse
func (controller *Controller) UpdateF0123456789012345678901234567890(c *gin.Context) {

	mutexF0123456789012345678901234567890.Lock()
	defer mutexF0123456789012345678901234567890.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateF0123456789012345678901234567890", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	// Validate input
	var input orm.F0123456789012345678901234567890API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var f0123456789012345678901234567890DB orm.F0123456789012345678901234567890DB

	// fetch the f0123456789012345678901234567890
	_, err := db.First(&f0123456789012345678901234567890DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	f0123456789012345678901234567890DB.CopyBasicFieldsFromF0123456789012345678901234567890_WOP(&input.F0123456789012345678901234567890_WOP)
	f0123456789012345678901234567890DB.F0123456789012345678901234567890PointersEncoding = input.F0123456789012345678901234567890PointersEncoding

	db, _ = db.Model(&f0123456789012345678901234567890DB)
	_, err = db.Updates(&f0123456789012345678901234567890DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	f0123456789012345678901234567890New := new(models.F0123456789012345678901234567890)
	f0123456789012345678901234567890DB.CopyBasicFieldsToF0123456789012345678901234567890(f0123456789012345678901234567890New)

	// redeem pointers
	f0123456789012345678901234567890DB.DecodePointers(backRepo, f0123456789012345678901234567890New)

	// get stage instance from DB instance, and call callback function
	f0123456789012345678901234567890Old := backRepo.BackRepoF0123456789012345678901234567890.Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890Ptr[f0123456789012345678901234567890DB.ID]
	if f0123456789012345678901234567890Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), f0123456789012345678901234567890Old, f0123456789012345678901234567890New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the f0123456789012345678901234567890DB
	c.JSON(http.StatusOK, f0123456789012345678901234567890DB)
}

// DeleteF0123456789012345678901234567890
//
// swagger:route DELETE /f0123456789012345678901234567890s/{ID} f0123456789012345678901234567890s deleteF0123456789012345678901234567890
//
// # Delete a f0123456789012345678901234567890
//
// default: genericError
//
//	200: f0123456789012345678901234567890DBResponse
func (controller *Controller) DeleteF0123456789012345678901234567890(c *gin.Context) {

	mutexF0123456789012345678901234567890.Lock()
	defer mutexF0123456789012345678901234567890.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteF0123456789012345678901234567890", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoF0123456789012345678901234567890.GetDB()

	// Get model if exist
	var f0123456789012345678901234567890DB orm.F0123456789012345678901234567890DB
	if _, err := db.First(&f0123456789012345678901234567890DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&f0123456789012345678901234567890DB)

	// get an instance (not staged) from DB instance, and call callback function
	f0123456789012345678901234567890Deleted := new(models.F0123456789012345678901234567890)
	f0123456789012345678901234567890DB.CopyBasicFieldsToF0123456789012345678901234567890(f0123456789012345678901234567890Deleted)

	// get stage instance from DB instance, and call callback function
	f0123456789012345678901234567890Staged := backRepo.BackRepoF0123456789012345678901234567890.Map_F0123456789012345678901234567890DBID_F0123456789012345678901234567890Ptr[f0123456789012345678901234567890DB.ID]
	if f0123456789012345678901234567890Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), f0123456789012345678901234567890Staged, f0123456789012345678901234567890Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
