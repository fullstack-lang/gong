// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __TorusGeometry__dummysDeclaration__ models.TorusGeometry
var _ = __TorusGeometry__dummysDeclaration__
var __TorusGeometry_time__dummyDeclaration time.Duration
var _ = __TorusGeometry_time__dummyDeclaration

var mutexTorusGeometry sync.Mutex

// An TorusGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTorusGeometry updateTorusGeometry deleteTorusGeometry
type TorusGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TorusGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTorusGeometry updateTorusGeometry
type TorusGeometryInput struct {
	// The TorusGeometry to submit or modify
	// in: body
	TorusGeometry *orm.TorusGeometryAPI
}

// GetTorusGeometrys
//
// swagger:route GET /torusgeometrys torusgeometrys getTorusGeometrys
//
// # Get all torusgeometrys
//
// Responses:
// default: genericError
//
//	200: torusgeometryDBResponse
func (controller *Controller) GetTorusGeometrys(c *gin.Context) {

	// source slice
	var torusgeometryDBs []orm.TorusGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTorusGeometrys", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTorusGeometry.GetDB()

	_, err := db.Find(&torusgeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	torusgeometryAPIs := make([]orm.TorusGeometryAPI, 0)

	// for each torusgeometry, update fields from the database nullable fields
	for idx := range torusgeometryDBs {
		torusgeometryDB := &torusgeometryDBs[idx]
		_ = torusgeometryDB
		var torusgeometryAPI orm.TorusGeometryAPI

		// insertion point for updating fields
		torusgeometryAPI.ID = torusgeometryDB.ID
		torusgeometryDB.CopyBasicFieldsToTorusGeometry_WOP(&torusgeometryAPI.TorusGeometry_WOP)
		torusgeometryAPI.TorusGeometryPointersEncoding = torusgeometryDB.TorusGeometryPointersEncoding
		torusgeometryAPIs = append(torusgeometryAPIs, torusgeometryAPI)
	}

	c.JSON(http.StatusOK, torusgeometryAPIs)
}

// PostTorusGeometry
//
// swagger:route POST /torusgeometrys torusgeometrys postTorusGeometry
//
// Creates a torusgeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTorusGeometry(c *gin.Context) {

	mutexTorusGeometry.Lock()
	defer mutexTorusGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTorusGeometrys", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTorusGeometry.GetDB()

	// Validate input
	var input orm.TorusGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create torusgeometry
	torusgeometryDB := orm.TorusGeometryDB{}
	torusgeometryDB.TorusGeometryPointersEncoding = input.TorusGeometryPointersEncoding
	torusgeometryDB.CopyBasicFieldsFromTorusGeometry_WOP(&input.TorusGeometry_WOP)

	_, err = db.Create(&torusgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTorusGeometry.CheckoutPhaseOneInstance(&torusgeometryDB)
	torusgeometry := backRepo.BackRepoTorusGeometry.Map_TorusGeometryDBID_TorusGeometryPtr[torusgeometryDB.ID]

	if torusgeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), torusgeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, torusgeometryDB)
}

// GetTorusGeometry
//
// swagger:route GET /torusgeometrys/{ID} torusgeometrys getTorusGeometry
//
// Gets the details for a torusgeometry.
//
// Responses:
// default: genericError
//
//	200: torusgeometryDBResponse
func (controller *Controller) GetTorusGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTorusGeometry", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTorusGeometry.GetDB()

	// Get torusgeometryDB in DB
	var torusgeometryDB orm.TorusGeometryDB
	if _, err := db.First(&torusgeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var torusgeometryAPI orm.TorusGeometryAPI
	torusgeometryAPI.ID = torusgeometryDB.ID
	torusgeometryAPI.TorusGeometryPointersEncoding = torusgeometryDB.TorusGeometryPointersEncoding
	torusgeometryDB.CopyBasicFieldsToTorusGeometry_WOP(&torusgeometryAPI.TorusGeometry_WOP)

	c.JSON(http.StatusOK, torusgeometryAPI)
}

// UpdateTorusGeometry
//
// swagger:route PATCH /torusgeometrys/{ID} torusgeometrys updateTorusGeometry
//
// # Update a torusgeometry
//
// Responses:
// default: genericError
//
//	200: torusgeometryDBResponse
func (controller *Controller) UpdateTorusGeometry(c *gin.Context) {

	mutexTorusGeometry.Lock()
	defer mutexTorusGeometry.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTorusGeometry.GetDB()

	// Validate input
	var input orm.TorusGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var torusgeometryDB orm.TorusGeometryDB

	// fetch the torusgeometry
	_, err := db.First(&torusgeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	torusgeometryDB.CopyBasicFieldsFromTorusGeometry_WOP(&input.TorusGeometry_WOP)
	torusgeometryDB.TorusGeometryPointersEncoding = input.TorusGeometryPointersEncoding

	db, _ = db.Model(&torusgeometryDB)
	_, err = db.Updates(&torusgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	torusgeometryNew := new(models.TorusGeometry)
	torusgeometryDB.CopyBasicFieldsToTorusGeometry(torusgeometryNew)

	// redeem pointers
	torusgeometryDB.DecodePointers(backRepo, torusgeometryNew)

	// get stage instance from DB instance, and call callback function
	torusgeometryOld := backRepo.BackRepoTorusGeometry.Map_TorusGeometryDBID_TorusGeometryPtr[torusgeometryDB.ID]
	if torusgeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), torusgeometryOld, torusgeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the torusgeometryDB
	c.JSON(http.StatusOK, torusgeometryDB)
}

// DeleteTorusGeometry
//
// swagger:route DELETE /torusgeometrys/{ID} torusgeometrys deleteTorusGeometry
//
// # Delete a torusgeometry
//
// default: genericError
//
//	200: torusgeometryDBResponse
func (controller *Controller) DeleteTorusGeometry(c *gin.Context) {

	mutexTorusGeometry.Lock()
	defer mutexTorusGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTorusGeometry", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTorusGeometry.GetDB()

	// Get model if exist
	var torusgeometryDB orm.TorusGeometryDB
	if _, err := db.First(&torusgeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&torusgeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	torusgeometryDeleted := new(models.TorusGeometry)
	torusgeometryDB.CopyBasicFieldsToTorusGeometry(torusgeometryDeleted)

	// get stage instance from DB instance, and call callback function
	torusgeometryStaged := backRepo.BackRepoTorusGeometry.Map_TorusGeometryDBID_TorusGeometryPtr[torusgeometryDB.ID]
	if torusgeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), torusgeometryStaged, torusgeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
