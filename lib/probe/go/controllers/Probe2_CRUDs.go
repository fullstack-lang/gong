// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/probe/go/models"
	"github.com/fullstack-lang/gong/lib/probe/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Probe2__dummysDeclaration__ models.Probe2
var __Probe2_time__dummyDeclaration time.Duration

var mutexProbe2 sync.Mutex

// An Probe2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getProbe2 updateProbe2 deleteProbe2
type Probe2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Probe2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postProbe2 updateProbe2
type Probe2Input struct {
	// The Probe2 to submit or modify
	// in: body
	Probe2 *orm.Probe2API
}

// GetProbe2s
//
// swagger:route GET /probe2s probe2s getProbe2s
//
// # Get all probe2s
//
// Responses:
// default: genericError
//
//	200: probe2DBResponse
func (controller *Controller) GetProbe2s(c *gin.Context) {

	// source slice
	var probe2DBs []orm.Probe2DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetProbe2s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/probe/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoProbe2.GetDB()

	_, err := db.Find(&probe2DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	probe2APIs := make([]orm.Probe2API, 0)

	// for each probe2, update fields from the database nullable fields
	for idx := range probe2DBs {
		probe2DB := &probe2DBs[idx]
		_ = probe2DB
		var probe2API orm.Probe2API

		// insertion point for updating fields
		probe2API.ID = probe2DB.ID
		probe2DB.CopyBasicFieldsToProbe2_WOP(&probe2API.Probe2_WOP)
		probe2API.Probe2PointersEncoding = probe2DB.Probe2PointersEncoding
		probe2APIs = append(probe2APIs, probe2API)
	}

	c.JSON(http.StatusOK, probe2APIs)
}

// PostProbe2
//
// swagger:route POST /probe2s probe2s postProbe2
//
// Creates a probe2
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostProbe2(c *gin.Context) {

	mutexProbe2.Lock()
	defer mutexProbe2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostProbe2s", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/probe/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoProbe2.GetDB()

	// Validate input
	var input orm.Probe2API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create probe2
	probe2DB := orm.Probe2DB{}
	probe2DB.Probe2PointersEncoding = input.Probe2PointersEncoding
	probe2DB.CopyBasicFieldsFromProbe2_WOP(&input.Probe2_WOP)

	_, err = db.Create(&probe2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoProbe2.CheckoutPhaseOneInstance(&probe2DB)
	probe2 := backRepo.BackRepoProbe2.Map_Probe2DBID_Probe2Ptr[probe2DB.ID]

	if probe2 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), probe2)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, probe2DB)
}

// GetProbe2
//
// swagger:route GET /probe2s/{ID} probe2s getProbe2
//
// Gets the details for a probe2.
//
// Responses:
// default: genericError
//
//	200: probe2DBResponse
func (controller *Controller) GetProbe2(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetProbe2", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/probe/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoProbe2.GetDB()

	// Get probe2DB in DB
	var probe2DB orm.Probe2DB
	if _, err := db.First(&probe2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var probe2API orm.Probe2API
	probe2API.ID = probe2DB.ID
	probe2API.Probe2PointersEncoding = probe2DB.Probe2PointersEncoding
	probe2DB.CopyBasicFieldsToProbe2_WOP(&probe2API.Probe2_WOP)

	c.JSON(http.StatusOK, probe2API)
}

// UpdateProbe2
//
// swagger:route PATCH /probe2s/{ID} probe2s updateProbe2
//
// # Update a probe2
//
// Responses:
// default: genericError
//
//	200: probe2DBResponse
func (controller *Controller) UpdateProbe2(c *gin.Context) {

	mutexProbe2.Lock()
	defer mutexProbe2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateProbe2", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/probe/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoProbe2.GetDB()

	// Validate input
	var input orm.Probe2API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var probe2DB orm.Probe2DB

	// fetch the probe2
	_, err := db.First(&probe2DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	probe2DB.CopyBasicFieldsFromProbe2_WOP(&input.Probe2_WOP)
	probe2DB.Probe2PointersEncoding = input.Probe2PointersEncoding

	db, _ = db.Model(&probe2DB)
	_, err = db.Updates(&probe2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	probe2New := new(models.Probe2)
	probe2DB.CopyBasicFieldsToProbe2(probe2New)

	// redeem pointers
	probe2DB.DecodePointers(backRepo, probe2New)

	// get stage instance from DB instance, and call callback function
	probe2Old := backRepo.BackRepoProbe2.Map_Probe2DBID_Probe2Ptr[probe2DB.ID]
	if probe2Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), probe2Old, probe2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the probe2DB
	c.JSON(http.StatusOK, probe2DB)
}

// DeleteProbe2
//
// swagger:route DELETE /probe2s/{ID} probe2s deleteProbe2
//
// # Delete a probe2
//
// default: genericError
//
//	200: probe2DBResponse
func (controller *Controller) DeleteProbe2(c *gin.Context) {

	mutexProbe2.Lock()
	defer mutexProbe2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteProbe2", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/probe/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoProbe2.GetDB()

	// Get model if exist
	var probe2DB orm.Probe2DB
	if _, err := db.First(&probe2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&probe2DB)

	// get an instance (not staged) from DB instance, and call callback function
	probe2Deleted := new(models.Probe2)
	probe2DB.CopyBasicFieldsToProbe2(probe2Deleted)

	// get stage instance from DB instance, and call callback function
	probe2Staged := backRepo.BackRepoProbe2.Map_Probe2DBID_Probe2Ptr[probe2DB.ID]
	if probe2Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), probe2Staged, probe2Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
