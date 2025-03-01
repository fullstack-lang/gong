// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/slider/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Checkbox:
		checkboxInstance := any(concreteInstance).(*models.Checkbox)
		ret2 := backRepo.BackRepoCheckbox.GetCheckboxDBFromCheckboxPtr(checkboxInstance)
		ret = any(ret2).(*T2)
	case *models.Group:
		groupInstance := any(concreteInstance).(*models.Group)
		ret2 := backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupInstance)
		ret = any(ret2).(*T2)
	case *models.Layout:
		layoutInstance := any(concreteInstance).(*models.Layout)
		ret2 := backRepo.BackRepoLayout.GetLayoutDBFromLayoutPtr(layoutInstance)
		ret = any(ret2).(*T2)
	case *models.Slider:
		sliderInstance := any(concreteInstance).(*models.Slider)
		ret2 := backRepo.BackRepoSlider.GetSliderDBFromSliderPtr(sliderInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Checkbox:
		tmp := GetInstanceDBFromInstance[models.Checkbox, CheckboxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Layout:
		tmp := GetInstanceDBFromInstance[models.Layout, LayoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slider:
		tmp := GetInstanceDBFromInstance[models.Slider, SliderDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Checkbox:
		tmp := GetInstanceDBFromInstance[models.Checkbox, CheckboxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Layout:
		tmp := GetInstanceDBFromInstance[models.Layout, LayoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slider:
		tmp := GetInstanceDBFromInstance[models.Slider, SliderDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
