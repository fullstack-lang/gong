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
var __Path__dummysDeclaration__ models.Path
var __Path_time__dummyDeclaration time.Duration

var mutexPath sync.Mutex

// An PathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPath updatePath deletePath
type PathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPath updatePath
type PathInput struct {
	// The Path to submit or modify
	// in: body
	Path *orm.PathAPI
}

// GetPaths
//
// swagger:route GET /paths paths getPaths
//
// # Get all paths
//
// Responses:
// default: genericError
//
//	200: pathDBResponse
func (controller *Controller) GetPaths(c *gin.Context) {

	// source slice
	var pathDBs []orm.PathDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPaths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPath.GetDB()

	_, err := db.Find(&pathDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pathAPIs := make([]orm.PathAPI, 0)

	// for each path, update fields from the database nullable fields
	for idx := range pathDBs {
		pathDB := &pathDBs[idx]
		_ = pathDB
		var pathAPI orm.PathAPI

		// insertion point for updating fields
		pathAPI.ID = pathDB.ID
		pathDB.CopyBasicFieldsToPath_WOP(&pathAPI.Path_WOP)
		pathAPI.PathPointersEncoding = pathDB.PathPointersEncoding
		pathAPIs = append(pathAPIs, pathAPI)
	}

	c.JSON(http.StatusOK, pathAPIs)
}

// PostPath
//
// swagger:route POST /paths paths postPath
//
// Creates a path
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPath(c *gin.Context) {

	mutexPath.Lock()
	defer mutexPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPaths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPath.GetDB()

	// Validate input
	var input orm.PathAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create path
	pathDB := orm.PathDB{}
	pathDB.PathPointersEncoding = input.PathPointersEncoding
	pathDB.CopyBasicFieldsFromPath_WOP(&input.Path_WOP)

	_, err = db.Create(&pathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPath.CheckoutPhaseOneInstance(&pathDB)
	path := backRepo.BackRepoPath.Map_PathDBID_PathPtr[pathDB.ID]

	if path != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), path)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pathDB)
}

// GetPath
//
// swagger:route GET /paths/{ID} paths getPath
//
// Gets the details for a path.
//
// Responses:
// default: genericError
//
//	200: pathDBResponse
func (controller *Controller) GetPath(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPath.GetDB()

	// Get pathDB in DB
	var pathDB orm.PathDB
	if _, err := db.First(&pathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pathAPI orm.PathAPI
	pathAPI.ID = pathDB.ID
	pathAPI.PathPointersEncoding = pathDB.PathPointersEncoding
	pathDB.CopyBasicFieldsToPath_WOP(&pathAPI.Path_WOP)

	c.JSON(http.StatusOK, pathAPI)
}

// UpdatePath
//
// swagger:route PATCH /paths/{ID} paths updatePath
//
// # Update a path
//
// Responses:
// default: genericError
//
//	200: pathDBResponse
func (controller *Controller) UpdatePath(c *gin.Context) {

	mutexPath.Lock()
	defer mutexPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPath.GetDB()

	// Validate input
	var input orm.PathAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pathDB orm.PathDB

	// fetch the path
	_, err := db.First(&pathDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pathDB.CopyBasicFieldsFromPath_WOP(&input.Path_WOP)
	pathDB.PathPointersEncoding = input.PathPointersEncoding

	db, _ = db.Model(&pathDB)
	_, err = db.Updates(&pathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pathNew := new(models.Path)
	pathDB.CopyBasicFieldsToPath(pathNew)

	// redeem pointers
	pathDB.DecodePointers(backRepo, pathNew)

	// get stage instance from DB instance, and call callback function
	pathOld := backRepo.BackRepoPath.Map_PathDBID_PathPtr[pathDB.ID]
	if pathOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pathOld, pathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pathDB
	c.JSON(http.StatusOK, pathDB)
}

// DeletePath
//
// swagger:route DELETE /paths/{ID} paths deletePath
//
// # Delete a path
//
// default: genericError
//
//	200: pathDBResponse
func (controller *Controller) DeletePath(c *gin.Context) {

	mutexPath.Lock()
	defer mutexPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePath", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPath.GetDB()

	// Get model if exist
	var pathDB orm.PathDB
	if _, err := db.First(&pathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&pathDB)

	// get an instance (not staged) from DB instance, and call callback function
	pathDeleted := new(models.Path)
	pathDB.CopyBasicFieldsToPath(pathDeleted)

	// get stage instance from DB instance, and call callback function
	pathStaged := backRepo.BackRepoPath.Map_PathDBID_PathPtr[pathDB.ID]
	if pathStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pathStaged, pathDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
