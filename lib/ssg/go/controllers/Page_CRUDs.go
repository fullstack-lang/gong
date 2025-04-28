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
var __Page__dummysDeclaration__ models.Page
var __Page_time__dummyDeclaration time.Duration

var mutexPage sync.Mutex

// An PageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPage updatePage deletePage
type PageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPage updatePage
type PageInput struct {
	// The Page to submit or modify
	// in: body
	Page *orm.PageAPI
}

// GetPages
//
// swagger:route GET /pages pages getPages
//
// # Get all pages
//
// Responses:
// default: genericError
//
//	200: pageDBResponse
func (controller *Controller) GetPages(c *gin.Context) {

	// source slice
	var pageDBs []orm.PageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPages", "Name", stackPath)
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
	db := backRepo.BackRepoPage.GetDB()

	_, err := db.Find(&pageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pageAPIs := make([]orm.PageAPI, 0)

	// for each page, update fields from the database nullable fields
	for idx := range pageDBs {
		pageDB := &pageDBs[idx]
		_ = pageDB
		var pageAPI orm.PageAPI

		// insertion point for updating fields
		pageAPI.ID = pageDB.ID
		pageDB.CopyBasicFieldsToPage_WOP(&pageAPI.Page_WOP)
		pageAPI.PagePointersEncoding = pageDB.PagePointersEncoding
		pageAPIs = append(pageAPIs, pageAPI)
	}

	c.JSON(http.StatusOK, pageAPIs)
}

// PostPage
//
// swagger:route POST /pages pages postPage
//
// Creates a page
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPage(c *gin.Context) {

	mutexPage.Lock()
	defer mutexPage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPages", "Name", stackPath)
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
	db := backRepo.BackRepoPage.GetDB()

	// Validate input
	var input orm.PageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create page
	pageDB := orm.PageDB{}
	pageDB.PagePointersEncoding = input.PagePointersEncoding
	pageDB.CopyBasicFieldsFromPage_WOP(&input.Page_WOP)

	_, err = db.Create(&pageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPage.CheckoutPhaseOneInstance(&pageDB)
	page := backRepo.BackRepoPage.Map_PageDBID_PagePtr[pageDB.ID]

	if page != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), page)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pageDB)
}

// GetPage
//
// swagger:route GET /pages/{ID} pages getPage
//
// Gets the details for a page.
//
// Responses:
// default: genericError
//
//	200: pageDBResponse
func (controller *Controller) GetPage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPage", "Name", stackPath)
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
	db := backRepo.BackRepoPage.GetDB()

	// Get pageDB in DB
	var pageDB orm.PageDB
	if _, err := db.First(&pageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pageAPI orm.PageAPI
	pageAPI.ID = pageDB.ID
	pageAPI.PagePointersEncoding = pageDB.PagePointersEncoding
	pageDB.CopyBasicFieldsToPage_WOP(&pageAPI.Page_WOP)

	c.JSON(http.StatusOK, pageAPI)
}

// UpdatePage
//
// swagger:route PATCH /pages/{ID} pages updatePage
//
// # Update a page
//
// Responses:
// default: genericError
//
//	200: pageDBResponse
func (controller *Controller) UpdatePage(c *gin.Context) {

	mutexPage.Lock()
	defer mutexPage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePage", "Name", stackPath)
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
	db := backRepo.BackRepoPage.GetDB()

	// Validate input
	var input orm.PageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pageDB orm.PageDB

	// fetch the page
	_, err := db.First(&pageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pageDB.CopyBasicFieldsFromPage_WOP(&input.Page_WOP)
	pageDB.PagePointersEncoding = input.PagePointersEncoding

	db, _ = db.Model(&pageDB)
	_, err = db.Updates(&pageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pageNew := new(models.Page)
	pageDB.CopyBasicFieldsToPage(pageNew)

	// redeem pointers
	pageDB.DecodePointers(backRepo, pageNew)

	// get stage instance from DB instance, and call callback function
	pageOld := backRepo.BackRepoPage.Map_PageDBID_PagePtr[pageDB.ID]
	if pageOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pageOld, pageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pageDB
	c.JSON(http.StatusOK, pageDB)
}

// DeletePage
//
// swagger:route DELETE /pages/{ID} pages deletePage
//
// # Delete a page
//
// default: genericError
//
//	200: pageDBResponse
func (controller *Controller) DeletePage(c *gin.Context) {

	mutexPage.Lock()
	defer mutexPage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePage", "Name", stackPath)
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
	db := backRepo.BackRepoPage.GetDB()

	// Get model if exist
	var pageDB orm.PageDB
	if _, err := db.First(&pageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&pageDB)

	// get an instance (not staged) from DB instance, and call callback function
	pageDeleted := new(models.Page)
	pageDB.CopyBasicFieldsToPage(pageDeleted)

	// get stage instance from DB instance, and call callback function
	pageStaged := backRepo.BackRepoPage.Map_PageDBID_PagePtr[pageDB.ID]
	if pageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pageStaged, pageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
