// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"

	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
	"github.com/fullstack-lang/gong/stacks/gongsim/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// declaration in order to justify use of the models import
var __Engine__dummysDeclaration__ models.Engine

// An EngineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEngine updateEngine deleteEngine
type EngineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EngineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEngine updateEngine
type EngineInput struct {
	// The Engine to submit or modify
	// in: body
	Engine *orm.EngineAPI
}

// GetEngines
//
// swagger:route GET /engines engines getEngines
//
// Get all engines
//
// Responses:
//    default: genericError
//        200: engineDBsResponse
func GetEngines(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var engines []orm.EngineDB
	query := db.Find(&engines)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// for each engine, update fields from the database nullable fields
	for idx := range engines {
		engine := &engines[idx]
		_ = engine
		// insertion point for updating fields
		if engine.Name_Data.Valid {
			engine.Name = engine.Name_Data.String
		}

		if engine.EndTime_Data.Valid {
			engine.EndTime = engine.EndTime_Data.String
		}

		if engine.CurrentTime_Data.Valid {
			engine.CurrentTime = engine.CurrentTime_Data.String
		}

		if engine.SecondsSinceStart_Data.Valid {
			engine.SecondsSinceStart = engine.SecondsSinceStart_Data.Float64
		}

		if engine.Fired_Data.Valid {
			engine.Fired = int(engine.Fired_Data.Int64)
		}

		if engine.ControlMode_Data.Valid {
			engine.ControlMode = models.ControlMode(engine.ControlMode_Data.String)
		}

		if engine.State_Data.Valid {
			engine.State = models.EngineState(engine.State_Data.String)
		}

		if engine.Speed_Data.Valid {
			engine.Speed = engine.Speed_Data.Float64
		}

	}

	c.JSON(http.StatusOK, engines)
}

// PostEngine
//
// swagger:route POST /engines engines postEngine
//
// Creates a engine
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: engineDBResponse
func PostEngine(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.EngineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create engine
	engineDB := orm.EngineDB{}
	engineDB.EngineAPI = input
	// insertion point for nullable field set
	engineDB.Name_Data.String = input.Name
	engineDB.Name_Data.Valid = true

	engineDB.EndTime_Data.String = input.EndTime
	engineDB.EndTime_Data.Valid = true

	engineDB.CurrentTime_Data.String = input.CurrentTime
	engineDB.CurrentTime_Data.Valid = true

	engineDB.SecondsSinceStart_Data.Float64 = input.SecondsSinceStart
	engineDB.SecondsSinceStart_Data.Valid = true

	engineDB.Fired_Data.Int64 = int64(input.Fired)
	engineDB.Fired_Data.Valid = true

	engineDB.ControlMode_Data.String = string(input.ControlMode)
	engineDB.ControlMode_Data.Valid = true

	engineDB.State_Data.String = string(input.State)
	engineDB.State_Data.Valid = true

	engineDB.Speed_Data.Float64 = input.Speed
	engineDB.Speed_Data.Valid = true

	query := db.Create(&engineDB)
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

	c.JSON(http.StatusOK, engineDB)
}

// GetEngine
//
// swagger:route GET /engines/{ID} engines getEngine
//
// Gets the details for a engine.
//
// Responses:
//    default: genericError
//        200: engineDBResponse
func GetEngine(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get engine in DB
	var engine orm.EngineDB
	if err := db.First(&engine, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// insertion point for fields value set from nullable fields
	if engine.Name_Data.Valid {
		engine.Name = engine.Name_Data.String
	}

	if engine.EndTime_Data.Valid {
		engine.EndTime = engine.EndTime_Data.String
	}

	if engine.CurrentTime_Data.Valid {
		engine.CurrentTime = engine.CurrentTime_Data.String
	}

	if engine.SecondsSinceStart_Data.Valid {
		engine.SecondsSinceStart = engine.SecondsSinceStart_Data.Float64
	}

	if engine.Fired_Data.Valid {
		engine.Fired = int(engine.Fired_Data.Int64)
	}

	if engine.ControlMode_Data.Valid {
		engine.ControlMode = models.ControlMode(engine.ControlMode_Data.String)
	}

	if engine.State_Data.Valid {
		engine.State = models.EngineState(engine.State_Data.String)
	}

	if engine.Speed_Data.Valid {
		engine.Speed = engine.Speed_Data.Float64
	}

	c.JSON(http.StatusOK, engine)
}

// UpdateEngine
//
// swagger:route PATCH /engines/{ID} engines updateEngine
//
// Update a engine
//
// Responses:
//    default: genericError
//        200: engineDBResponse
func UpdateEngine(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var engineDB orm.EngineDB

	// fetch the engine
	query := db.First(&engineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.EngineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	// insertion point for nullable field set
	input.Name_Data.String = input.Name
	input.Name_Data.Valid = true

	input.EndTime_Data.String = input.EndTime
	input.EndTime_Data.Valid = true

	input.CurrentTime_Data.String = input.CurrentTime
	input.CurrentTime_Data.Valid = true

	input.SecondsSinceStart_Data.Float64 = input.SecondsSinceStart
	input.SecondsSinceStart_Data.Valid = true

	input.Fired_Data.Int64 = int64(input.Fired)
	input.Fired_Data.Valid = true

	input.ControlMode_Data.String = string(input.ControlMode)
	input.ControlMode_Data.Valid = true

	input.State_Data.String = string(input.State)
	input.State_Data.Valid = true

	input.Speed_Data.Float64 = input.Speed
	input.Speed_Data.Valid = true

	query = db.Model(&engineDB).Updates(input)
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

	// return status OK with the marshalling of the the engineDB
	c.JSON(http.StatusOK, engineDB)
}

// DeleteEngine
//
// swagger:route DELETE /engines/{ID} engines deleteEngine
//
// Delete a engine
//
// Responses:
//    default: genericError
func DeleteEngine(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var engineDB orm.EngineDB
	if err := db.First(&engineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&engineDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}