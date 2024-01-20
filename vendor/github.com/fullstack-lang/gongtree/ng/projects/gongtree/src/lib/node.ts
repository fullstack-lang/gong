// generated code - do not edit

import { NodeDB } from './node-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { SVGIcon } from './svgicon'
import { Button } from './button'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Node {

	static GONGSTRUCT_NAME = "Node"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	FontStyle: string = ""
	BackgroundColor: string = ""
	IsExpanded: boolean = false
	HasCheckboxButton: boolean = false
	IsChecked: boolean = false
	IsCheckboxDisabled: boolean = false
	IsInEditMode: boolean = false
	IsNodeClickable: boolean = false
	IsWithPreceedingIcon: boolean = false
	PreceedingIcon: string = ""

	// insertion point for pointers and slices of pointers declarations
	PreceedingSVGIcon?: SVGIcon

	Children: Array<Node> = []
	Buttons: Array<Button> = []
}

export function CopyNodeToNodeDB(node: Node, nodeDB: NodeDB) {

	nodeDB.CreatedAt = node.CreatedAt
	nodeDB.DeletedAt = node.DeletedAt
	nodeDB.ID = node.ID
	
	// insertion point for basic fields copy operations
	nodeDB.Name = node.Name
	nodeDB.FontStyle = node.FontStyle
	nodeDB.BackgroundColor = node.BackgroundColor
	nodeDB.IsExpanded = node.IsExpanded
	nodeDB.HasCheckboxButton = node.HasCheckboxButton
	nodeDB.IsChecked = node.IsChecked
	nodeDB.IsCheckboxDisabled = node.IsCheckboxDisabled
	nodeDB.IsInEditMode = node.IsInEditMode
	nodeDB.IsNodeClickable = node.IsNodeClickable
	nodeDB.IsWithPreceedingIcon = node.IsWithPreceedingIcon
	nodeDB.PreceedingIcon = node.PreceedingIcon

	// insertion point for pointer fields encoding
    nodeDB.NodePointersEncoding.PreceedingSVGIconID.Valid = true
	if (node.PreceedingSVGIcon != undefined) {
		nodeDB.NodePointersEncoding.PreceedingSVGIconID.Int64 = node.PreceedingSVGIcon.ID  
    } else {
		nodeDB.NodePointersEncoding.PreceedingSVGIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	nodeDB.NodePointersEncoding.Children = []
    for (let _node of node.Children) {
		nodeDB.NodePointersEncoding.Children.push(_node.ID)
    }
	
	nodeDB.NodePointersEncoding.Buttons = []
    for (let _button of node.Buttons) {
		nodeDB.NodePointersEncoding.Buttons.push(_button.ID)
    }
	
}

export function CopyNodeDBToNode(nodeDB: NodeDB, node: Node, frontRepo: FrontRepo) {

	node.CreatedAt = nodeDB.CreatedAt
	node.DeletedAt = nodeDB.DeletedAt
	node.ID = nodeDB.ID
	
	// insertion point for basic fields copy operations
	node.Name = nodeDB.Name
	node.FontStyle = nodeDB.FontStyle
	node.BackgroundColor = nodeDB.BackgroundColor
	node.IsExpanded = nodeDB.IsExpanded
	node.HasCheckboxButton = nodeDB.HasCheckboxButton
	node.IsChecked = nodeDB.IsChecked
	node.IsCheckboxDisabled = nodeDB.IsCheckboxDisabled
	node.IsInEditMode = nodeDB.IsInEditMode
	node.IsNodeClickable = nodeDB.IsNodeClickable
	node.IsWithPreceedingIcon = nodeDB.IsWithPreceedingIcon
	node.PreceedingIcon = nodeDB.PreceedingIcon

	// insertion point for pointer fields encoding
	node.PreceedingSVGIcon = frontRepo.SVGIcons.get(nodeDB.NodePointersEncoding.PreceedingSVGIconID.Int64)

	// insertion point for slice of pointers fields encoding
	node.Children = new Array<Node>()
	for (let _id of nodeDB.NodePointersEncoding.Children) {
	  let _node = frontRepo.Nodes.get(_id)
	  if (_node != undefined) {
		node.Children.push(_node!)
	  }
	}
	node.Buttons = new Array<Button>()
	for (let _id of nodeDB.NodePointersEncoding.Buttons) {
	  let _button = frontRepo.Buttons.get(_id)
	  if (_button != undefined) {
		node.Buttons.push(_button!)
	  }
	}
}
