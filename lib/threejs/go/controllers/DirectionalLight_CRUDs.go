// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __DirectionalLight__dummysDeclaration__ models.DirectionalLight
var _ = __DirectionalLight__dummysDeclaration__
var __DirectionalLight_time__dummyDeclaration time.Duration
var _ = __DirectionalLight_time__dummyDeclaration

var mutexDirectionalLight sync.Mutex

// An DirectionalLightID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDirectionalLight updateDirectionalLight deleteDirectionalLight
type DirectionalLightID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DirectionalLightInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDirectionalLight updateDirectionalLight
type DirectionalLightInput struct {
	// The DirectionalLight to submit or modify
	// in: body
	DirectionalLight *orm.DirectionalLightAPI
}

// GetDirectionalLights
//
// swagger:route GET /directionallights directionallights getDirectionalLights
//
// # Get all directionallights
//
// Responses:
// default: genericError
//
//	200: directionallightDBResponse
func (controller *Controller) GetDirectionalLights(c *gin.Context) {

	// source slice
	var directionallightDBs []orm.DirectionalLightDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirectionalLights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	_, err := db.Find(&directionallightDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	directionallightAPIs := make([]orm.DirectionalLightAPI, 0)

	// for each directionallight, update fields from the database nullable fields
	for idx := range directionallightDBs {
		directionallightDB := &directionallightDBs[idx]
		_ = directionallightDB
		var directionallightAPI orm.DirectionalLightAPI

		// insertion point for updating fields
		directionallightAPI.ID = directionallightDB.ID
		directionallightDB.CopyBasicFieldsToDirectionalLight_WOP(&directionallightAPI.DirectionalLight_WOP)
		directionallightAPI.DirectionalLightPointersEncoding = directionallightDB.DirectionalLightPointersEncoding
		directionallightAPIs = append(directionallightAPIs, directionallightAPI)
	}

	c.JSON(http.StatusOK, directionallightAPIs)
}

// PostDirectionalLight
//
// swagger:route POST /directionallights directionallights postDirectionalLight
//
// Creates a directionallight
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDirectionalLight(c *gin.Context) {

	mutexDirectionalLight.Lock()
	defer mutexDirectionalLight.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDirectionalLights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	// Validate input
	var input orm.DirectionalLightAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create directionallight
	directionallightDB := orm.DirectionalLightDB{}
	directionallightDB.DirectionalLightPointersEncoding = input.DirectionalLightPointersEncoding
	directionallightDB.CopyBasicFieldsFromDirectionalLight_WOP(&input.DirectionalLight_WOP)

	_, err = db.Create(&directionallightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDirectionalLight.CheckoutPhaseOneInstance(&directionallightDB)
	directionallight := backRepo.BackRepoDirectionalLight.Map_DirectionalLightDBID_DirectionalLightPtr[directionallightDB.ID]

	if directionallight != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), directionallight)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, directionallightDB)
}

// GetDirectionalLight
//
// swagger:route GET /directionallights/{ID} directionallights getDirectionalLight
//
// Gets the details for a directionallight.
//
// Responses:
// default: genericError
//
//	200: directionallightDBResponse
func (controller *Controller) GetDirectionalLight(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirectionalLight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	// Get directionallightDB in DB
	var directionallightDB orm.DirectionalLightDB
	if _, err := db.First(&directionallightDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var directionallightAPI orm.DirectionalLightAPI
	directionallightAPI.ID = directionallightDB.ID
	directionallightAPI.DirectionalLightPointersEncoding = directionallightDB.DirectionalLightPointersEncoding
	directionallightDB.CopyBasicFieldsToDirectionalLight_WOP(&directionallightAPI.DirectionalLight_WOP)

	c.JSON(http.StatusOK, directionallightAPI)
}

// UpdateDirectionalLight
//
// swagger:route PATCH /directionallights/{ID} directionallights updateDirectionalLight
//
// # Update a directionallight
//
// Responses:
// default: genericError
//
//	200: directionallightDBResponse
func (controller *Controller) UpdateDirectionalLight(c *gin.Context) {

	mutexDirectionalLight.Lock()
	defer mutexDirectionalLight.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	// Validate input
	var input orm.DirectionalLightAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var directionallightDB orm.DirectionalLightDB

	// fetch the directionallight
	_, err := db.First(&directionallightDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	directionallightDB.CopyBasicFieldsFromDirectionalLight_WOP(&input.DirectionalLight_WOP)
	directionallightDB.DirectionalLightPointersEncoding = input.DirectionalLightPointersEncoding

	db, _ = db.Model(&directionallightDB)
	_, err = db.Updates(&directionallightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	directionallightNew := new(models.DirectionalLight)
	directionallightDB.CopyBasicFieldsToDirectionalLight(directionallightNew)

	// redeem pointers
	directionallightDB.DecodePointers(backRepo, directionallightNew)

	// get stage instance from DB instance, and call callback function
	directionallightOld := backRepo.BackRepoDirectionalLight.Map_DirectionalLightDBID_DirectionalLightPtr[directionallightDB.ID]
	if directionallightOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), directionallightOld, directionallightNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the directionallightDB
	c.JSON(http.StatusOK, directionallightDB)
}

// DeleteDirectionalLight
//
// swagger:route DELETE /directionallights/{ID} directionallights deleteDirectionalLight
//
// # Delete a directionallight
//
// default: genericError
//
//	200: directionallightDBResponse
func (controller *Controller) DeleteDirectionalLight(c *gin.Context) {

	mutexDirectionalLight.Lock()
	defer mutexDirectionalLight.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDirectionalLight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDirectionalLight.GetDB()

	// Get model if exist
	var directionallightDB orm.DirectionalLightDB
	if _, err := db.First(&directionallightDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&directionallightDB)

	// get an instance (not staged) from DB instance, and call callback function
	directionallightDeleted := new(models.DirectionalLight)
	directionallightDB.CopyBasicFieldsToDirectionalLight(directionallightDeleted)

	// get stage instance from DB instance, and call callback function
	directionallightStaged := backRepo.BackRepoDirectionalLight.Map_DirectionalLightDBID_DirectionalLightPtr[directionallightDB.ID]
	if directionallightStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), directionallightStaged, directionallightDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
