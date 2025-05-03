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
var __Lane__dummysDeclaration__ models.Lane
var __Lane_time__dummyDeclaration time.Duration

var mutexLane sync.Mutex

// An LaneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLane updateLane deleteLane
type LaneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LaneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLane updateLane
type LaneInput struct {
	// The Lane to submit or modify
	// in: body
	Lane *orm.LaneAPI
}

// GetLanes
//
// swagger:route GET /lanes lanes getLanes
//
// # Get all lanes
//
// Responses:
// default: genericError
//
//	200: laneDBResponse
func (controller *Controller) GetLanes(c *gin.Context) {

	// source slice
	var laneDBs []orm.LaneDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLanes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLane.GetDB()

	_, err := db.Find(&laneDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	laneAPIs := make([]orm.LaneAPI, 0)

	// for each lane, update fields from the database nullable fields
	for idx := range laneDBs {
		laneDB := &laneDBs[idx]
		_ = laneDB
		var laneAPI orm.LaneAPI

		// insertion point for updating fields
		laneAPI.ID = laneDB.ID
		laneDB.CopyBasicFieldsToLane_WOP(&laneAPI.Lane_WOP)
		laneAPI.LanePointersEncoding = laneDB.LanePointersEncoding
		laneAPIs = append(laneAPIs, laneAPI)
	}

	c.JSON(http.StatusOK, laneAPIs)
}

// PostLane
//
// swagger:route POST /lanes lanes postLane
//
// Creates a lane
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLane(c *gin.Context) {

	mutexLane.Lock()
	defer mutexLane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLanes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLane.GetDB()

	// Validate input
	var input orm.LaneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lane
	laneDB := orm.LaneDB{}
	laneDB.LanePointersEncoding = input.LanePointersEncoding
	laneDB.CopyBasicFieldsFromLane_WOP(&input.Lane_WOP)

	_, err = db.Create(&laneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLane.CheckoutPhaseOneInstance(&laneDB)
	lane := backRepo.BackRepoLane.Map_LaneDBID_LanePtr[laneDB.ID]

	if lane != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lane)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, laneDB)
}

// GetLane
//
// swagger:route GET /lanes/{ID} lanes getLane
//
// Gets the details for a lane.
//
// Responses:
// default: genericError
//
//	200: laneDBResponse
func (controller *Controller) GetLane(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLane", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLane.GetDB()

	// Get laneDB in DB
	var laneDB orm.LaneDB
	if _, err := db.First(&laneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var laneAPI orm.LaneAPI
	laneAPI.ID = laneDB.ID
	laneAPI.LanePointersEncoding = laneDB.LanePointersEncoding
	laneDB.CopyBasicFieldsToLane_WOP(&laneAPI.Lane_WOP)

	c.JSON(http.StatusOK, laneAPI)
}

// UpdateLane
//
// swagger:route PATCH /lanes/{ID} lanes updateLane
//
// # Update a lane
//
// Responses:
// default: genericError
//
//	200: laneDBResponse
func (controller *Controller) UpdateLane(c *gin.Context) {

	mutexLane.Lock()
	defer mutexLane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLane", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLane.GetDB()

	// Validate input
	var input orm.LaneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var laneDB orm.LaneDB

	// fetch the lane
	_, err := db.First(&laneDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	laneDB.CopyBasicFieldsFromLane_WOP(&input.Lane_WOP)
	laneDB.LanePointersEncoding = input.LanePointersEncoding

	db, _ = db.Model(&laneDB)
	_, err = db.Updates(&laneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	laneNew := new(models.Lane)
	laneDB.CopyBasicFieldsToLane(laneNew)

	// redeem pointers
	laneDB.DecodePointers(backRepo, laneNew)

	// get stage instance from DB instance, and call callback function
	laneOld := backRepo.BackRepoLane.Map_LaneDBID_LanePtr[laneDB.ID]
	if laneOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), laneOld, laneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the laneDB
	c.JSON(http.StatusOK, laneDB)
}

// DeleteLane
//
// swagger:route DELETE /lanes/{ID} lanes deleteLane
//
// # Delete a lane
//
// default: genericError
//
//	200: laneDBResponse
func (controller *Controller) DeleteLane(c *gin.Context) {

	mutexLane.Lock()
	defer mutexLane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLane", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoLane.GetDB()

	// Get model if exist
	var laneDB orm.LaneDB
	if _, err := db.First(&laneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&laneDB)

	// get an instance (not staged) from DB instance, and call callback function
	laneDeleted := new(models.Lane)
	laneDB.CopyBasicFieldsToLane(laneDeleted)

	// get stage instance from DB instance, and call callback function
	laneStaged := backRepo.BackRepoLane.Map_LaneDBID_LanePtr[laneDB.ID]
	if laneStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), laneStaged, laneDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
