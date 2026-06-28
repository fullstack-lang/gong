package specobjects

import (
	"fmt"
	"log"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecObjectsTreeStageUpdater struct {
}

func (o *SpecObjectsTreeStageUpdater) UpdateAndCommitSpecObjectsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecObjectsTreeStage()

	treeStage.Reset()

	sliceOfSpecObjectNodes := make([]*tree.Node, 0)
	nameNode := &tree.Node{
		Name:      "Spec Objects",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecObjectNodes = append(sliceOfSpecObjectNodes, nameNode)

	objects := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	// prepare one node per spec object type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	if spectypes == nil {
		treeStage.Commit()
		return
	}

	map_specObjectType_node := make(map[*m.SPEC_OBJECT_TYPE]*tree.Node)
	map_specObjectType_nbInstances := make(map[*m.SPEC_OBJECT_TYPE]int)
	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		nodeSpecObjectType := &tree.Node{
			Name: specObjectType.Name,
		}
		sliceOfSpecObjectNodes = append(sliceOfSpecObjectNodes, nodeSpecObjectType)
		map_specObjectType_node[specObjectType] = nodeSpecObjectType
	}

	for _, specObject := range objects.SPEC_OBJECT {

		specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
		if !ok {
			log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
				"unknown object type")
		}

		objectNode := &tree.Node{
			Name: specObject.Name,
		}
		map_specObjectType_node[specObjectType].Children =
			append(map_specObjectType_node[specObjectType].Children, objectNode)
		map_specObjectType_nbInstances[specObjectType] = map_specObjectType_nbInstances[specObjectType] + 1

		AddAttributeNodes(stager, objectNode, specObject)
	}

	// update the node with the number of instances
	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		nbInstances := map_specObjectType_nbInstances[specObjectType]
		nodeSpecType := map_specObjectType_node[specObjectType]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			RootNodes: sliceOfSpecObjectNodes,
		},
	)

	treeStage.Commit()
}
