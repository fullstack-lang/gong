// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *AsSplit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *AsSplitArea:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplit":
			switch reverseField.Fieldname {
			case "AsSplitAreas":
				if _assplit, ok := stage.AsSplit_AsSplitAreas_reverseMap[inst]; ok {
					res = _assplit.Name
				}
			}
		case "View":
			switch reverseField.Fieldname {
			case "RootAsSplitAreas":
				if _view, ok := stage.View_RootAsSplitAreas_reverseMap[inst]; ok {
					res = _view.Name
				}
			}
		}

	case *Button:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Cursor:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Doc:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Form:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Load:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Slider:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Split:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Svg:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Tone:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Tree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *View:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *AsSplit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *AsSplitArea:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplit":
			switch reverseField.Fieldname {
			case "AsSplitAreas":
				res = stage.AsSplit_AsSplitAreas_reverseMap[inst]
			}
		case "View":
			switch reverseField.Fieldname {
			case "RootAsSplitAreas":
				res = stage.View_RootAsSplitAreas_reverseMap[inst]
			}
		}

	case *Button:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Cursor:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Doc:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Form:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Load:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Slider:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Split:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Svg:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Tone:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Tree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *View:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
