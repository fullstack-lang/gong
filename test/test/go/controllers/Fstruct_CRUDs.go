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
var __Fstruct__dummysDeclaration__ models.Fstruct
var __Fstruct_time__dummyDeclaration time.Duration

var mutexFstruct sync.Mutex

// An FstructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFstruct updateFstruct deleteFstruct
type FstructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FstructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFstruct updateFstruct
type FstructInput struct {
	// The Fstruct to submit or modify
	// in: body
	Fstruct *orm.FstructAPI
}

// GetFstructs
//
// swagger:route GET /fstructs fstructs getFstructs
//
// # Get all fstructs
//
// Responses:
// default: genericError
//
//	200: fstructDBResponse
func (controller *Controller) GetFstructs(c *gin.Context) {

	// source slice
	var fstructDBs []orm.FstructDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFstructs", "Name", stackPath)
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
	db := backRepo.BackRepoFstruct.GetDB()

	_, err := db.Find(&fstructDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fstructAPIs := make([]orm.FstructAPI, 0)

	// for each fstruct, update fields from the database nullable fields
	for idx := range fstructDBs {
		fstructDB := &fstructDBs[idx]
		_ = fstructDB
		var fstructAPI orm.FstructAPI

		// insertion point for updating fields
		fstructAPI.ID = fstructDB.ID
		fstructDB.CopyBasicFieldsToFstruct_WOP(&fstructAPI.Fstruct_WOP)
		fstructAPI.FstructPointersEncoding = fstructDB.FstructPointersEncoding
		fstructAPIs = append(fstructAPIs, fstructAPI)
	}

	c.JSON(http.StatusOK, fstructAPIs)
}

// PostFstruct
//
// swagger:route POST /fstructs fstructs postFstruct
//
// Creates a fstruct
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFstruct(c *gin.Context) {

	mutexFstruct.Lock()
	defer mutexFstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFstructs", "Name", stackPath)
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
	db := backRepo.BackRepoFstruct.GetDB()

	// Validate input
	var input orm.FstructAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fstruct
	fstructDB := orm.FstructDB{}
	fstructDB.FstructPointersEncoding = input.FstructPointersEncoding
	fstructDB.CopyBasicFieldsFromFstruct_WOP(&input.Fstruct_WOP)

	_, err = db.Create(&fstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFstruct.CheckoutPhaseOneInstance(&fstructDB)
	fstruct := backRepo.BackRepoFstruct.Map_FstructDBID_FstructPtr[fstructDB.ID]

	if fstruct != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fstruct)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fstructDB)
}

// GetFstruct
//
// swagger:route GET /fstructs/{ID} fstructs getFstruct
//
// Gets the details for a fstruct.
//
// Responses:
// default: genericError
//
//	200: fstructDBResponse
func (controller *Controller) GetFstruct(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFstruct", "Name", stackPath)
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
	db := backRepo.BackRepoFstruct.GetDB()

	// Get fstructDB in DB
	var fstructDB orm.FstructDB
	if _, err := db.First(&fstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fstructAPI orm.FstructAPI
	fstructAPI.ID = fstructDB.ID
	fstructAPI.FstructPointersEncoding = fstructDB.FstructPointersEncoding
	fstructDB.CopyBasicFieldsToFstruct_WOP(&fstructAPI.Fstruct_WOP)

	c.JSON(http.StatusOK, fstructAPI)
}

// UpdateFstruct
//
// swagger:route PATCH /fstructs/{ID} fstructs updateFstruct
//
// # Update a fstruct
//
// Responses:
// default: genericError
//
//	200: fstructDBResponse
func (controller *Controller) UpdateFstruct(c *gin.Context) {

	mutexFstruct.Lock()
	defer mutexFstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFstruct", "Name", stackPath)
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
	db := backRepo.BackRepoFstruct.GetDB()

	// Validate input
	var input orm.FstructAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fstructDB orm.FstructDB

	// fetch the fstruct
	_, err := db.First(&fstructDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fstructDB.CopyBasicFieldsFromFstruct_WOP(&input.Fstruct_WOP)
	fstructDB.FstructPointersEncoding = input.FstructPointersEncoding

	db, _ = db.Model(&fstructDB)
	_, err = db.Updates(&fstructDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fstructNew := new(models.Fstruct)
	fstructDB.CopyBasicFieldsToFstruct(fstructNew)

	// redeem pointers
	fstructDB.DecodePointers(backRepo, fstructNew)

	// get stage instance from DB instance, and call callback function
	fstructOld := backRepo.BackRepoFstruct.Map_FstructDBID_FstructPtr[fstructDB.ID]
	if fstructOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fstructOld, fstructNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fstructDB
	c.JSON(http.StatusOK, fstructDB)
}

// DeleteFstruct
//
// swagger:route DELETE /fstructs/{ID} fstructs deleteFstruct
//
// # Delete a fstruct
//
// default: genericError
//
//	200: fstructDBResponse
func (controller *Controller) DeleteFstruct(c *gin.Context) {

	mutexFstruct.Lock()
	defer mutexFstruct.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFstruct", "Name", stackPath)
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
	db := backRepo.BackRepoFstruct.GetDB()

	// Get model if exist
	var fstructDB orm.FstructDB
	if _, err := db.First(&fstructDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&fstructDB)

	// get an instance (not staged) from DB instance, and call callback function
	fstructDeleted := new(models.Fstruct)
	fstructDB.CopyBasicFieldsToFstruct(fstructDeleted)

	// get stage instance from DB instance, and call callback function
	fstructStaged := backRepo.BackRepoFstruct.Map_FstructDBID_FstructPtr[fstructDB.ID]
	if fstructStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fstructStaged, fstructDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
