// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Svg__dummysDeclaration__ models.Svg
var __Svg_time__dummyDeclaration time.Duration

var mutexSvg sync.Mutex

// An SvgID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSvg updateSvg deleteSvg
type SvgID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SvgInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSvg updateSvg
type SvgInput struct {
	// The Svg to submit or modify
	// in: body
	Svg *orm.SvgAPI
}

// GetSvgs
//
// swagger:route GET /svgs svgs getSvgs
//
// # Get all svgs
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) GetSvgs(c *gin.Context) {

	// source slice
	var svgDBs []orm.SvgDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvgs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSvg.GetDB()

	_, err := db.Find(&svgDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	svgAPIs := make([]orm.SvgAPI, 0)

	// for each svg, update fields from the database nullable fields
	for idx := range svgDBs {
		svgDB := &svgDBs[idx]
		_ = svgDB
		var svgAPI orm.SvgAPI

		// insertion point for updating fields
		svgAPI.ID = svgDB.ID
		svgDB.CopyBasicFieldsToSvg_WOP(&svgAPI.Svg_WOP)
		svgAPI.SvgPointersEncoding = svgDB.SvgPointersEncoding
		svgAPIs = append(svgAPIs, svgAPI)
	}

	c.JSON(http.StatusOK, svgAPIs)
}

// PostSvg
//
// swagger:route POST /svgs svgs postSvg
//
// Creates a svg
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSvg(c *gin.Context) {

	mutexSvg.Lock()
	defer mutexSvg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSvgs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSvg.GetDB()

	// Validate input
	var input orm.SvgAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create svg
	svgDB := orm.SvgDB{}
	svgDB.SvgPointersEncoding = input.SvgPointersEncoding
	svgDB.CopyBasicFieldsFromSvg_WOP(&input.Svg_WOP)

	_, err = db.Create(&svgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSvg.CheckoutPhaseOneInstance(&svgDB)
	svg := backRepo.BackRepoSvg.Map_SvgDBID_SvgPtr[svgDB.ID]

	if svg != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), svg)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, svgDB)
}

// GetSvg
//
// swagger:route GET /svgs/{ID} svgs getSvg
//
// Gets the details for a svg.
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) GetSvg(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSvg", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSvg.GetDB()

	// Get svgDB in DB
	var svgDB orm.SvgDB
	if _, err := db.First(&svgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var svgAPI orm.SvgAPI
	svgAPI.ID = svgDB.ID
	svgAPI.SvgPointersEncoding = svgDB.SvgPointersEncoding
	svgDB.CopyBasicFieldsToSvg_WOP(&svgAPI.Svg_WOP)

	c.JSON(http.StatusOK, svgAPI)
}

// UpdateSvg
//
// swagger:route PATCH /svgs/{ID} svgs updateSvg
//
// # Update a svg
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) UpdateSvg(c *gin.Context) {

	mutexSvg.Lock()
	defer mutexSvg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSvg", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSvg.GetDB()

	// Validate input
	var input orm.SvgAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgDB orm.SvgDB

	// fetch the svg
	_, err := db.First(&svgDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgDB.CopyBasicFieldsFromSvg_WOP(&input.Svg_WOP)
	svgDB.SvgPointersEncoding = input.SvgPointersEncoding

	db, _ = db.Model(&svgDB)
	_, err = db.Updates(&svgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgNew := new(models.Svg)
	svgDB.CopyBasicFieldsToSvg(svgNew)

	// redeem pointers
	svgDB.DecodePointers(backRepo, svgNew)

	// get stage instance from DB instance, and call callback function
	svgOld := backRepo.BackRepoSvg.Map_SvgDBID_SvgPtr[svgDB.ID]
	if svgOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), svgOld, svgNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgDB
	c.JSON(http.StatusOK, svgDB)
}

// DeleteSvg
//
// swagger:route DELETE /svgs/{ID} svgs deleteSvg
//
// # Delete a svg
//
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) DeleteSvg(c *gin.Context) {

	mutexSvg.Lock()
	defer mutexSvg.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSvg", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSvg.GetDB()

	// Get model if exist
	var svgDB orm.SvgDB
	if _, err := db.First(&svgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&svgDB)

	// get an instance (not staged) from DB instance, and call callback function
	svgDeleted := new(models.Svg)
	svgDB.CopyBasicFieldsToSvg(svgDeleted)

	// get stage instance from DB instance, and call callback function
	svgStaged := backRepo.BackRepoSvg.Map_SvgDBID_SvgPtr[svgDB.ID]
	if svgStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), svgStaged, svgDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
