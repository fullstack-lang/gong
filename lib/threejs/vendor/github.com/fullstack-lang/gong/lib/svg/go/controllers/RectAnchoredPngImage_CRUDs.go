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
var __RectAnchoredPngImage__dummysDeclaration__ models.RectAnchoredPngImage
var _ = __RectAnchoredPngImage__dummysDeclaration__
var __RectAnchoredPngImage_time__dummyDeclaration time.Duration
var _ = __RectAnchoredPngImage_time__dummyDeclaration

var mutexRectAnchoredPngImage sync.Mutex

// An RectAnchoredPngImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredPngImage updateRectAnchoredPngImage deleteRectAnchoredPngImage
type RectAnchoredPngImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredPngImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredPngImage updateRectAnchoredPngImage
type RectAnchoredPngImageInput struct {
	// The RectAnchoredPngImage to submit or modify
	// in: body
	RectAnchoredPngImage *orm.RectAnchoredPngImageAPI
}

// GetRectAnchoredPngImages
//
// swagger:route GET /rectanchoredpngimages rectanchoredpngimages getRectAnchoredPngImages
//
// # Get all rectanchoredpngimages
//
// Responses:
// default: genericError
//
//	200: rectanchoredpngimageDBResponse
func (controller *Controller) GetRectAnchoredPngImages(c *gin.Context) {

	// source slice
	var rectanchoredpngimageDBs []orm.RectAnchoredPngImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredPngImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	_, err := db.Find(&rectanchoredpngimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredpngimageAPIs := make([]orm.RectAnchoredPngImageAPI, 0)

	// for each rectanchoredpngimage, update fields from the database nullable fields
	for idx := range rectanchoredpngimageDBs {
		rectanchoredpngimageDB := &rectanchoredpngimageDBs[idx]
		_ = rectanchoredpngimageDB
		var rectanchoredpngimageAPI orm.RectAnchoredPngImageAPI

		// insertion point for updating fields
		rectanchoredpngimageAPI.ID = rectanchoredpngimageDB.ID
		rectanchoredpngimageDB.CopyBasicFieldsToRectAnchoredPngImage_WOP(&rectanchoredpngimageAPI.RectAnchoredPngImage_WOP)
		rectanchoredpngimageAPI.RectAnchoredPngImagePointersEncoding = rectanchoredpngimageDB.RectAnchoredPngImagePointersEncoding
		rectanchoredpngimageAPIs = append(rectanchoredpngimageAPIs, rectanchoredpngimageAPI)
	}

	c.JSON(http.StatusOK, rectanchoredpngimageAPIs)
}

// PostRectAnchoredPngImage
//
// swagger:route POST /rectanchoredpngimages rectanchoredpngimages postRectAnchoredPngImage
//
// Creates a rectanchoredpngimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredPngImage(c *gin.Context) {

	mutexRectAnchoredPngImage.Lock()
	defer mutexRectAnchoredPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredPngImages", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	// Validate input
	var input orm.RectAnchoredPngImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredpngimage
	rectanchoredpngimageDB := orm.RectAnchoredPngImageDB{}
	rectanchoredpngimageDB.RectAnchoredPngImagePointersEncoding = input.RectAnchoredPngImagePointersEncoding
	rectanchoredpngimageDB.CopyBasicFieldsFromRectAnchoredPngImage_WOP(&input.RectAnchoredPngImage_WOP)

	_, err = db.Create(&rectanchoredpngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredPngImage.CheckoutPhaseOneInstance(&rectanchoredpngimageDB)
	rectanchoredpngimage := backRepo.BackRepoRectAnchoredPngImage.Map_RectAnchoredPngImageDBID_RectAnchoredPngImagePtr[rectanchoredpngimageDB.ID]

	if rectanchoredpngimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredpngimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredpngimageDB)
}

// GetRectAnchoredPngImage
//
// swagger:route GET /rectanchoredpngimages/{ID} rectanchoredpngimages getRectAnchoredPngImage
//
// Gets the details for a rectanchoredpngimage.
//
// Responses:
// default: genericError
//
//	200: rectanchoredpngimageDBResponse
func (controller *Controller) GetRectAnchoredPngImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredPngImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	// Get rectanchoredpngimageDB in DB
	var rectanchoredpngimageDB orm.RectAnchoredPngImageDB
	if _, err := db.First(&rectanchoredpngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredpngimageAPI orm.RectAnchoredPngImageAPI
	rectanchoredpngimageAPI.ID = rectanchoredpngimageDB.ID
	rectanchoredpngimageAPI.RectAnchoredPngImagePointersEncoding = rectanchoredpngimageDB.RectAnchoredPngImagePointersEncoding
	rectanchoredpngimageDB.CopyBasicFieldsToRectAnchoredPngImage_WOP(&rectanchoredpngimageAPI.RectAnchoredPngImage_WOP)

	c.JSON(http.StatusOK, rectanchoredpngimageAPI)
}

// UpdateRectAnchoredPngImage
//
// swagger:route PATCH /rectanchoredpngimages/{ID} rectanchoredpngimages updateRectAnchoredPngImage
//
// # Update a rectanchoredpngimage
//
// Responses:
// default: genericError
//
//	200: rectanchoredpngimageDBResponse
func (controller *Controller) UpdateRectAnchoredPngImage(c *gin.Context) {

	mutexRectAnchoredPngImage.Lock()
	defer mutexRectAnchoredPngImage.Unlock()

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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	// Validate input
	var input orm.RectAnchoredPngImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredpngimageDB orm.RectAnchoredPngImageDB

	// fetch the rectanchoredpngimage
	_, err := db.First(&rectanchoredpngimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredpngimageDB.CopyBasicFieldsFromRectAnchoredPngImage_WOP(&input.RectAnchoredPngImage_WOP)
	rectanchoredpngimageDB.RectAnchoredPngImagePointersEncoding = input.RectAnchoredPngImagePointersEncoding

	db, _ = db.Model(&rectanchoredpngimageDB)
	_, err = db.Updates(&rectanchoredpngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpngimageNew := new(models.RectAnchoredPngImage)
	rectanchoredpngimageDB.CopyBasicFieldsToRectAnchoredPngImage(rectanchoredpngimageNew)

	// redeem pointers
	rectanchoredpngimageDB.DecodePointers(backRepo, rectanchoredpngimageNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredpngimageOld := backRepo.BackRepoRectAnchoredPngImage.Map_RectAnchoredPngImageDBID_RectAnchoredPngImagePtr[rectanchoredpngimageDB.ID]
	if rectanchoredpngimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), rectanchoredpngimageOld, rectanchoredpngimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredpngimageDB
	c.JSON(http.StatusOK, rectanchoredpngimageDB)
}

// DeleteRectAnchoredPngImage
//
// swagger:route DELETE /rectanchoredpngimages/{ID} rectanchoredpngimages deleteRectAnchoredPngImage
//
// # Delete a rectanchoredpngimage
//
// default: genericError
//
//	200: rectanchoredpngimageDBResponse
func (controller *Controller) DeleteRectAnchoredPngImage(c *gin.Context) {

	mutexRectAnchoredPngImage.Lock()
	defer mutexRectAnchoredPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredPngImage", "Name", stackPath)
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
	db := backRepo.BackRepoRectAnchoredPngImage.GetDB()

	// Get model if exist
	var rectanchoredpngimageDB orm.RectAnchoredPngImageDB
	if _, err := db.First(&rectanchoredpngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredpngimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredpngimageDeleted := new(models.RectAnchoredPngImage)
	rectanchoredpngimageDB.CopyBasicFieldsToRectAnchoredPngImage(rectanchoredpngimageDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredpngimageStaged := backRepo.BackRepoRectAnchoredPngImage.Map_RectAnchoredPngImageDBID_RectAnchoredPngImagePtr[rectanchoredpngimageDB.ID]
	if rectanchoredpngimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredpngimageStaged, rectanchoredpngimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
