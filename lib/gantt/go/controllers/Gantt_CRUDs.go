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
var __Gantt__dummysDeclaration__ models.Gantt
var __Gantt_time__dummyDeclaration time.Duration

var mutexGantt sync.Mutex

// An GanttID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGantt updateGantt deleteGantt
type GanttID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GanttInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGantt updateGantt
type GanttInput struct {
	// The Gantt to submit or modify
	// in: body
	Gantt *orm.GanttAPI
}

// GetGantts
//
// swagger:route GET /gantts gantts getGantts
//
// # Get all gantts
//
// Responses:
// default: genericError
//
//	200: ganttDBResponse
func (controller *Controller) GetGantts(c *gin.Context) {

	// source slice
	var ganttDBs []orm.GanttDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGantts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGantt.GetDB()

	_, err := db.Find(&ganttDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ganttAPIs := make([]orm.GanttAPI, 0)

	// for each gantt, update fields from the database nullable fields
	for idx := range ganttDBs {
		ganttDB := &ganttDBs[idx]
		_ = ganttDB
		var ganttAPI orm.GanttAPI

		// insertion point for updating fields
		ganttAPI.ID = ganttDB.ID
		ganttDB.CopyBasicFieldsToGantt_WOP(&ganttAPI.Gantt_WOP)
		ganttAPI.GanttPointersEncoding = ganttDB.GanttPointersEncoding
		ganttAPIs = append(ganttAPIs, ganttAPI)
	}

	c.JSON(http.StatusOK, ganttAPIs)
}

// PostGantt
//
// swagger:route POST /gantts gantts postGantt
//
// Creates a gantt
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGantt(c *gin.Context) {

	mutexGantt.Lock()
	defer mutexGantt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGantts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGantt.GetDB()

	// Validate input
	var input orm.GanttAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gantt
	ganttDB := orm.GanttDB{}
	ganttDB.GanttPointersEncoding = input.GanttPointersEncoding
	ganttDB.CopyBasicFieldsFromGantt_WOP(&input.Gantt_WOP)

	_, err = db.Create(&ganttDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGantt.CheckoutPhaseOneInstance(&ganttDB)
	gantt := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[ganttDB.ID]

	if gantt != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gantt)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, ganttDB)
}

// GetGantt
//
// swagger:route GET /gantts/{ID} gantts getGantt
//
// Gets the details for a gantt.
//
// Responses:
// default: genericError
//
//	200: ganttDBResponse
func (controller *Controller) GetGantt(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGantt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGantt.GetDB()

	// Get ganttDB in DB
	var ganttDB orm.GanttDB
	if _, err := db.First(&ganttDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var ganttAPI orm.GanttAPI
	ganttAPI.ID = ganttDB.ID
	ganttAPI.GanttPointersEncoding = ganttDB.GanttPointersEncoding
	ganttDB.CopyBasicFieldsToGantt_WOP(&ganttAPI.Gantt_WOP)

	c.JSON(http.StatusOK, ganttAPI)
}

// UpdateGantt
//
// swagger:route PATCH /gantts/{ID} gantts updateGantt
//
// # Update a gantt
//
// Responses:
// default: genericError
//
//	200: ganttDBResponse
func (controller *Controller) UpdateGantt(c *gin.Context) {

	mutexGantt.Lock()
	defer mutexGantt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGantt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGantt.GetDB()

	// Validate input
	var input orm.GanttAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ganttDB orm.GanttDB

	// fetch the gantt
	_, err := db.First(&ganttDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ganttDB.CopyBasicFieldsFromGantt_WOP(&input.Gantt_WOP)
	ganttDB.GanttPointersEncoding = input.GanttPointersEncoding

	db, _ = db.Model(&ganttDB)
	_, err = db.Updates(&ganttDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ganttNew := new(models.Gantt)
	ganttDB.CopyBasicFieldsToGantt(ganttNew)

	// redeem pointers
	ganttDB.DecodePointers(backRepo, ganttNew)

	// get stage instance from DB instance, and call callback function
	ganttOld := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[ganttDB.ID]
	if ganttOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), ganttOld, ganttNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ganttDB
	c.JSON(http.StatusOK, ganttDB)
}

// DeleteGantt
//
// swagger:route DELETE /gantts/{ID} gantts deleteGantt
//
// # Delete a gantt
//
// default: genericError
//
//	200: ganttDBResponse
func (controller *Controller) DeleteGantt(c *gin.Context) {

	mutexGantt.Lock()
	defer mutexGantt.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGantt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoGantt.GetDB()

	// Get model if exist
	var ganttDB orm.GanttDB
	if _, err := db.First(&ganttDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&ganttDB)

	// get an instance (not staged) from DB instance, and call callback function
	ganttDeleted := new(models.Gantt)
	ganttDB.CopyBasicFieldsToGantt(ganttDeleted)

	// get stage instance from DB instance, and call callback function
	ganttStaged := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[ganttDB.ID]
	if ganttStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), ganttStaged, ganttDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
