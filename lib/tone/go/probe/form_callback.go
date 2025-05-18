// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__FreqencyFormCallback(
	freqency *models.Freqency,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (freqencyFormCallback *FreqencyFormCallback) OnSave() {

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
			// WARNING : this form deals with the N-N association "Note.Frequencies []*Freqency" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Freqency). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Note
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Note"
				rf.Fieldname = "Frequencies"
				formerAssociationSource := models.GetReverseFieldOwner(
					freqencyFormCallback.probe.stageOfInterest,
					freqency_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Note)
					if !ok {
						log.Fatalln("Source of Note.Frequencies []*Freqency, is not an Note instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Frequencies, freqency_)
					formerSource.Frequencies = slices.Delete(formerSource.Frequencies, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
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
			for _note := range *models.GetGongstructInstancesSet[models.Note](freqencyFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _note.GetName() == fieldValue.GetName() {
							newNoteOwner := _note // we have a match
							
							// we remove the freqency_ instance from the pastNoteOwner field
							if pastNoteOwner != nil {
								if newNoteOwner != pastNoteOwner {
									idx := slices.Index(pastNoteOwner.Frequencies, freqency_)
									pastNoteOwner.Frequencies = slices.Delete(pastNoteOwner.Frequencies, idx, idx+1)
									newNoteOwner.Frequencies = append(newNoteOwner.Frequencies, freqency_)
								}
							} else {
								newNoteOwner.Frequencies = append(newNoteOwner.Frequencies, freqency_)
							}
						}
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
	updateAndCommitTable[models.Freqency](
		freqencyFormCallback.probe,
	)
	freqencyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if freqencyFormCallback.CreationMode || freqencyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		freqencyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(freqencyFormCallback.probe)
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			note_.Frequencies = instanceSlice

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
	updateAndCommitTable[models.Note](
		noteFormCallback.probe,
	)
	noteFormCallback.probe.tableStage.Commit()

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
func __gong__New__PlayerFormCallback(
	player *models.Player,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (playerFormCallback *PlayerFormCallback) OnSave() {

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
	updateAndCommitTable[models.Player](
		playerFormCallback.probe,
	)
	playerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if playerFormCallback.CreationMode || playerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		playerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(playerFormCallback.probe)
}
