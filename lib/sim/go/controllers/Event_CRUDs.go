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
var __Event__dummysDeclaration__ models.Event
var __Event_time__dummyDeclaration time.Duration

var mutexEvent sync.Mutex

// An EventID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEvent updateEvent deleteEvent
type EventID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EventInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEvent updateEvent
type EventInput struct {
	// The Event to submit or modify
	// in: body
	Event *orm.EventAPI
}

// GetEvents
//
// swagger:route GET /events events getEvents
//
// # Get all events
//
// Responses:
// default: genericError
//
//	200: eventDBResponse
func (controller *Controller) GetEvents(c *gin.Context) {

	// source slice
	var eventDBs []orm.EventDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEvents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEvent.GetDB()

	_, err := db.Find(&eventDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	eventAPIs := make([]orm.EventAPI, 0)

	// for each event, update fields from the database nullable fields
	for idx := range eventDBs {
		eventDB := &eventDBs[idx]
		_ = eventDB
		var eventAPI orm.EventAPI

		// insertion point for updating fields
		eventAPI.ID = eventDB.ID
		eventDB.CopyBasicFieldsToEvent_WOP(&eventAPI.Event_WOP)
		eventAPI.EventPointersEncoding = eventDB.EventPointersEncoding
		eventAPIs = append(eventAPIs, eventAPI)
	}

	c.JSON(http.StatusOK, eventAPIs)
}

// PostEvent
//
// swagger:route POST /events events postEvent
//
// Creates a event
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEvent(c *gin.Context) {

	mutexEvent.Lock()
	defer mutexEvent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEvents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEvent.GetDB()

	// Validate input
	var input orm.EventAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create event
	eventDB := orm.EventDB{}
	eventDB.EventPointersEncoding = input.EventPointersEncoding
	eventDB.CopyBasicFieldsFromEvent_WOP(&input.Event_WOP)

	_, err = db.Create(&eventDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEvent.CheckoutPhaseOneInstance(&eventDB)
	event := backRepo.BackRepoEvent.Map_EventDBID_EventPtr[eventDB.ID]

	if event != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), event)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, eventDB)
}

// GetEvent
//
// swagger:route GET /events/{ID} events getEvent
//
// Gets the details for a event.
//
// Responses:
// default: genericError
//
//	200: eventDBResponse
func (controller *Controller) GetEvent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEvent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEvent.GetDB()

	// Get eventDB in DB
	var eventDB orm.EventDB
	if _, err := db.First(&eventDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var eventAPI orm.EventAPI
	eventAPI.ID = eventDB.ID
	eventAPI.EventPointersEncoding = eventDB.EventPointersEncoding
	eventDB.CopyBasicFieldsToEvent_WOP(&eventAPI.Event_WOP)

	c.JSON(http.StatusOK, eventAPI)
}

// UpdateEvent
//
// swagger:route PATCH /events/{ID} events updateEvent
//
// # Update a event
//
// Responses:
// default: genericError
//
//	200: eventDBResponse
func (controller *Controller) UpdateEvent(c *gin.Context) {

	mutexEvent.Lock()
	defer mutexEvent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEvent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEvent.GetDB()

	// Validate input
	var input orm.EventAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var eventDB orm.EventDB

	// fetch the event
	_, err := db.First(&eventDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	eventDB.CopyBasicFieldsFromEvent_WOP(&input.Event_WOP)
	eventDB.EventPointersEncoding = input.EventPointersEncoding

	db, _ = db.Model(&eventDB)
	_, err = db.Updates(&eventDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	eventNew := new(models.Event)
	eventDB.CopyBasicFieldsToEvent(eventNew)

	// redeem pointers
	eventDB.DecodePointers(backRepo, eventNew)

	// get stage instance from DB instance, and call callback function
	eventOld := backRepo.BackRepoEvent.Map_EventDBID_EventPtr[eventDB.ID]
	if eventOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), eventOld, eventNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the eventDB
	c.JSON(http.StatusOK, eventDB)
}

// DeleteEvent
//
// swagger:route DELETE /events/{ID} events deleteEvent
//
// # Delete a event
//
// default: genericError
//
//	200: eventDBResponse
func (controller *Controller) DeleteEvent(c *gin.Context) {

	mutexEvent.Lock()
	defer mutexEvent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEvent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/sim/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoEvent.GetDB()

	// Get model if exist
	var eventDB orm.EventDB
	if _, err := db.First(&eventDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&eventDB)

	// get an instance (not staged) from DB instance, and call callback function
	eventDeleted := new(models.Event)
	eventDB.CopyBasicFieldsToEvent(eventDeleted)

	// get stage instance from DB instance, and call callback function
	eventStaged := backRepo.BackRepoEvent.Map_EventDBID_EventPtr[eventDB.ID]
	if eventStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), eventStaged, eventDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
