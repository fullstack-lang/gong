package orm

import "github.com/fullstack-lang/gong/test/go/models"

func GetReverseFieldStringValue[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		if tmp.Astruct_AnarrayofbDBID.Int64 != 0 {
			id := uint(tmp.Astruct_AnarrayofbDBID.Int64)
			reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
			res = reservePointerTarget.Name
		}
	default:
		_ = inst
	}
	return
}
