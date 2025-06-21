// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	gong "github.com/fullstack-lang/gong/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Stager struct {
	stage     *Stage
	treeStage *tree.Stage
	svgStage  *svg.Stage
	gongStage *gong.Stage
	formStage *table.Stage
	ssgStage  *ssg.Stage

	embeddedDiagrams bool

	map_GongstructShape_Rect map[*GongStructShape]*svg_models.Rect
	map_GongenumShape_Rect   map[*GongEnumShape]*svg_models.Rect
	map_NoteShape_Rect       map[*GongNoteShape]*svg_models.Rect
	map_Structname_Rect      map[string]*svg_models.Rect
	map_Fieldname_Link       map[string]*svg_models.Link

	// this is a map is the managed by the callee thread
	// to inform of the number of instance by gongstruct names
	// this map is managed by callee stage struct
	map_GongStructName_InstancesNb map[string]int
	showNbInstances                bool
}

func NewStager(
	r *gin.Engine,
	receivingAsSplitArea *split.AsSplitArea,
	stage *Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
	gongStage *gong.Stage,
	formStage *table.Stage,
	ssgStage *ssg.Stage,

	embeddedDiagrams bool,

	map_GongStructName_InstancesNb map[string]int,

) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.treeStage = treeStage
	stager.svgStage = svgStage
	stager.gongStage = gongStage
	stager.formStage = formStage
	stager.ssgStage = ssgStage

	stager.embeddedDiagrams = embeddedDiagrams

	stager.map_GongStructName_InstancesNb = map_GongStructName_InstancesNb

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	receivingAsSplitArea.AsSplit = &split.AsSplit{
		Name:      "Root As Split for doc2 receiving area",
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "AsSplitArea 50% for Slit (Tree & Svg)",
				ShowNameInHeader: false,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 25,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "doc2 Tree",
										ShowNameInHeader: false,
										Size:             66,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
											TreeName:  stager.stage.GetProbeTreeSidebarStageName(),
										},
									},
									{
										Name: "temporary form stack",
										Size: 34,
										Form: &split.Form{
											StackName: stager.formStage.GetName(),
											FormName:  stager.formStage.GetName(), // convention
										},
									},
								},
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
		if !embeddedDiagrams {
			diagramPackage = (&DiagramPackage{
				Name: fmt.Sprintf("Diagram Package created the %s", time.Now().Local().UTC().Format(time.RFC3339)),
			}).Stage(stage)
			stage.Commit()
		} else {
			return
		}
	}

	// refresh all notes body from the original gong note in the package models
	// because, note are not synchronized via the gopls renaming request
	//
	// if a can be traced, this is probably for a lack of diagram maintenance
	gongNotes := *gong.GetGongstructInstancesMap[gong.GongNote](stager.gongStage)
	for _, classdiagram := range diagramPackage.Classdiagrams {

		for _, gongNoteShape := range classdiagram.GongNoteShapes {

			gongNote, ok := gongNotes[IdentifierToGongStructName(gongNoteShape.Identifier)]

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
	stager.UpdateAndCommitFormStage()

	return
}
