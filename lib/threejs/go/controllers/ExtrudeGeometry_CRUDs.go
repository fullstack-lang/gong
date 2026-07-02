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
var __ExtrudeGeometry__dummysDeclaration__ models.ExtrudeGeometry
var _ = __ExtrudeGeometry__dummysDeclaration__
var __ExtrudeGeometry_time__dummyDeclaration time.Duration
var _ = __ExtrudeGeometry_time__dummyDeclaration

var mutexExtrudeGeometry sync.Mutex

// An ExtrudeGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getExtrudeGeometry updateExtrudeGeometry deleteExtrudeGeometry
type ExtrudeGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExtrudeGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postExtrudeGeometry updateExtrudeGeometry
type ExtrudeGeometryInput struct {
	// The ExtrudeGeometry to submit or modify
	// in: body
	ExtrudeGeometry *orm.ExtrudeGeometryAPI
}

// GetExtrudeGeometrys
//
// swagger:route GET /extrudegeometrys extrudegeometrys getExtrudeGeometrys
//
// # Get all extrudegeometrys
//
// Responses:
// default: genericError
//
//	200: extrudegeometryDBResponse
func (controller *Controller) GetExtrudeGeometrys(c *gin.Context) {

	// source slice
	var extrudegeometryDBs []orm.ExtrudeGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtrudeGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	_, err := db.Find(&extrudegeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	extrudegeometryAPIs := make([]orm.ExtrudeGeometryAPI, 0)

	// for each extrudegeometry, update fields from the database nullable fields
	for idx := range extrudegeometryDBs {
		extrudegeometryDB := &extrudegeometryDBs[idx]
		_ = extrudegeometryDB
		var extrudegeometryAPI orm.ExtrudeGeometryAPI

		// insertion point for updating fields
		extrudegeometryAPI.ID = extrudegeometryDB.ID
		extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry_WOP(&extrudegeometryAPI.ExtrudeGeometry_WOP)
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding = extrudegeometryDB.ExtrudeGeometryPointersEncoding
		extrudegeometryAPIs = append(extrudegeometryAPIs, extrudegeometryAPI)
	}

	c.JSON(http.StatusOK, extrudegeometryAPIs)
}

// PostExtrudeGeometry
//
// swagger:route POST /extrudegeometrys extrudegeometrys postExtrudeGeometry
//
// Creates a extrudegeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostExtrudeGeometry(c *gin.Context) {

	mutexExtrudeGeometry.Lock()
	defer mutexExtrudeGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostExtrudeGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	// Validate input
	var input orm.ExtrudeGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create extrudegeometry
	extrudegeometryDB := orm.ExtrudeGeometryDB{}
	extrudegeometryDB.ExtrudeGeometryPointersEncoding = input.ExtrudeGeometryPointersEncoding
	extrudegeometryDB.CopyBasicFieldsFromExtrudeGeometry_WOP(&input.ExtrudeGeometry_WOP)

	_, err = db.Create(&extrudegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoExtrudeGeometry.CheckoutPhaseOneInstance(&extrudegeometryDB)
	extrudegeometry := backRepo.BackRepoExtrudeGeometry.Map_ExtrudeGeometryDBID_ExtrudeGeometryPtr[extrudegeometryDB.ID]

	if extrudegeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), extrudegeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, extrudegeometryDB)
}

// GetExtrudeGeometry
//
// swagger:route GET /extrudegeometrys/{ID} extrudegeometrys getExtrudeGeometry
//
// Gets the details for a extrudegeometry.
//
// Responses:
// default: genericError
//
//	200: extrudegeometryDBResponse
func (controller *Controller) GetExtrudeGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtrudeGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	// Get extrudegeometryDB in DB
	var extrudegeometryDB orm.ExtrudeGeometryDB
	if _, err := db.First(&extrudegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var extrudegeometryAPI orm.ExtrudeGeometryAPI
	extrudegeometryAPI.ID = extrudegeometryDB.ID
	extrudegeometryAPI.ExtrudeGeometryPointersEncoding = extrudegeometryDB.ExtrudeGeometryPointersEncoding
	extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry_WOP(&extrudegeometryAPI.ExtrudeGeometry_WOP)

	c.JSON(http.StatusOK, extrudegeometryAPI)
}

// UpdateExtrudeGeometry
//
// swagger:route PATCH /extrudegeometrys/{ID} extrudegeometrys updateExtrudeGeometry
//
// # Update a extrudegeometry
//
// Responses:
// default: genericError
//
//	200: extrudegeometryDBResponse
func (controller *Controller) UpdateExtrudeGeometry(c *gin.Context) {

	mutexExtrudeGeometry.Lock()
	defer mutexExtrudeGeometry.Unlock()

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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	// Validate input
	var input orm.ExtrudeGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var extrudegeometryDB orm.ExtrudeGeometryDB

	// fetch the extrudegeometry
	_, err := db.First(&extrudegeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	extrudegeometryDB.CopyBasicFieldsFromExtrudeGeometry_WOP(&input.ExtrudeGeometry_WOP)
	extrudegeometryDB.ExtrudeGeometryPointersEncoding = input.ExtrudeGeometryPointersEncoding

	db, _ = db.Model(&extrudegeometryDB)
	_, err = db.Updates(&extrudegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	extrudegeometryNew := new(models.ExtrudeGeometry)
	extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry(extrudegeometryNew)

	// redeem pointers
	extrudegeometryDB.DecodePointers(backRepo, extrudegeometryNew)

	// get stage instance from DB instance, and call callback function
	extrudegeometryOld := backRepo.BackRepoExtrudeGeometry.Map_ExtrudeGeometryDBID_ExtrudeGeometryPtr[extrudegeometryDB.ID]
	if extrudegeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), extrudegeometryOld, extrudegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the extrudegeometryDB
	c.JSON(http.StatusOK, extrudegeometryDB)
}

// DeleteExtrudeGeometry
//
// swagger:route DELETE /extrudegeometrys/{ID} extrudegeometrys deleteExtrudeGeometry
//
// # Delete a extrudegeometry
//
// default: genericError
//
//	200: extrudegeometryDBResponse
func (controller *Controller) DeleteExtrudeGeometry(c *gin.Context) {

	mutexExtrudeGeometry.Lock()
	defer mutexExtrudeGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteExtrudeGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoExtrudeGeometry.GetDB()

	// Get model if exist
	var extrudegeometryDB orm.ExtrudeGeometryDB
	if _, err := db.First(&extrudegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&extrudegeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	extrudegeometryDeleted := new(models.ExtrudeGeometry)
	extrudegeometryDB.CopyBasicFieldsToExtrudeGeometry(extrudegeometryDeleted)

	// get stage instance from DB instance, and call callback function
	extrudegeometryStaged := backRepo.BackRepoExtrudeGeometry.Map_ExtrudeGeometryDBID_ExtrudeGeometryPtr[extrudegeometryDB.ID]
	if extrudegeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), extrudegeometryStaged, extrudegeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
