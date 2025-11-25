// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Object__dummysDeclaration__ models.Object
var __Object_time__dummyDeclaration time.Duration

var mutexObject sync.Mutex

// An ObjectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getObject updateObject deleteObject
type ObjectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ObjectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postObject updateObject
type ObjectInput struct {
	// The Object to submit or modify
	// in: body
	Object *orm.ObjectAPI
}

// GetObjects
//
// swagger:route GET /objects objects getObjects
//
// # Get all objects
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) GetObjects(c *gin.Context) {

	// source slice
	var objectDBs []orm.ObjectDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetObjects", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoObject.GetDB()

	_, err := db.Find(&objectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	objectAPIs := make([]orm.ObjectAPI, 0)

	// for each object, update fields from the database nullable fields
	for idx := range objectDBs {
		objectDB := &objectDBs[idx]
		_ = objectDB
		var objectAPI orm.ObjectAPI

		// insertion point for updating fields
		objectAPI.ID = objectDB.ID
		objectDB.CopyBasicFieldsToObject_WOP(&objectAPI.Object_WOP)
		objectAPI.ObjectPointersEncoding = objectDB.ObjectPointersEncoding
		objectAPIs = append(objectAPIs, objectAPI)
	}

	c.JSON(http.StatusOK, objectAPIs)
}

// PostObject
//
// swagger:route POST /objects objects postObject
//
// Creates a object
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostObject(c *gin.Context) {

	mutexObject.Lock()
	defer mutexObject.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostObjects", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoObject.GetDB()

	// Validate input
	var input orm.ObjectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create object
	objectDB := orm.ObjectDB{}
	objectDB.ObjectPointersEncoding = input.ObjectPointersEncoding
	objectDB.CopyBasicFieldsFromObject_WOP(&input.Object_WOP)

	_, err = db.Create(&objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoObject.CheckoutPhaseOneInstance(&objectDB)
	object := backRepo.BackRepoObject.Map_ObjectDBID_ObjectPtr[objectDB.ID]

	if object != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), object)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, objectDB)
}

// GetObject
//
// swagger:route GET /objects/{ID} objects getObject
//
// Gets the details for a object.
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) GetObject(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetObject", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoObject.GetDB()

	// Get objectDB in DB
	var objectDB orm.ObjectDB
	if _, err := db.First(&objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var objectAPI orm.ObjectAPI
	objectAPI.ID = objectDB.ID
	objectAPI.ObjectPointersEncoding = objectDB.ObjectPointersEncoding
	objectDB.CopyBasicFieldsToObject_WOP(&objectAPI.Object_WOP)

	c.JSON(http.StatusOK, objectAPI)
}

// UpdateObject
//
// swagger:route PATCH /objects/{ID} objects updateObject
//
// # Update a object
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) UpdateObject(c *gin.Context) {

	mutexObject.Lock()
	defer mutexObject.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoObject.GetDB()

	// Validate input
	var input orm.ObjectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var objectDB orm.ObjectDB

	// fetch the object
	_, err := db.First(&objectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	objectDB.CopyBasicFieldsFromObject_WOP(&input.Object_WOP)
	objectDB.ObjectPointersEncoding = input.ObjectPointersEncoding

	db, _ = db.Model(&objectDB)
	_, err = db.Updates(&objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	objectNew := new(models.Object)
	objectDB.CopyBasicFieldsToObject(objectNew)

	// redeem pointers
	objectDB.DecodePointers(backRepo, objectNew)

	// get stage instance from DB instance, and call callback function
	objectOld := backRepo.BackRepoObject.Map_ObjectDBID_ObjectPtr[objectDB.ID]
	if objectOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), objectOld, objectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the objectDB
	c.JSON(http.StatusOK, objectDB)
}

// DeleteObject
//
// swagger:route DELETE /objects/{ID} objects deleteObject
//
// # Delete a object
//
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) DeleteObject(c *gin.Context) {

	mutexObject.Lock()
	defer mutexObject.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteObject", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/test/statemachines/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoObject.GetDB()

	// Get model if exist
	var objectDB orm.ObjectDB
	if _, err := db.First(&objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&objectDB)

	// get an instance (not staged) from DB instance, and call callback function
	objectDeleted := new(models.Object)
	objectDB.CopyBasicFieldsToObject(objectDeleted)

	// get stage instance from DB instance, and call callback function
	objectStaged := backRepo.BackRepoObject.Map_ObjectDBID_ObjectPtr[objectDB.ID]
	if objectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), objectStaged, objectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
