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
var __Node__dummysDeclaration__ models.Node
var _ = __Node__dummysDeclaration__
var __Node_time__dummyDeclaration time.Duration
var _ = __Node_time__dummyDeclaration

var mutexNode sync.Mutex

// An NodeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters updateNode
type NodeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NodeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters updateNode
type NodeInput struct {
	// The Node to submit or modify
	// in: body
	Node *orm.NodeAPI
}

// UpdateNode
//
// swagger:route PATCH /nodes/{ID} nodes updateNode
//
// # Update a node
//
// Responses:
// default: genericError
//
//	200: nodeDBResponse
func (controller *Controller) UpdateNode(w http.ResponseWriter, r *http.Request) {

	mutexNode.Lock()
	defer mutexNode.Unlock()

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
	db := backRepo.BackRepoNode.GetDB()

	// Validate input
	var input orm.NodeAPI
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, H{"error": err.Error()})
		return
	}

	// Get model if exist
	var nodeDB orm.NodeDB

	// fetch the node
	_, err := db.First(&nodeDB, r.PathValue("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	nodeDB.CopyBasicFieldsFromNode_WOP(&input.Node_WOP)
	nodeDB.NodePointersEncoding = input.NodePointersEncoding

	db, _ = db.Model(&nodeDB)
	_, err = db.Updates(&nodeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	nodeNew := new(models.Node)
	nodeDB.CopyBasicFieldsToNode(nodeNew)

	// redeem pointers
	nodeDB.DecodePointers(backRepo, nodeNew)

	// get stage instance from DB instance, and call callback function
	nodeOld := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
	if nodeOld != nil {
		backRepo.GetStage().OnAfterUpdateFromFront(nodeOld, nodeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the nodeDB
	writeJSON(w, http.StatusOK, nodeDB)
}
