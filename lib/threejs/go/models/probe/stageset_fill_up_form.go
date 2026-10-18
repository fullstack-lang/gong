// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.AmbiantLight:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Intensity", inst.Intensity, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Canvass {
				if src.AmbiantLight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Canvas", "AmbiantLight", refNames, formGroup, probe.formStage)
		}
	case *models.BoxGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WidthSegments", inst.WidthSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightSegments", inst.HeightSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DepthSegments", inst.DepthSegments, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.BoxGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "BoxGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.BufferGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Vertices
			div := (&form.FormDiv{Name: "Vertices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Vertices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Vertices",
				Label: "Vertices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Faces
			div := (&form.FormDiv{Name: "Faces"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Faces {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Faces",
				Label: "Faces",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.BufferGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "BufferGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.Camera:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Z", inst.Z, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetX", inst.TargetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetY", inst.TargetY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetZ", inst.TargetZ, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fov", inst.Fov, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Canvass {
				if src.Camera == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Canvas", "Camera", refNames, formGroup, probe.formStage)
		}
	case *models.Canvas:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: DirectionalLights
			div := (&form.FormDiv{Name: "DirectionalLights"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DirectionalLights {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DirectionalLights",
				Label: "DirectionalLights",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AmbiantLight", inst.AmbiantLight, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.AmbiantLight](), probe.formStage)

		{
			// Slice of pointers: Meshs
			div := (&form.FormDiv{Name: "Meshs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Meshs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Meshs",
				Label: "Meshs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Camera", inst.Camera, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Camera](), probe.formStage)
		StageSetBasicFieldtoForm("IsWithLastRenderingUpdate", inst.IsWithLastRenderingUpdate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastRendering", inst.LastRendering, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame64BitsEncoded", inst.Frame64BitsEncoded, probe.formStage, formGroup)
	case *models.Curve:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Points
			div := (&form.FormDiv{Name: "Points"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Points {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Points",
				Label: "Points",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ExtrudeGeometrys {
				if src.ExtrudePath == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ExtrudeGeometry", "ExtrudePath", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeGeometrys {
				if src.Path == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeGeometry", "Path", refNames, formGroup, probe.formStage)
		}
	case *models.CylinderGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadiusTop", inst.RadiusTop, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadiusBottom", inst.RadiusBottom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialSegments", inst.RadialSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightSegments", inst.HeightSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OpenEnded", inst.OpenEnded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ThetaStart", inst.ThetaStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ThetaLength", inst.ThetaLength, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.CylinderGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "CylinderGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.DirectionalLight:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Z", inst.Z, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Intensity", inst.Intensity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithCastShadow", inst.IsWithCastShadow, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Canvass {
				for _, target := range src.DirectionalLights {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Canvas", "DirectionalLights", refNames, formGroup, probe.formStage)
		}
	case *models.ExtrudeGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Shape", inst.Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Shape](), probe.formStage)
		StageSetAssociationFieldToForm("ExtrudePath", inst.ExtrudePath, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Curve](), probe.formStage)
		StageSetBasicFieldtoForm("Steps", inst.Steps, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.ExtrudeGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "ExtrudeGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.Mesh:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Z", inst.Z, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MeshMaterialBasic", inst.MeshMaterialBasic, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MeshMaterialBasic](), probe.formStage)
		StageSetAssociationFieldToForm("MeshPhysicalMaterial", inst.MeshPhysicalMaterial, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MeshPhysicalMaterial](), probe.formStage)
		StageSetAssociationFieldToForm("CylinderGeometry", inst.CylinderGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CylinderGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("BoxGeometry", inst.BoxGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.BoxGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("SphereGeometry", inst.SphereGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SphereGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("TorusGeometry", inst.TorusGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TorusGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("PlaneGeometry", inst.PlaneGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PlaneGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("TubeGeometry", inst.TubeGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TubeGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("ExtrudeGeometry", inst.ExtrudeGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ExtrudeGeometry](), probe.formStage)
		StageSetAssociationFieldToForm("BufferGeometry", inst.BufferGeometry, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.BufferGeometry](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Canvass {
				for _, target := range src.Meshs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Canvas", "Meshs", refNames, formGroup, probe.formStage)
		}
	case *models.MeshMaterialBasic:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.MeshMaterialBasic == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "MeshMaterialBasic", refNames, formGroup, probe.formStage)
		}
	case *models.MeshPhysicalMaterial:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Wireframe", inst.Wireframe, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Opacity", inst.Opacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transparent", inst.Transparent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Visible", inst.Visible, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.MeshPhysicalMaterial == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "MeshPhysicalMaterial", refNames, formGroup, probe.formStage)
		}
	case *models.PlaneGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WidthSegments", inst.WidthSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightSegments", inst.HeightSegments, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.PlaneGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "PlaneGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Points
			div := (&form.FormDiv{Name: "Points"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Points {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Points",
				Label: "Points",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ExtrudeGeometrys {
				if src.Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ExtrudeGeometry", "Shape", refNames, formGroup, probe.formStage)
		}
	case *models.SphereGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Radius", inst.Radius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WidthSegments", inst.WidthSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightSegments", inst.HeightSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PhiStart", inst.PhiStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PhiLength", inst.PhiLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ThetaStart", inst.ThetaStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ThetaLength", inst.ThetaLength, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.SphereGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "SphereGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.TorusGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Radius", inst.Radius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tube", inst.Tube, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialSegments", inst.RadialSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TubularSegments", inst.TubularSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arc", inst.Arc, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.TorusGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "TorusGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.Triangle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("V1", inst.V1, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("V2", inst.V2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("V3", inst.V3, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.BufferGeometrys {
				for _, target := range src.Faces {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BufferGeometry", "Faces", refNames, formGroup, probe.formStage)
		}
	case *models.TubeGeometry:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Path", inst.Path, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Curve](), probe.formStage)
		StageSetBasicFieldtoForm("TubularSegments", inst.TubularSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Radius", inst.Radius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialSegments", inst.RadialSegments, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Closed", inst.Closed, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Meshs {
				if src.TubeGeometry == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Mesh", "TubeGeometry", refNames, formGroup, probe.formStage)
		}
	case *models.Vector2:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Shapes {
				for _, target := range src.Points {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Shape", "Points", refNames, formGroup, probe.formStage)
		}
	case *models.Vector3:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Z", inst.Z, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.BufferGeometrys {
				for _, target := range src.Vertices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BufferGeometry", "Vertices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Curves {
				for _, target := range src.Points {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Curve", "Points", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
