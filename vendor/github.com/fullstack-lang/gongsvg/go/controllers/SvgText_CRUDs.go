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
var __SvgText__dummysDeclaration__ models.SvgText
var __SvgText_time__dummyDeclaration time.Duration

var mutexSvgText sync.Mutex

// An SvgTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSvgText updateSvgText deleteSvgText
type SvgTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SvgTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSvgText updateSvgText
type SvgTextInput struct {
	// The SvgText to submit or modify
	// in: body
	SvgText *orm.SvgTextAPI
}

// GetSvgTexts
//
// swagger:route GET /svgtexts svgtexts getSvgTexts
//
// # Get all svgtexts
//
// Responses:
// default: genericError
//
//	200: svgtextDBResponse
func (controller *Controller) GetSvgTexts(c *gin.Context) {

	// source slice
	var svgtextDBs []orm.SvgTextDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvgTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSvgText.GetDB()

	_, err := db.Find(&svgtextDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	svgtextAPIs := make([]orm.SvgTextAPI, 0)

	// for each svgtext, update fields from the database nullable fields
	for idx := range svgtextDBs {
		svgtextDB := &svgtextDBs[idx]
		_ = svgtextDB
		var svgtextAPI orm.SvgTextAPI

		// insertion point for updating fields
		svgtextAPI.ID = svgtextDB.ID
		svgtextDB.CopyBasicFieldsToSvgText_WOP(&svgtextAPI.SvgText_WOP)
		svgtextAPI.SvgTextPointersEncoding = svgtextDB.SvgTextPointersEncoding
		svgtextAPIs = append(svgtextAPIs, svgtextAPI)
	}

	c.JSON(http.StatusOK, svgtextAPIs)
}

// PostSvgText
//
// swagger:route POST /svgtexts svgtexts postSvgText
//
// Creates a svgtext
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSvgText(c *gin.Context) {

	mutexSvgText.Lock()
	defer mutexSvgText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSvgTexts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSvgText.GetDB()

	// Validate input
	var input orm.SvgTextAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create svgtext
	svgtextDB := orm.SvgTextDB{}
	svgtextDB.SvgTextPointersEncoding = input.SvgTextPointersEncoding
	svgtextDB.CopyBasicFieldsFromSvgText_WOP(&input.SvgText_WOP)

	_, err = db.Create(&svgtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSvgText.CheckoutPhaseOneInstance(&svgtextDB)
	svgtext := backRepo.BackRepoSvgText.Map_SvgTextDBID_SvgTextPtr[svgtextDB.ID]

	if svgtext != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), svgtext)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, svgtextDB)
}

// GetSvgText
//
// swagger:route GET /svgtexts/{ID} svgtexts getSvgText
//
// Gets the details for a svgtext.
//
// Responses:
// default: genericError
//
//	200: svgtextDBResponse
func (controller *Controller) GetSvgText(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvgText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSvgText.GetDB()

	// Get svgtextDB in DB
	var svgtextDB orm.SvgTextDB
	if _, err := db.First(&svgtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var svgtextAPI orm.SvgTextAPI
	svgtextAPI.ID = svgtextDB.ID
	svgtextAPI.SvgTextPointersEncoding = svgtextDB.SvgTextPointersEncoding
	svgtextDB.CopyBasicFieldsToSvgText_WOP(&svgtextAPI.SvgText_WOP)

	c.JSON(http.StatusOK, svgtextAPI)
}

// UpdateSvgText
//
// swagger:route PATCH /svgtexts/{ID} svgtexts updateSvgText
//
// # Update a svgtext
//
// Responses:
// default: genericError
//
//	200: svgtextDBResponse
func (controller *Controller) UpdateSvgText(c *gin.Context) {

	mutexSvgText.Lock()
	defer mutexSvgText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSvgText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSvgText.GetDB()

	// Validate input
	var input orm.SvgTextAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgtextDB orm.SvgTextDB

	// fetch the svgtext
	_, err := db.First(&svgtextDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgtextDB.CopyBasicFieldsFromSvgText_WOP(&input.SvgText_WOP)
	svgtextDB.SvgTextPointersEncoding = input.SvgTextPointersEncoding

	db, _ = db.Model(&svgtextDB)
	_, err = db.Updates(&svgtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgtextNew := new(models.SvgText)
	svgtextDB.CopyBasicFieldsToSvgText(svgtextNew)

	// redeem pointers
	svgtextDB.DecodePointers(backRepo, svgtextNew)

	// get stage instance from DB instance, and call callback function
	svgtextOld := backRepo.BackRepoSvgText.Map_SvgTextDBID_SvgTextPtr[svgtextDB.ID]
	if svgtextOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), svgtextOld, svgtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgtextDB
	c.JSON(http.StatusOK, svgtextDB)
}

// DeleteSvgText
//
// swagger:route DELETE /svgtexts/{ID} svgtexts deleteSvgText
//
// # Delete a svgtext
//
// default: genericError
//
//	200: svgtextDBResponse
func (controller *Controller) DeleteSvgText(c *gin.Context) {

	mutexSvgText.Lock()
	defer mutexSvgText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSvgText", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSvgText.GetDB()

	// Get model if exist
	var svgtextDB orm.SvgTextDB
	if _, err := db.First(&svgtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&svgtextDB)

	// get an instance (not staged) from DB instance, and call callback function
	svgtextDeleted := new(models.SvgText)
	svgtextDB.CopyBasicFieldsToSvgText(svgtextDeleted)

	// get stage instance from DB instance, and call callback function
	svgtextStaged := backRepo.BackRepoSvgText.Map_SvgTextDBID_SvgTextPtr[svgtextDB.ID]
	if svgtextStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), svgtextStaged, svgtextDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
