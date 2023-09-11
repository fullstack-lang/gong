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
func EvictInOtherSlices[T PointerToGongstruct, TF PointerToGongstruct](
	stage *StageStruct,
	owningInstance T,
	sliceField []TF,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[TF]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Animate:
		// insertion point per field

	case *Circle:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animations" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Circle)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animations).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animations = make([]*Animate, 0)
				if any(_instance).(*Circle) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Ellipse:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Ellipse)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Ellipse) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Layer:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Rects" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Rects).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Rects = make([]*Rect, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Texts" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Texts).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Texts = make([]*Text, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Circles" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Circles).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Circles = make([]*Circle, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Lines" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Lines).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Lines = make([]*Line, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Ellipses" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Ellipses).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Ellipses = make([]*Ellipse, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Polylines" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Polylines).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Polylines = make([]*Polyline, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Polygones" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Polygones).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Polygones = make([]*Polygone, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Paths" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Paths).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Paths = make([]*Path, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Links" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Links).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Links = make([]*Link, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "RectLinkLinks" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Layer)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.RectLinkLinks).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.RectLinkLinks = make([]*RectLinkLink, 0)
				if any(_instance).(*Layer) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Line:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Line)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Line) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Link:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "TextAtArrowEnd" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Link)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.TextAtArrowEnd).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.TextAtArrowEnd = make([]*LinkAnchoredText, 0)
				if any(_instance).(*Link) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "TextAtArrowStart" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Link)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.TextAtArrowStart).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.TextAtArrowStart = make([]*LinkAnchoredText, 0)
				if any(_instance).(*Link) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "ControlPoints" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Link)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.ControlPoints).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.ControlPoints = make([]*Point, 0)
				if any(_instance).(*Link) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *LinkAnchoredText:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*LinkAnchoredText)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*LinkAnchoredText) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Path:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Path)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Path) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Point:
		// insertion point per field

	case *Polygone:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Polygone)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Polygone) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Polyline:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Polyline)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Polyline) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Rect:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animations" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Rect)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animations).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animations = make([]*Animate, 0)
				if any(_instance).(*Rect) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "RectAnchoredTexts" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Rect)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.RectAnchoredTexts).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.RectAnchoredTexts = make([]*RectAnchoredText, 0)
				if any(_instance).(*Rect) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "RectAnchoredRects" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Rect)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.RectAnchoredRects).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.RectAnchoredRects = make([]*RectAnchoredRect, 0)
				if any(_instance).(*Rect) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *RectAnchoredRect:
		// insertion point per field

	case *RectAnchoredText:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*RectAnchoredText)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*RectAnchoredText) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *RectLinkLink:
		// insertion point per field

	case *SVG:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Layers" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*SVG)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Layers).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Layers = make([]*Layer, 0)
				if any(_instance).(*SVG) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Text:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Animates" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Text)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Animates).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Animates = make([]*Animate, 0)
				if any(_instance).(*Text) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
