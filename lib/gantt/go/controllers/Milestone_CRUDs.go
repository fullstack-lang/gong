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
var __Milestone__dummysDeclaration__ models.Milestone
var __Milestone_time__dummyDeclaration time.Duration

var mutexMilestone sync.Mutex

// An MilestoneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMilestone updateMilestone deleteMilestone
type MilestoneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MilestoneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMilestone updateMilestone
type MilestoneInput struct {
	// The Milestone to submit or modify
	// in: body
	Milestone *orm.MilestoneAPI
}

// GetMilestones
//
// swagger:route GET /milestones milestones getMilestones
//
// # Get all milestones
//
// Responses:
// default: genericError
//
//	200: milestoneDBResponse
func (controller *Controller) GetMilestones(c *gin.Context) {

	// source slice
	var milestoneDBs []orm.MilestoneDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMilestones", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoMilestone.GetDB()

	_, err := db.Find(&milestoneDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	milestoneAPIs := make([]orm.MilestoneAPI, 0)

	// for each milestone, update fields from the database nullable fields
	for idx := range milestoneDBs {
		milestoneDB := &milestoneDBs[idx]
		_ = milestoneDB
		var milestoneAPI orm.MilestoneAPI

		// insertion point for updating fields
		milestoneAPI.ID = milestoneDB.ID
		milestoneDB.CopyBasicFieldsToMilestone_WOP(&milestoneAPI.Milestone_WOP)
		milestoneAPI.MilestonePointersEncoding = milestoneDB.MilestonePointersEncoding
		milestoneAPIs = append(milestoneAPIs, milestoneAPI)
	}

	c.JSON(http.StatusOK, milestoneAPIs)
}

// PostMilestone
//
// swagger:route POST /milestones milestones postMilestone
//
// Creates a milestone
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMilestone(c *gin.Context) {

	mutexMilestone.Lock()
	defer mutexMilestone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMilestones", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoMilestone.GetDB()

	// Validate input
	var input orm.MilestoneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create milestone
	milestoneDB := orm.MilestoneDB{}
	milestoneDB.MilestonePointersEncoding = input.MilestonePointersEncoding
	milestoneDB.CopyBasicFieldsFromMilestone_WOP(&input.Milestone_WOP)

	_, err = db.Create(&milestoneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMilestone.CheckoutPhaseOneInstance(&milestoneDB)
	milestone := backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]

	if milestone != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), milestone)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, milestoneDB)
}

// GetMilestone
//
// swagger:route GET /milestones/{ID} milestones getMilestone
//
// Gets the details for a milestone.
//
// Responses:
// default: genericError
//
//	200: milestoneDBResponse
func (controller *Controller) GetMilestone(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMilestone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoMilestone.GetDB()

	// Get milestoneDB in DB
	var milestoneDB orm.MilestoneDB
	if _, err := db.First(&milestoneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var milestoneAPI orm.MilestoneAPI
	milestoneAPI.ID = milestoneDB.ID
	milestoneAPI.MilestonePointersEncoding = milestoneDB.MilestonePointersEncoding
	milestoneDB.CopyBasicFieldsToMilestone_WOP(&milestoneAPI.Milestone_WOP)

	c.JSON(http.StatusOK, milestoneAPI)
}

// UpdateMilestone
//
// swagger:route PATCH /milestones/{ID} milestones updateMilestone
//
// # Update a milestone
//
// Responses:
// default: genericError
//
//	200: milestoneDBResponse
func (controller *Controller) UpdateMilestone(c *gin.Context) {

	mutexMilestone.Lock()
	defer mutexMilestone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMilestone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoMilestone.GetDB()

	// Validate input
	var input orm.MilestoneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var milestoneDB orm.MilestoneDB

	// fetch the milestone
	_, err := db.First(&milestoneDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	milestoneDB.CopyBasicFieldsFromMilestone_WOP(&input.Milestone_WOP)
	milestoneDB.MilestonePointersEncoding = input.MilestonePointersEncoding

	db, _ = db.Model(&milestoneDB)
	_, err = db.Updates(&milestoneDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	milestoneNew := new(models.Milestone)
	milestoneDB.CopyBasicFieldsToMilestone(milestoneNew)

	// redeem pointers
	milestoneDB.DecodePointers(backRepo, milestoneNew)

	// get stage instance from DB instance, and call callback function
	milestoneOld := backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]
	if milestoneOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), milestoneOld, milestoneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the milestoneDB
	c.JSON(http.StatusOK, milestoneDB)
}

// DeleteMilestone
//
// swagger:route DELETE /milestones/{ID} milestones deleteMilestone
//
// # Delete a milestone
//
// default: genericError
//
//	200: milestoneDBResponse
func (controller *Controller) DeleteMilestone(c *gin.Context) {

	mutexMilestone.Lock()
	defer mutexMilestone.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMilestone", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoMilestone.GetDB()

	// Get model if exist
	var milestoneDB orm.MilestoneDB
	if _, err := db.First(&milestoneDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&milestoneDB)

	// get an instance (not staged) from DB instance, and call callback function
	milestoneDeleted := new(models.Milestone)
	milestoneDB.CopyBasicFieldsToMilestone(milestoneDeleted)

	// get stage instance from DB instance, and call callback function
	milestoneStaged := backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestonePtr[milestoneDB.ID]
	if milestoneStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), milestoneStaged, milestoneDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
