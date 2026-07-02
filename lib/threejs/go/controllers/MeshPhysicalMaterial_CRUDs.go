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
var __MeshPhysicalMaterial__dummysDeclaration__ models.MeshPhysicalMaterial
var _ = __MeshPhysicalMaterial__dummysDeclaration__
var __MeshPhysicalMaterial_time__dummyDeclaration time.Duration
var _ = __MeshPhysicalMaterial_time__dummyDeclaration

var mutexMeshPhysicalMaterial sync.Mutex

// An MeshPhysicalMaterialID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeshPhysicalMaterial updateMeshPhysicalMaterial deleteMeshPhysicalMaterial
type MeshPhysicalMaterialID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshPhysicalMaterialInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeshPhysicalMaterial updateMeshPhysicalMaterial
type MeshPhysicalMaterialInput struct {
	// The MeshPhysicalMaterial to submit or modify
	// in: body
	MeshPhysicalMaterial *orm.MeshPhysicalMaterialAPI
}

// GetMeshPhysicalMaterials
//
// swagger:route GET /meshphysicalmaterials meshphysicalmaterials getMeshPhysicalMaterials
//
// # Get all meshphysicalmaterials
//
// Responses:
// default: genericError
//
//	200: meshphysicalmaterialDBResponse
func (controller *Controller) GetMeshPhysicalMaterials(c *gin.Context) {

	// source slice
	var meshphysicalmaterialDBs []orm.MeshPhysicalMaterialDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeshPhysicalMaterials", "Name", stackPath)
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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	_, err := db.Find(&meshphysicalmaterialDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	meshphysicalmaterialAPIs := make([]orm.MeshPhysicalMaterialAPI, 0)

	// for each meshphysicalmaterial, update fields from the database nullable fields
	for idx := range meshphysicalmaterialDBs {
		meshphysicalmaterialDB := &meshphysicalmaterialDBs[idx]
		_ = meshphysicalmaterialDB
		var meshphysicalmaterialAPI orm.MeshPhysicalMaterialAPI

		// insertion point for updating fields
		meshphysicalmaterialAPI.ID = meshphysicalmaterialDB.ID
		meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial_WOP(&meshphysicalmaterialAPI.MeshPhysicalMaterial_WOP)
		meshphysicalmaterialAPI.MeshPhysicalMaterialPointersEncoding = meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding
		meshphysicalmaterialAPIs = append(meshphysicalmaterialAPIs, meshphysicalmaterialAPI)
	}

	c.JSON(http.StatusOK, meshphysicalmaterialAPIs)
}

// PostMeshPhysicalMaterial
//
// swagger:route POST /meshphysicalmaterials meshphysicalmaterials postMeshPhysicalMaterial
//
// Creates a meshphysicalmaterial
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeshPhysicalMaterial(c *gin.Context) {

	mutexMeshPhysicalMaterial.Lock()
	defer mutexMeshPhysicalMaterial.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeshPhysicalMaterials", "Name", stackPath)
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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	// Validate input
	var input orm.MeshPhysicalMaterialAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create meshphysicalmaterial
	meshphysicalmaterialDB := orm.MeshPhysicalMaterialDB{}
	meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding = input.MeshPhysicalMaterialPointersEncoding
	meshphysicalmaterialDB.CopyBasicFieldsFromMeshPhysicalMaterial_WOP(&input.MeshPhysicalMaterial_WOP)

	_, err = db.Create(&meshphysicalmaterialDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeshPhysicalMaterial.CheckoutPhaseOneInstance(&meshphysicalmaterialDB)
	meshphysicalmaterial := backRepo.BackRepoMeshPhysicalMaterial.Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialPtr[meshphysicalmaterialDB.ID]

	if meshphysicalmaterial != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), meshphysicalmaterial)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, meshphysicalmaterialDB)
}

// GetMeshPhysicalMaterial
//
// swagger:route GET /meshphysicalmaterials/{ID} meshphysicalmaterials getMeshPhysicalMaterial
//
// Gets the details for a meshphysicalmaterial.
//
// Responses:
// default: genericError
//
//	200: meshphysicalmaterialDBResponse
func (controller *Controller) GetMeshPhysicalMaterial(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeshPhysicalMaterial", "Name", stackPath)
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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	// Get meshphysicalmaterialDB in DB
	var meshphysicalmaterialDB orm.MeshPhysicalMaterialDB
	if _, err := db.First(&meshphysicalmaterialDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var meshphysicalmaterialAPI orm.MeshPhysicalMaterialAPI
	meshphysicalmaterialAPI.ID = meshphysicalmaterialDB.ID
	meshphysicalmaterialAPI.MeshPhysicalMaterialPointersEncoding = meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding
	meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial_WOP(&meshphysicalmaterialAPI.MeshPhysicalMaterial_WOP)

	c.JSON(http.StatusOK, meshphysicalmaterialAPI)
}

// UpdateMeshPhysicalMaterial
//
// swagger:route PATCH /meshphysicalmaterials/{ID} meshphysicalmaterials updateMeshPhysicalMaterial
//
// # Update a meshphysicalmaterial
//
// Responses:
// default: genericError
//
//	200: meshphysicalmaterialDBResponse
func (controller *Controller) UpdateMeshPhysicalMaterial(c *gin.Context) {

	mutexMeshPhysicalMaterial.Lock()
	defer mutexMeshPhysicalMaterial.Unlock()

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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	// Validate input
	var input orm.MeshPhysicalMaterialAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshphysicalmaterialDB orm.MeshPhysicalMaterialDB

	// fetch the meshphysicalmaterial
	_, err := db.First(&meshphysicalmaterialDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshphysicalmaterialDB.CopyBasicFieldsFromMeshPhysicalMaterial_WOP(&input.MeshPhysicalMaterial_WOP)
	meshphysicalmaterialDB.MeshPhysicalMaterialPointersEncoding = input.MeshPhysicalMaterialPointersEncoding

	db, _ = db.Model(&meshphysicalmaterialDB)
	_, err = db.Updates(&meshphysicalmaterialDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshphysicalmaterialNew := new(models.MeshPhysicalMaterial)
	meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial(meshphysicalmaterialNew)

	// redeem pointers
	meshphysicalmaterialDB.DecodePointers(backRepo, meshphysicalmaterialNew)

	// get stage instance from DB instance, and call callback function
	meshphysicalmaterialOld := backRepo.BackRepoMeshPhysicalMaterial.Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialPtr[meshphysicalmaterialDB.ID]
	if meshphysicalmaterialOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), meshphysicalmaterialOld, meshphysicalmaterialNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshphysicalmaterialDB
	c.JSON(http.StatusOK, meshphysicalmaterialDB)
}

// DeleteMeshPhysicalMaterial
//
// swagger:route DELETE /meshphysicalmaterials/{ID} meshphysicalmaterials deleteMeshPhysicalMaterial
//
// # Delete a meshphysicalmaterial
//
// default: genericError
//
//	200: meshphysicalmaterialDBResponse
func (controller *Controller) DeleteMeshPhysicalMaterial(c *gin.Context) {

	mutexMeshPhysicalMaterial.Lock()
	defer mutexMeshPhysicalMaterial.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeshPhysicalMaterial", "Name", stackPath)
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
	db := backRepo.BackRepoMeshPhysicalMaterial.GetDB()

	// Get model if exist
	var meshphysicalmaterialDB orm.MeshPhysicalMaterialDB
	if _, err := db.First(&meshphysicalmaterialDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&meshphysicalmaterialDB)

	// get an instance (not staged) from DB instance, and call callback function
	meshphysicalmaterialDeleted := new(models.MeshPhysicalMaterial)
	meshphysicalmaterialDB.CopyBasicFieldsToMeshPhysicalMaterial(meshphysicalmaterialDeleted)

	// get stage instance from DB instance, and call callback function
	meshphysicalmaterialStaged := backRepo.BackRepoMeshPhysicalMaterial.Map_MeshPhysicalMaterialDBID_MeshPhysicalMaterialPtr[meshphysicalmaterialDB.ID]
	if meshphysicalmaterialStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), meshphysicalmaterialStaged, meshphysicalmaterialDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
