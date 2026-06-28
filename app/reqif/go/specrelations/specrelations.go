package specrelations

import (
	"fmt"
	"log"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/specobjects"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecRelationsTreeStageUpdater struct {
}

func (specRelationsTreeUpdater *SpecRelationsTreeStageUpdater) UpdateAndCommitSpecRelationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecRelationsTreeStage()
	treeStage.Reset()

	sliceOfSpecRelationNodes := make([]*tree.Node, 0)
	specRelationsNode := &tree.Node{
		Name:      "Spec Relations",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecRelationNodes = append(sliceOfSpecRelationNodes, specRelationsNode)

	relations := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS

	// prepare one node per spec relation type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	if spectypes == nil {
		treeStage.Commit()
		return
	}

	map_specRelationType_node := make(map[*m.SPEC_RELATION_TYPE]*tree.Node)
	map_specRelationType_nbInstances := make(map[*m.SPEC_RELATION_TYPE]int)
	for _, specRelationType := range spectypes.SPEC_RELATION_TYPE {
		nodeSpecRelationType := &tree.Node{
			Name: specRelationType.Name,
		}
		sliceOfSpecRelationNodes = append(sliceOfSpecRelationNodes, nodeSpecRelationType)
		map_specRelationType_node[specRelationType] = nodeSpecRelationType
	}

	if relations != nil {
		for _, specRelation := range relations.SPEC_RELATION {
			relationNode := specRelationsTreeUpdater.addRelationNode(stager, specRelation, map_specRelationType_node, map_specRelationType_nbInstances)
			specobjects.AddAttributeNodes(stager, relationNode, specRelation)
		}
	}

	// update the node with the number of instances
	for _, specRelationType := range spectypes.SPEC_RELATION_TYPE {
		nbInstances := map_specRelationType_nbInstances[specRelationType]
		nodeSpecType := map_specRelationType_node[specRelationType]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	specRelationGroupsNode := &tree.Node{
		Name:      "Spec Relations",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecRelationNodes = append(sliceOfSpecRelationNodes, specRelationGroupsNode)

	relationGroups := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATION_GROUPS

	if relationGroups != nil {
		for _, specRelationGroup := range relationGroups.RELATION_GROUP {

			specRelationGroupNode := &tree.Node{
				Name:      specRelationGroup.GetName(),
				FontStyle: tree.ITALIC,
			}
			specRelationGroupsNode.Children = append(specRelationGroupsNode.Children, specRelationGroupNode)

			// for _, specRelation := range specRelationGroup.SPEC_RELATIONS {

			// 	relationNode := specRelationsTreeUpdater.addRelationNode(stager, specRelation, map_specRelationType_node, map_specRelationType_nbInstances)

			// 	specobjects.AddAttributeNodes(stager, relationNode, specRelation)
			// }
		}
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			RootNodes: sliceOfSpecRelationNodes,
		},
	)

	treeStage.Commit()
}

func (*SpecRelationsTreeStageUpdater) addRelationNode(
	stager *m.Stager,
	specRelation *m.SPEC_RELATION,
	map_specRelationType_node map[*m.SPEC_RELATION_TYPE]*tree.Node,
	map_specRelationType_nbInstances map[*m.SPEC_RELATION_TYPE]int) *tree.Node {
	specRelationType, ok := stager.Map_id_SPEC_RELATION_TYPE[specRelation.TYPE.SPEC_RELATION_TYPE_REF]
	if !ok {
		log.Panic("specRelation.TYPE.SPEC_RELATION_TYPE_REF", specRelation.TYPE.SPEC_RELATION_TYPE_REF,
			"unknown relation type")
	}

	relationNode := &tree.Node{
		Name: specRelation.Name,
	}
	map_specRelationType_node[specRelationType].Children =
		append(map_specRelationType_node[specRelationType].Children, relationNode)
	map_specRelationType_nbInstances[specRelationType] = map_specRelationType_nbInstances[specRelationType] + 1

	{
		specObject, ok := stager.Map_id_SPEC_OBJECT[specRelation.SOURCE.SPEC_OBJECT_REF]
		if !ok {
			log.Println("specRelation.SOURCE.SPEC_OBJECT_REF", specRelation.SOURCE.SPEC_OBJECT_REF,
				"unknown ref")
			node := &tree.Node{
				Name:       "Source unkown ref : " + specRelation.TARGET.SPEC_OBJECT_REF,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, node)
		} else {
			specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
					"unknown object type")
			}
			sourceNode := &tree.Node{
				Name:       "Source: " + specObject.Name + " : " + specObjectType.Name,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, sourceNode)

			specobjects.AddAttributeNodes(stager, sourceNode, specObject)
		}
	}

	{
		specObject, ok := stager.Map_id_SPEC_OBJECT[specRelation.TARGET.SPEC_OBJECT_REF]
		if !ok {
			log.Println("specRelation.TARGET.SPEC_OBJECT_REF", specRelation.TARGET.SPEC_OBJECT_REF,
				"unknown ref")
			node := &tree.Node{
				Name:       "Target unkown ref : " + specRelation.TARGET.SPEC_OBJECT_REF,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, node)
		} else {
			specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Println("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
					"unknown object type")
			}
			targetNode := &tree.Node{
				Name:       "Target: " + specObject.Name + " : " + specObjectType.Name,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, targetNode)

			specobjects.AddAttributeNodes(stager, targetNode, specObject)
		}
	}
	return relationNode
}
