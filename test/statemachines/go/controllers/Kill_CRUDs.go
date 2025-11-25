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
var __Kill__dummysDeclaration__ models.Kill
var __Kill_time__dummyDeclaration time.Duration

var mutexKill sync.Mutex

// An KillID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKill updateKill deleteKill
type KillID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KillInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKill updateKill
type KillInput struct {
	// The Kill to submit or modify
	// in: body
	Kill *orm.KillAPI
}

// GetKills
//
// swagger:route GET /kills kills getKills
//
// # Get all kills
//
// Responses:
// default: genericError
//
//	200: killDBResponse
func (controller *Controller) GetKills(c *gin.Context) {

	// source slice
	var killDBs []orm.KillDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKills", "Name", stackPath)
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
	db := backRepo.BackRepoKill.GetDB()

	_, err := db.Find(&killDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	killAPIs := make([]orm.KillAPI, 0)

	// for each kill, update fields from the database nullable fields
	for idx := range killDBs {
		killDB := &killDBs[idx]
		_ = killDB
		var killAPI orm.KillAPI

		// insertion point for updating fields
		killAPI.ID = killDB.ID
		killDB.CopyBasicFieldsToKill_WOP(&killAPI.Kill_WOP)
		killAPI.KillPointersEncoding = killDB.KillPointersEncoding
		killAPIs = append(killAPIs, killAPI)
	}

	c.JSON(http.StatusOK, killAPIs)
}

// PostKill
//
// swagger:route POST /kills kills postKill
//
// Creates a kill
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKill(c *gin.Context) {

	mutexKill.Lock()
	defer mutexKill.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKills", "Name", stackPath)
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
	db := backRepo.BackRepoKill.GetDB()

	// Validate input
	var input orm.KillAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create kill
	killDB := orm.KillDB{}
	killDB.KillPointersEncoding = input.KillPointersEncoding
	killDB.CopyBasicFieldsFromKill_WOP(&input.Kill_WOP)

	_, err = db.Create(&killDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKill.CheckoutPhaseOneInstance(&killDB)
	kill := backRepo.BackRepoKill.Map_KillDBID_KillPtr[killDB.ID]

	if kill != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), kill)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, killDB)
}

// GetKill
//
// swagger:route GET /kills/{ID} kills getKill
//
// Gets the details for a kill.
//
// Responses:
// default: genericError
//
//	200: killDBResponse
func (controller *Controller) GetKill(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKill", "Name", stackPath)
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
	db := backRepo.BackRepoKill.GetDB()

	// Get killDB in DB
	var killDB orm.KillDB
	if _, err := db.First(&killDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var killAPI orm.KillAPI
	killAPI.ID = killDB.ID
	killAPI.KillPointersEncoding = killDB.KillPointersEncoding
	killDB.CopyBasicFieldsToKill_WOP(&killAPI.Kill_WOP)

	c.JSON(http.StatusOK, killAPI)
}

// UpdateKill
//
// swagger:route PATCH /kills/{ID} kills updateKill
//
// # Update a kill
//
// Responses:
// default: genericError
//
//	200: killDBResponse
func (controller *Controller) UpdateKill(c *gin.Context) {

	mutexKill.Lock()
	defer mutexKill.Unlock()

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
	db := backRepo.BackRepoKill.GetDB()

	// Validate input
	var input orm.KillAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var killDB orm.KillDB

	// fetch the kill
	_, err := db.First(&killDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	killDB.CopyBasicFieldsFromKill_WOP(&input.Kill_WOP)
	killDB.KillPointersEncoding = input.KillPointersEncoding

	db, _ = db.Model(&killDB)
	_, err = db.Updates(&killDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	killNew := new(models.Kill)
	killDB.CopyBasicFieldsToKill(killNew)

	// redeem pointers
	killDB.DecodePointers(backRepo, killNew)

	// get stage instance from DB instance, and call callback function
	killOld := backRepo.BackRepoKill.Map_KillDBID_KillPtr[killDB.ID]
	if killOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), killOld, killNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the killDB
	c.JSON(http.StatusOK, killDB)
}

// DeleteKill
//
// swagger:route DELETE /kills/{ID} kills deleteKill
//
// # Delete a kill
//
// default: genericError
//
//	200: killDBResponse
func (controller *Controller) DeleteKill(c *gin.Context) {

	mutexKill.Lock()
	defer mutexKill.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKill", "Name", stackPath)
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
	db := backRepo.BackRepoKill.GetDB()

	// Get model if exist
	var killDB orm.KillDB
	if _, err := db.First(&killDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&killDB)

	// get an instance (not staged) from DB instance, and call callback function
	killDeleted := new(models.Kill)
	killDB.CopyBasicFieldsToKill(killDeleted)

	// get stage instance from DB instance, and call callback function
	killStaged := backRepo.BackRepoKill.Map_KillDBID_KillPtr[killDB.ID]
	if killStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), killStaged, killDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
