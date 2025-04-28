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
var __XLCell__dummysDeclaration__ models.XLCell
var __XLCell_time__dummyDeclaration time.Duration

var mutexXLCell sync.Mutex

// An XLCellID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXLCell updateXLCell deleteXLCell
type XLCellID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XLCellInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXLCell updateXLCell
type XLCellInput struct {
	// The XLCell to submit or modify
	// in: body
	XLCell *orm.XLCellAPI
}

// GetXLCells
//
// swagger:route GET /xlcells xlcells getXLCells
//
// # Get all xlcells
//
// Responses:
// default: genericError
//
//	200: xlcellDBResponse
func (controller *Controller) GetXLCells(c *gin.Context) {

	// source slice
	var xlcellDBs []orm.XLCellDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLCells", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXLCell.GetDB()

	_, err := db.Find(&xlcellDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlcellAPIs := make([]orm.XLCellAPI, 0)

	// for each xlcell, update fields from the database nullable fields
	for idx := range xlcellDBs {
		xlcellDB := &xlcellDBs[idx]
		_ = xlcellDB
		var xlcellAPI orm.XLCellAPI

		// insertion point for updating fields
		xlcellAPI.ID = xlcellDB.ID
		xlcellDB.CopyBasicFieldsToXLCell_WOP(&xlcellAPI.XLCell_WOP)
		xlcellAPI.XLCellPointersEncoding = xlcellDB.XLCellPointersEncoding
		xlcellAPIs = append(xlcellAPIs, xlcellAPI)
	}

	c.JSON(http.StatusOK, xlcellAPIs)
}

// PostXLCell
//
// swagger:route POST /xlcells xlcells postXLCell
//
// Creates a xlcell
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXLCell(c *gin.Context) {

	mutexXLCell.Lock()
	defer mutexXLCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXLCells", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXLCell.GetDB()

	// Validate input
	var input orm.XLCellAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlcell
	xlcellDB := orm.XLCellDB{}
	xlcellDB.XLCellPointersEncoding = input.XLCellPointersEncoding
	xlcellDB.CopyBasicFieldsFromXLCell_WOP(&input.XLCell_WOP)

	_, err = db.Create(&xlcellDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXLCell.CheckoutPhaseOneInstance(&xlcellDB)
	xlcell := backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellPtr[xlcellDB.ID]

	if xlcell != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlcell)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlcellDB)
}

// GetXLCell
//
// swagger:route GET /xlcells/{ID} xlcells getXLCell
//
// Gets the details for a xlcell.
//
// Responses:
// default: genericError
//
//	200: xlcellDBResponse
func (controller *Controller) GetXLCell(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLCell", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXLCell.GetDB()

	// Get xlcellDB in DB
	var xlcellDB orm.XLCellDB
	if _, err := db.First(&xlcellDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlcellAPI orm.XLCellAPI
	xlcellAPI.ID = xlcellDB.ID
	xlcellAPI.XLCellPointersEncoding = xlcellDB.XLCellPointersEncoding
	xlcellDB.CopyBasicFieldsToXLCell_WOP(&xlcellAPI.XLCell_WOP)

	c.JSON(http.StatusOK, xlcellAPI)
}

// UpdateXLCell
//
// swagger:route PATCH /xlcells/{ID} xlcells updateXLCell
//
// # Update a xlcell
//
// Responses:
// default: genericError
//
//	200: xlcellDBResponse
func (controller *Controller) UpdateXLCell(c *gin.Context) {

	mutexXLCell.Lock()
	defer mutexXLCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXLCell", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXLCell.GetDB()

	// Validate input
	var input orm.XLCellAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlcellDB orm.XLCellDB

	// fetch the xlcell
	_, err := db.First(&xlcellDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlcellDB.CopyBasicFieldsFromXLCell_WOP(&input.XLCell_WOP)
	xlcellDB.XLCellPointersEncoding = input.XLCellPointersEncoding

	db, _ = db.Model(&xlcellDB)
	_, err = db.Updates(&xlcellDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlcellNew := new(models.XLCell)
	xlcellDB.CopyBasicFieldsToXLCell(xlcellNew)

	// redeem pointers
	xlcellDB.DecodePointers(backRepo, xlcellNew)

	// get stage instance from DB instance, and call callback function
	xlcellOld := backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellPtr[xlcellDB.ID]
	if xlcellOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlcellOld, xlcellNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlcellDB
	c.JSON(http.StatusOK, xlcellDB)
}

// DeleteXLCell
//
// swagger:route DELETE /xlcells/{ID} xlcells deleteXLCell
//
// # Delete a xlcell
//
// default: genericError
//
//	200: xlcellDBResponse
func (controller *Controller) DeleteXLCell(c *gin.Context) {

	mutexXLCell.Lock()
	defer mutexXLCell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXLCell", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/xlsx/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXLCell.GetDB()

	// Get model if exist
	var xlcellDB orm.XLCellDB
	if _, err := db.First(&xlcellDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xlcellDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlcellDeleted := new(models.XLCell)
	xlcellDB.CopyBasicFieldsToXLCell(xlcellDeleted)

	// get stage instance from DB instance, and call callback function
	xlcellStaged := backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellPtr[xlcellDB.ID]
	if xlcellStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlcellStaged, xlcellDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
