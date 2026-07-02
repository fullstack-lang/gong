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
var __BoxGeometry__dummysDeclaration__ models.BoxGeometry
var _ = __BoxGeometry__dummysDeclaration__
var __BoxGeometry_time__dummyDeclaration time.Duration
var _ = __BoxGeometry_time__dummyDeclaration

var mutexBoxGeometry sync.Mutex

// An BoxGeometryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBoxGeometry updateBoxGeometry deleteBoxGeometry
type BoxGeometryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BoxGeometryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBoxGeometry updateBoxGeometry
type BoxGeometryInput struct {
	// The BoxGeometry to submit or modify
	// in: body
	BoxGeometry *orm.BoxGeometryAPI
}

// GetBoxGeometrys
//
// swagger:route GET /boxgeometrys boxgeometrys getBoxGeometrys
//
// # Get all boxgeometrys
//
// Responses:
// default: genericError
//
//	200: boxgeometryDBResponse
func (controller *Controller) GetBoxGeometrys(c *gin.Context) {

	// source slice
	var boxgeometryDBs []orm.BoxGeometryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBoxGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	_, err := db.Find(&boxgeometryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	boxgeometryAPIs := make([]orm.BoxGeometryAPI, 0)

	// for each boxgeometry, update fields from the database nullable fields
	for idx := range boxgeometryDBs {
		boxgeometryDB := &boxgeometryDBs[idx]
		_ = boxgeometryDB
		var boxgeometryAPI orm.BoxGeometryAPI

		// insertion point for updating fields
		boxgeometryAPI.ID = boxgeometryDB.ID
		boxgeometryDB.CopyBasicFieldsToBoxGeometry_WOP(&boxgeometryAPI.BoxGeometry_WOP)
		boxgeometryAPI.BoxGeometryPointersEncoding = boxgeometryDB.BoxGeometryPointersEncoding
		boxgeometryAPIs = append(boxgeometryAPIs, boxgeometryAPI)
	}

	c.JSON(http.StatusOK, boxgeometryAPIs)
}

// PostBoxGeometry
//
// swagger:route POST /boxgeometrys boxgeometrys postBoxGeometry
//
// Creates a boxgeometry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBoxGeometry(c *gin.Context) {

	mutexBoxGeometry.Lock()
	defer mutexBoxGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBoxGeometrys", "Name", stackPath)
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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	// Validate input
	var input orm.BoxGeometryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create boxgeometry
	boxgeometryDB := orm.BoxGeometryDB{}
	boxgeometryDB.BoxGeometryPointersEncoding = input.BoxGeometryPointersEncoding
	boxgeometryDB.CopyBasicFieldsFromBoxGeometry_WOP(&input.BoxGeometry_WOP)

	_, err = db.Create(&boxgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBoxGeometry.CheckoutPhaseOneInstance(&boxgeometryDB)
	boxgeometry := backRepo.BackRepoBoxGeometry.Map_BoxGeometryDBID_BoxGeometryPtr[boxgeometryDB.ID]

	if boxgeometry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), boxgeometry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, boxgeometryDB)
}

// GetBoxGeometry
//
// swagger:route GET /boxgeometrys/{ID} boxgeometrys getBoxGeometry
//
// Gets the details for a boxgeometry.
//
// Responses:
// default: genericError
//
//	200: boxgeometryDBResponse
func (controller *Controller) GetBoxGeometry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBoxGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	// Get boxgeometryDB in DB
	var boxgeometryDB orm.BoxGeometryDB
	if _, err := db.First(&boxgeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var boxgeometryAPI orm.BoxGeometryAPI
	boxgeometryAPI.ID = boxgeometryDB.ID
	boxgeometryAPI.BoxGeometryPointersEncoding = boxgeometryDB.BoxGeometryPointersEncoding
	boxgeometryDB.CopyBasicFieldsToBoxGeometry_WOP(&boxgeometryAPI.BoxGeometry_WOP)

	c.JSON(http.StatusOK, boxgeometryAPI)
}

// UpdateBoxGeometry
//
// swagger:route PATCH /boxgeometrys/{ID} boxgeometrys updateBoxGeometry
//
// # Update a boxgeometry
//
// Responses:
// default: genericError
//
//	200: boxgeometryDBResponse
func (controller *Controller) UpdateBoxGeometry(c *gin.Context) {

	mutexBoxGeometry.Lock()
	defer mutexBoxGeometry.Unlock()

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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	// Validate input
	var input orm.BoxGeometryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var boxgeometryDB orm.BoxGeometryDB

	// fetch the boxgeometry
	_, err := db.First(&boxgeometryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	boxgeometryDB.CopyBasicFieldsFromBoxGeometry_WOP(&input.BoxGeometry_WOP)
	boxgeometryDB.BoxGeometryPointersEncoding = input.BoxGeometryPointersEncoding

	db, _ = db.Model(&boxgeometryDB)
	_, err = db.Updates(&boxgeometryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	boxgeometryNew := new(models.BoxGeometry)
	boxgeometryDB.CopyBasicFieldsToBoxGeometry(boxgeometryNew)

	// redeem pointers
	boxgeometryDB.DecodePointers(backRepo, boxgeometryNew)

	// get stage instance from DB instance, and call callback function
	boxgeometryOld := backRepo.BackRepoBoxGeometry.Map_BoxGeometryDBID_BoxGeometryPtr[boxgeometryDB.ID]
	if boxgeometryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), boxgeometryOld, boxgeometryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the boxgeometryDB
	c.JSON(http.StatusOK, boxgeometryDB)
}

// DeleteBoxGeometry
//
// swagger:route DELETE /boxgeometrys/{ID} boxgeometrys deleteBoxGeometry
//
// # Delete a boxgeometry
//
// default: genericError
//
//	200: boxgeometryDBResponse
func (controller *Controller) DeleteBoxGeometry(c *gin.Context) {

	mutexBoxGeometry.Lock()
	defer mutexBoxGeometry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBoxGeometry", "Name", stackPath)
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
	db := backRepo.BackRepoBoxGeometry.GetDB()

	// Get model if exist
	var boxgeometryDB orm.BoxGeometryDB
	if _, err := db.First(&boxgeometryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&boxgeometryDB)

	// get an instance (not staged) from DB instance, and call callback function
	boxgeometryDeleted := new(models.BoxGeometry)
	boxgeometryDB.CopyBasicFieldsToBoxGeometry(boxgeometryDeleted)

	// get stage instance from DB instance, and call callback function
	boxgeometryStaged := backRepo.BackRepoBoxGeometry.Map_BoxGeometryDBID_BoxGeometryPtr[boxgeometryDB.ID]
	if boxgeometryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), boxgeometryStaged, boxgeometryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
