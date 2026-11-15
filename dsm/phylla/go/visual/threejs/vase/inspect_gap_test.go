package vase

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func TestInspectGap(t *testing.T) {
	runCase := func(label string, n, m, radReps int, rotRatio, botPlane, relThick, curv, carved float64, constThick, hideBase bool) {
		stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
		stage := models.NewStage("inspect_" + label)
		if err := stage.ParseAstFile(stagePath, true); err != nil {
			t.Fatalf("parse: %v", err)
		}
		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()

		var plant *models.PlantAbstract
		for p := range stage.PlantAbstracts {
			if p.Name == "Vase Trapeze" {
				p.IsSelected = true
				plant = p
			} else {
				p.IsSelected = false
			}
		}
		plant.N = n
		plant.M = m
		plant.StackHeight = 3
		plant.RhombusInsideAngle = 130.0
		plant.RhombusSideLength = 120.0

		vase := plant.TubeVaseAbstract
		vase.Z_Ribbon = 7.57
		vase.RibbonVerticalScale = 6.33
		vase.TopPlaneHeight = 0.59
		vase.BottomPlaneHeight = botPlane
		vase.BaseRingBottomPlaneHeight = 0.25
		vase.ProjectionAngle = -14.9
		vase.BasePlateHeight = 0.26
		vase.CarvedOutTopRingsParameter = carved
		vase.BulbousStartTangentMagnitude = 7.79
		vase.BulbousEndAngle = 50.0
		vase.BulbousEndTangentMagnitude = 3.95
		vase.HasConstantThickness = constThick
		vase.HasRingBottom = true
		vase.HasCurvedWall = true
		vase.WallCurvature = curv
		vase.HasLip = false
		vase.LipHeight = 0.45
		vase.LipWidth = 0.12
		vase.RelativeVerticalThickness = relThick
		vase.RelativeRadialThickness = 0.01
		vase.RelativeRotatedTorusSeparation = 0.0
		vase.RotationRatio = rotRatio
		vase.RadialRepetitions = radReps

		var diag *models.TubeVase3DDiagram
		for d := range stage.TubeVase3DDiagrams {
			if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
				d.IsChecked = true
				d.IsHiddenTiledFloor3DShape = true
				d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
				d.IsHiddenVaseTrapezeBasePlateShape = hideBase
				diag = d
			}
		}

		stager := models.NewStagerForTest(stage)
		stager.EnforceSemanticForTest()
		stager.SetCurrentPlant(plant)
		threejsStage := threejs.NewStage("threejs_" + label)
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		meshes := make(map[string]*threejs.Mesh)
		for mesh := range threejsStage.Meshs {
			meshes[mesh.Name] = mesh
		}
		baseMesh := meshes["Vase Trapeze Vase Trapeze Base Plate Mesh"]
		r1Mesh := meshes["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
		r2Mesh := meshes["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h2 Mesh"]

		// Extract wall profiles indexed by world angle theta in [-pi, pi)
		type colProfile struct {
			theta float64
			pts   [][2]float64 // sorted by Y ascending: (r, y)
		}
		extractByWorldTheta := func(m *threejs.Mesh, substr string) []colProfile {
			if m == nil {
				return nil
			}
			byIdx := make(map[int][][3]float64) // colIndex -> (theta, r, y)
			for _, v := range m.BufferGeometry.Vertices {
				if !strings.Contains(v.Name, substr) {
					continue
				}
				var col int
				fields := strings.Fields(v.Name)
				if _, err := fmt.Sscanf(fields[len(fields)-1], "%d", &col); err != nil {
					continue
				}
				r := math.Hypot(v.X, v.Z)
				th := math.Atan2(v.Z, v.X)
				byIdx[col] = append(byIdx[col], [3]float64{th, r, v.Y})
			}
			var cols []colProfile
			for _, list := range byIdx {
				sort.Slice(list, func(i, j int) bool { return list[i][2] < list[j][2] })
				var pts [][2]float64
				for _, item := range list {
					if len(pts) > 0 && math.Abs(item[2]-pts[len(pts)-1][1]) < 1e-6 {
						continue
					}
					pts = append(pts, [2]float64{item[1], item[2]})
				}
				cols = append(cols, colProfile{theta: list[0][0], pts: pts})
			}
			sort.Slice(cols, func(i, j int) bool { return cols[i].theta < cols[j].theta })
			return cols
		}

		evalRAtThetaY := func(cols []colProfile, theta, y float64) (float64, bool) {
			if len(cols) < 2 {
				return 0, false
			}
			for theta < -math.Pi {
				theta += 2 * math.Pi
			}
			for theta >= math.Pi {
				theta -= 2 * math.Pi
			}
			// Find bracketing columns on circle
			nC := len(cols)
			i1 := sort.Search(nC, func(i int) bool { return cols[i].theta >= theta })
			var c0, c1 colProfile
			var dth, t float64
			if i1 == 0 {
				c0 = cols[nC-1]
				c1 = cols[0]
				dth = (c1.theta + 2*math.Pi) - c0.theta
				t = (theta + 2*math.Pi - c0.theta) / dth
			} else if i1 == nC {
				c0 = cols[nC-1]
				c1 = cols[0]
				dth = (c1.theta + 2*math.Pi) - c0.theta
				t = (theta - c0.theta) / dth
			} else {
				c0 = cols[i1-1]
				c1 = cols[i1]
				dth = c1.theta - c0.theta
				t = (theta - c0.theta) / dth
			}
			r0, ok0 := interpRAtMonotoneY(c0.pts, y)
			r1, ok1 := interpRAtMonotoneY(c1.pts, y)
			if !ok0 || !ok1 {
				return 0, false
			}
			return (1-t)*r0 + t*r1, true
		}

		topYAtTheta := func(cols []colProfile, theta float64) (float64, float64, bool) {
			if len(cols) < 2 {
				return 0, 0, false
			}
			nC := len(cols)
			i1 := sort.Search(nC, func(i int) bool { return cols[i].theta >= theta })
			var c0, c1 colProfile
			var dth, t float64
			if i1 == 0 {
				c0, c1 = cols[nC-1], cols[0]
				dth = (c1.theta + 2*math.Pi) - c0.theta
				t = (theta + 2*math.Pi - c0.theta) / dth
			} else if i1 == nC {
				c0, c1 = cols[nC-1], cols[0]
				dth = (c1.theta + 2*math.Pi) - c0.theta
				t = (theta - c0.theta) / dth
			} else {
				c0, c1 = cols[i1-1], cols[i1]
				dth = c1.theta - c0.theta
				t = (theta - c0.theta) / dth
			}
			p0 := c0.pts[len(c0.pts)-1]
			p1 := c1.pts[len(c1.pts)-1]
			return (1-t)*p0[0] + t*p1[0], (1-t)*p0[1] + t*p1[1], true
		}

		r1In := extractByWorldTheta(r1Mesh, "Inner Wall Face")
		r2Out := extractByWorldTheta(r2Mesh, "Outer Wall Face")

		minRimGap, maxRimGap := math.MaxFloat64, -math.MaxFloat64
		minMinOverlapGap, maxMinOverlapGap := math.MaxFloat64, -math.MaxFloat64
		var sampleTable []string

		for deg := -180; deg < 180; deg += 5 {
			th := float64(deg) * math.Pi / 180.0
			rInTop, yInTop, okTop := topYAtTheta(r1In, th)
			if !okTop {
				continue
			}
			if rOutAtRim, ok := evalRAtThetaY(r2Out, th, yInTop); ok {
				gap := rInTop - rOutAtRim
				if gap < minRimGap {
					minRimGap = gap
				}
				if gap > maxRimGap {
					maxRimGap = gap
				}
				// Also min gap across all Y along this world theta!
				bestOverlap := math.MaxFloat64
				for ky := 0; ky <= 60; ky++ {
					y := (yInTop - 350.0) + float64(ky)*(350.0/60.0)
					ri, ok1 := evalRAtThetaY(r1In, th, y)
					ro, ok2 := evalRAtThetaY(r2Out, th, y)
					if ok1 && ok2 {
						if g := ri - ro; g < bestOverlap {
							bestOverlap = g
						}
					}
				}
				if bestOverlap < minMinOverlapGap {
					minMinOverlapGap = bestOverlap
				}
				if bestOverlap < 1e9 && bestOverlap > maxMinOverlapGap {
					maxMinOverlapGap = bestOverlap
				}
				if deg%30 == 0 {
					sampleTable = append(sampleTable, fmt.Sprintf("th=%4d°: rimGap=%7.3f mm, minOverlapGap=%7.3f mm", deg, gap, bestOverlap))
				}
			}
		}

		yBounds := func(m *threejs.Mesh) (float64, float64) {
			if m == nil {
				return 0, 0
			}
			lo, hi := math.MaxFloat64, -math.MaxFloat64
			for _, v := range m.BufferGeometry.Vertices {
				if v.Y < lo {
					lo = v.Y
				}
				if v.Y > hi {
					hi = v.Y
				}
			}
			return lo, hi
		}
		_, max1 := yBounds(r1Mesh)
		min2, max2 := yBounds(r2Mesh)
		_ = baseMesh

		t.Logf("\n=== WORLD-THETA ANALYSIS: %s ===", label)
		t.Logf("  r1 topY=%.2f, r2 Y=[%.2f, %.2f] (stepDY=%.2f)", max1, min2, max2, max2-max1)
		t.Logf("  World-theta Rim radial gap (r1In(top) - r2Out(yTop)): min=%.4f mm, max=%.4f mm", minRimGap, maxRimGap)
		t.Logf("  World-theta Min-over-Y radial gap: min=%.4f mm, max=%.4f mm", minMinOverlapGap, maxMinOverlapGap)
		for _, s := range sampleTable {
			t.Logf("    %s", s)
		}

		// Also check BasePlate inner wall vs Ring 1 outer wall if baseMesh != nil
		if baseMesh != nil {
			bIn := extractByWorldTheta(baseMesh, "Ring Inner")
			r1Out := extractByWorldTheta(r1Mesh, "Outer Wall Face")
			minBGap, maxBGap := math.MaxFloat64, -math.MaxFloat64
			for deg := -180; deg < 180; deg += 5 {
				th := float64(deg) * math.Pi / 180.0
				rBTop, yBTop, okB := topYAtTheta(bIn, th)
				if !okB {
					continue
				}
				if r1AtB, ok := evalRAtThetaY(r1Out, th, yBTop); ok {
					g := rBTop - r1AtB
					if g < minBGap {
						minBGap = g
					}
					if g > maxBGap {
						maxBGap = g
					}
				}
			}
			t.Logf("  Base->Ring1 Rim radial gap: min=%.4f mm, max=%.4f mm", minBGap, maxBGap)
		}

		// Render software Z-buffer image to /tmp/sim_<label>.png
		if diag != nil {
			renderSoftwarePreview("/tmp/sim_"+label+".png", threejsStage)
		}
	}

	runCase("HEAD_unrot_const1_curv-0.25", 1, 3, 2, 0.0, 0.16, 0.08, -0.25, 0.0, true, false)
	runCase("HEAD_unrot_const1_curv0.0", 1, 3, 2, 0.0, 0.16, 0.08, 0.0, 0.0, true, false)
	runCase("HEAD_unrot_const0_curv-0.25", 1, 3, 2, 0.0, 0.16, 0.08, -0.25, 0.0, false, false)
	runCase("HEAD_unrot_const0_curv0.0", 1, 3, 2, 0.0, 0.16, 0.08, 0.0, 0.0, false, false)
	runCase("HEAD_rot1_const1_curv-0.25_carv0.3", 1, 3, 2, 1.0, 0.16, 0.08, -0.25, 0.3, true, false)
	runCase("HEAD_rot1_const1_curv0.0_carv0.3", 1, 3, 2, 1.0, 0.16, 0.08, 0.0, 0.3, true, false)
}

func renderSoftwarePreview(outPath string, st *threejs.Stage) {
	const W, H = 500, 500
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	zbuf := make([]float64, W*H)
	for i := range W * H {
		zbuf[i] = math.MaxFloat64
		img.SetRGBA(i%W, i/W, color.RGBA{255, 255, 255, 255})
	}

	// HEAD camera from git show HEAD:dsm/phylla/go/cmd/phylla/data/stage.go
	camX, camY, camZ := 507.81, 1990.69, -1329.81
	tgtX, tgtY, tgtZ := -3.27, 1491.90, 72.69
	fx, fy, fz := tgtX-camX, tgtY-camY, tgtZ-camZ
	fLen := math.Sqrt(fx*fx + fy*fy + fz*fz)
	fx, fy, fz = fx/fLen, fy/fLen, fz/fLen
	// right = normalize(cross(f, up=(0,1,0))) = normalize((-fz, 0, fx))
	rx, ry, rz := -fz, 0.0, fx
	rLen := math.Sqrt(rx*rx + ry*ry + rz*rz)
	rx, ry, rz = rx/rLen, ry/rLen, rz/rLen
	// u = cross(r, f)
	ux, uy, uz := ry*fz-rz*fy, rz*fx-rx*fz, rx*fy-ry*fx

	fovScale := 1.0 / math.Tan(22.5*math.Pi/180.0) // ~45 deg FOV

	project := func(x, y, z float64) (float64, float64, float64, bool) {
		dx, dy, dz := x-camX, y-camY, z-camZ
		cz := dx*fx + dy*fy + dz*fz
		if cz <= 10.0 {
			return 0, 0, 0, false
		}
		cx := dx*rx + dy*ry + dz*rz
		cy := dx*ux + dy*uy + dz*uz
		sx := float64(W)*0.5 + (cx/cz)*fovScale*float64(H)*0.5
		sy := float64(H)*0.5 - (cy/cz)*fovScale*float64(H)*0.5
		return sx, sy, cz, true
	}

	lx, ly, lz := 0.55, 0.65, -0.52
	lLen := math.Sqrt(lx*lx + ly*ly + lz*lz)
	lx, ly, lz = lx/lLen, ly/lLen, lz/lLen

	for m := range st.Meshs {
		if m.BufferGeometry == nil {
			continue
		}
		verts := m.BufferGeometry.Vertices
		baseR, baseG, baseB := 190.0, 68.0, 52.0
		if strings.Contains(m.Name, "h1") {
			baseR, baseG, baseB = 180.0, 62.0, 46.0
		}
		for _, f := range m.BufferGeometry.Faces {
			v1, v2, v3 := verts[f.V1], verts[f.V2], verts[f.V3]
			x1, y1, z1, ok1 := project(v1.X, v1.Y, v1.Z)
			x2, y2, z2, ok2 := project(v2.X, v2.Y, v2.Z)
			x3, y3, z3, ok3 := project(v3.X, v3.Y, v3.Z)
			if !ok1 || !ok2 || !ok3 {
				continue
			}
			// World normal
			abx, aby, abz := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
			acx, acy, acz := v3.X-v1.X, v3.Y-v1.Y, v3.Z-v1.Z
			nx := aby*acz - abz*acy
			ny := abz*acx - abx*acz
			nz := abx*acy - aby*acx
			nLen := math.Sqrt(nx*nx + ny*ny + nz*nz)
			if nLen < 1e-9 {
				continue
			}
			nx, ny, nz = nx/nLen, ny/nLen, nz/nLen
			// Flip toward camera if backfacing (double-sided)
			vx, vy, vz := camX-v1.X, camY-v1.Y, camZ-v1.Z
			if nx*vx+ny*vy+nz*vz < 0 {
				nx, ny, nz = -nx, -ny, -nz
			}
			diff := math.Max(0.0, nx*lx+ny*ly+nz*lz)
			shade := 0.30 + 0.70*diff

			minX := max(0, int(math.Floor(min(x1, x2, x3))))
			maxX := min(W-1, int(math.Ceil(max(x1, x2, x3))))
			minY := max(0, int(math.Floor(min(y1, y2, y3))))
			maxY := min(H-1, int(math.Ceil(max(y1, y2, y3))))
			denom := (y2-y3)*(x1-x3) + (x3-x2)*(y1-y3)
			if math.Abs(denom) < 1e-6 {
				continue
			}
			for py := minY; py <= maxY; py++ {
				for px := minX; px <= maxX; px++ {
					fxp, fyp := float64(px)+0.5, float64(py)+0.5
					w1 := ((y2-y3)*(fxp-x3) + (x3-x2)*(fyp-y3)) / denom
					w2 := ((y3-y1)*(fxp-x3) + (x1-x3)*(fyp-y3)) / denom
					w3 := 1.0 - w1 - w2
					if w1 >= 0 && w2 >= 0 && w3 >= 0 {
						cz := w1*z1 + w2*z2 + w3*z3
						idx := py*W + px
						if cz < zbuf[idx] {
							zbuf[idx] = cz
							img.SetRGBA(px, py, color.RGBA{
								R: uint8(min(255.0, baseR*shade)),
								G: uint8(min(255.0, baseG*shade)),
								B: uint8(min(255.0, baseB*shade)),
								A: 255,
							})
						}
					}
				}
			}
		}
	}
	f, err := os.Create(outPath)
	if err == nil {
		_ = png.Encode(f, img)
		_ = f.Close()
	}
}
