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
var __FormGroup__dummysDeclaration__ models.FormGroup
var _ = __FormGroup__dummysDeclaration__
var __FormGroup_time__dummyDeclaration time.Duration
var _ = __FormGroup_time__dummyDeclaration

var mutexFormGroup sync.Mutex

// An FormGroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateFormGroup
type FormGroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormGroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateFormGroup
type FormGroupInput struct {
	// The FormGroup to submit or modify
	// in: body
	FormGroup *orm.FormGroupAPI
}

// UpdateFormGroup
//
// swagger:route PATCH /formgroups/{ID} formgroups updateFormGroup
//
// # Update a formgroup
//
// Responses:
// default: genericError
//
//	200: formgroupDBResponse
func (controller *Controller) UpdateFormGroup(w http.ResponseWriter, r *http.Request) {

	mutexFormGroup.Lock()
	defer mutexFormGroup.Unlock()

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
	db := backRepo.BackRepoFormGroup.GetDB()

	// Validate input
	var input orm.FormGroupAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formgroupDB orm.FormGroupDB

	// fetch the formgroup
	_, err := db.First(&formgroupDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formgroupDB.CopyBasicFieldsFromFormGroup_WOP(&input.FormGroup_WOP)
	formgroupDB.FormGroupPointersEncoding = input.FormGroupPointersEncoding

	db, _ = db.Model(&formgroupDB)
	_, err = db.Updates(&formgroupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formgroupNew := new(models.FormGroup)
	formgroupDB.CopyBasicFieldsToFormGroup(formgroupNew)

	// redeem pointers
	formgroupDB.DecodePointers(backRepo, formgroupNew)

	// get stage instance from DB instance, and call callback function
	formgroupOld := backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
	if formgroupOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(formgroupOld, formgroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formgroupDB
	writeJSON(w, http.StatusOK, formgroupDB)
}
