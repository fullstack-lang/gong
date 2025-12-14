// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
	"github.com/fullstack-lang/gong/lib/markdown/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __JpgImage__dummysDeclaration__ models.JpgImage
var _ = __JpgImage__dummysDeclaration__
var __JpgImage_time__dummyDeclaration time.Duration
var _ = __JpgImage_time__dummyDeclaration

var mutexJpgImage sync.Mutex

// An JpgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getJpgImage updateJpgImage deleteJpgImage
type JpgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// JpgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postJpgImage updateJpgImage
type JpgImageInput struct {
	// The JpgImage to submit or modify
	// in: body
	JpgImage *orm.JpgImageAPI
}

// GetJpgImages
//
// swagger:route GET /jpgimages jpgimages getJpgImages
//
// # Get all jpgimages
//
// Responses:
// default: genericError
//
//	200: jpgimageDBResponse
func (controller *Controller) GetJpgImages(c *gin.Context) {

	// source slice
	var jpgimageDBs []orm.JpgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetJpgImages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoJpgImage.GetDB()

	_, err := db.Find(&jpgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	jpgimageAPIs := make([]orm.JpgImageAPI, 0)

	// for each jpgimage, update fields from the database nullable fields
	for idx := range jpgimageDBs {
		jpgimageDB := &jpgimageDBs[idx]
		_ = jpgimageDB
		var jpgimageAPI orm.JpgImageAPI

		// insertion point for updating fields
		jpgimageAPI.ID = jpgimageDB.ID
		jpgimageDB.CopyBasicFieldsToJpgImage_WOP(&jpgimageAPI.JpgImage_WOP)
		jpgimageAPI.JpgImagePointersEncoding = jpgimageDB.JpgImagePointersEncoding
		jpgimageAPIs = append(jpgimageAPIs, jpgimageAPI)
	}

	c.JSON(http.StatusOK, jpgimageAPIs)
}

// PostJpgImage
//
// swagger:route POST /jpgimages jpgimages postJpgImage
//
// Creates a jpgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostJpgImage(c *gin.Context) {

	mutexJpgImage.Lock()
	defer mutexJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostJpgImages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoJpgImage.GetDB()

	// Validate input
	var input orm.JpgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create jpgimage
	jpgimageDB := orm.JpgImageDB{}
	jpgimageDB.JpgImagePointersEncoding = input.JpgImagePointersEncoding
	jpgimageDB.CopyBasicFieldsFromJpgImage_WOP(&input.JpgImage_WOP)

	_, err = db.Create(&jpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoJpgImage.CheckoutPhaseOneInstance(&jpgimageDB)
	jpgimage := backRepo.BackRepoJpgImage.Map_JpgImageDBID_JpgImagePtr[jpgimageDB.ID]

	if jpgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), jpgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, jpgimageDB)
}

// GetJpgImage
//
// swagger:route GET /jpgimages/{ID} jpgimages getJpgImage
//
// Gets the details for a jpgimage.
//
// Responses:
// default: genericError
//
//	200: jpgimageDBResponse
func (controller *Controller) GetJpgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetJpgImage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoJpgImage.GetDB()

	// Get jpgimageDB in DB
	var jpgimageDB orm.JpgImageDB
	if _, err := db.First(&jpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var jpgimageAPI orm.JpgImageAPI
	jpgimageAPI.ID = jpgimageDB.ID
	jpgimageAPI.JpgImagePointersEncoding = jpgimageDB.JpgImagePointersEncoding
	jpgimageDB.CopyBasicFieldsToJpgImage_WOP(&jpgimageAPI.JpgImage_WOP)

	c.JSON(http.StatusOK, jpgimageAPI)
}

// UpdateJpgImage
//
// swagger:route PATCH /jpgimages/{ID} jpgimages updateJpgImage
//
// # Update a jpgimage
//
// Responses:
// default: genericError
//
//	200: jpgimageDBResponse
func (controller *Controller) UpdateJpgImage(c *gin.Context) {

	mutexJpgImage.Lock()
	defer mutexJpgImage.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoJpgImage.GetDB()

	// Validate input
	var input orm.JpgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var jpgimageDB orm.JpgImageDB

	// fetch the jpgimage
	_, err := db.First(&jpgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	jpgimageDB.CopyBasicFieldsFromJpgImage_WOP(&input.JpgImage_WOP)
	jpgimageDB.JpgImagePointersEncoding = input.JpgImagePointersEncoding

	db, _ = db.Model(&jpgimageDB)
	_, err = db.Updates(&jpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	jpgimageNew := new(models.JpgImage)
	jpgimageDB.CopyBasicFieldsToJpgImage(jpgimageNew)

	// redeem pointers
	jpgimageDB.DecodePointers(backRepo, jpgimageNew)

	// get stage instance from DB instance, and call callback function
	jpgimageOld := backRepo.BackRepoJpgImage.Map_JpgImageDBID_JpgImagePtr[jpgimageDB.ID]
	if jpgimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), jpgimageOld, jpgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the jpgimageDB
	c.JSON(http.StatusOK, jpgimageDB)
}

// DeleteJpgImage
//
// swagger:route DELETE /jpgimages/{ID} jpgimages deleteJpgImage
//
// # Delete a jpgimage
//
// default: genericError
//
//	200: jpgimageDBResponse
func (controller *Controller) DeleteJpgImage(c *gin.Context) {

	mutexJpgImage.Lock()
	defer mutexJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteJpgImage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/markdown/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoJpgImage.GetDB()

	// Get model if exist
	var jpgimageDB orm.JpgImageDB
	if _, err := db.First(&jpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&jpgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	jpgimageDeleted := new(models.JpgImage)
	jpgimageDB.CopyBasicFieldsToJpgImage(jpgimageDeleted)

	// get stage instance from DB instance, and call callback function
	jpgimageStaged := backRepo.BackRepoJpgImage.Map_JpgImageDBID_JpgImagePtr[jpgimageDB.ID]
	if jpgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), jpgimageStaged, jpgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
