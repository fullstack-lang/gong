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
var __Tree__dummysDeclaration__ models.Tree
var __Tree_time__dummyDeclaration time.Duration

var mutexTree sync.Mutex

// An TreeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTree updateTree deleteTree
type TreeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TreeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTree updateTree
type TreeInput struct {
	// The Tree to submit or modify
	// in: body
	Tree *orm.TreeAPI
}

// GetTrees
//
// swagger:route GET /trees trees getTrees
//
// # Get all trees
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func (controller *Controller) GetTrees(c *gin.Context) {

	// source slice
	var treeDBs []orm.TreeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTrees", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTree.GetDB()

	_, err := db.Find(&treeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	treeAPIs := make([]orm.TreeAPI, 0)

	// for each tree, update fields from the database nullable fields
	for idx := range treeDBs {
		treeDB := &treeDBs[idx]
		_ = treeDB
		var treeAPI orm.TreeAPI

		// insertion point for updating fields
		treeAPI.ID = treeDB.ID
		treeDB.CopyBasicFieldsToTree_WOP(&treeAPI.Tree_WOP)
		treeAPI.TreePointersEncoding = treeDB.TreePointersEncoding
		treeAPIs = append(treeAPIs, treeAPI)
	}

	c.JSON(http.StatusOK, treeAPIs)
}

// PostTree
//
// swagger:route POST /trees trees postTree
//
// Creates a tree
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTree(c *gin.Context) {

	mutexTree.Lock()
	defer mutexTree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTrees", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTree.GetDB()

	// Validate input
	var input orm.TreeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tree
	treeDB := orm.TreeDB{}
	treeDB.TreePointersEncoding = input.TreePointersEncoding
	treeDB.CopyBasicFieldsFromTree_WOP(&input.Tree_WOP)

	_, err = db.Create(&treeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTree.CheckoutPhaseOneInstance(&treeDB)
	tree := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]

	if tree != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tree)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, treeDB)
}

// GetTree
//
// swagger:route GET /trees/{ID} trees getTree
//
// Gets the details for a tree.
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func (controller *Controller) GetTree(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTree.GetDB()

	// Get treeDB in DB
	var treeDB orm.TreeDB
	if _, err := db.First(&treeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var treeAPI orm.TreeAPI
	treeAPI.ID = treeDB.ID
	treeAPI.TreePointersEncoding = treeDB.TreePointersEncoding
	treeDB.CopyBasicFieldsToTree_WOP(&treeAPI.Tree_WOP)

	c.JSON(http.StatusOK, treeAPI)
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
func (controller *Controller) UpdateTree(c *gin.Context) {

	mutexTree.Lock()
	defer mutexTree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTree.GetDB()

	// Validate input
	var input orm.TreeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var treeDB orm.TreeDB

	// fetch the tree
	_, err := db.First(&treeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
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
		c.JSON(http.StatusBadRequest, returnError.Body)
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
		models.AfterUpdateFromFront(backRepo.GetStage(), treeOld, treeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the treeDB
	c.JSON(http.StatusOK, treeDB)
}

// DeleteTree
//
// swagger:route DELETE /trees/{ID} trees deleteTree
//
// # Delete a tree
//
// default: genericError
//
//	200: treeDBResponse
func (controller *Controller) DeleteTree(c *gin.Context) {

	mutexTree.Lock()
	defer mutexTree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTree.GetDB()

	// Get model if exist
	var treeDB orm.TreeDB
	if _, err := db.First(&treeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&treeDB)

	// get an instance (not staged) from DB instance, and call callback function
	treeDeleted := new(models.Tree)
	treeDB.CopyBasicFieldsToTree(treeDeleted)

	// get stage instance from DB instance, and call callback function
	treeStaged := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]
	if treeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), treeStaged, treeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
