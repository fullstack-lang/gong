// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Diagram__dummysDeclaration__ models.Diagram
var __Diagram_time__dummyDeclaration time.Duration

var mutexDiagram sync.Mutex

// An DiagramID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDiagram updateDiagram deleteDiagram
type DiagramID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DiagramInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDiagram updateDiagram
type DiagramInput struct {
	// The Diagram to submit or modify
	// in: body
	Diagram *orm.DiagramAPI
}

// GetDiagrams
//
// swagger:route GET /diagrams diagrams getDiagrams
//
// # Get all diagrams
//
// Responses:
// default: genericError
//
//	200: diagramDBResponse
func (controller *Controller) GetDiagrams(c *gin.Context) {

	// source slice
	var diagramDBs []orm.DiagramDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDiagrams", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDiagram.GetDB()

	_, err := db.Find(&diagramDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	diagramAPIs := make([]orm.DiagramAPI, 0)

	// for each diagram, update fields from the database nullable fields
	for idx := range diagramDBs {
		diagramDB := &diagramDBs[idx]
		_ = diagramDB
		var diagramAPI orm.DiagramAPI

		// insertion point for updating fields
		diagramAPI.ID = diagramDB.ID
		diagramDB.CopyBasicFieldsToDiagram_WOP(&diagramAPI.Diagram_WOP)
		diagramAPI.DiagramPointersEncoding = diagramDB.DiagramPointersEncoding
		diagramAPIs = append(diagramAPIs, diagramAPI)
	}

	c.JSON(http.StatusOK, diagramAPIs)
}

// PostDiagram
//
// swagger:route POST /diagrams diagrams postDiagram
//
// Creates a diagram
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDiagram(c *gin.Context) {

	mutexDiagram.Lock()
	defer mutexDiagram.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDiagrams", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDiagram.GetDB()

	// Validate input
	var input orm.DiagramAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create diagram
	diagramDB := orm.DiagramDB{}
	diagramDB.DiagramPointersEncoding = input.DiagramPointersEncoding
	diagramDB.CopyBasicFieldsFromDiagram_WOP(&input.Diagram_WOP)

	_, err = db.Create(&diagramDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDiagram.CheckoutPhaseOneInstance(&diagramDB)
	diagram := backRepo.BackRepoDiagram.Map_DiagramDBID_DiagramPtr[diagramDB.ID]

	if diagram != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), diagram)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, diagramDB)
}

// GetDiagram
//
// swagger:route GET /diagrams/{ID} diagrams getDiagram
//
// Gets the details for a diagram.
//
// Responses:
// default: genericError
//
//	200: diagramDBResponse
func (controller *Controller) GetDiagram(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDiagram", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDiagram.GetDB()

	// Get diagramDB in DB
	var diagramDB orm.DiagramDB
	if _, err := db.First(&diagramDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var diagramAPI orm.DiagramAPI
	diagramAPI.ID = diagramDB.ID
	diagramAPI.DiagramPointersEncoding = diagramDB.DiagramPointersEncoding
	diagramDB.CopyBasicFieldsToDiagram_WOP(&diagramAPI.Diagram_WOP)

	c.JSON(http.StatusOK, diagramAPI)
}

// UpdateDiagram
//
// swagger:route PATCH /diagrams/{ID} diagrams updateDiagram
//
// # Update a diagram
//
// Responses:
// default: genericError
//
//	200: diagramDBResponse
func (controller *Controller) UpdateDiagram(c *gin.Context) {

	mutexDiagram.Lock()
	defer mutexDiagram.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDiagram.GetDB()

	// Validate input
	var input orm.DiagramAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var diagramDB orm.DiagramDB

	// fetch the diagram
	_, err := db.First(&diagramDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	diagramDB.CopyBasicFieldsFromDiagram_WOP(&input.Diagram_WOP)
	diagramDB.DiagramPointersEncoding = input.DiagramPointersEncoding

	db, _ = db.Model(&diagramDB)
	_, err = db.Updates(&diagramDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	diagramNew := new(models.Diagram)
	diagramDB.CopyBasicFieldsToDiagram(diagramNew)

	// redeem pointers
	diagramDB.DecodePointers(backRepo, diagramNew)

	// get stage instance from DB instance, and call callback function
	diagramOld := backRepo.BackRepoDiagram.Map_DiagramDBID_DiagramPtr[diagramDB.ID]
	if diagramOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), diagramOld, diagramNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the diagramDB
	c.JSON(http.StatusOK, diagramDB)
}

// DeleteDiagram
//
// swagger:route DELETE /diagrams/{ID} diagrams deleteDiagram
//
// # Delete a diagram
//
// default: genericError
//
//	200: diagramDBResponse
func (controller *Controller) DeleteDiagram(c *gin.Context) {

	mutexDiagram.Lock()
	defer mutexDiagram.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDiagram", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoDiagram.GetDB()

	// Get model if exist
	var diagramDB orm.DiagramDB
	if _, err := db.First(&diagramDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&diagramDB)

	// get an instance (not staged) from DB instance, and call callback function
	diagramDeleted := new(models.Diagram)
	diagramDB.CopyBasicFieldsToDiagram(diagramDeleted)

	// get stage instance from DB instance, and call callback function
	diagramStaged := backRepo.BackRepoDiagram.Map_DiagramDBID_DiagramPtr[diagramDB.ID]
	if diagramStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), diagramStaged, diagramDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
