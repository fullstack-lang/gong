// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AllFormCallback(
	all *models.All,
	probe *Probe,
	formGroup *table.FormGroup,
) (allFormCallback *AllFormCallback) {
	allFormCallback = new(AllFormCallback)
	allFormCallback.probe = probe
	allFormCallback.all = all
	allFormCallback.formGroup = formGroup

	allFormCallback.CreationMode = (all == nil)

	return
}

type AllFormCallback struct {
	all *models.All

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (allFormCallback *AllFormCallback) OnSave() {
	allFormCallback.probe.stageOfInterest.Lock()
	defer allFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AllFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	allFormCallback.probe.formStage.Checkout()

	if allFormCallback.all == nil {
		allFormCallback.all = new(models.All).Stage(allFormCallback.probe.stageOfInterest)
	}
	all_ := allFormCallback.all
	_ = all_

	for _, formDiv := range allFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(all_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(all_.Annotation), allFormCallback.probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(all_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](allFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					allFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](allFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			all_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](allFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					allFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](allFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			all_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](allFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					allFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](allFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			all_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](allFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					allFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](allFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			all_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](allFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					allFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](allFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			all_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(all_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(all_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(all_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(all_.MaxOccurs), formDiv)
		case "All:Alls":
			// WARNING : this form deals with the N-N association "All.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.All
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "All"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.All)
					if !ok {
						log.Fatalln("Source of All.Alls []*All, is not an All instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.All
			for _all := range *models.GetGongstructInstancesSet[models.All](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _all.GetName() == newSourceName.GetName() {
					newSource = _all // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of All.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		case "Choice:Alls":
			// WARNING : this form deals with the N-N association "Choice.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Choice
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Choice"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Choice)
					if !ok {
						log.Fatalln("Source of Choice.Alls []*All, is not an Choice instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.Choice
			for _choice := range *models.GetGongstructInstancesSet[models.Choice](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _choice.GetName() == newSourceName.GetName() {
					newSource = _choice // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Choice.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		case "ComplexType:Alls":
			// WARNING : this form deals with the N-N association "ComplexType.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Alls []*All, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		case "Extension:Alls":
			// WARNING : this form deals with the N-N association "Extension.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Alls []*All, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		case "Group:Alls":
			// WARNING : this form deals with the N-N association "Group.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Alls []*All, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		case "Sequence:Alls":
			// WARNING : this form deals with the N-N association "Sequence.Alls []*All" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of All). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Sequence
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Sequence"
				rf.Fieldname = "Alls"
				formerAssociationSource := all_.GongGetReverseFieldOwner(
					allFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Sequence)
					if !ok {
						log.Fatalln("Source of Sequence.Alls []*All, is not an Sequence instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Alls, all_)
					formerSource.Alls = slices.Delete(formerSource.Alls, idx, idx+1)
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
			var newSource *models.Sequence
			for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](allFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _sequence.GetName() == newSourceName.GetName() {
					newSource = _sequence // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Sequence.Alls []*All, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Alls = append(newSource.Alls, all_)
		}
	}

	// manage the suppress operation
	if allFormCallback.formGroup.HasSuppressButtonBeenPressed {
		all_.Unstage(allFormCallback.probe.stageOfInterest)
	}

	allFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.All](
		allFormCallback.probe,
	)

	// display a new form by reset the form stage
	if allFormCallback.CreationMode || allFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(allFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AllFormCallback(
			nil,
			allFormCallback.probe,
			newFormGroup,
		)
		all := new(models.All)
		FillUpForm(all, newFormGroup, allFormCallback.probe)
		allFormCallback.probe.formStage.Commit()
	}

	allFormCallback.probe.ux_tree()
}
func __gong__New__AnnotationFormCallback(
	annotation *models.Annotation,
	probe *Probe,
	formGroup *table.FormGroup,
) (annotationFormCallback *AnnotationFormCallback) {
	annotationFormCallback = new(AnnotationFormCallback)
	annotationFormCallback.probe = probe
	annotationFormCallback.annotation = annotation
	annotationFormCallback.formGroup = formGroup

	annotationFormCallback.CreationMode = (annotation == nil)

	return
}

type AnnotationFormCallback struct {
	annotation *models.Annotation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (annotationFormCallback *AnnotationFormCallback) OnSave() {
	annotationFormCallback.probe.stageOfInterest.Lock()
	defer annotationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AnnotationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	annotationFormCallback.probe.formStage.Checkout()

	if annotationFormCallback.annotation == nil {
		annotationFormCallback.annotation = new(models.Annotation).Stage(annotationFormCallback.probe.stageOfInterest)
	}
	annotation_ := annotationFormCallback.annotation
	_ = annotation_

	for _, formDiv := range annotationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(annotation_.Name), formDiv)
		case "Documentations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Documentation](annotationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Documentation, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Documentation)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					annotationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Documentation](annotationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			annotation_.Documentations = instanceSlice

		}
	}

	// manage the suppress operation
	if annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotation_.Unstage(annotationFormCallback.probe.stageOfInterest)
	}

	annotationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Annotation](
		annotationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if annotationFormCallback.CreationMode || annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(annotationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnnotationFormCallback(
			nil,
			annotationFormCallback.probe,
			newFormGroup,
		)
		annotation := new(models.Annotation)
		FillUpForm(annotation, newFormGroup, annotationFormCallback.probe)
		annotationFormCallback.probe.formStage.Commit()
	}

	annotationFormCallback.probe.ux_tree()
}
func __gong__New__AttributeFormCallback(
	attribute *models.Attribute,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributeFormCallback *AttributeFormCallback) {
	attributeFormCallback = new(AttributeFormCallback)
	attributeFormCallback.probe = probe
	attributeFormCallback.attribute = attribute
	attributeFormCallback.formGroup = formGroup

	attributeFormCallback.CreationMode = (attribute == nil)

	return
}

type AttributeFormCallback struct {
	attribute *models.Attribute

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributeFormCallback *AttributeFormCallback) OnSave() {
	attributeFormCallback.probe.stageOfInterest.Lock()
	defer attributeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AttributeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributeFormCallback.probe.formStage.Checkout()

	if attributeFormCallback.attribute == nil {
		attributeFormCallback.attribute = new(models.Attribute).Stage(attributeFormCallback.probe.stageOfInterest)
	}
	attribute_ := attributeFormCallback.attribute
	_ = attribute_

	for _, formDiv := range attributeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(attribute_.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(attribute_.Type), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(attribute_.Annotation), attributeFormCallback.probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(attribute_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(attribute_.GoIdentifier), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(attribute_.Default), formDiv)
		case "Use":
			FormDivBasicFieldToField(&(attribute_.Use), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(attribute_.Form), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(attribute_.Fixed), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(attribute_.Ref), formDiv)
		case "TargetNamespace":
			FormDivBasicFieldToField(&(attribute_.TargetNamespace), formDiv)
		case "SimpleType":
			FormDivBasicFieldToField(&(attribute_.SimpleType), formDiv)
		case "IDXSD":
			FormDivBasicFieldToField(&(attribute_.IDXSD), formDiv)
		case "AttributeGroup:Attributes":
			// WARNING : this form deals with the N-N association "AttributeGroup.Attributes []*Attribute" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Attribute). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.AttributeGroup
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "AttributeGroup"
				rf.Fieldname = "Attributes"
				formerAssociationSource := attribute_.GongGetReverseFieldOwner(
					attributeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.AttributeGroup)
					if !ok {
						log.Fatalln("Source of AttributeGroup.Attributes []*Attribute, is not an AttributeGroup instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Attributes, attribute_)
					formerSource.Attributes = slices.Delete(formerSource.Attributes, idx, idx+1)
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
			var newSource *models.AttributeGroup
			for _attributegroup := range *models.GetGongstructInstancesSet[models.AttributeGroup](attributeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _attributegroup.GetName() == newSourceName.GetName() {
					newSource = _attributegroup // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of AttributeGroup.Attributes []*Attribute, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Attributes = append(newSource.Attributes, attribute_)
		case "ComplexType:Attributes":
			// WARNING : this form deals with the N-N association "ComplexType.Attributes []*Attribute" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Attribute). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Attributes"
				formerAssociationSource := attribute_.GongGetReverseFieldOwner(
					attributeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Attributes []*Attribute, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Attributes, attribute_)
					formerSource.Attributes = slices.Delete(formerSource.Attributes, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](attributeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Attributes []*Attribute, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Attributes = append(newSource.Attributes, attribute_)
		case "Extension:Attributes":
			// WARNING : this form deals with the N-N association "Extension.Attributes []*Attribute" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Attribute). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Attributes"
				formerAssociationSource := attribute_.GongGetReverseFieldOwner(
					attributeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Attributes []*Attribute, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Attributes, attribute_)
					formerSource.Attributes = slices.Delete(formerSource.Attributes, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](attributeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Attributes []*Attribute, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Attributes = append(newSource.Attributes, attribute_)
		}
	}

	// manage the suppress operation
	if attributeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_.Unstage(attributeFormCallback.probe.stageOfInterest)
	}

	attributeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Attribute](
		attributeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attributeFormCallback.CreationMode || attributeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attributeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeFormCallback(
			nil,
			attributeFormCallback.probe,
			newFormGroup,
		)
		attribute := new(models.Attribute)
		FillUpForm(attribute, newFormGroup, attributeFormCallback.probe)
		attributeFormCallback.probe.formStage.Commit()
	}

	attributeFormCallback.probe.ux_tree()
}
func __gong__New__AttributeGroupFormCallback(
	attributegroup *models.AttributeGroup,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributegroupFormCallback *AttributeGroupFormCallback) {
	attributegroupFormCallback = new(AttributeGroupFormCallback)
	attributegroupFormCallback.probe = probe
	attributegroupFormCallback.attributegroup = attributegroup
	attributegroupFormCallback.formGroup = formGroup

	attributegroupFormCallback.CreationMode = (attributegroup == nil)

	return
}

type AttributeGroupFormCallback struct {
	attributegroup *models.AttributeGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributegroupFormCallback *AttributeGroupFormCallback) OnSave() {
	attributegroupFormCallback.probe.stageOfInterest.Lock()
	defer attributegroupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AttributeGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributegroupFormCallback.probe.formStage.Checkout()

	if attributegroupFormCallback.attributegroup == nil {
		attributegroupFormCallback.attributegroup = new(models.AttributeGroup).Stage(attributegroupFormCallback.probe.stageOfInterest)
	}
	attributegroup_ := attributegroupFormCallback.attributegroup
	_ = attributegroup_

	for _, formDiv := range attributegroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributegroup_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(attributegroup_.NameXSD), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(attributegroup_.Annotation), attributegroupFormCallback.probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(attributegroup_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(attributegroup_.GoIdentifier), formDiv)
		case "AttributeGroups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AttributeGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AttributeGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					attributegroupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			attributegroup_.AttributeGroups = instanceSlice

		case "Ref":
			FormDivBasicFieldToField(&(attributegroup_.Ref), formDiv)
		case "Attributes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Attribute](attributegroupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Attribute, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Attribute)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					attributegroupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Attribute](attributegroupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			attributegroup_.Attributes = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(attributegroup_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(attributegroup_.Depth), formDiv)
		case "AttributeGroup:AttributeGroups":
			// WARNING : this form deals with the N-N association "AttributeGroup.AttributeGroups []*AttributeGroup" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AttributeGroup). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.AttributeGroup
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "AttributeGroup"
				rf.Fieldname = "AttributeGroups"
				formerAssociationSource := attributegroup_.GongGetReverseFieldOwner(
					attributegroupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.AttributeGroup)
					if !ok {
						log.Fatalln("Source of AttributeGroup.AttributeGroups []*AttributeGroup, is not an AttributeGroup instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.AttributeGroups, attributegroup_)
					formerSource.AttributeGroups = slices.Delete(formerSource.AttributeGroups, idx, idx+1)
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
			var newSource *models.AttributeGroup
			for _attributegroup := range *models.GetGongstructInstancesSet[models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _attributegroup.GetName() == newSourceName.GetName() {
					newSource = _attributegroup // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of AttributeGroup.AttributeGroups []*AttributeGroup, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.AttributeGroups = append(newSource.AttributeGroups, attributegroup_)
		case "ComplexType:AttributeGroups":
			// WARNING : this form deals with the N-N association "ComplexType.AttributeGroups []*AttributeGroup" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AttributeGroup). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "AttributeGroups"
				formerAssociationSource := attributegroup_.GongGetReverseFieldOwner(
					attributegroupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.AttributeGroups []*AttributeGroup, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.AttributeGroups, attributegroup_)
					formerSource.AttributeGroups = slices.Delete(formerSource.AttributeGroups, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](attributegroupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.AttributeGroups []*AttributeGroup, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.AttributeGroups = append(newSource.AttributeGroups, attributegroup_)
		case "Extension:AttributeGroups":
			// WARNING : this form deals with the N-N association "Extension.AttributeGroups []*AttributeGroup" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AttributeGroup). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "AttributeGroups"
				formerAssociationSource := attributegroup_.GongGetReverseFieldOwner(
					attributegroupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.AttributeGroups []*AttributeGroup, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.AttributeGroups, attributegroup_)
					formerSource.AttributeGroups = slices.Delete(formerSource.AttributeGroups, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](attributegroupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.AttributeGroups []*AttributeGroup, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.AttributeGroups = append(newSource.AttributeGroups, attributegroup_)
		case "Schema:AttributeGroups":
			// WARNING : this form deals with the N-N association "Schema.AttributeGroups []*AttributeGroup" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AttributeGroup). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Schema
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Schema"
				rf.Fieldname = "AttributeGroups"
				formerAssociationSource := attributegroup_.GongGetReverseFieldOwner(
					attributegroupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Schema)
					if !ok {
						log.Fatalln("Source of Schema.AttributeGroups []*AttributeGroup, is not an Schema instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.AttributeGroups, attributegroup_)
					formerSource.AttributeGroups = slices.Delete(formerSource.AttributeGroups, idx, idx+1)
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
			var newSource *models.Schema
			for _schema := range *models.GetGongstructInstancesSet[models.Schema](attributegroupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _schema.GetName() == newSourceName.GetName() {
					newSource = _schema // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Schema.AttributeGroups []*AttributeGroup, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.AttributeGroups = append(newSource.AttributeGroups, attributegroup_)
		}
	}

	// manage the suppress operation
	if attributegroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributegroup_.Unstage(attributegroupFormCallback.probe.stageOfInterest)
	}

	attributegroupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AttributeGroup](
		attributegroupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attributegroupFormCallback.CreationMode || attributegroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributegroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attributegroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeGroupFormCallback(
			nil,
			attributegroupFormCallback.probe,
			newFormGroup,
		)
		attributegroup := new(models.AttributeGroup)
		FillUpForm(attributegroup, newFormGroup, attributegroupFormCallback.probe)
		attributegroupFormCallback.probe.formStage.Commit()
	}

	attributegroupFormCallback.probe.ux_tree()
}
func __gong__New__ChoiceFormCallback(
	choice *models.Choice,
	probe *Probe,
	formGroup *table.FormGroup,
) (choiceFormCallback *ChoiceFormCallback) {
	choiceFormCallback = new(ChoiceFormCallback)
	choiceFormCallback.probe = probe
	choiceFormCallback.choice = choice
	choiceFormCallback.formGroup = formGroup

	choiceFormCallback.CreationMode = (choice == nil)

	return
}

type ChoiceFormCallback struct {
	choice *models.Choice

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (choiceFormCallback *ChoiceFormCallback) OnSave() {
	choiceFormCallback.probe.stageOfInterest.Lock()
	defer choiceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ChoiceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	choiceFormCallback.probe.formStage.Checkout()

	if choiceFormCallback.choice == nil {
		choiceFormCallback.choice = new(models.Choice).Stage(choiceFormCallback.probe.stageOfInterest)
	}
	choice_ := choiceFormCallback.choice
	_ = choice_

	for _, formDiv := range choiceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(choice_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(choice_.Annotation), choiceFormCallback.probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(choice_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](choiceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					choiceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](choiceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			choice_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](choiceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					choiceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](choiceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			choice_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](choiceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					choiceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](choiceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			choice_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](choiceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					choiceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](choiceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			choice_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](choiceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					choiceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](choiceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			choice_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(choice_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(choice_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(choice_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(choice_.MaxOccurs), formDiv)
		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(choice_.IsDuplicatedInXSD), formDiv)
		case "All:Choices":
			// WARNING : this form deals with the N-N association "All.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.All
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "All"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.All)
					if !ok {
						log.Fatalln("Source of All.Choices []*Choice, is not an All instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.All
			for _all := range *models.GetGongstructInstancesSet[models.All](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _all.GetName() == newSourceName.GetName() {
					newSource = _all // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of All.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		case "Choice:Choices":
			// WARNING : this form deals with the N-N association "Choice.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Choice
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Choice"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Choice)
					if !ok {
						log.Fatalln("Source of Choice.Choices []*Choice, is not an Choice instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.Choice
			for _choice := range *models.GetGongstructInstancesSet[models.Choice](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _choice.GetName() == newSourceName.GetName() {
					newSource = _choice // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Choice.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		case "ComplexType:Choices":
			// WARNING : this form deals with the N-N association "ComplexType.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Choices []*Choice, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		case "Extension:Choices":
			// WARNING : this form deals with the N-N association "Extension.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Choices []*Choice, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		case "Group:Choices":
			// WARNING : this form deals with the N-N association "Group.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Choices []*Choice, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		case "Sequence:Choices":
			// WARNING : this form deals with the N-N association "Sequence.Choices []*Choice" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Choice). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Sequence
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Sequence"
				rf.Fieldname = "Choices"
				formerAssociationSource := choice_.GongGetReverseFieldOwner(
					choiceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Sequence)
					if !ok {
						log.Fatalln("Source of Sequence.Choices []*Choice, is not an Sequence instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Choices, choice_)
					formerSource.Choices = slices.Delete(formerSource.Choices, idx, idx+1)
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
			var newSource *models.Sequence
			for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](choiceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _sequence.GetName() == newSourceName.GetName() {
					newSource = _sequence // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Sequence.Choices []*Choice, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Choices = append(newSource.Choices, choice_)
		}
	}

	// manage the suppress operation
	if choiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		choice_.Unstage(choiceFormCallback.probe.stageOfInterest)
	}

	choiceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Choice](
		choiceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if choiceFormCallback.CreationMode || choiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		choiceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(choiceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ChoiceFormCallback(
			nil,
			choiceFormCallback.probe,
			newFormGroup,
		)
		choice := new(models.Choice)
		FillUpForm(choice, newFormGroup, choiceFormCallback.probe)
		choiceFormCallback.probe.formStage.Commit()
	}

	choiceFormCallback.probe.ux_tree()
}
func __gong__New__ComplexContentFormCallback(
	complexcontent *models.ComplexContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (complexcontentFormCallback *ComplexContentFormCallback) {
	complexcontentFormCallback = new(ComplexContentFormCallback)
	complexcontentFormCallback.probe = probe
	complexcontentFormCallback.complexcontent = complexcontent
	complexcontentFormCallback.formGroup = formGroup

	complexcontentFormCallback.CreationMode = (complexcontent == nil)

	return
}

type ComplexContentFormCallback struct {
	complexcontent *models.ComplexContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complexcontentFormCallback *ComplexContentFormCallback) OnSave() {
	complexcontentFormCallback.probe.stageOfInterest.Lock()
	defer complexcontentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ComplexContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complexcontentFormCallback.probe.formStage.Checkout()

	if complexcontentFormCallback.complexcontent == nil {
		complexcontentFormCallback.complexcontent = new(models.ComplexContent).Stage(complexcontentFormCallback.probe.stageOfInterest)
	}
	complexcontent_ := complexcontentFormCallback.complexcontent
	_ = complexcontent_

	for _, formDiv := range complexcontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complexcontent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontent_.Unstage(complexcontentFormCallback.probe.stageOfInterest)
	}

	complexcontentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ComplexContent](
		complexcontentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if complexcontentFormCallback.CreationMode || complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(complexcontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexContentFormCallback(
			nil,
			complexcontentFormCallback.probe,
			newFormGroup,
		)
		complexcontent := new(models.ComplexContent)
		FillUpForm(complexcontent, newFormGroup, complexcontentFormCallback.probe)
		complexcontentFormCallback.probe.formStage.Commit()
	}

	complexcontentFormCallback.probe.ux_tree()
}
func __gong__New__ComplexTypeFormCallback(
	complextype *models.ComplexType,
	probe *Probe,
	formGroup *table.FormGroup,
) (complextypeFormCallback *ComplexTypeFormCallback) {
	complextypeFormCallback = new(ComplexTypeFormCallback)
	complextypeFormCallback.probe = probe
	complextypeFormCallback.complextype = complextype
	complextypeFormCallback.formGroup = formGroup

	complextypeFormCallback.CreationMode = (complextype == nil)

	return
}

type ComplexTypeFormCallback struct {
	complextype *models.ComplexType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complextypeFormCallback *ComplexTypeFormCallback) OnSave() {
	complextypeFormCallback.probe.stageOfInterest.Lock()
	defer complextypeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ComplexTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complextypeFormCallback.probe.formStage.Checkout()

	if complextypeFormCallback.complextype == nil {
		complextypeFormCallback.complextype = new(models.ComplexType).Stage(complextypeFormCallback.probe.stageOfInterest)
	}
	complextype_ := complextypeFormCallback.complextype
	_ = complextype_

	for _, formDiv := range complextypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complextype_.Name), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(complextype_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(complextype_.GoIdentifier), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(complextype_.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(complextype_.OuterElement), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(complextype_.Annotation), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(complextype_.NameXSD), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(complextype_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(complextype_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(complextype_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(complextype_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(complextype_.MaxOccurs), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(complextype_.Extension), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "SimpleContent":
			FormDivSelectFieldToField(&(complextype_.SimpleContent), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "ComplexContent":
			FormDivSelectFieldToField(&(complextype_.ComplexContent), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "Attributes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Attribute](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Attribute, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Attribute)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Attribute](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.Attributes = instanceSlice

		case "AttributeGroups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](complextypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AttributeGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AttributeGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					complextypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](complextypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			complextype_.AttributeGroups = instanceSlice

		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(complextype_.IsDuplicatedInXSD), formDiv)
		case "Schema:ComplexTypes":
			// WARNING : this form deals with the N-N association "Schema.ComplexTypes []*ComplexType" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ComplexType). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Schema
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Schema"
				rf.Fieldname = "ComplexTypes"
				formerAssociationSource := complextype_.GongGetReverseFieldOwner(
					complextypeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Schema)
					if !ok {
						log.Fatalln("Source of Schema.ComplexTypes []*ComplexType, is not an Schema instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ComplexTypes, complextype_)
					formerSource.ComplexTypes = slices.Delete(formerSource.ComplexTypes, idx, idx+1)
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
			var newSource *models.Schema
			for _schema := range *models.GetGongstructInstancesSet[models.Schema](complextypeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _schema.GetName() == newSourceName.GetName() {
					newSource = _schema // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Schema.ComplexTypes []*ComplexType, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ComplexTypes = append(newSource.ComplexTypes, complextype_)
		}
	}

	// manage the suppress operation
	if complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextype_.Unstage(complextypeFormCallback.probe.stageOfInterest)
	}

	complextypeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ComplexType](
		complextypeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if complextypeFormCallback.CreationMode || complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(complextypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			nil,
			complextypeFormCallback.probe,
			newFormGroup,
		)
		complextype := new(models.ComplexType)
		FillUpForm(complextype, newFormGroup, complextypeFormCallback.probe)
		complextypeFormCallback.probe.formStage.Commit()
	}

	complextypeFormCallback.probe.ux_tree()
}
func __gong__New__DocumentationFormCallback(
	documentation *models.Documentation,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentationFormCallback *DocumentationFormCallback) {
	documentationFormCallback = new(DocumentationFormCallback)
	documentationFormCallback.probe = probe
	documentationFormCallback.documentation = documentation
	documentationFormCallback.formGroup = formGroup

	documentationFormCallback.CreationMode = (documentation == nil)

	return
}

type DocumentationFormCallback struct {
	documentation *models.Documentation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentationFormCallback *DocumentationFormCallback) OnSave() {
	documentationFormCallback.probe.stageOfInterest.Lock()
	defer documentationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DocumentationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentationFormCallback.probe.formStage.Checkout()

	if documentationFormCallback.documentation == nil {
		documentationFormCallback.documentation = new(models.Documentation).Stage(documentationFormCallback.probe.stageOfInterest)
	}
	documentation_ := documentationFormCallback.documentation
	_ = documentation_

	for _, formDiv := range documentationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(documentation_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(documentation_.Text), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(documentation_.Source), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(documentation_.Lang), formDiv)
		case "Annotation:Documentations":
			// WARNING : this form deals with the N-N association "Annotation.Documentations []*Documentation" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Documentation). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Annotation
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Annotation"
				rf.Fieldname = "Documentations"
				formerAssociationSource := documentation_.GongGetReverseFieldOwner(
					documentationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Annotation)
					if !ok {
						log.Fatalln("Source of Annotation.Documentations []*Documentation, is not an Annotation instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Documentations, documentation_)
					formerSource.Documentations = slices.Delete(formerSource.Documentations, idx, idx+1)
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
			var newSource *models.Annotation
			for _annotation := range *models.GetGongstructInstancesSet[models.Annotation](documentationFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _annotation.GetName() == newSourceName.GetName() {
					newSource = _annotation // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Annotation.Documentations []*Documentation, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Documentations = append(newSource.Documentations, documentation_)
		}
	}

	// manage the suppress operation
	if documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentation_.Unstage(documentationFormCallback.probe.stageOfInterest)
	}

	documentationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Documentation](
		documentationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if documentationFormCallback.CreationMode || documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(documentationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentationFormCallback(
			nil,
			documentationFormCallback.probe,
			newFormGroup,
		)
		documentation := new(models.Documentation)
		FillUpForm(documentation, newFormGroup, documentationFormCallback.probe)
		documentationFormCallback.probe.formStage.Commit()
	}

	documentationFormCallback.probe.ux_tree()
}
func __gong__New__ElementFormCallback(
	element *models.Element,
	probe *Probe,
	formGroup *table.FormGroup,
) (elementFormCallback *ElementFormCallback) {
	elementFormCallback = new(ElementFormCallback)
	elementFormCallback.probe = probe
	elementFormCallback.element = element
	elementFormCallback.formGroup = formGroup

	elementFormCallback.CreationMode = (element == nil)

	return
}

type ElementFormCallback struct {
	element *models.Element

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (elementFormCallback *ElementFormCallback) OnSave() {
	elementFormCallback.probe.stageOfInterest.Lock()
	defer elementFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ElementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	elementFormCallback.probe.formStage.Checkout()

	if elementFormCallback.element == nil {
		elementFormCallback.element = new(models.Element).Stage(elementFormCallback.probe.stageOfInterest)
	}
	element_ := elementFormCallback.element
	_ = element_

	for _, formDiv := range elementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(element_.Name), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(element_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(element_.Depth), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(element_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(element_.GoIdentifier), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(element_.Annotation), elementFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(element_.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(element_.Type), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(element_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(element_.MaxOccurs), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(element_.Default), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(element_.Fixed), formDiv)
		case "Nillable":
			FormDivBasicFieldToField(&(element_.Nillable), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(element_.Ref), formDiv)
		case "Abstract":
			FormDivBasicFieldToField(&(element_.Abstract), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(element_.Form), formDiv)
		case "Block":
			FormDivBasicFieldToField(&(element_.Block), formDiv)
		case "Final":
			FormDivBasicFieldToField(&(element_.Final), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(element_.SimpleType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "ComplexType":
			FormDivSelectFieldToField(&(element_.ComplexType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](elementFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					elementFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](elementFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			element_.Groups = instanceSlice

		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(element_.IsDuplicatedInXSD), formDiv)
		case "All:Elements":
			// WARNING : this form deals with the N-N association "All.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.All
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "All"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.All)
					if !ok {
						log.Fatalln("Source of All.Elements []*Element, is not an All instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.All
			for _all := range *models.GetGongstructInstancesSet[models.All](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _all.GetName() == newSourceName.GetName() {
					newSource = _all // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of All.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "Choice:Elements":
			// WARNING : this form deals with the N-N association "Choice.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Choice
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Choice"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Choice)
					if !ok {
						log.Fatalln("Source of Choice.Elements []*Element, is not an Choice instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.Choice
			for _choice := range *models.GetGongstructInstancesSet[models.Choice](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _choice.GetName() == newSourceName.GetName() {
					newSource = _choice // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Choice.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "ComplexType:Elements":
			// WARNING : this form deals with the N-N association "ComplexType.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Elements []*Element, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "Extension:Elements":
			// WARNING : this form deals with the N-N association "Extension.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Elements []*Element, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "Group:Elements":
			// WARNING : this form deals with the N-N association "Group.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Elements []*Element, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "Schema:Elements":
			// WARNING : this form deals with the N-N association "Schema.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Schema
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Schema"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Schema)
					if !ok {
						log.Fatalln("Source of Schema.Elements []*Element, is not an Schema instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.Schema
			for _schema := range *models.GetGongstructInstancesSet[models.Schema](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _schema.GetName() == newSourceName.GetName() {
					newSource = _schema // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Schema.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		case "Sequence:Elements":
			// WARNING : this form deals with the N-N association "Sequence.Elements []*Element" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Element). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Sequence
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Sequence"
				rf.Fieldname = "Elements"
				formerAssociationSource := element_.GongGetReverseFieldOwner(
					elementFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Sequence)
					if !ok {
						log.Fatalln("Source of Sequence.Elements []*Element, is not an Sequence instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Elements, element_)
					formerSource.Elements = slices.Delete(formerSource.Elements, idx, idx+1)
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
			var newSource *models.Sequence
			for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](elementFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _sequence.GetName() == newSourceName.GetName() {
					newSource = _sequence // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Sequence.Elements []*Element, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Elements = append(newSource.Elements, element_)
		}
	}

	// manage the suppress operation
	if elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		element_.Unstage(elementFormCallback.probe.stageOfInterest)
	}

	elementFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Element](
		elementFormCallback.probe,
	)

	// display a new form by reset the form stage
	if elementFormCallback.CreationMode || elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(elementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			elementFormCallback.probe,
			newFormGroup,
		)
		element := new(models.Element)
		FillUpForm(element, newFormGroup, elementFormCallback.probe)
		elementFormCallback.probe.formStage.Commit()
	}

	elementFormCallback.probe.ux_tree()
}
func __gong__New__EnumerationFormCallback(
	enumeration *models.Enumeration,
	probe *Probe,
	formGroup *table.FormGroup,
) (enumerationFormCallback *EnumerationFormCallback) {
	enumerationFormCallback = new(EnumerationFormCallback)
	enumerationFormCallback.probe = probe
	enumerationFormCallback.enumeration = enumeration
	enumerationFormCallback.formGroup = formGroup

	enumerationFormCallback.CreationMode = (enumeration == nil)

	return
}

type EnumerationFormCallback struct {
	enumeration *models.Enumeration

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enumerationFormCallback *EnumerationFormCallback) OnSave() {
	enumerationFormCallback.probe.stageOfInterest.Lock()
	defer enumerationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EnumerationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enumerationFormCallback.probe.formStage.Checkout()

	if enumerationFormCallback.enumeration == nil {
		enumerationFormCallback.enumeration = new(models.Enumeration).Stage(enumerationFormCallback.probe.stageOfInterest)
	}
	enumeration_ := enumerationFormCallback.enumeration
	_ = enumeration_

	for _, formDiv := range enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enumeration_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(enumeration_.Annotation), enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(enumeration_.Value), formDiv)
		case "Restriction:Enumerations":
			// WARNING : this form deals with the N-N association "Restriction.Enumerations []*Enumeration" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Enumeration). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Restriction
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Restriction"
				rf.Fieldname = "Enumerations"
				formerAssociationSource := enumeration_.GongGetReverseFieldOwner(
					enumerationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Restriction)
					if !ok {
						log.Fatalln("Source of Restriction.Enumerations []*Enumeration, is not an Restriction instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Enumerations, enumeration_)
					formerSource.Enumerations = slices.Delete(formerSource.Enumerations, idx, idx+1)
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
			var newSource *models.Restriction
			for _restriction := range *models.GetGongstructInstancesSet[models.Restriction](enumerationFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _restriction.GetName() == newSourceName.GetName() {
					newSource = _restriction // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Restriction.Enumerations []*Enumeration, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Enumerations = append(newSource.Enumerations, enumeration_)
		}
	}

	// manage the suppress operation
	if enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumeration_.Unstage(enumerationFormCallback.probe.stageOfInterest)
	}

	enumerationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Enumeration](
		enumerationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if enumerationFormCallback.CreationMode || enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EnumerationFormCallback(
			nil,
			enumerationFormCallback.probe,
			newFormGroup,
		)
		enumeration := new(models.Enumeration)
		FillUpForm(enumeration, newFormGroup, enumerationFormCallback.probe)
		enumerationFormCallback.probe.formStage.Commit()
	}

	enumerationFormCallback.probe.ux_tree()
}
func __gong__New__ExtensionFormCallback(
	extension *models.Extension,
	probe *Probe,
	formGroup *table.FormGroup,
) (extensionFormCallback *ExtensionFormCallback) {
	extensionFormCallback = new(ExtensionFormCallback)
	extensionFormCallback.probe = probe
	extensionFormCallback.extension = extension
	extensionFormCallback.formGroup = formGroup

	extensionFormCallback.CreationMode = (extension == nil)

	return
}

type ExtensionFormCallback struct {
	extension *models.Extension

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (extensionFormCallback *ExtensionFormCallback) OnSave() {
	extensionFormCallback.probe.stageOfInterest.Lock()
	defer extensionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ExtensionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	extensionFormCallback.probe.formStage.Checkout()

	if extensionFormCallback.extension == nil {
		extensionFormCallback.extension = new(models.Extension).Stage(extensionFormCallback.probe.stageOfInterest)
	}
	extension_ := extensionFormCallback.extension
	_ = extension_

	for _, formDiv := range extensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(extension_.Name), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(extension_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(extension_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(extension_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(extension_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(extension_.MaxOccurs), formDiv)
		case "Base":
			FormDivBasicFieldToField(&(extension_.Base), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(extension_.Ref), formDiv)
		case "Attributes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Attribute](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Attribute, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Attribute)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Attribute](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.Attributes = instanceSlice

		case "AttributeGroups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](extensionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AttributeGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AttributeGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					extensionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](extensionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			extension_.AttributeGroups = instanceSlice

		}
	}

	// manage the suppress operation
	if extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extension_.Unstage(extensionFormCallback.probe.stageOfInterest)
	}

	extensionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Extension](
		extensionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if extensionFormCallback.CreationMode || extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(extensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExtensionFormCallback(
			nil,
			extensionFormCallback.probe,
			newFormGroup,
		)
		extension := new(models.Extension)
		FillUpForm(extension, newFormGroup, extensionFormCallback.probe)
		extensionFormCallback.probe.formStage.Commit()
	}

	extensionFormCallback.probe.ux_tree()
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.probe = probe
	groupFormCallback.group = group
	groupFormCallback.formGroup = formGroup

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {
	groupFormCallback.probe.stageOfInterest.Lock()
	defer groupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.probe.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.probe.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	for _, formDiv := range groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(group_.Annotation), groupFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(group_.NameXSD), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(group_.Ref), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(group_.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(group_.OuterElement), groupFormCallback.probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(group_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(group_.GoIdentifier), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(group_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(group_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(group_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(group_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(group_.MaxOccurs), formDiv)
		case "All:Groups":
			// WARNING : this form deals with the N-N association "All.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.All
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "All"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.All)
					if !ok {
						log.Fatalln("Source of All.Groups []*Group, is not an All instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.All
			for _all := range *models.GetGongstructInstancesSet[models.All](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _all.GetName() == newSourceName.GetName() {
					newSource = _all // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of All.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Choice:Groups":
			// WARNING : this form deals with the N-N association "Choice.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Choice
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Choice"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Choice)
					if !ok {
						log.Fatalln("Source of Choice.Groups []*Group, is not an Choice instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Choice
			for _choice := range *models.GetGongstructInstancesSet[models.Choice](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _choice.GetName() == newSourceName.GetName() {
					newSource = _choice // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Choice.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "ComplexType:Groups":
			// WARNING : this form deals with the N-N association "ComplexType.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Groups []*Group, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Element:Groups":
			// WARNING : this form deals with the N-N association "Element.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Element
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Element"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Element)
					if !ok {
						log.Fatalln("Source of Element.Groups []*Group, is not an Element instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Element
			for _element := range *models.GetGongstructInstancesSet[models.Element](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _element.GetName() == newSourceName.GetName() {
					newSource = _element // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Element.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Extension:Groups":
			// WARNING : this form deals with the N-N association "Extension.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Groups []*Group, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Group:Groups":
			// WARNING : this form deals with the N-N association "Group.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Groups []*Group, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Schema:Groups":
			// WARNING : this form deals with the N-N association "Schema.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Schema
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Schema"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Schema)
					if !ok {
						log.Fatalln("Source of Schema.Groups []*Group, is not an Schema instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Schema
			for _schema := range *models.GetGongstructInstancesSet[models.Schema](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _schema.GetName() == newSourceName.GetName() {
					newSource = _schema // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Schema.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		case "Sequence:Groups":
			// WARNING : this form deals with the N-N association "Sequence.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Sequence
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Sequence"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Sequence)
					if !ok {
						log.Fatalln("Source of Sequence.Groups []*Group, is not an Sequence instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
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
			var newSource *models.Sequence
			for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _sequence.GetName() == newSourceName.GetName() {
					newSource = _sequence // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Sequence.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Group](
		groupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			groupFormCallback.probe,
			newFormGroup,
		)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.probe)
		groupFormCallback.probe.formStage.Commit()
	}

	groupFormCallback.probe.ux_tree()
}
func __gong__New__LengthFormCallback(
	length *models.Length,
	probe *Probe,
	formGroup *table.FormGroup,
) (lengthFormCallback *LengthFormCallback) {
	lengthFormCallback = new(LengthFormCallback)
	lengthFormCallback.probe = probe
	lengthFormCallback.length = length
	lengthFormCallback.formGroup = formGroup

	lengthFormCallback.CreationMode = (length == nil)

	return
}

type LengthFormCallback struct {
	length *models.Length

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lengthFormCallback *LengthFormCallback) OnSave() {
	lengthFormCallback.probe.stageOfInterest.Lock()
	defer lengthFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lengthFormCallback.probe.formStage.Checkout()

	if lengthFormCallback.length == nil {
		lengthFormCallback.length = new(models.Length).Stage(lengthFormCallback.probe.stageOfInterest)
	}
	length_ := lengthFormCallback.length
	_ = length_

	for _, formDiv := range lengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(length_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(length_.Annotation), lengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(length_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if lengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		length_.Unstage(lengthFormCallback.probe.stageOfInterest)
	}

	lengthFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Length](
		lengthFormCallback.probe,
	)

	// display a new form by reset the form stage
	if lengthFormCallback.CreationMode || lengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(lengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LengthFormCallback(
			nil,
			lengthFormCallback.probe,
			newFormGroup,
		)
		length := new(models.Length)
		FillUpForm(length, newFormGroup, lengthFormCallback.probe)
		lengthFormCallback.probe.formStage.Commit()
	}

	lengthFormCallback.probe.ux_tree()
}
func __gong__New__MaxInclusiveFormCallback(
	maxinclusive *models.MaxInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (maxinclusiveFormCallback *MaxInclusiveFormCallback) {
	maxinclusiveFormCallback = new(MaxInclusiveFormCallback)
	maxinclusiveFormCallback.probe = probe
	maxinclusiveFormCallback.maxinclusive = maxinclusive
	maxinclusiveFormCallback.formGroup = formGroup

	maxinclusiveFormCallback.CreationMode = (maxinclusive == nil)

	return
}

type MaxInclusiveFormCallback struct {
	maxinclusive *models.MaxInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (maxinclusiveFormCallback *MaxInclusiveFormCallback) OnSave() {
	maxinclusiveFormCallback.probe.stageOfInterest.Lock()
	defer maxinclusiveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MaxInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	maxinclusiveFormCallback.probe.formStage.Checkout()

	if maxinclusiveFormCallback.maxinclusive == nil {
		maxinclusiveFormCallback.maxinclusive = new(models.MaxInclusive).Stage(maxinclusiveFormCallback.probe.stageOfInterest)
	}
	maxinclusive_ := maxinclusiveFormCallback.maxinclusive
	_ = maxinclusive_

	for _, formDiv := range maxinclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(maxinclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(maxinclusive_.Annotation), maxinclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(maxinclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusive_.Unstage(maxinclusiveFormCallback.probe.stageOfInterest)
	}

	maxinclusiveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MaxInclusive](
		maxinclusiveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if maxinclusiveFormCallback.CreationMode || maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(maxinclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			nil,
			maxinclusiveFormCallback.probe,
			newFormGroup,
		)
		maxinclusive := new(models.MaxInclusive)
		FillUpForm(maxinclusive, newFormGroup, maxinclusiveFormCallback.probe)
		maxinclusiveFormCallback.probe.formStage.Commit()
	}

	maxinclusiveFormCallback.probe.ux_tree()
}
func __gong__New__MaxLengthFormCallback(
	maxlength *models.MaxLength,
	probe *Probe,
	formGroup *table.FormGroup,
) (maxlengthFormCallback *MaxLengthFormCallback) {
	maxlengthFormCallback = new(MaxLengthFormCallback)
	maxlengthFormCallback.probe = probe
	maxlengthFormCallback.maxlength = maxlength
	maxlengthFormCallback.formGroup = formGroup

	maxlengthFormCallback.CreationMode = (maxlength == nil)

	return
}

type MaxLengthFormCallback struct {
	maxlength *models.MaxLength

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (maxlengthFormCallback *MaxLengthFormCallback) OnSave() {
	maxlengthFormCallback.probe.stageOfInterest.Lock()
	defer maxlengthFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MaxLengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	maxlengthFormCallback.probe.formStage.Checkout()

	if maxlengthFormCallback.maxlength == nil {
		maxlengthFormCallback.maxlength = new(models.MaxLength).Stage(maxlengthFormCallback.probe.stageOfInterest)
	}
	maxlength_ := maxlengthFormCallback.maxlength
	_ = maxlength_

	for _, formDiv := range maxlengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(maxlength_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(maxlength_.Annotation), maxlengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(maxlength_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if maxlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxlength_.Unstage(maxlengthFormCallback.probe.stageOfInterest)
	}

	maxlengthFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MaxLength](
		maxlengthFormCallback.probe,
	)

	// display a new form by reset the form stage
	if maxlengthFormCallback.CreationMode || maxlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxlengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(maxlengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MaxLengthFormCallback(
			nil,
			maxlengthFormCallback.probe,
			newFormGroup,
		)
		maxlength := new(models.MaxLength)
		FillUpForm(maxlength, newFormGroup, maxlengthFormCallback.probe)
		maxlengthFormCallback.probe.formStage.Commit()
	}

	maxlengthFormCallback.probe.ux_tree()
}
func __gong__New__MinInclusiveFormCallback(
	mininclusive *models.MinInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (mininclusiveFormCallback *MinInclusiveFormCallback) {
	mininclusiveFormCallback = new(MinInclusiveFormCallback)
	mininclusiveFormCallback.probe = probe
	mininclusiveFormCallback.mininclusive = mininclusive
	mininclusiveFormCallback.formGroup = formGroup

	mininclusiveFormCallback.CreationMode = (mininclusive == nil)

	return
}

type MinInclusiveFormCallback struct {
	mininclusive *models.MinInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mininclusiveFormCallback *MinInclusiveFormCallback) OnSave() {
	mininclusiveFormCallback.probe.stageOfInterest.Lock()
	defer mininclusiveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MinInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mininclusiveFormCallback.probe.formStage.Checkout()

	if mininclusiveFormCallback.mininclusive == nil {
		mininclusiveFormCallback.mininclusive = new(models.MinInclusive).Stage(mininclusiveFormCallback.probe.stageOfInterest)
	}
	mininclusive_ := mininclusiveFormCallback.mininclusive
	_ = mininclusive_

	for _, formDiv := range mininclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mininclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(mininclusive_.Annotation), mininclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(mininclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusive_.Unstage(mininclusiveFormCallback.probe.stageOfInterest)
	}

	mininclusiveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MinInclusive](
		mininclusiveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if mininclusiveFormCallback.CreationMode || mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(mininclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			nil,
			mininclusiveFormCallback.probe,
			newFormGroup,
		)
		mininclusive := new(models.MinInclusive)
		FillUpForm(mininclusive, newFormGroup, mininclusiveFormCallback.probe)
		mininclusiveFormCallback.probe.formStage.Commit()
	}

	mininclusiveFormCallback.probe.ux_tree()
}
func __gong__New__MinLengthFormCallback(
	minlength *models.MinLength,
	probe *Probe,
	formGroup *table.FormGroup,
) (minlengthFormCallback *MinLengthFormCallback) {
	minlengthFormCallback = new(MinLengthFormCallback)
	minlengthFormCallback.probe = probe
	minlengthFormCallback.minlength = minlength
	minlengthFormCallback.formGroup = formGroup

	minlengthFormCallback.CreationMode = (minlength == nil)

	return
}

type MinLengthFormCallback struct {
	minlength *models.MinLength

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (minlengthFormCallback *MinLengthFormCallback) OnSave() {
	minlengthFormCallback.probe.stageOfInterest.Lock()
	defer minlengthFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MinLengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	minlengthFormCallback.probe.formStage.Checkout()

	if minlengthFormCallback.minlength == nil {
		minlengthFormCallback.minlength = new(models.MinLength).Stage(minlengthFormCallback.probe.stageOfInterest)
	}
	minlength_ := minlengthFormCallback.minlength
	_ = minlength_

	for _, formDiv := range minlengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(minlength_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(minlength_.Annotation), minlengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(minlength_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if minlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		minlength_.Unstage(minlengthFormCallback.probe.stageOfInterest)
	}

	minlengthFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MinLength](
		minlengthFormCallback.probe,
	)

	// display a new form by reset the form stage
	if minlengthFormCallback.CreationMode || minlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		minlengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(minlengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MinLengthFormCallback(
			nil,
			minlengthFormCallback.probe,
			newFormGroup,
		)
		minlength := new(models.MinLength)
		FillUpForm(minlength, newFormGroup, minlengthFormCallback.probe)
		minlengthFormCallback.probe.formStage.Commit()
	}

	minlengthFormCallback.probe.ux_tree()
}
func __gong__New__PatternFormCallback(
	pattern *models.Pattern,
	probe *Probe,
	formGroup *table.FormGroup,
) (patternFormCallback *PatternFormCallback) {
	patternFormCallback = new(PatternFormCallback)
	patternFormCallback.probe = probe
	patternFormCallback.pattern = pattern
	patternFormCallback.formGroup = formGroup

	patternFormCallback.CreationMode = (pattern == nil)

	return
}

type PatternFormCallback struct {
	pattern *models.Pattern

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (patternFormCallback *PatternFormCallback) OnSave() {
	patternFormCallback.probe.stageOfInterest.Lock()
	defer patternFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PatternFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	patternFormCallback.probe.formStage.Checkout()

	if patternFormCallback.pattern == nil {
		patternFormCallback.pattern = new(models.Pattern).Stage(patternFormCallback.probe.stageOfInterest)
	}
	pattern_ := patternFormCallback.pattern
	_ = pattern_

	for _, formDiv := range patternFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pattern_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(pattern_.Annotation), patternFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(pattern_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pattern_.Unstage(patternFormCallback.probe.stageOfInterest)
	}

	patternFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Pattern](
		patternFormCallback.probe,
	)

	// display a new form by reset the form stage
	if patternFormCallback.CreationMode || patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		patternFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(patternFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PatternFormCallback(
			nil,
			patternFormCallback.probe,
			newFormGroup,
		)
		pattern := new(models.Pattern)
		FillUpForm(pattern, newFormGroup, patternFormCallback.probe)
		patternFormCallback.probe.formStage.Commit()
	}

	patternFormCallback.probe.ux_tree()
}
func __gong__New__RestrictionFormCallback(
	restriction *models.Restriction,
	probe *Probe,
	formGroup *table.FormGroup,
) (restrictionFormCallback *RestrictionFormCallback) {
	restrictionFormCallback = new(RestrictionFormCallback)
	restrictionFormCallback.probe = probe
	restrictionFormCallback.restriction = restriction
	restrictionFormCallback.formGroup = formGroup

	restrictionFormCallback.CreationMode = (restriction == nil)

	return
}

type RestrictionFormCallback struct {
	restriction *models.Restriction

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (restrictionFormCallback *RestrictionFormCallback) OnSave() {
	restrictionFormCallback.probe.stageOfInterest.Lock()
	defer restrictionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RestrictionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	restrictionFormCallback.probe.formStage.Checkout()

	if restrictionFormCallback.restriction == nil {
		restrictionFormCallback.restriction = new(models.Restriction).Stage(restrictionFormCallback.probe.stageOfInterest)
	}
	restriction_ := restrictionFormCallback.restriction
	_ = restriction_

	for _, formDiv := range restrictionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(restriction_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(restriction_.Annotation), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Base":
			FormDivBasicFieldToField(&(restriction_.Base), formDiv)
		case "Enumerations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Enumeration](restrictionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Enumeration, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Enumeration)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					restrictionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Enumeration](restrictionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			restriction_.Enumerations = instanceSlice

		case "MinInclusive":
			FormDivSelectFieldToField(&(restriction_.MinInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MaxInclusive":
			FormDivSelectFieldToField(&(restriction_.MaxInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Pattern":
			FormDivSelectFieldToField(&(restriction_.Pattern), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "WhiteSpace":
			FormDivSelectFieldToField(&(restriction_.WhiteSpace), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MinLength":
			FormDivSelectFieldToField(&(restriction_.MinLength), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MaxLength":
			FormDivSelectFieldToField(&(restriction_.MaxLength), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Length":
			FormDivSelectFieldToField(&(restriction_.Length), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "TotalDigit":
			FormDivSelectFieldToField(&(restriction_.TotalDigit), restrictionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restriction_.Unstage(restrictionFormCallback.probe.stageOfInterest)
	}

	restrictionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Restriction](
		restrictionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if restrictionFormCallback.CreationMode || restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restrictionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(restrictionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RestrictionFormCallback(
			nil,
			restrictionFormCallback.probe,
			newFormGroup,
		)
		restriction := new(models.Restriction)
		FillUpForm(restriction, newFormGroup, restrictionFormCallback.probe)
		restrictionFormCallback.probe.formStage.Commit()
	}

	restrictionFormCallback.probe.ux_tree()
}
func __gong__New__SchemaFormCallback(
	schema *models.Schema,
	probe *Probe,
	formGroup *table.FormGroup,
) (schemaFormCallback *SchemaFormCallback) {
	schemaFormCallback = new(SchemaFormCallback)
	schemaFormCallback.probe = probe
	schemaFormCallback.schema = schema
	schemaFormCallback.formGroup = formGroup

	schemaFormCallback.CreationMode = (schema == nil)

	return
}

type SchemaFormCallback struct {
	schema *models.Schema

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (schemaFormCallback *SchemaFormCallback) OnSave() {
	schemaFormCallback.probe.stageOfInterest.Lock()
	defer schemaFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SchemaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	schemaFormCallback.probe.formStage.Checkout()

	if schemaFormCallback.schema == nil {
		schemaFormCallback.schema = new(models.Schema).Stage(schemaFormCallback.probe.stageOfInterest)
	}
	schema_ := schemaFormCallback.schema
	_ = schema_

	for _, formDiv := range schemaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(schema_.Name), formDiv)
		case "Xs":
			FormDivBasicFieldToField(&(schema_.Xs), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(schema_.Annotation), schemaFormCallback.probe.stageOfInterest, formDiv)
		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](schemaFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					schemaFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](schemaFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			schema_.Elements = instanceSlice

		case "SimpleTypes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SimpleType](schemaFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SimpleType, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SimpleType)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					schemaFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SimpleType](schemaFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			schema_.SimpleTypes = instanceSlice

		case "ComplexTypes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](schemaFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ComplexType, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ComplexType)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					schemaFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](schemaFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			schema_.ComplexTypes = instanceSlice

		case "AttributeGroups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](schemaFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AttributeGroup, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AttributeGroup)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					schemaFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](schemaFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			schema_.AttributeGroups = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](schemaFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					schemaFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](schemaFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			schema_.Groups = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(schema_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(schema_.Depth), formDiv)
		}
	}

	// manage the suppress operation
	if schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schema_.Unstage(schemaFormCallback.probe.stageOfInterest)
	}

	schemaFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Schema](
		schemaFormCallback.probe,
	)

	// display a new form by reset the form stage
	if schemaFormCallback.CreationMode || schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schemaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(schemaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SchemaFormCallback(
			nil,
			schemaFormCallback.probe,
			newFormGroup,
		)
		schema := new(models.Schema)
		FillUpForm(schema, newFormGroup, schemaFormCallback.probe)
		schemaFormCallback.probe.formStage.Commit()
	}

	schemaFormCallback.probe.ux_tree()
}
func __gong__New__SequenceFormCallback(
	sequence *models.Sequence,
	probe *Probe,
	formGroup *table.FormGroup,
) (sequenceFormCallback *SequenceFormCallback) {
	sequenceFormCallback = new(SequenceFormCallback)
	sequenceFormCallback.probe = probe
	sequenceFormCallback.sequence = sequence
	sequenceFormCallback.formGroup = formGroup

	sequenceFormCallback.CreationMode = (sequence == nil)

	return
}

type SequenceFormCallback struct {
	sequence *models.Sequence

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sequenceFormCallback *SequenceFormCallback) OnSave() {
	sequenceFormCallback.probe.stageOfInterest.Lock()
	defer sequenceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SequenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sequenceFormCallback.probe.formStage.Checkout()

	if sequenceFormCallback.sequence == nil {
		sequenceFormCallback.sequence = new(models.Sequence).Stage(sequenceFormCallback.probe.stageOfInterest)
	}
	sequence_ := sequenceFormCallback.sequence
	_ = sequence_

	for _, formDiv := range sequenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sequence_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(sequence_.Annotation), sequenceFormCallback.probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(sequence_.OuterElementName), formDiv)
		case "Sequences":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](sequenceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Sequence, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Sequence)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					sequenceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](sequenceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			sequence_.Sequences = instanceSlice

		case "Alls":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.All](sequenceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.All, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.All)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					sequenceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.All](sequenceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			sequence_.Alls = instanceSlice

		case "Choices":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Choice](sequenceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Choice, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Choice)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					sequenceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](sequenceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			sequence_.Choices = instanceSlice

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](sequenceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					sequenceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](sequenceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			sequence_.Groups = instanceSlice

		case "Elements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Element](sequenceFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Element, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Element)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					sequenceFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Element](sequenceFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			sequence_.Elements = instanceSlice

		case "Order":
			FormDivBasicFieldToField(&(sequence_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(sequence_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(sequence_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(sequence_.MaxOccurs), formDiv)
		case "All:Sequences":
			// WARNING : this form deals with the N-N association "All.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.All
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "All"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.All)
					if !ok {
						log.Fatalln("Source of All.Sequences []*Sequence, is not an All instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.All
			for _all := range *models.GetGongstructInstancesSet[models.All](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _all.GetName() == newSourceName.GetName() {
					newSource = _all // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of All.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		case "Choice:Sequences":
			// WARNING : this form deals with the N-N association "Choice.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Choice
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Choice"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Choice)
					if !ok {
						log.Fatalln("Source of Choice.Sequences []*Sequence, is not an Choice instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.Choice
			for _choice := range *models.GetGongstructInstancesSet[models.Choice](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _choice.GetName() == newSourceName.GetName() {
					newSource = _choice // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Choice.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		case "ComplexType:Sequences":
			// WARNING : this form deals with the N-N association "ComplexType.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.ComplexType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "ComplexType"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.ComplexType)
					if !ok {
						log.Fatalln("Source of ComplexType.Sequences []*Sequence, is not an ComplexType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.ComplexType
			for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _complextype.GetName() == newSourceName.GetName() {
					newSource = _complextype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of ComplexType.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		case "Extension:Sequences":
			// WARNING : this form deals with the N-N association "Extension.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Extension
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Extension"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Extension)
					if !ok {
						log.Fatalln("Source of Extension.Sequences []*Sequence, is not an Extension instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.Extension
			for _extension := range *models.GetGongstructInstancesSet[models.Extension](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _extension.GetName() == newSourceName.GetName() {
					newSource = _extension // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Extension.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		case "Group:Sequences":
			// WARNING : this form deals with the N-N association "Group.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Sequences []*Sequence, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		case "Sequence:Sequences":
			// WARNING : this form deals with the N-N association "Sequence.Sequences []*Sequence" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Sequence). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Sequence
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Sequence"
				rf.Fieldname = "Sequences"
				formerAssociationSource := sequence_.GongGetReverseFieldOwner(
					sequenceFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Sequence)
					if !ok {
						log.Fatalln("Source of Sequence.Sequences []*Sequence, is not an Sequence instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sequences, sequence_)
					formerSource.Sequences = slices.Delete(formerSource.Sequences, idx, idx+1)
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
			var newSource *models.Sequence
			for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](sequenceFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _sequence.GetName() == newSourceName.GetName() {
					newSource = _sequence // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Sequence.Sequences []*Sequence, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sequences = append(newSource.Sequences, sequence_)
		}
	}

	// manage the suppress operation
	if sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequence_.Unstage(sequenceFormCallback.probe.stageOfInterest)
	}

	sequenceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Sequence](
		sequenceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if sequenceFormCallback.CreationMode || sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(sequenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SequenceFormCallback(
			nil,
			sequenceFormCallback.probe,
			newFormGroup,
		)
		sequence := new(models.Sequence)
		FillUpForm(sequence, newFormGroup, sequenceFormCallback.probe)
		sequenceFormCallback.probe.formStage.Commit()
	}

	sequenceFormCallback.probe.ux_tree()
}
func __gong__New__SimpleContentFormCallback(
	simplecontent *models.SimpleContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (simplecontentFormCallback *SimpleContentFormCallback) {
	simplecontentFormCallback = new(SimpleContentFormCallback)
	simplecontentFormCallback.probe = probe
	simplecontentFormCallback.simplecontent = simplecontent
	simplecontentFormCallback.formGroup = formGroup

	simplecontentFormCallback.CreationMode = (simplecontent == nil)

	return
}

type SimpleContentFormCallback struct {
	simplecontent *models.SimpleContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (simplecontentFormCallback *SimpleContentFormCallback) OnSave() {
	simplecontentFormCallback.probe.stageOfInterest.Lock()
	defer simplecontentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SimpleContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	simplecontentFormCallback.probe.formStage.Checkout()

	if simplecontentFormCallback.simplecontent == nil {
		simplecontentFormCallback.simplecontent = new(models.SimpleContent).Stage(simplecontentFormCallback.probe.stageOfInterest)
	}
	simplecontent_ := simplecontentFormCallback.simplecontent
	_ = simplecontent_

	for _, formDiv := range simplecontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(simplecontent_.Name), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(simplecontent_.Extension), simplecontentFormCallback.probe.stageOfInterest, formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(simplecontent_.Restriction), simplecontentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if simplecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simplecontent_.Unstage(simplecontentFormCallback.probe.stageOfInterest)
	}

	simplecontentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SimpleContent](
		simplecontentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if simplecontentFormCallback.CreationMode || simplecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simplecontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(simplecontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SimpleContentFormCallback(
			nil,
			simplecontentFormCallback.probe,
			newFormGroup,
		)
		simplecontent := new(models.SimpleContent)
		FillUpForm(simplecontent, newFormGroup, simplecontentFormCallback.probe)
		simplecontentFormCallback.probe.formStage.Commit()
	}

	simplecontentFormCallback.probe.ux_tree()
}
func __gong__New__SimpleTypeFormCallback(
	simpletype *models.SimpleType,
	probe *Probe,
	formGroup *table.FormGroup,
) (simpletypeFormCallback *SimpleTypeFormCallback) {
	simpletypeFormCallback = new(SimpleTypeFormCallback)
	simpletypeFormCallback.probe = probe
	simpletypeFormCallback.simpletype = simpletype
	simpletypeFormCallback.formGroup = formGroup

	simpletypeFormCallback.CreationMode = (simpletype == nil)

	return
}

type SimpleTypeFormCallback struct {
	simpletype *models.SimpleType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (simpletypeFormCallback *SimpleTypeFormCallback) OnSave() {
	simpletypeFormCallback.probe.stageOfInterest.Lock()
	defer simpletypeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SimpleTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	simpletypeFormCallback.probe.formStage.Checkout()

	if simpletypeFormCallback.simpletype == nil {
		simpletypeFormCallback.simpletype = new(models.SimpleType).Stage(simpletypeFormCallback.probe.stageOfInterest)
	}
	simpletype_ := simpletypeFormCallback.simpletype
	_ = simpletype_

	for _, formDiv := range simpletypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(simpletype_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(simpletype_.Annotation), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(simpletype_.NameXSD), formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(simpletype_.Restriction), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "Union":
			FormDivSelectFieldToField(&(simpletype_.Union), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "Order":
			FormDivBasicFieldToField(&(simpletype_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(simpletype_.Depth), formDiv)
		case "Schema:SimpleTypes":
			// WARNING : this form deals with the N-N association "Schema.SimpleTypes []*SimpleType" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SimpleType). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Schema
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Schema"
				rf.Fieldname = "SimpleTypes"
				formerAssociationSource := simpletype_.GongGetReverseFieldOwner(
					simpletypeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Schema)
					if !ok {
						log.Fatalln("Source of Schema.SimpleTypes []*SimpleType, is not an Schema instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SimpleTypes, simpletype_)
					formerSource.SimpleTypes = slices.Delete(formerSource.SimpleTypes, idx, idx+1)
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
			var newSource *models.Schema
			for _schema := range *models.GetGongstructInstancesSet[models.Schema](simpletypeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _schema.GetName() == newSourceName.GetName() {
					newSource = _schema // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Schema.SimpleTypes []*SimpleType, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SimpleTypes = append(newSource.SimpleTypes, simpletype_)
		}
	}

	// manage the suppress operation
	if simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletype_.Unstage(simpletypeFormCallback.probe.stageOfInterest)
	}

	simpletypeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SimpleType](
		simpletypeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if simpletypeFormCallback.CreationMode || simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(simpletypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			nil,
			simpletypeFormCallback.probe,
			newFormGroup,
		)
		simpletype := new(models.SimpleType)
		FillUpForm(simpletype, newFormGroup, simpletypeFormCallback.probe)
		simpletypeFormCallback.probe.formStage.Commit()
	}

	simpletypeFormCallback.probe.ux_tree()
}
func __gong__New__TotalDigitFormCallback(
	totaldigit *models.TotalDigit,
	probe *Probe,
	formGroup *table.FormGroup,
) (totaldigitFormCallback *TotalDigitFormCallback) {
	totaldigitFormCallback = new(TotalDigitFormCallback)
	totaldigitFormCallback.probe = probe
	totaldigitFormCallback.totaldigit = totaldigit
	totaldigitFormCallback.formGroup = formGroup

	totaldigitFormCallback.CreationMode = (totaldigit == nil)

	return
}

type TotalDigitFormCallback struct {
	totaldigit *models.TotalDigit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (totaldigitFormCallback *TotalDigitFormCallback) OnSave() {
	totaldigitFormCallback.probe.stageOfInterest.Lock()
	defer totaldigitFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TotalDigitFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	totaldigitFormCallback.probe.formStage.Checkout()

	if totaldigitFormCallback.totaldigit == nil {
		totaldigitFormCallback.totaldigit = new(models.TotalDigit).Stage(totaldigitFormCallback.probe.stageOfInterest)
	}
	totaldigit_ := totaldigitFormCallback.totaldigit
	_ = totaldigit_

	for _, formDiv := range totaldigitFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(totaldigit_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(totaldigit_.Annotation), totaldigitFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(totaldigit_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if totaldigitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		totaldigit_.Unstage(totaldigitFormCallback.probe.stageOfInterest)
	}

	totaldigitFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TotalDigit](
		totaldigitFormCallback.probe,
	)

	// display a new form by reset the form stage
	if totaldigitFormCallback.CreationMode || totaldigitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		totaldigitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(totaldigitFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TotalDigitFormCallback(
			nil,
			totaldigitFormCallback.probe,
			newFormGroup,
		)
		totaldigit := new(models.TotalDigit)
		FillUpForm(totaldigit, newFormGroup, totaldigitFormCallback.probe)
		totaldigitFormCallback.probe.formStage.Commit()
	}

	totaldigitFormCallback.probe.ux_tree()
}
func __gong__New__UnionFormCallback(
	union *models.Union,
	probe *Probe,
	formGroup *table.FormGroup,
) (unionFormCallback *UnionFormCallback) {
	unionFormCallback = new(UnionFormCallback)
	unionFormCallback.probe = probe
	unionFormCallback.union = union
	unionFormCallback.formGroup = formGroup

	unionFormCallback.CreationMode = (union == nil)

	return
}

type UnionFormCallback struct {
	union *models.Union

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (unionFormCallback *UnionFormCallback) OnSave() {
	unionFormCallback.probe.stageOfInterest.Lock()
	defer unionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("UnionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	unionFormCallback.probe.formStage.Checkout()

	if unionFormCallback.union == nil {
		unionFormCallback.union = new(models.Union).Stage(unionFormCallback.probe.stageOfInterest)
	}
	union_ := unionFormCallback.union
	_ = union_

	for _, formDiv := range unionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(union_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(union_.Annotation), unionFormCallback.probe.stageOfInterest, formDiv)
		case "MemberTypes":
			FormDivBasicFieldToField(&(union_.MemberTypes), formDiv)
		}
	}

	// manage the suppress operation
	if unionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		union_.Unstage(unionFormCallback.probe.stageOfInterest)
	}

	unionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Union](
		unionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if unionFormCallback.CreationMode || unionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		unionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(unionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UnionFormCallback(
			nil,
			unionFormCallback.probe,
			newFormGroup,
		)
		union := new(models.Union)
		FillUpForm(union, newFormGroup, unionFormCallback.probe)
		unionFormCallback.probe.formStage.Commit()
	}

	unionFormCallback.probe.ux_tree()
}
func __gong__New__WhiteSpaceFormCallback(
	whitespace *models.WhiteSpace,
	probe *Probe,
	formGroup *table.FormGroup,
) (whitespaceFormCallback *WhiteSpaceFormCallback) {
	whitespaceFormCallback = new(WhiteSpaceFormCallback)
	whitespaceFormCallback.probe = probe
	whitespaceFormCallback.whitespace = whitespace
	whitespaceFormCallback.formGroup = formGroup

	whitespaceFormCallback.CreationMode = (whitespace == nil)

	return
}

type WhiteSpaceFormCallback struct {
	whitespace *models.WhiteSpace

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (whitespaceFormCallback *WhiteSpaceFormCallback) OnSave() {
	whitespaceFormCallback.probe.stageOfInterest.Lock()
	defer whitespaceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("WhiteSpaceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	whitespaceFormCallback.probe.formStage.Checkout()

	if whitespaceFormCallback.whitespace == nil {
		whitespaceFormCallback.whitespace = new(models.WhiteSpace).Stage(whitespaceFormCallback.probe.stageOfInterest)
	}
	whitespace_ := whitespaceFormCallback.whitespace
	_ = whitespace_

	for _, formDiv := range whitespaceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(whitespace_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(whitespace_.Annotation), whitespaceFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(whitespace_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if whitespaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whitespace_.Unstage(whitespaceFormCallback.probe.stageOfInterest)
	}

	whitespaceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.WhiteSpace](
		whitespaceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if whitespaceFormCallback.CreationMode || whitespaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whitespaceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(whitespaceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WhiteSpaceFormCallback(
			nil,
			whitespaceFormCallback.probe,
			newFormGroup,
		)
		whitespace := new(models.WhiteSpace)
		FillUpForm(whitespace, newFormGroup, whitespaceFormCallback.probe)
		whitespaceFormCallback.probe.formStage.Commit()
	}

	whitespaceFormCallback.probe.ux_tree()
}
