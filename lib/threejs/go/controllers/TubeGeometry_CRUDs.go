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
var __TubeGeometry__dummysDeclaration__ models.TubeGeometry
var _ = __TubeGeometry__dummysDeclaration__
var __TubeGeometry_time__dummyDeclaration time.Duration
var _ = __TubeGeometry_time__dummyDeclaration

var mutexTubeGeometry sync.Mutex

// An TubeGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTubeGeometry updateTubeGeometry deleteTubeGeometry
type TubeGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TubeGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTubeGeometry updateTubeGeometry
type TubeGeometryInput struct {
	// The TubeGeometry to submit or modify
	// in: body
	TubeGeometry *orm.TubeGeometryAPI
}

// GetTubeGeometrys
//
// swagger:route GET /tubegeometrys tubegeometrys getTubeGeometrys
//
// # Get all tubegeometrys
//
// Responses:
// default: genericError
//
//	200: tubegeometryDBResponse
func (controller *Controller) GetTubeGeometrys(c *gin.Context) {

	// source slice
	var tubegeometryDBs []orm.TubeGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTubeGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	_, err := db.Find(&tubegeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tubegeometryAPIs := make([]orm.TubeGeometryAPI, 0)

	// for each tubegeometry, update fields from the database nullable fields
	for idx := range tubegeometryDBs {
		tubegeometryDB := &tubegeometryDBs[idx]
		_ = tubegeometryDB
		var tubegeometryAPI orm.TubeGeometryAPI

		// insertion point for updating fields
		tubegeometryAPI.ID = tubegeometryDB.ID
		tubegeometryDB.CopyBasicFieldsToTubeGeometry_WOP(&tubegeometryAPI.TubeGeometry_WOP)
		tubegeometryAPI.TubeGeometryPointersEncoding = tubegeometryDB.TubeGeometryPointersEncoding
		tubegeometryAPIs = append(tubegeometryAPIs, tubegeometryAPI)
	}

	c.JSON(http.StatusOK, tubegeometryAPIs)
}

// PostTubeGeometry
//
// swagger:route POST /tubegeometrys tubegeometrys postTubeGeometry
//
// Creates a tubegeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTubeGeometry(c *gin.Context) {

	mutexTubeGeometry.Lock()
	defer mutexTubeGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTubeGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	// Validate input
	var input orm.TubeGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tubegeometry
	tubegeometryDB := orm.TubeGeometryDB{}
	tubegeometryDB.TubeGeometryPointersEncoding = input.TubeGeometryPointersEncoding
	tubegeometryDB.CopyBasicFieldsFromTubeGeometry_WOP(&input.TubeGeometry_WOP)

	_, err = db.Create(&tubegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTubeGeometry.CheckoutPhaseOneInstance(&tubegeometryDB)
	tubegeometry := backRepo.BackRepoTubeGeometry.Map_TubeGeometryDBID_TubeGeometryPtr[tubegeometryDB.ID]

	if tubegeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tubegeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tubegeometryDB)
}

// GetTubeGeometry
//
// swagger:route GET /tubegeometrys/{ID} tubegeometrys getTubeGeometry
//
// Gets the details for a tubegeometry.
//
// Responses:
// default: genericError
//
//	200: tubegeometryDBResponse
func (controller *Controller) GetTubeGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTubeGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	// Get tubegeometryDB in DB
	var tubegeometryDB orm.TubeGeometryDB
	if _, err := db.First(&tubegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tubegeometryAPI orm.TubeGeometryAPI
	tubegeometryAPI.ID = tubegeometryDB.ID
	tubegeometryAPI.TubeGeometryPointersEncoding = tubegeometryDB.TubeGeometryPointersEncoding
	tubegeometryDB.CopyBasicFieldsToTubeGeometry_WOP(&tubegeometryAPI.TubeGeometry_WOP)

	c.JSON(http.StatusOK, tubegeometryAPI)
}

// UpdateTubeGeometry
//
// swagger:route PATCH /tubegeometrys/{ID} tubegeometrys updateTubeGeometry
//
// # Update a tubegeometry
//
// Responses:
// default: genericError
//
//	200: tubegeometryDBResponse
func (controller *Controller) UpdateTubeGeometry(c *gin.Context) {

	mutexTubeGeometry.Lock()
	defer mutexTubeGeometry.Unlock()

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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	// Validate input
	var input orm.TubeGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tubegeometryDB orm.TubeGeometryDB

	// fetch the tubegeometry
	_, err := db.First(&tubegeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tubegeometryDB.CopyBasicFieldsFromTubeGeometry_WOP(&input.TubeGeometry_WOP)
	tubegeometryDB.TubeGeometryPointersEncoding = input.TubeGeometryPointersEncoding

	db, _ = db.Model(&tubegeometryDB)
	_, err = db.Updates(&tubegeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tubegeometryNew := new(models.TubeGeometry)
	tubegeometryDB.CopyBasicFieldsToTubeGeometry(tubegeometryNew)

	// redeem pointers
	tubegeometryDB.DecodePointers(backRepo, tubegeometryNew)

	// get stage instance from DB instance, and call callback function
	tubegeometryOld := backRepo.BackRepoTubeGeometry.Map_TubeGeometryDBID_TubeGeometryPtr[tubegeometryDB.ID]
	if tubegeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), tubegeometryOld, tubegeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tubegeometryDB
	c.JSON(http.StatusOK, tubegeometryDB)
}

// DeleteTubeGeometry
//
// swagger:route DELETE /tubegeometrys/{ID} tubegeometrys deleteTubeGeometry
//
// # Delete a tubegeometry
//
// default: genericError
//
//	200: tubegeometryDBResponse
func (controller *Controller) DeleteTubeGeometry(c *gin.Context) {

	mutexTubeGeometry.Lock()
	defer mutexTubeGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTubeGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoTubeGeometry.GetDB()

	// Get model if exist
	var tubegeometryDB orm.TubeGeometryDB
	if _, err := db.First(&tubegeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tubegeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	tubegeometryDeleted := new(models.TubeGeometry)
	tubegeometryDB.CopyBasicFieldsToTubeGeometry(tubegeometryDeleted)

	// get stage instance from DB instance, and call callback function
	tubegeometryStaged := backRepo.BackRepoTubeGeometry.Map_TubeGeometryDBID_TubeGeometryPtr[tubegeometryDB.ID]
	if tubegeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tubegeometryStaged, tubegeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
