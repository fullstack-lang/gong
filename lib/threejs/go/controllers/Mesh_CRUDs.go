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
var __Mesh__dummysDeclaration__ models.Mesh
var _ = __Mesh__dummysDeclaration__
var __Mesh_time__dummyDeclaration time.Duration
var _ = __Mesh_time__dummyDeclaration

var mutexMesh sync.Mutex

// An MeshID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMesh updateMesh deleteMesh
type MeshID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MeshInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMesh updateMesh
type MeshInput struct {
	// The Mesh to submit or modify
	// in: body
	Mesh *orm.MeshAPI
}

// GetMeshs
//
// swagger:route GET /meshs meshs getMeshs
//
// # Get all meshs
//
// Responses:
// default: genericError
//
//	200: meshDBResponse
func (controller *Controller) GetMeshs(c *gin.Context) {

	// source slice
	var meshDBs []orm.MeshDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeshs", "Name", stackPath)
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
	db := backRepo.BackRepoMesh.GetDB()

	_, err := db.Find(&meshDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	meshAPIs := make([]orm.MeshAPI, 0)

	// for each mesh, update fields from the database nullable fields
	for idx := range meshDBs {
		meshDB := &meshDBs[idx]
		_ = meshDB
		var meshAPI orm.MeshAPI

		// insertion point for updating fields
		meshAPI.ID = meshDB.ID
		meshDB.CopyBasicFieldsToMesh_WOP(&meshAPI.Mesh_WOP)
		meshAPI.MeshPointersEncoding = meshDB.MeshPointersEncoding
		meshAPIs = append(meshAPIs, meshAPI)
	}

	c.JSON(http.StatusOK, meshAPIs)
}

// PostMesh
//
// swagger:route POST /meshs meshs postMesh
//
// Creates a mesh
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMesh(c *gin.Context) {

	mutexMesh.Lock()
	defer mutexMesh.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeshs", "Name", stackPath)
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
	db := backRepo.BackRepoMesh.GetDB()

	// Validate input
	var input orm.MeshAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mesh
	meshDB := orm.MeshDB{}
	meshDB.MeshPointersEncoding = input.MeshPointersEncoding
	meshDB.CopyBasicFieldsFromMesh_WOP(&input.Mesh_WOP)

	_, err = db.Create(&meshDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMesh.CheckoutPhaseOneInstance(&meshDB)
	mesh := backRepo.BackRepoMesh.Map_MeshDBID_MeshPtr[meshDB.ID]

	if mesh != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mesh)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, meshDB)
}

// GetMesh
//
// swagger:route GET /meshs/{ID} meshs getMesh
//
// Gets the details for a mesh.
//
// Responses:
// default: genericError
//
//	200: meshDBResponse
func (controller *Controller) GetMesh(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMesh", "Name", stackPath)
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
	db := backRepo.BackRepoMesh.GetDB()

	// Get meshDB in DB
	var meshDB orm.MeshDB
	if _, err := db.First(&meshDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var meshAPI orm.MeshAPI
	meshAPI.ID = meshDB.ID
	meshAPI.MeshPointersEncoding = meshDB.MeshPointersEncoding
	meshDB.CopyBasicFieldsToMesh_WOP(&meshAPI.Mesh_WOP)

	c.JSON(http.StatusOK, meshAPI)
}

// UpdateMesh
//
// swagger:route PATCH /meshs/{ID} meshs updateMesh
//
// # Update a mesh
//
// Responses:
// default: genericError
//
//	200: meshDBResponse
func (controller *Controller) UpdateMesh(c *gin.Context) {

	mutexMesh.Lock()
	defer mutexMesh.Unlock()

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
	db := backRepo.BackRepoMesh.GetDB()

	// Validate input
	var input orm.MeshAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var meshDB orm.MeshDB

	// fetch the mesh
	_, err := db.First(&meshDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	meshDB.CopyBasicFieldsFromMesh_WOP(&input.Mesh_WOP)
	meshDB.MeshPointersEncoding = input.MeshPointersEncoding

	db, _ = db.Model(&meshDB)
	_, err = db.Updates(&meshDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	meshNew := new(models.Mesh)
	meshDB.CopyBasicFieldsToMesh(meshNew)

	// redeem pointers
	meshDB.DecodePointers(backRepo, meshNew)

	// get stage instance from DB instance, and call callback function
	meshOld := backRepo.BackRepoMesh.Map_MeshDBID_MeshPtr[meshDB.ID]
	if meshOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), meshOld, meshNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the meshDB
	c.JSON(http.StatusOK, meshDB)
}

// DeleteMesh
//
// swagger:route DELETE /meshs/{ID} meshs deleteMesh
//
// # Delete a mesh
//
// default: genericError
//
//	200: meshDBResponse
func (controller *Controller) DeleteMesh(c *gin.Context) {

	mutexMesh.Lock()
	defer mutexMesh.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMesh", "Name", stackPath)
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
	db := backRepo.BackRepoMesh.GetDB()

	// Get model if exist
	var meshDB orm.MeshDB
	if _, err := db.First(&meshDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&meshDB)

	// get an instance (not staged) from DB instance, and call callback function
	meshDeleted := new(models.Mesh)
	meshDB.CopyBasicFieldsToMesh(meshDeleted)

	// get stage instance from DB instance, and call callback function
	meshStaged := backRepo.BackRepoMesh.Map_MeshDBID_MeshPtr[meshDB.ID]
	if meshStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), meshStaged, meshDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
