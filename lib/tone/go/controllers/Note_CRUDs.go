// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Note__dummysDeclaration__ models.Note
var __Note_time__dummyDeclaration time.Duration

var mutexNote sync.Mutex

// An NoteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNote updateNote deleteNote
type NoteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNote updateNote
type NoteInput struct {
	// The Note to submit or modify
	// in: body
	Note *orm.NoteAPI
}

// GetNotes
//
// swagger:route GET /notes notes getNotes
//
// # Get all notes
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) GetNotes(c *gin.Context) {

	// source slice
	var noteDBs []orm.NoteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	_, err := db.Find(&noteDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteAPIs := make([]orm.NoteAPI, 0)

	// for each note, update fields from the database nullable fields
	for idx := range noteDBs {
		noteDB := &noteDBs[idx]
		_ = noteDB
		var noteAPI orm.NoteAPI

		// insertion point for updating fields
		noteAPI.ID = noteDB.ID
		noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)
		noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
		noteAPIs = append(noteAPIs, noteAPI)
	}

	c.JSON(http.StatusOK, noteAPIs)
}

// PostNote
//
// swagger:route POST /notes notes postNote
//
// Creates a note
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNotes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Validate input
	var input orm.NoteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create note
	noteDB := orm.NoteDB{}
	noteDB.NotePointersEncoding = input.NotePointersEncoding
	noteDB.CopyBasicFieldsFromNote_WOP(&input.Note_WOP)

	_, err = db.Create(&noteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNote.CheckoutPhaseOneInstance(&noteDB)
	note := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]

	if note != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), note)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteDB)
}

// GetNote
//
// swagger:route GET /notes/{ID} notes getNote
//
// Gets the details for a note.
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) GetNote(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Get noteDB in DB
	var noteDB orm.NoteDB
	if _, err := db.First(&noteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteAPI orm.NoteAPI
	noteAPI.ID = noteDB.ID
	noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
	noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)

	c.JSON(http.StatusOK, noteAPI)
}

// UpdateNote
//
// swagger:route PATCH /notes/{ID} notes updateNote
//
// # Update a note
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) UpdateNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Validate input
	var input orm.NoteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteDB orm.NoteDB

	// fetch the note
	_, err := db.First(&noteDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteDB.CopyBasicFieldsFromNote_WOP(&input.Note_WOP)
	noteDB.NotePointersEncoding = input.NotePointersEncoding

	db, _ = db.Model(&noteDB)
	_, err = db.Updates(&noteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteNew := new(models.Note)
	noteDB.CopyBasicFieldsToNote(noteNew)

	// redeem pointers
	noteDB.DecodePointers(backRepo, noteNew)

	// get stage instance from DB instance, and call callback function
	noteOld := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]
	if noteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteOld, noteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteDB
	c.JSON(http.StatusOK, noteDB)
}

// DeleteNote
//
// swagger:route DELETE /notes/{ID} notes deleteNote
//
// # Delete a note
//
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) DeleteNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Get model if exist
	var noteDB orm.NoteDB
	if _, err := db.First(&noteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&noteDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteDeleted := new(models.Note)
	noteDB.CopyBasicFieldsToNote(noteDeleted)

	// get stage instance from DB instance, and call callback function
	noteStaged := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]
	if noteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteStaged, noteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
