// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Slider__dummysDeclaration__ models.Slider
var __Slider_time__dummyDeclaration time.Duration

var mutexSlider sync.Mutex

// An SliderID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSlider updateSlider deleteSlider
type SliderID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SliderInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSlider updateSlider
type SliderInput struct {
	// The Slider to submit or modify
	// in: body
	Slider *orm.SliderAPI
}

// GetSliders
//
// swagger:route GET /sliders sliders getSliders
//
// # Get all sliders
//
// Responses:
// default: genericError
//
//	200: sliderDBResponse
func (controller *Controller) GetSliders(c *gin.Context) {

	// source slice
	var sliderDBs []orm.SliderDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSliders", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSlider.GetDB()

	_, err := db.Find(&sliderDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sliderAPIs := make([]orm.SliderAPI, 0)

	// for each slider, update fields from the database nullable fields
	for idx := range sliderDBs {
		sliderDB := &sliderDBs[idx]
		_ = sliderDB
		var sliderAPI orm.SliderAPI

		// insertion point for updating fields
		sliderAPI.ID = sliderDB.ID
		sliderDB.CopyBasicFieldsToSlider_WOP(&sliderAPI.Slider_WOP)
		sliderAPI.SliderPointersEncoding = sliderDB.SliderPointersEncoding
		sliderAPIs = append(sliderAPIs, sliderAPI)
	}

	c.JSON(http.StatusOK, sliderAPIs)
}

// PostSlider
//
// swagger:route POST /sliders sliders postSlider
//
// Creates a slider
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSlider(c *gin.Context) {

	mutexSlider.Lock()
	defer mutexSlider.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSliders", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSlider.GetDB()

	// Validate input
	var input orm.SliderAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create slider
	sliderDB := orm.SliderDB{}
	sliderDB.SliderPointersEncoding = input.SliderPointersEncoding
	sliderDB.CopyBasicFieldsFromSlider_WOP(&input.Slider_WOP)

	_, err = db.Create(&sliderDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSlider.CheckoutPhaseOneInstance(&sliderDB)
	slider := backRepo.BackRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]

	if slider != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), slider)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sliderDB)
}

// GetSlider
//
// swagger:route GET /sliders/{ID} sliders getSlider
//
// Gets the details for a slider.
//
// Responses:
// default: genericError
//
//	200: sliderDBResponse
func (controller *Controller) GetSlider(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlider", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSlider.GetDB()

	// Get sliderDB in DB
	var sliderDB orm.SliderDB
	if _, err := db.First(&sliderDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sliderAPI orm.SliderAPI
	sliderAPI.ID = sliderDB.ID
	sliderAPI.SliderPointersEncoding = sliderDB.SliderPointersEncoding
	sliderDB.CopyBasicFieldsToSlider_WOP(&sliderAPI.Slider_WOP)

	c.JSON(http.StatusOK, sliderAPI)
}

// UpdateSlider
//
// swagger:route PATCH /sliders/{ID} sliders updateSlider
//
// # Update a slider
//
// Responses:
// default: genericError
//
//	200: sliderDBResponse
func (controller *Controller) UpdateSlider(c *gin.Context) {

	mutexSlider.Lock()
	defer mutexSlider.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSlider", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSlider.GetDB()

	// Validate input
	var input orm.SliderAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sliderDB orm.SliderDB

	// fetch the slider
	_, err := db.First(&sliderDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sliderDB.CopyBasicFieldsFromSlider_WOP(&input.Slider_WOP)
	sliderDB.SliderPointersEncoding = input.SliderPointersEncoding

	db, _ = db.Model(&sliderDB)
	_, err = db.Updates(&sliderDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sliderNew := new(models.Slider)
	sliderDB.CopyBasicFieldsToSlider(sliderNew)

	// redeem pointers
	sliderDB.DecodePointers(backRepo, sliderNew)

	// get stage instance from DB instance, and call callback function
	sliderOld := backRepo.BackRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]
	if sliderOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sliderOld, sliderNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sliderDB
	c.JSON(http.StatusOK, sliderDB)
}

// DeleteSlider
//
// swagger:route DELETE /sliders/{ID} sliders deleteSlider
//
// # Delete a slider
//
// default: genericError
//
//	200: sliderDBResponse
func (controller *Controller) DeleteSlider(c *gin.Context) {

	mutexSlider.Lock()
	defer mutexSlider.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSlider", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/slider/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSlider.GetDB()

	// Get model if exist
	var sliderDB orm.SliderDB
	if _, err := db.First(&sliderDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&sliderDB)

	// get an instance (not staged) from DB instance, and call callback function
	sliderDeleted := new(models.Slider)
	sliderDB.CopyBasicFieldsToSlider(sliderDeleted)

	// get stage instance from DB instance, and call callback function
	sliderStaged := backRepo.BackRepoSlider.Map_SliderDBID_SliderPtr[sliderDB.ID]
	if sliderStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sliderStaged, sliderDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
