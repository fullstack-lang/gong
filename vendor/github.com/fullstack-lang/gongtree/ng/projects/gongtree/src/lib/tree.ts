// generated code - do not edit

import { TreeDB } from './tree-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Node } from './node'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Tree {

	static GONGSTRUCT_NAME = "Tree"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	RootNodes: Array<Node> = []
}

export function CopyTreeToTreeDB(tree: Tree, treeDB: TreeDB) {

	treeDB.CreatedAt = tree.CreatedAt
	treeDB.DeletedAt = tree.DeletedAt
	treeDB.ID = tree.ID

	// insertion point for basic fields copy operations
	treeDB.Name = tree.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	treeDB.TreePointersEncoding.RootNodes = []
	for (let _node of tree.RootNodes) {
		treeDB.TreePointersEncoding.RootNodes.push(_node.ID)
	}

}

// CopyTreeDBToTree update basic, pointers and slice of pointers fields of tree
// from respectively the basic fields and encoded fields of pointers and slices of pointers of treeDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTreeDBToTree(treeDB: TreeDB, tree: Tree, frontRepo: FrontRepo) {

	tree.CreatedAt = treeDB.CreatedAt
	tree.DeletedAt = treeDB.DeletedAt
	tree.ID = treeDB.ID

	// insertion point for basic fields copy operations
	tree.Name = treeDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	tree.RootNodes = new Array<Node>()
	for (let _id of treeDB.TreePointersEncoding.RootNodes) {
		let _node = frontRepo.map_ID_Node.get(_id)
		if (_node != undefined) {
			tree.RootNodes.push(_node!)
		}
	}
}
