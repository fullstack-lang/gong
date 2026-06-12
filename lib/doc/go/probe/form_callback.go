// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AttributeShapeFormCallback(
	attributeshape *models.AttributeShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (attributeshapeFormCallback *AttributeShapeFormCallback) {
	attributeshapeFormCallback = new(AttributeShapeFormCallback)
	attributeshapeFormCallback.probe = probe
	attributeshapeFormCallback.attributeshape = attributeshape
	attributeshapeFormCallback.formGroup = formGroup

	attributeshapeFormCallback.CreationMode = (attributeshape == nil)

	return
}

type AttributeShapeFormCallback struct {
	attributeshape *models.AttributeShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (attributeshapeFormCallback *AttributeShapeFormCallback) OnSave() {
	attributeshapeFormCallback.probe.stageOfInterest.Lock()
	defer attributeshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AttributeShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributeshapeFormCallback.probe.formStage.Checkout()

	if attributeshapeFormCallback.attributeshape == nil {
		attributeshapeFormCallback.attributeshape = new(models.AttributeShape).Stage(attributeshapeFormCallback.probe.stageOfInterest)
	}
	attributeshape_ := attributeshapeFormCallback.attributeshape
	_ = attributeshape_

	for _, formDiv := range attributeshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributeshape_.Name), formDiv)
		case "FieldTypeAsString":
			FormDivBasicFieldToField(&(attributeshape_.FieldTypeAsString), formDiv)
		case "Structname":
			FormDivBasicFieldToField(&(attributeshape_.Structname), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(attributeshape_.Fieldtypename), formDiv)
		case "GongStructShape:AttributeShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStructShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStructShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStructShape](attributeshapeFormCallback.probe.stageOfInterest)
			targetGongStructShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStructShape instances and update their AttributeShapes slice
			for _gongstructshape := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStructShape](attributeshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(attributeshapeFormCallback.probe.stageOfInterest, _gongstructshape)
				
				// if GongStructShape is selected
				if targetGongStructShapeIDs[id] {
					// ensure attributeshape_ is in _gongstructshape.AttributeShapes
					found := false
					for _, _b := range _gongstructshape.AttributeShapes {
						if _b == attributeshape_ {
							found = true
							break
						}
					}
					if !found {
						_gongstructshape.AttributeShapes = append(_gongstructshape.AttributeShapes, attributeshape_)
						attributeshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongstructshape, "AttributeShapes", &_gongstructshape.AttributeShapes)
					}
				} else {
					// ensure attributeshape_ is NOT in _gongstructshape.AttributeShapes
					idx := slices.Index(_gongstructshape.AttributeShapes, attributeshape_)
					if idx != -1 {
						_gongstructshape.AttributeShapes = slices.Delete(_gongstructshape.AttributeShapes, idx, idx+1)
						attributeshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongstructshape, "AttributeShapes", &_gongstructshape.AttributeShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeshape_.Unstage(attributeshapeFormCallback.probe.stageOfInterest)
	}

	attributeshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AttributeShape](
		attributeshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attributeshapeFormCallback.CreationMode || attributeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(attributeshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeShapeFormCallback(
			nil,
			attributeshapeFormCallback.probe,
			newFormGroup,
		)
		attributeshape := new(models.AttributeShape)
		FillUpForm(attributeshape, newFormGroup, attributeshapeFormCallback.probe)
		attributeshapeFormCallback.probe.formStage.Commit()
	}

	attributeshapeFormCallback.probe.ux_tree()
}
func __gong__New__ClassdiagramFormCallback(
	classdiagram *models.Classdiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (classdiagramFormCallback *ClassdiagramFormCallback) {
	classdiagramFormCallback = new(ClassdiagramFormCallback)
	classdiagramFormCallback.probe = probe
	classdiagramFormCallback.classdiagram = classdiagram
	classdiagramFormCallback.formGroup = formGroup

	classdiagramFormCallback.CreationMode = (classdiagram == nil)

	return
}

type ClassdiagramFormCallback struct {
	classdiagram *models.Classdiagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (classdiagramFormCallback *ClassdiagramFormCallback) OnSave() {
	classdiagramFormCallback.probe.stageOfInterest.Lock()
	defer classdiagramFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ClassdiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.probe.formStage.Checkout()

	if classdiagramFormCallback.classdiagram == nil {
		classdiagramFormCallback.classdiagram = new(models.Classdiagram).Stage(classdiagramFormCallback.probe.stageOfInterest)
	}
	classdiagram_ := classdiagramFormCallback.classdiagram
	_ = classdiagram_

	for _, formDiv := range classdiagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(classdiagram_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(classdiagram_.Description), formDiv)
		case "IsIncludedInStaticWebSite":
			FormDivBasicFieldToField(&(classdiagram_.IsIncludedInStaticWebSite), formDiv)
		case "GongStructShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongStructShape](classdiagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongStructShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongStructShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					classdiagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongStructShape](classdiagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			classdiagram_.GongStructShapes = instanceSlice
			classdiagramFormCallback.probe.UpdateSliceOfPointersCallback(classdiagram_, "GongStructShapes", &classdiagram_.GongStructShapes)

		case "GongEnumShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumShape](classdiagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongEnumShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongEnumShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					classdiagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongEnumShape](classdiagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			classdiagram_.GongEnumShapes = instanceSlice
			classdiagramFormCallback.probe.UpdateSliceOfPointersCallback(classdiagram_, "GongEnumShapes", &classdiagram_.GongEnumShapes)

		case "GongNoteShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteShape](classdiagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongNoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongNoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					classdiagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongNoteShape](classdiagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			classdiagram_.GongNoteShapes = instanceSlice
			classdiagramFormCallback.probe.UpdateSliceOfPointersCallback(classdiagram_, "GongNoteShapes", &classdiagram_.GongNoteShapes)

		case "ShowNbInstances":
			FormDivBasicFieldToField(&(classdiagram_.ShowNbInstances), formDiv)
		case "ShowMultiplicity":
			FormDivBasicFieldToField(&(classdiagram_.ShowMultiplicity), formDiv)
		case "ShowLinkNames":
			FormDivBasicFieldToField(&(classdiagram_.ShowLinkNames), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.IsExpanded), formDiv)
		case "NodeGongStructsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructsIsExpanded), formDiv)
		case "NodeGongStructNodeExpansion":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructNodeExpansion), formDiv)
		case "NodeGongEnumsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongEnumsIsExpanded), formDiv)
		case "NodeGongEnumNodeExpansion":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongEnumNodeExpansion), formDiv)
		case "NodeGongNotesIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongNotesIsExpanded), formDiv)
		case "NodeGongNoteNodeExpansion":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongNoteNodeExpansion), formDiv)
		case "DiagramPackage:Classdiagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramPackage instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramPackage instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest)
			targetDiagramPackageIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramPackageIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramPackage instances and update their Classdiagrams slice
			for _diagrampackage := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(classdiagramFormCallback.probe.stageOfInterest, _diagrampackage)
				
				// if DiagramPackage is selected
				if targetDiagramPackageIDs[id] {
					// ensure classdiagram_ is in _diagrampackage.Classdiagrams
					found := false
					for _, _b := range _diagrampackage.Classdiagrams {
						if _b == classdiagram_ {
							found = true
							break
						}
					}
					if !found {
						_diagrampackage.Classdiagrams = append(_diagrampackage.Classdiagrams, classdiagram_)
						classdiagramFormCallback.probe.UpdateSliceOfPointersCallback(_diagrampackage, "Classdiagrams", &_diagrampackage.Classdiagrams)
					}
				} else {
					// ensure classdiagram_ is NOT in _diagrampackage.Classdiagrams
					idx := slices.Index(_diagrampackage.Classdiagrams, classdiagram_)
					if idx != -1 {
						_diagrampackage.Classdiagrams = slices.Delete(_diagrampackage.Classdiagrams, idx, idx+1)
						classdiagramFormCallback.probe.UpdateSliceOfPointersCallback(_diagrampackage, "Classdiagrams", &_diagrampackage.Classdiagrams)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if classdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		classdiagram_.Unstage(classdiagramFormCallback.probe.stageOfInterest)
	}

	classdiagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Classdiagram](
		classdiagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if classdiagramFormCallback.CreationMode || classdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		classdiagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(classdiagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ClassdiagramFormCallback(
			nil,
			classdiagramFormCallback.probe,
			newFormGroup,
		)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, newFormGroup, classdiagramFormCallback.probe)
		classdiagramFormCallback.probe.formStage.Commit()
	}

	classdiagramFormCallback.probe.ux_tree()
}
func __gong__New__DiagramPackageFormCallback(
	diagrampackage *models.DiagramPackage,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagrampackageFormCallback *DiagramPackageFormCallback) {
	diagrampackageFormCallback = new(DiagramPackageFormCallback)
	diagrampackageFormCallback.probe = probe
	diagrampackageFormCallback.diagrampackage = diagrampackage
	diagrampackageFormCallback.formGroup = formGroup

	diagrampackageFormCallback.CreationMode = (diagrampackage == nil)

	return
}

type DiagramPackageFormCallback struct {
	diagrampackage *models.DiagramPackage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagrampackageFormCallback *DiagramPackageFormCallback) OnSave() {
	diagrampackageFormCallback.probe.stageOfInterest.Lock()
	defer diagrampackageFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramPackageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagrampackageFormCallback.probe.formStage.Checkout()

	if diagrampackageFormCallback.diagrampackage == nil {
		diagrampackageFormCallback.diagrampackage = new(models.DiagramPackage).Stage(diagrampackageFormCallback.probe.stageOfInterest)
	}
	diagrampackage_ := diagrampackageFormCallback.diagrampackage
	_ = diagrampackage_

	for _, formDiv := range diagrampackageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagrampackage_.Name), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(diagrampackage_.Path), formDiv)
		case "GongModelPath":
			FormDivBasicFieldToField(&(diagrampackage_.GongModelPath), formDiv)
		case "Classdiagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](diagrampackageFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Classdiagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Classdiagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagrampackageFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Classdiagram](diagrampackageFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagrampackage_.Classdiagrams = instanceSlice
			diagrampackageFormCallback.probe.UpdateSliceOfPointersCallback(diagrampackage_, "Classdiagrams", &diagrampackage_.Classdiagrams)

		case "SelectedClassdiagram":
			FormDivSelectFieldToField(&(diagrampackage_.SelectedClassdiagram), diagrampackageFormCallback.probe.stageOfInterest, formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	// manage the suppress operation
	if diagrampackageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagrampackage_.Unstage(diagrampackageFormCallback.probe.stageOfInterest)
	}

	diagrampackageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramPackage](
		diagrampackageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagrampackageFormCallback.CreationMode || diagrampackageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagrampackageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagrampackageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramPackageFormCallback(
			nil,
			diagrampackageFormCallback.probe,
			newFormGroup,
		)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, newFormGroup, diagrampackageFormCallback.probe)
		diagrampackageFormCallback.probe.formStage.Commit()
	}

	diagrampackageFormCallback.probe.ux_tree()
}
func __gong__New__GongEnumShapeFormCallback(
	gongenumshape *models.GongEnumShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumshapeFormCallback *GongEnumShapeFormCallback) {
	gongenumshapeFormCallback = new(GongEnumShapeFormCallback)
	gongenumshapeFormCallback.probe = probe
	gongenumshapeFormCallback.gongenumshape = gongenumshape
	gongenumshapeFormCallback.formGroup = formGroup

	gongenumshapeFormCallback.CreationMode = (gongenumshape == nil)

	return
}

type GongEnumShapeFormCallback struct {
	gongenumshape *models.GongEnumShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongenumshapeFormCallback *GongEnumShapeFormCallback) OnSave() {
	gongenumshapeFormCallback.probe.stageOfInterest.Lock()
	defer gongenumshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongEnumShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumshapeFormCallback.probe.formStage.Checkout()

	if gongenumshapeFormCallback.gongenumshape == nil {
		gongenumshapeFormCallback.gongenumshape = new(models.GongEnumShape).Stage(gongenumshapeFormCallback.probe.stageOfInterest)
	}
	gongenumshape_ := gongenumshapeFormCallback.gongenumshape
	_ = gongenumshape_

	for _, formDiv := range gongenumshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongenumshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongenumshape_.Y), formDiv)
		case "GongEnumValueShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumValueShape](gongenumshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongEnumValueShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongEnumValueShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongenumshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongEnumValueShape](gongenumshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongenumshape_.GongEnumValueShapes = instanceSlice
			gongenumshapeFormCallback.probe.UpdateSliceOfPointersCallback(gongenumshape_, "GongEnumValueShapes", &gongenumshape_.GongEnumValueShapes)

		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongenumshape_.Height), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(gongenumshape_.IsExpanded), formDiv)
		case "Classdiagram:GongEnumShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Classdiagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Classdiagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest)
			targetClassdiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetClassdiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Classdiagram instances and update their GongEnumShapes slice
			for _classdiagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongenumshapeFormCallback.probe.stageOfInterest, _classdiagram)
				
				// if Classdiagram is selected
				if targetClassdiagramIDs[id] {
					// ensure gongenumshape_ is in _classdiagram.GongEnumShapes
					found := false
					for _, _b := range _classdiagram.GongEnumShapes {
						if _b == gongenumshape_ {
							found = true
							break
						}
					}
					if !found {
						_classdiagram.GongEnumShapes = append(_classdiagram.GongEnumShapes, gongenumshape_)
						gongenumshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongEnumShapes", &_classdiagram.GongEnumShapes)
					}
				} else {
					// ensure gongenumshape_ is NOT in _classdiagram.GongEnumShapes
					idx := slices.Index(_classdiagram.GongEnumShapes, gongenumshape_)
					if idx != -1 {
						_classdiagram.GongEnumShapes = slices.Delete(_classdiagram.GongEnumShapes, idx, idx+1)
						gongenumshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongEnumShapes", &_classdiagram.GongEnumShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumshape_.Unstage(gongenumshapeFormCallback.probe.stageOfInterest)
	}

	gongenumshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongEnumShape](
		gongenumshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongenumshapeFormCallback.CreationMode || gongenumshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongenumshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumShapeFormCallback(
			nil,
			gongenumshapeFormCallback.probe,
			newFormGroup,
		)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, newFormGroup, gongenumshapeFormCallback.probe)
		gongenumshapeFormCallback.probe.formStage.Commit()
	}

	gongenumshapeFormCallback.probe.ux_tree()
}
func __gong__New__GongEnumValueShapeFormCallback(
	gongenumvalueshape *models.GongEnumValueShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumvalueshapeFormCallback *GongEnumValueShapeFormCallback) {
	gongenumvalueshapeFormCallback = new(GongEnumValueShapeFormCallback)
	gongenumvalueshapeFormCallback.probe = probe
	gongenumvalueshapeFormCallback.gongenumvalueshape = gongenumvalueshape
	gongenumvalueshapeFormCallback.formGroup = formGroup

	gongenumvalueshapeFormCallback.CreationMode = (gongenumvalueshape == nil)

	return
}

type GongEnumValueShapeFormCallback struct {
	gongenumvalueshape *models.GongEnumValueShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongenumvalueshapeFormCallback *GongEnumValueShapeFormCallback) OnSave() {
	gongenumvalueshapeFormCallback.probe.stageOfInterest.Lock()
	defer gongenumvalueshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongEnumValueShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueshapeFormCallback.probe.formStage.Checkout()

	if gongenumvalueshapeFormCallback.gongenumvalueshape == nil {
		gongenumvalueshapeFormCallback.gongenumvalueshape = new(models.GongEnumValueShape).Stage(gongenumvalueshapeFormCallback.probe.stageOfInterest)
	}
	gongenumvalueshape_ := gongenumvalueshapeFormCallback.gongenumvalueshape
	_ = gongenumvalueshape_

	for _, formDiv := range gongenumvalueshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueshape_.Name), formDiv)
		case "GongEnumShape:GongEnumValueShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongEnumShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongEnumShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongEnumShape](gongenumvalueshapeFormCallback.probe.stageOfInterest)
			targetGongEnumShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongEnumShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongEnumShape instances and update their GongEnumValueShapes slice
			for _gongenumshape := range *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumShape](gongenumvalueshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongenumvalueshapeFormCallback.probe.stageOfInterest, _gongenumshape)
				
				// if GongEnumShape is selected
				if targetGongEnumShapeIDs[id] {
					// ensure gongenumvalueshape_ is in _gongenumshape.GongEnumValueShapes
					found := false
					for _, _b := range _gongenumshape.GongEnumValueShapes {
						if _b == gongenumvalueshape_ {
							found = true
							break
						}
					}
					if !found {
						_gongenumshape.GongEnumValueShapes = append(_gongenumshape.GongEnumValueShapes, gongenumvalueshape_)
						gongenumvalueshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongenumshape, "GongEnumValueShapes", &_gongenumshape.GongEnumValueShapes)
					}
				} else {
					// ensure gongenumvalueshape_ is NOT in _gongenumshape.GongEnumValueShapes
					idx := slices.Index(_gongenumshape.GongEnumValueShapes, gongenumvalueshape_)
					if idx != -1 {
						_gongenumshape.GongEnumValueShapes = slices.Delete(_gongenumshape.GongEnumValueShapes, idx, idx+1)
						gongenumvalueshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongenumshape, "GongEnumValueShapes", &_gongenumshape.GongEnumValueShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumvalueshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueshape_.Unstage(gongenumvalueshapeFormCallback.probe.stageOfInterest)
	}

	gongenumvalueshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongEnumValueShape](
		gongenumvalueshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongenumvalueshapeFormCallback.CreationMode || gongenumvalueshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongenumvalueshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumValueShapeFormCallback(
			nil,
			gongenumvalueshapeFormCallback.probe,
			newFormGroup,
		)
		gongenumvalueshape := new(models.GongEnumValueShape)
		FillUpForm(gongenumvalueshape, newFormGroup, gongenumvalueshapeFormCallback.probe)
		gongenumvalueshapeFormCallback.probe.formStage.Commit()
	}

	gongenumvalueshapeFormCallback.probe.ux_tree()
}
func __gong__New__GongNoteLinkShapeFormCallback(
	gongnotelinkshape *models.GongNoteLinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnotelinkshapeFormCallback *GongNoteLinkShapeFormCallback) {
	gongnotelinkshapeFormCallback = new(GongNoteLinkShapeFormCallback)
	gongnotelinkshapeFormCallback.probe = probe
	gongnotelinkshapeFormCallback.gongnotelinkshape = gongnotelinkshape
	gongnotelinkshapeFormCallback.formGroup = formGroup

	gongnotelinkshapeFormCallback.CreationMode = (gongnotelinkshape == nil)

	return
}

type GongNoteLinkShapeFormCallback struct {
	gongnotelinkshape *models.GongNoteLinkShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongnotelinkshapeFormCallback *GongNoteLinkShapeFormCallback) OnSave() {
	gongnotelinkshapeFormCallback.probe.stageOfInterest.Lock()
	defer gongnotelinkshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongNoteLinkShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnotelinkshapeFormCallback.probe.formStage.Checkout()

	if gongnotelinkshapeFormCallback.gongnotelinkshape == nil {
		gongnotelinkshapeFormCallback.gongnotelinkshape = new(models.GongNoteLinkShape).Stage(gongnotelinkshapeFormCallback.probe.stageOfInterest)
	}
	gongnotelinkshape_ := gongnotelinkshapeFormCallback.gongnotelinkshape
	_ = gongnotelinkshape_

	for _, formDiv := range gongnotelinkshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnotelinkshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongnotelinkshape_.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(gongnotelinkshape_.Type), formDiv)
		case "GongNoteShape:GongNoteLinkShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongNoteShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongNoteShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongNoteShape](gongnotelinkshapeFormCallback.probe.stageOfInterest)
			targetGongNoteShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongNoteShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongNoteShape instances and update their GongNoteLinkShapes slice
			for _gongnoteshape := range *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteShape](gongnotelinkshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongnotelinkshapeFormCallback.probe.stageOfInterest, _gongnoteshape)
				
				// if GongNoteShape is selected
				if targetGongNoteShapeIDs[id] {
					// ensure gongnotelinkshape_ is in _gongnoteshape.GongNoteLinkShapes
					found := false
					for _, _b := range _gongnoteshape.GongNoteLinkShapes {
						if _b == gongnotelinkshape_ {
							found = true
							break
						}
					}
					if !found {
						_gongnoteshape.GongNoteLinkShapes = append(_gongnoteshape.GongNoteLinkShapes, gongnotelinkshape_)
						gongnotelinkshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongnoteshape, "GongNoteLinkShapes", &_gongnoteshape.GongNoteLinkShapes)
					}
				} else {
					// ensure gongnotelinkshape_ is NOT in _gongnoteshape.GongNoteLinkShapes
					idx := slices.Index(_gongnoteshape.GongNoteLinkShapes, gongnotelinkshape_)
					if idx != -1 {
						_gongnoteshape.GongNoteLinkShapes = slices.Delete(_gongnoteshape.GongNoteLinkShapes, idx, idx+1)
						gongnotelinkshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongnoteshape, "GongNoteLinkShapes", &_gongnoteshape.GongNoteLinkShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongnotelinkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnotelinkshape_.Unstage(gongnotelinkshapeFormCallback.probe.stageOfInterest)
	}

	gongnotelinkshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongNoteLinkShape](
		gongnotelinkshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongnotelinkshapeFormCallback.CreationMode || gongnotelinkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnotelinkshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongnotelinkshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteLinkShapeFormCallback(
			nil,
			gongnotelinkshapeFormCallback.probe,
			newFormGroup,
		)
		gongnotelinkshape := new(models.GongNoteLinkShape)
		FillUpForm(gongnotelinkshape, newFormGroup, gongnotelinkshapeFormCallback.probe)
		gongnotelinkshapeFormCallback.probe.formStage.Commit()
	}

	gongnotelinkshapeFormCallback.probe.ux_tree()
}
func __gong__New__GongNoteShapeFormCallback(
	gongnoteshape *models.GongNoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnoteshapeFormCallback *GongNoteShapeFormCallback) {
	gongnoteshapeFormCallback = new(GongNoteShapeFormCallback)
	gongnoteshapeFormCallback.probe = probe
	gongnoteshapeFormCallback.gongnoteshape = gongnoteshape
	gongnoteshapeFormCallback.formGroup = formGroup

	gongnoteshapeFormCallback.CreationMode = (gongnoteshape == nil)

	return
}

type GongNoteShapeFormCallback struct {
	gongnoteshape *models.GongNoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongnoteshapeFormCallback *GongNoteShapeFormCallback) OnSave() {
	gongnoteshapeFormCallback.probe.stageOfInterest.Lock()
	defer gongnoteshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongNoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteshapeFormCallback.probe.formStage.Checkout()

	if gongnoteshapeFormCallback.gongnoteshape == nil {
		gongnoteshapeFormCallback.gongnoteshape = new(models.GongNoteShape).Stage(gongnoteshapeFormCallback.probe.stageOfInterest)
	}
	gongnoteshape_ := gongnoteshapeFormCallback.gongnoteshape
	_ = gongnoteshape_

	for _, formDiv := range gongnoteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnoteshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongnoteshape_.Identifier), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnoteshape_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnoteshape_.BodyHTML), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongnoteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongnoteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongnoteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongnoteshape_.Height), formDiv)
		case "Matched":
			FormDivBasicFieldToField(&(gongnoteshape_.Matched), formDiv)
		case "GongNoteLinkShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteLinkShape](gongnoteshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongNoteLinkShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongNoteLinkShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongnoteshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GongNoteLinkShape](gongnoteshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongnoteshape_.GongNoteLinkShapes = instanceSlice
			gongnoteshapeFormCallback.probe.UpdateSliceOfPointersCallback(gongnoteshape_, "GongNoteLinkShapes", &gongnoteshape_.GongNoteLinkShapes)

		case "IsExpanded":
			FormDivBasicFieldToField(&(gongnoteshape_.IsExpanded), formDiv)
		case "Classdiagram:GongNoteShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Classdiagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Classdiagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Classdiagram](gongnoteshapeFormCallback.probe.stageOfInterest)
			targetClassdiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetClassdiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Classdiagram instances and update their GongNoteShapes slice
			for _classdiagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](gongnoteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongnoteshapeFormCallback.probe.stageOfInterest, _classdiagram)
				
				// if Classdiagram is selected
				if targetClassdiagramIDs[id] {
					// ensure gongnoteshape_ is in _classdiagram.GongNoteShapes
					found := false
					for _, _b := range _classdiagram.GongNoteShapes {
						if _b == gongnoteshape_ {
							found = true
							break
						}
					}
					if !found {
						_classdiagram.GongNoteShapes = append(_classdiagram.GongNoteShapes, gongnoteshape_)
						gongnoteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongNoteShapes", &_classdiagram.GongNoteShapes)
					}
				} else {
					// ensure gongnoteshape_ is NOT in _classdiagram.GongNoteShapes
					idx := slices.Index(_classdiagram.GongNoteShapes, gongnoteshape_)
					if idx != -1 {
						_classdiagram.GongNoteShapes = slices.Delete(_classdiagram.GongNoteShapes, idx, idx+1)
						gongnoteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongNoteShapes", &_classdiagram.GongNoteShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongnoteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteshape_.Unstage(gongnoteshapeFormCallback.probe.stageOfInterest)
	}

	gongnoteshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongNoteShape](
		gongnoteshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongnoteshapeFormCallback.CreationMode || gongnoteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongnoteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteShapeFormCallback(
			nil,
			gongnoteshapeFormCallback.probe,
			newFormGroup,
		)
		gongnoteshape := new(models.GongNoteShape)
		FillUpForm(gongnoteshape, newFormGroup, gongnoteshapeFormCallback.probe)
		gongnoteshapeFormCallback.probe.formStage.Commit()
	}

	gongnoteshapeFormCallback.probe.ux_tree()
}
func __gong__New__GongStructShapeFormCallback(
	gongstructshape *models.GongStructShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongstructshapeFormCallback *GongStructShapeFormCallback) {
	gongstructshapeFormCallback = new(GongStructShapeFormCallback)
	gongstructshapeFormCallback.probe = probe
	gongstructshapeFormCallback.gongstructshape = gongstructshape
	gongstructshapeFormCallback.formGroup = formGroup

	gongstructshapeFormCallback.CreationMode = (gongstructshape == nil)

	return
}

type GongStructShapeFormCallback struct {
	gongstructshape *models.GongStructShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gongstructshapeFormCallback *GongStructShapeFormCallback) OnSave() {
	gongstructshapeFormCallback.probe.stageOfInterest.Lock()
	defer gongstructshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GongStructShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructshapeFormCallback.probe.formStage.Checkout()

	if gongstructshapeFormCallback.gongstructshape == nil {
		gongstructshapeFormCallback.gongstructshape = new(models.GongStructShape).Stage(gongstructshapeFormCallback.probe.stageOfInterest)
	}
	gongstructshape_ := gongstructshapeFormCallback.gongstructshape
	_ = gongstructshape_

	for _, formDiv := range gongstructshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstructshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongstructshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongstructshape_.Y), formDiv)
		case "AttributeShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeShape](gongstructshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AttributeShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AttributeShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AttributeShape](gongstructshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstructshape_.AttributeShapes = instanceSlice
			gongstructshapeFormCallback.probe.UpdateSliceOfPointersCallback(gongstructshape_, "AttributeShapes", &gongstructshape_.AttributeShapes)

		case "LinkShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkShape](gongstructshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.LinkShape](gongstructshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gongstructshape_.LinkShapes = instanceSlice
			gongstructshapeFormCallback.probe.UpdateSliceOfPointersCallback(gongstructshape_, "LinkShapes", &gongstructshape_.LinkShapes)

		case "Width":
			FormDivBasicFieldToField(&(gongstructshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongstructshape_.Height), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(gongstructshape_.IsSelected), formDiv)
		case "Classdiagram:GongStructShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Classdiagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Classdiagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest)
			targetClassdiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetClassdiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Classdiagram instances and update their GongStructShapes slice
			for _classdiagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(gongstructshapeFormCallback.probe.stageOfInterest, _classdiagram)
				
				// if Classdiagram is selected
				if targetClassdiagramIDs[id] {
					// ensure gongstructshape_ is in _classdiagram.GongStructShapes
					found := false
					for _, _b := range _classdiagram.GongStructShapes {
						if _b == gongstructshape_ {
							found = true
							break
						}
					}
					if !found {
						_classdiagram.GongStructShapes = append(_classdiagram.GongStructShapes, gongstructshape_)
						gongstructshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongStructShapes", &_classdiagram.GongStructShapes)
					}
				} else {
					// ensure gongstructshape_ is NOT in _classdiagram.GongStructShapes
					idx := slices.Index(_classdiagram.GongStructShapes, gongstructshape_)
					if idx != -1 {
						_classdiagram.GongStructShapes = slices.Delete(_classdiagram.GongStructShapes, idx, idx+1)
						gongstructshapeFormCallback.probe.UpdateSliceOfPointersCallback(_classdiagram, "GongStructShapes", &_classdiagram.GongStructShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongstructshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructshape_.Unstage(gongstructshapeFormCallback.probe.stageOfInterest)
	}

	gongstructshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GongStructShape](
		gongstructshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gongstructshapeFormCallback.CreationMode || gongstructshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gongstructshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongStructShapeFormCallback(
			nil,
			gongstructshapeFormCallback.probe,
			newFormGroup,
		)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, newFormGroup, gongstructshapeFormCallback.probe)
		gongstructshapeFormCallback.probe.formStage.Commit()
	}

	gongstructshapeFormCallback.probe.ux_tree()
}
func __gong__New__LinkShapeFormCallback(
	linkshape *models.LinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkshapeFormCallback *LinkShapeFormCallback) {
	linkshapeFormCallback = new(LinkShapeFormCallback)
	linkshapeFormCallback.probe = probe
	linkshapeFormCallback.linkshape = linkshape
	linkshapeFormCallback.formGroup = formGroup

	linkshapeFormCallback.CreationMode = (linkshape == nil)

	return
}

type LinkShapeFormCallback struct {
	linkshape *models.LinkShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkshapeFormCallback *LinkShapeFormCallback) OnSave() {
	linkshapeFormCallback.probe.stageOfInterest.Lock()
	defer linkshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkshapeFormCallback.probe.formStage.Checkout()

	if linkshapeFormCallback.linkshape == nil {
		linkshapeFormCallback.linkshape = new(models.LinkShape).Stage(linkshapeFormCallback.probe.stageOfInterest)
	}
	linkshape_ := linkshapeFormCallback.linkshape
	_ = linkshape_

	for _, formDiv := range linkshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkshape_.Name), formDiv)
		case "FieldOffsetX":
			FormDivBasicFieldToField(&(linkshape_.FieldOffsetX), formDiv)
		case "FieldOffsetY":
			FormDivBasicFieldToField(&(linkshape_.FieldOffsetY), formDiv)
		case "TargetMultiplicity":
			FormDivEnumStringFieldToField(&(linkshape_.TargetMultiplicity), formDiv)
		case "TargetMultiplicityOffsetX":
			FormDivBasicFieldToField(&(linkshape_.TargetMultiplicityOffsetX), formDiv)
		case "TargetMultiplicityOffsetY":
			FormDivBasicFieldToField(&(linkshape_.TargetMultiplicityOffsetY), formDiv)
		case "SourceMultiplicity":
			FormDivEnumStringFieldToField(&(linkshape_.SourceMultiplicity), formDiv)
		case "SourceMultiplicityOffsetX":
			FormDivBasicFieldToField(&(linkshape_.SourceMultiplicityOffsetX), formDiv)
		case "SourceMultiplicityOffsetY":
			FormDivBasicFieldToField(&(linkshape_.SourceMultiplicityOffsetY), formDiv)
		case "X":
			FormDivBasicFieldToField(&(linkshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(linkshape_.Y), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(linkshape_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(linkshape_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(linkshape_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(linkshape_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(linkshape_.CornerOffsetRatio), formDiv)
		case "GongStructShape:LinkShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GongStructShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GongStructShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GongStructShape](linkshapeFormCallback.probe.stageOfInterest)
			targetGongStructShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGongStructShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GongStructShape instances and update their LinkShapes slice
			for _gongstructshape := range *models.GetGongstructInstancesSetFromPointerType[*models.GongStructShape](linkshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkshapeFormCallback.probe.stageOfInterest, _gongstructshape)
				
				// if GongStructShape is selected
				if targetGongStructShapeIDs[id] {
					// ensure linkshape_ is in _gongstructshape.LinkShapes
					found := false
					for _, _b := range _gongstructshape.LinkShapes {
						if _b == linkshape_ {
							found = true
							break
						}
					}
					if !found {
						_gongstructshape.LinkShapes = append(_gongstructshape.LinkShapes, linkshape_)
						linkshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongstructshape, "LinkShapes", &_gongstructshape.LinkShapes)
					}
				} else {
					// ensure linkshape_ is NOT in _gongstructshape.LinkShapes
					idx := slices.Index(_gongstructshape.LinkShapes, linkshape_)
					if idx != -1 {
						_gongstructshape.LinkShapes = slices.Delete(_gongstructshape.LinkShapes, idx, idx+1)
						linkshapeFormCallback.probe.UpdateSliceOfPointersCallback(_gongstructshape, "LinkShapes", &_gongstructshape.LinkShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkshape_.Unstage(linkshapeFormCallback.probe.stageOfInterest)
	}

	linkshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.LinkShape](
		linkshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkshapeFormCallback.CreationMode || linkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkShapeFormCallback(
			nil,
			linkshapeFormCallback.probe,
			newFormGroup,
		)
		linkshape := new(models.LinkShape)
		FillUpForm(linkshape, newFormGroup, linkshapeFormCallback.probe)
		linkshapeFormCallback.probe.formStage.Commit()
	}

	linkshapeFormCallback.probe.ux_tree()
}
