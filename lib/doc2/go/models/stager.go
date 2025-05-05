// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"

	gong "github.com/fullstack-lang/gong/go/models"
)

type Stager struct {
	stage     *Stage
	treeStage *tree.Stage
	svgStage  *svg.Stage
	gongStage *gong.Stage

	embeddedDiagrams bool
}

func NewStager(
	r *gin.Engine,
	receivingAsSplitArea *split.AsSplitArea,
	stage *Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
	gongStage *gong.Stage,

	embeddedDiagrams bool,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.treeStage = treeStage
	stager.svgStage = svgStage
	stager.gongStage = gongStage

	stager.embeddedDiagrams = embeddedDiagrams

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	receivingAsSplitArea.AsSplit = &split.AsSplit{
		Name:      "Root As Split for doc2 receiving area",
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "AsSplitArea 50% for Slit (Tree & Svg)",
				ShowNameInHeader: false,
				Size:             50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "doc2 Tree",
							ShowNameInHeader: false,
							Size:             25,
							Tree: &split.Tree{
								StackName: stager.treeStage.GetName(),
								TreeName:  stager.stage.GetProbeTreeSidebarStageName(),
							},
						},
						{
							Name:             "doc2 SVG",
							ShowNameInHeader: false,
							Size:             75,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				}),
			},
		},
	}

	// if no diagram package is present, creates one
	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	var diagramPackage *DiagramPackage
	for k := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		diagramPackage = (&DiagramPackage{
			Name: fmt.Sprintf("Diagram Package created the %s", time.Now().Local().UTC().Format(time.RFC3339)),
		}).Stage(stage)
		stage.Commit()
	}

	// refresh all notes body from the original gong note in the package models
	// because, note are not synchronized via the gopls renaming request
	//
	// if a can be traced, this is probably for a lack of diagram maintenance
	gongNotes := *gong.GetGongstructInstancesMap[gong.GongNote](stager.gongStage)
	for _, classdiagram := range diagramPackage.Classdiagrams {

		for _, gongNoteShape := range classdiagram.GongNoteShapes {

			gongNote, ok := gongNotes[IdentifierToGongObjectName(gongNoteShape.Identifier)]

			if !ok {
				log.Println("UnmarshallOneDiagram: In diagram", classdiagram.Name, "unknown note related to note shape", gongNoteShape.Identifier)
				gongNoteShape.Unstage(stage)

				if contains(classdiagram.GongNoteShapes, gongNoteShape) {
					classdiagram.GongNoteShapes = remove(classdiagram.GongNoteShapes, gongNoteShape)
				}
				continue
			}

			gongNoteShape.Body = gongNote.Body
			gongNoteShape.BodyHTML = gongNote.BodyHTML
		}

	}

	stager.UpdateAndCommitSVGStage()
	stager.UpdateAndCommitTreeStage()

	return
}
