// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	"github.com/fullstack-lang/gong/lib/gantt/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __LaneUse__dummysDeclaration__ models.LaneUse
var __LaneUse_time__dummyDeclaration time.Duration

var mutexLaneUse sync.Mutex

// An LaneUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLaneUse updateLaneUse deleteLaneUse
type LaneUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LaneUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLaneUse updateLaneUse
type LaneUseInput struct {
	// The LaneUse to submit or modify
	// in: body
	LaneUse *orm.LaneUseAPI
}

// GetLaneUses
//
// swagger:route GET /laneuses laneuses getLaneUses
//
// # Get all laneuses
//
// Responses:
// default: genericError
//
//	200: laneuseDBResponse
func (controller *Controller) GetLaneUses(c *gin.Context) {

	// source slice
	var laneuseDBs []orm.LaneUseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLaneUses", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLaneUse.GetDB()

	_, err := db.Find(&laneuseDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	laneuseAPIs := make([]orm.LaneUseAPI, 0)

	// for each laneuse, update fields from the database nullable fields
	for idx := range laneuseDBs {
		laneuseDB := &laneuseDBs[idx]
		_ = laneuseDB
		var laneuseAPI orm.LaneUseAPI

		// insertion point for updating fields
		laneuseAPI.ID = laneuseDB.ID
		laneuseDB.CopyBasicFieldsToLaneUse_WOP(&laneuseAPI.LaneUse_WOP)
		laneuseAPI.LaneUsePointersEncoding = laneuseDB.LaneUsePointersEncoding
		laneuseAPIs = append(laneuseAPIs, laneuseAPI)
	}

	c.JSON(http.StatusOK, laneuseAPIs)
}

// PostLaneUse
//
// swagger:route POST /laneuses laneuses postLaneUse
//
// Creates a laneuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLaneUse(c *gin.Context) {

	mutexLaneUse.Lock()
	defer mutexLaneUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLaneUses", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLaneUse.GetDB()

	// Validate input
	var input orm.LaneUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create laneuse
	laneuseDB := orm.LaneUseDB{}
	laneuseDB.LaneUsePointersEncoding = input.LaneUsePointersEncoding
	laneuseDB.CopyBasicFieldsFromLaneUse_WOP(&input.LaneUse_WOP)

	_, err = db.Create(&laneuseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLaneUse.CheckoutPhaseOneInstance(&laneuseDB)
	laneuse := backRepo.BackRepoLaneUse.Map_LaneUseDBID_LaneUsePtr[laneuseDB.ID]

	if laneuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), laneuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, laneuseDB)
}

// GetLaneUse
//
// swagger:route GET /laneuses/{ID} laneuses getLaneUse
//
// Gets the details for a laneuse.
//
// Responses:
// default: genericError
//
//	200: laneuseDBResponse
func (controller *Controller) GetLaneUse(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLaneUse", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLaneUse.GetDB()

	// Get laneuseDB in DB
	var laneuseDB orm.LaneUseDB
	if _, err := db.First(&laneuseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var laneuseAPI orm.LaneUseAPI
	laneuseAPI.ID = laneuseDB.ID
	laneuseAPI.LaneUsePointersEncoding = laneuseDB.LaneUsePointersEncoding
	laneuseDB.CopyBasicFieldsToLaneUse_WOP(&laneuseAPI.LaneUse_WOP)

	c.JSON(http.StatusOK, laneuseAPI)
}

// UpdateLaneUse
//
// swagger:route PATCH /laneuses/{ID} laneuses updateLaneUse
//
// # Update a laneuse
//
// Responses:
// default: genericError
//
//	200: laneuseDBResponse
func (controller *Controller) UpdateLaneUse(c *gin.Context) {

	mutexLaneUse.Lock()
	defer mutexLaneUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLaneUse", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLaneUse.GetDB()

	// Validate input
	var input orm.LaneUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var laneuseDB orm.LaneUseDB

	// fetch the laneuse
	_, err := db.First(&laneuseDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	laneuseDB.CopyBasicFieldsFromLaneUse_WOP(&input.LaneUse_WOP)
	laneuseDB.LaneUsePointersEncoding = input.LaneUsePointersEncoding

	db, _ = db.Model(&laneuseDB)
	_, err = db.Updates(&laneuseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	laneuseNew := new(models.LaneUse)
	laneuseDB.CopyBasicFieldsToLaneUse(laneuseNew)

	// redeem pointers
	laneuseDB.DecodePointers(backRepo, laneuseNew)

	// get stage instance from DB instance, and call callback function
	laneuseOld := backRepo.BackRepoLaneUse.Map_LaneUseDBID_LaneUsePtr[laneuseDB.ID]
	if laneuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), laneuseOld, laneuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the laneuseDB
	c.JSON(http.StatusOK, laneuseDB)
}

// DeleteLaneUse
//
// swagger:route DELETE /laneuses/{ID} laneuses deleteLaneUse
//
// # Delete a laneuse
//
// default: genericError
//
//	200: laneuseDBResponse
func (controller *Controller) DeleteLaneUse(c *gin.Context) {

	mutexLaneUse.Lock()
	defer mutexLaneUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLaneUse", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLaneUse.GetDB()

	// Get model if exist
	var laneuseDB orm.LaneUseDB
	if _, err := db.First(&laneuseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&laneuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	laneuseDeleted := new(models.LaneUse)
	laneuseDB.CopyBasicFieldsToLaneUse(laneuseDeleted)

	// get stage instance from DB instance, and call callback function
	laneuseStaged := backRepo.BackRepoLaneUse.Map_LaneUseDBID_LaneUsePtr[laneuseDB.ID]
	if laneuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), laneuseStaged, laneuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
