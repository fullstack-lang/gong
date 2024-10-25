// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Animate__dummysDeclaration__ models.Animate
var __Animate_time__dummyDeclaration time.Duration

var mutexAnimate sync.Mutex

// An AnimateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAnimate updateAnimate deleteAnimate
type AnimateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AnimateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAnimate updateAnimate
type AnimateInput struct {
	// The Animate to submit or modify
	// in: body
	Animate *orm.AnimateAPI
}

// GetAnimates
//
// swagger:route GET /animates animates getAnimates
//
// # Get all animates
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) GetAnimates(c *gin.Context) {

	// source slice
	var animateDBs []orm.AnimateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnimates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	_, err := db.Find(&animateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	animateAPIs := make([]orm.AnimateAPI, 0)

	// for each animate, update fields from the database nullable fields
	for idx := range animateDBs {
		animateDB := &animateDBs[idx]
		_ = animateDB
		var animateAPI orm.AnimateAPI

		// insertion point for updating fields
		animateAPI.ID = animateDB.ID
		animateDB.CopyBasicFieldsToAnimate_WOP(&animateAPI.Animate_WOP)
		animateAPI.AnimatePointersEncoding = animateDB.AnimatePointersEncoding
		animateAPIs = append(animateAPIs, animateAPI)
	}

	c.JSON(http.StatusOK, animateAPIs)
}

// PostAnimate
//
// swagger:route POST /animates animates postAnimate
//
// Creates a animate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAnimate(c *gin.Context) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAnimates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	// Validate input
	var input orm.AnimateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create animate
	animateDB := orm.AnimateDB{}
	animateDB.AnimatePointersEncoding = input.AnimatePointersEncoding
	animateDB.CopyBasicFieldsFromAnimate_WOP(&input.Animate_WOP)

	_, err = db.Create(&animateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAnimate.CheckoutPhaseOneInstance(&animateDB)
	animate := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]

	if animate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), animate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, animateDB)
}

// GetAnimate
//
// swagger:route GET /animates/{ID} animates getAnimate
//
// Gets the details for a animate.
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) GetAnimate(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnimate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	// Get animateDB in DB
	var animateDB orm.AnimateDB
	if _, err := db.First(&animateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var animateAPI orm.AnimateAPI
	animateAPI.ID = animateDB.ID
	animateAPI.AnimatePointersEncoding = animateDB.AnimatePointersEncoding
	animateDB.CopyBasicFieldsToAnimate_WOP(&animateAPI.Animate_WOP)

	c.JSON(http.StatusOK, animateAPI)
}

// UpdateAnimate
//
// swagger:route PATCH /animates/{ID} animates updateAnimate
//
// # Update a animate
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) UpdateAnimate(c *gin.Context) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAnimate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	// Validate input
	var input orm.AnimateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var animateDB orm.AnimateDB

	// fetch the animate
	_, err := db.First(&animateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	animateDB.CopyBasicFieldsFromAnimate_WOP(&input.Animate_WOP)
	animateDB.AnimatePointersEncoding = input.AnimatePointersEncoding

	db, _ = db.Model(&animateDB)
	_, err = db.Updates(&animateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	animateNew := new(models.Animate)
	animateDB.CopyBasicFieldsToAnimate(animateNew)

	// redeem pointers
	animateDB.DecodePointers(backRepo, animateNew)

	// get stage instance from DB instance, and call callback function
	animateOld := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	if animateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), animateOld, animateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the animateDB
	c.JSON(http.StatusOK, animateDB)
}

// DeleteAnimate
//
// swagger:route DELETE /animates/{ID} animates deleteAnimate
//
// # Delete a animate
//
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) DeleteAnimate(c *gin.Context) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAnimate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	// Get model if exist
	var animateDB orm.AnimateDB
	if _, err := db.First(&animateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&animateDB)

	// get an instance (not staged) from DB instance, and call callback function
	animateDeleted := new(models.Animate)
	animateDB.CopyBasicFieldsToAnimate(animateDeleted)

	// get stage instance from DB instance, and call callback function
	animateStaged := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	if animateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), animateStaged, animateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
