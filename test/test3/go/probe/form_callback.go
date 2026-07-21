// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AFormCallback(
	a *models.A,
	probe *Probe,
	formGroup *form.FormGroup,
) (aFormCallback *AFormCallback) {
	aFormCallback = new(AFormCallback)
	aFormCallback.probe = probe
	aFormCallback.a = a
	aFormCallback.formGroup = formGroup

	aFormCallback.CreationMode = (a == nil)

	return
}

type AFormCallback struct {
	a *models.A

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (aFormCallback *AFormCallback) OnSave() {
	aFormCallback.probe.stageOfInterest.Lock()
	defer aFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	aFormCallback.probe.formStage.Checkout()

	if aFormCallback.a == nil {
		aFormCallback.a = new(models.A).Stage(aFormCallback.probe.stageOfInterest)
	}
	a_ := aFormCallback.a
	_ = a_

	for _, formDiv := range aFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_.Name), formDiv)
		case "Date":
			FormDivTimeFieldToField(&(a_.Date), formDiv, true)
		case "Duration":
			FormDivBasicFieldToField(&(a_.Duration), formDiv)
		case "FloatValue":
			FormDivBasicFieldToField(&(a_.FloatValue), formDiv)
		case "IntValue":
			FormDivBasicFieldToField(&(a_.IntValue), formDiv)
		case "EnumString":
			FormDivEnumStringFieldToField(&(a_.EnumString), formDiv)
		case "EnumInt":
			FormDivEnumIntFieldToField(&(a_.EnumInt), formDiv)
		case "B":
			FormDivSelectFieldToField(&(a_.B), aFormCallback.probe.stageOfInterest, formDiv)
		case "Bs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.B](aFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.B, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.B)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					aFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.B](aFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_.Bs = instanceSlice
			aFormCallback.probe.UpdateSliceOfPointersCallback(a_, "Bs", &a_.Bs)

		case "C":
			FormDivSelectFieldToField(&(a_.C), aFormCallback.probe.stageOfInterest, formDiv)
		case "Cs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.C](aFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.C, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.C)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					aFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.C](aFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_.Cs = instanceSlice
			aFormCallback.probe.UpdateSliceOfPointersCallback(a_, "Cs", &a_.Cs)

		case "UUID":
			FormDivBasicFieldToField(&(a_.UUID), formDiv)
		}
	}

	// manage the suppress operation
	if aFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_.Unstage(aFormCallback.probe.stageOfInterest)
	}

	aFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A](
		aFormCallback.probe,
	)

	// display a new form by reset the form stage
	if aFormCallback.CreationMode || aFormCallback.formGroup.HasSuppressButtonBeenPressed {
		aFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(aFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AFormCallback(
			nil,
			aFormCallback.probe,
			newFormGroup,
		)
		a := new(models.A)
		FillUpForm(a, newFormGroup, aFormCallback.probe)
		aFormCallback.probe.formStage.Commit()
	}

	aFormCallback.probe.ux_tree()
}
func __gong__New__BFormCallback(
	b *models.B,
	probe *Probe,
	formGroup *form.FormGroup,
) (bFormCallback *BFormCallback) {
	bFormCallback = new(BFormCallback)
	bFormCallback.probe = probe
	bFormCallback.b = b
	bFormCallback.formGroup = formGroup

	bFormCallback.CreationMode = (b == nil)

	return
}

type BFormCallback struct {
	b *models.B

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bFormCallback *BFormCallback) OnSave() {
	bFormCallback.probe.stageOfInterest.Lock()
	defer bFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bFormCallback.probe.formStage.Checkout()

	if bFormCallback.b == nil {
		bFormCallback.b = new(models.B).Stage(bFormCallback.probe.stageOfInterest)
	}
	b_ := bFormCallback.b
	_ = b_

	for _, formDiv := range bFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(b_.Name), formDiv)
		case "A:Bs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the A instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target A instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.A](bFormCallback.probe.stageOfInterest)
			targetAIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all A instances and update their Bs slice
			for _a := range *models.GetGongstructInstancesSetFromPointerType[*models.A](bFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bFormCallback.probe.stageOfInterest, _a)
				
				// if A is selected
				if targetAIDs[id] {
					// ensure b_ is in _a.Bs
					found := false
					for _, _b := range _a.Bs {
						if _b == b_ {
							found = true
							break
						}
					}
					if !found {
						_a.Bs = append(_a.Bs, b_)
						bFormCallback.probe.UpdateSliceOfPointersCallback(_a, "Bs", &_a.Bs)
					}
				} else {
					// ensure b_ is NOT in _a.Bs
					idx := slices.Index(_a.Bs, b_)
					if idx != -1 {
						_a.Bs = slices.Delete(_a.Bs, idx, idx+1)
						bFormCallback.probe.UpdateSliceOfPointersCallback(_a, "Bs", &_a.Bs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		b_.Unstage(bFormCallback.probe.stageOfInterest)
	}

	bFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.B](
		bFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bFormCallback.CreationMode || bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BFormCallback(
			nil,
			bFormCallback.probe,
			newFormGroup,
		)
		b := new(models.B)
		FillUpForm(b, newFormGroup, bFormCallback.probe)
		bFormCallback.probe.formStage.Commit()
	}

	bFormCallback.probe.ux_tree()
}
func __gong__New__CFormCallback(
	c *models.C,
	probe *Probe,
	formGroup *form.FormGroup,
) (cFormCallback *CFormCallback) {
	cFormCallback = new(CFormCallback)
	cFormCallback.probe = probe
	cFormCallback.c = c
	cFormCallback.formGroup = formGroup

	cFormCallback.CreationMode = (c == nil)

	return
}

type CFormCallback struct {
	c *models.C

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (cFormCallback *CFormCallback) OnSave() {
	cFormCallback.probe.stageOfInterest.Lock()
	defer cFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cFormCallback.probe.formStage.Checkout()

	if cFormCallback.c == nil {
		cFormCallback.c = new(models.C).Stage(cFormCallback.probe.stageOfInterest)
	}
	c_ := cFormCallback.c
	_ = c_

	for _, formDiv := range cFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(c_.Name), formDiv)
		case "A:Cs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the A instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target A instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.A](cFormCallback.probe.stageOfInterest)
			targetAIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all A instances and update their Cs slice
			for _a := range *models.GetGongstructInstancesSetFromPointerType[*models.A](cFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(cFormCallback.probe.stageOfInterest, _a)
				
				// if A is selected
				if targetAIDs[id] {
					// ensure c_ is in _a.Cs
					found := false
					for _, _b := range _a.Cs {
						if _b == c_ {
							found = true
							break
						}
					}
					if !found {
						_a.Cs = append(_a.Cs, c_)
						cFormCallback.probe.UpdateSliceOfPointersCallback(_a, "Cs", &_a.Cs)
					}
				} else {
					// ensure c_ is NOT in _a.Cs
					idx := slices.Index(_a.Cs, c_)
					if idx != -1 {
						_a.Cs = slices.Delete(_a.Cs, idx, idx+1)
						cFormCallback.probe.UpdateSliceOfPointersCallback(_a, "Cs", &_a.Cs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if cFormCallback.formGroup.HasSuppressButtonBeenPressed {
		c_.Unstage(cFormCallback.probe.stageOfInterest)
	}

	cFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.C](
		cFormCallback.probe,
	)

	// display a new form by reset the form stage
	if cFormCallback.CreationMode || cFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CFormCallback(
			nil,
			cFormCallback.probe,
			newFormGroup,
		)
		c := new(models.C)
		FillUpForm(c, newFormGroup, cFormCallback.probe)
		cFormCallback.probe.formStage.Commit()
	}

	cFormCallback.probe.ux_tree()
}
