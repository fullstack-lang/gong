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
var __Vector3__dummysDeclaration__ models.Vector3
var _ = __Vector3__dummysDeclaration__
var __Vector3_time__dummyDeclaration time.Duration
var _ = __Vector3_time__dummyDeclaration

var mutexVector3 sync.Mutex

// An Vector3ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVector3 updateVector3 deleteVector3
type Vector3ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Vector3Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVector3 updateVector3
type Vector3Input struct {
	// The Vector3 to submit or modify
	// in: body
	Vector3 *orm.Vector3API
}

// GetVector3s
//
// swagger:route GET /vector3s vector3s getVector3s
//
// # Get all vector3s
//
// Responses:
// default: genericError
//
//	200: vector3DBResponse
func (controller *Controller) GetVector3s(c *gin.Context) {

	// source slice
	var vector3DBs []orm.Vector3DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVector3s", "Name", stackPath)
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
	db := backRepo.BackRepoVector3.GetDB()

	_, err := db.Find(&vector3DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	vector3APIs := make([]orm.Vector3API, 0)

	// for each vector3, update fields from the database nullable fields
	for idx := range vector3DBs {
		vector3DB := &vector3DBs[idx]
		_ = vector3DB
		var vector3API orm.Vector3API

		// insertion point for updating fields
		vector3API.ID = vector3DB.ID
		vector3DB.CopyBasicFieldsToVector3_WOP(&vector3API.Vector3_WOP)
		vector3API.Vector3PointersEncoding = vector3DB.Vector3PointersEncoding
		vector3APIs = append(vector3APIs, vector3API)
	}

	c.JSON(http.StatusOK, vector3APIs)
}

// PostVector3
//
// swagger:route POST /vector3s vector3s postVector3
//
// Creates a vector3
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVector3(c *gin.Context) {

	mutexVector3.Lock()
	defer mutexVector3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVector3s", "Name", stackPath)
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
	db := backRepo.BackRepoVector3.GetDB()

	// Validate input
	var input orm.Vector3API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vector3
	vector3DB := orm.Vector3DB{}
	vector3DB.Vector3PointersEncoding = input.Vector3PointersEncoding
	vector3DB.CopyBasicFieldsFromVector3_WOP(&input.Vector3_WOP)

	_, err = db.Create(&vector3DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVector3.CheckoutPhaseOneInstance(&vector3DB)
	vector3 := backRepo.BackRepoVector3.Map_Vector3DBID_Vector3Ptr[vector3DB.ID]

	if vector3 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), vector3)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, vector3DB)
}

// GetVector3
//
// swagger:route GET /vector3s/{ID} vector3s getVector3
//
// Gets the details for a vector3.
//
// Responses:
// default: genericError
//
//	200: vector3DBResponse
func (controller *Controller) GetVector3(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVector3", "Name", stackPath)
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
	db := backRepo.BackRepoVector3.GetDB()

	// Get vector3DB in DB
	var vector3DB orm.Vector3DB
	if _, err := db.First(&vector3DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var vector3API orm.Vector3API
	vector3API.ID = vector3DB.ID
	vector3API.Vector3PointersEncoding = vector3DB.Vector3PointersEncoding
	vector3DB.CopyBasicFieldsToVector3_WOP(&vector3API.Vector3_WOP)

	c.JSON(http.StatusOK, vector3API)
}

// UpdateVector3
//
// swagger:route PATCH /vector3s/{ID} vector3s updateVector3
//
// # Update a vector3
//
// Responses:
// default: genericError
//
//	200: vector3DBResponse
func (controller *Controller) UpdateVector3(c *gin.Context) {

	mutexVector3.Lock()
	defer mutexVector3.Unlock()

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
	db := backRepo.BackRepoVector3.GetDB()

	// Validate input
	var input orm.Vector3API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var vector3DB orm.Vector3DB

	// fetch the vector3
	_, err := db.First(&vector3DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	vector3DB.CopyBasicFieldsFromVector3_WOP(&input.Vector3_WOP)
	vector3DB.Vector3PointersEncoding = input.Vector3PointersEncoding

	db, _ = db.Model(&vector3DB)
	_, err = db.Updates(&vector3DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	vector3New := new(models.Vector3)
	vector3DB.CopyBasicFieldsToVector3(vector3New)

	// redeem pointers
	vector3DB.DecodePointers(backRepo, vector3New)

	// get stage instance from DB instance, and call callback function
	vector3Old := backRepo.BackRepoVector3.Map_Vector3DBID_Vector3Ptr[vector3DB.ID]
	if vector3Old != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), vector3Old, vector3New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the vector3DB
	c.JSON(http.StatusOK, vector3DB)
}

// DeleteVector3
//
// swagger:route DELETE /vector3s/{ID} vector3s deleteVector3
//
// # Delete a vector3
//
// default: genericError
//
//	200: vector3DBResponse
func (controller *Controller) DeleteVector3(c *gin.Context) {

	mutexVector3.Lock()
	defer mutexVector3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVector3", "Name", stackPath)
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
	db := backRepo.BackRepoVector3.GetDB()

	// Get model if exist
	var vector3DB orm.Vector3DB
	if _, err := db.First(&vector3DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&vector3DB)

	// get an instance (not staged) from DB instance, and call callback function
	vector3Deleted := new(models.Vector3)
	vector3DB.CopyBasicFieldsToVector3(vector3Deleted)

	// get stage instance from DB instance, and call callback function
	vector3Staged := backRepo.BackRepoVector3.Map_Vector3DBID_Vector3Ptr[vector3DB.ID]
	if vector3Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), vector3Staged, vector3Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
