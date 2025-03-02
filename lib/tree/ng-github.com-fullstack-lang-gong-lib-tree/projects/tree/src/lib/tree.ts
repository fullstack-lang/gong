// generated code - do not edit

import { TreeAPI } from './tree-api'
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

export function CopyTreeToTreeAPI(tree: Tree, treeAPI: TreeAPI) {

	treeAPI.CreatedAt = tree.CreatedAt
	treeAPI.DeletedAt = tree.DeletedAt
	treeAPI.ID = tree.ID

	// insertion point for basic fields copy operations
	treeAPI.Name = tree.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	treeAPI.TreePointersEncoding.RootNodes = []
	for (let _node of tree.RootNodes) {
		treeAPI.TreePointersEncoding.RootNodes.push(_node.ID)
	}

}

// CopyTreeAPIToTree update basic, pointers and slice of pointers fields of tree
// from respectively the basic fields and encoded fields of pointers and slices of pointers of treeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTreeAPIToTree(treeAPI: TreeAPI, tree: Tree, frontRepo: FrontRepo) {

	tree.CreatedAt = treeAPI.CreatedAt
	tree.DeletedAt = treeAPI.DeletedAt
	tree.ID = treeAPI.ID

	// insertion point for basic fields copy operations
	tree.Name = treeAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(treeAPI.TreePointersEncoding.RootNodes)) {
		console.error('Rects is not an array:', treeAPI.TreePointersEncoding.RootNodes);
		return;
	}

	tree.RootNodes = new Array<Node>()
	for (let _id of treeAPI.TreePointersEncoding.RootNodes) {
		let _node = frontRepo.map_ID_Node.get(_id)
		if (_node != undefined) {
			tree.RootNodes.push(_node!)
		}
	}
}
