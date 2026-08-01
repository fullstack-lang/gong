package models

import (
	"fmt"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)


func (stager *Stager) addPointSpheres(points []*threejs.Vector3, color string, canvas *threejs.Canvas, namePrefix string, dy float64) {
	for i, pt := range points {
		sphereColor := color
		radius := 2.0
		if i % 40 == 0 {
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
			}).Stage(stager.threejsStage),
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name: fmt.Sprintf("%s SphereMat %d", namePrefix, i),
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{
					Color: sphereColor,
				},
			}).Stage(stager.threejsStage),
		}).Stage(stager.threejsStage)
		canvas.Meshs = append(canvas.Meshs, sphere)
	}
}
