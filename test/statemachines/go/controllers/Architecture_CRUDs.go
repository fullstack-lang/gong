// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Architecture__dummysDeclaration__ models.Architecture
var __Architecture_time__dummyDeclaration time.Duration

var mutexArchitecture sync.Mutex

// An ArchitectureID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getArchitecture updateArchitecture deleteArchitecture
type ArchitectureID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ArchitectureInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postArchitecture updateArchitecture
type ArchitectureInput struct {
	// The Architecture to submit or modify
	// in: body
	Architecture *orm.ArchitectureAPI
}

// GetArchitectures
//
// swagger:route GET /architectures architectures getArchitectures
//
// # Get all architectures
//
// Responses:
// default: genericError
//
//	200: architectureDBResponse
func (controller *Controller) GetArchitectures(c *gin.Context) {

	// source slice
	var architectureDBs []orm.ArchitectureDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArchitectures", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoArchitecture.GetDB()

	_, err := db.Find(&architectureDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	architectureAPIs := make([]orm.ArchitectureAPI, 0)

	// for each architecture, update fields from the database nullable fields
	for idx := range architectureDBs {
		architectureDB := &architectureDBs[idx]
		_ = architectureDB
		var architectureAPI orm.ArchitectureAPI

		// insertion point for updating fields
		architectureAPI.ID = architectureDB.ID
		architectureDB.CopyBasicFieldsToArchitecture_WOP(&architectureAPI.Architecture_WOP)
		architectureAPI.ArchitecturePointersEncoding = architectureDB.ArchitecturePointersEncoding
		architectureAPIs = append(architectureAPIs, architectureAPI)
	}

	c.JSON(http.StatusOK, architectureAPIs)
}

// PostArchitecture
//
// swagger:route POST /architectures architectures postArchitecture
//
// Creates a architecture
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostArchitecture(c *gin.Context) {

	mutexArchitecture.Lock()
	defer mutexArchitecture.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostArchitectures", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoArchitecture.GetDB()

	// Validate input
	var input orm.ArchitectureAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create architecture
	architectureDB := orm.ArchitectureDB{}
	architectureDB.ArchitecturePointersEncoding = input.ArchitecturePointersEncoding
	architectureDB.CopyBasicFieldsFromArchitecture_WOP(&input.Architecture_WOP)

	_, err = db.Create(&architectureDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoArchitecture.CheckoutPhaseOneInstance(&architectureDB)
	architecture := backRepo.BackRepoArchitecture.Map_ArchitectureDBID_ArchitecturePtr[architectureDB.ID]

	if architecture != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), architecture)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, architectureDB)
}

// GetArchitecture
//
// swagger:route GET /architectures/{ID} architectures getArchitecture
//
// Gets the details for a architecture.
//
// Responses:
// default: genericError
//
//	200: architectureDBResponse
func (controller *Controller) GetArchitecture(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArchitecture", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoArchitecture.GetDB()

	// Get architectureDB in DB
	var architectureDB orm.ArchitectureDB
	if _, err := db.First(&architectureDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var architectureAPI orm.ArchitectureAPI
	architectureAPI.ID = architectureDB.ID
	architectureAPI.ArchitecturePointersEncoding = architectureDB.ArchitecturePointersEncoding
	architectureDB.CopyBasicFieldsToArchitecture_WOP(&architectureAPI.Architecture_WOP)

	c.JSON(http.StatusOK, architectureAPI)
}

// UpdateArchitecture
//
// swagger:route PATCH /architectures/{ID} architectures updateArchitecture
//
// # Update a architecture
//
// Responses:
// default: genericError
//
//	200: architectureDBResponse
func (controller *Controller) UpdateArchitecture(c *gin.Context) {

	mutexArchitecture.Lock()
	defer mutexArchitecture.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoArchitecture.GetDB()

	// Validate input
	var input orm.ArchitectureAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var architectureDB orm.ArchitectureDB

	// fetch the architecture
	_, err := db.First(&architectureDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	architectureDB.CopyBasicFieldsFromArchitecture_WOP(&input.Architecture_WOP)
	architectureDB.ArchitecturePointersEncoding = input.ArchitecturePointersEncoding

	db, _ = db.Model(&architectureDB)
	_, err = db.Updates(&architectureDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	architectureNew := new(models.Architecture)
	architectureDB.CopyBasicFieldsToArchitecture(architectureNew)

	// redeem pointers
	architectureDB.DecodePointers(backRepo, architectureNew)

	// get stage instance from DB instance, and call callback function
	architectureOld := backRepo.BackRepoArchitecture.Map_ArchitectureDBID_ArchitecturePtr[architectureDB.ID]
	if architectureOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), architectureOld, architectureNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the architectureDB
	c.JSON(http.StatusOK, architectureDB)
}

// DeleteArchitecture
//
// swagger:route DELETE /architectures/{ID} architectures deleteArchitecture
//
// # Delete a architecture
//
// default: genericError
//
//	200: architectureDBResponse
func (controller *Controller) DeleteArchitecture(c *gin.Context) {

	mutexArchitecture.Lock()
	defer mutexArchitecture.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteArchitecture", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoArchitecture.GetDB()

	// Get model if exist
	var architectureDB orm.ArchitectureDB
	if _, err := db.First(&architectureDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&architectureDB)

	// get an instance (not staged) from DB instance, and call callback function
	architectureDeleted := new(models.Architecture)
	architectureDB.CopyBasicFieldsToArchitecture(architectureDeleted)

	// get stage instance from DB instance, and call callback function
	architectureStaged := backRepo.BackRepoArchitecture.Map_ArchitectureDBID_ArchitecturePtr[architectureDB.ID]
	if architectureStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), architectureStaged, architectureDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
