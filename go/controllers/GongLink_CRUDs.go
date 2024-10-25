// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongLink__dummysDeclaration__ models.GongLink
var __GongLink_time__dummyDeclaration time.Duration

var mutexGongLink sync.Mutex

// An GongLinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongLink updateGongLink deleteGongLink
type GongLinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongLinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongLink updateGongLink
type GongLinkInput struct {
	// The GongLink to submit or modify
	// in: body
	GongLink *orm.GongLinkAPI
}

// GetGongLinks
//
// swagger:route GET /gonglinks gonglinks getGongLinks
//
// # Get all gonglinks
//
// Responses:
// default: genericError
//
//	200: gonglinkDBResponse
func (controller *Controller) GetGongLinks(c *gin.Context) {

	// source slice
	var gonglinkDBs []orm.GongLinkDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongLink.GetDB()

	_, err := db.Find(&gonglinkDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gonglinkAPIs := make([]orm.GongLinkAPI, 0)

	// for each gonglink, update fields from the database nullable fields
	for idx := range gonglinkDBs {
		gonglinkDB := &gonglinkDBs[idx]
		_ = gonglinkDB
		var gonglinkAPI orm.GongLinkAPI

		// insertion point for updating fields
		gonglinkAPI.ID = gonglinkDB.ID
		gonglinkDB.CopyBasicFieldsToGongLink_WOP(&gonglinkAPI.GongLink_WOP)
		gonglinkAPI.GongLinkPointersEncoding = gonglinkDB.GongLinkPointersEncoding
		gonglinkAPIs = append(gonglinkAPIs, gonglinkAPI)
	}

	c.JSON(http.StatusOK, gonglinkAPIs)
}

// PostGongLink
//
// swagger:route POST /gonglinks gonglinks postGongLink
//
// Creates a gonglink
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongLink(c *gin.Context) {

	mutexGongLink.Lock()
	defer mutexGongLink.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongLink.GetDB()

	// Validate input
	var input orm.GongLinkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gonglink
	gonglinkDB := orm.GongLinkDB{}
	gonglinkDB.GongLinkPointersEncoding = input.GongLinkPointersEncoding
	gonglinkDB.CopyBasicFieldsFromGongLink_WOP(&input.GongLink_WOP)

	_, err = db.Create(&gonglinkDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongLink.CheckoutPhaseOneInstance(&gonglinkDB)
	gonglink := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkPtr[gonglinkDB.ID]

	if gonglink != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gonglink)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gonglinkDB)
}

// GetGongLink
//
// swagger:route GET /gonglinks/{ID} gonglinks getGongLink
//
// Gets the details for a gonglink.
//
// Responses:
// default: genericError
//
//	200: gonglinkDBResponse
func (controller *Controller) GetGongLink(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongLink.GetDB()

	// Get gonglinkDB in DB
	var gonglinkDB orm.GongLinkDB
	if _, err := db.First(&gonglinkDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gonglinkAPI orm.GongLinkAPI
	gonglinkAPI.ID = gonglinkDB.ID
	gonglinkAPI.GongLinkPointersEncoding = gonglinkDB.GongLinkPointersEncoding
	gonglinkDB.CopyBasicFieldsToGongLink_WOP(&gonglinkAPI.GongLink_WOP)

	c.JSON(http.StatusOK, gonglinkAPI)
}

// UpdateGongLink
//
// swagger:route PATCH /gonglinks/{ID} gonglinks updateGongLink
//
// # Update a gonglink
//
// Responses:
// default: genericError
//
//	200: gonglinkDBResponse
func (controller *Controller) UpdateGongLink(c *gin.Context) {

	mutexGongLink.Lock()
	defer mutexGongLink.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongLink.GetDB()

	// Validate input
	var input orm.GongLinkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gonglinkDB orm.GongLinkDB

	// fetch the gonglink
	_, err := db.First(&gonglinkDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gonglinkDB.CopyBasicFieldsFromGongLink_WOP(&input.GongLink_WOP)
	gonglinkDB.GongLinkPointersEncoding = input.GongLinkPointersEncoding

	db, _ = db.Model(&gonglinkDB)
	_, err = db.Updates(&gonglinkDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gonglinkNew := new(models.GongLink)
	gonglinkDB.CopyBasicFieldsToGongLink(gonglinkNew)

	// redeem pointers
	gonglinkDB.DecodePointers(backRepo, gonglinkNew)

	// get stage instance from DB instance, and call callback function
	gonglinkOld := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkPtr[gonglinkDB.ID]
	if gonglinkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gonglinkOld, gonglinkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gonglinkDB
	c.JSON(http.StatusOK, gonglinkDB)
}

// DeleteGongLink
//
// swagger:route DELETE /gonglinks/{ID} gonglinks deleteGongLink
//
// # Delete a gonglink
//
// default: genericError
//
//	200: gonglinkDBResponse
func (controller *Controller) DeleteGongLink(c *gin.Context) {

	mutexGongLink.Lock()
	defer mutexGongLink.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongLink.GetDB()

	// Get model if exist
	var gonglinkDB orm.GongLinkDB
	if _, err := db.First(&gonglinkDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gonglinkDB)

	// get an instance (not staged) from DB instance, and call callback function
	gonglinkDeleted := new(models.GongLink)
	gonglinkDB.CopyBasicFieldsToGongLink(gonglinkDeleted)

	// get stage instance from DB instance, and call callback function
	gonglinkStaged := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkPtr[gonglinkDB.ID]
	if gonglinkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gonglinkStaged, gonglinkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
