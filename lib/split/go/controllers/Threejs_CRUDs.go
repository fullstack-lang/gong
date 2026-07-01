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
var __Threejs__dummysDeclaration__ models.Threejs
var _ = __Threejs__dummysDeclaration__
var __Threejs_time__dummyDeclaration time.Duration
var _ = __Threejs_time__dummyDeclaration

var mutexThreejs sync.Mutex

// An ThreejsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getThreejs updateThreejs deleteThreejs
type ThreejsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ThreejsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postThreejs updateThreejs
type ThreejsInput struct {
	// The Threejs to submit or modify
	// in: body
	Threejs *orm.ThreejsAPI
}

// GetThreejss
//
// swagger:route GET /threejss threejss getThreejss
//
// # Get all threejss
//
// Responses:
// default: genericError
//
//	200: threejsDBResponse
func (controller *Controller) GetThreejss(c *gin.Context) {

	// source slice
	var threejsDBs []orm.ThreejsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetThreejss", "Name", stackPath)
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
	db := backRepo.BackRepoThreejs.GetDB()

	_, err := db.Find(&threejsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	threejsAPIs := make([]orm.ThreejsAPI, 0)

	// for each threejs, update fields from the database nullable fields
	for idx := range threejsDBs {
		threejsDB := &threejsDBs[idx]
		_ = threejsDB
		var threejsAPI orm.ThreejsAPI

		// insertion point for updating fields
		threejsAPI.ID = threejsDB.ID
		threejsDB.CopyBasicFieldsToThreejs_WOP(&threejsAPI.Threejs_WOP)
		threejsAPI.ThreejsPointersEncoding = threejsDB.ThreejsPointersEncoding
		threejsAPIs = append(threejsAPIs, threejsAPI)
	}

	c.JSON(http.StatusOK, threejsAPIs)
}

// PostThreejs
//
// swagger:route POST /threejss threejss postThreejs
//
// Creates a threejs
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostThreejs(c *gin.Context) {

	mutexThreejs.Lock()
	defer mutexThreejs.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostThreejss", "Name", stackPath)
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
	db := backRepo.BackRepoThreejs.GetDB()

	// Validate input
	var input orm.ThreejsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create threejs
	threejsDB := orm.ThreejsDB{}
	threejsDB.ThreejsPointersEncoding = input.ThreejsPointersEncoding
	threejsDB.CopyBasicFieldsFromThreejs_WOP(&input.Threejs_WOP)

	_, err = db.Create(&threejsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoThreejs.CheckoutPhaseOneInstance(&threejsDB)
	threejs := backRepo.BackRepoThreejs.Map_ThreejsDBID_ThreejsPtr[threejsDB.ID]

	if threejs != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), threejs)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, threejsDB)
}

// GetThreejs
//
// swagger:route GET /threejss/{ID} threejss getThreejs
//
// Gets the details for a threejs.
//
// Responses:
// default: genericError
//
//	200: threejsDBResponse
func (controller *Controller) GetThreejs(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetThreejs", "Name", stackPath)
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
	db := backRepo.BackRepoThreejs.GetDB()

	// Get threejsDB in DB
	var threejsDB orm.ThreejsDB
	if _, err := db.First(&threejsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var threejsAPI orm.ThreejsAPI
	threejsAPI.ID = threejsDB.ID
	threejsAPI.ThreejsPointersEncoding = threejsDB.ThreejsPointersEncoding
	threejsDB.CopyBasicFieldsToThreejs_WOP(&threejsAPI.Threejs_WOP)

	c.JSON(http.StatusOK, threejsAPI)
}

// UpdateThreejs
//
// swagger:route PATCH /threejss/{ID} threejss updateThreejs
//
// # Update a threejs
//
// Responses:
// default: genericError
//
//	200: threejsDBResponse
func (controller *Controller) UpdateThreejs(c *gin.Context) {

	mutexThreejs.Lock()
	defer mutexThreejs.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoThreejs.GetDB()

	// Validate input
	var input orm.ThreejsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var threejsDB orm.ThreejsDB

	// fetch the threejs
	_, err := db.First(&threejsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	threejsDB.CopyBasicFieldsFromThreejs_WOP(&input.Threejs_WOP)
	threejsDB.ThreejsPointersEncoding = input.ThreejsPointersEncoding

	db, _ = db.Model(&threejsDB)
	_, err = db.Updates(&threejsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	threejsNew := new(models.Threejs)
	threejsDB.CopyBasicFieldsToThreejs(threejsNew)

	// redeem pointers
	threejsDB.DecodePointers(backRepo, threejsNew)

	// get stage instance from DB instance, and call callback function
	threejsOld := backRepo.BackRepoThreejs.Map_ThreejsDBID_ThreejsPtr[threejsDB.ID]
	if threejsOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), threejsOld, threejsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the threejsDB
	c.JSON(http.StatusOK, threejsDB)
}

// DeleteThreejs
//
// swagger:route DELETE /threejss/{ID} threejss deleteThreejs
//
// # Delete a threejs
//
// default: genericError
//
//	200: threejsDBResponse
func (controller *Controller) DeleteThreejs(c *gin.Context) {

	mutexThreejs.Lock()
	defer mutexThreejs.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteThreejs", "Name", stackPath)
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
	db := backRepo.BackRepoThreejs.GetDB()

	// Get model if exist
	var threejsDB orm.ThreejsDB
	if _, err := db.First(&threejsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&threejsDB)

	// get an instance (not staged) from DB instance, and call callback function
	threejsDeleted := new(models.Threejs)
	threejsDB.CopyBasicFieldsToThreejs(threejsDeleted)

	// get stage instance from DB instance, and call callback function
	threejsStaged := backRepo.BackRepoThreejs.Map_ThreejsDBID_ThreejsPtr[threejsDB.ID]
	if threejsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), threejsStaged, threejsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
