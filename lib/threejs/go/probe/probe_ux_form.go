// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AmbiantLightFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AmbiantLight", true)
			} else {
				FillUpFormFromGongstruct(onSave.ambiantlight, probe)
			}
		case *BoxGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BoxGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.boxgeometry, probe)
			}
		case *CanvasFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Canvas", true)
			} else {
				FillUpFormFromGongstruct(onSave.canvas, probe)
			}
		case *CurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Curve", true)
			} else {
				FillUpFormFromGongstruct(onSave.curve, probe)
			}
		case *CylinderGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CylinderGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.cylindergeometry, probe)
			}
		case *DirectionalLightFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DirectionalLight", true)
			} else {
				FillUpFormFromGongstruct(onSave.directionallight, probe)
			}
		case *ExtrudeGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ExtrudeGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.extrudegeometry, probe)
			}
		case *MeshFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Mesh", true)
			} else {
				FillUpFormFromGongstruct(onSave.mesh, probe)
			}
		case *MeshMaterialBasicFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MeshMaterialBasic", true)
			} else {
				FillUpFormFromGongstruct(onSave.meshmaterialbasic, probe)
			}
		case *MeshPhysicalMaterialFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MeshPhysicalMaterial", true)
			} else {
				FillUpFormFromGongstruct(onSave.meshphysicalmaterial, probe)
			}
		case *PlaneGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PlaneGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.planegeometry, probe)
			}
		case *ShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Shape", true)
			} else {
				FillUpFormFromGongstruct(onSave.shape, probe)
			}
		case *SphereGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SphereGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.spheregeometry, probe)
			}
		case *TorusGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TorusGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.torusgeometry, probe)
			}
		case *TubeGeometryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TubeGeometry", true)
			} else {
				FillUpFormFromGongstruct(onSave.tubegeometry, probe)
			}
		case *Vector2FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Vector2", true)
			} else {
				FillUpFormFromGongstruct(onSave.vector2, probe)
			}
		case *Vector3FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Vector3", true)
			} else {
				FillUpFormFromGongstruct(onSave.vector3, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "AmbiantLight":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AmbiantLight Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AmbiantLightFormCallback(
			nil,
			probe,
			formGroup,
		)
		ambiantlight := new(models.AmbiantLight)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ambiantlight, formGroup, probe)
	case "BoxGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BoxGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BoxGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		boxgeometry := new(models.BoxGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(boxgeometry, formGroup, probe)
	case "Canvas":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Canvas Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CanvasFormCallback(
			nil,
			probe,
			formGroup,
		)
		canvas := new(models.Canvas)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(canvas, formGroup, probe)
	case "Curve":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Curve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		curve := new(models.Curve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(curve, formGroup, probe)
	case "CylinderGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CylinderGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CylinderGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		cylindergeometry := new(models.CylinderGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cylindergeometry, formGroup, probe)
	case "DirectionalLight":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DirectionalLight Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DirectionalLightFormCallback(
			nil,
			probe,
			formGroup,
		)
		directionallight := new(models.DirectionalLight)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(directionallight, formGroup, probe)
	case "ExtrudeGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ExtrudeGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtrudeGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		extrudegeometry := new(models.ExtrudeGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extrudegeometry, formGroup, probe)
	case "Mesh":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Mesh Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MeshFormCallback(
			nil,
			probe,
			formGroup,
		)
		mesh := new(models.Mesh)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mesh, formGroup, probe)
	case "MeshMaterialBasic":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MeshMaterialBasic Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MeshMaterialBasicFormCallback(
			nil,
			probe,
			formGroup,
		)
		meshmaterialbasic := new(models.MeshMaterialBasic)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(meshmaterialbasic, formGroup, probe)
	case "MeshPhysicalMaterial":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MeshPhysicalMaterial Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MeshPhysicalMaterialFormCallback(
			nil,
			probe,
			formGroup,
		)
		meshphysicalmaterial := new(models.MeshPhysicalMaterial)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(meshphysicalmaterial, formGroup, probe)
	case "PlaneGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PlaneGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlaneGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		planegeometry := new(models.PlaneGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(planegeometry, formGroup, probe)
	case "Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		shape := new(models.Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shape, formGroup, probe)
	case "SphereGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SphereGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SphereGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		spheregeometry := new(models.SphereGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spheregeometry, formGroup, probe)
	case "TorusGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TorusGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TorusGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		torusgeometry := new(models.TorusGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(torusgeometry, formGroup, probe)
	case "TubeGeometry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TubeGeometry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TubeGeometryFormCallback(
			nil,
			probe,
			formGroup,
		)
		tubegeometry := new(models.TubeGeometry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tubegeometry, formGroup, probe)
	case "Vector2":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Vector2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Vector2FormCallback(
			nil,
			probe,
			formGroup,
		)
		vector2 := new(models.Vector2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vector2, formGroup, probe)
	case "Vector3":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Vector3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Vector3FormCallback(
			nil,
			probe,
			formGroup,
		)
		vector3 := new(models.Vector3)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vector3, formGroup, probe)
	}
	formStage.Commit()
}
