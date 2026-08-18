package vasestage3d

import (
	"fmt"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) addPointSpheres(stager *models.Stager, points []*threejs.Vector3, color string, canvas *threejs.Canvas, namePrefix string, dy float64, numPointsPerRep int) {
	threejsStage := stager.GetThreejsStage()

	for i, pt := range points {
		sphereColor := color
		radius := 2.0

		localIdx := i
		if numPointsPerRep > 0 {
			localIdx = i % numPointsPerRep
		}

		if localIdx%20 == 0 {
			sphereColor = "yellow"
			radius = 4.0
		}

		sphere := (&threejs.Mesh{
			Name: fmt.Sprintf("%s Sphere %d", namePrefix, i),
			Position: threejs.Position{
				X: pt.X,
				Y: pt.Y + dy,
				Z: pt.Z,
			},
			SphereGeometry: (&threejs.SphereGeometry{
				Name:   fmt.Sprintf("%s SphereGeom %d", namePrefix, i),
				Radius: radius,
			}).Stage(threejsStage),
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name: fmt.Sprintf("%s SphereMat %d", namePrefix, i),
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{
					Color: sphereColor,
				},
			}).Stage(threejsStage),
		}).Stage(threejsStage)
		canvas.Meshs = append(canvas.Meshs, sphere)
	}
}
