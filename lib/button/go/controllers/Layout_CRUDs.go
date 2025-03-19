// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Layout__dummysDeclaration__ models.Layout
var __Layout_time__dummyDeclaration time.Duration

var mutexLayout sync.Mutex

// An LayoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLayout updateLayout deleteLayout
type LayoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLayout updateLayout
type LayoutInput struct {
	// The Layout to submit or modify
	// in: body
	Layout *orm.LayoutAPI
}

// GetLayouts
//
// swagger:route GET /layouts layouts getLayouts
//
// # Get all layouts
//
// Responses:
// default: genericError
//
//	200: layoutDBResponse
func (controller *Controller) GetLayouts(c *gin.Context) {

	// source slice
	var layoutDBs []orm.LayoutDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayouts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	_, err := db.Find(&layoutDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	layoutAPIs := make([]orm.LayoutAPI, 0)

	// for each layout, update fields from the database nullable fields
	for idx := range layoutDBs {
		layoutDB := &layoutDBs[idx]
		_ = layoutDB
		var layoutAPI orm.LayoutAPI

		// insertion point for updating fields
		layoutAPI.ID = layoutDB.ID
		layoutDB.CopyBasicFieldsToLayout_WOP(&layoutAPI.Layout_WOP)
		layoutAPI.LayoutPointersEncoding = layoutDB.LayoutPointersEncoding
		layoutAPIs = append(layoutAPIs, layoutAPI)
	}

	c.JSON(http.StatusOK, layoutAPIs)
}

// PostLayout
//
// swagger:route POST /layouts layouts postLayout
//
// Creates a layout
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLayout(c *gin.Context) {

	mutexLayout.Lock()
	defer mutexLayout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLayouts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	// Validate input
	var input orm.LayoutAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create layout
	layoutDB := orm.LayoutDB{}
	layoutDB.LayoutPointersEncoding = input.LayoutPointersEncoding
	layoutDB.CopyBasicFieldsFromLayout_WOP(&input.Layout_WOP)

	_, err = db.Create(&layoutDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLayout.CheckoutPhaseOneInstance(&layoutDB)
	layout := backRepo.BackRepoLayout.Map_LayoutDBID_LayoutPtr[layoutDB.ID]

	if layout != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), layout)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, layoutDB)
}

// GetLayout
//
// swagger:route GET /layouts/{ID} layouts getLayout
//
// Gets the details for a layout.
//
// Responses:
// default: genericError
//
//	200: layoutDBResponse
func (controller *Controller) GetLayout(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayout", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	// Get layoutDB in DB
	var layoutDB orm.LayoutDB
	if _, err := db.First(&layoutDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var layoutAPI orm.LayoutAPI
	layoutAPI.ID = layoutDB.ID
	layoutAPI.LayoutPointersEncoding = layoutDB.LayoutPointersEncoding
	layoutDB.CopyBasicFieldsToLayout_WOP(&layoutAPI.Layout_WOP)

	c.JSON(http.StatusOK, layoutAPI)
}

// UpdateLayout
//
// swagger:route PATCH /layouts/{ID} layouts updateLayout
//
// # Update a layout
//
// Responses:
// default: genericError
//
//	200: layoutDBResponse
func (controller *Controller) UpdateLayout(c *gin.Context) {

	mutexLayout.Lock()
	defer mutexLayout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLayout", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	// Validate input
	var input orm.LayoutAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layoutDB orm.LayoutDB

	// fetch the layout
	_, err := db.First(&layoutDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layoutDB.CopyBasicFieldsFromLayout_WOP(&input.Layout_WOP)
	layoutDB.LayoutPointersEncoding = input.LayoutPointersEncoding

	db, _ = db.Model(&layoutDB)
	_, err = db.Updates(&layoutDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layoutNew := new(models.Layout)
	layoutDB.CopyBasicFieldsToLayout(layoutNew)

	// redeem pointers
	layoutDB.DecodePointers(backRepo, layoutNew)

	// get stage instance from DB instance, and call callback function
	layoutOld := backRepo.BackRepoLayout.Map_LayoutDBID_LayoutPtr[layoutDB.ID]
	if layoutOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), layoutOld, layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layoutDB
	c.JSON(http.StatusOK, layoutDB)
}

// DeleteLayout
//
// swagger:route DELETE /layouts/{ID} layouts deleteLayout
//
// # Delete a layout
//
// default: genericError
//
//	200: layoutDBResponse
func (controller *Controller) DeleteLayout(c *gin.Context) {

	mutexLayout.Lock()
	defer mutexLayout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLayout", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoLayout.GetDB()

	// Get model if exist
	var layoutDB orm.LayoutDB
	if _, err := db.First(&layoutDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&layoutDB)

	// get an instance (not staged) from DB instance, and call callback function
	layoutDeleted := new(models.Layout)
	layoutDB.CopyBasicFieldsToLayout(layoutDeleted)

	// get stage instance from DB instance, and call callback function
	layoutStaged := backRepo.BackRepoLayout.Map_LayoutDBID_LayoutPtr[layoutDB.ID]
	if layoutStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), layoutStaged, layoutDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
