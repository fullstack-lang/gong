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
var __LinkAnchoredText__dummysDeclaration__ models.LinkAnchoredText
var __LinkAnchoredText_time__dummyDeclaration time.Duration

var mutexLinkAnchoredText sync.Mutex

// An LinkAnchoredTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLinkAnchoredText updateLinkAnchoredText deleteLinkAnchoredText
type LinkAnchoredTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkAnchoredTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLinkAnchoredText updateLinkAnchoredText
type LinkAnchoredTextInput struct {
	// The LinkAnchoredText to submit or modify
	// in: body
	LinkAnchoredText *orm.LinkAnchoredTextAPI
}

// GetLinkAnchoredTexts
//
// swagger:route GET /linkanchoredtexts linkanchoredtexts getLinkAnchoredTexts
//
// # Get all linkanchoredtexts
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) GetLinkAnchoredTexts(c *gin.Context) {

	// source slice
	var linkanchoredtextDBs []orm.LinkAnchoredTextDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	_, err := db.Find(&linkanchoredtextDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linkanchoredtextAPIs := make([]orm.LinkAnchoredTextAPI, 0)

	// for each linkanchoredtext, update fields from the database nullable fields
	for idx := range linkanchoredtextDBs {
		linkanchoredtextDB := &linkanchoredtextDBs[idx]
		_ = linkanchoredtextDB
		var linkanchoredtextAPI orm.LinkAnchoredTextAPI

		// insertion point for updating fields
		linkanchoredtextAPI.ID = linkanchoredtextDB.ID
		linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText_WOP(&linkanchoredtextAPI.LinkAnchoredText_WOP)
		linkanchoredtextAPI.LinkAnchoredTextPointersEncoding = linkanchoredtextDB.LinkAnchoredTextPointersEncoding
		linkanchoredtextAPIs = append(linkanchoredtextAPIs, linkanchoredtextAPI)
	}

	c.JSON(http.StatusOK, linkanchoredtextAPIs)
}

// PostLinkAnchoredText
//
// swagger:route POST /linkanchoredtexts linkanchoredtexts postLinkAnchoredText
//
// Creates a linkanchoredtext
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLinkAnchoredText(c *gin.Context) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLinkAnchoredTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Validate input
	var input orm.LinkAnchoredTextAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create linkanchoredtext
	linkanchoredtextDB := orm.LinkAnchoredTextDB{}
	linkanchoredtextDB.LinkAnchoredTextPointersEncoding = input.LinkAnchoredTextPointersEncoding
	linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText_WOP(&input.LinkAnchoredText_WOP)

	_, err = db.Create(&linkanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLinkAnchoredText.CheckoutPhaseOneInstance(&linkanchoredtextDB)
	linkanchoredtext := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]

	if linkanchoredtext != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), linkanchoredtext)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, linkanchoredtextDB)
}

// GetLinkAnchoredText
//
// swagger:route GET /linkanchoredtexts/{ID} linkanchoredtexts getLinkAnchoredText
//
// Gets the details for a linkanchoredtext.
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) GetLinkAnchoredText(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Get linkanchoredtextDB in DB
	var linkanchoredtextDB orm.LinkAnchoredTextDB
	if _, err := db.First(&linkanchoredtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var linkanchoredtextAPI orm.LinkAnchoredTextAPI
	linkanchoredtextAPI.ID = linkanchoredtextDB.ID
	linkanchoredtextAPI.LinkAnchoredTextPointersEncoding = linkanchoredtextDB.LinkAnchoredTextPointersEncoding
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText_WOP(&linkanchoredtextAPI.LinkAnchoredText_WOP)

	c.JSON(http.StatusOK, linkanchoredtextAPI)
}

// UpdateLinkAnchoredText
//
// swagger:route PATCH /linkanchoredtexts/{ID} linkanchoredtexts updateLinkAnchoredText
//
// # Update a linkanchoredtext
//
// Responses:
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) UpdateLinkAnchoredText(c *gin.Context) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLinkAnchoredText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Validate input
	var input orm.LinkAnchoredTextAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkanchoredtextDB orm.LinkAnchoredTextDB

	// fetch the linkanchoredtext
	_, err := db.First(&linkanchoredtextDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText_WOP(&input.LinkAnchoredText_WOP)
	linkanchoredtextDB.LinkAnchoredTextPointersEncoding = input.LinkAnchoredTextPointersEncoding

	db, _ = db.Model(&linkanchoredtextDB)
	_, err = db.Updates(&linkanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredtextNew := new(models.LinkAnchoredText)
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText(linkanchoredtextNew)

	// redeem pointers
	linkanchoredtextDB.DecodePointers(backRepo, linkanchoredtextNew)

	// get stage instance from DB instance, and call callback function
	linkanchoredtextOld := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	if linkanchoredtextOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), linkanchoredtextOld, linkanchoredtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkanchoredtextDB
	c.JSON(http.StatusOK, linkanchoredtextDB)
}

// DeleteLinkAnchoredText
//
// swagger:route DELETE /linkanchoredtexts/{ID} linkanchoredtexts deleteLinkAnchoredText
//
// # Delete a linkanchoredtext
//
// default: genericError
//
//	200: linkanchoredtextDBResponse
func (controller *Controller) DeleteLinkAnchoredText(c *gin.Context) {

	mutexLinkAnchoredText.Lock()
	defer mutexLinkAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLinkAnchoredText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLinkAnchoredText.GetDB()

	// Get model if exist
	var linkanchoredtextDB orm.LinkAnchoredTextDB
	if _, err := db.First(&linkanchoredtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&linkanchoredtextDB)

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredtextDeleted := new(models.LinkAnchoredText)
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText(linkanchoredtextDeleted)

	// get stage instance from DB instance, and call callback function
	linkanchoredtextStaged := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	if linkanchoredtextStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), linkanchoredtextStaged, linkanchoredtextDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
