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
var __Doc__dummysDeclaration__ models.Doc
var __Doc_time__dummyDeclaration time.Duration

var mutexDoc sync.Mutex

// An DocID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDoc updateDoc deleteDoc
type DocID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DocInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDoc updateDoc
type DocInput struct {
	// The Doc to submit or modify
	// in: body
	Doc *orm.DocAPI
}

// GetDocs
//
// swagger:route GET /docs docs getDocs
//
// # Get all docs
//
// Responses:
// default: genericError
//
//	200: docDBResponse
func (controller *Controller) GetDocs(c *gin.Context) {

	// source slice
	var docDBs []orm.DocDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDoc.GetDB()

	_, err := db.Find(&docDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	docAPIs := make([]orm.DocAPI, 0)

	// for each doc, update fields from the database nullable fields
	for idx := range docDBs {
		docDB := &docDBs[idx]
		_ = docDB
		var docAPI orm.DocAPI

		// insertion point for updating fields
		docAPI.ID = docDB.ID
		docDB.CopyBasicFieldsToDoc_WOP(&docAPI.Doc_WOP)
		docAPI.DocPointersEncoding = docDB.DocPointersEncoding
		docAPIs = append(docAPIs, docAPI)
	}

	c.JSON(http.StatusOK, docAPIs)
}

// PostDoc
//
// swagger:route POST /docs docs postDoc
//
// Creates a doc
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDoc(c *gin.Context) {

	mutexDoc.Lock()
	defer mutexDoc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDocs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDoc.GetDB()

	// Validate input
	var input orm.DocAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create doc
	docDB := orm.DocDB{}
	docDB.DocPointersEncoding = input.DocPointersEncoding
	docDB.CopyBasicFieldsFromDoc_WOP(&input.Doc_WOP)

	_, err = db.Create(&docDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDoc.CheckoutPhaseOneInstance(&docDB)
	doc := backRepo.BackRepoDoc.Map_DocDBID_DocPtr[docDB.ID]

	if doc != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), doc)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, docDB)
}

// GetDoc
//
// swagger:route GET /docs/{ID} docs getDoc
//
// Gets the details for a doc.
//
// Responses:
// default: genericError
//
//	200: docDBResponse
func (controller *Controller) GetDoc(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDoc", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDoc.GetDB()

	// Get docDB in DB
	var docDB orm.DocDB
	if _, err := db.First(&docDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var docAPI orm.DocAPI
	docAPI.ID = docDB.ID
	docAPI.DocPointersEncoding = docDB.DocPointersEncoding
	docDB.CopyBasicFieldsToDoc_WOP(&docAPI.Doc_WOP)

	c.JSON(http.StatusOK, docAPI)
}

// UpdateDoc
//
// swagger:route PATCH /docs/{ID} docs updateDoc
//
// # Update a doc
//
// Responses:
// default: genericError
//
//	200: docDBResponse
func (controller *Controller) UpdateDoc(c *gin.Context) {

	mutexDoc.Lock()
	defer mutexDoc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDoc", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDoc.GetDB()

	// Validate input
	var input orm.DocAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var docDB orm.DocDB

	// fetch the doc
	_, err := db.First(&docDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	docDB.CopyBasicFieldsFromDoc_WOP(&input.Doc_WOP)
	docDB.DocPointersEncoding = input.DocPointersEncoding

	db, _ = db.Model(&docDB)
	_, err = db.Updates(&docDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	docNew := new(models.Doc)
	docDB.CopyBasicFieldsToDoc(docNew)

	// redeem pointers
	docDB.DecodePointers(backRepo, docNew)

	// get stage instance from DB instance, and call callback function
	docOld := backRepo.BackRepoDoc.Map_DocDBID_DocPtr[docDB.ID]
	if docOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), docOld, docNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the docDB
	c.JSON(http.StatusOK, docDB)
}

// DeleteDoc
//
// swagger:route DELETE /docs/{ID} docs deleteDoc
//
// # Delete a doc
//
// default: genericError
//
//	200: docDBResponse
func (controller *Controller) DeleteDoc(c *gin.Context) {

	mutexDoc.Lock()
	defer mutexDoc.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDoc", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoDoc.GetDB()

	// Get model if exist
	var docDB orm.DocDB
	if _, err := db.First(&docDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&docDB)

	// get an instance (not staged) from DB instance, and call callback function
	docDeleted := new(models.Doc)
	docDB.CopyBasicFieldsToDoc(docDeleted)

	// get stage instance from DB instance, and call callback function
	docStaged := backRepo.BackRepoDoc.Map_DocDBID_DocPtr[docDB.ID]
	if docStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), docStaged, docDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
