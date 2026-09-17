// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"
)

// declaration in order to justify use of the models import
var __Slider__dummysDeclaration__ models.Slider
var _ = __Slider__dummysDeclaration__
var __Slider_time__dummyDeclaration time.Duration
var _ = __Slider_time__dummyDeclaration

var mutexSlider sync.Mutex

// An SliderID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateSlider
type SliderID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SliderInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateSlider
type SliderInput struct {
	// The Slider to submit or modify
	// in: body
	Slider *orm.SliderAPI
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
func (controller *Controller) UpdateSlider(w http.ResponseWriter, r *http.Request) {

	mutexSlider.Lock()
	defer mutexSlider.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sliderDB orm.SliderDB

	// fetch the slider
	_, err := db.First(&sliderDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
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
		writeJSON(w, http.StatusBadRequest, returnError.Body)
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
		backRepo.GetStage().OnAfterUpdateFromFront(sliderOld, sliderNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sliderDB
	writeJSON(w, http.StatusOK, sliderDB)
}
