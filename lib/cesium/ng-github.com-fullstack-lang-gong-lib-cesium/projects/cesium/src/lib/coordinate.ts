// generated code - do not edit

import { CoordinateAPI } from './coordinate-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Coordinate {

	static GONGSTRUCT_NAME = "Coordinate"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Longitude: number = 0
	Latitude: number = 0
	Altitude: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCoordinateToCoordinateAPI(coordinate: Coordinate, coordinateAPI: CoordinateAPI) {

	coordinateAPI.CreatedAt = coordinate.CreatedAt
	coordinateAPI.DeletedAt = coordinate.DeletedAt
	coordinateAPI.ID = coordinate.ID

	// insertion point for basic fields copy operations
	coordinateAPI.Name = coordinate.Name
	coordinateAPI.Longitude = coordinate.Longitude
	coordinateAPI.Latitude = coordinate.Latitude
	coordinateAPI.Altitude = coordinate.Altitude

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCoordinateAPIToCoordinate update basic, pointers and slice of pointers fields of coordinate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of coordinateAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCoordinateAPIToCoordinate(coordinateAPI: CoordinateAPI, coordinate: Coordinate, frontRepo: FrontRepo) {

	coordinate.CreatedAt = coordinateAPI.CreatedAt
	coordinate.DeletedAt = coordinateAPI.DeletedAt
	coordinate.ID = coordinateAPI.ID

	// insertion point for basic fields copy operations
	coordinate.Name = coordinateAPI.Name
	coordinate.Longitude = coordinateAPI.Longitude
	coordinate.Latitude = coordinateAPI.Latitude
	coordinate.Altitude = coordinateAPI.Altitude

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
