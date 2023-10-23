// generated code - do not edit
package x

import "github.com/fullstack-lang/gong/test2/go/models/x"

func GetReverseFieldOwnerName[T x.Gongstruct](
	stage *x.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *x.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *x.A:
		tmp := GetInstanceDBFromInstance[x.A, ADB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}

	case *x.B:
		tmp := GetInstanceDBFromInstance[x.B, BDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				if _a, ok := stage.A_Bs_reverseMap[inst]; ok {
					res = _a.Name
				}
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T x.Gongstruct](
	stage *x.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *x.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *x.A:
		tmp := GetInstanceDBFromInstance[x.A, ADB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}

	case *x.B:
		tmp := GetInstanceDBFromInstance[x.B, BDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				res = stage.A_Bs_reverseMap[inst]
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
