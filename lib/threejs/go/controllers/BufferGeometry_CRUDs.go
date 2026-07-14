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
var __BufferGeometry__dummysDeclaration__ models.BufferGeometry
var _ = __BufferGeometry__dummysDeclaration__
var __BufferGeometry_time__dummyDeclaration time.Duration
var _ = __BufferGeometry_time__dummyDeclaration

var mutexBufferGeometry sync.Mutex

// An BufferGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBufferGeometry updateBufferGeometry deleteBufferGeometry
type BufferGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BufferGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBufferGeometry updateBufferGeometry
type BufferGeometryInput struct {
	// The BufferGeometry to submit or modify
	// in: body
	BufferGeometry *orm.BufferGeometryAPI
}

// GetBufferGeometrys
//
// swagger:route GET /buffergeometrys buffergeometrys getBufferGeometrys
//
// # Get all buffergeometrys
//
// Responses:
// default: genericError
//
//	200: buffergeometryDBResponse
func (controller *Controller) GetBufferGeometrys(c *gin.Context) {

	// source slice
	var buffergeometryDBs []orm.BufferGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBufferGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoBufferGeometry.GetDB()

	_, err := db.Find(&buffergeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	buffergeometryAPIs := make([]orm.BufferGeometryAPI, 0)

	// for each buffergeometry, update fields from the database nullable fields
	for idx := range buffergeometryDBs {
		buffergeometryDB := &buffergeometryDBs[idx]
		_ = buffergeometryDB
		var buffergeometryAPI orm.BufferGeometryAPI

		// insertion point for updating fields
		buffergeometryAPI.ID = buffergeometryDB.ID
		buffergeometryDB.CopyBasicFieldsToBufferGeometry_WOP(&buffergeometryAPI.BufferGeometry_WOP)
		buffergeometryAPI.BufferGeometryPointersEncoding = buffergeometryDB.BufferGeometryPointersEncoding
		buffergeometryAPIs = append(buffergeometryAPIs, buffergeometryAPI)
	}

	c.JSON(http.StatusOK, buffergeometryAPIs)
}

// PostBufferGeometry
//
// swagger:route POST /buffergeometrys buffergeometrys postBufferGeometry
//
// Creates a buffergeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBufferGeometry(c *gin.Context) {

	mutexBufferGeometry.Lock()
	defer mutexBufferGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBufferGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoBufferGeometry.GetDB()

	// Validate input
	var input orm.BufferGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create buffergeometry
	buffergeometryDB := orm.BufferGeometryDB{}
	buffergeometryDB.BufferGeometryPointersEncoding = input.BufferGeometryPointersEncoding
	buffergeometryDB.CopyBasicFieldsFromBufferGeometry_WOP(&input.BufferGeometry_WOP)

	_, err = db.Create(&buffergeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBufferGeometry.CheckoutPhaseOneInstance(&buffergeometryDB)
	buffergeometry := backRepo.BackRepoBufferGeometry.Map_BufferGeometryDBID_BufferGeometryPtr[buffergeometryDB.ID]

	if buffergeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), buffergeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, buffergeometryDB)
}

// GetBufferGeometry
//
// swagger:route GET /buffergeometrys/{ID} buffergeometrys getBufferGeometry
//
// Gets the details for a buffergeometry.
//
// Responses:
// default: genericError
//
//	200: buffergeometryDBResponse
func (controller *Controller) GetBufferGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBufferGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoBufferGeometry.GetDB()

	// Get buffergeometryDB in DB
	var buffergeometryDB orm.BufferGeometryDB
	if _, err := db.First(&buffergeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var buffergeometryAPI orm.BufferGeometryAPI
	buffergeometryAPI.ID = buffergeometryDB.ID
	buffergeometryAPI.BufferGeometryPointersEncoding = buffergeometryDB.BufferGeometryPointersEncoding
	buffergeometryDB.CopyBasicFieldsToBufferGeometry_WOP(&buffergeometryAPI.BufferGeometry_WOP)

	c.JSON(http.StatusOK, buffergeometryAPI)
}

// UpdateBufferGeometry
//
// swagger:route PATCH /buffergeometrys/{ID} buffergeometrys updateBufferGeometry
//
// # Update a buffergeometry
//
// Responses:
// default: genericError
//
//	200: buffergeometryDBResponse
func (controller *Controller) UpdateBufferGeometry(c *gin.Context) {

	mutexBufferGeometry.Lock()
	defer mutexBufferGeometry.Unlock()

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
	db := backRepo.BackRepoBufferGeometry.GetDB()

	// Validate input
	var input orm.BufferGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var buffergeometryDB orm.BufferGeometryDB

	// fetch the buffergeometry
	_, err := db.First(&buffergeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	buffergeometryDB.CopyBasicFieldsFromBufferGeometry_WOP(&input.BufferGeometry_WOP)
	buffergeometryDB.BufferGeometryPointersEncoding = input.BufferGeometryPointersEncoding

	db, _ = db.Model(&buffergeometryDB)
	_, err = db.Updates(&buffergeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	buffergeometryNew := new(models.BufferGeometry)
	buffergeometryDB.CopyBasicFieldsToBufferGeometry(buffergeometryNew)

	// redeem pointers
	buffergeometryDB.DecodePointers(backRepo, buffergeometryNew)

	// get stage instance from DB instance, and call callback function
	buffergeometryOld := backRepo.BackRepoBufferGeometry.Map_BufferGeometryDBID_BufferGeometryPtr[buffergeometryDB.ID]
	if buffergeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), buffergeometryOld, buffergeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the buffergeometryDB
	c.JSON(http.StatusOK, buffergeometryDB)
}

// DeleteBufferGeometry
//
// swagger:route DELETE /buffergeometrys/{ID} buffergeometrys deleteBufferGeometry
//
// # Delete a buffergeometry
//
// default: genericError
//
//	200: buffergeometryDBResponse
func (controller *Controller) DeleteBufferGeometry(c *gin.Context) {

	mutexBufferGeometry.Lock()
	defer mutexBufferGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBufferGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoBufferGeometry.GetDB()

	// Get model if exist
	var buffergeometryDB orm.BufferGeometryDB
	if _, err := db.First(&buffergeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&buffergeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	buffergeometryDeleted := new(models.BufferGeometry)
	buffergeometryDB.CopyBasicFieldsToBufferGeometry(buffergeometryDeleted)

	// get stage instance from DB instance, and call callback function
	buffergeometryStaged := backRepo.BackRepoBufferGeometry.Map_BufferGeometryDBID_BufferGeometryPtr[buffergeometryDB.ID]
	if buffergeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), buffergeometryStaged, buffergeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
