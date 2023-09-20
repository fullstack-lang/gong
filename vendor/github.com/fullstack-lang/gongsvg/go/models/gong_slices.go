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
