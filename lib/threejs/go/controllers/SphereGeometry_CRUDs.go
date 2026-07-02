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
var __SphereGeometry__dummysDeclaration__ models.SphereGeometry
var _ = __SphereGeometry__dummysDeclaration__
var __SphereGeometry_time__dummyDeclaration time.Duration
var _ = __SphereGeometry_time__dummyDeclaration

var mutexSphereGeometry sync.Mutex

// An SphereGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSphereGeometry updateSphereGeometry deleteSphereGeometry
type SphereGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SphereGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSphereGeometry updateSphereGeometry
type SphereGeometryInput struct {
	// The SphereGeometry to submit or modify
	// in: body
	SphereGeometry *orm.SphereGeometryAPI
}

// GetSphereGeometrys
//
// swagger:route GET /spheregeometrys spheregeometrys getSphereGeometrys
//
// # Get all spheregeometrys
//
// Responses:
// default: genericError
//
//	200: spheregeometryDBResponse
func (controller *Controller) GetSphereGeometrys(c *gin.Context) {

	// source slice
	var spheregeometryDBs []orm.SphereGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSphereGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	_, err := db.Find(&spheregeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spheregeometryAPIs := make([]orm.SphereGeometryAPI, 0)

	// for each spheregeometry, update fields from the database nullable fields
	for idx := range spheregeometryDBs {
		spheregeometryDB := &spheregeometryDBs[idx]
		_ = spheregeometryDB
		var spheregeometryAPI orm.SphereGeometryAPI

		// insertion point for updating fields
		spheregeometryAPI.ID = spheregeometryDB.ID
		spheregeometryDB.CopyBasicFieldsToSphereGeometry_WOP(&spheregeometryAPI.SphereGeometry_WOP)
		spheregeometryAPI.SphereGeometryPointersEncoding = spheregeometryDB.SphereGeometryPointersEncoding
		spheregeometryAPIs = append(spheregeometryAPIs, spheregeometryAPI)
	}

	c.JSON(http.StatusOK, spheregeometryAPIs)
}

// PostSphereGeometry
//
// swagger:route POST /spheregeometrys spheregeometrys postSphereGeometry
//
// Creates a spheregeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSphereGeometry(c *gin.Context) {

	mutexSphereGeometry.Lock()
	defer mutexSphereGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSphereGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	// Validate input
	var input orm.SphereGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spheregeometry
	spheregeometryDB := orm.SphereGeometryDB{}
	spheregeometryDB.SphereGeometryPointersEncoding = input.SphereGeometryPointersEncoding
	spheregeometryDB.CopyBasicFieldsFromSphereGeometry_WOP(&input.SphereGeometry_WOP)

	_, err = db.Create(&spheregeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSphereGeometry.CheckoutPhaseOneInstance(&spheregeometryDB)
	spheregeometry := backRepo.BackRepoSphereGeometry.Map_SphereGeometryDBID_SphereGeometryPtr[spheregeometryDB.ID]

	if spheregeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spheregeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spheregeometryDB)
}

// GetSphereGeometry
//
// swagger:route GET /spheregeometrys/{ID} spheregeometrys getSphereGeometry
//
// Gets the details for a spheregeometry.
//
// Responses:
// default: genericError
//
//	200: spheregeometryDBResponse
func (controller *Controller) GetSphereGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSphereGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	// Get spheregeometryDB in DB
	var spheregeometryDB orm.SphereGeometryDB
	if _, err := db.First(&spheregeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spheregeometryAPI orm.SphereGeometryAPI
	spheregeometryAPI.ID = spheregeometryDB.ID
	spheregeometryAPI.SphereGeometryPointersEncoding = spheregeometryDB.SphereGeometryPointersEncoding
	spheregeometryDB.CopyBasicFieldsToSphereGeometry_WOP(&spheregeometryAPI.SphereGeometry_WOP)

	c.JSON(http.StatusOK, spheregeometryAPI)
}

// UpdateSphereGeometry
//
// swagger:route PATCH /spheregeometrys/{ID} spheregeometrys updateSphereGeometry
//
// # Update a spheregeometry
//
// Responses:
// default: genericError
//
//	200: spheregeometryDBResponse
func (controller *Controller) UpdateSphereGeometry(c *gin.Context) {

	mutexSphereGeometry.Lock()
	defer mutexSphereGeometry.Unlock()

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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	// Validate input
	var input orm.SphereGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spheregeometryDB orm.SphereGeometryDB

	// fetch the spheregeometry
	_, err := db.First(&spheregeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spheregeometryDB.CopyBasicFieldsFromSphereGeometry_WOP(&input.SphereGeometry_WOP)
	spheregeometryDB.SphereGeometryPointersEncoding = input.SphereGeometryPointersEncoding

	db, _ = db.Model(&spheregeometryDB)
	_, err = db.Updates(&spheregeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spheregeometryNew := new(models.SphereGeometry)
	spheregeometryDB.CopyBasicFieldsToSphereGeometry(spheregeometryNew)

	// redeem pointers
	spheregeometryDB.DecodePointers(backRepo, spheregeometryNew)

	// get stage instance from DB instance, and call callback function
	spheregeometryOld := backRepo.BackRepoSphereGeometry.Map_SphereGeometryDBID_SphereGeometryPtr[spheregeometryDB.ID]
	if spheregeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), spheregeometryOld, spheregeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spheregeometryDB
	c.JSON(http.StatusOK, spheregeometryDB)
}

// DeleteSphereGeometry
//
// swagger:route DELETE /spheregeometrys/{ID} spheregeometrys deleteSphereGeometry
//
// # Delete a spheregeometry
//
// default: genericError
//
//	200: spheregeometryDBResponse
func (controller *Controller) DeleteSphereGeometry(c *gin.Context) {

	mutexSphereGeometry.Lock()
	defer mutexSphereGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSphereGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoSphereGeometry.GetDB()

	// Get model if exist
	var spheregeometryDB orm.SphereGeometryDB
	if _, err := db.First(&spheregeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spheregeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	spheregeometryDeleted := new(models.SphereGeometry)
	spheregeometryDB.CopyBasicFieldsToSphereGeometry(spheregeometryDeleted)

	// get stage instance from DB instance, and call callback function
	spheregeometryStaged := backRepo.BackRepoSphereGeometry.Map_SphereGeometryDBID_SphereGeometryPtr[spheregeometryDB.ID]
	if spheregeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spheregeometryStaged, spheregeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
