// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __LinkAnchoredPath__dummysDeclaration__ models.LinkAnchoredPath
var _ = __LinkAnchoredPath__dummysDeclaration__
var __LinkAnchoredPath_time__dummyDeclaration time.Duration
var _ = __LinkAnchoredPath_time__dummyDeclaration

var mutexLinkAnchoredPath sync.Mutex

// An LinkAnchoredPathID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLinkAnchoredPath updateLinkAnchoredPath deleteLinkAnchoredPath
type LinkAnchoredPathID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkAnchoredPathInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLinkAnchoredPath updateLinkAnchoredPath
type LinkAnchoredPathInput struct {
	// The LinkAnchoredPath to submit or modify
	// in: body
	LinkAnchoredPath *orm.LinkAnchoredPathAPI
}

// GetLinkAnchoredPaths
//
// swagger:route GET /linkanchoredpaths linkanchoredpaths getLinkAnchoredPaths
//
// # Get all linkanchoredpaths
//
// Responses:
// default: genericError
//
//	200: linkanchoredpathDBResponse
func (controller *Controller) GetLinkAnchoredPaths(c *gin.Context) {

	// source slice
	var linkanchoredpathDBs []orm.LinkAnchoredPathDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredPaths", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	_, err := db.Find(&linkanchoredpathDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linkanchoredpathAPIs := make([]orm.LinkAnchoredPathAPI, 0)

	// for each linkanchoredpath, update fields from the database nullable fields
	for idx := range linkanchoredpathDBs {
		linkanchoredpathDB := &linkanchoredpathDBs[idx]
		_ = linkanchoredpathDB
		var linkanchoredpathAPI orm.LinkAnchoredPathAPI

		// insertion point for updating fields
		linkanchoredpathAPI.ID = linkanchoredpathDB.ID
		linkanchoredpathDB.CopyBasicFieldsToLinkAnchoredPath_WOP(&linkanchoredpathAPI.LinkAnchoredPath_WOP)
		linkanchoredpathAPI.LinkAnchoredPathPointersEncoding = linkanchoredpathDB.LinkAnchoredPathPointersEncoding
		linkanchoredpathAPIs = append(linkanchoredpathAPIs, linkanchoredpathAPI)
	}

	c.JSON(http.StatusOK, linkanchoredpathAPIs)
}

// PostLinkAnchoredPath
//
// swagger:route POST /linkanchoredpaths linkanchoredpaths postLinkAnchoredPath
//
// Creates a linkanchoredpath
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLinkAnchoredPath(c *gin.Context) {

	mutexLinkAnchoredPath.Lock()
	defer mutexLinkAnchoredPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLinkAnchoredPaths", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	// Validate input
	var input orm.LinkAnchoredPathAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create linkanchoredpath
	linkanchoredpathDB := orm.LinkAnchoredPathDB{}
	linkanchoredpathDB.LinkAnchoredPathPointersEncoding = input.LinkAnchoredPathPointersEncoding
	linkanchoredpathDB.CopyBasicFieldsFromLinkAnchoredPath_WOP(&input.LinkAnchoredPath_WOP)

	_, err = db.Create(&linkanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLinkAnchoredPath.CheckoutPhaseOneInstance(&linkanchoredpathDB)
	linkanchoredpath := backRepo.BackRepoLinkAnchoredPath.Map_LinkAnchoredPathDBID_LinkAnchoredPathPtr[linkanchoredpathDB.ID]

	if linkanchoredpath != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), linkanchoredpath)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, linkanchoredpathDB)
}

// GetLinkAnchoredPath
//
// swagger:route GET /linkanchoredpaths/{ID} linkanchoredpaths getLinkAnchoredPath
//
// Gets the details for a linkanchoredpath.
//
// Responses:
// default: genericError
//
//	200: linkanchoredpathDBResponse
func (controller *Controller) GetLinkAnchoredPath(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLinkAnchoredPath", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	// Get linkanchoredpathDB in DB
	var linkanchoredpathDB orm.LinkAnchoredPathDB
	if _, err := db.First(&linkanchoredpathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var linkanchoredpathAPI orm.LinkAnchoredPathAPI
	linkanchoredpathAPI.ID = linkanchoredpathDB.ID
	linkanchoredpathAPI.LinkAnchoredPathPointersEncoding = linkanchoredpathDB.LinkAnchoredPathPointersEncoding
	linkanchoredpathDB.CopyBasicFieldsToLinkAnchoredPath_WOP(&linkanchoredpathAPI.LinkAnchoredPath_WOP)

	c.JSON(http.StatusOK, linkanchoredpathAPI)
}

// UpdateLinkAnchoredPath
//
// swagger:route PATCH /linkanchoredpaths/{ID} linkanchoredpaths updateLinkAnchoredPath
//
// # Update a linkanchoredpath
//
// Responses:
// default: genericError
//
//	200: linkanchoredpathDBResponse
func (controller *Controller) UpdateLinkAnchoredPath(c *gin.Context) {

	mutexLinkAnchoredPath.Lock()
	defer mutexLinkAnchoredPath.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	// Validate input
	var input orm.LinkAnchoredPathAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linkanchoredpathDB orm.LinkAnchoredPathDB

	// fetch the linkanchoredpath
	_, err := db.First(&linkanchoredpathDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linkanchoredpathDB.CopyBasicFieldsFromLinkAnchoredPath_WOP(&input.LinkAnchoredPath_WOP)
	linkanchoredpathDB.LinkAnchoredPathPointersEncoding = input.LinkAnchoredPathPointersEncoding

	db, _ = db.Model(&linkanchoredpathDB)
	_, err = db.Updates(&linkanchoredpathDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredpathNew := new(models.LinkAnchoredPath)
	linkanchoredpathDB.CopyBasicFieldsToLinkAnchoredPath(linkanchoredpathNew)

	// redeem pointers
	linkanchoredpathDB.DecodePointers(backRepo, linkanchoredpathNew)

	// get stage instance from DB instance, and call callback function
	linkanchoredpathOld := backRepo.BackRepoLinkAnchoredPath.Map_LinkAnchoredPathDBID_LinkAnchoredPathPtr[linkanchoredpathDB.ID]
	if linkanchoredpathOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), linkanchoredpathOld, linkanchoredpathNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkanchoredpathDB
	c.JSON(http.StatusOK, linkanchoredpathDB)
}

// DeleteLinkAnchoredPath
//
// swagger:route DELETE /linkanchoredpaths/{ID} linkanchoredpaths deleteLinkAnchoredPath
//
// # Delete a linkanchoredpath
//
// default: genericError
//
//	200: linkanchoredpathDBResponse
func (controller *Controller) DeleteLinkAnchoredPath(c *gin.Context) {

	mutexLinkAnchoredPath.Lock()
	defer mutexLinkAnchoredPath.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLinkAnchoredPath", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoLinkAnchoredPath.GetDB()

	// Get model if exist
	var linkanchoredpathDB orm.LinkAnchoredPathDB
	if _, err := db.First(&linkanchoredpathDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&linkanchoredpathDB)

	// get an instance (not staged) from DB instance, and call callback function
	linkanchoredpathDeleted := new(models.LinkAnchoredPath)
	linkanchoredpathDB.CopyBasicFieldsToLinkAnchoredPath(linkanchoredpathDeleted)

	// get stage instance from DB instance, and call callback function
	linkanchoredpathStaged := backRepo.BackRepoLinkAnchoredPath.Map_LinkAnchoredPathDBID_LinkAnchoredPathPtr[linkanchoredpathDB.ID]
	if linkanchoredpathStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), linkanchoredpathStaged, linkanchoredpathDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
