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
var __Role__dummysDeclaration__ models.Role
var __Role_time__dummyDeclaration time.Duration

var mutexRole sync.Mutex

// An RoleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRole updateRole deleteRole
type RoleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RoleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRole updateRole
type RoleInput struct {
	// The Role to submit or modify
	// in: body
	Role *orm.RoleAPI
}

// GetRoles
//
// swagger:route GET /roles roles getRoles
//
// # Get all roles
//
// Responses:
// default: genericError
//
//	200: roleDBResponse
func (controller *Controller) GetRoles(c *gin.Context) {

	// source slice
	var roleDBs []orm.RoleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRoles", "Name", stackPath)
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
	db := backRepo.BackRepoRole.GetDB()

	_, err := db.Find(&roleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	roleAPIs := make([]orm.RoleAPI, 0)

	// for each role, update fields from the database nullable fields
	for idx := range roleDBs {
		roleDB := &roleDBs[idx]
		_ = roleDB
		var roleAPI orm.RoleAPI

		// insertion point for updating fields
		roleAPI.ID = roleDB.ID
		roleDB.CopyBasicFieldsToRole_WOP(&roleAPI.Role_WOP)
		roleAPI.RolePointersEncoding = roleDB.RolePointersEncoding
		roleAPIs = append(roleAPIs, roleAPI)
	}

	c.JSON(http.StatusOK, roleAPIs)
}

// PostRole
//
// swagger:route POST /roles roles postRole
//
// Creates a role
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRole(c *gin.Context) {

	mutexRole.Lock()
	defer mutexRole.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRoles", "Name", stackPath)
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
	db := backRepo.BackRepoRole.GetDB()

	// Validate input
	var input orm.RoleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create role
	roleDB := orm.RoleDB{}
	roleDB.RolePointersEncoding = input.RolePointersEncoding
	roleDB.CopyBasicFieldsFromRole_WOP(&input.Role_WOP)

	_, err = db.Create(&roleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRole.CheckoutPhaseOneInstance(&roleDB)
	role := backRepo.BackRepoRole.Map_RoleDBID_RolePtr[roleDB.ID]

	if role != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), role)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, roleDB)
}

// GetRole
//
// swagger:route GET /roles/{ID} roles getRole
//
// Gets the details for a role.
//
// Responses:
// default: genericError
//
//	200: roleDBResponse
func (controller *Controller) GetRole(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRole", "Name", stackPath)
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
	db := backRepo.BackRepoRole.GetDB()

	// Get roleDB in DB
	var roleDB orm.RoleDB
	if _, err := db.First(&roleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var roleAPI orm.RoleAPI
	roleAPI.ID = roleDB.ID
	roleAPI.RolePointersEncoding = roleDB.RolePointersEncoding
	roleDB.CopyBasicFieldsToRole_WOP(&roleAPI.Role_WOP)

	c.JSON(http.StatusOK, roleAPI)
}

// UpdateRole
//
// swagger:route PATCH /roles/{ID} roles updateRole
//
// # Update a role
//
// Responses:
// default: genericError
//
//	200: roleDBResponse
func (controller *Controller) UpdateRole(c *gin.Context) {

	mutexRole.Lock()
	defer mutexRole.Unlock()

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
	db := backRepo.BackRepoRole.GetDB()

	// Validate input
	var input orm.RoleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var roleDB orm.RoleDB

	// fetch the role
	_, err := db.First(&roleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	roleDB.CopyBasicFieldsFromRole_WOP(&input.Role_WOP)
	roleDB.RolePointersEncoding = input.RolePointersEncoding

	db, _ = db.Model(&roleDB)
	_, err = db.Updates(&roleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	roleNew := new(models.Role)
	roleDB.CopyBasicFieldsToRole(roleNew)

	// redeem pointers
	roleDB.DecodePointers(backRepo, roleNew)

	// get stage instance from DB instance, and call callback function
	roleOld := backRepo.BackRepoRole.Map_RoleDBID_RolePtr[roleDB.ID]
	if roleOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), roleOld, roleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the roleDB
	c.JSON(http.StatusOK, roleDB)
}

// DeleteRole
//
// swagger:route DELETE /roles/{ID} roles deleteRole
//
// # Delete a role
//
// default: genericError
//
//	200: roleDBResponse
func (controller *Controller) DeleteRole(c *gin.Context) {

	mutexRole.Lock()
	defer mutexRole.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRole", "Name", stackPath)
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
	db := backRepo.BackRepoRole.GetDB()

	// Get model if exist
	var roleDB orm.RoleDB
	if _, err := db.First(&roleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&roleDB)

	// get an instance (not staged) from DB instance, and call callback function
	roleDeleted := new(models.Role)
	roleDB.CopyBasicFieldsToRole(roleDeleted)

	// get stage instance from DB instance, and call callback function
	roleStaged := backRepo.BackRepoRole.Map_RoleDBID_RolePtr[roleDB.ID]
	if roleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), roleStaged, roleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
