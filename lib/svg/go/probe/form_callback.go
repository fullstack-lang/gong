// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AnimateFormCallback(
	animate *models.Animate,
	probe *Probe,
	formGroup *form.FormGroup,
) (animateFormCallback *AnimateFormCallback) {
	animateFormCallback = new(AnimateFormCallback)
	animateFormCallback.probe = probe
	animateFormCallback.animate = animate
	animateFormCallback.formGroup = formGroup

	animateFormCallback.CreationMode = (animate == nil)

	return
}

type AnimateFormCallback struct {
	animate *models.Animate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (animateFormCallback *AnimateFormCallback) OnSave() {
	animateFormCallback.probe.stageOfInterest.Lock()
	defer animateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AnimateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	animateFormCallback.probe.formStage.Checkout()

	if animateFormCallback.animate == nil {
		animateFormCallback.animate = new(models.Animate).Stage(animateFormCallback.probe.stageOfInterest)
	}
	animate_ := animateFormCallback.animate
	_ = animate_

	for _, formDiv := range animateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(animate_.Name), formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&(animate_.AttributeName), formDiv)
		case "Values":
			FormDivBasicFieldToField(&(animate_.Values), formDiv)
		case "From":
			FormDivBasicFieldToField(&(animate_.From), formDiv)
		case "To":
			FormDivBasicFieldToField(&(animate_.To), formDiv)
		case "Dur":
			FormDivBasicFieldToField(&(animate_.Dur), formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&(animate_.RepeatCount), formDiv)
		case "Circle:Animations":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Circle instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Circle instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Circle](animateFormCallback.probe.stageOfInterest)
			targetCircleIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCircleIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Circle instances and update their Animations slice
			for _circle := range *models.GetGongstructInstancesSetFromPointerType[*models.Circle](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _circle)
				
				// if Circle is selected
				if targetCircleIDs[id] {
					// ensure animate_ is in _circle.Animations
					found := false
					for _, _b := range _circle.Animations {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_circle.Animations = append(_circle.Animations, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_circle, "Animations", &_circle.Animations)
					}
				} else {
					// ensure animate_ is NOT in _circle.Animations
					idx := slices.Index(_circle.Animations, animate_)
					if idx != -1 {
						_circle.Animations = slices.Delete(_circle.Animations, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_circle, "Animations", &_circle.Animations)
					}
				}
			}
		case "Ellipse:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Ellipse instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Ellipse instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Ellipse](animateFormCallback.probe.stageOfInterest)
			targetEllipseIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetEllipseIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Ellipse instances and update their Animates slice
			for _ellipse := range *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _ellipse)
				
				// if Ellipse is selected
				if targetEllipseIDs[id] {
					// ensure animate_ is in _ellipse.Animates
					found := false
					for _, _b := range _ellipse.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_ellipse.Animates = append(_ellipse.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_ellipse, "Animates", &_ellipse.Animates)
					}
				} else {
					// ensure animate_ is NOT in _ellipse.Animates
					idx := slices.Index(_ellipse.Animates, animate_)
					if idx != -1 {
						_ellipse.Animates = slices.Delete(_ellipse.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_ellipse, "Animates", &_ellipse.Animates)
					}
				}
			}
		case "Line:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Line instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Line instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Line](animateFormCallback.probe.stageOfInterest)
			targetLineIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLineIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Line instances and update their Animates slice
			for _line := range *models.GetGongstructInstancesSetFromPointerType[*models.Line](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _line)
				
				// if Line is selected
				if targetLineIDs[id] {
					// ensure animate_ is in _line.Animates
					found := false
					for _, _b := range _line.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_line.Animates = append(_line.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_line, "Animates", &_line.Animates)
					}
				} else {
					// ensure animate_ is NOT in _line.Animates
					idx := slices.Index(_line.Animates, animate_)
					if idx != -1 {
						_line.Animates = slices.Delete(_line.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_line, "Animates", &_line.Animates)
					}
				}
			}
		case "LinkAnchoredText:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the LinkAnchoredText instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target LinkAnchoredText instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredText](animateFormCallback.probe.stageOfInterest)
			targetLinkAnchoredTextIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkAnchoredTextIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all LinkAnchoredText instances and update their Animates slice
			for _linkanchoredtext := range *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _linkanchoredtext)
				
				// if LinkAnchoredText is selected
				if targetLinkAnchoredTextIDs[id] {
					// ensure animate_ is in _linkanchoredtext.Animates
					found := false
					for _, _b := range _linkanchoredtext.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_linkanchoredtext.Animates = append(_linkanchoredtext.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_linkanchoredtext, "Animates", &_linkanchoredtext.Animates)
					}
				} else {
					// ensure animate_ is NOT in _linkanchoredtext.Animates
					idx := slices.Index(_linkanchoredtext.Animates, animate_)
					if idx != -1 {
						_linkanchoredtext.Animates = slices.Delete(_linkanchoredtext.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_linkanchoredtext, "Animates", &_linkanchoredtext.Animates)
					}
				}
			}
		case "Path:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Path instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Path instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Path](animateFormCallback.probe.stageOfInterest)
			targetPathIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPathIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Path instances and update their Animates slice
			for _path := range *models.GetGongstructInstancesSetFromPointerType[*models.Path](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _path)
				
				// if Path is selected
				if targetPathIDs[id] {
					// ensure animate_ is in _path.Animates
					found := false
					for _, _b := range _path.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_path.Animates = append(_path.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_path, "Animates", &_path.Animates)
					}
				} else {
					// ensure animate_ is NOT in _path.Animates
					idx := slices.Index(_path.Animates, animate_)
					if idx != -1 {
						_path.Animates = slices.Delete(_path.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_path, "Animates", &_path.Animates)
					}
				}
			}
		case "Polygone:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Polygone instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Polygone instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Polygone](animateFormCallback.probe.stageOfInterest)
			targetPolygoneIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPolygoneIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Polygone instances and update their Animates slice
			for _polygone := range *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _polygone)
				
				// if Polygone is selected
				if targetPolygoneIDs[id] {
					// ensure animate_ is in _polygone.Animates
					found := false
					for _, _b := range _polygone.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_polygone.Animates = append(_polygone.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_polygone, "Animates", &_polygone.Animates)
					}
				} else {
					// ensure animate_ is NOT in _polygone.Animates
					idx := slices.Index(_polygone.Animates, animate_)
					if idx != -1 {
						_polygone.Animates = slices.Delete(_polygone.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_polygone, "Animates", &_polygone.Animates)
					}
				}
			}
		case "Polyline:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Polyline instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Polyline instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Polyline](animateFormCallback.probe.stageOfInterest)
			targetPolylineIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPolylineIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Polyline instances and update their Animates slice
			for _polyline := range *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _polyline)
				
				// if Polyline is selected
				if targetPolylineIDs[id] {
					// ensure animate_ is in _polyline.Animates
					found := false
					for _, _b := range _polyline.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_polyline.Animates = append(_polyline.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_polyline, "Animates", &_polyline.Animates)
					}
				} else {
					// ensure animate_ is NOT in _polyline.Animates
					idx := slices.Index(_polyline.Animates, animate_)
					if idx != -1 {
						_polyline.Animates = slices.Delete(_polyline.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_polyline, "Animates", &_polyline.Animates)
					}
				}
			}
		case "Rect:Animations":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](animateFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their Animations slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure animate_ is in _rect.Animations
					found := false
					for _, _b := range _rect.Animations {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_rect.Animations = append(_rect.Animations, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Animations", &_rect.Animations)
					}
				} else {
					// ensure animate_ is NOT in _rect.Animations
					idx := slices.Index(_rect.Animations, animate_)
					if idx != -1 {
						_rect.Animations = slices.Delete(_rect.Animations, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Animations", &_rect.Animations)
					}
				}
			}
		case "RectAnchoredText:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the RectAnchoredText instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target RectAnchoredText instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.RectAnchoredText](animateFormCallback.probe.stageOfInterest)
			targetRectAnchoredTextIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectAnchoredTextIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all RectAnchoredText instances and update their Animates slice
			for _rectanchoredtext := range *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _rectanchoredtext)
				
				// if RectAnchoredText is selected
				if targetRectAnchoredTextIDs[id] {
					// ensure animate_ is in _rectanchoredtext.Animates
					found := false
					for _, _b := range _rectanchoredtext.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_rectanchoredtext.Animates = append(_rectanchoredtext.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_rectanchoredtext, "Animates", &_rectanchoredtext.Animates)
					}
				} else {
					// ensure animate_ is NOT in _rectanchoredtext.Animates
					idx := slices.Index(_rectanchoredtext.Animates, animate_)
					if idx != -1 {
						_rectanchoredtext.Animates = slices.Delete(_rectanchoredtext.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_rectanchoredtext, "Animates", &_rectanchoredtext.Animates)
					}
				}
			}
		case "Text:Animates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Text instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Text instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Text](animateFormCallback.probe.stageOfInterest)
			targetTextIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTextIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Text instances and update their Animates slice
			for _text := range *models.GetGongstructInstancesSetFromPointerType[*models.Text](animateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(animateFormCallback.probe.stageOfInterest, _text)
				
				// if Text is selected
				if targetTextIDs[id] {
					// ensure animate_ is in _text.Animates
					found := false
					for _, _b := range _text.Animates {
						if _b == animate_ {
							found = true
							break
						}
					}
					if !found {
						_text.Animates = append(_text.Animates, animate_)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_text, "Animates", &_text.Animates)
					}
				} else {
					// ensure animate_ is NOT in _text.Animates
					idx := slices.Index(_text.Animates, animate_)
					if idx != -1 {
						_text.Animates = slices.Delete(_text.Animates, idx, idx+1)
						animateFormCallback.probe.UpdateSliceOfPointersCallback(_text, "Animates", &_text.Animates)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if animateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		animate_.Unstage(animateFormCallback.probe.stageOfInterest)
	}

	animateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Animate](
		animateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if animateFormCallback.CreationMode || animateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		animateFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(animateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnimateFormCallback(
			nil,
			animateFormCallback.probe,
			newFormGroup,
		)
		animate := new(models.Animate)
		FillUpForm(animate, newFormGroup, animateFormCallback.probe)
		animateFormCallback.probe.formStage.Commit()
	}

	animateFormCallback.probe.ux_tree()
}
func __gong__New__CircleFormCallback(
	circle *models.Circle,
	probe *Probe,
	formGroup *form.FormGroup,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.probe = probe
	circleFormCallback.circle = circle
	circleFormCallback.formGroup = formGroup

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (circleFormCallback *CircleFormCallback) OnSave() {
	circleFormCallback.probe.stageOfInterest.Lock()
	defer circleFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.probe.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.probe.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	for _, formDiv := range circleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(circle_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(circle_.CY), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(circle_.Radius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(circle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(circle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(circle_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(circle_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(circle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(circle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(circle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(circle_.Transform), formDiv)
		case "Animations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](circleFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					circleFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](circleFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			circle_.Animations = instanceSlice
			circleFormCallback.probe.UpdateSliceOfPointersCallback(circle_, "Animations", &circle_.Animations)

		case "Layer:Circles":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](circleFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Circles slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](circleFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(circleFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure circle_ is in _layer.Circles
					found := false
					for _, _b := range _layer.Circles {
						if _b == circle_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Circles = append(_layer.Circles, circle_)
						circleFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Circles", &_layer.Circles)
					}
				} else {
					// ensure circle_ is NOT in _layer.Circles
					idx := slices.Index(_layer.Circles, circle_)
					if idx != -1 {
						_layer.Circles = slices.Delete(_layer.Circles, idx, idx+1)
						circleFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Circles", &_layer.Circles)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circle_.Unstage(circleFormCallback.probe.stageOfInterest)
	}

	circleFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Circle](
		circleFormCallback.probe,
	)

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode || circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(circleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			circleFormCallback.probe,
			newFormGroup,
		)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.probe)
		circleFormCallback.probe.formStage.Commit()
	}

	circleFormCallback.probe.ux_tree()
}
func __gong__New__ConditionFormCallback(
	condition *models.Condition,
	probe *Probe,
	formGroup *form.FormGroup,
) (conditionFormCallback *ConditionFormCallback) {
	conditionFormCallback = new(ConditionFormCallback)
	conditionFormCallback.probe = probe
	conditionFormCallback.condition = condition
	conditionFormCallback.formGroup = formGroup

	conditionFormCallback.CreationMode = (condition == nil)

	return
}

type ConditionFormCallback struct {
	condition *models.Condition

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (conditionFormCallback *ConditionFormCallback) OnSave() {
	conditionFormCallback.probe.stageOfInterest.Lock()
	defer conditionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConditionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	conditionFormCallback.probe.formStage.Checkout()

	if conditionFormCallback.condition == nil {
		conditionFormCallback.condition = new(models.Condition).Stage(conditionFormCallback.probe.stageOfInterest)
	}
	condition_ := conditionFormCallback.condition
	_ = condition_

	for _, formDiv := range conditionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(condition_.Name), formDiv)
		case "Rect:HoveringTrigger":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](conditionFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their HoveringTrigger slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](conditionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conditionFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure condition_ is in _rect.HoveringTrigger
					found := false
					for _, _b := range _rect.HoveringTrigger {
						if _b == condition_ {
							found = true
							break
						}
					}
					if !found {
						_rect.HoveringTrigger = append(_rect.HoveringTrigger, condition_)
						conditionFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "HoveringTrigger", &_rect.HoveringTrigger)
					}
				} else {
					// ensure condition_ is NOT in _rect.HoveringTrigger
					idx := slices.Index(_rect.HoveringTrigger, condition_)
					if idx != -1 {
						_rect.HoveringTrigger = slices.Delete(_rect.HoveringTrigger, idx, idx+1)
						conditionFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "HoveringTrigger", &_rect.HoveringTrigger)
					}
				}
			}
		case "Rect:DisplayConditions":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](conditionFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their DisplayConditions slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](conditionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conditionFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure condition_ is in _rect.DisplayConditions
					found := false
					for _, _b := range _rect.DisplayConditions {
						if _b == condition_ {
							found = true
							break
						}
					}
					if !found {
						_rect.DisplayConditions = append(_rect.DisplayConditions, condition_)
						conditionFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "DisplayConditions", &_rect.DisplayConditions)
					}
				} else {
					// ensure condition_ is NOT in _rect.DisplayConditions
					idx := slices.Index(_rect.DisplayConditions, condition_)
					if idx != -1 {
						_rect.DisplayConditions = slices.Delete(_rect.DisplayConditions, idx, idx+1)
						conditionFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "DisplayConditions", &_rect.DisplayConditions)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if conditionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		condition_.Unstage(conditionFormCallback.probe.stageOfInterest)
	}

	conditionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Condition](
		conditionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if conditionFormCallback.CreationMode || conditionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		conditionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(conditionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConditionFormCallback(
			nil,
			conditionFormCallback.probe,
			newFormGroup,
		)
		condition := new(models.Condition)
		FillUpForm(condition, newFormGroup, conditionFormCallback.probe)
		conditionFormCallback.probe.formStage.Commit()
	}

	conditionFormCallback.probe.ux_tree()
}
func __gong__New__ControlPointFormCallback(
	controlpoint *models.ControlPoint,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointFormCallback *ControlPointFormCallback) {
	controlpointFormCallback = new(ControlPointFormCallback)
	controlpointFormCallback.probe = probe
	controlpointFormCallback.controlpoint = controlpoint
	controlpointFormCallback.formGroup = formGroup

	controlpointFormCallback.CreationMode = (controlpoint == nil)

	return
}

type ControlPointFormCallback struct {
	controlpoint *models.ControlPoint

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (controlpointFormCallback *ControlPointFormCallback) OnSave() {
	controlpointFormCallback.probe.stageOfInterest.Lock()
	defer controlpointFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ControlPointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlpointFormCallback.probe.formStage.Checkout()

	if controlpointFormCallback.controlpoint == nil {
		controlpointFormCallback.controlpoint = new(models.ControlPoint).Stage(controlpointFormCallback.probe.stageOfInterest)
	}
	controlpoint_ := controlpointFormCallback.controlpoint
	_ = controlpoint_

	for _, formDiv := range controlpointFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlpoint_.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(controlpoint_.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(controlpoint_.Y_Relative), formDiv)
		case "ClosestRect":
			FormDivSelectFieldToField(&(controlpoint_.ClosestRect), controlpointFormCallback.probe.stageOfInterest, formDiv)
		case "Link:ControlPoints":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](controlpointFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their ControlPoints slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](controlpointFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure controlpoint_ is in _link.ControlPoints
					found := false
					for _, _b := range _link.ControlPoints {
						if _b == controlpoint_ {
							found = true
							break
						}
					}
					if !found {
						_link.ControlPoints = append(_link.ControlPoints, controlpoint_)
						controlpointFormCallback.probe.UpdateSliceOfPointersCallback(_link, "ControlPoints", &_link.ControlPoints)
					}
				} else {
					// ensure controlpoint_ is NOT in _link.ControlPoints
					idx := slices.Index(_link.ControlPoints, controlpoint_)
					if idx != -1 {
						_link.ControlPoints = slices.Delete(_link.ControlPoints, idx, idx+1)
						controlpointFormCallback.probe.UpdateSliceOfPointersCallback(_link, "ControlPoints", &_link.ControlPoints)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if controlpointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpoint_.Unstage(controlpointFormCallback.probe.stageOfInterest)
	}

	controlpointFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlPoint](
		controlpointFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlpointFormCallback.CreationMode || controlpointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(controlpointFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlPointFormCallback(
			nil,
			controlpointFormCallback.probe,
			newFormGroup,
		)
		controlpoint := new(models.ControlPoint)
		FillUpForm(controlpoint, newFormGroup, controlpointFormCallback.probe)
		controlpointFormCallback.probe.formStage.Commit()
	}

	controlpointFormCallback.probe.ux_tree()
}
func __gong__New__EllipseFormCallback(
	ellipse *models.Ellipse,
	probe *Probe,
	formGroup *form.FormGroup,
) (ellipseFormCallback *EllipseFormCallback) {
	ellipseFormCallback = new(EllipseFormCallback)
	ellipseFormCallback.probe = probe
	ellipseFormCallback.ellipse = ellipse
	ellipseFormCallback.formGroup = formGroup

	ellipseFormCallback.CreationMode = (ellipse == nil)

	return
}

type EllipseFormCallback struct {
	ellipse *models.Ellipse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (ellipseFormCallback *EllipseFormCallback) OnSave() {
	ellipseFormCallback.probe.stageOfInterest.Lock()
	defer ellipseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EllipseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ellipseFormCallback.probe.formStage.Checkout()

	if ellipseFormCallback.ellipse == nil {
		ellipseFormCallback.ellipse = new(models.Ellipse).Stage(ellipseFormCallback.probe.stageOfInterest)
	}
	ellipse_ := ellipseFormCallback.ellipse
	_ = ellipse_

	for _, formDiv := range ellipseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ellipse_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(ellipse_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(ellipse_.CY), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(ellipse_.RX), formDiv)
		case "RY":
			FormDivBasicFieldToField(&(ellipse_.RY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(ellipse_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(ellipse_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(ellipse_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(ellipse_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(ellipse_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(ellipse_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](ellipseFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ellipseFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](ellipseFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			ellipse_.Animates = instanceSlice
			ellipseFormCallback.probe.UpdateSliceOfPointersCallback(ellipse_, "Animates", &ellipse_.Animates)

		case "Layer:Ellipses":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](ellipseFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Ellipses slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](ellipseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(ellipseFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure ellipse_ is in _layer.Ellipses
					found := false
					for _, _b := range _layer.Ellipses {
						if _b == ellipse_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Ellipses = append(_layer.Ellipses, ellipse_)
						ellipseFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Ellipses", &_layer.Ellipses)
					}
				} else {
					// ensure ellipse_ is NOT in _layer.Ellipses
					idx := slices.Index(_layer.Ellipses, ellipse_)
					if idx != -1 {
						_layer.Ellipses = slices.Delete(_layer.Ellipses, idx, idx+1)
						ellipseFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Ellipses", &_layer.Ellipses)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if ellipseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ellipse_.Unstage(ellipseFormCallback.probe.stageOfInterest)
	}

	ellipseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Ellipse](
		ellipseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if ellipseFormCallback.CreationMode || ellipseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ellipseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(ellipseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EllipseFormCallback(
			nil,
			ellipseFormCallback.probe,
			newFormGroup,
		)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, newFormGroup, ellipseFormCallback.probe)
		ellipseFormCallback.probe.formStage.Commit()
	}

	ellipseFormCallback.probe.ux_tree()
}
func __gong__New__FileToDownloadFormCallback(
	filetodownload *models.FileToDownload,
	probe *Probe,
	formGroup *form.FormGroup,
) (filetodownloadFormCallback *FileToDownloadFormCallback) {
	filetodownloadFormCallback = new(FileToDownloadFormCallback)
	filetodownloadFormCallback.probe = probe
	filetodownloadFormCallback.filetodownload = filetodownload
	filetodownloadFormCallback.formGroup = formGroup

	filetodownloadFormCallback.CreationMode = (filetodownload == nil)

	return
}

type FileToDownloadFormCallback struct {
	filetodownload *models.FileToDownload

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (filetodownloadFormCallback *FileToDownloadFormCallback) OnSave() {
	filetodownloadFormCallback.probe.stageOfInterest.Lock()
	defer filetodownloadFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FileToDownloadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	filetodownloadFormCallback.probe.formStage.Checkout()

	if filetodownloadFormCallback.filetodownload == nil {
		filetodownloadFormCallback.filetodownload = new(models.FileToDownload).Stage(filetodownloadFormCallback.probe.stageOfInterest)
	}
	filetodownload_ := filetodownloadFormCallback.filetodownload
	_ = filetodownload_

	for _, formDiv := range filetodownloadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(filetodownload_.Name), formDiv)
		case "Base64EncodedContent":
			FormDivBasicFieldToField(&(filetodownload_.Base64EncodedContent), formDiv)
		}
	}

	// manage the suppress operation
	if filetodownloadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetodownload_.Unstage(filetodownloadFormCallback.probe.stageOfInterest)
	}

	filetodownloadFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FileToDownload](
		filetodownloadFormCallback.probe,
	)

	// display a new form by reset the form stage
	if filetodownloadFormCallback.CreationMode || filetodownloadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetodownloadFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(filetodownloadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FileToDownloadFormCallback(
			nil,
			filetodownloadFormCallback.probe,
			newFormGroup,
		)
		filetodownload := new(models.FileToDownload)
		FillUpForm(filetodownload, newFormGroup, filetodownloadFormCallback.probe)
		filetodownloadFormCallback.probe.formStage.Commit()
	}

	filetodownloadFormCallback.probe.ux_tree()
}
func __gong__New__LayerFormCallback(
	layer *models.Layer,
	probe *Probe,
	formGroup *form.FormGroup,
) (layerFormCallback *LayerFormCallback) {
	layerFormCallback = new(LayerFormCallback)
	layerFormCallback.probe = probe
	layerFormCallback.layer = layer
	layerFormCallback.formGroup = formGroup

	layerFormCallback.CreationMode = (layer == nil)

	return
}

type LayerFormCallback struct {
	layer *models.Layer

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (layerFormCallback *LayerFormCallback) OnSave() {
	layerFormCallback.probe.stageOfInterest.Lock()
	defer layerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layerFormCallback.probe.formStage.Checkout()

	if layerFormCallback.layer == nil {
		layerFormCallback.layer = new(models.Layer).Stage(layerFormCallback.probe.stageOfInterest)
	}
	layer_ := layerFormCallback.layer
	_ = layer_

	for _, formDiv := range layerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layer_.Name), formDiv)
		case "Rects":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Rect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Rect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Rects = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Rects", &layer_.Rects)

		case "Texts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Text](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Text, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Text)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Text](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Texts = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Texts", &layer_.Texts)

		case "Circles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Circle](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Circle, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Circle)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Circle](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Circles = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Circles", &layer_.Circles)

		case "Lines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Line](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Line, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Line)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Line](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Lines = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Lines", &layer_.Lines)

		case "Ellipses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Ellipse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Ellipse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Ellipse](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Ellipses = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Ellipses", &layer_.Ellipses)

		case "Polylines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Polyline, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Polyline)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Polyline](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Polylines = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Polylines", &layer_.Polylines)

		case "Polygones":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Polygone, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Polygone)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Polygone](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Polygones = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Polygones", &layer_.Polygones)

		case "Paths":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Path](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Path, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Path)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Path](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Paths = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Paths", &layer_.Paths)

		case "Links":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Link](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.Links = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "Links", &layer_.Links)

		case "RectLinkLinks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectLinkLink](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectLinkLink, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectLinkLink)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RectLinkLink](layerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			layer_.RectLinkLinks = instanceSlice
			layerFormCallback.probe.UpdateSliceOfPointersCallback(layer_, "RectLinkLinks", &layer_.RectLinkLinks)

		case "SVG:Layers":
			// 1. Decode the AssociationStorage which contains the rowIDs of the SVG instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target SVG instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.SVG](layerFormCallback.probe.stageOfInterest)
			targetSVGIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSVGIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all SVG instances and update their Layers slice
			for _svg := range *models.GetGongstructInstancesSetFromPointerType[*models.SVG](layerFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(layerFormCallback.probe.stageOfInterest, _svg)
				
				// if SVG is selected
				if targetSVGIDs[id] {
					// ensure layer_ is in _svg.Layers
					found := false
					for _, _b := range _svg.Layers {
						if _b == layer_ {
							found = true
							break
						}
					}
					if !found {
						_svg.Layers = append(_svg.Layers, layer_)
						layerFormCallback.probe.UpdateSliceOfPointersCallback(_svg, "Layers", &_svg.Layers)
					}
				} else {
					// ensure layer_ is NOT in _svg.Layers
					idx := slices.Index(_svg.Layers, layer_)
					if idx != -1 {
						_svg.Layers = slices.Delete(_svg.Layers, idx, idx+1)
						layerFormCallback.probe.UpdateSliceOfPointersCallback(_svg, "Layers", &_svg.Layers)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if layerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layer_.Unstage(layerFormCallback.probe.stageOfInterest)
	}

	layerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Layer](
		layerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if layerFormCallback.CreationMode || layerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layerFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(layerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LayerFormCallback(
			nil,
			layerFormCallback.probe,
			newFormGroup,
		)
		layer := new(models.Layer)
		FillUpForm(layer, newFormGroup, layerFormCallback.probe)
		layerFormCallback.probe.formStage.Commit()
	}

	layerFormCallback.probe.ux_tree()
}
func __gong__New__LineFormCallback(
	line *models.Line,
	probe *Probe,
	formGroup *form.FormGroup,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.probe = probe
	lineFormCallback.line = line
	lineFormCallback.formGroup = formGroup

	lineFormCallback.CreationMode = (line == nil)

	return
}

type LineFormCallback struct {
	line *models.Line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (lineFormCallback *LineFormCallback) OnSave() {
	lineFormCallback.probe.stageOfInterest.Lock()
	defer lineFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.probe.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.probe.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	for _, formDiv := range lineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(line_.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(line_.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(line_.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(line_.Y2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(line_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(line_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(line_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(line_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(line_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(line_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(line_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(line_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](lineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					lineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](lineFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			line_.Animates = instanceSlice
			lineFormCallback.probe.UpdateSliceOfPointersCallback(line_, "Animates", &line_.Animates)

		case "MouseClickX":
			FormDivBasicFieldToField(&(line_.MouseClickX), formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&(line_.MouseClickY), formDiv)
		case "Layer:Lines":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](lineFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Lines slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](lineFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(lineFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure line_ is in _layer.Lines
					found := false
					for _, _b := range _layer.Lines {
						if _b == line_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Lines = append(_layer.Lines, line_)
						lineFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Lines", &_layer.Lines)
					}
				} else {
					// ensure line_ is NOT in _layer.Lines
					idx := slices.Index(_layer.Lines, line_)
					if idx != -1 {
						_layer.Lines = slices.Delete(_layer.Lines, idx, idx+1)
						lineFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Lines", &_layer.Lines)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_.Unstage(lineFormCallback.probe.stageOfInterest)
	}

	lineFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Line](
		lineFormCallback.probe,
	)

	// display a new form by reset the form stage
	if lineFormCallback.CreationMode || lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lineFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(lineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LineFormCallback(
			nil,
			lineFormCallback.probe,
			newFormGroup,
		)
		line := new(models.Line)
		FillUpForm(line, newFormGroup, lineFormCallback.probe)
		lineFormCallback.probe.formStage.Commit()
	}

	lineFormCallback.probe.ux_tree()
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {
	linkFormCallback.probe.stageOfInterest.Lock()
	defer linkFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(link_.Type), formDiv)
		case "IsBezierCurve":
			FormDivBasicFieldToField(&(link_.IsBezierCurve), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(link_.Start), linkFormCallback.probe.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(link_.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(link_.End), linkFormCallback.probe.stageOfInterest, formDiv)
		case "EndAnchorType":
			FormDivEnumStringFieldToField(&(link_.EndAnchorType), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&(link_.CornerRadius), formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&(link_.HasEndArrow), formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&(link_.EndArrowSize), formDiv)
		case "EndArrowOffset":
			FormDivBasicFieldToField(&(link_.EndArrowOffset), formDiv)
		case "HasStartArrow":
			FormDivBasicFieldToField(&(link_.HasStartArrow), formDiv)
		case "StartArrowSize":
			FormDivBasicFieldToField(&(link_.StartArrowSize), formDiv)
		case "StartArrowOffset":
			FormDivBasicFieldToField(&(link_.StartArrowOffset), formDiv)
		case "TextAtArrowStart":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.TextAtArrowStart = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "TextAtArrowStart", &link_.TextAtArrowStart)

		case "TextAtArrowEnd":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.TextAtArrowEnd = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "TextAtArrowEnd", &link_.TextAtArrowEnd)

		case "TextAtCorner":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.TextAtCorner = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "TextAtCorner", &link_.TextAtCorner)

		case "PathAtArrowStart":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.PathAtArrowStart = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "PathAtArrowStart", &link_.PathAtArrowStart)

		case "PathAtArrowEnd":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.PathAtArrowEnd = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "PathAtArrowEnd", &link_.PathAtArrowEnd)

		case "PathAtCorner":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAnchoredPath](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.PathAtCorner = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "PathAtCorner", &link_.PathAtCorner)

		case "ControlPoints":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPoint](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPoint, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPoint)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPoint](linkFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			link_.ControlPoints = instanceSlice
			linkFormCallback.probe.UpdateSliceOfPointersCallback(link_, "ControlPoints", &link_.ControlPoints)

		case "Color":
			FormDivBasicFieldToField(&(link_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(link_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(link_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(link_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(link_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(link_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(link_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(link_.Transform), formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&(link_.MouseX), formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&(link_.MouseY), formDiv)
		case "MouseEventKey":
			FormDivEnumStringFieldToField(&(link_.MouseEventKey), formDiv)
		case "Layer:Links":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](linkFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Links slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](linkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure link_ is in _layer.Links
					found := false
					for _, _b := range _layer.Links {
						if _b == link_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Links = append(_layer.Links, link_)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Links", &_layer.Links)
					}
				} else {
					// ensure link_ is NOT in _layer.Links
					idx := slices.Index(_layer.Links, link_)
					if idx != -1 {
						_layer.Links = slices.Delete(_layer.Links, idx, idx+1)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Links", &_layer.Links)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Link](
		linkFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	linkFormCallback.probe.ux_tree()
}
func __gong__New__LinkAnchoredPathFormCallback(
	linkanchoredpath *models.LinkAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkanchoredpathFormCallback *LinkAnchoredPathFormCallback) {
	linkanchoredpathFormCallback = new(LinkAnchoredPathFormCallback)
	linkanchoredpathFormCallback.probe = probe
	linkanchoredpathFormCallback.linkanchoredpath = linkanchoredpath
	linkanchoredpathFormCallback.formGroup = formGroup

	linkanchoredpathFormCallback.CreationMode = (linkanchoredpath == nil)

	return
}

type LinkAnchoredPathFormCallback struct {
	linkanchoredpath *models.LinkAnchoredPath

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkanchoredpathFormCallback *LinkAnchoredPathFormCallback) OnSave() {
	linkanchoredpathFormCallback.probe.stageOfInterest.Lock()
	defer linkanchoredpathFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkAnchoredPathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredpathFormCallback.probe.formStage.Checkout()

	if linkanchoredpathFormCallback.linkanchoredpath == nil {
		linkanchoredpathFormCallback.linkanchoredpath = new(models.LinkAnchoredPath).Stage(linkanchoredpathFormCallback.probe.stageOfInterest)
	}
	linkanchoredpath_ := linkanchoredpathFormCallback.linkanchoredpath
	_ = linkanchoredpath_

	for _, formDiv := range linkanchoredpathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkanchoredpath_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(linkanchoredpath_.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(linkanchoredpath_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(linkanchoredpath_.Y_Offset), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(linkanchoredpath_.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(linkanchoredpath_.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(linkanchoredpath_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(linkanchoredpath_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(linkanchoredpath_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(linkanchoredpath_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(linkanchoredpath_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(linkanchoredpath_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(linkanchoredpath_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(linkanchoredpath_.Transform), formDiv)
		case "Link:PathAtArrowStart":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their PathAtArrowStart slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredpathFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredpath_ is in _link.PathAtArrowStart
					found := false
					for _, _b := range _link.PathAtArrowStart {
						if _b == linkanchoredpath_ {
							found = true
							break
						}
					}
					if !found {
						_link.PathAtArrowStart = append(_link.PathAtArrowStart, linkanchoredpath_)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtArrowStart", &_link.PathAtArrowStart)
					}
				} else {
					// ensure linkanchoredpath_ is NOT in _link.PathAtArrowStart
					idx := slices.Index(_link.PathAtArrowStart, linkanchoredpath_)
					if idx != -1 {
						_link.PathAtArrowStart = slices.Delete(_link.PathAtArrowStart, idx, idx+1)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtArrowStart", &_link.PathAtArrowStart)
					}
				}
			}
		case "Link:PathAtArrowEnd":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their PathAtArrowEnd slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredpathFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredpath_ is in _link.PathAtArrowEnd
					found := false
					for _, _b := range _link.PathAtArrowEnd {
						if _b == linkanchoredpath_ {
							found = true
							break
						}
					}
					if !found {
						_link.PathAtArrowEnd = append(_link.PathAtArrowEnd, linkanchoredpath_)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtArrowEnd", &_link.PathAtArrowEnd)
					}
				} else {
					// ensure linkanchoredpath_ is NOT in _link.PathAtArrowEnd
					idx := slices.Index(_link.PathAtArrowEnd, linkanchoredpath_)
					if idx != -1 {
						_link.PathAtArrowEnd = slices.Delete(_link.PathAtArrowEnd, idx, idx+1)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtArrowEnd", &_link.PathAtArrowEnd)
					}
				}
			}
		case "Link:PathAtCorner":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their PathAtCorner slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredpathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredpathFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredpath_ is in _link.PathAtCorner
					found := false
					for _, _b := range _link.PathAtCorner {
						if _b == linkanchoredpath_ {
							found = true
							break
						}
					}
					if !found {
						_link.PathAtCorner = append(_link.PathAtCorner, linkanchoredpath_)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtCorner", &_link.PathAtCorner)
					}
				} else {
					// ensure linkanchoredpath_ is NOT in _link.PathAtCorner
					idx := slices.Index(_link.PathAtCorner, linkanchoredpath_)
					if idx != -1 {
						_link.PathAtCorner = slices.Delete(_link.PathAtCorner, idx, idx+1)
						linkanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_link, "PathAtCorner", &_link.PathAtCorner)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredpath_.Unstage(linkanchoredpathFormCallback.probe.stageOfInterest)
	}

	linkanchoredpathFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.LinkAnchoredPath](
		linkanchoredpathFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkanchoredpathFormCallback.CreationMode || linkanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredpathFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkanchoredpathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkAnchoredPathFormCallback(
			nil,
			linkanchoredpathFormCallback.probe,
			newFormGroup,
		)
		linkanchoredpath := new(models.LinkAnchoredPath)
		FillUpForm(linkanchoredpath, newFormGroup, linkanchoredpathFormCallback.probe)
		linkanchoredpathFormCallback.probe.formStage.Commit()
	}

	linkanchoredpathFormCallback.probe.ux_tree()
}
func __gong__New__LinkAnchoredTextFormCallback(
	linkanchoredtext *models.LinkAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) {
	linkanchoredtextFormCallback = new(LinkAnchoredTextFormCallback)
	linkanchoredtextFormCallback.probe = probe
	linkanchoredtextFormCallback.linkanchoredtext = linkanchoredtext
	linkanchoredtextFormCallback.formGroup = formGroup

	linkanchoredtextFormCallback.CreationMode = (linkanchoredtext == nil)

	return
}

type LinkAnchoredTextFormCallback struct {
	linkanchoredtext *models.LinkAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) OnSave() {
	linkanchoredtextFormCallback.probe.stageOfInterest.Lock()
	defer linkanchoredtextFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredtextFormCallback.probe.formStage.Checkout()

	if linkanchoredtextFormCallback.linkanchoredtext == nil {
		linkanchoredtextFormCallback.linkanchoredtext = new(models.LinkAnchoredText).Stage(linkanchoredtextFormCallback.probe.stageOfInterest)
	}
	linkanchoredtext_ := linkanchoredtextFormCallback.linkanchoredtext
	_ = linkanchoredtext_

	for _, formDiv := range linkanchoredtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(linkanchoredtext_.Content), formDiv)
		case "AutomaticLayout":
			FormDivBasicFieldToField(&(linkanchoredtext_.AutomaticLayout), formDiv)
		case "LinkAnchorType":
			FormDivEnumStringFieldToField(&(linkanchoredtext_.LinkAnchorType), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.Y_Offset), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(linkanchoredtext_.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(linkanchoredtext_.WhiteSpace), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(linkanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(linkanchoredtext_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(linkanchoredtext_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](linkanchoredtextFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkanchoredtextFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](linkanchoredtextFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			linkanchoredtext_.Animates = instanceSlice
			linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(linkanchoredtext_, "Animates", &linkanchoredtext_.Animates)

		case "Link:TextAtArrowStart":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their TextAtArrowStart slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredtextFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredtext_ is in _link.TextAtArrowStart
					found := false
					for _, _b := range _link.TextAtArrowStart {
						if _b == linkanchoredtext_ {
							found = true
							break
						}
					}
					if !found {
						_link.TextAtArrowStart = append(_link.TextAtArrowStart, linkanchoredtext_)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtArrowStart", &_link.TextAtArrowStart)
					}
				} else {
					// ensure linkanchoredtext_ is NOT in _link.TextAtArrowStart
					idx := slices.Index(_link.TextAtArrowStart, linkanchoredtext_)
					if idx != -1 {
						_link.TextAtArrowStart = slices.Delete(_link.TextAtArrowStart, idx, idx+1)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtArrowStart", &_link.TextAtArrowStart)
					}
				}
			}
		case "Link:TextAtArrowEnd":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their TextAtArrowEnd slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredtextFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredtext_ is in _link.TextAtArrowEnd
					found := false
					for _, _b := range _link.TextAtArrowEnd {
						if _b == linkanchoredtext_ {
							found = true
							break
						}
					}
					if !found {
						_link.TextAtArrowEnd = append(_link.TextAtArrowEnd, linkanchoredtext_)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtArrowEnd", &_link.TextAtArrowEnd)
					}
				} else {
					// ensure linkanchoredtext_ is NOT in _link.TextAtArrowEnd
					idx := slices.Index(_link.TextAtArrowEnd, linkanchoredtext_)
					if idx != -1 {
						_link.TextAtArrowEnd = slices.Delete(_link.TextAtArrowEnd, idx, idx+1)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtArrowEnd", &_link.TextAtArrowEnd)
					}
				}
			}
		case "Link:TextAtCorner":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Link instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Link instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest)
			targetLinkIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLinkIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Link instances and update their TextAtCorner slice
			for _link := range *models.GetGongstructInstancesSetFromPointerType[*models.Link](linkanchoredtextFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkanchoredtextFormCallback.probe.stageOfInterest, _link)
				
				// if Link is selected
				if targetLinkIDs[id] {
					// ensure linkanchoredtext_ is in _link.TextAtCorner
					found := false
					for _, _b := range _link.TextAtCorner {
						if _b == linkanchoredtext_ {
							found = true
							break
						}
					}
					if !found {
						_link.TextAtCorner = append(_link.TextAtCorner, linkanchoredtext_)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtCorner", &_link.TextAtCorner)
					}
				} else {
					// ensure linkanchoredtext_ is NOT in _link.TextAtCorner
					idx := slices.Index(_link.TextAtCorner, linkanchoredtext_)
					if idx != -1 {
						_link.TextAtCorner = slices.Delete(_link.TextAtCorner, idx, idx+1)
						linkanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_link, "TextAtCorner", &_link.TextAtCorner)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredtext_.Unstage(linkanchoredtextFormCallback.probe.stageOfInterest)
	}

	linkanchoredtextFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.LinkAnchoredText](
		linkanchoredtextFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkanchoredtextFormCallback.CreationMode || linkanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkanchoredtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkAnchoredTextFormCallback(
			nil,
			linkanchoredtextFormCallback.probe,
			newFormGroup,
		)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, newFormGroup, linkanchoredtextFormCallback.probe)
		linkanchoredtextFormCallback.probe.formStage.Commit()
	}

	linkanchoredtextFormCallback.probe.ux_tree()
}
func __gong__New__PathFormCallback(
	path *models.Path,
	probe *Probe,
	formGroup *form.FormGroup,
) (pathFormCallback *PathFormCallback) {
	pathFormCallback = new(PathFormCallback)
	pathFormCallback.probe = probe
	pathFormCallback.path = path
	pathFormCallback.formGroup = formGroup

	pathFormCallback.CreationMode = (path == nil)

	return
}

type PathFormCallback struct {
	path *models.Path

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (pathFormCallback *PathFormCallback) OnSave() {
	pathFormCallback.probe.stageOfInterest.Lock()
	defer pathFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pathFormCallback.probe.formStage.Checkout()

	if pathFormCallback.path == nil {
		pathFormCallback.path = new(models.Path).Stage(pathFormCallback.probe.stageOfInterest)
	}
	path_ := pathFormCallback.path
	_ = path_

	for _, formDiv := range pathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(path_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(path_.Definition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(path_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(path_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(path_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(path_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(path_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(path_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(path_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(path_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](pathFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					pathFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](pathFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			path_.Animates = instanceSlice
			pathFormCallback.probe.UpdateSliceOfPointersCallback(path_, "Animates", &path_.Animates)

		case "Layer:Paths":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](pathFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Paths slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](pathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(pathFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure path_ is in _layer.Paths
					found := false
					for _, _b := range _layer.Paths {
						if _b == path_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Paths = append(_layer.Paths, path_)
						pathFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Paths", &_layer.Paths)
					}
				} else {
					// ensure path_ is NOT in _layer.Paths
					idx := slices.Index(_layer.Paths, path_)
					if idx != -1 {
						_layer.Paths = slices.Delete(_layer.Paths, idx, idx+1)
						pathFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Paths", &_layer.Paths)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if pathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		path_.Unstage(pathFormCallback.probe.stageOfInterest)
	}

	pathFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Path](
		pathFormCallback.probe,
	)

	// display a new form by reset the form stage
	if pathFormCallback.CreationMode || pathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pathFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(pathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PathFormCallback(
			nil,
			pathFormCallback.probe,
			newFormGroup,
		)
		path := new(models.Path)
		FillUpForm(path, newFormGroup, pathFormCallback.probe)
		pathFormCallback.probe.formStage.Commit()
	}

	pathFormCallback.probe.ux_tree()
}
func __gong__New__PointFormCallback(
	point *models.Point,
	probe *Probe,
	formGroup *form.FormGroup,
) (pointFormCallback *PointFormCallback) {
	pointFormCallback = new(PointFormCallback)
	pointFormCallback.probe = probe
	pointFormCallback.point = point
	pointFormCallback.formGroup = formGroup

	pointFormCallback.CreationMode = (point == nil)

	return
}

type PointFormCallback struct {
	point *models.Point

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (pointFormCallback *PointFormCallback) OnSave() {
	pointFormCallback.probe.stageOfInterest.Lock()
	defer pointFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointFormCallback.probe.formStage.Checkout()

	if pointFormCallback.point == nil {
		pointFormCallback.point = new(models.Point).Stage(pointFormCallback.probe.stageOfInterest)
	}
	point_ := pointFormCallback.point
	_ = point_

	for _, formDiv := range pointFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(point_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(point_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(point_.Y), formDiv)
		}
	}

	// manage the suppress operation
	if pointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		point_.Unstage(pointFormCallback.probe.stageOfInterest)
	}

	pointFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Point](
		pointFormCallback.probe,
	)

	// display a new form by reset the form stage
	if pointFormCallback.CreationMode || pointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(pointFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PointFormCallback(
			nil,
			pointFormCallback.probe,
			newFormGroup,
		)
		point := new(models.Point)
		FillUpForm(point, newFormGroup, pointFormCallback.probe)
		pointFormCallback.probe.formStage.Commit()
	}

	pointFormCallback.probe.ux_tree()
}
func __gong__New__PolygoneFormCallback(
	polygone *models.Polygone,
	probe *Probe,
	formGroup *form.FormGroup,
) (polygoneFormCallback *PolygoneFormCallback) {
	polygoneFormCallback = new(PolygoneFormCallback)
	polygoneFormCallback.probe = probe
	polygoneFormCallback.polygone = polygone
	polygoneFormCallback.formGroup = formGroup

	polygoneFormCallback.CreationMode = (polygone == nil)

	return
}

type PolygoneFormCallback struct {
	polygone *models.Polygone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (polygoneFormCallback *PolygoneFormCallback) OnSave() {
	polygoneFormCallback.probe.stageOfInterest.Lock()
	defer polygoneFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PolygoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polygoneFormCallback.probe.formStage.Checkout()

	if polygoneFormCallback.polygone == nil {
		polygoneFormCallback.polygone = new(models.Polygone).Stage(polygoneFormCallback.probe.stageOfInterest)
	}
	polygone_ := polygoneFormCallback.polygone
	_ = polygone_

	for _, formDiv := range polygoneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polygone_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polygone_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polygone_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polygone_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polygone_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(polygone_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polygone_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polygone_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](polygoneFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					polygoneFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](polygoneFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			polygone_.Animates = instanceSlice
			polygoneFormCallback.probe.UpdateSliceOfPointersCallback(polygone_, "Animates", &polygone_.Animates)

		case "Layer:Polygones":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](polygoneFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Polygones slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](polygoneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(polygoneFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure polygone_ is in _layer.Polygones
					found := false
					for _, _b := range _layer.Polygones {
						if _b == polygone_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Polygones = append(_layer.Polygones, polygone_)
						polygoneFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Polygones", &_layer.Polygones)
					}
				} else {
					// ensure polygone_ is NOT in _layer.Polygones
					idx := slices.Index(_layer.Polygones, polygone_)
					if idx != -1 {
						_layer.Polygones = slices.Delete(_layer.Polygones, idx, idx+1)
						polygoneFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Polygones", &_layer.Polygones)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if polygoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polygone_.Unstage(polygoneFormCallback.probe.stageOfInterest)
	}

	polygoneFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Polygone](
		polygoneFormCallback.probe,
	)

	// display a new form by reset the form stage
	if polygoneFormCallback.CreationMode || polygoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polygoneFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(polygoneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PolygoneFormCallback(
			nil,
			polygoneFormCallback.probe,
			newFormGroup,
		)
		polygone := new(models.Polygone)
		FillUpForm(polygone, newFormGroup, polygoneFormCallback.probe)
		polygoneFormCallback.probe.formStage.Commit()
	}

	polygoneFormCallback.probe.ux_tree()
}
func __gong__New__PolylineFormCallback(
	polyline *models.Polyline,
	probe *Probe,
	formGroup *form.FormGroup,
) (polylineFormCallback *PolylineFormCallback) {
	polylineFormCallback = new(PolylineFormCallback)
	polylineFormCallback.probe = probe
	polylineFormCallback.polyline = polyline
	polylineFormCallback.formGroup = formGroup

	polylineFormCallback.CreationMode = (polyline == nil)

	return
}

type PolylineFormCallback struct {
	polyline *models.Polyline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (polylineFormCallback *PolylineFormCallback) OnSave() {
	polylineFormCallback.probe.stageOfInterest.Lock()
	defer polylineFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PolylineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polylineFormCallback.probe.formStage.Checkout()

	if polylineFormCallback.polyline == nil {
		polylineFormCallback.polyline = new(models.Polyline).Stage(polylineFormCallback.probe.stageOfInterest)
	}
	polyline_ := polylineFormCallback.polyline
	_ = polyline_

	for _, formDiv := range polylineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polyline_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polyline_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polyline_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polyline_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polyline_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(polyline_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polyline_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polyline_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](polylineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					polylineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](polylineFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			polyline_.Animates = instanceSlice
			polylineFormCallback.probe.UpdateSliceOfPointersCallback(polyline_, "Animates", &polyline_.Animates)

		case "Layer:Polylines":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](polylineFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Polylines slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](polylineFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(polylineFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure polyline_ is in _layer.Polylines
					found := false
					for _, _b := range _layer.Polylines {
						if _b == polyline_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Polylines = append(_layer.Polylines, polyline_)
						polylineFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Polylines", &_layer.Polylines)
					}
				} else {
					// ensure polyline_ is NOT in _layer.Polylines
					idx := slices.Index(_layer.Polylines, polyline_)
					if idx != -1 {
						_layer.Polylines = slices.Delete(_layer.Polylines, idx, idx+1)
						polylineFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Polylines", &_layer.Polylines)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if polylineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polyline_.Unstage(polylineFormCallback.probe.stageOfInterest)
	}

	polylineFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Polyline](
		polylineFormCallback.probe,
	)

	// display a new form by reset the form stage
	if polylineFormCallback.CreationMode || polylineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polylineFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(polylineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PolylineFormCallback(
			nil,
			polylineFormCallback.probe,
			newFormGroup,
		)
		polyline := new(models.Polyline)
		FillUpForm(polyline, newFormGroup, polylineFormCallback.probe)
		polylineFormCallback.probe.formStage.Commit()
	}

	polylineFormCallback.probe.ux_tree()
}
func __gong__New__RectFormCallback(
	rect *models.Rect,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectFormCallback *RectFormCallback) {
	rectFormCallback = new(RectFormCallback)
	rectFormCallback.probe = probe
	rectFormCallback.rect = rect
	rectFormCallback.formGroup = formGroup

	rectFormCallback.CreationMode = (rect == nil)

	return
}

type RectFormCallback struct {
	rect *models.Rect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectFormCallback *RectFormCallback) OnSave() {
	rectFormCallback.probe.stageOfInterest.Lock()
	defer rectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectFormCallback.probe.formStage.Checkout()

	if rectFormCallback.rect == nil {
		rectFormCallback.rect = new(models.Rect).Stage(rectFormCallback.probe.stageOfInterest)
	}
	rect_ := rectFormCallback.rect
	_ = rect_

	for _, formDiv := range rectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rect_.RX), formDiv)
		case "Peers":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Rect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Rect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.Peers = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "Peers", &rect_.Peers)

		case "EnclosingRect":
			FormDivSelectFieldToField(&(rect_.EnclosingRect), rectFormCallback.probe.stageOfInterest, formDiv)
		case "Obstacles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Rect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Rect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.Obstacles = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "Obstacles", &rect_.Obstacles)

		case "AnchoredTo":
			FormDivSelectFieldToField(&(rect_.AnchoredTo), rectFormCallback.probe.stageOfInterest, formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rect_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rect_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rect_.Transform), formDiv)
		case "HoveringTrigger":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Condition](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Condition, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Condition)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Condition](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.HoveringTrigger = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "HoveringTrigger", &rect_.HoveringTrigger)

		case "DisplayConditions":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Condition](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Condition, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Condition)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Condition](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.DisplayConditions = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "DisplayConditions", &rect_.DisplayConditions)

		case "Animations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.Animations = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "Animations", &rect_.Animations)

		case "IsSelectable":
			FormDivBasicFieldToField(&(rect_.IsSelectable), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(rect_.IsSelected), formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveLeftHandle), formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&(rect_.HasLeftHandle), formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveRightHandle), formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&(rect_.HasRightHandle), formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveTopHandle), formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&(rect_.HasTopHandle), formDiv)
		case "IsScalingProportionally":
			FormDivBasicFieldToField(&(rect_.IsScalingProportionally), formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveBottomHandle), formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&(rect_.HasBottomHandle), formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&(rect_.CanMoveHorizontaly), formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&(rect_.CanMoveVerticaly), formDiv)
		case "RectAnchoredTexts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RectAnchoredText](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.RectAnchoredTexts = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "RectAnchoredTexts", &rect_.RectAnchoredTexts)

		case "RectAnchoredRects":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredRect](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredRect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredRect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RectAnchoredRect](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.RectAnchoredRects = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "RectAnchoredRects", &rect_.RectAnchoredRects)

		case "RectAnchoredPaths":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPath](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RectAnchoredPath](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.RectAnchoredPaths = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "RectAnchoredPaths", &rect_.RectAnchoredPaths)

		case "RectAnchoredPngImages":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPngImage](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredPngImage, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredPngImage)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RectAnchoredPngImage](rectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rect_.RectAnchoredPngImages = instanceSlice
			rectFormCallback.probe.UpdateSliceOfPointersCallback(rect_, "RectAnchoredPngImages", &rect_.RectAnchoredPngImages)

		case "ChangeColorWhenHovered":
			FormDivBasicFieldToField(&(rect_.ChangeColorWhenHovered), formDiv)
		case "ColorWhenHovered":
			FormDivBasicFieldToField(&(rect_.ColorWhenHovered), formDiv)
		case "OriginalColor":
			FormDivBasicFieldToField(&(rect_.OriginalColor), formDiv)
		case "FillOpacityWhenHovered":
			FormDivBasicFieldToField(&(rect_.FillOpacityWhenHovered), formDiv)
		case "OriginalFillOpacity":
			FormDivBasicFieldToField(&(rect_.OriginalFillOpacity), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(rect_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(rect_.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(rect_.ToolTipPosition), formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&(rect_.MouseX), formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&(rect_.MouseY), formDiv)
		case "MouseEventKey":
			FormDivEnumStringFieldToField(&(rect_.MouseEventKey), formDiv)
		case "URLPath":
			FormDivBasicFieldToField(&(rect_.URLPath), formDiv)
		case "URLTarget":
			FormDivEnumStringFieldToField(&(rect_.URLTarget), formDiv)
		case "Layer:Rects":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](rectFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Rects slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](rectFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure rect_ is in _layer.Rects
					found := false
					for _, _b := range _layer.Rects {
						if _b == rect_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Rects = append(_layer.Rects, rect_)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Rects", &_layer.Rects)
					}
				} else {
					// ensure rect_ is NOT in _layer.Rects
					idx := slices.Index(_layer.Rects, rect_)
					if idx != -1 {
						_layer.Rects = slices.Delete(_layer.Rects, idx, idx+1)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Rects", &_layer.Rects)
					}
				}
			}
		case "Rect:Peers":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their Peers slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rect_ is in _rect.Peers
					found := false
					for _, _b := range _rect.Peers {
						if _b == rect_ {
							found = true
							break
						}
					}
					if !found {
						_rect.Peers = append(_rect.Peers, rect_)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Peers", &_rect.Peers)
					}
				} else {
					// ensure rect_ is NOT in _rect.Peers
					idx := slices.Index(_rect.Peers, rect_)
					if idx != -1 {
						_rect.Peers = slices.Delete(_rect.Peers, idx, idx+1)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Peers", &_rect.Peers)
					}
				}
			}
		case "Rect:Obstacles":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their Obstacles slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rect_ is in _rect.Obstacles
					found := false
					for _, _b := range _rect.Obstacles {
						if _b == rect_ {
							found = true
							break
						}
					}
					if !found {
						_rect.Obstacles = append(_rect.Obstacles, rect_)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Obstacles", &_rect.Obstacles)
					}
				} else {
					// ensure rect_ is NOT in _rect.Obstacles
					idx := slices.Index(_rect.Obstacles, rect_)
					if idx != -1 {
						_rect.Obstacles = slices.Delete(_rect.Obstacles, idx, idx+1)
						rectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "Obstacles", &_rect.Obstacles)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rect_.Unstage(rectFormCallback.probe.stageOfInterest)
	}

	rectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Rect](
		rectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectFormCallback.CreationMode || rectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectFormCallback(
			nil,
			rectFormCallback.probe,
			newFormGroup,
		)
		rect := new(models.Rect)
		FillUpForm(rect, newFormGroup, rectFormCallback.probe)
		rectFormCallback.probe.formStage.Commit()
	}

	rectFormCallback.probe.ux_tree()
}
func __gong__New__RectAnchoredPathFormCallback(
	rectanchoredpath *models.RectAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredpathFormCallback *RectAnchoredPathFormCallback) {
	rectanchoredpathFormCallback = new(RectAnchoredPathFormCallback)
	rectanchoredpathFormCallback.probe = probe
	rectanchoredpathFormCallback.rectanchoredpath = rectanchoredpath
	rectanchoredpathFormCallback.formGroup = formGroup

	rectanchoredpathFormCallback.CreationMode = (rectanchoredpath == nil)

	return
}

type RectAnchoredPathFormCallback struct {
	rectanchoredpath *models.RectAnchoredPath

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectanchoredpathFormCallback *RectAnchoredPathFormCallback) OnSave() {
	rectanchoredpathFormCallback.probe.stageOfInterest.Lock()
	defer rectanchoredpathFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectAnchoredPathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredpathFormCallback.probe.formStage.Checkout()

	if rectanchoredpathFormCallback.rectanchoredpath == nil {
		rectanchoredpathFormCallback.rectanchoredpath = new(models.RectAnchoredPath).Stage(rectanchoredpathFormCallback.probe.stageOfInterest)
	}
	rectanchoredpath_ := rectanchoredpathFormCallback.rectanchoredpath
	_ = rectanchoredpath_

	for _, formDiv := range rectanchoredpathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredpath_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(rectanchoredpath_.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredpath_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredpath_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredpath_.RectAnchorType), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(rectanchoredpath_.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(rectanchoredpath_.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredpath_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredpath_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredpath_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredpath_.Transform), formDiv)
		case "Rect:RectAnchoredPaths":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectanchoredpathFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their RectAnchoredPaths slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectanchoredpathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectanchoredpathFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rectanchoredpath_ is in _rect.RectAnchoredPaths
					found := false
					for _, _b := range _rect.RectAnchoredPaths {
						if _b == rectanchoredpath_ {
							found = true
							break
						}
					}
					if !found {
						_rect.RectAnchoredPaths = append(_rect.RectAnchoredPaths, rectanchoredpath_)
						rectanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredPaths", &_rect.RectAnchoredPaths)
					}
				} else {
					// ensure rectanchoredpath_ is NOT in _rect.RectAnchoredPaths
					idx := slices.Index(_rect.RectAnchoredPaths, rectanchoredpath_)
					if idx != -1 {
						_rect.RectAnchoredPaths = slices.Delete(_rect.RectAnchoredPaths, idx, idx+1)
						rectanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredPaths", &_rect.RectAnchoredPaths)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpath_.Unstage(rectanchoredpathFormCallback.probe.stageOfInterest)
	}

	rectanchoredpathFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RectAnchoredPath](
		rectanchoredpathFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectanchoredpathFormCallback.CreationMode || rectanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpathFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredpathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredPathFormCallback(
			nil,
			rectanchoredpathFormCallback.probe,
			newFormGroup,
		)
		rectanchoredpath := new(models.RectAnchoredPath)
		FillUpForm(rectanchoredpath, newFormGroup, rectanchoredpathFormCallback.probe)
		rectanchoredpathFormCallback.probe.formStage.Commit()
	}

	rectanchoredpathFormCallback.probe.ux_tree()
}
func __gong__New__RectAnchoredPngImageFormCallback(
	rectanchoredpngimage *models.RectAnchoredPngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredpngimageFormCallback *RectAnchoredPngImageFormCallback) {
	rectanchoredpngimageFormCallback = new(RectAnchoredPngImageFormCallback)
	rectanchoredpngimageFormCallback.probe = probe
	rectanchoredpngimageFormCallback.rectanchoredpngimage = rectanchoredpngimage
	rectanchoredpngimageFormCallback.formGroup = formGroup

	rectanchoredpngimageFormCallback.CreationMode = (rectanchoredpngimage == nil)

	return
}

type RectAnchoredPngImageFormCallback struct {
	rectanchoredpngimage *models.RectAnchoredPngImage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectanchoredpngimageFormCallback *RectAnchoredPngImageFormCallback) OnSave() {
	rectanchoredpngimageFormCallback.probe.stageOfInterest.Lock()
	defer rectanchoredpngimageFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectAnchoredPngImageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredpngimageFormCallback.probe.formStage.Checkout()

	if rectanchoredpngimageFormCallback.rectanchoredpngimage == nil {
		rectanchoredpngimageFormCallback.rectanchoredpngimage = new(models.RectAnchoredPngImage).Stage(rectanchoredpngimageFormCallback.probe.stageOfInterest)
	}
	rectanchoredpngimage_ := rectanchoredpngimageFormCallback.rectanchoredpngimage
	_ = rectanchoredpngimage_

	for _, formDiv := range rectanchoredpngimageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredpngimage_.RectAnchorType), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(rectanchoredpngimage_.Base64Content), formDiv)
		case "Rect:RectAnchoredPngImages":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectanchoredpngimageFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their RectAnchoredPngImages slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectanchoredpngimageFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectanchoredpngimageFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rectanchoredpngimage_ is in _rect.RectAnchoredPngImages
					found := false
					for _, _b := range _rect.RectAnchoredPngImages {
						if _b == rectanchoredpngimage_ {
							found = true
							break
						}
					}
					if !found {
						_rect.RectAnchoredPngImages = append(_rect.RectAnchoredPngImages, rectanchoredpngimage_)
						rectanchoredpngimageFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredPngImages", &_rect.RectAnchoredPngImages)
					}
				} else {
					// ensure rectanchoredpngimage_ is NOT in _rect.RectAnchoredPngImages
					idx := slices.Index(_rect.RectAnchoredPngImages, rectanchoredpngimage_)
					if idx != -1 {
						_rect.RectAnchoredPngImages = slices.Delete(_rect.RectAnchoredPngImages, idx, idx+1)
						rectanchoredpngimageFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredPngImages", &_rect.RectAnchoredPngImages)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectanchoredpngimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpngimage_.Unstage(rectanchoredpngimageFormCallback.probe.stageOfInterest)
	}

	rectanchoredpngimageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RectAnchoredPngImage](
		rectanchoredpngimageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectanchoredpngimageFormCallback.CreationMode || rectanchoredpngimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpngimageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredpngimageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredPngImageFormCallback(
			nil,
			rectanchoredpngimageFormCallback.probe,
			newFormGroup,
		)
		rectanchoredpngimage := new(models.RectAnchoredPngImage)
		FillUpForm(rectanchoredpngimage, newFormGroup, rectanchoredpngimageFormCallback.probe)
		rectanchoredpngimageFormCallback.probe.formStage.Commit()
	}

	rectanchoredpngimageFormCallback.probe.ux_tree()
}
func __gong__New__RectAnchoredRectFormCallback(
	rectanchoredrect *models.RectAnchoredRect,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) {
	rectanchoredrectFormCallback = new(RectAnchoredRectFormCallback)
	rectanchoredrectFormCallback.probe = probe
	rectanchoredrectFormCallback.rectanchoredrect = rectanchoredrect
	rectanchoredrectFormCallback.formGroup = formGroup

	rectanchoredrectFormCallback.CreationMode = (rectanchoredrect == nil)

	return
}

type RectAnchoredRectFormCallback struct {
	rectanchoredrect *models.RectAnchoredRect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) OnSave() {
	rectanchoredrectFormCallback.probe.stageOfInterest.Lock()
	defer rectanchoredrectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectAnchoredRectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredrectFormCallback.probe.formStage.Checkout()

	if rectanchoredrectFormCallback.rectanchoredrect == nil {
		rectanchoredrectFormCallback.rectanchoredrect = new(models.RectAnchoredRect).Stage(rectanchoredrectFormCallback.probe.stageOfInterest)
	}
	rectanchoredrect_ := rectanchoredrectFormCallback.rectanchoredrect
	_ = rectanchoredrect_

	for _, formDiv := range rectanchoredrectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredrect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rectanchoredrect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rectanchoredrect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rectanchoredrect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rectanchoredrect_.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredrect_.RectAnchorType), formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.WidthFollowRect), formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.HeightFollowRect), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(rectanchoredrect_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(rectanchoredrect_.ToolTipText), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredrect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredrect_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredrect_.Transform), formDiv)
		case "Rect:RectAnchoredRects":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectanchoredrectFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their RectAnchoredRects slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectanchoredrectFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectanchoredrectFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rectanchoredrect_ is in _rect.RectAnchoredRects
					found := false
					for _, _b := range _rect.RectAnchoredRects {
						if _b == rectanchoredrect_ {
							found = true
							break
						}
					}
					if !found {
						_rect.RectAnchoredRects = append(_rect.RectAnchoredRects, rectanchoredrect_)
						rectanchoredrectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredRects", &_rect.RectAnchoredRects)
					}
				} else {
					// ensure rectanchoredrect_ is NOT in _rect.RectAnchoredRects
					idx := slices.Index(_rect.RectAnchoredRects, rectanchoredrect_)
					if idx != -1 {
						_rect.RectAnchoredRects = slices.Delete(_rect.RectAnchoredRects, idx, idx+1)
						rectanchoredrectFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredRects", &_rect.RectAnchoredRects)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectanchoredrectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredrect_.Unstage(rectanchoredrectFormCallback.probe.stageOfInterest)
	}

	rectanchoredrectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RectAnchoredRect](
		rectanchoredrectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectanchoredrectFormCallback.CreationMode || rectanchoredrectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredrectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredrectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredRectFormCallback(
			nil,
			rectanchoredrectFormCallback.probe,
			newFormGroup,
		)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, newFormGroup, rectanchoredrectFormCallback.probe)
		rectanchoredrectFormCallback.probe.formStage.Commit()
	}

	rectanchoredrectFormCallback.probe.ux_tree()
}
func __gong__New__RectAnchoredTextFormCallback(
	rectanchoredtext *models.RectAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) {
	rectanchoredtextFormCallback = new(RectAnchoredTextFormCallback)
	rectanchoredtextFormCallback.probe = probe
	rectanchoredtextFormCallback.rectanchoredtext = rectanchoredtext
	rectanchoredtextFormCallback.formGroup = formGroup

	rectanchoredtextFormCallback.CreationMode = (rectanchoredtext == nil)

	return
}

type RectAnchoredTextFormCallback struct {
	rectanchoredtext *models.RectAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) OnSave() {
	rectanchoredtextFormCallback.probe.stageOfInterest.Lock()
	defer rectanchoredtextFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredtextFormCallback.probe.formStage.Checkout()

	if rectanchoredtextFormCallback.rectanchoredtext == nil {
		rectanchoredtextFormCallback.rectanchoredtext = new(models.RectAnchoredText).Stage(rectanchoredtextFormCallback.probe.stageOfInterest)
	}
	rectanchoredtext_ := rectanchoredtextFormCallback.rectanchoredtext
	_ = rectanchoredtext_

	for _, formDiv := range rectanchoredtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rectanchoredtext_.Content), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(rectanchoredtext_.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.WhiteSpace), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.RectAnchorType), formDiv)
		case "TextAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.TextAnchorType), formDiv)
		case "DominantBaseline":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.DominantBaseline), formDiv)
		case "WritingMode":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.WritingMode), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredtext_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredtext_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](rectanchoredtextFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectanchoredtextFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](rectanchoredtextFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rectanchoredtext_.Animates = instanceSlice
			rectanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(rectanchoredtext_, "Animates", &rectanchoredtext_.Animates)

		case "URLPath":
			FormDivBasicFieldToField(&(rectanchoredtext_.URLPath), formDiv)
		case "URLTarget":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.URLTarget), formDiv)
		case "Rect:RectAnchoredTexts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Rect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Rect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Rect](rectanchoredtextFormCallback.probe.stageOfInterest)
			targetRectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Rect instances and update their RectAnchoredTexts slice
			for _rect := range *models.GetGongstructInstancesSetFromPointerType[*models.Rect](rectanchoredtextFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectanchoredtextFormCallback.probe.stageOfInterest, _rect)
				
				// if Rect is selected
				if targetRectIDs[id] {
					// ensure rectanchoredtext_ is in _rect.RectAnchoredTexts
					found := false
					for _, _b := range _rect.RectAnchoredTexts {
						if _b == rectanchoredtext_ {
							found = true
							break
						}
					}
					if !found {
						_rect.RectAnchoredTexts = append(_rect.RectAnchoredTexts, rectanchoredtext_)
						rectanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredTexts", &_rect.RectAnchoredTexts)
					}
				} else {
					// ensure rectanchoredtext_ is NOT in _rect.RectAnchoredTexts
					idx := slices.Index(_rect.RectAnchoredTexts, rectanchoredtext_)
					if idx != -1 {
						_rect.RectAnchoredTexts = slices.Delete(_rect.RectAnchoredTexts, idx, idx+1)
						rectanchoredtextFormCallback.probe.UpdateSliceOfPointersCallback(_rect, "RectAnchoredTexts", &_rect.RectAnchoredTexts)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredtext_.Unstage(rectanchoredtextFormCallback.probe.stageOfInterest)
	}

	rectanchoredtextFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RectAnchoredText](
		rectanchoredtextFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectanchoredtextFormCallback.CreationMode || rectanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredTextFormCallback(
			nil,
			rectanchoredtextFormCallback.probe,
			newFormGroup,
		)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, newFormGroup, rectanchoredtextFormCallback.probe)
		rectanchoredtextFormCallback.probe.formStage.Commit()
	}

	rectanchoredtextFormCallback.probe.ux_tree()
}
func __gong__New__RectLinkLinkFormCallback(
	rectlinklink *models.RectLinkLink,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectlinklinkFormCallback *RectLinkLinkFormCallback) {
	rectlinklinkFormCallback = new(RectLinkLinkFormCallback)
	rectlinklinkFormCallback.probe = probe
	rectlinklinkFormCallback.rectlinklink = rectlinklink
	rectlinklinkFormCallback.formGroup = formGroup

	rectlinklinkFormCallback.CreationMode = (rectlinklink == nil)

	return
}

type RectLinkLinkFormCallback struct {
	rectlinklink *models.RectLinkLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rectlinklinkFormCallback *RectLinkLinkFormCallback) OnSave() {
	rectlinklinkFormCallback.probe.stageOfInterest.Lock()
	defer rectlinklinkFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RectLinkLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectlinklinkFormCallback.probe.formStage.Checkout()

	if rectlinklinkFormCallback.rectlinklink == nil {
		rectlinklinkFormCallback.rectlinklink = new(models.RectLinkLink).Stage(rectlinklinkFormCallback.probe.stageOfInterest)
	}
	rectlinklink_ := rectlinklinkFormCallback.rectlinklink
	_ = rectlinklink_

	for _, formDiv := range rectlinklinkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectlinklink_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(rectlinklink_.Start), rectlinklinkFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(rectlinklink_.End), rectlinklinkFormCallback.probe.stageOfInterest, formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&(rectlinklink_.TargetAnchorPosition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectlinklink_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectlinklink_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectlinklink_.Transform), formDiv)
		case "Layer:RectLinkLinks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](rectlinklinkFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their RectLinkLinks slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](rectlinklinkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rectlinklinkFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure rectlinklink_ is in _layer.RectLinkLinks
					found := false
					for _, _b := range _layer.RectLinkLinks {
						if _b == rectlinklink_ {
							found = true
							break
						}
					}
					if !found {
						_layer.RectLinkLinks = append(_layer.RectLinkLinks, rectlinklink_)
						rectlinklinkFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "RectLinkLinks", &_layer.RectLinkLinks)
					}
				} else {
					// ensure rectlinklink_ is NOT in _layer.RectLinkLinks
					idx := slices.Index(_layer.RectLinkLinks, rectlinklink_)
					if idx != -1 {
						_layer.RectLinkLinks = slices.Delete(_layer.RectLinkLinks, idx, idx+1)
						rectlinklinkFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "RectLinkLinks", &_layer.RectLinkLinks)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rectlinklinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectlinklink_.Unstage(rectlinklinkFormCallback.probe.stageOfInterest)
	}

	rectlinklinkFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RectLinkLink](
		rectlinklinkFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rectlinklinkFormCallback.CreationMode || rectlinklinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectlinklinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rectlinklinkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectLinkLinkFormCallback(
			nil,
			rectlinklinkFormCallback.probe,
			newFormGroup,
		)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, newFormGroup, rectlinklinkFormCallback.probe)
		rectlinklinkFormCallback.probe.formStage.Commit()
	}

	rectlinklinkFormCallback.probe.ux_tree()
}
func __gong__New__SVGFormCallback(
	svg *models.SVG,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgFormCallback *SVGFormCallback) {
	svgFormCallback = new(SVGFormCallback)
	svgFormCallback.probe = probe
	svgFormCallback.svg = svg
	svgFormCallback.formGroup = formGroup

	svgFormCallback.CreationMode = (svg == nil)

	return
}

type SVGFormCallback struct {
	svg *models.SVG

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (svgFormCallback *SVGFormCallback) OnSave() {
	svgFormCallback.probe.stageOfInterest.Lock()
	defer svgFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SVGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.probe.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.SVG).Stage(svgFormCallback.probe.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	for _, formDiv := range svgFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "Layers":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Layer](svgFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Layer, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Layer)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					svgFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](svgFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			svg_.Layers = instanceSlice
			svgFormCallback.probe.UpdateSliceOfPointersCallback(svg_, "Layers", &svg_.Layers)

		case "DrawingState":
			FormDivEnumStringFieldToField(&(svg_.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(svg_.StartRect), svgFormCallback.probe.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(svg_.EndRect), svgFormCallback.probe.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(svg_.IsEditable), formDiv)
		case "IsSVGFrontEndFileGenerated":
			FormDivBasicFieldToField(&(svg_.IsSVGFrontEndFileGenerated), formDiv)
		case "IsSVGBackEndFileGenerated":
			FormDivBasicFieldToField(&(svg_.IsSVGBackEndFileGenerated), formDiv)
		case "DefaultDirectoryForGeneratedImages":
			FormDivBasicFieldToField(&(svg_.DefaultDirectoryForGeneratedImages), formDiv)
		case "IsControlBannerHidden":
			FormDivBasicFieldToField(&(svg_.IsControlBannerHidden), formDiv)
		case "OverrideWidth":
			FormDivBasicFieldToField(&(svg_.OverrideWidth), formDiv)
		case "OverriddenWidth":
			FormDivBasicFieldToField(&(svg_.OverriddenWidth), formDiv)
		case "OverrideHeight":
			FormDivBasicFieldToField(&(svg_.OverrideHeight), formDiv)
		case "OverriddenHeight":
			FormDivBasicFieldToField(&(svg_.OverriddenHeight), formDiv)
		}
	}

	// manage the suppress operation
	if svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svg_.Unstage(svgFormCallback.probe.stageOfInterest)
	}

	svgFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SVG](
		svgFormCallback.probe,
	)

	// display a new form by reset the form stage
	if svgFormCallback.CreationMode || svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(svgFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SVGFormCallback(
			nil,
			svgFormCallback.probe,
			newFormGroup,
		)
		svg := new(models.SVG)
		FillUpForm(svg, newFormGroup, svgFormCallback.probe)
		svgFormCallback.probe.formStage.Commit()
	}

	svgFormCallback.probe.ux_tree()
}
func __gong__New__SvgTextFormCallback(
	svgtext *models.SvgText,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgtextFormCallback *SvgTextFormCallback) {
	svgtextFormCallback = new(SvgTextFormCallback)
	svgtextFormCallback.probe = probe
	svgtextFormCallback.svgtext = svgtext
	svgtextFormCallback.formGroup = formGroup

	svgtextFormCallback.CreationMode = (svgtext == nil)

	return
}

type SvgTextFormCallback struct {
	svgtext *models.SvgText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (svgtextFormCallback *SvgTextFormCallback) OnSave() {
	svgtextFormCallback.probe.stageOfInterest.Lock()
	defer svgtextFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SvgTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgtextFormCallback.probe.formStage.Checkout()

	if svgtextFormCallback.svgtext == nil {
		svgtextFormCallback.svgtext = new(models.SvgText).Stage(svgtextFormCallback.probe.stageOfInterest)
	}
	svgtext_ := svgtextFormCallback.svgtext
	_ = svgtext_

	for _, formDiv := range svgtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgtext_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(svgtext_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if svgtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgtext_.Unstage(svgtextFormCallback.probe.stageOfInterest)
	}

	svgtextFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SvgText](
		svgtextFormCallback.probe,
	)

	// display a new form by reset the form stage
	if svgtextFormCallback.CreationMode || svgtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(svgtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SvgTextFormCallback(
			nil,
			svgtextFormCallback.probe,
			newFormGroup,
		)
		svgtext := new(models.SvgText)
		FillUpForm(svgtext, newFormGroup, svgtextFormCallback.probe)
		svgtextFormCallback.probe.formStage.Commit()
	}

	svgtextFormCallback.probe.ux_tree()
}
func __gong__New__TextFormCallback(
	text *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.probe = probe
	textFormCallback.text = text
	textFormCallback.formGroup = formGroup

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (textFormCallback *TextFormCallback) OnSave() {
	textFormCallback.probe.stageOfInterest.Lock()
	defer textFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.probe.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.probe.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	for _, formDiv := range textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(text_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(text_.Y), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(text_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(text_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(text_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(text_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(text_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(text_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(text_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(text_.Transform), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(text_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(text_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(text_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(text_.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(text_.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(text_.WhiteSpace), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](textFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					textFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Animate](textFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			text_.Animates = instanceSlice
			textFormCallback.probe.UpdateSliceOfPointersCallback(text_, "Animates", &text_.Animates)

		case "Layer:Texts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Layer instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Layer instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Layer](textFormCallback.probe.stageOfInterest)
			targetLayerIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLayerIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Layer instances and update their Texts slice
			for _layer := range *models.GetGongstructInstancesSetFromPointerType[*models.Layer](textFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(textFormCallback.probe.stageOfInterest, _layer)
				
				// if Layer is selected
				if targetLayerIDs[id] {
					// ensure text_ is in _layer.Texts
					found := false
					for _, _b := range _layer.Texts {
						if _b == text_ {
							found = true
							break
						}
					}
					if !found {
						_layer.Texts = append(_layer.Texts, text_)
						textFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Texts", &_layer.Texts)
					}
				} else {
					// ensure text_ is NOT in _layer.Texts
					idx := slices.Index(_layer.Texts, text_)
					if idx != -1 {
						_layer.Texts = slices.Delete(_layer.Texts, idx, idx+1)
						textFormCallback.probe.UpdateSliceOfPointersCallback(_layer, "Texts", &_layer.Texts)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		text_.Unstage(textFormCallback.probe.stageOfInterest)
	}

	textFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Text](
		textFormCallback.probe,
	)

	// display a new form by reset the form stage
	if textFormCallback.CreationMode || textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		textFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TextFormCallback(
			nil,
			textFormCallback.probe,
			newFormGroup,
		)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.probe)
		textFormCallback.probe.formStage.Commit()
	}

	textFormCallback.probe.ux_tree()
}
