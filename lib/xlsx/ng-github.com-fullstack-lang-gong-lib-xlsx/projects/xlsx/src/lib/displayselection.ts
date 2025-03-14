// generated code - do not edit

import { DisplaySelectionAPI } from './displayselection-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLFile } from './xlfile'
import { XLSheet } from './xlsheet'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplaySelection {

	static GONGSTRUCT_NAME = "DisplaySelection"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	XLFile?: XLFile

	XLSheet?: XLSheet

}

export function CopyDisplaySelectionToDisplaySelectionAPI(displayselection: DisplaySelection, displayselectionAPI: DisplaySelectionAPI) {

	displayselectionAPI.CreatedAt = displayselection.CreatedAt
	displayselectionAPI.DeletedAt = displayselection.DeletedAt
	displayselectionAPI.ID = displayselection.ID

	// insertion point for basic fields copy operations
	displayselectionAPI.Name = displayselection.Name

	// insertion point for pointer fields encoding
	displayselectionAPI.DisplaySelectionPointersEncoding.XLFileID.Valid = true
	if (displayselection.XLFile != undefined) {
		displayselectionAPI.DisplaySelectionPointersEncoding.XLFileID.Int64 = displayselection.XLFile.ID  
	} else {
		displayselectionAPI.DisplaySelectionPointersEncoding.XLFileID.Int64 = 0 		
	}

	displayselectionAPI.DisplaySelectionPointersEncoding.XLSheetID.Valid = true
	if (displayselection.XLSheet != undefined) {
		displayselectionAPI.DisplaySelectionPointersEncoding.XLSheetID.Int64 = displayselection.XLSheet.ID  
	} else {
		displayselectionAPI.DisplaySelectionPointersEncoding.XLSheetID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyDisplaySelectionAPIToDisplaySelection update basic, pointers and slice of pointers fields of displayselection
// from respectively the basic fields and encoded fields of pointers and slices of pointers of displayselectionAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDisplaySelectionAPIToDisplaySelection(displayselectionAPI: DisplaySelectionAPI, displayselection: DisplaySelection, frontRepo: FrontRepo) {

	displayselection.CreatedAt = displayselectionAPI.CreatedAt
	displayselection.DeletedAt = displayselectionAPI.DeletedAt
	displayselection.ID = displayselectionAPI.ID

	// insertion point for basic fields copy operations
	displayselection.Name = displayselectionAPI.Name

	// insertion point for pointer fields encoding
	displayselection.XLFile = frontRepo.map_ID_XLFile.get(displayselectionAPI.DisplaySelectionPointersEncoding.XLFileID.Int64)
	displayselection.XLSheet = frontRepo.map_ID_XLSheet.get(displayselectionAPI.DisplaySelectionPointersEncoding.XLSheetID.Int64)

	// insertion point for slice of pointers fields encoding
}
