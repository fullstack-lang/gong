// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	"github.com/fullstack-lang/gong/lib/load/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FileToDownload__dummysDeclaration__ models.FileToDownload
var __FileToDownload_time__dummyDeclaration time.Duration

var mutexFileToDownload sync.Mutex

// An FileToDownloadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFileToDownload updateFileToDownload deleteFileToDownload
type FileToDownloadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FileToDownloadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFileToDownload updateFileToDownload
type FileToDownloadInput struct {
	// The FileToDownload to submit or modify
	// in: body
	FileToDownload *orm.FileToDownloadAPI
}

// GetFileToDownloads
//
// swagger:route GET /filetodownloads filetodownloads getFileToDownloads
//
// # Get all filetodownloads
//
// Responses:
// default: genericError
//
//	200: filetodownloadDBResponse
func (controller *Controller) GetFileToDownloads(c *gin.Context) {

	// source slice
	var filetodownloadDBs []orm.FileToDownloadDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFileToDownloads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	_, err := db.Find(&filetodownloadDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	filetodownloadAPIs := make([]orm.FileToDownloadAPI, 0)

	// for each filetodownload, update fields from the database nullable fields
	for idx := range filetodownloadDBs {
		filetodownloadDB := &filetodownloadDBs[idx]
		_ = filetodownloadDB
		var filetodownloadAPI orm.FileToDownloadAPI

		// insertion point for updating fields
		filetodownloadAPI.ID = filetodownloadDB.ID
		filetodownloadDB.CopyBasicFieldsToFileToDownload_WOP(&filetodownloadAPI.FileToDownload_WOP)
		filetodownloadAPI.FileToDownloadPointersEncoding = filetodownloadDB.FileToDownloadPointersEncoding
		filetodownloadAPIs = append(filetodownloadAPIs, filetodownloadAPI)
	}

	c.JSON(http.StatusOK, filetodownloadAPIs)
}

// PostFileToDownload
//
// swagger:route POST /filetodownloads filetodownloads postFileToDownload
//
// Creates a filetodownload
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFileToDownload(c *gin.Context) {

	mutexFileToDownload.Lock()
	defer mutexFileToDownload.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFileToDownloads", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	// Validate input
	var input orm.FileToDownloadAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create filetodownload
	filetodownloadDB := orm.FileToDownloadDB{}
	filetodownloadDB.FileToDownloadPointersEncoding = input.FileToDownloadPointersEncoding
	filetodownloadDB.CopyBasicFieldsFromFileToDownload_WOP(&input.FileToDownload_WOP)

	_, err = db.Create(&filetodownloadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFileToDownload.CheckoutPhaseOneInstance(&filetodownloadDB)
	filetodownload := backRepo.BackRepoFileToDownload.Map_FileToDownloadDBID_FileToDownloadPtr[filetodownloadDB.ID]

	if filetodownload != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), filetodownload)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, filetodownloadDB)
}

// GetFileToDownload
//
// swagger:route GET /filetodownloads/{ID} filetodownloads getFileToDownload
//
// Gets the details for a filetodownload.
//
// Responses:
// default: genericError
//
//	200: filetodownloadDBResponse
func (controller *Controller) GetFileToDownload(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFileToDownload", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	// Get filetodownloadDB in DB
	var filetodownloadDB orm.FileToDownloadDB
	if _, err := db.First(&filetodownloadDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var filetodownloadAPI orm.FileToDownloadAPI
	filetodownloadAPI.ID = filetodownloadDB.ID
	filetodownloadAPI.FileToDownloadPointersEncoding = filetodownloadDB.FileToDownloadPointersEncoding
	filetodownloadDB.CopyBasicFieldsToFileToDownload_WOP(&filetodownloadAPI.FileToDownload_WOP)

	c.JSON(http.StatusOK, filetodownloadAPI)
}

// UpdateFileToDownload
//
// swagger:route PATCH /filetodownloads/{ID} filetodownloads updateFileToDownload
//
// # Update a filetodownload
//
// Responses:
// default: genericError
//
//	200: filetodownloadDBResponse
func (controller *Controller) UpdateFileToDownload(c *gin.Context) {

	mutexFileToDownload.Lock()
	defer mutexFileToDownload.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFileToDownload", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	// Validate input
	var input orm.FileToDownloadAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var filetodownloadDB orm.FileToDownloadDB

	// fetch the filetodownload
	_, err := db.First(&filetodownloadDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	filetodownloadDB.CopyBasicFieldsFromFileToDownload_WOP(&input.FileToDownload_WOP)
	filetodownloadDB.FileToDownloadPointersEncoding = input.FileToDownloadPointersEncoding

	db, _ = db.Model(&filetodownloadDB)
	_, err = db.Updates(&filetodownloadDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	filetodownloadNew := new(models.FileToDownload)
	filetodownloadDB.CopyBasicFieldsToFileToDownload(filetodownloadNew)

	// redeem pointers
	filetodownloadDB.DecodePointers(backRepo, filetodownloadNew)

	// get stage instance from DB instance, and call callback function
	filetodownloadOld := backRepo.BackRepoFileToDownload.Map_FileToDownloadDBID_FileToDownloadPtr[filetodownloadDB.ID]
	if filetodownloadOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), filetodownloadOld, filetodownloadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the filetodownloadDB
	c.JSON(http.StatusOK, filetodownloadDB)
}

// DeleteFileToDownload
//
// swagger:route DELETE /filetodownloads/{ID} filetodownloads deleteFileToDownload
//
// # Delete a filetodownload
//
// default: genericError
//
//	200: filetodownloadDBResponse
func (controller *Controller) DeleteFileToDownload(c *gin.Context) {

	mutexFileToDownload.Lock()
	defer mutexFileToDownload.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFileToDownload", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/load/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoFileToDownload.GetDB()

	// Get model if exist
	var filetodownloadDB orm.FileToDownloadDB
	if _, err := db.First(&filetodownloadDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&filetodownloadDB)

	// get an instance (not staged) from DB instance, and call callback function
	filetodownloadDeleted := new(models.FileToDownload)
	filetodownloadDB.CopyBasicFieldsToFileToDownload(filetodownloadDeleted)

	// get stage instance from DB instance, and call callback function
	filetodownloadStaged := backRepo.BackRepoFileToDownload.Map_FileToDownloadDBID_FileToDownloadPtr[filetodownloadDB.ID]
	if filetodownloadStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), filetodownloadStaged, filetodownloadDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
