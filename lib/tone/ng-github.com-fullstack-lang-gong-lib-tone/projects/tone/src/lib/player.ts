// generated code - do not edit

import { PlayerAPI } from './player-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Player {

	static GONGSTRUCT_NAME = "Player"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Status: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyPlayerToPlayerAPI(player: Player, playerAPI: PlayerAPI) {

	playerAPI.CreatedAt = player.CreatedAt
	playerAPI.DeletedAt = player.DeletedAt
	playerAPI.ID = player.ID

	// insertion point for basic fields copy operations
	playerAPI.Name = player.Name
	playerAPI.Status = player.Status

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyPlayerAPIToPlayer update basic, pointers and slice of pointers fields of player
// from respectively the basic fields and encoded fields of pointers and slices of pointers of playerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPlayerAPIToPlayer(playerAPI: PlayerAPI, player: Player, frontRepo: FrontRepo) {

	player.CreatedAt = playerAPI.CreatedAt
	player.DeletedAt = playerAPI.DeletedAt
	player.ID = playerAPI.ID

	// insertion point for basic fields copy operations
	player.Name = playerAPI.Name
	player.Status = playerAPI.Status

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
