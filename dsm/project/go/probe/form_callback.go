// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__DiagramHierarchyFormCallback(
	diagramhierarchy *models.DiagramHierarchy,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramhierarchyFormCallback *DiagramHierarchyFormCallback) {
	diagramhierarchyFormCallback = new(DiagramHierarchyFormCallback)
	diagramhierarchyFormCallback.probe = probe
	diagramhierarchyFormCallback.diagramhierarchy = diagramhierarchy
	diagramhierarchyFormCallback.formGroup = formGroup

	diagramhierarchyFormCallback.CreationMode = (diagramhierarchy == nil)

	return
}

type DiagramHierarchyFormCallback struct {
	diagramhierarchy *models.DiagramHierarchy

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramhierarchyFormCallback *DiagramHierarchyFormCallback) OnSave() {
	diagramhierarchyFormCallback.probe.stageOfInterest.Lock()
	defer diagramhierarchyFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramHierarchyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramhierarchyFormCallback.probe.formStage.Checkout()

	if diagramhierarchyFormCallback.diagramhierarchy == nil {
		diagramhierarchyFormCallback.diagramhierarchy = new(models.DiagramHierarchy).Stage(diagramhierarchyFormCallback.probe.stageOfInterest)
	}
	diagramhierarchy_ := diagramhierarchyFormCallback.diagramhierarchy
	_ = diagramhierarchy_

	for _, formDiv := range diagramhierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramhierarchy_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramhierarchy_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramhierarchy_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramhierarchy_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramhierarchy_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramhierarchy_.Height), formDiv)
		case "Product_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.Product_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "Product_Shapes", &diagramhierarchy_.Product_Shapes)

		case "ProductsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.ProductsWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "ProductsWhoseNodeIsExpanded", &diagramhierarchy_.ProductsWhoseNodeIsExpanded)

		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsPBSNodeExpanded), formDiv)
		case "ProductComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.ProductComposition_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "ProductComposition_Shapes", &diagramhierarchy_.ProductComposition_Shapes)

		case "IsWBSNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsWBSNodeExpanded), formDiv)
		case "Task_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.Task_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "Task_Shapes", &diagramhierarchy_.Task_Shapes)

		case "TasksWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TasksWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TasksWhoseNodeIsExpanded", &diagramhierarchy_.TasksWhoseNodeIsExpanded)

		case "TasksWhoseInputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TasksWhoseInputNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TasksWhoseInputNodeIsExpanded", &diagramhierarchy_.TasksWhoseInputNodeIsExpanded)

		case "TasksWhoseOutputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TasksWhoseOutputNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TasksWhoseOutputNodeIsExpanded", &diagramhierarchy_.TasksWhoseOutputNodeIsExpanded)

		case "IsTaskGroupsNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsTaskGroupsNodeExpanded), formDiv)
		case "TaskGroupShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskGroupShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskGroupShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskGroupShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskGroupShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TaskGroupShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TaskGroupShapes", &diagramhierarchy_.TaskGroupShapes)

		case "TaskGroupsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskGroup](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskGroup](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TaskGroupsWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TaskGroupsWhoseNodeIsExpanded", &diagramhierarchy_.TaskGroupsWhoseNodeIsExpanded)

		case "IsMilestonesNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsMilestonesNodeExpanded), formDiv)
		case "MilestoneShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.MilestoneShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.MilestoneShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.MilestoneShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.MilestoneShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.MilestoneShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "MilestoneShapes", &diagramhierarchy_.MilestoneShapes)

		case "MilestonesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Milestone, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Milestone)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Milestone](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.MilestonesWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "MilestonesWhoseNodeIsExpanded", &diagramhierarchy_.MilestonesWhoseNodeIsExpanded)

		case "DateFormat":
			FormDivBasicFieldToField(&(diagramhierarchy_.DateFormat), formDiv)
		case "TaskComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TaskComposition_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TaskComposition_Shapes", &diagramhierarchy_.TaskComposition_Shapes)

		case "TaskInputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskInputShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskInputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskInputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskInputShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TaskInputShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TaskInputShapes", &diagramhierarchy_.TaskInputShapes)

		case "TaskOutputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskOutputShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskOutputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskOutputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskOutputShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.TaskOutputShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "TaskOutputShapes", &diagramhierarchy_.TaskOutputShapes)

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.Note_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "Note_Shapes", &diagramhierarchy_.Note_Shapes)

		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "NotesWhoseNodeIsExpanded", &diagramhierarchy_.NotesWhoseNodeIsExpanded)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsNotesNodeExpanded), formDiv)
		case "NoteProductShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteProductShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteProductShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.NoteProductShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "NoteProductShapes", &diagramhierarchy_.NoteProductShapes)

		case "NoteTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteTaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.NoteTaskShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "NoteTaskShapes", &diagramhierarchy_.NoteTaskShapes)

		case "NoteResourceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteResourceShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteResourceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteResourceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteResourceShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.NoteResourceShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "NoteResourceShapes", &diagramhierarchy_.NoteResourceShapes)

		case "Resource_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.Resource_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "Resource_Shapes", &diagramhierarchy_.Resource_Shapes)

		case "ResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.ResourcesWhoseNodeIsExpanded = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "ResourcesWhoseNodeIsExpanded", &diagramhierarchy_.ResourcesWhoseNodeIsExpanded)

		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsResourcesNodeExpanded), formDiv)
		case "ResourceComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceCompositionShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.ResourceComposition_Shapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "ResourceComposition_Shapes", &diagramhierarchy_.ResourceComposition_Shapes)

		case "ResourceTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceTaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceTaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramhierarchyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceTaskShape](diagramhierarchyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramhierarchy_.ResourceTaskShapes = instanceSlice
			diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(diagramhierarchy_, "ResourceTaskShapes", &diagramhierarchy_.ResourceTaskShapes)

		case "IsTimeDiagram":
			FormDivBasicFieldToField(&(diagramhierarchy_.IsTimeDiagram), formDiv)
		case "ComputedStart":
			FormDivBasicFieldToField(&(diagramhierarchy_.ComputedStart), formDiv)
		case "ComputedEnd":
			FormDivBasicFieldToField(&(diagramhierarchy_.ComputedEnd), formDiv)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(diagramhierarchy_.ComputedDuration), formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&(diagramhierarchy_.UseManualStartAndEndDates), formDiv)
		case "ManualStart":
			FormDivBasicFieldToField(&(diagramhierarchy_.ManualStart), formDiv)
		case "ManualEnd":
			FormDivBasicFieldToField(&(diagramhierarchy_.ManualEnd), formDiv)
		case "TimeStep":
			FormDivBasicFieldToField(&(diagramhierarchy_.TimeStep), formDiv)
		case "TimeStepScale":
			FormDivEnumStringFieldToField(&(diagramhierarchy_.TimeStepScale), formDiv)
		case "LaneHeight":
			FormDivBasicFieldToField(&(diagramhierarchy_.LaneHeight), formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&(diagramhierarchy_.RatioBarToLaneHeight), formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&(diagramhierarchy_.YTopMargin), formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&(diagramhierarchy_.XLeftText), formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&(diagramhierarchy_.TextHeight), formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&(diagramhierarchy_.XLeftLanes), formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&(diagramhierarchy_.XRightMargin), formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&(diagramhierarchy_.ArrowLengthToTheRightOfStartBar), formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&(diagramhierarchy_.ArrowTipLenght), formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&(diagramhierarchy_.TimeLine_Color), formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&(diagramhierarchy_.TimeLine_FillOpacity), formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&(diagramhierarchy_.TimeLine_Stroke), formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&(diagramhierarchy_.TimeLine_StrokeWidth), formDiv)
		case "DrawVerticalTimeLines":
			FormDivBasicFieldToField(&(diagramhierarchy_.DrawVerticalTimeLines), formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&(diagramhierarchy_.Group_Stroke), formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&(diagramhierarchy_.Group_StrokeWidth), formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&(diagramhierarchy_.Group_StrokeDashArray), formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&(diagramhierarchy_.DateYOffset), formDiv)
		case "AlignOnStartEndOnYearStart":
			FormDivBasicFieldToField(&(diagramhierarchy_.AlignOnStartEndOnYearStart), formDiv)
		case "Library:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](diagramhierarchyFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Diagrams slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](diagramhierarchyFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramhierarchyFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure diagramhierarchy_ is in _library.Diagrams
					found := false
					for _, _b := range _library.Diagrams {
						if _b == diagramhierarchy_ {
							found = true
							break
						}
					}
					if !found {
						_library.Diagrams = append(_library.Diagrams, diagramhierarchy_)
						diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				} else {
					// ensure diagramhierarchy_ is NOT in _library.Diagrams
					idx := slices.Index(_library.Diagrams, diagramhierarchy_)
					if idx != -1 {
						_library.Diagrams = slices.Delete(_library.Diagrams, idx, idx+1)
						diagramhierarchyFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramhierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramhierarchy_.Unstage(diagramhierarchyFormCallback.probe.stageOfInterest)
	}

	diagramhierarchyFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramHierarchy](
		diagramhierarchyFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramhierarchyFormCallback.CreationMode || diagramhierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramhierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramhierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramHierarchyFormCallback(
			nil,
			diagramhierarchyFormCallback.probe,
			newFormGroup,
		)
		diagramhierarchy := new(models.DiagramHierarchy)
		FillUpForm(diagramhierarchy, newFormGroup, diagramhierarchyFormCallback.probe)
		diagramhierarchyFormCallback.probe.formStage.Commit()
	}

	diagramhierarchyFormCallback.probe.ux_tree()
}
func __gong__New__LibraryFormCallback(
	library *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *LibraryFormCallback) {
	libraryFormCallback = new(LibraryFormCallback)
	libraryFormCallback.probe = probe
	libraryFormCallback.library = library
	libraryFormCallback.formGroup = formGroup

	libraryFormCallback.CreationMode = (library == nil)

	return
}

type LibraryFormCallback struct {
	library *models.Library

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (libraryFormCallback *LibraryFormCallback) OnSave() {
	libraryFormCallback.probe.stageOfInterest.Lock()
	defer libraryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LibraryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	libraryFormCallback.probe.formStage.Checkout()

	if libraryFormCallback.library == nil {
		libraryFormCallback.library = new(models.Library).Stage(libraryFormCallback.probe.stageOfInterest)
	}
	library_ := libraryFormCallback.library
	_ = library_

	for _, formDiv := range libraryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(library_.Name), formDiv)
		case "SubLibraries":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibraries = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibraries", &library_.SubLibraries)

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "RootProducts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootProducts = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootProducts", &library_.RootProducts)

		case "RootTasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootTasks = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootTasks", &library_.RootTasks)

		case "RootTaskGroups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskGroup](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskGroup](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootTaskGroups = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootTaskGroups", &library_.RootTaskGroups)

		case "RootMilestones":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Milestone, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Milestone)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Milestone](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootMilestones = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootMilestones", &library_.RootMilestones)

		case "RootResources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootResources = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootResources", &library_.RootResources)

		case "Notes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Notes = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Notes", &library_.Notes)

		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramHierarchy, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramHierarchy)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Diagrams = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Diagrams", &library_.Diagrams)

		case "Library:SubLibraries":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibraries slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibraries
					found := false
					for _, _b := range _library.SubLibraries {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibraries = append(_library.SubLibraries, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibraries
					idx := slices.Index(_library.SubLibraries, library_)
					if idx != -1 {
						_library.SubLibraries = slices.Delete(_library.SubLibraries, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		library_.Unstage(libraryFormCallback.probe.stageOfInterest)
	}

	libraryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Library](
		libraryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if libraryFormCallback.CreationMode || libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		libraryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(libraryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			libraryFormCallback.probe,
			newFormGroup,
		)
		library := new(models.Library)
		FillUpForm(library, newFormGroup, libraryFormCallback.probe)
		libraryFormCallback.probe.formStage.Commit()
	}

	libraryFormCallback.probe.ux_tree()
}
func __gong__New__MilestoneFormCallback(
	milestone *models.Milestone,
	probe *Probe,
	formGroup *form.FormGroup,
) (milestoneFormCallback *MilestoneFormCallback) {
	milestoneFormCallback = new(MilestoneFormCallback)
	milestoneFormCallback.probe = probe
	milestoneFormCallback.milestone = milestone
	milestoneFormCallback.formGroup = formGroup

	milestoneFormCallback.CreationMode = (milestone == nil)

	return
}

type MilestoneFormCallback struct {
	milestone *models.Milestone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (milestoneFormCallback *MilestoneFormCallback) OnSave() {
	milestoneFormCallback.probe.stageOfInterest.Lock()
	defer milestoneFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MilestoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	milestoneFormCallback.probe.formStage.Checkout()

	if milestoneFormCallback.milestone == nil {
		milestoneFormCallback.milestone = new(models.Milestone).Stage(milestoneFormCallback.probe.stageOfInterest)
	}
	milestone_ := milestoneFormCallback.milestone
	_ = milestone_

	for _, formDiv := range milestoneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(milestone_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(milestone_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(milestone_.IsExpanded), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(milestone_.Date), formDiv)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&(milestone_.DisplayVerticalBar), formDiv)
		case "TaskGroupsToDisplay":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskGroup](milestoneFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					milestoneFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskGroup](milestoneFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			milestone_.TaskGroupsToDisplay = instanceSlice
			milestoneFormCallback.probe.UpdateSliceOfPointersCallback(milestone_, "TaskGroupsToDisplay", &milestone_.TaskGroupsToDisplay)

		case "DiagramHierarchy:MilestonesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](milestoneFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their MilestonesWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](milestoneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(milestoneFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure milestone_ is in _diagramhierarchy.MilestonesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.MilestonesWhoseNodeIsExpanded {
						if _b == milestone_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.MilestonesWhoseNodeIsExpanded = append(_diagramhierarchy.MilestonesWhoseNodeIsExpanded, milestone_)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "MilestonesWhoseNodeIsExpanded", &_diagramhierarchy.MilestonesWhoseNodeIsExpanded)
					}
				} else {
					// ensure milestone_ is NOT in _diagramhierarchy.MilestonesWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.MilestonesWhoseNodeIsExpanded, milestone_)
					if idx != -1 {
						_diagramhierarchy.MilestonesWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.MilestonesWhoseNodeIsExpanded, idx, idx+1)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "MilestonesWhoseNodeIsExpanded", &_diagramhierarchy.MilestonesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootMilestones":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](milestoneFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootMilestones slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](milestoneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(milestoneFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure milestone_ is in _library.RootMilestones
					found := false
					for _, _b := range _library.RootMilestones {
						if _b == milestone_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootMilestones = append(_library.RootMilestones, milestone_)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootMilestones", &_library.RootMilestones)
					}
				} else {
					// ensure milestone_ is NOT in _library.RootMilestones
					idx := slices.Index(_library.RootMilestones, milestone_)
					if idx != -1 {
						_library.RootMilestones = slices.Delete(_library.RootMilestones, idx, idx+1)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootMilestones", &_library.RootMilestones)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestone_.Unstage(milestoneFormCallback.probe.stageOfInterest)
	}

	milestoneFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Milestone](
		milestoneFormCallback.probe,
	)

	// display a new form by reset the form stage
	if milestoneFormCallback.CreationMode || milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestoneFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(milestoneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MilestoneFormCallback(
			nil,
			milestoneFormCallback.probe,
			newFormGroup,
		)
		milestone := new(models.Milestone)
		FillUpForm(milestone, newFormGroup, milestoneFormCallback.probe)
		milestoneFormCallback.probe.formStage.Commit()
	}

	milestoneFormCallback.probe.ux_tree()
}
func __gong__New__MilestoneShapeFormCallback(
	milestoneshape *models.MilestoneShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (milestoneshapeFormCallback *MilestoneShapeFormCallback) {
	milestoneshapeFormCallback = new(MilestoneShapeFormCallback)
	milestoneshapeFormCallback.probe = probe
	milestoneshapeFormCallback.milestoneshape = milestoneshape
	milestoneshapeFormCallback.formGroup = formGroup

	milestoneshapeFormCallback.CreationMode = (milestoneshape == nil)

	return
}

type MilestoneShapeFormCallback struct {
	milestoneshape *models.MilestoneShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (milestoneshapeFormCallback *MilestoneShapeFormCallback) OnSave() {
	milestoneshapeFormCallback.probe.stageOfInterest.Lock()
	defer milestoneshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MilestoneShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	milestoneshapeFormCallback.probe.formStage.Checkout()

	if milestoneshapeFormCallback.milestoneshape == nil {
		milestoneshapeFormCallback.milestoneshape = new(models.MilestoneShape).Stage(milestoneshapeFormCallback.probe.stageOfInterest)
	}
	milestoneshape_ := milestoneshapeFormCallback.milestoneshape
	_ = milestoneshape_

	for _, formDiv := range milestoneshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(milestoneshape_.Name), formDiv)
		case "Milestone":
			FormDivSelectFieldToField(&(milestoneshape_.Milestone), milestoneshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(milestoneshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(milestoneshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(milestoneshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(milestoneshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(milestoneshape_.IsHidden), formDiv)
		case "DiagramHierarchy:MilestoneShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](milestoneshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their MilestoneShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](milestoneshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(milestoneshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure milestoneshape_ is in _diagramhierarchy.MilestoneShapes
					found := false
					for _, _b := range _diagramhierarchy.MilestoneShapes {
						if _b == milestoneshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.MilestoneShapes = append(_diagramhierarchy.MilestoneShapes, milestoneshape_)
						milestoneshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "MilestoneShapes", &_diagramhierarchy.MilestoneShapes)
					}
				} else {
					// ensure milestoneshape_ is NOT in _diagramhierarchy.MilestoneShapes
					idx := slices.Index(_diagramhierarchy.MilestoneShapes, milestoneshape_)
					if idx != -1 {
						_diagramhierarchy.MilestoneShapes = slices.Delete(_diagramhierarchy.MilestoneShapes, idx, idx+1)
						milestoneshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "MilestoneShapes", &_diagramhierarchy.MilestoneShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if milestoneshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestoneshape_.Unstage(milestoneshapeFormCallback.probe.stageOfInterest)
	}

	milestoneshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MilestoneShape](
		milestoneshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if milestoneshapeFormCallback.CreationMode || milestoneshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestoneshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(milestoneshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MilestoneShapeFormCallback(
			nil,
			milestoneshapeFormCallback.probe,
			newFormGroup,
		)
		milestoneshape := new(models.MilestoneShape)
		FillUpForm(milestoneshape, newFormGroup, milestoneshapeFormCallback.probe)
		milestoneshapeFormCallback.probe.formStage.Commit()
	}

	milestoneshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *NoteFormCallback) {
	noteFormCallback = new(NoteFormCallback)
	noteFormCallback.probe = probe
	noteFormCallback.note = note
	noteFormCallback.formGroup = formGroup

	noteFormCallback.CreationMode = (note == nil)

	return
}

type NoteFormCallback struct {
	note *models.Note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteFormCallback *NoteFormCallback) OnSave() {
	noteFormCallback.probe.stageOfInterest.Lock()
	defer noteFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteFormCallback.probe.formStage.Checkout()

	if noteFormCallback.note == nil {
		noteFormCallback.note = new(models.Note).Stage(noteFormCallback.probe.stageOfInterest)
	}
	note_ := noteFormCallback.note
	_ = note_

	for _, formDiv := range noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "Products":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Products = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Products", &note_.Products)

		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Tasks = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Tasks", &note_.Tasks)

		case "Resources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Resources = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Resources", &note_.Resources)

		case "DiagramHierarchy:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](noteFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their NotesWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure note_ is in _diagramhierarchy.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.NotesWhoseNodeIsExpanded = append(_diagramhierarchy.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NotesWhoseNodeIsExpanded", &_diagramhierarchy.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagramhierarchy.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagramhierarchy.NotesWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NotesWhoseNodeIsExpanded", &_diagramhierarchy.NotesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:Notes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](noteFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Notes slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.Notes
					found := false
					for _, _b := range _library.Notes {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.Notes = append(_library.Notes, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Notes", &_library.Notes)
					}
				} else {
					// ensure note_ is NOT in _library.Notes
					idx := slices.Index(_library.Notes, note_)
					if idx != -1 {
						_library.Notes = slices.Delete(_library.Notes, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Notes", &_library.Notes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_.Unstage(noteFormCallback.probe.stageOfInterest)
	}

	noteFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Note](
		noteFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteFormCallback.CreationMode || noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			noteFormCallback.probe,
			newFormGroup,
		)
		note := new(models.Note)
		FillUpForm(note, newFormGroup, noteFormCallback.probe)
		noteFormCallback.probe.formStage.Commit()
	}

	noteFormCallback.probe.ux_tree()
}
func __gong__New__NoteProductShapeFormCallback(
	noteproductshape *models.NoteProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteproductshapeFormCallback *NoteProductShapeFormCallback) {
	noteproductshapeFormCallback = new(NoteProductShapeFormCallback)
	noteproductshapeFormCallback.probe = probe
	noteproductshapeFormCallback.noteproductshape = noteproductshape
	noteproductshapeFormCallback.formGroup = formGroup

	noteproductshapeFormCallback.CreationMode = (noteproductshape == nil)

	return
}

type NoteProductShapeFormCallback struct {
	noteproductshape *models.NoteProductShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteproductshapeFormCallback *NoteProductShapeFormCallback) OnSave() {
	noteproductshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteproductshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteProductShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteproductshapeFormCallback.probe.formStage.Checkout()

	if noteproductshapeFormCallback.noteproductshape == nil {
		noteproductshapeFormCallback.noteproductshape = new(models.NoteProductShape).Stage(noteproductshapeFormCallback.probe.stageOfInterest)
	}
	noteproductshape_ := noteproductshapeFormCallback.noteproductshape
	_ = noteproductshape_

	for _, formDiv := range noteproductshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteproductshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteproductshape_.Note), noteproductshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(noteproductshape_.Product), noteproductshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteproductshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteproductshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteproductshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteproductshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteproductshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteproductshape_.IsHidden), formDiv)
		case "DiagramHierarchy:NoteProductShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](noteproductshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their NoteProductShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](noteproductshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteproductshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure noteproductshape_ is in _diagramhierarchy.NoteProductShapes
					found := false
					for _, _b := range _diagramhierarchy.NoteProductShapes {
						if _b == noteproductshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.NoteProductShapes = append(_diagramhierarchy.NoteProductShapes, noteproductshape_)
						noteproductshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteProductShapes", &_diagramhierarchy.NoteProductShapes)
					}
				} else {
					// ensure noteproductshape_ is NOT in _diagramhierarchy.NoteProductShapes
					idx := slices.Index(_diagramhierarchy.NoteProductShapes, noteproductshape_)
					if idx != -1 {
						_diagramhierarchy.NoteProductShapes = slices.Delete(_diagramhierarchy.NoteProductShapes, idx, idx+1)
						noteproductshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteProductShapes", &_diagramhierarchy.NoteProductShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteproductshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteproductshape_.Unstage(noteproductshapeFormCallback.probe.stageOfInterest)
	}

	noteproductshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteProductShape](
		noteproductshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteproductshapeFormCallback.CreationMode || noteproductshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteproductshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteproductshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteProductShapeFormCallback(
			nil,
			noteproductshapeFormCallback.probe,
			newFormGroup,
		)
		noteproductshape := new(models.NoteProductShape)
		FillUpForm(noteproductshape, newFormGroup, noteproductshapeFormCallback.probe)
		noteproductshapeFormCallback.probe.formStage.Commit()
	}

	noteproductshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteResourceShapeFormCallback(
	noteresourceshape *models.NoteResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteresourceshapeFormCallback *NoteResourceShapeFormCallback) {
	noteresourceshapeFormCallback = new(NoteResourceShapeFormCallback)
	noteresourceshapeFormCallback.probe = probe
	noteresourceshapeFormCallback.noteresourceshape = noteresourceshape
	noteresourceshapeFormCallback.formGroup = formGroup

	noteresourceshapeFormCallback.CreationMode = (noteresourceshape == nil)

	return
}

type NoteResourceShapeFormCallback struct {
	noteresourceshape *models.NoteResourceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteresourceshapeFormCallback *NoteResourceShapeFormCallback) OnSave() {
	noteresourceshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteresourceshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteResourceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteresourceshapeFormCallback.probe.formStage.Checkout()

	if noteresourceshapeFormCallback.noteresourceshape == nil {
		noteresourceshapeFormCallback.noteresourceshape = new(models.NoteResourceShape).Stage(noteresourceshapeFormCallback.probe.stageOfInterest)
	}
	noteresourceshape_ := noteresourceshapeFormCallback.noteresourceshape
	_ = noteresourceshape_

	for _, formDiv := range noteresourceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteresourceshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteresourceshape_.Note), noteresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(noteresourceshape_.Resource), noteresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteresourceshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteresourceshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteresourceshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteresourceshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteresourceshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteresourceshape_.IsHidden), formDiv)
		case "DiagramHierarchy:NoteResourceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](noteresourceshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their NoteResourceShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](noteresourceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteresourceshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure noteresourceshape_ is in _diagramhierarchy.NoteResourceShapes
					found := false
					for _, _b := range _diagramhierarchy.NoteResourceShapes {
						if _b == noteresourceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.NoteResourceShapes = append(_diagramhierarchy.NoteResourceShapes, noteresourceshape_)
						noteresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteResourceShapes", &_diagramhierarchy.NoteResourceShapes)
					}
				} else {
					// ensure noteresourceshape_ is NOT in _diagramhierarchy.NoteResourceShapes
					idx := slices.Index(_diagramhierarchy.NoteResourceShapes, noteresourceshape_)
					if idx != -1 {
						_diagramhierarchy.NoteResourceShapes = slices.Delete(_diagramhierarchy.NoteResourceShapes, idx, idx+1)
						noteresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteResourceShapes", &_diagramhierarchy.NoteResourceShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteresourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteresourceshape_.Unstage(noteresourceshapeFormCallback.probe.stageOfInterest)
	}

	noteresourceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteResourceShape](
		noteresourceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteresourceshapeFormCallback.CreationMode || noteresourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteresourceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteresourceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteResourceShapeFormCallback(
			nil,
			noteresourceshapeFormCallback.probe,
			newFormGroup,
		)
		noteresourceshape := new(models.NoteResourceShape)
		FillUpForm(noteresourceshape, newFormGroup, noteresourceshapeFormCallback.probe)
		noteresourceshapeFormCallback.probe.formStage.Commit()
	}

	noteresourceshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.probe = probe
	noteshapeFormCallback.noteshape = noteshape
	noteshapeFormCallback.formGroup = formGroup

	noteshapeFormCallback.CreationMode = (noteshape == nil)

	return
}

type NoteShapeFormCallback struct {
	noteshape *models.NoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {
	noteshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.probe.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.probe.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	for _, formDiv := range noteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteshape_.Note), noteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(noteshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteshape_.IsHidden), formDiv)
		case "DiagramHierarchy:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their Note_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure noteshape_ is in _diagramhierarchy.Note_Shapes
					found := false
					for _, _b := range _diagramhierarchy.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.Note_Shapes = append(_diagramhierarchy.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Note_Shapes", &_diagramhierarchy.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagramhierarchy.Note_Shapes
					idx := slices.Index(_diagramhierarchy.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagramhierarchy.Note_Shapes = slices.Delete(_diagramhierarchy.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Note_Shapes", &_diagramhierarchy.Note_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshape_.Unstage(noteshapeFormCallback.probe.stageOfInterest)
	}

	noteshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteShape](
		noteshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteshapeFormCallback.CreationMode || noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			noteshapeFormCallback.probe,
			newFormGroup,
		)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, newFormGroup, noteshapeFormCallback.probe)
		noteshapeFormCallback.probe.formStage.Commit()
	}

	noteshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteTaskShapeFormCallback(
	notetaskshape *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetaskshapeFormCallback *NoteTaskShapeFormCallback) {
	notetaskshapeFormCallback = new(NoteTaskShapeFormCallback)
	notetaskshapeFormCallback.probe = probe
	notetaskshapeFormCallback.notetaskshape = notetaskshape
	notetaskshapeFormCallback.formGroup = formGroup

	notetaskshapeFormCallback.CreationMode = (notetaskshape == nil)

	return
}

type NoteTaskShapeFormCallback struct {
	notetaskshape *models.NoteTaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notetaskshapeFormCallback *NoteTaskShapeFormCallback) OnSave() {
	notetaskshapeFormCallback.probe.stageOfInterest.Lock()
	defer notetaskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteTaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notetaskshapeFormCallback.probe.formStage.Checkout()

	if notetaskshapeFormCallback.notetaskshape == nil {
		notetaskshapeFormCallback.notetaskshape = new(models.NoteTaskShape).Stage(notetaskshapeFormCallback.probe.stageOfInterest)
	}
	notetaskshape_ := notetaskshapeFormCallback.notetaskshape
	_ = notetaskshape_

	for _, formDiv := range notetaskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notetaskshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notetaskshape_.Note), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(notetaskshape_.Task), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notetaskshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notetaskshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notetaskshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notetaskshape_.IsHidden), formDiv)
		case "DiagramHierarchy:NoteTaskShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](notetaskshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their NoteTaskShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](notetaskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notetaskshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure notetaskshape_ is in _diagramhierarchy.NoteTaskShapes
					found := false
					for _, _b := range _diagramhierarchy.NoteTaskShapes {
						if _b == notetaskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.NoteTaskShapes = append(_diagramhierarchy.NoteTaskShapes, notetaskshape_)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteTaskShapes", &_diagramhierarchy.NoteTaskShapes)
					}
				} else {
					// ensure notetaskshape_ is NOT in _diagramhierarchy.NoteTaskShapes
					idx := slices.Index(_diagramhierarchy.NoteTaskShapes, notetaskshape_)
					if idx != -1 {
						_diagramhierarchy.NoteTaskShapes = slices.Delete(_diagramhierarchy.NoteTaskShapes, idx, idx+1)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "NoteTaskShapes", &_diagramhierarchy.NoteTaskShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshape_.Unstage(notetaskshapeFormCallback.probe.stageOfInterest)
	}

	notetaskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteTaskShape](
		notetaskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notetaskshapeFormCallback.CreationMode || notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notetaskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			nil,
			notetaskshapeFormCallback.probe,
			newFormGroup,
		)
		notetaskshape := new(models.NoteTaskShape)
		FillUpForm(notetaskshape, newFormGroup, notetaskshapeFormCallback.probe)
		notetaskshapeFormCallback.probe.formStage.Commit()
	}

	notetaskshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProductFormCallback(
	product *models.Product,
	probe *Probe,
	formGroup *form.FormGroup,
) (productFormCallback *ProductFormCallback) {
	productFormCallback = new(ProductFormCallback)
	productFormCallback.probe = probe
	productFormCallback.product = product
	productFormCallback.formGroup = formGroup

	productFormCallback.CreationMode = (product == nil)

	return
}

type ProductFormCallback struct {
	product *models.Product

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (productFormCallback *ProductFormCallback) OnSave() {
	productFormCallback.probe.stageOfInterest.Lock()
	defer productFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProductFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	productFormCallback.probe.formStage.Checkout()

	if productFormCallback.product == nil {
		productFormCallback.product = new(models.Product).Stage(productFormCallback.probe.stageOfInterest)
	}
	product_ := productFormCallback.product
	_ = product_

	for _, formDiv := range productFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(product_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(product_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(product_.IsExpanded), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(product_.IsImport), formDiv)
		case "ReferencedProduct":
			FormDivSelectFieldToField(&(product_.ReferencedProduct), productFormCallback.probe.stageOfInterest, formDiv)
		case "Description":
			FormDivBasicFieldToField(&(product_.Description), formDiv)
		case "SubProducts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](productFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					productFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](productFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			product_.SubProducts = instanceSlice
			productFormCallback.probe.UpdateSliceOfPointersCallback(product_, "SubProducts", &product_.SubProducts)

		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&(product_.IsProducersNodeExpanded), formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&(product_.IsConsumersNodeExpanded), formDiv)
		case "DiagramHierarchy:ProductsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](productFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their ProductsWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure product_ is in _diagramhierarchy.ProductsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.ProductsWhoseNodeIsExpanded {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.ProductsWhoseNodeIsExpanded = append(_diagramhierarchy.ProductsWhoseNodeIsExpanded, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ProductsWhoseNodeIsExpanded", &_diagramhierarchy.ProductsWhoseNodeIsExpanded)
					}
				} else {
					// ensure product_ is NOT in _diagramhierarchy.ProductsWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.ProductsWhoseNodeIsExpanded, product_)
					if idx != -1 {
						_diagramhierarchy.ProductsWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.ProductsWhoseNodeIsExpanded, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ProductsWhoseNodeIsExpanded", &_diagramhierarchy.ProductsWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootProducts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](productFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootProducts slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure product_ is in _library.RootProducts
					found := false
					for _, _b := range _library.RootProducts {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootProducts = append(_library.RootProducts, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootProducts", &_library.RootProducts)
					}
				} else {
					// ensure product_ is NOT in _library.RootProducts
					idx := slices.Index(_library.RootProducts, product_)
					if idx != -1 {
						_library.RootProducts = slices.Delete(_library.RootProducts, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootProducts", &_library.RootProducts)
					}
				}
			}
		case "Note:Products":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](productFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Products slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure product_ is in _note.Products
					found := false
					for _, _b := range _note.Products {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_note.Products = append(_note.Products, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Products", &_note.Products)
					}
				} else {
					// ensure product_ is NOT in _note.Products
					idx := slices.Index(_note.Products, product_)
					if idx != -1 {
						_note.Products = slices.Delete(_note.Products, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Products", &_note.Products)
					}
				}
			}
		case "Product:SubProducts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Product instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Product instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Product](productFormCallback.probe.stageOfInterest)
			targetProductIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProductIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Product instances and update their SubProducts slice
			for _product := range *models.GetGongstructInstancesSetFromPointerType[*models.Product](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _product)
				
				// if Product is selected
				if targetProductIDs[id] {
					// ensure product_ is in _product.SubProducts
					found := false
					for _, _b := range _product.SubProducts {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_product.SubProducts = append(_product.SubProducts, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_product, "SubProducts", &_product.SubProducts)
					}
				} else {
					// ensure product_ is NOT in _product.SubProducts
					idx := slices.Index(_product.SubProducts, product_)
					if idx != -1 {
						_product.SubProducts = slices.Delete(_product.SubProducts, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_product, "SubProducts", &_product.SubProducts)
					}
				}
			}
		case "Task:Inputs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Task instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Task instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Task](productFormCallback.probe.stageOfInterest)
			targetTaskIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTaskIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Task instances and update their Inputs slice
			for _task := range *models.GetGongstructInstancesSetFromPointerType[*models.Task](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _task)
				
				// if Task is selected
				if targetTaskIDs[id] {
					// ensure product_ is in _task.Inputs
					found := false
					for _, _b := range _task.Inputs {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_task.Inputs = append(_task.Inputs, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_task, "Inputs", &_task.Inputs)
					}
				} else {
					// ensure product_ is NOT in _task.Inputs
					idx := slices.Index(_task.Inputs, product_)
					if idx != -1 {
						_task.Inputs = slices.Delete(_task.Inputs, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_task, "Inputs", &_task.Inputs)
					}
				}
			}
		case "Task:Outputs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Task instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Task instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Task](productFormCallback.probe.stageOfInterest)
			targetTaskIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTaskIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Task instances and update their Outputs slice
			for _task := range *models.GetGongstructInstancesSetFromPointerType[*models.Task](productFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productFormCallback.probe.stageOfInterest, _task)
				
				// if Task is selected
				if targetTaskIDs[id] {
					// ensure product_ is in _task.Outputs
					found := false
					for _, _b := range _task.Outputs {
						if _b == product_ {
							found = true
							break
						}
					}
					if !found {
						_task.Outputs = append(_task.Outputs, product_)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_task, "Outputs", &_task.Outputs)
					}
				} else {
					// ensure product_ is NOT in _task.Outputs
					idx := slices.Index(_task.Outputs, product_)
					if idx != -1 {
						_task.Outputs = slices.Delete(_task.Outputs, idx, idx+1)
						productFormCallback.probe.UpdateSliceOfPointersCallback(_task, "Outputs", &_task.Outputs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if productFormCallback.formGroup.HasSuppressButtonBeenPressed {
		product_.Unstage(productFormCallback.probe.stageOfInterest)
	}

	productFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Product](
		productFormCallback.probe,
	)

	// display a new form by reset the form stage
	if productFormCallback.CreationMode || productFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(productFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProductFormCallback(
			nil,
			productFormCallback.probe,
			newFormGroup,
		)
		product := new(models.Product)
		FillUpForm(product, newFormGroup, productFormCallback.probe)
		productFormCallback.probe.formStage.Commit()
	}

	productFormCallback.probe.ux_tree()
}
func __gong__New__ProductCompositionShapeFormCallback(
	productcompositionshape *models.ProductCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productcompositionshapeFormCallback *ProductCompositionShapeFormCallback) {
	productcompositionshapeFormCallback = new(ProductCompositionShapeFormCallback)
	productcompositionshapeFormCallback.probe = probe
	productcompositionshapeFormCallback.productcompositionshape = productcompositionshape
	productcompositionshapeFormCallback.formGroup = formGroup

	productcompositionshapeFormCallback.CreationMode = (productcompositionshape == nil)

	return
}

type ProductCompositionShapeFormCallback struct {
	productcompositionshape *models.ProductCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (productcompositionshapeFormCallback *ProductCompositionShapeFormCallback) OnSave() {
	productcompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer productcompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProductCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	productcompositionshapeFormCallback.probe.formStage.Checkout()

	if productcompositionshapeFormCallback.productcompositionshape == nil {
		productcompositionshapeFormCallback.productcompositionshape = new(models.ProductCompositionShape).Stage(productcompositionshapeFormCallback.probe.stageOfInterest)
	}
	productcompositionshape_ := productcompositionshapeFormCallback.productcompositionshape
	_ = productcompositionshape_

	for _, formDiv := range productcompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(productcompositionshape_.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(productcompositionshape_.Product), productcompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(productcompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(productcompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(productcompositionshape_.IsHidden), formDiv)
		case "DiagramHierarchy:ProductComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](productcompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their ProductComposition_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](productcompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productcompositionshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure productcompositionshape_ is in _diagramhierarchy.ProductComposition_Shapes
					found := false
					for _, _b := range _diagramhierarchy.ProductComposition_Shapes {
						if _b == productcompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.ProductComposition_Shapes = append(_diagramhierarchy.ProductComposition_Shapes, productcompositionshape_)
						productcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ProductComposition_Shapes", &_diagramhierarchy.ProductComposition_Shapes)
					}
				} else {
					// ensure productcompositionshape_ is NOT in _diagramhierarchy.ProductComposition_Shapes
					idx := slices.Index(_diagramhierarchy.ProductComposition_Shapes, productcompositionshape_)
					if idx != -1 {
						_diagramhierarchy.ProductComposition_Shapes = slices.Delete(_diagramhierarchy.ProductComposition_Shapes, idx, idx+1)
						productcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ProductComposition_Shapes", &_diagramhierarchy.ProductComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if productcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productcompositionshape_.Unstage(productcompositionshapeFormCallback.probe.stageOfInterest)
	}

	productcompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProductCompositionShape](
		productcompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if productcompositionshapeFormCallback.CreationMode || productcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productcompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(productcompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProductCompositionShapeFormCallback(
			nil,
			productcompositionshapeFormCallback.probe,
			newFormGroup,
		)
		productcompositionshape := new(models.ProductCompositionShape)
		FillUpForm(productcompositionshape, newFormGroup, productcompositionshapeFormCallback.probe)
		productcompositionshapeFormCallback.probe.formStage.Commit()
	}

	productcompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProductShapeFormCallback(
	productshape *models.ProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productshapeFormCallback *ProductShapeFormCallback) {
	productshapeFormCallback = new(ProductShapeFormCallback)
	productshapeFormCallback.probe = probe
	productshapeFormCallback.productshape = productshape
	productshapeFormCallback.formGroup = formGroup

	productshapeFormCallback.CreationMode = (productshape == nil)

	return
}

type ProductShapeFormCallback struct {
	productshape *models.ProductShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (productshapeFormCallback *ProductShapeFormCallback) OnSave() {
	productshapeFormCallback.probe.stageOfInterest.Lock()
	defer productshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProductShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	productshapeFormCallback.probe.formStage.Checkout()

	if productshapeFormCallback.productshape == nil {
		productshapeFormCallback.productshape = new(models.ProductShape).Stage(productshapeFormCallback.probe.stageOfInterest)
	}
	productshape_ := productshapeFormCallback.productshape
	_ = productshape_

	for _, formDiv := range productshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(productshape_.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(productshape_.Product), productshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(productshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(productshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(productshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(productshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(productshape_.IsHidden), formDiv)
		case "DiagramHierarchy:Product_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](productshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their Product_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](productshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure productshape_ is in _diagramhierarchy.Product_Shapes
					found := false
					for _, _b := range _diagramhierarchy.Product_Shapes {
						if _b == productshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.Product_Shapes = append(_diagramhierarchy.Product_Shapes, productshape_)
						productshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Product_Shapes", &_diagramhierarchy.Product_Shapes)
					}
				} else {
					// ensure productshape_ is NOT in _diagramhierarchy.Product_Shapes
					idx := slices.Index(_diagramhierarchy.Product_Shapes, productshape_)
					if idx != -1 {
						_diagramhierarchy.Product_Shapes = slices.Delete(_diagramhierarchy.Product_Shapes, idx, idx+1)
						productshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Product_Shapes", &_diagramhierarchy.Product_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if productshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productshape_.Unstage(productshapeFormCallback.probe.stageOfInterest)
	}

	productshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProductShape](
		productshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if productshapeFormCallback.CreationMode || productshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(productshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProductShapeFormCallback(
			nil,
			productshapeFormCallback.probe,
			newFormGroup,
		)
		productshape := new(models.ProductShape)
		FillUpForm(productshape, newFormGroup, productshapeFormCallback.probe)
		productshapeFormCallback.probe.formStage.Commit()
	}

	productshapeFormCallback.probe.ux_tree()
}
func __gong__New__ResourceFormCallback(
	resource *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceFormCallback *ResourceFormCallback) {
	resourceFormCallback = new(ResourceFormCallback)
	resourceFormCallback.probe = probe
	resourceFormCallback.resource = resource
	resourceFormCallback.formGroup = formGroup

	resourceFormCallback.CreationMode = (resource == nil)

	return
}

type ResourceFormCallback struct {
	resource *models.Resource

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (resourceFormCallback *ResourceFormCallback) OnSave() {
	resourceFormCallback.probe.stageOfInterest.Lock()
	defer resourceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ResourceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	resourceFormCallback.probe.formStage.Checkout()

	if resourceFormCallback.resource == nil {
		resourceFormCallback.resource = new(models.Resource).Stage(resourceFormCallback.probe.stageOfInterest)
	}
	resource_ := resourceFormCallback.resource
	_ = resource_

	for _, formDiv := range resourceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(resource_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(resource_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(resource_.IsExpanded), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(resource_.IsImport), formDiv)
		case "ReferencedResource":
			FormDivSelectFieldToField(&(resource_.ReferencedResource), resourceFormCallback.probe.stageOfInterest, formDiv)
		case "Description":
			FormDivBasicFieldToField(&(resource_.Description), formDiv)
		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](resourceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					resourceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](resourceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			resource_.Tasks = instanceSlice
			resourceFormCallback.probe.UpdateSliceOfPointersCallback(resource_, "Tasks", &resource_.Tasks)

		case "SubResources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](resourceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					resourceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](resourceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			resource_.SubResources = instanceSlice
			resourceFormCallback.probe.UpdateSliceOfPointersCallback(resource_, "SubResources", &resource_.SubResources)

		case "DiagramHierarchy:ResourcesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](resourceFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their ResourcesWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure resource_ is in _diagramhierarchy.ResourcesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.ResourcesWhoseNodeIsExpanded {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.ResourcesWhoseNodeIsExpanded = append(_diagramhierarchy.ResourcesWhoseNodeIsExpanded, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourcesWhoseNodeIsExpanded", &_diagramhierarchy.ResourcesWhoseNodeIsExpanded)
					}
				} else {
					// ensure resource_ is NOT in _diagramhierarchy.ResourcesWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.ResourcesWhoseNodeIsExpanded, resource_)
					if idx != -1 {
						_diagramhierarchy.ResourcesWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.ResourcesWhoseNodeIsExpanded, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourcesWhoseNodeIsExpanded", &_diagramhierarchy.ResourcesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootResources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](resourceFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootResources slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure resource_ is in _library.RootResources
					found := false
					for _, _b := range _library.RootResources {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootResources = append(_library.RootResources, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootResources", &_library.RootResources)
					}
				} else {
					// ensure resource_ is NOT in _library.RootResources
					idx := slices.Index(_library.RootResources, resource_)
					if idx != -1 {
						_library.RootResources = slices.Delete(_library.RootResources, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootResources", &_library.RootResources)
					}
				}
			}
		case "Note:Resources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](resourceFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Resources slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure resource_ is in _note.Resources
					found := false
					for _, _b := range _note.Resources {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_note.Resources = append(_note.Resources, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Resources", &_note.Resources)
					}
				} else {
					// ensure resource_ is NOT in _note.Resources
					idx := slices.Index(_note.Resources, resource_)
					if idx != -1 {
						_note.Resources = slices.Delete(_note.Resources, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Resources", &_note.Resources)
					}
				}
			}
		case "Resource:SubResources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Resource instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Resource instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](resourceFormCallback.probe.stageOfInterest)
			targetResourceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetResourceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Resource instances and update their SubResources slice
			for _resource := range *models.GetGongstructInstancesSetFromPointerType[*models.Resource](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _resource)
				
				// if Resource is selected
				if targetResourceIDs[id] {
					// ensure resource_ is in _resource.SubResources
					found := false
					for _, _b := range _resource.SubResources {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_resource.SubResources = append(_resource.SubResources, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_resource, "SubResources", &_resource.SubResources)
					}
				} else {
					// ensure resource_ is NOT in _resource.SubResources
					idx := slices.Index(_resource.SubResources, resource_)
					if idx != -1 {
						_resource.SubResources = slices.Delete(_resource.SubResources, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_resource, "SubResources", &_resource.SubResources)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if resourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resource_.Unstage(resourceFormCallback.probe.stageOfInterest)
	}

	resourceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Resource](
		resourceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if resourceFormCallback.CreationMode || resourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourceFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(resourceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ResourceFormCallback(
			nil,
			resourceFormCallback.probe,
			newFormGroup,
		)
		resource := new(models.Resource)
		FillUpForm(resource, newFormGroup, resourceFormCallback.probe)
		resourceFormCallback.probe.formStage.Commit()
	}

	resourceFormCallback.probe.ux_tree()
}
func __gong__New__ResourceCompositionShapeFormCallback(
	resourcecompositionshape *models.ResourceCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourcecompositionshapeFormCallback *ResourceCompositionShapeFormCallback) {
	resourcecompositionshapeFormCallback = new(ResourceCompositionShapeFormCallback)
	resourcecompositionshapeFormCallback.probe = probe
	resourcecompositionshapeFormCallback.resourcecompositionshape = resourcecompositionshape
	resourcecompositionshapeFormCallback.formGroup = formGroup

	resourcecompositionshapeFormCallback.CreationMode = (resourcecompositionshape == nil)

	return
}

type ResourceCompositionShapeFormCallback struct {
	resourcecompositionshape *models.ResourceCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (resourcecompositionshapeFormCallback *ResourceCompositionShapeFormCallback) OnSave() {
	resourcecompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer resourcecompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ResourceCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	resourcecompositionshapeFormCallback.probe.formStage.Checkout()

	if resourcecompositionshapeFormCallback.resourcecompositionshape == nil {
		resourcecompositionshapeFormCallback.resourcecompositionshape = new(models.ResourceCompositionShape).Stage(resourcecompositionshapeFormCallback.probe.stageOfInterest)
	}
	resourcecompositionshape_ := resourcecompositionshapeFormCallback.resourcecompositionshape
	_ = resourcecompositionshape_

	for _, formDiv := range resourcecompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(resourcecompositionshape_.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(resourcecompositionshape_.Resource), resourcecompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(resourcecompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(resourcecompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(resourcecompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(resourcecompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(resourcecompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(resourcecompositionshape_.IsHidden), formDiv)
		case "DiagramHierarchy:ResourceComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](resourcecompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their ResourceComposition_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](resourcecompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourcecompositionshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure resourcecompositionshape_ is in _diagramhierarchy.ResourceComposition_Shapes
					found := false
					for _, _b := range _diagramhierarchy.ResourceComposition_Shapes {
						if _b == resourcecompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.ResourceComposition_Shapes = append(_diagramhierarchy.ResourceComposition_Shapes, resourcecompositionshape_)
						resourcecompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourceComposition_Shapes", &_diagramhierarchy.ResourceComposition_Shapes)
					}
				} else {
					// ensure resourcecompositionshape_ is NOT in _diagramhierarchy.ResourceComposition_Shapes
					idx := slices.Index(_diagramhierarchy.ResourceComposition_Shapes, resourcecompositionshape_)
					if idx != -1 {
						_diagramhierarchy.ResourceComposition_Shapes = slices.Delete(_diagramhierarchy.ResourceComposition_Shapes, idx, idx+1)
						resourcecompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourceComposition_Shapes", &_diagramhierarchy.ResourceComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if resourcecompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourcecompositionshape_.Unstage(resourcecompositionshapeFormCallback.probe.stageOfInterest)
	}

	resourcecompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ResourceCompositionShape](
		resourcecompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if resourcecompositionshapeFormCallback.CreationMode || resourcecompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourcecompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(resourcecompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ResourceCompositionShapeFormCallback(
			nil,
			resourcecompositionshapeFormCallback.probe,
			newFormGroup,
		)
		resourcecompositionshape := new(models.ResourceCompositionShape)
		FillUpForm(resourcecompositionshape, newFormGroup, resourcecompositionshapeFormCallback.probe)
		resourcecompositionshapeFormCallback.probe.formStage.Commit()
	}

	resourcecompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ResourceShapeFormCallback(
	resourceshape *models.ResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceshapeFormCallback *ResourceShapeFormCallback) {
	resourceshapeFormCallback = new(ResourceShapeFormCallback)
	resourceshapeFormCallback.probe = probe
	resourceshapeFormCallback.resourceshape = resourceshape
	resourceshapeFormCallback.formGroup = formGroup

	resourceshapeFormCallback.CreationMode = (resourceshape == nil)

	return
}

type ResourceShapeFormCallback struct {
	resourceshape *models.ResourceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (resourceshapeFormCallback *ResourceShapeFormCallback) OnSave() {
	resourceshapeFormCallback.probe.stageOfInterest.Lock()
	defer resourceshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ResourceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	resourceshapeFormCallback.probe.formStage.Checkout()

	if resourceshapeFormCallback.resourceshape == nil {
		resourceshapeFormCallback.resourceshape = new(models.ResourceShape).Stage(resourceshapeFormCallback.probe.stageOfInterest)
	}
	resourceshape_ := resourceshapeFormCallback.resourceshape
	_ = resourceshape_

	for _, formDiv := range resourceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(resourceshape_.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(resourceshape_.Resource), resourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(resourceshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(resourceshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(resourceshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(resourceshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(resourceshape_.IsHidden), formDiv)
		case "DiagramHierarchy:Resource_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](resourceshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their Resource_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](resourceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure resourceshape_ is in _diagramhierarchy.Resource_Shapes
					found := false
					for _, _b := range _diagramhierarchy.Resource_Shapes {
						if _b == resourceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.Resource_Shapes = append(_diagramhierarchy.Resource_Shapes, resourceshape_)
						resourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Resource_Shapes", &_diagramhierarchy.Resource_Shapes)
					}
				} else {
					// ensure resourceshape_ is NOT in _diagramhierarchy.Resource_Shapes
					idx := slices.Index(_diagramhierarchy.Resource_Shapes, resourceshape_)
					if idx != -1 {
						_diagramhierarchy.Resource_Shapes = slices.Delete(_diagramhierarchy.Resource_Shapes, idx, idx+1)
						resourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Resource_Shapes", &_diagramhierarchy.Resource_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if resourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourceshape_.Unstage(resourceshapeFormCallback.probe.stageOfInterest)
	}

	resourceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ResourceShape](
		resourceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if resourceshapeFormCallback.CreationMode || resourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(resourceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ResourceShapeFormCallback(
			nil,
			resourceshapeFormCallback.probe,
			newFormGroup,
		)
		resourceshape := new(models.ResourceShape)
		FillUpForm(resourceshape, newFormGroup, resourceshapeFormCallback.probe)
		resourceshapeFormCallback.probe.formStage.Commit()
	}

	resourceshapeFormCallback.probe.ux_tree()
}
func __gong__New__ResourceTaskShapeFormCallback(
	resourcetaskshape *models.ResourceTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourcetaskshapeFormCallback *ResourceTaskShapeFormCallback) {
	resourcetaskshapeFormCallback = new(ResourceTaskShapeFormCallback)
	resourcetaskshapeFormCallback.probe = probe
	resourcetaskshapeFormCallback.resourcetaskshape = resourcetaskshape
	resourcetaskshapeFormCallback.formGroup = formGroup

	resourcetaskshapeFormCallback.CreationMode = (resourcetaskshape == nil)

	return
}

type ResourceTaskShapeFormCallback struct {
	resourcetaskshape *models.ResourceTaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (resourcetaskshapeFormCallback *ResourceTaskShapeFormCallback) OnSave() {
	resourcetaskshapeFormCallback.probe.stageOfInterest.Lock()
	defer resourcetaskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ResourceTaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	resourcetaskshapeFormCallback.probe.formStage.Checkout()

	if resourcetaskshapeFormCallback.resourcetaskshape == nil {
		resourcetaskshapeFormCallback.resourcetaskshape = new(models.ResourceTaskShape).Stage(resourcetaskshapeFormCallback.probe.stageOfInterest)
	}
	resourcetaskshape_ := resourcetaskshapeFormCallback.resourcetaskshape
	_ = resourcetaskshape_

	for _, formDiv := range resourcetaskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(resourcetaskshape_.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(resourcetaskshape_.Resource), resourcetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(resourcetaskshape_.Task), resourcetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(resourcetaskshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(resourcetaskshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(resourcetaskshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(resourcetaskshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(resourcetaskshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(resourcetaskshape_.IsHidden), formDiv)
		case "DiagramHierarchy:ResourceTaskShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](resourcetaskshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their ResourceTaskShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](resourcetaskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourcetaskshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure resourcetaskshape_ is in _diagramhierarchy.ResourceTaskShapes
					found := false
					for _, _b := range _diagramhierarchy.ResourceTaskShapes {
						if _b == resourcetaskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.ResourceTaskShapes = append(_diagramhierarchy.ResourceTaskShapes, resourcetaskshape_)
						resourcetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourceTaskShapes", &_diagramhierarchy.ResourceTaskShapes)
					}
				} else {
					// ensure resourcetaskshape_ is NOT in _diagramhierarchy.ResourceTaskShapes
					idx := slices.Index(_diagramhierarchy.ResourceTaskShapes, resourcetaskshape_)
					if idx != -1 {
						_diagramhierarchy.ResourceTaskShapes = slices.Delete(_diagramhierarchy.ResourceTaskShapes, idx, idx+1)
						resourcetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "ResourceTaskShapes", &_diagramhierarchy.ResourceTaskShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if resourcetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourcetaskshape_.Unstage(resourcetaskshapeFormCallback.probe.stageOfInterest)
	}

	resourcetaskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ResourceTaskShape](
		resourcetaskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if resourcetaskshapeFormCallback.CreationMode || resourcetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourcetaskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(resourcetaskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ResourceTaskShapeFormCallback(
			nil,
			resourcetaskshapeFormCallback.probe,
			newFormGroup,
		)
		resourcetaskshape := new(models.ResourceTaskShape)
		FillUpForm(resourcetaskshape, newFormGroup, resourcetaskshapeFormCallback.probe)
		resourcetaskshapeFormCallback.probe.formStage.Commit()
	}

	resourcetaskshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskFormCallback(
	task *models.Task,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskFormCallback *TaskFormCallback) {
	taskFormCallback = new(TaskFormCallback)
	taskFormCallback.probe = probe
	taskFormCallback.task = task
	taskFormCallback.formGroup = formGroup

	taskFormCallback.CreationMode = (task == nil)

	return
}

type TaskFormCallback struct {
	task *models.Task

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskFormCallback *TaskFormCallback) OnSave() {
	taskFormCallback.probe.stageOfInterest.Lock()
	defer taskFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskFormCallback.probe.formStage.Checkout()

	if taskFormCallback.task == nil {
		taskFormCallback.task = new(models.Task).Stage(taskFormCallback.probe.stageOfInterest)
	}
	task_ := taskFormCallback.task
	_ = task_

	for _, formDiv := range taskFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(task_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(task_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(task_.IsExpanded), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(task_.IsImport), formDiv)
		case "ReferencedTask":
			FormDivSelectFieldToField(&(task_.ReferencedTask), taskFormCallback.probe.stageOfInterest, formDiv)
		case "Start":
			FormDivBasicFieldToField(&(task_.Start), formDiv)
		case "End":
			FormDivBasicFieldToField(&(task_.End), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(task_.Description), formDiv)
		case "SubTasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](taskFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					taskFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](taskFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			task_.SubTasks = instanceSlice
			taskFormCallback.probe.UpdateSliceOfPointersCallback(task_, "SubTasks", &task_.SubTasks)

		case "Inputs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](taskFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					taskFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](taskFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			task_.Inputs = instanceSlice
			taskFormCallback.probe.UpdateSliceOfPointersCallback(task_, "Inputs", &task_.Inputs)

		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&(task_.IsInputsNodeExpanded), formDiv)
		case "Outputs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](taskFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					taskFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](taskFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			task_.Outputs = instanceSlice
			taskFormCallback.probe.UpdateSliceOfPointersCallback(task_, "Outputs", &task_.Outputs)

		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&(task_.IsOutputsNodeExpanded), formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&(task_.IsWithCompletion), formDiv)
		case "Completion":
			FormDivEnumStringFieldToField(&(task_.Completion), formDiv)
		case "DiagramHierarchy:TasksWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TasksWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure task_ is in _diagramhierarchy.TasksWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.TasksWhoseNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TasksWhoseNodeIsExpanded = append(_diagramhierarchy.TasksWhoseNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseNodeIsExpanded", &_diagramhierarchy.TasksWhoseNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _diagramhierarchy.TasksWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.TasksWhoseNodeIsExpanded, task_)
					if idx != -1 {
						_diagramhierarchy.TasksWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.TasksWhoseNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseNodeIsExpanded", &_diagramhierarchy.TasksWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramHierarchy:TasksWhoseInputNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TasksWhoseInputNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure task_ is in _diagramhierarchy.TasksWhoseInputNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.TasksWhoseInputNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TasksWhoseInputNodeIsExpanded = append(_diagramhierarchy.TasksWhoseInputNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseInputNodeIsExpanded", &_diagramhierarchy.TasksWhoseInputNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _diagramhierarchy.TasksWhoseInputNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.TasksWhoseInputNodeIsExpanded, task_)
					if idx != -1 {
						_diagramhierarchy.TasksWhoseInputNodeIsExpanded = slices.Delete(_diagramhierarchy.TasksWhoseInputNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseInputNodeIsExpanded", &_diagramhierarchy.TasksWhoseInputNodeIsExpanded)
					}
				}
			}
		case "DiagramHierarchy:TasksWhoseOutputNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TasksWhoseOutputNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure task_ is in _diagramhierarchy.TasksWhoseOutputNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.TasksWhoseOutputNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TasksWhoseOutputNodeIsExpanded = append(_diagramhierarchy.TasksWhoseOutputNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseOutputNodeIsExpanded", &_diagramhierarchy.TasksWhoseOutputNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _diagramhierarchy.TasksWhoseOutputNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.TasksWhoseOutputNodeIsExpanded, task_)
					if idx != -1 {
						_diagramhierarchy.TasksWhoseOutputNodeIsExpanded = slices.Delete(_diagramhierarchy.TasksWhoseOutputNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TasksWhoseOutputNodeIsExpanded", &_diagramhierarchy.TasksWhoseOutputNodeIsExpanded)
					}
				}
			}
		case "Library:RootTasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](taskFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootTasks slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure task_ is in _library.RootTasks
					found := false
					for _, _b := range _library.RootTasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootTasks = append(_library.RootTasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootTasks", &_library.RootTasks)
					}
				} else {
					// ensure task_ is NOT in _library.RootTasks
					idx := slices.Index(_library.RootTasks, task_)
					if idx != -1 {
						_library.RootTasks = slices.Delete(_library.RootTasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootTasks", &_library.RootTasks)
					}
				}
			}
		case "Note:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](taskFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Tasks slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure task_ is in _note.Tasks
					found := false
					for _, _b := range _note.Tasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_note.Tasks = append(_note.Tasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				} else {
					// ensure task_ is NOT in _note.Tasks
					idx := slices.Index(_note.Tasks, task_)
					if idx != -1 {
						_note.Tasks = slices.Delete(_note.Tasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				}
			}
		case "Resource:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Resource instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Resource instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](taskFormCallback.probe.stageOfInterest)
			targetResourceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetResourceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Resource instances and update their Tasks slice
			for _resource := range *models.GetGongstructInstancesSetFromPointerType[*models.Resource](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _resource)
				
				// if Resource is selected
				if targetResourceIDs[id] {
					// ensure task_ is in _resource.Tasks
					found := false
					for _, _b := range _resource.Tasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_resource.Tasks = append(_resource.Tasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_resource, "Tasks", &_resource.Tasks)
					}
				} else {
					// ensure task_ is NOT in _resource.Tasks
					idx := slices.Index(_resource.Tasks, task_)
					if idx != -1 {
						_resource.Tasks = slices.Delete(_resource.Tasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_resource, "Tasks", &_resource.Tasks)
					}
				}
			}
		case "Task:SubTasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Task instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Task instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Task](taskFormCallback.probe.stageOfInterest)
			targetTaskIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTaskIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Task instances and update their SubTasks slice
			for _task := range *models.GetGongstructInstancesSetFromPointerType[*models.Task](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _task)
				
				// if Task is selected
				if targetTaskIDs[id] {
					// ensure task_ is in _task.SubTasks
					found := false
					for _, _b := range _task.SubTasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_task.SubTasks = append(_task.SubTasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_task, "SubTasks", &_task.SubTasks)
					}
				} else {
					// ensure task_ is NOT in _task.SubTasks
					idx := slices.Index(_task.SubTasks, task_)
					if idx != -1 {
						_task.SubTasks = slices.Delete(_task.SubTasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_task, "SubTasks", &_task.SubTasks)
					}
				}
			}
		case "TaskGroup:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TaskGroup instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TaskGroup instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TaskGroup](taskFormCallback.probe.stageOfInterest)
			targetTaskGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTaskGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TaskGroup instances and update their Tasks slice
			for _taskgroup := range *models.GetGongstructInstancesSetFromPointerType[*models.TaskGroup](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _taskgroup)
				
				// if TaskGroup is selected
				if targetTaskGroupIDs[id] {
					// ensure task_ is in _taskgroup.Tasks
					found := false
					for _, _b := range _taskgroup.Tasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_taskgroup.Tasks = append(_taskgroup.Tasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_taskgroup, "Tasks", &_taskgroup.Tasks)
					}
				} else {
					// ensure task_ is NOT in _taskgroup.Tasks
					idx := slices.Index(_taskgroup.Tasks, task_)
					if idx != -1 {
						_taskgroup.Tasks = slices.Delete(_taskgroup.Tasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_taskgroup, "Tasks", &_taskgroup.Tasks)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		task_.Unstage(taskFormCallback.probe.stageOfInterest)
	}

	taskFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Task](
		taskFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskFormCallback.CreationMode || taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			taskFormCallback.probe,
			newFormGroup,
		)
		task := new(models.Task)
		FillUpForm(task, newFormGroup, taskFormCallback.probe)
		taskFormCallback.probe.formStage.Commit()
	}

	taskFormCallback.probe.ux_tree()
}
func __gong__New__TaskCompositionShapeFormCallback(
	taskcompositionshape *models.TaskCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskcompositionshapeFormCallback *TaskCompositionShapeFormCallback) {
	taskcompositionshapeFormCallback = new(TaskCompositionShapeFormCallback)
	taskcompositionshapeFormCallback.probe = probe
	taskcompositionshapeFormCallback.taskcompositionshape = taskcompositionshape
	taskcompositionshapeFormCallback.formGroup = formGroup

	taskcompositionshapeFormCallback.CreationMode = (taskcompositionshape == nil)

	return
}

type TaskCompositionShapeFormCallback struct {
	taskcompositionshape *models.TaskCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskcompositionshapeFormCallback *TaskCompositionShapeFormCallback) OnSave() {
	taskcompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskcompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskcompositionshapeFormCallback.probe.formStage.Checkout()

	if taskcompositionshapeFormCallback.taskcompositionshape == nil {
		taskcompositionshapeFormCallback.taskcompositionshape = new(models.TaskCompositionShape).Stage(taskcompositionshapeFormCallback.probe.stageOfInterest)
	}
	taskcompositionshape_ := taskcompositionshapeFormCallback.taskcompositionshape
	_ = taskcompositionshape_

	for _, formDiv := range taskcompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskcompositionshape_.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(taskcompositionshape_.Task), taskcompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(taskcompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(taskcompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(taskcompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(taskcompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(taskcompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskcompositionshape_.IsHidden), formDiv)
		case "DiagramHierarchy:TaskComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskcompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TaskComposition_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskcompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskcompositionshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskcompositionshape_ is in _diagramhierarchy.TaskComposition_Shapes
					found := false
					for _, _b := range _diagramhierarchy.TaskComposition_Shapes {
						if _b == taskcompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TaskComposition_Shapes = append(_diagramhierarchy.TaskComposition_Shapes, taskcompositionshape_)
						taskcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskComposition_Shapes", &_diagramhierarchy.TaskComposition_Shapes)
					}
				} else {
					// ensure taskcompositionshape_ is NOT in _diagramhierarchy.TaskComposition_Shapes
					idx := slices.Index(_diagramhierarchy.TaskComposition_Shapes, taskcompositionshape_)
					if idx != -1 {
						_diagramhierarchy.TaskComposition_Shapes = slices.Delete(_diagramhierarchy.TaskComposition_Shapes, idx, idx+1)
						taskcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskComposition_Shapes", &_diagramhierarchy.TaskComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskcompositionshape_.Unstage(taskcompositionshapeFormCallback.probe.stageOfInterest)
	}

	taskcompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskCompositionShape](
		taskcompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskcompositionshapeFormCallback.CreationMode || taskcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskcompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskcompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskCompositionShapeFormCallback(
			nil,
			taskcompositionshapeFormCallback.probe,
			newFormGroup,
		)
		taskcompositionshape := new(models.TaskCompositionShape)
		FillUpForm(taskcompositionshape, newFormGroup, taskcompositionshapeFormCallback.probe)
		taskcompositionshapeFormCallback.probe.formStage.Commit()
	}

	taskcompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskGroupFormCallback(
	taskgroup *models.TaskGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskgroupFormCallback *TaskGroupFormCallback) {
	taskgroupFormCallback = new(TaskGroupFormCallback)
	taskgroupFormCallback.probe = probe
	taskgroupFormCallback.taskgroup = taskgroup
	taskgroupFormCallback.formGroup = formGroup

	taskgroupFormCallback.CreationMode = (taskgroup == nil)

	return
}

type TaskGroupFormCallback struct {
	taskgroup *models.TaskGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskgroupFormCallback *TaskGroupFormCallback) OnSave() {
	taskgroupFormCallback.probe.stageOfInterest.Lock()
	defer taskgroupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskgroupFormCallback.probe.formStage.Checkout()

	if taskgroupFormCallback.taskgroup == nil {
		taskgroupFormCallback.taskgroup = new(models.TaskGroup).Stage(taskgroupFormCallback.probe.stageOfInterest)
	}
	taskgroup_ := taskgroupFormCallback.taskgroup
	_ = taskgroup_

	for _, formDiv := range taskgroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskgroup_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(taskgroup_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(taskgroup_.IsExpanded), formDiv)
		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](taskgroupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					taskgroupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](taskgroupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			taskgroup_.Tasks = instanceSlice
			taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(taskgroup_, "Tasks", &taskgroup_.Tasks)

		case "DiagramHierarchy:TaskGroupsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskgroupFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TaskGroupsWhoseNodeIsExpanded slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskgroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskgroupFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskgroup_ is in _diagramhierarchy.TaskGroupsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramhierarchy.TaskGroupsWhoseNodeIsExpanded {
						if _b == taskgroup_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded = append(_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded, taskgroup_)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskGroupsWhoseNodeIsExpanded", &_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded)
					}
				} else {
					// ensure taskgroup_ is NOT in _diagramhierarchy.TaskGroupsWhoseNodeIsExpanded
					idx := slices.Index(_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded, taskgroup_)
					if idx != -1 {
						_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded = slices.Delete(_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded, idx, idx+1)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskGroupsWhoseNodeIsExpanded", &_diagramhierarchy.TaskGroupsWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootTaskGroups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](taskgroupFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootTaskGroups slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](taskgroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskgroupFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure taskgroup_ is in _library.RootTaskGroups
					found := false
					for _, _b := range _library.RootTaskGroups {
						if _b == taskgroup_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootTaskGroups = append(_library.RootTaskGroups, taskgroup_)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootTaskGroups", &_library.RootTaskGroups)
					}
				} else {
					// ensure taskgroup_ is NOT in _library.RootTaskGroups
					idx := slices.Index(_library.RootTaskGroups, taskgroup_)
					if idx != -1 {
						_library.RootTaskGroups = slices.Delete(_library.RootTaskGroups, idx, idx+1)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootTaskGroups", &_library.RootTaskGroups)
					}
				}
			}
		case "Milestone:TaskGroupsToDisplay":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Milestone instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Milestone instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Milestone](taskgroupFormCallback.probe.stageOfInterest)
			targetMilestoneIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetMilestoneIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Milestone instances and update their TaskGroupsToDisplay slice
			for _milestone := range *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](taskgroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskgroupFormCallback.probe.stageOfInterest, _milestone)
				
				// if Milestone is selected
				if targetMilestoneIDs[id] {
					// ensure taskgroup_ is in _milestone.TaskGroupsToDisplay
					found := false
					for _, _b := range _milestone.TaskGroupsToDisplay {
						if _b == taskgroup_ {
							found = true
							break
						}
					}
					if !found {
						_milestone.TaskGroupsToDisplay = append(_milestone.TaskGroupsToDisplay, taskgroup_)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_milestone, "TaskGroupsToDisplay", &_milestone.TaskGroupsToDisplay)
					}
				} else {
					// ensure taskgroup_ is NOT in _milestone.TaskGroupsToDisplay
					idx := slices.Index(_milestone.TaskGroupsToDisplay, taskgroup_)
					if idx != -1 {
						_milestone.TaskGroupsToDisplay = slices.Delete(_milestone.TaskGroupsToDisplay, idx, idx+1)
						taskgroupFormCallback.probe.UpdateSliceOfPointersCallback(_milestone, "TaskGroupsToDisplay", &_milestone.TaskGroupsToDisplay)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskgroup_.Unstage(taskgroupFormCallback.probe.stageOfInterest)
	}

	taskgroupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskGroup](
		taskgroupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskgroupFormCallback.CreationMode || taskgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskgroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskGroupFormCallback(
			nil,
			taskgroupFormCallback.probe,
			newFormGroup,
		)
		taskgroup := new(models.TaskGroup)
		FillUpForm(taskgroup, newFormGroup, taskgroupFormCallback.probe)
		taskgroupFormCallback.probe.formStage.Commit()
	}

	taskgroupFormCallback.probe.ux_tree()
}
func __gong__New__TaskGroupShapeFormCallback(
	taskgroupshape *models.TaskGroupShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskgroupshapeFormCallback *TaskGroupShapeFormCallback) {
	taskgroupshapeFormCallback = new(TaskGroupShapeFormCallback)
	taskgroupshapeFormCallback.probe = probe
	taskgroupshapeFormCallback.taskgroupshape = taskgroupshape
	taskgroupshapeFormCallback.formGroup = formGroup

	taskgroupshapeFormCallback.CreationMode = (taskgroupshape == nil)

	return
}

type TaskGroupShapeFormCallback struct {
	taskgroupshape *models.TaskGroupShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskgroupshapeFormCallback *TaskGroupShapeFormCallback) OnSave() {
	taskgroupshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskgroupshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskGroupShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskgroupshapeFormCallback.probe.formStage.Checkout()

	if taskgroupshapeFormCallback.taskgroupshape == nil {
		taskgroupshapeFormCallback.taskgroupshape = new(models.TaskGroupShape).Stage(taskgroupshapeFormCallback.probe.stageOfInterest)
	}
	taskgroupshape_ := taskgroupshapeFormCallback.taskgroupshape
	_ = taskgroupshape_

	for _, formDiv := range taskgroupshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskgroupshape_.Name), formDiv)
		case "TaskGroup":
			FormDivSelectFieldToField(&(taskgroupshape_.TaskGroup), taskgroupshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(taskgroupshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(taskgroupshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(taskgroupshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(taskgroupshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskgroupshape_.IsHidden), formDiv)
		case "DiagramHierarchy:TaskGroupShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskgroupshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TaskGroupShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskgroupshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskgroupshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskgroupshape_ is in _diagramhierarchy.TaskGroupShapes
					found := false
					for _, _b := range _diagramhierarchy.TaskGroupShapes {
						if _b == taskgroupshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TaskGroupShapes = append(_diagramhierarchy.TaskGroupShapes, taskgroupshape_)
						taskgroupshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskGroupShapes", &_diagramhierarchy.TaskGroupShapes)
					}
				} else {
					// ensure taskgroupshape_ is NOT in _diagramhierarchy.TaskGroupShapes
					idx := slices.Index(_diagramhierarchy.TaskGroupShapes, taskgroupshape_)
					if idx != -1 {
						_diagramhierarchy.TaskGroupShapes = slices.Delete(_diagramhierarchy.TaskGroupShapes, idx, idx+1)
						taskgroupshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskGroupShapes", &_diagramhierarchy.TaskGroupShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskgroupshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskgroupshape_.Unstage(taskgroupshapeFormCallback.probe.stageOfInterest)
	}

	taskgroupshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskGroupShape](
		taskgroupshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskgroupshapeFormCallback.CreationMode || taskgroupshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskgroupshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskgroupshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskGroupShapeFormCallback(
			nil,
			taskgroupshapeFormCallback.probe,
			newFormGroup,
		)
		taskgroupshape := new(models.TaskGroupShape)
		FillUpForm(taskgroupshape, newFormGroup, taskgroupshapeFormCallback.probe)
		taskgroupshapeFormCallback.probe.formStage.Commit()
	}

	taskgroupshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskInputShapeFormCallback(
	taskinputshape *models.TaskInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskinputshapeFormCallback *TaskInputShapeFormCallback) {
	taskinputshapeFormCallback = new(TaskInputShapeFormCallback)
	taskinputshapeFormCallback.probe = probe
	taskinputshapeFormCallback.taskinputshape = taskinputshape
	taskinputshapeFormCallback.formGroup = formGroup

	taskinputshapeFormCallback.CreationMode = (taskinputshape == nil)

	return
}

type TaskInputShapeFormCallback struct {
	taskinputshape *models.TaskInputShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskinputshapeFormCallback *TaskInputShapeFormCallback) OnSave() {
	taskinputshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskinputshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskInputShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskinputshapeFormCallback.probe.formStage.Checkout()

	if taskinputshapeFormCallback.taskinputshape == nil {
		taskinputshapeFormCallback.taskinputshape = new(models.TaskInputShape).Stage(taskinputshapeFormCallback.probe.stageOfInterest)
	}
	taskinputshape_ := taskinputshapeFormCallback.taskinputshape
	_ = taskinputshape_

	for _, formDiv := range taskinputshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskinputshape_.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(taskinputshape_.Product), taskinputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(taskinputshape_.Task), taskinputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(taskinputshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(taskinputshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(taskinputshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(taskinputshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(taskinputshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskinputshape_.IsHidden), formDiv)
		case "DiagramHierarchy:TaskInputShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskinputshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TaskInputShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskinputshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskinputshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskinputshape_ is in _diagramhierarchy.TaskInputShapes
					found := false
					for _, _b := range _diagramhierarchy.TaskInputShapes {
						if _b == taskinputshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TaskInputShapes = append(_diagramhierarchy.TaskInputShapes, taskinputshape_)
						taskinputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskInputShapes", &_diagramhierarchy.TaskInputShapes)
					}
				} else {
					// ensure taskinputshape_ is NOT in _diagramhierarchy.TaskInputShapes
					idx := slices.Index(_diagramhierarchy.TaskInputShapes, taskinputshape_)
					if idx != -1 {
						_diagramhierarchy.TaskInputShapes = slices.Delete(_diagramhierarchy.TaskInputShapes, idx, idx+1)
						taskinputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskInputShapes", &_diagramhierarchy.TaskInputShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskinputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskinputshape_.Unstage(taskinputshapeFormCallback.probe.stageOfInterest)
	}

	taskinputshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskInputShape](
		taskinputshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskinputshapeFormCallback.CreationMode || taskinputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskinputshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskinputshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskInputShapeFormCallback(
			nil,
			taskinputshapeFormCallback.probe,
			newFormGroup,
		)
		taskinputshape := new(models.TaskInputShape)
		FillUpForm(taskinputshape, newFormGroup, taskinputshapeFormCallback.probe)
		taskinputshapeFormCallback.probe.formStage.Commit()
	}

	taskinputshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskOutputShapeFormCallback(
	taskoutputshape *models.TaskOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskoutputshapeFormCallback *TaskOutputShapeFormCallback) {
	taskoutputshapeFormCallback = new(TaskOutputShapeFormCallback)
	taskoutputshapeFormCallback.probe = probe
	taskoutputshapeFormCallback.taskoutputshape = taskoutputshape
	taskoutputshapeFormCallback.formGroup = formGroup

	taskoutputshapeFormCallback.CreationMode = (taskoutputshape == nil)

	return
}

type TaskOutputShapeFormCallback struct {
	taskoutputshape *models.TaskOutputShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskoutputshapeFormCallback *TaskOutputShapeFormCallback) OnSave() {
	taskoutputshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskoutputshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskOutputShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskoutputshapeFormCallback.probe.formStage.Checkout()

	if taskoutputshapeFormCallback.taskoutputshape == nil {
		taskoutputshapeFormCallback.taskoutputshape = new(models.TaskOutputShape).Stage(taskoutputshapeFormCallback.probe.stageOfInterest)
	}
	taskoutputshape_ := taskoutputshapeFormCallback.taskoutputshape
	_ = taskoutputshape_

	for _, formDiv := range taskoutputshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskoutputshape_.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(taskoutputshape_.Task), taskoutputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(taskoutputshape_.Product), taskoutputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(taskoutputshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(taskoutputshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(taskoutputshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(taskoutputshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(taskoutputshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskoutputshape_.IsHidden), formDiv)
		case "DiagramHierarchy:TaskOutputShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskoutputshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their TaskOutputShapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskoutputshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskoutputshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskoutputshape_ is in _diagramhierarchy.TaskOutputShapes
					found := false
					for _, _b := range _diagramhierarchy.TaskOutputShapes {
						if _b == taskoutputshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.TaskOutputShapes = append(_diagramhierarchy.TaskOutputShapes, taskoutputshape_)
						taskoutputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskOutputShapes", &_diagramhierarchy.TaskOutputShapes)
					}
				} else {
					// ensure taskoutputshape_ is NOT in _diagramhierarchy.TaskOutputShapes
					idx := slices.Index(_diagramhierarchy.TaskOutputShapes, taskoutputshape_)
					if idx != -1 {
						_diagramhierarchy.TaskOutputShapes = slices.Delete(_diagramhierarchy.TaskOutputShapes, idx, idx+1)
						taskoutputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "TaskOutputShapes", &_diagramhierarchy.TaskOutputShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskoutputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskoutputshape_.Unstage(taskoutputshapeFormCallback.probe.stageOfInterest)
	}

	taskoutputshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskOutputShape](
		taskoutputshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskoutputshapeFormCallback.CreationMode || taskoutputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskoutputshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskoutputshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskOutputShapeFormCallback(
			nil,
			taskoutputshapeFormCallback.probe,
			newFormGroup,
		)
		taskoutputshape := new(models.TaskOutputShape)
		FillUpForm(taskoutputshape, newFormGroup, taskoutputshapeFormCallback.probe)
		taskoutputshapeFormCallback.probe.formStage.Commit()
	}

	taskoutputshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskShapeFormCallback(
	taskshape *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskshapeFormCallback *TaskShapeFormCallback) {
	taskshapeFormCallback = new(TaskShapeFormCallback)
	taskshapeFormCallback.probe = probe
	taskshapeFormCallback.taskshape = taskshape
	taskshapeFormCallback.formGroup = formGroup

	taskshapeFormCallback.CreationMode = (taskshape == nil)

	return
}

type TaskShapeFormCallback struct {
	taskshape *models.TaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskshapeFormCallback *TaskShapeFormCallback) OnSave() {
	taskshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskshapeFormCallback.probe.formStage.Checkout()

	if taskshapeFormCallback.taskshape == nil {
		taskshapeFormCallback.taskshape = new(models.TaskShape).Stage(taskshapeFormCallback.probe.stageOfInterest)
	}
	taskshape_ := taskshapeFormCallback.taskshape
	_ = taskshape_

	for _, formDiv := range taskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskshape_.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(taskshape_.Task), taskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsShowDate":
			FormDivBasicFieldToField(&(taskshape_.IsShowDate), formDiv)
		case "X":
			FormDivBasicFieldToField(&(taskshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(taskshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(taskshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(taskshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskshape_.IsHidden), formDiv)
		case "DiagramHierarchy:Task_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramHierarchy instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramHierarchy instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramHierarchy](taskshapeFormCallback.probe.stageOfInterest)
			targetDiagramHierarchyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramHierarchyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramHierarchy instances and update their Task_Shapes slice
			for _diagramhierarchy := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramHierarchy](taskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskshapeFormCallback.probe.stageOfInterest, _diagramhierarchy)
				
				// if DiagramHierarchy is selected
				if targetDiagramHierarchyIDs[id] {
					// ensure taskshape_ is in _diagramhierarchy.Task_Shapes
					found := false
					for _, _b := range _diagramhierarchy.Task_Shapes {
						if _b == taskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramhierarchy.Task_Shapes = append(_diagramhierarchy.Task_Shapes, taskshape_)
						taskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Task_Shapes", &_diagramhierarchy.Task_Shapes)
					}
				} else {
					// ensure taskshape_ is NOT in _diagramhierarchy.Task_Shapes
					idx := slices.Index(_diagramhierarchy.Task_Shapes, taskshape_)
					if idx != -1 {
						_diagramhierarchy.Task_Shapes = slices.Delete(_diagramhierarchy.Task_Shapes, idx, idx+1)
						taskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramhierarchy, "Task_Shapes", &_diagramhierarchy.Task_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskshape_.Unstage(taskshapeFormCallback.probe.stageOfInterest)
	}

	taskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskShape](
		taskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskshapeFormCallback.CreationMode || taskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskShapeFormCallback(
			nil,
			taskshapeFormCallback.probe,
			newFormGroup,
		)
		taskshape := new(models.TaskShape)
		FillUpForm(taskshape, newFormGroup, taskshapeFormCallback.probe)
		taskshapeFormCallback.probe.formStage.Commit()
	}

	taskshapeFormCallback.probe.ux_tree()
}
