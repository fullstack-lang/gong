// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagram_.ComputedPrefix), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(diagram_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagram_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagram_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagram_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagram_.Height), formDiv)
		case "Process_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProcessShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProcessShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProcessShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ProcessShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Process_Shapes = instanceSlice

		case "ProcesssWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Process](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProcesssWhoseNodeIsExpanded = instanceSlice

		case "IsProcesssNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsProcesssNodeExpanded), formDiv)
		case "ProcessComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProcessCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProcessCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProcessCompositionShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ProcessCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProcessComposition_Shapes = instanceSlice

		case "Library:Diagrams":
			// WARNING : this form deals with the N-N association "Library.Diagrams []*Diagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Diagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "Diagrams"
				formerAssociationSource := diagram_.GongGetReverseFieldOwner(
					diagramFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.Diagrams []*Diagram, is not an Library instance")
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](diagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.Diagrams []*Diagram, with name", newSourceName, ", does not exist")
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
		newFormGroup := (&form.FormGroup{
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

	diagramFormCallback.probe.ux_tree()
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(library_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Diagrams = instanceSlice

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

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "Library:SubLibraries":
			// WARNING : this form deals with the N-N association "Library.SubLibraries []*Library" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Library). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "SubLibraries"
				formerAssociationSource := library_.GongGetReverseFieldOwner(
					libraryFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.SubLibraries []*Library, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubLibraries, library_)
					formerSource.SubLibraries = slices.Delete(formerSource.SubLibraries, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](libraryFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.SubLibraries []*Library, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubLibraries = append(newSource.SubLibraries, library_)
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
func __gong__New__ProcessFormCallback(
	process *models.Process,
	probe *Probe,
	formGroup *form.FormGroup,
) (processFormCallback *ProcessFormCallback) {
	processFormCallback = new(ProcessFormCallback)
	processFormCallback.probe = probe
	processFormCallback.process = process
	processFormCallback.formGroup = formGroup

	processFormCallback.CreationMode = (process == nil)

	return
}

type ProcessFormCallback struct {
	process *models.Process

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processFormCallback *ProcessFormCallback) OnSave() {
	processFormCallback.probe.stageOfInterest.Lock()
	defer processFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processFormCallback.probe.formStage.Checkout()

	if processFormCallback.process == nil {
		processFormCallback.process = new(models.Process).Stage(processFormCallback.probe.stageOfInterest)
	}
	process_ := processFormCallback.process
	_ = process_

	for _, formDiv := range processFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(process_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(process_.ComputedPrefix), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(process_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(process_.IsExpanded), formDiv)
		case "SubProcesses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.SubProcesses = instanceSlice

		case "Diagram:ProcesssWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Diagram.ProcesssWhoseNodeIsExpanded []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ProcesssWhoseNodeIsExpanded []*Process, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProcesssWhoseNodeIsExpanded, process_)
					formerSource.ProcesssWhoseNodeIsExpanded = slices.Delete(formerSource.ProcesssWhoseNodeIsExpanded, idx, idx+1)
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
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ProcesssWhoseNodeIsExpanded []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProcesssWhoseNodeIsExpanded = append(newSource.ProcesssWhoseNodeIsExpanded, process_)
		case "Process:SubProcesses":
			// WARNING : this form deals with the N-N association "Process.SubProcesses []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "SubProcesses"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.SubProcesses []*Process, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubProcesses, process_)
					formerSource.SubProcesses = slices.Delete(formerSource.SubProcesses, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.SubProcesses []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubProcesses = append(newSource.SubProcesses, process_)
		}
	}

	// manage the suppress operation
	if processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		process_.Unstage(processFormCallback.probe.stageOfInterest)
	}

	processFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Process](
		processFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processFormCallback.CreationMode || processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			processFormCallback.probe,
			newFormGroup,
		)
		process := new(models.Process)
		FillUpForm(process, newFormGroup, processFormCallback.probe)
		processFormCallback.probe.formStage.Commit()
	}

	processFormCallback.probe.ux_tree()
}
func __gong__New__ProcessCompositionShapeFormCallback(
	processcompositionshape *models.ProcessCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (processcompositionshapeFormCallback *ProcessCompositionShapeFormCallback) {
	processcompositionshapeFormCallback = new(ProcessCompositionShapeFormCallback)
	processcompositionshapeFormCallback.probe = probe
	processcompositionshapeFormCallback.processcompositionshape = processcompositionshape
	processcompositionshapeFormCallback.formGroup = formGroup

	processcompositionshapeFormCallback.CreationMode = (processcompositionshape == nil)

	return
}

type ProcessCompositionShapeFormCallback struct {
	processcompositionshape *models.ProcessCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processcompositionshapeFormCallback *ProcessCompositionShapeFormCallback) OnSave() {
	processcompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer processcompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processcompositionshapeFormCallback.probe.formStage.Checkout()

	if processcompositionshapeFormCallback.processcompositionshape == nil {
		processcompositionshapeFormCallback.processcompositionshape = new(models.ProcessCompositionShape).Stage(processcompositionshapeFormCallback.probe.stageOfInterest)
	}
	processcompositionshape_ := processcompositionshapeFormCallback.processcompositionshape
	_ = processcompositionshape_

	for _, formDiv := range processcompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(processcompositionshape_.Name), formDiv)
		case "Process":
			FormDivSelectFieldToField(&(processcompositionshape_.Process), processcompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(processcompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(processcompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(processcompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(processcompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(processcompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(processcompositionshape_.IsHidden), formDiv)
		case "Diagram:ProcessComposition_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.ProcessComposition_Shapes []*ProcessCompositionShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProcessCompositionShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ProcessComposition_Shapes"
				formerAssociationSource := processcompositionshape_.GongGetReverseFieldOwner(
					processcompositionshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ProcessComposition_Shapes []*ProcessCompositionShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProcessComposition_Shapes, processcompositionshape_)
					formerSource.ProcessComposition_Shapes = slices.Delete(formerSource.ProcessComposition_Shapes, idx, idx+1)
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
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](processcompositionshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ProcessComposition_Shapes []*ProcessCompositionShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProcessComposition_Shapes = append(newSource.ProcessComposition_Shapes, processcompositionshape_)
		}
	}

	// manage the suppress operation
	if processcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processcompositionshape_.Unstage(processcompositionshapeFormCallback.probe.stageOfInterest)
	}

	processcompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProcessCompositionShape](
		processcompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processcompositionshapeFormCallback.CreationMode || processcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processcompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processcompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessCompositionShapeFormCallback(
			nil,
			processcompositionshapeFormCallback.probe,
			newFormGroup,
		)
		processcompositionshape := new(models.ProcessCompositionShape)
		FillUpForm(processcompositionshape, newFormGroup, processcompositionshapeFormCallback.probe)
		processcompositionshapeFormCallback.probe.formStage.Commit()
	}

	processcompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProcessShapeFormCallback(
	processshape *models.ProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (processshapeFormCallback *ProcessShapeFormCallback) {
	processshapeFormCallback = new(ProcessShapeFormCallback)
	processshapeFormCallback.probe = probe
	processshapeFormCallback.processshape = processshape
	processshapeFormCallback.formGroup = formGroup

	processshapeFormCallback.CreationMode = (processshape == nil)

	return
}

type ProcessShapeFormCallback struct {
	processshape *models.ProcessShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processshapeFormCallback *ProcessShapeFormCallback) OnSave() {
	processshapeFormCallback.probe.stageOfInterest.Lock()
	defer processshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processshapeFormCallback.probe.formStage.Checkout()

	if processshapeFormCallback.processshape == nil {
		processshapeFormCallback.processshape = new(models.ProcessShape).Stage(processshapeFormCallback.probe.stageOfInterest)
	}
	processshape_ := processshapeFormCallback.processshape
	_ = processshape_

	for _, formDiv := range processshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(processshape_.Name), formDiv)
		case "Process":
			FormDivSelectFieldToField(&(processshape_.Process), processshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(processshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(processshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(processshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(processshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(processshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(processshape_.IsHidden), formDiv)
		case "Diagram:Process_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Process_Shapes []*ProcessShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProcessShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Process_Shapes"
				formerAssociationSource := processshape_.GongGetReverseFieldOwner(
					processshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Process_Shapes []*ProcessShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Process_Shapes, processshape_)
					formerSource.Process_Shapes = slices.Delete(formerSource.Process_Shapes, idx, idx+1)
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
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](processshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Process_Shapes []*ProcessShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Process_Shapes = append(newSource.Process_Shapes, processshape_)
		}
	}

	// manage the suppress operation
	if processshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processshape_.Unstage(processshapeFormCallback.probe.stageOfInterest)
	}

	processshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProcessShape](
		processshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processshapeFormCallback.CreationMode || processshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessShapeFormCallback(
			nil,
			processshapeFormCallback.probe,
			newFormGroup,
		)
		processshape := new(models.ProcessShape)
		FillUpForm(processshape, newFormGroup, processshapeFormCallback.probe)
		processshapeFormCallback.probe.formStage.Commit()
	}

	processshapeFormCallback.probe.ux_tree()
}
