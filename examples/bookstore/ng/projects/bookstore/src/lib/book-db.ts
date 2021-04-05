// insertion point for imports
import { AreaDB } from './area-db'
import { EditorDB } from './editor-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class BookDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Author?: string
	City?: string
	Year?: number
	Price?: number
	Recommanded?: string

	// insertion point for other declarations
	Area?: AreaDB
	AreaID?: NullInt64
	AreaName?: string

	Editor_BooksDBID?: NullInt64
	Editor_Books_reverse?: EditorDB

}
