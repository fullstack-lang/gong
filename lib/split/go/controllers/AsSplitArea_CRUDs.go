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
var __AsSplitArea__dummysDeclaration__ models.AsSplitArea
var __AsSplitArea_time__dummyDeclaration time.Duration

var mutexAsSplitArea sync.Mutex

// An AsSplitAreaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAsSplitArea updateAsSplitArea deleteAsSplitArea
type AsSplitAreaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AsSplitAreaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAsSplitArea updateAsSplitArea
type AsSplitAreaInput struct {
	// The AsSplitArea to submit or modify
	// in: body
	AsSplitArea *orm.AsSplitAreaAPI
}

// GetAsSplitAreas
//
// swagger:route GET /assplitareas assplitareas getAsSplitAreas
//
// # Get all assplitareas
//
// Responses:
// default: genericError
//
//	200: assplitareaDBResponse
func (controller *Controller) GetAsSplitAreas(c *gin.Context) {

	// source slice
	var assplitareaDBs []orm.AsSplitAreaDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAsSplitAreas", "Name", stackPath)
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
	db := backRepo.BackRepoAsSplitArea.GetDB()

	_, err := db.Find(&assplitareaDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	assplitareaAPIs := make([]orm.AsSplitAreaAPI, 0)

	// for each assplitarea, update fields from the database nullable fields
	for idx := range assplitareaDBs {
		assplitareaDB := &assplitareaDBs[idx]
		_ = assplitareaDB
		var assplitareaAPI orm.AsSplitAreaAPI

		// insertion point for updating fields
		assplitareaAPI.ID = assplitareaDB.ID
		assplitareaDB.CopyBasicFieldsToAsSplitArea_WOP(&assplitareaAPI.AsSplitArea_WOP)
		assplitareaAPI.AsSplitAreaPointersEncoding = assplitareaDB.AsSplitAreaPointersEncoding
		assplitareaAPIs = append(assplitareaAPIs, assplitareaAPI)
	}

	c.JSON(http.StatusOK, assplitareaAPIs)
}

// PostAsSplitArea
//
// swagger:route POST /assplitareas assplitareas postAsSplitArea
//
// Creates a assplitarea
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAsSplitArea(c *gin.Context) {

	mutexAsSplitArea.Lock()
	defer mutexAsSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAsSplitAreas", "Name", stackPath)
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
	db := backRepo.BackRepoAsSplitArea.GetDB()

	// Validate input
	var input orm.AsSplitAreaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create assplitarea
	assplitareaDB := orm.AsSplitAreaDB{}
	assplitareaDB.AsSplitAreaPointersEncoding = input.AsSplitAreaPointersEncoding
	assplitareaDB.CopyBasicFieldsFromAsSplitArea_WOP(&input.AsSplitArea_WOP)

	_, err = db.Create(&assplitareaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAsSplitArea.CheckoutPhaseOneInstance(&assplitareaDB)
	assplitarea := backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaPtr[assplitareaDB.ID]

	if assplitarea != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), assplitarea)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, assplitareaDB)
}

// GetAsSplitArea
//
// swagger:route GET /assplitareas/{ID} assplitareas getAsSplitArea
//
// Gets the details for a assplitarea.
//
// Responses:
// default: genericError
//
//	200: assplitareaDBResponse
func (controller *Controller) GetAsSplitArea(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAsSplitArea", "Name", stackPath)
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
	db := backRepo.BackRepoAsSplitArea.GetDB()

	// Get assplitareaDB in DB
	var assplitareaDB orm.AsSplitAreaDB
	if _, err := db.First(&assplitareaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var assplitareaAPI orm.AsSplitAreaAPI
	assplitareaAPI.ID = assplitareaDB.ID
	assplitareaAPI.AsSplitAreaPointersEncoding = assplitareaDB.AsSplitAreaPointersEncoding
	assplitareaDB.CopyBasicFieldsToAsSplitArea_WOP(&assplitareaAPI.AsSplitArea_WOP)

	c.JSON(http.StatusOK, assplitareaAPI)
}

// UpdateAsSplitArea
//
// swagger:route PATCH /assplitareas/{ID} assplitareas updateAsSplitArea
//
// # Update a assplitarea
//
// Responses:
// default: genericError
//
//	200: assplitareaDBResponse
func (controller *Controller) UpdateAsSplitArea(c *gin.Context) {

	mutexAsSplitArea.Lock()
	defer mutexAsSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAsSplitArea", "Name", stackPath)
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
	db := backRepo.BackRepoAsSplitArea.GetDB()

	// Validate input
	var input orm.AsSplitAreaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var assplitareaDB orm.AsSplitAreaDB

	// fetch the assplitarea
	_, err := db.First(&assplitareaDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	assplitareaDB.CopyBasicFieldsFromAsSplitArea_WOP(&input.AsSplitArea_WOP)
	assplitareaDB.AsSplitAreaPointersEncoding = input.AsSplitAreaPointersEncoding

	db, _ = db.Model(&assplitareaDB)
	_, err = db.Updates(&assplitareaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	assplitareaNew := new(models.AsSplitArea)
	assplitareaDB.CopyBasicFieldsToAsSplitArea(assplitareaNew)

	// redeem pointers
	assplitareaDB.DecodePointers(backRepo, assplitareaNew)

	// get stage instance from DB instance, and call callback function
	assplitareaOld := backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaPtr[assplitareaDB.ID]
	if assplitareaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), assplitareaOld, assplitareaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the assplitareaDB
	c.JSON(http.StatusOK, assplitareaDB)
}

// DeleteAsSplitArea
//
// swagger:route DELETE /assplitareas/{ID} assplitareas deleteAsSplitArea
//
// # Delete a assplitarea
//
// default: genericError
//
//	200: assplitareaDBResponse
func (controller *Controller) DeleteAsSplitArea(c *gin.Context) {

	mutexAsSplitArea.Lock()
	defer mutexAsSplitArea.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAsSplitArea", "Name", stackPath)
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
	db := backRepo.BackRepoAsSplitArea.GetDB()

	// Get model if exist
	var assplitareaDB orm.AsSplitAreaDB
	if _, err := db.First(&assplitareaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&assplitareaDB)

	// get an instance (not staged) from DB instance, and call callback function
	assplitareaDeleted := new(models.AsSplitArea)
	assplitareaDB.CopyBasicFieldsToAsSplitArea(assplitareaDeleted)

	// get stage instance from DB instance, and call callback function
	assplitareaStaged := backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaPtr[assplitareaDB.ID]
	if assplitareaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), assplitareaStaged, assplitareaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
