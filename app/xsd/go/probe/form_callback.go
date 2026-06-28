// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AllFormCallback(
	all *models.All,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			allFormCallback.probe.UpdateSliceOfPointersCallback(all_, "Sequences", &all_.Sequences)

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
			allFormCallback.probe.UpdateSliceOfPointersCallback(all_, "Alls", &all_.Alls)

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
			allFormCallback.probe.UpdateSliceOfPointersCallback(all_, "Choices", &all_.Choices)

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
			allFormCallback.probe.UpdateSliceOfPointersCallback(all_, "Groups", &all_.Groups)

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
			allFormCallback.probe.UpdateSliceOfPointersCallback(all_, "Elements", &all_.Elements)

		case "Order":
			FormDivBasicFieldToField(&(all_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(all_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(all_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(all_.MaxOccurs), formDiv)
		case "All:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the All instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target All instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.All](allFormCallback.probe.stageOfInterest)
			targetAllIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAllIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all All instances and update their Alls slice
			for _all := range *models.GetGongstructInstancesSetFromPointerType[*models.All](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _all)
				
				// if All is selected
				if targetAllIDs[id] {
					// ensure all_ is in _all.Alls
					found := false
					for _, _b := range _all.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_all.Alls = append(_all.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Alls", &_all.Alls)
					}
				} else {
					// ensure all_ is NOT in _all.Alls
					idx := slices.Index(_all.Alls, all_)
					if idx != -1 {
						_all.Alls = slices.Delete(_all.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Alls", &_all.Alls)
					}
				}
			}
		case "Choice:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Choice instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Choice instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](allFormCallback.probe.stageOfInterest)
			targetChoiceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetChoiceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Choice instances and update their Alls slice
			for _choice := range *models.GetGongstructInstancesSetFromPointerType[*models.Choice](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _choice)
				
				// if Choice is selected
				if targetChoiceIDs[id] {
					// ensure all_ is in _choice.Alls
					found := false
					for _, _b := range _choice.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_choice.Alls = append(_choice.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Alls", &_choice.Alls)
					}
				} else {
					// ensure all_ is NOT in _choice.Alls
					idx := slices.Index(_choice.Alls, all_)
					if idx != -1 {
						_choice.Alls = slices.Delete(_choice.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Alls", &_choice.Alls)
					}
				}
			}
		case "ComplexType:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](allFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Alls slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure all_ is in _complextype.Alls
					found := false
					for _, _b := range _complextype.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Alls = append(_complextype.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Alls", &_complextype.Alls)
					}
				} else {
					// ensure all_ is NOT in _complextype.Alls
					idx := slices.Index(_complextype.Alls, all_)
					if idx != -1 {
						_complextype.Alls = slices.Delete(_complextype.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Alls", &_complextype.Alls)
					}
				}
			}
		case "Extension:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](allFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Alls slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure all_ is in _extension.Alls
					found := false
					for _, _b := range _extension.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Alls = append(_extension.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Alls", &_extension.Alls)
					}
				} else {
					// ensure all_ is NOT in _extension.Alls
					idx := slices.Index(_extension.Alls, all_)
					if idx != -1 {
						_extension.Alls = slices.Delete(_extension.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Alls", &_extension.Alls)
					}
				}
			}
		case "Group:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](allFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their Alls slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure all_ is in _group.Alls
					found := false
					for _, _b := range _group.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_group.Alls = append(_group.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Alls", &_group.Alls)
					}
				} else {
					// ensure all_ is NOT in _group.Alls
					idx := slices.Index(_group.Alls, all_)
					if idx != -1 {
						_group.Alls = slices.Delete(_group.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Alls", &_group.Alls)
					}
				}
			}
		case "Sequence:Alls":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Sequence instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Sequence instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](allFormCallback.probe.stageOfInterest)
			targetSequenceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSequenceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Sequence instances and update their Alls slice
			for _sequence := range *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](allFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allFormCallback.probe.stageOfInterest, _sequence)
				
				// if Sequence is selected
				if targetSequenceIDs[id] {
					// ensure all_ is in _sequence.Alls
					found := false
					for _, _b := range _sequence.Alls {
						if _b == all_ {
							found = true
							break
						}
					}
					if !found {
						_sequence.Alls = append(_sequence.Alls, all_)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Alls", &_sequence.Alls)
					}
				} else {
					// ensure all_ is NOT in _sequence.Alls
					idx := slices.Index(_sequence.Alls, all_)
					if idx != -1 {
						_sequence.Alls = slices.Delete(_sequence.Alls, idx, idx+1)
						allFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Alls", &_sequence.Alls)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			annotationFormCallback.probe.UpdateSliceOfPointersCallback(annotation_, "Documentations", &annotation_.Documentations)

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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the AttributeGroup instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target AttributeGroup instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](attributeFormCallback.probe.stageOfInterest)
			targetAttributeGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAttributeGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all AttributeGroup instances and update their Attributes slice
			for _attributegroup := range *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](attributeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributeFormCallback.probe.stageOfInterest, _attributegroup)
				
				// if AttributeGroup is selected
				if targetAttributeGroupIDs[id] {
					// ensure attribute_ is in _attributegroup.Attributes
					found := false
					for _, _b := range _attributegroup.Attributes {
						if _b == attribute_ {
							found = true
							break
						}
					}
					if !found {
						_attributegroup.Attributes = append(_attributegroup.Attributes, attribute_)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_attributegroup, "Attributes", &_attributegroup.Attributes)
					}
				} else {
					// ensure attribute_ is NOT in _attributegroup.Attributes
					idx := slices.Index(_attributegroup.Attributes, attribute_)
					if idx != -1 {
						_attributegroup.Attributes = slices.Delete(_attributegroup.Attributes, idx, idx+1)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_attributegroup, "Attributes", &_attributegroup.Attributes)
					}
				}
			}
		case "ComplexType:Attributes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](attributeFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Attributes slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](attributeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributeFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure attribute_ is in _complextype.Attributes
					found := false
					for _, _b := range _complextype.Attributes {
						if _b == attribute_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Attributes = append(_complextype.Attributes, attribute_)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Attributes", &_complextype.Attributes)
					}
				} else {
					// ensure attribute_ is NOT in _complextype.Attributes
					idx := slices.Index(_complextype.Attributes, attribute_)
					if idx != -1 {
						_complextype.Attributes = slices.Delete(_complextype.Attributes, idx, idx+1)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Attributes", &_complextype.Attributes)
					}
				}
			}
		case "Extension:Attributes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](attributeFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Attributes slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](attributeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributeFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure attribute_ is in _extension.Attributes
					found := false
					for _, _b := range _extension.Attributes {
						if _b == attribute_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Attributes = append(_extension.Attributes, attribute_)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Attributes", &_extension.Attributes)
					}
				} else {
					// ensure attribute_ is NOT in _extension.Attributes
					idx := slices.Index(_extension.Attributes, attribute_)
					if idx != -1 {
						_extension.Attributes = slices.Delete(_extension.Attributes, idx, idx+1)
						attributeFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Attributes", &_extension.Attributes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(attributegroup_, "AttributeGroups", &attributegroup_.AttributeGroups)

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
			attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(attributegroup_, "Attributes", &attributegroup_.Attributes)

		case "Order":
			FormDivBasicFieldToField(&(attributegroup_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(attributegroup_.Depth), formDiv)
		case "AttributeGroup:AttributeGroups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the AttributeGroup instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target AttributeGroup instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest)
			targetAttributeGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAttributeGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all AttributeGroup instances and update their AttributeGroups slice
			for _attributegroup := range *models.GetGongstructInstancesSetFromPointerType[*models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributegroupFormCallback.probe.stageOfInterest, _attributegroup)
				
				// if AttributeGroup is selected
				if targetAttributeGroupIDs[id] {
					// ensure attributegroup_ is in _attributegroup.AttributeGroups
					found := false
					for _, _b := range _attributegroup.AttributeGroups {
						if _b == attributegroup_ {
							found = true
							break
						}
					}
					if !found {
						_attributegroup.AttributeGroups = append(_attributegroup.AttributeGroups, attributegroup_)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_attributegroup, "AttributeGroups", &_attributegroup.AttributeGroups)
					}
				} else {
					// ensure attributegroup_ is NOT in _attributegroup.AttributeGroups
					idx := slices.Index(_attributegroup.AttributeGroups, attributegroup_)
					if idx != -1 {
						_attributegroup.AttributeGroups = slices.Delete(_attributegroup.AttributeGroups, idx, idx+1)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_attributegroup, "AttributeGroups", &_attributegroup.AttributeGroups)
					}
				}
			}
		case "ComplexType:AttributeGroups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](attributegroupFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their AttributeGroups slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](attributegroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributegroupFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure attributegroup_ is in _complextype.AttributeGroups
					found := false
					for _, _b := range _complextype.AttributeGroups {
						if _b == attributegroup_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.AttributeGroups = append(_complextype.AttributeGroups, attributegroup_)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "AttributeGroups", &_complextype.AttributeGroups)
					}
				} else {
					// ensure attributegroup_ is NOT in _complextype.AttributeGroups
					idx := slices.Index(_complextype.AttributeGroups, attributegroup_)
					if idx != -1 {
						_complextype.AttributeGroups = slices.Delete(_complextype.AttributeGroups, idx, idx+1)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "AttributeGroups", &_complextype.AttributeGroups)
					}
				}
			}
		case "Extension:AttributeGroups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](attributegroupFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their AttributeGroups slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](attributegroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributegroupFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure attributegroup_ is in _extension.AttributeGroups
					found := false
					for _, _b := range _extension.AttributeGroups {
						if _b == attributegroup_ {
							found = true
							break
						}
					}
					if !found {
						_extension.AttributeGroups = append(_extension.AttributeGroups, attributegroup_)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "AttributeGroups", &_extension.AttributeGroups)
					}
				} else {
					// ensure attributegroup_ is NOT in _extension.AttributeGroups
					idx := slices.Index(_extension.AttributeGroups, attributegroup_)
					if idx != -1 {
						_extension.AttributeGroups = slices.Delete(_extension.AttributeGroups, idx, idx+1)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "AttributeGroups", &_extension.AttributeGroups)
					}
				}
			}
		case "Schema:AttributeGroups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Schema instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Schema instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Schema](attributegroupFormCallback.probe.stageOfInterest)
			targetSchemaIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSchemaIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Schema instances and update their AttributeGroups slice
			for _schema := range *models.GetGongstructInstancesSetFromPointerType[*models.Schema](attributegroupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributegroupFormCallback.probe.stageOfInterest, _schema)
				
				// if Schema is selected
				if targetSchemaIDs[id] {
					// ensure attributegroup_ is in _schema.AttributeGroups
					found := false
					for _, _b := range _schema.AttributeGroups {
						if _b == attributegroup_ {
							found = true
							break
						}
					}
					if !found {
						_schema.AttributeGroups = append(_schema.AttributeGroups, attributegroup_)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "AttributeGroups", &_schema.AttributeGroups)
					}
				} else {
					// ensure attributegroup_ is NOT in _schema.AttributeGroups
					idx := slices.Index(_schema.AttributeGroups, attributegroup_)
					if idx != -1 {
						_schema.AttributeGroups = slices.Delete(_schema.AttributeGroups, idx, idx+1)
						attributegroupFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "AttributeGroups", &_schema.AttributeGroups)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			choiceFormCallback.probe.UpdateSliceOfPointersCallback(choice_, "Sequences", &choice_.Sequences)

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
			choiceFormCallback.probe.UpdateSliceOfPointersCallback(choice_, "Alls", &choice_.Alls)

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
			choiceFormCallback.probe.UpdateSliceOfPointersCallback(choice_, "Choices", &choice_.Choices)

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
			choiceFormCallback.probe.UpdateSliceOfPointersCallback(choice_, "Groups", &choice_.Groups)

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
			choiceFormCallback.probe.UpdateSliceOfPointersCallback(choice_, "Elements", &choice_.Elements)

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the All instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target All instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.All](choiceFormCallback.probe.stageOfInterest)
			targetAllIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAllIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all All instances and update their Choices slice
			for _all := range *models.GetGongstructInstancesSetFromPointerType[*models.All](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _all)
				
				// if All is selected
				if targetAllIDs[id] {
					// ensure choice_ is in _all.Choices
					found := false
					for _, _b := range _all.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_all.Choices = append(_all.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Choices", &_all.Choices)
					}
				} else {
					// ensure choice_ is NOT in _all.Choices
					idx := slices.Index(_all.Choices, choice_)
					if idx != -1 {
						_all.Choices = slices.Delete(_all.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Choices", &_all.Choices)
					}
				}
			}
		case "Choice:Choices":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Choice instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Choice instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](choiceFormCallback.probe.stageOfInterest)
			targetChoiceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetChoiceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Choice instances and update their Choices slice
			for _choice := range *models.GetGongstructInstancesSetFromPointerType[*models.Choice](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _choice)
				
				// if Choice is selected
				if targetChoiceIDs[id] {
					// ensure choice_ is in _choice.Choices
					found := false
					for _, _b := range _choice.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_choice.Choices = append(_choice.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Choices", &_choice.Choices)
					}
				} else {
					// ensure choice_ is NOT in _choice.Choices
					idx := slices.Index(_choice.Choices, choice_)
					if idx != -1 {
						_choice.Choices = slices.Delete(_choice.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Choices", &_choice.Choices)
					}
				}
			}
		case "ComplexType:Choices":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](choiceFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Choices slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure choice_ is in _complextype.Choices
					found := false
					for _, _b := range _complextype.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Choices = append(_complextype.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Choices", &_complextype.Choices)
					}
				} else {
					// ensure choice_ is NOT in _complextype.Choices
					idx := slices.Index(_complextype.Choices, choice_)
					if idx != -1 {
						_complextype.Choices = slices.Delete(_complextype.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Choices", &_complextype.Choices)
					}
				}
			}
		case "Extension:Choices":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](choiceFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Choices slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure choice_ is in _extension.Choices
					found := false
					for _, _b := range _extension.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Choices = append(_extension.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Choices", &_extension.Choices)
					}
				} else {
					// ensure choice_ is NOT in _extension.Choices
					idx := slices.Index(_extension.Choices, choice_)
					if idx != -1 {
						_extension.Choices = slices.Delete(_extension.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Choices", &_extension.Choices)
					}
				}
			}
		case "Group:Choices":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](choiceFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their Choices slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure choice_ is in _group.Choices
					found := false
					for _, _b := range _group.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_group.Choices = append(_group.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Choices", &_group.Choices)
					}
				} else {
					// ensure choice_ is NOT in _group.Choices
					idx := slices.Index(_group.Choices, choice_)
					if idx != -1 {
						_group.Choices = slices.Delete(_group.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Choices", &_group.Choices)
					}
				}
			}
		case "Sequence:Choices":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Sequence instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Sequence instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](choiceFormCallback.probe.stageOfInterest)
			targetSequenceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSequenceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Sequence instances and update their Choices slice
			for _sequence := range *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](choiceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(choiceFormCallback.probe.stageOfInterest, _sequence)
				
				// if Sequence is selected
				if targetSequenceIDs[id] {
					// ensure choice_ is in _sequence.Choices
					found := false
					for _, _b := range _sequence.Choices {
						if _b == choice_ {
							found = true
							break
						}
					}
					if !found {
						_sequence.Choices = append(_sequence.Choices, choice_)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Choices", &_sequence.Choices)
					}
				} else {
					// ensure choice_ is NOT in _sequence.Choices
					idx := slices.Index(_sequence.Choices, choice_)
					if idx != -1 {
						_sequence.Choices = slices.Delete(_sequence.Choices, idx, idx+1)
						choiceFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Choices", &_sequence.Choices)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Sequences", &complextype_.Sequences)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Alls", &complextype_.Alls)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Choices", &complextype_.Choices)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Groups", &complextype_.Groups)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Elements", &complextype_.Elements)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "Attributes", &complextype_.Attributes)

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
			complextypeFormCallback.probe.UpdateSliceOfPointersCallback(complextype_, "AttributeGroups", &complextype_.AttributeGroups)

		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(complextype_.IsDuplicatedInXSD), formDiv)
		case "Schema:ComplexTypes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Schema instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Schema instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Schema](complextypeFormCallback.probe.stageOfInterest)
			targetSchemaIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSchemaIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Schema instances and update their ComplexTypes slice
			for _schema := range *models.GetGongstructInstancesSetFromPointerType[*models.Schema](complextypeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complextypeFormCallback.probe.stageOfInterest, _schema)
				
				// if Schema is selected
				if targetSchemaIDs[id] {
					// ensure complextype_ is in _schema.ComplexTypes
					found := false
					for _, _b := range _schema.ComplexTypes {
						if _b == complextype_ {
							found = true
							break
						}
					}
					if !found {
						_schema.ComplexTypes = append(_schema.ComplexTypes, complextype_)
						complextypeFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "ComplexTypes", &_schema.ComplexTypes)
					}
				} else {
					// ensure complextype_ is NOT in _schema.ComplexTypes
					idx := slices.Index(_schema.ComplexTypes, complextype_)
					if idx != -1 {
						_schema.ComplexTypes = slices.Delete(_schema.ComplexTypes, idx, idx+1)
						complextypeFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "ComplexTypes", &_schema.ComplexTypes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Annotation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Annotation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Annotation](documentationFormCallback.probe.stageOfInterest)
			targetAnnotationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAnnotationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Annotation instances and update their Documentations slice
			for _annotation := range *models.GetGongstructInstancesSetFromPointerType[*models.Annotation](documentationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(documentationFormCallback.probe.stageOfInterest, _annotation)
				
				// if Annotation is selected
				if targetAnnotationIDs[id] {
					// ensure documentation_ is in _annotation.Documentations
					found := false
					for _, _b := range _annotation.Documentations {
						if _b == documentation_ {
							found = true
							break
						}
					}
					if !found {
						_annotation.Documentations = append(_annotation.Documentations, documentation_)
						documentationFormCallback.probe.UpdateSliceOfPointersCallback(_annotation, "Documentations", &_annotation.Documentations)
					}
				} else {
					// ensure documentation_ is NOT in _annotation.Documentations
					idx := slices.Index(_annotation.Documentations, documentation_)
					if idx != -1 {
						_annotation.Documentations = slices.Delete(_annotation.Documentations, idx, idx+1)
						documentationFormCallback.probe.UpdateSliceOfPointersCallback(_annotation, "Documentations", &_annotation.Documentations)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			elementFormCallback.probe.UpdateSliceOfPointersCallback(element_, "Groups", &element_.Groups)

		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(element_.IsDuplicatedInXSD), formDiv)
		case "All:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the All instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target All instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.All](elementFormCallback.probe.stageOfInterest)
			targetAllIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAllIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all All instances and update their Elements slice
			for _all := range *models.GetGongstructInstancesSetFromPointerType[*models.All](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _all)
				
				// if All is selected
				if targetAllIDs[id] {
					// ensure element_ is in _all.Elements
					found := false
					for _, _b := range _all.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_all.Elements = append(_all.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Elements", &_all.Elements)
					}
				} else {
					// ensure element_ is NOT in _all.Elements
					idx := slices.Index(_all.Elements, element_)
					if idx != -1 {
						_all.Elements = slices.Delete(_all.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Elements", &_all.Elements)
					}
				}
			}
		case "Choice:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Choice instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Choice instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](elementFormCallback.probe.stageOfInterest)
			targetChoiceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetChoiceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Choice instances and update their Elements slice
			for _choice := range *models.GetGongstructInstancesSetFromPointerType[*models.Choice](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _choice)
				
				// if Choice is selected
				if targetChoiceIDs[id] {
					// ensure element_ is in _choice.Elements
					found := false
					for _, _b := range _choice.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_choice.Elements = append(_choice.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Elements", &_choice.Elements)
					}
				} else {
					// ensure element_ is NOT in _choice.Elements
					idx := slices.Index(_choice.Elements, element_)
					if idx != -1 {
						_choice.Elements = slices.Delete(_choice.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Elements", &_choice.Elements)
					}
				}
			}
		case "ComplexType:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](elementFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Elements slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure element_ is in _complextype.Elements
					found := false
					for _, _b := range _complextype.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Elements = append(_complextype.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Elements", &_complextype.Elements)
					}
				} else {
					// ensure element_ is NOT in _complextype.Elements
					idx := slices.Index(_complextype.Elements, element_)
					if idx != -1 {
						_complextype.Elements = slices.Delete(_complextype.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Elements", &_complextype.Elements)
					}
				}
			}
		case "Extension:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](elementFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Elements slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure element_ is in _extension.Elements
					found := false
					for _, _b := range _extension.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Elements = append(_extension.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Elements", &_extension.Elements)
					}
				} else {
					// ensure element_ is NOT in _extension.Elements
					idx := slices.Index(_extension.Elements, element_)
					if idx != -1 {
						_extension.Elements = slices.Delete(_extension.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Elements", &_extension.Elements)
					}
				}
			}
		case "Group:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](elementFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their Elements slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure element_ is in _group.Elements
					found := false
					for _, _b := range _group.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_group.Elements = append(_group.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Elements", &_group.Elements)
					}
				} else {
					// ensure element_ is NOT in _group.Elements
					idx := slices.Index(_group.Elements, element_)
					if idx != -1 {
						_group.Elements = slices.Delete(_group.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Elements", &_group.Elements)
					}
				}
			}
		case "Schema:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Schema instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Schema instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Schema](elementFormCallback.probe.stageOfInterest)
			targetSchemaIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSchemaIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Schema instances and update their Elements slice
			for _schema := range *models.GetGongstructInstancesSetFromPointerType[*models.Schema](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _schema)
				
				// if Schema is selected
				if targetSchemaIDs[id] {
					// ensure element_ is in _schema.Elements
					found := false
					for _, _b := range _schema.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_schema.Elements = append(_schema.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "Elements", &_schema.Elements)
					}
				} else {
					// ensure element_ is NOT in _schema.Elements
					idx := slices.Index(_schema.Elements, element_)
					if idx != -1 {
						_schema.Elements = slices.Delete(_schema.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "Elements", &_schema.Elements)
					}
				}
			}
		case "Sequence:Elements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Sequence instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Sequence instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](elementFormCallback.probe.stageOfInterest)
			targetSequenceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSequenceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Sequence instances and update their Elements slice
			for _sequence := range *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](elementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(elementFormCallback.probe.stageOfInterest, _sequence)
				
				// if Sequence is selected
				if targetSequenceIDs[id] {
					// ensure element_ is in _sequence.Elements
					found := false
					for _, _b := range _sequence.Elements {
						if _b == element_ {
							found = true
							break
						}
					}
					if !found {
						_sequence.Elements = append(_sequence.Elements, element_)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Elements", &_sequence.Elements)
					}
				} else {
					// ensure element_ is NOT in _sequence.Elements
					idx := slices.Index(_sequence.Elements, element_)
					if idx != -1 {
						_sequence.Elements = slices.Delete(_sequence.Elements, idx, idx+1)
						elementFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Elements", &_sequence.Elements)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Restriction instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Restriction instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Restriction](enumerationFormCallback.probe.stageOfInterest)
			targetRestrictionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRestrictionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Restriction instances and update their Enumerations slice
			for _restriction := range *models.GetGongstructInstancesSetFromPointerType[*models.Restriction](enumerationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(enumerationFormCallback.probe.stageOfInterest, _restriction)
				
				// if Restriction is selected
				if targetRestrictionIDs[id] {
					// ensure enumeration_ is in _restriction.Enumerations
					found := false
					for _, _b := range _restriction.Enumerations {
						if _b == enumeration_ {
							found = true
							break
						}
					}
					if !found {
						_restriction.Enumerations = append(_restriction.Enumerations, enumeration_)
						enumerationFormCallback.probe.UpdateSliceOfPointersCallback(_restriction, "Enumerations", &_restriction.Enumerations)
					}
				} else {
					// ensure enumeration_ is NOT in _restriction.Enumerations
					idx := slices.Index(_restriction.Enumerations, enumeration_)
					if idx != -1 {
						_restriction.Enumerations = slices.Delete(_restriction.Enumerations, idx, idx+1)
						enumerationFormCallback.probe.UpdateSliceOfPointersCallback(_restriction, "Enumerations", &_restriction.Enumerations)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Sequences", &extension_.Sequences)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Alls", &extension_.Alls)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Choices", &extension_.Choices)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Groups", &extension_.Groups)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Elements", &extension_.Elements)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "Attributes", &extension_.Attributes)

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
			extensionFormCallback.probe.UpdateSliceOfPointersCallback(extension_, "AttributeGroups", &extension_.AttributeGroups)

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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "Sequences", &group_.Sequences)

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
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "Alls", &group_.Alls)

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
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "Choices", &group_.Choices)

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
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "Groups", &group_.Groups)

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
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "Elements", &group_.Elements)

		case "Order":
			FormDivBasicFieldToField(&(group_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(group_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(group_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(group_.MaxOccurs), formDiv)
		case "All:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the All instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target All instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.All](groupFormCallback.probe.stageOfInterest)
			targetAllIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAllIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all All instances and update their Groups slice
			for _all := range *models.GetGongstructInstancesSetFromPointerType[*models.All](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _all)
				
				// if All is selected
				if targetAllIDs[id] {
					// ensure group_ is in _all.Groups
					found := false
					for _, _b := range _all.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_all.Groups = append(_all.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Groups", &_all.Groups)
					}
				} else {
					// ensure group_ is NOT in _all.Groups
					idx := slices.Index(_all.Groups, group_)
					if idx != -1 {
						_all.Groups = slices.Delete(_all.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Groups", &_all.Groups)
					}
				}
			}
		case "Choice:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Choice instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Choice instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](groupFormCallback.probe.stageOfInterest)
			targetChoiceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetChoiceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Choice instances and update their Groups slice
			for _choice := range *models.GetGongstructInstancesSetFromPointerType[*models.Choice](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _choice)
				
				// if Choice is selected
				if targetChoiceIDs[id] {
					// ensure group_ is in _choice.Groups
					found := false
					for _, _b := range _choice.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_choice.Groups = append(_choice.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Groups", &_choice.Groups)
					}
				} else {
					// ensure group_ is NOT in _choice.Groups
					idx := slices.Index(_choice.Groups, group_)
					if idx != -1 {
						_choice.Groups = slices.Delete(_choice.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Groups", &_choice.Groups)
					}
				}
			}
		case "ComplexType:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](groupFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Groups slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure group_ is in _complextype.Groups
					found := false
					for _, _b := range _complextype.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Groups = append(_complextype.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Groups", &_complextype.Groups)
					}
				} else {
					// ensure group_ is NOT in _complextype.Groups
					idx := slices.Index(_complextype.Groups, group_)
					if idx != -1 {
						_complextype.Groups = slices.Delete(_complextype.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Groups", &_complextype.Groups)
					}
				}
			}
		case "Element:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Element instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Element instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Element](groupFormCallback.probe.stageOfInterest)
			targetElementIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetElementIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Element instances and update their Groups slice
			for _element := range *models.GetGongstructInstancesSetFromPointerType[*models.Element](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _element)
				
				// if Element is selected
				if targetElementIDs[id] {
					// ensure group_ is in _element.Groups
					found := false
					for _, _b := range _element.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_element.Groups = append(_element.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_element, "Groups", &_element.Groups)
					}
				} else {
					// ensure group_ is NOT in _element.Groups
					idx := slices.Index(_element.Groups, group_)
					if idx != -1 {
						_element.Groups = slices.Delete(_element.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_element, "Groups", &_element.Groups)
					}
				}
			}
		case "Extension:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](groupFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Groups slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure group_ is in _extension.Groups
					found := false
					for _, _b := range _extension.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Groups = append(_extension.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Groups", &_extension.Groups)
					}
				} else {
					// ensure group_ is NOT in _extension.Groups
					idx := slices.Index(_extension.Groups, group_)
					if idx != -1 {
						_extension.Groups = slices.Delete(_extension.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Groups", &_extension.Groups)
					}
				}
			}
		case "Group:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](groupFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their Groups slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure group_ is in _group.Groups
					found := false
					for _, _b := range _group.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_group.Groups = append(_group.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Groups", &_group.Groups)
					}
				} else {
					// ensure group_ is NOT in _group.Groups
					idx := slices.Index(_group.Groups, group_)
					if idx != -1 {
						_group.Groups = slices.Delete(_group.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Groups", &_group.Groups)
					}
				}
			}
		case "Schema:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Schema instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Schema instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Schema](groupFormCallback.probe.stageOfInterest)
			targetSchemaIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSchemaIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Schema instances and update their Groups slice
			for _schema := range *models.GetGongstructInstancesSetFromPointerType[*models.Schema](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _schema)
				
				// if Schema is selected
				if targetSchemaIDs[id] {
					// ensure group_ is in _schema.Groups
					found := false
					for _, _b := range _schema.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_schema.Groups = append(_schema.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "Groups", &_schema.Groups)
					}
				} else {
					// ensure group_ is NOT in _schema.Groups
					idx := slices.Index(_schema.Groups, group_)
					if idx != -1 {
						_schema.Groups = slices.Delete(_schema.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "Groups", &_schema.Groups)
					}
				}
			}
		case "Sequence:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Sequence instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Sequence instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](groupFormCallback.probe.stageOfInterest)
			targetSequenceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSequenceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Sequence instances and update their Groups slice
			for _sequence := range *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _sequence)
				
				// if Sequence is selected
				if targetSequenceIDs[id] {
					// ensure group_ is in _sequence.Groups
					found := false
					for _, _b := range _sequence.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_sequence.Groups = append(_sequence.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Groups", &_sequence.Groups)
					}
				} else {
					// ensure group_ is NOT in _sequence.Groups
					idx := slices.Index(_sequence.Groups, group_)
					if idx != -1 {
						_sequence.Groups = slices.Delete(_sequence.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Groups", &_sequence.Groups)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			restrictionFormCallback.probe.UpdateSliceOfPointersCallback(restriction_, "Enumerations", &restriction_.Enumerations)

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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			schemaFormCallback.probe.UpdateSliceOfPointersCallback(schema_, "Elements", &schema_.Elements)

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
			schemaFormCallback.probe.UpdateSliceOfPointersCallback(schema_, "SimpleTypes", &schema_.SimpleTypes)

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
			schemaFormCallback.probe.UpdateSliceOfPointersCallback(schema_, "ComplexTypes", &schema_.ComplexTypes)

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
			schemaFormCallback.probe.UpdateSliceOfPointersCallback(schema_, "AttributeGroups", &schema_.AttributeGroups)

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
			schemaFormCallback.probe.UpdateSliceOfPointersCallback(schema_, "Groups", &schema_.Groups)

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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			sequenceFormCallback.probe.UpdateSliceOfPointersCallback(sequence_, "Sequences", &sequence_.Sequences)

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
			sequenceFormCallback.probe.UpdateSliceOfPointersCallback(sequence_, "Alls", &sequence_.Alls)

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
			sequenceFormCallback.probe.UpdateSliceOfPointersCallback(sequence_, "Choices", &sequence_.Choices)

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
			sequenceFormCallback.probe.UpdateSliceOfPointersCallback(sequence_, "Groups", &sequence_.Groups)

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
			sequenceFormCallback.probe.UpdateSliceOfPointersCallback(sequence_, "Elements", &sequence_.Elements)

		case "Order":
			FormDivBasicFieldToField(&(sequence_.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(sequence_.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(sequence_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(sequence_.MaxOccurs), formDiv)
		case "All:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the All instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target All instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.All](sequenceFormCallback.probe.stageOfInterest)
			targetAllIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAllIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all All instances and update their Sequences slice
			for _all := range *models.GetGongstructInstancesSetFromPointerType[*models.All](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _all)
				
				// if All is selected
				if targetAllIDs[id] {
					// ensure sequence_ is in _all.Sequences
					found := false
					for _, _b := range _all.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_all.Sequences = append(_all.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Sequences", &_all.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _all.Sequences
					idx := slices.Index(_all.Sequences, sequence_)
					if idx != -1 {
						_all.Sequences = slices.Delete(_all.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_all, "Sequences", &_all.Sequences)
					}
				}
			}
		case "Choice:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Choice instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Choice instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Choice](sequenceFormCallback.probe.stageOfInterest)
			targetChoiceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetChoiceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Choice instances and update their Sequences slice
			for _choice := range *models.GetGongstructInstancesSetFromPointerType[*models.Choice](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _choice)
				
				// if Choice is selected
				if targetChoiceIDs[id] {
					// ensure sequence_ is in _choice.Sequences
					found := false
					for _, _b := range _choice.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_choice.Sequences = append(_choice.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Sequences", &_choice.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _choice.Sequences
					idx := slices.Index(_choice.Sequences, sequence_)
					if idx != -1 {
						_choice.Sequences = slices.Delete(_choice.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_choice, "Sequences", &_choice.Sequences)
					}
				}
			}
		case "ComplexType:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ComplexType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ComplexType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ComplexType](sequenceFormCallback.probe.stageOfInterest)
			targetComplexTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetComplexTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ComplexType instances and update their Sequences slice
			for _complextype := range *models.GetGongstructInstancesSetFromPointerType[*models.ComplexType](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _complextype)
				
				// if ComplexType is selected
				if targetComplexTypeIDs[id] {
					// ensure sequence_ is in _complextype.Sequences
					found := false
					for _, _b := range _complextype.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_complextype.Sequences = append(_complextype.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Sequences", &_complextype.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _complextype.Sequences
					idx := slices.Index(_complextype.Sequences, sequence_)
					if idx != -1 {
						_complextype.Sequences = slices.Delete(_complextype.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_complextype, "Sequences", &_complextype.Sequences)
					}
				}
			}
		case "Extension:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Extension instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Extension instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Extension](sequenceFormCallback.probe.stageOfInterest)
			targetExtensionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetExtensionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Extension instances and update their Sequences slice
			for _extension := range *models.GetGongstructInstancesSetFromPointerType[*models.Extension](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _extension)
				
				// if Extension is selected
				if targetExtensionIDs[id] {
					// ensure sequence_ is in _extension.Sequences
					found := false
					for _, _b := range _extension.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_extension.Sequences = append(_extension.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Sequences", &_extension.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _extension.Sequences
					idx := slices.Index(_extension.Sequences, sequence_)
					if idx != -1 {
						_extension.Sequences = slices.Delete(_extension.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_extension, "Sequences", &_extension.Sequences)
					}
				}
			}
		case "Group:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](sequenceFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their Sequences slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure sequence_ is in _group.Sequences
					found := false
					for _, _b := range _group.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_group.Sequences = append(_group.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Sequences", &_group.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _group.Sequences
					idx := slices.Index(_group.Sequences, sequence_)
					if idx != -1 {
						_group.Sequences = slices.Delete(_group.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_group, "Sequences", &_group.Sequences)
					}
				}
			}
		case "Sequence:Sequences":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Sequence instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Sequence instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Sequence](sequenceFormCallback.probe.stageOfInterest)
			targetSequenceIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSequenceIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Sequence instances and update their Sequences slice
			for _sequence := range *models.GetGongstructInstancesSetFromPointerType[*models.Sequence](sequenceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(sequenceFormCallback.probe.stageOfInterest, _sequence)
				
				// if Sequence is selected
				if targetSequenceIDs[id] {
					// ensure sequence_ is in _sequence.Sequences
					found := false
					for _, _b := range _sequence.Sequences {
						if _b == sequence_ {
							found = true
							break
						}
					}
					if !found {
						_sequence.Sequences = append(_sequence.Sequences, sequence_)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Sequences", &_sequence.Sequences)
					}
				} else {
					// ensure sequence_ is NOT in _sequence.Sequences
					idx := slices.Index(_sequence.Sequences, sequence_)
					if idx != -1 {
						_sequence.Sequences = slices.Delete(_sequence.Sequences, idx, idx+1)
						sequenceFormCallback.probe.UpdateSliceOfPointersCallback(_sequence, "Sequences", &_sequence.Sequences)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Schema instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Schema instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Schema](simpletypeFormCallback.probe.stageOfInterest)
			targetSchemaIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSchemaIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Schema instances and update their SimpleTypes slice
			for _schema := range *models.GetGongstructInstancesSetFromPointerType[*models.Schema](simpletypeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(simpletypeFormCallback.probe.stageOfInterest, _schema)
				
				// if Schema is selected
				if targetSchemaIDs[id] {
					// ensure simpletype_ is in _schema.SimpleTypes
					found := false
					for _, _b := range _schema.SimpleTypes {
						if _b == simpletype_ {
							found = true
							break
						}
					}
					if !found {
						_schema.SimpleTypes = append(_schema.SimpleTypes, simpletype_)
						simpletypeFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "SimpleTypes", &_schema.SimpleTypes)
					}
				} else {
					// ensure simpletype_ is NOT in _schema.SimpleTypes
					idx := slices.Index(_schema.SimpleTypes, simpletype_)
					if idx != -1 {
						_schema.SimpleTypes = slices.Delete(_schema.SimpleTypes, idx, idx+1)
						simpletypeFormCallback.probe.UpdateSliceOfPointersCallback(_schema, "SimpleTypes", &_schema.SimpleTypes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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
