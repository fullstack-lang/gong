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
		case "Root:OrphanedProducts":
			// WARNING : this form deals with the N-N association "Root.OrphanedProducts []*Product" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Product). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Root
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Root"
				rf.Fieldname = "OrphanedProducts"
				formerAssociationSource := product_.GongGetReverseFieldOwner(
					productFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Root)
					if !ok {
						log.Fatalln("Source of Root.OrphanedProducts []*Product, is not an Root instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.OrphanedProducts, product_)
					formerSource.OrphanedProducts = slices.Delete(formerSource.OrphanedProducts, idx, idx+1)
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
			for _root := range *models.GetGongstructInstancesSet[models.Root](productFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _root.GetName() == newSourceName.GetName() {
					newSource = _root // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Root.OrphanedProducts []*Product, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.OrphanedProducts = append(newSource.OrphanedProducts, product_)
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

		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(project_.IsPBSNodeExpanded), formDiv)
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

		case "IsWBSNodeExpanded":
			FormDivBasicFieldToField(&(project_.IsWBSNodeExpanded), formDiv)
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

		case "IsDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(project_.IsDiagramsNodeExpanded), formDiv)
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

		case "OrphanedProducts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Product](rootFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Product, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Product)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Product](rootFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			root_.OrphanedProducts = instanceSlice

		case "OrphanedTasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](rootFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Task](rootFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			root_.OrphanedTasks = instanceSlice

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
		case "Root:OrphanedTasks":
			// WARNING : this form deals with the N-N association "Root.OrphanedTasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Root
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Root"
				rf.Fieldname = "OrphanedTasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Root)
					if !ok {
						log.Fatalln("Source of Root.OrphanedTasks []*Task, is not an Root instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.OrphanedTasks, task_)
					formerSource.OrphanedTasks = slices.Delete(formerSource.OrphanedTasks, idx, idx+1)
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
			for _root := range *models.GetGongstructInstancesSet[models.Root](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _root.GetName() == newSourceName.GetName() {
					newSource = _root // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Root.OrphanedTasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.OrphanedTasks = append(newSource.OrphanedTasks, task_)
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
