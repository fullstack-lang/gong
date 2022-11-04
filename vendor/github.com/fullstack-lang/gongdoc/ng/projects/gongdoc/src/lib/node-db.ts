// insertion point for imports
import { ClassdiagramDB } from './classdiagram-db'
import { UmlscDB } from './umlsc-db'
import { TreeDB } from './tree-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NodeDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: string = ""
	IsExpanded: boolean = false
	HasCheckboxButton: boolean = false
	IsChecked: boolean = false
	IsCheckboxDisabled: boolean = false
	HasAddChildButton: boolean = false
	HasEditButton: boolean = false
	IsInEditMode: boolean = false
	HasDrawButton: boolean = false
	HasDrawOffButton: boolean = false
	IsInDrawMode: boolean = false
	IsSaved: boolean = false
	HasDeleteButton: boolean = false

	// insertion point for other declarations
	Classdiagram?: ClassdiagramDB
	ClassdiagramID: NullInt64 = new NullInt64 // if pointer is null, Classdiagram.ID = 0

	Umlsc?: UmlscDB
	UmlscID: NullInt64 = new NullInt64 // if pointer is null, Umlsc.ID = 0

	Children?: Array<NodeDB>
	Node_ChildrenDBID: NullInt64 = new NullInt64
	Node_ChildrenDBID_Index: NullInt64  = new NullInt64 // store the index of the node instance in Node.Children
	Node_Children_reverse?: NodeDB 

	Tree_RootNodesDBID: NullInt64 = new NullInt64
	Tree_RootNodesDBID_Index: NullInt64  = new NullInt64 // store the index of the node instance in Tree.RootNodes
	Tree_RootNodes_reverse?: TreeDB 

}
