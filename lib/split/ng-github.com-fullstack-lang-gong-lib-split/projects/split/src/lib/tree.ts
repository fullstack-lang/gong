// generated code - do not edit

import { TreeAPI } from './tree-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Tree {

	static GONGSTRUCT_NAME = "Tree"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	TreeName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyTreeToTreeAPI(tree: Tree, treeAPI: TreeAPI) {

	treeAPI.CreatedAt = tree.CreatedAt
	treeAPI.DeletedAt = tree.DeletedAt
	treeAPI.ID = tree.ID

	// insertion point for basic fields copy operations
	treeAPI.Name = tree.Name
	treeAPI.StackName = tree.StackName
	treeAPI.TreeName = tree.TreeName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
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
	tree.StackName = treeAPI.StackName
	tree.TreeName = treeAPI.TreeName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
