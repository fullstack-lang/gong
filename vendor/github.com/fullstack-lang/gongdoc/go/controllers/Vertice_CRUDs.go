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
var __Vertice__dummysDeclaration__ models.Vertice
var __Vertice_time__dummyDeclaration time.Duration

var mutexVertice sync.Mutex

// An VerticeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVertice updateVertice deleteVertice
type VerticeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VerticeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVertice updateVertice
type VerticeInput struct {
	// The Vertice to submit or modify
	// in: body
	Vertice *orm.VerticeAPI
}

// GetVertices
//
// swagger:route GET /vertices vertices getVertices
//
// # Get all vertices
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func (controller *Controller) GetVertices(c *gin.Context) {

	// source slice
	var verticeDBs []orm.VerticeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVertices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVertice.GetDB()

	_, err := db.Find(&verticeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	verticeAPIs := make([]orm.VerticeAPI, 0)

	// for each vertice, update fields from the database nullable fields
	for idx := range verticeDBs {
		verticeDB := &verticeDBs[idx]
		_ = verticeDB
		var verticeAPI orm.VerticeAPI

		// insertion point for updating fields
		verticeAPI.ID = verticeDB.ID
		verticeDB.CopyBasicFieldsToVertice_WOP(&verticeAPI.Vertice_WOP)
		verticeAPI.VerticePointersEncoding = verticeDB.VerticePointersEncoding
		verticeAPIs = append(verticeAPIs, verticeAPI)
	}

	c.JSON(http.StatusOK, verticeAPIs)
}

// PostVertice
//
// swagger:route POST /vertices vertices postVertice
//
// Creates a vertice
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVertice(c *gin.Context) {

	mutexVertice.Lock()
	defer mutexVertice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVertices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVertice.GetDB()

	// Validate input
	var input orm.VerticeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vertice
	verticeDB := orm.VerticeDB{}
	verticeDB.VerticePointersEncoding = input.VerticePointersEncoding
	verticeDB.CopyBasicFieldsFromVertice_WOP(&input.Vertice_WOP)

	_, err = db.Create(&verticeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVertice.CheckoutPhaseOneInstance(&verticeDB)
	vertice := backRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr[verticeDB.ID]

	if vertice != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), vertice)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, verticeDB)
}

// GetVertice
//
// swagger:route GET /vertices/{ID} vertices getVertice
//
// Gets the details for a vertice.
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func (controller *Controller) GetVertice(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVertice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVertice.GetDB()

	// Get verticeDB in DB
	var verticeDB orm.VerticeDB
	if _, err := db.First(&verticeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var verticeAPI orm.VerticeAPI
	verticeAPI.ID = verticeDB.ID
	verticeAPI.VerticePointersEncoding = verticeDB.VerticePointersEncoding
	verticeDB.CopyBasicFieldsToVertice_WOP(&verticeAPI.Vertice_WOP)

	c.JSON(http.StatusOK, verticeAPI)
}

// UpdateVertice
//
// swagger:route PATCH /vertices/{ID} vertices updateVertice
//
// # Update a vertice
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func (controller *Controller) UpdateVertice(c *gin.Context) {

	mutexVertice.Lock()
	defer mutexVertice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVertice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVertice.GetDB()

	// Validate input
	var input orm.VerticeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var verticeDB orm.VerticeDB

	// fetch the vertice
	_, err := db.First(&verticeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	verticeDB.CopyBasicFieldsFromVertice_WOP(&input.Vertice_WOP)
	verticeDB.VerticePointersEncoding = input.VerticePointersEncoding

	db, _ = db.Model(&verticeDB)
	_, err = db.Updates(&verticeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	verticeNew := new(models.Vertice)
	verticeDB.CopyBasicFieldsToVertice(verticeNew)

	// redeem pointers
	verticeDB.DecodePointers(backRepo, verticeNew)

	// get stage instance from DB instance, and call callback function
	verticeOld := backRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr[verticeDB.ID]
	if verticeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), verticeOld, verticeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the verticeDB
	c.JSON(http.StatusOK, verticeDB)
}

// DeleteVertice
//
// swagger:route DELETE /vertices/{ID} vertices deleteVertice
//
// # Delete a vertice
//
// default: genericError
//
//	200: verticeDBResponse
func (controller *Controller) DeleteVertice(c *gin.Context) {

	mutexVertice.Lock()
	defer mutexVertice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVertice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVertice.GetDB()

	// Get model if exist
	var verticeDB orm.VerticeDB
	if _, err := db.First(&verticeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&verticeDB)

	// get an instance (not staged) from DB instance, and call callback function
	verticeDeleted := new(models.Vertice)
	verticeDB.CopyBasicFieldsToVertice(verticeDeleted)

	// get stage instance from DB instance, and call callback function
	verticeStaged := backRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr[verticeDB.ID]
	if verticeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), verticeStaged, verticeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
