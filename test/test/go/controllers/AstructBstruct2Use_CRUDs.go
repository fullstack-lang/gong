// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test/go/models"
	"github.com/fullstack-lang/gong/test/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __AstructBstruct2Use__dummysDeclaration__ models.AstructBstruct2Use
var __AstructBstruct2Use_time__dummyDeclaration time.Duration

var mutexAstructBstruct2Use sync.Mutex

// An AstructBstruct2UseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAstructBstruct2Use updateAstructBstruct2Use deleteAstructBstruct2Use
type AstructBstruct2UseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructBstruct2UseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAstructBstruct2Use updateAstructBstruct2Use
type AstructBstruct2UseInput struct {
	// The AstructBstruct2Use to submit or modify
	// in: body
	AstructBstruct2Use *orm.AstructBstruct2UseAPI
}

// GetAstructBstruct2Uses
//
// swagger:route GET /astructbstruct2uses astructbstruct2uses getAstructBstruct2Uses
//
// # Get all astructbstruct2uses
//
// Responses:
// default: genericError
//
//	200: astructbstruct2useDBResponse
func (controller *Controller) GetAstructBstruct2Uses(c *gin.Context) {

	// source slice
	var astructbstruct2useDBs []orm.AstructBstruct2UseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstructBstruct2Uses", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	_, err := db.Find(&astructbstruct2useDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	astructbstruct2useAPIs := make([]orm.AstructBstruct2UseAPI, 0)

	// for each astructbstruct2use, update fields from the database nullable fields
	for idx := range astructbstruct2useDBs {
		astructbstruct2useDB := &astructbstruct2useDBs[idx]
		_ = astructbstruct2useDB
		var astructbstruct2useAPI orm.AstructBstruct2UseAPI

		// insertion point for updating fields
		astructbstruct2useAPI.ID = astructbstruct2useDB.ID
		astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use_WOP(&astructbstruct2useAPI.AstructBstruct2Use_WOP)
		astructbstruct2useAPI.AstructBstruct2UsePointersEncoding = astructbstruct2useDB.AstructBstruct2UsePointersEncoding
		astructbstruct2useAPIs = append(astructbstruct2useAPIs, astructbstruct2useAPI)
	}

	c.JSON(http.StatusOK, astructbstruct2useAPIs)
}

// PostAstructBstruct2Use
//
// swagger:route POST /astructbstruct2uses astructbstruct2uses postAstructBstruct2Use
//
// Creates a astructbstruct2use
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAstructBstruct2Use(c *gin.Context) {

	mutexAstructBstruct2Use.Lock()
	defer mutexAstructBstruct2Use.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAstructBstruct2Uses", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	// Validate input
	var input orm.AstructBstruct2UseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create astructbstruct2use
	astructbstruct2useDB := orm.AstructBstruct2UseDB{}
	astructbstruct2useDB.AstructBstruct2UsePointersEncoding = input.AstructBstruct2UsePointersEncoding
	astructbstruct2useDB.CopyBasicFieldsFromAstructBstruct2Use_WOP(&input.AstructBstruct2Use_WOP)

	_, err = db.Create(&astructbstruct2useDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAstructBstruct2Use.CheckoutPhaseOneInstance(&astructbstruct2useDB)
	astructbstruct2use := backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UsePtr[astructbstruct2useDB.ID]

	if astructbstruct2use != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), astructbstruct2use)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, astructbstruct2useDB)
}

// GetAstructBstruct2Use
//
// swagger:route GET /astructbstruct2uses/{ID} astructbstruct2uses getAstructBstruct2Use
//
// Gets the details for a astructbstruct2use.
//
// Responses:
// default: genericError
//
//	200: astructbstruct2useDBResponse
func (controller *Controller) GetAstructBstruct2Use(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstructBstruct2Use", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	// Get astructbstruct2useDB in DB
	var astructbstruct2useDB orm.AstructBstruct2UseDB
	if _, err := db.First(&astructbstruct2useDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var astructbstruct2useAPI orm.AstructBstruct2UseAPI
	astructbstruct2useAPI.ID = astructbstruct2useDB.ID
	astructbstruct2useAPI.AstructBstruct2UsePointersEncoding = astructbstruct2useDB.AstructBstruct2UsePointersEncoding
	astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use_WOP(&astructbstruct2useAPI.AstructBstruct2Use_WOP)

	c.JSON(http.StatusOK, astructbstruct2useAPI)
}

// UpdateAstructBstruct2Use
//
// swagger:route PATCH /astructbstruct2uses/{ID} astructbstruct2uses updateAstructBstruct2Use
//
// # Update a astructbstruct2use
//
// Responses:
// default: genericError
//
//	200: astructbstruct2useDBResponse
func (controller *Controller) UpdateAstructBstruct2Use(c *gin.Context) {

	mutexAstructBstruct2Use.Lock()
	defer mutexAstructBstruct2Use.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAstructBstruct2Use", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	// Validate input
	var input orm.AstructBstruct2UseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructbstruct2useDB orm.AstructBstruct2UseDB

	// fetch the astructbstruct2use
	_, err := db.First(&astructbstruct2useDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructbstruct2useDB.CopyBasicFieldsFromAstructBstruct2Use_WOP(&input.AstructBstruct2Use_WOP)
	astructbstruct2useDB.AstructBstruct2UsePointersEncoding = input.AstructBstruct2UsePointersEncoding

	db, _ = db.Model(&astructbstruct2useDB)
	_, err = db.Updates(&astructbstruct2useDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructbstruct2useNew := new(models.AstructBstruct2Use)
	astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use(astructbstruct2useNew)

	// redeem pointers
	astructbstruct2useDB.DecodePointers(backRepo, astructbstruct2useNew)

	// get stage instance from DB instance, and call callback function
	astructbstruct2useOld := backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UsePtr[astructbstruct2useDB.ID]
	if astructbstruct2useOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), astructbstruct2useOld, astructbstruct2useNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructbstruct2useDB
	c.JSON(http.StatusOK, astructbstruct2useDB)
}

// DeleteAstructBstruct2Use
//
// swagger:route DELETE /astructbstruct2uses/{ID} astructbstruct2uses deleteAstructBstruct2Use
//
// # Delete a astructbstruct2use
//
// default: genericError
//
//	200: astructbstruct2useDBResponse
func (controller *Controller) DeleteAstructBstruct2Use(c *gin.Context) {

	mutexAstructBstruct2Use.Lock()
	defer mutexAstructBstruct2Use.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAstructBstruct2Use", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/test/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAstructBstruct2Use.GetDB()

	// Get model if exist
	var astructbstruct2useDB orm.AstructBstruct2UseDB
	if _, err := db.First(&astructbstruct2useDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&astructbstruct2useDB)

	// get an instance (not staged) from DB instance, and call callback function
	astructbstruct2useDeleted := new(models.AstructBstruct2Use)
	astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use(astructbstruct2useDeleted)

	// get stage instance from DB instance, and call callback function
	astructbstruct2useStaged := backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UsePtr[astructbstruct2useDB.ID]
	if astructbstruct2useStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), astructbstruct2useStaged, astructbstruct2useDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
