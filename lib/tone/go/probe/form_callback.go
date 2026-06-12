// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__FreqencyFormCallback(
	freqency *models.Freqency,
	probe *Probe,
	formGroup *form.FormGroup,
) (freqencyFormCallback *FreqencyFormCallback) {
	freqencyFormCallback = new(FreqencyFormCallback)
	freqencyFormCallback.probe = probe
	freqencyFormCallback.freqency = freqency
	freqencyFormCallback.formGroup = formGroup

	freqencyFormCallback.CreationMode = (freqency == nil)

	return
}

type FreqencyFormCallback struct {
	freqency *models.Freqency

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (freqencyFormCallback *FreqencyFormCallback) OnSave() {
	freqencyFormCallback.probe.stageOfInterest.Lock()
	defer freqencyFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FreqencyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	freqencyFormCallback.probe.formStage.Checkout()

	if freqencyFormCallback.freqency == nil {
		freqencyFormCallback.freqency = new(models.Freqency).Stage(freqencyFormCallback.probe.stageOfInterest)
	}
	freqency_ := freqencyFormCallback.freqency
	_ = freqency_

	for _, formDiv := range freqencyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(freqency_.Name), formDiv)
		case "Note:Frequencies":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](freqencyFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Frequencies slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](freqencyFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(freqencyFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure freqency_ is in _note.Frequencies
					found := false
					for _, _b := range _note.Frequencies {
						if _b == freqency_ {
							found = true
							break
						}
					}
					if !found {
						_note.Frequencies = append(_note.Frequencies, freqency_)
						freqencyFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Frequencies", &_note.Frequencies)
					}
				} else {
					// ensure freqency_ is NOT in _note.Frequencies
					idx := slices.Index(_note.Frequencies, freqency_)
					if idx != -1 {
						_note.Frequencies = slices.Delete(_note.Frequencies, idx, idx+1)
						freqencyFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Frequencies", &_note.Frequencies)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if freqencyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		freqency_.Unstage(freqencyFormCallback.probe.stageOfInterest)
	}

	freqencyFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Freqency](
		freqencyFormCallback.probe,
	)

	// display a new form by reset the form stage
	if freqencyFormCallback.CreationMode || freqencyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		freqencyFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(freqencyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FreqencyFormCallback(
			nil,
			freqencyFormCallback.probe,
			newFormGroup,
		)
		freqency := new(models.Freqency)
		FillUpForm(freqency, newFormGroup, freqencyFormCallback.probe)
		freqencyFormCallback.probe.formStage.Commit()
	}

	freqencyFormCallback.probe.ux_tree()
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
		case "Frequencies":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Freqency](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Freqency, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Freqency)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Freqency](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Frequencies = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Frequencies", &note_.Frequencies)

		case "Start":
			FormDivBasicFieldToField(&(note_.Start), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(note_.Duration), formDiv)
		case "Velocity":
			FormDivBasicFieldToField(&(note_.Velocity), formDiv)
		case "Info":
			FormDivBasicFieldToField(&(note_.Info), formDiv)
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
func __gong__New__PlayerFormCallback(
	player *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) (playerFormCallback *PlayerFormCallback) {
	playerFormCallback = new(PlayerFormCallback)
	playerFormCallback.probe = probe
	playerFormCallback.player = player
	playerFormCallback.formGroup = formGroup

	playerFormCallback.CreationMode = (player == nil)

	return
}

type PlayerFormCallback struct {
	player *models.Player

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (playerFormCallback *PlayerFormCallback) OnSave() {
	playerFormCallback.probe.stageOfInterest.Lock()
	defer playerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PlayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	playerFormCallback.probe.formStage.Checkout()

	if playerFormCallback.player == nil {
		playerFormCallback.player = new(models.Player).Stage(playerFormCallback.probe.stageOfInterest)
	}
	player_ := playerFormCallback.player
	_ = player_

	for _, formDiv := range playerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(player_.Name), formDiv)
		case "Status":
			FormDivEnumStringFieldToField(&(player_.Status), formDiv)
		}
	}

	// manage the suppress operation
	if playerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		player_.Unstage(playerFormCallback.probe.stageOfInterest)
	}

	playerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Player](
		playerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if playerFormCallback.CreationMode || playerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		playerFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(playerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlayerFormCallback(
			nil,
			playerFormCallback.probe,
			newFormGroup,
		)
		player := new(models.Player)
		FillUpForm(player, newFormGroup, playerFormCallback.probe)
		playerFormCallback.probe.formStage.Commit()
	}

	playerFormCallback.probe.ux_tree()
}
