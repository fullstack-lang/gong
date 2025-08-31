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
var __Astruct__dummysDeclaration__ models.Astruct
var __Astruct_time__dummyDeclaration time.Duration

var mutexAstruct sync.Mutex

// An AstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAstruct updateAstruct deleteAstruct
type AstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAstruct updateAstruct
type AstructInput struct {
	// The Astruct to submit or modify
	// in: body
	Astruct *orm.AstructAPI
}

// GetAstructs
//
// swagger:route GET /astructs astructs getAstructs
//
// # Get all astructs
//
// Responses:
// default: genericError
//
//	200: astructDBResponse
func (controller *Controller) GetAstructs(c *gin.Context) {

	// source slice
	var astructDBs []orm.AstructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstructs", "Name", stackPath)
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
	db := backRepo.BackRepoAstruct.GetDB()

	_, err := db.Find(&astructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	astructAPIs := make([]orm.AstructAPI, 0)

	// for each astruct, update fields from the database nullable fields
	for idx := range astructDBs {
		astructDB := &astructDBs[idx]
		_ = astructDB
		var astructAPI orm.AstructAPI

		// insertion point for updating fields
		astructAPI.ID = astructDB.ID
		astructDB.CopyBasicFieldsToAstruct_WOP(&astructAPI.Astruct_WOP)
		astructAPI.AstructPointersEncoding = astructDB.AstructPointersEncoding
		astructAPIs = append(astructAPIs, astructAPI)
	}

	c.JSON(http.StatusOK, astructAPIs)
}

// PostAstruct
//
// swagger:route POST /astructs astructs postAstruct
//
// Creates a astruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAstruct(c *gin.Context) {

	mutexAstruct.Lock()
	defer mutexAstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAstructs", "Name", stackPath)
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
	db := backRepo.BackRepoAstruct.GetDB()

	// Validate input
	var input orm.AstructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create astruct
	astructDB := orm.AstructDB{}
	astructDB.AstructPointersEncoding = input.AstructPointersEncoding
	astructDB.CopyBasicFieldsFromAstruct_WOP(&input.Astruct_WOP)

	_, err = db.Create(&astructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAstruct.CheckoutPhaseOneInstance(&astructDB)
	astruct := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[astructDB.ID]

	if astruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), astruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, astructDB)
}

// GetAstruct
//
// swagger:route GET /astructs/{ID} astructs getAstruct
//
// Gets the details for a astruct.
//
// Responses:
// default: genericError
//
//	200: astructDBResponse
func (controller *Controller) GetAstruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstruct", "Name", stackPath)
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
	db := backRepo.BackRepoAstruct.GetDB()

	// Get astructDB in DB
	var astructDB orm.AstructDB
	if _, err := db.First(&astructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var astructAPI orm.AstructAPI
	astructAPI.ID = astructDB.ID
	astructAPI.AstructPointersEncoding = astructDB.AstructPointersEncoding
	astructDB.CopyBasicFieldsToAstruct_WOP(&astructAPI.Astruct_WOP)

	c.JSON(http.StatusOK, astructAPI)
}

// UpdateAstruct
//
// swagger:route PATCH /astructs/{ID} astructs updateAstruct
//
// # Update a astruct
//
// Responses:
// default: genericError
//
//	200: astructDBResponse
func (controller *Controller) UpdateAstruct(c *gin.Context) {

	mutexAstruct.Lock()
	defer mutexAstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
			// log.Println("UpdateAstruct", "Name", stackPath)
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
			// log.Println("UpdateAstruct", "Name", stackPath)
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
	db := backRepo.BackRepoAstruct.GetDB()

	// Validate input
	var input orm.AstructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructDB orm.AstructDB

	// fetch the astruct
	_, err := db.First(&astructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructDB.CopyBasicFieldsFromAstruct_WOP(&input.Astruct_WOP)
	astructDB.AstructPointersEncoding = input.AstructPointersEncoding

	db, _ = db.Model(&astructDB)
	_, err = db.Updates(&astructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructNew := new(models.Astruct)
	astructDB.CopyBasicFieldsToAstruct(astructNew)

	// redeem pointers
	astructDB.DecodePointers(backRepo, astructNew)

	// get stage instance from DB instance, and call callback function
	astructOld := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[astructDB.ID]
	if astructOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), astructOld, astructNew, nil)
		} else {
			mouseEvent := &models.MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), astructOld, astructNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructDB
	c.JSON(http.StatusOK, astructDB)
}

// DeleteAstruct
//
// swagger:route DELETE /astructs/{ID} astructs deleteAstruct
//
// # Delete a astruct
//
// default: genericError
//
//	200: astructDBResponse
func (controller *Controller) DeleteAstruct(c *gin.Context) {

	mutexAstruct.Lock()
	defer mutexAstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAstruct", "Name", stackPath)
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
	db := backRepo.BackRepoAstruct.GetDB()

	// Get model if exist
	var astructDB orm.AstructDB
	if _, err := db.First(&astructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&astructDB)

	// get an instance (not staged) from DB instance, and call callback function
	astructDeleted := new(models.Astruct)
	astructDB.CopyBasicFieldsToAstruct(astructDeleted)

	// get stage instance from DB instance, and call callback function
	astructStaged := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[astructDB.ID]
	if astructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), astructStaged, astructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
