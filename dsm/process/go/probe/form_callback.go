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
func __gong__New__DiagramProcessFormCallback(
	diagramprocess *models.DiagramProcess,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramprocessFormCallback *DiagramProcessFormCallback) {
	diagramprocessFormCallback = new(DiagramProcessFormCallback)
	diagramprocessFormCallback.probe = probe
	diagramprocessFormCallback.diagramprocess = diagramprocess
	diagramprocessFormCallback.formGroup = formGroup

	diagramprocessFormCallback.CreationMode = (diagramprocess == nil)

	return
}

type DiagramProcessFormCallback struct {
	diagramprocess *models.DiagramProcess

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramprocessFormCallback *DiagramProcessFormCallback) OnSave() {
	diagramprocessFormCallback.probe.stageOfInterest.Lock()
	defer diagramprocessFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramProcessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramprocessFormCallback.probe.formStage.Checkout()

	if diagramprocessFormCallback.diagramprocess == nil {
		diagramprocessFormCallback.diagramprocess = new(models.DiagramProcess).Stage(diagramprocessFormCallback.probe.stageOfInterest)
	}
	diagramprocess_ := diagramprocessFormCallback.diagramprocess
	_ = diagramprocess_

	for _, formDiv := range diagramprocessFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramprocess_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramprocess_.ComputedPrefix), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramprocess_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramprocess_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramprocess_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramprocess_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramprocess_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramprocess_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramprocess_.Height), formDiv)
		case "Process_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProcessShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProcessShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProcessShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProcessShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Process_Shapes = instanceSlice

		case "ProcesssWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ProcesssWhoseNodeIsExpanded = instanceSlice

		case "IsProcesssNodeExpanded":
			FormDivBasicFieldToField(&(diagramprocess_.IsProcesssNodeExpanded), formDiv)
		case "ProcessComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProcessCompositionShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProcessCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProcessCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProcessCompositionShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ProcessComposition_Shapes = instanceSlice

		case "Process:DiagramProcesss":
			// WARNING : this form deals with the N-N association "Process.DiagramProcesss []*DiagramProcess" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DiagramProcess). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "DiagramProcesss"
				formerAssociationSource := diagramprocess_.GongGetReverseFieldOwner(
					diagramprocessFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.DiagramProcesss []*DiagramProcess, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DiagramProcesss, diagramprocess_)
					formerSource.DiagramProcesss = slices.Delete(formerSource.DiagramProcesss, idx, idx+1)
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
			for _process := range *models.GetGongstructInstancesSet[models.Process](diagramprocessFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.DiagramProcesss []*DiagramProcess, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DiagramProcesss = append(newSource.DiagramProcesss, diagramprocess_)
		}
	}

	// manage the suppress operation
	if diagramprocessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramprocess_.Unstage(diagramprocessFormCallback.probe.stageOfInterest)
	}

	diagramprocessFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramProcess](
		diagramprocessFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramprocessFormCallback.CreationMode || diagramprocessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramprocessFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramprocessFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramProcessFormCallback(
			nil,
			diagramprocessFormCallback.probe,
			newFormGroup,
		)
		diagramprocess := new(models.DiagramProcess)
		FillUpForm(diagramprocess, newFormGroup, diagramprocessFormCallback.probe)
		diagramprocessFormCallback.probe.formStage.Commit()
	}

	diagramprocessFormCallback.probe.ux_tree()
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
		case "RootProcesses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Process](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootProcesses = instanceSlice

		case "ProcesssWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Process](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.ProcesssWhoseNodeIsExpanded = instanceSlice

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
func __gong__New__ParticipantFormCallback(
	participant *models.Participant,
	probe *Probe,
	formGroup *form.FormGroup,
) (participantFormCallback *ParticipantFormCallback) {
	participantFormCallback = new(ParticipantFormCallback)
	participantFormCallback.probe = probe
	participantFormCallback.participant = participant
	participantFormCallback.formGroup = formGroup

	participantFormCallback.CreationMode = (participant == nil)

	return
}

type ParticipantFormCallback struct {
	participant *models.Participant

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (participantFormCallback *ParticipantFormCallback) OnSave() {
	participantFormCallback.probe.stageOfInterest.Lock()
	defer participantFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParticipantFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	participantFormCallback.probe.formStage.Checkout()

	if participantFormCallback.participant == nil {
		participantFormCallback.participant = new(models.Participant).Stage(participantFormCallback.probe.stageOfInterest)
	}
	participant_ := participantFormCallback.participant
	_ = participant_

	for _, formDiv := range participantFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(participant_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(participant_.ComputedPrefix), formDiv)
		case "Process:Participants":
			// WARNING : this form deals with the N-N association "Process.Participants []*Participant" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Participant). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "Participants"
				formerAssociationSource := participant_.GongGetReverseFieldOwner(
					participantFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.Participants []*Participant, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Participants, participant_)
					formerSource.Participants = slices.Delete(formerSource.Participants, idx, idx+1)
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
			for _process := range *models.GetGongstructInstancesSet[models.Process](participantFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.Participants []*Participant, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Participants = append(newSource.Participants, participant_)
		case "Process:ParticipantWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Process.ParticipantWhoseNodeIsExpanded []*Participant" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Participant). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
				formerAssociationSource := participant_.GongGetReverseFieldOwner(
					participantFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.ParticipantWhoseNodeIsExpanded []*Participant, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ParticipantWhoseNodeIsExpanded, participant_)
					formerSource.ParticipantWhoseNodeIsExpanded = slices.Delete(formerSource.ParticipantWhoseNodeIsExpanded, idx, idx+1)
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
			for _process := range *models.GetGongstructInstancesSet[models.Process](participantFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.ParticipantWhoseNodeIsExpanded []*Participant, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ParticipantWhoseNodeIsExpanded = append(newSource.ParticipantWhoseNodeIsExpanded, participant_)
		}
	}

	// manage the suppress operation
	if participantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participant_.Unstage(participantFormCallback.probe.stageOfInterest)
	}

	participantFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Participant](
		participantFormCallback.probe,
	)

	// display a new form by reset the form stage
	if participantFormCallback.CreationMode || participantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participantFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(participantFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParticipantFormCallback(
			nil,
			participantFormCallback.probe,
			newFormGroup,
		)
		participant := new(models.Participant)
		FillUpForm(participant, newFormGroup, participantFormCallback.probe)
		participantFormCallback.probe.formStage.Commit()
	}

	participantFormCallback.probe.ux_tree()
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
		case "DiagramProcesss":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramProcess, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramProcess)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.DiagramProcesss = instanceSlice

		case "IsSubProcessNodeExpanded":
			FormDivBasicFieldToField(&(process_.IsSubProcessNodeExpanded), formDiv)
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

		case "IsParticipantsNodeExpanded":
			FormDivBasicFieldToField(&(process_.IsParticipantsNodeExpanded), formDiv)
		case "Participants":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.Participants = instanceSlice

		case "ParticipantWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.ParticipantWhoseNodeIsExpanded = instanceSlice

		case "DiagramProcess:ProcesssWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process, is not an DiagramProcess instance")
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProcesssWhoseNodeIsExpanded = append(newSource.ProcesssWhoseNodeIsExpanded, process_)
		case "Library:RootProcesses":
			// WARNING : this form deals with the N-N association "Library.RootProcesses []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "RootProcesses"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.RootProcesses []*Process, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootProcesses, process_)
					formerSource.RootProcesses = slices.Delete(formerSource.RootProcesses, idx, idx+1)
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
			for _library := range *models.GetGongstructInstancesSet[models.Library](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.RootProcesses []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootProcesses = append(newSource.RootProcesses, process_)
		case "Library:ProcesssWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Library.ProcesssWhoseNodeIsExpanded []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.ProcesssWhoseNodeIsExpanded []*Process, is not an Library instance")
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.ProcesssWhoseNodeIsExpanded []*Process, with name", newSourceName, ", does not exist")
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
		case "DiagramProcess:ProcessComposition_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.ProcessComposition_Shapes []*ProcessCompositionShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProcessCompositionShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ProcessComposition_Shapes"
				formerAssociationSource := processcompositionshape_.GongGetReverseFieldOwner(
					processcompositionshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ProcessComposition_Shapes []*ProcessCompositionShape, is not an DiagramProcess instance")
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](processcompositionshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ProcessComposition_Shapes []*ProcessCompositionShape, with name", newSourceName, ", does not exist")
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
		case "DiagramProcess:Process_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.Process_Shapes []*ProcessShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProcessShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "Process_Shapes"
				formerAssociationSource := processshape_.GongGetReverseFieldOwner(
					processshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.Process_Shapes []*ProcessShape, is not an DiagramProcess instance")
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](processshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.Process_Shapes []*ProcessShape, with name", newSourceName, ", does not exist")
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
