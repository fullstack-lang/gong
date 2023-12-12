// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Animate:
		// insertion point per field

	case *Circle:
		// insertion point per field
		if fieldName == "Animations" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Circle) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Circle)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animations).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animations = _inferedTypeInstance.Animations[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animations =
								append(_inferedTypeInstance.Animations, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Ellipse:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Ellipse) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Ellipse)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Layer:
		// insertion point per field
		if fieldName == "Rects" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rects).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rects = _inferedTypeInstance.Rects[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rects =
								append(_inferedTypeInstance.Rects, any(fieldInstance).(*Rect))
						}
					}
				}
			}
		}
		if fieldName == "Texts" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Texts).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Texts = _inferedTypeInstance.Texts[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Texts =
								append(_inferedTypeInstance.Texts, any(fieldInstance).(*Text))
						}
					}
				}
			}
		}
		if fieldName == "Circles" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Circles).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Circles = _inferedTypeInstance.Circles[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Circles =
								append(_inferedTypeInstance.Circles, any(fieldInstance).(*Circle))
						}
					}
				}
			}
		}
		if fieldName == "Lines" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lines).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lines = _inferedTypeInstance.Lines[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lines =
								append(_inferedTypeInstance.Lines, any(fieldInstance).(*Line))
						}
					}
				}
			}
		}
		if fieldName == "Ellipses" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Ellipses).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Ellipses = _inferedTypeInstance.Ellipses[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Ellipses =
								append(_inferedTypeInstance.Ellipses, any(fieldInstance).(*Ellipse))
						}
					}
				}
			}
		}
		if fieldName == "Polylines" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Polylines).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Polylines = _inferedTypeInstance.Polylines[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Polylines =
								append(_inferedTypeInstance.Polylines, any(fieldInstance).(*Polyline))
						}
					}
				}
			}
		}
		if fieldName == "Polygones" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Polygones).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Polygones = _inferedTypeInstance.Polygones[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Polygones =
								append(_inferedTypeInstance.Polygones, any(fieldInstance).(*Polygone))
						}
					}
				}
			}
		}
		if fieldName == "Paths" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Paths).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Paths = _inferedTypeInstance.Paths[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Paths =
								append(_inferedTypeInstance.Paths, any(fieldInstance).(*Path))
						}
					}
				}
			}
		}
		if fieldName == "Links" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Links).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Links = _inferedTypeInstance.Links[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Links =
								append(_inferedTypeInstance.Links, any(fieldInstance).(*Link))
						}
					}
				}
			}
		}
		if fieldName == "RectLinkLinks" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Layer) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Layer)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RectLinkLinks).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RectLinkLinks = _inferedTypeInstance.RectLinkLinks[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RectLinkLinks =
								append(_inferedTypeInstance.RectLinkLinks, any(fieldInstance).(*RectLinkLink))
						}
					}
				}
			}
		}

	case *Line:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Line) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Line)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Link:
		// insertion point per field
		if fieldName == "TextAtArrowEnd" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Link) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Link)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TextAtArrowEnd).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TextAtArrowEnd = _inferedTypeInstance.TextAtArrowEnd[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TextAtArrowEnd =
								append(_inferedTypeInstance.TextAtArrowEnd, any(fieldInstance).(*LinkAnchoredText))
						}
					}
				}
			}
		}
		if fieldName == "TextAtArrowStart" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Link) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Link)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TextAtArrowStart).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TextAtArrowStart = _inferedTypeInstance.TextAtArrowStart[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TextAtArrowStart =
								append(_inferedTypeInstance.TextAtArrowStart, any(fieldInstance).(*LinkAnchoredText))
						}
					}
				}
			}
		}
		if fieldName == "ControlPoints" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Link) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Link)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ControlPoints).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ControlPoints = _inferedTypeInstance.ControlPoints[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ControlPoints =
								append(_inferedTypeInstance.ControlPoints, any(fieldInstance).(*Point))
						}
					}
				}
			}
		}

	case *LinkAnchoredText:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*LinkAnchoredText) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*LinkAnchoredText)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Path:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Path) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Path)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Point:
		// insertion point per field

	case *Polygone:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Polygone) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Polygone)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Polyline:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Polyline) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Polyline)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *Rect:
		// insertion point per field
		if fieldName == "Animations" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Rect) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Rect)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animations).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animations = _inferedTypeInstance.Animations[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animations =
								append(_inferedTypeInstance.Animations, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}
		if fieldName == "RectAnchoredTexts" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Rect) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Rect)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RectAnchoredTexts).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RectAnchoredTexts = _inferedTypeInstance.RectAnchoredTexts[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RectAnchoredTexts =
								append(_inferedTypeInstance.RectAnchoredTexts, any(fieldInstance).(*RectAnchoredText))
						}
					}
				}
			}
		}
		if fieldName == "RectAnchoredRects" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Rect) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Rect)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RectAnchoredRects).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RectAnchoredRects = _inferedTypeInstance.RectAnchoredRects[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RectAnchoredRects =
								append(_inferedTypeInstance.RectAnchoredRects, any(fieldInstance).(*RectAnchoredRect))
						}
					}
				}
			}
		}
		if fieldName == "RectAnchoredPaths" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Rect) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Rect)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RectAnchoredPaths).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RectAnchoredPaths = _inferedTypeInstance.RectAnchoredPaths[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RectAnchoredPaths =
								append(_inferedTypeInstance.RectAnchoredPaths, any(fieldInstance).(*RectAnchoredPath))
						}
					}
				}
			}
		}

	case *RectAnchoredPath:
		// insertion point per field

	case *RectAnchoredRect:
		// insertion point per field

	case *RectAnchoredText:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RectAnchoredText) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RectAnchoredText)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	case *RectLinkLink:
		// insertion point per field

	case *SVG:
		// insertion point per field
		if fieldName == "Layers" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SVG) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SVG)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Layers).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Layers = _inferedTypeInstance.Layers[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Layers =
								append(_inferedTypeInstance.Layers, any(fieldInstance).(*Layer))
						}
					}
				}
			}
		}

	case *Text:
		// insertion point per field
		if fieldName == "Animates" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Text) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Text)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Animates).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Animates = _inferedTypeInstance.Animates[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Animates =
								append(_inferedTypeInstance.Animates, any(fieldInstance).(*Animate))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Animate
	// insertion point per field

	// Compute reverse map for named struct Circle
	// insertion point per field
	clear(stage.Circle_Animations_reverseMap)
	stage.Circle_Animations_reverseMap = make(map[*Animate]*Circle)
	for circle := range stage.Circles {
		_ = circle
		for _, _animate := range circle.Animations {
			stage.Circle_Animations_reverseMap[_animate] = circle
		}
	}

	// Compute reverse map for named struct Ellipse
	// insertion point per field
	clear(stage.Ellipse_Animates_reverseMap)
	stage.Ellipse_Animates_reverseMap = make(map[*Animate]*Ellipse)
	for ellipse := range stage.Ellipses {
		_ = ellipse
		for _, _animate := range ellipse.Animates {
			stage.Ellipse_Animates_reverseMap[_animate] = ellipse
		}
	}

	// Compute reverse map for named struct Layer
	// insertion point per field
	clear(stage.Layer_Rects_reverseMap)
	stage.Layer_Rects_reverseMap = make(map[*Rect]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _rect := range layer.Rects {
			stage.Layer_Rects_reverseMap[_rect] = layer
		}
	}
	clear(stage.Layer_Texts_reverseMap)
	stage.Layer_Texts_reverseMap = make(map[*Text]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _text := range layer.Texts {
			stage.Layer_Texts_reverseMap[_text] = layer
		}
	}
	clear(stage.Layer_Circles_reverseMap)
	stage.Layer_Circles_reverseMap = make(map[*Circle]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _circle := range layer.Circles {
			stage.Layer_Circles_reverseMap[_circle] = layer
		}
	}
	clear(stage.Layer_Lines_reverseMap)
	stage.Layer_Lines_reverseMap = make(map[*Line]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _line := range layer.Lines {
			stage.Layer_Lines_reverseMap[_line] = layer
		}
	}
	clear(stage.Layer_Ellipses_reverseMap)
	stage.Layer_Ellipses_reverseMap = make(map[*Ellipse]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _ellipse := range layer.Ellipses {
			stage.Layer_Ellipses_reverseMap[_ellipse] = layer
		}
	}
	clear(stage.Layer_Polylines_reverseMap)
	stage.Layer_Polylines_reverseMap = make(map[*Polyline]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polyline := range layer.Polylines {
			stage.Layer_Polylines_reverseMap[_polyline] = layer
		}
	}
	clear(stage.Layer_Polygones_reverseMap)
	stage.Layer_Polygones_reverseMap = make(map[*Polygone]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polygone := range layer.Polygones {
			stage.Layer_Polygones_reverseMap[_polygone] = layer
		}
	}
	clear(stage.Layer_Paths_reverseMap)
	stage.Layer_Paths_reverseMap = make(map[*Path]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _path := range layer.Paths {
			stage.Layer_Paths_reverseMap[_path] = layer
		}
	}
	clear(stage.Layer_Links_reverseMap)
	stage.Layer_Links_reverseMap = make(map[*Link]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _link := range layer.Links {
			stage.Layer_Links_reverseMap[_link] = layer
		}
	}
	clear(stage.Layer_RectLinkLinks_reverseMap)
	stage.Layer_RectLinkLinks_reverseMap = make(map[*RectLinkLink]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _rectlinklink := range layer.RectLinkLinks {
			stage.Layer_RectLinkLinks_reverseMap[_rectlinklink] = layer
		}
	}

	// Compute reverse map for named struct Line
	// insertion point per field
	clear(stage.Line_Animates_reverseMap)
	stage.Line_Animates_reverseMap = make(map[*Animate]*Line)
	for line := range stage.Lines {
		_ = line
		for _, _animate := range line.Animates {
			stage.Line_Animates_reverseMap[_animate] = line
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field
	clear(stage.Link_TextAtArrowEnd_reverseMap)
	stage.Link_TextAtArrowEnd_reverseMap = make(map[*LinkAnchoredText]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			stage.Link_TextAtArrowEnd_reverseMap[_linkanchoredtext] = link
		}
	}
	clear(stage.Link_TextAtArrowStart_reverseMap)
	stage.Link_TextAtArrowStart_reverseMap = make(map[*LinkAnchoredText]*Link)
	for link := range stage.Links {
		_ = link
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			stage.Link_TextAtArrowStart_reverseMap[_linkanchoredtext] = link
		}
	}
	clear(stage.Link_ControlPoints_reverseMap)
	stage.Link_ControlPoints_reverseMap = make(map[*Point]*Link)
	for link := range stage.Links {
		_ = link
		for _, _point := range link.ControlPoints {
			stage.Link_ControlPoints_reverseMap[_point] = link
		}
	}

	// Compute reverse map for named struct LinkAnchoredText
	// insertion point per field
	clear(stage.LinkAnchoredText_Animates_reverseMap)
	stage.LinkAnchoredText_Animates_reverseMap = make(map[*Animate]*LinkAnchoredText)
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		_ = linkanchoredtext
		for _, _animate := range linkanchoredtext.Animates {
			stage.LinkAnchoredText_Animates_reverseMap[_animate] = linkanchoredtext
		}
	}

	// Compute reverse map for named struct Path
	// insertion point per field
	clear(stage.Path_Animates_reverseMap)
	stage.Path_Animates_reverseMap = make(map[*Animate]*Path)
	for path := range stage.Paths {
		_ = path
		for _, _animate := range path.Animates {
			stage.Path_Animates_reverseMap[_animate] = path
		}
	}

	// Compute reverse map for named struct Point
	// insertion point per field

	// Compute reverse map for named struct Polygone
	// insertion point per field
	clear(stage.Polygone_Animates_reverseMap)
	stage.Polygone_Animates_reverseMap = make(map[*Animate]*Polygone)
	for polygone := range stage.Polygones {
		_ = polygone
		for _, _animate := range polygone.Animates {
			stage.Polygone_Animates_reverseMap[_animate] = polygone
		}
	}

	// Compute reverse map for named struct Polyline
	// insertion point per field
	clear(stage.Polyline_Animates_reverseMap)
	stage.Polyline_Animates_reverseMap = make(map[*Animate]*Polyline)
	for polyline := range stage.Polylines {
		_ = polyline
		for _, _animate := range polyline.Animates {
			stage.Polyline_Animates_reverseMap[_animate] = polyline
		}
	}

	// Compute reverse map for named struct Rect
	// insertion point per field
	clear(stage.Rect_Animations_reverseMap)
	stage.Rect_Animations_reverseMap = make(map[*Animate]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _animate := range rect.Animations {
			stage.Rect_Animations_reverseMap[_animate] = rect
		}
	}
	clear(stage.Rect_RectAnchoredTexts_reverseMap)
	stage.Rect_RectAnchoredTexts_reverseMap = make(map[*RectAnchoredText]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			stage.Rect_RectAnchoredTexts_reverseMap[_rectanchoredtext] = rect
		}
	}
	clear(stage.Rect_RectAnchoredRects_reverseMap)
	stage.Rect_RectAnchoredRects_reverseMap = make(map[*RectAnchoredRect]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			stage.Rect_RectAnchoredRects_reverseMap[_rectanchoredrect] = rect
		}
	}
	clear(stage.Rect_RectAnchoredPaths_reverseMap)
	stage.Rect_RectAnchoredPaths_reverseMap = make(map[*RectAnchoredPath]*Rect)
	for rect := range stage.Rects {
		_ = rect
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			stage.Rect_RectAnchoredPaths_reverseMap[_rectanchoredpath] = rect
		}
	}

	// Compute reverse map for named struct RectAnchoredPath
	// insertion point per field

	// Compute reverse map for named struct RectAnchoredRect
	// insertion point per field

	// Compute reverse map for named struct RectAnchoredText
	// insertion point per field
	clear(stage.RectAnchoredText_Animates_reverseMap)
	stage.RectAnchoredText_Animates_reverseMap = make(map[*Animate]*RectAnchoredText)
	for rectanchoredtext := range stage.RectAnchoredTexts {
		_ = rectanchoredtext
		for _, _animate := range rectanchoredtext.Animates {
			stage.RectAnchoredText_Animates_reverseMap[_animate] = rectanchoredtext
		}
	}

	// Compute reverse map for named struct RectLinkLink
	// insertion point per field

	// Compute reverse map for named struct SVG
	// insertion point per field
	clear(stage.SVG_Layers_reverseMap)
	stage.SVG_Layers_reverseMap = make(map[*Layer]*SVG)
	for svg := range stage.SVGs {
		_ = svg
		for _, _layer := range svg.Layers {
			stage.SVG_Layers_reverseMap[_layer] = svg
		}
	}

	// Compute reverse map for named struct Text
	// insertion point per field
	clear(stage.Text_Animates_reverseMap)
	stage.Text_Animates_reverseMap = make(map[*Animate]*Text)
	for text := range stage.Texts {
		_ = text
		for _, _animate := range text.Animates {
			stage.Text_Animates_reverseMap[_animate] = text
		}
	}

}
