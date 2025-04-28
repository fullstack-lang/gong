// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
	"github.com/fullstack-lang/gong/lib/ssg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Content__dummysDeclaration__ models.Content
var __Content_time__dummyDeclaration time.Duration

var mutexContent sync.Mutex

// An ContentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getContent updateContent deleteContent
type ContentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ContentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postContent updateContent
type ContentInput struct {
	// The Content to submit or modify
	// in: body
	Content *orm.ContentAPI
}

// GetContents
//
// swagger:route GET /contents contents getContents
//
// # Get all contents
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) GetContents(c *gin.Context) {

	// source slice
	var contentDBs []orm.ContentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetContents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	_, err := db.Find(&contentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	contentAPIs := make([]orm.ContentAPI, 0)

	// for each content, update fields from the database nullable fields
	for idx := range contentDBs {
		contentDB := &contentDBs[idx]
		_ = contentDB
		var contentAPI orm.ContentAPI

		// insertion point for updating fields
		contentAPI.ID = contentDB.ID
		contentDB.CopyBasicFieldsToContent_WOP(&contentAPI.Content_WOP)
		contentAPI.ContentPointersEncoding = contentDB.ContentPointersEncoding
		contentAPIs = append(contentAPIs, contentAPI)
	}

	c.JSON(http.StatusOK, contentAPIs)
}

// PostContent
//
// swagger:route POST /contents contents postContent
//
// Creates a content
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostContent(c *gin.Context) {

	mutexContent.Lock()
	defer mutexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostContents", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	// Validate input
	var input orm.ContentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create content
	contentDB := orm.ContentDB{}
	contentDB.ContentPointersEncoding = input.ContentPointersEncoding
	contentDB.CopyBasicFieldsFromContent_WOP(&input.Content_WOP)

	_, err = db.Create(&contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoContent.CheckoutPhaseOneInstance(&contentDB)
	content := backRepo.BackRepoContent.Map_ContentDBID_ContentPtr[contentDB.ID]

	if content != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), content)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, contentDB)
}

// GetContent
//
// swagger:route GET /contents/{ID} contents getContent
//
// Gets the details for a content.
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) GetContent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetContent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	// Get contentDB in DB
	var contentDB orm.ContentDB
	if _, err := db.First(&contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var contentAPI orm.ContentAPI
	contentAPI.ID = contentDB.ID
	contentAPI.ContentPointersEncoding = contentDB.ContentPointersEncoding
	contentDB.CopyBasicFieldsToContent_WOP(&contentAPI.Content_WOP)

	c.JSON(http.StatusOK, contentAPI)
}

// UpdateContent
//
// swagger:route PATCH /contents/{ID} contents updateContent
//
// # Update a content
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) UpdateContent(c *gin.Context) {

	mutexContent.Lock()
	defer mutexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateContent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	// Validate input
	var input orm.ContentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var contentDB orm.ContentDB

	// fetch the content
	_, err := db.First(&contentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	contentDB.CopyBasicFieldsFromContent_WOP(&input.Content_WOP)
	contentDB.ContentPointersEncoding = input.ContentPointersEncoding

	db, _ = db.Model(&contentDB)
	_, err = db.Updates(&contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	contentNew := new(models.Content)
	contentDB.CopyBasicFieldsToContent(contentNew)

	// redeem pointers
	contentDB.DecodePointers(backRepo, contentNew)

	// get stage instance from DB instance, and call callback function
	contentOld := backRepo.BackRepoContent.Map_ContentDBID_ContentPtr[contentDB.ID]
	if contentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), contentOld, contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the contentDB
	c.JSON(http.StatusOK, contentDB)
}

// DeleteContent
//
// swagger:route DELETE /contents/{ID} contents deleteContent
//
// # Delete a content
//
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) DeleteContent(c *gin.Context) {

	mutexContent.Lock()
	defer mutexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteContent", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoContent.GetDB()

	// Get model if exist
	var contentDB orm.ContentDB
	if _, err := db.First(&contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&contentDB)

	// get an instance (not staged) from DB instance, and call callback function
	contentDeleted := new(models.Content)
	contentDB.CopyBasicFieldsToContent(contentDeleted)

	// get stage instance from DB instance, and call callback function
	contentStaged := backRepo.BackRepoContent.Map_ContentDBID_ContentPtr[contentDB.ID]
	if contentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), contentStaged, contentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
