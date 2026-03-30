// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/reqif/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {
	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.Stage,
	stagedNode, frontNode *gongtree_models.Node,
) {
	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {
	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {
	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	// log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "ALTERNATIVE_ID" {
		updateProbeTable[*models.ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_BOOLEAN" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_DATE" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_ENUMERATION" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_INTEGER" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_REAL" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_STRING" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_XHTML" {
		updateProbeTable[*models.ATTRIBUTE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_BOOLEAN" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_DATE" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_ENUMERATION" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_INTEGER" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_REAL" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_STRING" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_XHTML" {
		updateProbeTable[*models.ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ALTERNATIVE_ID" {
		updateProbeTable[*models.A_ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_DATE_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_INTEGER_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_REAL_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_STRING_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_XHTML_REF" {
		updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_BOOLEAN" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_DATE" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_ENUMERATION" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_INTEGER" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_REAL" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_STRING" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML_1" {
		updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CHILDREN" {
		updateProbeTable[*models.A_CHILDREN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CORE_CONTENT" {
		updateProbeTable[*models.A_CORE_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPES" {
		updateProbeTable[*models.A_DATATYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_BOOLEAN_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_DATE_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_ENUMERATION_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_INTEGER_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_REAL_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_STRING_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_XHTML_REF" {
		updateProbeTable[*models.A_DATATYPE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_EDITABLE_ATTS" {
		updateProbeTable[*models.A_EDITABLE_ATTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ENUM_VALUE_REF" {
		updateProbeTable[*models.A_ENUM_VALUE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_OBJECT" {
		updateProbeTable[*models.A_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_PROPERTIES" {
		updateProbeTable[*models.A_PROPERTIES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_RELATION_GROUP_TYPE_REF" {
		updateProbeTable[*models.A_RELATION_GROUP_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE_1" {
		updateProbeTable[*models.A_SOURCE_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE_SPECIFICATION_1" {
		updateProbeTable[*models.A_SOURCE_SPECIFICATION_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATIONS" {
		updateProbeTable[*models.A_SPECIFICATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATION_TYPE_REF" {
		updateProbeTable[*models.A_SPECIFICATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFIED_VALUES" {
		updateProbeTable[*models.A_SPECIFIED_VALUES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_ATTRIBUTES" {
		updateProbeTable[*models.A_SPEC_ATTRIBUTES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECTS" {
		updateProbeTable[*models.A_SPEC_OBJECTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECT_TYPE_REF" {
		updateProbeTable[*models.A_SPEC_OBJECT_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATIONS" {
		updateProbeTable[*models.A_SPEC_RELATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_GROUPS" {
		updateProbeTable[*models.A_SPEC_RELATION_GROUPS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_REF" {
		updateProbeTable[*models.A_SPEC_RELATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_TYPE_REF" {
		updateProbeTable[*models.A_SPEC_RELATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_TYPES" {
		updateProbeTable[*models.A_SPEC_TYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_THE_HEADER" {
		updateProbeTable[*models.A_THE_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_TOOL_EXTENSIONS" {
		updateProbeTable[*models.A_TOOL_EXTENSIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_BOOLEAN" {
		updateProbeTable[*models.DATATYPE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_DATE" {
		updateProbeTable[*models.DATATYPE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_ENUMERATION" {
		updateProbeTable[*models.DATATYPE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_INTEGER" {
		updateProbeTable[*models.DATATYPE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_REAL" {
		updateProbeTable[*models.DATATYPE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_STRING" {
		updateProbeTable[*models.DATATYPE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_XHTML" {
		updateProbeTable[*models.DATATYPE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EMBEDDED_VALUE" {
		updateProbeTable[*models.EMBEDDED_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ENUM_VALUE" {
		updateProbeTable[*models.ENUM_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP" {
		updateProbeTable[*models.RELATION_GROUP](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP_TYPE" {
		updateProbeTable[*models.RELATION_GROUP_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF" {
		updateProbeTable[*models.REQ_IF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_CONTENT" {
		updateProbeTable[*models.REQ_IF_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_HEADER" {
		updateProbeTable[*models.REQ_IF_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_TOOL_EXTENSION" {
		updateProbeTable[*models.REQ_IF_TOOL_EXTENSION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION" {
		updateProbeTable[*models.SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION_TYPE" {
		updateProbeTable[*models.SPECIFICATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_HIERARCHY" {
		updateProbeTable[*models.SPEC_HIERARCHY](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT" {
		updateProbeTable[*models.SPEC_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT_TYPE" {
		updateProbeTable[*models.SPEC_OBJECT_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION" {
		updateProbeTable[*models.SPEC_RELATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION_TYPE" {
		updateProbeTable[*models.SPEC_RELATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "XHTML_CONTENT" {
		updateProbeTable[*models.XHTML_CONTENT](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()
}
