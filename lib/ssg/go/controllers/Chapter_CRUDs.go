// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
	"github.com/fullstack-lang/gong/lib/ssg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Chapter__dummysDeclaration__ models.Chapter
var __Chapter_time__dummyDeclaration time.Duration

var mutexChapter sync.Mutex

// An ChapterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getChapter updateChapter deleteChapter
type ChapterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ChapterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postChapter updateChapter
type ChapterInput struct {
	// The Chapter to submit or modify
	// in: body
	Chapter *orm.ChapterAPI
}

// GetChapters
//
// swagger:route GET /chapters chapters getChapters
//
// # Get all chapters
//
// Responses:
// default: genericError
//
//	200: chapterDBResponse
func (controller *Controller) GetChapters(c *gin.Context) {

	// source slice
	var chapterDBs []orm.ChapterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetChapters", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoChapter.GetDB()

	_, err := db.Find(&chapterDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	chapterAPIs := make([]orm.ChapterAPI, 0)

	// for each chapter, update fields from the database nullable fields
	for idx := range chapterDBs {
		chapterDB := &chapterDBs[idx]
		_ = chapterDB
		var chapterAPI orm.ChapterAPI

		// insertion point for updating fields
		chapterAPI.ID = chapterDB.ID
		chapterDB.CopyBasicFieldsToChapter_WOP(&chapterAPI.Chapter_WOP)
		chapterAPI.ChapterPointersEncoding = chapterDB.ChapterPointersEncoding
		chapterAPIs = append(chapterAPIs, chapterAPI)
	}

	c.JSON(http.StatusOK, chapterAPIs)
}

// PostChapter
//
// swagger:route POST /chapters chapters postChapter
//
// Creates a chapter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostChapter(c *gin.Context) {

	mutexChapter.Lock()
	defer mutexChapter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostChapters", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoChapter.GetDB()

	// Validate input
	var input orm.ChapterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create chapter
	chapterDB := orm.ChapterDB{}
	chapterDB.ChapterPointersEncoding = input.ChapterPointersEncoding
	chapterDB.CopyBasicFieldsFromChapter_WOP(&input.Chapter_WOP)

	_, err = db.Create(&chapterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoChapter.CheckoutPhaseOneInstance(&chapterDB)
	chapter := backRepo.BackRepoChapter.Map_ChapterDBID_ChapterPtr[chapterDB.ID]

	if chapter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), chapter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, chapterDB)
}

// GetChapter
//
// swagger:route GET /chapters/{ID} chapters getChapter
//
// Gets the details for a chapter.
//
// Responses:
// default: genericError
//
//	200: chapterDBResponse
func (controller *Controller) GetChapter(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetChapter", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoChapter.GetDB()

	// Get chapterDB in DB
	var chapterDB orm.ChapterDB
	if _, err := db.First(&chapterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var chapterAPI orm.ChapterAPI
	chapterAPI.ID = chapterDB.ID
	chapterAPI.ChapterPointersEncoding = chapterDB.ChapterPointersEncoding
	chapterDB.CopyBasicFieldsToChapter_WOP(&chapterAPI.Chapter_WOP)

	c.JSON(http.StatusOK, chapterAPI)
}

// UpdateChapter
//
// swagger:route PATCH /chapters/{ID} chapters updateChapter
//
// # Update a chapter
//
// Responses:
// default: genericError
//
//	200: chapterDBResponse
func (controller *Controller) UpdateChapter(c *gin.Context) {

	mutexChapter.Lock()
	defer mutexChapter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateChapter", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoChapter.GetDB()

	// Validate input
	var input orm.ChapterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var chapterDB orm.ChapterDB

	// fetch the chapter
	_, err := db.First(&chapterDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	chapterDB.CopyBasicFieldsFromChapter_WOP(&input.Chapter_WOP)
	chapterDB.ChapterPointersEncoding = input.ChapterPointersEncoding

	db, _ = db.Model(&chapterDB)
	_, err = db.Updates(&chapterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	chapterNew := new(models.Chapter)
	chapterDB.CopyBasicFieldsToChapter(chapterNew)

	// redeem pointers
	chapterDB.DecodePointers(backRepo, chapterNew)

	// get stage instance from DB instance, and call callback function
	chapterOld := backRepo.BackRepoChapter.Map_ChapterDBID_ChapterPtr[chapterDB.ID]
	if chapterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), chapterOld, chapterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the chapterDB
	c.JSON(http.StatusOK, chapterDB)
}

// DeleteChapter
//
// swagger:route DELETE /chapters/{ID} chapters deleteChapter
//
// # Delete a chapter
//
// default: genericError
//
//	200: chapterDBResponse
func (controller *Controller) DeleteChapter(c *gin.Context) {

	mutexChapter.Lock()
	defer mutexChapter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteChapter", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/ssg/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	db := backRepo.BackRepoChapter.GetDB()

	// Get model if exist
	var chapterDB orm.ChapterDB
	if _, err := db.First(&chapterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&chapterDB)

	// get an instance (not staged) from DB instance, and call callback function
	chapterDeleted := new(models.Chapter)
	chapterDB.CopyBasicFieldsToChapter(chapterDeleted)

	// get stage instance from DB instance, and call callback function
	chapterStaged := backRepo.BackRepoChapter.Map_ChapterDBID_ChapterPtr[chapterDB.ID]
	if chapterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), chapterStaged, chapterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
