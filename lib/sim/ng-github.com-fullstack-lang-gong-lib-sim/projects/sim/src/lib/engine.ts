// generated code - do not edit

import { EngineAPI } from './engine-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Engine {

	static GONGSTRUCT_NAME = "Engine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EndTime: string = ""
	CurrentTime: string = ""
	DisplayFormat: string = ""
	SecondsSinceStart: number = 0
	Fired: number = 0
	ControlMode: string = ""
	State: string = ""
	Speed: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyEngineToEngineAPI(engine: Engine, engineAPI: EngineAPI) {

	engineAPI.CreatedAt = engine.CreatedAt
	engineAPI.DeletedAt = engine.DeletedAt
	engineAPI.ID = engine.ID

	// insertion point for basic fields copy operations
	engineAPI.Name = engine.Name
	engineAPI.EndTime = engine.EndTime
	engineAPI.CurrentTime = engine.CurrentTime
	engineAPI.DisplayFormat = engine.DisplayFormat
	engineAPI.SecondsSinceStart = engine.SecondsSinceStart
	engineAPI.Fired = engine.Fired
	engineAPI.ControlMode = engine.ControlMode
	engineAPI.State = engine.State
	engineAPI.Speed = engine.Speed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyEngineAPIToEngine update basic, pointers and slice of pointers fields of engine
// from respectively the basic fields and encoded fields of pointers and slices of pointers of engineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyEngineAPIToEngine(engineAPI: EngineAPI, engine: Engine, frontRepo: FrontRepo) {

	engine.CreatedAt = engineAPI.CreatedAt
	engine.DeletedAt = engineAPI.DeletedAt
	engine.ID = engineAPI.ID

	// insertion point for basic fields copy operations
	engine.Name = engineAPI.Name
	engine.EndTime = engineAPI.EndTime
	engine.CurrentTime = engineAPI.CurrentTime
	engine.DisplayFormat = engineAPI.DisplayFormat
	engine.SecondsSinceStart = engineAPI.SecondsSinceStart
	engine.Fired = engineAPI.Fired
	engine.ControlMode = engineAPI.ControlMode
	engine.State = engineAPI.State
	engine.Speed = engineAPI.Speed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
