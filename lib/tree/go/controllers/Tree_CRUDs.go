// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/orm"
)

// declaration in order to justify use of the models import
var __Tree__dummysDeclaration__ models.Tree
var _ = __Tree__dummysDeclaration__
var __Tree_time__dummyDeclaration time.Duration
var _ = __Tree_time__dummyDeclaration

var mutexTree sync.Mutex

// An TreeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateTree
type TreeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TreeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateTree
type TreeInput struct {
	// The Tree to submit or modify
	// in: body
	Tree *orm.TreeAPI
}

// UpdateTree
//
// swagger:route PATCH /trees/{ID} trees updateTree
//
// # Update a tree
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func (controller *Controller) UpdateTree(w http.ResponseWriter, r *http.Request) {

	mutexTree.Lock()
	defer mutexTree.Unlock()

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
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/tree/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoTree.GetDB()

	// Validate input
	var input orm.TreeAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var treeDB orm.TreeDB

	// fetch the tree
	_, err := db.First(&treeDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	treeDB.CopyBasicFieldsFromTree_WOP(&input.Tree_WOP)
	treeDB.TreePointersEncoding = input.TreePointersEncoding

	db, _ = db.Model(&treeDB)
	_, err = db.Updates(&treeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	treeNew := new(models.Tree)
	treeDB.CopyBasicFieldsToTree(treeNew)

	// redeem pointers
	treeDB.DecodePointers(backRepo, treeNew)

	// get stage instance from DB instance, and call callback function
	treeOld := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]
	if treeOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(treeOld, treeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the treeDB
	writeJSON(w, http.StatusOK, treeDB)
}
