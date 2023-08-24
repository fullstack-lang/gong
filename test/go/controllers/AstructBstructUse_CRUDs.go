// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __AstructBstructUse__dummysDeclaration__ models.AstructBstructUse
var __AstructBstructUse_time__dummyDeclaration time.Duration

var mutexAstructBstructUse sync.Mutex

// An AstructBstructUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAstructBstructUse updateAstructBstructUse deleteAstructBstructUse
type AstructBstructUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AstructBstructUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAstructBstructUse updateAstructBstructUse
type AstructBstructUseInput struct {
	// The AstructBstructUse to submit or modify
	// in: body
	AstructBstructUse *orm.AstructBstructUseAPI
}

// GetAstructBstructUses
//
// swagger:route GET /astructbstructuses astructbstructuses getAstructBstructUses
//
// # Get all astructbstructuses
//
// Responses:
// default: genericError
//
//	200: astructbstructuseDBResponse
func (controller *Controller) GetAstructBstructUses(c *gin.Context) {

	// source slice
	var astructbstructuseDBs []orm.AstructBstructUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstructBstructUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	query := db.Find(&astructbstructuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	astructbstructuseAPIs := make([]orm.AstructBstructUseAPI, 0)

	// for each astructbstructuse, update fields from the database nullable fields
	for idx := range astructbstructuseDBs {
		astructbstructuseDB := &astructbstructuseDBs[idx]
		_ = astructbstructuseDB
		var astructbstructuseAPI orm.AstructBstructUseAPI

		// insertion point for updating fields
		astructbstructuseAPI.ID = astructbstructuseDB.ID
		astructbstructuseDB.CopyBasicFieldsToAstructBstructUse(&astructbstructuseAPI.AstructBstructUse)
		astructbstructuseAPI.AstructBstructUsePointersEnconding = astructbstructuseDB.AstructBstructUsePointersEnconding
		astructbstructuseAPIs = append(astructbstructuseAPIs, astructbstructuseAPI)
	}

	c.JSON(http.StatusOK, astructbstructuseAPIs)
}

// PostAstructBstructUse
//
// swagger:route POST /astructbstructuses astructbstructuses postAstructBstructUse
//
// Creates a astructbstructuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAstructBstructUse(c *gin.Context) {

	mutexAstructBstructUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAstructBstructUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	// Validate input
	var input orm.AstructBstructUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create astructbstructuse
	astructbstructuseDB := orm.AstructBstructUseDB{}
	astructbstructuseDB.AstructBstructUsePointersEnconding = input.AstructBstructUsePointersEnconding
	astructbstructuseDB.CopyBasicFieldsFromAstructBstructUse(&input.AstructBstructUse)

	query := db.Create(&astructbstructuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAstructBstructUse.CheckoutPhaseOneInstance(&astructbstructuseDB)
	astructbstructuse := backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUsePtr[astructbstructuseDB.ID]

	if astructbstructuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), astructbstructuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, astructbstructuseDB)

	mutexAstructBstructUse.Unlock()
}

// GetAstructBstructUse
//
// swagger:route GET /astructbstructuses/{ID} astructbstructuses getAstructBstructUse
//
// Gets the details for a astructbstructuse.
//
// Responses:
// default: genericError
//
//	200: astructbstructuseDBResponse
func (controller *Controller) GetAstructBstructUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAstructBstructUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	// Get astructbstructuseDB in DB
	var astructbstructuseDB orm.AstructBstructUseDB
	if err := db.First(&astructbstructuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var astructbstructuseAPI orm.AstructBstructUseAPI
	astructbstructuseAPI.ID = astructbstructuseDB.ID
	astructbstructuseAPI.AstructBstructUsePointersEnconding = astructbstructuseDB.AstructBstructUsePointersEnconding
	astructbstructuseDB.CopyBasicFieldsToAstructBstructUse(&astructbstructuseAPI.AstructBstructUse)

	c.JSON(http.StatusOK, astructbstructuseAPI)
}

// UpdateAstructBstructUse
//
// swagger:route PATCH /astructbstructuses/{ID} astructbstructuses updateAstructBstructUse
//
// # Update a astructbstructuse
//
// Responses:
// default: genericError
//
//	200: astructbstructuseDBResponse
func (controller *Controller) UpdateAstructBstructUse(c *gin.Context) {

	mutexAstructBstructUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAstructBstructUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	// Validate input
	var input orm.AstructBstructUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var astructbstructuseDB orm.AstructBstructUseDB

	// fetch the astructbstructuse
	query := db.First(&astructbstructuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	astructbstructuseDB.CopyBasicFieldsFromAstructBstructUse(&input.AstructBstructUse)
	astructbstructuseDB.AstructBstructUsePointersEnconding = input.AstructBstructUsePointersEnconding

	query = db.Model(&astructbstructuseDB).Updates(astructbstructuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	astructbstructuseNew := new(models.AstructBstructUse)
	astructbstructuseDB.CopyBasicFieldsToAstructBstructUse(astructbstructuseNew)

	// get stage instance from DB instance, and call callback function
	astructbstructuseOld := backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUsePtr[astructbstructuseDB.ID]
	if astructbstructuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), astructbstructuseOld, astructbstructuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the astructbstructuseDB
	c.JSON(http.StatusOK, astructbstructuseDB)

	mutexAstructBstructUse.Unlock()
}

// DeleteAstructBstructUse
//
// swagger:route DELETE /astructbstructuses/{ID} astructbstructuses deleteAstructBstructUse
//
// # Delete a astructbstructuse
//
// default: genericError
//
//	200: astructbstructuseDBResponse
func (controller *Controller) DeleteAstructBstructUse(c *gin.Context) {

	mutexAstructBstructUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAstructBstructUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoAstructBstructUse.GetDB()

	// Get model if exist
	var astructbstructuseDB orm.AstructBstructUseDB
	if err := db.First(&astructbstructuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&astructbstructuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	astructbstructuseDeleted := new(models.AstructBstructUse)
	astructbstructuseDB.CopyBasicFieldsToAstructBstructUse(astructbstructuseDeleted)

	// get stage instance from DB instance, and call callback function
	astructbstructuseStaged := backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUsePtr[astructbstructuseDB.ID]
	if astructbstructuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), astructbstructuseStaged, astructbstructuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexAstructBstructUse.Unlock()
}
