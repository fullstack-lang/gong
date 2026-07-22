// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/probe"

	embeddedgo "github.com/fullstack-lang/gong/dsm/phylla/go"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	if stage.GetGongMarshallingMode() == models.GongMarshallingAppendCommit {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
	}

	// the ".go" is not provided
	filename := impl.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	packageName := impl.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/dsm/phylla/go/models", packageName)
}

type Level1Stack struct {
	Stage *models.Stage
	Probe *probe.Probe
	R     *gin.Engine
}

func NewLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (level1Stack *Level1Stack) {
	return NewLevel1StackDelta(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, false)
}

func NewLevel1StackDelta(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
	deltaMode bool,
) (level1Stack *Level1Stack) {

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)

	if deltaMode {
		stage.SetDeltaMode(true)
	}

	level1Stack.Stage = stage

	level1Stack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
		level1Stack.Probe = probe.NewProbe(
			level1Stack.R,
			embeddedgo.GoModelsDir,
			embeddedgo.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)

		stage.SetProbeIF(level1Stack.Probe)
	}

	if unmarshallFromCode != "" {
		err := models.ParseAstFile(stage, unmarshallFromCode, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = hook
	}

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.ArcNormalVectorShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ArcNormalVectorShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.AxesShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.BaseVectorShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.BaseVectorShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.CircleGridShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.EndArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.EndArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.EndHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.EndHalfwayArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.ExplanationTextShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GridPathShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbon](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbonEndShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbonStartShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurveRhombusGridShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthCurveRhombusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.GrowthVectorShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.InitialRhombusGridShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.InitialRhombusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Library](stage)
	models.SetOrchestratorOnAfterUpdate[models.MidArcVectorShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.MidArcVectorShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbon](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbonEndShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbonStartShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectory](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.PartiallyRotatedTorusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.PerpendicularVector](stage)
	models.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorGridHalfway](stage)
	models.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorHalfway](stage)
	models.SetOrchestratorOnAfterUpdate[models.Plant](stage)
	models.SetOrchestratorOnAfterUpdate[models.PlantCircumferenceShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.PlantDiagram](stage)
	models.SetOrchestratorOnAfterUpdate[models.Rendered3DShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.RhombusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.RhombusStuff](stage)
	models.SetOrchestratorOnAfterUpdate[models.RotatedRhombusGridShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.RotatedRhombusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedBottomTopStartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedBottomTopStartArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackGrowthCurveEndArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackGrowthCurveStartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackNormalVector](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackOfGrowthCurve](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackOfNormalVector](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbon](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbonEndShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbonStartShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DEndHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DRibbonEndShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DRibbonStartShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DStartHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackOfGrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackOfGrowthCurve2DRibbon](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackOfPartiallyRotatedTorusShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackOfRotatedGrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackOfRotatedGrowthCurve2DRibbon](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DEndArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DRibbonEndShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DRibbonStartShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DStartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StartArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.StartHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StartHalfwayArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopEndArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopEndArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopEndHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopEndHalfwayArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopGrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopMidArcVectorShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopMidArcVectorShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackGrowthCurve2DEndHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackGrowthCurve2DStartHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackOfGrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2D](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2DEndArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2DStartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStartArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStartArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStartHalfwayArcShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TopStartHalfwayArcShapeGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.TorusStackShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.VerticalTorusStackShape](stage)

	return
}
