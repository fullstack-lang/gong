// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Status__dummysDeclaration__ models.Status
var __Status_time__dummyDeclaration time.Duration

var mutexStatus sync.Mutex

// An StatusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStatus updateStatus deleteStatus
type StatusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StatusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStatus updateStatus
type StatusInput struct {
	// The Status to submit or modify
	// in: body
	Status *orm.StatusAPI
}

// GetStatuss
//
// swagger:route GET /statuss statuss getStatuss
//
// # Get all statuss
//
// Responses:
// default: genericError
//
//	200: statusDBResponse
func (controller *Controller) GetStatuss(c *gin.Context) {

	// source slice
	var statusDBs []orm.StatusDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStatuss", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoStatus.GetDB()

	_, err := db.Find(&statusDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	statusAPIs := make([]orm.StatusAPI, 0)

	// for each status, update fields from the database nullable fields
	for idx := range statusDBs {
		statusDB := &statusDBs[idx]
		_ = statusDB
		var statusAPI orm.StatusAPI

		// insertion point for updating fields
		statusAPI.ID = statusDB.ID
		statusDB.CopyBasicFieldsToStatus_WOP(&statusAPI.Status_WOP)
		statusAPI.StatusPointersEncoding = statusDB.StatusPointersEncoding
		statusAPIs = append(statusAPIs, statusAPI)
	}

	c.JSON(http.StatusOK, statusAPIs)
}

// PostStatus
//
// swagger:route POST /statuss statuss postStatus
//
// Creates a status
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStatus(c *gin.Context) {

	mutexStatus.Lock()
	defer mutexStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStatuss", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoStatus.GetDB()

	// Validate input
	var input orm.StatusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create status
	statusDB := orm.StatusDB{}
	statusDB.StatusPointersEncoding = input.StatusPointersEncoding
	statusDB.CopyBasicFieldsFromStatus_WOP(&input.Status_WOP)

	_, err = db.Create(&statusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStatus.CheckoutPhaseOneInstance(&statusDB)
	status := backRepo.BackRepoStatus.Map_StatusDBID_StatusPtr[statusDB.ID]

	if status != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), status)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, statusDB)
}

// GetStatus
//
// swagger:route GET /statuss/{ID} statuss getStatus
//
// Gets the details for a status.
//
// Responses:
// default: genericError
//
//	200: statusDBResponse
func (controller *Controller) GetStatus(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStatus", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoStatus.GetDB()

	// Get statusDB in DB
	var statusDB orm.StatusDB
	if _, err := db.First(&statusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var statusAPI orm.StatusAPI
	statusAPI.ID = statusDB.ID
	statusAPI.StatusPointersEncoding = statusDB.StatusPointersEncoding
	statusDB.CopyBasicFieldsToStatus_WOP(&statusAPI.Status_WOP)

	c.JSON(http.StatusOK, statusAPI)
}

// UpdateStatus
//
// swagger:route PATCH /statuss/{ID} statuss updateStatus
//
// # Update a status
//
// Responses:
// default: genericError
//
//	200: statusDBResponse
func (controller *Controller) UpdateStatus(c *gin.Context) {

	mutexStatus.Lock()
	defer mutexStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStatus", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoStatus.GetDB()

	// Validate input
	var input orm.StatusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var statusDB orm.StatusDB

	// fetch the status
	_, err := db.First(&statusDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	statusDB.CopyBasicFieldsFromStatus_WOP(&input.Status_WOP)
	statusDB.StatusPointersEncoding = input.StatusPointersEncoding

	db, _ = db.Model(&statusDB)
	_, err = db.Updates(&statusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	statusNew := new(models.Status)
	statusDB.CopyBasicFieldsToStatus(statusNew)

	// redeem pointers
	statusDB.DecodePointers(backRepo, statusNew)

	// get stage instance from DB instance, and call callback function
	statusOld := backRepo.BackRepoStatus.Map_StatusDBID_StatusPtr[statusDB.ID]
	if statusOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), statusOld, statusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the statusDB
	c.JSON(http.StatusOK, statusDB)
}

// DeleteStatus
//
// swagger:route DELETE /statuss/{ID} statuss deleteStatus
//
// # Delete a status
//
// default: genericError
//
//	200: statusDBResponse
func (controller *Controller) DeleteStatus(c *gin.Context) {

	mutexStatus.Lock()
	defer mutexStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStatus", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoStatus.GetDB()

	// Get model if exist
	var statusDB orm.StatusDB
	if _, err := db.First(&statusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&statusDB)

	// get an instance (not staged) from DB instance, and call callback function
	statusDeleted := new(models.Status)
	statusDB.CopyBasicFieldsToStatus(statusDeleted)

	// get stage instance from DB instance, and call callback function
	statusStaged := backRepo.BackRepoStatus.Map_StatusDBID_StatusPtr[statusDB.ID]
	if statusStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), statusStaged, statusDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
