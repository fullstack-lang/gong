// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// declaration in order to justify use of the models import
var __Animate__dummysDeclaration__ models.Animate
var _ = __Animate__dummysDeclaration__
var __Animate_time__dummyDeclaration time.Duration
var _ = __Animate_time__dummyDeclaration

var mutexAnimate sync.Mutex

// An AnimateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAnimate updateAnimate deleteAnimate
type AnimateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AnimateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAnimate updateAnimate
type AnimateInput struct {
	// The Animate to submit or modify
	// in: body
	Animate *orm.AnimateAPI
}

// GetAnimates
//
// swagger:route GET /animates animates getAnimates
//
// # Get all animates
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) GetAnimates(w http.ResponseWriter, r *http.Request) {

	// source slice
	var animateDBs []orm.AnimateDB

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnimates", "Name", stackPath)
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
	db := backRepo.BackRepoAnimate.GetDB()

	_, err := db.Find(&animateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	animateAPIs := make([]orm.AnimateAPI, 0)

	// for each animate, update fields from the database nullable fields
	for idx := range animateDBs {
		animateDB := &animateDBs[idx]
		_ = animateDB
		var animateAPI orm.AnimateAPI

		// insertion point for updating fields
		animateAPI.ID = animateDB.ID
		animateDB.CopyBasicFieldsToAnimate_WOP(&animateAPI.Animate_WOP)
		animateAPI.AnimatePointersEncoding = animateDB.AnimatePointersEncoding
		animateAPIs = append(animateAPIs, animateAPI)
	}

	writeJSON(w, http.StatusOK, animateAPIs)
}

// PostAnimate
//
// swagger:route POST /animates animates postAnimate
//
// Creates a animate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAnimate(w http.ResponseWriter, r *http.Request) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAnimates", "Name", stackPath)
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
	db := backRepo.BackRepoAnimate.GetDB()

	// Validate input
	var input orm.AnimateAPI

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// Create animate
	animateDB := orm.AnimateDB{}
	animateDB.AnimatePointersEncoding = input.AnimatePointersEncoding
	animateDB.CopyBasicFieldsFromAnimate_WOP(&input.Animate_WOP)

	_, err = db.Create(&animateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAnimate.CheckoutPhaseOneInstance(&animateDB)
	animate := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]

	if animate != nil {
		backRepo.GetStage().AfterCreateFromFront(animate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, animateDB)
}

// GetAnimate
//
// swagger:route GET /animates/{ID} animates getAnimate
//
// Gets the details for a animate.
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) GetAnimate(w http.ResponseWriter, r *http.Request) {

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnimate", "Name", stackPath)
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
	db := backRepo.BackRepoAnimate.GetDB()

	// Get animateDB in DB
	var animateDB orm.AnimateDB
	if _, err := db.First(&animateDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	var animateAPI orm.AnimateAPI
	animateAPI.ID = animateDB.ID
	animateAPI.AnimatePointersEncoding = animateDB.AnimatePointersEncoding
	animateDB.CopyBasicFieldsToAnimate_WOP(&animateAPI.Animate_WOP)

	writeJSON(w, http.StatusOK, animateAPI)
}

// UpdateAnimate
//
// swagger:route PATCH /animates/{ID} animates updateAnimate
//
// # Update a animate
//
// Responses:
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) UpdateAnimate(w http.ResponseWriter, r *http.Request) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoAnimate.GetDB()

	// Validate input
	var input orm.AnimateAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var animateDB orm.AnimateDB

	// fetch the animate
	_, err := db.First(&animateDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	animateDB.CopyBasicFieldsFromAnimate_WOP(&input.Animate_WOP)
	animateDB.AnimatePointersEncoding = input.AnimatePointersEncoding

	db, _ = db.Model(&animateDB)
	_, err = db.Updates(&animateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	animateNew := new(models.Animate)
	animateDB.CopyBasicFieldsToAnimate(animateNew)

	// redeem pointers
	animateDB.DecodePointers(backRepo, animateNew)

	// get stage instance from DB instance, and call callback function
	animateOld := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	if animateOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(animateOld, animateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the animateDB
	writeJSON(w, http.StatusOK, animateDB)
}

// DeleteAnimate
//
// swagger:route DELETE /animates/{ID} animates deleteAnimate
//
// # Delete a animate
//
// default: genericError
//
//	200: animateDBResponse
func (controller *Controller) DeleteAnimate(w http.ResponseWriter, r *http.Request) {

	mutexAnimate.Lock()
	defer mutexAnimate.Unlock()

	_values := r.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAnimate", "Name", stackPath)
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
	db := backRepo.BackRepoAnimate.GetDB()

	// Get model if exist
	var animateDB orm.AnimateDB
	if _, err := db.First(&animateDB, r.PathValue("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&animateDB)

	// get an instance (not staged) from DB instance, and call callback function
	animateDeleted := new(models.Animate)
	animateDB.CopyBasicFieldsToAnimate(animateDeleted)

	// get stage instance from DB instance, and call callback function
	animateStaged := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	if animateStaged != nil {
		backRepo.GetStage().AfterDeleteFromFront(animateStaged, animateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	writeJSON(w, http.StatusOK, H{"data": true})
}
