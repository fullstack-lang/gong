// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	"github.com/fullstack-lang/gong/lib/doc2/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Classdiagram__dummysDeclaration__ models.Classdiagram
var __Classdiagram_time__dummyDeclaration time.Duration

var mutexClassdiagram sync.Mutex

// An ClassdiagramID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getClassdiagram updateClassdiagram deleteClassdiagram
type ClassdiagramID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClassdiagramInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postClassdiagram updateClassdiagram
type ClassdiagramInput struct {
	// The Classdiagram to submit or modify
	// in: body
	Classdiagram *orm.ClassdiagramAPI
}

// GetClassdiagrams
//
// swagger:route GET /classdiagrams classdiagrams getClassdiagrams
//
// # Get all classdiagrams
//
// Responses:
// default: genericError
//
//	200: classdiagramDBResponse
func (controller *Controller) GetClassdiagrams(c *gin.Context) {

	// source slice
	var classdiagramDBs []orm.ClassdiagramDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetClassdiagrams", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoClassdiagram.GetDB()

	_, err := db.Find(&classdiagramDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	classdiagramAPIs := make([]orm.ClassdiagramAPI, 0)

	// for each classdiagram, update fields from the database nullable fields
	for idx := range classdiagramDBs {
		classdiagramDB := &classdiagramDBs[idx]
		_ = classdiagramDB
		var classdiagramAPI orm.ClassdiagramAPI

		// insertion point for updating fields
		classdiagramAPI.ID = classdiagramDB.ID
		classdiagramDB.CopyBasicFieldsToClassdiagram_WOP(&classdiagramAPI.Classdiagram_WOP)
		classdiagramAPI.ClassdiagramPointersEncoding = classdiagramDB.ClassdiagramPointersEncoding
		classdiagramAPIs = append(classdiagramAPIs, classdiagramAPI)
	}

	c.JSON(http.StatusOK, classdiagramAPIs)
}

// PostClassdiagram
//
// swagger:route POST /classdiagrams classdiagrams postClassdiagram
//
// Creates a classdiagram
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostClassdiagram(c *gin.Context) {

	mutexClassdiagram.Lock()
	defer mutexClassdiagram.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostClassdiagrams", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoClassdiagram.GetDB()

	// Validate input
	var input orm.ClassdiagramAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create classdiagram
	classdiagramDB := orm.ClassdiagramDB{}
	classdiagramDB.ClassdiagramPointersEncoding = input.ClassdiagramPointersEncoding
	classdiagramDB.CopyBasicFieldsFromClassdiagram_WOP(&input.Classdiagram_WOP)

	_, err = db.Create(&classdiagramDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoClassdiagram.CheckoutPhaseOneInstance(&classdiagramDB)
	classdiagram := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]

	if classdiagram != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), classdiagram)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, classdiagramDB)
}

// GetClassdiagram
//
// swagger:route GET /classdiagrams/{ID} classdiagrams getClassdiagram
//
// Gets the details for a classdiagram.
//
// Responses:
// default: genericError
//
//	200: classdiagramDBResponse
func (controller *Controller) GetClassdiagram(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetClassdiagram", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoClassdiagram.GetDB()

	// Get classdiagramDB in DB
	var classdiagramDB orm.ClassdiagramDB
	if _, err := db.First(&classdiagramDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var classdiagramAPI orm.ClassdiagramAPI
	classdiagramAPI.ID = classdiagramDB.ID
	classdiagramAPI.ClassdiagramPointersEncoding = classdiagramDB.ClassdiagramPointersEncoding
	classdiagramDB.CopyBasicFieldsToClassdiagram_WOP(&classdiagramAPI.Classdiagram_WOP)

	c.JSON(http.StatusOK, classdiagramAPI)
}

// UpdateClassdiagram
//
// swagger:route PATCH /classdiagrams/{ID} classdiagrams updateClassdiagram
//
// # Update a classdiagram
//
// Responses:
// default: genericError
//
//	200: classdiagramDBResponse
func (controller *Controller) UpdateClassdiagram(c *gin.Context) {

	mutexClassdiagram.Lock()
	defer mutexClassdiagram.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateClassdiagram", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoClassdiagram.GetDB()

	// Validate input
	var input orm.ClassdiagramAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var classdiagramDB orm.ClassdiagramDB

	// fetch the classdiagram
	_, err := db.First(&classdiagramDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	classdiagramDB.CopyBasicFieldsFromClassdiagram_WOP(&input.Classdiagram_WOP)
	classdiagramDB.ClassdiagramPointersEncoding = input.ClassdiagramPointersEncoding

	db, _ = db.Model(&classdiagramDB)
	_, err = db.Updates(&classdiagramDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	classdiagramNew := new(models.Classdiagram)
	classdiagramDB.CopyBasicFieldsToClassdiagram(classdiagramNew)

	// redeem pointers
	classdiagramDB.DecodePointers(backRepo, classdiagramNew)

	// get stage instance from DB instance, and call callback function
	classdiagramOld := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]
	if classdiagramOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), classdiagramOld, classdiagramNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the classdiagramDB
	c.JSON(http.StatusOK, classdiagramDB)
}

// DeleteClassdiagram
//
// swagger:route DELETE /classdiagrams/{ID} classdiagrams deleteClassdiagram
//
// # Delete a classdiagram
//
// default: genericError
//
//	200: classdiagramDBResponse
func (controller *Controller) DeleteClassdiagram(c *gin.Context) {

	mutexClassdiagram.Lock()
	defer mutexClassdiagram.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteClassdiagram", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoClassdiagram.GetDB()

	// Get model if exist
	var classdiagramDB orm.ClassdiagramDB
	if _, err := db.First(&classdiagramDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&classdiagramDB)

	// get an instance (not staged) from DB instance, and call callback function
	classdiagramDeleted := new(models.Classdiagram)
	classdiagramDB.CopyBasicFieldsToClassdiagram(classdiagramDeleted)

	// get stage instance from DB instance, and call callback function
	classdiagramStaged := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[classdiagramDB.ID]
	if classdiagramStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), classdiagramStaged, classdiagramDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
