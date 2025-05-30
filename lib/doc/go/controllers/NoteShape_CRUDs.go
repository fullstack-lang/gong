// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __NoteShape__dummysDeclaration__ models.NoteShape
var __NoteShape_time__dummyDeclaration time.Duration

var mutexNoteShape sync.Mutex

// An NoteShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNoteShape updateNoteShape deleteNoteShape
type NoteShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNoteShape updateNoteShape
type NoteShapeInput struct {
	// The NoteShape to submit or modify
	// in: body
	NoteShape *orm.NoteShapeAPI
}

// GetNoteShapes
//
// swagger:route GET /noteshapes noteshapes getNoteShapes
//
// # Get all noteshapes
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func (controller *Controller) GetNoteShapes(c *gin.Context) {

	// source slice
	var noteshapeDBs []orm.NoteShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/doc/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNoteShape.GetDB()

	_, err := db.Find(&noteshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteshapeAPIs := make([]orm.NoteShapeAPI, 0)

	// for each noteshape, update fields from the database nullable fields
	for idx := range noteshapeDBs {
		noteshapeDB := &noteshapeDBs[idx]
		_ = noteshapeDB
		var noteshapeAPI orm.NoteShapeAPI

		// insertion point for updating fields
		noteshapeAPI.ID = noteshapeDB.ID
		noteshapeDB.CopyBasicFieldsToNoteShape_WOP(&noteshapeAPI.NoteShape_WOP)
		noteshapeAPI.NoteShapePointersEncoding = noteshapeDB.NoteShapePointersEncoding
		noteshapeAPIs = append(noteshapeAPIs, noteshapeAPI)
	}

	c.JSON(http.StatusOK, noteshapeAPIs)
}

// PostNoteShape
//
// swagger:route POST /noteshapes noteshapes postNoteShape
//
// Creates a noteshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNoteShape(c *gin.Context) {

	mutexNoteShape.Lock()
	defer mutexNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNoteShapes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/doc/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNoteShape.GetDB()

	// Validate input
	var input orm.NoteShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create noteshape
	noteshapeDB := orm.NoteShapeDB{}
	noteshapeDB.NoteShapePointersEncoding = input.NoteShapePointersEncoding
	noteshapeDB.CopyBasicFieldsFromNoteShape_WOP(&input.NoteShape_WOP)

	_, err = db.Create(&noteshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNoteShape.CheckoutPhaseOneInstance(&noteshapeDB)
	noteshape := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]

	if noteshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), noteshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteshapeDB)
}

// GetNoteShape
//
// swagger:route GET /noteshapes/{ID} noteshapes getNoteShape
//
// Gets the details for a noteshape.
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func (controller *Controller) GetNoteShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/doc/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNoteShape.GetDB()

	// Get noteshapeDB in DB
	var noteshapeDB orm.NoteShapeDB
	if _, err := db.First(&noteshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteshapeAPI orm.NoteShapeAPI
	noteshapeAPI.ID = noteshapeDB.ID
	noteshapeAPI.NoteShapePointersEncoding = noteshapeDB.NoteShapePointersEncoding
	noteshapeDB.CopyBasicFieldsToNoteShape_WOP(&noteshapeAPI.NoteShape_WOP)

	c.JSON(http.StatusOK, noteshapeAPI)
}

// UpdateNoteShape
//
// swagger:route PATCH /noteshapes/{ID} noteshapes updateNoteShape
//
// # Update a noteshape
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func (controller *Controller) UpdateNoteShape(c *gin.Context) {

	mutexNoteShape.Lock()
	defer mutexNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNoteShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/doc/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNoteShape.GetDB()

	// Validate input
	var input orm.NoteShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteshapeDB orm.NoteShapeDB

	// fetch the noteshape
	_, err := db.First(&noteshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteshapeDB.CopyBasicFieldsFromNoteShape_WOP(&input.NoteShape_WOP)
	noteshapeDB.NoteShapePointersEncoding = input.NoteShapePointersEncoding

	db, _ = db.Model(&noteshapeDB)
	_, err = db.Updates(&noteshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteshapeNew := new(models.NoteShape)
	noteshapeDB.CopyBasicFieldsToNoteShape(noteshapeNew)

	// redeem pointers
	noteshapeDB.DecodePointers(backRepo, noteshapeNew)

	// get stage instance from DB instance, and call callback function
	noteshapeOld := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]
	if noteshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteshapeOld, noteshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteshapeDB
	c.JSON(http.StatusOK, noteshapeDB)
}

// DeleteNoteShape
//
// swagger:route DELETE /noteshapes/{ID} noteshapes deleteNoteShape
//
// # Delete a noteshape
//
// default: genericError
//
//	200: noteshapeDBResponse
func (controller *Controller) DeleteNoteShape(c *gin.Context) {

	mutexNoteShape.Lock()
	defer mutexNoteShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNoteShape", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/doc/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNoteShape.GetDB()

	// Get model if exist
	var noteshapeDB orm.NoteShapeDB
	if _, err := db.First(&noteshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&noteshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteshapeDeleted := new(models.NoteShape)
	noteshapeDB.CopyBasicFieldsToNoteShape(noteshapeDeleted)

	// get stage instance from DB instance, and call callback function
	noteshapeStaged := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[noteshapeDB.ID]
	if noteshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteshapeStaged, noteshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
