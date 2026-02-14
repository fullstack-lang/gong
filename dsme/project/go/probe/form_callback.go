// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *table.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {
	diagramFormCallback.probe.stageOfInterest.Lock()
	defer diagramFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagram_.IsEditable_), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(diagram_.IsInRenameMode), formDiv)
		case "ShowPrefix":
			FormDivBasicFieldToField(&(diagram_.ShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxHeigth), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagram_.ComputedPrefix), formDiv)
		case "Product_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Product_Shapes = instanceSlice

		case "ProductsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProductsWhoseNodeIsExpanded = instanceSlice

		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsPBSNodeExpanded), formDiv)
		case "ProductComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProductComposition_Shapes = instanceSlice

		case "IsWBSNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsWBSNodeExpanded), formDiv)
		case "Task_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Task_Shapes = instanceSlice

		case "TasksWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TasksWhoseNodeIsExpanded = instanceSlice

		case "TasksWhoseInputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TasksWhoseInputNodeIsExpanded = instanceSlice

		case "TasksWhoseOutputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TasksWhoseOutputNodeIsExpanded = instanceSlice

		case "TaskComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TaskComposition_Shapes = instanceSlice

		case "TaskInputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskInputShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskInputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskInputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskInputShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TaskInputShapes = instanceSlice

		case "TaskOutputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskOutputShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskOutputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskOutputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskOutputShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.TaskOutputShapes = instanceSlice

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Note_Shapes = instanceSlice

		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NotesWhoseNodeIsExpanded = instanceSlice

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsNotesNodeExpanded), formDiv)
		case "NoteProductShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteProductShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteProductShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteProductShapes = instanceSlice

		case "NoteTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteTaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTaskShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteTaskShapes = instanceSlice

		case "Resource_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Resource_Shapes = instanceSlice

		case "ResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ResourcesWhoseNodeIsExpanded = instanceSlice

		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsResourcesNodeExpanded), formDiv)
		case "ResourceComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ResourceComposition_Shapes = instanceSlice

		case "ResourceTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ResourceTaskShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ResourceTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ResourceTaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ResourceTaskShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ResourceTaskShapes = instanceSlice

		case "Project:Diagrams":
			// WARNING : this form deals with the N-N association "Project.Diagrams []*Diagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Diagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Project
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Project"
				rf.Fieldname = "Diagrams"
				formerAssociationSource := diagram_.GongGetReverseFieldOwner(
					diagramFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Project)
					if !ok {
						log.Fatalln("Source of Project.Diagrams []*Diagram, is not an Project instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Diagrams, diagram_)
					formerSource.Diagrams = slices.Delete(formerSource.Diagrams, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Project
			for _project := range *models.GetGongstructInstancesSet[models.Project](diagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _project.GetName() == newSourceName.GetName() {
					newSource = _project // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Project.Diagrams []*Diagram, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Diagrams = append(newSource.Diagrams, diagram_)
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Diagram](
		diagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(diagramFormCallback.probe)
}
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "Diagram:NotesWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.NotesWhoseNodeIsExpanded []*Note" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Note). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "NotesWhoseNodeIsExpanded"
				formerAssociationSource := note_.GongGetReverseFieldOwner(
					noteFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.NotesWhoseNodeIsExpanded []*Note, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.NotesWhoseNodeIsExpanded, note_)
					formerSource.NotesWhoseNodeIsExpanded = slices.Delete(formerSource.NotesWhoseNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](noteFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.NotesWhoseNodeIsExpanded []*Note, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.NotesWhoseNodeIsExpanded = append(newSource.NotesWhoseNodeIsExpanded, note_)
		case "Project:Notes":
			// WARNING : this form deals with the N-N association "Project.Notes []*Note" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Note). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Project
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Project"
				rf.Fieldname = "Notes"
				formerAssociationSource := note_.GongGetReverseFieldOwner(
					noteFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Project)
					if !ok {
						log.Fatalln("Source of Project.Notes []*Note, is not an Project instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Notes, note_)
					formerSource.Notes = slices.Delete(formerSource.Notes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Project
			for _project := range *models.GetGongstructInstancesSet[models.Project](noteFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _project.GetName() == newSourceName.GetName() {
					newSource = _project // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Project.Notes []*Note, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Notes = append(newSource.Notes, note_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(noteFormCallback.probe)
}
func __gong__New__NoteProductShapeFormCallback(
	noteproductshape *models.NoteProductShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:NoteProductShapes":
			// WARNING : this form deals with the N-N association "Diagram.NoteProductShapes []*NoteProductShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of NoteProductShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "NoteProductShapes"
				formerAssociationSource := noteproductshape_.GongGetReverseFieldOwner(
					noteproductshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.NoteProductShapes []*NoteProductShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.NoteProductShapes, noteproductshape_)
					formerSource.NoteProductShapes = slices.Delete(formerSource.NoteProductShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](noteproductshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.NoteProductShapes []*NoteProductShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.NoteProductShapes = append(newSource.NoteProductShapes, noteproductshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(noteproductshapeFormCallback.probe)
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(noteshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(noteshape_.Height), formDiv)
		case "Diagram:Note_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Note_Shapes []*NoteShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of NoteShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Note_Shapes"
				formerAssociationSource := noteshape_.GongGetReverseFieldOwner(
					noteshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Note_Shapes []*NoteShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Note_Shapes, noteshape_)
					formerSource.Note_Shapes = slices.Delete(formerSource.Note_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](noteshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Note_Shapes []*NoteShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Note_Shapes = append(newSource.Note_Shapes, noteshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(noteshapeFormCallback.probe)
}
func __gong__New__NoteTaskShapeFormCallback(
	notetaskshape *models.NoteTaskShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:NoteTaskShapes":
			// WARNING : this form deals with the N-N association "Diagram.NoteTaskShapes []*NoteTaskShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of NoteTaskShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "NoteTaskShapes"
				formerAssociationSource := notetaskshape_.GongGetReverseFieldOwner(
					notetaskshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.NoteTaskShapes []*NoteTaskShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.NoteTaskShapes, notetaskshape_)
					formerSource.NoteTaskShapes = slices.Delete(formerSource.NoteTaskShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](notetaskshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.NoteTaskShapes []*NoteTaskShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.NoteTaskShapes = append(newSource.NoteTaskShapes, notetaskshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(notetaskshapeFormCallback.probe)
}
func __gong__New__ProductFormCallback(
	product *models.Product,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "IsExpanded":
			FormDivBasicFieldToField(&(product_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(product_.ComputedPrefix), formDiv)
		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&(product_.IsProducersNodeExpanded), formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&(product_.IsConsumersNodeExpanded), formDiv)
		case "Diagram:ProductsWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.ProductsWhoseNodeIsExpanded []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ProductsWhoseNodeIsExpanded"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ProductsWhoseNodeIsExpanded []*Product, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProductsWhoseNodeIsExpanded, product_)
					formerSource.ProductsWhoseNodeIsExpanded = slices.Delete(formerSource.ProductsWhoseNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ProductsWhoseNodeIsExpanded []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProductsWhoseNodeIsExpanded = append(newSource.ProductsWhoseNodeIsExpanded, product_)
		case "Note:Products":
			// WARNING : this form deals with the N-N association "Note.Products []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Note
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Note"
				rf.Fieldname = "Products"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Note)
					if !ok {
						log.Fatalln("Source of Note.Products []*Product, is not an Note instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Products, product_)
					formerSource.Products = slices.Delete(formerSource.Products, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Note
			for _note := range *models.GetGongstructInstancesSet[models.Note](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _note.GetName() == newSourceName.GetName() {
					newSource = _note // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Note.Products []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Products = append(newSource.Products, product_)
		case "Product:SubProducts":
			// WARNING : this form deals with the N-N association "Product.SubProducts []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Product
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Product"
				rf.Fieldname = "SubProducts"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Product)
					if !ok {
						log.Fatalln("Source of Product.SubProducts []*Product, is not an Product instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubProducts, product_)
					formerSource.SubProducts = slices.Delete(formerSource.SubProducts, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Product
			for _product := range *models.GetGongstructInstancesSet[models.Product](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _product.GetName() == newSourceName.GetName() {
					newSource = _product // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Product.SubProducts []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubProducts = append(newSource.SubProducts, product_)
		case "Project:RootProducts":
			// WARNING : this form deals with the N-N association "Project.RootProducts []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Project
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Project"
				rf.Fieldname = "RootProducts"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Project)
					if !ok {
						log.Fatalln("Source of Project.RootProducts []*Product, is not an Project instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootProducts, product_)
					formerSource.RootProducts = slices.Delete(formerSource.RootProducts, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Project
			for _project := range *models.GetGongstructInstancesSet[models.Project](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _project.GetName() == newSourceName.GetName() {
					newSource = _project // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Project.RootProducts []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootProducts = append(newSource.RootProducts, product_)
		case "Task:Inputs":
			// WARNING : this form deals with the N-N association "Task.Inputs []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Task
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Task"
				rf.Fieldname = "Inputs"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Task)
					if !ok {
						log.Fatalln("Source of Task.Inputs []*Product, is not an Task instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Inputs, product_)
					formerSource.Inputs = slices.Delete(formerSource.Inputs, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Task
			for _task := range *models.GetGongstructInstancesSet[models.Task](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _task.GetName() == newSourceName.GetName() {
					newSource = _task // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Task.Inputs []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Inputs = append(newSource.Inputs, product_)
		case "Task:Outputs":
			// WARNING : this form deals with the N-N association "Task.Outputs []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Task
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Task"
				rf.Fieldname = "Outputs"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Task)
					if !ok {
						log.Fatalln("Source of Task.Outputs []*Product, is not an Task instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Outputs, product_)
					formerSource.Outputs = slices.Delete(formerSource.Outputs, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Task
			for _task := range *models.GetGongstructInstancesSet[models.Task](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _task.GetName() == newSourceName.GetName() {
					newSource = _task // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Task.Outputs []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Outputs = append(newSource.Outputs, product_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(productFormCallback.probe)
}
func __gong__New__ProductCompositionShapeFormCallback(
	productcompositionshape *models.ProductCompositionShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:ProductComposition_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.ProductComposition_Shapes []*ProductCompositionShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProductCompositionShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ProductComposition_Shapes"
				formerAssociationSource := productcompositionshape_.GongGetReverseFieldOwner(
					productcompositionshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ProductComposition_Shapes []*ProductCompositionShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProductComposition_Shapes, productcompositionshape_)
					formerSource.ProductComposition_Shapes = slices.Delete(formerSource.ProductComposition_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](productcompositionshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ProductComposition_Shapes []*ProductCompositionShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProductComposition_Shapes = append(newSource.ProductComposition_Shapes, productcompositionshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(productcompositionshapeFormCallback.probe)
}
func __gong__New__ProductShapeFormCallback(
	productshape *models.ProductShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(productshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(productshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(productshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(productshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(productshape_.Height), formDiv)
		case "Diagram:Product_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Product_Shapes []*ProductShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProductShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Product_Shapes"
				formerAssociationSource := productshape_.GongGetReverseFieldOwner(
					productshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Product_Shapes []*ProductShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Product_Shapes, productshape_)
					formerSource.Product_Shapes = slices.Delete(formerSource.Product_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](productshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Product_Shapes []*ProductShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Product_Shapes = append(newSource.Product_Shapes, productshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(productshapeFormCallback.probe)
}
func __gong__New__ProjectFormCallback(
	project *models.Project,
	probe *Probe,
	formGroup *table.FormGroup,
) (projectFormCallback *ProjectFormCallback) {
	projectFormCallback = new(ProjectFormCallback)
	projectFormCallback.probe = probe
	projectFormCallback.project = project
	projectFormCallback.formGroup = formGroup

	projectFormCallback.CreationMode = (project == nil)

	return
}

type ProjectFormCallback struct {
	project *models.Project

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (projectFormCallback *ProjectFormCallback) OnSave() {
	projectFormCallback.probe.stageOfInterest.Lock()
	defer projectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	projectFormCallback.probe.formStage.Checkout()

	if projectFormCallback.project == nil {
		projectFormCallback.project = new(models.Project).Stage(projectFormCallback.probe.stageOfInterest)
	}
	project_ := projectFormCallback.project
	_ = project_

	for _, formDiv := range projectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(project_.Name), formDiv)
		case "RootProducts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](projectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					projectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Product](projectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			project_.RootProducts = instanceSlice

		case "RootTasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](projectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					projectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](projectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			project_.RootTasks = instanceSlice

		case "RootResources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](projectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					projectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](projectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			project_.RootResources = instanceSlice

		case "Notes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](projectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					projectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](projectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			project_.Notes = instanceSlice

		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](projectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					projectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](projectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			project_.Diagrams = instanceSlice

		case "IsExpanded":
			FormDivBasicFieldToField(&(project_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(project_.ComputedPrefix), formDiv)
		case "Root:Projects":
			// WARNING : this form deals with the N-N association "Root.Projects []*Project" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Project). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Root
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Root"
				rf.Fieldname = "Projects"
				formerAssociationSource := project_.GongGetReverseFieldOwner(
					projectFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Root)
					if !ok {
						log.Fatalln("Source of Root.Projects []*Project, is not an Root instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Projects, project_)
					formerSource.Projects = slices.Delete(formerSource.Projects, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Root
			for _root := range *models.GetGongstructInstancesSet[models.Root](projectFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _root.GetName() == newSourceName.GetName() {
					newSource = _root // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Root.Projects []*Project, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Projects = append(newSource.Projects, project_)
		}
	}

	// manage the suppress operation
	if projectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		project_.Unstage(projectFormCallback.probe.stageOfInterest)
	}

	projectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Project](
		projectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if projectFormCallback.CreationMode || projectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		projectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(projectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProjectFormCallback(
			nil,
			projectFormCallback.probe,
			newFormGroup,
		)
		project := new(models.Project)
		FillUpForm(project, newFormGroup, projectFormCallback.probe)
		projectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(projectFormCallback.probe)
}
func __gong__New__ResourceFormCallback(
	resource *models.Resource,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "IsExpanded":
			FormDivBasicFieldToField(&(resource_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(resource_.ComputedPrefix), formDiv)
		case "Diagram:ResourcesWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.ResourcesWhoseNodeIsExpanded []*Resource" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Resource). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ResourcesWhoseNodeIsExpanded"
				formerAssociationSource := resource_.GongGetReverseFieldOwner(
					resourceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ResourcesWhoseNodeIsExpanded []*Resource, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ResourcesWhoseNodeIsExpanded, resource_)
					formerSource.ResourcesWhoseNodeIsExpanded = slices.Delete(formerSource.ResourcesWhoseNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](resourceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ResourcesWhoseNodeIsExpanded []*Resource, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ResourcesWhoseNodeIsExpanded = append(newSource.ResourcesWhoseNodeIsExpanded, resource_)
		case "Project:RootResources":
			// WARNING : this form deals with the N-N association "Project.RootResources []*Resource" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Resource). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Project
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Project"
				rf.Fieldname = "RootResources"
				formerAssociationSource := resource_.GongGetReverseFieldOwner(
					resourceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Project)
					if !ok {
						log.Fatalln("Source of Project.RootResources []*Resource, is not an Project instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootResources, resource_)
					formerSource.RootResources = slices.Delete(formerSource.RootResources, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Project
			for _project := range *models.GetGongstructInstancesSet[models.Project](resourceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _project.GetName() == newSourceName.GetName() {
					newSource = _project // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Project.RootResources []*Resource, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootResources = append(newSource.RootResources, resource_)
		case "Resource:SubResources":
			// WARNING : this form deals with the N-N association "Resource.SubResources []*Resource" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Resource). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Resource
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Resource"
				rf.Fieldname = "SubResources"
				formerAssociationSource := resource_.GongGetReverseFieldOwner(
					resourceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Resource)
					if !ok {
						log.Fatalln("Source of Resource.SubResources []*Resource, is not an Resource instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubResources, resource_)
					formerSource.SubResources = slices.Delete(formerSource.SubResources, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Resource
			for _resource := range *models.GetGongstructInstancesSet[models.Resource](resourceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _resource.GetName() == newSourceName.GetName() {
					newSource = _resource // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Resource.SubResources []*Resource, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubResources = append(newSource.SubResources, resource_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(resourceFormCallback.probe)
}
func __gong__New__ResourceCompositionShapeFormCallback(
	resourcecompositionshape *models.ResourceCompositionShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:ResourceComposition_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.ResourceComposition_Shapes []*ResourceCompositionShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ResourceCompositionShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ResourceComposition_Shapes"
				formerAssociationSource := resourcecompositionshape_.GongGetReverseFieldOwner(
					resourcecompositionshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ResourceComposition_Shapes []*ResourceCompositionShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ResourceComposition_Shapes, resourcecompositionshape_)
					formerSource.ResourceComposition_Shapes = slices.Delete(formerSource.ResourceComposition_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](resourcecompositionshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ResourceComposition_Shapes []*ResourceCompositionShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ResourceComposition_Shapes = append(newSource.ResourceComposition_Shapes, resourcecompositionshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(resourcecompositionshapeFormCallback.probe)
}
func __gong__New__ResourceShapeFormCallback(
	resourceshape *models.ResourceShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(resourceshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(resourceshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(resourceshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(resourceshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(resourceshape_.Height), formDiv)
		case "Diagram:Resource_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Resource_Shapes []*ResourceShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ResourceShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Resource_Shapes"
				formerAssociationSource := resourceshape_.GongGetReverseFieldOwner(
					resourceshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Resource_Shapes []*ResourceShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Resource_Shapes, resourceshape_)
					formerSource.Resource_Shapes = slices.Delete(formerSource.Resource_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](resourceshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Resource_Shapes []*ResourceShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Resource_Shapes = append(newSource.Resource_Shapes, resourceshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(resourceshapeFormCallback.probe)
}
func __gong__New__ResourceTaskShapeFormCallback(
	resourcetaskshape *models.ResourceTaskShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:ResourceTaskShapes":
			// WARNING : this form deals with the N-N association "Diagram.ResourceTaskShapes []*ResourceTaskShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ResourceTaskShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ResourceTaskShapes"
				formerAssociationSource := resourcetaskshape_.GongGetReverseFieldOwner(
					resourcetaskshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ResourceTaskShapes []*ResourceTaskShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ResourceTaskShapes, resourcetaskshape_)
					formerSource.ResourceTaskShapes = slices.Delete(formerSource.ResourceTaskShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](resourcetaskshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ResourceTaskShapes []*ResourceTaskShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ResourceTaskShapes = append(newSource.ResourceTaskShapes, resourcetaskshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(resourcetaskshapeFormCallback.probe)
}
func __gong__New__RootFormCallback(
	root *models.Root,
	probe *Probe,
	formGroup *table.FormGroup,
) (rootFormCallback *RootFormCallback) {
	rootFormCallback = new(RootFormCallback)
	rootFormCallback.probe = probe
	rootFormCallback.root = root
	rootFormCallback.formGroup = formGroup

	rootFormCallback.CreationMode = (root == nil)

	return
}

type RootFormCallback struct {
	root *models.Root

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rootFormCallback *RootFormCallback) OnSave() {
	rootFormCallback.probe.stageOfInterest.Lock()
	defer rootFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RootFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rootFormCallback.probe.formStage.Checkout()

	if rootFormCallback.root == nil {
		rootFormCallback.root = new(models.Root).Stage(rootFormCallback.probe.stageOfInterest)
	}
	root_ := rootFormCallback.root
	_ = root_

	for _, formDiv := range rootFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(root_.Name), formDiv)
		case "Projects":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Project](rootFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Project, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Project)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rootFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Project](rootFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			root_.Projects = instanceSlice

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(root_.NbPixPerCharacter), formDiv)
		}
	}

	// manage the suppress operation
	if rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		root_.Unstage(rootFormCallback.probe.stageOfInterest)
	}

	rootFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Root](
		rootFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rootFormCallback.CreationMode || rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rootFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rootFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RootFormCallback(
			nil,
			rootFormCallback.probe,
			newFormGroup,
		)
		root := new(models.Root)
		FillUpForm(root, newFormGroup, rootFormCallback.probe)
		rootFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rootFormCallback.probe)
}
func __gong__New__TaskFormCallback(
	task *models.Task,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "IsExpanded":
			FormDivBasicFieldToField(&(task_.IsExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(task_.ComputedPrefix), formDiv)
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

		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&(task_.IsOutputsNodeExpanded), formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&(task_.IsWithCompletion), formDiv)
		case "Completion":
			FormDivEnumStringFieldToField(&(task_.Completion), formDiv)
		case "Diagram:TasksWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.TasksWhoseNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TasksWhoseNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TasksWhoseNodeIsExpanded []*Task, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TasksWhoseNodeIsExpanded, task_)
					formerSource.TasksWhoseNodeIsExpanded = slices.Delete(formerSource.TasksWhoseNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TasksWhoseNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TasksWhoseNodeIsExpanded = append(newSource.TasksWhoseNodeIsExpanded, task_)
		case "Diagram:TasksWhoseInputNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.TasksWhoseInputNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TasksWhoseInputNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TasksWhoseInputNodeIsExpanded []*Task, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TasksWhoseInputNodeIsExpanded, task_)
					formerSource.TasksWhoseInputNodeIsExpanded = slices.Delete(formerSource.TasksWhoseInputNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TasksWhoseInputNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TasksWhoseInputNodeIsExpanded = append(newSource.TasksWhoseInputNodeIsExpanded, task_)
		case "Diagram:TasksWhoseOutputNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.TasksWhoseOutputNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TasksWhoseOutputNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TasksWhoseOutputNodeIsExpanded []*Task, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TasksWhoseOutputNodeIsExpanded, task_)
					formerSource.TasksWhoseOutputNodeIsExpanded = slices.Delete(formerSource.TasksWhoseOutputNodeIsExpanded, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TasksWhoseOutputNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TasksWhoseOutputNodeIsExpanded = append(newSource.TasksWhoseOutputNodeIsExpanded, task_)
		case "Note:Tasks":
			// WARNING : this form deals with the N-N association "Note.Tasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Note
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Note"
				rf.Fieldname = "Tasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Note)
					if !ok {
						log.Fatalln("Source of Note.Tasks []*Task, is not an Note instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Tasks, task_)
					formerSource.Tasks = slices.Delete(formerSource.Tasks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Note
			for _note := range *models.GetGongstructInstancesSet[models.Note](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _note.GetName() == newSourceName.GetName() {
					newSource = _note // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Note.Tasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Tasks = append(newSource.Tasks, task_)
		case "Project:RootTasks":
			// WARNING : this form deals with the N-N association "Project.RootTasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Project
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Project"
				rf.Fieldname = "RootTasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Project)
					if !ok {
						log.Fatalln("Source of Project.RootTasks []*Task, is not an Project instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootTasks, task_)
					formerSource.RootTasks = slices.Delete(formerSource.RootTasks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Project
			for _project := range *models.GetGongstructInstancesSet[models.Project](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _project.GetName() == newSourceName.GetName() {
					newSource = _project // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Project.RootTasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootTasks = append(newSource.RootTasks, task_)
		case "Resource:Tasks":
			// WARNING : this form deals with the N-N association "Resource.Tasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Resource
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Resource"
				rf.Fieldname = "Tasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Resource)
					if !ok {
						log.Fatalln("Source of Resource.Tasks []*Task, is not an Resource instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Tasks, task_)
					formerSource.Tasks = slices.Delete(formerSource.Tasks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Resource
			for _resource := range *models.GetGongstructInstancesSet[models.Resource](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _resource.GetName() == newSourceName.GetName() {
					newSource = _resource // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Resource.Tasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Tasks = append(newSource.Tasks, task_)
		case "Task:SubTasks":
			// WARNING : this form deals with the N-N association "Task.SubTasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Task
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Task"
				rf.Fieldname = "SubTasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Task)
					if !ok {
						log.Fatalln("Source of Task.SubTasks []*Task, is not an Task instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubTasks, task_)
					formerSource.SubTasks = slices.Delete(formerSource.SubTasks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Task
			for _task := range *models.GetGongstructInstancesSet[models.Task](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _task.GetName() == newSourceName.GetName() {
					newSource = _task // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Task.SubTasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubTasks = append(newSource.SubTasks, task_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(taskFormCallback.probe)
}
func __gong__New__TaskCompositionShapeFormCallback(
	taskcompositionshape *models.TaskCompositionShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:TaskComposition_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.TaskComposition_Shapes []*TaskCompositionShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TaskCompositionShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TaskComposition_Shapes"
				formerAssociationSource := taskcompositionshape_.GongGetReverseFieldOwner(
					taskcompositionshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TaskComposition_Shapes []*TaskCompositionShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TaskComposition_Shapes, taskcompositionshape_)
					formerSource.TaskComposition_Shapes = slices.Delete(formerSource.TaskComposition_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskcompositionshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TaskComposition_Shapes []*TaskCompositionShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TaskComposition_Shapes = append(newSource.TaskComposition_Shapes, taskcompositionshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(taskcompositionshapeFormCallback.probe)
}
func __gong__New__TaskInputShapeFormCallback(
	taskinputshape *models.TaskInputShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:TaskInputShapes":
			// WARNING : this form deals with the N-N association "Diagram.TaskInputShapes []*TaskInputShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TaskInputShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TaskInputShapes"
				formerAssociationSource := taskinputshape_.GongGetReverseFieldOwner(
					taskinputshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TaskInputShapes []*TaskInputShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TaskInputShapes, taskinputshape_)
					formerSource.TaskInputShapes = slices.Delete(formerSource.TaskInputShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskinputshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TaskInputShapes []*TaskInputShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TaskInputShapes = append(newSource.TaskInputShapes, taskinputshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(taskinputshapeFormCallback.probe)
}
func __gong__New__TaskOutputShapeFormCallback(
	taskoutputshape *models.TaskOutputShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "Diagram:TaskOutputShapes":
			// WARNING : this form deals with the N-N association "Diagram.TaskOutputShapes []*TaskOutputShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TaskOutputShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "TaskOutputShapes"
				formerAssociationSource := taskoutputshape_.GongGetReverseFieldOwner(
					taskoutputshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.TaskOutputShapes []*TaskOutputShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TaskOutputShapes, taskoutputshape_)
					formerSource.TaskOutputShapes = slices.Delete(formerSource.TaskOutputShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskoutputshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.TaskOutputShapes []*TaskOutputShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TaskOutputShapes = append(newSource.TaskOutputShapes, taskoutputshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(taskoutputshapeFormCallback.probe)
}
func __gong__New__TaskShapeFormCallback(
	taskshape *models.TaskShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(taskshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(taskshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(taskshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(taskshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(taskshape_.Height), formDiv)
		case "Diagram:Task_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Task_Shapes []*TaskShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TaskShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Task_Shapes"
				formerAssociationSource := taskshape_.GongGetReverseFieldOwner(
					taskshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Task_Shapes []*TaskShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Task_Shapes, taskshape_)
					formerSource.Task_Shapes = slices.Delete(formerSource.Task_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](taskshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Task_Shapes []*TaskShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Task_Shapes = append(newSource.Task_Shapes, taskshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(taskshapeFormCallback.probe)
}
