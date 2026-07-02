package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__AmbiantLight__00000000_ := (&models.AmbiantLight{Name: `Ambiant Light`}).Stage(stage)

	__BoxGeometry__00000000_ := (&models.BoxGeometry{Name: `Box`}).Stage(stage)

	__Canvas__00000000_ := (&models.Canvas{Name: `Singloton`}).Stage(stage)

	__Curve__00000000_ := (&models.Curve{Name: `Helix Path`}).Stage(stage)
	__Curve__00000001_ := (&models.Curve{Name: `Torus Like Curve`}).Stage(stage)

	__CylinderGeometry__00000000_ := (&models.CylinderGeometry{Name: `Cylinder`}).Stage(stage)
	__CylinderGeometry__00000002_ := (&models.CylinderGeometry{Name: `Cylinder 2`}).Stage(stage)
	__CylinderGeometry__00000003_ := (&models.CylinderGeometry{Name: `Cylinder 3`}).Stage(stage)

	__DirectionalLight__00000000_ := (&models.DirectionalLight{Name: `Direction Light`}).Stage(stage)

	__ExtrudeGeometry__00000000_ := (&models.ExtrudeGeometry{Name: `Torus Extrude Geometry`}).Stage(stage)

	__Mesh__00000000_ := (&models.Mesh{Name: `Box Mesh`}).Stage(stage)
	__Mesh__00000001_ := (&models.Mesh{Name: `Cylinder Mesh`}).Stage(stage)
	__Mesh__00000002_ := (&models.Mesh{Name: `Sphere Mesh`}).Stage(stage)
	__Mesh__00000003_ := (&models.Mesh{Name: `Torus Mesh`}).Stage(stage)
	__Mesh__00000004_ := (&models.Mesh{Name: `Plane Mesh`}).Stage(stage)
	__Mesh__00000005_ := (&models.Mesh{Name: `Cylinder Mesh 2`}).Stage(stage)
	__Mesh__00000006_ := (&models.Mesh{Name: `Cylinder Mesh 3`}).Stage(stage)
	__Mesh__00000007_ := (&models.Mesh{Name: `Helix Mesh`}).Stage(stage)
	__Mesh__00000008_ := (&models.Mesh{Name: `Extrude Mesh`}).Stage(stage)

	__MeshMaterialBasic__00000000_ := (&models.MeshMaterialBasic{Name: `Pink`}).Stage(stage)
	__MeshMaterialBasic__00000001_ := (&models.MeshMaterialBasic{Name: `Gold`}).Stage(stage)
	__MeshMaterialBasic__00000002_ := (&models.MeshMaterialBasic{Name: `White`}).Stage(stage)
	__MeshMaterialBasic__00000003_ := (&models.MeshMaterialBasic{Name: `Gray`}).Stage(stage)
	__MeshMaterialBasic__00000004_ := (&models.MeshMaterialBasic{Name: `Light Gray`}).Stage(stage)
	__MeshMaterialBasic__00000005_ := (&models.MeshMaterialBasic{Name: `Silver`}).Stage(stage)
	__MeshMaterialBasic__00000006_ := (&models.MeshMaterialBasic{Name: `Magenta`}).Stage(stage)

	__MeshPhysicalMaterial__00000000_ := (&models.MeshPhysicalMaterial{Name: `Cyan Glass`}).Stage(stage)

	__PlaneGeometry__00000000_ := (&models.PlaneGeometry{Name: `Plane`}).Stage(stage)

	__Shape__00000000_ := (&models.Shape{Name: `Square Shape`}).Stage(stage)

	__SphereGeometry__00000000_ := (&models.SphereGeometry{Name: `Sphere`}).Stage(stage)

	__TorusGeometry__00000000_ := (&models.TorusGeometry{Name: `Torus`}).Stage(stage)

	__TubeGeometry__00000000_ := (&models.TubeGeometry{Name: `Helix Tube Geometry`}).Stage(stage)

	__Vector2__00000000_ := (&models.Vector2{Name: `SquarePoint 0`}).Stage(stage)
	__Vector2__00000001_ := (&models.Vector2{Name: `SquarePoint 1`}).Stage(stage)
	__Vector2__00000002_ := (&models.Vector2{Name: `SquarePoint 2`}).Stage(stage)
	__Vector2__00000003_ := (&models.Vector2{Name: `SquarePoint 3`}).Stage(stage)
	__Vector2__00000004_ := (&models.Vector2{Name: `SquarePoint 4`}).Stage(stage)

	__Vector3__00000000_ := (&models.Vector3{Name: `Point 0`}).Stage(stage)
	__Vector3__00000001_ := (&models.Vector3{Name: `Point 1`}).Stage(stage)
	__Vector3__00000002_ := (&models.Vector3{Name: `Point 2`}).Stage(stage)
	__Vector3__00000003_ := (&models.Vector3{Name: `Point 3`}).Stage(stage)
	__Vector3__00000004_ := (&models.Vector3{Name: `Point 4`}).Stage(stage)
	__Vector3__00000005_ := (&models.Vector3{Name: `Point 5`}).Stage(stage)
	__Vector3__00000006_ := (&models.Vector3{Name: `Point 6`}).Stage(stage)
	__Vector3__00000007_ := (&models.Vector3{Name: `Point 7`}).Stage(stage)
	__Vector3__00000008_ := (&models.Vector3{Name: `Point 8`}).Stage(stage)
	__Vector3__00000009_ := (&models.Vector3{Name: `Point 9`}).Stage(stage)
	__Vector3__00000010_ := (&models.Vector3{Name: `Point 10`}).Stage(stage)
	__Vector3__00000011_ := (&models.Vector3{Name: `Point 11`}).Stage(stage)
	__Vector3__00000012_ := (&models.Vector3{Name: `Point 12`}).Stage(stage)
	__Vector3__00000013_ := (&models.Vector3{Name: `Point 13`}).Stage(stage)
	__Vector3__00000014_ := (&models.Vector3{Name: `Point 14`}).Stage(stage)
	__Vector3__00000015_ := (&models.Vector3{Name: `Point 15`}).Stage(stage)
	__Vector3__00000016_ := (&models.Vector3{Name: `Point 16`}).Stage(stage)
	__Vector3__00000017_ := (&models.Vector3{Name: `Point 17`}).Stage(stage)
	__Vector3__00000018_ := (&models.Vector3{Name: `Point 18`}).Stage(stage)
	__Vector3__00000019_ := (&models.Vector3{Name: `Point 19`}).Stage(stage)
	__Vector3__00000020_ := (&models.Vector3{Name: `Point 20`}).Stage(stage)
	__Vector3__00000021_ := (&models.Vector3{Name: `Point 21`}).Stage(stage)
	__Vector3__00000022_ := (&models.Vector3{Name: `Point 22`}).Stage(stage)
	__Vector3__00000023_ := (&models.Vector3{Name: `Point 23`}).Stage(stage)
	__Vector3__00000024_ := (&models.Vector3{Name: `Point 24`}).Stage(stage)
	__Vector3__00000025_ := (&models.Vector3{Name: `Point 25`}).Stage(stage)
	__Vector3__00000026_ := (&models.Vector3{Name: `Point 26`}).Stage(stage)
	__Vector3__00000027_ := (&models.Vector3{Name: `Point 27`}).Stage(stage)
	__Vector3__00000028_ := (&models.Vector3{Name: `Point 28`}).Stage(stage)
	__Vector3__00000029_ := (&models.Vector3{Name: `Point 29`}).Stage(stage)
	__Vector3__00000030_ := (&models.Vector3{Name: `Point 30`}).Stage(stage)
	__Vector3__00000031_ := (&models.Vector3{Name: `Point 31`}).Stage(stage)
	__Vector3__00000032_ := (&models.Vector3{Name: `Point 32`}).Stage(stage)
	__Vector3__00000033_ := (&models.Vector3{Name: `Point 33`}).Stage(stage)
	__Vector3__00000034_ := (&models.Vector3{Name: `Point 34`}).Stage(stage)
	__Vector3__00000035_ := (&models.Vector3{Name: `Point 35`}).Stage(stage)
	__Vector3__00000036_ := (&models.Vector3{Name: `Point 36`}).Stage(stage)
	__Vector3__00000037_ := (&models.Vector3{Name: `Point 37`}).Stage(stage)
	__Vector3__00000038_ := (&models.Vector3{Name: `Point 38`}).Stage(stage)
	__Vector3__00000039_ := (&models.Vector3{Name: `Point 39`}).Stage(stage)
	__Vector3__00000040_ := (&models.Vector3{Name: `Point 40`}).Stage(stage)
	__Vector3__00000041_ := (&models.Vector3{Name: `Point 41`}).Stage(stage)
	__Vector3__00000042_ := (&models.Vector3{Name: `Point 42`}).Stage(stage)
	__Vector3__00000043_ := (&models.Vector3{Name: `Point 43`}).Stage(stage)
	__Vector3__00000044_ := (&models.Vector3{Name: `Point 44`}).Stage(stage)
	__Vector3__00000045_ := (&models.Vector3{Name: `Point 45`}).Stage(stage)
	__Vector3__00000046_ := (&models.Vector3{Name: `Point 46`}).Stage(stage)
	__Vector3__00000047_ := (&models.Vector3{Name: `Point 47`}).Stage(stage)
	__Vector3__00000048_ := (&models.Vector3{Name: `Point 48`}).Stage(stage)
	__Vector3__00000049_ := (&models.Vector3{Name: `Point 49`}).Stage(stage)
	__Vector3__00000050_ := (&models.Vector3{Name: `Point 50`}).Stage(stage)
	__Vector3__00000051_ := (&models.Vector3{Name: `Point 51`}).Stage(stage)
	__Vector3__00000052_ := (&models.Vector3{Name: `Point 52`}).Stage(stage)
	__Vector3__00000053_ := (&models.Vector3{Name: `Point 53`}).Stage(stage)
	__Vector3__00000054_ := (&models.Vector3{Name: `Point 54`}).Stage(stage)
	__Vector3__00000055_ := (&models.Vector3{Name: `Point 55`}).Stage(stage)
	__Vector3__00000056_ := (&models.Vector3{Name: `Point 56`}).Stage(stage)
	__Vector3__00000057_ := (&models.Vector3{Name: `Point 57`}).Stage(stage)
	__Vector3__00000058_ := (&models.Vector3{Name: `Point 58`}).Stage(stage)
	__Vector3__00000059_ := (&models.Vector3{Name: `Point 59`}).Stage(stage)
	__Vector3__00000060_ := (&models.Vector3{Name: `Point 60`}).Stage(stage)
	__Vector3__00000061_ := (&models.Vector3{Name: `Point 61`}).Stage(stage)
	__Vector3__00000062_ := (&models.Vector3{Name: `Point 62`}).Stage(stage)
	__Vector3__00000063_ := (&models.Vector3{Name: `Point 63`}).Stage(stage)
	__Vector3__00000064_ := (&models.Vector3{Name: `Point 64`}).Stage(stage)
	__Vector3__00000065_ := (&models.Vector3{Name: `Point 65`}).Stage(stage)
	__Vector3__00000066_ := (&models.Vector3{Name: `Point 66`}).Stage(stage)
	__Vector3__00000067_ := (&models.Vector3{Name: `Point 67`}).Stage(stage)
	__Vector3__00000068_ := (&models.Vector3{Name: `Point 68`}).Stage(stage)
	__Vector3__00000069_ := (&models.Vector3{Name: `Point 69`}).Stage(stage)
	__Vector3__00000070_ := (&models.Vector3{Name: `Point 70`}).Stage(stage)
	__Vector3__00000071_ := (&models.Vector3{Name: `Point 71`}).Stage(stage)
	__Vector3__00000072_ := (&models.Vector3{Name: `Point 72`}).Stage(stage)
	__Vector3__00000073_ := (&models.Vector3{Name: `Point 73`}).Stage(stage)
	__Vector3__00000074_ := (&models.Vector3{Name: `Point 74`}).Stage(stage)
	__Vector3__00000075_ := (&models.Vector3{Name: `Point 75`}).Stage(stage)
	__Vector3__00000076_ := (&models.Vector3{Name: `Point 76`}).Stage(stage)
	__Vector3__00000077_ := (&models.Vector3{Name: `Point 77`}).Stage(stage)
	__Vector3__00000078_ := (&models.Vector3{Name: `Point 78`}).Stage(stage)
	__Vector3__00000079_ := (&models.Vector3{Name: `Point 79`}).Stage(stage)
	__Vector3__00000080_ := (&models.Vector3{Name: `Point 80`}).Stage(stage)
	__Vector3__00000081_ := (&models.Vector3{Name: `Point 81`}).Stage(stage)
	__Vector3__00000082_ := (&models.Vector3{Name: `Point 82`}).Stage(stage)
	__Vector3__00000083_ := (&models.Vector3{Name: `Point 83`}).Stage(stage)
	__Vector3__00000084_ := (&models.Vector3{Name: `Point 84`}).Stage(stage)
	__Vector3__00000085_ := (&models.Vector3{Name: `Point 85`}).Stage(stage)
	__Vector3__00000086_ := (&models.Vector3{Name: `Point 86`}).Stage(stage)
	__Vector3__00000087_ := (&models.Vector3{Name: `Point 87`}).Stage(stage)
	__Vector3__00000088_ := (&models.Vector3{Name: `Point 88`}).Stage(stage)
	__Vector3__00000089_ := (&models.Vector3{Name: `Point 89`}).Stage(stage)
	__Vector3__00000090_ := (&models.Vector3{Name: `Point 90`}).Stage(stage)
	__Vector3__00000091_ := (&models.Vector3{Name: `Point 91`}).Stage(stage)
	__Vector3__00000092_ := (&models.Vector3{Name: `Point 92`}).Stage(stage)
	__Vector3__00000093_ := (&models.Vector3{Name: `Point 93`}).Stage(stage)
	__Vector3__00000094_ := (&models.Vector3{Name: `Point 94`}).Stage(stage)
	__Vector3__00000095_ := (&models.Vector3{Name: `Point 95`}).Stage(stage)
	__Vector3__00000096_ := (&models.Vector3{Name: `Point 96`}).Stage(stage)
	__Vector3__00000097_ := (&models.Vector3{Name: `Point 97`}).Stage(stage)
	__Vector3__00000098_ := (&models.Vector3{Name: `Point 98`}).Stage(stage)
	__Vector3__00000099_ := (&models.Vector3{Name: `Point 99`}).Stage(stage)
	__Vector3__00000100_ := (&models.Vector3{Name: `Point 100`}).Stage(stage)
	__Vector3__00000101_ := (&models.Vector3{Name: `Point 101`}).Stage(stage)
	__Vector3__00000102_ := (&models.Vector3{Name: `Point 102`}).Stage(stage)
	__Vector3__00000103_ := (&models.Vector3{Name: `Point 103`}).Stage(stage)
	__Vector3__00000104_ := (&models.Vector3{Name: `Point 104`}).Stage(stage)
	__Vector3__00000105_ := (&models.Vector3{Name: `Point 105`}).Stage(stage)
	__Vector3__00000106_ := (&models.Vector3{Name: `Point 106`}).Stage(stage)
	__Vector3__00000107_ := (&models.Vector3{Name: `Point 107`}).Stage(stage)
	__Vector3__00000108_ := (&models.Vector3{Name: `Point 108`}).Stage(stage)
	__Vector3__00000109_ := (&models.Vector3{Name: `Point 109`}).Stage(stage)
	__Vector3__00000110_ := (&models.Vector3{Name: `Point 110`}).Stage(stage)
	__Vector3__00000111_ := (&models.Vector3{Name: `Point 111`}).Stage(stage)
	__Vector3__00000112_ := (&models.Vector3{Name: `Point 112`}).Stage(stage)
	__Vector3__00000113_ := (&models.Vector3{Name: `Point 113`}).Stage(stage)
	__Vector3__00000114_ := (&models.Vector3{Name: `Point 114`}).Stage(stage)
	__Vector3__00000115_ := (&models.Vector3{Name: `Point 115`}).Stage(stage)
	__Vector3__00000116_ := (&models.Vector3{Name: `Point 116`}).Stage(stage)
	__Vector3__00000117_ := (&models.Vector3{Name: `Point 117`}).Stage(stage)
	__Vector3__00000118_ := (&models.Vector3{Name: `Point 118`}).Stage(stage)
	__Vector3__00000119_ := (&models.Vector3{Name: `Point 119`}).Stage(stage)
	__Vector3__00000120_ := (&models.Vector3{Name: `Point 120`}).Stage(stage)
	__Vector3__00000121_ := (&models.Vector3{Name: `Point 121`}).Stage(stage)
	__Vector3__00000122_ := (&models.Vector3{Name: `Point 122`}).Stage(stage)
	__Vector3__00000123_ := (&models.Vector3{Name: `Point 123`}).Stage(stage)
	__Vector3__00000124_ := (&models.Vector3{Name: `Point 124`}).Stage(stage)
	__Vector3__00000125_ := (&models.Vector3{Name: `Point 125`}).Stage(stage)
	__Vector3__00000126_ := (&models.Vector3{Name: `Point 126`}).Stage(stage)
	__Vector3__00000127_ := (&models.Vector3{Name: `Point 127`}).Stage(stage)
	__Vector3__00000128_ := (&models.Vector3{Name: `Point 128`}).Stage(stage)
	__Vector3__00000129_ := (&models.Vector3{Name: `Point 129`}).Stage(stage)
	__Vector3__00000130_ := (&models.Vector3{Name: `Point 130`}).Stage(stage)
	__Vector3__00000131_ := (&models.Vector3{Name: `Point 131`}).Stage(stage)
	__Vector3__00000132_ := (&models.Vector3{Name: `Point 132`}).Stage(stage)
	__Vector3__00000133_ := (&models.Vector3{Name: `Point 133`}).Stage(stage)
	__Vector3__00000134_ := (&models.Vector3{Name: `Point 134`}).Stage(stage)
	__Vector3__00000135_ := (&models.Vector3{Name: `Point 135`}).Stage(stage)
	__Vector3__00000136_ := (&models.Vector3{Name: `Point 136`}).Stage(stage)
	__Vector3__00000137_ := (&models.Vector3{Name: `Point 137`}).Stage(stage)
	__Vector3__00000138_ := (&models.Vector3{Name: `Point 138`}).Stage(stage)
	__Vector3__00000139_ := (&models.Vector3{Name: `Point 139`}).Stage(stage)
	__Vector3__00000140_ := (&models.Vector3{Name: `Point 140`}).Stage(stage)
	__Vector3__00000141_ := (&models.Vector3{Name: `Point 141`}).Stage(stage)
	__Vector3__00000142_ := (&models.Vector3{Name: `Point 142`}).Stage(stage)
	__Vector3__00000143_ := (&models.Vector3{Name: `Point 143`}).Stage(stage)
	__Vector3__00000144_ := (&models.Vector3{Name: `Point 144`}).Stage(stage)
	__Vector3__00000145_ := (&models.Vector3{Name: `Point 145`}).Stage(stage)
	__Vector3__00000146_ := (&models.Vector3{Name: `Point 146`}).Stage(stage)
	__Vector3__00000147_ := (&models.Vector3{Name: `Point 147`}).Stage(stage)
	__Vector3__00000148_ := (&models.Vector3{Name: `Point 148`}).Stage(stage)
	__Vector3__00000149_ := (&models.Vector3{Name: `Point 149`}).Stage(stage)
	__Vector3__00000150_ := (&models.Vector3{Name: `Point 150`}).Stage(stage)
	__Vector3__00000151_ := (&models.Vector3{Name: `Point 151`}).Stage(stage)
	__Vector3__00000152_ := (&models.Vector3{Name: `Point 152`}).Stage(stage)
	__Vector3__00000153_ := (&models.Vector3{Name: `Point 153`}).Stage(stage)
	__Vector3__00000154_ := (&models.Vector3{Name: `Point 154`}).Stage(stage)
	__Vector3__00000155_ := (&models.Vector3{Name: `Point 155`}).Stage(stage)
	__Vector3__00000156_ := (&models.Vector3{Name: `Point 156`}).Stage(stage)
	__Vector3__00000157_ := (&models.Vector3{Name: `Point 157`}).Stage(stage)
	__Vector3__00000158_ := (&models.Vector3{Name: `Point 158`}).Stage(stage)
	__Vector3__00000159_ := (&models.Vector3{Name: `Point 159`}).Stage(stage)
	__Vector3__00000160_ := (&models.Vector3{Name: `Point 160`}).Stage(stage)
	__Vector3__00000161_ := (&models.Vector3{Name: `Point 161`}).Stage(stage)
	__Vector3__00000162_ := (&models.Vector3{Name: `Point 162`}).Stage(stage)
	__Vector3__00000163_ := (&models.Vector3{Name: `Point 163`}).Stage(stage)
	__Vector3__00000164_ := (&models.Vector3{Name: `Point 164`}).Stage(stage)
	__Vector3__00000165_ := (&models.Vector3{Name: `Point 165`}).Stage(stage)
	__Vector3__00000166_ := (&models.Vector3{Name: `Point 166`}).Stage(stage)
	__Vector3__00000167_ := (&models.Vector3{Name: `Point 167`}).Stage(stage)
	__Vector3__00000168_ := (&models.Vector3{Name: `Point 168`}).Stage(stage)
	__Vector3__00000169_ := (&models.Vector3{Name: `Point 169`}).Stage(stage)
	__Vector3__00000170_ := (&models.Vector3{Name: `Point 170`}).Stage(stage)
	__Vector3__00000171_ := (&models.Vector3{Name: `Point 171`}).Stage(stage)
	__Vector3__00000172_ := (&models.Vector3{Name: `Point 172`}).Stage(stage)
	__Vector3__00000173_ := (&models.Vector3{Name: `Point 173`}).Stage(stage)
	__Vector3__00000174_ := (&models.Vector3{Name: `Point 174`}).Stage(stage)
	__Vector3__00000175_ := (&models.Vector3{Name: `Point 175`}).Stage(stage)
	__Vector3__00000176_ := (&models.Vector3{Name: `Point 176`}).Stage(stage)
	__Vector3__00000177_ := (&models.Vector3{Name: `Point 177`}).Stage(stage)
	__Vector3__00000178_ := (&models.Vector3{Name: `Point 178`}).Stage(stage)
	__Vector3__00000179_ := (&models.Vector3{Name: `Point 179`}).Stage(stage)
	__Vector3__00000180_ := (&models.Vector3{Name: `Point 180`}).Stage(stage)
	__Vector3__00000181_ := (&models.Vector3{Name: `Point 181`}).Stage(stage)
	__Vector3__00000182_ := (&models.Vector3{Name: `Point 182`}).Stage(stage)
	__Vector3__00000183_ := (&models.Vector3{Name: `Point 183`}).Stage(stage)
	__Vector3__00000184_ := (&models.Vector3{Name: `Point 184`}).Stage(stage)
	__Vector3__00000185_ := (&models.Vector3{Name: `Point 185`}).Stage(stage)
	__Vector3__00000186_ := (&models.Vector3{Name: `Point 186`}).Stage(stage)
	__Vector3__00000187_ := (&models.Vector3{Name: `Point 187`}).Stage(stage)
	__Vector3__00000188_ := (&models.Vector3{Name: `Point 188`}).Stage(stage)
	__Vector3__00000189_ := (&models.Vector3{Name: `Point 189`}).Stage(stage)
	__Vector3__00000190_ := (&models.Vector3{Name: `Point 190`}).Stage(stage)
	__Vector3__00000191_ := (&models.Vector3{Name: `Point 191`}).Stage(stage)
	__Vector3__00000192_ := (&models.Vector3{Name: `Point 192`}).Stage(stage)
	__Vector3__00000193_ := (&models.Vector3{Name: `Point 193`}).Stage(stage)
	__Vector3__00000194_ := (&models.Vector3{Name: `Point 194`}).Stage(stage)
	__Vector3__00000195_ := (&models.Vector3{Name: `Point 195`}).Stage(stage)
	__Vector3__00000196_ := (&models.Vector3{Name: `Point 196`}).Stage(stage)
	__Vector3__00000197_ := (&models.Vector3{Name: `Point 197`}).Stage(stage)
	__Vector3__00000198_ := (&models.Vector3{Name: `Point 198`}).Stage(stage)
	__Vector3__00000199_ := (&models.Vector3{Name: `Point 199`}).Stage(stage)
	__Vector3__00000200_ := (&models.Vector3{Name: `Point 200`}).Stage(stage)
	__Vector3__00000201_ := (&models.Vector3{Name: `TorusPoint 0`}).Stage(stage)
	__Vector3__00000202_ := (&models.Vector3{Name: `TorusPoint 1`}).Stage(stage)
	__Vector3__00000203_ := (&models.Vector3{Name: `TorusPoint 2`}).Stage(stage)
	__Vector3__00000204_ := (&models.Vector3{Name: `TorusPoint 3`}).Stage(stage)
	__Vector3__00000205_ := (&models.Vector3{Name: `TorusPoint 4`}).Stage(stage)
	__Vector3__00000206_ := (&models.Vector3{Name: `TorusPoint 5`}).Stage(stage)
	__Vector3__00000207_ := (&models.Vector3{Name: `TorusPoint 6`}).Stage(stage)
	__Vector3__00000208_ := (&models.Vector3{Name: `TorusPoint 7`}).Stage(stage)
	__Vector3__00000209_ := (&models.Vector3{Name: `TorusPoint 8`}).Stage(stage)
	__Vector3__00000210_ := (&models.Vector3{Name: `TorusPoint 9`}).Stage(stage)
	__Vector3__00000211_ := (&models.Vector3{Name: `TorusPoint 10`}).Stage(stage)
	__Vector3__00000212_ := (&models.Vector3{Name: `TorusPoint 11`}).Stage(stage)
	__Vector3__00000213_ := (&models.Vector3{Name: `TorusPoint 12`}).Stage(stage)
	__Vector3__00000214_ := (&models.Vector3{Name: `TorusPoint 13`}).Stage(stage)
	__Vector3__00000215_ := (&models.Vector3{Name: `TorusPoint 14`}).Stage(stage)
	__Vector3__00000216_ := (&models.Vector3{Name: `TorusPoint 15`}).Stage(stage)
	__Vector3__00000217_ := (&models.Vector3{Name: `TorusPoint 16`}).Stage(stage)
	__Vector3__00000218_ := (&models.Vector3{Name: `TorusPoint 17`}).Stage(stage)
	__Vector3__00000219_ := (&models.Vector3{Name: `TorusPoint 18`}).Stage(stage)
	__Vector3__00000220_ := (&models.Vector3{Name: `TorusPoint 19`}).Stage(stage)
	__Vector3__00000221_ := (&models.Vector3{Name: `TorusPoint 20`}).Stage(stage)
	__Vector3__00000222_ := (&models.Vector3{Name: `TorusPoint 21`}).Stage(stage)
	__Vector3__00000223_ := (&models.Vector3{Name: `TorusPoint 22`}).Stage(stage)
	__Vector3__00000224_ := (&models.Vector3{Name: `TorusPoint 23`}).Stage(stage)
	__Vector3__00000225_ := (&models.Vector3{Name: `TorusPoint 24`}).Stage(stage)
	__Vector3__00000226_ := (&models.Vector3{Name: `TorusPoint 25`}).Stage(stage)
	__Vector3__00000227_ := (&models.Vector3{Name: `TorusPoint 26`}).Stage(stage)
	__Vector3__00000228_ := (&models.Vector3{Name: `TorusPoint 27`}).Stage(stage)
	__Vector3__00000229_ := (&models.Vector3{Name: `TorusPoint 28`}).Stage(stage)
	__Vector3__00000230_ := (&models.Vector3{Name: `TorusPoint 29`}).Stage(stage)
	__Vector3__00000231_ := (&models.Vector3{Name: `TorusPoint 30`}).Stage(stage)
	__Vector3__00000232_ := (&models.Vector3{Name: `TorusPoint 31`}).Stage(stage)
	__Vector3__00000233_ := (&models.Vector3{Name: `TorusPoint 32`}).Stage(stage)
	__Vector3__00000234_ := (&models.Vector3{Name: `TorusPoint 33`}).Stage(stage)
	__Vector3__00000235_ := (&models.Vector3{Name: `TorusPoint 34`}).Stage(stage)
	__Vector3__00000236_ := (&models.Vector3{Name: `TorusPoint 35`}).Stage(stage)
	__Vector3__00000237_ := (&models.Vector3{Name: `TorusPoint 36`}).Stage(stage)
	__Vector3__00000238_ := (&models.Vector3{Name: `TorusPoint 37`}).Stage(stage)
	__Vector3__00000239_ := (&models.Vector3{Name: `TorusPoint 38`}).Stage(stage)
	__Vector3__00000240_ := (&models.Vector3{Name: `TorusPoint 39`}).Stage(stage)
	__Vector3__00000241_ := (&models.Vector3{Name: `TorusPoint 40`}).Stage(stage)
	__Vector3__00000242_ := (&models.Vector3{Name: `TorusPoint 41`}).Stage(stage)
	__Vector3__00000243_ := (&models.Vector3{Name: `TorusPoint 42`}).Stage(stage)
	__Vector3__00000244_ := (&models.Vector3{Name: `TorusPoint 43`}).Stage(stage)
	__Vector3__00000245_ := (&models.Vector3{Name: `TorusPoint 44`}).Stage(stage)
	__Vector3__00000246_ := (&models.Vector3{Name: `TorusPoint 45`}).Stage(stage)
	__Vector3__00000247_ := (&models.Vector3{Name: `TorusPoint 46`}).Stage(stage)
	__Vector3__00000248_ := (&models.Vector3{Name: `TorusPoint 47`}).Stage(stage)
	__Vector3__00000249_ := (&models.Vector3{Name: `TorusPoint 48`}).Stage(stage)
	__Vector3__00000250_ := (&models.Vector3{Name: `TorusPoint 49`}).Stage(stage)
	__Vector3__00000251_ := (&models.Vector3{Name: `TorusPoint 50`}).Stage(stage)
	__Vector3__00000252_ := (&models.Vector3{Name: `TorusPoint 51`}).Stage(stage)
	__Vector3__00000253_ := (&models.Vector3{Name: `TorusPoint 52`}).Stage(stage)
	__Vector3__00000254_ := (&models.Vector3{Name: `TorusPoint 53`}).Stage(stage)
	__Vector3__00000255_ := (&models.Vector3{Name: `TorusPoint 54`}).Stage(stage)
	__Vector3__00000256_ := (&models.Vector3{Name: `TorusPoint 55`}).Stage(stage)
	__Vector3__00000257_ := (&models.Vector3{Name: `TorusPoint 56`}).Stage(stage)
	__Vector3__00000258_ := (&models.Vector3{Name: `TorusPoint 57`}).Stage(stage)
	__Vector3__00000259_ := (&models.Vector3{Name: `TorusPoint 58`}).Stage(stage)
	__Vector3__00000260_ := (&models.Vector3{Name: `TorusPoint 59`}).Stage(stage)
	__Vector3__00000261_ := (&models.Vector3{Name: `TorusPoint 60`}).Stage(stage)
	__Vector3__00000262_ := (&models.Vector3{Name: `TorusPoint 61`}).Stage(stage)
	__Vector3__00000263_ := (&models.Vector3{Name: `TorusPoint 62`}).Stage(stage)
	__Vector3__00000264_ := (&models.Vector3{Name: `TorusPoint 63`}).Stage(stage)
	__Vector3__00000265_ := (&models.Vector3{Name: `TorusPoint 64`}).Stage(stage)
	__Vector3__00000266_ := (&models.Vector3{Name: `TorusPoint 65`}).Stage(stage)
	__Vector3__00000267_ := (&models.Vector3{Name: `TorusPoint 66`}).Stage(stage)
	__Vector3__00000268_ := (&models.Vector3{Name: `TorusPoint 67`}).Stage(stage)
	__Vector3__00000269_ := (&models.Vector3{Name: `TorusPoint 68`}).Stage(stage)
	__Vector3__00000270_ := (&models.Vector3{Name: `TorusPoint 69`}).Stage(stage)
	__Vector3__00000271_ := (&models.Vector3{Name: `TorusPoint 70`}).Stage(stage)
	__Vector3__00000272_ := (&models.Vector3{Name: `TorusPoint 71`}).Stage(stage)
	__Vector3__00000273_ := (&models.Vector3{Name: `TorusPoint 72`}).Stage(stage)
	__Vector3__00000274_ := (&models.Vector3{Name: `TorusPoint 73`}).Stage(stage)
	__Vector3__00000275_ := (&models.Vector3{Name: `TorusPoint 74`}).Stage(stage)
	__Vector3__00000276_ := (&models.Vector3{Name: `TorusPoint 75`}).Stage(stage)
	__Vector3__00000277_ := (&models.Vector3{Name: `TorusPoint 76`}).Stage(stage)
	__Vector3__00000278_ := (&models.Vector3{Name: `TorusPoint 77`}).Stage(stage)
	__Vector3__00000279_ := (&models.Vector3{Name: `TorusPoint 78`}).Stage(stage)
	__Vector3__00000280_ := (&models.Vector3{Name: `TorusPoint 79`}).Stage(stage)
	__Vector3__00000281_ := (&models.Vector3{Name: `TorusPoint 80`}).Stage(stage)
	__Vector3__00000282_ := (&models.Vector3{Name: `TorusPoint 81`}).Stage(stage)
	__Vector3__00000283_ := (&models.Vector3{Name: `TorusPoint 82`}).Stage(stage)
	__Vector3__00000284_ := (&models.Vector3{Name: `TorusPoint 83`}).Stage(stage)
	__Vector3__00000285_ := (&models.Vector3{Name: `TorusPoint 84`}).Stage(stage)
	__Vector3__00000286_ := (&models.Vector3{Name: `TorusPoint 85`}).Stage(stage)
	__Vector3__00000287_ := (&models.Vector3{Name: `TorusPoint 86`}).Stage(stage)
	__Vector3__00000288_ := (&models.Vector3{Name: `TorusPoint 87`}).Stage(stage)
	__Vector3__00000289_ := (&models.Vector3{Name: `TorusPoint 88`}).Stage(stage)
	__Vector3__00000290_ := (&models.Vector3{Name: `TorusPoint 89`}).Stage(stage)
	__Vector3__00000291_ := (&models.Vector3{Name: `TorusPoint 90`}).Stage(stage)
	__Vector3__00000292_ := (&models.Vector3{Name: `TorusPoint 91`}).Stage(stage)
	__Vector3__00000293_ := (&models.Vector3{Name: `TorusPoint 92`}).Stage(stage)
	__Vector3__00000294_ := (&models.Vector3{Name: `TorusPoint 93`}).Stage(stage)
	__Vector3__00000295_ := (&models.Vector3{Name: `TorusPoint 94`}).Stage(stage)
	__Vector3__00000296_ := (&models.Vector3{Name: `TorusPoint 95`}).Stage(stage)
	__Vector3__00000297_ := (&models.Vector3{Name: `TorusPoint 96`}).Stage(stage)
	__Vector3__00000298_ := (&models.Vector3{Name: `TorusPoint 97`}).Stage(stage)
	__Vector3__00000299_ := (&models.Vector3{Name: `TorusPoint 98`}).Stage(stage)
	__Vector3__00000300_ := (&models.Vector3{Name: `TorusPoint 99`}).Stage(stage)
	__Vector3__00000301_ := (&models.Vector3{Name: `TorusPoint 100`}).Stage(stage)
	__Vector3__00000302_ := (&models.Vector3{Name: `TorusPoint 101`}).Stage(stage)
	__Vector3__00000303_ := (&models.Vector3{Name: `TorusPoint 102`}).Stage(stage)
	__Vector3__00000304_ := (&models.Vector3{Name: `TorusPoint 103`}).Stage(stage)
	__Vector3__00000305_ := (&models.Vector3{Name: `TorusPoint 104`}).Stage(stage)
	__Vector3__00000306_ := (&models.Vector3{Name: `TorusPoint 105`}).Stage(stage)
	__Vector3__00000307_ := (&models.Vector3{Name: `TorusPoint 106`}).Stage(stage)
	__Vector3__00000308_ := (&models.Vector3{Name: `TorusPoint 107`}).Stage(stage)
	__Vector3__00000309_ := (&models.Vector3{Name: `TorusPoint 108`}).Stage(stage)
	__Vector3__00000310_ := (&models.Vector3{Name: `TorusPoint 109`}).Stage(stage)
	__Vector3__00000311_ := (&models.Vector3{Name: `TorusPoint 110`}).Stage(stage)
	__Vector3__00000312_ := (&models.Vector3{Name: `TorusPoint 111`}).Stage(stage)
	__Vector3__00000313_ := (&models.Vector3{Name: `TorusPoint 112`}).Stage(stage)
	__Vector3__00000314_ := (&models.Vector3{Name: `TorusPoint 113`}).Stage(stage)
	__Vector3__00000315_ := (&models.Vector3{Name: `TorusPoint 114`}).Stage(stage)
	__Vector3__00000316_ := (&models.Vector3{Name: `TorusPoint 115`}).Stage(stage)
	__Vector3__00000317_ := (&models.Vector3{Name: `TorusPoint 116`}).Stage(stage)
	__Vector3__00000318_ := (&models.Vector3{Name: `TorusPoint 117`}).Stage(stage)
	__Vector3__00000319_ := (&models.Vector3{Name: `TorusPoint 118`}).Stage(stage)
	__Vector3__00000320_ := (&models.Vector3{Name: `TorusPoint 119`}).Stage(stage)
	__Vector3__00000321_ := (&models.Vector3{Name: `TorusPoint 120`}).Stage(stage)
	__Vector3__00000322_ := (&models.Vector3{Name: `TorusPoint 121`}).Stage(stage)
	__Vector3__00000323_ := (&models.Vector3{Name: `TorusPoint 122`}).Stage(stage)
	__Vector3__00000324_ := (&models.Vector3{Name: `TorusPoint 123`}).Stage(stage)
	__Vector3__00000325_ := (&models.Vector3{Name: `TorusPoint 124`}).Stage(stage)
	__Vector3__00000326_ := (&models.Vector3{Name: `TorusPoint 125`}).Stage(stage)
	__Vector3__00000327_ := (&models.Vector3{Name: `TorusPoint 126`}).Stage(stage)
	__Vector3__00000328_ := (&models.Vector3{Name: `TorusPoint 127`}).Stage(stage)
	__Vector3__00000329_ := (&models.Vector3{Name: `TorusPoint 128`}).Stage(stage)
	__Vector3__00000330_ := (&models.Vector3{Name: `TorusPoint 129`}).Stage(stage)
	__Vector3__00000331_ := (&models.Vector3{Name: `TorusPoint 130`}).Stage(stage)
	__Vector3__00000332_ := (&models.Vector3{Name: `TorusPoint 131`}).Stage(stage)
	__Vector3__00000333_ := (&models.Vector3{Name: `TorusPoint 132`}).Stage(stage)
	__Vector3__00000334_ := (&models.Vector3{Name: `TorusPoint 133`}).Stage(stage)
	__Vector3__00000335_ := (&models.Vector3{Name: `TorusPoint 134`}).Stage(stage)
	__Vector3__00000336_ := (&models.Vector3{Name: `TorusPoint 135`}).Stage(stage)
	__Vector3__00000337_ := (&models.Vector3{Name: `TorusPoint 136`}).Stage(stage)
	__Vector3__00000338_ := (&models.Vector3{Name: `TorusPoint 137`}).Stage(stage)
	__Vector3__00000339_ := (&models.Vector3{Name: `TorusPoint 138`}).Stage(stage)
	__Vector3__00000340_ := (&models.Vector3{Name: `TorusPoint 139`}).Stage(stage)
	__Vector3__00000341_ := (&models.Vector3{Name: `TorusPoint 140`}).Stage(stage)
	__Vector3__00000342_ := (&models.Vector3{Name: `TorusPoint 141`}).Stage(stage)
	__Vector3__00000343_ := (&models.Vector3{Name: `TorusPoint 142`}).Stage(stage)
	__Vector3__00000344_ := (&models.Vector3{Name: `TorusPoint 143`}).Stage(stage)
	__Vector3__00000345_ := (&models.Vector3{Name: `TorusPoint 144`}).Stage(stage)
	__Vector3__00000346_ := (&models.Vector3{Name: `TorusPoint 145`}).Stage(stage)
	__Vector3__00000347_ := (&models.Vector3{Name: `TorusPoint 146`}).Stage(stage)
	__Vector3__00000348_ := (&models.Vector3{Name: `TorusPoint 147`}).Stage(stage)
	__Vector3__00000349_ := (&models.Vector3{Name: `TorusPoint 148`}).Stage(stage)
	__Vector3__00000350_ := (&models.Vector3{Name: `TorusPoint 149`}).Stage(stage)
	__Vector3__00000351_ := (&models.Vector3{Name: `TorusPoint 150`}).Stage(stage)
	__Vector3__00000352_ := (&models.Vector3{Name: `TorusPoint 151`}).Stage(stage)
	__Vector3__00000353_ := (&models.Vector3{Name: `TorusPoint 152`}).Stage(stage)
	__Vector3__00000354_ := (&models.Vector3{Name: `TorusPoint 153`}).Stage(stage)
	__Vector3__00000355_ := (&models.Vector3{Name: `TorusPoint 154`}).Stage(stage)
	__Vector3__00000356_ := (&models.Vector3{Name: `TorusPoint 155`}).Stage(stage)
	__Vector3__00000357_ := (&models.Vector3{Name: `TorusPoint 156`}).Stage(stage)
	__Vector3__00000358_ := (&models.Vector3{Name: `TorusPoint 157`}).Stage(stage)
	__Vector3__00000359_ := (&models.Vector3{Name: `TorusPoint 158`}).Stage(stage)
	__Vector3__00000360_ := (&models.Vector3{Name: `TorusPoint 159`}).Stage(stage)
	__Vector3__00000361_ := (&models.Vector3{Name: `TorusPoint 160`}).Stage(stage)
	__Vector3__00000362_ := (&models.Vector3{Name: `TorusPoint 161`}).Stage(stage)
	__Vector3__00000363_ := (&models.Vector3{Name: `TorusPoint 162`}).Stage(stage)
	__Vector3__00000364_ := (&models.Vector3{Name: `TorusPoint 163`}).Stage(stage)
	__Vector3__00000365_ := (&models.Vector3{Name: `TorusPoint 164`}).Stage(stage)
	__Vector3__00000366_ := (&models.Vector3{Name: `TorusPoint 165`}).Stage(stage)
	__Vector3__00000367_ := (&models.Vector3{Name: `TorusPoint 166`}).Stage(stage)
	__Vector3__00000368_ := (&models.Vector3{Name: `TorusPoint 167`}).Stage(stage)
	__Vector3__00000369_ := (&models.Vector3{Name: `TorusPoint 168`}).Stage(stage)
	__Vector3__00000370_ := (&models.Vector3{Name: `TorusPoint 169`}).Stage(stage)
	__Vector3__00000371_ := (&models.Vector3{Name: `TorusPoint 170`}).Stage(stage)
	__Vector3__00000372_ := (&models.Vector3{Name: `TorusPoint 171`}).Stage(stage)
	__Vector3__00000373_ := (&models.Vector3{Name: `TorusPoint 172`}).Stage(stage)
	__Vector3__00000374_ := (&models.Vector3{Name: `TorusPoint 173`}).Stage(stage)
	__Vector3__00000375_ := (&models.Vector3{Name: `TorusPoint 174`}).Stage(stage)
	__Vector3__00000376_ := (&models.Vector3{Name: `TorusPoint 175`}).Stage(stage)
	__Vector3__00000377_ := (&models.Vector3{Name: `TorusPoint 176`}).Stage(stage)
	__Vector3__00000378_ := (&models.Vector3{Name: `TorusPoint 177`}).Stage(stage)
	__Vector3__00000379_ := (&models.Vector3{Name: `TorusPoint 178`}).Stage(stage)
	__Vector3__00000380_ := (&models.Vector3{Name: `TorusPoint 179`}).Stage(stage)
	__Vector3__00000381_ := (&models.Vector3{Name: `TorusPoint 180`}).Stage(stage)
	__Vector3__00000382_ := (&models.Vector3{Name: `TorusPoint 181`}).Stage(stage)
	__Vector3__00000383_ := (&models.Vector3{Name: `TorusPoint 182`}).Stage(stage)
	__Vector3__00000384_ := (&models.Vector3{Name: `TorusPoint 183`}).Stage(stage)
	__Vector3__00000385_ := (&models.Vector3{Name: `TorusPoint 184`}).Stage(stage)
	__Vector3__00000386_ := (&models.Vector3{Name: `TorusPoint 185`}).Stage(stage)
	__Vector3__00000387_ := (&models.Vector3{Name: `TorusPoint 186`}).Stage(stage)
	__Vector3__00000388_ := (&models.Vector3{Name: `TorusPoint 187`}).Stage(stage)
	__Vector3__00000389_ := (&models.Vector3{Name: `TorusPoint 188`}).Stage(stage)
	__Vector3__00000390_ := (&models.Vector3{Name: `TorusPoint 189`}).Stage(stage)
	__Vector3__00000391_ := (&models.Vector3{Name: `TorusPoint 190`}).Stage(stage)
	__Vector3__00000392_ := (&models.Vector3{Name: `TorusPoint 191`}).Stage(stage)
	__Vector3__00000393_ := (&models.Vector3{Name: `TorusPoint 192`}).Stage(stage)
	__Vector3__00000394_ := (&models.Vector3{Name: `TorusPoint 193`}).Stage(stage)
	__Vector3__00000395_ := (&models.Vector3{Name: `TorusPoint 194`}).Stage(stage)
	__Vector3__00000396_ := (&models.Vector3{Name: `TorusPoint 195`}).Stage(stage)
	__Vector3__00000397_ := (&models.Vector3{Name: `TorusPoint 196`}).Stage(stage)
	__Vector3__00000398_ := (&models.Vector3{Name: `TorusPoint 197`}).Stage(stage)
	__Vector3__00000399_ := (&models.Vector3{Name: `TorusPoint 198`}).Stage(stage)
	__Vector3__00000400_ := (&models.Vector3{Name: `TorusPoint 199`}).Stage(stage)
	__Vector3__00000401_ := (&models.Vector3{Name: `TorusPoint 200`}).Stage(stage)
	__Vector3__00000402_ := (&models.Vector3{Name: `TorusPoint 201`}).Stage(stage)
	__Vector3__00000403_ := (&models.Vector3{Name: `TorusPoint 202`}).Stage(stage)
	__Vector3__00000404_ := (&models.Vector3{Name: `TorusPoint 203`}).Stage(stage)
	__Vector3__00000405_ := (&models.Vector3{Name: `TorusPoint 204`}).Stage(stage)
	__Vector3__00000406_ := (&models.Vector3{Name: `TorusPoint 205`}).Stage(stage)
	__Vector3__00000407_ := (&models.Vector3{Name: `TorusPoint 206`}).Stage(stage)
	__Vector3__00000408_ := (&models.Vector3{Name: `TorusPoint 207`}).Stage(stage)
	__Vector3__00000409_ := (&models.Vector3{Name: `TorusPoint 208`}).Stage(stage)
	__Vector3__00000410_ := (&models.Vector3{Name: `TorusPoint 209`}).Stage(stage)
	__Vector3__00000411_ := (&models.Vector3{Name: `TorusPoint 210`}).Stage(stage)
	__Vector3__00000412_ := (&models.Vector3{Name: `TorusPoint 211`}).Stage(stage)
	__Vector3__00000413_ := (&models.Vector3{Name: `TorusPoint 212`}).Stage(stage)
	__Vector3__00000414_ := (&models.Vector3{Name: `TorusPoint 213`}).Stage(stage)
	__Vector3__00000415_ := (&models.Vector3{Name: `TorusPoint 214`}).Stage(stage)
	__Vector3__00000416_ := (&models.Vector3{Name: `TorusPoint 215`}).Stage(stage)
	__Vector3__00000417_ := (&models.Vector3{Name: `TorusPoint 216`}).Stage(stage)
	__Vector3__00000418_ := (&models.Vector3{Name: `TorusPoint 217`}).Stage(stage)
	__Vector3__00000419_ := (&models.Vector3{Name: `TorusPoint 218`}).Stage(stage)
	__Vector3__00000420_ := (&models.Vector3{Name: `TorusPoint 219`}).Stage(stage)
	__Vector3__00000421_ := (&models.Vector3{Name: `TorusPoint 220`}).Stage(stage)
	__Vector3__00000422_ := (&models.Vector3{Name: `TorusPoint 221`}).Stage(stage)
	__Vector3__00000423_ := (&models.Vector3{Name: `TorusPoint 222`}).Stage(stage)
	__Vector3__00000424_ := (&models.Vector3{Name: `TorusPoint 223`}).Stage(stage)
	__Vector3__00000425_ := (&models.Vector3{Name: `TorusPoint 224`}).Stage(stage)
	__Vector3__00000426_ := (&models.Vector3{Name: `TorusPoint 225`}).Stage(stage)
	__Vector3__00000427_ := (&models.Vector3{Name: `TorusPoint 226`}).Stage(stage)
	__Vector3__00000428_ := (&models.Vector3{Name: `TorusPoint 227`}).Stage(stage)
	__Vector3__00000429_ := (&models.Vector3{Name: `TorusPoint 228`}).Stage(stage)
	__Vector3__00000430_ := (&models.Vector3{Name: `TorusPoint 229`}).Stage(stage)
	__Vector3__00000431_ := (&models.Vector3{Name: `TorusPoint 230`}).Stage(stage)
	__Vector3__00000432_ := (&models.Vector3{Name: `TorusPoint 231`}).Stage(stage)
	__Vector3__00000433_ := (&models.Vector3{Name: `TorusPoint 232`}).Stage(stage)
	__Vector3__00000434_ := (&models.Vector3{Name: `TorusPoint 233`}).Stage(stage)
	__Vector3__00000435_ := (&models.Vector3{Name: `TorusPoint 234`}).Stage(stage)
	__Vector3__00000436_ := (&models.Vector3{Name: `TorusPoint 235`}).Stage(stage)
	__Vector3__00000437_ := (&models.Vector3{Name: `TorusPoint 236`}).Stage(stage)
	__Vector3__00000438_ := (&models.Vector3{Name: `TorusPoint 237`}).Stage(stage)
	__Vector3__00000439_ := (&models.Vector3{Name: `TorusPoint 238`}).Stage(stage)
	__Vector3__00000440_ := (&models.Vector3{Name: `TorusPoint 239`}).Stage(stage)
	__Vector3__00000441_ := (&models.Vector3{Name: `TorusPoint 240`}).Stage(stage)
	__Vector3__00000442_ := (&models.Vector3{Name: `TorusPoint 241`}).Stage(stage)
	__Vector3__00000443_ := (&models.Vector3{Name: `TorusPoint 242`}).Stage(stage)
	__Vector3__00000444_ := (&models.Vector3{Name: `TorusPoint 243`}).Stage(stage)
	__Vector3__00000445_ := (&models.Vector3{Name: `TorusPoint 244`}).Stage(stage)
	__Vector3__00000446_ := (&models.Vector3{Name: `TorusPoint 245`}).Stage(stage)
	__Vector3__00000447_ := (&models.Vector3{Name: `TorusPoint 246`}).Stage(stage)
	__Vector3__00000448_ := (&models.Vector3{Name: `TorusPoint 247`}).Stage(stage)
	__Vector3__00000449_ := (&models.Vector3{Name: `TorusPoint 248`}).Stage(stage)
	__Vector3__00000450_ := (&models.Vector3{Name: `TorusPoint 249`}).Stage(stage)
	__Vector3__00000451_ := (&models.Vector3{Name: `TorusPoint 250`}).Stage(stage)
	__Vector3__00000452_ := (&models.Vector3{Name: `TorusPoint 251`}).Stage(stage)
	__Vector3__00000453_ := (&models.Vector3{Name: `TorusPoint 252`}).Stage(stage)
	__Vector3__00000454_ := (&models.Vector3{Name: `TorusPoint 253`}).Stage(stage)
	__Vector3__00000455_ := (&models.Vector3{Name: `TorusPoint 254`}).Stage(stage)
	__Vector3__00000456_ := (&models.Vector3{Name: `TorusPoint 255`}).Stage(stage)
	__Vector3__00000457_ := (&models.Vector3{Name: `TorusPoint 256`}).Stage(stage)
	__Vector3__00000458_ := (&models.Vector3{Name: `TorusPoint 257`}).Stage(stage)
	__Vector3__00000459_ := (&models.Vector3{Name: `TorusPoint 258`}).Stage(stage)
	__Vector3__00000460_ := (&models.Vector3{Name: `TorusPoint 259`}).Stage(stage)
	__Vector3__00000461_ := (&models.Vector3{Name: `TorusPoint 260`}).Stage(stage)
	__Vector3__00000462_ := (&models.Vector3{Name: `TorusPoint 261`}).Stage(stage)
	__Vector3__00000463_ := (&models.Vector3{Name: `TorusPoint 262`}).Stage(stage)
	__Vector3__00000464_ := (&models.Vector3{Name: `TorusPoint 263`}).Stage(stage)
	__Vector3__00000465_ := (&models.Vector3{Name: `TorusPoint 264`}).Stage(stage)
	__Vector3__00000466_ := (&models.Vector3{Name: `TorusPoint 265`}).Stage(stage)
	__Vector3__00000467_ := (&models.Vector3{Name: `TorusPoint 266`}).Stage(stage)
	__Vector3__00000468_ := (&models.Vector3{Name: `TorusPoint 267`}).Stage(stage)
	__Vector3__00000469_ := (&models.Vector3{Name: `TorusPoint 268`}).Stage(stage)
	__Vector3__00000470_ := (&models.Vector3{Name: `TorusPoint 269`}).Stage(stage)
	__Vector3__00000471_ := (&models.Vector3{Name: `TorusPoint 270`}).Stage(stage)
	__Vector3__00000472_ := (&models.Vector3{Name: `TorusPoint 271`}).Stage(stage)
	__Vector3__00000473_ := (&models.Vector3{Name: `TorusPoint 272`}).Stage(stage)
	__Vector3__00000474_ := (&models.Vector3{Name: `TorusPoint 273`}).Stage(stage)
	__Vector3__00000475_ := (&models.Vector3{Name: `TorusPoint 274`}).Stage(stage)
	__Vector3__00000476_ := (&models.Vector3{Name: `TorusPoint 275`}).Stage(stage)
	__Vector3__00000477_ := (&models.Vector3{Name: `TorusPoint 276`}).Stage(stage)
	__Vector3__00000478_ := (&models.Vector3{Name: `TorusPoint 277`}).Stage(stage)
	__Vector3__00000479_ := (&models.Vector3{Name: `TorusPoint 278`}).Stage(stage)
	__Vector3__00000480_ := (&models.Vector3{Name: `TorusPoint 279`}).Stage(stage)
	__Vector3__00000481_ := (&models.Vector3{Name: `TorusPoint 280`}).Stage(stage)
	__Vector3__00000482_ := (&models.Vector3{Name: `TorusPoint 281`}).Stage(stage)
	__Vector3__00000483_ := (&models.Vector3{Name: `TorusPoint 282`}).Stage(stage)
	__Vector3__00000484_ := (&models.Vector3{Name: `TorusPoint 283`}).Stage(stage)
	__Vector3__00000485_ := (&models.Vector3{Name: `TorusPoint 284`}).Stage(stage)
	__Vector3__00000486_ := (&models.Vector3{Name: `TorusPoint 285`}).Stage(stage)
	__Vector3__00000487_ := (&models.Vector3{Name: `TorusPoint 286`}).Stage(stage)
	__Vector3__00000488_ := (&models.Vector3{Name: `TorusPoint 287`}).Stage(stage)
	__Vector3__00000489_ := (&models.Vector3{Name: `TorusPoint 288`}).Stage(stage)
	__Vector3__00000490_ := (&models.Vector3{Name: `TorusPoint 289`}).Stage(stage)
	__Vector3__00000491_ := (&models.Vector3{Name: `TorusPoint 290`}).Stage(stage)
	__Vector3__00000492_ := (&models.Vector3{Name: `TorusPoint 291`}).Stage(stage)
	__Vector3__00000493_ := (&models.Vector3{Name: `TorusPoint 292`}).Stage(stage)
	__Vector3__00000494_ := (&models.Vector3{Name: `TorusPoint 293`}).Stage(stage)
	__Vector3__00000495_ := (&models.Vector3{Name: `TorusPoint 294`}).Stage(stage)
	__Vector3__00000496_ := (&models.Vector3{Name: `TorusPoint 295`}).Stage(stage)
	__Vector3__00000497_ := (&models.Vector3{Name: `TorusPoint 296`}).Stage(stage)
	__Vector3__00000498_ := (&models.Vector3{Name: `TorusPoint 297`}).Stage(stage)
	__Vector3__00000499_ := (&models.Vector3{Name: `TorusPoint 298`}).Stage(stage)
	__Vector3__00000500_ := (&models.Vector3{Name: `TorusPoint 299`}).Stage(stage)
	__Vector3__00000501_ := (&models.Vector3{Name: `TorusPoint 300`}).Stage(stage)
	__Vector3__00000502_ := (&models.Vector3{Name: `TorusPoint 301`}).Stage(stage)
	__Vector3__00000503_ := (&models.Vector3{Name: `TorusPoint 302`}).Stage(stage)
	__Vector3__00000504_ := (&models.Vector3{Name: `TorusPoint 303`}).Stage(stage)
	__Vector3__00000505_ := (&models.Vector3{Name: `TorusPoint 304`}).Stage(stage)
	__Vector3__00000506_ := (&models.Vector3{Name: `TorusPoint 305`}).Stage(stage)
	__Vector3__00000507_ := (&models.Vector3{Name: `TorusPoint 306`}).Stage(stage)
	__Vector3__00000508_ := (&models.Vector3{Name: `TorusPoint 307`}).Stage(stage)
	__Vector3__00000509_ := (&models.Vector3{Name: `TorusPoint 308`}).Stage(stage)
	__Vector3__00000510_ := (&models.Vector3{Name: `TorusPoint 309`}).Stage(stage)
	__Vector3__00000511_ := (&models.Vector3{Name: `TorusPoint 310`}).Stage(stage)
	__Vector3__00000512_ := (&models.Vector3{Name: `TorusPoint 311`}).Stage(stage)
	__Vector3__00000513_ := (&models.Vector3{Name: `TorusPoint 312`}).Stage(stage)
	__Vector3__00000514_ := (&models.Vector3{Name: `TorusPoint 313`}).Stage(stage)
	__Vector3__00000515_ := (&models.Vector3{Name: `TorusPoint 314`}).Stage(stage)
	__Vector3__00000516_ := (&models.Vector3{Name: `TorusPoint 315`}).Stage(stage)
	__Vector3__00000517_ := (&models.Vector3{Name: `TorusPoint 316`}).Stage(stage)
	__Vector3__00000518_ := (&models.Vector3{Name: `TorusPoint 317`}).Stage(stage)
	__Vector3__00000519_ := (&models.Vector3{Name: `TorusPoint 318`}).Stage(stage)
	__Vector3__00000520_ := (&models.Vector3{Name: `TorusPoint 319`}).Stage(stage)
	__Vector3__00000521_ := (&models.Vector3{Name: `TorusPoint 320`}).Stage(stage)
	__Vector3__00000522_ := (&models.Vector3{Name: `TorusPoint 321`}).Stage(stage)
	__Vector3__00000523_ := (&models.Vector3{Name: `TorusPoint 322`}).Stage(stage)
	__Vector3__00000524_ := (&models.Vector3{Name: `TorusPoint 323`}).Stage(stage)
	__Vector3__00000525_ := (&models.Vector3{Name: `TorusPoint 324`}).Stage(stage)
	__Vector3__00000526_ := (&models.Vector3{Name: `TorusPoint 325`}).Stage(stage)
	__Vector3__00000527_ := (&models.Vector3{Name: `TorusPoint 326`}).Stage(stage)
	__Vector3__00000528_ := (&models.Vector3{Name: `TorusPoint 327`}).Stage(stage)
	__Vector3__00000529_ := (&models.Vector3{Name: `TorusPoint 328`}).Stage(stage)
	__Vector3__00000530_ := (&models.Vector3{Name: `TorusPoint 329`}).Stage(stage)
	__Vector3__00000531_ := (&models.Vector3{Name: `TorusPoint 330`}).Stage(stage)
	__Vector3__00000532_ := (&models.Vector3{Name: `TorusPoint 331`}).Stage(stage)
	__Vector3__00000533_ := (&models.Vector3{Name: `TorusPoint 332`}).Stage(stage)
	__Vector3__00000534_ := (&models.Vector3{Name: `TorusPoint 333`}).Stage(stage)
	__Vector3__00000535_ := (&models.Vector3{Name: `TorusPoint 334`}).Stage(stage)
	__Vector3__00000536_ := (&models.Vector3{Name: `TorusPoint 335`}).Stage(stage)
	__Vector3__00000537_ := (&models.Vector3{Name: `TorusPoint 336`}).Stage(stage)
	__Vector3__00000538_ := (&models.Vector3{Name: `TorusPoint 337`}).Stage(stage)
	__Vector3__00000539_ := (&models.Vector3{Name: `TorusPoint 338`}).Stage(stage)
	__Vector3__00000540_ := (&models.Vector3{Name: `TorusPoint 339`}).Stage(stage)
	__Vector3__00000541_ := (&models.Vector3{Name: `TorusPoint 340`}).Stage(stage)
	__Vector3__00000542_ := (&models.Vector3{Name: `TorusPoint 341`}).Stage(stage)
	__Vector3__00000543_ := (&models.Vector3{Name: `TorusPoint 342`}).Stage(stage)
	__Vector3__00000544_ := (&models.Vector3{Name: `TorusPoint 343`}).Stage(stage)
	__Vector3__00000545_ := (&models.Vector3{Name: `TorusPoint 344`}).Stage(stage)
	__Vector3__00000546_ := (&models.Vector3{Name: `TorusPoint 345`}).Stage(stage)
	__Vector3__00000547_ := (&models.Vector3{Name: `TorusPoint 346`}).Stage(stage)
	__Vector3__00000548_ := (&models.Vector3{Name: `TorusPoint 347`}).Stage(stage)
	__Vector3__00000549_ := (&models.Vector3{Name: `TorusPoint 348`}).Stage(stage)
	__Vector3__00000550_ := (&models.Vector3{Name: `TorusPoint 349`}).Stage(stage)
	__Vector3__00000551_ := (&models.Vector3{Name: `TorusPoint 350`}).Stage(stage)
	__Vector3__00000552_ := (&models.Vector3{Name: `TorusPoint 351`}).Stage(stage)
	__Vector3__00000553_ := (&models.Vector3{Name: `TorusPoint 352`}).Stage(stage)
	__Vector3__00000554_ := (&models.Vector3{Name: `TorusPoint 353`}).Stage(stage)
	__Vector3__00000555_ := (&models.Vector3{Name: `TorusPoint 354`}).Stage(stage)
	__Vector3__00000556_ := (&models.Vector3{Name: `TorusPoint 355`}).Stage(stage)
	__Vector3__00000557_ := (&models.Vector3{Name: `TorusPoint 356`}).Stage(stage)
	__Vector3__00000558_ := (&models.Vector3{Name: `TorusPoint 357`}).Stage(stage)
	__Vector3__00000559_ := (&models.Vector3{Name: `TorusPoint 358`}).Stage(stage)
	__Vector3__00000560_ := (&models.Vector3{Name: `TorusPoint 359`}).Stage(stage)
	__Vector3__00000561_ := (&models.Vector3{Name: `TorusPoint 360`}).Stage(stage)
	__Vector3__00000562_ := (&models.Vector3{Name: `TorusPoint 361`}).Stage(stage)
	__Vector3__00000563_ := (&models.Vector3{Name: `TorusPoint 362`}).Stage(stage)
	__Vector3__00000564_ := (&models.Vector3{Name: `TorusPoint 363`}).Stage(stage)
	__Vector3__00000565_ := (&models.Vector3{Name: `TorusPoint 364`}).Stage(stage)
	__Vector3__00000566_ := (&models.Vector3{Name: `TorusPoint 365`}).Stage(stage)
	__Vector3__00000567_ := (&models.Vector3{Name: `TorusPoint 366`}).Stage(stage)
	__Vector3__00000568_ := (&models.Vector3{Name: `TorusPoint 367`}).Stage(stage)
	__Vector3__00000569_ := (&models.Vector3{Name: `TorusPoint 368`}).Stage(stage)
	__Vector3__00000570_ := (&models.Vector3{Name: `TorusPoint 369`}).Stage(stage)
	__Vector3__00000571_ := (&models.Vector3{Name: `TorusPoint 370`}).Stage(stage)
	__Vector3__00000572_ := (&models.Vector3{Name: `TorusPoint 371`}).Stage(stage)
	__Vector3__00000573_ := (&models.Vector3{Name: `TorusPoint 372`}).Stage(stage)
	__Vector3__00000574_ := (&models.Vector3{Name: `TorusPoint 373`}).Stage(stage)
	__Vector3__00000575_ := (&models.Vector3{Name: `TorusPoint 374`}).Stage(stage)
	__Vector3__00000576_ := (&models.Vector3{Name: `TorusPoint 375`}).Stage(stage)
	__Vector3__00000577_ := (&models.Vector3{Name: `TorusPoint 376`}).Stage(stage)
	__Vector3__00000578_ := (&models.Vector3{Name: `TorusPoint 377`}).Stage(stage)
	__Vector3__00000579_ := (&models.Vector3{Name: `TorusPoint 378`}).Stage(stage)
	__Vector3__00000580_ := (&models.Vector3{Name: `TorusPoint 379`}).Stage(stage)
	__Vector3__00000581_ := (&models.Vector3{Name: `TorusPoint 380`}).Stage(stage)
	__Vector3__00000582_ := (&models.Vector3{Name: `TorusPoint 381`}).Stage(stage)
	__Vector3__00000583_ := (&models.Vector3{Name: `TorusPoint 382`}).Stage(stage)
	__Vector3__00000584_ := (&models.Vector3{Name: `TorusPoint 383`}).Stage(stage)
	__Vector3__00000585_ := (&models.Vector3{Name: `TorusPoint 384`}).Stage(stage)
	__Vector3__00000586_ := (&models.Vector3{Name: `TorusPoint 385`}).Stage(stage)
	__Vector3__00000587_ := (&models.Vector3{Name: `TorusPoint 386`}).Stage(stage)
	__Vector3__00000588_ := (&models.Vector3{Name: `TorusPoint 387`}).Stage(stage)
	__Vector3__00000589_ := (&models.Vector3{Name: `TorusPoint 388`}).Stage(stage)
	__Vector3__00000590_ := (&models.Vector3{Name: `TorusPoint 389`}).Stage(stage)
	__Vector3__00000591_ := (&models.Vector3{Name: `TorusPoint 390`}).Stage(stage)
	__Vector3__00000592_ := (&models.Vector3{Name: `TorusPoint 391`}).Stage(stage)
	__Vector3__00000593_ := (&models.Vector3{Name: `TorusPoint 392`}).Stage(stage)
	__Vector3__00000594_ := (&models.Vector3{Name: `TorusPoint 393`}).Stage(stage)
	__Vector3__00000595_ := (&models.Vector3{Name: `TorusPoint 394`}).Stage(stage)
	__Vector3__00000596_ := (&models.Vector3{Name: `TorusPoint 395`}).Stage(stage)
	__Vector3__00000597_ := (&models.Vector3{Name: `TorusPoint 396`}).Stage(stage)
	__Vector3__00000598_ := (&models.Vector3{Name: `TorusPoint 397`}).Stage(stage)
	__Vector3__00000599_ := (&models.Vector3{Name: `TorusPoint 398`}).Stage(stage)
	__Vector3__00000600_ := (&models.Vector3{Name: `TorusPoint 399`}).Stage(stage)
	__Vector3__00000601_ := (&models.Vector3{Name: `TorusPoint 400`}).Stage(stage)
	__Vector3__00000602_ := (&models.Vector3{Name: `TorusPoint 401`}).Stage(stage)
	__Vector3__00000603_ := (&models.Vector3{Name: `TorusPoint 402`}).Stage(stage)
	__Vector3__00000604_ := (&models.Vector3{Name: `TorusPoint 403`}).Stage(stage)
	__Vector3__00000605_ := (&models.Vector3{Name: `TorusPoint 404`}).Stage(stage)
	__Vector3__00000606_ := (&models.Vector3{Name: `TorusPoint 405`}).Stage(stage)
	__Vector3__00000607_ := (&models.Vector3{Name: `TorusPoint 406`}).Stage(stage)
	__Vector3__00000608_ := (&models.Vector3{Name: `TorusPoint 407`}).Stage(stage)
	__Vector3__00000609_ := (&models.Vector3{Name: `TorusPoint 408`}).Stage(stage)
	__Vector3__00000610_ := (&models.Vector3{Name: `TorusPoint 409`}).Stage(stage)
	__Vector3__00000611_ := (&models.Vector3{Name: `TorusPoint 410`}).Stage(stage)
	__Vector3__00000612_ := (&models.Vector3{Name: `TorusPoint 411`}).Stage(stage)
	__Vector3__00000613_ := (&models.Vector3{Name: `TorusPoint 412`}).Stage(stage)
	__Vector3__00000614_ := (&models.Vector3{Name: `TorusPoint 413`}).Stage(stage)
	__Vector3__00000615_ := (&models.Vector3{Name: `TorusPoint 414`}).Stage(stage)
	__Vector3__00000616_ := (&models.Vector3{Name: `TorusPoint 415`}).Stage(stage)
	__Vector3__00000617_ := (&models.Vector3{Name: `TorusPoint 416`}).Stage(stage)
	__Vector3__00000618_ := (&models.Vector3{Name: `TorusPoint 417`}).Stage(stage)
	__Vector3__00000619_ := (&models.Vector3{Name: `TorusPoint 418`}).Stage(stage)
	__Vector3__00000620_ := (&models.Vector3{Name: `TorusPoint 419`}).Stage(stage)
	__Vector3__00000621_ := (&models.Vector3{Name: `TorusPoint 420`}).Stage(stage)
	__Vector3__00000622_ := (&models.Vector3{Name: `TorusPoint 421`}).Stage(stage)
	__Vector3__00000623_ := (&models.Vector3{Name: `TorusPoint 422`}).Stage(stage)
	__Vector3__00000624_ := (&models.Vector3{Name: `TorusPoint 423`}).Stage(stage)
	__Vector3__00000625_ := (&models.Vector3{Name: `TorusPoint 424`}).Stage(stage)
	__Vector3__00000626_ := (&models.Vector3{Name: `TorusPoint 425`}).Stage(stage)
	__Vector3__00000627_ := (&models.Vector3{Name: `TorusPoint 426`}).Stage(stage)
	__Vector3__00000628_ := (&models.Vector3{Name: `TorusPoint 427`}).Stage(stage)
	__Vector3__00000629_ := (&models.Vector3{Name: `TorusPoint 428`}).Stage(stage)
	__Vector3__00000630_ := (&models.Vector3{Name: `TorusPoint 429`}).Stage(stage)
	__Vector3__00000631_ := (&models.Vector3{Name: `TorusPoint 430`}).Stage(stage)
	__Vector3__00000632_ := (&models.Vector3{Name: `TorusPoint 431`}).Stage(stage)
	__Vector3__00000633_ := (&models.Vector3{Name: `TorusPoint 432`}).Stage(stage)
	__Vector3__00000634_ := (&models.Vector3{Name: `TorusPoint 433`}).Stage(stage)
	__Vector3__00000635_ := (&models.Vector3{Name: `TorusPoint 434`}).Stage(stage)
	__Vector3__00000636_ := (&models.Vector3{Name: `TorusPoint 435`}).Stage(stage)
	__Vector3__00000637_ := (&models.Vector3{Name: `TorusPoint 436`}).Stage(stage)
	__Vector3__00000638_ := (&models.Vector3{Name: `TorusPoint 437`}).Stage(stage)
	__Vector3__00000639_ := (&models.Vector3{Name: `TorusPoint 438`}).Stage(stage)
	__Vector3__00000640_ := (&models.Vector3{Name: `TorusPoint 439`}).Stage(stage)
	__Vector3__00000641_ := (&models.Vector3{Name: `TorusPoint 440`}).Stage(stage)
	__Vector3__00000642_ := (&models.Vector3{Name: `TorusPoint 441`}).Stage(stage)
	__Vector3__00000643_ := (&models.Vector3{Name: `TorusPoint 442`}).Stage(stage)
	__Vector3__00000644_ := (&models.Vector3{Name: `TorusPoint 443`}).Stage(stage)
	__Vector3__00000645_ := (&models.Vector3{Name: `TorusPoint 444`}).Stage(stage)
	__Vector3__00000646_ := (&models.Vector3{Name: `TorusPoint 445`}).Stage(stage)
	__Vector3__00000647_ := (&models.Vector3{Name: `TorusPoint 446`}).Stage(stage)
	__Vector3__00000648_ := (&models.Vector3{Name: `TorusPoint 447`}).Stage(stage)
	__Vector3__00000649_ := (&models.Vector3{Name: `TorusPoint 448`}).Stage(stage)
	__Vector3__00000650_ := (&models.Vector3{Name: `TorusPoint 449`}).Stage(stage)
	__Vector3__00000651_ := (&models.Vector3{Name: `TorusPoint 450`}).Stage(stage)
	__Vector3__00000652_ := (&models.Vector3{Name: `TorusPoint 451`}).Stage(stage)
	__Vector3__00000653_ := (&models.Vector3{Name: `TorusPoint 452`}).Stage(stage)
	__Vector3__00000654_ := (&models.Vector3{Name: `TorusPoint 453`}).Stage(stage)
	__Vector3__00000655_ := (&models.Vector3{Name: `TorusPoint 454`}).Stage(stage)
	__Vector3__00000656_ := (&models.Vector3{Name: `TorusPoint 455`}).Stage(stage)
	__Vector3__00000657_ := (&models.Vector3{Name: `TorusPoint 456`}).Stage(stage)
	__Vector3__00000658_ := (&models.Vector3{Name: `TorusPoint 457`}).Stage(stage)
	__Vector3__00000659_ := (&models.Vector3{Name: `TorusPoint 458`}).Stage(stage)
	__Vector3__00000660_ := (&models.Vector3{Name: `TorusPoint 459`}).Stage(stage)
	__Vector3__00000661_ := (&models.Vector3{Name: `TorusPoint 460`}).Stage(stage)
	__Vector3__00000662_ := (&models.Vector3{Name: `TorusPoint 461`}).Stage(stage)
	__Vector3__00000663_ := (&models.Vector3{Name: `TorusPoint 462`}).Stage(stage)
	__Vector3__00000664_ := (&models.Vector3{Name: `TorusPoint 463`}).Stage(stage)
	__Vector3__00000665_ := (&models.Vector3{Name: `TorusPoint 464`}).Stage(stage)
	__Vector3__00000666_ := (&models.Vector3{Name: `TorusPoint 465`}).Stage(stage)
	__Vector3__00000667_ := (&models.Vector3{Name: `TorusPoint 466`}).Stage(stage)
	__Vector3__00000668_ := (&models.Vector3{Name: `TorusPoint 467`}).Stage(stage)
	__Vector3__00000669_ := (&models.Vector3{Name: `TorusPoint 468`}).Stage(stage)
	__Vector3__00000670_ := (&models.Vector3{Name: `TorusPoint 469`}).Stage(stage)
	__Vector3__00000671_ := (&models.Vector3{Name: `TorusPoint 470`}).Stage(stage)
	__Vector3__00000672_ := (&models.Vector3{Name: `TorusPoint 471`}).Stage(stage)
	__Vector3__00000673_ := (&models.Vector3{Name: `TorusPoint 472`}).Stage(stage)
	__Vector3__00000674_ := (&models.Vector3{Name: `TorusPoint 473`}).Stage(stage)
	__Vector3__00000675_ := (&models.Vector3{Name: `TorusPoint 474`}).Stage(stage)
	__Vector3__00000676_ := (&models.Vector3{Name: `TorusPoint 475`}).Stage(stage)
	__Vector3__00000677_ := (&models.Vector3{Name: `TorusPoint 476`}).Stage(stage)
	__Vector3__00000678_ := (&models.Vector3{Name: `TorusPoint 477`}).Stage(stage)
	__Vector3__00000679_ := (&models.Vector3{Name: `TorusPoint 478`}).Stage(stage)
	__Vector3__00000680_ := (&models.Vector3{Name: `TorusPoint 479`}).Stage(stage)
	__Vector3__00000681_ := (&models.Vector3{Name: `TorusPoint 480`}).Stage(stage)
	__Vector3__00000682_ := (&models.Vector3{Name: `TorusPoint 481`}).Stage(stage)
	__Vector3__00000683_ := (&models.Vector3{Name: `TorusPoint 482`}).Stage(stage)
	__Vector3__00000684_ := (&models.Vector3{Name: `TorusPoint 483`}).Stage(stage)
	__Vector3__00000685_ := (&models.Vector3{Name: `TorusPoint 484`}).Stage(stage)
	__Vector3__00000686_ := (&models.Vector3{Name: `TorusPoint 485`}).Stage(stage)
	__Vector3__00000687_ := (&models.Vector3{Name: `TorusPoint 486`}).Stage(stage)
	__Vector3__00000688_ := (&models.Vector3{Name: `TorusPoint 487`}).Stage(stage)
	__Vector3__00000689_ := (&models.Vector3{Name: `TorusPoint 488`}).Stage(stage)
	__Vector3__00000690_ := (&models.Vector3{Name: `TorusPoint 489`}).Stage(stage)
	__Vector3__00000691_ := (&models.Vector3{Name: `TorusPoint 490`}).Stage(stage)
	__Vector3__00000692_ := (&models.Vector3{Name: `TorusPoint 491`}).Stage(stage)
	__Vector3__00000693_ := (&models.Vector3{Name: `TorusPoint 492`}).Stage(stage)
	__Vector3__00000694_ := (&models.Vector3{Name: `TorusPoint 493`}).Stage(stage)
	__Vector3__00000695_ := (&models.Vector3{Name: `TorusPoint 494`}).Stage(stage)
	__Vector3__00000696_ := (&models.Vector3{Name: `TorusPoint 495`}).Stage(stage)
	__Vector3__00000697_ := (&models.Vector3{Name: `TorusPoint 496`}).Stage(stage)
	__Vector3__00000698_ := (&models.Vector3{Name: `TorusPoint 497`}).Stage(stage)
	__Vector3__00000699_ := (&models.Vector3{Name: `TorusPoint 498`}).Stage(stage)
	__Vector3__00000700_ := (&models.Vector3{Name: `TorusPoint 499`}).Stage(stage)
	__Vector3__00000701_ := (&models.Vector3{Name: `TorusPoint 500`}).Stage(stage)

	// insertion point for initialization of values

	__AmbiantLight__00000000_.Name = `Ambiant Light`
	__AmbiantLight__00000000_.Intensity = 0.500000

	__BoxGeometry__00000000_.Name = `Box`
	__BoxGeometry__00000000_.Width = 5.000000
	__BoxGeometry__00000000_.Height = 5.000000
	__BoxGeometry__00000000_.Depth = 5.000000
	__BoxGeometry__00000000_.WidthSegments = 0
	__BoxGeometry__00000000_.HeightSegments = 0
	__BoxGeometry__00000000_.DepthSegments = 0

	__Canvas__00000000_.Name = `Singloton`

	__Curve__00000000_.Name = `Helix Path`

	__Curve__00000001_.Name = `Torus Like Curve`

	__CylinderGeometry__00000000_.Name = `Cylinder`
	__CylinderGeometry__00000000_.RadiusTop = 2.000000
	__CylinderGeometry__00000000_.RadiusBottom = 2.000000
	__CylinderGeometry__00000000_.Height = 5.000000
	__CylinderGeometry__00000000_.RadialSegments = 32
	__CylinderGeometry__00000000_.HeightSegments = 0
	__CylinderGeometry__00000000_.OpenEnded = false
	__CylinderGeometry__00000000_.ThetaStart = 0.000000
	__CylinderGeometry__00000000_.ThetaLength = 0.000000

	__CylinderGeometry__00000002_.Name = `Cylinder 2`
	__CylinderGeometry__00000002_.RadiusTop = 3.000000
	__CylinderGeometry__00000002_.RadiusBottom = 3.000000
	__CylinderGeometry__00000002_.Height = 2.000000
	__CylinderGeometry__00000002_.RadialSegments = 8
	__CylinderGeometry__00000002_.HeightSegments = 0
	__CylinderGeometry__00000002_.OpenEnded = false
	__CylinderGeometry__00000002_.ThetaStart = 0.000000
	__CylinderGeometry__00000002_.ThetaLength = 0.000000

	__CylinderGeometry__00000003_.Name = `Cylinder 3`
	__CylinderGeometry__00000003_.RadiusTop = 2.000000
	__CylinderGeometry__00000003_.RadiusBottom = 2.000000
	__CylinderGeometry__00000003_.Height = 4.000000
	__CylinderGeometry__00000003_.RadialSegments = 32
	__CylinderGeometry__00000003_.HeightSegments = 0
	__CylinderGeometry__00000003_.OpenEnded = false
	__CylinderGeometry__00000003_.ThetaStart = 0.000000
	__CylinderGeometry__00000003_.ThetaLength = 0.000000

	__DirectionalLight__00000000_.Name = `Direction Light`
	__DirectionalLight__00000000_.X = 10.000000
	__DirectionalLight__00000000_.Y = 20.000000
	__DirectionalLight__00000000_.Z = 10.000000
	__DirectionalLight__00000000_.Intensity = 1.500000
	__DirectionalLight__00000000_.IsWithCastShadow = true

	__ExtrudeGeometry__00000000_.Name = `Torus Extrude Geometry`
	__ExtrudeGeometry__00000000_.Steps = 500

	__Mesh__00000000_.Name = `Box Mesh`
	__Mesh__00000000_.X = 0.000000
	__Mesh__00000000_.Y = 0.000000
	__Mesh__00000000_.Z = 0.000000

	__Mesh__00000001_.Name = `Cylinder Mesh`
	__Mesh__00000001_.X = 8.000000
	__Mesh__00000001_.Y = 0.000000
	__Mesh__00000001_.Z = 0.000000

	__Mesh__00000002_.Name = `Sphere Mesh`
	__Mesh__00000002_.X = -8.000000
	__Mesh__00000002_.Y = 0.000000
	__Mesh__00000002_.Z = 0.000000

	__Mesh__00000003_.Name = `Torus Mesh`
	__Mesh__00000003_.X = 0.000000
	__Mesh__00000003_.Y = 5.000000
	__Mesh__00000003_.Z = 0.000000

	__Mesh__00000004_.Name = `Plane Mesh`
	__Mesh__00000004_.X = 4.000000
	__Mesh__00000004_.Y = -2.500000
	__Mesh__00000004_.Z = 0.000000

	__Mesh__00000005_.Name = `Cylinder Mesh 2`
	__Mesh__00000005_.X = 0.000000
	__Mesh__00000005_.Y = -1.500000
	__Mesh__00000005_.Z = -10.000000

	__Mesh__00000006_.Name = `Cylinder Mesh 3`
	__Mesh__00000006_.X = 0.000000
	__Mesh__00000006_.Y = 1.500000
	__Mesh__00000006_.Z = -10.000000

	__Mesh__00000007_.Name = `Helix Mesh`
	__Mesh__00000007_.X = 0.000000
	__Mesh__00000007_.Y = 3.500000
	__Mesh__00000007_.Z = -10.000000

	__Mesh__00000008_.Name = `Extrude Mesh`
	__Mesh__00000008_.X = 0.000000
	__Mesh__00000008_.Y = 5.000000
	__Mesh__00000008_.Z = 0.000000

	__MeshMaterialBasic__00000000_.Name = `Pink`
	__MeshMaterialBasic__00000000_.Color = `hotpink`

	__MeshMaterialBasic__00000001_.Name = `Gold`
	__MeshMaterialBasic__00000001_.Color = `gold`

	__MeshMaterialBasic__00000002_.Name = `White`
	__MeshMaterialBasic__00000002_.Color = `white`

	__MeshMaterialBasic__00000003_.Name = `Gray`
	__MeshMaterialBasic__00000003_.Color = `gray`

	__MeshMaterialBasic__00000004_.Name = `Light Gray`
	__MeshMaterialBasic__00000004_.Color = `lightgray`

	__MeshMaterialBasic__00000005_.Name = `Silver`
	__MeshMaterialBasic__00000005_.Color = `silver`

	__MeshMaterialBasic__00000006_.Name = `Magenta`
	__MeshMaterialBasic__00000006_.Color = `magenta`

	__MeshPhysicalMaterial__00000000_.Name = `Cyan Glass`
	__MeshPhysicalMaterial__00000000_.Color = `cyan`
	__MeshPhysicalMaterial__00000000_.Wireframe = false
	__MeshPhysicalMaterial__00000000_.Opacity = 1.000000
	__MeshPhysicalMaterial__00000000_.Transparent = true
	__MeshPhysicalMaterial__00000000_.Visible = false

	__PlaneGeometry__00000000_.Name = `Plane`
	__PlaneGeometry__00000000_.Width = 50.000000
	__PlaneGeometry__00000000_.Height = 50.000000
	__PlaneGeometry__00000000_.WidthSegments = 0
	__PlaneGeometry__00000000_.HeightSegments = 0

	__Shape__00000000_.Name = `Square Shape`

	__SphereGeometry__00000000_.Name = `Sphere`
	__SphereGeometry__00000000_.Radius = 2.500000
	__SphereGeometry__00000000_.WidthSegments = 32
	__SphereGeometry__00000000_.HeightSegments = 32
	__SphereGeometry__00000000_.PhiStart = 0.000000
	__SphereGeometry__00000000_.PhiLength = 0.000000
	__SphereGeometry__00000000_.ThetaStart = 0.000000
	__SphereGeometry__00000000_.ThetaLength = 0.000000

	__TorusGeometry__00000000_.Name = `Torus`
	__TorusGeometry__00000000_.Radius = 14.000000
	__TorusGeometry__00000000_.Tube = 0.100000
	__TorusGeometry__00000000_.RadialSegments = 16
	__TorusGeometry__00000000_.TubularSegments = 100
	__TorusGeometry__00000000_.Arc = 0.000000

	__TubeGeometry__00000000_.Name = `Helix Tube Geometry`
	__TubeGeometry__00000000_.TubularSegments = 200
	__TubeGeometry__00000000_.Radius = 0.400000
	__TubeGeometry__00000000_.RadialSegments = 16
	__TubeGeometry__00000000_.Closed = false

	__Vector2__00000000_.Name = `SquarePoint 0`
	__Vector2__00000000_.X = -0.750000
	__Vector2__00000000_.Y = -0.750000

	__Vector2__00000001_.Name = `SquarePoint 1`
	__Vector2__00000001_.X = 0.750000
	__Vector2__00000001_.Y = -0.750000

	__Vector2__00000002_.Name = `SquarePoint 2`
	__Vector2__00000002_.X = 0.750000
	__Vector2__00000002_.Y = 0.750000

	__Vector2__00000003_.Name = `SquarePoint 3`
	__Vector2__00000003_.X = -0.750000
	__Vector2__00000003_.Y = 0.750000

	__Vector2__00000004_.Name = `SquarePoint 4`
	__Vector2__00000004_.X = -0.750000
	__Vector2__00000004_.Y = -0.750000

	__Vector3__00000000_.Name = `Point 0`
	__Vector3__00000000_.X = 2.000000
	__Vector3__00000000_.Y = 0.000000
	__Vector3__00000000_.Z = 0.000000

	__Vector3__00000001_.Name = `Point 1`
	__Vector3__00000001_.X = 1.984229
	__Vector3__00000001_.Y = 0.030000
	__Vector3__00000001_.Z = 0.250666

	__Vector3__00000002_.Name = `Point 2`
	__Vector3__00000002_.X = 1.937166
	__Vector3__00000002_.Y = 0.060000
	__Vector3__00000002_.Z = 0.497380

	__Vector3__00000003_.Name = `Point 3`
	__Vector3__00000003_.X = 1.859553
	__Vector3__00000003_.Y = 0.090000
	__Vector3__00000003_.Z = 0.736249

	__Vector3__00000004_.Name = `Point 4`
	__Vector3__00000004_.X = 1.752613
	__Vector3__00000004_.Y = 0.120000
	__Vector3__00000004_.Z = 0.963507

	__Vector3__00000005_.Name = `Point 5`
	__Vector3__00000005_.X = 1.618034
	__Vector3__00000005_.Y = 0.150000
	__Vector3__00000005_.Z = 1.175571

	__Vector3__00000006_.Name = `Point 6`
	__Vector3__00000006_.X = 1.457937
	__Vector3__00000006_.Y = 0.180000
	__Vector3__00000006_.Z = 1.369094

	__Vector3__00000007_.Name = `Point 7`
	__Vector3__00000007_.X = 1.274848
	__Vector3__00000007_.Y = 0.210000
	__Vector3__00000007_.Z = 1.541026

	__Vector3__00000008_.Name = `Point 8`
	__Vector3__00000008_.X = 1.071654
	__Vector3__00000008_.Y = 0.240000
	__Vector3__00000008_.Z = 1.688656

	__Vector3__00000009_.Name = `Point 9`
	__Vector3__00000009_.X = 0.851559
	__Vector3__00000009_.Y = 0.270000
	__Vector3__00000009_.Z = 1.809654

	__Vector3__00000010_.Name = `Point 10`
	__Vector3__00000010_.X = 0.618034
	__Vector3__00000010_.Y = 0.300000
	__Vector3__00000010_.Z = 1.902113

	__Vector3__00000011_.Name = `Point 11`
	__Vector3__00000011_.X = 0.374763
	__Vector3__00000011_.Y = 0.330000
	__Vector3__00000011_.Z = 1.964575

	__Vector3__00000012_.Name = `Point 12`
	__Vector3__00000012_.X = 0.125581
	__Vector3__00000012_.Y = 0.360000
	__Vector3__00000012_.Z = 1.996053

	__Vector3__00000013_.Name = `Point 13`
	__Vector3__00000013_.X = -0.125581
	__Vector3__00000013_.Y = 0.390000
	__Vector3__00000013_.Z = 1.996053

	__Vector3__00000014_.Name = `Point 14`
	__Vector3__00000014_.X = -0.374763
	__Vector3__00000014_.Y = 0.420000
	__Vector3__00000014_.Z = 1.964575

	__Vector3__00000015_.Name = `Point 15`
	__Vector3__00000015_.X = -0.618034
	__Vector3__00000015_.Y = 0.450000
	__Vector3__00000015_.Z = 1.902113

	__Vector3__00000016_.Name = `Point 16`
	__Vector3__00000016_.X = -0.851559
	__Vector3__00000016_.Y = 0.480000
	__Vector3__00000016_.Z = 1.809654

	__Vector3__00000017_.Name = `Point 17`
	__Vector3__00000017_.X = -1.071654
	__Vector3__00000017_.Y = 0.510000
	__Vector3__00000017_.Z = 1.688656

	__Vector3__00000018_.Name = `Point 18`
	__Vector3__00000018_.X = -1.274848
	__Vector3__00000018_.Y = 0.540000
	__Vector3__00000018_.Z = 1.541026

	__Vector3__00000019_.Name = `Point 19`
	__Vector3__00000019_.X = -1.457937
	__Vector3__00000019_.Y = 0.570000
	__Vector3__00000019_.Z = 1.369094

	__Vector3__00000020_.Name = `Point 20`
	__Vector3__00000020_.X = -1.618034
	__Vector3__00000020_.Y = 0.600000
	__Vector3__00000020_.Z = 1.175571

	__Vector3__00000021_.Name = `Point 21`
	__Vector3__00000021_.X = -1.752613
	__Vector3__00000021_.Y = 0.630000
	__Vector3__00000021_.Z = 0.963507

	__Vector3__00000022_.Name = `Point 22`
	__Vector3__00000022_.X = -1.859553
	__Vector3__00000022_.Y = 0.660000
	__Vector3__00000022_.Z = 0.736249

	__Vector3__00000023_.Name = `Point 23`
	__Vector3__00000023_.X = -1.937166
	__Vector3__00000023_.Y = 0.690000
	__Vector3__00000023_.Z = 0.497380

	__Vector3__00000024_.Name = `Point 24`
	__Vector3__00000024_.X = -1.984229
	__Vector3__00000024_.Y = 0.720000
	__Vector3__00000024_.Z = 0.250666

	__Vector3__00000025_.Name = `Point 25`
	__Vector3__00000025_.X = -2.000000
	__Vector3__00000025_.Y = 0.750000
	__Vector3__00000025_.Z = 0.000000

	__Vector3__00000026_.Name = `Point 26`
	__Vector3__00000026_.X = -1.984229
	__Vector3__00000026_.Y = 0.780000
	__Vector3__00000026_.Z = -0.250666

	__Vector3__00000027_.Name = `Point 27`
	__Vector3__00000027_.X = -1.937166
	__Vector3__00000027_.Y = 0.810000
	__Vector3__00000027_.Z = -0.497380

	__Vector3__00000028_.Name = `Point 28`
	__Vector3__00000028_.X = -1.859553
	__Vector3__00000028_.Y = 0.840000
	__Vector3__00000028_.Z = -0.736249

	__Vector3__00000029_.Name = `Point 29`
	__Vector3__00000029_.X = -1.752613
	__Vector3__00000029_.Y = 0.870000
	__Vector3__00000029_.Z = -0.963507

	__Vector3__00000030_.Name = `Point 30`
	__Vector3__00000030_.X = -1.618034
	__Vector3__00000030_.Y = 0.900000
	__Vector3__00000030_.Z = -1.175571

	__Vector3__00000031_.Name = `Point 31`
	__Vector3__00000031_.X = -1.457937
	__Vector3__00000031_.Y = 0.930000
	__Vector3__00000031_.Z = -1.369094

	__Vector3__00000032_.Name = `Point 32`
	__Vector3__00000032_.X = -1.274848
	__Vector3__00000032_.Y = 0.960000
	__Vector3__00000032_.Z = -1.541026

	__Vector3__00000033_.Name = `Point 33`
	__Vector3__00000033_.X = -1.071654
	__Vector3__00000033_.Y = 0.990000
	__Vector3__00000033_.Z = -1.688656

	__Vector3__00000034_.Name = `Point 34`
	__Vector3__00000034_.X = -0.851559
	__Vector3__00000034_.Y = 1.020000
	__Vector3__00000034_.Z = -1.809654

	__Vector3__00000035_.Name = `Point 35`
	__Vector3__00000035_.X = -0.618034
	__Vector3__00000035_.Y = 1.050000
	__Vector3__00000035_.Z = -1.902113

	__Vector3__00000036_.Name = `Point 36`
	__Vector3__00000036_.X = -0.374763
	__Vector3__00000036_.Y = 1.080000
	__Vector3__00000036_.Z = -1.964575

	__Vector3__00000037_.Name = `Point 37`
	__Vector3__00000037_.X = -0.125581
	__Vector3__00000037_.Y = 1.110000
	__Vector3__00000037_.Z = -1.996053

	__Vector3__00000038_.Name = `Point 38`
	__Vector3__00000038_.X = 0.125581
	__Vector3__00000038_.Y = 1.140000
	__Vector3__00000038_.Z = -1.996053

	__Vector3__00000039_.Name = `Point 39`
	__Vector3__00000039_.X = 0.374763
	__Vector3__00000039_.Y = 1.170000
	__Vector3__00000039_.Z = -1.964575

	__Vector3__00000040_.Name = `Point 40`
	__Vector3__00000040_.X = 0.618034
	__Vector3__00000040_.Y = 1.200000
	__Vector3__00000040_.Z = -1.902113

	__Vector3__00000041_.Name = `Point 41`
	__Vector3__00000041_.X = 0.851559
	__Vector3__00000041_.Y = 1.230000
	__Vector3__00000041_.Z = -1.809654

	__Vector3__00000042_.Name = `Point 42`
	__Vector3__00000042_.X = 1.071654
	__Vector3__00000042_.Y = 1.260000
	__Vector3__00000042_.Z = -1.688656

	__Vector3__00000043_.Name = `Point 43`
	__Vector3__00000043_.X = 1.274848
	__Vector3__00000043_.Y = 1.290000
	__Vector3__00000043_.Z = -1.541026

	__Vector3__00000044_.Name = `Point 44`
	__Vector3__00000044_.X = 1.457937
	__Vector3__00000044_.Y = 1.320000
	__Vector3__00000044_.Z = -1.369094

	__Vector3__00000045_.Name = `Point 45`
	__Vector3__00000045_.X = 1.618034
	__Vector3__00000045_.Y = 1.350000
	__Vector3__00000045_.Z = -1.175571

	__Vector3__00000046_.Name = `Point 46`
	__Vector3__00000046_.X = 1.752613
	__Vector3__00000046_.Y = 1.380000
	__Vector3__00000046_.Z = -0.963507

	__Vector3__00000047_.Name = `Point 47`
	__Vector3__00000047_.X = 1.859553
	__Vector3__00000047_.Y = 1.410000
	__Vector3__00000047_.Z = -0.736249

	__Vector3__00000048_.Name = `Point 48`
	__Vector3__00000048_.X = 1.937166
	__Vector3__00000048_.Y = 1.440000
	__Vector3__00000048_.Z = -0.497380

	__Vector3__00000049_.Name = `Point 49`
	__Vector3__00000049_.X = 1.984229
	__Vector3__00000049_.Y = 1.470000
	__Vector3__00000049_.Z = -0.250666

	__Vector3__00000050_.Name = `Point 50`
	__Vector3__00000050_.X = 2.000000
	__Vector3__00000050_.Y = 1.500000
	__Vector3__00000050_.Z = -0.000000

	__Vector3__00000051_.Name = `Point 51`
	__Vector3__00000051_.X = 1.984229
	__Vector3__00000051_.Y = 1.530000
	__Vector3__00000051_.Z = 0.250666

	__Vector3__00000052_.Name = `Point 52`
	__Vector3__00000052_.X = 1.937166
	__Vector3__00000052_.Y = 1.560000
	__Vector3__00000052_.Z = 0.497380

	__Vector3__00000053_.Name = `Point 53`
	__Vector3__00000053_.X = 1.859553
	__Vector3__00000053_.Y = 1.590000
	__Vector3__00000053_.Z = 0.736249

	__Vector3__00000054_.Name = `Point 54`
	__Vector3__00000054_.X = 1.752613
	__Vector3__00000054_.Y = 1.620000
	__Vector3__00000054_.Z = 0.963507

	__Vector3__00000055_.Name = `Point 55`
	__Vector3__00000055_.X = 1.618034
	__Vector3__00000055_.Y = 1.650000
	__Vector3__00000055_.Z = 1.175571

	__Vector3__00000056_.Name = `Point 56`
	__Vector3__00000056_.X = 1.457937
	__Vector3__00000056_.Y = 1.680000
	__Vector3__00000056_.Z = 1.369094

	__Vector3__00000057_.Name = `Point 57`
	__Vector3__00000057_.X = 1.274848
	__Vector3__00000057_.Y = 1.710000
	__Vector3__00000057_.Z = 1.541026

	__Vector3__00000058_.Name = `Point 58`
	__Vector3__00000058_.X = 1.071654
	__Vector3__00000058_.Y = 1.740000
	__Vector3__00000058_.Z = 1.688656

	__Vector3__00000059_.Name = `Point 59`
	__Vector3__00000059_.X = 0.851559
	__Vector3__00000059_.Y = 1.770000
	__Vector3__00000059_.Z = 1.809654

	__Vector3__00000060_.Name = `Point 60`
	__Vector3__00000060_.X = 0.618034
	__Vector3__00000060_.Y = 1.800000
	__Vector3__00000060_.Z = 1.902113

	__Vector3__00000061_.Name = `Point 61`
	__Vector3__00000061_.X = 0.374763
	__Vector3__00000061_.Y = 1.830000
	__Vector3__00000061_.Z = 1.964575

	__Vector3__00000062_.Name = `Point 62`
	__Vector3__00000062_.X = 0.125581
	__Vector3__00000062_.Y = 1.860000
	__Vector3__00000062_.Z = 1.996053

	__Vector3__00000063_.Name = `Point 63`
	__Vector3__00000063_.X = -0.125581
	__Vector3__00000063_.Y = 1.890000
	__Vector3__00000063_.Z = 1.996053

	__Vector3__00000064_.Name = `Point 64`
	__Vector3__00000064_.X = -0.374763
	__Vector3__00000064_.Y = 1.920000
	__Vector3__00000064_.Z = 1.964575

	__Vector3__00000065_.Name = `Point 65`
	__Vector3__00000065_.X = -0.618034
	__Vector3__00000065_.Y = 1.950000
	__Vector3__00000065_.Z = 1.902113

	__Vector3__00000066_.Name = `Point 66`
	__Vector3__00000066_.X = -0.851559
	__Vector3__00000066_.Y = 1.980000
	__Vector3__00000066_.Z = 1.809654

	__Vector3__00000067_.Name = `Point 67`
	__Vector3__00000067_.X = -1.071654
	__Vector3__00000067_.Y = 2.010000
	__Vector3__00000067_.Z = 1.688656

	__Vector3__00000068_.Name = `Point 68`
	__Vector3__00000068_.X = -1.274848
	__Vector3__00000068_.Y = 2.040000
	__Vector3__00000068_.Z = 1.541026

	__Vector3__00000069_.Name = `Point 69`
	__Vector3__00000069_.X = -1.457937
	__Vector3__00000069_.Y = 2.070000
	__Vector3__00000069_.Z = 1.369094

	__Vector3__00000070_.Name = `Point 70`
	__Vector3__00000070_.X = -1.618034
	__Vector3__00000070_.Y = 2.100000
	__Vector3__00000070_.Z = 1.175571

	__Vector3__00000071_.Name = `Point 71`
	__Vector3__00000071_.X = -1.752613
	__Vector3__00000071_.Y = 2.130000
	__Vector3__00000071_.Z = 0.963507

	__Vector3__00000072_.Name = `Point 72`
	__Vector3__00000072_.X = -1.859553
	__Vector3__00000072_.Y = 2.160000
	__Vector3__00000072_.Z = 0.736249

	__Vector3__00000073_.Name = `Point 73`
	__Vector3__00000073_.X = -1.937166
	__Vector3__00000073_.Y = 2.190000
	__Vector3__00000073_.Z = 0.497380

	__Vector3__00000074_.Name = `Point 74`
	__Vector3__00000074_.X = -1.984229
	__Vector3__00000074_.Y = 2.220000
	__Vector3__00000074_.Z = 0.250666

	__Vector3__00000075_.Name = `Point 75`
	__Vector3__00000075_.X = -2.000000
	__Vector3__00000075_.Y = 2.250000
	__Vector3__00000075_.Z = 0.000000

	__Vector3__00000076_.Name = `Point 76`
	__Vector3__00000076_.X = -1.984229
	__Vector3__00000076_.Y = 2.280000
	__Vector3__00000076_.Z = -0.250666

	__Vector3__00000077_.Name = `Point 77`
	__Vector3__00000077_.X = -1.937166
	__Vector3__00000077_.Y = 2.310000
	__Vector3__00000077_.Z = -0.497380

	__Vector3__00000078_.Name = `Point 78`
	__Vector3__00000078_.X = -1.859553
	__Vector3__00000078_.Y = 2.340000
	__Vector3__00000078_.Z = -0.736249

	__Vector3__00000079_.Name = `Point 79`
	__Vector3__00000079_.X = -1.752613
	__Vector3__00000079_.Y = 2.370000
	__Vector3__00000079_.Z = -0.963507

	__Vector3__00000080_.Name = `Point 80`
	__Vector3__00000080_.X = -1.618034
	__Vector3__00000080_.Y = 2.400000
	__Vector3__00000080_.Z = -1.175571

	__Vector3__00000081_.Name = `Point 81`
	__Vector3__00000081_.X = -1.457937
	__Vector3__00000081_.Y = 2.430000
	__Vector3__00000081_.Z = -1.369094

	__Vector3__00000082_.Name = `Point 82`
	__Vector3__00000082_.X = -1.274848
	__Vector3__00000082_.Y = 2.460000
	__Vector3__00000082_.Z = -1.541026

	__Vector3__00000083_.Name = `Point 83`
	__Vector3__00000083_.X = -1.071654
	__Vector3__00000083_.Y = 2.490000
	__Vector3__00000083_.Z = -1.688656

	__Vector3__00000084_.Name = `Point 84`
	__Vector3__00000084_.X = -0.851559
	__Vector3__00000084_.Y = 2.520000
	__Vector3__00000084_.Z = -1.809654

	__Vector3__00000085_.Name = `Point 85`
	__Vector3__00000085_.X = -0.618034
	__Vector3__00000085_.Y = 2.550000
	__Vector3__00000085_.Z = -1.902113

	__Vector3__00000086_.Name = `Point 86`
	__Vector3__00000086_.X = -0.374763
	__Vector3__00000086_.Y = 2.580000
	__Vector3__00000086_.Z = -1.964575

	__Vector3__00000087_.Name = `Point 87`
	__Vector3__00000087_.X = -0.125581
	__Vector3__00000087_.Y = 2.610000
	__Vector3__00000087_.Z = -1.996053

	__Vector3__00000088_.Name = `Point 88`
	__Vector3__00000088_.X = 0.125581
	__Vector3__00000088_.Y = 2.640000
	__Vector3__00000088_.Z = -1.996053

	__Vector3__00000089_.Name = `Point 89`
	__Vector3__00000089_.X = 0.374763
	__Vector3__00000089_.Y = 2.670000
	__Vector3__00000089_.Z = -1.964575

	__Vector3__00000090_.Name = `Point 90`
	__Vector3__00000090_.X = 0.618034
	__Vector3__00000090_.Y = 2.700000
	__Vector3__00000090_.Z = -1.902113

	__Vector3__00000091_.Name = `Point 91`
	__Vector3__00000091_.X = 0.851559
	__Vector3__00000091_.Y = 2.730000
	__Vector3__00000091_.Z = -1.809654

	__Vector3__00000092_.Name = `Point 92`
	__Vector3__00000092_.X = 1.071654
	__Vector3__00000092_.Y = 2.760000
	__Vector3__00000092_.Z = -1.688656

	__Vector3__00000093_.Name = `Point 93`
	__Vector3__00000093_.X = 1.274848
	__Vector3__00000093_.Y = 2.790000
	__Vector3__00000093_.Z = -1.541026

	__Vector3__00000094_.Name = `Point 94`
	__Vector3__00000094_.X = 1.457937
	__Vector3__00000094_.Y = 2.820000
	__Vector3__00000094_.Z = -1.369094

	__Vector3__00000095_.Name = `Point 95`
	__Vector3__00000095_.X = 1.618034
	__Vector3__00000095_.Y = 2.850000
	__Vector3__00000095_.Z = -1.175571

	__Vector3__00000096_.Name = `Point 96`
	__Vector3__00000096_.X = 1.752613
	__Vector3__00000096_.Y = 2.880000
	__Vector3__00000096_.Z = -0.963507

	__Vector3__00000097_.Name = `Point 97`
	__Vector3__00000097_.X = 1.859553
	__Vector3__00000097_.Y = 2.910000
	__Vector3__00000097_.Z = -0.736249

	__Vector3__00000098_.Name = `Point 98`
	__Vector3__00000098_.X = 1.937166
	__Vector3__00000098_.Y = 2.940000
	__Vector3__00000098_.Z = -0.497380

	__Vector3__00000099_.Name = `Point 99`
	__Vector3__00000099_.X = 1.984229
	__Vector3__00000099_.Y = 2.970000
	__Vector3__00000099_.Z = -0.250666

	__Vector3__00000100_.Name = `Point 100`
	__Vector3__00000100_.X = 2.000000
	__Vector3__00000100_.Y = 3.000000
	__Vector3__00000100_.Z = -0.000000

	__Vector3__00000101_.Name = `Point 101`
	__Vector3__00000101_.X = 1.984229
	__Vector3__00000101_.Y = 3.030000
	__Vector3__00000101_.Z = 0.250666

	__Vector3__00000102_.Name = `Point 102`
	__Vector3__00000102_.X = 1.937166
	__Vector3__00000102_.Y = 3.060000
	__Vector3__00000102_.Z = 0.497380

	__Vector3__00000103_.Name = `Point 103`
	__Vector3__00000103_.X = 1.859553
	__Vector3__00000103_.Y = 3.090000
	__Vector3__00000103_.Z = 0.736249

	__Vector3__00000104_.Name = `Point 104`
	__Vector3__00000104_.X = 1.752613
	__Vector3__00000104_.Y = 3.120000
	__Vector3__00000104_.Z = 0.963507

	__Vector3__00000105_.Name = `Point 105`
	__Vector3__00000105_.X = 1.618034
	__Vector3__00000105_.Y = 3.150000
	__Vector3__00000105_.Z = 1.175571

	__Vector3__00000106_.Name = `Point 106`
	__Vector3__00000106_.X = 1.457937
	__Vector3__00000106_.Y = 3.180000
	__Vector3__00000106_.Z = 1.369094

	__Vector3__00000107_.Name = `Point 107`
	__Vector3__00000107_.X = 1.274848
	__Vector3__00000107_.Y = 3.210000
	__Vector3__00000107_.Z = 1.541026

	__Vector3__00000108_.Name = `Point 108`
	__Vector3__00000108_.X = 1.071654
	__Vector3__00000108_.Y = 3.240000
	__Vector3__00000108_.Z = 1.688656

	__Vector3__00000109_.Name = `Point 109`
	__Vector3__00000109_.X = 0.851559
	__Vector3__00000109_.Y = 3.270000
	__Vector3__00000109_.Z = 1.809654

	__Vector3__00000110_.Name = `Point 110`
	__Vector3__00000110_.X = 0.618034
	__Vector3__00000110_.Y = 3.300000
	__Vector3__00000110_.Z = 1.902113

	__Vector3__00000111_.Name = `Point 111`
	__Vector3__00000111_.X = 0.374763
	__Vector3__00000111_.Y = 3.330000
	__Vector3__00000111_.Z = 1.964575

	__Vector3__00000112_.Name = `Point 112`
	__Vector3__00000112_.X = 0.125581
	__Vector3__00000112_.Y = 3.360000
	__Vector3__00000112_.Z = 1.996053

	__Vector3__00000113_.Name = `Point 113`
	__Vector3__00000113_.X = -0.125581
	__Vector3__00000113_.Y = 3.390000
	__Vector3__00000113_.Z = 1.996053

	__Vector3__00000114_.Name = `Point 114`
	__Vector3__00000114_.X = -0.374763
	__Vector3__00000114_.Y = 3.420000
	__Vector3__00000114_.Z = 1.964575

	__Vector3__00000115_.Name = `Point 115`
	__Vector3__00000115_.X = -0.618034
	__Vector3__00000115_.Y = 3.450000
	__Vector3__00000115_.Z = 1.902113

	__Vector3__00000116_.Name = `Point 116`
	__Vector3__00000116_.X = -0.851559
	__Vector3__00000116_.Y = 3.480000
	__Vector3__00000116_.Z = 1.809654

	__Vector3__00000117_.Name = `Point 117`
	__Vector3__00000117_.X = -1.071654
	__Vector3__00000117_.Y = 3.510000
	__Vector3__00000117_.Z = 1.688656

	__Vector3__00000118_.Name = `Point 118`
	__Vector3__00000118_.X = -1.274848
	__Vector3__00000118_.Y = 3.540000
	__Vector3__00000118_.Z = 1.541026

	__Vector3__00000119_.Name = `Point 119`
	__Vector3__00000119_.X = -1.457937
	__Vector3__00000119_.Y = 3.570000
	__Vector3__00000119_.Z = 1.369094

	__Vector3__00000120_.Name = `Point 120`
	__Vector3__00000120_.X = -1.618034
	__Vector3__00000120_.Y = 3.600000
	__Vector3__00000120_.Z = 1.175571

	__Vector3__00000121_.Name = `Point 121`
	__Vector3__00000121_.X = -1.752613
	__Vector3__00000121_.Y = 3.630000
	__Vector3__00000121_.Z = 0.963507

	__Vector3__00000122_.Name = `Point 122`
	__Vector3__00000122_.X = -1.859553
	__Vector3__00000122_.Y = 3.660000
	__Vector3__00000122_.Z = 0.736249

	__Vector3__00000123_.Name = `Point 123`
	__Vector3__00000123_.X = -1.937166
	__Vector3__00000123_.Y = 3.690000
	__Vector3__00000123_.Z = 0.497380

	__Vector3__00000124_.Name = `Point 124`
	__Vector3__00000124_.X = -1.984229
	__Vector3__00000124_.Y = 3.720000
	__Vector3__00000124_.Z = 0.250666

	__Vector3__00000125_.Name = `Point 125`
	__Vector3__00000125_.X = -2.000000
	__Vector3__00000125_.Y = 3.750000
	__Vector3__00000125_.Z = 0.000000

	__Vector3__00000126_.Name = `Point 126`
	__Vector3__00000126_.X = -1.984229
	__Vector3__00000126_.Y = 3.780000
	__Vector3__00000126_.Z = -0.250666

	__Vector3__00000127_.Name = `Point 127`
	__Vector3__00000127_.X = -1.937166
	__Vector3__00000127_.Y = 3.810000
	__Vector3__00000127_.Z = -0.497380

	__Vector3__00000128_.Name = `Point 128`
	__Vector3__00000128_.X = -1.859553
	__Vector3__00000128_.Y = 3.840000
	__Vector3__00000128_.Z = -0.736249

	__Vector3__00000129_.Name = `Point 129`
	__Vector3__00000129_.X = -1.752613
	__Vector3__00000129_.Y = 3.870000
	__Vector3__00000129_.Z = -0.963507

	__Vector3__00000130_.Name = `Point 130`
	__Vector3__00000130_.X = -1.618034
	__Vector3__00000130_.Y = 3.900000
	__Vector3__00000130_.Z = -1.175571

	__Vector3__00000131_.Name = `Point 131`
	__Vector3__00000131_.X = -1.457937
	__Vector3__00000131_.Y = 3.930000
	__Vector3__00000131_.Z = -1.369094

	__Vector3__00000132_.Name = `Point 132`
	__Vector3__00000132_.X = -1.274848
	__Vector3__00000132_.Y = 3.960000
	__Vector3__00000132_.Z = -1.541026

	__Vector3__00000133_.Name = `Point 133`
	__Vector3__00000133_.X = -1.071654
	__Vector3__00000133_.Y = 3.990000
	__Vector3__00000133_.Z = -1.688656

	__Vector3__00000134_.Name = `Point 134`
	__Vector3__00000134_.X = -0.851559
	__Vector3__00000134_.Y = 4.020000
	__Vector3__00000134_.Z = -1.809654

	__Vector3__00000135_.Name = `Point 135`
	__Vector3__00000135_.X = -0.618034
	__Vector3__00000135_.Y = 4.050000
	__Vector3__00000135_.Z = -1.902113

	__Vector3__00000136_.Name = `Point 136`
	__Vector3__00000136_.X = -0.374763
	__Vector3__00000136_.Y = 4.080000
	__Vector3__00000136_.Z = -1.964575

	__Vector3__00000137_.Name = `Point 137`
	__Vector3__00000137_.X = -0.125581
	__Vector3__00000137_.Y = 4.110000
	__Vector3__00000137_.Z = -1.996053

	__Vector3__00000138_.Name = `Point 138`
	__Vector3__00000138_.X = 0.125581
	__Vector3__00000138_.Y = 4.140000
	__Vector3__00000138_.Z = -1.996053

	__Vector3__00000139_.Name = `Point 139`
	__Vector3__00000139_.X = 0.374763
	__Vector3__00000139_.Y = 4.170000
	__Vector3__00000139_.Z = -1.964575

	__Vector3__00000140_.Name = `Point 140`
	__Vector3__00000140_.X = 0.618034
	__Vector3__00000140_.Y = 4.200000
	__Vector3__00000140_.Z = -1.902113

	__Vector3__00000141_.Name = `Point 141`
	__Vector3__00000141_.X = 0.851559
	__Vector3__00000141_.Y = 4.230000
	__Vector3__00000141_.Z = -1.809654

	__Vector3__00000142_.Name = `Point 142`
	__Vector3__00000142_.X = 1.071654
	__Vector3__00000142_.Y = 4.260000
	__Vector3__00000142_.Z = -1.688656

	__Vector3__00000143_.Name = `Point 143`
	__Vector3__00000143_.X = 1.274848
	__Vector3__00000143_.Y = 4.290000
	__Vector3__00000143_.Z = -1.541026

	__Vector3__00000144_.Name = `Point 144`
	__Vector3__00000144_.X = 1.457937
	__Vector3__00000144_.Y = 4.320000
	__Vector3__00000144_.Z = -1.369094

	__Vector3__00000145_.Name = `Point 145`
	__Vector3__00000145_.X = 1.618034
	__Vector3__00000145_.Y = 4.350000
	__Vector3__00000145_.Z = -1.175571

	__Vector3__00000146_.Name = `Point 146`
	__Vector3__00000146_.X = 1.752613
	__Vector3__00000146_.Y = 4.380000
	__Vector3__00000146_.Z = -0.963507

	__Vector3__00000147_.Name = `Point 147`
	__Vector3__00000147_.X = 1.859553
	__Vector3__00000147_.Y = 4.410000
	__Vector3__00000147_.Z = -0.736249

	__Vector3__00000148_.Name = `Point 148`
	__Vector3__00000148_.X = 1.937166
	__Vector3__00000148_.Y = 4.440000
	__Vector3__00000148_.Z = -0.497380

	__Vector3__00000149_.Name = `Point 149`
	__Vector3__00000149_.X = 1.984229
	__Vector3__00000149_.Y = 4.470000
	__Vector3__00000149_.Z = -0.250666

	__Vector3__00000150_.Name = `Point 150`
	__Vector3__00000150_.X = 2.000000
	__Vector3__00000150_.Y = 4.500000
	__Vector3__00000150_.Z = -0.000000

	__Vector3__00000151_.Name = `Point 151`
	__Vector3__00000151_.X = 1.984229
	__Vector3__00000151_.Y = 4.530000
	__Vector3__00000151_.Z = 0.250666

	__Vector3__00000152_.Name = `Point 152`
	__Vector3__00000152_.X = 1.937166
	__Vector3__00000152_.Y = 4.560000
	__Vector3__00000152_.Z = 0.497380

	__Vector3__00000153_.Name = `Point 153`
	__Vector3__00000153_.X = 1.859553
	__Vector3__00000153_.Y = 4.590000
	__Vector3__00000153_.Z = 0.736249

	__Vector3__00000154_.Name = `Point 154`
	__Vector3__00000154_.X = 1.752613
	__Vector3__00000154_.Y = 4.620000
	__Vector3__00000154_.Z = 0.963507

	__Vector3__00000155_.Name = `Point 155`
	__Vector3__00000155_.X = 1.618034
	__Vector3__00000155_.Y = 4.650000
	__Vector3__00000155_.Z = 1.175571

	__Vector3__00000156_.Name = `Point 156`
	__Vector3__00000156_.X = 1.457937
	__Vector3__00000156_.Y = 4.680000
	__Vector3__00000156_.Z = 1.369094

	__Vector3__00000157_.Name = `Point 157`
	__Vector3__00000157_.X = 1.274848
	__Vector3__00000157_.Y = 4.710000
	__Vector3__00000157_.Z = 1.541026

	__Vector3__00000158_.Name = `Point 158`
	__Vector3__00000158_.X = 1.071654
	__Vector3__00000158_.Y = 4.740000
	__Vector3__00000158_.Z = 1.688656

	__Vector3__00000159_.Name = `Point 159`
	__Vector3__00000159_.X = 0.851559
	__Vector3__00000159_.Y = 4.770000
	__Vector3__00000159_.Z = 1.809654

	__Vector3__00000160_.Name = `Point 160`
	__Vector3__00000160_.X = 0.618034
	__Vector3__00000160_.Y = 4.800000
	__Vector3__00000160_.Z = 1.902113

	__Vector3__00000161_.Name = `Point 161`
	__Vector3__00000161_.X = 0.374763
	__Vector3__00000161_.Y = 4.830000
	__Vector3__00000161_.Z = 1.964575

	__Vector3__00000162_.Name = `Point 162`
	__Vector3__00000162_.X = 0.125581
	__Vector3__00000162_.Y = 4.860000
	__Vector3__00000162_.Z = 1.996053

	__Vector3__00000163_.Name = `Point 163`
	__Vector3__00000163_.X = -0.125581
	__Vector3__00000163_.Y = 4.890000
	__Vector3__00000163_.Z = 1.996053

	__Vector3__00000164_.Name = `Point 164`
	__Vector3__00000164_.X = -0.374763
	__Vector3__00000164_.Y = 4.920000
	__Vector3__00000164_.Z = 1.964575

	__Vector3__00000165_.Name = `Point 165`
	__Vector3__00000165_.X = -0.618034
	__Vector3__00000165_.Y = 4.950000
	__Vector3__00000165_.Z = 1.902113

	__Vector3__00000166_.Name = `Point 166`
	__Vector3__00000166_.X = -0.851559
	__Vector3__00000166_.Y = 4.980000
	__Vector3__00000166_.Z = 1.809654

	__Vector3__00000167_.Name = `Point 167`
	__Vector3__00000167_.X = -1.071654
	__Vector3__00000167_.Y = 5.010000
	__Vector3__00000167_.Z = 1.688656

	__Vector3__00000168_.Name = `Point 168`
	__Vector3__00000168_.X = -1.274848
	__Vector3__00000168_.Y = 5.040000
	__Vector3__00000168_.Z = 1.541026

	__Vector3__00000169_.Name = `Point 169`
	__Vector3__00000169_.X = -1.457937
	__Vector3__00000169_.Y = 5.070000
	__Vector3__00000169_.Z = 1.369094

	__Vector3__00000170_.Name = `Point 170`
	__Vector3__00000170_.X = -1.618034
	__Vector3__00000170_.Y = 5.100000
	__Vector3__00000170_.Z = 1.175571

	__Vector3__00000171_.Name = `Point 171`
	__Vector3__00000171_.X = -1.752613
	__Vector3__00000171_.Y = 5.130000
	__Vector3__00000171_.Z = 0.963507

	__Vector3__00000172_.Name = `Point 172`
	__Vector3__00000172_.X = -1.859553
	__Vector3__00000172_.Y = 5.160000
	__Vector3__00000172_.Z = 0.736249

	__Vector3__00000173_.Name = `Point 173`
	__Vector3__00000173_.X = -1.937166
	__Vector3__00000173_.Y = 5.190000
	__Vector3__00000173_.Z = 0.497380

	__Vector3__00000174_.Name = `Point 174`
	__Vector3__00000174_.X = -1.984229
	__Vector3__00000174_.Y = 5.220000
	__Vector3__00000174_.Z = 0.250666

	__Vector3__00000175_.Name = `Point 175`
	__Vector3__00000175_.X = -2.000000
	__Vector3__00000175_.Y = 5.250000
	__Vector3__00000175_.Z = 0.000000

	__Vector3__00000176_.Name = `Point 176`
	__Vector3__00000176_.X = -1.984229
	__Vector3__00000176_.Y = 5.280000
	__Vector3__00000176_.Z = -0.250666

	__Vector3__00000177_.Name = `Point 177`
	__Vector3__00000177_.X = -1.937166
	__Vector3__00000177_.Y = 5.310000
	__Vector3__00000177_.Z = -0.497380

	__Vector3__00000178_.Name = `Point 178`
	__Vector3__00000178_.X = -1.859553
	__Vector3__00000178_.Y = 5.340000
	__Vector3__00000178_.Z = -0.736249

	__Vector3__00000179_.Name = `Point 179`
	__Vector3__00000179_.X = -1.752613
	__Vector3__00000179_.Y = 5.370000
	__Vector3__00000179_.Z = -0.963507

	__Vector3__00000180_.Name = `Point 180`
	__Vector3__00000180_.X = -1.618034
	__Vector3__00000180_.Y = 5.400000
	__Vector3__00000180_.Z = -1.175571

	__Vector3__00000181_.Name = `Point 181`
	__Vector3__00000181_.X = -1.457937
	__Vector3__00000181_.Y = 5.430000
	__Vector3__00000181_.Z = -1.369094

	__Vector3__00000182_.Name = `Point 182`
	__Vector3__00000182_.X = -1.274848
	__Vector3__00000182_.Y = 5.460000
	__Vector3__00000182_.Z = -1.541026

	__Vector3__00000183_.Name = `Point 183`
	__Vector3__00000183_.X = -1.071654
	__Vector3__00000183_.Y = 5.490000
	__Vector3__00000183_.Z = -1.688656

	__Vector3__00000184_.Name = `Point 184`
	__Vector3__00000184_.X = -0.851559
	__Vector3__00000184_.Y = 5.520000
	__Vector3__00000184_.Z = -1.809654

	__Vector3__00000185_.Name = `Point 185`
	__Vector3__00000185_.X = -0.618034
	__Vector3__00000185_.Y = 5.550000
	__Vector3__00000185_.Z = -1.902113

	__Vector3__00000186_.Name = `Point 186`
	__Vector3__00000186_.X = -0.374763
	__Vector3__00000186_.Y = 5.580000
	__Vector3__00000186_.Z = -1.964575

	__Vector3__00000187_.Name = `Point 187`
	__Vector3__00000187_.X = -0.125581
	__Vector3__00000187_.Y = 5.610000
	__Vector3__00000187_.Z = -1.996053

	__Vector3__00000188_.Name = `Point 188`
	__Vector3__00000188_.X = 0.125581
	__Vector3__00000188_.Y = 5.640000
	__Vector3__00000188_.Z = -1.996053

	__Vector3__00000189_.Name = `Point 189`
	__Vector3__00000189_.X = 0.374763
	__Vector3__00000189_.Y = 5.670000
	__Vector3__00000189_.Z = -1.964575

	__Vector3__00000190_.Name = `Point 190`
	__Vector3__00000190_.X = 0.618034
	__Vector3__00000190_.Y = 5.700000
	__Vector3__00000190_.Z = -1.902113

	__Vector3__00000191_.Name = `Point 191`
	__Vector3__00000191_.X = 0.851559
	__Vector3__00000191_.Y = 5.730000
	__Vector3__00000191_.Z = -1.809654

	__Vector3__00000192_.Name = `Point 192`
	__Vector3__00000192_.X = 1.071654
	__Vector3__00000192_.Y = 5.760000
	__Vector3__00000192_.Z = -1.688656

	__Vector3__00000193_.Name = `Point 193`
	__Vector3__00000193_.X = 1.274848
	__Vector3__00000193_.Y = 5.790000
	__Vector3__00000193_.Z = -1.541026

	__Vector3__00000194_.Name = `Point 194`
	__Vector3__00000194_.X = 1.457937
	__Vector3__00000194_.Y = 5.820000
	__Vector3__00000194_.Z = -1.369094

	__Vector3__00000195_.Name = `Point 195`
	__Vector3__00000195_.X = 1.618034
	__Vector3__00000195_.Y = 5.850000
	__Vector3__00000195_.Z = -1.175571

	__Vector3__00000196_.Name = `Point 196`
	__Vector3__00000196_.X = 1.752613
	__Vector3__00000196_.Y = 5.880000
	__Vector3__00000196_.Z = -0.963507

	__Vector3__00000197_.Name = `Point 197`
	__Vector3__00000197_.X = 1.859553
	__Vector3__00000197_.Y = 5.910000
	__Vector3__00000197_.Z = -0.736249

	__Vector3__00000198_.Name = `Point 198`
	__Vector3__00000198_.X = 1.937166
	__Vector3__00000198_.Y = 5.940000
	__Vector3__00000198_.Z = -0.497380

	__Vector3__00000199_.Name = `Point 199`
	__Vector3__00000199_.X = 1.984229
	__Vector3__00000199_.Y = 5.970000
	__Vector3__00000199_.Z = -0.250666

	__Vector3__00000200_.Name = `Point 200`
	__Vector3__00000200_.X = 2.000000
	__Vector3__00000200_.Y = 6.000000
	__Vector3__00000200_.Z = -0.000000

	__Vector3__00000201_.Name = `TorusPoint 0`
	__Vector3__00000201_.X = 10.000000
	__Vector3__00000201_.Y = 0.000000
	__Vector3__00000201_.Z = 0.000000

	__Vector3__00000202_.Name = `TorusPoint 1`
	__Vector3__00000202_.X = 9.999210
	__Vector3__00000202_.Y = 0.125581
	__Vector3__00000202_.Z = 0.125660

	__Vector3__00000203_.Name = `TorusPoint 2`
	__Vector3__00000203_.X = 9.996842
	__Vector3__00000203_.Y = 0.250666
	__Vector3__00000203_.Z = 0.251301

	__Vector3__00000204_.Name = `TorusPoint 3`
	__Vector3__00000204_.X = 9.992895
	__Vector3__00000204_.Y = 0.374763
	__Vector3__00000204_.Z = 0.376902

	__Vector3__00000205_.Name = `TorusPoint 4`
	__Vector3__00000205_.X = 9.987370
	__Vector3__00000205_.Y = 0.497380
	__Vector3__00000205_.Z = 0.502443

	__Vector3__00000206_.Name = `TorusPoint 5`
	__Vector3__00000206_.X = 9.980267
	__Vector3__00000206_.Y = 0.618034
	__Vector3__00000206_.Z = 0.627905

	__Vector3__00000207_.Name = `TorusPoint 6`
	__Vector3__00000207_.X = 9.971589
	__Vector3__00000207_.Y = 0.736249
	__Vector3__00000207_.Z = 0.753268

	__Vector3__00000208_.Name = `TorusPoint 7`
	__Vector3__00000208_.X = 9.961336
	__Vector3__00000208_.Y = 0.851559
	__Vector3__00000208_.Z = 0.878512

	__Vector3__00000209_.Name = `TorusPoint 8`
	__Vector3__00000209_.X = 9.949510
	__Vector3__00000209_.Y = 0.963507
	__Vector3__00000209_.Z = 1.003617

	__Vector3__00000210_.Name = `TorusPoint 9`
	__Vector3__00000210_.X = 9.936113
	__Vector3__00000210_.Y = 1.071654
	__Vector3__00000210_.Z = 1.128564

	__Vector3__00000211_.Name = `TorusPoint 10`
	__Vector3__00000211_.X = 9.921147
	__Vector3__00000211_.Y = 1.175571
	__Vector3__00000211_.Z = 1.253332

	__Vector3__00000212_.Name = `TorusPoint 11`
	__Vector3__00000212_.X = 9.904614
	__Vector3__00000212_.Y = 1.274848
	__Vector3__00000212_.Z = 1.377903

	__Vector3__00000213_.Name = `TorusPoint 12`
	__Vector3__00000213_.X = 9.886517
	__Vector3__00000213_.Y = 1.369094
	__Vector3__00000213_.Z = 1.502256

	__Vector3__00000214_.Name = `TorusPoint 13`
	__Vector3__00000214_.X = 9.866859
	__Vector3__00000214_.Y = 1.457937
	__Vector3__00000214_.Z = 1.626372

	__Vector3__00000215_.Name = `TorusPoint 14`
	__Vector3__00000215_.X = 9.845643
	__Vector3__00000215_.Y = 1.541026
	__Vector3__00000215_.Z = 1.750231

	__Vector3__00000216_.Name = `TorusPoint 15`
	__Vector3__00000216_.X = 9.822873
	__Vector3__00000216_.Y = 1.618034
	__Vector3__00000216_.Z = 1.873813

	__Vector3__00000217_.Name = `TorusPoint 16`
	__Vector3__00000217_.X = 9.798551
	__Vector3__00000217_.Y = 1.688656
	__Vector3__00000217_.Z = 1.997100

	__Vector3__00000218_.Name = `TorusPoint 17`
	__Vector3__00000218_.X = 9.772681
	__Vector3__00000218_.Y = 1.752613
	__Vector3__00000218_.Z = 2.120071

	__Vector3__00000219_.Name = `TorusPoint 18`
	__Vector3__00000219_.X = 9.745269
	__Vector3__00000219_.Y = 1.809654
	__Vector3__00000219_.Z = 2.242708

	__Vector3__00000220_.Name = `TorusPoint 19`
	__Vector3__00000220_.X = 9.716317
	__Vector3__00000220_.Y = 1.859553
	__Vector3__00000220_.Z = 2.364990

	__Vector3__00000221_.Name = `TorusPoint 20`
	__Vector3__00000221_.X = 9.685832
	__Vector3__00000221_.Y = 1.902113
	__Vector3__00000221_.Z = 2.486899

	__Vector3__00000222_.Name = `TorusPoint 21`
	__Vector3__00000222_.X = 9.653816
	__Vector3__00000222_.Y = 1.937166
	__Vector3__00000222_.Z = 2.608415

	__Vector3__00000223_.Name = `TorusPoint 22`
	__Vector3__00000223_.X = 9.620277
	__Vector3__00000223_.Y = 1.964575
	__Vector3__00000223_.Z = 2.729519

	__Vector3__00000224_.Name = `TorusPoint 23`
	__Vector3__00000224_.X = 9.585218
	__Vector3__00000224_.Y = 1.984229
	__Vector3__00000224_.Z = 2.850193

	__Vector3__00000225_.Name = `TorusPoint 24`
	__Vector3__00000225_.X = 9.548645
	__Vector3__00000225_.Y = 1.996053
	__Vector3__00000225_.Z = 2.970416

	__Vector3__00000226_.Name = `TorusPoint 25`
	__Vector3__00000226_.X = 9.510565
	__Vector3__00000226_.Y = 2.000000
	__Vector3__00000226_.Z = 3.090170

	__Vector3__00000227_.Name = `TorusPoint 26`
	__Vector3__00000227_.X = 9.470983
	__Vector3__00000227_.Y = 1.996053
	__Vector3__00000227_.Z = 3.209436

	__Vector3__00000228_.Name = `TorusPoint 27`
	__Vector3__00000228_.X = 9.429905
	__Vector3__00000228_.Y = 1.984229
	__Vector3__00000228_.Z = 3.328195

	__Vector3__00000229_.Name = `TorusPoint 28`
	__Vector3__00000229_.X = 9.387339
	__Vector3__00000229_.Y = 1.964575
	__Vector3__00000229_.Z = 3.446429

	__Vector3__00000230_.Name = `TorusPoint 29`
	__Vector3__00000230_.X = 9.343289
	__Vector3__00000230_.Y = 1.937166
	__Vector3__00000230_.Z = 3.564119

	__Vector3__00000231_.Name = `TorusPoint 30`
	__Vector3__00000231_.X = 9.297765
	__Vector3__00000231_.Y = 1.902113
	__Vector3__00000231_.Z = 3.681246

	__Vector3__00000232_.Name = `TorusPoint 31`
	__Vector3__00000232_.X = 9.250772
	__Vector3__00000232_.Y = 1.859553
	__Vector3__00000232_.Z = 3.797791

	__Vector3__00000233_.Name = `TorusPoint 32`
	__Vector3__00000233_.X = 9.202318
	__Vector3__00000233_.Y = 1.809654
	__Vector3__00000233_.Z = 3.913737

	__Vector3__00000234_.Name = `TorusPoint 33`
	__Vector3__00000234_.X = 9.152412
	__Vector3__00000234_.Y = 1.752613
	__Vector3__00000234_.Z = 4.029064

	__Vector3__00000235_.Name = `TorusPoint 34`
	__Vector3__00000235_.X = 9.101060
	__Vector3__00000235_.Y = 1.688656
	__Vector3__00000235_.Z = 4.143756

	__Vector3__00000236_.Name = `TorusPoint 35`
	__Vector3__00000236_.X = 9.048271
	__Vector3__00000236_.Y = 1.618034
	__Vector3__00000236_.Z = 4.257793

	__Vector3__00000237_.Name = `TorusPoint 36`
	__Vector3__00000237_.X = 8.994053
	__Vector3__00000237_.Y = 1.541026
	__Vector3__00000237_.Z = 4.371158

	__Vector3__00000238_.Name = `TorusPoint 37`
	__Vector3__00000238_.X = 8.938414
	__Vector3__00000238_.Y = 1.457937
	__Vector3__00000238_.Z = 4.483832

	__Vector3__00000239_.Name = `TorusPoint 38`
	__Vector3__00000239_.X = 8.881364
	__Vector3__00000239_.Y = 1.369094
	__Vector3__00000239_.Z = 4.595799

	__Vector3__00000240_.Name = `TorusPoint 39`
	__Vector3__00000240_.X = 8.822912
	__Vector3__00000240_.Y = 1.274848
	__Vector3__00000240_.Z = 4.707039

	__Vector3__00000241_.Name = `TorusPoint 40`
	__Vector3__00000241_.X = 8.763067
	__Vector3__00000241_.Y = 1.175571
	__Vector3__00000241_.Z = 4.817537

	__Vector3__00000242_.Name = `TorusPoint 41`
	__Vector3__00000242_.X = 8.701838
	__Vector3__00000242_.Y = 1.071654
	__Vector3__00000242_.Z = 4.927273

	__Vector3__00000243_.Name = `TorusPoint 42`
	__Vector3__00000243_.X = 8.639234
	__Vector3__00000243_.Y = 0.963507
	__Vector3__00000243_.Z = 5.036232

	__Vector3__00000244_.Name = `TorusPoint 43`
	__Vector3__00000244_.X = 8.575267
	__Vector3__00000244_.Y = 0.851559
	__Vector3__00000244_.Z = 5.144395

	__Vector3__00000245_.Name = `TorusPoint 44`
	__Vector3__00000245_.X = 8.509945
	__Vector3__00000245_.Y = 0.736249
	__Vector3__00000245_.Z = 5.251746

	__Vector3__00000246_.Name = `TorusPoint 45`
	__Vector3__00000246_.X = 8.443279
	__Vector3__00000246_.Y = 0.618034
	__Vector3__00000246_.Z = 5.358268

	__Vector3__00000247_.Name = `TorusPoint 46`
	__Vector3__00000247_.X = 8.375280
	__Vector3__00000247_.Y = 0.497380
	__Vector3__00000247_.Z = 5.463943

	__Vector3__00000248_.Name = `TorusPoint 47`
	__Vector3__00000248_.X = 8.305959
	__Vector3__00000248_.Y = 0.374763
	__Vector3__00000248_.Z = 5.568756

	__Vector3__00000249_.Name = `TorusPoint 48`
	__Vector3__00000249_.X = 8.235326
	__Vector3__00000249_.Y = 0.250666
	__Vector3__00000249_.Z = 5.672689

	__Vector3__00000250_.Name = `TorusPoint 49`
	__Vector3__00000250_.X = 8.163393
	__Vector3__00000250_.Y = 0.125581
	__Vector3__00000250_.Z = 5.775727

	__Vector3__00000251_.Name = `TorusPoint 50`
	__Vector3__00000251_.X = 8.090170
	__Vector3__00000251_.Y = 0.000000
	__Vector3__00000251_.Z = 5.877853

	__Vector3__00000252_.Name = `TorusPoint 51`
	__Vector3__00000252_.X = 8.015670
	__Vector3__00000252_.Y = -0.125581
	__Vector3__00000252_.Z = 5.979050

	__Vector3__00000253_.Name = `TorusPoint 52`
	__Vector3__00000253_.X = 7.939904
	__Vector3__00000253_.Y = -0.250666
	__Vector3__00000253_.Z = 6.079303

	__Vector3__00000254_.Name = `TorusPoint 53`
	__Vector3__00000254_.X = 7.862884
	__Vector3__00000254_.Y = -0.374763
	__Vector3__00000254_.Z = 6.178596

	__Vector3__00000255_.Name = `TorusPoint 54`
	__Vector3__00000255_.X = 7.784623
	__Vector3__00000255_.Y = -0.497380
	__Vector3__00000255_.Z = 6.276914

	__Vector3__00000256_.Name = `TorusPoint 55`
	__Vector3__00000256_.X = 7.705132
	__Vector3__00000256_.Y = -0.618034
	__Vector3__00000256_.Z = 6.374240

	__Vector3__00000257_.Name = `TorusPoint 56`
	__Vector3__00000257_.X = 7.624425
	__Vector3__00000257_.Y = -0.736249
	__Vector3__00000257_.Z = 6.470560

	__Vector3__00000258_.Name = `TorusPoint 57`
	__Vector3__00000258_.X = 7.542514
	__Vector3__00000258_.Y = -0.851559
	__Vector3__00000258_.Z = 6.565858

	__Vector3__00000259_.Name = `TorusPoint 58`
	__Vector3__00000259_.X = 7.459411
	__Vector3__00000259_.Y = -0.963507
	__Vector3__00000259_.Z = 6.660119

	__Vector3__00000260_.Name = `TorusPoint 59`
	__Vector3__00000260_.X = 7.375131
	__Vector3__00000260_.Y = -1.071654
	__Vector3__00000260_.Z = 6.753328

	__Vector3__00000261_.Name = `TorusPoint 60`
	__Vector3__00000261_.X = 7.289686
	__Vector3__00000261_.Y = -1.175571
	__Vector3__00000261_.Z = 6.845471

	__Vector3__00000262_.Name = `TorusPoint 61`
	__Vector3__00000262_.X = 7.203090
	__Vector3__00000262_.Y = -1.274848
	__Vector3__00000262_.Z = 6.936533

	__Vector3__00000263_.Name = `TorusPoint 62`
	__Vector3__00000263_.X = 7.115357
	__Vector3__00000263_.Y = -1.369094
	__Vector3__00000263_.Z = 7.026500

	__Vector3__00000264_.Name = `TorusPoint 63`
	__Vector3__00000264_.X = 7.026500
	__Vector3__00000264_.Y = -1.457937
	__Vector3__00000264_.Z = 7.115357

	__Vector3__00000265_.Name = `TorusPoint 64`
	__Vector3__00000265_.X = 6.936533
	__Vector3__00000265_.Y = -1.541026
	__Vector3__00000265_.Z = 7.203090

	__Vector3__00000266_.Name = `TorusPoint 65`
	__Vector3__00000266_.X = 6.845471
	__Vector3__00000266_.Y = -1.618034
	__Vector3__00000266_.Z = 7.289686

	__Vector3__00000267_.Name = `TorusPoint 66`
	__Vector3__00000267_.X = 6.753328
	__Vector3__00000267_.Y = -1.688656
	__Vector3__00000267_.Z = 7.375131

	__Vector3__00000268_.Name = `TorusPoint 67`
	__Vector3__00000268_.X = 6.660119
	__Vector3__00000268_.Y = -1.752613
	__Vector3__00000268_.Z = 7.459411

	__Vector3__00000269_.Name = `TorusPoint 68`
	__Vector3__00000269_.X = 6.565858
	__Vector3__00000269_.Y = -1.809654
	__Vector3__00000269_.Z = 7.542514

	__Vector3__00000270_.Name = `TorusPoint 69`
	__Vector3__00000270_.X = 6.470560
	__Vector3__00000270_.Y = -1.859553
	__Vector3__00000270_.Z = 7.624425

	__Vector3__00000271_.Name = `TorusPoint 70`
	__Vector3__00000271_.X = 6.374240
	__Vector3__00000271_.Y = -1.902113
	__Vector3__00000271_.Z = 7.705132

	__Vector3__00000272_.Name = `TorusPoint 71`
	__Vector3__00000272_.X = 6.276914
	__Vector3__00000272_.Y = -1.937166
	__Vector3__00000272_.Z = 7.784623

	__Vector3__00000273_.Name = `TorusPoint 72`
	__Vector3__00000273_.X = 6.178596
	__Vector3__00000273_.Y = -1.964575
	__Vector3__00000273_.Z = 7.862884

	__Vector3__00000274_.Name = `TorusPoint 73`
	__Vector3__00000274_.X = 6.079303
	__Vector3__00000274_.Y = -1.984229
	__Vector3__00000274_.Z = 7.939904

	__Vector3__00000275_.Name = `TorusPoint 74`
	__Vector3__00000275_.X = 5.979050
	__Vector3__00000275_.Y = -1.996053
	__Vector3__00000275_.Z = 8.015670

	__Vector3__00000276_.Name = `TorusPoint 75`
	__Vector3__00000276_.X = 5.877853
	__Vector3__00000276_.Y = -2.000000
	__Vector3__00000276_.Z = 8.090170

	__Vector3__00000277_.Name = `TorusPoint 76`
	__Vector3__00000277_.X = 5.775727
	__Vector3__00000277_.Y = -1.996053
	__Vector3__00000277_.Z = 8.163393

	__Vector3__00000278_.Name = `TorusPoint 77`
	__Vector3__00000278_.X = 5.672689
	__Vector3__00000278_.Y = -1.984229
	__Vector3__00000278_.Z = 8.235326

	__Vector3__00000279_.Name = `TorusPoint 78`
	__Vector3__00000279_.X = 5.568756
	__Vector3__00000279_.Y = -1.964575
	__Vector3__00000279_.Z = 8.305959

	__Vector3__00000280_.Name = `TorusPoint 79`
	__Vector3__00000280_.X = 5.463943
	__Vector3__00000280_.Y = -1.937166
	__Vector3__00000280_.Z = 8.375280

	__Vector3__00000281_.Name = `TorusPoint 80`
	__Vector3__00000281_.X = 5.358268
	__Vector3__00000281_.Y = -1.902113
	__Vector3__00000281_.Z = 8.443279

	__Vector3__00000282_.Name = `TorusPoint 81`
	__Vector3__00000282_.X = 5.251746
	__Vector3__00000282_.Y = -1.859553
	__Vector3__00000282_.Z = 8.509945

	__Vector3__00000283_.Name = `TorusPoint 82`
	__Vector3__00000283_.X = 5.144395
	__Vector3__00000283_.Y = -1.809654
	__Vector3__00000283_.Z = 8.575267

	__Vector3__00000284_.Name = `TorusPoint 83`
	__Vector3__00000284_.X = 5.036232
	__Vector3__00000284_.Y = -1.752613
	__Vector3__00000284_.Z = 8.639234

	__Vector3__00000285_.Name = `TorusPoint 84`
	__Vector3__00000285_.X = 4.927273
	__Vector3__00000285_.Y = -1.688656
	__Vector3__00000285_.Z = 8.701838

	__Vector3__00000286_.Name = `TorusPoint 85`
	__Vector3__00000286_.X = 4.817537
	__Vector3__00000286_.Y = -1.618034
	__Vector3__00000286_.Z = 8.763067

	__Vector3__00000287_.Name = `TorusPoint 86`
	__Vector3__00000287_.X = 4.707039
	__Vector3__00000287_.Y = -1.541026
	__Vector3__00000287_.Z = 8.822912

	__Vector3__00000288_.Name = `TorusPoint 87`
	__Vector3__00000288_.X = 4.595799
	__Vector3__00000288_.Y = -1.457937
	__Vector3__00000288_.Z = 8.881364

	__Vector3__00000289_.Name = `TorusPoint 88`
	__Vector3__00000289_.X = 4.483832
	__Vector3__00000289_.Y = -1.369094
	__Vector3__00000289_.Z = 8.938414

	__Vector3__00000290_.Name = `TorusPoint 89`
	__Vector3__00000290_.X = 4.371158
	__Vector3__00000290_.Y = -1.274848
	__Vector3__00000290_.Z = 8.994053

	__Vector3__00000291_.Name = `TorusPoint 90`
	__Vector3__00000291_.X = 4.257793
	__Vector3__00000291_.Y = -1.175571
	__Vector3__00000291_.Z = 9.048271

	__Vector3__00000292_.Name = `TorusPoint 91`
	__Vector3__00000292_.X = 4.143756
	__Vector3__00000292_.Y = -1.071654
	__Vector3__00000292_.Z = 9.101060

	__Vector3__00000293_.Name = `TorusPoint 92`
	__Vector3__00000293_.X = 4.029064
	__Vector3__00000293_.Y = -0.963507
	__Vector3__00000293_.Z = 9.152412

	__Vector3__00000294_.Name = `TorusPoint 93`
	__Vector3__00000294_.X = 3.913737
	__Vector3__00000294_.Y = -0.851559
	__Vector3__00000294_.Z = 9.202318

	__Vector3__00000295_.Name = `TorusPoint 94`
	__Vector3__00000295_.X = 3.797791
	__Vector3__00000295_.Y = -0.736249
	__Vector3__00000295_.Z = 9.250772

	__Vector3__00000296_.Name = `TorusPoint 95`
	__Vector3__00000296_.X = 3.681246
	__Vector3__00000296_.Y = -0.618034
	__Vector3__00000296_.Z = 9.297765

	__Vector3__00000297_.Name = `TorusPoint 96`
	__Vector3__00000297_.X = 3.564119
	__Vector3__00000297_.Y = -0.497380
	__Vector3__00000297_.Z = 9.343289

	__Vector3__00000298_.Name = `TorusPoint 97`
	__Vector3__00000298_.X = 3.446429
	__Vector3__00000298_.Y = -0.374763
	__Vector3__00000298_.Z = 9.387339

	__Vector3__00000299_.Name = `TorusPoint 98`
	__Vector3__00000299_.X = 3.328195
	__Vector3__00000299_.Y = -0.250666
	__Vector3__00000299_.Z = 9.429905

	__Vector3__00000300_.Name = `TorusPoint 99`
	__Vector3__00000300_.X = 3.209436
	__Vector3__00000300_.Y = -0.125581
	__Vector3__00000300_.Z = 9.470983

	__Vector3__00000301_.Name = `TorusPoint 100`
	__Vector3__00000301_.X = 3.090170
	__Vector3__00000301_.Y = -0.000000
	__Vector3__00000301_.Z = 9.510565

	__Vector3__00000302_.Name = `TorusPoint 101`
	__Vector3__00000302_.X = 2.970416
	__Vector3__00000302_.Y = 0.125581
	__Vector3__00000302_.Z = 9.548645

	__Vector3__00000303_.Name = `TorusPoint 102`
	__Vector3__00000303_.X = 2.850193
	__Vector3__00000303_.Y = 0.250666
	__Vector3__00000303_.Z = 9.585218

	__Vector3__00000304_.Name = `TorusPoint 103`
	__Vector3__00000304_.X = 2.729519
	__Vector3__00000304_.Y = 0.374763
	__Vector3__00000304_.Z = 9.620277

	__Vector3__00000305_.Name = `TorusPoint 104`
	__Vector3__00000305_.X = 2.608415
	__Vector3__00000305_.Y = 0.497380
	__Vector3__00000305_.Z = 9.653816

	__Vector3__00000306_.Name = `TorusPoint 105`
	__Vector3__00000306_.X = 2.486899
	__Vector3__00000306_.Y = 0.618034
	__Vector3__00000306_.Z = 9.685832

	__Vector3__00000307_.Name = `TorusPoint 106`
	__Vector3__00000307_.X = 2.364990
	__Vector3__00000307_.Y = 0.736249
	__Vector3__00000307_.Z = 9.716317

	__Vector3__00000308_.Name = `TorusPoint 107`
	__Vector3__00000308_.X = 2.242708
	__Vector3__00000308_.Y = 0.851559
	__Vector3__00000308_.Z = 9.745269

	__Vector3__00000309_.Name = `TorusPoint 108`
	__Vector3__00000309_.X = 2.120071
	__Vector3__00000309_.Y = 0.963507
	__Vector3__00000309_.Z = 9.772681

	__Vector3__00000310_.Name = `TorusPoint 109`
	__Vector3__00000310_.X = 1.997100
	__Vector3__00000310_.Y = 1.071654
	__Vector3__00000310_.Z = 9.798551

	__Vector3__00000311_.Name = `TorusPoint 110`
	__Vector3__00000311_.X = 1.873813
	__Vector3__00000311_.Y = 1.175571
	__Vector3__00000311_.Z = 9.822873

	__Vector3__00000312_.Name = `TorusPoint 111`
	__Vector3__00000312_.X = 1.750231
	__Vector3__00000312_.Y = 1.274848
	__Vector3__00000312_.Z = 9.845643

	__Vector3__00000313_.Name = `TorusPoint 112`
	__Vector3__00000313_.X = 1.626372
	__Vector3__00000313_.Y = 1.369094
	__Vector3__00000313_.Z = 9.866859

	__Vector3__00000314_.Name = `TorusPoint 113`
	__Vector3__00000314_.X = 1.502256
	__Vector3__00000314_.Y = 1.457937
	__Vector3__00000314_.Z = 9.886517

	__Vector3__00000315_.Name = `TorusPoint 114`
	__Vector3__00000315_.X = 1.377903
	__Vector3__00000315_.Y = 1.541026
	__Vector3__00000315_.Z = 9.904614

	__Vector3__00000316_.Name = `TorusPoint 115`
	__Vector3__00000316_.X = 1.253332
	__Vector3__00000316_.Y = 1.618034
	__Vector3__00000316_.Z = 9.921147

	__Vector3__00000317_.Name = `TorusPoint 116`
	__Vector3__00000317_.X = 1.128564
	__Vector3__00000317_.Y = 1.688656
	__Vector3__00000317_.Z = 9.936113

	__Vector3__00000318_.Name = `TorusPoint 117`
	__Vector3__00000318_.X = 1.003617
	__Vector3__00000318_.Y = 1.752613
	__Vector3__00000318_.Z = 9.949510

	__Vector3__00000319_.Name = `TorusPoint 118`
	__Vector3__00000319_.X = 0.878512
	__Vector3__00000319_.Y = 1.809654
	__Vector3__00000319_.Z = 9.961336

	__Vector3__00000320_.Name = `TorusPoint 119`
	__Vector3__00000320_.X = 0.753268
	__Vector3__00000320_.Y = 1.859553
	__Vector3__00000320_.Z = 9.971589

	__Vector3__00000321_.Name = `TorusPoint 120`
	__Vector3__00000321_.X = 0.627905
	__Vector3__00000321_.Y = 1.902113
	__Vector3__00000321_.Z = 9.980267

	__Vector3__00000322_.Name = `TorusPoint 121`
	__Vector3__00000322_.X = 0.502443
	__Vector3__00000322_.Y = 1.937166
	__Vector3__00000322_.Z = 9.987370

	__Vector3__00000323_.Name = `TorusPoint 122`
	__Vector3__00000323_.X = 0.376902
	__Vector3__00000323_.Y = 1.964575
	__Vector3__00000323_.Z = 9.992895

	__Vector3__00000324_.Name = `TorusPoint 123`
	__Vector3__00000324_.X = 0.251301
	__Vector3__00000324_.Y = 1.984229
	__Vector3__00000324_.Z = 9.996842

	__Vector3__00000325_.Name = `TorusPoint 124`
	__Vector3__00000325_.X = 0.125660
	__Vector3__00000325_.Y = 1.996053
	__Vector3__00000325_.Z = 9.999210

	__Vector3__00000326_.Name = `TorusPoint 125`
	__Vector3__00000326_.X = 0.000000
	__Vector3__00000326_.Y = 2.000000
	__Vector3__00000326_.Z = 10.000000

	__Vector3__00000327_.Name = `TorusPoint 126`
	__Vector3__00000327_.X = -0.125660
	__Vector3__00000327_.Y = 1.996053
	__Vector3__00000327_.Z = 9.999210

	__Vector3__00000328_.Name = `TorusPoint 127`
	__Vector3__00000328_.X = -0.251301
	__Vector3__00000328_.Y = 1.984229
	__Vector3__00000328_.Z = 9.996842

	__Vector3__00000329_.Name = `TorusPoint 128`
	__Vector3__00000329_.X = -0.376902
	__Vector3__00000329_.Y = 1.964575
	__Vector3__00000329_.Z = 9.992895

	__Vector3__00000330_.Name = `TorusPoint 129`
	__Vector3__00000330_.X = -0.502443
	__Vector3__00000330_.Y = 1.937166
	__Vector3__00000330_.Z = 9.987370

	__Vector3__00000331_.Name = `TorusPoint 130`
	__Vector3__00000331_.X = -0.627905
	__Vector3__00000331_.Y = 1.902113
	__Vector3__00000331_.Z = 9.980267

	__Vector3__00000332_.Name = `TorusPoint 131`
	__Vector3__00000332_.X = -0.753268
	__Vector3__00000332_.Y = 1.859553
	__Vector3__00000332_.Z = 9.971589

	__Vector3__00000333_.Name = `TorusPoint 132`
	__Vector3__00000333_.X = -0.878512
	__Vector3__00000333_.Y = 1.809654
	__Vector3__00000333_.Z = 9.961336

	__Vector3__00000334_.Name = `TorusPoint 133`
	__Vector3__00000334_.X = -1.003617
	__Vector3__00000334_.Y = 1.752613
	__Vector3__00000334_.Z = 9.949510

	__Vector3__00000335_.Name = `TorusPoint 134`
	__Vector3__00000335_.X = -1.128564
	__Vector3__00000335_.Y = 1.688656
	__Vector3__00000335_.Z = 9.936113

	__Vector3__00000336_.Name = `TorusPoint 135`
	__Vector3__00000336_.X = -1.253332
	__Vector3__00000336_.Y = 1.618034
	__Vector3__00000336_.Z = 9.921147

	__Vector3__00000337_.Name = `TorusPoint 136`
	__Vector3__00000337_.X = -1.377903
	__Vector3__00000337_.Y = 1.541026
	__Vector3__00000337_.Z = 9.904614

	__Vector3__00000338_.Name = `TorusPoint 137`
	__Vector3__00000338_.X = -1.502256
	__Vector3__00000338_.Y = 1.457937
	__Vector3__00000338_.Z = 9.886517

	__Vector3__00000339_.Name = `TorusPoint 138`
	__Vector3__00000339_.X = -1.626372
	__Vector3__00000339_.Y = 1.369094
	__Vector3__00000339_.Z = 9.866859

	__Vector3__00000340_.Name = `TorusPoint 139`
	__Vector3__00000340_.X = -1.750231
	__Vector3__00000340_.Y = 1.274848
	__Vector3__00000340_.Z = 9.845643

	__Vector3__00000341_.Name = `TorusPoint 140`
	__Vector3__00000341_.X = -1.873813
	__Vector3__00000341_.Y = 1.175571
	__Vector3__00000341_.Z = 9.822873

	__Vector3__00000342_.Name = `TorusPoint 141`
	__Vector3__00000342_.X = -1.997100
	__Vector3__00000342_.Y = 1.071654
	__Vector3__00000342_.Z = 9.798551

	__Vector3__00000343_.Name = `TorusPoint 142`
	__Vector3__00000343_.X = -2.120071
	__Vector3__00000343_.Y = 0.963507
	__Vector3__00000343_.Z = 9.772681

	__Vector3__00000344_.Name = `TorusPoint 143`
	__Vector3__00000344_.X = -2.242708
	__Vector3__00000344_.Y = 0.851559
	__Vector3__00000344_.Z = 9.745269

	__Vector3__00000345_.Name = `TorusPoint 144`
	__Vector3__00000345_.X = -2.364990
	__Vector3__00000345_.Y = 0.736249
	__Vector3__00000345_.Z = 9.716317

	__Vector3__00000346_.Name = `TorusPoint 145`
	__Vector3__00000346_.X = -2.486899
	__Vector3__00000346_.Y = 0.618034
	__Vector3__00000346_.Z = 9.685832

	__Vector3__00000347_.Name = `TorusPoint 146`
	__Vector3__00000347_.X = -2.608415
	__Vector3__00000347_.Y = 0.497380
	__Vector3__00000347_.Z = 9.653816

	__Vector3__00000348_.Name = `TorusPoint 147`
	__Vector3__00000348_.X = -2.729519
	__Vector3__00000348_.Y = 0.374763
	__Vector3__00000348_.Z = 9.620277

	__Vector3__00000349_.Name = `TorusPoint 148`
	__Vector3__00000349_.X = -2.850193
	__Vector3__00000349_.Y = 0.250666
	__Vector3__00000349_.Z = 9.585218

	__Vector3__00000350_.Name = `TorusPoint 149`
	__Vector3__00000350_.X = -2.970416
	__Vector3__00000350_.Y = 0.125581
	__Vector3__00000350_.Z = 9.548645

	__Vector3__00000351_.Name = `TorusPoint 150`
	__Vector3__00000351_.X = -3.090170
	__Vector3__00000351_.Y = 0.000000
	__Vector3__00000351_.Z = 9.510565

	__Vector3__00000352_.Name = `TorusPoint 151`
	__Vector3__00000352_.X = -3.209436
	__Vector3__00000352_.Y = -0.125581
	__Vector3__00000352_.Z = 9.470983

	__Vector3__00000353_.Name = `TorusPoint 152`
	__Vector3__00000353_.X = -3.328195
	__Vector3__00000353_.Y = -0.250666
	__Vector3__00000353_.Z = 9.429905

	__Vector3__00000354_.Name = `TorusPoint 153`
	__Vector3__00000354_.X = -3.446429
	__Vector3__00000354_.Y = -0.374763
	__Vector3__00000354_.Z = 9.387339

	__Vector3__00000355_.Name = `TorusPoint 154`
	__Vector3__00000355_.X = -3.564119
	__Vector3__00000355_.Y = -0.497380
	__Vector3__00000355_.Z = 9.343289

	__Vector3__00000356_.Name = `TorusPoint 155`
	__Vector3__00000356_.X = -3.681246
	__Vector3__00000356_.Y = -0.618034
	__Vector3__00000356_.Z = 9.297765

	__Vector3__00000357_.Name = `TorusPoint 156`
	__Vector3__00000357_.X = -3.797791
	__Vector3__00000357_.Y = -0.736249
	__Vector3__00000357_.Z = 9.250772

	__Vector3__00000358_.Name = `TorusPoint 157`
	__Vector3__00000358_.X = -3.913737
	__Vector3__00000358_.Y = -0.851559
	__Vector3__00000358_.Z = 9.202318

	__Vector3__00000359_.Name = `TorusPoint 158`
	__Vector3__00000359_.X = -4.029064
	__Vector3__00000359_.Y = -0.963507
	__Vector3__00000359_.Z = 9.152412

	__Vector3__00000360_.Name = `TorusPoint 159`
	__Vector3__00000360_.X = -4.143756
	__Vector3__00000360_.Y = -1.071654
	__Vector3__00000360_.Z = 9.101060

	__Vector3__00000361_.Name = `TorusPoint 160`
	__Vector3__00000361_.X = -4.257793
	__Vector3__00000361_.Y = -1.175571
	__Vector3__00000361_.Z = 9.048271

	__Vector3__00000362_.Name = `TorusPoint 161`
	__Vector3__00000362_.X = -4.371158
	__Vector3__00000362_.Y = -1.274848
	__Vector3__00000362_.Z = 8.994053

	__Vector3__00000363_.Name = `TorusPoint 162`
	__Vector3__00000363_.X = -4.483832
	__Vector3__00000363_.Y = -1.369094
	__Vector3__00000363_.Z = 8.938414

	__Vector3__00000364_.Name = `TorusPoint 163`
	__Vector3__00000364_.X = -4.595799
	__Vector3__00000364_.Y = -1.457937
	__Vector3__00000364_.Z = 8.881364

	__Vector3__00000365_.Name = `TorusPoint 164`
	__Vector3__00000365_.X = -4.707039
	__Vector3__00000365_.Y = -1.541026
	__Vector3__00000365_.Z = 8.822912

	__Vector3__00000366_.Name = `TorusPoint 165`
	__Vector3__00000366_.X = -4.817537
	__Vector3__00000366_.Y = -1.618034
	__Vector3__00000366_.Z = 8.763067

	__Vector3__00000367_.Name = `TorusPoint 166`
	__Vector3__00000367_.X = -4.927273
	__Vector3__00000367_.Y = -1.688656
	__Vector3__00000367_.Z = 8.701838

	__Vector3__00000368_.Name = `TorusPoint 167`
	__Vector3__00000368_.X = -5.036232
	__Vector3__00000368_.Y = -1.752613
	__Vector3__00000368_.Z = 8.639234

	__Vector3__00000369_.Name = `TorusPoint 168`
	__Vector3__00000369_.X = -5.144395
	__Vector3__00000369_.Y = -1.809654
	__Vector3__00000369_.Z = 8.575267

	__Vector3__00000370_.Name = `TorusPoint 169`
	__Vector3__00000370_.X = -5.251746
	__Vector3__00000370_.Y = -1.859553
	__Vector3__00000370_.Z = 8.509945

	__Vector3__00000371_.Name = `TorusPoint 170`
	__Vector3__00000371_.X = -5.358268
	__Vector3__00000371_.Y = -1.902113
	__Vector3__00000371_.Z = 8.443279

	__Vector3__00000372_.Name = `TorusPoint 171`
	__Vector3__00000372_.X = -5.463943
	__Vector3__00000372_.Y = -1.937166
	__Vector3__00000372_.Z = 8.375280

	__Vector3__00000373_.Name = `TorusPoint 172`
	__Vector3__00000373_.X = -5.568756
	__Vector3__00000373_.Y = -1.964575
	__Vector3__00000373_.Z = 8.305959

	__Vector3__00000374_.Name = `TorusPoint 173`
	__Vector3__00000374_.X = -5.672689
	__Vector3__00000374_.Y = -1.984229
	__Vector3__00000374_.Z = 8.235326

	__Vector3__00000375_.Name = `TorusPoint 174`
	__Vector3__00000375_.X = -5.775727
	__Vector3__00000375_.Y = -1.996053
	__Vector3__00000375_.Z = 8.163393

	__Vector3__00000376_.Name = `TorusPoint 175`
	__Vector3__00000376_.X = -5.877853
	__Vector3__00000376_.Y = -2.000000
	__Vector3__00000376_.Z = 8.090170

	__Vector3__00000377_.Name = `TorusPoint 176`
	__Vector3__00000377_.X = -5.979050
	__Vector3__00000377_.Y = -1.996053
	__Vector3__00000377_.Z = 8.015670

	__Vector3__00000378_.Name = `TorusPoint 177`
	__Vector3__00000378_.X = -6.079303
	__Vector3__00000378_.Y = -1.984229
	__Vector3__00000378_.Z = 7.939904

	__Vector3__00000379_.Name = `TorusPoint 178`
	__Vector3__00000379_.X = -6.178596
	__Vector3__00000379_.Y = -1.964575
	__Vector3__00000379_.Z = 7.862884

	__Vector3__00000380_.Name = `TorusPoint 179`
	__Vector3__00000380_.X = -6.276914
	__Vector3__00000380_.Y = -1.937166
	__Vector3__00000380_.Z = 7.784623

	__Vector3__00000381_.Name = `TorusPoint 180`
	__Vector3__00000381_.X = -6.374240
	__Vector3__00000381_.Y = -1.902113
	__Vector3__00000381_.Z = 7.705132

	__Vector3__00000382_.Name = `TorusPoint 181`
	__Vector3__00000382_.X = -6.470560
	__Vector3__00000382_.Y = -1.859553
	__Vector3__00000382_.Z = 7.624425

	__Vector3__00000383_.Name = `TorusPoint 182`
	__Vector3__00000383_.X = -6.565858
	__Vector3__00000383_.Y = -1.809654
	__Vector3__00000383_.Z = 7.542514

	__Vector3__00000384_.Name = `TorusPoint 183`
	__Vector3__00000384_.X = -6.660119
	__Vector3__00000384_.Y = -1.752613
	__Vector3__00000384_.Z = 7.459411

	__Vector3__00000385_.Name = `TorusPoint 184`
	__Vector3__00000385_.X = -6.753328
	__Vector3__00000385_.Y = -1.688656
	__Vector3__00000385_.Z = 7.375131

	__Vector3__00000386_.Name = `TorusPoint 185`
	__Vector3__00000386_.X = -6.845471
	__Vector3__00000386_.Y = -1.618034
	__Vector3__00000386_.Z = 7.289686

	__Vector3__00000387_.Name = `TorusPoint 186`
	__Vector3__00000387_.X = -6.936533
	__Vector3__00000387_.Y = -1.541026
	__Vector3__00000387_.Z = 7.203090

	__Vector3__00000388_.Name = `TorusPoint 187`
	__Vector3__00000388_.X = -7.026500
	__Vector3__00000388_.Y = -1.457937
	__Vector3__00000388_.Z = 7.115357

	__Vector3__00000389_.Name = `TorusPoint 188`
	__Vector3__00000389_.X = -7.115357
	__Vector3__00000389_.Y = -1.369094
	__Vector3__00000389_.Z = 7.026500

	__Vector3__00000390_.Name = `TorusPoint 189`
	__Vector3__00000390_.X = -7.203090
	__Vector3__00000390_.Y = -1.274848
	__Vector3__00000390_.Z = 6.936533

	__Vector3__00000391_.Name = `TorusPoint 190`
	__Vector3__00000391_.X = -7.289686
	__Vector3__00000391_.Y = -1.175571
	__Vector3__00000391_.Z = 6.845471

	__Vector3__00000392_.Name = `TorusPoint 191`
	__Vector3__00000392_.X = -7.375131
	__Vector3__00000392_.Y = -1.071654
	__Vector3__00000392_.Z = 6.753328

	__Vector3__00000393_.Name = `TorusPoint 192`
	__Vector3__00000393_.X = -7.459411
	__Vector3__00000393_.Y = -0.963507
	__Vector3__00000393_.Z = 6.660119

	__Vector3__00000394_.Name = `TorusPoint 193`
	__Vector3__00000394_.X = -7.542514
	__Vector3__00000394_.Y = -0.851559
	__Vector3__00000394_.Z = 6.565858

	__Vector3__00000395_.Name = `TorusPoint 194`
	__Vector3__00000395_.X = -7.624425
	__Vector3__00000395_.Y = -0.736249
	__Vector3__00000395_.Z = 6.470560

	__Vector3__00000396_.Name = `TorusPoint 195`
	__Vector3__00000396_.X = -7.705132
	__Vector3__00000396_.Y = -0.618034
	__Vector3__00000396_.Z = 6.374240

	__Vector3__00000397_.Name = `TorusPoint 196`
	__Vector3__00000397_.X = -7.784623
	__Vector3__00000397_.Y = -0.497380
	__Vector3__00000397_.Z = 6.276914

	__Vector3__00000398_.Name = `TorusPoint 197`
	__Vector3__00000398_.X = -7.862884
	__Vector3__00000398_.Y = -0.374763
	__Vector3__00000398_.Z = 6.178596

	__Vector3__00000399_.Name = `TorusPoint 198`
	__Vector3__00000399_.X = -7.939904
	__Vector3__00000399_.Y = -0.250666
	__Vector3__00000399_.Z = 6.079303

	__Vector3__00000400_.Name = `TorusPoint 199`
	__Vector3__00000400_.X = -8.015670
	__Vector3__00000400_.Y = -0.125581
	__Vector3__00000400_.Z = 5.979050

	__Vector3__00000401_.Name = `TorusPoint 200`
	__Vector3__00000401_.X = -8.090170
	__Vector3__00000401_.Y = -0.000000
	__Vector3__00000401_.Z = 5.877853

	__Vector3__00000402_.Name = `TorusPoint 201`
	__Vector3__00000402_.X = -8.163393
	__Vector3__00000402_.Y = 0.125581
	__Vector3__00000402_.Z = 5.775727

	__Vector3__00000403_.Name = `TorusPoint 202`
	__Vector3__00000403_.X = -8.235326
	__Vector3__00000403_.Y = 0.250666
	__Vector3__00000403_.Z = 5.672689

	__Vector3__00000404_.Name = `TorusPoint 203`
	__Vector3__00000404_.X = -8.305959
	__Vector3__00000404_.Y = 0.374763
	__Vector3__00000404_.Z = 5.568756

	__Vector3__00000405_.Name = `TorusPoint 204`
	__Vector3__00000405_.X = -8.375280
	__Vector3__00000405_.Y = 0.497380
	__Vector3__00000405_.Z = 5.463943

	__Vector3__00000406_.Name = `TorusPoint 205`
	__Vector3__00000406_.X = -8.443279
	__Vector3__00000406_.Y = 0.618034
	__Vector3__00000406_.Z = 5.358268

	__Vector3__00000407_.Name = `TorusPoint 206`
	__Vector3__00000407_.X = -8.509945
	__Vector3__00000407_.Y = 0.736249
	__Vector3__00000407_.Z = 5.251746

	__Vector3__00000408_.Name = `TorusPoint 207`
	__Vector3__00000408_.X = -8.575267
	__Vector3__00000408_.Y = 0.851559
	__Vector3__00000408_.Z = 5.144395

	__Vector3__00000409_.Name = `TorusPoint 208`
	__Vector3__00000409_.X = -8.639234
	__Vector3__00000409_.Y = 0.963507
	__Vector3__00000409_.Z = 5.036232

	__Vector3__00000410_.Name = `TorusPoint 209`
	__Vector3__00000410_.X = -8.701838
	__Vector3__00000410_.Y = 1.071654
	__Vector3__00000410_.Z = 4.927273

	__Vector3__00000411_.Name = `TorusPoint 210`
	__Vector3__00000411_.X = -8.763067
	__Vector3__00000411_.Y = 1.175571
	__Vector3__00000411_.Z = 4.817537

	__Vector3__00000412_.Name = `TorusPoint 211`
	__Vector3__00000412_.X = -8.822912
	__Vector3__00000412_.Y = 1.274848
	__Vector3__00000412_.Z = 4.707039

	__Vector3__00000413_.Name = `TorusPoint 212`
	__Vector3__00000413_.X = -8.881364
	__Vector3__00000413_.Y = 1.369094
	__Vector3__00000413_.Z = 4.595799

	__Vector3__00000414_.Name = `TorusPoint 213`
	__Vector3__00000414_.X = -8.938414
	__Vector3__00000414_.Y = 1.457937
	__Vector3__00000414_.Z = 4.483832

	__Vector3__00000415_.Name = `TorusPoint 214`
	__Vector3__00000415_.X = -8.994053
	__Vector3__00000415_.Y = 1.541026
	__Vector3__00000415_.Z = 4.371158

	__Vector3__00000416_.Name = `TorusPoint 215`
	__Vector3__00000416_.X = -9.048271
	__Vector3__00000416_.Y = 1.618034
	__Vector3__00000416_.Z = 4.257793

	__Vector3__00000417_.Name = `TorusPoint 216`
	__Vector3__00000417_.X = -9.101060
	__Vector3__00000417_.Y = 1.688656
	__Vector3__00000417_.Z = 4.143756

	__Vector3__00000418_.Name = `TorusPoint 217`
	__Vector3__00000418_.X = -9.152412
	__Vector3__00000418_.Y = 1.752613
	__Vector3__00000418_.Z = 4.029064

	__Vector3__00000419_.Name = `TorusPoint 218`
	__Vector3__00000419_.X = -9.202318
	__Vector3__00000419_.Y = 1.809654
	__Vector3__00000419_.Z = 3.913737

	__Vector3__00000420_.Name = `TorusPoint 219`
	__Vector3__00000420_.X = -9.250772
	__Vector3__00000420_.Y = 1.859553
	__Vector3__00000420_.Z = 3.797791

	__Vector3__00000421_.Name = `TorusPoint 220`
	__Vector3__00000421_.X = -9.297765
	__Vector3__00000421_.Y = 1.902113
	__Vector3__00000421_.Z = 3.681246

	__Vector3__00000422_.Name = `TorusPoint 221`
	__Vector3__00000422_.X = -9.343289
	__Vector3__00000422_.Y = 1.937166
	__Vector3__00000422_.Z = 3.564119

	__Vector3__00000423_.Name = `TorusPoint 222`
	__Vector3__00000423_.X = -9.387339
	__Vector3__00000423_.Y = 1.964575
	__Vector3__00000423_.Z = 3.446429

	__Vector3__00000424_.Name = `TorusPoint 223`
	__Vector3__00000424_.X = -9.429905
	__Vector3__00000424_.Y = 1.984229
	__Vector3__00000424_.Z = 3.328195

	__Vector3__00000425_.Name = `TorusPoint 224`
	__Vector3__00000425_.X = -9.470983
	__Vector3__00000425_.Y = 1.996053
	__Vector3__00000425_.Z = 3.209436

	__Vector3__00000426_.Name = `TorusPoint 225`
	__Vector3__00000426_.X = -9.510565
	__Vector3__00000426_.Y = 2.000000
	__Vector3__00000426_.Z = 3.090170

	__Vector3__00000427_.Name = `TorusPoint 226`
	__Vector3__00000427_.X = -9.548645
	__Vector3__00000427_.Y = 1.996053
	__Vector3__00000427_.Z = 2.970416

	__Vector3__00000428_.Name = `TorusPoint 227`
	__Vector3__00000428_.X = -9.585218
	__Vector3__00000428_.Y = 1.984229
	__Vector3__00000428_.Z = 2.850193

	__Vector3__00000429_.Name = `TorusPoint 228`
	__Vector3__00000429_.X = -9.620277
	__Vector3__00000429_.Y = 1.964575
	__Vector3__00000429_.Z = 2.729519

	__Vector3__00000430_.Name = `TorusPoint 229`
	__Vector3__00000430_.X = -9.653816
	__Vector3__00000430_.Y = 1.937166
	__Vector3__00000430_.Z = 2.608415

	__Vector3__00000431_.Name = `TorusPoint 230`
	__Vector3__00000431_.X = -9.685832
	__Vector3__00000431_.Y = 1.902113
	__Vector3__00000431_.Z = 2.486899

	__Vector3__00000432_.Name = `TorusPoint 231`
	__Vector3__00000432_.X = -9.716317
	__Vector3__00000432_.Y = 1.859553
	__Vector3__00000432_.Z = 2.364990

	__Vector3__00000433_.Name = `TorusPoint 232`
	__Vector3__00000433_.X = -9.745269
	__Vector3__00000433_.Y = 1.809654
	__Vector3__00000433_.Z = 2.242708

	__Vector3__00000434_.Name = `TorusPoint 233`
	__Vector3__00000434_.X = -9.772681
	__Vector3__00000434_.Y = 1.752613
	__Vector3__00000434_.Z = 2.120071

	__Vector3__00000435_.Name = `TorusPoint 234`
	__Vector3__00000435_.X = -9.798551
	__Vector3__00000435_.Y = 1.688656
	__Vector3__00000435_.Z = 1.997100

	__Vector3__00000436_.Name = `TorusPoint 235`
	__Vector3__00000436_.X = -9.822873
	__Vector3__00000436_.Y = 1.618034
	__Vector3__00000436_.Z = 1.873813

	__Vector3__00000437_.Name = `TorusPoint 236`
	__Vector3__00000437_.X = -9.845643
	__Vector3__00000437_.Y = 1.541026
	__Vector3__00000437_.Z = 1.750231

	__Vector3__00000438_.Name = `TorusPoint 237`
	__Vector3__00000438_.X = -9.866859
	__Vector3__00000438_.Y = 1.457937
	__Vector3__00000438_.Z = 1.626372

	__Vector3__00000439_.Name = `TorusPoint 238`
	__Vector3__00000439_.X = -9.886517
	__Vector3__00000439_.Y = 1.369094
	__Vector3__00000439_.Z = 1.502256

	__Vector3__00000440_.Name = `TorusPoint 239`
	__Vector3__00000440_.X = -9.904614
	__Vector3__00000440_.Y = 1.274848
	__Vector3__00000440_.Z = 1.377903

	__Vector3__00000441_.Name = `TorusPoint 240`
	__Vector3__00000441_.X = -9.921147
	__Vector3__00000441_.Y = 1.175571
	__Vector3__00000441_.Z = 1.253332

	__Vector3__00000442_.Name = `TorusPoint 241`
	__Vector3__00000442_.X = -9.936113
	__Vector3__00000442_.Y = 1.071654
	__Vector3__00000442_.Z = 1.128564

	__Vector3__00000443_.Name = `TorusPoint 242`
	__Vector3__00000443_.X = -9.949510
	__Vector3__00000443_.Y = 0.963507
	__Vector3__00000443_.Z = 1.003617

	__Vector3__00000444_.Name = `TorusPoint 243`
	__Vector3__00000444_.X = -9.961336
	__Vector3__00000444_.Y = 0.851559
	__Vector3__00000444_.Z = 0.878512

	__Vector3__00000445_.Name = `TorusPoint 244`
	__Vector3__00000445_.X = -9.971589
	__Vector3__00000445_.Y = 0.736249
	__Vector3__00000445_.Z = 0.753268

	__Vector3__00000446_.Name = `TorusPoint 245`
	__Vector3__00000446_.X = -9.980267
	__Vector3__00000446_.Y = 0.618034
	__Vector3__00000446_.Z = 0.627905

	__Vector3__00000447_.Name = `TorusPoint 246`
	__Vector3__00000447_.X = -9.987370
	__Vector3__00000447_.Y = 0.497380
	__Vector3__00000447_.Z = 0.502443

	__Vector3__00000448_.Name = `TorusPoint 247`
	__Vector3__00000448_.X = -9.992895
	__Vector3__00000448_.Y = 0.374763
	__Vector3__00000448_.Z = 0.376902

	__Vector3__00000449_.Name = `TorusPoint 248`
	__Vector3__00000449_.X = -9.996842
	__Vector3__00000449_.Y = 0.250666
	__Vector3__00000449_.Z = 0.251301

	__Vector3__00000450_.Name = `TorusPoint 249`
	__Vector3__00000450_.X = -9.999210
	__Vector3__00000450_.Y = 0.125581
	__Vector3__00000450_.Z = 0.125660

	__Vector3__00000451_.Name = `TorusPoint 250`
	__Vector3__00000451_.X = -10.000000
	__Vector3__00000451_.Y = 0.000000
	__Vector3__00000451_.Z = 0.000000

	__Vector3__00000452_.Name = `TorusPoint 251`
	__Vector3__00000452_.X = -9.999210
	__Vector3__00000452_.Y = -0.125581
	__Vector3__00000452_.Z = -0.125660

	__Vector3__00000453_.Name = `TorusPoint 252`
	__Vector3__00000453_.X = -9.996842
	__Vector3__00000453_.Y = -0.250666
	__Vector3__00000453_.Z = -0.251301

	__Vector3__00000454_.Name = `TorusPoint 253`
	__Vector3__00000454_.X = -9.992895
	__Vector3__00000454_.Y = -0.374763
	__Vector3__00000454_.Z = -0.376902

	__Vector3__00000455_.Name = `TorusPoint 254`
	__Vector3__00000455_.X = -9.987370
	__Vector3__00000455_.Y = -0.497380
	__Vector3__00000455_.Z = -0.502443

	__Vector3__00000456_.Name = `TorusPoint 255`
	__Vector3__00000456_.X = -9.980267
	__Vector3__00000456_.Y = -0.618034
	__Vector3__00000456_.Z = -0.627905

	__Vector3__00000457_.Name = `TorusPoint 256`
	__Vector3__00000457_.X = -9.971589
	__Vector3__00000457_.Y = -0.736249
	__Vector3__00000457_.Z = -0.753268

	__Vector3__00000458_.Name = `TorusPoint 257`
	__Vector3__00000458_.X = -9.961336
	__Vector3__00000458_.Y = -0.851559
	__Vector3__00000458_.Z = -0.878512

	__Vector3__00000459_.Name = `TorusPoint 258`
	__Vector3__00000459_.X = -9.949510
	__Vector3__00000459_.Y = -0.963507
	__Vector3__00000459_.Z = -1.003617

	__Vector3__00000460_.Name = `TorusPoint 259`
	__Vector3__00000460_.X = -9.936113
	__Vector3__00000460_.Y = -1.071654
	__Vector3__00000460_.Z = -1.128564

	__Vector3__00000461_.Name = `TorusPoint 260`
	__Vector3__00000461_.X = -9.921147
	__Vector3__00000461_.Y = -1.175571
	__Vector3__00000461_.Z = -1.253332

	__Vector3__00000462_.Name = `TorusPoint 261`
	__Vector3__00000462_.X = -9.904614
	__Vector3__00000462_.Y = -1.274848
	__Vector3__00000462_.Z = -1.377903

	__Vector3__00000463_.Name = `TorusPoint 262`
	__Vector3__00000463_.X = -9.886517
	__Vector3__00000463_.Y = -1.369094
	__Vector3__00000463_.Z = -1.502256

	__Vector3__00000464_.Name = `TorusPoint 263`
	__Vector3__00000464_.X = -9.866859
	__Vector3__00000464_.Y = -1.457937
	__Vector3__00000464_.Z = -1.626372

	__Vector3__00000465_.Name = `TorusPoint 264`
	__Vector3__00000465_.X = -9.845643
	__Vector3__00000465_.Y = -1.541026
	__Vector3__00000465_.Z = -1.750231

	__Vector3__00000466_.Name = `TorusPoint 265`
	__Vector3__00000466_.X = -9.822873
	__Vector3__00000466_.Y = -1.618034
	__Vector3__00000466_.Z = -1.873813

	__Vector3__00000467_.Name = `TorusPoint 266`
	__Vector3__00000467_.X = -9.798551
	__Vector3__00000467_.Y = -1.688656
	__Vector3__00000467_.Z = -1.997100

	__Vector3__00000468_.Name = `TorusPoint 267`
	__Vector3__00000468_.X = -9.772681
	__Vector3__00000468_.Y = -1.752613
	__Vector3__00000468_.Z = -2.120071

	__Vector3__00000469_.Name = `TorusPoint 268`
	__Vector3__00000469_.X = -9.745269
	__Vector3__00000469_.Y = -1.809654
	__Vector3__00000469_.Z = -2.242708

	__Vector3__00000470_.Name = `TorusPoint 269`
	__Vector3__00000470_.X = -9.716317
	__Vector3__00000470_.Y = -1.859553
	__Vector3__00000470_.Z = -2.364990

	__Vector3__00000471_.Name = `TorusPoint 270`
	__Vector3__00000471_.X = -9.685832
	__Vector3__00000471_.Y = -1.902113
	__Vector3__00000471_.Z = -2.486899

	__Vector3__00000472_.Name = `TorusPoint 271`
	__Vector3__00000472_.X = -9.653816
	__Vector3__00000472_.Y = -1.937166
	__Vector3__00000472_.Z = -2.608415

	__Vector3__00000473_.Name = `TorusPoint 272`
	__Vector3__00000473_.X = -9.620277
	__Vector3__00000473_.Y = -1.964575
	__Vector3__00000473_.Z = -2.729519

	__Vector3__00000474_.Name = `TorusPoint 273`
	__Vector3__00000474_.X = -9.585218
	__Vector3__00000474_.Y = -1.984229
	__Vector3__00000474_.Z = -2.850193

	__Vector3__00000475_.Name = `TorusPoint 274`
	__Vector3__00000475_.X = -9.548645
	__Vector3__00000475_.Y = -1.996053
	__Vector3__00000475_.Z = -2.970416

	__Vector3__00000476_.Name = `TorusPoint 275`
	__Vector3__00000476_.X = -9.510565
	__Vector3__00000476_.Y = -2.000000
	__Vector3__00000476_.Z = -3.090170

	__Vector3__00000477_.Name = `TorusPoint 276`
	__Vector3__00000477_.X = -9.470983
	__Vector3__00000477_.Y = -1.996053
	__Vector3__00000477_.Z = -3.209436

	__Vector3__00000478_.Name = `TorusPoint 277`
	__Vector3__00000478_.X = -9.429905
	__Vector3__00000478_.Y = -1.984229
	__Vector3__00000478_.Z = -3.328195

	__Vector3__00000479_.Name = `TorusPoint 278`
	__Vector3__00000479_.X = -9.387339
	__Vector3__00000479_.Y = -1.964575
	__Vector3__00000479_.Z = -3.446429

	__Vector3__00000480_.Name = `TorusPoint 279`
	__Vector3__00000480_.X = -9.343289
	__Vector3__00000480_.Y = -1.937166
	__Vector3__00000480_.Z = -3.564119

	__Vector3__00000481_.Name = `TorusPoint 280`
	__Vector3__00000481_.X = -9.297765
	__Vector3__00000481_.Y = -1.902113
	__Vector3__00000481_.Z = -3.681246

	__Vector3__00000482_.Name = `TorusPoint 281`
	__Vector3__00000482_.X = -9.250772
	__Vector3__00000482_.Y = -1.859553
	__Vector3__00000482_.Z = -3.797791

	__Vector3__00000483_.Name = `TorusPoint 282`
	__Vector3__00000483_.X = -9.202318
	__Vector3__00000483_.Y = -1.809654
	__Vector3__00000483_.Z = -3.913737

	__Vector3__00000484_.Name = `TorusPoint 283`
	__Vector3__00000484_.X = -9.152412
	__Vector3__00000484_.Y = -1.752613
	__Vector3__00000484_.Z = -4.029064

	__Vector3__00000485_.Name = `TorusPoint 284`
	__Vector3__00000485_.X = -9.101060
	__Vector3__00000485_.Y = -1.688656
	__Vector3__00000485_.Z = -4.143756

	__Vector3__00000486_.Name = `TorusPoint 285`
	__Vector3__00000486_.X = -9.048271
	__Vector3__00000486_.Y = -1.618034
	__Vector3__00000486_.Z = -4.257793

	__Vector3__00000487_.Name = `TorusPoint 286`
	__Vector3__00000487_.X = -8.994053
	__Vector3__00000487_.Y = -1.541026
	__Vector3__00000487_.Z = -4.371158

	__Vector3__00000488_.Name = `TorusPoint 287`
	__Vector3__00000488_.X = -8.938414
	__Vector3__00000488_.Y = -1.457937
	__Vector3__00000488_.Z = -4.483832

	__Vector3__00000489_.Name = `TorusPoint 288`
	__Vector3__00000489_.X = -8.881364
	__Vector3__00000489_.Y = -1.369094
	__Vector3__00000489_.Z = -4.595799

	__Vector3__00000490_.Name = `TorusPoint 289`
	__Vector3__00000490_.X = -8.822912
	__Vector3__00000490_.Y = -1.274848
	__Vector3__00000490_.Z = -4.707039

	__Vector3__00000491_.Name = `TorusPoint 290`
	__Vector3__00000491_.X = -8.763067
	__Vector3__00000491_.Y = -1.175571
	__Vector3__00000491_.Z = -4.817537

	__Vector3__00000492_.Name = `TorusPoint 291`
	__Vector3__00000492_.X = -8.701838
	__Vector3__00000492_.Y = -1.071654
	__Vector3__00000492_.Z = -4.927273

	__Vector3__00000493_.Name = `TorusPoint 292`
	__Vector3__00000493_.X = -8.639234
	__Vector3__00000493_.Y = -0.963507
	__Vector3__00000493_.Z = -5.036232

	__Vector3__00000494_.Name = `TorusPoint 293`
	__Vector3__00000494_.X = -8.575267
	__Vector3__00000494_.Y = -0.851559
	__Vector3__00000494_.Z = -5.144395

	__Vector3__00000495_.Name = `TorusPoint 294`
	__Vector3__00000495_.X = -8.509945
	__Vector3__00000495_.Y = -0.736249
	__Vector3__00000495_.Z = -5.251746

	__Vector3__00000496_.Name = `TorusPoint 295`
	__Vector3__00000496_.X = -8.443279
	__Vector3__00000496_.Y = -0.618034
	__Vector3__00000496_.Z = -5.358268

	__Vector3__00000497_.Name = `TorusPoint 296`
	__Vector3__00000497_.X = -8.375280
	__Vector3__00000497_.Y = -0.497380
	__Vector3__00000497_.Z = -5.463943

	__Vector3__00000498_.Name = `TorusPoint 297`
	__Vector3__00000498_.X = -8.305959
	__Vector3__00000498_.Y = -0.374763
	__Vector3__00000498_.Z = -5.568756

	__Vector3__00000499_.Name = `TorusPoint 298`
	__Vector3__00000499_.X = -8.235326
	__Vector3__00000499_.Y = -0.250666
	__Vector3__00000499_.Z = -5.672689

	__Vector3__00000500_.Name = `TorusPoint 299`
	__Vector3__00000500_.X = -8.163393
	__Vector3__00000500_.Y = -0.125581
	__Vector3__00000500_.Z = -5.775727

	__Vector3__00000501_.Name = `TorusPoint 300`
	__Vector3__00000501_.X = -8.090170
	__Vector3__00000501_.Y = -0.000000
	__Vector3__00000501_.Z = -5.877853

	__Vector3__00000502_.Name = `TorusPoint 301`
	__Vector3__00000502_.X = -8.015670
	__Vector3__00000502_.Y = 0.125581
	__Vector3__00000502_.Z = -5.979050

	__Vector3__00000503_.Name = `TorusPoint 302`
	__Vector3__00000503_.X = -7.939904
	__Vector3__00000503_.Y = 0.250666
	__Vector3__00000503_.Z = -6.079303

	__Vector3__00000504_.Name = `TorusPoint 303`
	__Vector3__00000504_.X = -7.862884
	__Vector3__00000504_.Y = 0.374763
	__Vector3__00000504_.Z = -6.178596

	__Vector3__00000505_.Name = `TorusPoint 304`
	__Vector3__00000505_.X = -7.784623
	__Vector3__00000505_.Y = 0.497380
	__Vector3__00000505_.Z = -6.276914

	__Vector3__00000506_.Name = `TorusPoint 305`
	__Vector3__00000506_.X = -7.705132
	__Vector3__00000506_.Y = 0.618034
	__Vector3__00000506_.Z = -6.374240

	__Vector3__00000507_.Name = `TorusPoint 306`
	__Vector3__00000507_.X = -7.624425
	__Vector3__00000507_.Y = 0.736249
	__Vector3__00000507_.Z = -6.470560

	__Vector3__00000508_.Name = `TorusPoint 307`
	__Vector3__00000508_.X = -7.542514
	__Vector3__00000508_.Y = 0.851559
	__Vector3__00000508_.Z = -6.565858

	__Vector3__00000509_.Name = `TorusPoint 308`
	__Vector3__00000509_.X = -7.459411
	__Vector3__00000509_.Y = 0.963507
	__Vector3__00000509_.Z = -6.660119

	__Vector3__00000510_.Name = `TorusPoint 309`
	__Vector3__00000510_.X = -7.375131
	__Vector3__00000510_.Y = 1.071654
	__Vector3__00000510_.Z = -6.753328

	__Vector3__00000511_.Name = `TorusPoint 310`
	__Vector3__00000511_.X = -7.289686
	__Vector3__00000511_.Y = 1.175571
	__Vector3__00000511_.Z = -6.845471

	__Vector3__00000512_.Name = `TorusPoint 311`
	__Vector3__00000512_.X = -7.203090
	__Vector3__00000512_.Y = 1.274848
	__Vector3__00000512_.Z = -6.936533

	__Vector3__00000513_.Name = `TorusPoint 312`
	__Vector3__00000513_.X = -7.115357
	__Vector3__00000513_.Y = 1.369094
	__Vector3__00000513_.Z = -7.026500

	__Vector3__00000514_.Name = `TorusPoint 313`
	__Vector3__00000514_.X = -7.026500
	__Vector3__00000514_.Y = 1.457937
	__Vector3__00000514_.Z = -7.115357

	__Vector3__00000515_.Name = `TorusPoint 314`
	__Vector3__00000515_.X = -6.936533
	__Vector3__00000515_.Y = 1.541026
	__Vector3__00000515_.Z = -7.203090

	__Vector3__00000516_.Name = `TorusPoint 315`
	__Vector3__00000516_.X = -6.845471
	__Vector3__00000516_.Y = 1.618034
	__Vector3__00000516_.Z = -7.289686

	__Vector3__00000517_.Name = `TorusPoint 316`
	__Vector3__00000517_.X = -6.753328
	__Vector3__00000517_.Y = 1.688656
	__Vector3__00000517_.Z = -7.375131

	__Vector3__00000518_.Name = `TorusPoint 317`
	__Vector3__00000518_.X = -6.660119
	__Vector3__00000518_.Y = 1.752613
	__Vector3__00000518_.Z = -7.459411

	__Vector3__00000519_.Name = `TorusPoint 318`
	__Vector3__00000519_.X = -6.565858
	__Vector3__00000519_.Y = 1.809654
	__Vector3__00000519_.Z = -7.542514

	__Vector3__00000520_.Name = `TorusPoint 319`
	__Vector3__00000520_.X = -6.470560
	__Vector3__00000520_.Y = 1.859553
	__Vector3__00000520_.Z = -7.624425

	__Vector3__00000521_.Name = `TorusPoint 320`
	__Vector3__00000521_.X = -6.374240
	__Vector3__00000521_.Y = 1.902113
	__Vector3__00000521_.Z = -7.705132

	__Vector3__00000522_.Name = `TorusPoint 321`
	__Vector3__00000522_.X = -6.276914
	__Vector3__00000522_.Y = 1.937166
	__Vector3__00000522_.Z = -7.784623

	__Vector3__00000523_.Name = `TorusPoint 322`
	__Vector3__00000523_.X = -6.178596
	__Vector3__00000523_.Y = 1.964575
	__Vector3__00000523_.Z = -7.862884

	__Vector3__00000524_.Name = `TorusPoint 323`
	__Vector3__00000524_.X = -6.079303
	__Vector3__00000524_.Y = 1.984229
	__Vector3__00000524_.Z = -7.939904

	__Vector3__00000525_.Name = `TorusPoint 324`
	__Vector3__00000525_.X = -5.979050
	__Vector3__00000525_.Y = 1.996053
	__Vector3__00000525_.Z = -8.015670

	__Vector3__00000526_.Name = `TorusPoint 325`
	__Vector3__00000526_.X = -5.877853
	__Vector3__00000526_.Y = 2.000000
	__Vector3__00000526_.Z = -8.090170

	__Vector3__00000527_.Name = `TorusPoint 326`
	__Vector3__00000527_.X = -5.775727
	__Vector3__00000527_.Y = 1.996053
	__Vector3__00000527_.Z = -8.163393

	__Vector3__00000528_.Name = `TorusPoint 327`
	__Vector3__00000528_.X = -5.672689
	__Vector3__00000528_.Y = 1.984229
	__Vector3__00000528_.Z = -8.235326

	__Vector3__00000529_.Name = `TorusPoint 328`
	__Vector3__00000529_.X = -5.568756
	__Vector3__00000529_.Y = 1.964575
	__Vector3__00000529_.Z = -8.305959

	__Vector3__00000530_.Name = `TorusPoint 329`
	__Vector3__00000530_.X = -5.463943
	__Vector3__00000530_.Y = 1.937166
	__Vector3__00000530_.Z = -8.375280

	__Vector3__00000531_.Name = `TorusPoint 330`
	__Vector3__00000531_.X = -5.358268
	__Vector3__00000531_.Y = 1.902113
	__Vector3__00000531_.Z = -8.443279

	__Vector3__00000532_.Name = `TorusPoint 331`
	__Vector3__00000532_.X = -5.251746
	__Vector3__00000532_.Y = 1.859553
	__Vector3__00000532_.Z = -8.509945

	__Vector3__00000533_.Name = `TorusPoint 332`
	__Vector3__00000533_.X = -5.144395
	__Vector3__00000533_.Y = 1.809654
	__Vector3__00000533_.Z = -8.575267

	__Vector3__00000534_.Name = `TorusPoint 333`
	__Vector3__00000534_.X = -5.036232
	__Vector3__00000534_.Y = 1.752613
	__Vector3__00000534_.Z = -8.639234

	__Vector3__00000535_.Name = `TorusPoint 334`
	__Vector3__00000535_.X = -4.927273
	__Vector3__00000535_.Y = 1.688656
	__Vector3__00000535_.Z = -8.701838

	__Vector3__00000536_.Name = `TorusPoint 335`
	__Vector3__00000536_.X = -4.817537
	__Vector3__00000536_.Y = 1.618034
	__Vector3__00000536_.Z = -8.763067

	__Vector3__00000537_.Name = `TorusPoint 336`
	__Vector3__00000537_.X = -4.707039
	__Vector3__00000537_.Y = 1.541026
	__Vector3__00000537_.Z = -8.822912

	__Vector3__00000538_.Name = `TorusPoint 337`
	__Vector3__00000538_.X = -4.595799
	__Vector3__00000538_.Y = 1.457937
	__Vector3__00000538_.Z = -8.881364

	__Vector3__00000539_.Name = `TorusPoint 338`
	__Vector3__00000539_.X = -4.483832
	__Vector3__00000539_.Y = 1.369094
	__Vector3__00000539_.Z = -8.938414

	__Vector3__00000540_.Name = `TorusPoint 339`
	__Vector3__00000540_.X = -4.371158
	__Vector3__00000540_.Y = 1.274848
	__Vector3__00000540_.Z = -8.994053

	__Vector3__00000541_.Name = `TorusPoint 340`
	__Vector3__00000541_.X = -4.257793
	__Vector3__00000541_.Y = 1.175571
	__Vector3__00000541_.Z = -9.048271

	__Vector3__00000542_.Name = `TorusPoint 341`
	__Vector3__00000542_.X = -4.143756
	__Vector3__00000542_.Y = 1.071654
	__Vector3__00000542_.Z = -9.101060

	__Vector3__00000543_.Name = `TorusPoint 342`
	__Vector3__00000543_.X = -4.029064
	__Vector3__00000543_.Y = 0.963507
	__Vector3__00000543_.Z = -9.152412

	__Vector3__00000544_.Name = `TorusPoint 343`
	__Vector3__00000544_.X = -3.913737
	__Vector3__00000544_.Y = 0.851559
	__Vector3__00000544_.Z = -9.202318

	__Vector3__00000545_.Name = `TorusPoint 344`
	__Vector3__00000545_.X = -3.797791
	__Vector3__00000545_.Y = 0.736249
	__Vector3__00000545_.Z = -9.250772

	__Vector3__00000546_.Name = `TorusPoint 345`
	__Vector3__00000546_.X = -3.681246
	__Vector3__00000546_.Y = 0.618034
	__Vector3__00000546_.Z = -9.297765

	__Vector3__00000547_.Name = `TorusPoint 346`
	__Vector3__00000547_.X = -3.564119
	__Vector3__00000547_.Y = 0.497380
	__Vector3__00000547_.Z = -9.343289

	__Vector3__00000548_.Name = `TorusPoint 347`
	__Vector3__00000548_.X = -3.446429
	__Vector3__00000548_.Y = 0.374763
	__Vector3__00000548_.Z = -9.387339

	__Vector3__00000549_.Name = `TorusPoint 348`
	__Vector3__00000549_.X = -3.328195
	__Vector3__00000549_.Y = 0.250666
	__Vector3__00000549_.Z = -9.429905

	__Vector3__00000550_.Name = `TorusPoint 349`
	__Vector3__00000550_.X = -3.209436
	__Vector3__00000550_.Y = 0.125581
	__Vector3__00000550_.Z = -9.470983

	__Vector3__00000551_.Name = `TorusPoint 350`
	__Vector3__00000551_.X = -3.090170
	__Vector3__00000551_.Y = 0.000000
	__Vector3__00000551_.Z = -9.510565

	__Vector3__00000552_.Name = `TorusPoint 351`
	__Vector3__00000552_.X = -2.970416
	__Vector3__00000552_.Y = -0.125581
	__Vector3__00000552_.Z = -9.548645

	__Vector3__00000553_.Name = `TorusPoint 352`
	__Vector3__00000553_.X = -2.850193
	__Vector3__00000553_.Y = -0.250666
	__Vector3__00000553_.Z = -9.585218

	__Vector3__00000554_.Name = `TorusPoint 353`
	__Vector3__00000554_.X = -2.729519
	__Vector3__00000554_.Y = -0.374763
	__Vector3__00000554_.Z = -9.620277

	__Vector3__00000555_.Name = `TorusPoint 354`
	__Vector3__00000555_.X = -2.608415
	__Vector3__00000555_.Y = -0.497380
	__Vector3__00000555_.Z = -9.653816

	__Vector3__00000556_.Name = `TorusPoint 355`
	__Vector3__00000556_.X = -2.486899
	__Vector3__00000556_.Y = -0.618034
	__Vector3__00000556_.Z = -9.685832

	__Vector3__00000557_.Name = `TorusPoint 356`
	__Vector3__00000557_.X = -2.364990
	__Vector3__00000557_.Y = -0.736249
	__Vector3__00000557_.Z = -9.716317

	__Vector3__00000558_.Name = `TorusPoint 357`
	__Vector3__00000558_.X = -2.242708
	__Vector3__00000558_.Y = -0.851559
	__Vector3__00000558_.Z = -9.745269

	__Vector3__00000559_.Name = `TorusPoint 358`
	__Vector3__00000559_.X = -2.120071
	__Vector3__00000559_.Y = -0.963507
	__Vector3__00000559_.Z = -9.772681

	__Vector3__00000560_.Name = `TorusPoint 359`
	__Vector3__00000560_.X = -1.997100
	__Vector3__00000560_.Y = -1.071654
	__Vector3__00000560_.Z = -9.798551

	__Vector3__00000561_.Name = `TorusPoint 360`
	__Vector3__00000561_.X = -1.873813
	__Vector3__00000561_.Y = -1.175571
	__Vector3__00000561_.Z = -9.822873

	__Vector3__00000562_.Name = `TorusPoint 361`
	__Vector3__00000562_.X = -1.750231
	__Vector3__00000562_.Y = -1.274848
	__Vector3__00000562_.Z = -9.845643

	__Vector3__00000563_.Name = `TorusPoint 362`
	__Vector3__00000563_.X = -1.626372
	__Vector3__00000563_.Y = -1.369094
	__Vector3__00000563_.Z = -9.866859

	__Vector3__00000564_.Name = `TorusPoint 363`
	__Vector3__00000564_.X = -1.502256
	__Vector3__00000564_.Y = -1.457937
	__Vector3__00000564_.Z = -9.886517

	__Vector3__00000565_.Name = `TorusPoint 364`
	__Vector3__00000565_.X = -1.377903
	__Vector3__00000565_.Y = -1.541026
	__Vector3__00000565_.Z = -9.904614

	__Vector3__00000566_.Name = `TorusPoint 365`
	__Vector3__00000566_.X = -1.253332
	__Vector3__00000566_.Y = -1.618034
	__Vector3__00000566_.Z = -9.921147

	__Vector3__00000567_.Name = `TorusPoint 366`
	__Vector3__00000567_.X = -1.128564
	__Vector3__00000567_.Y = -1.688656
	__Vector3__00000567_.Z = -9.936113

	__Vector3__00000568_.Name = `TorusPoint 367`
	__Vector3__00000568_.X = -1.003617
	__Vector3__00000568_.Y = -1.752613
	__Vector3__00000568_.Z = -9.949510

	__Vector3__00000569_.Name = `TorusPoint 368`
	__Vector3__00000569_.X = -0.878512
	__Vector3__00000569_.Y = -1.809654
	__Vector3__00000569_.Z = -9.961336

	__Vector3__00000570_.Name = `TorusPoint 369`
	__Vector3__00000570_.X = -0.753268
	__Vector3__00000570_.Y = -1.859553
	__Vector3__00000570_.Z = -9.971589

	__Vector3__00000571_.Name = `TorusPoint 370`
	__Vector3__00000571_.X = -0.627905
	__Vector3__00000571_.Y = -1.902113
	__Vector3__00000571_.Z = -9.980267

	__Vector3__00000572_.Name = `TorusPoint 371`
	__Vector3__00000572_.X = -0.502443
	__Vector3__00000572_.Y = -1.937166
	__Vector3__00000572_.Z = -9.987370

	__Vector3__00000573_.Name = `TorusPoint 372`
	__Vector3__00000573_.X = -0.376902
	__Vector3__00000573_.Y = -1.964575
	__Vector3__00000573_.Z = -9.992895

	__Vector3__00000574_.Name = `TorusPoint 373`
	__Vector3__00000574_.X = -0.251301
	__Vector3__00000574_.Y = -1.984229
	__Vector3__00000574_.Z = -9.996842

	__Vector3__00000575_.Name = `TorusPoint 374`
	__Vector3__00000575_.X = -0.125660
	__Vector3__00000575_.Y = -1.996053
	__Vector3__00000575_.Z = -9.999210

	__Vector3__00000576_.Name = `TorusPoint 375`
	__Vector3__00000576_.X = -0.000000
	__Vector3__00000576_.Y = -2.000000
	__Vector3__00000576_.Z = -10.000000

	__Vector3__00000577_.Name = `TorusPoint 376`
	__Vector3__00000577_.X = 0.125660
	__Vector3__00000577_.Y = -1.996053
	__Vector3__00000577_.Z = -9.999210

	__Vector3__00000578_.Name = `TorusPoint 377`
	__Vector3__00000578_.X = 0.251301
	__Vector3__00000578_.Y = -1.984229
	__Vector3__00000578_.Z = -9.996842

	__Vector3__00000579_.Name = `TorusPoint 378`
	__Vector3__00000579_.X = 0.376902
	__Vector3__00000579_.Y = -1.964575
	__Vector3__00000579_.Z = -9.992895

	__Vector3__00000580_.Name = `TorusPoint 379`
	__Vector3__00000580_.X = 0.502443
	__Vector3__00000580_.Y = -1.937166
	__Vector3__00000580_.Z = -9.987370

	__Vector3__00000581_.Name = `TorusPoint 380`
	__Vector3__00000581_.X = 0.627905
	__Vector3__00000581_.Y = -1.902113
	__Vector3__00000581_.Z = -9.980267

	__Vector3__00000582_.Name = `TorusPoint 381`
	__Vector3__00000582_.X = 0.753268
	__Vector3__00000582_.Y = -1.859553
	__Vector3__00000582_.Z = -9.971589

	__Vector3__00000583_.Name = `TorusPoint 382`
	__Vector3__00000583_.X = 0.878512
	__Vector3__00000583_.Y = -1.809654
	__Vector3__00000583_.Z = -9.961336

	__Vector3__00000584_.Name = `TorusPoint 383`
	__Vector3__00000584_.X = 1.003617
	__Vector3__00000584_.Y = -1.752613
	__Vector3__00000584_.Z = -9.949510

	__Vector3__00000585_.Name = `TorusPoint 384`
	__Vector3__00000585_.X = 1.128564
	__Vector3__00000585_.Y = -1.688656
	__Vector3__00000585_.Z = -9.936113

	__Vector3__00000586_.Name = `TorusPoint 385`
	__Vector3__00000586_.X = 1.253332
	__Vector3__00000586_.Y = -1.618034
	__Vector3__00000586_.Z = -9.921147

	__Vector3__00000587_.Name = `TorusPoint 386`
	__Vector3__00000587_.X = 1.377903
	__Vector3__00000587_.Y = -1.541026
	__Vector3__00000587_.Z = -9.904614

	__Vector3__00000588_.Name = `TorusPoint 387`
	__Vector3__00000588_.X = 1.502256
	__Vector3__00000588_.Y = -1.457937
	__Vector3__00000588_.Z = -9.886517

	__Vector3__00000589_.Name = `TorusPoint 388`
	__Vector3__00000589_.X = 1.626372
	__Vector3__00000589_.Y = -1.369094
	__Vector3__00000589_.Z = -9.866859

	__Vector3__00000590_.Name = `TorusPoint 389`
	__Vector3__00000590_.X = 1.750231
	__Vector3__00000590_.Y = -1.274848
	__Vector3__00000590_.Z = -9.845643

	__Vector3__00000591_.Name = `TorusPoint 390`
	__Vector3__00000591_.X = 1.873813
	__Vector3__00000591_.Y = -1.175571
	__Vector3__00000591_.Z = -9.822873

	__Vector3__00000592_.Name = `TorusPoint 391`
	__Vector3__00000592_.X = 1.997100
	__Vector3__00000592_.Y = -1.071654
	__Vector3__00000592_.Z = -9.798551

	__Vector3__00000593_.Name = `TorusPoint 392`
	__Vector3__00000593_.X = 2.120071
	__Vector3__00000593_.Y = -0.963507
	__Vector3__00000593_.Z = -9.772681

	__Vector3__00000594_.Name = `TorusPoint 393`
	__Vector3__00000594_.X = 2.242708
	__Vector3__00000594_.Y = -0.851559
	__Vector3__00000594_.Z = -9.745269

	__Vector3__00000595_.Name = `TorusPoint 394`
	__Vector3__00000595_.X = 2.364990
	__Vector3__00000595_.Y = -0.736249
	__Vector3__00000595_.Z = -9.716317

	__Vector3__00000596_.Name = `TorusPoint 395`
	__Vector3__00000596_.X = 2.486899
	__Vector3__00000596_.Y = -0.618034
	__Vector3__00000596_.Z = -9.685832

	__Vector3__00000597_.Name = `TorusPoint 396`
	__Vector3__00000597_.X = 2.608415
	__Vector3__00000597_.Y = -0.497380
	__Vector3__00000597_.Z = -9.653816

	__Vector3__00000598_.Name = `TorusPoint 397`
	__Vector3__00000598_.X = 2.729519
	__Vector3__00000598_.Y = -0.374763
	__Vector3__00000598_.Z = -9.620277

	__Vector3__00000599_.Name = `TorusPoint 398`
	__Vector3__00000599_.X = 2.850193
	__Vector3__00000599_.Y = -0.250666
	__Vector3__00000599_.Z = -9.585218

	__Vector3__00000600_.Name = `TorusPoint 399`
	__Vector3__00000600_.X = 2.970416
	__Vector3__00000600_.Y = -0.125581
	__Vector3__00000600_.Z = -9.548645

	__Vector3__00000601_.Name = `TorusPoint 400`
	__Vector3__00000601_.X = 3.090170
	__Vector3__00000601_.Y = -0.000000
	__Vector3__00000601_.Z = -9.510565

	__Vector3__00000602_.Name = `TorusPoint 401`
	__Vector3__00000602_.X = 3.209436
	__Vector3__00000602_.Y = 0.125581
	__Vector3__00000602_.Z = -9.470983

	__Vector3__00000603_.Name = `TorusPoint 402`
	__Vector3__00000603_.X = 3.328195
	__Vector3__00000603_.Y = 0.250666
	__Vector3__00000603_.Z = -9.429905

	__Vector3__00000604_.Name = `TorusPoint 403`
	__Vector3__00000604_.X = 3.446429
	__Vector3__00000604_.Y = 0.374763
	__Vector3__00000604_.Z = -9.387339

	__Vector3__00000605_.Name = `TorusPoint 404`
	__Vector3__00000605_.X = 3.564119
	__Vector3__00000605_.Y = 0.497380
	__Vector3__00000605_.Z = -9.343289

	__Vector3__00000606_.Name = `TorusPoint 405`
	__Vector3__00000606_.X = 3.681246
	__Vector3__00000606_.Y = 0.618034
	__Vector3__00000606_.Z = -9.297765

	__Vector3__00000607_.Name = `TorusPoint 406`
	__Vector3__00000607_.X = 3.797791
	__Vector3__00000607_.Y = 0.736249
	__Vector3__00000607_.Z = -9.250772

	__Vector3__00000608_.Name = `TorusPoint 407`
	__Vector3__00000608_.X = 3.913737
	__Vector3__00000608_.Y = 0.851559
	__Vector3__00000608_.Z = -9.202318

	__Vector3__00000609_.Name = `TorusPoint 408`
	__Vector3__00000609_.X = 4.029064
	__Vector3__00000609_.Y = 0.963507
	__Vector3__00000609_.Z = -9.152412

	__Vector3__00000610_.Name = `TorusPoint 409`
	__Vector3__00000610_.X = 4.143756
	__Vector3__00000610_.Y = 1.071654
	__Vector3__00000610_.Z = -9.101060

	__Vector3__00000611_.Name = `TorusPoint 410`
	__Vector3__00000611_.X = 4.257793
	__Vector3__00000611_.Y = 1.175571
	__Vector3__00000611_.Z = -9.048271

	__Vector3__00000612_.Name = `TorusPoint 411`
	__Vector3__00000612_.X = 4.371158
	__Vector3__00000612_.Y = 1.274848
	__Vector3__00000612_.Z = -8.994053

	__Vector3__00000613_.Name = `TorusPoint 412`
	__Vector3__00000613_.X = 4.483832
	__Vector3__00000613_.Y = 1.369094
	__Vector3__00000613_.Z = -8.938414

	__Vector3__00000614_.Name = `TorusPoint 413`
	__Vector3__00000614_.X = 4.595799
	__Vector3__00000614_.Y = 1.457937
	__Vector3__00000614_.Z = -8.881364

	__Vector3__00000615_.Name = `TorusPoint 414`
	__Vector3__00000615_.X = 4.707039
	__Vector3__00000615_.Y = 1.541026
	__Vector3__00000615_.Z = -8.822912

	__Vector3__00000616_.Name = `TorusPoint 415`
	__Vector3__00000616_.X = 4.817537
	__Vector3__00000616_.Y = 1.618034
	__Vector3__00000616_.Z = -8.763067

	__Vector3__00000617_.Name = `TorusPoint 416`
	__Vector3__00000617_.X = 4.927273
	__Vector3__00000617_.Y = 1.688656
	__Vector3__00000617_.Z = -8.701838

	__Vector3__00000618_.Name = `TorusPoint 417`
	__Vector3__00000618_.X = 5.036232
	__Vector3__00000618_.Y = 1.752613
	__Vector3__00000618_.Z = -8.639234

	__Vector3__00000619_.Name = `TorusPoint 418`
	__Vector3__00000619_.X = 5.144395
	__Vector3__00000619_.Y = 1.809654
	__Vector3__00000619_.Z = -8.575267

	__Vector3__00000620_.Name = `TorusPoint 419`
	__Vector3__00000620_.X = 5.251746
	__Vector3__00000620_.Y = 1.859553
	__Vector3__00000620_.Z = -8.509945

	__Vector3__00000621_.Name = `TorusPoint 420`
	__Vector3__00000621_.X = 5.358268
	__Vector3__00000621_.Y = 1.902113
	__Vector3__00000621_.Z = -8.443279

	__Vector3__00000622_.Name = `TorusPoint 421`
	__Vector3__00000622_.X = 5.463943
	__Vector3__00000622_.Y = 1.937166
	__Vector3__00000622_.Z = -8.375280

	__Vector3__00000623_.Name = `TorusPoint 422`
	__Vector3__00000623_.X = 5.568756
	__Vector3__00000623_.Y = 1.964575
	__Vector3__00000623_.Z = -8.305959

	__Vector3__00000624_.Name = `TorusPoint 423`
	__Vector3__00000624_.X = 5.672689
	__Vector3__00000624_.Y = 1.984229
	__Vector3__00000624_.Z = -8.235326

	__Vector3__00000625_.Name = `TorusPoint 424`
	__Vector3__00000625_.X = 5.775727
	__Vector3__00000625_.Y = 1.996053
	__Vector3__00000625_.Z = -8.163393

	__Vector3__00000626_.Name = `TorusPoint 425`
	__Vector3__00000626_.X = 5.877853
	__Vector3__00000626_.Y = 2.000000
	__Vector3__00000626_.Z = -8.090170

	__Vector3__00000627_.Name = `TorusPoint 426`
	__Vector3__00000627_.X = 5.979050
	__Vector3__00000627_.Y = 1.996053
	__Vector3__00000627_.Z = -8.015670

	__Vector3__00000628_.Name = `TorusPoint 427`
	__Vector3__00000628_.X = 6.079303
	__Vector3__00000628_.Y = 1.984229
	__Vector3__00000628_.Z = -7.939904

	__Vector3__00000629_.Name = `TorusPoint 428`
	__Vector3__00000629_.X = 6.178596
	__Vector3__00000629_.Y = 1.964575
	__Vector3__00000629_.Z = -7.862884

	__Vector3__00000630_.Name = `TorusPoint 429`
	__Vector3__00000630_.X = 6.276914
	__Vector3__00000630_.Y = 1.937166
	__Vector3__00000630_.Z = -7.784623

	__Vector3__00000631_.Name = `TorusPoint 430`
	__Vector3__00000631_.X = 6.374240
	__Vector3__00000631_.Y = 1.902113
	__Vector3__00000631_.Z = -7.705132

	__Vector3__00000632_.Name = `TorusPoint 431`
	__Vector3__00000632_.X = 6.470560
	__Vector3__00000632_.Y = 1.859553
	__Vector3__00000632_.Z = -7.624425

	__Vector3__00000633_.Name = `TorusPoint 432`
	__Vector3__00000633_.X = 6.565858
	__Vector3__00000633_.Y = 1.809654
	__Vector3__00000633_.Z = -7.542514

	__Vector3__00000634_.Name = `TorusPoint 433`
	__Vector3__00000634_.X = 6.660119
	__Vector3__00000634_.Y = 1.752613
	__Vector3__00000634_.Z = -7.459411

	__Vector3__00000635_.Name = `TorusPoint 434`
	__Vector3__00000635_.X = 6.753328
	__Vector3__00000635_.Y = 1.688656
	__Vector3__00000635_.Z = -7.375131

	__Vector3__00000636_.Name = `TorusPoint 435`
	__Vector3__00000636_.X = 6.845471
	__Vector3__00000636_.Y = 1.618034
	__Vector3__00000636_.Z = -7.289686

	__Vector3__00000637_.Name = `TorusPoint 436`
	__Vector3__00000637_.X = 6.936533
	__Vector3__00000637_.Y = 1.541026
	__Vector3__00000637_.Z = -7.203090

	__Vector3__00000638_.Name = `TorusPoint 437`
	__Vector3__00000638_.X = 7.026500
	__Vector3__00000638_.Y = 1.457937
	__Vector3__00000638_.Z = -7.115357

	__Vector3__00000639_.Name = `TorusPoint 438`
	__Vector3__00000639_.X = 7.115357
	__Vector3__00000639_.Y = 1.369094
	__Vector3__00000639_.Z = -7.026500

	__Vector3__00000640_.Name = `TorusPoint 439`
	__Vector3__00000640_.X = 7.203090
	__Vector3__00000640_.Y = 1.274848
	__Vector3__00000640_.Z = -6.936533

	__Vector3__00000641_.Name = `TorusPoint 440`
	__Vector3__00000641_.X = 7.289686
	__Vector3__00000641_.Y = 1.175571
	__Vector3__00000641_.Z = -6.845471

	__Vector3__00000642_.Name = `TorusPoint 441`
	__Vector3__00000642_.X = 7.375131
	__Vector3__00000642_.Y = 1.071654
	__Vector3__00000642_.Z = -6.753328

	__Vector3__00000643_.Name = `TorusPoint 442`
	__Vector3__00000643_.X = 7.459411
	__Vector3__00000643_.Y = 0.963507
	__Vector3__00000643_.Z = -6.660119

	__Vector3__00000644_.Name = `TorusPoint 443`
	__Vector3__00000644_.X = 7.542514
	__Vector3__00000644_.Y = 0.851559
	__Vector3__00000644_.Z = -6.565858

	__Vector3__00000645_.Name = `TorusPoint 444`
	__Vector3__00000645_.X = 7.624425
	__Vector3__00000645_.Y = 0.736249
	__Vector3__00000645_.Z = -6.470560

	__Vector3__00000646_.Name = `TorusPoint 445`
	__Vector3__00000646_.X = 7.705132
	__Vector3__00000646_.Y = 0.618034
	__Vector3__00000646_.Z = -6.374240

	__Vector3__00000647_.Name = `TorusPoint 446`
	__Vector3__00000647_.X = 7.784623
	__Vector3__00000647_.Y = 0.497380
	__Vector3__00000647_.Z = -6.276914

	__Vector3__00000648_.Name = `TorusPoint 447`
	__Vector3__00000648_.X = 7.862884
	__Vector3__00000648_.Y = 0.374763
	__Vector3__00000648_.Z = -6.178596

	__Vector3__00000649_.Name = `TorusPoint 448`
	__Vector3__00000649_.X = 7.939904
	__Vector3__00000649_.Y = 0.250666
	__Vector3__00000649_.Z = -6.079303

	__Vector3__00000650_.Name = `TorusPoint 449`
	__Vector3__00000650_.X = 8.015670
	__Vector3__00000650_.Y = 0.125581
	__Vector3__00000650_.Z = -5.979050

	__Vector3__00000651_.Name = `TorusPoint 450`
	__Vector3__00000651_.X = 8.090170
	__Vector3__00000651_.Y = 0.000000
	__Vector3__00000651_.Z = -5.877853

	__Vector3__00000652_.Name = `TorusPoint 451`
	__Vector3__00000652_.X = 8.163393
	__Vector3__00000652_.Y = -0.125581
	__Vector3__00000652_.Z = -5.775727

	__Vector3__00000653_.Name = `TorusPoint 452`
	__Vector3__00000653_.X = 8.235326
	__Vector3__00000653_.Y = -0.250666
	__Vector3__00000653_.Z = -5.672689

	__Vector3__00000654_.Name = `TorusPoint 453`
	__Vector3__00000654_.X = 8.305959
	__Vector3__00000654_.Y = -0.374763
	__Vector3__00000654_.Z = -5.568756

	__Vector3__00000655_.Name = `TorusPoint 454`
	__Vector3__00000655_.X = 8.375280
	__Vector3__00000655_.Y = -0.497380
	__Vector3__00000655_.Z = -5.463943

	__Vector3__00000656_.Name = `TorusPoint 455`
	__Vector3__00000656_.X = 8.443279
	__Vector3__00000656_.Y = -0.618034
	__Vector3__00000656_.Z = -5.358268

	__Vector3__00000657_.Name = `TorusPoint 456`
	__Vector3__00000657_.X = 8.509945
	__Vector3__00000657_.Y = -0.736249
	__Vector3__00000657_.Z = -5.251746

	__Vector3__00000658_.Name = `TorusPoint 457`
	__Vector3__00000658_.X = 8.575267
	__Vector3__00000658_.Y = -0.851559
	__Vector3__00000658_.Z = -5.144395

	__Vector3__00000659_.Name = `TorusPoint 458`
	__Vector3__00000659_.X = 8.639234
	__Vector3__00000659_.Y = -0.963507
	__Vector3__00000659_.Z = -5.036232

	__Vector3__00000660_.Name = `TorusPoint 459`
	__Vector3__00000660_.X = 8.701838
	__Vector3__00000660_.Y = -1.071654
	__Vector3__00000660_.Z = -4.927273

	__Vector3__00000661_.Name = `TorusPoint 460`
	__Vector3__00000661_.X = 8.763067
	__Vector3__00000661_.Y = -1.175571
	__Vector3__00000661_.Z = -4.817537

	__Vector3__00000662_.Name = `TorusPoint 461`
	__Vector3__00000662_.X = 8.822912
	__Vector3__00000662_.Y = -1.274848
	__Vector3__00000662_.Z = -4.707039

	__Vector3__00000663_.Name = `TorusPoint 462`
	__Vector3__00000663_.X = 8.881364
	__Vector3__00000663_.Y = -1.369094
	__Vector3__00000663_.Z = -4.595799

	__Vector3__00000664_.Name = `TorusPoint 463`
	__Vector3__00000664_.X = 8.938414
	__Vector3__00000664_.Y = -1.457937
	__Vector3__00000664_.Z = -4.483832

	__Vector3__00000665_.Name = `TorusPoint 464`
	__Vector3__00000665_.X = 8.994053
	__Vector3__00000665_.Y = -1.541026
	__Vector3__00000665_.Z = -4.371158

	__Vector3__00000666_.Name = `TorusPoint 465`
	__Vector3__00000666_.X = 9.048271
	__Vector3__00000666_.Y = -1.618034
	__Vector3__00000666_.Z = -4.257793

	__Vector3__00000667_.Name = `TorusPoint 466`
	__Vector3__00000667_.X = 9.101060
	__Vector3__00000667_.Y = -1.688656
	__Vector3__00000667_.Z = -4.143756

	__Vector3__00000668_.Name = `TorusPoint 467`
	__Vector3__00000668_.X = 9.152412
	__Vector3__00000668_.Y = -1.752613
	__Vector3__00000668_.Z = -4.029064

	__Vector3__00000669_.Name = `TorusPoint 468`
	__Vector3__00000669_.X = 9.202318
	__Vector3__00000669_.Y = -1.809654
	__Vector3__00000669_.Z = -3.913737

	__Vector3__00000670_.Name = `TorusPoint 469`
	__Vector3__00000670_.X = 9.250772
	__Vector3__00000670_.Y = -1.859553
	__Vector3__00000670_.Z = -3.797791

	__Vector3__00000671_.Name = `TorusPoint 470`
	__Vector3__00000671_.X = 9.297765
	__Vector3__00000671_.Y = -1.902113
	__Vector3__00000671_.Z = -3.681246

	__Vector3__00000672_.Name = `TorusPoint 471`
	__Vector3__00000672_.X = 9.343289
	__Vector3__00000672_.Y = -1.937166
	__Vector3__00000672_.Z = -3.564119

	__Vector3__00000673_.Name = `TorusPoint 472`
	__Vector3__00000673_.X = 9.387339
	__Vector3__00000673_.Y = -1.964575
	__Vector3__00000673_.Z = -3.446429

	__Vector3__00000674_.Name = `TorusPoint 473`
	__Vector3__00000674_.X = 9.429905
	__Vector3__00000674_.Y = -1.984229
	__Vector3__00000674_.Z = -3.328195

	__Vector3__00000675_.Name = `TorusPoint 474`
	__Vector3__00000675_.X = 9.470983
	__Vector3__00000675_.Y = -1.996053
	__Vector3__00000675_.Z = -3.209436

	__Vector3__00000676_.Name = `TorusPoint 475`
	__Vector3__00000676_.X = 9.510565
	__Vector3__00000676_.Y = -2.000000
	__Vector3__00000676_.Z = -3.090170

	__Vector3__00000677_.Name = `TorusPoint 476`
	__Vector3__00000677_.X = 9.548645
	__Vector3__00000677_.Y = -1.996053
	__Vector3__00000677_.Z = -2.970416

	__Vector3__00000678_.Name = `TorusPoint 477`
	__Vector3__00000678_.X = 9.585218
	__Vector3__00000678_.Y = -1.984229
	__Vector3__00000678_.Z = -2.850193

	__Vector3__00000679_.Name = `TorusPoint 478`
	__Vector3__00000679_.X = 9.620277
	__Vector3__00000679_.Y = -1.964575
	__Vector3__00000679_.Z = -2.729519

	__Vector3__00000680_.Name = `TorusPoint 479`
	__Vector3__00000680_.X = 9.653816
	__Vector3__00000680_.Y = -1.937166
	__Vector3__00000680_.Z = -2.608415

	__Vector3__00000681_.Name = `TorusPoint 480`
	__Vector3__00000681_.X = 9.685832
	__Vector3__00000681_.Y = -1.902113
	__Vector3__00000681_.Z = -2.486899

	__Vector3__00000682_.Name = `TorusPoint 481`
	__Vector3__00000682_.X = 9.716317
	__Vector3__00000682_.Y = -1.859553
	__Vector3__00000682_.Z = -2.364990

	__Vector3__00000683_.Name = `TorusPoint 482`
	__Vector3__00000683_.X = 9.745269
	__Vector3__00000683_.Y = -1.809654
	__Vector3__00000683_.Z = -2.242708

	__Vector3__00000684_.Name = `TorusPoint 483`
	__Vector3__00000684_.X = 9.772681
	__Vector3__00000684_.Y = -1.752613
	__Vector3__00000684_.Z = -2.120071

	__Vector3__00000685_.Name = `TorusPoint 484`
	__Vector3__00000685_.X = 9.798551
	__Vector3__00000685_.Y = -1.688656
	__Vector3__00000685_.Z = -1.997100

	__Vector3__00000686_.Name = `TorusPoint 485`
	__Vector3__00000686_.X = 9.822873
	__Vector3__00000686_.Y = -1.618034
	__Vector3__00000686_.Z = -1.873813

	__Vector3__00000687_.Name = `TorusPoint 486`
	__Vector3__00000687_.X = 9.845643
	__Vector3__00000687_.Y = -1.541026
	__Vector3__00000687_.Z = -1.750231

	__Vector3__00000688_.Name = `TorusPoint 487`
	__Vector3__00000688_.X = 9.866859
	__Vector3__00000688_.Y = -1.457937
	__Vector3__00000688_.Z = -1.626372

	__Vector3__00000689_.Name = `TorusPoint 488`
	__Vector3__00000689_.X = 9.886517
	__Vector3__00000689_.Y = -1.369094
	__Vector3__00000689_.Z = -1.502256

	__Vector3__00000690_.Name = `TorusPoint 489`
	__Vector3__00000690_.X = 9.904614
	__Vector3__00000690_.Y = -1.274848
	__Vector3__00000690_.Z = -1.377903

	__Vector3__00000691_.Name = `TorusPoint 490`
	__Vector3__00000691_.X = 9.921147
	__Vector3__00000691_.Y = -1.175571
	__Vector3__00000691_.Z = -1.253332

	__Vector3__00000692_.Name = `TorusPoint 491`
	__Vector3__00000692_.X = 9.936113
	__Vector3__00000692_.Y = -1.071654
	__Vector3__00000692_.Z = -1.128564

	__Vector3__00000693_.Name = `TorusPoint 492`
	__Vector3__00000693_.X = 9.949510
	__Vector3__00000693_.Y = -0.963507
	__Vector3__00000693_.Z = -1.003617

	__Vector3__00000694_.Name = `TorusPoint 493`
	__Vector3__00000694_.X = 9.961336
	__Vector3__00000694_.Y = -0.851559
	__Vector3__00000694_.Z = -0.878512

	__Vector3__00000695_.Name = `TorusPoint 494`
	__Vector3__00000695_.X = 9.971589
	__Vector3__00000695_.Y = -0.736249
	__Vector3__00000695_.Z = -0.753268

	__Vector3__00000696_.Name = `TorusPoint 495`
	__Vector3__00000696_.X = 9.980267
	__Vector3__00000696_.Y = -0.618034
	__Vector3__00000696_.Z = -0.627905

	__Vector3__00000697_.Name = `TorusPoint 496`
	__Vector3__00000697_.X = 9.987370
	__Vector3__00000697_.Y = -0.497380
	__Vector3__00000697_.Z = -0.502443

	__Vector3__00000698_.Name = `TorusPoint 497`
	__Vector3__00000698_.X = 9.992895
	__Vector3__00000698_.Y = -0.374763
	__Vector3__00000698_.Z = -0.376902

	__Vector3__00000699_.Name = `TorusPoint 498`
	__Vector3__00000699_.X = 9.996842
	__Vector3__00000699_.Y = -0.250666
	__Vector3__00000699_.Z = -0.251301

	__Vector3__00000700_.Name = `TorusPoint 499`
	__Vector3__00000700_.X = 9.999210
	__Vector3__00000700_.Y = -0.125581
	__Vector3__00000700_.Z = -0.125660

	__Vector3__00000701_.Name = `TorusPoint 500`
	__Vector3__00000701_.X = 10.000000
	__Vector3__00000701_.Y = -0.000000
	__Vector3__00000701_.Z = -0.000000

	// insertion point for setup of pointers
	__Canvas__00000000_.DirectionalLights = append(__Canvas__00000000_.DirectionalLights, __DirectionalLight__00000000_)
	__Canvas__00000000_.AmbiantLight = __AmbiantLight__00000000_
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000000_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000001_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000002_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000003_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000004_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000005_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000006_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000007_)
	__Canvas__00000000_.Meshs = append(__Canvas__00000000_.Meshs, __Mesh__00000008_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000000_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000001_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000002_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000003_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000004_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000005_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000006_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000007_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000008_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000009_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000010_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000011_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000012_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000013_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000014_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000015_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000016_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000017_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000018_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000019_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000020_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000021_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000022_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000023_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000024_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000025_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000026_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000027_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000028_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000029_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000030_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000031_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000032_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000033_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000034_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000035_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000036_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000037_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000038_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000039_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000040_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000041_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000042_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000043_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000044_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000045_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000046_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000047_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000048_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000049_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000050_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000051_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000052_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000053_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000054_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000055_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000056_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000057_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000058_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000059_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000060_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000061_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000062_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000063_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000064_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000065_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000066_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000067_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000068_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000069_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000070_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000071_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000072_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000073_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000074_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000075_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000076_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000077_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000078_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000079_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000080_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000081_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000082_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000083_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000084_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000085_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000086_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000087_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000088_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000089_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000090_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000091_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000092_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000093_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000094_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000095_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000096_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000097_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000098_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000099_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000100_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000101_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000102_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000103_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000104_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000105_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000106_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000107_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000108_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000109_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000110_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000111_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000112_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000113_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000114_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000115_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000116_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000117_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000118_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000119_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000120_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000121_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000122_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000123_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000124_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000125_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000126_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000127_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000128_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000129_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000130_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000131_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000132_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000133_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000134_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000135_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000136_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000137_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000138_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000139_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000140_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000141_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000142_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000143_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000144_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000145_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000146_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000147_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000148_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000149_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000150_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000151_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000152_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000153_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000154_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000155_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000156_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000157_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000158_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000159_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000160_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000161_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000162_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000163_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000164_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000165_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000166_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000167_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000168_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000169_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000170_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000171_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000172_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000173_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000174_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000175_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000176_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000177_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000178_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000179_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000180_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000181_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000182_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000183_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000184_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000185_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000186_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000187_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000188_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000189_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000190_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000191_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000192_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000193_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000194_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000195_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000196_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000197_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000198_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000199_)
	__Curve__00000000_.Points = append(__Curve__00000000_.Points, __Vector3__00000200_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000201_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000202_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000203_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000204_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000205_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000206_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000207_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000208_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000209_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000210_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000211_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000212_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000213_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000214_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000215_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000216_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000217_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000218_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000219_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000220_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000221_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000222_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000223_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000224_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000225_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000226_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000227_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000228_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000229_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000230_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000231_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000232_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000233_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000234_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000235_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000236_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000237_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000238_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000239_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000240_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000241_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000242_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000243_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000244_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000245_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000246_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000247_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000248_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000249_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000250_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000251_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000252_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000253_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000254_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000255_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000256_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000257_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000258_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000259_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000260_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000261_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000262_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000263_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000264_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000265_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000266_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000267_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000268_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000269_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000270_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000271_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000272_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000273_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000274_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000275_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000276_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000277_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000278_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000279_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000280_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000281_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000282_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000283_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000284_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000285_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000286_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000287_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000288_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000289_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000290_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000291_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000292_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000293_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000294_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000295_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000296_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000297_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000298_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000299_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000300_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000301_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000302_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000303_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000304_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000305_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000306_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000307_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000308_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000309_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000310_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000311_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000312_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000313_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000314_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000315_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000316_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000317_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000318_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000319_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000320_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000321_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000322_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000323_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000324_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000325_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000326_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000327_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000328_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000329_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000330_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000331_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000332_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000333_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000334_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000335_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000336_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000337_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000338_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000339_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000340_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000341_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000342_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000343_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000344_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000345_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000346_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000347_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000348_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000349_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000350_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000351_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000352_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000353_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000354_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000355_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000356_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000357_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000358_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000359_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000360_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000361_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000362_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000363_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000364_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000365_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000366_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000367_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000368_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000369_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000370_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000371_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000372_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000373_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000374_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000375_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000376_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000377_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000378_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000379_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000380_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000381_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000382_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000383_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000384_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000385_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000386_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000387_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000388_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000389_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000390_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000391_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000392_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000393_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000394_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000395_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000396_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000397_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000398_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000399_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000400_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000401_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000402_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000403_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000404_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000405_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000406_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000407_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000408_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000409_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000410_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000411_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000412_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000413_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000414_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000415_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000416_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000417_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000418_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000419_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000420_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000421_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000422_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000423_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000424_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000425_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000426_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000427_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000428_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000429_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000430_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000431_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000432_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000433_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000434_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000435_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000436_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000437_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000438_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000439_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000440_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000441_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000442_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000443_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000444_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000445_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000446_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000447_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000448_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000449_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000450_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000451_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000452_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000453_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000454_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000455_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000456_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000457_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000458_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000459_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000460_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000461_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000462_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000463_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000464_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000465_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000466_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000467_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000468_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000469_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000470_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000471_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000472_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000473_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000474_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000475_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000476_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000477_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000478_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000479_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000480_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000481_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000482_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000483_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000484_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000485_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000486_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000487_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000488_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000489_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000490_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000491_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000492_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000493_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000494_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000495_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000496_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000497_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000498_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000499_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000500_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000501_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000502_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000503_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000504_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000505_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000506_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000507_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000508_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000509_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000510_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000511_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000512_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000513_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000514_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000515_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000516_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000517_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000518_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000519_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000520_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000521_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000522_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000523_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000524_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000525_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000526_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000527_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000528_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000529_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000530_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000531_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000532_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000533_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000534_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000535_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000536_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000537_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000538_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000539_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000540_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000541_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000542_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000543_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000544_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000545_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000546_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000547_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000548_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000549_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000550_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000551_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000552_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000553_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000554_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000555_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000556_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000557_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000558_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000559_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000560_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000561_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000562_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000563_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000564_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000565_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000566_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000567_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000568_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000569_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000570_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000571_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000572_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000573_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000574_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000575_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000576_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000577_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000578_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000579_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000580_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000581_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000582_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000583_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000584_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000585_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000586_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000587_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000588_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000589_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000590_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000591_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000592_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000593_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000594_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000595_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000596_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000597_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000598_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000599_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000600_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000601_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000602_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000603_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000604_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000605_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000606_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000607_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000608_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000609_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000610_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000611_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000612_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000613_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000614_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000615_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000616_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000617_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000618_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000619_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000620_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000621_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000622_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000623_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000624_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000625_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000626_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000627_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000628_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000629_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000630_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000631_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000632_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000633_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000634_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000635_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000636_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000637_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000638_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000639_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000640_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000641_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000642_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000643_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000644_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000645_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000646_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000647_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000648_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000649_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000650_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000651_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000652_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000653_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000654_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000655_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000656_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000657_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000658_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000659_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000660_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000661_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000662_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000663_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000664_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000665_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000666_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000667_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000668_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000669_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000670_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000671_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000672_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000673_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000674_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000675_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000676_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000677_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000678_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000679_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000680_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000681_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000682_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000683_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000684_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000685_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000686_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000687_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000688_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000689_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000690_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000691_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000692_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000693_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000694_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000695_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000696_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000697_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000698_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000699_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000700_)
	__Curve__00000001_.Points = append(__Curve__00000001_.Points, __Vector3__00000701_)
	__ExtrudeGeometry__00000000_.Shape = __Shape__00000000_
	__ExtrudeGeometry__00000000_.ExtrudePath = __Curve__00000001_
	__Mesh__00000000_.MeshMaterialBasic = nil
	__Mesh__00000000_.MeshPhysicalMaterial = nil
	__Mesh__00000000_.CylinderGeometry = nil
	__Mesh__00000000_.BoxGeometry = __BoxGeometry__00000000_
	__Mesh__00000000_.SphereGeometry = nil
	__Mesh__00000000_.TorusGeometry = nil
	__Mesh__00000000_.PlaneGeometry = nil
	__Mesh__00000000_.TubeGeometry = nil
	__Mesh__00000000_.ExtrudeGeometry = nil
	__Mesh__00000001_.MeshMaterialBasic = __MeshMaterialBasic__00000000_
	__Mesh__00000001_.MeshPhysicalMaterial = nil
	__Mesh__00000001_.CylinderGeometry = __CylinderGeometry__00000000_
	__Mesh__00000001_.BoxGeometry = nil
	__Mesh__00000001_.SphereGeometry = nil
	__Mesh__00000001_.TorusGeometry = nil
	__Mesh__00000001_.PlaneGeometry = nil
	__Mesh__00000001_.TubeGeometry = nil
	__Mesh__00000001_.ExtrudeGeometry = nil
	__Mesh__00000002_.MeshMaterialBasic = nil
	__Mesh__00000002_.MeshPhysicalMaterial = __MeshPhysicalMaterial__00000000_
	__Mesh__00000002_.CylinderGeometry = nil
	__Mesh__00000002_.BoxGeometry = nil
	__Mesh__00000002_.SphereGeometry = __SphereGeometry__00000000_
	__Mesh__00000002_.TorusGeometry = nil
	__Mesh__00000002_.PlaneGeometry = nil
	__Mesh__00000002_.TubeGeometry = nil
	__Mesh__00000002_.ExtrudeGeometry = nil
	__Mesh__00000003_.MeshMaterialBasic = __MeshMaterialBasic__00000001_
	__Mesh__00000003_.MeshPhysicalMaterial = nil
	__Mesh__00000003_.CylinderGeometry = nil
	__Mesh__00000003_.BoxGeometry = nil
	__Mesh__00000003_.SphereGeometry = nil
	__Mesh__00000003_.TorusGeometry = __TorusGeometry__00000000_
	__Mesh__00000003_.PlaneGeometry = nil
	__Mesh__00000003_.TubeGeometry = nil
	__Mesh__00000003_.ExtrudeGeometry = nil
	__Mesh__00000004_.MeshMaterialBasic = __MeshMaterialBasic__00000002_
	__Mesh__00000004_.MeshPhysicalMaterial = nil
	__Mesh__00000004_.CylinderGeometry = nil
	__Mesh__00000004_.BoxGeometry = nil
	__Mesh__00000004_.SphereGeometry = nil
	__Mesh__00000004_.TorusGeometry = nil
	__Mesh__00000004_.PlaneGeometry = __PlaneGeometry__00000000_
	__Mesh__00000004_.TubeGeometry = nil
	__Mesh__00000004_.ExtrudeGeometry = nil
	__Mesh__00000005_.MeshMaterialBasic = __MeshMaterialBasic__00000003_
	__Mesh__00000005_.MeshPhysicalMaterial = nil
	__Mesh__00000005_.CylinderGeometry = __CylinderGeometry__00000002_
	__Mesh__00000005_.BoxGeometry = nil
	__Mesh__00000005_.SphereGeometry = nil
	__Mesh__00000005_.TorusGeometry = nil
	__Mesh__00000005_.PlaneGeometry = nil
	__Mesh__00000005_.TubeGeometry = nil
	__Mesh__00000005_.ExtrudeGeometry = nil
	__Mesh__00000006_.MeshMaterialBasic = __MeshMaterialBasic__00000004_
	__Mesh__00000006_.MeshPhysicalMaterial = nil
	__Mesh__00000006_.CylinderGeometry = __CylinderGeometry__00000003_
	__Mesh__00000006_.BoxGeometry = nil
	__Mesh__00000006_.SphereGeometry = nil
	__Mesh__00000006_.TorusGeometry = nil
	__Mesh__00000006_.PlaneGeometry = nil
	__Mesh__00000006_.TubeGeometry = nil
	__Mesh__00000006_.ExtrudeGeometry = nil
	__Mesh__00000007_.MeshMaterialBasic = __MeshMaterialBasic__00000005_
	__Mesh__00000007_.MeshPhysicalMaterial = nil
	__Mesh__00000007_.CylinderGeometry = nil
	__Mesh__00000007_.BoxGeometry = nil
	__Mesh__00000007_.SphereGeometry = nil
	__Mesh__00000007_.TorusGeometry = nil
	__Mesh__00000007_.PlaneGeometry = nil
	__Mesh__00000007_.TubeGeometry = __TubeGeometry__00000000_
	__Mesh__00000007_.ExtrudeGeometry = nil
	__Mesh__00000008_.MeshMaterialBasic = __MeshMaterialBasic__00000006_
	__Mesh__00000008_.MeshPhysicalMaterial = nil
	__Mesh__00000008_.CylinderGeometry = nil
	__Mesh__00000008_.BoxGeometry = nil
	__Mesh__00000008_.SphereGeometry = nil
	__Mesh__00000008_.TorusGeometry = nil
	__Mesh__00000008_.PlaneGeometry = nil
	__Mesh__00000008_.TubeGeometry = nil
	__Mesh__00000008_.ExtrudeGeometry = __ExtrudeGeometry__00000000_
	__Shape__00000000_.Points = append(__Shape__00000000_.Points, __Vector2__00000000_)
	__Shape__00000000_.Points = append(__Shape__00000000_.Points, __Vector2__00000001_)
	__Shape__00000000_.Points = append(__Shape__00000000_.Points, __Vector2__00000002_)
	__Shape__00000000_.Points = append(__Shape__00000000_.Points, __Vector2__00000003_)
	__Shape__00000000_.Points = append(__Shape__00000000_.Points, __Vector2__00000004_)
	__TubeGeometry__00000000_.Path = __Curve__00000000_
}
