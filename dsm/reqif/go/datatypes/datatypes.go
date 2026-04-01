package datatypes

import (
	"fmt"
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"
)

type DataTypeTreeStageUpdater struct {
}

func (dataTypeTreeStageUpdater *DataTypeTreeStageUpdater) UpdateAndCommitDataTypeTreeStage(stager *m.Stager) {

	stager.GetDataTypeTreeStage().Reset()

	datatypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.DATATYPES

	if datatypes == nil {
		return
	}

	rootNode := &tree.Node{
		Name:       "Data types",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "XHTML",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_XHTML]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_XHTML_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_XHTML[x.DATATYPE_DEFINITION_XHTML_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_XHTML_REF", x.DATATYPE_DEFINITION_XHTML_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_XHTML {

			datatypeCategory.Children = append(datatypeCategory.Children,
				&tree.Node{
					Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
				},
			)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "String",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_STRING]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_STRING_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_STRING[x.DATATYPE_DEFINITION_STRING_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_STRING_REF", x.DATATYPE_DEFINITION_STRING_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_STRING {
			datatypeCategory.Children = append(datatypeCategory.Children,
				&tree.Node{
					Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
				},
			)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Enumeration",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_ENUMERATION]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_ENUMERATION_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_ENUMERATION[x.DATATYPE_DEFINITION_ENUMERATION_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_ENUMERATION_REF", x.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_ENUMERATION {
			node := &tree.Node{
				Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)

			for _, enum := range datatype.SPECIFIED_VALUES.ENUM_VALUE {
				nodeEnum := &tree.Node{
					Name: enum.LONG_NAME,
				}
				node.Children = append(node.Children, nodeEnum)
			}
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Boolean",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_BOOLEAN]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_BOOLEAN_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_BOOLEAN[x.DATATYPE_DEFINITION_BOOLEAN_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_BOOLEAN_REF", x.DATATYPE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_BOOLEAN {
			node := &tree.Node{
				Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Integer",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_INTEGER]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_INTEGER_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_INTEGER[x.DATATYPE_DEFINITION_INTEGER_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_INTEGER_REF", x.DATATYPE_DEFINITION_INTEGER_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_INTEGER {
			node := &tree.Node{
				Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Date",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_DATE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_DATE_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_DATE[x.DATATYPE_DEFINITION_DATE_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_DATE_REF", x.DATATYPE_DEFINITION_DATE_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_DATE {
			node := &tree.Node{
				Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Real",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)

		// compute the number of time this datatype is used
		map_datatypeDefinition_nbInstance := make(map[*m.DATATYPE_DEFINITION_REAL]int)

		for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_REAL_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_DATATYPE_DEFINITION_REAL[x.DATATYPE_DEFINITION_REAL_REF]
			if !ok {
				log.Panic("x.DATATYPE_DEFINITION_REAL_REF", x.DATATYPE_DEFINITION_REAL_REF,
					"unknown ref")
			} else {
				map_datatypeDefinition_nbInstance[datatypeDefinition]++
			}
		}

		for _, datatype := range datatypes.DATATYPE_DEFINITION_REAL {
			node := &tree.Node{
				Name: datatype.LONG_NAME + fmt.Sprintf(" (%d)", map_datatypeDefinition_nbInstance[datatype]),
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	tree.StageBranch(stager.GetDataTypeTreeStage(),
		&tree.Tree{
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.GetDataTypeTreeStage().Commit()
}
