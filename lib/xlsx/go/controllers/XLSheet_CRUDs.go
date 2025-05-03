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
var __XLSheet__dummysDeclaration__ models.XLSheet
var __XLSheet_time__dummyDeclaration time.Duration

var mutexXLSheet sync.Mutex

// An XLSheetID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXLSheet updateXLSheet deleteXLSheet
type XLSheetID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XLSheetInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXLSheet updateXLSheet
type XLSheetInput struct {
	// The XLSheet to submit or modify
	// in: body
	XLSheet *orm.XLSheetAPI
}

// GetXLSheets
//
// swagger:route GET /xlsheets xlsheets getXLSheets
//
// # Get all xlsheets
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) GetXLSheets(c *gin.Context) {

	// source slice
	var xlsheetDBs []orm.XLSheetDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLSheets", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	_, err := db.Find(&xlsheetDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlsheetAPIs := make([]orm.XLSheetAPI, 0)

	// for each xlsheet, update fields from the database nullable fields
	for idx := range xlsheetDBs {
		xlsheetDB := &xlsheetDBs[idx]
		_ = xlsheetDB
		var xlsheetAPI orm.XLSheetAPI

		// insertion point for updating fields
		xlsheetAPI.ID = xlsheetDB.ID
		xlsheetDB.CopyBasicFieldsToXLSheet_WOP(&xlsheetAPI.XLSheet_WOP)
		xlsheetAPI.XLSheetPointersEncoding = xlsheetDB.XLSheetPointersEncoding
		xlsheetAPIs = append(xlsheetAPIs, xlsheetAPI)
	}

	c.JSON(http.StatusOK, xlsheetAPIs)
}

// PostXLSheet
//
// swagger:route POST /xlsheets xlsheets postXLSheet
//
// Creates a xlsheet
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()
	defer mutexXLSheet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXLSheets", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Validate input
	var input orm.XLSheetAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlsheet
	xlsheetDB := orm.XLSheetDB{}
	xlsheetDB.XLSheetPointersEncoding = input.XLSheetPointersEncoding
	xlsheetDB.CopyBasicFieldsFromXLSheet_WOP(&input.XLSheet_WOP)

	_, err = db.Create(&xlsheetDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXLSheet.CheckoutPhaseOneInstance(&xlsheetDB)
	xlsheet := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]

	if xlsheet != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlsheet)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlsheetDB)
}

// GetXLSheet
//
// swagger:route GET /xlsheets/{ID} xlsheets getXLSheet
//
// Gets the details for a xlsheet.
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) GetXLSheet(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLSheet", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Get xlsheetDB in DB
	var xlsheetDB orm.XLSheetDB
	if _, err := db.First(&xlsheetDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlsheetAPI orm.XLSheetAPI
	xlsheetAPI.ID = xlsheetDB.ID
	xlsheetAPI.XLSheetPointersEncoding = xlsheetDB.XLSheetPointersEncoding
	xlsheetDB.CopyBasicFieldsToXLSheet_WOP(&xlsheetAPI.XLSheet_WOP)

	c.JSON(http.StatusOK, xlsheetAPI)
}

// UpdateXLSheet
//
// swagger:route PATCH /xlsheets/{ID} xlsheets updateXLSheet
//
// # Update a xlsheet
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) UpdateXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()
	defer mutexXLSheet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXLSheet", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Validate input
	var input orm.XLSheetAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlsheetDB orm.XLSheetDB

	// fetch the xlsheet
	_, err := db.First(&xlsheetDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlsheetDB.CopyBasicFieldsFromXLSheet_WOP(&input.XLSheet_WOP)
	xlsheetDB.XLSheetPointersEncoding = input.XLSheetPointersEncoding

	db, _ = db.Model(&xlsheetDB)
	_, err = db.Updates(&xlsheetDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlsheetNew := new(models.XLSheet)
	xlsheetDB.CopyBasicFieldsToXLSheet(xlsheetNew)

	// redeem pointers
	xlsheetDB.DecodePointers(backRepo, xlsheetNew)

	// get stage instance from DB instance, and call callback function
	xlsheetOld := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]
	if xlsheetOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlsheetOld, xlsheetNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlsheetDB
	c.JSON(http.StatusOK, xlsheetDB)
}

// DeleteXLSheet
//
// swagger:route DELETE /xlsheets/{ID} xlsheets deleteXLSheet
//
// # Delete a xlsheet
//
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) DeleteXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()
	defer mutexXLSheet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXLSheet", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Get model if exist
	var xlsheetDB orm.XLSheetDB
	if _, err := db.First(&xlsheetDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xlsheetDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlsheetDeleted := new(models.XLSheet)
	xlsheetDB.CopyBasicFieldsToXLSheet(xlsheetDeleted)

	// get stage instance from DB instance, and call callback function
	xlsheetStaged := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]
	if xlsheetStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlsheetStaged, xlsheetDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
