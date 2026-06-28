package spectypes

import (
	"fmt"
	"log"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type nodeWithRank struct {
	Node *tree.Node
	Rank int
}

func addAttributeNode[
	AttrDef m.AttributeDefinition,
	AttrDefRef m.AttributeDefinitionRef,
	DatatypeDef m.DatatypeDefinition](
	stager *m.Stager,

	attributeDefinitions []AttrDef, // the attributes definitions to display
	map_Id_AttributeDefinition map[string]AttrDef, // map to find attribute def from the ref
	map_AtttributeDefinition_Spec_nbInstance map[AttrDef]int, // number of use in the selected specification

	attributesDefinitionRefs map[AttrDefRef]struct{}, // the set of all reference to this kind of attribute definition
	map_Id_DatatypeDefinition map[string]DatatypeDef,

	collectedNodes *[]nodeWithRank,
) {
	if len(attributeDefinitions) > 0 {

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[AttrDef]int)
		var attributeDefinition AttrDef
		for x := range attributesDefinitionRefs {

			var ok bool
			attributeDefinition, ok = map_Id_AttributeDefinition[x.GetRef()]
			if !ok {
				log.Panic("x.GetRef()", x.GetRef(),
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attributeDefinition := range attributeDefinitions {
			attrDefRendering := GetAttrDefRendering(stager, attributeDefinition)

			var attributeType string
			if datatype, ok := map_Id_DatatypeDefinition[attributeDefinition.GetDatatypeDefinitionRef()]; ok {
				attributeType = datatype.GetLongName()
			} else {
				log.Panic("attributeDefinition: ",
					attributeDefinition.GetLongName(), " ref:",
					attributeDefinition.GetDatatypeDefinitionRef(), ": ",
					"unknown ref")
			}

			nbInstances := map_AtttributeDefinition_Spec_nbInstance[attributeDefinition]
			nodeAttribute := &tree.Node{
				Name: attributeDefinition.GetLongName() + ":" + attributeType +
					fmt.Sprintf(" (%d/%d) Rank(%d)",
						nbInstances, map_attributeDefinition_nbInstance[attributeDefinition],
						*attrDefRendering.GetRankPtr()),
			}
			rank := configureAndAddAttributeNode(
				stager,
				nodeAttribute,
				nbInstances,
				attributeDefinition.GetIsEditable(),
				attributeDefinition.GetLongName(),
				attributeDefinition,
			)

			*collectedNodes = append(*collectedNodes, nodeWithRank{
				Node: nodeAttribute,
				Rank: rank,
			})
		}
	}
}
