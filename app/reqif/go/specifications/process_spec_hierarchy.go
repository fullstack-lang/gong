package specifications

import (
	"log"
	"strings"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/app/reqif/go/spectypes"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func processSpecHierarchy(
	stager *m.Stager,
	specHierarchy *m.SPEC_HIERARCHY,
	hierarchyParentNode *tree.Node,
	outerDepth int,
	markDownContent *string,
) {
	stage := stager.GetStage()

	specObject, ok := stager.Map_id_SPEC_OBJECT[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}

	specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
	if !ok {
		log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
			"unknown ref")
	}

	markdownBoldStartingMark := `
**`
	markdownBoldEndingMark := `**
`

	specObjectTypeRendering := spectypes.GetSpecObjectTypeRendering(stage, specObjectType)

	if specObjectTypeRendering.IsHeading {

		*markDownContent += `
`

		if specHierarchy.CHILDREN == nil || len(specHierarchy.CHILDREN.SPEC_HIERARCHY) == 0 {
			*markDownContent += "#"
		}

		if true {
			for range outerDepth {
				*markDownContent += "#"
			}
			*markDownContent += " "
		}
	} else {
		*markDownContent += markdownBoldStartingMark
	}

	specobjects.AddIdentifierAndNameToTitle(stager, specObjectType, markDownContent, specObject)

	titleComplement := specobjects.FillUpStringWithAttributes(stager, specObject, specobjects.Title)

	subjectComplement := specobjects.FillUpStringWithAttributes(stager, specObject, specobjects.Subject)

	if !specObjectTypeRendering.IsHeading && !strings.HasSuffix(*markDownContent, "**") && titleComplement != "" {
		*markDownContent += " - "
	}

	*markDownContent += titleComplement

	if specObjectTypeRendering.IsHeading {
		*markDownContent += `
`
	} else {
		// ending mark for bold
		*markDownContent += markdownBoldEndingMark
	}

	// remove "****" if no title is present
	if strings.HasSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark) {
		*markDownContent = strings.TrimSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark)
	}

	// add the subject after the title
	if subjectComplement != "" {
		*markDownContent += `
` + subjectComplement + `
`
	}

	specObjectNode := &tree.Node{
		Name: specObject.Name + " : " + specObjectType.Name,
	}
	hierarchyParentNode.Children = append(hierarchyParentNode.Children, specObjectNode)

	specobjects.AddAttributeNodes(stager, specObjectNode, specObject)
	specobjects.AppendAttributesToMarkdown(stager, specObject, markDownContent)

	m.AddIconForEditabilityOfAttribute(specHierarchy.IS_EDITABLE, specObject.Name, specObjectNode)

	if specHierarchy.CHILDREN != nil {
		for _, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchyChildren,
				specObjectNode,
				outerDepth+1,
				markDownContent)
		}
	}
}

type ProxySpecification struct {
	stager        *m.Stager
	specification *m.SPECIFICATION
}

func (p *ProxySpecification) OnAfterUpdate(treeStage *tree.Stage, stageNode, frontNode *tree.Node) {
	stage := p.stager.GetStage()
	if frontNode.IsChecked && !stageNode.IsChecked {
		frontNode.IsChecked = stageNode.IsChecked

		// log.Println("Specification", proxy.specification.Name, "selected")
		SetSelectedSpecification(stage, p.specification)

		p.stager.Map_SPECIFICATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPECIFICATION_TYPE]()
		p.stager.Map_SPEC_RELATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPEC_RELATION_TYPE]()

		p.stager.GetSpecificationsTreeUpdater().UpdateAttributeDefinitionNb(p.stager)
		p.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(p.stager)
		p.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsTreeStage(p.stager)
		p.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(p.stager)
	}

	if !frontNode.IsChecked && stageNode.IsChecked {
		frontNode.IsChecked = stageNode.IsChecked
	}

	if frontNode.IsExpanded && !stageNode.IsExpanded {
		specificationRendering := GetSpecificationRendering(stage, p.specification)
		stageNode.IsExpanded = frontNode.IsExpanded
		specificationRendering.IsNodeExpanded = true
	}

	if !frontNode.IsExpanded && stageNode.IsExpanded {
		specificationRendering := GetSpecificationRendering(stage, p.specification)
		stageNode.IsExpanded = frontNode.IsExpanded
		specificationRendering.IsNodeExpanded = false
	}
}
