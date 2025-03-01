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
var __Group__dummysDeclaration__ models.Group
var __Group_time__dummyDeclaration time.Duration

var mutexGroup sync.Mutex

// An GroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGroup updateGroup deleteGroup
type GroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroup updateGroup
type GroupInput struct {
	// The Group to submit or modify
	// in: body
	Group *orm.GroupAPI
}

// GetGroups
//
// swagger:route GET /groups groups getGroups
//
// # Get all groups
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) GetGroups(c *gin.Context) {

	// source slice
	var groupDBs []orm.GroupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/button/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup.GetDB()

	_, err := db.Find(&groupDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	groupAPIs := make([]orm.GroupAPI, 0)

	// for each group, update fields from the database nullable fields
	for idx := range groupDBs {
		groupDB := &groupDBs[idx]
		_ = groupDB
		var groupAPI orm.GroupAPI

		// insertion point for updating fields
		groupAPI.ID = groupDB.ID
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupAPIs = append(groupAPIs, groupAPI)
	}

	c.JSON(http.StatusOK, groupAPIs)
}

// PostGroup
//
// swagger:route POST /groups groups postGroup
//
// Creates a group
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroup(c *gin.Context) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/button/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Validate input
	var input orm.GroupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create group
	groupDB := orm.GroupDB{}
	groupDB.GroupPointersEncoding = input.GroupPointersEncoding
	groupDB.CopyBasicFieldsFromGroup_WOP(&input.Group_WOP)

	_, err = db.Create(&groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroup.CheckoutPhaseOneInstance(&groupDB)
	group := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]

	if group != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), group)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, groupDB)
}

// GetGroup
//
// swagger:route GET /groups/{ID} groups getGroup
//
// Gets the details for a group.
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) GetGroup(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/button/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Get groupDB in DB
	var groupDB orm.GroupDB
	if _, err := db.First(&groupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var groupAPI orm.GroupAPI
	groupAPI.ID = groupDB.ID
	groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
	groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

	c.JSON(http.StatusOK, groupAPI)
}

// UpdateGroup
//
// swagger:route PATCH /groups/{ID} groups updateGroup
//
// # Update a group
//
// Responses:
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) UpdateGroup(c *gin.Context) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/button/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Validate input
	var input orm.GroupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var groupDB orm.GroupDB

	// fetch the group
	_, err := db.First(&groupDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	groupDB.CopyBasicFieldsFromGroup_WOP(&input.Group_WOP)
	groupDB.GroupPointersEncoding = input.GroupPointersEncoding

	db, _ = db.Model(&groupDB)
	_, err = db.Updates(&groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	groupNew := new(models.Group)
	groupDB.CopyBasicFieldsToGroup(groupNew)

	// redeem pointers
	groupDB.DecodePointers(backRepo, groupNew)

	// get stage instance from DB instance, and call callback function
	groupOld := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]
	if groupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), groupOld, groupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the groupDB
	c.JSON(http.StatusOK, groupDB)
}

// DeleteGroup
//
// swagger:route DELETE /groups/{ID} groups deleteGroup
//
// # Delete a group
//
// default: genericError
//
//	200: groupDBResponse
func (controller *Controller) DeleteGroup(c *gin.Context) {

	mutexGroup.Lock()
	defer mutexGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/lib/button/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup.GetDB()

	// Get model if exist
	var groupDB orm.GroupDB
	if _, err := db.First(&groupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&groupDB)

	// get an instance (not staged) from DB instance, and call callback function
	groupDeleted := new(models.Group)
	groupDB.CopyBasicFieldsToGroup(groupDeleted)

	// get stage instance from DB instance, and call callback function
	groupStaged := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[groupDB.ID]
	if groupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), groupStaged, groupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
