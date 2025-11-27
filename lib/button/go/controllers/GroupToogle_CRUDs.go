// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GroupToogle__dummysDeclaration__ models.GroupToogle
var __GroupToogle_time__dummyDeclaration time.Duration

var mutexGroupToogle sync.Mutex

// An GroupToogleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGroupToogle updateGroupToogle deleteGroupToogle
type GroupToogleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupToogleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroupToogle updateGroupToogle
type GroupToogleInput struct {
	// The GroupToogle to submit or modify
	// in: body
	GroupToogle *orm.GroupToogleAPI
}

// GetGroupToogles
//
// swagger:route GET /grouptoogles grouptoogles getGroupToogles
//
// # Get all grouptoogles
//
// Responses:
// default: genericError
//
//	200: grouptoogleDBResponse
func (controller *Controller) GetGroupToogles(c *gin.Context) {

	// source slice
	var grouptoogleDBs []orm.GroupToogleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroupToogles", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroupToogle.GetDB()

	_, err := db.Find(&grouptoogleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	grouptoogleAPIs := make([]orm.GroupToogleAPI, 0)

	// for each grouptoogle, update fields from the database nullable fields
	for idx := range grouptoogleDBs {
		grouptoogleDB := &grouptoogleDBs[idx]
		_ = grouptoogleDB
		var grouptoogleAPI orm.GroupToogleAPI

		// insertion point for updating fields
		grouptoogleAPI.ID = grouptoogleDB.ID
		grouptoogleDB.CopyBasicFieldsToGroupToogle_WOP(&grouptoogleAPI.GroupToogle_WOP)
		grouptoogleAPI.GroupTooglePointersEncoding = grouptoogleDB.GroupTooglePointersEncoding
		grouptoogleAPIs = append(grouptoogleAPIs, grouptoogleAPI)
	}

	c.JSON(http.StatusOK, grouptoogleAPIs)
}

// PostGroupToogle
//
// swagger:route POST /grouptoogles grouptoogles postGroupToogle
//
// Creates a grouptoogle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroupToogle(c *gin.Context) {

	mutexGroupToogle.Lock()
	defer mutexGroupToogle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroupToogles", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroupToogle.GetDB()

	// Validate input
	var input orm.GroupToogleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create grouptoogle
	grouptoogleDB := orm.GroupToogleDB{}
	grouptoogleDB.GroupTooglePointersEncoding = input.GroupTooglePointersEncoding
	grouptoogleDB.CopyBasicFieldsFromGroupToogle_WOP(&input.GroupToogle_WOP)

	_, err = db.Create(&grouptoogleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroupToogle.CheckoutPhaseOneInstance(&grouptoogleDB)
	grouptoogle := backRepo.BackRepoGroupToogle.Map_GroupToogleDBID_GroupTooglePtr[grouptoogleDB.ID]

	if grouptoogle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), grouptoogle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, grouptoogleDB)
}

// GetGroupToogle
//
// swagger:route GET /grouptoogles/{ID} grouptoogles getGroupToogle
//
// Gets the details for a grouptoogle.
//
// Responses:
// default: genericError
//
//	200: grouptoogleDBResponse
func (controller *Controller) GetGroupToogle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroupToogle", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroupToogle.GetDB()

	// Get grouptoogleDB in DB
	var grouptoogleDB orm.GroupToogleDB
	if _, err := db.First(&grouptoogleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var grouptoogleAPI orm.GroupToogleAPI
	grouptoogleAPI.ID = grouptoogleDB.ID
	grouptoogleAPI.GroupTooglePointersEncoding = grouptoogleDB.GroupTooglePointersEncoding
	grouptoogleDB.CopyBasicFieldsToGroupToogle_WOP(&grouptoogleAPI.GroupToogle_WOP)

	c.JSON(http.StatusOK, grouptoogleAPI)
}

// UpdateGroupToogle
//
// swagger:route PATCH /grouptoogles/{ID} grouptoogles updateGroupToogle
//
// # Update a grouptoogle
//
// Responses:
// default: genericError
//
//	200: grouptoogleDBResponse
func (controller *Controller) UpdateGroupToogle(c *gin.Context) {

	mutexGroupToogle.Lock()
	defer mutexGroupToogle.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroupToogle.GetDB()

	// Validate input
	var input orm.GroupToogleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var grouptoogleDB orm.GroupToogleDB

	// fetch the grouptoogle
	_, err := db.First(&grouptoogleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	grouptoogleDB.CopyBasicFieldsFromGroupToogle_WOP(&input.GroupToogle_WOP)
	grouptoogleDB.GroupTooglePointersEncoding = input.GroupTooglePointersEncoding

	db, _ = db.Model(&grouptoogleDB)
	_, err = db.Updates(&grouptoogleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	grouptoogleNew := new(models.GroupToogle)
	grouptoogleDB.CopyBasicFieldsToGroupToogle(grouptoogleNew)

	// redeem pointers
	grouptoogleDB.DecodePointers(backRepo, grouptoogleNew)

	// get stage instance from DB instance, and call callback function
	grouptoogleOld := backRepo.BackRepoGroupToogle.Map_GroupToogleDBID_GroupTooglePtr[grouptoogleDB.ID]
	if grouptoogleOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), grouptoogleOld, grouptoogleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the grouptoogleDB
	c.JSON(http.StatusOK, grouptoogleDB)
}

// DeleteGroupToogle
//
// swagger:route DELETE /grouptoogles/{ID} grouptoogles deleteGroupToogle
//
// # Delete a grouptoogle
//
// default: genericError
//
//	200: grouptoogleDBResponse
func (controller *Controller) DeleteGroupToogle(c *gin.Context) {

	mutexGroupToogle.Lock()
	defer mutexGroupToogle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroupToogle", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/button/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGroupToogle.GetDB()

	// Get model if exist
	var grouptoogleDB orm.GroupToogleDB
	if _, err := db.First(&grouptoogleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&grouptoogleDB)

	// get an instance (not staged) from DB instance, and call callback function
	grouptoogleDeleted := new(models.GroupToogle)
	grouptoogleDB.CopyBasicFieldsToGroupToogle(grouptoogleDeleted)

	// get stage instance from DB instance, and call callback function
	grouptoogleStaged := backRepo.BackRepoGroupToogle.Map_GroupToogleDBID_GroupTooglePtr[grouptoogleDB.ID]
	if grouptoogleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), grouptoogleStaged, grouptoogleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
