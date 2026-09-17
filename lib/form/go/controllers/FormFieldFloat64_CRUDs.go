// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	"github.com/fullstack-lang/gong/lib/form/go/orm"
)

// declaration in order to justify use of the models import
var __FormFieldFloat64__dummysDeclaration__ models.FormFieldFloat64
var _ = __FormFieldFloat64__dummysDeclaration__
var __FormFieldFloat64_time__dummyDeclaration time.Duration
var _ = __FormFieldFloat64_time__dummyDeclaration

var mutexFormFieldFloat64 sync.Mutex

// An FormFieldFloat64ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormFieldFloat64
type FormFieldFloat64ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldFloat64Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormFieldFloat64
type FormFieldFloat64Input struct {
	// The FormFieldFloat64 to submit or modify
	// in: body
	FormFieldFloat64 *orm.FormFieldFloat64API
}

// UpdateFormFieldFloat64
//
// swagger:route PATCH /formfieldfloat64s/{ID} formfieldfloat64s updateFormFieldFloat64
//
// # Update a formfieldfloat64
//
// Responses:
// default: genericError
//
//	200: formfieldfloat64DBResponse
func (controller *Controller) UpdateFormFieldFloat64(w http.ResponseWriter, r *http.Request) {

	mutexFormFieldFloat64.Lock()
	defer mutexFormFieldFloat64.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/form/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoFormFieldFloat64.GetDB()

	// Validate input
	var input orm.FormFieldFloat64API
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldfloat64DB orm.FormFieldFloat64DB

	// fetch the formfieldfloat64
	_, err := db.First(&formfieldfloat64DB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64_WOP(&input.FormFieldFloat64_WOP)
	formfieldfloat64DB.FormFieldFloat64PointersEncoding = input.FormFieldFloat64PointersEncoding

	db, _ = db.Model(&formfieldfloat64DB)
	_, err = db.Updates(&formfieldfloat64DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldfloat64New := new(models.FormFieldFloat64)
	formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64(formfieldfloat64New)

	// redeem pointers
	formfieldfloat64DB.DecodePointers(backRepo, formfieldfloat64New)

	// get stage instance from DB instance, and call callback function
	formfieldfloat64Old := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
	if formfieldfloat64Old != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formfieldfloat64Old, formfieldfloat64New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldfloat64DB
	writeJSON(w, http.StatusOK, formfieldfloat64DB)
}
