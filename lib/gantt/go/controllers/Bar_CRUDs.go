// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	"github.com/fullstack-lang/gong/lib/gantt/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Bar__dummysDeclaration__ models.Bar
var __Bar_time__dummyDeclaration time.Duration

var mutexBar sync.Mutex

// An BarID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBar updateBar deleteBar
type BarID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BarInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBar updateBar
type BarInput struct {
	// The Bar to submit or modify
	// in: body
	Bar *orm.BarAPI
}

// GetBars
//
// swagger:route GET /bars bars getBars
//
// # Get all bars
//
// Responses:
// default: genericError
//
//	200: barDBResponse
func (controller *Controller) GetBars(c *gin.Context) {

	// source slice
	var barDBs []orm.BarDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBars", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBar.GetDB()

	_, err := db.Find(&barDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	barAPIs := make([]orm.BarAPI, 0)

	// for each bar, update fields from the database nullable fields
	for idx := range barDBs {
		barDB := &barDBs[idx]
		_ = barDB
		var barAPI orm.BarAPI

		// insertion point for updating fields
		barAPI.ID = barDB.ID
		barDB.CopyBasicFieldsToBar_WOP(&barAPI.Bar_WOP)
		barAPI.BarPointersEncoding = barDB.BarPointersEncoding
		barAPIs = append(barAPIs, barAPI)
	}

	c.JSON(http.StatusOK, barAPIs)
}

// PostBar
//
// swagger:route POST /bars bars postBar
//
// Creates a bar
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBar(c *gin.Context) {

	mutexBar.Lock()
	defer mutexBar.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBars", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBar.GetDB()

	// Validate input
	var input orm.BarAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bar
	barDB := orm.BarDB{}
	barDB.BarPointersEncoding = input.BarPointersEncoding
	barDB.CopyBasicFieldsFromBar_WOP(&input.Bar_WOP)

	_, err = db.Create(&barDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBar.CheckoutPhaseOneInstance(&barDB)
	bar := backRepo.BackRepoBar.Map_BarDBID_BarPtr[barDB.ID]

	if bar != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bar)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, barDB)
}

// GetBar
//
// swagger:route GET /bars/{ID} bars getBar
//
// Gets the details for a bar.
//
// Responses:
// default: genericError
//
//	200: barDBResponse
func (controller *Controller) GetBar(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBar", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBar.GetDB()

	// Get barDB in DB
	var barDB orm.BarDB
	if _, err := db.First(&barDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var barAPI orm.BarAPI
	barAPI.ID = barDB.ID
	barAPI.BarPointersEncoding = barDB.BarPointersEncoding
	barDB.CopyBasicFieldsToBar_WOP(&barAPI.Bar_WOP)

	c.JSON(http.StatusOK, barAPI)
}

// UpdateBar
//
// swagger:route PATCH /bars/{ID} bars updateBar
//
// # Update a bar
//
// Responses:
// default: genericError
//
//	200: barDBResponse
func (controller *Controller) UpdateBar(c *gin.Context) {

	mutexBar.Lock()
	defer mutexBar.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBar", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBar.GetDB()

	// Validate input
	var input orm.BarAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var barDB orm.BarDB

	// fetch the bar
	_, err := db.First(&barDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	barDB.CopyBasicFieldsFromBar_WOP(&input.Bar_WOP)
	barDB.BarPointersEncoding = input.BarPointersEncoding

	db, _ = db.Model(&barDB)
	_, err = db.Updates(&barDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	barNew := new(models.Bar)
	barDB.CopyBasicFieldsToBar(barNew)

	// redeem pointers
	barDB.DecodePointers(backRepo, barNew)

	// get stage instance from DB instance, and call callback function
	barOld := backRepo.BackRepoBar.Map_BarDBID_BarPtr[barDB.ID]
	if barOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), barOld, barNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the barDB
	c.JSON(http.StatusOK, barDB)
}

// DeleteBar
//
// swagger:route DELETE /bars/{ID} bars deleteBar
//
// # Delete a bar
//
// default: genericError
//
//	200: barDBResponse
func (controller *Controller) DeleteBar(c *gin.Context) {

	mutexBar.Lock()
	defer mutexBar.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBar", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/gantt/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoBar.GetDB()

	// Get model if exist
	var barDB orm.BarDB
	if _, err := db.First(&barDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&barDB)

	// get an instance (not staged) from DB instance, and call callback function
	barDeleted := new(models.Bar)
	barDB.CopyBasicFieldsToBar(barDeleted)

	// get stage instance from DB instance, and call callback function
	barStaged := backRepo.BackRepoBar.Map_BarDBID_BarPtr[barDB.ID]
	if barStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), barStaged, barDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
