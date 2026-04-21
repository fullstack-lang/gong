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
var __RectAnchoredSvgImage__dummysDeclaration__ models.RectAnchoredSvgImage
var _ = __RectAnchoredSvgImage__dummysDeclaration__
var __RectAnchoredSvgImage_time__dummyDeclaration time.Duration
var _ = __RectAnchoredSvgImage_time__dummyDeclaration

var mutexRectAnchoredSvgImage sync.Mutex

// An RectAnchoredSvgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredSvgImage updateRectAnchoredSvgImage deleteRectAnchoredSvgImage
type RectAnchoredSvgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredSvgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredSvgImage updateRectAnchoredSvgImage
type RectAnchoredSvgImageInput struct {
	// The RectAnchoredSvgImage to submit or modify
	// in: body
	RectAnchoredSvgImage *orm.RectAnchoredSvgImageAPI
}

// GetRectAnchoredSvgImages
//
// swagger:route GET /rectanchoredsvgimages rectanchoredsvgimages getRectAnchoredSvgImages
//
// # Get all rectanchoredsvgimages
//
// Responses:
// default: genericError
//
//	200: rectanchoredsvgimageDBResponse
func (controller *Controller) GetRectAnchoredSvgImages(c *gin.Context) {

	// source slice
	var rectanchoredsvgimageDBs []orm.RectAnchoredSvgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredSvgImage.GetDB()

	_, err := db.Find(&rectanchoredsvgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredsvgimageAPIs := make([]orm.RectAnchoredSvgImageAPI, 0)

	// for each rectanchoredsvgimage, update fields from the database nullable fields
	for idx := range rectanchoredsvgimageDBs {
		rectanchoredsvgimageDB := &rectanchoredsvgimageDBs[idx]
		_ = rectanchoredsvgimageDB
		var rectanchoredsvgimageAPI orm.RectAnchoredSvgImageAPI

		// insertion point for updating fields
		rectanchoredsvgimageAPI.ID = rectanchoredsvgimageDB.ID
		rectanchoredsvgimageDB.CopyBasicFieldsToRectAnchoredSvgImage_WOP(&rectanchoredsvgimageAPI.RectAnchoredSvgImage_WOP)
		rectanchoredsvgimageAPI.RectAnchoredSvgImagePointersEncoding = rectanchoredsvgimageDB.RectAnchoredSvgImagePointersEncoding
		rectanchoredsvgimageAPIs = append(rectanchoredsvgimageAPIs, rectanchoredsvgimageAPI)
	}

	c.JSON(http.StatusOK, rectanchoredsvgimageAPIs)
}

// PostRectAnchoredSvgImage
//
// swagger:route POST /rectanchoredsvgimages rectanchoredsvgimages postRectAnchoredSvgImage
//
// Creates a rectanchoredsvgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredSvgImage(c *gin.Context) {

	mutexRectAnchoredSvgImage.Lock()
	defer mutexRectAnchoredSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredSvgImage.GetDB()

	// Validate input
	var input orm.RectAnchoredSvgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredsvgimage
	rectanchoredsvgimageDB := orm.RectAnchoredSvgImageDB{}
	rectanchoredsvgimageDB.RectAnchoredSvgImagePointersEncoding = input.RectAnchoredSvgImagePointersEncoding
	rectanchoredsvgimageDB.CopyBasicFieldsFromRectAnchoredSvgImage_WOP(&input.RectAnchoredSvgImage_WOP)

	_, err = db.Create(&rectanchoredsvgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredSvgImage.CheckoutPhaseOneInstance(&rectanchoredsvgimageDB)
	rectanchoredsvgimage := backRepo.BackRepoRectAnchoredSvgImage.Map_RectAnchoredSvgImageDBID_RectAnchoredSvgImagePtr[rectanchoredsvgimageDB.ID]

	if rectanchoredsvgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredsvgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredsvgimageDB)
}

// GetRectAnchoredSvgImage
//
// swagger:route GET /rectanchoredsvgimages/{ID} rectanchoredsvgimages getRectAnchoredSvgImage
//
// Gets the details for a rectanchoredsvgimage.
//
// Responses:
// default: genericError
//
//	200: rectanchoredsvgimageDBResponse
func (controller *Controller) GetRectAnchoredSvgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredSvgImage.GetDB()

	// Get rectanchoredsvgimageDB in DB
	var rectanchoredsvgimageDB orm.RectAnchoredSvgImageDB
	if _, err := db.First(&rectanchoredsvgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredsvgimageAPI orm.RectAnchoredSvgImageAPI
	rectanchoredsvgimageAPI.ID = rectanchoredsvgimageDB.ID
	rectanchoredsvgimageAPI.RectAnchoredSvgImagePointersEncoding = rectanchoredsvgimageDB.RectAnchoredSvgImagePointersEncoding
	rectanchoredsvgimageDB.CopyBasicFieldsToRectAnchoredSvgImage_WOP(&rectanchoredsvgimageAPI.RectAnchoredSvgImage_WOP)

	c.JSON(http.StatusOK, rectanchoredsvgimageAPI)
}

// UpdateRectAnchoredSvgImage
//
// swagger:route PATCH /rectanchoredsvgimages/{ID} rectanchoredsvgimages updateRectAnchoredSvgImage
//
// # Update a rectanchoredsvgimage
//
// Responses:
// default: genericError
//
//	200: rectanchoredsvgimageDBResponse
func (controller *Controller) UpdateRectAnchoredSvgImage(c *gin.Context) {

	mutexRectAnchoredSvgImage.Lock()
	defer mutexRectAnchoredSvgImage.Unlock()

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
	db := backRepo.BackRepoRectAnchoredSvgImage.GetDB()

	// Validate input
	var input orm.RectAnchoredSvgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredsvgimageDB orm.RectAnchoredSvgImageDB

	// fetch the rectanchoredsvgimage
	_, err := db.First(&rectanchoredsvgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredsvgimageDB.CopyBasicFieldsFromRectAnchoredSvgImage_WOP(&input.RectAnchoredSvgImage_WOP)
	rectanchoredsvgimageDB.RectAnchoredSvgImagePointersEncoding = input.RectAnchoredSvgImagePointersEncoding

	db, _ = db.Model(&rectanchoredsvgimageDB)
	_, err = db.Updates(&rectanchoredsvgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredsvgimageNew := new(models.RectAnchoredSvgImage)
	rectanchoredsvgimageDB.CopyBasicFieldsToRectAnchoredSvgImage(rectanchoredsvgimageNew)

	// redeem pointers
	rectanchoredsvgimageDB.DecodePointers(backRepo, rectanchoredsvgimageNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredsvgimageOld := backRepo.BackRepoRectAnchoredSvgImage.Map_RectAnchoredSvgImageDBID_RectAnchoredSvgImagePtr[rectanchoredsvgimageDB.ID]
	if rectanchoredsvgimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), rectanchoredsvgimageOld, rectanchoredsvgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredsvgimageDB
	c.JSON(http.StatusOK, rectanchoredsvgimageDB)
}

// DeleteRectAnchoredSvgImage
//
// swagger:route DELETE /rectanchoredsvgimages/{ID} rectanchoredsvgimages deleteRectAnchoredSvgImage
//
// # Delete a rectanchoredsvgimage
//
// default: genericError
//
//	200: rectanchoredsvgimageDBResponse
func (controller *Controller) DeleteRectAnchoredSvgImage(c *gin.Context) {

	mutexRectAnchoredSvgImage.Lock()
	defer mutexRectAnchoredSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredSvgImage.GetDB()

	// Get model if exist
	var rectanchoredsvgimageDB orm.RectAnchoredSvgImageDB
	if _, err := db.First(&rectanchoredsvgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredsvgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredsvgimageDeleted := new(models.RectAnchoredSvgImage)
	rectanchoredsvgimageDB.CopyBasicFieldsToRectAnchoredSvgImage(rectanchoredsvgimageDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredsvgimageStaged := backRepo.BackRepoRectAnchoredSvgImage.Map_RectAnchoredSvgImageDBID_RectAnchoredSvgImagePtr[rectanchoredsvgimageDB.ID]
	if rectanchoredsvgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredsvgimageStaged, rectanchoredsvgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
