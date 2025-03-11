// generated code - do not edit

import { CommandAPI } from './command-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Engine } from './engine'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Command {

	static GONGSTRUCT_NAME = "Command"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Command: string = ""
	CommandDate: string = ""

	// insertion point for pointers and slices of pointers declarations
	Engine?: Engine

}

export function CopyCommandToCommandAPI(command: Command, commandAPI: CommandAPI) {

	commandAPI.CreatedAt = command.CreatedAt
	commandAPI.DeletedAt = command.DeletedAt
	commandAPI.ID = command.ID

	// insertion point for basic fields copy operations
	commandAPI.Name = command.Name
	commandAPI.Command = command.Command
	commandAPI.CommandDate = command.CommandDate

	// insertion point for pointer fields encoding
	commandAPI.CommandPointersEncoding.EngineID.Valid = true
	if (command.Engine != undefined) {
		commandAPI.CommandPointersEncoding.EngineID.Int64 = command.Engine.ID  
	} else {
		commandAPI.CommandPointersEncoding.EngineID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyCommandAPIToCommand update basic, pointers and slice of pointers fields of command
// from respectively the basic fields and encoded fields of pointers and slices of pointers of commandAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCommandAPIToCommand(commandAPI: CommandAPI, command: Command, frontRepo: FrontRepo) {

	command.CreatedAt = commandAPI.CreatedAt
	command.DeletedAt = commandAPI.DeletedAt
	command.ID = commandAPI.ID

	// insertion point for basic fields copy operations
	command.Name = commandAPI.Name
	command.Command = commandAPI.Command
	command.CommandDate = commandAPI.CommandDate

	// insertion point for pointer fields encoding
	command.Engine = frontRepo.map_ID_Engine.get(commandAPI.CommandPointersEncoding.EngineID.Int64)

	// insertion point for slice of pointers fields encoding
}
