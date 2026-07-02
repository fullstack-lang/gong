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
var __Vector2__dummysDeclaration__ models.Vector2
var _ = __Vector2__dummysDeclaration__
var __Vector2_time__dummyDeclaration time.Duration
var _ = __Vector2_time__dummyDeclaration

var mutexVector2 sync.Mutex

// An Vector2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVector2 updateVector2 deleteVector2
type Vector2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Vector2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVector2 updateVector2
type Vector2Input struct {
	// The Vector2 to submit or modify
	// in: body
	Vector2 *orm.Vector2API
}

// GetVector2s
//
// swagger:route GET /vector2s vector2s getVector2s
//
// # Get all vector2s
//
// Responses:
// default: genericError
//
//	200: vector2DBResponse
func (controller *Controller) GetVector2s(c *gin.Context) {

	// source slice
	var vector2DBs []orm.Vector2DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVector2s", "Name", stackPath)
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
	db := backRepo.BackRepoVector2.GetDB()

	_, err := db.Find(&vector2DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	vector2APIs := make([]orm.Vector2API, 0)

	// for each vector2, update fields from the database nullable fields
	for idx := range vector2DBs {
		vector2DB := &vector2DBs[idx]
		_ = vector2DB
		var vector2API orm.Vector2API

		// insertion point for updating fields
		vector2API.ID = vector2DB.ID
		vector2DB.CopyBasicFieldsToVector2_WOP(&vector2API.Vector2_WOP)
		vector2API.Vector2PointersEncoding = vector2DB.Vector2PointersEncoding
		vector2APIs = append(vector2APIs, vector2API)
	}

	c.JSON(http.StatusOK, vector2APIs)
}

// PostVector2
//
// swagger:route POST /vector2s vector2s postVector2
//
// Creates a vector2
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVector2(c *gin.Context) {

	mutexVector2.Lock()
	defer mutexVector2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVector2s", "Name", stackPath)
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
	db := backRepo.BackRepoVector2.GetDB()

	// Validate input
	var input orm.Vector2API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vector2
	vector2DB := orm.Vector2DB{}
	vector2DB.Vector2PointersEncoding = input.Vector2PointersEncoding
	vector2DB.CopyBasicFieldsFromVector2_WOP(&input.Vector2_WOP)

	_, err = db.Create(&vector2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVector2.CheckoutPhaseOneInstance(&vector2DB)
	vector2 := backRepo.BackRepoVector2.Map_Vector2DBID_Vector2Ptr[vector2DB.ID]

	if vector2 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), vector2)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, vector2DB)
}

// GetVector2
//
// swagger:route GET /vector2s/{ID} vector2s getVector2
//
// Gets the details for a vector2.
//
// Responses:
// default: genericError
//
//	200: vector2DBResponse
func (controller *Controller) GetVector2(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVector2", "Name", stackPath)
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
	db := backRepo.BackRepoVector2.GetDB()

	// Get vector2DB in DB
	var vector2DB orm.Vector2DB
	if _, err := db.First(&vector2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var vector2API orm.Vector2API
	vector2API.ID = vector2DB.ID
	vector2API.Vector2PointersEncoding = vector2DB.Vector2PointersEncoding
	vector2DB.CopyBasicFieldsToVector2_WOP(&vector2API.Vector2_WOP)

	c.JSON(http.StatusOK, vector2API)
}

// UpdateVector2
//
// swagger:route PATCH /vector2s/{ID} vector2s updateVector2
//
// # Update a vector2
//
// Responses:
// default: genericError
//
//	200: vector2DBResponse
func (controller *Controller) UpdateVector2(c *gin.Context) {

	mutexVector2.Lock()
	defer mutexVector2.Unlock()

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
	db := backRepo.BackRepoVector2.GetDB()

	// Validate input
	var input orm.Vector2API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var vector2DB orm.Vector2DB

	// fetch the vector2
	_, err := db.First(&vector2DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	vector2DB.CopyBasicFieldsFromVector2_WOP(&input.Vector2_WOP)
	vector2DB.Vector2PointersEncoding = input.Vector2PointersEncoding

	db, _ = db.Model(&vector2DB)
	_, err = db.Updates(&vector2DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	vector2New := new(models.Vector2)
	vector2DB.CopyBasicFieldsToVector2(vector2New)

	// redeem pointers
	vector2DB.DecodePointers(backRepo, vector2New)

	// get stage instance from DB instance, and call callback function
	vector2Old := backRepo.BackRepoVector2.Map_Vector2DBID_Vector2Ptr[vector2DB.ID]
	if vector2Old != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), vector2Old, vector2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the vector2DB
	c.JSON(http.StatusOK, vector2DB)
}

// DeleteVector2
//
// swagger:route DELETE /vector2s/{ID} vector2s deleteVector2
//
// # Delete a vector2
//
// default: genericError
//
//	200: vector2DBResponse
func (controller *Controller) DeleteVector2(c *gin.Context) {

	mutexVector2.Lock()
	defer mutexVector2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVector2", "Name", stackPath)
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
	db := backRepo.BackRepoVector2.GetDB()

	// Get model if exist
	var vector2DB orm.Vector2DB
	if _, err := db.First(&vector2DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&vector2DB)

	// get an instance (not staged) from DB instance, and call callback function
	vector2Deleted := new(models.Vector2)
	vector2DB.CopyBasicFieldsToVector2(vector2Deleted)

	// get stage instance from DB instance, and call callback function
	vector2Staged := backRepo.BackRepoVector2.Map_Vector2DBID_Vector2Ptr[vector2DB.ID]
	if vector2Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), vector2Staged, vector2Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
