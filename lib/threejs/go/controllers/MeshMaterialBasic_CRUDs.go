// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __MeshMaterialBasic__dummysDeclaration__ models.MeshMaterialBasic
var _ = __MeshMaterialBasic__dummysDeclaration__
var __MeshMaterialBasic_time__dummyDeclaration time.Duration
var _ = __MeshMaterialBasic_time__dummyDeclaration

var mutexMeshMaterialBasic sync.Mutex

// An MeshMaterialBasicID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeshMaterialBasic updateMeshMaterialBasic deleteMeshMaterialBasic
type MeshMaterialBasicID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshMaterialBasicInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeshMaterialBasic updateMeshMaterialBasic
type MeshMaterialBasicInput struct {
	// The MeshMaterialBasic to submit or modify
	// in: body
	MeshMaterialBasic *orm.MeshMaterialBasicAPI
}

// GetMeshMaterialBasics
//
// swagger:route GET /meshmaterialbasics meshmaterialbasics getMeshMaterialBasics
//
// # Get all meshmaterialbasics
//
// Responses:
// default: genericError
//
//	200: meshmaterialbasicDBResponse
func (controller *Controller) GetMeshMaterialBasics(c *gin.Context) {

	// source slice
	var meshmaterialbasicDBs []orm.MeshMaterialBasicDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeshMaterialBasics", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	_, err := db.Find(&meshmaterialbasicDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	meshmaterialbasicAPIs := make([]orm.MeshMaterialBasicAPI, 0)

	// for each meshmaterialbasic, update fields from the database nullable fields
	for idx := range meshmaterialbasicDBs {
		meshmaterialbasicDB := &meshmaterialbasicDBs[idx]
		_ = meshmaterialbasicDB
		var meshmaterialbasicAPI orm.MeshMaterialBasicAPI

		// insertion point for updating fields
		meshmaterialbasicAPI.ID = meshmaterialbasicDB.ID
		meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic_WOP(&meshmaterialbasicAPI.MeshMaterialBasic_WOP)
		meshmaterialbasicAPI.MeshMaterialBasicPointersEncoding = meshmaterialbasicDB.MeshMaterialBasicPointersEncoding
		meshmaterialbasicAPIs = append(meshmaterialbasicAPIs, meshmaterialbasicAPI)
	}

	c.JSON(http.StatusOK, meshmaterialbasicAPIs)
}

// PostMeshMaterialBasic
//
// swagger:route POST /meshmaterialbasics meshmaterialbasics postMeshMaterialBasic
//
// Creates a meshmaterialbasic
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeshMaterialBasic(c *gin.Context) {

	mutexMeshMaterialBasic.Lock()
	defer mutexMeshMaterialBasic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeshMaterialBasics", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	// Validate input
	var input orm.MeshMaterialBasicAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create meshmaterialbasic
	meshmaterialbasicDB := orm.MeshMaterialBasicDB{}
	meshmaterialbasicDB.MeshMaterialBasicPointersEncoding = input.MeshMaterialBasicPointersEncoding
	meshmaterialbasicDB.CopyBasicFieldsFromMeshMaterialBasic_WOP(&input.MeshMaterialBasic_WOP)

	_, err = db.Create(&meshmaterialbasicDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeshMaterialBasic.CheckoutPhaseOneInstance(&meshmaterialbasicDB)
	meshmaterialbasic := backRepo.BackRepoMeshMaterialBasic.Map_MeshMaterialBasicDBID_MeshMaterialBasicPtr[meshmaterialbasicDB.ID]

	if meshmaterialbasic != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), meshmaterialbasic)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, meshmaterialbasicDB)
}

// GetMeshMaterialBasic
//
// swagger:route GET /meshmaterialbasics/{ID} meshmaterialbasics getMeshMaterialBasic
//
// Gets the details for a meshmaterialbasic.
//
// Responses:
// default: genericError
//
//	200: meshmaterialbasicDBResponse
func (controller *Controller) GetMeshMaterialBasic(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeshMaterialBasic", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	// Get meshmaterialbasicDB in DB
	var meshmaterialbasicDB orm.MeshMaterialBasicDB
	if _, err := db.First(&meshmaterialbasicDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var meshmaterialbasicAPI orm.MeshMaterialBasicAPI
	meshmaterialbasicAPI.ID = meshmaterialbasicDB.ID
	meshmaterialbasicAPI.MeshMaterialBasicPointersEncoding = meshmaterialbasicDB.MeshMaterialBasicPointersEncoding
	meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic_WOP(&meshmaterialbasicAPI.MeshMaterialBasic_WOP)

	c.JSON(http.StatusOK, meshmaterialbasicAPI)
}

// UpdateMeshMaterialBasic
//
// swagger:route PATCH /meshmaterialbasics/{ID} meshmaterialbasics updateMeshMaterialBasic
//
// # Update a meshmaterialbasic
//
// Responses:
// default: genericError
//
//	200: meshmaterialbasicDBResponse
func (controller *Controller) UpdateMeshMaterialBasic(c *gin.Context) {

	mutexMeshMaterialBasic.Lock()
	defer mutexMeshMaterialBasic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	// Validate input
	var input orm.MeshMaterialBasicAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshmaterialbasicDB orm.MeshMaterialBasicDB

	// fetch the meshmaterialbasic
	_, err := db.First(&meshmaterialbasicDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshmaterialbasicDB.CopyBasicFieldsFromMeshMaterialBasic_WOP(&input.MeshMaterialBasic_WOP)
	meshmaterialbasicDB.MeshMaterialBasicPointersEncoding = input.MeshMaterialBasicPointersEncoding

	db, _ = db.Model(&meshmaterialbasicDB)
	_, err = db.Updates(&meshmaterialbasicDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshmaterialbasicNew := new(models.MeshMaterialBasic)
	meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic(meshmaterialbasicNew)

	// redeem pointers
	meshmaterialbasicDB.DecodePointers(backRepo, meshmaterialbasicNew)

	// get stage instance from DB instance, and call callback function
	meshmaterialbasicOld := backRepo.BackRepoMeshMaterialBasic.Map_MeshMaterialBasicDBID_MeshMaterialBasicPtr[meshmaterialbasicDB.ID]
	if meshmaterialbasicOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), meshmaterialbasicOld, meshmaterialbasicNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshmaterialbasicDB
	c.JSON(http.StatusOK, meshmaterialbasicDB)
}

// DeleteMeshMaterialBasic
//
// swagger:route DELETE /meshmaterialbasics/{ID} meshmaterialbasics deleteMeshMaterialBasic
//
// # Delete a meshmaterialbasic
//
// default: genericError
//
//	200: meshmaterialbasicDBResponse
func (controller *Controller) DeleteMeshMaterialBasic(c *gin.Context) {

	mutexMeshMaterialBasic.Lock()
	defer mutexMeshMaterialBasic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeshMaterialBasic", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMeshMaterialBasic.GetDB()

	// Get model if exist
	var meshmaterialbasicDB orm.MeshMaterialBasicDB
	if _, err := db.First(&meshmaterialbasicDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&meshmaterialbasicDB)

	// get an instance (not staged) from DB instance, and call callback function
	meshmaterialbasicDeleted := new(models.MeshMaterialBasic)
	meshmaterialbasicDB.CopyBasicFieldsToMeshMaterialBasic(meshmaterialbasicDeleted)

	// get stage instance from DB instance, and call callback function
	meshmaterialbasicStaged := backRepo.BackRepoMeshMaterialBasic.Map_MeshMaterialBasicDBID_MeshMaterialBasicPtr[meshmaterialbasicDB.ID]
	if meshmaterialbasicStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), meshmaterialbasicStaged, meshmaterialbasicDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
