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
var __Logo__dummysDeclaration__ models.Logo
var __Logo_time__dummyDeclaration time.Duration

var mutexLogo sync.Mutex

// An LogoID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLogo updateLogo deleteLogo
type LogoID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LogoInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLogo updateLogo
type LogoInput struct {
	// The Logo to submit or modify
	// in: body
	Logo *orm.LogoAPI
}

// GetLogos
//
// swagger:route GET /logos logos getLogos
//
// # Get all logos
//
// Responses:
// default: genericError
//
//	200: logoDBResponse
func (controller *Controller) GetLogos(c *gin.Context) {

	// source slice
	var logoDBs []orm.LogoDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogos", "Name", stackPath)
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
	db := backRepo.BackRepoLogo.GetDB()

	_, err := db.Find(&logoDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	logoAPIs := make([]orm.LogoAPI, 0)

	// for each logo, update fields from the database nullable fields
	for idx := range logoDBs {
		logoDB := &logoDBs[idx]
		_ = logoDB
		var logoAPI orm.LogoAPI

		// insertion point for updating fields
		logoAPI.ID = logoDB.ID
		logoDB.CopyBasicFieldsToLogo_WOP(&logoAPI.Logo_WOP)
		logoAPI.LogoPointersEncoding = logoDB.LogoPointersEncoding
		logoAPIs = append(logoAPIs, logoAPI)
	}

	c.JSON(http.StatusOK, logoAPIs)
}

// PostLogo
//
// swagger:route POST /logos logos postLogo
//
// Creates a logo
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLogo(c *gin.Context) {

	mutexLogo.Lock()
	defer mutexLogo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLogos", "Name", stackPath)
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
	db := backRepo.BackRepoLogo.GetDB()

	// Validate input
	var input orm.LogoAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create logo
	logoDB := orm.LogoDB{}
	logoDB.LogoPointersEncoding = input.LogoPointersEncoding
	logoDB.CopyBasicFieldsFromLogo_WOP(&input.Logo_WOP)

	_, err = db.Create(&logoDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLogo.CheckoutPhaseOneInstance(&logoDB)
	logo := backRepo.BackRepoLogo.Map_LogoDBID_LogoPtr[logoDB.ID]

	if logo != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), logo)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, logoDB)
}

// GetLogo
//
// swagger:route GET /logos/{ID} logos getLogo
//
// Gets the details for a logo.
//
// Responses:
// default: genericError
//
//	200: logoDBResponse
func (controller *Controller) GetLogo(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogo", "Name", stackPath)
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
	db := backRepo.BackRepoLogo.GetDB()

	// Get logoDB in DB
	var logoDB orm.LogoDB
	if _, err := db.First(&logoDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var logoAPI orm.LogoAPI
	logoAPI.ID = logoDB.ID
	logoAPI.LogoPointersEncoding = logoDB.LogoPointersEncoding
	logoDB.CopyBasicFieldsToLogo_WOP(&logoAPI.Logo_WOP)

	c.JSON(http.StatusOK, logoAPI)
}

// UpdateLogo
//
// swagger:route PATCH /logos/{ID} logos updateLogo
//
// # Update a logo
//
// Responses:
// default: genericError
//
//	200: logoDBResponse
func (controller *Controller) UpdateLogo(c *gin.Context) {

	mutexLogo.Lock()
	defer mutexLogo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLogo", "Name", stackPath)
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
	db := backRepo.BackRepoLogo.GetDB()

	// Validate input
	var input orm.LogoAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var logoDB orm.LogoDB

	// fetch the logo
	_, err := db.First(&logoDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	logoDB.CopyBasicFieldsFromLogo_WOP(&input.Logo_WOP)
	logoDB.LogoPointersEncoding = input.LogoPointersEncoding

	db, _ = db.Model(&logoDB)
	_, err = db.Updates(&logoDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	logoNew := new(models.Logo)
	logoDB.CopyBasicFieldsToLogo(logoNew)

	// redeem pointers
	logoDB.DecodePointers(backRepo, logoNew)

	// get stage instance from DB instance, and call callback function
	logoOld := backRepo.BackRepoLogo.Map_LogoDBID_LogoPtr[logoDB.ID]
	if logoOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), logoOld, logoNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the logoDB
	c.JSON(http.StatusOK, logoDB)
}

// DeleteLogo
//
// swagger:route DELETE /logos/{ID} logos deleteLogo
//
// # Delete a logo
//
// default: genericError
//
//	200: logoDBResponse
func (controller *Controller) DeleteLogo(c *gin.Context) {

	mutexLogo.Lock()
	defer mutexLogo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLogo", "Name", stackPath)
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
	db := backRepo.BackRepoLogo.GetDB()

	// Get model if exist
	var logoDB orm.LogoDB
	if _, err := db.First(&logoDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&logoDB)

	// get an instance (not staged) from DB instance, and call callback function
	logoDeleted := new(models.Logo)
	logoDB.CopyBasicFieldsToLogo(logoDeleted)

	// get stage instance from DB instance, and call callback function
	logoStaged := backRepo.BackRepoLogo.Map_LogoDBID_LogoPtr[logoDB.ID]
	if logoStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), logoStaged, logoDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
