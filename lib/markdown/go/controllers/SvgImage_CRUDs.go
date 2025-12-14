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
var __SvgImage__dummysDeclaration__ models.SvgImage
var _ = __SvgImage__dummysDeclaration__
var __SvgImage_time__dummyDeclaration time.Duration
var _ = __SvgImage_time__dummyDeclaration

var mutexSvgImage sync.Mutex

// An SvgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSvgImage updateSvgImage deleteSvgImage
type SvgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SvgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSvgImage updateSvgImage
type SvgImageInput struct {
	// The SvgImage to submit or modify
	// in: body
	SvgImage *orm.SvgImageAPI
}

// GetSvgImages
//
// swagger:route GET /svgimages svgimages getSvgImages
//
// # Get all svgimages
//
// Responses:
// default: genericError
//
//	200: svgimageDBResponse
func (controller *Controller) GetSvgImages(c *gin.Context) {

	// source slice
	var svgimageDBs []orm.SvgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoSvgImage.GetDB()

	_, err := db.Find(&svgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	svgimageAPIs := make([]orm.SvgImageAPI, 0)

	// for each svgimage, update fields from the database nullable fields
	for idx := range svgimageDBs {
		svgimageDB := &svgimageDBs[idx]
		_ = svgimageDB
		var svgimageAPI orm.SvgImageAPI

		// insertion point for updating fields
		svgimageAPI.ID = svgimageDB.ID
		svgimageDB.CopyBasicFieldsToSvgImage_WOP(&svgimageAPI.SvgImage_WOP)
		svgimageAPI.SvgImagePointersEncoding = svgimageDB.SvgImagePointersEncoding
		svgimageAPIs = append(svgimageAPIs, svgimageAPI)
	}

	c.JSON(http.StatusOK, svgimageAPIs)
}

// PostSvgImage
//
// swagger:route POST /svgimages svgimages postSvgImage
//
// Creates a svgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSvgImage(c *gin.Context) {

	mutexSvgImage.Lock()
	defer mutexSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoSvgImage.GetDB()

	// Validate input
	var input orm.SvgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create svgimage
	svgimageDB := orm.SvgImageDB{}
	svgimageDB.SvgImagePointersEncoding = input.SvgImagePointersEncoding
	svgimageDB.CopyBasicFieldsFromSvgImage_WOP(&input.SvgImage_WOP)

	_, err = db.Create(&svgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSvgImage.CheckoutPhaseOneInstance(&svgimageDB)
	svgimage := backRepo.BackRepoSvgImage.Map_SvgImageDBID_SvgImagePtr[svgimageDB.ID]

	if svgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), svgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, svgimageDB)
}

// GetSvgImage
//
// swagger:route GET /svgimages/{ID} svgimages getSvgImage
//
// Gets the details for a svgimage.
//
// Responses:
// default: genericError
//
//	200: svgimageDBResponse
func (controller *Controller) GetSvgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoSvgImage.GetDB()

	// Get svgimageDB in DB
	var svgimageDB orm.SvgImageDB
	if _, err := db.First(&svgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var svgimageAPI orm.SvgImageAPI
	svgimageAPI.ID = svgimageDB.ID
	svgimageAPI.SvgImagePointersEncoding = svgimageDB.SvgImagePointersEncoding
	svgimageDB.CopyBasicFieldsToSvgImage_WOP(&svgimageAPI.SvgImage_WOP)

	c.JSON(http.StatusOK, svgimageAPI)
}

// UpdateSvgImage
//
// swagger:route PATCH /svgimages/{ID} svgimages updateSvgImage
//
// # Update a svgimage
//
// Responses:
// default: genericError
//
//	200: svgimageDBResponse
func (controller *Controller) UpdateSvgImage(c *gin.Context) {

	mutexSvgImage.Lock()
	defer mutexSvgImage.Unlock()

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
	db := backRepo.BackRepoSvgImage.GetDB()

	// Validate input
	var input orm.SvgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgimageDB orm.SvgImageDB

	// fetch the svgimage
	_, err := db.First(&svgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgimageDB.CopyBasicFieldsFromSvgImage_WOP(&input.SvgImage_WOP)
	svgimageDB.SvgImagePointersEncoding = input.SvgImagePointersEncoding

	db, _ = db.Model(&svgimageDB)
	_, err = db.Updates(&svgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgimageNew := new(models.SvgImage)
	svgimageDB.CopyBasicFieldsToSvgImage(svgimageNew)

	// redeem pointers
	svgimageDB.DecodePointers(backRepo, svgimageNew)

	// get stage instance from DB instance, and call callback function
	svgimageOld := backRepo.BackRepoSvgImage.Map_SvgImageDBID_SvgImagePtr[svgimageDB.ID]
	if svgimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), svgimageOld, svgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgimageDB
	c.JSON(http.StatusOK, svgimageDB)
}

// DeleteSvgImage
//
// swagger:route DELETE /svgimages/{ID} svgimages deleteSvgImage
//
// # Delete a svgimage
//
// default: genericError
//
//	200: svgimageDBResponse
func (controller *Controller) DeleteSvgImage(c *gin.Context) {

	mutexSvgImage.Lock()
	defer mutexSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoSvgImage.GetDB()

	// Get model if exist
	var svgimageDB orm.SvgImageDB
	if _, err := db.First(&svgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&svgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	svgimageDeleted := new(models.SvgImage)
	svgimageDB.CopyBasicFieldsToSvgImage(svgimageDeleted)

	// get stage instance from DB instance, and call callback function
	svgimageStaged := backRepo.BackRepoSvgImage.Map_SvgImageDBID_SvgImagePtr[svgimageDB.ID]
	if svgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), svgimageStaged, svgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
