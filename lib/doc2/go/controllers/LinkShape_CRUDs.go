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
var __LinkShape__dummysDeclaration__ models.LinkShape
var __LinkShape_time__dummyDeclaration time.Duration

var mutexLinkShape sync.Mutex

// An LinkShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLinkShape updateLinkShape deleteLinkShape
type LinkShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLinkShape updateLinkShape
type LinkShapeInput struct {
	// The LinkShape to submit or modify
	// in: body
	LinkShape *orm.LinkShapeAPI
}

// GetLinkShapes
//
// swagger:route GET /linkshapes linkshapes getLinkShapes
//
// # Get all linkshapes
//
// Responses:
// default: genericError
//
//	200: linkshapeDBResponse
func (controller *Controller) GetLinkShapes(c *gin.Context) {

	// source slice
	var linkshapeDBs []orm.LinkShapeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkShapes", "Name", stackPath)
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
	db := backRepo.BackRepoLinkShape.GetDB()

	_, err := db.Find(&linkshapeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linkshapeAPIs := make([]orm.LinkShapeAPI, 0)

	// for each linkshape, update fields from the database nullable fields
	for idx := range linkshapeDBs {
		linkshapeDB := &linkshapeDBs[idx]
		_ = linkshapeDB
		var linkshapeAPI orm.LinkShapeAPI

		// insertion point for updating fields
		linkshapeAPI.ID = linkshapeDB.ID
		linkshapeDB.CopyBasicFieldsToLinkShape_WOP(&linkshapeAPI.LinkShape_WOP)
		linkshapeAPI.LinkShapePointersEncoding = linkshapeDB.LinkShapePointersEncoding
		linkshapeAPIs = append(linkshapeAPIs, linkshapeAPI)
	}

	c.JSON(http.StatusOK, linkshapeAPIs)
}

// PostLinkShape
//
// swagger:route POST /linkshapes linkshapes postLinkShape
//
// Creates a linkshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLinkShape(c *gin.Context) {

	mutexLinkShape.Lock()
	defer mutexLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLinkShapes", "Name", stackPath)
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
	db := backRepo.BackRepoLinkShape.GetDB()

	// Validate input
	var input orm.LinkShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create linkshape
	linkshapeDB := orm.LinkShapeDB{}
	linkshapeDB.LinkShapePointersEncoding = input.LinkShapePointersEncoding
	linkshapeDB.CopyBasicFieldsFromLinkShape_WOP(&input.LinkShape_WOP)

	_, err = db.Create(&linkshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLinkShape.CheckoutPhaseOneInstance(&linkshapeDB)
	linkshape := backRepo.BackRepoLinkShape.Map_LinkShapeDBID_LinkShapePtr[linkshapeDB.ID]

	if linkshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), linkshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, linkshapeDB)
}

// GetLinkShape
//
// swagger:route GET /linkshapes/{ID} linkshapes getLinkShape
//
// Gets the details for a linkshape.
//
// Responses:
// default: genericError
//
//	200: linkshapeDBResponse
func (controller *Controller) GetLinkShape(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoLinkShape.GetDB()

	// Get linkshapeDB in DB
	var linkshapeDB orm.LinkShapeDB
	if _, err := db.First(&linkshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var linkshapeAPI orm.LinkShapeAPI
	linkshapeAPI.ID = linkshapeDB.ID
	linkshapeAPI.LinkShapePointersEncoding = linkshapeDB.LinkShapePointersEncoding
	linkshapeDB.CopyBasicFieldsToLinkShape_WOP(&linkshapeAPI.LinkShape_WOP)

	c.JSON(http.StatusOK, linkshapeAPI)
}

// UpdateLinkShape
//
// swagger:route PATCH /linkshapes/{ID} linkshapes updateLinkShape
//
// # Update a linkshape
//
// Responses:
// default: genericError
//
//	200: linkshapeDBResponse
func (controller *Controller) UpdateLinkShape(c *gin.Context) {

	mutexLinkShape.Lock()
	defer mutexLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoLinkShape.GetDB()

	// Validate input
	var input orm.LinkShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkshapeDB orm.LinkShapeDB

	// fetch the linkshape
	_, err := db.First(&linkshapeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkshapeDB.CopyBasicFieldsFromLinkShape_WOP(&input.LinkShape_WOP)
	linkshapeDB.LinkShapePointersEncoding = input.LinkShapePointersEncoding

	db, _ = db.Model(&linkshapeDB)
	_, err = db.Updates(&linkshapeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkshapeNew := new(models.LinkShape)
	linkshapeDB.CopyBasicFieldsToLinkShape(linkshapeNew)

	// redeem pointers
	linkshapeDB.DecodePointers(backRepo, linkshapeNew)

	// get stage instance from DB instance, and call callback function
	linkshapeOld := backRepo.BackRepoLinkShape.Map_LinkShapeDBID_LinkShapePtr[linkshapeDB.ID]
	if linkshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), linkshapeOld, linkshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkshapeDB
	c.JSON(http.StatusOK, linkshapeDB)
}

// DeleteLinkShape
//
// swagger:route DELETE /linkshapes/{ID} linkshapes deleteLinkShape
//
// # Delete a linkshape
//
// default: genericError
//
//	200: linkshapeDBResponse
func (controller *Controller) DeleteLinkShape(c *gin.Context) {

	mutexLinkShape.Lock()
	defer mutexLinkShape.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLinkShape", "Name", stackPath)
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
	db := backRepo.BackRepoLinkShape.GetDB()

	// Get model if exist
	var linkshapeDB orm.LinkShapeDB
	if _, err := db.First(&linkshapeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&linkshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	linkshapeDeleted := new(models.LinkShape)
	linkshapeDB.CopyBasicFieldsToLinkShape(linkshapeDeleted)

	// get stage instance from DB instance, and call callback function
	linkshapeStaged := backRepo.BackRepoLinkShape.Map_LinkShapeDBID_LinkShapePtr[linkshapeDB.ID]
	if linkshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), linkshapeStaged, linkshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
