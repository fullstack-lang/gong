// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"

	"github.com/fullstack-lang/gong/examples/simple/go/models"
	"github.com/fullstack-lang/gong/examples/simple/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// declaration in order to justify use of the models import
var __Bclass__dummysDeclaration__ models.Bclass

// An BclassID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBclass updateBclass deleteBclass
type BclassID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BclassInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBclass updateBclass
type BclassInput struct {
	// The Bclass to submit or modify
	// in: body
	Bclass *orm.BclassAPI
}

// GetBclasss
//
// swagger:route GET /bclasss bclasss getBclasss
//
// Get all bclasss
//
// Responses:
//    default: genericError
//        200: bclassDBsResponse
func GetBclasss(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var bclasss []orm.BclassDB
	query := db.Find(&bclasss)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// for each bclass, update fields from the database nullable fields
	for idx := range bclasss {
		bclass := &bclasss[idx]
		_ = bclass
		// insertion point for updating fields
		if bclass.Name_Data.Valid {
			bclass.Name = bclass.Name_Data.String
		}

		if bclass.Floatfield_Data.Valid {
			bclass.Floatfield = bclass.Floatfield_Data.Float64
		}

		if bclass.Intfield_Data.Valid {
			bclass.Intfield = int(bclass.Intfield_Data.Int64)
		}

	}

	c.JSON(http.StatusOK, bclasss)
}

// PostBclass
//
// swagger:route POST /bclasss bclasss postBclass
//
// Creates a bclass
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: bclassDBResponse
func PostBclass(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.BclassAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bclass
	bclassDB := orm.BclassDB{}
	bclassDB.BclassAPI = input
	// insertion point for nullable field set
	bclassDB.Name_Data.String = input.Name
	bclassDB.Name_Data.Valid = true

	bclassDB.Floatfield_Data.Float64 = input.Floatfield
	bclassDB.Floatfield_Data.Valid = true

	bclassDB.Intfield_Data.Int64 = int64(input.Intfield)
	bclassDB.Intfield_Data.Valid = true

	query := db.Create(&bclassDB)
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

	c.JSON(http.StatusOK, bclassDB)
}

// GetBclass
//
// swagger:route GET /bclasss/{ID} bclasss getBclass
//
// Gets the details for a bclass.
//
// Responses:
//    default: genericError
//        200: bclassDBResponse
func GetBclass(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get bclass in DB
	var bclass orm.BclassDB
	if err := db.First(&bclass, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// insertion point for fields value set from nullable fields
	if bclass.Name_Data.Valid {
		bclass.Name = bclass.Name_Data.String
	}

	if bclass.Floatfield_Data.Valid {
		bclass.Floatfield = bclass.Floatfield_Data.Float64
	}

	if bclass.Intfield_Data.Valid {
		bclass.Intfield = int(bclass.Intfield_Data.Int64)
	}

	c.JSON(http.StatusOK, bclass)
}

// UpdateBclass
//
// swagger:route PATCH /bclasss/{ID} bclasss updateBclass
//
// Update a bclass
//
// Responses:
//    default: genericError
//        200: bclassDBResponse
func UpdateBclass(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var bclassDB orm.BclassDB

	// fetch the bclass
	query := db.First(&bclassDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.BclassAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	// insertion point for nullable field set
	input.Name_Data.String = input.Name
	input.Name_Data.Valid = true

	input.Floatfield_Data.Float64 = input.Floatfield
	input.Floatfield_Data.Valid = true

	input.Intfield_Data.Int64 = int64(input.Intfield)
	input.Intfield_Data.Valid = true

	query = db.Model(&bclassDB).Updates(input)
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

	// return status OK with the marshalling of the the bclassDB
	c.JSON(http.StatusOK, bclassDB)
}

// DeleteBclass
//
// swagger:route DELETE /bclasss/{ID} bclasss deleteBclass
//
// Delete a bclass
//
// Responses:
//    default: genericError
func DeleteBclass(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var bclassDB orm.BclassDB
	if err := db.First(&bclassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bclassDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}