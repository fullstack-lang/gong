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
var __PngImage__dummysDeclaration__ models.PngImage
var __PngImage_time__dummyDeclaration time.Duration

var mutexPngImage sync.Mutex

// An PngImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPngImage updatePngImage deletePngImage
type PngImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PngImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPngImage updatePngImage
type PngImageInput struct {
	// The PngImage to submit or modify
	// in: body
	PngImage *orm.PngImageAPI
}

// GetPngImages
//
// swagger:route GET /pngimages pngimages getPngImages
//
// # Get all pngimages
//
// Responses:
// default: genericError
//
//	200: pngimageDBResponse
func (controller *Controller) GetPngImages(c *gin.Context) {

	// source slice
	var pngimageDBs []orm.PngImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPngImages", "Name", stackPath)
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
	db := backRepo.BackRepoPngImage.GetDB()

	_, err := db.Find(&pngimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pngimageAPIs := make([]orm.PngImageAPI, 0)

	// for each pngimage, update fields from the database nullable fields
	for idx := range pngimageDBs {
		pngimageDB := &pngimageDBs[idx]
		_ = pngimageDB
		var pngimageAPI orm.PngImageAPI

		// insertion point for updating fields
		pngimageAPI.ID = pngimageDB.ID
		pngimageDB.CopyBasicFieldsToPngImage_WOP(&pngimageAPI.PngImage_WOP)
		pngimageAPI.PngImagePointersEncoding = pngimageDB.PngImagePointersEncoding
		pngimageAPIs = append(pngimageAPIs, pngimageAPI)
	}

	c.JSON(http.StatusOK, pngimageAPIs)
}

// PostPngImage
//
// swagger:route POST /pngimages pngimages postPngImage
//
// Creates a pngimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPngImage(c *gin.Context) {

	mutexPngImage.Lock()
	defer mutexPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPngImages", "Name", stackPath)
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
	db := backRepo.BackRepoPngImage.GetDB()

	// Validate input
	var input orm.PngImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pngimage
	pngimageDB := orm.PngImageDB{}
	pngimageDB.PngImagePointersEncoding = input.PngImagePointersEncoding
	pngimageDB.CopyBasicFieldsFromPngImage_WOP(&input.PngImage_WOP)

	_, err = db.Create(&pngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPngImage.CheckoutPhaseOneInstance(&pngimageDB)
	pngimage := backRepo.BackRepoPngImage.Map_PngImageDBID_PngImagePtr[pngimageDB.ID]

	if pngimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pngimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pngimageDB)
}

// GetPngImage
//
// swagger:route GET /pngimages/{ID} pngimages getPngImage
//
// Gets the details for a pngimage.
//
// Responses:
// default: genericError
//
//	200: pngimageDBResponse
func (controller *Controller) GetPngImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPngImage", "Name", stackPath)
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
	db := backRepo.BackRepoPngImage.GetDB()

	// Get pngimageDB in DB
	var pngimageDB orm.PngImageDB
	if _, err := db.First(&pngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pngimageAPI orm.PngImageAPI
	pngimageAPI.ID = pngimageDB.ID
	pngimageAPI.PngImagePointersEncoding = pngimageDB.PngImagePointersEncoding
	pngimageDB.CopyBasicFieldsToPngImage_WOP(&pngimageAPI.PngImage_WOP)

	c.JSON(http.StatusOK, pngimageAPI)
}

// UpdatePngImage
//
// swagger:route PATCH /pngimages/{ID} pngimages updatePngImage
//
// # Update a pngimage
//
// Responses:
// default: genericError
//
//	200: pngimageDBResponse
func (controller *Controller) UpdatePngImage(c *gin.Context) {

	mutexPngImage.Lock()
	defer mutexPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
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
	db := backRepo.BackRepoPngImage.GetDB()

	// Validate input
	var input orm.PngImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pngimageDB orm.PngImageDB

	// fetch the pngimage
	_, err := db.First(&pngimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pngimageDB.CopyBasicFieldsFromPngImage_WOP(&input.PngImage_WOP)
	pngimageDB.PngImagePointersEncoding = input.PngImagePointersEncoding

	db, _ = db.Model(&pngimageDB)
	_, err = db.Updates(&pngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pngimageNew := new(models.PngImage)
	pngimageDB.CopyBasicFieldsToPngImage(pngimageNew)

	// redeem pointers
	pngimageDB.DecodePointers(backRepo, pngimageNew)

	// get stage instance from DB instance, and call callback function
	pngimageOld := backRepo.BackRepoPngImage.Map_PngImageDBID_PngImagePtr[pngimageDB.ID]
	if pngimageOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), pngimageOld, pngimageNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), pngimageOld, pngimageNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pngimageDB
	c.JSON(http.StatusOK, pngimageDB)
}

// DeletePngImage
//
// swagger:route DELETE /pngimages/{ID} pngimages deletePngImage
//
// # Delete a pngimage
//
// default: genericError
//
//	200: pngimageDBResponse
func (controller *Controller) DeletePngImage(c *gin.Context) {

	mutexPngImage.Lock()
	defer mutexPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePngImage", "Name", stackPath)
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
	db := backRepo.BackRepoPngImage.GetDB()

	// Get model if exist
	var pngimageDB orm.PngImageDB
	if _, err := db.First(&pngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&pngimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	pngimageDeleted := new(models.PngImage)
	pngimageDB.CopyBasicFieldsToPngImage(pngimageDeleted)

	// get stage instance from DB instance, and call callback function
	pngimageStaged := backRepo.BackRepoPngImage.Map_PngImageDBID_PngImagePtr[pngimageDB.ID]
	if pngimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pngimageStaged, pngimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
