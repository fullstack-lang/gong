// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
	"github.com/fullstack-lang/gong/lib/xlsx/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __DisplaySelection__dummysDeclaration__ models.DisplaySelection
var __DisplaySelection_time__dummyDeclaration time.Duration

var mutexDisplaySelection sync.Mutex

// An DisplaySelectionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDisplaySelection updateDisplaySelection deleteDisplaySelection
type DisplaySelectionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DisplaySelectionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDisplaySelection updateDisplaySelection
type DisplaySelectionInput struct {
	// The DisplaySelection to submit or modify
	// in: body
	DisplaySelection *orm.DisplaySelectionAPI
}

// GetDisplaySelections
//
// swagger:route GET /displayselections displayselections getDisplaySelections
//
// # Get all displayselections
//
// Responses:
// default: genericError
//
//	200: displayselectionDBResponse
func (controller *Controller) GetDisplaySelections(c *gin.Context) {

	// source slice
	var displayselectionDBs []orm.DisplaySelectionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDisplaySelections", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDisplaySelection.GetDB()

	_, err := db.Find(&displayselectionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	displayselectionAPIs := make([]orm.DisplaySelectionAPI, 0)

	// for each displayselection, update fields from the database nullable fields
	for idx := range displayselectionDBs {
		displayselectionDB := &displayselectionDBs[idx]
		_ = displayselectionDB
		var displayselectionAPI orm.DisplaySelectionAPI

		// insertion point for updating fields
		displayselectionAPI.ID = displayselectionDB.ID
		displayselectionDB.CopyBasicFieldsToDisplaySelection_WOP(&displayselectionAPI.DisplaySelection_WOP)
		displayselectionAPI.DisplaySelectionPointersEncoding = displayselectionDB.DisplaySelectionPointersEncoding
		displayselectionAPIs = append(displayselectionAPIs, displayselectionAPI)
	}

	c.JSON(http.StatusOK, displayselectionAPIs)
}

// PostDisplaySelection
//
// swagger:route POST /displayselections displayselections postDisplaySelection
//
// Creates a displayselection
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDisplaySelection(c *gin.Context) {

	mutexDisplaySelection.Lock()
	defer mutexDisplaySelection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDisplaySelections", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDisplaySelection.GetDB()

	// Validate input
	var input orm.DisplaySelectionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create displayselection
	displayselectionDB := orm.DisplaySelectionDB{}
	displayselectionDB.DisplaySelectionPointersEncoding = input.DisplaySelectionPointersEncoding
	displayselectionDB.CopyBasicFieldsFromDisplaySelection_WOP(&input.DisplaySelection_WOP)

	_, err = db.Create(&displayselectionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDisplaySelection.CheckoutPhaseOneInstance(&displayselectionDB)
	displayselection := backRepo.BackRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr[displayselectionDB.ID]

	if displayselection != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), displayselection)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, displayselectionDB)
}

// GetDisplaySelection
//
// swagger:route GET /displayselections/{ID} displayselections getDisplaySelection
//
// Gets the details for a displayselection.
//
// Responses:
// default: genericError
//
//	200: displayselectionDBResponse
func (controller *Controller) GetDisplaySelection(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDisplaySelection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDisplaySelection.GetDB()

	// Get displayselectionDB in DB
	var displayselectionDB orm.DisplaySelectionDB
	if _, err := db.First(&displayselectionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var displayselectionAPI orm.DisplaySelectionAPI
	displayselectionAPI.ID = displayselectionDB.ID
	displayselectionAPI.DisplaySelectionPointersEncoding = displayselectionDB.DisplaySelectionPointersEncoding
	displayselectionDB.CopyBasicFieldsToDisplaySelection_WOP(&displayselectionAPI.DisplaySelection_WOP)

	c.JSON(http.StatusOK, displayselectionAPI)
}

// UpdateDisplaySelection
//
// swagger:route PATCH /displayselections/{ID} displayselections updateDisplaySelection
//
// # Update a displayselection
//
// Responses:
// default: genericError
//
//	200: displayselectionDBResponse
func (controller *Controller) UpdateDisplaySelection(c *gin.Context) {

	mutexDisplaySelection.Lock()
	defer mutexDisplaySelection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDisplaySelection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDisplaySelection.GetDB()

	// Validate input
	var input orm.DisplaySelectionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var displayselectionDB orm.DisplaySelectionDB

	// fetch the displayselection
	_, err := db.First(&displayselectionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	displayselectionDB.CopyBasicFieldsFromDisplaySelection_WOP(&input.DisplaySelection_WOP)
	displayselectionDB.DisplaySelectionPointersEncoding = input.DisplaySelectionPointersEncoding

	db, _ = db.Model(&displayselectionDB)
	_, err = db.Updates(&displayselectionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	displayselectionNew := new(models.DisplaySelection)
	displayselectionDB.CopyBasicFieldsToDisplaySelection(displayselectionNew)

	// redeem pointers
	displayselectionDB.DecodePointers(backRepo, displayselectionNew)

	// get stage instance from DB instance, and call callback function
	displayselectionOld := backRepo.BackRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr[displayselectionDB.ID]
	if displayselectionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), displayselectionOld, displayselectionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the displayselectionDB
	c.JSON(http.StatusOK, displayselectionDB)
}

// DeleteDisplaySelection
//
// swagger:route DELETE /displayselections/{ID} displayselections deleteDisplaySelection
//
// # Delete a displayselection
//
// default: genericError
//
//	200: displayselectionDBResponse
func (controller *Controller) DeleteDisplaySelection(c *gin.Context) {

	mutexDisplaySelection.Lock()
	defer mutexDisplaySelection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDisplaySelection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDisplaySelection.GetDB()

	// Get model if exist
	var displayselectionDB orm.DisplaySelectionDB
	if _, err := db.First(&displayselectionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&displayselectionDB)

	// get an instance (not staged) from DB instance, and call callback function
	displayselectionDeleted := new(models.DisplaySelection)
	displayselectionDB.CopyBasicFieldsToDisplaySelection(displayselectionDeleted)

	// get stage instance from DB instance, and call callback function
	displayselectionStaged := backRepo.BackRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr[displayselectionDB.ID]
	if displayselectionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), displayselectionStaged, displayselectionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
