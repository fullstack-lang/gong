package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const LegacyDiagramUmarshalling = false

// DiagramPackage stores all diagrams related to a gong package
// swagger:model DiagramPackage
type DiagramPackage struct {
	Name string

	// Stage_ where the DiagamPackage lives
	Stage_ *Stage

	// Path to the "diagrams" directory
	Path string

	// GongModelPath is the package path, e.g. "fullstack-lang/gongxlsx/go/models"
	GongModelPath string

	// Classdiagrams store UML Classdiagrams
	Classdiagrams []*Classdiagram

	// SelectedClassdiagram is the diagram of interest
	SelectedClassdiagram *Classdiagram

	// pointer to the model package
	ModelPkg                     *gong_models.ModelPkg
	AbsolutePathToDiagramPackage string

	// swagger:ignore
	Map_Identifier_NbInstances map[string]int
}

func getTheDiagramPackage(stage *Stage) (diagramPackage *DiagramPackage) {

	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	for k := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		log.Fatalln("There should be at least one diagram package on the stage")
	}
	return
}
