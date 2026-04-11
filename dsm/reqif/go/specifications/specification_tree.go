package specifications

import (
	"fmt"
	"log"
	"strings"

	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecificationsTreeStage()
	treeStage.Reset()
	stage := stager.GetStage()

	sliceOfSpecificationNodes := make([]*tree.Node, 0)
	nameNode := &tree.Node{
		Name:      "Specifications",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nameNode)

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS

	if specifications == nil {
		treeStage.Commit()
		return
	}

	// prepare one node per specification type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specificationType_node := make(map[*m.SPECIFICATION_TYPE]*tree.Node)
	map_specificationType_nbInstances := make(map[*m.SPECIFICATION_TYPE]int)
	for _, specificationType := range spectypes.SPECIFICATION_TYPE {
		nodeSpecificationType := &tree.Node{
			Name:       specificationType.Name,
			IsExpanded: true,
		}
		sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nodeSpecificationType)
		map_specificationType_node[specificationType] = nodeSpecificationType
	}

	selectedSpecification := GetSelectedSpecification(stage)
	for _, specification := range specifications.SPECIFICATION {

		specificationType, ok := stager.Map_id_SPECIFICATION_TYPE[specification.TYPE.SPECIFICATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPECIFICATION_TYPE_REF", specification.TYPE.SPECIFICATION_TYPE_REF,
				"unknown relation type")
		}

		isSelectedSpecification := specification == selectedSpecification
		specificationRendering := GetSpecificationRendering(stage, specification)

		specificationNode := &tree.Node{
			Name:              specification.Name,
			HasCheckboxButton: true,
			IsChecked:         isSelectedSpecification,
			IsExpanded:        specificationRendering.IsNodeExpanded,
			Impl: &ProxySpecification{
				stager:        stager,
				specification: specification,
			},
		}

		{
			button := &tree.Button{
				Name: "Show/Unshow Numbering",
				OnClick: func() {
					proxy := &toggleButtonProxy{
						stager:      stager,
						toggleValue: &specificationRendering.IsWithHeadingNumbering,
					}
					proxy.ButtonUpdated(nil, nil, nil)
				},
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
			}

			if !specificationRendering.IsWithHeadingNumbering {
				button.ToolTipText = "Show numbering"
				button.Icon = string(buttons.BUTTON_format_list_numbered)
			} else {
				button.ToolTipText = "Hide numbering"
				button.Icon = string(buttons.BUTTON_subject)
			}
			specificationNode.Buttons = append(specificationNode.Buttons, button)
		}

		markDownContent := "# *" + specification.Name + "*"
		map_specificationType_node[specificationType].Children = append(map_specificationType_node[specificationType].Children, specificationNode)
		map_specificationType_nbInstances[specificationType] = map_specificationType_nbInstances[specificationType] + 1

		{
			// specificationAttributeCategoryXHTML := &tree.Node{
			// 	Name:       "XHTML",
			// 	IsExpanded: true,
			// 	FontStyle:  tree.ITALIC,
			// }
			// specificationNode.Children = append(specificationNode.Children, specificationAttributeCategoryXHTML)
			if specification.VALUES != nil {
				for _, attribute := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
					// provide the type
					var attributeDefinition string
					if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
						attributeDefinition = datatype.LONG_NAME
					} else {
						log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF,
							"unknown ref")
					}

					enclosedText := attribute.THE_VALUE.EnclosedText

					enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:div>", " ")
					enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:div>", "\n")
					enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br >", "-")
					enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:br >", "\n")
					enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br />", "\n")
					nodeXHTMLAttribute := &tree.Node{
						Name: attributeDefinition + " : " + enclosedText,
					}
					specificationNode.Children = append(specificationNode.Children, nodeXHTMLAttribute)
				}
			}
		}
		{
			hierarchyParentNode := &tree.Node{
				Name:       "Hierarchy",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			// specificationNode.Children = append(specificationNode.Children, hierarchyParentNode)

			depth := 0 // depth of chapters

			for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
				processSpecHierarchy(
					stager,
					specHierarchy,
					hierarchyParentNode,
					depth,
					&markDownContent)
			}

			// log.Println(markDownContent)
		}

		specobjects.AddAttributeNodes(stager, specificationNode, specification)

	}

	// update the node with the number of instances
	for _, type_ := range spectypes.SPECIFICATION_TYPE {
		nbInstances := map_specificationType_nbInstances[type_]
		nodeSpecType := map_specificationType_node[type_]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			RootNodes: sliceOfSpecificationNodes,
		},
	)

	treeStage.Commit()
}
