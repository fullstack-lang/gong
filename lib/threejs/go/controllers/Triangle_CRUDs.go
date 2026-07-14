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
var __Triangle__dummysDeclaration__ models.Triangle
var _ = __Triangle__dummysDeclaration__
var __Triangle_time__dummyDeclaration time.Duration
var _ = __Triangle_time__dummyDeclaration

var mutexTriangle sync.Mutex

// An TriangleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTriangle updateTriangle deleteTriangle
type TriangleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TriangleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTriangle updateTriangle
type TriangleInput struct {
	// The Triangle to submit or modify
	// in: body
	Triangle *orm.TriangleAPI
}

// GetTriangles
//
// swagger:route GET /triangles triangles getTriangles
//
// # Get all triangles
//
// Responses:
// default: genericError
//
//	200: triangleDBResponse
func (controller *Controller) GetTriangles(c *gin.Context) {

	// source slice
	var triangleDBs []orm.TriangleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTriangles", "Name", stackPath)
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
	db := backRepo.BackRepoTriangle.GetDB()

	_, err := db.Find(&triangleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	triangleAPIs := make([]orm.TriangleAPI, 0)

	// for each triangle, update fields from the database nullable fields
	for idx := range triangleDBs {
		triangleDB := &triangleDBs[idx]
		_ = triangleDB
		var triangleAPI orm.TriangleAPI

		// insertion point for updating fields
		triangleAPI.ID = triangleDB.ID
		triangleDB.CopyBasicFieldsToTriangle_WOP(&triangleAPI.Triangle_WOP)
		triangleAPI.TrianglePointersEncoding = triangleDB.TrianglePointersEncoding
		triangleAPIs = append(triangleAPIs, triangleAPI)
	}

	c.JSON(http.StatusOK, triangleAPIs)
}

// PostTriangle
//
// swagger:route POST /triangles triangles postTriangle
//
// Creates a triangle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTriangle(c *gin.Context) {

	mutexTriangle.Lock()
	defer mutexTriangle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTriangles", "Name", stackPath)
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
	db := backRepo.BackRepoTriangle.GetDB()

	// Validate input
	var input orm.TriangleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create triangle
	triangleDB := orm.TriangleDB{}
	triangleDB.TrianglePointersEncoding = input.TrianglePointersEncoding
	triangleDB.CopyBasicFieldsFromTriangle_WOP(&input.Triangle_WOP)

	_, err = db.Create(&triangleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTriangle.CheckoutPhaseOneInstance(&triangleDB)
	triangle := backRepo.BackRepoTriangle.Map_TriangleDBID_TrianglePtr[triangleDB.ID]

	if triangle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), triangle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, triangleDB)
}

// GetTriangle
//
// swagger:route GET /triangles/{ID} triangles getTriangle
//
// Gets the details for a triangle.
//
// Responses:
// default: genericError
//
//	200: triangleDBResponse
func (controller *Controller) GetTriangle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTriangle", "Name", stackPath)
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
	db := backRepo.BackRepoTriangle.GetDB()

	// Get triangleDB in DB
	var triangleDB orm.TriangleDB
	if _, err := db.First(&triangleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var triangleAPI orm.TriangleAPI
	triangleAPI.ID = triangleDB.ID
	triangleAPI.TrianglePointersEncoding = triangleDB.TrianglePointersEncoding
	triangleDB.CopyBasicFieldsToTriangle_WOP(&triangleAPI.Triangle_WOP)

	c.JSON(http.StatusOK, triangleAPI)
}

// UpdateTriangle
//
// swagger:route PATCH /triangles/{ID} triangles updateTriangle
//
// # Update a triangle
//
// Responses:
// default: genericError
//
//	200: triangleDBResponse
func (controller *Controller) UpdateTriangle(c *gin.Context) {

	mutexTriangle.Lock()
	defer mutexTriangle.Unlock()

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
	db := backRepo.BackRepoTriangle.GetDB()

	// Validate input
	var input orm.TriangleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var triangleDB orm.TriangleDB

	// fetch the triangle
	_, err := db.First(&triangleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	triangleDB.CopyBasicFieldsFromTriangle_WOP(&input.Triangle_WOP)
	triangleDB.TrianglePointersEncoding = input.TrianglePointersEncoding

	db, _ = db.Model(&triangleDB)
	_, err = db.Updates(&triangleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	triangleNew := new(models.Triangle)
	triangleDB.CopyBasicFieldsToTriangle(triangleNew)

	// redeem pointers
	triangleDB.DecodePointers(backRepo, triangleNew)

	// get stage instance from DB instance, and call callback function
	triangleOld := backRepo.BackRepoTriangle.Map_TriangleDBID_TrianglePtr[triangleDB.ID]
	if triangleOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), triangleOld, triangleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the triangleDB
	c.JSON(http.StatusOK, triangleDB)
}

// DeleteTriangle
//
// swagger:route DELETE /triangles/{ID} triangles deleteTriangle
//
// # Delete a triangle
//
// default: genericError
//
//	200: triangleDBResponse
func (controller *Controller) DeleteTriangle(c *gin.Context) {

	mutexTriangle.Lock()
	defer mutexTriangle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTriangle", "Name", stackPath)
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
	db := backRepo.BackRepoTriangle.GetDB()

	// Get model if exist
	var triangleDB orm.TriangleDB
	if _, err := db.First(&triangleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&triangleDB)

	// get an instance (not staged) from DB instance, and call callback function
	triangleDeleted := new(models.Triangle)
	triangleDB.CopyBasicFieldsToTriangle(triangleDeleted)

	// get stage instance from DB instance, and call callback function
	triangleStaged := backRepo.BackRepoTriangle.Map_TriangleDBID_TrianglePtr[triangleDB.ID]
	if triangleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), triangleStaged, triangleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
