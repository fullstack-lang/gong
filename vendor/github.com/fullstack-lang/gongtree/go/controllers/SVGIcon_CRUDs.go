// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __SVGIcon__dummysDeclaration__ models.SVGIcon
var __SVGIcon_time__dummyDeclaration time.Duration

var mutexSVGIcon sync.Mutex

// An SVGIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSVGIcon updateSVGIcon deleteSVGIcon
type SVGIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SVGIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSVGIcon updateSVGIcon
type SVGIconInput struct {
	// The SVGIcon to submit or modify
	// in: body
	SVGIcon *orm.SVGIconAPI
}

// GetSVGIcons
//
// swagger:route GET /svgicons svgicons getSVGIcons
//
// # Get all svgicons
//
// Responses:
// default: genericError
//
//	200: svgiconDBResponse
func (controller *Controller) GetSVGIcons(c *gin.Context) {

	// source slice
	var svgiconDBs []orm.SVGIconDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSVGIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	_, err := db.Find(&svgiconDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	svgiconAPIs := make([]orm.SVGIconAPI, 0)

	// for each svgicon, update fields from the database nullable fields
	for idx := range svgiconDBs {
		svgiconDB := &svgiconDBs[idx]
		_ = svgiconDB
		var svgiconAPI orm.SVGIconAPI

		// insertion point for updating fields
		svgiconAPI.ID = svgiconDB.ID
		svgiconDB.CopyBasicFieldsToSVGIcon_WOP(&svgiconAPI.SVGIcon_WOP)
		svgiconAPI.SVGIconPointersEncoding = svgiconDB.SVGIconPointersEncoding
		svgiconAPIs = append(svgiconAPIs, svgiconAPI)
	}

	c.JSON(http.StatusOK, svgiconAPIs)
}

// PostSVGIcon
//
// swagger:route POST /svgicons svgicons postSVGIcon
//
// Creates a svgicon
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSVGIcon(c *gin.Context) {

	mutexSVGIcon.Lock()
	defer mutexSVGIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSVGIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	// Validate input
	var input orm.SVGIconAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create svgicon
	svgiconDB := orm.SVGIconDB{}
	svgiconDB.SVGIconPointersEncoding = input.SVGIconPointersEncoding
	svgiconDB.CopyBasicFieldsFromSVGIcon_WOP(&input.SVGIcon_WOP)

	_, err = db.Create(&svgiconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSVGIcon.CheckoutPhaseOneInstance(&svgiconDB)
	svgicon := backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]

	if svgicon != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), svgicon)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, svgiconDB)
}

// GetSVGIcon
//
// swagger:route GET /svgicons/{ID} svgicons getSVGIcon
//
// Gets the details for a svgicon.
//
// Responses:
// default: genericError
//
//	200: svgiconDBResponse
func (controller *Controller) GetSVGIcon(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSVGIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	// Get svgiconDB in DB
	var svgiconDB orm.SVGIconDB
	if _, err := db.First(&svgiconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var svgiconAPI orm.SVGIconAPI
	svgiconAPI.ID = svgiconDB.ID
	svgiconAPI.SVGIconPointersEncoding = svgiconDB.SVGIconPointersEncoding
	svgiconDB.CopyBasicFieldsToSVGIcon_WOP(&svgiconAPI.SVGIcon_WOP)

	c.JSON(http.StatusOK, svgiconAPI)
}

// UpdateSVGIcon
//
// swagger:route PATCH /svgicons/{ID} svgicons updateSVGIcon
//
// # Update a svgicon
//
// Responses:
// default: genericError
//
//	200: svgiconDBResponse
func (controller *Controller) UpdateSVGIcon(c *gin.Context) {

	mutexSVGIcon.Lock()
	defer mutexSVGIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSVGIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	// Validate input
	var input orm.SVGIconAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgiconDB orm.SVGIconDB

	// fetch the svgicon
	_, err := db.First(&svgiconDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgiconDB.CopyBasicFieldsFromSVGIcon_WOP(&input.SVGIcon_WOP)
	svgiconDB.SVGIconPointersEncoding = input.SVGIconPointersEncoding

	db, _ = db.Model(&svgiconDB)
	_, err = db.Updates(&svgiconDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgiconNew := new(models.SVGIcon)
	svgiconDB.CopyBasicFieldsToSVGIcon(svgiconNew)

	// redeem pointers
	svgiconDB.DecodePointers(backRepo, svgiconNew)

	// get stage instance from DB instance, and call callback function
	svgiconOld := backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]
	if svgiconOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), svgiconOld, svgiconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgiconDB
	c.JSON(http.StatusOK, svgiconDB)
}

// DeleteSVGIcon
//
// swagger:route DELETE /svgicons/{ID} svgicons deleteSVGIcon
//
// # Delete a svgicon
//
// default: genericError
//
//	200: svgiconDBResponse
func (controller *Controller) DeleteSVGIcon(c *gin.Context) {

	mutexSVGIcon.Lock()
	defer mutexSVGIcon.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSVGIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSVGIcon.GetDB()

	// Get model if exist
	var svgiconDB orm.SVGIconDB
	if _, err := db.First(&svgiconDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&svgiconDB)

	// get an instance (not staged) from DB instance, and call callback function
	svgiconDeleted := new(models.SVGIcon)
	svgiconDB.CopyBasicFieldsToSVGIcon(svgiconDeleted)

	// get stage instance from DB instance, and call callback function
	svgiconStaged := backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]
	if svgiconStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), svgiconStaged, svgiconDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
