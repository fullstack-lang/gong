// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Menu__dummysDeclaration__ models.Menu
var _ = __Menu__dummysDeclaration__
var __Menu_time__dummyDeclaration time.Duration
var _ = __Menu_time__dummyDeclaration

var mutexMenu sync.Mutex

// An MenuID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMenu updateMenu deleteMenu
type MenuID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MenuInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMenu updateMenu
type MenuInput struct {
	// The Menu to submit or modify
	// in: body
	Menu *orm.MenuAPI
}

// GetMenus
//
// swagger:route GET /menus menus getMenus
//
// # Get all menus
//
// Responses:
// default: genericError
//
//	200: menuDBResponse
func (controller *Controller) GetMenus(c *gin.Context) {

	// source slice
	var menuDBs []orm.MenuDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMenus", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMenu.GetDB()

	_, err := db.Find(&menuDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	menuAPIs := make([]orm.MenuAPI, 0)

	// for each menu, update fields from the database nullable fields
	for idx := range menuDBs {
		menuDB := &menuDBs[idx]
		_ = menuDB
		var menuAPI orm.MenuAPI

		// insertion point for updating fields
		menuAPI.ID = menuDB.ID
		menuDB.CopyBasicFieldsToMenu_WOP(&menuAPI.Menu_WOP)
		menuAPI.MenuPointersEncoding = menuDB.MenuPointersEncoding
		menuAPIs = append(menuAPIs, menuAPI)
	}

	c.JSON(http.StatusOK, menuAPIs)
}

// PostMenu
//
// swagger:route POST /menus menus postMenu
//
// Creates a menu
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMenu(c *gin.Context) {

	mutexMenu.Lock()
	defer mutexMenu.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMenus", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMenu.GetDB()

	// Validate input
	var input orm.MenuAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create menu
	menuDB := orm.MenuDB{}
	menuDB.MenuPointersEncoding = input.MenuPointersEncoding
	menuDB.CopyBasicFieldsFromMenu_WOP(&input.Menu_WOP)

	_, err = db.Create(&menuDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMenu.CheckoutPhaseOneInstance(&menuDB)
	menu := backRepo.BackRepoMenu.Map_MenuDBID_MenuPtr[menuDB.ID]

	if menu != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), menu)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, menuDB)
}

// GetMenu
//
// swagger:route GET /menus/{ID} menus getMenu
//
// Gets the details for a menu.
//
// Responses:
// default: genericError
//
//	200: menuDBResponse
func (controller *Controller) GetMenu(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMenu", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMenu.GetDB()

	// Get menuDB in DB
	var menuDB orm.MenuDB
	if _, err := db.First(&menuDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var menuAPI orm.MenuAPI
	menuAPI.ID = menuDB.ID
	menuAPI.MenuPointersEncoding = menuDB.MenuPointersEncoding
	menuDB.CopyBasicFieldsToMenu_WOP(&menuAPI.Menu_WOP)

	c.JSON(http.StatusOK, menuAPI)
}

// UpdateMenu
//
// swagger:route PATCH /menus/{ID} menus updateMenu
//
// # Update a menu
//
// Responses:
// default: genericError
//
//	200: menuDBResponse
func (controller *Controller) UpdateMenu(c *gin.Context) {

	mutexMenu.Lock()
	defer mutexMenu.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMenu.GetDB()

	// Validate input
	var input orm.MenuAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var menuDB orm.MenuDB

	// fetch the menu
	_, err := db.First(&menuDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	menuDB.CopyBasicFieldsFromMenu_WOP(&input.Menu_WOP)
	menuDB.MenuPointersEncoding = input.MenuPointersEncoding

	db, _ = db.Model(&menuDB)
	_, err = db.Updates(&menuDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	menuNew := new(models.Menu)
	menuDB.CopyBasicFieldsToMenu(menuNew)

	// redeem pointers
	menuDB.DecodePointers(backRepo, menuNew)

	// get stage instance from DB instance, and call callback function
	menuOld := backRepo.BackRepoMenu.Map_MenuDBID_MenuPtr[menuDB.ID]
	if menuOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), menuOld, menuNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the menuDB
	c.JSON(http.StatusOK, menuDB)
}

// DeleteMenu
//
// swagger:route DELETE /menus/{ID} menus deleteMenu
//
// # Delete a menu
//
// default: genericError
//
//	200: menuDBResponse
func (controller *Controller) DeleteMenu(c *gin.Context) {

	mutexMenu.Lock()
	defer mutexMenu.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMenu", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMenu.GetDB()

	// Get model if exist
	var menuDB orm.MenuDB
	if _, err := db.First(&menuDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&menuDB)

	// get an instance (not staged) from DB instance, and call callback function
	menuDeleted := new(models.Menu)
	menuDB.CopyBasicFieldsToMenu(menuDeleted)

	// get stage instance from DB instance, and call callback function
	menuStaged := backRepo.BackRepoMenu.Map_MenuDBID_MenuPtr[menuDB.ID]
	if menuStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), menuStaged, menuDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
