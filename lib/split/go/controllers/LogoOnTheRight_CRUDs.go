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
var __LogoOnTheRight__dummysDeclaration__ models.LogoOnTheRight
var __LogoOnTheRight_time__dummyDeclaration time.Duration

var mutexLogoOnTheRight sync.Mutex

// An LogoOnTheRightID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLogoOnTheRight updateLogoOnTheRight deleteLogoOnTheRight
type LogoOnTheRightID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LogoOnTheRightInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLogoOnTheRight updateLogoOnTheRight
type LogoOnTheRightInput struct {
	// The LogoOnTheRight to submit or modify
	// in: body
	LogoOnTheRight *orm.LogoOnTheRightAPI
}

// GetLogoOnTheRights
//
// swagger:route GET /logoontherights logoontherights getLogoOnTheRights
//
// # Get all logoontherights
//
// Responses:
// default: genericError
//
//	200: logoontherightDBResponse
func (controller *Controller) GetLogoOnTheRights(c *gin.Context) {

	// source slice
	var logoontherightDBs []orm.LogoOnTheRightDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogoOnTheRights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	_, err := db.Find(&logoontherightDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	logoontherightAPIs := make([]orm.LogoOnTheRightAPI, 0)

	// for each logoontheright, update fields from the database nullable fields
	for idx := range logoontherightDBs {
		logoontherightDB := &logoontherightDBs[idx]
		_ = logoontherightDB
		var logoontherightAPI orm.LogoOnTheRightAPI

		// insertion point for updating fields
		logoontherightAPI.ID = logoontherightDB.ID
		logoontherightDB.CopyBasicFieldsToLogoOnTheRight_WOP(&logoontherightAPI.LogoOnTheRight_WOP)
		logoontherightAPI.LogoOnTheRightPointersEncoding = logoontherightDB.LogoOnTheRightPointersEncoding
		logoontherightAPIs = append(logoontherightAPIs, logoontherightAPI)
	}

	c.JSON(http.StatusOK, logoontherightAPIs)
}

// PostLogoOnTheRight
//
// swagger:route POST /logoontherights logoontherights postLogoOnTheRight
//
// Creates a logoontheright
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLogoOnTheRight(c *gin.Context) {

	mutexLogoOnTheRight.Lock()
	defer mutexLogoOnTheRight.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLogoOnTheRights", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	// Validate input
	var input orm.LogoOnTheRightAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create logoontheright
	logoontherightDB := orm.LogoOnTheRightDB{}
	logoontherightDB.LogoOnTheRightPointersEncoding = input.LogoOnTheRightPointersEncoding
	logoontherightDB.CopyBasicFieldsFromLogoOnTheRight_WOP(&input.LogoOnTheRight_WOP)

	_, err = db.Create(&logoontherightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLogoOnTheRight.CheckoutPhaseOneInstance(&logoontherightDB)
	logoontheright := backRepo.BackRepoLogoOnTheRight.Map_LogoOnTheRightDBID_LogoOnTheRightPtr[logoontherightDB.ID]

	if logoontheright != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), logoontheright)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, logoontherightDB)
}

// GetLogoOnTheRight
//
// swagger:route GET /logoontherights/{ID} logoontherights getLogoOnTheRight
//
// Gets the details for a logoontheright.
//
// Responses:
// default: genericError
//
//	200: logoontherightDBResponse
func (controller *Controller) GetLogoOnTheRight(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogoOnTheRight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	// Get logoontherightDB in DB
	var logoontherightDB orm.LogoOnTheRightDB
	if _, err := db.First(&logoontherightDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var logoontherightAPI orm.LogoOnTheRightAPI
	logoontherightAPI.ID = logoontherightDB.ID
	logoontherightAPI.LogoOnTheRightPointersEncoding = logoontherightDB.LogoOnTheRightPointersEncoding
	logoontherightDB.CopyBasicFieldsToLogoOnTheRight_WOP(&logoontherightAPI.LogoOnTheRight_WOP)

	c.JSON(http.StatusOK, logoontherightAPI)
}

// UpdateLogoOnTheRight
//
// swagger:route PATCH /logoontherights/{ID} logoontherights updateLogoOnTheRight
//
// # Update a logoontheright
//
// Responses:
// default: genericError
//
//	200: logoontherightDBResponse
func (controller *Controller) UpdateLogoOnTheRight(c *gin.Context) {

	mutexLogoOnTheRight.Lock()
	defer mutexLogoOnTheRight.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLogoOnTheRight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	// Validate input
	var input orm.LogoOnTheRightAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var logoontherightDB orm.LogoOnTheRightDB

	// fetch the logoontheright
	_, err := db.First(&logoontherightDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	logoontherightDB.CopyBasicFieldsFromLogoOnTheRight_WOP(&input.LogoOnTheRight_WOP)
	logoontherightDB.LogoOnTheRightPointersEncoding = input.LogoOnTheRightPointersEncoding

	db, _ = db.Model(&logoontherightDB)
	_, err = db.Updates(&logoontherightDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	logoontherightNew := new(models.LogoOnTheRight)
	logoontherightDB.CopyBasicFieldsToLogoOnTheRight(logoontherightNew)

	// redeem pointers
	logoontherightDB.DecodePointers(backRepo, logoontherightNew)

	// get stage instance from DB instance, and call callback function
	logoontherightOld := backRepo.BackRepoLogoOnTheRight.Map_LogoOnTheRightDBID_LogoOnTheRightPtr[logoontherightDB.ID]
	if logoontherightOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), logoontherightOld, logoontherightNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the logoontherightDB
	c.JSON(http.StatusOK, logoontherightDB)
}

// DeleteLogoOnTheRight
//
// swagger:route DELETE /logoontherights/{ID} logoontherights deleteLogoOnTheRight
//
// # Delete a logoontheright
//
// default: genericError
//
//	200: logoontherightDBResponse
func (controller *Controller) DeleteLogoOnTheRight(c *gin.Context) {

	mutexLogoOnTheRight.Lock()
	defer mutexLogoOnTheRight.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLogoOnTheRight", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLogoOnTheRight.GetDB()

	// Get model if exist
	var logoontherightDB orm.LogoOnTheRightDB
	if _, err := db.First(&logoontherightDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&logoontherightDB)

	// get an instance (not staged) from DB instance, and call callback function
	logoontherightDeleted := new(models.LogoOnTheRight)
	logoontherightDB.CopyBasicFieldsToLogoOnTheRight(logoontherightDeleted)

	// get stage instance from DB instance, and call callback function
	logoontherightStaged := backRepo.BackRepoLogoOnTheRight.Map_LogoOnTheRightDBID_LogoOnTheRightPtr[logoontherightDB.ID]
	if logoontherightStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), logoontherightStaged, logoontherightDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
