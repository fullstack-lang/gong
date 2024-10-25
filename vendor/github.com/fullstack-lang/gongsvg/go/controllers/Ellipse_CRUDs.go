// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Ellipse__dummysDeclaration__ models.Ellipse
var __Ellipse_time__dummyDeclaration time.Duration

var mutexEllipse sync.Mutex

// An EllipseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEllipse updateEllipse deleteEllipse
type EllipseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EllipseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEllipse updateEllipse
type EllipseInput struct {
	// The Ellipse to submit or modify
	// in: body
	Ellipse *orm.EllipseAPI
}

// GetEllipses
//
// swagger:route GET /ellipses ellipses getEllipses
//
// # Get all ellipses
//
// Responses:
// default: genericError
//
//	200: ellipseDBResponse
func (controller *Controller) GetEllipses(c *gin.Context) {

	// source slice
	var ellipseDBs []orm.EllipseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEllipses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEllipse.GetDB()

	_, err := db.Find(&ellipseDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ellipseAPIs := make([]orm.EllipseAPI, 0)

	// for each ellipse, update fields from the database nullable fields
	for idx := range ellipseDBs {
		ellipseDB := &ellipseDBs[idx]
		_ = ellipseDB
		var ellipseAPI orm.EllipseAPI

		// insertion point for updating fields
		ellipseAPI.ID = ellipseDB.ID
		ellipseDB.CopyBasicFieldsToEllipse_WOP(&ellipseAPI.Ellipse_WOP)
		ellipseAPI.EllipsePointersEncoding = ellipseDB.EllipsePointersEncoding
		ellipseAPIs = append(ellipseAPIs, ellipseAPI)
	}

	c.JSON(http.StatusOK, ellipseAPIs)
}

// PostEllipse
//
// swagger:route POST /ellipses ellipses postEllipse
//
// Creates a ellipse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEllipse(c *gin.Context) {

	mutexEllipse.Lock()
	defer mutexEllipse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEllipses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEllipse.GetDB()

	// Validate input
	var input orm.EllipseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ellipse
	ellipseDB := orm.EllipseDB{}
	ellipseDB.EllipsePointersEncoding = input.EllipsePointersEncoding
	ellipseDB.CopyBasicFieldsFromEllipse_WOP(&input.Ellipse_WOP)

	_, err = db.Create(&ellipseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEllipse.CheckoutPhaseOneInstance(&ellipseDB)
	ellipse := backRepo.BackRepoEllipse.Map_EllipseDBID_EllipsePtr[ellipseDB.ID]

	if ellipse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), ellipse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, ellipseDB)
}

// GetEllipse
//
// swagger:route GET /ellipses/{ID} ellipses getEllipse
//
// Gets the details for a ellipse.
//
// Responses:
// default: genericError
//
//	200: ellipseDBResponse
func (controller *Controller) GetEllipse(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEllipse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEllipse.GetDB()

	// Get ellipseDB in DB
	var ellipseDB orm.EllipseDB
	if _, err := db.First(&ellipseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var ellipseAPI orm.EllipseAPI
	ellipseAPI.ID = ellipseDB.ID
	ellipseAPI.EllipsePointersEncoding = ellipseDB.EllipsePointersEncoding
	ellipseDB.CopyBasicFieldsToEllipse_WOP(&ellipseAPI.Ellipse_WOP)

	c.JSON(http.StatusOK, ellipseAPI)
}

// UpdateEllipse
//
// swagger:route PATCH /ellipses/{ID} ellipses updateEllipse
//
// # Update a ellipse
//
// Responses:
// default: genericError
//
//	200: ellipseDBResponse
func (controller *Controller) UpdateEllipse(c *gin.Context) {

	mutexEllipse.Lock()
	defer mutexEllipse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEllipse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEllipse.GetDB()

	// Validate input
	var input orm.EllipseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ellipseDB orm.EllipseDB

	// fetch the ellipse
	_, err := db.First(&ellipseDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ellipseDB.CopyBasicFieldsFromEllipse_WOP(&input.Ellipse_WOP)
	ellipseDB.EllipsePointersEncoding = input.EllipsePointersEncoding

	db, _ = db.Model(&ellipseDB)
	_, err = db.Updates(&ellipseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ellipseNew := new(models.Ellipse)
	ellipseDB.CopyBasicFieldsToEllipse(ellipseNew)

	// redeem pointers
	ellipseDB.DecodePointers(backRepo, ellipseNew)

	// get stage instance from DB instance, and call callback function
	ellipseOld := backRepo.BackRepoEllipse.Map_EllipseDBID_EllipsePtr[ellipseDB.ID]
	if ellipseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), ellipseOld, ellipseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ellipseDB
	c.JSON(http.StatusOK, ellipseDB)
}

// DeleteEllipse
//
// swagger:route DELETE /ellipses/{ID} ellipses deleteEllipse
//
// # Delete a ellipse
//
// default: genericError
//
//	200: ellipseDBResponse
func (controller *Controller) DeleteEllipse(c *gin.Context) {

	mutexEllipse.Lock()
	defer mutexEllipse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEllipse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEllipse.GetDB()

	// Get model if exist
	var ellipseDB orm.EllipseDB
	if _, err := db.First(&ellipseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&ellipseDB)

	// get an instance (not staged) from DB instance, and call callback function
	ellipseDeleted := new(models.Ellipse)
	ellipseDB.CopyBasicFieldsToEllipse(ellipseDeleted)

	// get stage instance from DB instance, and call callback function
	ellipseStaged := backRepo.BackRepoEllipse.Map_EllipseDBID_EllipsePtr[ellipseDB.ID]
	if ellipseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), ellipseStaged, ellipseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
