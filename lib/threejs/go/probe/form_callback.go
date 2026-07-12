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

// insertion point
func __gong__New__AmbiantLightFormCallback(
	ambiantlight *models.AmbiantLight,
	probe *Probe,
	formGroup *form.FormGroup,
) (ambiantlightFormCallback *AmbiantLightFormCallback) {
	ambiantlightFormCallback = new(AmbiantLightFormCallback)
	ambiantlightFormCallback.probe = probe
	ambiantlightFormCallback.ambiantlight = ambiantlight
	ambiantlightFormCallback.formGroup = formGroup

	ambiantlightFormCallback.CreationMode = (ambiantlight == nil)

	return
}

type AmbiantLightFormCallback struct {
	ambiantlight *models.AmbiantLight

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (ambiantlightFormCallback *AmbiantLightFormCallback) OnSave() {
	ambiantlightFormCallback.probe.stageOfInterest.Lock()
	defer ambiantlightFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AmbiantLightFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ambiantlightFormCallback.probe.formStage.Checkout()

	if ambiantlightFormCallback.ambiantlight == nil {
		ambiantlightFormCallback.ambiantlight = new(models.AmbiantLight).Stage(ambiantlightFormCallback.probe.stageOfInterest)
	}
	ambiantlight_ := ambiantlightFormCallback.ambiantlight
	_ = ambiantlight_

	for _, formDiv := range ambiantlightFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ambiantlight_.Name), formDiv)
		case "Intensity":
			FormDivBasicFieldToField(&(ambiantlight_.Intensity), formDiv)
		}
	}

	// manage the suppress operation
	if ambiantlightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ambiantlight_.Unstage(ambiantlightFormCallback.probe.stageOfInterest)
	}

	ambiantlightFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AmbiantLight](
		ambiantlightFormCallback.probe,
	)

	// display a new form by reset the form stage
	if ambiantlightFormCallback.CreationMode || ambiantlightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ambiantlightFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(ambiantlightFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AmbiantLightFormCallback(
			nil,
			ambiantlightFormCallback.probe,
			newFormGroup,
		)
		ambiantlight := new(models.AmbiantLight)
		FillUpForm(ambiantlight, newFormGroup, ambiantlightFormCallback.probe)
		ambiantlightFormCallback.probe.formStage.Commit()
	}

	ambiantlightFormCallback.probe.ux_tree()
}
func __gong__New__BoxGeometryFormCallback(
	boxgeometry *models.BoxGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (boxgeometryFormCallback *BoxGeometryFormCallback) {
	boxgeometryFormCallback = new(BoxGeometryFormCallback)
	boxgeometryFormCallback.probe = probe
	boxgeometryFormCallback.boxgeometry = boxgeometry
	boxgeometryFormCallback.formGroup = formGroup

	boxgeometryFormCallback.CreationMode = (boxgeometry == nil)

	return
}

type BoxGeometryFormCallback struct {
	boxgeometry *models.BoxGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (boxgeometryFormCallback *BoxGeometryFormCallback) OnSave() {
	boxgeometryFormCallback.probe.stageOfInterest.Lock()
	defer boxgeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BoxGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	boxgeometryFormCallback.probe.formStage.Checkout()

	if boxgeometryFormCallback.boxgeometry == nil {
		boxgeometryFormCallback.boxgeometry = new(models.BoxGeometry).Stage(boxgeometryFormCallback.probe.stageOfInterest)
	}
	boxgeometry_ := boxgeometryFormCallback.boxgeometry
	_ = boxgeometry_

	for _, formDiv := range boxgeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(boxgeometry_.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(boxgeometry_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(boxgeometry_.Height), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(boxgeometry_.Depth), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(boxgeometry_.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(boxgeometry_.HeightSegments), formDiv)
		case "DepthSegments":
			FormDivBasicFieldToField(&(boxgeometry_.DepthSegments), formDiv)
		}
	}

	// manage the suppress operation
	if boxgeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		boxgeometry_.Unstage(boxgeometryFormCallback.probe.stageOfInterest)
	}

	boxgeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BoxGeometry](
		boxgeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if boxgeometryFormCallback.CreationMode || boxgeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		boxgeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(boxgeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BoxGeometryFormCallback(
			nil,
			boxgeometryFormCallback.probe,
			newFormGroup,
		)
		boxgeometry := new(models.BoxGeometry)
		FillUpForm(boxgeometry, newFormGroup, boxgeometryFormCallback.probe)
		boxgeometryFormCallback.probe.formStage.Commit()
	}

	boxgeometryFormCallback.probe.ux_tree()
}
func __gong__New__CameraFormCallback(
	camera *models.Camera,
	probe *Probe,
	formGroup *form.FormGroup,
) (cameraFormCallback *CameraFormCallback) {
	cameraFormCallback = new(CameraFormCallback)
	cameraFormCallback.probe = probe
	cameraFormCallback.camera = camera
	cameraFormCallback.formGroup = formGroup

	cameraFormCallback.CreationMode = (camera == nil)

	return
}

type CameraFormCallback struct {
	camera *models.Camera

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (cameraFormCallback *CameraFormCallback) OnSave() {
	cameraFormCallback.probe.stageOfInterest.Lock()
	defer cameraFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CameraFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cameraFormCallback.probe.formStage.Checkout()

	if cameraFormCallback.camera == nil {
		cameraFormCallback.camera = new(models.Camera).Stage(cameraFormCallback.probe.stageOfInterest)
	}
	camera_ := cameraFormCallback.camera
	_ = camera_

	for _, formDiv := range cameraFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(camera_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(camera_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(camera_.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(camera_.Z), formDiv)
		case "TargetX":
			FormDivBasicFieldToField(&(camera_.TargetX), formDiv)
		case "TargetY":
			FormDivBasicFieldToField(&(camera_.TargetY), formDiv)
		case "TargetZ":
			FormDivBasicFieldToField(&(camera_.TargetZ), formDiv)
		case "Fov":
			FormDivBasicFieldToField(&(camera_.Fov), formDiv)
		}
	}

	// manage the suppress operation
	if cameraFormCallback.formGroup.HasSuppressButtonBeenPressed {
		camera_.Unstage(cameraFormCallback.probe.stageOfInterest)
	}

	cameraFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Camera](
		cameraFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cameraFormCallback.CreationMode || cameraFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cameraFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cameraFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CameraFormCallback(
			nil,
			cameraFormCallback.probe,
			newFormGroup,
		)
		camera := new(models.Camera)
		FillUpForm(camera, newFormGroup, cameraFormCallback.probe)
		cameraFormCallback.probe.formStage.Commit()
	}

	cameraFormCallback.probe.ux_tree()
}
func __gong__New__CanvasFormCallback(
	canvas *models.Canvas,
	probe *Probe,
	formGroup *form.FormGroup,
) (canvasFormCallback *CanvasFormCallback) {
	canvasFormCallback = new(CanvasFormCallback)
	canvasFormCallback.probe = probe
	canvasFormCallback.canvas = canvas
	canvasFormCallback.formGroup = formGroup

	canvasFormCallback.CreationMode = (canvas == nil)

	return
}

type CanvasFormCallback struct {
	canvas *models.Canvas

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (canvasFormCallback *CanvasFormCallback) OnSave() {
	canvasFormCallback.probe.stageOfInterest.Lock()
	defer canvasFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CanvasFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	canvasFormCallback.probe.formStage.Checkout()

	if canvasFormCallback.canvas == nil {
		canvasFormCallback.canvas = new(models.Canvas).Stage(canvasFormCallback.probe.stageOfInterest)
	}
	canvas_ := canvasFormCallback.canvas
	_ = canvas_

	for _, formDiv := range canvasFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(canvas_.Name), formDiv)
		case "DirectionalLights":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DirectionalLight](canvasFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DirectionalLight, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DirectionalLight)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					canvasFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DirectionalLight](canvasFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			canvas_.DirectionalLights = instanceSlice
			canvasFormCallback.probe.UpdateSliceOfPointersCallback(canvas_, "DirectionalLights", &canvas_.DirectionalLights)

		case "AmbiantLight":
			FormDivSelectFieldToField(&(canvas_.AmbiantLight), canvasFormCallback.probe.stageOfInterest, formDiv)
		case "Meshs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Mesh](canvasFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Mesh, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Mesh)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					canvasFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Mesh](canvasFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			canvas_.Meshs = instanceSlice
			canvasFormCallback.probe.UpdateSliceOfPointersCallback(canvas_, "Meshs", &canvas_.Meshs)

		case "Camera":
			FormDivSelectFieldToField(&(canvas_.Camera), canvasFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if canvasFormCallback.formGroup.HasSuppressButtonBeenPressed {
		canvas_.Unstage(canvasFormCallback.probe.stageOfInterest)
	}

	canvasFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Canvas](
		canvasFormCallback.probe,
	)

	// display a new form by reset the form stage
	if canvasFormCallback.CreationMode || canvasFormCallback.formGroup.HasSuppressButtonBeenPressed {
		canvasFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(canvasFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CanvasFormCallback(
			nil,
			canvasFormCallback.probe,
			newFormGroup,
		)
		canvas := new(models.Canvas)
		FillUpForm(canvas, newFormGroup, canvasFormCallback.probe)
		canvasFormCallback.probe.formStage.Commit()
	}

	canvasFormCallback.probe.ux_tree()
}
func __gong__New__CurveFormCallback(
	curve *models.Curve,
	probe *Probe,
	formGroup *form.FormGroup,
) (curveFormCallback *CurveFormCallback) {
	curveFormCallback = new(CurveFormCallback)
	curveFormCallback.probe = probe
	curveFormCallback.curve = curve
	curveFormCallback.formGroup = formGroup

	curveFormCallback.CreationMode = (curve == nil)

	return
}

type CurveFormCallback struct {
	curve *models.Curve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (curveFormCallback *CurveFormCallback) OnSave() {
	curveFormCallback.probe.stageOfInterest.Lock()
	defer curveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	curveFormCallback.probe.formStage.Checkout()

	if curveFormCallback.curve == nil {
		curveFormCallback.curve = new(models.Curve).Stage(curveFormCallback.probe.stageOfInterest)
	}
	curve_ := curveFormCallback.curve
	_ = curve_

	for _, formDiv := range curveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(curve_.Name), formDiv)
		case "Points":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Vector3](curveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Vector3, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Vector3)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					curveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Vector3](curveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			curve_.Points = instanceSlice
			curveFormCallback.probe.UpdateSliceOfPointersCallback(curve_, "Points", &curve_.Points)

		}
	}

	// manage the suppress operation
	if curveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		curve_.Unstage(curveFormCallback.probe.stageOfInterest)
	}

	curveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Curve](
		curveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if curveFormCallback.CreationMode || curveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		curveFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(curveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CurveFormCallback(
			nil,
			curveFormCallback.probe,
			newFormGroup,
		)
		curve := new(models.Curve)
		FillUpForm(curve, newFormGroup, curveFormCallback.probe)
		curveFormCallback.probe.formStage.Commit()
	}

	curveFormCallback.probe.ux_tree()
}
func __gong__New__CylinderGeometryFormCallback(
	cylindergeometry *models.CylinderGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (cylindergeometryFormCallback *CylinderGeometryFormCallback) {
	cylindergeometryFormCallback = new(CylinderGeometryFormCallback)
	cylindergeometryFormCallback.probe = probe
	cylindergeometryFormCallback.cylindergeometry = cylindergeometry
	cylindergeometryFormCallback.formGroup = formGroup

	cylindergeometryFormCallback.CreationMode = (cylindergeometry == nil)

	return
}

type CylinderGeometryFormCallback struct {
	cylindergeometry *models.CylinderGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (cylindergeometryFormCallback *CylinderGeometryFormCallback) OnSave() {
	cylindergeometryFormCallback.probe.stageOfInterest.Lock()
	defer cylindergeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CylinderGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cylindergeometryFormCallback.probe.formStage.Checkout()

	if cylindergeometryFormCallback.cylindergeometry == nil {
		cylindergeometryFormCallback.cylindergeometry = new(models.CylinderGeometry).Stage(cylindergeometryFormCallback.probe.stageOfInterest)
	}
	cylindergeometry_ := cylindergeometryFormCallback.cylindergeometry
	_ = cylindergeometry_

	for _, formDiv := range cylindergeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cylindergeometry_.Name), formDiv)
		case "RadiusTop":
			FormDivBasicFieldToField(&(cylindergeometry_.RadiusTop), formDiv)
		case "RadiusBottom":
			FormDivBasicFieldToField(&(cylindergeometry_.RadiusBottom), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(cylindergeometry_.Height), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(cylindergeometry_.RadialSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(cylindergeometry_.HeightSegments), formDiv)
		case "OpenEnded":
			FormDivBasicFieldToField(&(cylindergeometry_.OpenEnded), formDiv)
		case "ThetaStart":
			FormDivBasicFieldToField(&(cylindergeometry_.ThetaStart), formDiv)
		case "ThetaLength":
			FormDivBasicFieldToField(&(cylindergeometry_.ThetaLength), formDiv)
		}
	}

	// manage the suppress operation
	if cylindergeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cylindergeometry_.Unstage(cylindergeometryFormCallback.probe.stageOfInterest)
	}

	cylindergeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.CylinderGeometry](
		cylindergeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cylindergeometryFormCallback.CreationMode || cylindergeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cylindergeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cylindergeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CylinderGeometryFormCallback(
			nil,
			cylindergeometryFormCallback.probe,
			newFormGroup,
		)
		cylindergeometry := new(models.CylinderGeometry)
		FillUpForm(cylindergeometry, newFormGroup, cylindergeometryFormCallback.probe)
		cylindergeometryFormCallback.probe.formStage.Commit()
	}

	cylindergeometryFormCallback.probe.ux_tree()
}
func __gong__New__DirectionalLightFormCallback(
	directionallight *models.DirectionalLight,
	probe *Probe,
	formGroup *form.FormGroup,
) (directionallightFormCallback *DirectionalLightFormCallback) {
	directionallightFormCallback = new(DirectionalLightFormCallback)
	directionallightFormCallback.probe = probe
	directionallightFormCallback.directionallight = directionallight
	directionallightFormCallback.formGroup = formGroup

	directionallightFormCallback.CreationMode = (directionallight == nil)

	return
}

type DirectionalLightFormCallback struct {
	directionallight *models.DirectionalLight

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (directionallightFormCallback *DirectionalLightFormCallback) OnSave() {
	directionallightFormCallback.probe.stageOfInterest.Lock()
	defer directionallightFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DirectionalLightFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	directionallightFormCallback.probe.formStage.Checkout()

	if directionallightFormCallback.directionallight == nil {
		directionallightFormCallback.directionallight = new(models.DirectionalLight).Stage(directionallightFormCallback.probe.stageOfInterest)
	}
	directionallight_ := directionallightFormCallback.directionallight
	_ = directionallight_

	for _, formDiv := range directionallightFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(directionallight_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(directionallight_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(directionallight_.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(directionallight_.Z), formDiv)
		case "Intensity":
			FormDivBasicFieldToField(&(directionallight_.Intensity), formDiv)
		case "IsWithCastShadow":
			FormDivBasicFieldToField(&(directionallight_.IsWithCastShadow), formDiv)
		case "Canvas:DirectionalLights":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Canvas instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Canvas instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Canvas](directionallightFormCallback.probe.stageOfInterest)
			targetCanvasIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCanvasIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Canvas instances and update their DirectionalLights slice
			for _canvas := range *models.GetGongstructInstancesSetFromPointerType[*models.Canvas](directionallightFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(directionallightFormCallback.probe.stageOfInterest, _canvas)
				
				// if Canvas is selected
				if targetCanvasIDs[id] {
					// ensure directionallight_ is in _canvas.DirectionalLights
					found := false
					for _, _b := range _canvas.DirectionalLights {
						if _b == directionallight_ {
							found = true
							break
						}
					}
					if !found {
						_canvas.DirectionalLights = append(_canvas.DirectionalLights, directionallight_)
						directionallightFormCallback.probe.UpdateSliceOfPointersCallback(_canvas, "DirectionalLights", &_canvas.DirectionalLights)
					}
				} else {
					// ensure directionallight_ is NOT in _canvas.DirectionalLights
					idx := slices.Index(_canvas.DirectionalLights, directionallight_)
					if idx != -1 {
						_canvas.DirectionalLights = slices.Delete(_canvas.DirectionalLights, idx, idx+1)
						directionallightFormCallback.probe.UpdateSliceOfPointersCallback(_canvas, "DirectionalLights", &_canvas.DirectionalLights)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if directionallightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		directionallight_.Unstage(directionallightFormCallback.probe.stageOfInterest)
	}

	directionallightFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DirectionalLight](
		directionallightFormCallback.probe,
	)

	// display a new form by reset the form stage
	if directionallightFormCallback.CreationMode || directionallightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		directionallightFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(directionallightFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DirectionalLightFormCallback(
			nil,
			directionallightFormCallback.probe,
			newFormGroup,
		)
		directionallight := new(models.DirectionalLight)
		FillUpForm(directionallight, newFormGroup, directionallightFormCallback.probe)
		directionallightFormCallback.probe.formStage.Commit()
	}

	directionallightFormCallback.probe.ux_tree()
}
func __gong__New__ExtrudeGeometryFormCallback(
	extrudegeometry *models.ExtrudeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (extrudegeometryFormCallback *ExtrudeGeometryFormCallback) {
	extrudegeometryFormCallback = new(ExtrudeGeometryFormCallback)
	extrudegeometryFormCallback.probe = probe
	extrudegeometryFormCallback.extrudegeometry = extrudegeometry
	extrudegeometryFormCallback.formGroup = formGroup

	extrudegeometryFormCallback.CreationMode = (extrudegeometry == nil)

	return
}

type ExtrudeGeometryFormCallback struct {
	extrudegeometry *models.ExtrudeGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (extrudegeometryFormCallback *ExtrudeGeometryFormCallback) OnSave() {
	extrudegeometryFormCallback.probe.stageOfInterest.Lock()
	defer extrudegeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ExtrudeGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	extrudegeometryFormCallback.probe.formStage.Checkout()

	if extrudegeometryFormCallback.extrudegeometry == nil {
		extrudegeometryFormCallback.extrudegeometry = new(models.ExtrudeGeometry).Stage(extrudegeometryFormCallback.probe.stageOfInterest)
	}
	extrudegeometry_ := extrudegeometryFormCallback.extrudegeometry
	_ = extrudegeometry_

	for _, formDiv := range extrudegeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(extrudegeometry_.Name), formDiv)
		case "Shape":
			FormDivSelectFieldToField(&(extrudegeometry_.Shape), extrudegeometryFormCallback.probe.stageOfInterest, formDiv)
		case "ExtrudePath":
			FormDivSelectFieldToField(&(extrudegeometry_.ExtrudePath), extrudegeometryFormCallback.probe.stageOfInterest, formDiv)
		case "Steps":
			FormDivBasicFieldToField(&(extrudegeometry_.Steps), formDiv)
		}
	}

	// manage the suppress operation
	if extrudegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extrudegeometry_.Unstage(extrudegeometryFormCallback.probe.stageOfInterest)
	}

	extrudegeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ExtrudeGeometry](
		extrudegeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if extrudegeometryFormCallback.CreationMode || extrudegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extrudegeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(extrudegeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExtrudeGeometryFormCallback(
			nil,
			extrudegeometryFormCallback.probe,
			newFormGroup,
		)
		extrudegeometry := new(models.ExtrudeGeometry)
		FillUpForm(extrudegeometry, newFormGroup, extrudegeometryFormCallback.probe)
		extrudegeometryFormCallback.probe.formStage.Commit()
	}

	extrudegeometryFormCallback.probe.ux_tree()
}
func __gong__New__MeshFormCallback(
	mesh *models.Mesh,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshFormCallback *MeshFormCallback) {
	meshFormCallback = new(MeshFormCallback)
	meshFormCallback.probe = probe
	meshFormCallback.mesh = mesh
	meshFormCallback.formGroup = formGroup

	meshFormCallback.CreationMode = (mesh == nil)

	return
}

type MeshFormCallback struct {
	mesh *models.Mesh

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (meshFormCallback *MeshFormCallback) OnSave() {
	meshFormCallback.probe.stageOfInterest.Lock()
	defer meshFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MeshFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	meshFormCallback.probe.formStage.Checkout()

	if meshFormCallback.mesh == nil {
		meshFormCallback.mesh = new(models.Mesh).Stage(meshFormCallback.probe.stageOfInterest)
	}
	mesh_ := meshFormCallback.mesh
	_ = mesh_

	for _, formDiv := range meshFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mesh_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(mesh_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(mesh_.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(mesh_.Z), formDiv)
		case "MeshMaterialBasic":
			FormDivSelectFieldToField(&(mesh_.MeshMaterialBasic), meshFormCallback.probe.stageOfInterest, formDiv)
		case "MeshPhysicalMaterial":
			FormDivSelectFieldToField(&(mesh_.MeshPhysicalMaterial), meshFormCallback.probe.stageOfInterest, formDiv)
		case "CylinderGeometry":
			FormDivSelectFieldToField(&(mesh_.CylinderGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "BoxGeometry":
			FormDivSelectFieldToField(&(mesh_.BoxGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "SphereGeometry":
			FormDivSelectFieldToField(&(mesh_.SphereGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "TorusGeometry":
			FormDivSelectFieldToField(&(mesh_.TorusGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "PlaneGeometry":
			FormDivSelectFieldToField(&(mesh_.PlaneGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "TubeGeometry":
			FormDivSelectFieldToField(&(mesh_.TubeGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "ExtrudeGeometry":
			FormDivSelectFieldToField(&(mesh_.ExtrudeGeometry), meshFormCallback.probe.stageOfInterest, formDiv)
		case "Canvas:Meshs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Canvas instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Canvas instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Canvas](meshFormCallback.probe.stageOfInterest)
			targetCanvasIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCanvasIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Canvas instances and update their Meshs slice
			for _canvas := range *models.GetGongstructInstancesSetFromPointerType[*models.Canvas](meshFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(meshFormCallback.probe.stageOfInterest, _canvas)
				
				// if Canvas is selected
				if targetCanvasIDs[id] {
					// ensure mesh_ is in _canvas.Meshs
					found := false
					for _, _b := range _canvas.Meshs {
						if _b == mesh_ {
							found = true
							break
						}
					}
					if !found {
						_canvas.Meshs = append(_canvas.Meshs, mesh_)
						meshFormCallback.probe.UpdateSliceOfPointersCallback(_canvas, "Meshs", &_canvas.Meshs)
					}
				} else {
					// ensure mesh_ is NOT in _canvas.Meshs
					idx := slices.Index(_canvas.Meshs, mesh_)
					if idx != -1 {
						_canvas.Meshs = slices.Delete(_canvas.Meshs, idx, idx+1)
						meshFormCallback.probe.UpdateSliceOfPointersCallback(_canvas, "Meshs", &_canvas.Meshs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if meshFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mesh_.Unstage(meshFormCallback.probe.stageOfInterest)
	}

	meshFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Mesh](
		meshFormCallback.probe,
	)

	// display a new form by reset the form stage
	if meshFormCallback.CreationMode || meshFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meshFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(meshFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MeshFormCallback(
			nil,
			meshFormCallback.probe,
			newFormGroup,
		)
		mesh := new(models.Mesh)
		FillUpForm(mesh, newFormGroup, meshFormCallback.probe)
		meshFormCallback.probe.formStage.Commit()
	}

	meshFormCallback.probe.ux_tree()
}
func __gong__New__MeshMaterialBasicFormCallback(
	meshmaterialbasic *models.MeshMaterialBasic,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshmaterialbasicFormCallback *MeshMaterialBasicFormCallback) {
	meshmaterialbasicFormCallback = new(MeshMaterialBasicFormCallback)
	meshmaterialbasicFormCallback.probe = probe
	meshmaterialbasicFormCallback.meshmaterialbasic = meshmaterialbasic
	meshmaterialbasicFormCallback.formGroup = formGroup

	meshmaterialbasicFormCallback.CreationMode = (meshmaterialbasic == nil)

	return
}

type MeshMaterialBasicFormCallback struct {
	meshmaterialbasic *models.MeshMaterialBasic

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (meshmaterialbasicFormCallback *MeshMaterialBasicFormCallback) OnSave() {
	meshmaterialbasicFormCallback.probe.stageOfInterest.Lock()
	defer meshmaterialbasicFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MeshMaterialBasicFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	meshmaterialbasicFormCallback.probe.formStage.Checkout()

	if meshmaterialbasicFormCallback.meshmaterialbasic == nil {
		meshmaterialbasicFormCallback.meshmaterialbasic = new(models.MeshMaterialBasic).Stage(meshmaterialbasicFormCallback.probe.stageOfInterest)
	}
	meshmaterialbasic_ := meshmaterialbasicFormCallback.meshmaterialbasic
	_ = meshmaterialbasic_

	for _, formDiv := range meshmaterialbasicFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(meshmaterialbasic_.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(meshmaterialbasic_.Color), formDiv)
		}
	}

	// manage the suppress operation
	if meshmaterialbasicFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meshmaterialbasic_.Unstage(meshmaterialbasicFormCallback.probe.stageOfInterest)
	}

	meshmaterialbasicFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MeshMaterialBasic](
		meshmaterialbasicFormCallback.probe,
	)

	// display a new form by reset the form stage
	if meshmaterialbasicFormCallback.CreationMode || meshmaterialbasicFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meshmaterialbasicFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(meshmaterialbasicFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MeshMaterialBasicFormCallback(
			nil,
			meshmaterialbasicFormCallback.probe,
			newFormGroup,
		)
		meshmaterialbasic := new(models.MeshMaterialBasic)
		FillUpForm(meshmaterialbasic, newFormGroup, meshmaterialbasicFormCallback.probe)
		meshmaterialbasicFormCallback.probe.formStage.Commit()
	}

	meshmaterialbasicFormCallback.probe.ux_tree()
}
func __gong__New__MeshPhysicalMaterialFormCallback(
	meshphysicalmaterial *models.MeshPhysicalMaterial,
	probe *Probe,
	formGroup *form.FormGroup,
) (meshphysicalmaterialFormCallback *MeshPhysicalMaterialFormCallback) {
	meshphysicalmaterialFormCallback = new(MeshPhysicalMaterialFormCallback)
	meshphysicalmaterialFormCallback.probe = probe
	meshphysicalmaterialFormCallback.meshphysicalmaterial = meshphysicalmaterial
	meshphysicalmaterialFormCallback.formGroup = formGroup

	meshphysicalmaterialFormCallback.CreationMode = (meshphysicalmaterial == nil)

	return
}

type MeshPhysicalMaterialFormCallback struct {
	meshphysicalmaterial *models.MeshPhysicalMaterial

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (meshphysicalmaterialFormCallback *MeshPhysicalMaterialFormCallback) OnSave() {
	meshphysicalmaterialFormCallback.probe.stageOfInterest.Lock()
	defer meshphysicalmaterialFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MeshPhysicalMaterialFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	meshphysicalmaterialFormCallback.probe.formStage.Checkout()

	if meshphysicalmaterialFormCallback.meshphysicalmaterial == nil {
		meshphysicalmaterialFormCallback.meshphysicalmaterial = new(models.MeshPhysicalMaterial).Stage(meshphysicalmaterialFormCallback.probe.stageOfInterest)
	}
	meshphysicalmaterial_ := meshphysicalmaterialFormCallback.meshphysicalmaterial
	_ = meshphysicalmaterial_

	for _, formDiv := range meshphysicalmaterialFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Color), formDiv)
		case "Wireframe":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Wireframe), formDiv)
		case "Opacity":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Opacity), formDiv)
		case "Transparent":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Transparent), formDiv)
		case "Visible":
			FormDivBasicFieldToField(&(meshphysicalmaterial_.Visible), formDiv)
		}
	}

	// manage the suppress operation
	if meshphysicalmaterialFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meshphysicalmaterial_.Unstage(meshphysicalmaterialFormCallback.probe.stageOfInterest)
	}

	meshphysicalmaterialFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MeshPhysicalMaterial](
		meshphysicalmaterialFormCallback.probe,
	)

	// display a new form by reset the form stage
	if meshphysicalmaterialFormCallback.CreationMode || meshphysicalmaterialFormCallback.formGroup.HasSuppressButtonBeenPressed {
		meshphysicalmaterialFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(meshphysicalmaterialFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MeshPhysicalMaterialFormCallback(
			nil,
			meshphysicalmaterialFormCallback.probe,
			newFormGroup,
		)
		meshphysicalmaterial := new(models.MeshPhysicalMaterial)
		FillUpForm(meshphysicalmaterial, newFormGroup, meshphysicalmaterialFormCallback.probe)
		meshphysicalmaterialFormCallback.probe.formStage.Commit()
	}

	meshphysicalmaterialFormCallback.probe.ux_tree()
}
func __gong__New__PlaneGeometryFormCallback(
	planegeometry *models.PlaneGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (planegeometryFormCallback *PlaneGeometryFormCallback) {
	planegeometryFormCallback = new(PlaneGeometryFormCallback)
	planegeometryFormCallback.probe = probe
	planegeometryFormCallback.planegeometry = planegeometry
	planegeometryFormCallback.formGroup = formGroup

	planegeometryFormCallback.CreationMode = (planegeometry == nil)

	return
}

type PlaneGeometryFormCallback struct {
	planegeometry *models.PlaneGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (planegeometryFormCallback *PlaneGeometryFormCallback) OnSave() {
	planegeometryFormCallback.probe.stageOfInterest.Lock()
	defer planegeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PlaneGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	planegeometryFormCallback.probe.formStage.Checkout()

	if planegeometryFormCallback.planegeometry == nil {
		planegeometryFormCallback.planegeometry = new(models.PlaneGeometry).Stage(planegeometryFormCallback.probe.stageOfInterest)
	}
	planegeometry_ := planegeometryFormCallback.planegeometry
	_ = planegeometry_

	for _, formDiv := range planegeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(planegeometry_.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(planegeometry_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(planegeometry_.Height), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(planegeometry_.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(planegeometry_.HeightSegments), formDiv)
		}
	}

	// manage the suppress operation
	if planegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		planegeometry_.Unstage(planegeometryFormCallback.probe.stageOfInterest)
	}

	planegeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PlaneGeometry](
		planegeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if planegeometryFormCallback.CreationMode || planegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		planegeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(planegeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlaneGeometryFormCallback(
			nil,
			planegeometryFormCallback.probe,
			newFormGroup,
		)
		planegeometry := new(models.PlaneGeometry)
		FillUpForm(planegeometry, newFormGroup, planegeometryFormCallback.probe)
		planegeometryFormCallback.probe.formStage.Commit()
	}

	planegeometryFormCallback.probe.ux_tree()
}
func __gong__New__ShapeFormCallback(
	shape *models.Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shapeFormCallback *ShapeFormCallback) {
	shapeFormCallback = new(ShapeFormCallback)
	shapeFormCallback.probe = probe
	shapeFormCallback.shape = shape
	shapeFormCallback.formGroup = formGroup

	shapeFormCallback.CreationMode = (shape == nil)

	return
}

type ShapeFormCallback struct {
	shape *models.Shape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shapeFormCallback *ShapeFormCallback) OnSave() {
	shapeFormCallback.probe.stageOfInterest.Lock()
	defer shapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shapeFormCallback.probe.formStage.Checkout()

	if shapeFormCallback.shape == nil {
		shapeFormCallback.shape = new(models.Shape).Stage(shapeFormCallback.probe.stageOfInterest)
	}
	shape_ := shapeFormCallback.shape
	_ = shape_

	for _, formDiv := range shapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shape_.Name), formDiv)
		case "Points":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Vector2](shapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Vector2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Vector2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Vector2](shapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shape_.Points = instanceSlice
			shapeFormCallback.probe.UpdateSliceOfPointersCallback(shape_, "Points", &shape_.Points)

		}
	}

	// manage the suppress operation
	if shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shape_.Unstage(shapeFormCallback.probe.stageOfInterest)
	}

	shapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Shape](
		shapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shapeFormCallback.CreationMode || shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShapeFormCallback(
			nil,
			shapeFormCallback.probe,
			newFormGroup,
		)
		shape := new(models.Shape)
		FillUpForm(shape, newFormGroup, shapeFormCallback.probe)
		shapeFormCallback.probe.formStage.Commit()
	}

	shapeFormCallback.probe.ux_tree()
}
func __gong__New__SphereGeometryFormCallback(
	spheregeometry *models.SphereGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (spheregeometryFormCallback *SphereGeometryFormCallback) {
	spheregeometryFormCallback = new(SphereGeometryFormCallback)
	spheregeometryFormCallback.probe = probe
	spheregeometryFormCallback.spheregeometry = spheregeometry
	spheregeometryFormCallback.formGroup = formGroup

	spheregeometryFormCallback.CreationMode = (spheregeometry == nil)

	return
}

type SphereGeometryFormCallback struct {
	spheregeometry *models.SphereGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (spheregeometryFormCallback *SphereGeometryFormCallback) OnSave() {
	spheregeometryFormCallback.probe.stageOfInterest.Lock()
	defer spheregeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SphereGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spheregeometryFormCallback.probe.formStage.Checkout()

	if spheregeometryFormCallback.spheregeometry == nil {
		spheregeometryFormCallback.spheregeometry = new(models.SphereGeometry).Stage(spheregeometryFormCallback.probe.stageOfInterest)
	}
	spheregeometry_ := spheregeometryFormCallback.spheregeometry
	_ = spheregeometry_

	for _, formDiv := range spheregeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spheregeometry_.Name), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(spheregeometry_.Radius), formDiv)
		case "WidthSegments":
			FormDivBasicFieldToField(&(spheregeometry_.WidthSegments), formDiv)
		case "HeightSegments":
			FormDivBasicFieldToField(&(spheregeometry_.HeightSegments), formDiv)
		case "PhiStart":
			FormDivBasicFieldToField(&(spheregeometry_.PhiStart), formDiv)
		case "PhiLength":
			FormDivBasicFieldToField(&(spheregeometry_.PhiLength), formDiv)
		case "ThetaStart":
			FormDivBasicFieldToField(&(spheregeometry_.ThetaStart), formDiv)
		case "ThetaLength":
			FormDivBasicFieldToField(&(spheregeometry_.ThetaLength), formDiv)
		}
	}

	// manage the suppress operation
	if spheregeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spheregeometry_.Unstage(spheregeometryFormCallback.probe.stageOfInterest)
	}

	spheregeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SphereGeometry](
		spheregeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spheregeometryFormCallback.CreationMode || spheregeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spheregeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(spheregeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SphereGeometryFormCallback(
			nil,
			spheregeometryFormCallback.probe,
			newFormGroup,
		)
		spheregeometry := new(models.SphereGeometry)
		FillUpForm(spheregeometry, newFormGroup, spheregeometryFormCallback.probe)
		spheregeometryFormCallback.probe.formStage.Commit()
	}

	spheregeometryFormCallback.probe.ux_tree()
}
func __gong__New__TorusGeometryFormCallback(
	torusgeometry *models.TorusGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (torusgeometryFormCallback *TorusGeometryFormCallback) {
	torusgeometryFormCallback = new(TorusGeometryFormCallback)
	torusgeometryFormCallback.probe = probe
	torusgeometryFormCallback.torusgeometry = torusgeometry
	torusgeometryFormCallback.formGroup = formGroup

	torusgeometryFormCallback.CreationMode = (torusgeometry == nil)

	return
}

type TorusGeometryFormCallback struct {
	torusgeometry *models.TorusGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (torusgeometryFormCallback *TorusGeometryFormCallback) OnSave() {
	torusgeometryFormCallback.probe.stageOfInterest.Lock()
	defer torusgeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TorusGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	torusgeometryFormCallback.probe.formStage.Checkout()

	if torusgeometryFormCallback.torusgeometry == nil {
		torusgeometryFormCallback.torusgeometry = new(models.TorusGeometry).Stage(torusgeometryFormCallback.probe.stageOfInterest)
	}
	torusgeometry_ := torusgeometryFormCallback.torusgeometry
	_ = torusgeometry_

	for _, formDiv := range torusgeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(torusgeometry_.Name), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(torusgeometry_.Radius), formDiv)
		case "Tube":
			FormDivBasicFieldToField(&(torusgeometry_.Tube), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(torusgeometry_.RadialSegments), formDiv)
		case "TubularSegments":
			FormDivBasicFieldToField(&(torusgeometry_.TubularSegments), formDiv)
		case "Arc":
			FormDivBasicFieldToField(&(torusgeometry_.Arc), formDiv)
		}
	}

	// manage the suppress operation
	if torusgeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		torusgeometry_.Unstage(torusgeometryFormCallback.probe.stageOfInterest)
	}

	torusgeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TorusGeometry](
		torusgeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if torusgeometryFormCallback.CreationMode || torusgeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		torusgeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(torusgeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TorusGeometryFormCallback(
			nil,
			torusgeometryFormCallback.probe,
			newFormGroup,
		)
		torusgeometry := new(models.TorusGeometry)
		FillUpForm(torusgeometry, newFormGroup, torusgeometryFormCallback.probe)
		torusgeometryFormCallback.probe.formStage.Commit()
	}

	torusgeometryFormCallback.probe.ux_tree()
}
func __gong__New__TubeGeometryFormCallback(
	tubegeometry *models.TubeGeometry,
	probe *Probe,
	formGroup *form.FormGroup,
) (tubegeometryFormCallback *TubeGeometryFormCallback) {
	tubegeometryFormCallback = new(TubeGeometryFormCallback)
	tubegeometryFormCallback.probe = probe
	tubegeometryFormCallback.tubegeometry = tubegeometry
	tubegeometryFormCallback.formGroup = formGroup

	tubegeometryFormCallback.CreationMode = (tubegeometry == nil)

	return
}

type TubeGeometryFormCallback struct {
	tubegeometry *models.TubeGeometry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (tubegeometryFormCallback *TubeGeometryFormCallback) OnSave() {
	tubegeometryFormCallback.probe.stageOfInterest.Lock()
	defer tubegeometryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TubeGeometryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tubegeometryFormCallback.probe.formStage.Checkout()

	if tubegeometryFormCallback.tubegeometry == nil {
		tubegeometryFormCallback.tubegeometry = new(models.TubeGeometry).Stage(tubegeometryFormCallback.probe.stageOfInterest)
	}
	tubegeometry_ := tubegeometryFormCallback.tubegeometry
	_ = tubegeometry_

	for _, formDiv := range tubegeometryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tubegeometry_.Name), formDiv)
		case "Path":
			FormDivSelectFieldToField(&(tubegeometry_.Path), tubegeometryFormCallback.probe.stageOfInterest, formDiv)
		case "TubularSegments":
			FormDivBasicFieldToField(&(tubegeometry_.TubularSegments), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(tubegeometry_.Radius), formDiv)
		case "RadialSegments":
			FormDivBasicFieldToField(&(tubegeometry_.RadialSegments), formDiv)
		case "Closed":
			FormDivBasicFieldToField(&(tubegeometry_.Closed), formDiv)
		}
	}

	// manage the suppress operation
	if tubegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tubegeometry_.Unstage(tubegeometryFormCallback.probe.stageOfInterest)
	}

	tubegeometryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TubeGeometry](
		tubegeometryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tubegeometryFormCallback.CreationMode || tubegeometryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tubegeometryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(tubegeometryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TubeGeometryFormCallback(
			nil,
			tubegeometryFormCallback.probe,
			newFormGroup,
		)
		tubegeometry := new(models.TubeGeometry)
		FillUpForm(tubegeometry, newFormGroup, tubegeometryFormCallback.probe)
		tubegeometryFormCallback.probe.formStage.Commit()
	}

	tubegeometryFormCallback.probe.ux_tree()
}
func __gong__New__Vector2FormCallback(
	vector2 *models.Vector2,
	probe *Probe,
	formGroup *form.FormGroup,
) (vector2FormCallback *Vector2FormCallback) {
	vector2FormCallback = new(Vector2FormCallback)
	vector2FormCallback.probe = probe
	vector2FormCallback.vector2 = vector2
	vector2FormCallback.formGroup = formGroup

	vector2FormCallback.CreationMode = (vector2 == nil)

	return
}

type Vector2FormCallback struct {
	vector2 *models.Vector2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (vector2FormCallback *Vector2FormCallback) OnSave() {
	vector2FormCallback.probe.stageOfInterest.Lock()
	defer vector2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("Vector2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	vector2FormCallback.probe.formStage.Checkout()

	if vector2FormCallback.vector2 == nil {
		vector2FormCallback.vector2 = new(models.Vector2).Stage(vector2FormCallback.probe.stageOfInterest)
	}
	vector2_ := vector2FormCallback.vector2
	_ = vector2_

	for _, formDiv := range vector2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(vector2_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(vector2_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(vector2_.Y), formDiv)
		case "Shape:Points":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Shape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Shape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Shape](vector2FormCallback.probe.stageOfInterest)
			targetShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Shape instances and update their Points slice
			for _shape := range *models.GetGongstructInstancesSetFromPointerType[*models.Shape](vector2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(vector2FormCallback.probe.stageOfInterest, _shape)
				
				// if Shape is selected
				if targetShapeIDs[id] {
					// ensure vector2_ is in _shape.Points
					found := false
					for _, _b := range _shape.Points {
						if _b == vector2_ {
							found = true
							break
						}
					}
					if !found {
						_shape.Points = append(_shape.Points, vector2_)
						vector2FormCallback.probe.UpdateSliceOfPointersCallback(_shape, "Points", &_shape.Points)
					}
				} else {
					// ensure vector2_ is NOT in _shape.Points
					idx := slices.Index(_shape.Points, vector2_)
					if idx != -1 {
						_shape.Points = slices.Delete(_shape.Points, idx, idx+1)
						vector2FormCallback.probe.UpdateSliceOfPointersCallback(_shape, "Points", &_shape.Points)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if vector2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector2_.Unstage(vector2FormCallback.probe.stageOfInterest)
	}

	vector2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Vector2](
		vector2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if vector2FormCallback.CreationMode || vector2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(vector2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Vector2FormCallback(
			nil,
			vector2FormCallback.probe,
			newFormGroup,
		)
		vector2 := new(models.Vector2)
		FillUpForm(vector2, newFormGroup, vector2FormCallback.probe)
		vector2FormCallback.probe.formStage.Commit()
	}

	vector2FormCallback.probe.ux_tree()
}
func __gong__New__Vector3FormCallback(
	vector3 *models.Vector3,
	probe *Probe,
	formGroup *form.FormGroup,
) (vector3FormCallback *Vector3FormCallback) {
	vector3FormCallback = new(Vector3FormCallback)
	vector3FormCallback.probe = probe
	vector3FormCallback.vector3 = vector3
	vector3FormCallback.formGroup = formGroup

	vector3FormCallback.CreationMode = (vector3 == nil)

	return
}

type Vector3FormCallback struct {
	vector3 *models.Vector3

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (vector3FormCallback *Vector3FormCallback) OnSave() {
	vector3FormCallback.probe.stageOfInterest.Lock()
	defer vector3FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("Vector3FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	vector3FormCallback.probe.formStage.Checkout()

	if vector3FormCallback.vector3 == nil {
		vector3FormCallback.vector3 = new(models.Vector3).Stage(vector3FormCallback.probe.stageOfInterest)
	}
	vector3_ := vector3FormCallback.vector3
	_ = vector3_

	for _, formDiv := range vector3FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(vector3_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(vector3_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(vector3_.Y), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(vector3_.Z), formDiv)
		case "Curve:Points":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Curve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Curve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Curve](vector3FormCallback.probe.stageOfInterest)
			targetCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Curve instances and update their Points slice
			for _curve := range *models.GetGongstructInstancesSetFromPointerType[*models.Curve](vector3FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(vector3FormCallback.probe.stageOfInterest, _curve)
				
				// if Curve is selected
				if targetCurveIDs[id] {
					// ensure vector3_ is in _curve.Points
					found := false
					for _, _b := range _curve.Points {
						if _b == vector3_ {
							found = true
							break
						}
					}
					if !found {
						_curve.Points = append(_curve.Points, vector3_)
						vector3FormCallback.probe.UpdateSliceOfPointersCallback(_curve, "Points", &_curve.Points)
					}
				} else {
					// ensure vector3_ is NOT in _curve.Points
					idx := slices.Index(_curve.Points, vector3_)
					if idx != -1 {
						_curve.Points = slices.Delete(_curve.Points, idx, idx+1)
						vector3FormCallback.probe.UpdateSliceOfPointersCallback(_curve, "Points", &_curve.Points)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if vector3FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector3_.Unstage(vector3FormCallback.probe.stageOfInterest)
	}

	vector3FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Vector3](
		vector3FormCallback.probe,
	)

	// display a new form by reset the form stage
	if vector3FormCallback.CreationMode || vector3FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector3FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(vector3FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Vector3FormCallback(
			nil,
			vector3FormCallback.probe,
			newFormGroup,
		)
		vector3 := new(models.Vector3)
		FillUpForm(vector3, newFormGroup, vector3FormCallback.probe)
		vector3FormCallback.probe.formStage.Commit()
	}

	vector3FormCallback.probe.ux_tree()
}
