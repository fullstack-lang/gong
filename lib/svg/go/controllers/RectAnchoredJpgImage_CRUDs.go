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
var __RectAnchoredJpgImage__dummysDeclaration__ models.RectAnchoredJpgImage
var _ = __RectAnchoredJpgImage__dummysDeclaration__
var __RectAnchoredJpgImage_time__dummyDeclaration time.Duration
var _ = __RectAnchoredJpgImage_time__dummyDeclaration

var mutexRectAnchoredJpgImage sync.Mutex

// An RectAnchoredJpgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredJpgImage updateRectAnchoredJpgImage deleteRectAnchoredJpgImage
type RectAnchoredJpgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredJpgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredJpgImage updateRectAnchoredJpgImage
type RectAnchoredJpgImageInput struct {
	// The RectAnchoredJpgImage to submit or modify
	// in: body
	RectAnchoredJpgImage *orm.RectAnchoredJpgImageAPI
}

// GetRectAnchoredJpgImages
//
// swagger:route GET /rectanchoredjpgimages rectanchoredjpgimages getRectAnchoredJpgImages
//
// # Get all rectanchoredjpgimages
//
// Responses:
// default: genericError
//
//	200: rectanchoredjpgimageDBResponse
func (controller *Controller) GetRectAnchoredJpgImages(c *gin.Context) {

	// source slice
	var rectanchoredjpgimageDBs []orm.RectAnchoredJpgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredJpgImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredJpgImage.GetDB()

	_, err := db.Find(&rectanchoredjpgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredjpgimageAPIs := make([]orm.RectAnchoredJpgImageAPI, 0)

	// for each rectanchoredjpgimage, update fields from the database nullable fields
	for idx := range rectanchoredjpgimageDBs {
		rectanchoredjpgimageDB := &rectanchoredjpgimageDBs[idx]
		_ = rectanchoredjpgimageDB
		var rectanchoredjpgimageAPI orm.RectAnchoredJpgImageAPI

		// insertion point for updating fields
		rectanchoredjpgimageAPI.ID = rectanchoredjpgimageDB.ID
		rectanchoredjpgimageDB.CopyBasicFieldsToRectAnchoredJpgImage_WOP(&rectanchoredjpgimageAPI.RectAnchoredJpgImage_WOP)
		rectanchoredjpgimageAPI.RectAnchoredJpgImagePointersEncoding = rectanchoredjpgimageDB.RectAnchoredJpgImagePointersEncoding
		rectanchoredjpgimageAPIs = append(rectanchoredjpgimageAPIs, rectanchoredjpgimageAPI)
	}

	c.JSON(http.StatusOK, rectanchoredjpgimageAPIs)
}

// PostRectAnchoredJpgImage
//
// swagger:route POST /rectanchoredjpgimages rectanchoredjpgimages postRectAnchoredJpgImage
//
// Creates a rectanchoredjpgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredJpgImage(c *gin.Context) {

	mutexRectAnchoredJpgImage.Lock()
	defer mutexRectAnchoredJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredJpgImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredJpgImage.GetDB()

	// Validate input
	var input orm.RectAnchoredJpgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredjpgimage
	rectanchoredjpgimageDB := orm.RectAnchoredJpgImageDB{}
	rectanchoredjpgimageDB.RectAnchoredJpgImagePointersEncoding = input.RectAnchoredJpgImagePointersEncoding
	rectanchoredjpgimageDB.CopyBasicFieldsFromRectAnchoredJpgImage_WOP(&input.RectAnchoredJpgImage_WOP)

	_, err = db.Create(&rectanchoredjpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredJpgImage.CheckoutPhaseOneInstance(&rectanchoredjpgimageDB)
	rectanchoredjpgimage := backRepo.BackRepoRectAnchoredJpgImage.Map_RectAnchoredJpgImageDBID_RectAnchoredJpgImagePtr[rectanchoredjpgimageDB.ID]

	if rectanchoredjpgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredjpgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredjpgimageDB)
}

// GetRectAnchoredJpgImage
//
// swagger:route GET /rectanchoredjpgimages/{ID} rectanchoredjpgimages getRectAnchoredJpgImage
//
// Gets the details for a rectanchoredjpgimage.
//
// Responses:
// default: genericError
//
//	200: rectanchoredjpgimageDBResponse
func (controller *Controller) GetRectAnchoredJpgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredJpgImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredJpgImage.GetDB()

	// Get rectanchoredjpgimageDB in DB
	var rectanchoredjpgimageDB orm.RectAnchoredJpgImageDB
	if _, err := db.First(&rectanchoredjpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredjpgimageAPI orm.RectAnchoredJpgImageAPI
	rectanchoredjpgimageAPI.ID = rectanchoredjpgimageDB.ID
	rectanchoredjpgimageAPI.RectAnchoredJpgImagePointersEncoding = rectanchoredjpgimageDB.RectAnchoredJpgImagePointersEncoding
	rectanchoredjpgimageDB.CopyBasicFieldsToRectAnchoredJpgImage_WOP(&rectanchoredjpgimageAPI.RectAnchoredJpgImage_WOP)

	c.JSON(http.StatusOK, rectanchoredjpgimageAPI)
}

// UpdateRectAnchoredJpgImage
//
// swagger:route PATCH /rectanchoredjpgimages/{ID} rectanchoredjpgimages updateRectAnchoredJpgImage
//
// # Update a rectanchoredjpgimage
//
// Responses:
// default: genericError
//
//	200: rectanchoredjpgimageDBResponse
func (controller *Controller) UpdateRectAnchoredJpgImage(c *gin.Context) {

	mutexRectAnchoredJpgImage.Lock()
	defer mutexRectAnchoredJpgImage.Unlock()

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
	db := backRepo.BackRepoRectAnchoredJpgImage.GetDB()

	// Validate input
	var input orm.RectAnchoredJpgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredjpgimageDB orm.RectAnchoredJpgImageDB

	// fetch the rectanchoredjpgimage
	_, err := db.First(&rectanchoredjpgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredjpgimageDB.CopyBasicFieldsFromRectAnchoredJpgImage_WOP(&input.RectAnchoredJpgImage_WOP)
	rectanchoredjpgimageDB.RectAnchoredJpgImagePointersEncoding = input.RectAnchoredJpgImagePointersEncoding

	db, _ = db.Model(&rectanchoredjpgimageDB)
	_, err = db.Updates(&rectanchoredjpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredjpgimageNew := new(models.RectAnchoredJpgImage)
	rectanchoredjpgimageDB.CopyBasicFieldsToRectAnchoredJpgImage(rectanchoredjpgimageNew)

	// redeem pointers
	rectanchoredjpgimageDB.DecodePointers(backRepo, rectanchoredjpgimageNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredjpgimageOld := backRepo.BackRepoRectAnchoredJpgImage.Map_RectAnchoredJpgImageDBID_RectAnchoredJpgImagePtr[rectanchoredjpgimageDB.ID]
	if rectanchoredjpgimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), rectanchoredjpgimageOld, rectanchoredjpgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredjpgimageDB
	c.JSON(http.StatusOK, rectanchoredjpgimageDB)
}

// DeleteRectAnchoredJpgImage
//
// swagger:route DELETE /rectanchoredjpgimages/{ID} rectanchoredjpgimages deleteRectAnchoredJpgImage
//
// # Delete a rectanchoredjpgimage
//
// default: genericError
//
//	200: rectanchoredjpgimageDBResponse
func (controller *Controller) DeleteRectAnchoredJpgImage(c *gin.Context) {

	mutexRectAnchoredJpgImage.Lock()
	defer mutexRectAnchoredJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredJpgImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredJpgImage.GetDB()

	// Get model if exist
	var rectanchoredjpgimageDB orm.RectAnchoredJpgImageDB
	if _, err := db.First(&rectanchoredjpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredjpgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredjpgimageDeleted := new(models.RectAnchoredJpgImage)
	rectanchoredjpgimageDB.CopyBasicFieldsToRectAnchoredJpgImage(rectanchoredjpgimageDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredjpgimageStaged := backRepo.BackRepoRectAnchoredJpgImage.Map_RectAnchoredJpgImageDBID_RectAnchoredJpgImagePtr[rectanchoredjpgimageDB.ID]
	if rectanchoredjpgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredjpgimageStaged, rectanchoredjpgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
