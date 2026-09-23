// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/probe"

	embeddedgo "github.com/fullstack-lang/gong/dsm/phylla/go"

	"net/http"

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
	R     *http.ServeMux
}

func (stack *Level1Stack) Run(addr string) error {
	return split_static.RunServer(stack.R, addr)
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
		// "go/models/diagrams/diagrams.go", the path is "../../models/diagrams/diagrams.go"
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
		err := stage.ParseAstFile(unmarshallFromCode, true)

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
	stage.SetOrchestratorOnAfterUpdate[models.Angle0Shape]()
	stage.SetOrchestratorOnAfterUpdate[models.ArcNormalVectorShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ArcNormalVectorShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.AxesShape]()
	stage.SetOrchestratorOnAfterUpdate[models.BaseVectorShape]()
	stage.SetOrchestratorOnAfterUpdate[models.BaseVectorShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.BottomCurvePlane1Shape]()
	stage.SetOrchestratorOnAfterUpdate[models.BottomCurvePlane2Shape]()
	stage.SetOrchestratorOnAfterUpdate[models.ChosenP1P2PairShape]()
	stage.SetOrchestratorOnAfterUpdate[models.CircleGridShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Circumference3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Clock2DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.Clock3DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.ClockAbstract]()
	stage.SetOrchestratorOnAfterUpdate[models.ClockTopCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.CutLine3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EndArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EndArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.EndHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EndHalfwayArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.ExplanationTextShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Eye3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EyeCornersSampledPoints3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EyeSampledPoints3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EyeSeatBottomCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EyeStoolBottomCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.EyeVolume3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GridPathShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurveRhombusGridShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthCurveRhombusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.GrowthVectorShape]()
	stage.SetOrchestratorOnAfterUpdate[models.InitialRhombusGridShape]()
	stage.SetOrchestratorOnAfterUpdate[models.InitialRhombusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Key3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.KeyHole3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.KeyHoleShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Leaves3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Library]()
	stage.SetOrchestratorOnAfterUpdate[models.MidArcVectorShape]()
	stage.SetOrchestratorOnAfterUpdate[models.MidArcVectorShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.MusicAbstract]()
	stage.SetOrchestratorOnAfterUpdate[models.OriginalPoints3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ParastichyMCurves3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ParastichyNCurves3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectory]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP1CurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP1P2]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP1PointShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP2CurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryP2PointShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyGrowthCurve2DTrajectoryShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyRotatedSeatBottomCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyRotatedSeatTopCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PartiallyRotatedTorusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PerpendicularVector]()
	stage.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorGridHalfway]()
	stage.SetOrchestratorOnAfterUpdate[models.PerpendicularVectorHalfway]()
	stage.SetOrchestratorOnAfterUpdate[models.Plant2DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.Plant3DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.PlantAbstract]()
	stage.SetOrchestratorOnAfterUpdate[models.PlantCircumferenceShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PointsAndLines3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.PxShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Rendered3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.RhombusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.RhombusStuff]()
	stage.SetOrchestratorOnAfterUpdate[models.RotatedRhombusGridShape]()
	stage.SetOrchestratorOnAfterUpdate[models.RotatedRhombusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.RotatedSampledPoints3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.RotatedSeatAndLegs3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.SampledPoints3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Seat3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.SeatAndLegs3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.SeatBottomCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.SeatTopCurveShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedBottomTopStartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedBottomTopStartArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftPartiallyGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackGrowthCurveEndArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackGrowthCurveStartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackNormalVector]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackOfGrowthCurve]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedLeftStackOfNormalVector]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ShiftedRightGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DEndHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackGrowthCurve2DStartHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfGrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfGrowthCurve2DByGrowthVector]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfPartiallyRotatedTorusShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfRotatedGrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.StackOfRotatedGrowthCurve2DRibbon]()
	stage.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DEndArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DRibbonEndShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DRibbonStartShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StackRotatedGrowthCurve2DStartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StartArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.StartHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.StartHalfwayArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.StemCylinder3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Stool2DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.Stool3DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.StoolAbstract]()
	stage.SetOrchestratorOnAfterUpdate[models.TiledFloor3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopCurvePlane1Shape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopCurvePlane2Shape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopEndArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopEndArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.TopEndHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopEndHalfwayArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.TopGrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.TopMidArcVectorShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopMidArcVectorShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackGrowthCurve2DEndHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackGrowthCurve2DStartHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackOfGrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2D]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2DEndArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStackOfRotatedGrowthCurve2DStartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStartArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStartArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStartHalfwayArcShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TopStartHalfwayArcShapeGrid]()
	stage.SetOrchestratorOnAfterUpdate[models.Torus3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TorusEdge3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TorusStackShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TrapezeVolume3DShape]()
	stage.SetOrchestratorOnAfterUpdate[models.TubeVase3DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.TubeVaseAbstract]()
	stage.SetOrchestratorOnAfterUpdate[models.Vase2DDiagram]()
	stage.SetOrchestratorOnAfterUpdate[models.VerticalTorusStackShape]()
	stage.SetOrchestratorOnAfterUpdate[models.VolumeKey3DShape]()

	return
}
