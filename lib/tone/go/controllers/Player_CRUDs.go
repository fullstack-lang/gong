// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"
)

// declaration in order to justify use of the models import
var __Player__dummysDeclaration__ models.Player
var _ = __Player__dummysDeclaration__
var __Player_time__dummyDeclaration time.Duration
var _ = __Player_time__dummyDeclaration

var mutexPlayer sync.Mutex

// An PlayerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updatePlayer
type PlayerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PlayerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updatePlayer
type PlayerInput struct {
	// The Player to submit or modify
	// in: body
	Player *orm.PlayerAPI
}

// UpdatePlayer
//
// swagger:route PATCH /players/{ID} players updatePlayer
//
// # Update a player
//
// Responses:
// default: genericError
//
//	200: playerDBResponse
func (controller *Controller) UpdatePlayer(w http.ResponseWriter, r *http.Request) {

	mutexPlayer.Lock()
	defer mutexPlayer.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoPlayer.GetDB()

	// Validate input
	var input orm.PlayerAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var playerDB orm.PlayerDB

	// fetch the player
	_, err := db.First(&playerDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	playerDB.CopyBasicFieldsFromPlayer_WOP(&input.Player_WOP)
	playerDB.PlayerPointersEncoding = input.PlayerPointersEncoding

	db, _ = db.Model(&playerDB)
	_, err = db.Updates(&playerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	playerNew := new(models.Player)
	playerDB.CopyBasicFieldsToPlayer(playerNew)

	// redeem pointers
	playerDB.DecodePointers(backRepo, playerNew)

	// get stage instance from DB instance, and call callback function
	playerOld := backRepo.BackRepoPlayer.Map_PlayerDBID_PlayerPtr[playerDB.ID]
	if playerOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(playerOld, playerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the playerDB
	writeJSON(w, http.StatusOK, playerDB)
}
