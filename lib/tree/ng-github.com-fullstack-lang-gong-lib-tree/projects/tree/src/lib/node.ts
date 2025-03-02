// generated code - do not edit

import { NodeAPI } from './node-api'
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
	HasSecondCheckboxButton: boolean = false
	IsSecondCheckboxChecked: boolean = false
	IsSecondCheckboxDisabled: boolean = false
	TextAfterSecondCheckbox: string = ""
	IsInEditMode: boolean = false
	IsNodeClickable: boolean = false
	IsWithPreceedingIcon: boolean = false
	PreceedingIcon: string = ""

	// insertion point for pointers and slices of pointers declarations
	PreceedingSVGIcon?: SVGIcon

	Children: Array<Node> = []
	Buttons: Array<Button> = []
}

export function CopyNodeToNodeAPI(node: Node, nodeAPI: NodeAPI) {

	nodeAPI.CreatedAt = node.CreatedAt
	nodeAPI.DeletedAt = node.DeletedAt
	nodeAPI.ID = node.ID

	// insertion point for basic fields copy operations
	nodeAPI.Name = node.Name
	nodeAPI.FontStyle = node.FontStyle
	nodeAPI.BackgroundColor = node.BackgroundColor
	nodeAPI.IsExpanded = node.IsExpanded
	nodeAPI.HasCheckboxButton = node.HasCheckboxButton
	nodeAPI.IsChecked = node.IsChecked
	nodeAPI.IsCheckboxDisabled = node.IsCheckboxDisabled
	nodeAPI.HasSecondCheckboxButton = node.HasSecondCheckboxButton
	nodeAPI.IsSecondCheckboxChecked = node.IsSecondCheckboxChecked
	nodeAPI.IsSecondCheckboxDisabled = node.IsSecondCheckboxDisabled
	nodeAPI.TextAfterSecondCheckbox = node.TextAfterSecondCheckbox
	nodeAPI.IsInEditMode = node.IsInEditMode
	nodeAPI.IsNodeClickable = node.IsNodeClickable
	nodeAPI.IsWithPreceedingIcon = node.IsWithPreceedingIcon
	nodeAPI.PreceedingIcon = node.PreceedingIcon

	// insertion point for pointer fields encoding
	nodeAPI.NodePointersEncoding.PreceedingSVGIconID.Valid = true
	if (node.PreceedingSVGIcon != undefined) {
		nodeAPI.NodePointersEncoding.PreceedingSVGIconID.Int64 = node.PreceedingSVGIcon.ID  
	} else {
		nodeAPI.NodePointersEncoding.PreceedingSVGIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	nodeAPI.NodePointersEncoding.Children = []
	for (let _node of node.Children) {
		nodeAPI.NodePointersEncoding.Children.push(_node.ID)
	}

	nodeAPI.NodePointersEncoding.Buttons = []
	for (let _button of node.Buttons) {
		nodeAPI.NodePointersEncoding.Buttons.push(_button.ID)
	}

}

// CopyNodeAPIToNode update basic, pointers and slice of pointers fields of node
// from respectively the basic fields and encoded fields of pointers and slices of pointers of nodeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyNodeAPIToNode(nodeAPI: NodeAPI, node: Node, frontRepo: FrontRepo) {

	node.CreatedAt = nodeAPI.CreatedAt
	node.DeletedAt = nodeAPI.DeletedAt
	node.ID = nodeAPI.ID

	// insertion point for basic fields copy operations
	node.Name = nodeAPI.Name
	node.FontStyle = nodeAPI.FontStyle
	node.BackgroundColor = nodeAPI.BackgroundColor
	node.IsExpanded = nodeAPI.IsExpanded
	node.HasCheckboxButton = nodeAPI.HasCheckboxButton
	node.IsChecked = nodeAPI.IsChecked
	node.IsCheckboxDisabled = nodeAPI.IsCheckboxDisabled
	node.HasSecondCheckboxButton = nodeAPI.HasSecondCheckboxButton
	node.IsSecondCheckboxChecked = nodeAPI.IsSecondCheckboxChecked
	node.IsSecondCheckboxDisabled = nodeAPI.IsSecondCheckboxDisabled
	node.TextAfterSecondCheckbox = nodeAPI.TextAfterSecondCheckbox
	node.IsInEditMode = nodeAPI.IsInEditMode
	node.IsNodeClickable = nodeAPI.IsNodeClickable
	node.IsWithPreceedingIcon = nodeAPI.IsWithPreceedingIcon
	node.PreceedingIcon = nodeAPI.PreceedingIcon

	// insertion point for pointer fields encoding
	node.PreceedingSVGIcon = frontRepo.map_ID_SVGIcon.get(nodeAPI.NodePointersEncoding.PreceedingSVGIconID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(nodeAPI.NodePointersEncoding.Children)) {
		console.error('Rects is not an array:', nodeAPI.NodePointersEncoding.Children);
		return;
	}

	node.Children = new Array<Node>()
	for (let _id of nodeAPI.NodePointersEncoding.Children) {
		let _node = frontRepo.map_ID_Node.get(_id)
		if (_node != undefined) {
			node.Children.push(_node!)
		}
	}
	if (!Array.isArray(nodeAPI.NodePointersEncoding.Buttons)) {
		console.error('Rects is not an array:', nodeAPI.NodePointersEncoding.Buttons);
		return;
	}

	node.Buttons = new Array<Button>()
	for (let _id of nodeAPI.NodePointersEncoding.Buttons) {
		let _button = frontRepo.map_ID_Button.get(_id)
		if (_button != undefined) {
			node.Buttons.push(_button!)
		}
	}
}
