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
var __Curve__dummysDeclaration__ models.Curve
var _ = __Curve__dummysDeclaration__
var __Curve_time__dummyDeclaration time.Duration
var _ = __Curve_time__dummyDeclaration

var mutexCurve sync.Mutex

// An CurveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCurve updateCurve deleteCurve
type CurveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CurveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCurve updateCurve
type CurveInput struct {
	// The Curve to submit or modify
	// in: body
	Curve *orm.CurveAPI
}

// GetCurves
//
// swagger:route GET /curves curves getCurves
//
// # Get all curves
//
// Responses:
// default: genericError
//
//	200: curveDBResponse
func (controller *Controller) GetCurves(c *gin.Context) {

	// source slice
	var curveDBs []orm.CurveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCurves", "Name", stackPath)
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
	db := backRepo.BackRepoCurve.GetDB()

	_, err := db.Find(&curveDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	curveAPIs := make([]orm.CurveAPI, 0)

	// for each curve, update fields from the database nullable fields
	for idx := range curveDBs {
		curveDB := &curveDBs[idx]
		_ = curveDB
		var curveAPI orm.CurveAPI

		// insertion point for updating fields
		curveAPI.ID = curveDB.ID
		curveDB.CopyBasicFieldsToCurve_WOP(&curveAPI.Curve_WOP)
		curveAPI.CurvePointersEncoding = curveDB.CurvePointersEncoding
		curveAPIs = append(curveAPIs, curveAPI)
	}

	c.JSON(http.StatusOK, curveAPIs)
}

// PostCurve
//
// swagger:route POST /curves curves postCurve
//
// Creates a curve
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCurve(c *gin.Context) {

	mutexCurve.Lock()
	defer mutexCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCurves", "Name", stackPath)
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
	db := backRepo.BackRepoCurve.GetDB()

	// Validate input
	var input orm.CurveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create curve
	curveDB := orm.CurveDB{}
	curveDB.CurvePointersEncoding = input.CurvePointersEncoding
	curveDB.CopyBasicFieldsFromCurve_WOP(&input.Curve_WOP)

	_, err = db.Create(&curveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCurve.CheckoutPhaseOneInstance(&curveDB)
	curve := backRepo.BackRepoCurve.Map_CurveDBID_CurvePtr[curveDB.ID]

	if curve != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), curve)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, curveDB)
}

// GetCurve
//
// swagger:route GET /curves/{ID} curves getCurve
//
// Gets the details for a curve.
//
// Responses:
// default: genericError
//
//	200: curveDBResponse
func (controller *Controller) GetCurve(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCurve", "Name", stackPath)
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
	db := backRepo.BackRepoCurve.GetDB()

	// Get curveDB in DB
	var curveDB orm.CurveDB
	if _, err := db.First(&curveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var curveAPI orm.CurveAPI
	curveAPI.ID = curveDB.ID
	curveAPI.CurvePointersEncoding = curveDB.CurvePointersEncoding
	curveDB.CopyBasicFieldsToCurve_WOP(&curveAPI.Curve_WOP)

	c.JSON(http.StatusOK, curveAPI)
}

// UpdateCurve
//
// swagger:route PATCH /curves/{ID} curves updateCurve
//
// # Update a curve
//
// Responses:
// default: genericError
//
//	200: curveDBResponse
func (controller *Controller) UpdateCurve(c *gin.Context) {

	mutexCurve.Lock()
	defer mutexCurve.Unlock()

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
	db := backRepo.BackRepoCurve.GetDB()

	// Validate input
	var input orm.CurveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var curveDB orm.CurveDB

	// fetch the curve
	_, err := db.First(&curveDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	curveDB.CopyBasicFieldsFromCurve_WOP(&input.Curve_WOP)
	curveDB.CurvePointersEncoding = input.CurvePointersEncoding

	db, _ = db.Model(&curveDB)
	_, err = db.Updates(&curveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	curveNew := new(models.Curve)
	curveDB.CopyBasicFieldsToCurve(curveNew)

	// redeem pointers
	curveDB.DecodePointers(backRepo, curveNew)

	// get stage instance from DB instance, and call callback function
	curveOld := backRepo.BackRepoCurve.Map_CurveDBID_CurvePtr[curveDB.ID]
	if curveOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), curveOld, curveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the curveDB
	c.JSON(http.StatusOK, curveDB)
}

// DeleteCurve
//
// swagger:route DELETE /curves/{ID} curves deleteCurve
//
// # Delete a curve
//
// default: genericError
//
//	200: curveDBResponse
func (controller *Controller) DeleteCurve(c *gin.Context) {

	mutexCurve.Lock()
	defer mutexCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCurve", "Name", stackPath)
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
	db := backRepo.BackRepoCurve.GetDB()

	// Get model if exist
	var curveDB orm.CurveDB
	if _, err := db.First(&curveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&curveDB)

	// get an instance (not staged) from DB instance, and call callback function
	curveDeleted := new(models.Curve)
	curveDB.CopyBasicFieldsToCurve(curveDeleted)

	// get stage instance from DB instance, and call callback function
	curveStaged := backRepo.BackRepoCurve.Map_CurveDBID_CurvePtr[curveDB.ID]
	if curveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), curveStaged, curveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
