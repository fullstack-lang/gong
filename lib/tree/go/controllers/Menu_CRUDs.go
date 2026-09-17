// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/orm"
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
// swagger:parameters updateMenu
type MenuID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MenuInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateMenu
type MenuInput struct {
	// The Menu to submit or modify
	// in: body
	Menu *orm.MenuAPI
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
func (controller *Controller) UpdateMenu(w http.ResponseWriter, r *http.Request) {

	mutexMenu.Lock()
	defer mutexMenu.Unlock()

	_values := r.URL.Query()
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
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var menuDB orm.MenuDB

	// fetch the menu
	_, err := db.First(&menuDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
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
		writeJSON(w, http.StatusBadRequest, returnError.Body)
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
		backRepo.GetStage().OnAfterUpdateFromFront(menuOld, menuNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the menuDB
	writeJSON(w, http.StatusOK, menuDB)
}
