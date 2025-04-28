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
var __Xlsx__dummysDeclaration__ models.Xlsx
var __Xlsx_time__dummyDeclaration time.Duration

var mutexXlsx sync.Mutex

// An XlsxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXlsx updateXlsx deleteXlsx
type XlsxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XlsxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXlsx updateXlsx
type XlsxInput struct {
	// The Xlsx to submit or modify
	// in: body
	Xlsx *orm.XlsxAPI
}

// GetXlsxs
//
// swagger:route GET /xlsxs xlsxs getXlsxs
//
// # Get all xlsxs
//
// Responses:
// default: genericError
//
//	200: xlsxDBResponse
func (controller *Controller) GetXlsxs(c *gin.Context) {

	// source slice
	var xlsxDBs []orm.XlsxDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXlsxs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	_, err := db.Find(&xlsxDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlsxAPIs := make([]orm.XlsxAPI, 0)

	// for each xlsx, update fields from the database nullable fields
	for idx := range xlsxDBs {
		xlsxDB := &xlsxDBs[idx]
		_ = xlsxDB
		var xlsxAPI orm.XlsxAPI

		// insertion point for updating fields
		xlsxAPI.ID = xlsxDB.ID
		xlsxDB.CopyBasicFieldsToXlsx_WOP(&xlsxAPI.Xlsx_WOP)
		xlsxAPI.XlsxPointersEncoding = xlsxDB.XlsxPointersEncoding
		xlsxAPIs = append(xlsxAPIs, xlsxAPI)
	}

	c.JSON(http.StatusOK, xlsxAPIs)
}

// PostXlsx
//
// swagger:route POST /xlsxs xlsxs postXlsx
//
// Creates a xlsx
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXlsx(c *gin.Context) {

	mutexXlsx.Lock()
	defer mutexXlsx.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXlsxs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	// Validate input
	var input orm.XlsxAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlsx
	xlsxDB := orm.XlsxDB{}
	xlsxDB.XlsxPointersEncoding = input.XlsxPointersEncoding
	xlsxDB.CopyBasicFieldsFromXlsx_WOP(&input.Xlsx_WOP)

	_, err = db.Create(&xlsxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXlsx.CheckoutPhaseOneInstance(&xlsxDB)
	xlsx := backRepo.BackRepoXlsx.Map_XlsxDBID_XlsxPtr[xlsxDB.ID]

	if xlsx != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlsx)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlsxDB)
}

// GetXlsx
//
// swagger:route GET /xlsxs/{ID} xlsxs getXlsx
//
// Gets the details for a xlsx.
//
// Responses:
// default: genericError
//
//	200: xlsxDBResponse
func (controller *Controller) GetXlsx(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXlsx", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	// Get xlsxDB in DB
	var xlsxDB orm.XlsxDB
	if _, err := db.First(&xlsxDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlsxAPI orm.XlsxAPI
	xlsxAPI.ID = xlsxDB.ID
	xlsxAPI.XlsxPointersEncoding = xlsxDB.XlsxPointersEncoding
	xlsxDB.CopyBasicFieldsToXlsx_WOP(&xlsxAPI.Xlsx_WOP)

	c.JSON(http.StatusOK, xlsxAPI)
}

// UpdateXlsx
//
// swagger:route PATCH /xlsxs/{ID} xlsxs updateXlsx
//
// # Update a xlsx
//
// Responses:
// default: genericError
//
//	200: xlsxDBResponse
func (controller *Controller) UpdateXlsx(c *gin.Context) {

	mutexXlsx.Lock()
	defer mutexXlsx.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXlsx", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	// Validate input
	var input orm.XlsxAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlsxDB orm.XlsxDB

	// fetch the xlsx
	_, err := db.First(&xlsxDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlsxDB.CopyBasicFieldsFromXlsx_WOP(&input.Xlsx_WOP)
	xlsxDB.XlsxPointersEncoding = input.XlsxPointersEncoding

	db, _ = db.Model(&xlsxDB)
	_, err = db.Updates(&xlsxDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlsxNew := new(models.Xlsx)
	xlsxDB.CopyBasicFieldsToXlsx(xlsxNew)

	// redeem pointers
	xlsxDB.DecodePointers(backRepo, xlsxNew)

	// get stage instance from DB instance, and call callback function
	xlsxOld := backRepo.BackRepoXlsx.Map_XlsxDBID_XlsxPtr[xlsxDB.ID]
	if xlsxOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlsxOld, xlsxNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlsxDB
	c.JSON(http.StatusOK, xlsxDB)
}

// DeleteXlsx
//
// swagger:route DELETE /xlsxs/{ID} xlsxs deleteXlsx
//
// # Delete a xlsx
//
// default: genericError
//
//	200: xlsxDBResponse
func (controller *Controller) DeleteXlsx(c *gin.Context) {

	mutexXlsx.Lock()
	defer mutexXlsx.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXlsx", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoXlsx.GetDB()

	// Get model if exist
	var xlsxDB orm.XlsxDB
	if _, err := db.First(&xlsxDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xlsxDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlsxDeleted := new(models.Xlsx)
	xlsxDB.CopyBasicFieldsToXlsx(xlsxDeleted)

	// get stage instance from DB instance, and call callback function
	xlsxStaged := backRepo.BackRepoXlsx.Map_XlsxDBID_XlsxPtr[xlsxDB.ID]
	if xlsxStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlsxStaged, xlsxDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
