// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Markdown__dummysDeclaration__ models.Markdown
var _ = __Markdown__dummysDeclaration__
var __Markdown_time__dummyDeclaration time.Duration
var _ = __Markdown_time__dummyDeclaration

var mutexMarkdown sync.Mutex

// An MarkdownID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMarkdown updateMarkdown deleteMarkdown
type MarkdownID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MarkdownInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMarkdown updateMarkdown
type MarkdownInput struct {
	// The Markdown to submit or modify
	// in: body
	Markdown *orm.MarkdownAPI
}

// GetMarkdowns
//
// swagger:route GET /markdowns markdowns getMarkdowns
//
// # Get all markdowns
//
// Responses:
// default: genericError
//
//	200: markdownDBResponse
func (controller *Controller) GetMarkdowns(c *gin.Context) {

	// source slice
	var markdownDBs []orm.MarkdownDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarkdowns", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	_, err := db.Find(&markdownDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	markdownAPIs := make([]orm.MarkdownAPI, 0)

	// for each markdown, update fields from the database nullable fields
	for idx := range markdownDBs {
		markdownDB := &markdownDBs[idx]
		_ = markdownDB
		var markdownAPI orm.MarkdownAPI

		// insertion point for updating fields
		markdownAPI.ID = markdownDB.ID
		markdownDB.CopyBasicFieldsToMarkdown_WOP(&markdownAPI.Markdown_WOP)
		markdownAPI.MarkdownPointersEncoding = markdownDB.MarkdownPointersEncoding
		markdownAPIs = append(markdownAPIs, markdownAPI)
	}

	c.JSON(http.StatusOK, markdownAPIs)
}

// PostMarkdown
//
// swagger:route POST /markdowns markdowns postMarkdown
//
// Creates a markdown
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMarkdown(c *gin.Context) {

	mutexMarkdown.Lock()
	defer mutexMarkdown.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMarkdowns", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	// Validate input
	var input orm.MarkdownAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create markdown
	markdownDB := orm.MarkdownDB{}
	markdownDB.MarkdownPointersEncoding = input.MarkdownPointersEncoding
	markdownDB.CopyBasicFieldsFromMarkdown_WOP(&input.Markdown_WOP)

	_, err = db.Create(&markdownDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMarkdown.CheckoutPhaseOneInstance(&markdownDB)
	markdown := backRepo.BackRepoMarkdown.Map_MarkdownDBID_MarkdownPtr[markdownDB.ID]

	if markdown != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), markdown)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, markdownDB)
}

// GetMarkdown
//
// swagger:route GET /markdowns/{ID} markdowns getMarkdown
//
// Gets the details for a markdown.
//
// Responses:
// default: genericError
//
//	200: markdownDBResponse
func (controller *Controller) GetMarkdown(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarkdown", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	// Get markdownDB in DB
	var markdownDB orm.MarkdownDB
	if _, err := db.First(&markdownDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var markdownAPI orm.MarkdownAPI
	markdownAPI.ID = markdownDB.ID
	markdownAPI.MarkdownPointersEncoding = markdownDB.MarkdownPointersEncoding
	markdownDB.CopyBasicFieldsToMarkdown_WOP(&markdownAPI.Markdown_WOP)

	c.JSON(http.StatusOK, markdownAPI)
}

// UpdateMarkdown
//
// swagger:route PATCH /markdowns/{ID} markdowns updateMarkdown
//
// # Update a markdown
//
// Responses:
// default: genericError
//
//	200: markdownDBResponse
func (controller *Controller) UpdateMarkdown(c *gin.Context) {

	mutexMarkdown.Lock()
	defer mutexMarkdown.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	// Validate input
	var input orm.MarkdownAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var markdownDB orm.MarkdownDB

	// fetch the markdown
	_, err := db.First(&markdownDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	markdownDB.CopyBasicFieldsFromMarkdown_WOP(&input.Markdown_WOP)
	markdownDB.MarkdownPointersEncoding = input.MarkdownPointersEncoding

	db, _ = db.Model(&markdownDB)
	_, err = db.Updates(&markdownDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	markdownNew := new(models.Markdown)
	markdownDB.CopyBasicFieldsToMarkdown(markdownNew)

	// redeem pointers
	markdownDB.DecodePointers(backRepo, markdownNew)

	// get stage instance from DB instance, and call callback function
	markdownOld := backRepo.BackRepoMarkdown.Map_MarkdownDBID_MarkdownPtr[markdownDB.ID]
	if markdownOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), markdownOld, markdownNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the markdownDB
	c.JSON(http.StatusOK, markdownDB)
}

// DeleteMarkdown
//
// swagger:route DELETE /markdowns/{ID} markdowns deleteMarkdown
//
// # Delete a markdown
//
// default: genericError
//
//	200: markdownDBResponse
func (controller *Controller) DeleteMarkdown(c *gin.Context) {

	mutexMarkdown.Lock()
	defer mutexMarkdown.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMarkdown", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMarkdown.GetDB()

	// Get model if exist
	var markdownDB orm.MarkdownDB
	if _, err := db.First(&markdownDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&markdownDB)

	// get an instance (not staged) from DB instance, and call callback function
	markdownDeleted := new(models.Markdown)
	markdownDB.CopyBasicFieldsToMarkdown(markdownDeleted)

	// get stage instance from DB instance, and call callback function
	markdownStaged := backRepo.BackRepoMarkdown.Map_MarkdownDBID_MarkdownPtr[markdownDB.ID]
	if markdownStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), markdownStaged, markdownDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
