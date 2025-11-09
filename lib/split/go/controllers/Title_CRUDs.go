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
var __Title__dummysDeclaration__ models.Title
var __Title_time__dummyDeclaration time.Duration

var mutexTitle sync.Mutex

// An TitleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTitle updateTitle deleteTitle
type TitleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TitleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTitle updateTitle
type TitleInput struct {
	// The Title to submit or modify
	// in: body
	Title *orm.TitleAPI
}

// GetTitles
//
// swagger:route GET /titles titles getTitles
//
// # Get all titles
//
// Responses:
// default: genericError
//
//	200: titleDBResponse
func (controller *Controller) GetTitles(c *gin.Context) {

	// source slice
	var titleDBs []orm.TitleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTitles", "Name", stackPath)
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
	db := backRepo.BackRepoTitle.GetDB()

	_, err := db.Find(&titleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	titleAPIs := make([]orm.TitleAPI, 0)

	// for each title, update fields from the database nullable fields
	for idx := range titleDBs {
		titleDB := &titleDBs[idx]
		_ = titleDB
		var titleAPI orm.TitleAPI

		// insertion point for updating fields
		titleAPI.ID = titleDB.ID
		titleDB.CopyBasicFieldsToTitle_WOP(&titleAPI.Title_WOP)
		titleAPI.TitlePointersEncoding = titleDB.TitlePointersEncoding
		titleAPIs = append(titleAPIs, titleAPI)
	}

	c.JSON(http.StatusOK, titleAPIs)
}

// PostTitle
//
// swagger:route POST /titles titles postTitle
//
// Creates a title
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTitle(c *gin.Context) {

	mutexTitle.Lock()
	defer mutexTitle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTitles", "Name", stackPath)
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
	db := backRepo.BackRepoTitle.GetDB()

	// Validate input
	var input orm.TitleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create title
	titleDB := orm.TitleDB{}
	titleDB.TitlePointersEncoding = input.TitlePointersEncoding
	titleDB.CopyBasicFieldsFromTitle_WOP(&input.Title_WOP)

	_, err = db.Create(&titleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTitle.CheckoutPhaseOneInstance(&titleDB)
	title := backRepo.BackRepoTitle.Map_TitleDBID_TitlePtr[titleDB.ID]

	if title != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), title)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, titleDB)
}

// GetTitle
//
// swagger:route GET /titles/{ID} titles getTitle
//
// Gets the details for a title.
//
// Responses:
// default: genericError
//
//	200: titleDBResponse
func (controller *Controller) GetTitle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTitle", "Name", stackPath)
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
	db := backRepo.BackRepoTitle.GetDB()

	// Get titleDB in DB
	var titleDB orm.TitleDB
	if _, err := db.First(&titleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var titleAPI orm.TitleAPI
	titleAPI.ID = titleDB.ID
	titleAPI.TitlePointersEncoding = titleDB.TitlePointersEncoding
	titleDB.CopyBasicFieldsToTitle_WOP(&titleAPI.Title_WOP)

	c.JSON(http.StatusOK, titleAPI)
}

// UpdateTitle
//
// swagger:route PATCH /titles/{ID} titles updateTitle
//
// # Update a title
//
// Responses:
// default: genericError
//
//	200: titleDBResponse
func (controller *Controller) UpdateTitle(c *gin.Context) {

	mutexTitle.Lock()
	defer mutexTitle.Unlock()

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
	db := backRepo.BackRepoTitle.GetDB()

	// Validate input
	var input orm.TitleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var titleDB orm.TitleDB

	// fetch the title
	_, err := db.First(&titleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	titleDB.CopyBasicFieldsFromTitle_WOP(&input.Title_WOP)
	titleDB.TitlePointersEncoding = input.TitlePointersEncoding

	db, _ = db.Model(&titleDB)
	_, err = db.Updates(&titleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	titleNew := new(models.Title)
	titleDB.CopyBasicFieldsToTitle(titleNew)

	// redeem pointers
	titleDB.DecodePointers(backRepo, titleNew)

	// get stage instance from DB instance, and call callback function
	titleOld := backRepo.BackRepoTitle.Map_TitleDBID_TitlePtr[titleDB.ID]
	if titleOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), titleOld, titleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the titleDB
	c.JSON(http.StatusOK, titleDB)
}

// DeleteTitle
//
// swagger:route DELETE /titles/{ID} titles deleteTitle
//
// # Delete a title
//
// default: genericError
//
//	200: titleDBResponse
func (controller *Controller) DeleteTitle(c *gin.Context) {

	mutexTitle.Lock()
	defer mutexTitle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTitle", "Name", stackPath)
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
	db := backRepo.BackRepoTitle.GetDB()

	// Get model if exist
	var titleDB orm.TitleDB
	if _, err := db.First(&titleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&titleDB)

	// get an instance (not staged) from DB instance, and call callback function
	titleDeleted := new(models.Title)
	titleDB.CopyBasicFieldsToTitle(titleDeleted)

	// get stage instance from DB instance, and call callback function
	titleStaged := backRepo.BackRepoTitle.Map_TitleDBID_TitlePtr[titleDB.ID]
	if titleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), titleStaged, titleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
