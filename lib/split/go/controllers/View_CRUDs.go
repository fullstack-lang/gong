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
var __View__dummysDeclaration__ models.View
var __View_time__dummyDeclaration time.Duration

var mutexView sync.Mutex

// An ViewID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getView updateView deleteView
type ViewID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ViewInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postView updateView
type ViewInput struct {
	// The View to submit or modify
	// in: body
	View *orm.ViewAPI
}

// GetViews
//
// swagger:route GET /views views getViews
//
// # Get all views
//
// Responses:
// default: genericError
//
//	200: viewDBResponse
func (controller *Controller) GetViews(c *gin.Context) {

	// source slice
	var viewDBs []orm.ViewDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetViews", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoView.GetDB()

	_, err := db.Find(&viewDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	viewAPIs := make([]orm.ViewAPI, 0)

	// for each view, update fields from the database nullable fields
	for idx := range viewDBs {
		viewDB := &viewDBs[idx]
		_ = viewDB
		var viewAPI orm.ViewAPI

		// insertion point for updating fields
		viewAPI.ID = viewDB.ID
		viewDB.CopyBasicFieldsToView_WOP(&viewAPI.View_WOP)
		viewAPI.ViewPointersEncoding = viewDB.ViewPointersEncoding
		viewAPIs = append(viewAPIs, viewAPI)
	}

	c.JSON(http.StatusOK, viewAPIs)
}

// PostView
//
// swagger:route POST /views views postView
//
// Creates a view
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostView(c *gin.Context) {

	mutexView.Lock()
	defer mutexView.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostViews", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoView.GetDB()

	// Validate input
	var input orm.ViewAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create view
	viewDB := orm.ViewDB{}
	viewDB.ViewPointersEncoding = input.ViewPointersEncoding
	viewDB.CopyBasicFieldsFromView_WOP(&input.View_WOP)

	_, err = db.Create(&viewDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoView.CheckoutPhaseOneInstance(&viewDB)
	view := backRepo.BackRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]

	if view != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), view)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, viewDB)
}

// GetView
//
// swagger:route GET /views/{ID} views getView
//
// Gets the details for a view.
//
// Responses:
// default: genericError
//
//	200: viewDBResponse
func (controller *Controller) GetView(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetView", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoView.GetDB()

	// Get viewDB in DB
	var viewDB orm.ViewDB
	if _, err := db.First(&viewDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var viewAPI orm.ViewAPI
	viewAPI.ID = viewDB.ID
	viewAPI.ViewPointersEncoding = viewDB.ViewPointersEncoding
	viewDB.CopyBasicFieldsToView_WOP(&viewAPI.View_WOP)

	c.JSON(http.StatusOK, viewAPI)
}

// UpdateView
//
// swagger:route PATCH /views/{ID} views updateView
//
// # Update a view
//
// Responses:
// default: genericError
//
//	200: viewDBResponse
func (controller *Controller) UpdateView(c *gin.Context) {

	mutexView.Lock()
	defer mutexView.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateView", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoView.GetDB()

	// Validate input
	var input orm.ViewAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var viewDB orm.ViewDB

	// fetch the view
	_, err := db.First(&viewDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	viewDB.CopyBasicFieldsFromView_WOP(&input.View_WOP)
	viewDB.ViewPointersEncoding = input.ViewPointersEncoding

	db, _ = db.Model(&viewDB)
	_, err = db.Updates(&viewDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	viewNew := new(models.View)
	viewDB.CopyBasicFieldsToView(viewNew)

	// redeem pointers
	viewDB.DecodePointers(backRepo, viewNew)

	// get stage instance from DB instance, and call callback function
	viewOld := backRepo.BackRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]
	if viewOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), viewOld, viewNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the viewDB
	c.JSON(http.StatusOK, viewDB)
}

// DeleteView
//
// swagger:route DELETE /views/{ID} views deleteView
//
// # Delete a view
//
// default: genericError
//
//	200: viewDBResponse
func (controller *Controller) DeleteView(c *gin.Context) {

	mutexView.Lock()
	defer mutexView.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteView", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/split/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoView.GetDB()

	// Get model if exist
	var viewDB orm.ViewDB
	if _, err := db.First(&viewDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&viewDB)

	// get an instance (not staged) from DB instance, and call callback function
	viewDeleted := new(models.View)
	viewDB.CopyBasicFieldsToView(viewDeleted)

	// get stage instance from DB instance, and call callback function
	viewStaged := backRepo.BackRepoView.Map_ViewDBID_ViewPtr[viewDB.ID]
	if viewStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), viewStaged, viewDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
