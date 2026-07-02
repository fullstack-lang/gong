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
var __PlaneGeometry__dummysDeclaration__ models.PlaneGeometry
var _ = __PlaneGeometry__dummysDeclaration__
var __PlaneGeometry_time__dummyDeclaration time.Duration
var _ = __PlaneGeometry_time__dummyDeclaration

var mutexPlaneGeometry sync.Mutex

// An PlaneGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPlaneGeometry updatePlaneGeometry deletePlaneGeometry
type PlaneGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PlaneGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPlaneGeometry updatePlaneGeometry
type PlaneGeometryInput struct {
	// The PlaneGeometry to submit or modify
	// in: body
	PlaneGeometry *orm.PlaneGeometryAPI
}

// GetPlaneGeometrys
//
// swagger:route GET /planegeometrys planegeometrys getPlaneGeometrys
//
// # Get all planegeometrys
//
// Responses:
// default: genericError
//
//	200: planegeometryDBResponse
func (controller *Controller) GetPlaneGeometrys(c *gin.Context) {

	// source slice
	var planegeometryDBs []orm.PlaneGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlaneGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	_, err := db.Find(&planegeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	planegeometryAPIs := make([]orm.PlaneGeometryAPI, 0)

	// for each planegeometry, update fields from the database nullable fields
	for idx := range planegeometryDBs {
		planegeometryDB := &planegeometryDBs[idx]
		_ = planegeometryDB
		var planegeometryAPI orm.PlaneGeometryAPI

		// insertion point for updating fields
		planegeometryAPI.ID = planegeometryDB.ID
		planegeometryDB.CopyBasicFieldsToPlaneGeometry_WOP(&planegeometryAPI.PlaneGeometry_WOP)
		planegeometryAPI.PlaneGeometryPointersEncoding = planegeometryDB.PlaneGeometryPointersEncoding
		planegeometryAPIs = append(planegeometryAPIs, planegeometryAPI)
	}

	c.JSON(http.StatusOK, planegeometryAPIs)
}

// PostPlaneGeometry
//
// swagger:route POST /planegeometrys planegeometrys postPlaneGeometry
//
// Creates a planegeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPlaneGeometry(c *gin.Context) {

	mutexPlaneGeometry.Lock()
	defer mutexPlaneGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPlaneGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	// Validate input
	var input orm.PlaneGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create planegeometry
	planegeometryDB := orm.PlaneGeometryDB{}
	planegeometryDB.PlaneGeometryPointersEncoding = input.PlaneGeometryPointersEncoding
	planegeometryDB.CopyBasicFieldsFromPlaneGeometry_WOP(&input.PlaneGeometry_WOP)

	_, err = db.Create(&planegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPlaneGeometry.CheckoutPhaseOneInstance(&planegeometryDB)
	planegeometry := backRepo.BackRepoPlaneGeometry.Map_PlaneGeometryDBID_PlaneGeometryPtr[planegeometryDB.ID]

	if planegeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), planegeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, planegeometryDB)
}

// GetPlaneGeometry
//
// swagger:route GET /planegeometrys/{ID} planegeometrys getPlaneGeometry
//
// Gets the details for a planegeometry.
//
// Responses:
// default: genericError
//
//	200: planegeometryDBResponse
func (controller *Controller) GetPlaneGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlaneGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	// Get planegeometryDB in DB
	var planegeometryDB orm.PlaneGeometryDB
	if _, err := db.First(&planegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var planegeometryAPI orm.PlaneGeometryAPI
	planegeometryAPI.ID = planegeometryDB.ID
	planegeometryAPI.PlaneGeometryPointersEncoding = planegeometryDB.PlaneGeometryPointersEncoding
	planegeometryDB.CopyBasicFieldsToPlaneGeometry_WOP(&planegeometryAPI.PlaneGeometry_WOP)

	c.JSON(http.StatusOK, planegeometryAPI)
}

// UpdatePlaneGeometry
//
// swagger:route PATCH /planegeometrys/{ID} planegeometrys updatePlaneGeometry
//
// # Update a planegeometry
//
// Responses:
// default: genericError
//
//	200: planegeometryDBResponse
func (controller *Controller) UpdatePlaneGeometry(c *gin.Context) {

	mutexPlaneGeometry.Lock()
	defer mutexPlaneGeometry.Unlock()

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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	// Validate input
	var input orm.PlaneGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var planegeometryDB orm.PlaneGeometryDB

	// fetch the planegeometry
	_, err := db.First(&planegeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	planegeometryDB.CopyBasicFieldsFromPlaneGeometry_WOP(&input.PlaneGeometry_WOP)
	planegeometryDB.PlaneGeometryPointersEncoding = input.PlaneGeometryPointersEncoding

	db, _ = db.Model(&planegeometryDB)
	_, err = db.Updates(&planegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	planegeometryNew := new(models.PlaneGeometry)
	planegeometryDB.CopyBasicFieldsToPlaneGeometry(planegeometryNew)

	// redeem pointers
	planegeometryDB.DecodePointers(backRepo, planegeometryNew)

	// get stage instance from DB instance, and call callback function
	planegeometryOld := backRepo.BackRepoPlaneGeometry.Map_PlaneGeometryDBID_PlaneGeometryPtr[planegeometryDB.ID]
	if planegeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), planegeometryOld, planegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the planegeometryDB
	c.JSON(http.StatusOK, planegeometryDB)
}

// DeletePlaneGeometry
//
// swagger:route DELETE /planegeometrys/{ID} planegeometrys deletePlaneGeometry
//
// # Delete a planegeometry
//
// default: genericError
//
//	200: planegeometryDBResponse
func (controller *Controller) DeletePlaneGeometry(c *gin.Context) {

	mutexPlaneGeometry.Lock()
	defer mutexPlaneGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePlaneGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoPlaneGeometry.GetDB()

	// Get model if exist
	var planegeometryDB orm.PlaneGeometryDB
	if _, err := db.First(&planegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&planegeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	planegeometryDeleted := new(models.PlaneGeometry)
	planegeometryDB.CopyBasicFieldsToPlaneGeometry(planegeometryDeleted)

	// get stage instance from DB instance, and call callback function
	planegeometryStaged := backRepo.BackRepoPlaneGeometry.Map_PlaneGeometryDBID_PlaneGeometryPtr[planegeometryDB.ID]
	if planegeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), planegeometryStaged, planegeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
