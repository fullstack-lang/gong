// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// declaration in order to justify use of the models import
var __GongdocStatus__dummysDeclaration__ models.GongdocStatus

// An GongdocStatusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongdocStatus updateGongdocStatus deleteGongdocStatus
type GongdocStatusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongdocStatusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongdocStatus updateGongdocStatus
type GongdocStatusInput struct {
	// The GongdocStatus to submit or modify
	// in: body
	GongdocStatus *orm.GongdocStatusAPI
}

// GetGongdocStatuss
//
// swagger:route GET /gongdocstatuss gongdocstatuss getGongdocStatuss
//
// Get all gongdocstatuss
//
// Responses:
//    default: genericError
//        200: gongdocstatusDBsResponse
func GetGongdocStatuss(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var gongdocstatuss []orm.GongdocStatusDB
	query := db.Find(&gongdocstatuss)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// for each gongdocstatus, update fields from the database nullable fields
	for idx := range gongdocstatuss {
		gongdocstatus := &gongdocstatuss[idx]
		_ = gongdocstatus
		// insertion point for updating fields
		if gongdocstatus.Name_Data.Valid {
			gongdocstatus.Name = gongdocstatus.Name_Data.String
		}

		if gongdocstatus.Status_Data.Valid {
			gongdocstatus.Status = models.GongdocCommandType(gongdocstatus.Status_Data.String)
		}

		if gongdocstatus.CommandCompletionDate_Data.Valid {
			gongdocstatus.CommandCompletionDate = gongdocstatus.CommandCompletionDate_Data.String
		}

	}

	c.JSON(http.StatusOK, gongdocstatuss)
}

// PostGongdocStatus
//
// swagger:route POST /gongdocstatuss gongdocstatuss postGongdocStatus
//
// Creates a gongdocstatus
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: gongdocstatusDBResponse
func PostGongdocStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.GongdocStatusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongdocstatus
	gongdocstatusDB := orm.GongdocStatusDB{}
	gongdocstatusDB.GongdocStatusAPI = input
	// insertion point for nullable field set
	gongdocstatusDB.Name_Data.String = input.Name
	gongdocstatusDB.Name_Data.Valid = true

	gongdocstatusDB.Status_Data.String = string(input.Status)
	gongdocstatusDB.Status_Data.Valid = true

	gongdocstatusDB.CommandCompletionDate_Data.String = input.CommandCompletionDate
	gongdocstatusDB.CommandCompletionDate_Data.Valid = true

	query := db.Create(&gongdocstatusDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gongdocstatusDB)
}

// GetGongdocStatus
//
// swagger:route GET /gongdocstatuss/{ID} gongdocstatuss getGongdocStatus
//
// Gets the details for a gongdocstatus.
//
// Responses:
//    default: genericError
//        200: gongdocstatusDBResponse
func GetGongdocStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get gongdocstatus in DB
	var gongdocstatus orm.GongdocStatusDB
	if err := db.First(&gongdocstatus, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// insertion point for fields value set from nullable fields
	if gongdocstatus.Name_Data.Valid {
		gongdocstatus.Name = gongdocstatus.Name_Data.String
	}

	if gongdocstatus.Status_Data.Valid {
		gongdocstatus.Status = models.GongdocCommandType(gongdocstatus.Status_Data.String)
	}

	if gongdocstatus.CommandCompletionDate_Data.Valid {
		gongdocstatus.CommandCompletionDate = gongdocstatus.CommandCompletionDate_Data.String
	}

	c.JSON(http.StatusOK, gongdocstatus)
}

// UpdateGongdocStatus
//
// swagger:route PATCH /gongdocstatuss/{ID} gongdocstatuss updateGongdocStatus
//
// Update a gongdocstatus
//
// Responses:
//    default: genericError
//        200: gongdocstatusDBResponse
func UpdateGongdocStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var gongdocstatusDB orm.GongdocStatusDB

	// fetch the gongdocstatus
	query := db.First(&gongdocstatusDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.GongdocStatusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	// insertion point for nullable field set
	input.Name_Data.String = input.Name
	input.Name_Data.Valid = true

	input.Status_Data.String = string(input.Status)
	input.Status_Data.Valid = true

	input.CommandCompletionDate_Data.String = input.CommandCompletionDate
	input.CommandCompletionDate_Data.Valid = true

	query = db.Model(&gongdocstatusDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	// return status OK with the marshalling of the the gongdocstatusDB
	c.JSON(http.StatusOK, gongdocstatusDB)
}

// DeleteGongdocStatus
//
// swagger:route DELETE /gongdocstatuss/{ID} gongdocstatuss deleteGongdocStatus
//
// Delete a gongdocstatus
//
// Responses:
//    default: genericError
func DeleteGongdocStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var gongdocstatusDB orm.GongdocStatusDB
	if err := db.First(&gongdocstatusDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongdocstatusDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}