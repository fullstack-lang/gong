// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FavIcon__dummysDeclaration__ models.FavIcon
var __FavIcon_time__dummyDeclaration time.Duration

var mutexFavIcon sync.Mutex

// An FavIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFavIcon updateFavIcon deleteFavIcon
type FavIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FavIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFavIcon updateFavIcon
type FavIconInput struct {
	// The FavIcon to submit or modify
	// in: body
	FavIcon *orm.FavIconAPI
}

// GetFavIcons
//
// swagger:route GET /favicons favicons getFavIcons
//
// # Get all favicons
//
// Responses:
// default: genericError
//
//	200: faviconDBResponse
func (controller *Controller) GetFavIcons(c *gin.Context) {

	// source slice
	var faviconDBs []orm.FavIconDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFavIcons", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFavIcon.GetDB()

	_, err := db.Find(&faviconDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	faviconAPIs := make([]orm.FavIconAPI, 0)

	// for each favicon, update fields from the database nullable fields
	for idx := range faviconDBs {
		faviconDB := &faviconDBs[idx]
		_ = faviconDB
		var faviconAPI orm.FavIconAPI

		// insertion point for updating fields
		faviconAPI.ID = faviconDB.ID
		faviconDB.CopyBasicFieldsToFavIcon_WOP(&faviconAPI.FavIcon_WOP)
		faviconAPI.FavIconPointersEncoding = faviconDB.FavIconPointersEncoding
		faviconAPIs = append(faviconAPIs, faviconAPI)
	}

	c.JSON(http.StatusOK, faviconAPIs)
}

// PostFavIcon
//
// swagger:route POST /favicons favicons postFavIcon
//
// Creates a favicon
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFavIcon(c *gin.Context) {

	mutexFavIcon.Lock()
	defer mutexFavIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFavIcons", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFavIcon.GetDB()

	// Validate input
	var input orm.FavIconAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create favicon
	faviconDB := orm.FavIconDB{}
	faviconDB.FavIconPointersEncoding = input.FavIconPointersEncoding
	faviconDB.CopyBasicFieldsFromFavIcon_WOP(&input.FavIcon_WOP)

	_, err = db.Create(&faviconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFavIcon.CheckoutPhaseOneInstance(&faviconDB)
	favicon := backRepo.BackRepoFavIcon.Map_FavIconDBID_FavIconPtr[faviconDB.ID]

	if favicon != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), favicon)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, faviconDB)
}

// GetFavIcon
//
// swagger:route GET /favicons/{ID} favicons getFavIcon
//
// Gets the details for a favicon.
//
// Responses:
// default: genericError
//
//	200: faviconDBResponse
func (controller *Controller) GetFavIcon(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFavIcon", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFavIcon.GetDB()

	// Get faviconDB in DB
	var faviconDB orm.FavIconDB
	if _, err := db.First(&faviconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var faviconAPI orm.FavIconAPI
	faviconAPI.ID = faviconDB.ID
	faviconAPI.FavIconPointersEncoding = faviconDB.FavIconPointersEncoding
	faviconDB.CopyBasicFieldsToFavIcon_WOP(&faviconAPI.FavIcon_WOP)

	c.JSON(http.StatusOK, faviconAPI)
}

// UpdateFavIcon
//
// swagger:route PATCH /favicons/{ID} favicons updateFavIcon
//
// # Update a favicon
//
// Responses:
// default: genericError
//
//	200: faviconDBResponse
func (controller *Controller) UpdateFavIcon(c *gin.Context) {

	mutexFavIcon.Lock()
	defer mutexFavIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFavIcon", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFavIcon.GetDB()

	// Validate input
	var input orm.FavIconAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var faviconDB orm.FavIconDB

	// fetch the favicon
	_, err := db.First(&faviconDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	faviconDB.CopyBasicFieldsFromFavIcon_WOP(&input.FavIcon_WOP)
	faviconDB.FavIconPointersEncoding = input.FavIconPointersEncoding

	db, _ = db.Model(&faviconDB)
	_, err = db.Updates(&faviconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	faviconNew := new(models.FavIcon)
	faviconDB.CopyBasicFieldsToFavIcon(faviconNew)

	// redeem pointers
	faviconDB.DecodePointers(backRepo, faviconNew)

	// get stage instance from DB instance, and call callback function
	faviconOld := backRepo.BackRepoFavIcon.Map_FavIconDBID_FavIconPtr[faviconDB.ID]
	if faviconOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), faviconOld, faviconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the faviconDB
	c.JSON(http.StatusOK, faviconDB)
}

// DeleteFavIcon
//
// swagger:route DELETE /favicons/{ID} favicons deleteFavIcon
//
// # Delete a favicon
//
// default: genericError
//
//	200: faviconDBResponse
func (controller *Controller) DeleteFavIcon(c *gin.Context) {

	mutexFavIcon.Lock()
	defer mutexFavIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFavIcon", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFavIcon.GetDB()

	// Get model if exist
	var faviconDB orm.FavIconDB
	if _, err := db.First(&faviconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&faviconDB)

	// get an instance (not staged) from DB instance, and call callback function
	faviconDeleted := new(models.FavIcon)
	faviconDB.CopyBasicFieldsToFavIcon(faviconDeleted)

	// get stage instance from DB instance, and call callback function
	faviconStaged := backRepo.BackRepoFavIcon.Map_FavIconDBID_FavIconPtr[faviconDB.ID]
	if faviconStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), faviconStaged, faviconDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
