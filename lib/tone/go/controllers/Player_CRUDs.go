// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
	"github.com/fullstack-lang/gong/lib/tone/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Player__dummysDeclaration__ models.Player
var __Player_time__dummyDeclaration time.Duration

var mutexPlayer sync.Mutex

// An PlayerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPlayer updatePlayer deletePlayer
type PlayerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PlayerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPlayer updatePlayer
type PlayerInput struct {
	// The Player to submit or modify
	// in: body
	Player *orm.PlayerAPI
}

// GetPlayers
//
// swagger:route GET /players players getPlayers
//
// # Get all players
//
// Responses:
// default: genericError
//
//	200: playerDBResponse
func (controller *Controller) GetPlayers(c *gin.Context) {

	// source slice
	var playerDBs []orm.PlayerDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlayers", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoPlayer.GetDB()

	_, err := db.Find(&playerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	playerAPIs := make([]orm.PlayerAPI, 0)

	// for each player, update fields from the database nullable fields
	for idx := range playerDBs {
		playerDB := &playerDBs[idx]
		_ = playerDB
		var playerAPI orm.PlayerAPI

		// insertion point for updating fields
		playerAPI.ID = playerDB.ID
		playerDB.CopyBasicFieldsToPlayer_WOP(&playerAPI.Player_WOP)
		playerAPI.PlayerPointersEncoding = playerDB.PlayerPointersEncoding
		playerAPIs = append(playerAPIs, playerAPI)
	}

	c.JSON(http.StatusOK, playerAPIs)
}

// PostPlayer
//
// swagger:route POST /players players postPlayer
//
// Creates a player
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPlayer(c *gin.Context) {

	mutexPlayer.Lock()
	defer mutexPlayer.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPlayers", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoPlayer.GetDB()

	// Validate input
	var input orm.PlayerAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create player
	playerDB := orm.PlayerDB{}
	playerDB.PlayerPointersEncoding = input.PlayerPointersEncoding
	playerDB.CopyBasicFieldsFromPlayer_WOP(&input.Player_WOP)

	_, err = db.Create(&playerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPlayer.CheckoutPhaseOneInstance(&playerDB)
	player := backRepo.BackRepoPlayer.Map_PlayerDBID_PlayerPtr[playerDB.ID]

	if player != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), player)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, playerDB)
}

// GetPlayer
//
// swagger:route GET /players/{ID} players getPlayer
//
// Gets the details for a player.
//
// Responses:
// default: genericError
//
//	200: playerDBResponse
func (controller *Controller) GetPlayer(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlayer", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoPlayer.GetDB()

	// Get playerDB in DB
	var playerDB orm.PlayerDB
	if _, err := db.First(&playerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var playerAPI orm.PlayerAPI
	playerAPI.ID = playerDB.ID
	playerAPI.PlayerPointersEncoding = playerDB.PlayerPointersEncoding
	playerDB.CopyBasicFieldsToPlayer_WOP(&playerAPI.Player_WOP)

	c.JSON(http.StatusOK, playerAPI)
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
func (controller *Controller) UpdatePlayer(c *gin.Context) {

	mutexPlayer.Lock()
	defer mutexPlayer.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePlayer", "Name", stackPath)
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
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var playerDB orm.PlayerDB

	// fetch the player
	_, err := db.First(&playerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
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
		c.JSON(http.StatusBadRequest, returnError.Body)
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
		models.AfterUpdateFromFront(backRepo.GetStage(), playerOld, playerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the playerDB
	c.JSON(http.StatusOK, playerDB)
}

// DeletePlayer
//
// swagger:route DELETE /players/{ID} players deletePlayer
//
// # Delete a player
//
// default: genericError
//
//	200: playerDBResponse
func (controller *Controller) DeletePlayer(c *gin.Context) {

	mutexPlayer.Lock()
	defer mutexPlayer.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePlayer", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/tone/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoPlayer.GetDB()

	// Get model if exist
	var playerDB orm.PlayerDB
	if _, err := db.First(&playerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&playerDB)

	// get an instance (not staged) from DB instance, and call callback function
	playerDeleted := new(models.Player)
	playerDB.CopyBasicFieldsToPlayer(playerDeleted)

	// get stage instance from DB instance, and call callback function
	playerStaged := backRepo.BackRepoPlayer.Map_PlayerDBID_PlayerPtr[playerDB.ID]
	if playerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), playerStaged, playerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
