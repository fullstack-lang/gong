// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	var zero T
	if cb.Instance == zero {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__AmbiantLightFormCallback(
	_instance *models.AmbiantLight,
	probe *Probe,
	formGroup *form.FormGroup,
) (ambiantlightFormCallback *FormCallback[*models.AmbiantLight]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAmbiantLightFields,
	)
}

type AmbiantLightFormCallback = FormCallback[*models.AmbiantLight]

func saveAmbiantLightFields(
	_instance *models.AmbiantLight,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Intensity":
			FormDivBasicFieldToField(&(_instance.Intensity), formDiv)
		}
	}
}

func __gong__New__BoxGeometryFormCallback(
	_instance *models.BoxGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (boxgeometryFormCallback *FormCallback[*models.BoxGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBoxGeometryFields,
	)
}

type BoxGeometryFormCallback = FormCallback[*models.BoxGeometry]

func saveBoxGeometryFields(
	_instance *models.BoxGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(_instance.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(_instance.HeightSegments), formDiv)
		case "DepthSegments":
			FormDivBasicFieldToField(&(_instance.DepthSegments), formDiv)
		}
	}
}

func __gong__New__BufferGeometryFormCallback(
	_instance *models.BufferGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (buffergeometryFormCallback *FormCallback[*models.BufferGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBufferGeometryFields,
	)
}

type BufferGeometryFormCallback = FormCallback[*models.BufferGeometry]

func saveBufferGeometryFields(
	_instance *models.BufferGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Vertices":
			FormDivSliceOfPointersToField(_instance, "Vertices", &(_instance.Vertices), formDiv, probe)
		case "Faces":
			FormDivSliceOfPointersToField(_instance, "Faces", &(_instance.Faces), formDiv, probe)
		}
	}
}

func __gong__New__CameraFormCallback(
	_instance *models.Camera,
	probe *Probe,
	formGroup *form.FormGroup,
) (cameraFormCallback *FormCallback[*models.Camera]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCameraFields,
	)
}

type CameraFormCallback = FormCallback[*models.Camera]

func saveCameraFields(
	_instance *models.Camera,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(_instance.Z), formDiv)
		case "TargetX":
			FormDivBasicFieldToField(&(_instance.TargetX), formDiv)
		case "TargetY":
			FormDivBasicFieldToField(&(_instance.TargetY), formDiv)
		case "TargetZ":
			FormDivBasicFieldToField(&(_instance.TargetZ), formDiv)
		case "Fov":
			FormDivBasicFieldToField(&(_instance.Fov), formDiv)
		}
	}
}

func __gong__New__CanvasFormCallback(
	_instance *models.Canvas,
	probe *Probe,
	formGroup *form.FormGroup,
) (canvasFormCallback *FormCallback[*models.Canvas]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCanvasFields,
	)
}

type CanvasFormCallback = FormCallback[*models.Canvas]

func saveCanvasFields(
	_instance *models.Canvas,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DirectionalLights":
			FormDivSliceOfPointersToField(_instance, "DirectionalLights", &(_instance.DirectionalLights), formDiv, probe)
		case "AmbiantLight":
			FormDivSelectFieldToField(&(_instance.AmbiantLight), probe.stageOfInterest, formDiv)
		case "Meshs":
			FormDivSliceOfPointersToField(_instance, "Meshs", &(_instance.Meshs), formDiv, probe)
		case "Camera":
			FormDivSelectFieldToField(&(_instance.Camera), probe.stageOfInterest, formDiv)
		case "IsWithLastRenderingUpdate":
			FormDivBasicFieldToField(&(_instance.IsWithLastRenderingUpdate), formDiv)
		case "LastRendering":
			FormDivTimeFieldToField(&(_instance.LastRendering), formDiv, false)
		case "Frame64BitsEncoded":
			FormDivBasicFieldToField(&(_instance.Frame64BitsEncoded), formDiv)
		}
	}
}

func __gong__New__CurveFormCallback(
	_instance *models.Curve,
	probe *Probe,
	formGroup *form.FormGroup,
) (curveFormCallback *FormCallback[*models.Curve]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCurveFields,
	)
}

type CurveFormCallback = FormCallback[*models.Curve]

func saveCurveFields(
	_instance *models.Curve,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Points":
			FormDivSliceOfPointersToField(_instance, "Points", &(_instance.Points), formDiv, probe)
		}
	}
}

func __gong__New__CylinderGeometryFormCallback(
	_instance *models.CylinderGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (cylindergeometryFormCallback *FormCallback[*models.CylinderGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCylinderGeometryFields,
	)
}

type CylinderGeometryFormCallback = FormCallback[*models.CylinderGeometry]

func saveCylinderGeometryFields(
	_instance *models.CylinderGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RadiusTop":
			FormDivBasicFieldToField(&(_instance.RadiusTop), formDiv)
		case "RadiusBottom":
			FormDivBasicFieldToField(&(_instance.RadiusBottom), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(_instance.RadialSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(_instance.HeightSegments), formDiv)
		case "OpenEnded":
			FormDivBasicFieldToField(&(_instance.OpenEnded), formDiv)
		case "ThetaStart":
			FormDivBasicFieldToField(&(_instance.ThetaStart), formDiv)
		case "ThetaLength":
			FormDivBasicFieldToField(&(_instance.ThetaLength), formDiv)
		}
	}
}

func __gong__New__DirectionalLightFormCallback(
	_instance *models.DirectionalLight,
	probe *Probe,
	formGroup *form.FormGroup,
) (directionallightFormCallback *FormCallback[*models.DirectionalLight]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDirectionalLightFields,
	)
}

type DirectionalLightFormCallback = FormCallback[*models.DirectionalLight]

func saveDirectionalLightFields(
	_instance *models.DirectionalLight,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(_instance.Z), formDiv)
		case "Intensity":
			FormDivBasicFieldToField(&(_instance.Intensity), formDiv)
		case "IsWithCastShadow":
			FormDivBasicFieldToField(&(_instance.IsWithCastShadow), formDiv)
		case "Canvas:DirectionalLights":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DirectionalLights", func(owner *models.Canvas) *[]*models.DirectionalLight { return &owner.DirectionalLights })
		}
	}
}

func __gong__New__ExtrudeGeometryFormCallback(
	_instance *models.ExtrudeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (extrudegeometryFormCallback *FormCallback[*models.ExtrudeGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExtrudeGeometryFields,
	)
}

type ExtrudeGeometryFormCallback = FormCallback[*models.ExtrudeGeometry]

func saveExtrudeGeometryFields(
	_instance *models.ExtrudeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Shape":
			FormDivSelectFieldToField(&(_instance.Shape), probe.stageOfInterest, formDiv)
		case "ExtrudePath":
			FormDivSelectFieldToField(&(_instance.ExtrudePath), probe.stageOfInterest, formDiv)
		case "Steps":
			FormDivBasicFieldToField(&(_instance.Steps), formDiv)
		}
	}
}

func __gong__New__MeshFormCallback(
	_instance *models.Mesh,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshFormCallback *FormCallback[*models.Mesh]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeshFields,
	)
}

type MeshFormCallback = FormCallback[*models.Mesh]

func saveMeshFields(
	_instance *models.Mesh,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(_instance.Z), formDiv)
		case "MeshMaterialBasic":
			FormDivSelectFieldToField(&(_instance.MeshMaterialBasic), probe.stageOfInterest, formDiv)
		case "MeshPhysicalMaterial":
			FormDivSelectFieldToField(&(_instance.MeshPhysicalMaterial), probe.stageOfInterest, formDiv)
		case "CylinderGeometry":
			FormDivSelectFieldToField(&(_instance.CylinderGeometry), probe.stageOfInterest, formDiv)
		case "BoxGeometry":
			FormDivSelectFieldToField(&(_instance.BoxGeometry), probe.stageOfInterest, formDiv)
		case "SphereGeometry":
			FormDivSelectFieldToField(&(_instance.SphereGeometry), probe.stageOfInterest, formDiv)
		case "TorusGeometry":
			FormDivSelectFieldToField(&(_instance.TorusGeometry), probe.stageOfInterest, formDiv)
		case "PlaneGeometry":
			FormDivSelectFieldToField(&(_instance.PlaneGeometry), probe.stageOfInterest, formDiv)
		case "TubeGeometry":
			FormDivSelectFieldToField(&(_instance.TubeGeometry), probe.stageOfInterest, formDiv)
		case "ExtrudeGeometry":
			FormDivSelectFieldToField(&(_instance.ExtrudeGeometry), probe.stageOfInterest, formDiv)
		case "BufferGeometry":
			FormDivSelectFieldToField(&(_instance.BufferGeometry), probe.stageOfInterest, formDiv)
		case "Canvas:Meshs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Meshs", func(owner *models.Canvas) *[]*models.Mesh { return &owner.Meshs })
		}
	}
}

func __gong__New__MeshMaterialBasicFormCallback(
	_instance *models.MeshMaterialBasic,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshmaterialbasicFormCallback *FormCallback[*models.MeshMaterialBasic]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeshMaterialBasicFields,
	)
}

type MeshMaterialBasicFormCallback = FormCallback[*models.MeshMaterialBasic]

func saveMeshMaterialBasicFields(
	_instance *models.MeshMaterialBasic,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		}
	}
}

func __gong__New__MeshPhysicalMaterialFormCallback(
	_instance *models.MeshPhysicalMaterial,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshphysicalmaterialFormCallback *FormCallback[*models.MeshPhysicalMaterial]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeshPhysicalMaterialFields,
	)
}

type MeshPhysicalMaterialFormCallback = FormCallback[*models.MeshPhysicalMaterial]

func saveMeshPhysicalMaterialFields(
	_instance *models.MeshPhysicalMaterial,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Wireframe":
			FormDivBasicFieldToField(&(_instance.Wireframe), formDiv)
		case "Opacity":
			FormDivBasicFieldToField(&(_instance.Opacity), formDiv)
		case "Transparent":
			FormDivBasicFieldToField(&(_instance.Transparent), formDiv)
		case "Visible":
			FormDivBasicFieldToField(&(_instance.Visible), formDiv)
		}
	}
}

func __gong__New__PlaneGeometryFormCallback(
	_instance *models.PlaneGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (planegeometryFormCallback *FormCallback[*models.PlaneGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlaneGeometryFields,
	)
}

type PlaneGeometryFormCallback = FormCallback[*models.PlaneGeometry]

func savePlaneGeometryFields(
	_instance *models.PlaneGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(_instance.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(_instance.HeightSegments), formDiv)
		}
	}
}

func __gong__New__ShapeFormCallback(
	_instance *models.Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shapeFormCallback *FormCallback[*models.Shape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShapeFields,
	)
}

type ShapeFormCallback = FormCallback[*models.Shape]

func saveShapeFields(
	_instance *models.Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Points":
			FormDivSliceOfPointersToField(_instance, "Points", &(_instance.Points), formDiv, probe)
		}
	}
}

func __gong__New__SphereGeometryFormCallback(
	_instance *models.SphereGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (spheregeometryFormCallback *FormCallback[*models.SphereGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSphereGeometryFields,
	)
}

type SphereGeometryFormCallback = FormCallback[*models.SphereGeometry]

func saveSphereGeometryFields(
	_instance *models.SphereGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(_instance.Radius), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(_instance.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(_instance.HeightSegments), formDiv)
		case "PhiStart":
			FormDivBasicFieldToField(&(_instance.PhiStart), formDiv)
		case "PhiLength":
			FormDivBasicFieldToField(&(_instance.PhiLength), formDiv)
		case "ThetaStart":
			FormDivBasicFieldToField(&(_instance.ThetaStart), formDiv)
		case "ThetaLength":
			FormDivBasicFieldToField(&(_instance.ThetaLength), formDiv)
		}
	}
}

func __gong__New__TorusGeometryFormCallback(
	_instance *models.TorusGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (torusgeometryFormCallback *FormCallback[*models.TorusGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTorusGeometryFields,
	)
}

type TorusGeometryFormCallback = FormCallback[*models.TorusGeometry]

func saveTorusGeometryFields(
	_instance *models.TorusGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(_instance.Radius), formDiv)
		case "Tube":
			FormDivBasicFieldToField(&(_instance.Tube), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(_instance.RadialSegments), formDiv)
		case "TubularSegments":
			FormDivBasicFieldToField(&(_instance.TubularSegments), formDiv)
		case "Arc":
			FormDivBasicFieldToField(&(_instance.Arc), formDiv)
		}
	}
}

func __gong__New__TriangleFormCallback(
	_instance *models.Triangle,
	probe *Probe,
	formGroup *form.FormGroup,
) (triangleFormCallback *FormCallback[*models.Triangle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTriangleFields,
	)
}

type TriangleFormCallback = FormCallback[*models.Triangle]

func saveTriangleFields(
	_instance *models.Triangle,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "V1":
			FormDivBasicFieldToField(&(_instance.V1), formDiv)
		case "V2":
			FormDivBasicFieldToField(&(_instance.V2), formDiv)
		case "V3":
			FormDivBasicFieldToField(&(_instance.V3), formDiv)
		case "BufferGeometry:Faces":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Faces", func(owner *models.BufferGeometry) *[]*models.Triangle { return &owner.Faces })
		}
	}
}

func __gong__New__TubeGeometryFormCallback(
	_instance *models.TubeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (tubegeometryFormCallback *FormCallback[*models.TubeGeometry]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTubeGeometryFields,
	)
}

type TubeGeometryFormCallback = FormCallback[*models.TubeGeometry]

func saveTubeGeometryFields(
	_instance *models.TubeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Path":
			FormDivSelectFieldToField(&(_instance.Path), probe.stageOfInterest, formDiv)
		case "TubularSegments":
			FormDivBasicFieldToField(&(_instance.TubularSegments), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(_instance.Radius), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(_instance.RadialSegments), formDiv)
		case "Closed":
			FormDivBasicFieldToField(&(_instance.Closed), formDiv)
		}
	}
}

func __gong__New__Vector2FormCallback(
	_instance *models.Vector2,
	probe *Probe,
	formGroup *form.FormGroup,
) (vector2FormCallback *FormCallback[*models.Vector2]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVector2Fields,
	)
}

type Vector2FormCallback = FormCallback[*models.Vector2]

func saveVector2Fields(
	_instance *models.Vector2,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Shape:Points":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Points", func(owner *models.Shape) *[]*models.Vector2 { return &owner.Points })
		}
	}
}

func __gong__New__Vector3FormCallback(
	_instance *models.Vector3,
	probe *Probe,
	formGroup *form.FormGroup,
) (vector3FormCallback *FormCallback[*models.Vector3]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVector3Fields,
	)
}

type Vector3FormCallback = FormCallback[*models.Vector3]

func saveVector3Fields(
	_instance *models.Vector3,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(_instance.Z), formDiv)
		case "BufferGeometry:Vertices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Vertices", func(owner *models.BufferGeometry) *[]*models.Vector3 { return &owner.Vertices })
		case "Curve:Points":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Points", func(owner *models.Curve) *[]*models.Vector3 { return &owner.Points })
		}
	}
}

