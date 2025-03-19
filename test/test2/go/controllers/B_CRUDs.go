// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __B__dummysDeclaration__ models.B
var __B_time__dummyDeclaration time.Duration

var mutexB sync.Mutex

// An BID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getB updateB deleteB
type BID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postB updateB
type BInput struct {
	// The B to submit or modify
	// in: body
	B *orm.BAPI
}

// GetBs
//
// swagger:route GET /bs bs getBs
//
// # Get all bs
//
// Responses:
// default: genericError
//
//	200: bDBResponse
func (controller *Controller) GetBs(c *gin.Context) {

	// source slice
	var bDBs []orm.BDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoB.GetDB()

	_, err := db.Find(&bDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bAPIs := make([]orm.BAPI, 0)

	// for each b, update fields from the database nullable fields
	for idx := range bDBs {
		bDB := &bDBs[idx]
		_ = bDB
		var bAPI orm.BAPI

		// insertion point for updating fields
		bAPI.ID = bDB.ID
		bDB.CopyBasicFieldsToB_WOP(&bAPI.B_WOP)
		bAPI.BPointersEncoding = bDB.BPointersEncoding
		bAPIs = append(bAPIs, bAPI)
	}

	c.JSON(http.StatusOK, bAPIs)
}

// PostB
//
// swagger:route POST /bs bs postB
//
// Creates a b
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostB(c *gin.Context) {

	mutexB.Lock()
	defer mutexB.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoB.GetDB()

	// Validate input
	var input orm.BAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create b
	bDB := orm.BDB{}
	bDB.BPointersEncoding = input.BPointersEncoding
	bDB.CopyBasicFieldsFromB_WOP(&input.B_WOP)

	_, err = db.Create(&bDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoB.CheckoutPhaseOneInstance(&bDB)
	b := backRepo.BackRepoB.Map_BDBID_BPtr[bDB.ID]

	if b != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), b)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bDB)
}

// GetB
//
// swagger:route GET /bs/{ID} bs getB
//
// Gets the details for a b.
//
// Responses:
// default: genericError
//
//	200: bDBResponse
func (controller *Controller) GetB(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetB", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoB.GetDB()

	// Get bDB in DB
	var bDB orm.BDB
	if _, err := db.First(&bDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bAPI orm.BAPI
	bAPI.ID = bDB.ID
	bAPI.BPointersEncoding = bDB.BPointersEncoding
	bDB.CopyBasicFieldsToB_WOP(&bAPI.B_WOP)

	c.JSON(http.StatusOK, bAPI)
}

// UpdateB
//
// swagger:route PATCH /bs/{ID} bs updateB
//
// # Update a b
//
// Responses:
// default: genericError
//
//	200: bDBResponse
func (controller *Controller) UpdateB(c *gin.Context) {

	mutexB.Lock()
	defer mutexB.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateB", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoB.GetDB()

	// Validate input
	var input orm.BAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bDB orm.BDB

	// fetch the b
	_, err := db.First(&bDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bDB.CopyBasicFieldsFromB_WOP(&input.B_WOP)
	bDB.BPointersEncoding = input.BPointersEncoding

	db, _ = db.Model(&bDB)
	_, err = db.Updates(&bDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bNew := new(models.B)
	bDB.CopyBasicFieldsToB(bNew)

	// redeem pointers
	bDB.DecodePointers(backRepo, bNew)

	// get stage instance from DB instance, and call callback function
	bOld := backRepo.BackRepoB.Map_BDBID_BPtr[bDB.ID]
	if bOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bOld, bNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bDB
	c.JSON(http.StatusOK, bDB)
}

// DeleteB
//
// swagger:route DELETE /bs/{ID} bs deleteB
//
// # Delete a b
//
// default: genericError
//
//	200: bDBResponse
func (controller *Controller) DeleteB(c *gin.Context) {

	mutexB.Lock()
	defer mutexB.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteB", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/test/test2/go, Unkown stack: \"" + stackPath + "\""
		log.Panic(message)
	}
	db := backRepo.BackRepoB.GetDB()

	// Get model if exist
	var bDB orm.BDB
	if _, err := db.First(&bDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&bDB)

	// get an instance (not staged) from DB instance, and call callback function
	bDeleted := new(models.B)
	bDB.CopyBasicFieldsToB(bDeleted)

	// get stage instance from DB instance, and call callback function
	bStaged := backRepo.BackRepoB.Map_BDBID_BPtr[bDB.ID]
	if bStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bStaged, bDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
