package vase

import (
	"fmt"
	"math"
	"path/filepath"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func TestInspectBasePlateAndRing(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stageSet := models.NewStageSet("test_stageset")
	err := stageSet.ParseAstFile(stagePath, true)
	if err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stageSet.ComputeReverseMaps()
	stageSet.ComputeInstancesNb()
	stageSet.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stageSet.Stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	stager := models.NewStagerForTest(stageSet.Stage)
	stager.SetStageSet(stageSet)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(targetPlant)
	for d := range stageSet.Stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenStackOfCarvedOutVaseTrapezeRingsShape = false
			d.IsHiddenVaseTrapezeBasePlateShape = false
		}
	}
	threejsStage := threejs.NewStage("threejs")
	stager.SetThreejsStage(threejsStage)

	updater := NewThreeJSStageUpdater()
	updater.UpdateThreeJSStage(stager)

	threejsStage = stager.GetThreejsStage()
	for c := range threejsStage.Curves {
		minY, maxY := math.MaxFloat64, -math.MaxFloat64
		for _, p := range c.Points {
			if p.Y < minY { minY = p.Y }
			if p.Y > maxY { maxY = p.Y }
		}
		t.Logf("Curve %s: %d points, Y in [%.2f, %.2f]", c.Name, len(c.Points), minY, maxY)
	}
	for mesh := range threejsStage.Meshs {
		t.Logf("Found Mesh: %q (Visible: %v)", mesh.Name, mesh.BufferGeometry != nil)
		if mesh.BufferGeometry != nil {
			t.Logf("Mesh %s: %d vertices, %d faces", mesh.Name, len(mesh.BufferGeometry.Vertices), len(mesh.BufferGeometry.Faces))
			if len(mesh.BufferGeometry.Vertices) > 0 {
				minY := math.MaxFloat64
				maxY := -math.MaxFloat64
				minR := math.MaxFloat64
				maxR := -math.MaxFloat64
				for _, v := range mesh.BufferGeometry.Vertices {
					if v.Y < minY {
						minY = v.Y
					}
					if v.Y > maxY {
						maxY = v.Y
					}
					r := math.Hypot(v.X, v.Z)
					if r < minR {
						minR = r
					}
					if r > maxR {
						maxR = r
					}
				}
				t.Logf("   Y range: [%.2f, %.2f], R range: [%.2f, %.2f]", minY, maxY, minR, maxR)
			}
		}
	}

	// Print comparison at the junction between ring and base plate
	var ringMesh, basePlateMesh *threejs.Mesh
	for mesh := range threejsStage.Meshs {
		if mesh.Name == "Vase Trapeze Stack Of Carved Out Vase Trapeze Rings h0 Mesh" {
			ringMesh = mesh
		}
		if mesh.Name == "Vase Trapeze Vase Trapeze Base Plate Mesh" {
			basePlateMesh = mesh
		}
	}

	if ringMesh != nil && basePlateMesh != nil {
		fmt.Printf("Found ringMesh and basePlateMesh!\n")
		// In ringMesh, find all vertices with Y close to p2H (409.25)
		t.Logf("RingMesh vertices total: %d", len(ringMesh.BufferGeometry.Vertices))
		t.Logf("BasePlateMesh vertices total: %d", len(basePlateMesh.BufferGeometry.Vertices))

		// Check the exterior wall of ringMesh: that was "Bottom Wall Face"
		// In buildRingMesh:
		// addQuadStrip(cBottomA, cTopA, "Plane 1 Face", false) -> 0..2M-1
		// addQuadStrip(cTopB, cBottomB, "Plane 2 Face", false) -> 2M..4M-1
		// addQuadStrip(cTopA, cTopB, "Top Wall Face", false) -> 4M..6M-1
		// addQuadStrip(cBottomB, cBottomA, "Bottom Wall Face", false) -> 6M..8M-1
		// For Bottom Wall Face:
		// vA is on cBottomB (Plane 2), vB is on cBottomA (Plane 1)
		// So vA has index 6M + 2*i
		M := 720
		// Strip 1: Plane 1 Face (0 .. 2M-1)
		// Strip 2: Plane 2 Face (2M .. 4M-1)
		// Strip 3: Top Wall Face (4M .. 6M-1)
		// Strip 4: Bottom Wall Face (6M .. 8M-1)
		for sIdx, sName := range []string{"Plane 1 Face", "Plane 2 Face", "Top Wall Face", "Bottom Wall Face"} {
			minY, maxY := math.MaxFloat64, -math.MaxFloat64
			minR, maxR := math.MaxFloat64, -math.MaxFloat64
			for i := 2 * sIdx * M; i < 2*(sIdx+1)*M; i++ {
				v := ringMesh.BufferGeometry.Vertices[i]
				if v.Y < minY { minY = v.Y }
				if v.Y > maxY { maxY = v.Y }
				r := math.Hypot(v.X, v.Z)
				if r < minR { minR = r }
				if r > maxR { maxR = r }
			}
			t.Logf("Ring Strip %d (%s): Y in [%.2f, %.2f], R in [%.2f, %.2f]", sIdx+1, sName, minY, maxY, minR, maxR)
		}

		// Also check BasePlateMesh: Outer s=0, s=1, etc.
		for s := 0; s <= 3; s++ {
			minY, maxY := math.MaxFloat64, -math.MaxFloat64
			minR, maxR := math.MaxFloat64, -math.MaxFloat64
			for i := 0; i < M; i++ {
				targetName := fmt.Sprintf("Vase Trapeze Vase Trapeze Base Plate Outer s%d %d", s, i)
				if s == 0 {
					targetName = fmt.Sprintf("Vase Trapeze Vase Trapeze Base Plate Outer Top %d", i)
				}
				for _, v := range basePlateMesh.BufferGeometry.Vertices {
					if v.Name == targetName {
						if v.Y < minY { minY = v.Y }
						if v.Y > maxY { maxY = v.Y }
						r := math.Hypot(v.X, v.Z)
						if r < minR { minR = r }
						if r > maxR { maxR = r }
						break
					}
				}
			}
			t.Logf("Base Outer s=%d: Y in [%.2f, %.2f], R in [%.2f, %.2f]", s, minY, maxY, minR, maxR)
		}
		// Compare rTop and rBot across all i
		cTopB := threejsStage.Curves
		var curveTopP2, curveBotP2 *threejs.Curve
		for c := range cTopB {
			if c.Name == "Vase Trapeze Top Curve Plane 2" {
				curveTopP2 = c
			}
			if c.Name == "Vase Trapeze Bottom Curve Plane 2" {
				curveBotP2 = c
			}
		}
		if curveTopP2 != nil && curveBotP2 != nil {
			var botGreaterCount, topGreaterCount int
			for i := 0; i < len(curveTopP2.Points) && i < len(curveBotP2.Points); i++ {
				rTop := math.Hypot(curveTopP2.Points[i].X, curveTopP2.Points[i].Z)
				rBot := math.Hypot(curveBotP2.Points[i].X, curveBotP2.Points[i].Z)
				if rBot > rTop {
					botGreaterCount++
				} else {
					topGreaterCount++
				}
			}
		if curveBotP2 != nil {
			t.Logf("curveBotP2 (ringBottomP2) has %d points. Sample points:", len(curveBotP2.Points))
			for i := 0; i < len(curveBotP2.Points); i += 60 {
				p := curveBotP2.Points[i]
				r := math.Hypot(p.X, p.Z)
				thetaDeg := math.Atan2(p.Z, p.X) * 180.0 / math.Pi
				t.Logf("  i=%3d: theta=%6.1f deg, R=%6.2f, X=%6.2f, Y=%6.2f, Z=%6.2f", i, thetaDeg, r, p.X, p.Y, p.Z)
			}

			// Profile at i=60 (valley) in BasePlateMesh
			t.Logf("BasePlate vertices at i=60:")
			for s := 0; s <= 24; s++ {
				targetName := fmt.Sprintf("Vase Trapeze Vase Trapeze Base Plate Outer s%d 60", s)
				if s == 0 {
					targetName = "Vase Trapeze Vase Trapeze Base Plate Outer Top 60"
				} else if s == 24 {
					targetName = "Vase Trapeze Vase Trapeze Base Plate Outer Floor 60"
				}
				for _, v := range basePlateMesh.BufferGeometry.Vertices {
					if v.Name == targetName {
						r := math.Hypot(v.X, v.Z)
						t.Logf("  s=%2d: Y=%.2f, R=%.2f (diff from r0: %.2f)", s, v.Y, r, r-183.62)
						break
					}
				}
			}
		}
	}

		// Assert materials are opaque when transparency is 0
		if ringMesh.MeshPhysicalMaterial.Transparent {
			t.Errorf("expected ringMesh to not be transparent when transparency is 0")
		}
		if basePlateMesh.MeshPhysicalMaterial.Transparent {
			t.Errorf("expected basePlateMesh to not be transparent when transparency is 0")
		}

		// Verify that all M boundary vertices at Plane 2 between Ring outer wall and BasePlate outer top match exactly
		var maxDiff float64
		for i := 0; i < M; i++ {
			vRing := ringMesh.BufferGeometry.Vertices[6*M + 2*i]
			targetName := fmt.Sprintf("Vase Trapeze Vase Trapeze Base Plate Outer Top %d", i)
			var vBase *threejs.Vector3
			for _, v := range basePlateMesh.BufferGeometry.Vertices {
				if v.Name == targetName {
					vBase = v
					break
				}
			}
			if vBase == nil {
				t.Fatalf("missing vertex %s in basePlateMesh", targetName)
			}
			diff := math.Hypot(vRing.X-vBase.X, math.Hypot(vRing.Y-vBase.Y, vRing.Z-vBase.Z))
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		t.Logf("Max coordinate difference between Ring and BasePlate at Plane 2: %e", maxDiff)
		if maxDiff > 1e-4 {
			t.Errorf("expected Ring and BasePlate to share exact same curve at Plane 2, but max diff is %f", maxDiff)
		}
	}
}
