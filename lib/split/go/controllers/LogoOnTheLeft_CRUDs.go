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
var __LogoOnTheLeft__dummysDeclaration__ models.LogoOnTheLeft
var __LogoOnTheLeft_time__dummyDeclaration time.Duration

var mutexLogoOnTheLeft sync.Mutex

// An LogoOnTheLeftID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLogoOnTheLeft updateLogoOnTheLeft deleteLogoOnTheLeft
type LogoOnTheLeftID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LogoOnTheLeftInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLogoOnTheLeft updateLogoOnTheLeft
type LogoOnTheLeftInput struct {
	// The LogoOnTheLeft to submit or modify
	// in: body
	LogoOnTheLeft *orm.LogoOnTheLeftAPI
}

// GetLogoOnTheLefts
//
// swagger:route GET /logoonthelefts logoonthelefts getLogoOnTheLefts
//
// # Get all logoonthelefts
//
// Responses:
// default: genericError
//
//	200: logoontheleftDBResponse
func (controller *Controller) GetLogoOnTheLefts(c *gin.Context) {

	// source slice
	var logoontheleftDBs []orm.LogoOnTheLeftDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogoOnTheLefts", "Name", stackPath)
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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	_, err := db.Find(&logoontheleftDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	logoontheleftAPIs := make([]orm.LogoOnTheLeftAPI, 0)

	// for each logoontheleft, update fields from the database nullable fields
	for idx := range logoontheleftDBs {
		logoontheleftDB := &logoontheleftDBs[idx]
		_ = logoontheleftDB
		var logoontheleftAPI orm.LogoOnTheLeftAPI

		// insertion point for updating fields
		logoontheleftAPI.ID = logoontheleftDB.ID
		logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft_WOP(&logoontheleftAPI.LogoOnTheLeft_WOP)
		logoontheleftAPI.LogoOnTheLeftPointersEncoding = logoontheleftDB.LogoOnTheLeftPointersEncoding
		logoontheleftAPIs = append(logoontheleftAPIs, logoontheleftAPI)
	}

	c.JSON(http.StatusOK, logoontheleftAPIs)
}

// PostLogoOnTheLeft
//
// swagger:route POST /logoonthelefts logoonthelefts postLogoOnTheLeft
//
// Creates a logoontheleft
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLogoOnTheLeft(c *gin.Context) {

	mutexLogoOnTheLeft.Lock()
	defer mutexLogoOnTheLeft.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLogoOnTheLefts", "Name", stackPath)
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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	// Validate input
	var input orm.LogoOnTheLeftAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create logoontheleft
	logoontheleftDB := orm.LogoOnTheLeftDB{}
	logoontheleftDB.LogoOnTheLeftPointersEncoding = input.LogoOnTheLeftPointersEncoding
	logoontheleftDB.CopyBasicFieldsFromLogoOnTheLeft_WOP(&input.LogoOnTheLeft_WOP)

	_, err = db.Create(&logoontheleftDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLogoOnTheLeft.CheckoutPhaseOneInstance(&logoontheleftDB)
	logoontheleft := backRepo.BackRepoLogoOnTheLeft.Map_LogoOnTheLeftDBID_LogoOnTheLeftPtr[logoontheleftDB.ID]

	if logoontheleft != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), logoontheleft)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, logoontheleftDB)
}

// GetLogoOnTheLeft
//
// swagger:route GET /logoonthelefts/{ID} logoonthelefts getLogoOnTheLeft
//
// Gets the details for a logoontheleft.
//
// Responses:
// default: genericError
//
//	200: logoontheleftDBResponse
func (controller *Controller) GetLogoOnTheLeft(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLogoOnTheLeft", "Name", stackPath)
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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	// Get logoontheleftDB in DB
	var logoontheleftDB orm.LogoOnTheLeftDB
	if _, err := db.First(&logoontheleftDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var logoontheleftAPI orm.LogoOnTheLeftAPI
	logoontheleftAPI.ID = logoontheleftDB.ID
	logoontheleftAPI.LogoOnTheLeftPointersEncoding = logoontheleftDB.LogoOnTheLeftPointersEncoding
	logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft_WOP(&logoontheleftAPI.LogoOnTheLeft_WOP)

	c.JSON(http.StatusOK, logoontheleftAPI)
}

// UpdateLogoOnTheLeft
//
// swagger:route PATCH /logoonthelefts/{ID} logoonthelefts updateLogoOnTheLeft
//
// # Update a logoontheleft
//
// Responses:
// default: genericError
//
//	200: logoontheleftDBResponse
func (controller *Controller) UpdateLogoOnTheLeft(c *gin.Context) {

	mutexLogoOnTheLeft.Lock()
	defer mutexLogoOnTheLeft.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	// Validate input
	var input orm.LogoOnTheLeftAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var logoontheleftDB orm.LogoOnTheLeftDB

	// fetch the logoontheleft
	_, err := db.First(&logoontheleftDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	logoontheleftDB.CopyBasicFieldsFromLogoOnTheLeft_WOP(&input.LogoOnTheLeft_WOP)
	logoontheleftDB.LogoOnTheLeftPointersEncoding = input.LogoOnTheLeftPointersEncoding

	db, _ = db.Model(&logoontheleftDB)
	_, err = db.Updates(&logoontheleftDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	logoontheleftNew := new(models.LogoOnTheLeft)
	logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft(logoontheleftNew)

	// redeem pointers
	logoontheleftDB.DecodePointers(backRepo, logoontheleftNew)

	// get stage instance from DB instance, and call callback function
	logoontheleftOld := backRepo.BackRepoLogoOnTheLeft.Map_LogoOnTheLeftDBID_LogoOnTheLeftPtr[logoontheleftDB.ID]
	if logoontheleftOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), logoontheleftOld, logoontheleftNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the logoontheleftDB
	c.JSON(http.StatusOK, logoontheleftDB)
}

// DeleteLogoOnTheLeft
//
// swagger:route DELETE /logoonthelefts/{ID} logoonthelefts deleteLogoOnTheLeft
//
// # Delete a logoontheleft
//
// default: genericError
//
//	200: logoontheleftDBResponse
func (controller *Controller) DeleteLogoOnTheLeft(c *gin.Context) {

	mutexLogoOnTheLeft.Lock()
	defer mutexLogoOnTheLeft.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLogoOnTheLeft", "Name", stackPath)
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
	db := backRepo.BackRepoLogoOnTheLeft.GetDB()

	// Get model if exist
	var logoontheleftDB orm.LogoOnTheLeftDB
	if _, err := db.First(&logoontheleftDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&logoontheleftDB)

	// get an instance (not staged) from DB instance, and call callback function
	logoontheleftDeleted := new(models.LogoOnTheLeft)
	logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft(logoontheleftDeleted)

	// get stage instance from DB instance, and call callback function
	logoontheleftStaged := backRepo.BackRepoLogoOnTheLeft.Map_LogoOnTheLeftDBID_LogoOnTheLeftPtr[logoontheleftDB.ID]
	if logoontheleftStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), logoontheleftStaged, logoontheleftDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
