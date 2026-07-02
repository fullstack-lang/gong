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
var __CylinderGeometry__dummysDeclaration__ models.CylinderGeometry
var _ = __CylinderGeometry__dummysDeclaration__
var __CylinderGeometry_time__dummyDeclaration time.Duration
var _ = __CylinderGeometry_time__dummyDeclaration

var mutexCylinderGeometry sync.Mutex

// An CylinderGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCylinderGeometry updateCylinderGeometry deleteCylinderGeometry
type CylinderGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CylinderGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCylinderGeometry updateCylinderGeometry
type CylinderGeometryInput struct {
	// The CylinderGeometry to submit or modify
	// in: body
	CylinderGeometry *orm.CylinderGeometryAPI
}

// GetCylinderGeometrys
//
// swagger:route GET /cylindergeometrys cylindergeometrys getCylinderGeometrys
//
// # Get all cylindergeometrys
//
// Responses:
// default: genericError
//
//	200: cylindergeometryDBResponse
func (controller *Controller) GetCylinderGeometrys(c *gin.Context) {

	// source slice
	var cylindergeometryDBs []orm.CylinderGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCylinderGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoCylinderGeometry.GetDB()

	_, err := db.Find(&cylindergeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cylindergeometryAPIs := make([]orm.CylinderGeometryAPI, 0)

	// for each cylindergeometry, update fields from the database nullable fields
	for idx := range cylindergeometryDBs {
		cylindergeometryDB := &cylindergeometryDBs[idx]
		_ = cylindergeometryDB
		var cylindergeometryAPI orm.CylinderGeometryAPI

		// insertion point for updating fields
		cylindergeometryAPI.ID = cylindergeometryDB.ID
		cylindergeometryDB.CopyBasicFieldsToCylinderGeometry_WOP(&cylindergeometryAPI.CylinderGeometry_WOP)
		cylindergeometryAPI.CylinderGeometryPointersEncoding = cylindergeometryDB.CylinderGeometryPointersEncoding
		cylindergeometryAPIs = append(cylindergeometryAPIs, cylindergeometryAPI)
	}

	c.JSON(http.StatusOK, cylindergeometryAPIs)
}

// PostCylinderGeometry
//
// swagger:route POST /cylindergeometrys cylindergeometrys postCylinderGeometry
//
// Creates a cylindergeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCylinderGeometry(c *gin.Context) {

	mutexCylinderGeometry.Lock()
	defer mutexCylinderGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCylinderGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoCylinderGeometry.GetDB()

	// Validate input
	var input orm.CylinderGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cylindergeometry
	cylindergeometryDB := orm.CylinderGeometryDB{}
	cylindergeometryDB.CylinderGeometryPointersEncoding = input.CylinderGeometryPointersEncoding
	cylindergeometryDB.CopyBasicFieldsFromCylinderGeometry_WOP(&input.CylinderGeometry_WOP)

	_, err = db.Create(&cylindergeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCylinderGeometry.CheckoutPhaseOneInstance(&cylindergeometryDB)
	cylindergeometry := backRepo.BackRepoCylinderGeometry.Map_CylinderGeometryDBID_CylinderGeometryPtr[cylindergeometryDB.ID]

	if cylindergeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cylindergeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cylindergeometryDB)
}

// GetCylinderGeometry
//
// swagger:route GET /cylindergeometrys/{ID} cylindergeometrys getCylinderGeometry
//
// Gets the details for a cylindergeometry.
//
// Responses:
// default: genericError
//
//	200: cylindergeometryDBResponse
func (controller *Controller) GetCylinderGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCylinderGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoCylinderGeometry.GetDB()

	// Get cylindergeometryDB in DB
	var cylindergeometryDB orm.CylinderGeometryDB
	if _, err := db.First(&cylindergeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cylindergeometryAPI orm.CylinderGeometryAPI
	cylindergeometryAPI.ID = cylindergeometryDB.ID
	cylindergeometryAPI.CylinderGeometryPointersEncoding = cylindergeometryDB.CylinderGeometryPointersEncoding
	cylindergeometryDB.CopyBasicFieldsToCylinderGeometry_WOP(&cylindergeometryAPI.CylinderGeometry_WOP)

	c.JSON(http.StatusOK, cylindergeometryAPI)
}

// UpdateCylinderGeometry
//
// swagger:route PATCH /cylindergeometrys/{ID} cylindergeometrys updateCylinderGeometry
//
// # Update a cylindergeometry
//
// Responses:
// default: genericError
//
//	200: cylindergeometryDBResponse
func (controller *Controller) UpdateCylinderGeometry(c *gin.Context) {

	mutexCylinderGeometry.Lock()
	defer mutexCylinderGeometry.Unlock()

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
	db := backRepo.BackRepoCylinderGeometry.GetDB()

	// Validate input
	var input orm.CylinderGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cylindergeometryDB orm.CylinderGeometryDB

	// fetch the cylindergeometry
	_, err := db.First(&cylindergeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cylindergeometryDB.CopyBasicFieldsFromCylinderGeometry_WOP(&input.CylinderGeometry_WOP)
	cylindergeometryDB.CylinderGeometryPointersEncoding = input.CylinderGeometryPointersEncoding

	db, _ = db.Model(&cylindergeometryDB)
	_, err = db.Updates(&cylindergeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cylindergeometryNew := new(models.CylinderGeometry)
	cylindergeometryDB.CopyBasicFieldsToCylinderGeometry(cylindergeometryNew)

	// redeem pointers
	cylindergeometryDB.DecodePointers(backRepo, cylindergeometryNew)

	// get stage instance from DB instance, and call callback function
	cylindergeometryOld := backRepo.BackRepoCylinderGeometry.Map_CylinderGeometryDBID_CylinderGeometryPtr[cylindergeometryDB.ID]
	if cylindergeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), cylindergeometryOld, cylindergeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cylindergeometryDB
	c.JSON(http.StatusOK, cylindergeometryDB)
}

// DeleteCylinderGeometry
//
// swagger:route DELETE /cylindergeometrys/{ID} cylindergeometrys deleteCylinderGeometry
//
// # Delete a cylindergeometry
//
// default: genericError
//
//	200: cylindergeometryDBResponse
func (controller *Controller) DeleteCylinderGeometry(c *gin.Context) {

	mutexCylinderGeometry.Lock()
	defer mutexCylinderGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCylinderGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoCylinderGeometry.GetDB()

	// Get model if exist
	var cylindergeometryDB orm.CylinderGeometryDB
	if _, err := db.First(&cylindergeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&cylindergeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	cylindergeometryDeleted := new(models.CylinderGeometry)
	cylindergeometryDB.CopyBasicFieldsToCylinderGeometry(cylindergeometryDeleted)

	// get stage instance from DB instance, and call callback function
	cylindergeometryStaged := backRepo.BackRepoCylinderGeometry.Map_CylinderGeometryDBID_CylinderGeometryPtr[cylindergeometryDB.ID]
	if cylindergeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cylindergeometryStaged, cylindergeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
