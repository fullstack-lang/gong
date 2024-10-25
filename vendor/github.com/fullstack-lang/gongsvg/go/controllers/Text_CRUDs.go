// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Text__dummysDeclaration__ models.Text
var __Text_time__dummyDeclaration time.Duration

var mutexText sync.Mutex

// An TextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getText updateText deleteText
type TextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postText updateText
type TextInput struct {
	// The Text to submit or modify
	// in: body
	Text *orm.TextAPI
}

// GetTexts
//
// swagger:route GET /texts texts getTexts
//
// # Get all texts
//
// Responses:
// default: genericError
//
//	200: textDBResponse
func (controller *Controller) GetTexts(c *gin.Context) {

	// source slice
	var textDBs []orm.TextDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText.GetDB()

	_, err := db.Find(&textDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	textAPIs := make([]orm.TextAPI, 0)

	// for each text, update fields from the database nullable fields
	for idx := range textDBs {
		textDB := &textDBs[idx]
		_ = textDB
		var textAPI orm.TextAPI

		// insertion point for updating fields
		textAPI.ID = textDB.ID
		textDB.CopyBasicFieldsToText_WOP(&textAPI.Text_WOP)
		textAPI.TextPointersEncoding = textDB.TextPointersEncoding
		textAPIs = append(textAPIs, textAPI)
	}

	c.JSON(http.StatusOK, textAPIs)
}

// PostText
//
// swagger:route POST /texts texts postText
//
// Creates a text
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostText(c *gin.Context) {

	mutexText.Lock()
	defer mutexText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText.GetDB()

	// Validate input
	var input orm.TextAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create text
	textDB := orm.TextDB{}
	textDB.TextPointersEncoding = input.TextPointersEncoding
	textDB.CopyBasicFieldsFromText_WOP(&input.Text_WOP)

	_, err = db.Create(&textDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoText.CheckoutPhaseOneInstance(&textDB)
	text := backRepo.BackRepoText.Map_TextDBID_TextPtr[textDB.ID]

	if text != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), text)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, textDB)
}

// GetText
//
// swagger:route GET /texts/{ID} texts getText
//
// Gets the details for a text.
//
// Responses:
// default: genericError
//
//	200: textDBResponse
func (controller *Controller) GetText(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText.GetDB()

	// Get textDB in DB
	var textDB orm.TextDB
	if _, err := db.First(&textDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var textAPI orm.TextAPI
	textAPI.ID = textDB.ID
	textAPI.TextPointersEncoding = textDB.TextPointersEncoding
	textDB.CopyBasicFieldsToText_WOP(&textAPI.Text_WOP)

	c.JSON(http.StatusOK, textAPI)
}

// UpdateText
//
// swagger:route PATCH /texts/{ID} texts updateText
//
// # Update a text
//
// Responses:
// default: genericError
//
//	200: textDBResponse
func (controller *Controller) UpdateText(c *gin.Context) {

	mutexText.Lock()
	defer mutexText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText.GetDB()

	// Validate input
	var input orm.TextAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var textDB orm.TextDB

	// fetch the text
	_, err := db.First(&textDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	textDB.CopyBasicFieldsFromText_WOP(&input.Text_WOP)
	textDB.TextPointersEncoding = input.TextPointersEncoding

	db, _ = db.Model(&textDB)
	_, err = db.Updates(&textDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	textNew := new(models.Text)
	textDB.CopyBasicFieldsToText(textNew)

	// redeem pointers
	textDB.DecodePointers(backRepo, textNew)

	// get stage instance from DB instance, and call callback function
	textOld := backRepo.BackRepoText.Map_TextDBID_TextPtr[textDB.ID]
	if textOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), textOld, textNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the textDB
	c.JSON(http.StatusOK, textDB)
}

// DeleteText
//
// swagger:route DELETE /texts/{ID} texts deleteText
//
// # Delete a text
//
// default: genericError
//
//	200: textDBResponse
func (controller *Controller) DeleteText(c *gin.Context) {

	mutexText.Lock()
	defer mutexText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText.GetDB()

	// Get model if exist
	var textDB orm.TextDB
	if _, err := db.First(&textDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&textDB)

	// get an instance (not staged) from DB instance, and call callback function
	textDeleted := new(models.Text)
	textDB.CopyBasicFieldsToText(textDeleted)

	// get stage instance from DB instance, and call callback function
	textStaged := backRepo.BackRepoText.Map_TextDBID_TextPtr[textDB.ID]
	if textStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), textStaged, textDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
