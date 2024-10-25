// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __DiagramPackage__dummysDeclaration__ models.DiagramPackage
var __DiagramPackage_time__dummyDeclaration time.Duration

var mutexDiagramPackage sync.Mutex

// An DiagramPackageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDiagramPackage updateDiagramPackage deleteDiagramPackage
type DiagramPackageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DiagramPackageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDiagramPackage updateDiagramPackage
type DiagramPackageInput struct {
	// The DiagramPackage to submit or modify
	// in: body
	DiagramPackage *orm.DiagramPackageAPI
}

// GetDiagramPackages
//
// swagger:route GET /diagrampackages diagrampackages getDiagramPackages
//
// # Get all diagrampackages
//
// Responses:
// default: genericError
//
//	200: diagrampackageDBResponse
func (controller *Controller) GetDiagramPackages(c *gin.Context) {

	// source slice
	var diagrampackageDBs []orm.DiagramPackageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDiagramPackages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDiagramPackage.GetDB()

	_, err := db.Find(&diagrampackageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	diagrampackageAPIs := make([]orm.DiagramPackageAPI, 0)

	// for each diagrampackage, update fields from the database nullable fields
	for idx := range diagrampackageDBs {
		diagrampackageDB := &diagrampackageDBs[idx]
		_ = diagrampackageDB
		var diagrampackageAPI orm.DiagramPackageAPI

		// insertion point for updating fields
		diagrampackageAPI.ID = diagrampackageDB.ID
		diagrampackageDB.CopyBasicFieldsToDiagramPackage_WOP(&diagrampackageAPI.DiagramPackage_WOP)
		diagrampackageAPI.DiagramPackagePointersEncoding = diagrampackageDB.DiagramPackagePointersEncoding
		diagrampackageAPIs = append(diagrampackageAPIs, diagrampackageAPI)
	}

	c.JSON(http.StatusOK, diagrampackageAPIs)
}

// PostDiagramPackage
//
// swagger:route POST /diagrampackages diagrampackages postDiagramPackage
//
// Creates a diagrampackage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDiagramPackage(c *gin.Context) {

	mutexDiagramPackage.Lock()
	defer mutexDiagramPackage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDiagramPackages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDiagramPackage.GetDB()

	// Validate input
	var input orm.DiagramPackageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create diagrampackage
	diagrampackageDB := orm.DiagramPackageDB{}
	diagrampackageDB.DiagramPackagePointersEncoding = input.DiagramPackagePointersEncoding
	diagrampackageDB.CopyBasicFieldsFromDiagramPackage_WOP(&input.DiagramPackage_WOP)

	_, err = db.Create(&diagrampackageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDiagramPackage.CheckoutPhaseOneInstance(&diagrampackageDB)
	diagrampackage := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[diagrampackageDB.ID]

	if diagrampackage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), diagrampackage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, diagrampackageDB)
}

// GetDiagramPackage
//
// swagger:route GET /diagrampackages/{ID} diagrampackages getDiagramPackage
//
// Gets the details for a diagrampackage.
//
// Responses:
// default: genericError
//
//	200: diagrampackageDBResponse
func (controller *Controller) GetDiagramPackage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDiagramPackage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDiagramPackage.GetDB()

	// Get diagrampackageDB in DB
	var diagrampackageDB orm.DiagramPackageDB
	if _, err := db.First(&diagrampackageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var diagrampackageAPI orm.DiagramPackageAPI
	diagrampackageAPI.ID = diagrampackageDB.ID
	diagrampackageAPI.DiagramPackagePointersEncoding = diagrampackageDB.DiagramPackagePointersEncoding
	diagrampackageDB.CopyBasicFieldsToDiagramPackage_WOP(&diagrampackageAPI.DiagramPackage_WOP)

	c.JSON(http.StatusOK, diagrampackageAPI)
}

// UpdateDiagramPackage
//
// swagger:route PATCH /diagrampackages/{ID} diagrampackages updateDiagramPackage
//
// # Update a diagrampackage
//
// Responses:
// default: genericError
//
//	200: diagrampackageDBResponse
func (controller *Controller) UpdateDiagramPackage(c *gin.Context) {

	mutexDiagramPackage.Lock()
	defer mutexDiagramPackage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDiagramPackage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDiagramPackage.GetDB()

	// Validate input
	var input orm.DiagramPackageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var diagrampackageDB orm.DiagramPackageDB

	// fetch the diagrampackage
	_, err := db.First(&diagrampackageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	diagrampackageDB.CopyBasicFieldsFromDiagramPackage_WOP(&input.DiagramPackage_WOP)
	diagrampackageDB.DiagramPackagePointersEncoding = input.DiagramPackagePointersEncoding

	db, _ = db.Model(&diagrampackageDB)
	_, err = db.Updates(&diagrampackageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	diagrampackageNew := new(models.DiagramPackage)
	diagrampackageDB.CopyBasicFieldsToDiagramPackage(diagrampackageNew)

	// redeem pointers
	diagrampackageDB.DecodePointers(backRepo, diagrampackageNew)

	// get stage instance from DB instance, and call callback function
	diagrampackageOld := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[diagrampackageDB.ID]
	if diagrampackageOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), diagrampackageOld, diagrampackageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the diagrampackageDB
	c.JSON(http.StatusOK, diagrampackageDB)
}

// DeleteDiagramPackage
//
// swagger:route DELETE /diagrampackages/{ID} diagrampackages deleteDiagramPackage
//
// # Delete a diagrampackage
//
// default: genericError
//
//	200: diagrampackageDBResponse
func (controller *Controller) DeleteDiagramPackage(c *gin.Context) {

	mutexDiagramPackage.Lock()
	defer mutexDiagramPackage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDiagramPackage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDiagramPackage.GetDB()

	// Get model if exist
	var diagrampackageDB orm.DiagramPackageDB
	if _, err := db.First(&diagrampackageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&diagrampackageDB)

	// get an instance (not staged) from DB instance, and call callback function
	diagrampackageDeleted := new(models.DiagramPackage)
	diagrampackageDB.CopyBasicFieldsToDiagramPackage(diagrampackageDeleted)

	// get stage instance from DB instance, and call callback function
	diagrampackageStaged := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[diagrampackageDB.ID]
	if diagrampackageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), diagrampackageStaged, diagrampackageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
