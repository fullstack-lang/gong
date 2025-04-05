// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Command:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *DummyAgent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Status:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *UpdateState:
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
	case *Command:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *DummyAgent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Status:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *UpdateState:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
