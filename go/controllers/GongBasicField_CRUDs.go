// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongBasicField__dummysDeclaration__ models.GongBasicField
var __GongBasicField_time__dummyDeclaration time.Duration

// An GongBasicFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongBasicField updateGongBasicField deleteGongBasicField
type GongBasicFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongBasicFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongBasicField updateGongBasicField
type GongBasicFieldInput struct {
	// The GongBasicField to submit or modify
	// in: body
	GongBasicField *orm.GongBasicFieldAPI
}

// GetGongBasicFields
//
// swagger:route GET /gongbasicfields gongbasicfields getGongBasicFields
//
// Get all gongbasicfields
//
// Responses:
//    default: genericError
//        200: gongbasicfieldDBsResponse
func GetGongBasicFields(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongBasicField.GetDB()

	// source slice
	var gongbasicfieldDBs []orm.GongBasicFieldDB
	query := db.Find(&gongbasicfieldDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongbasicfieldAPIs := make([]orm.GongBasicFieldAPI, 0)

	// for each gongbasicfield, update fields from the database nullable fields
	for idx := range gongbasicfieldDBs {
		gongbasicfieldDB := &gongbasicfieldDBs[idx]
		_ = gongbasicfieldDB
		var gongbasicfieldAPI orm.GongBasicFieldAPI

		// insertion point for updating fields
		gongbasicfieldAPI.ID = gongbasicfieldDB.ID
		gongbasicfieldDB.CopyBasicFieldsToGongBasicField(&gongbasicfieldAPI.GongBasicField)
		gongbasicfieldAPI.GongBasicFieldPointersEnconding = gongbasicfieldDB.GongBasicFieldPointersEnconding
		gongbasicfieldAPIs = append(gongbasicfieldAPIs, gongbasicfieldAPI)
	}

	c.JSON(http.StatusOK, gongbasicfieldAPIs)
}

// PostGongBasicField
//
// swagger:route POST /gongbasicfields gongbasicfields postGongBasicField
//
// Creates a gongbasicfield
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: gongbasicfieldDBResponse
func PostGongBasicField(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongBasicField.GetDB()

	// Validate input
	var input orm.GongBasicFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongbasicfield
	gongbasicfieldDB := orm.GongBasicFieldDB{}
	gongbasicfieldDB.GongBasicFieldPointersEnconding = input.GongBasicFieldPointersEnconding
	gongbasicfieldDB.CopyBasicFieldsFromGongBasicField(&input.GongBasicField)

	query := db.Create(&gongbasicfieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongbasicfieldDB)
}

// GetGongBasicField
//
// swagger:route GET /gongbasicfields/{ID} gongbasicfields getGongBasicField
//
// Gets the details for a gongbasicfield.
//
// Responses:
//    default: genericError
//        200: gongbasicfieldDBResponse
func GetGongBasicField(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongBasicField.GetDB()

	// Get gongbasicfieldDB in DB
	var gongbasicfieldDB orm.GongBasicFieldDB
	if err := db.First(&gongbasicfieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongbasicfieldAPI orm.GongBasicFieldAPI
	gongbasicfieldAPI.ID = gongbasicfieldDB.ID
	gongbasicfieldAPI.GongBasicFieldPointersEnconding = gongbasicfieldDB.GongBasicFieldPointersEnconding
	gongbasicfieldDB.CopyBasicFieldsToGongBasicField(&gongbasicfieldAPI.GongBasicField)

	c.JSON(http.StatusOK, gongbasicfieldAPI)
}

// UpdateGongBasicField
//
// swagger:route PATCH /gongbasicfields/{ID} gongbasicfields updateGongBasicField
//
// Update a gongbasicfield
//
// Responses:
//    default: genericError
//        200: gongbasicfieldDBResponse
func UpdateGongBasicField(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongBasicField.GetDB()

	// Get model if exist
	var gongbasicfieldDB orm.GongBasicFieldDB

	// fetch the gongbasicfield
	query := db.First(&gongbasicfieldDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.GongBasicFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	gongbasicfieldDB.CopyBasicFieldsFromGongBasicField(&input.GongBasicField)
	gongbasicfieldDB.GongBasicFieldPointersEnconding = input.GongBasicFieldPointersEnconding

	query = db.Model(&gongbasicfieldDB).Updates(gongbasicfieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongbasicfieldDB
	c.JSON(http.StatusOK, gongbasicfieldDB)
}

// DeleteGongBasicField
//
// swagger:route DELETE /gongbasicfields/{ID} gongbasicfields deleteGongBasicField
//
// Delete a gongbasicfield
//
// Responses:
//    default: genericError
func DeleteGongBasicField(c *gin.Context) {
	db := orm.BackRepo.BackRepoGongBasicField.GetDB()

	// Get model if exist
	var gongbasicfieldDB orm.GongBasicFieldDB
	if err := db.First(&gongbasicfieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongbasicfieldDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
