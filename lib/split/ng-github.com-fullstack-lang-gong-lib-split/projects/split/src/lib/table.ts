// generated code - do not edit

import { TableAPI } from './table-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Table {

	static GONGSTRUCT_NAME = "Table"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	TableName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyTableToTableAPI(table: Table, tableAPI: TableAPI) {

	tableAPI.CreatedAt = table.CreatedAt
	tableAPI.DeletedAt = table.DeletedAt
	tableAPI.ID = table.ID

	// insertion point for basic fields copy operations
	tableAPI.Name = table.Name
	tableAPI.StackName = table.StackName
	tableAPI.TableName = table.TableName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyTableAPIToTable update basic, pointers and slice of pointers fields of table
// from respectively the basic fields and encoded fields of pointers and slices of pointers of tableAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTableAPIToTable(tableAPI: TableAPI, table: Table, frontRepo: FrontRepo) {

	table.CreatedAt = tableAPI.CreatedAt
	table.DeletedAt = tableAPI.DeletedAt
	table.ID = tableAPI.ID

	// insertion point for basic fields copy operations
	table.Name = tableAPI.Name
	table.StackName = tableAPI.StackName
	table.TableName = tableAPI.TableName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
