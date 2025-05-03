// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Freqency__dummysDeclaration__ models.Freqency
var __Freqency_time__dummyDeclaration time.Duration

var mutexFreqency sync.Mutex

// An FreqencyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFreqency updateFreqency deleteFreqency
type FreqencyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FreqencyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFreqency updateFreqency
type FreqencyInput struct {
	// The Freqency to submit or modify
	// in: body
	Freqency *orm.FreqencyAPI
}

// GetFreqencys
//
// swagger:route GET /freqencys freqencys getFreqencys
//
// # Get all freqencys
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) GetFreqencys(c *gin.Context) {

	// source slice
	var freqencyDBs []orm.FreqencyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFreqencys", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	_, err := db.Find(&freqencyDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	freqencyAPIs := make([]orm.FreqencyAPI, 0)

	// for each freqency, update fields from the database nullable fields
	for idx := range freqencyDBs {
		freqencyDB := &freqencyDBs[idx]
		_ = freqencyDB
		var freqencyAPI orm.FreqencyAPI

		// insertion point for updating fields
		freqencyAPI.ID = freqencyDB.ID
		freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)
		freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
		freqencyAPIs = append(freqencyAPIs, freqencyAPI)
	}

	c.JSON(http.StatusOK, freqencyAPIs)
}

// PostFreqency
//
// swagger:route POST /freqencys freqencys postFreqency
//
// Creates a freqency
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFreqencys", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Validate input
	var input orm.FreqencyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create freqency
	freqencyDB := orm.FreqencyDB{}
	freqencyDB.FreqencyPointersEncoding = input.FreqencyPointersEncoding
	freqencyDB.CopyBasicFieldsFromFreqency_WOP(&input.Freqency_WOP)

	_, err = db.Create(&freqencyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFreqency.CheckoutPhaseOneInstance(&freqencyDB)
	freqency := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]

	if freqency != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), freqency)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, freqencyDB)
}

// GetFreqency
//
// swagger:route GET /freqencys/{ID} freqencys getFreqency
//
// Gets the details for a freqency.
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) GetFreqency(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFreqency", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Get freqencyDB in DB
	var freqencyDB orm.FreqencyDB
	if _, err := db.First(&freqencyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var freqencyAPI orm.FreqencyAPI
	freqencyAPI.ID = freqencyDB.ID
	freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
	freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)

	c.JSON(http.StatusOK, freqencyAPI)
}

// UpdateFreqency
//
// swagger:route PATCH /freqencys/{ID} freqencys updateFreqency
//
// # Update a freqency
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) UpdateFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFreqency", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Validate input
	var input orm.FreqencyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var freqencyDB orm.FreqencyDB

	// fetch the freqency
	_, err := db.First(&freqencyDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	freqencyDB.CopyBasicFieldsFromFreqency_WOP(&input.Freqency_WOP)
	freqencyDB.FreqencyPointersEncoding = input.FreqencyPointersEncoding

	db, _ = db.Model(&freqencyDB)
	_, err = db.Updates(&freqencyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	freqencyNew := new(models.Freqency)
	freqencyDB.CopyBasicFieldsToFreqency(freqencyNew)

	// redeem pointers
	freqencyDB.DecodePointers(backRepo, freqencyNew)

	// get stage instance from DB instance, and call callback function
	freqencyOld := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if freqencyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), freqencyOld, freqencyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the freqencyDB
	c.JSON(http.StatusOK, freqencyDB)
}

// DeleteFreqency
//
// swagger:route DELETE /freqencys/{ID} freqencys deleteFreqency
//
// # Delete a freqency
//
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) DeleteFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFreqency", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Get model if exist
	var freqencyDB orm.FreqencyDB
	if _, err := db.First(&freqencyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&freqencyDB)

	// get an instance (not staged) from DB instance, and call callback function
	freqencyDeleted := new(models.Freqency)
	freqencyDB.CopyBasicFieldsToFreqency(freqencyDeleted)

	// get stage instance from DB instance, and call callback function
	freqencyStaged := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if freqencyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), freqencyStaged, freqencyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
