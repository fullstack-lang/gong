// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	var zero T
	if cb.Instance == zero {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__A_directiveFormCallback(
	_instance *models.A_directive,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_directiveFormCallback *FormCallback[*models.A_directive]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_directiveFields,
	)
}

type A_directiveFormCallback = FormCallback[*models.A_directive]

func saveA_directiveFields(
	_instance *models.A_directive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Attributes:Directive":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Directive", func(owner *models.Attributes) *[]*models.A_directive { return &owner.Directive })
		}
	}
}

func __gong__New__A_measureFormCallback(
	_instance *models.A_measure,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_measureFormCallback *FormCallback[*models.A_measure]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_measureFields,
	)
}

type A_measureFormCallback = FormCallback[*models.A_measure]

func saveA_measureFields(
	_instance *models.A_measure,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Implicit":
			FormDivEnumStringFieldToField(&(_instance.Implicit), formDiv)
		case "Non_controlling":
			FormDivEnumStringFieldToField(&(_instance.Non_controlling), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Note":
			FormDivSliceOfPointersToField(_instance, "Note", &(_instance.Note), formDiv, probe)
		case "Backup":
			FormDivSliceOfPointersToField(_instance, "Backup", &(_instance.Backup), formDiv, probe)
		case "Forward":
			FormDivSliceOfPointersToField(_instance, "Forward", &(_instance.Forward), formDiv, probe)
		case "Direction":
			FormDivSliceOfPointersToField(_instance, "Direction", &(_instance.Direction), formDiv, probe)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "Harmony":
			FormDivSliceOfPointersToField(_instance, "Harmony", &(_instance.Harmony), formDiv, probe)
		case "Figured_bass":
			FormDivSliceOfPointersToField(_instance, "Figured_bass", &(_instance.Figured_bass), formDiv, probe)
		case "Print":
			FormDivSliceOfPointersToField(_instance, "Print", &(_instance.Print), formDiv, probe)
		case "Sound":
			FormDivSliceOfPointersToField(_instance, "Sound", &(_instance.Sound), formDiv, probe)
		case "Listening":
			FormDivSliceOfPointersToField(_instance, "Listening", &(_instance.Listening), formDiv, probe)
		case "Barline":
			FormDivSliceOfPointersToField(_instance, "Barline", &(_instance.Barline), formDiv, probe)
		case "Grouping":
			FormDivSliceOfPointersToField(_instance, "Grouping", &(_instance.Grouping), formDiv, probe)
		case "Link":
			FormDivSliceOfPointersToField(_instance, "Link", &(_instance.Link), formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(_instance, "Bookmark", &(_instance.Bookmark), formDiv, probe)
		case "A_part:Measure":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Measure", func(owner *models.A_part) *[]*models.A_measure { return &owner.Measure })
		}
	}
}

func __gong__New__A_measure_1FormCallback(
	_instance *models.A_measure_1,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_measure_1FormCallback *FormCallback[*models.A_measure_1]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_measure_1Fields,
	)
}

type A_measure_1FormCallback = FormCallback[*models.A_measure_1]

func saveA_measure_1Fields(
	_instance *models.A_measure_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Implicit":
			FormDivEnumStringFieldToField(&(_instance.Implicit), formDiv)
		case "Non_controlling":
			FormDivEnumStringFieldToField(&(_instance.Non_controlling), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Part":
			FormDivSliceOfPointersToField(_instance, "Part", &(_instance.Part), formDiv, probe)
		case "Score_timewise:Measure":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Measure", func(owner *models.Score_timewise) *[]*models.A_measure_1 { return &owner.Measure })
		}
	}
}

func __gong__New__A_partFormCallback(
	_instance *models.A_part,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_partFormCallback *FormCallback[*models.A_part]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_partFields,
	)
}

type A_partFormCallback = FormCallback[*models.A_part]

func saveA_partFields(
	_instance *models.A_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Measure":
			FormDivSliceOfPointersToField(_instance, "Measure", &(_instance.Measure), formDiv, probe)
		case "Score_partwise:Part":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Part", func(owner *models.Score_partwise) *[]*models.A_part { return &owner.Part })
		}
	}
}

func __gong__New__A_part_1FormCallback(
	_instance *models.A_part_1,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_part_1FormCallback *FormCallback[*models.A_part_1]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_part_1Fields,
	)
}

type A_part_1FormCallback = FormCallback[*models.A_part_1]

func saveA_part_1Fields(
	_instance *models.A_part_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Note":
			FormDivSliceOfPointersToField(_instance, "Note", &(_instance.Note), formDiv, probe)
		case "Backup":
			FormDivSliceOfPointersToField(_instance, "Backup", &(_instance.Backup), formDiv, probe)
		case "Forward":
			FormDivSliceOfPointersToField(_instance, "Forward", &(_instance.Forward), formDiv, probe)
		case "Direction":
			FormDivSliceOfPointersToField(_instance, "Direction", &(_instance.Direction), formDiv, probe)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "Harmony":
			FormDivSliceOfPointersToField(_instance, "Harmony", &(_instance.Harmony), formDiv, probe)
		case "Figured_bass":
			FormDivSliceOfPointersToField(_instance, "Figured_bass", &(_instance.Figured_bass), formDiv, probe)
		case "Print":
			FormDivSliceOfPointersToField(_instance, "Print", &(_instance.Print), formDiv, probe)
		case "Sound":
			FormDivSliceOfPointersToField(_instance, "Sound", &(_instance.Sound), formDiv, probe)
		case "Listening":
			FormDivSliceOfPointersToField(_instance, "Listening", &(_instance.Listening), formDiv, probe)
		case "Barline":
			FormDivSliceOfPointersToField(_instance, "Barline", &(_instance.Barline), formDiv, probe)
		case "Grouping":
			FormDivSliceOfPointersToField(_instance, "Grouping", &(_instance.Grouping), formDiv, probe)
		case "Link":
			FormDivSliceOfPointersToField(_instance, "Link", &(_instance.Link), formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(_instance, "Bookmark", &(_instance.Bookmark), formDiv, probe)
		case "A_measure_1:Part":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Part", func(owner *models.A_measure_1) *[]*models.A_part_1 { return &owner.Part })
		}
	}
}

func __gong__New__AccidentalFormCallback(
	_instance *models.Accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) (accidentalFormCallback *FormCallback[*models.Accidental]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAccidentalFields,
	)
}

type AccidentalFormCallback = FormCallback[*models.Accidental]

func saveAccidentalFields(
	_instance *models.Accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Cautionary":
			FormDivBasicFieldToField(&(_instance.Cautionary), formDiv)
		case "Editorial":
			FormDivEnumStringFieldToField(&(_instance.Editorial), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Bracket":
			FormDivEnumStringFieldToField(&(_instance.Bracket), formDiv)
		case "Size":
			FormDivEnumStringFieldToField(&(_instance.Size), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Accidental_markFormCallback(
	_instance *models.Accidental_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) (accidental_markFormCallback *FormCallback[*models.Accidental_mark]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAccidental_markFields,
	)
}

type Accidental_markFormCallback = FormCallback[*models.Accidental_mark]

func saveAccidental_markFields(
	_instance *models.Accidental_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Bracket":
			FormDivEnumStringFieldToField(&(_instance.Bracket), formDiv)
		case "Size":
			FormDivEnumStringFieldToField(&(_instance.Size), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		case "Notations:Accidental_mark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accidental_mark", func(owner *models.Notations) *[]*models.Accidental_mark { return &owner.Accidental_mark })
		case "Ornaments:Accidental_mark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accidental_mark", func(owner *models.Ornaments) *[]*models.Accidental_mark { return &owner.Accidental_mark })
		}
	}
}

func __gong__New__Accidental_textFormCallback(
	_instance *models.Accidental_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (accidental_textFormCallback *FormCallback[*models.Accidental_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAccidental_textFields,
	)
}

type Accidental_textFormCallback = FormCallback[*models.Accidental_text]

func saveAccidental_textFields(
	_instance *models.Accidental_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Space":
			FormDivBasicFieldToField(&(_instance.Space), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&(_instance.Line_height), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		case "Name_display:Accidental_text":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accidental_text", func(owner *models.Name_display) *[]*models.Accidental_text { return &owner.Accidental_text })
		case "Notehead_text:Accidental_text":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accidental_text", func(owner *models.Notehead_text) *[]*models.Accidental_text { return &owner.Accidental_text })
		}
	}
}

func __gong__New__AccordFormCallback(
	_instance *models.Accord,
	probe *Probe,
	formGroup *form.FormGroup,
) (accordFormCallback *FormCallback[*models.Accord]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAccordFields,
	)
}

type AccordFormCallback = FormCallback[*models.Accord]

func saveAccordFields(
	_instance *models.Accord,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "String":
			FormDivBasicFieldToField(&(_instance.String), formDiv)
		case "Tuning_step":
			FormDivEnumStringFieldToField(&(_instance.Tuning_step), formDiv)
		case "Tuning_alter":
			FormDivBasicFieldToField(&(_instance.Tuning_alter), formDiv)
		case "Tuning_octave":
			FormDivBasicFieldToField(&(_instance.Tuning_octave), formDiv)
		case "Scordatura:Accord":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accord", func(owner *models.Scordatura) *[]*models.Accord { return &owner.Accord })
		}
	}
}

func __gong__New__Accordion_registrationFormCallback(
	_instance *models.Accordion_registration,
	probe *Probe,
	formGroup *form.FormGroup,
) (accordion_registrationFormCallback *FormCallback[*models.Accordion_registration]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAccordion_registrationFields,
	)
}

type Accordion_registrationFormCallback = FormCallback[*models.Accordion_registration]

func saveAccordion_registrationFields(
	_instance *models.Accordion_registration,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Accordion_high":
			FormDivBasicFieldToField(&(_instance.Accordion_high), formDiv)
		case "Accordion_middle":
			FormDivBasicFieldToField(&(_instance.Accordion_middle), formDiv)
		case "Accordion_low":
			FormDivBasicFieldToField(&(_instance.Accordion_low), formDiv)
		}
	}
}

func __gong__New__AppearanceFormCallback(
	_instance *models.Appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) (appearanceFormCallback *FormCallback[*models.Appearance]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAppearanceFields,
	)
}

type AppearanceFormCallback = FormCallback[*models.Appearance]

func saveAppearanceFields(
	_instance *models.Appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Line_width":
			FormDivSliceOfPointersToField(_instance, "Line_width", &(_instance.Line_width), formDiv, probe)
		case "Note_size":
			FormDivSliceOfPointersToField(_instance, "Note_size", &(_instance.Note_size), formDiv, probe)
		case "Distance":
			FormDivSliceOfPointersToField(_instance, "Distance", &(_instance.Distance), formDiv, probe)
		case "Glyph":
			FormDivSliceOfPointersToField(_instance, "Glyph", &(_instance.Glyph), formDiv, probe)
		case "Other_appearance":
			FormDivSliceOfPointersToField(_instance, "Other_appearance", &(_instance.Other_appearance), formDiv, probe)
		}
	}
}

func __gong__New__ArpeggiateFormCallback(
	_instance *models.Arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) (arpeggiateFormCallback *FormCallback[*models.Arpeggiate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArpeggiateFields,
	)
}

type ArpeggiateFormCallback = FormCallback[*models.Arpeggiate]

func saveArpeggiateFields(
	_instance *models.Arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Direction":
			FormDivBasicFieldToField(&(_instance.Direction), formDiv)
		case "Unbroken":
			FormDivEnumStringFieldToField(&(_instance.Unbroken), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Notations:Arpeggiate":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Arpeggiate", func(owner *models.Notations) *[]*models.Arpeggiate { return &owner.Arpeggiate })
		}
	}
}

func __gong__New__ArrowFormCallback(
	_instance *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
) (arrowFormCallback *FormCallback[*models.Arrow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArrowFields,
	)
}

type ArrowFormCallback = FormCallback[*models.Arrow]

func saveArrowFields(
	_instance *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Arrow_direction":
			FormDivBasicFieldToField(&(_instance.Arrow_direction), formDiv)
		case "Arrow_style":
			FormDivBasicFieldToField(&(_instance.Arrow_style), formDiv)
		case "Arrowhead":
			FormDivBasicFieldToField(&(_instance.Arrowhead), formDiv)
		case "Circular_arrow":
			FormDivBasicFieldToField(&(_instance.Circular_arrow), formDiv)
		case "Technical:Arrow":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Arrow", func(owner *models.Technical) *[]*models.Arrow { return &owner.Arrow })
		}
	}
}

func __gong__New__ArticulationsFormCallback(
	_instance *models.Articulations,
	probe *Probe,
	formGroup *form.FormGroup,
) (articulationsFormCallback *FormCallback[*models.Articulations]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArticulationsFields,
	)
}

type ArticulationsFormCallback = FormCallback[*models.Articulations]

func saveArticulationsFields(
	_instance *models.Articulations,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Accent":
			FormDivSliceOfPointersToField(_instance, "Accent", &(_instance.Accent), formDiv, probe)
		case "Strong_accent":
			FormDivSliceOfPointersToField(_instance, "Strong_accent", &(_instance.Strong_accent), formDiv, probe)
		case "Staccato":
			FormDivSliceOfPointersToField(_instance, "Staccato", &(_instance.Staccato), formDiv, probe)
		case "Tenuto":
			FormDivSliceOfPointersToField(_instance, "Tenuto", &(_instance.Tenuto), formDiv, probe)
		case "Detached_legato":
			FormDivSliceOfPointersToField(_instance, "Detached_legato", &(_instance.Detached_legato), formDiv, probe)
		case "Staccatissimo":
			FormDivSliceOfPointersToField(_instance, "Staccatissimo", &(_instance.Staccatissimo), formDiv, probe)
		case "Spiccato":
			FormDivSliceOfPointersToField(_instance, "Spiccato", &(_instance.Spiccato), formDiv, probe)
		case "Scoop":
			FormDivSliceOfPointersToField(_instance, "Scoop", &(_instance.Scoop), formDiv, probe)
		case "Plop":
			FormDivSliceOfPointersToField(_instance, "Plop", &(_instance.Plop), formDiv, probe)
		case "Doit":
			FormDivSliceOfPointersToField(_instance, "Doit", &(_instance.Doit), formDiv, probe)
		case "Falloff":
			FormDivSliceOfPointersToField(_instance, "Falloff", &(_instance.Falloff), formDiv, probe)
		case "Breath_mark":
			FormDivSliceOfPointersToField(_instance, "Breath_mark", &(_instance.Breath_mark), formDiv, probe)
		case "Caesura":
			FormDivSliceOfPointersToField(_instance, "Caesura", &(_instance.Caesura), formDiv, probe)
		case "Stress":
			FormDivSliceOfPointersToField(_instance, "Stress", &(_instance.Stress), formDiv, probe)
		case "Unstress":
			FormDivSliceOfPointersToField(_instance, "Unstress", &(_instance.Unstress), formDiv, probe)
		case "Soft_accent":
			FormDivSliceOfPointersToField(_instance, "Soft_accent", &(_instance.Soft_accent), formDiv, probe)
		case "Other_articulation":
			FormDivSliceOfPointersToField(_instance, "Other_articulation", &(_instance.Other_articulation), formDiv, probe)
		case "Notations:Articulations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Articulations", func(owner *models.Notations) *[]*models.Articulations { return &owner.Articulations })
		}
	}
}

func __gong__New__AssessFormCallback(
	_instance *models.Assess,
	probe *Probe,
	formGroup *form.FormGroup,
) (assessFormCallback *FormCallback[*models.Assess]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAssessFields,
	)
}

type AssessFormCallback = FormCallback[*models.Assess]

func saveAssessFields(
	_instance *models.Assess,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Player":
			FormDivBasicFieldToField(&(_instance.Player), formDiv)
		case "Time_only":
			FormDivBasicFieldToField(&(_instance.Time_only), formDiv)
		case "Listen:Assess":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Assess", func(owner *models.Listen) *[]*models.Assess { return &owner.Assess })
		}
	}
}

func __gong__New__AttributesFormCallback(
	_instance *models.Attributes,
	probe *Probe,
	formGroup *form.FormGroup,
) (attributesFormCallback *FormCallback[*models.Attributes]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAttributesFields,
	)
}

type AttributesFormCallback = FormCallback[*models.Attributes]

func saveAttributesFields(
	_instance *models.Attributes,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&(_instance.Divisions), formDiv)
		case "Key":
			FormDivSliceOfPointersToField(_instance, "Key", &(_instance.Key), formDiv, probe)
		case "Time":
			FormDivSliceOfPointersToField(_instance, "Time", &(_instance.Time), formDiv, probe)
		case "Staves":
			FormDivBasicFieldToField(&(_instance.Staves), formDiv)
		case "Part_symbol":
			FormDivSelectFieldToField(&(_instance.Part_symbol), probe.stageOfInterest, formDiv)
		case "Instruments":
			FormDivBasicFieldToField(&(_instance.Instruments), formDiv)
		case "Clef":
			FormDivSliceOfPointersToField(_instance, "Clef", &(_instance.Clef), formDiv, probe)
		case "Staff_details":
			FormDivSliceOfPointersToField(_instance, "Staff_details", &(_instance.Staff_details), formDiv, probe)
		case "Transpose":
			FormDivSliceOfPointersToField(_instance, "Transpose", &(_instance.Transpose), formDiv, probe)
		case "For_part":
			FormDivSliceOfPointersToField(_instance, "For_part", &(_instance.For_part), formDiv, probe)
		case "Directive":
			FormDivSliceOfPointersToField(_instance, "Directive", &(_instance.Directive), formDiv, probe)
		case "Measure_style":
			FormDivSliceOfPointersToField(_instance, "Measure_style", &(_instance.Measure_style), formDiv, probe)
		case "A_measure:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.A_measure) *[]*models.Attributes { return &owner.Attributes })
		case "A_part_1:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.A_part_1) *[]*models.Attributes { return &owner.Attributes })
		}
	}
}

func __gong__New__BackupFormCallback(
	_instance *models.Backup,
	probe *Probe,
	formGroup *form.FormGroup,
) (backupFormCallback *FormCallback[*models.Backup]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBackupFields,
	)
}

type BackupFormCallback = FormCallback[*models.Backup]

func saveBackupFields(
	_instance *models.Backup,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "A_measure:Backup":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Backup", func(owner *models.A_measure) *[]*models.Backup { return &owner.Backup })
		case "A_part_1:Backup":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Backup", func(owner *models.A_part_1) *[]*models.Backup { return &owner.Backup })
		}
	}
}

func __gong__New__Bar_style_colorFormCallback(
	_instance *models.Bar_style_color,
	probe *Probe,
	formGroup *form.FormGroup,
) (bar_style_colorFormCallback *FormCallback[*models.Bar_style_color]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBar_style_colorFields,
	)
}

type Bar_style_colorFormCallback = FormCallback[*models.Bar_style_color]

func saveBar_style_colorFields(
	_instance *models.Bar_style_color,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__BarlineFormCallback(
	_instance *models.Barline,
	probe *Probe,
	formGroup *form.FormGroup,
) (barlineFormCallback *FormCallback[*models.Barline]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBarlineFields,
	)
}

type BarlineFormCallback = FormCallback[*models.Barline]

func saveBarlineFields(
	_instance *models.Barline,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Location":
			FormDivBasicFieldToField(&(_instance.Location), formDiv)
		case "Segno":
			FormDivBasicFieldToField(&(_instance.Segno), formDiv)
		case "Coda":
			FormDivBasicFieldToField(&(_instance.Coda), formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&(_instance.Divisions), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Bar_style":
			FormDivSelectFieldToField(&(_instance.Bar_style), probe.stageOfInterest, formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Wavy_line":
			FormDivSelectFieldToField(&(_instance.Wavy_line), probe.stageOfInterest, formDiv)
		case "Segno_1":
			FormDivSelectFieldToField(&(_instance.Segno_1), probe.stageOfInterest, formDiv)
		case "Coda_1":
			FormDivSelectFieldToField(&(_instance.Coda_1), probe.stageOfInterest, formDiv)
		case "Fermata":
			FormDivSelectFieldToField(&(_instance.Fermata), probe.stageOfInterest, formDiv)
		case "Ending":
			FormDivSelectFieldToField(&(_instance.Ending), probe.stageOfInterest, formDiv)
		case "Repeat":
			FormDivSelectFieldToField(&(_instance.Repeat), probe.stageOfInterest, formDiv)
		case "A_measure:Barline":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Barline", func(owner *models.A_measure) *[]*models.Barline { return &owner.Barline })
		case "A_part_1:Barline":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Barline", func(owner *models.A_part_1) *[]*models.Barline { return &owner.Barline })
		}
	}
}

func __gong__New__BarreFormCallback(
	_instance *models.Barre,
	probe *Probe,
	formGroup *form.FormGroup,
) (barreFormCallback *FormCallback[*models.Barre]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBarreFields,
	)
}

type BarreFormCallback = FormCallback[*models.Barre]

func saveBarreFields(
	_instance *models.Barre,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		}
	}
}

func __gong__New__BassFormCallback(
	_instance *models.Bass,
	probe *Probe,
	formGroup *form.FormGroup,
) (bassFormCallback *FormCallback[*models.Bass]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBassFields,
	)
}

type BassFormCallback = FormCallback[*models.Bass]

func saveBassFields(
	_instance *models.Bass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Arrangement":
			FormDivBasicFieldToField(&(_instance.Arrangement), formDiv)
		case "Bass_separator":
			FormDivSelectFieldToField(&(_instance.Bass_separator), probe.stageOfInterest, formDiv)
		case "Bass_step":
			FormDivSelectFieldToField(&(_instance.Bass_step), probe.stageOfInterest, formDiv)
		case "Bass_alter":
			FormDivSelectFieldToField(&(_instance.Bass_alter), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Bass_stepFormCallback(
	_instance *models.Bass_step,
	probe *Probe,
	formGroup *form.FormGroup,
) (bass_stepFormCallback *FormCallback[*models.Bass_step]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBass_stepFields,
	)
}

type Bass_stepFormCallback = FormCallback[*models.Bass_step]

func saveBass_stepFields(
	_instance *models.Bass_step,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__BeamFormCallback(
	_instance *models.Beam,
	probe *Probe,
	formGroup *form.FormGroup,
) (beamFormCallback *FormCallback[*models.Beam]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBeamFields,
	)
}

type BeamFormCallback = FormCallback[*models.Beam]

func saveBeamFields(
	_instance *models.Beam,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Repeater":
			FormDivEnumStringFieldToField(&(_instance.Repeater), formDiv)
		case "Fan":
			FormDivBasicFieldToField(&(_instance.Fan), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Beat_repeatFormCallback(
	_instance *models.Beat_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) (beat_repeatFormCallback *FormCallback[*models.Beat_repeat]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBeat_repeatFields,
	)
}

type Beat_repeatFormCallback = FormCallback[*models.Beat_repeat]

func saveBeat_repeatFields(
	_instance *models.Beat_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&(_instance.Slashes), formDiv)
		case "Use_dots":
			FormDivEnumStringFieldToField(&(_instance.Use_dots), formDiv)
		case "Slash_type":
			FormDivEnumStringFieldToField(&(_instance.Slash_type), formDiv)
		case "Slash_dot":
			FormDivBasicFieldToField(&(_instance.Slash_dot), formDiv)
		case "Except_voice":
			FormDivBasicFieldToField(&(_instance.Except_voice), formDiv)
		}
	}
}

func __gong__New__Beat_unit_tiedFormCallback(
	_instance *models.Beat_unit_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) (beat_unit_tiedFormCallback *FormCallback[*models.Beat_unit_tied]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBeat_unit_tiedFields,
	)
}

type Beat_unit_tiedFormCallback = FormCallback[*models.Beat_unit_tied]

func saveBeat_unit_tiedFields(
	_instance *models.Beat_unit_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Beat_unit":
			FormDivEnumStringFieldToField(&(_instance.Beat_unit), formDiv)
		case "Beat_unit_dot":
			FormDivBasicFieldToField(&(_instance.Beat_unit_dot), formDiv)
		case "Metronome:Beat_unit_tied":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Beat_unit_tied", func(owner *models.Metronome) *[]*models.Beat_unit_tied { return &owner.Beat_unit_tied })
		}
	}
}

func __gong__New__BeaterFormCallback(
	_instance *models.Beater,
	probe *Probe,
	formGroup *form.FormGroup,
) (beaterFormCallback *FormCallback[*models.Beater]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBeaterFields,
	)
}

type BeaterFormCallback = FormCallback[*models.Beater]

func saveBeaterFields(
	_instance *models.Beater,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tip":
			FormDivBasicFieldToField(&(_instance.Tip), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__BendFormCallback(
	_instance *models.Bend,
	probe *Probe,
	formGroup *form.FormGroup,
) (bendFormCallback *FormCallback[*models.Bend]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBendFields,
	)
}

type BendFormCallback = FormCallback[*models.Bend]

func saveBendFields(
	_instance *models.Bend,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Shape":
			FormDivBasicFieldToField(&(_instance.Shape), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Accelerate":
			FormDivEnumStringFieldToField(&(_instance.Accelerate), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "First_beat":
			FormDivBasicFieldToField(&(_instance.First_beat), formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&(_instance.Last_beat), formDiv)
		case "Bend_alter":
			FormDivBasicFieldToField(&(_instance.Bend_alter), formDiv)
		case "Pre_bend":
			FormDivBasicFieldToField(&(_instance.Pre_bend), formDiv)
		case "Release":
			FormDivSelectFieldToField(&(_instance.Release), probe.stageOfInterest, formDiv)
		case "With_bar":
			FormDivSelectFieldToField(&(_instance.With_bar), probe.stageOfInterest, formDiv)
		case "Technical:Bend":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bend", func(owner *models.Technical) *[]*models.Bend { return &owner.Bend })
		}
	}
}

func __gong__New__BookmarkFormCallback(
	_instance *models.Bookmark,
	probe *Probe,
	formGroup *form.FormGroup,
) (bookmarkFormCallback *FormCallback[*models.Bookmark]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBookmarkFields,
	)
}

type BookmarkFormCallback = FormCallback[*models.Bookmark]

func saveBookmarkFields(
	_instance *models.Bookmark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Element":
			FormDivBasicFieldToField(&(_instance.Element), formDiv)
		case "Position":
			FormDivBasicFieldToField(&(_instance.Position), formDiv)
		case "A_measure:Bookmark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bookmark", func(owner *models.A_measure) *[]*models.Bookmark { return &owner.Bookmark })
		case "A_part_1:Bookmark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bookmark", func(owner *models.A_part_1) *[]*models.Bookmark { return &owner.Bookmark })
		case "Credit:Bookmark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bookmark", func(owner *models.Credit) *[]*models.Bookmark { return &owner.Bookmark })
		}
	}
}

func __gong__New__BracketFormCallback(
	_instance *models.Bracket,
	probe *Probe,
	formGroup *form.FormGroup,
) (bracketFormCallback *FormCallback[*models.Bracket]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBracketFields,
	)
}

type BracketFormCallback = FormCallback[*models.Bracket]

func saveBracketFields(
	_instance *models.Bracket,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line_end":
			FormDivBasicFieldToField(&(_instance.Line_end), formDiv)
		case "End_length":
			FormDivBasicFieldToField(&(_instance.End_length), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__Breath_markFormCallback(
	_instance *models.Breath_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) (breath_markFormCallback *FormCallback[*models.Breath_mark]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBreath_markFields,
	)
}

type Breath_markFormCallback = FormCallback[*models.Breath_mark]

func saveBreath_markFields(
	_instance *models.Breath_mark,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Articulations:Breath_mark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Breath_mark", func(owner *models.Articulations) *[]*models.Breath_mark { return &owner.Breath_mark })
		}
	}
}

func __gong__New__CaesuraFormCallback(
	_instance *models.Caesura,
	probe *Probe,
	formGroup *form.FormGroup,
) (caesuraFormCallback *FormCallback[*models.Caesura]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCaesuraFields,
	)
}

type CaesuraFormCallback = FormCallback[*models.Caesura]

func saveCaesuraFields(
	_instance *models.Caesura,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Articulations:Caesura":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Caesura", func(owner *models.Articulations) *[]*models.Caesura { return &owner.Caesura })
		}
	}
}

func __gong__New__CancelFormCallback(
	_instance *models.Cancel,
	probe *Probe,
	formGroup *form.FormGroup,
) (cancelFormCallback *FormCallback[*models.Cancel]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCancelFields,
	)
}

type CancelFormCallback = FormCallback[*models.Cancel]

func saveCancelFields(
	_instance *models.Cancel,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Location":
			FormDivBasicFieldToField(&(_instance.Location), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__ClefFormCallback(
	_instance *models.Clef,
	probe *Probe,
	formGroup *form.FormGroup,
) (clefFormCallback *FormCallback[*models.Clef]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClefFields,
	)
}

type ClefFormCallback = FormCallback[*models.Clef]

func saveClefFields(
	_instance *models.Clef,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Additional":
			FormDivEnumStringFieldToField(&(_instance.Additional), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(_instance.Size), formDiv)
		case "After_barline":
			FormDivEnumStringFieldToField(&(_instance.After_barline), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Sign":
			FormDivBasicFieldToField(&(_instance.Sign), formDiv)
		case "Line":
			FormDivBasicFieldToField(&(_instance.Line), formDiv)
		case "Clef_octave_change":
			FormDivBasicFieldToField(&(_instance.Clef_octave_change), formDiv)
		case "Attributes:Clef":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Clef", func(owner *models.Attributes) *[]*models.Clef { return &owner.Clef })
		}
	}
}

func __gong__New__CodaFormCallback(
	_instance *models.Coda,
	probe *Probe,
	formGroup *form.FormGroup,
) (codaFormCallback *FormCallback[*models.Coda]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCodaFields,
	)
}

type CodaFormCallback = FormCallback[*models.Coda]

func saveCodaFields(
	_instance *models.Coda,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Direction_type:Coda":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Coda", func(owner *models.Direction_type) *[]*models.Coda { return &owner.Coda })
		}
	}
}

func __gong__New__CreditFormCallback(
	_instance *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) (creditFormCallback *FormCallback[*models.Credit]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCreditFields,
	)
}

type CreditFormCallback = FormCallback[*models.Credit]

func saveCreditFields(
	_instance *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Page":
			FormDivBasicFieldToField(&(_instance.Page), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Credit_type":
			FormDivBasicFieldToField(&(_instance.Credit_type), formDiv)
		case "Credit_image":
			FormDivSelectFieldToField(&(_instance.Credit_image), probe.stageOfInterest, formDiv)
		case "Link":
			FormDivSliceOfPointersToField(_instance, "Link", &(_instance.Link), formDiv, probe)
		case "Bookmark":
			FormDivSliceOfPointersToField(_instance, "Bookmark", &(_instance.Bookmark), formDiv, probe)
		case "Credit_words":
			FormDivSliceOfPointersToField(_instance, "Credit_words", &(_instance.Credit_words), formDiv, probe)
		case "Credit_symbol":
			FormDivSliceOfPointersToField(_instance, "Credit_symbol", &(_instance.Credit_symbol), formDiv, probe)
		case "Score_partwise:Credit":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Credit", func(owner *models.Score_partwise) *[]*models.Credit { return &owner.Credit })
		case "Score_timewise:Credit":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Credit", func(owner *models.Score_timewise) *[]*models.Credit { return &owner.Credit })
		}
	}
}

func __gong__New__DashesFormCallback(
	_instance *models.Dashes,
	probe *Probe,
	formGroup *form.FormGroup,
) (dashesFormCallback *FormCallback[*models.Dashes]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDashesFields,
	)
}

type DashesFormCallback = FormCallback[*models.Dashes]

func saveDashesFields(
	_instance *models.Dashes,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__DefaultsFormCallback(
	_instance *models.Defaults,
	probe *Probe,
	formGroup *form.FormGroup,
) (defaultsFormCallback *FormCallback[*models.Defaults]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDefaultsFields,
	)
}

type DefaultsFormCallback = FormCallback[*models.Defaults]

func saveDefaultsFields(
	_instance *models.Defaults,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Scaling":
			FormDivSelectFieldToField(&(_instance.Scaling), probe.stageOfInterest, formDiv)
		case "Concert_score":
			FormDivBasicFieldToField(&(_instance.Concert_score), formDiv)
		case "Page_layout":
			FormDivSelectFieldToField(&(_instance.Page_layout), probe.stageOfInterest, formDiv)
		case "System_layout":
			FormDivSelectFieldToField(&(_instance.System_layout), probe.stageOfInterest, formDiv)
		case "Staff_layout":
			FormDivSliceOfPointersToField(_instance, "Staff_layout", &(_instance.Staff_layout), formDiv, probe)
		case "Appearance":
			FormDivSelectFieldToField(&(_instance.Appearance), probe.stageOfInterest, formDiv)
		case "Music_font":
			FormDivSelectFieldToField(&(_instance.Music_font), probe.stageOfInterest, formDiv)
		case "Word_font":
			FormDivSelectFieldToField(&(_instance.Word_font), probe.stageOfInterest, formDiv)
		case "Lyric_font":
			FormDivSliceOfPointersToField(_instance, "Lyric_font", &(_instance.Lyric_font), formDiv, probe)
		case "Lyric_language":
			FormDivSliceOfPointersToField(_instance, "Lyric_language", &(_instance.Lyric_language), formDiv, probe)
		}
	}
}

func __gong__New__DegreeFormCallback(
	_instance *models.Degree,
	probe *Probe,
	formGroup *form.FormGroup,
) (degreeFormCallback *FormCallback[*models.Degree]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDegreeFields,
	)
}

type DegreeFormCallback = FormCallback[*models.Degree]

func saveDegreeFields(
	_instance *models.Degree,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Degree_value":
			FormDivSelectFieldToField(&(_instance.Degree_value), probe.stageOfInterest, formDiv)
		case "Degree_alter":
			FormDivSelectFieldToField(&(_instance.Degree_alter), probe.stageOfInterest, formDiv)
		case "Degree_type":
			FormDivSelectFieldToField(&(_instance.Degree_type), probe.stageOfInterest, formDiv)
		case "Harmony:Degree":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Degree", func(owner *models.Harmony) *[]*models.Degree { return &owner.Degree })
		}
	}
}

func __gong__New__Degree_alterFormCallback(
	_instance *models.Degree_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) (degree_alterFormCallback *FormCallback[*models.Degree_alter]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDegree_alterFields,
	)
}

type Degree_alterFormCallback = FormCallback[*models.Degree_alter]

func saveDegree_alterFields(
	_instance *models.Degree_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Plus_minus":
			FormDivEnumStringFieldToField(&(_instance.Plus_minus), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Degree_typeFormCallback(
	_instance *models.Degree_type,
	probe *Probe,
	formGroup *form.FormGroup,
) (degree_typeFormCallback *FormCallback[*models.Degree_type]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDegree_typeFields,
	)
}

type Degree_typeFormCallback = FormCallback[*models.Degree_type]

func saveDegree_typeFields(
	_instance *models.Degree_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Degree_valueFormCallback(
	_instance *models.Degree_value,
	probe *Probe,
	formGroup *form.FormGroup,
) (degree_valueFormCallback *FormCallback[*models.Degree_value]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDegree_valueFields,
	)
}

type Degree_valueFormCallback = FormCallback[*models.Degree_value]

func saveDegree_valueFields(
	_instance *models.Degree_value,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Symbol":
			FormDivBasicFieldToField(&(_instance.Symbol), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__DirectionFormCallback(
	_instance *models.Direction,
	probe *Probe,
	formGroup *form.FormGroup,
) (directionFormCallback *FormCallback[*models.Direction]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDirectionFields,
	)
}

type DirectionFormCallback = FormCallback[*models.Direction]

func saveDirectionFields(
	_instance *models.Direction,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Directive":
			FormDivEnumStringFieldToField(&(_instance.Directive), formDiv)
		case "System":
			FormDivEnumStringFieldToField(&(_instance.System), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Direction_type":
			FormDivSliceOfPointersToField(_instance, "Direction_type", &(_instance.Direction_type), formDiv, probe)
		case "Offset":
			FormDivSelectFieldToField(&(_instance.Offset), probe.stageOfInterest, formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Voice":
			FormDivBasicFieldToField(&(_instance.Voice), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&(_instance.Staff), formDiv)
		case "Sound":
			FormDivSelectFieldToField(&(_instance.Sound), probe.stageOfInterest, formDiv)
		case "Listening":
			FormDivSelectFieldToField(&(_instance.Listening), probe.stageOfInterest, formDiv)
		case "A_measure:Direction":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Direction", func(owner *models.A_measure) *[]*models.Direction { return &owner.Direction })
		case "A_part_1:Direction":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Direction", func(owner *models.A_part_1) *[]*models.Direction { return &owner.Direction })
		}
	}
}

func __gong__New__Direction_typeFormCallback(
	_instance *models.Direction_type,
	probe *Probe,
	formGroup *form.FormGroup,
) (direction_typeFormCallback *FormCallback[*models.Direction_type]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDirection_typeFields,
	)
}

type Direction_typeFormCallback = FormCallback[*models.Direction_type]

func saveDirection_typeFields(
	_instance *models.Direction_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Rehearsal":
			FormDivSliceOfPointersToField(_instance, "Rehearsal", &(_instance.Rehearsal), formDiv, probe)
		case "Segno":
			FormDivSliceOfPointersToField(_instance, "Segno", &(_instance.Segno), formDiv, probe)
		case "Coda":
			FormDivSliceOfPointersToField(_instance, "Coda", &(_instance.Coda), formDiv, probe)
		case "Words":
			FormDivSliceOfPointersToField(_instance, "Words", &(_instance.Words), formDiv, probe)
		case "Symbol":
			FormDivSliceOfPointersToField(_instance, "Symbol", &(_instance.Symbol), formDiv, probe)
		case "Wedge":
			FormDivSelectFieldToField(&(_instance.Wedge), probe.stageOfInterest, formDiv)
		case "Dynamics":
			FormDivSliceOfPointersToField(_instance, "Dynamics", &(_instance.Dynamics), formDiv, probe)
		case "Dashes":
			FormDivSelectFieldToField(&(_instance.Dashes), probe.stageOfInterest, formDiv)
		case "Bracket":
			FormDivSelectFieldToField(&(_instance.Bracket), probe.stageOfInterest, formDiv)
		case "Pedal":
			FormDivSelectFieldToField(&(_instance.Pedal), probe.stageOfInterest, formDiv)
		case "Metronome":
			FormDivSelectFieldToField(&(_instance.Metronome), probe.stageOfInterest, formDiv)
		case "Octave_shift":
			FormDivSelectFieldToField(&(_instance.Octave_shift), probe.stageOfInterest, formDiv)
		case "Harp_pedals":
			FormDivSelectFieldToField(&(_instance.Harp_pedals), probe.stageOfInterest, formDiv)
		case "Damp":
			FormDivSelectFieldToField(&(_instance.Damp), probe.stageOfInterest, formDiv)
		case "Damp_all":
			FormDivSelectFieldToField(&(_instance.Damp_all), probe.stageOfInterest, formDiv)
		case "Eyeglasses":
			FormDivSelectFieldToField(&(_instance.Eyeglasses), probe.stageOfInterest, formDiv)
		case "String_mute":
			FormDivSelectFieldToField(&(_instance.String_mute), probe.stageOfInterest, formDiv)
		case "Scordatura":
			FormDivSelectFieldToField(&(_instance.Scordatura), probe.stageOfInterest, formDiv)
		case "Image":
			FormDivSelectFieldToField(&(_instance.Image), probe.stageOfInterest, formDiv)
		case "Principal_voice":
			FormDivSelectFieldToField(&(_instance.Principal_voice), probe.stageOfInterest, formDiv)
		case "Percussion":
			FormDivSliceOfPointersToField(_instance, "Percussion", &(_instance.Percussion), formDiv, probe)
		case "Accordion_registration":
			FormDivSelectFieldToField(&(_instance.Accordion_registration), probe.stageOfInterest, formDiv)
		case "Staff_divide":
			FormDivSelectFieldToField(&(_instance.Staff_divide), probe.stageOfInterest, formDiv)
		case "Other_direction":
			FormDivSelectFieldToField(&(_instance.Other_direction), probe.stageOfInterest, formDiv)
		case "Direction:Direction_type":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Direction_type", func(owner *models.Direction) *[]*models.Direction_type { return &owner.Direction_type })
		}
	}
}

func __gong__New__DistanceFormCallback(
	_instance *models.Distance,
	probe *Probe,
	formGroup *form.FormGroup,
) (distanceFormCallback *FormCallback[*models.Distance]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDistanceFields,
	)
}

type DistanceFormCallback = FormCallback[*models.Distance]

func saveDistanceFields(
	_instance *models.Distance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Appearance:Distance":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Distance", func(owner *models.Appearance) *[]*models.Distance { return &owner.Distance })
		}
	}
}

func __gong__New__DoubleFormCallback(
	_instance *models.Double,
	probe *Probe,
	formGroup *form.FormGroup,
) (doubleFormCallback *FormCallback[*models.Double]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDoubleFields,
	)
}

type DoubleFormCallback = FormCallback[*models.Double]

func saveDoubleFields(
	_instance *models.Double,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Above":
			FormDivEnumStringFieldToField(&(_instance.Above), formDiv)
		}
	}
}

func __gong__New__DynamicsFormCallback(
	_instance *models.Dynamics,
	probe *Probe,
	formGroup *form.FormGroup,
) (dynamicsFormCallback *FormCallback[*models.Dynamics]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDynamicsFields,
	)
}

type DynamicsFormCallback = FormCallback[*models.Dynamics]

func saveDynamicsFields(
	_instance *models.Dynamics,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "P":
			FormDivBasicFieldToField(&(_instance.P), formDiv)
		case "Pp":
			FormDivBasicFieldToField(&(_instance.Pp), formDiv)
		case "Ppp":
			FormDivBasicFieldToField(&(_instance.Ppp), formDiv)
		case "Pppp":
			FormDivBasicFieldToField(&(_instance.Pppp), formDiv)
		case "Ppppp":
			FormDivBasicFieldToField(&(_instance.Ppppp), formDiv)
		case "Pppppp":
			FormDivBasicFieldToField(&(_instance.Pppppp), formDiv)
		case "F":
			FormDivBasicFieldToField(&(_instance.F), formDiv)
		case "Ff":
			FormDivBasicFieldToField(&(_instance.Ff), formDiv)
		case "Fff":
			FormDivBasicFieldToField(&(_instance.Fff), formDiv)
		case "Ffff":
			FormDivBasicFieldToField(&(_instance.Ffff), formDiv)
		case "Fffff":
			FormDivBasicFieldToField(&(_instance.Fffff), formDiv)
		case "Ffffff":
			FormDivBasicFieldToField(&(_instance.Ffffff), formDiv)
		case "Mp":
			FormDivBasicFieldToField(&(_instance.Mp), formDiv)
		case "Mf":
			FormDivBasicFieldToField(&(_instance.Mf), formDiv)
		case "Sf":
			FormDivBasicFieldToField(&(_instance.Sf), formDiv)
		case "Sfp":
			FormDivBasicFieldToField(&(_instance.Sfp), formDiv)
		case "Sfpp":
			FormDivBasicFieldToField(&(_instance.Sfpp), formDiv)
		case "Fp":
			FormDivBasicFieldToField(&(_instance.Fp), formDiv)
		case "Rf":
			FormDivBasicFieldToField(&(_instance.Rf), formDiv)
		case "Rfz":
			FormDivBasicFieldToField(&(_instance.Rfz), formDiv)
		case "Sfz":
			FormDivBasicFieldToField(&(_instance.Sfz), formDiv)
		case "Sffz":
			FormDivBasicFieldToField(&(_instance.Sffz), formDiv)
		case "Fz":
			FormDivBasicFieldToField(&(_instance.Fz), formDiv)
		case "N":
			FormDivBasicFieldToField(&(_instance.N), formDiv)
		case "Pf":
			FormDivBasicFieldToField(&(_instance.Pf), formDiv)
		case "Sfzp":
			FormDivBasicFieldToField(&(_instance.Sfzp), formDiv)
		case "Other_dynamics":
			FormDivSliceOfPointersToField(_instance, "Other_dynamics", &(_instance.Other_dynamics), formDiv, probe)
		case "Direction_type:Dynamics":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Dynamics", func(owner *models.Direction_type) *[]*models.Dynamics { return &owner.Dynamics })
		case "Notations:Dynamics":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Dynamics", func(owner *models.Notations) *[]*models.Dynamics { return &owner.Dynamics })
		}
	}
}

func __gong__New__EffectFormCallback(
	_instance *models.Effect,
	probe *Probe,
	formGroup *form.FormGroup,
) (effectFormCallback *FormCallback[*models.Effect]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEffectFields,
	)
}

type EffectFormCallback = FormCallback[*models.Effect]

func saveEffectFields(
	_instance *models.Effect,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__ElisionFormCallback(
	_instance *models.Elision,
	probe *Probe,
	formGroup *form.FormGroup,
) (elisionFormCallback *FormCallback[*models.Elision]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveElisionFields,
	)
}

type ElisionFormCallback = FormCallback[*models.Elision]

func saveElisionFields(
	_instance *models.Elision,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Lyric:Elision":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elision", func(owner *models.Lyric) *[]*models.Elision { return &owner.Elision })
		}
	}
}

func __gong__New__EmptyFormCallback(
	_instance *models.Empty,
	probe *Probe,
	formGroup *form.FormGroup,
) (emptyFormCallback *FormCallback[*models.Empty]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmptyFields,
	)
}

type EmptyFormCallback = FormCallback[*models.Empty]

func saveEmptyFields(
	_instance *models.Empty,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Empty_fontFormCallback(
	_instance *models.Empty_font,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_fontFormCallback *FormCallback[*models.Empty_font]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_fontFields,
	)
}

type Empty_fontFormCallback = FormCallback[*models.Empty_font]

func saveEmpty_fontFields(
	_instance *models.Empty_font,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		}
	}
}

func __gong__New__Empty_lineFormCallback(
	_instance *models.Empty_line,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_lineFormCallback *FormCallback[*models.Empty_line]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_lineFields,
	)
}

type Empty_lineFormCallback = FormCallback[*models.Empty_line]

func saveEmpty_lineFields(
	_instance *models.Empty_line,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Line_shape":
			FormDivBasicFieldToField(&(_instance.Line_shape), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Line_length":
			FormDivBasicFieldToField(&(_instance.Line_length), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Articulations:Scoop":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Scoop", func(owner *models.Articulations) *[]*models.Empty_line { return &owner.Scoop })
		case "Articulations:Plop":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Plop", func(owner *models.Articulations) *[]*models.Empty_line { return &owner.Plop })
		case "Articulations:Doit":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Doit", func(owner *models.Articulations) *[]*models.Empty_line { return &owner.Doit })
		case "Articulations:Falloff":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Falloff", func(owner *models.Articulations) *[]*models.Empty_line { return &owner.Falloff })
		}
	}
}

func __gong__New__Empty_placementFormCallback(
	_instance *models.Empty_placement,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_placementFormCallback *FormCallback[*models.Empty_placement]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_placementFields,
	)
}

type Empty_placementFormCallback = FormCallback[*models.Empty_placement]

func saveEmpty_placementFields(
	_instance *models.Empty_placement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Articulations:Accent":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Accent", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Accent })
		case "Articulations:Staccato":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staccato", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Staccato })
		case "Articulations:Tenuto":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tenuto", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Tenuto })
		case "Articulations:Detached_legato":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Detached_legato", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Detached_legato })
		case "Articulations:Staccatissimo":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staccatissimo", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Staccatissimo })
		case "Articulations:Spiccato":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Spiccato", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Spiccato })
		case "Articulations:Stress":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Stress", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Stress })
		case "Articulations:Unstress":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Unstress", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Unstress })
		case "Articulations:Soft_accent":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Soft_accent", func(owner *models.Articulations) *[]*models.Empty_placement { return &owner.Soft_accent })
		case "Note:Dot":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Dot", func(owner *models.Note) *[]*models.Empty_placement { return &owner.Dot })
		case "Ornaments:Schleifer":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Schleifer", func(owner *models.Ornaments) *[]*models.Empty_placement { return &owner.Schleifer })
		case "Technical:Up_bow":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Up_bow", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Up_bow })
		case "Technical:Down_bow":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Down_bow", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Down_bow })
		case "Technical:Open_string":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Open_string", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Open_string })
		case "Technical:Thumb_position":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Thumb_position", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Thumb_position })
		case "Technical:Double_tongue":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Double_tongue", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Double_tongue })
		case "Technical:Triple_tongue":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Triple_tongue", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Triple_tongue })
		case "Technical:Snap_pizzicato":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Snap_pizzicato", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Snap_pizzicato })
		case "Technical:Fingernails":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fingernails", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Fingernails })
		case "Technical:Brass_bend":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Brass_bend", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Brass_bend })
		case "Technical:Flip":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Flip", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Flip })
		case "Technical:Smear":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Smear", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Smear })
		case "Technical:Golpe":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Golpe", func(owner *models.Technical) *[]*models.Empty_placement { return &owner.Golpe })
		}
	}
}

func __gong__New__Empty_placement_smuflFormCallback(
	_instance *models.Empty_placement_smufl,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_placement_smuflFormCallback *FormCallback[*models.Empty_placement_smufl]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_placement_smuflFields,
	)
}

type Empty_placement_smuflFormCallback = FormCallback[*models.Empty_placement_smufl]

func saveEmpty_placement_smuflFields(
	_instance *models.Empty_placement_smufl,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Technical:Stopped":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Stopped", func(owner *models.Technical) *[]*models.Empty_placement_smufl { return &owner.Stopped })
		case "Technical:Open":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Open", func(owner *models.Technical) *[]*models.Empty_placement_smufl { return &owner.Open })
		case "Technical:Half_muted":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Half_muted", func(owner *models.Technical) *[]*models.Empty_placement_smufl { return &owner.Half_muted })
		}
	}
}

func __gong__New__Empty_print_object_style_alignFormCallback(
	_instance *models.Empty_print_object_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_print_object_style_alignFormCallback *FormCallback[*models.Empty_print_object_style_align]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_print_object_style_alignFields,
	)
}

type Empty_print_object_style_alignFormCallback = FormCallback[*models.Empty_print_object_style_align]

func saveEmpty_print_object_style_alignFields(
	_instance *models.Empty_print_object_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		}
	}
}

func __gong__New__Empty_print_styleFormCallback(
	_instance *models.Empty_print_style,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_print_styleFormCallback *FormCallback[*models.Empty_print_style]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_print_styleFields,
	)
}

type Empty_print_styleFormCallback = FormCallback[*models.Empty_print_style]

func saveEmpty_print_styleFields(
	_instance *models.Empty_print_style,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		}
	}
}

func __gong__New__Empty_print_style_alignFormCallback(
	_instance *models.Empty_print_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_print_style_alignFormCallback *FormCallback[*models.Empty_print_style_align]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_print_style_alignFields,
	)
}

type Empty_print_style_alignFormCallback = FormCallback[*models.Empty_print_style_align]

func saveEmpty_print_style_alignFields(
	_instance *models.Empty_print_style_align,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		}
	}
}

func __gong__New__Empty_print_style_align_idFormCallback(
	_instance *models.Empty_print_style_align_id,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_print_style_align_idFormCallback *FormCallback[*models.Empty_print_style_align_id]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_print_style_align_idFields,
	)
}

type Empty_print_style_align_idFormCallback = FormCallback[*models.Empty_print_style_align_id]

func saveEmpty_print_style_align_idFields(
	_instance *models.Empty_print_style_align_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__Empty_trill_soundFormCallback(
	_instance *models.Empty_trill_sound,
	probe *Probe,
	formGroup *form.FormGroup,
) (empty_trill_soundFormCallback *FormCallback[*models.Empty_trill_sound]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmpty_trill_soundFields,
	)
}

type Empty_trill_soundFormCallback = FormCallback[*models.Empty_trill_sound]

func saveEmpty_trill_soundFields(
	_instance *models.Empty_trill_sound,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Start_note":
			FormDivBasicFieldToField(&(_instance.Start_note), formDiv)
		case "Trill_step":
			FormDivBasicFieldToField(&(_instance.Trill_step), formDiv)
		case "Two_note_turn":
			FormDivBasicFieldToField(&(_instance.Two_note_turn), formDiv)
		case "Accelerate":
			FormDivEnumStringFieldToField(&(_instance.Accelerate), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&(_instance.Second_beat), formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&(_instance.Last_beat), formDiv)
		case "Ornaments:Trill_mark":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Trill_mark", func(owner *models.Ornaments) *[]*models.Empty_trill_sound { return &owner.Trill_mark })
		case "Ornaments:Vertical_turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Vertical_turn", func(owner *models.Ornaments) *[]*models.Empty_trill_sound { return &owner.Vertical_turn })
		case "Ornaments:Inverted_vertical_turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Inverted_vertical_turn", func(owner *models.Ornaments) *[]*models.Empty_trill_sound { return &owner.Inverted_vertical_turn })
		case "Ornaments:Shake":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Shake", func(owner *models.Ornaments) *[]*models.Empty_trill_sound { return &owner.Shake })
		case "Ornaments:Haydn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Haydn", func(owner *models.Ornaments) *[]*models.Empty_trill_sound { return &owner.Haydn })
		}
	}
}

func __gong__New__EncodingFormCallback(
	_instance *models.Encoding,
	probe *Probe,
	formGroup *form.FormGroup,
) (encodingFormCallback *FormCallback[*models.Encoding]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEncodingFields,
	)
}

type EncodingFormCallback = FormCallback[*models.Encoding]

func saveEncodingFields(
	_instance *models.Encoding,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Encoder":
			FormDivSliceOfPointersToField(_instance, "Encoder", &(_instance.Encoder), formDiv, probe)
		case "Software":
			FormDivBasicFieldToField(&(_instance.Software), formDiv)
		case "Encoding_description":
			FormDivBasicFieldToField(&(_instance.Encoding_description), formDiv)
		case "Supports":
			FormDivSliceOfPointersToField(_instance, "Supports", &(_instance.Supports), formDiv, probe)
		}
	}
}

func __gong__New__EndingFormCallback(
	_instance *models.Ending,
	probe *Probe,
	formGroup *form.FormGroup,
) (endingFormCallback *FormCallback[*models.Ending]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEndingFields,
	)
}

type EndingFormCallback = FormCallback[*models.Ending]

func saveEndingFields(
	_instance *models.Ending,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "End_length":
			FormDivBasicFieldToField(&(_instance.End_length), formDiv)
		case "Text_x":
			FormDivBasicFieldToField(&(_instance.Text_x), formDiv)
		case "Text_y":
			FormDivBasicFieldToField(&(_instance.Text_y), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "System":
			FormDivEnumStringFieldToField(&(_instance.System), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__ExtendFormCallback(
	_instance *models.Extend,
	probe *Probe,
	formGroup *form.FormGroup,
) (extendFormCallback *FormCallback[*models.Extend]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExtendFields,
	)
}

type ExtendFormCallback = FormCallback[*models.Extend]

func saveExtendFields(
	_instance *models.Extend,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		}
	}
}

func __gong__New__FeatureFormCallback(
	_instance *models.Feature,
	probe *Probe,
	formGroup *form.FormGroup,
) (featureFormCallback *FormCallback[*models.Feature]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFeatureFields,
	)
}

type FeatureFormCallback = FormCallback[*models.Feature]

func saveFeatureFields(
	_instance *models.Feature,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Grouping:Feature":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Feature", func(owner *models.Grouping) *[]*models.Feature { return &owner.Feature })
		}
	}
}

func __gong__New__FermataFormCallback(
	_instance *models.Fermata,
	probe *Probe,
	formGroup *form.FormGroup,
) (fermataFormCallback *FormCallback[*models.Fermata]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFermataFields,
	)
}

type FermataFormCallback = FormCallback[*models.Fermata]

func saveFermataFields(
	_instance *models.Fermata,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Notations:Fermata":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fermata", func(owner *models.Notations) *[]*models.Fermata { return &owner.Fermata })
		}
	}
}

func __gong__New__FigureFormCallback(
	_instance *models.Figure,
	probe *Probe,
	formGroup *form.FormGroup,
) (figureFormCallback *FormCallback[*models.Figure]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFigureFields,
	)
}

type FigureFormCallback = FormCallback[*models.Figure]

func saveFigureFields(
	_instance *models.Figure,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Prefix":
			FormDivSelectFieldToField(&(_instance.Prefix), probe.stageOfInterest, formDiv)
		case "Figure_number":
			FormDivSelectFieldToField(&(_instance.Figure_number), probe.stageOfInterest, formDiv)
		case "Suffix":
			FormDivSelectFieldToField(&(_instance.Suffix), probe.stageOfInterest, formDiv)
		case "Extend":
			FormDivSelectFieldToField(&(_instance.Extend), probe.stageOfInterest, formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Figured_bass:Figure":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Figure", func(owner *models.Figured_bass) *[]*models.Figure { return &owner.Figure })
		}
	}
}

func __gong__New__Figured_bassFormCallback(
	_instance *models.Figured_bass,
	probe *Probe,
	formGroup *form.FormGroup,
) (figured_bassFormCallback *FormCallback[*models.Figured_bass]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFigured_bassFields,
	)
}

type Figured_bassFormCallback = FormCallback[*models.Figured_bass]

func saveFigured_bassFields(
	_instance *models.Figured_bass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Print_dot":
			FormDivEnumStringFieldToField(&(_instance.Print_dot), formDiv)
		case "Print_lyric":
			FormDivEnumStringFieldToField(&(_instance.Print_lyric), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Print_spacing":
			FormDivEnumStringFieldToField(&(_instance.Print_spacing), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Figure":
			FormDivSliceOfPointersToField(_instance, "Figure", &(_instance.Figure), formDiv, probe)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "A_measure:Figured_bass":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Figured_bass", func(owner *models.A_measure) *[]*models.Figured_bass { return &owner.Figured_bass })
		case "A_part_1:Figured_bass":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Figured_bass", func(owner *models.A_part_1) *[]*models.Figured_bass { return &owner.Figured_bass })
		}
	}
}

func __gong__New__FingeringFormCallback(
	_instance *models.Fingering,
	probe *Probe,
	formGroup *form.FormGroup,
) (fingeringFormCallback *FormCallback[*models.Fingering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFingeringFields,
	)
}

type FingeringFormCallback = FormCallback[*models.Fingering]

func saveFingeringFields(
	_instance *models.Fingering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Substitution":
			FormDivEnumStringFieldToField(&(_instance.Substitution), formDiv)
		case "Alternate":
			FormDivEnumStringFieldToField(&(_instance.Alternate), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Fingering":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fingering", func(owner *models.Technical) *[]*models.Fingering { return &owner.Fingering })
		}
	}
}

func __gong__New__First_fretFormCallback(
	_instance *models.First_fret,
	probe *Probe,
	formGroup *form.FormGroup,
) (first_fretFormCallback *FormCallback[*models.First_fret]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFirst_fretFields,
	)
}

type First_fretFormCallback = FormCallback[*models.First_fret]

func saveFirst_fretFields(
	_instance *models.First_fret,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Location":
			FormDivBasicFieldToField(&(_instance.Location), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__For_partFormCallback(
	_instance *models.For_part,
	probe *Probe,
	formGroup *form.FormGroup,
) (for_partFormCallback *FormCallback[*models.For_part]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFor_partFields,
	)
}

type For_partFormCallback = FormCallback[*models.For_part]

func saveFor_partFields(
	_instance *models.For_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Part_clef":
			FormDivSelectFieldToField(&(_instance.Part_clef), probe.stageOfInterest, formDiv)
		case "Part_transpose":
			FormDivSelectFieldToField(&(_instance.Part_transpose), probe.stageOfInterest, formDiv)
		case "Attributes:For_part":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "For_part", func(owner *models.Attributes) *[]*models.For_part { return &owner.For_part })
		}
	}
}

func __gong__New__Formatted_symbolFormCallback(
	_instance *models.Formatted_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) (formatted_symbolFormCallback *FormCallback[*models.Formatted_symbol]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormatted_symbolFields,
	)
}

type Formatted_symbolFormCallback = FormCallback[*models.Formatted_symbol]

func saveFormatted_symbolFields(
	_instance *models.Formatted_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&(_instance.Line_height), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Formatted_symbol_idFormCallback(
	_instance *models.Formatted_symbol_id,
	probe *Probe,
	formGroup *form.FormGroup,
) (formatted_symbol_idFormCallback *FormCallback[*models.Formatted_symbol_id]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormatted_symbol_idFields,
	)
}

type Formatted_symbol_idFormCallback = FormCallback[*models.Formatted_symbol_id]

func saveFormatted_symbol_idFields(
	_instance *models.Formatted_symbol_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&(_instance.Line_height), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Credit:Credit_symbol":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Credit_symbol", func(owner *models.Credit) *[]*models.Formatted_symbol_id { return &owner.Credit_symbol })
		case "Direction_type:Symbol":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Symbol", func(owner *models.Direction_type) *[]*models.Formatted_symbol_id { return &owner.Symbol })
		}
	}
}

func __gong__New__Formatted_textFormCallback(
	_instance *models.Formatted_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (formatted_textFormCallback *FormCallback[*models.Formatted_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormatted_textFields,
	)
}

type Formatted_textFormCallback = FormCallback[*models.Formatted_text]

func saveFormatted_textFields(
	_instance *models.Formatted_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Space":
			FormDivBasicFieldToField(&(_instance.Space), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&(_instance.Line_height), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Name_display:Display_text":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Display_text", func(owner *models.Name_display) *[]*models.Formatted_text { return &owner.Display_text })
		case "Notehead_text:Display_text":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Display_text", func(owner *models.Notehead_text) *[]*models.Formatted_text { return &owner.Display_text })
		}
	}
}

func __gong__New__Formatted_text_idFormCallback(
	_instance *models.Formatted_text_id,
	probe *Probe,
	formGroup *form.FormGroup,
) (formatted_text_idFormCallback *FormCallback[*models.Formatted_text_id]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFormatted_text_idFields,
	)
}

type Formatted_text_idFormCallback = FormCallback[*models.Formatted_text_id]

func saveFormatted_text_idFields(
	_instance *models.Formatted_text_id,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Space":
			FormDivBasicFieldToField(&(_instance.Space), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Line_height":
			FormDivBasicFieldToField(&(_instance.Line_height), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Credit:Credit_words":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Credit_words", func(owner *models.Credit) *[]*models.Formatted_text_id { return &owner.Credit_words })
		case "Direction_type:Rehearsal":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Rehearsal", func(owner *models.Direction_type) *[]*models.Formatted_text_id { return &owner.Rehearsal })
		case "Direction_type:Words":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Words", func(owner *models.Direction_type) *[]*models.Formatted_text_id { return &owner.Words })
		}
	}
}

func __gong__New__ForwardFormCallback(
	_instance *models.Forward,
	probe *Probe,
	formGroup *form.FormGroup,
) (forwardFormCallback *FormCallback[*models.Forward]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveForwardFields,
	)
}

type ForwardFormCallback = FormCallback[*models.Forward]

func saveForwardFields(
	_instance *models.Forward,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Voice":
			FormDivBasicFieldToField(&(_instance.Voice), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&(_instance.Staff), formDiv)
		case "A_measure:Forward":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Forward", func(owner *models.A_measure) *[]*models.Forward { return &owner.Forward })
		case "A_part_1:Forward":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Forward", func(owner *models.A_part_1) *[]*models.Forward { return &owner.Forward })
		}
	}
}

func __gong__New__FrameFormCallback(
	_instance *models.Frame,
	probe *Probe,
	formGroup *form.FormGroup,
) (frameFormCallback *FormCallback[*models.Frame]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFrameFields,
	)
}

type FrameFormCallback = FormCallback[*models.Frame]

func saveFrameFields(
	_instance *models.Frame,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Unplayed":
			FormDivBasicFieldToField(&(_instance.Unplayed), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Frame_strings":
			FormDivBasicFieldToField(&(_instance.Frame_strings), formDiv)
		case "Frame_frets":
			FormDivBasicFieldToField(&(_instance.Frame_frets), formDiv)
		case "First_fret":
			FormDivSelectFieldToField(&(_instance.First_fret), probe.stageOfInterest, formDiv)
		case "Frame_note":
			FormDivSliceOfPointersToField(_instance, "Frame_note", &(_instance.Frame_note), formDiv, probe)
		}
	}
}

func __gong__New__Frame_noteFormCallback(
	_instance *models.Frame_note,
	probe *Probe,
	formGroup *form.FormGroup,
) (frame_noteFormCallback *FormCallback[*models.Frame_note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFrame_noteFields,
	)
}

type Frame_noteFormCallback = FormCallback[*models.Frame_note]

func saveFrame_noteFields(
	_instance *models.Frame_note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "String":
			FormDivSelectFieldToField(&(_instance.String), probe.stageOfInterest, formDiv)
		case "Fret":
			FormDivSelectFieldToField(&(_instance.Fret), probe.stageOfInterest, formDiv)
		case "Fingering":
			FormDivSelectFieldToField(&(_instance.Fingering), probe.stageOfInterest, formDiv)
		case "Barre":
			FormDivSelectFieldToField(&(_instance.Barre), probe.stageOfInterest, formDiv)
		case "Frame:Frame_note":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Frame_note", func(owner *models.Frame) *[]*models.Frame_note { return &owner.Frame_note })
		}
	}
}

func __gong__New__FretFormCallback(
	_instance *models.Fret,
	probe *Probe,
	formGroup *form.FormGroup,
) (fretFormCallback *FormCallback[*models.Fret]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFretFields,
	)
}

type FretFormCallback = FormCallback[*models.Fret]

func saveFretFields(
	_instance *models.Fret,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Fret":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fret", func(owner *models.Technical) *[]*models.Fret { return &owner.Fret })
		}
	}
}

func __gong__New__GlassFormCallback(
	_instance *models.Glass,
	probe *Probe,
	formGroup *form.FormGroup,
) (glassFormCallback *FormCallback[*models.Glass]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGlassFields,
	)
}

type GlassFormCallback = FormCallback[*models.Glass]

func saveGlassFields(
	_instance *models.Glass,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__GlissandoFormCallback(
	_instance *models.Glissando,
	probe *Probe,
	formGroup *form.FormGroup,
) (glissandoFormCallback *FormCallback[*models.Glissando]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGlissandoFields,
	)
}

type GlissandoFormCallback = FormCallback[*models.Glissando]

func saveGlissandoFields(
	_instance *models.Glissando,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Notations:Glissando":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Glissando", func(owner *models.Notations) *[]*models.Glissando { return &owner.Glissando })
		}
	}
}

func __gong__New__GlyphFormCallback(
	_instance *models.Glyph,
	probe *Probe,
	formGroup *form.FormGroup,
) (glyphFormCallback *FormCallback[*models.Glyph]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGlyphFields,
	)
}

type GlyphFormCallback = FormCallback[*models.Glyph]

func saveGlyphFields(
	_instance *models.Glyph,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Appearance:Glyph":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Glyph", func(owner *models.Appearance) *[]*models.Glyph { return &owner.Glyph })
		}
	}
}

func __gong__New__GraceFormCallback(
	_instance *models.Grace,
	probe *Probe,
	formGroup *form.FormGroup,
) (graceFormCallback *FormCallback[*models.Grace]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGraceFields,
	)
}

type GraceFormCallback = FormCallback[*models.Grace]

func saveGraceFields(
	_instance *models.Grace,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Steal_time_previous":
			FormDivBasicFieldToField(&(_instance.Steal_time_previous), formDiv)
		case "Steal_time_following":
			FormDivBasicFieldToField(&(_instance.Steal_time_following), formDiv)
		case "Make_time":
			FormDivBasicFieldToField(&(_instance.Make_time), formDiv)
		case "Slash":
			FormDivEnumStringFieldToField(&(_instance.Slash), formDiv)
		}
	}
}

func __gong__New__Group_barlineFormCallback(
	_instance *models.Group_barline,
	probe *Probe,
	formGroup *form.FormGroup,
) (group_barlineFormCallback *FormCallback[*models.Group_barline]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroup_barlineFields,
	)
}

type Group_barlineFormCallback = FormCallback[*models.Group_barline]

func saveGroup_barlineFields(
	_instance *models.Group_barline,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Group_nameFormCallback(
	_instance *models.Group_name,
	probe *Probe,
	formGroup *form.FormGroup,
) (group_nameFormCallback *FormCallback[*models.Group_name]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroup_nameFields,
	)
}

type Group_nameFormCallback = FormCallback[*models.Group_name]

func saveGroup_nameFields(
	_instance *models.Group_name,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Group_symbolFormCallback(
	_instance *models.Group_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) (group_symbolFormCallback *FormCallback[*models.Group_symbol]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroup_symbolFields,
	)
}

type Group_symbolFormCallback = FormCallback[*models.Group_symbol]

func saveGroup_symbolFields(
	_instance *models.Group_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__GroupingFormCallback(
	_instance *models.Grouping,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupingFormCallback *FormCallback[*models.Grouping]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupingFields,
	)
}

type GroupingFormCallback = FormCallback[*models.Grouping]

func saveGroupingFields(
	_instance *models.Grouping,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Member_of":
			FormDivBasicFieldToField(&(_instance.Member_of), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Feature":
			FormDivSliceOfPointersToField(_instance, "Feature", &(_instance.Feature), formDiv, probe)
		case "A_measure:Grouping":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Grouping", func(owner *models.A_measure) *[]*models.Grouping { return &owner.Grouping })
		case "A_part_1:Grouping":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Grouping", func(owner *models.A_part_1) *[]*models.Grouping { return &owner.Grouping })
		}
	}
}

func __gong__New__Hammer_on_pull_offFormCallback(
	_instance *models.Hammer_on_pull_off,
	probe *Probe,
	formGroup *form.FormGroup,
) (hammer_on_pull_offFormCallback *FormCallback[*models.Hammer_on_pull_off]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHammer_on_pull_offFields,
	)
}

type Hammer_on_pull_offFormCallback = FormCallback[*models.Hammer_on_pull_off]

func saveHammer_on_pull_offFields(
	_instance *models.Hammer_on_pull_off,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Hammer_on":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Hammer_on", func(owner *models.Technical) *[]*models.Hammer_on_pull_off { return &owner.Hammer_on })
		case "Technical:Pull_off":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Pull_off", func(owner *models.Technical) *[]*models.Hammer_on_pull_off { return &owner.Pull_off })
		}
	}
}

func __gong__New__HandbellFormCallback(
	_instance *models.Handbell,
	probe *Probe,
	formGroup *form.FormGroup,
) (handbellFormCallback *FormCallback[*models.Handbell]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHandbellFields,
	)
}

type HandbellFormCallback = FormCallback[*models.Handbell]

func saveHandbellFields(
	_instance *models.Handbell,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Handbell":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Handbell", func(owner *models.Technical) *[]*models.Handbell { return &owner.Handbell })
		}
	}
}

func __gong__New__Harmon_closedFormCallback(
	_instance *models.Harmon_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) (harmon_closedFormCallback *FormCallback[*models.Harmon_closed]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarmon_closedFields,
	)
}

type Harmon_closedFormCallback = FormCallback[*models.Harmon_closed]

func saveHarmon_closedFields(
	_instance *models.Harmon_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Location":
			FormDivBasicFieldToField(&(_instance.Location), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Harmon_muteFormCallback(
	_instance *models.Harmon_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) (harmon_muteFormCallback *FormCallback[*models.Harmon_mute]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarmon_muteFields,
	)
}

type Harmon_muteFormCallback = FormCallback[*models.Harmon_mute]

func saveHarmon_muteFields(
	_instance *models.Harmon_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Harmon_closed":
			FormDivSelectFieldToField(&(_instance.Harmon_closed), probe.stageOfInterest, formDiv)
		case "Technical:Harmon_mute":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Harmon_mute", func(owner *models.Technical) *[]*models.Harmon_mute { return &owner.Harmon_mute })
		}
	}
}

func __gong__New__HarmonicFormCallback(
	_instance *models.Harmonic,
	probe *Probe,
	formGroup *form.FormGroup,
) (harmonicFormCallback *FormCallback[*models.Harmonic]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarmonicFields,
	)
}

type HarmonicFormCallback = FormCallback[*models.Harmonic]

func saveHarmonicFields(
	_instance *models.Harmonic,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Natural":
			FormDivBasicFieldToField(&(_instance.Natural), formDiv)
		case "Artificial":
			FormDivBasicFieldToField(&(_instance.Artificial), formDiv)
		case "Base_pitch":
			FormDivBasicFieldToField(&(_instance.Base_pitch), formDiv)
		case "Touching_pitch":
			FormDivBasicFieldToField(&(_instance.Touching_pitch), formDiv)
		case "Sounding_pitch":
			FormDivBasicFieldToField(&(_instance.Sounding_pitch), formDiv)
		case "Technical:Harmonic":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Harmonic", func(owner *models.Technical) *[]*models.Harmonic { return &owner.Harmonic })
		}
	}
}

func __gong__New__HarmonyFormCallback(
	_instance *models.Harmony,
	probe *Probe,
	formGroup *form.FormGroup,
) (harmonyFormCallback *FormCallback[*models.Harmony]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarmonyFields,
	)
}

type HarmonyFormCallback = FormCallback[*models.Harmony]

func saveHarmonyFields(
	_instance *models.Harmony,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Print_frame":
			FormDivEnumStringFieldToField(&(_instance.Print_frame), formDiv)
		case "Arrangement":
			FormDivEnumStringFieldToField(&(_instance.Arrangement), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "System":
			FormDivEnumStringFieldToField(&(_instance.System), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Root":
			FormDivSelectFieldToField(&(_instance.Root), probe.stageOfInterest, formDiv)
		case "Numeral":
			FormDivSelectFieldToField(&(_instance.Numeral), probe.stageOfInterest, formDiv)
		case "Function":
			FormDivSelectFieldToField(&(_instance.Function), probe.stageOfInterest, formDiv)
		case "Kind":
			FormDivSelectFieldToField(&(_instance.Kind), probe.stageOfInterest, formDiv)
		case "Inversion":
			FormDivSelectFieldToField(&(_instance.Inversion), probe.stageOfInterest, formDiv)
		case "Bass":
			FormDivSelectFieldToField(&(_instance.Bass), probe.stageOfInterest, formDiv)
		case "Degree":
			FormDivSliceOfPointersToField(_instance, "Degree", &(_instance.Degree), formDiv, probe)
		case "Frame":
			FormDivSelectFieldToField(&(_instance.Frame), probe.stageOfInterest, formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(_instance.Offset), probe.stageOfInterest, formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&(_instance.Staff), formDiv)
		case "A_measure:Harmony":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Harmony", func(owner *models.A_measure) *[]*models.Harmony { return &owner.Harmony })
		case "A_part_1:Harmony":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Harmony", func(owner *models.A_part_1) *[]*models.Harmony { return &owner.Harmony })
		}
	}
}

func __gong__New__Harmony_alterFormCallback(
	_instance *models.Harmony_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) (harmony_alterFormCallback *FormCallback[*models.Harmony_alter]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarmony_alterFields,
	)
}

type Harmony_alterFormCallback = FormCallback[*models.Harmony_alter]

func saveHarmony_alterFields(
	_instance *models.Harmony_alter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Location":
			FormDivEnumStringFieldToField(&(_instance.Location), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Harp_pedalsFormCallback(
	_instance *models.Harp_pedals,
	probe *Probe,
	formGroup *form.FormGroup,
) (harp_pedalsFormCallback *FormCallback[*models.Harp_pedals]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHarp_pedalsFields,
	)
}

type Harp_pedalsFormCallback = FormCallback[*models.Harp_pedals]

func saveHarp_pedalsFields(
	_instance *models.Harp_pedals,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Pedal_tuning":
			FormDivSliceOfPointersToField(_instance, "Pedal_tuning", &(_instance.Pedal_tuning), formDiv, probe)
		}
	}
}

func __gong__New__Heel_toeFormCallback(
	_instance *models.Heel_toe,
	probe *Probe,
	formGroup *form.FormGroup,
) (heel_toeFormCallback *FormCallback[*models.Heel_toe]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHeel_toeFields,
	)
}

type Heel_toeFormCallback = FormCallback[*models.Heel_toe]

func saveHeel_toeFields(
	_instance *models.Heel_toe,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Technical:Heel":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Heel", func(owner *models.Technical) *[]*models.Heel_toe { return &owner.Heel })
		case "Technical:Toe":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Toe", func(owner *models.Technical) *[]*models.Heel_toe { return &owner.Toe })
		}
	}
}

func __gong__New__HoleFormCallback(
	_instance *models.Hole,
	probe *Probe,
	formGroup *form.FormGroup,
) (holeFormCallback *FormCallback[*models.Hole]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHoleFields,
	)
}

type HoleFormCallback = FormCallback[*models.Hole]

func saveHoleFields(
	_instance *models.Hole,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Hole_type":
			FormDivBasicFieldToField(&(_instance.Hole_type), formDiv)
		case "Hole_closed":
			FormDivSelectFieldToField(&(_instance.Hole_closed), probe.stageOfInterest, formDiv)
		case "Hole_shape":
			FormDivBasicFieldToField(&(_instance.Hole_shape), formDiv)
		case "Technical:Hole":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Hole", func(owner *models.Technical) *[]*models.Hole { return &owner.Hole })
		}
	}
}

func __gong__New__Hole_closedFormCallback(
	_instance *models.Hole_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) (hole_closedFormCallback *FormCallback[*models.Hole_closed]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHole_closedFields,
	)
}

type Hole_closedFormCallback = FormCallback[*models.Hole_closed]

func saveHole_closedFields(
	_instance *models.Hole_closed,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Location":
			FormDivBasicFieldToField(&(_instance.Location), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Horizontal_turnFormCallback(
	_instance *models.Horizontal_turn,
	probe *Probe,
	formGroup *form.FormGroup,
) (horizontal_turnFormCallback *FormCallback[*models.Horizontal_turn]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveHorizontal_turnFields,
	)
}

type Horizontal_turnFormCallback = FormCallback[*models.Horizontal_turn]

func saveHorizontal_turnFields(
	_instance *models.Horizontal_turn,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Slash":
			FormDivEnumStringFieldToField(&(_instance.Slash), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Start_note":
			FormDivBasicFieldToField(&(_instance.Start_note), formDiv)
		case "Trill_step":
			FormDivBasicFieldToField(&(_instance.Trill_step), formDiv)
		case "Two_note_turn":
			FormDivBasicFieldToField(&(_instance.Two_note_turn), formDiv)
		case "Accelerate":
			FormDivEnumStringFieldToField(&(_instance.Accelerate), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&(_instance.Second_beat), formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&(_instance.Last_beat), formDiv)
		case "Ornaments:Turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Turn", func(owner *models.Ornaments) *[]*models.Horizontal_turn { return &owner.Turn })
		case "Ornaments:Delayed_turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Delayed_turn", func(owner *models.Ornaments) *[]*models.Horizontal_turn { return &owner.Delayed_turn })
		case "Ornaments:Inverted_turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Inverted_turn", func(owner *models.Ornaments) *[]*models.Horizontal_turn { return &owner.Inverted_turn })
		case "Ornaments:Delayed_inverted_turn":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Delayed_inverted_turn", func(owner *models.Ornaments) *[]*models.Horizontal_turn { return &owner.Delayed_inverted_turn })
		}
	}
}

func __gong__New__IdentificationFormCallback(
	_instance *models.Identification,
	probe *Probe,
	formGroup *form.FormGroup,
) (identificationFormCallback *FormCallback[*models.Identification]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveIdentificationFields,
	)
}

type IdentificationFormCallback = FormCallback[*models.Identification]

func saveIdentificationFields(
	_instance *models.Identification,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Creator":
			FormDivSliceOfPointersToField(_instance, "Creator", &(_instance.Creator), formDiv, probe)
		case "Rights":
			FormDivSliceOfPointersToField(_instance, "Rights", &(_instance.Rights), formDiv, probe)
		case "Encoding":
			FormDivSelectFieldToField(&(_instance.Encoding), probe.stageOfInterest, formDiv)
		case "Source":
			FormDivBasicFieldToField(&(_instance.Source), formDiv)
		case "Relation":
			FormDivSliceOfPointersToField(_instance, "Relation", &(_instance.Relation), formDiv, probe)
		case "Miscellaneous":
			FormDivSelectFieldToField(&(_instance.Miscellaneous), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__ImageFormCallback(
	_instance *models.Image,
	probe *Probe,
	formGroup *form.FormGroup,
) (imageFormCallback *FormCallback[*models.Image]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveImageFields,
	)
}

type ImageFormCallback = FormCallback[*models.Image]

func saveImageFields(
	_instance *models.Image,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(_instance.Source), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__InstrumentFormCallback(
	_instance *models.Instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) (instrumentFormCallback *FormCallback[*models.Instrument]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInstrumentFields,
	)
}

type InstrumentFormCallback = FormCallback[*models.Instrument]

func saveInstrumentFields(
	_instance *models.Instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Note:Instrument":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Instrument", func(owner *models.Note) *[]*models.Instrument { return &owner.Instrument })
		}
	}
}

func __gong__New__Instrument_changeFormCallback(
	_instance *models.Instrument_change,
	probe *Probe,
	formGroup *form.FormGroup,
) (instrument_changeFormCallback *FormCallback[*models.Instrument_change]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInstrument_changeFields,
	)
}

type Instrument_changeFormCallback = FormCallback[*models.Instrument_change]

func saveInstrument_changeFields(
	_instance *models.Instrument_change,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Instrument_sound":
			FormDivBasicFieldToField(&(_instance.Instrument_sound), formDiv)
		case "Solo":
			FormDivBasicFieldToField(&(_instance.Solo), formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&(_instance.Ensemble), formDiv)
		case "Virtual_instrument":
			FormDivSelectFieldToField(&(_instance.Virtual_instrument), probe.stageOfInterest, formDiv)
		case "Sound:Instrument_change":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Instrument_change", func(owner *models.Sound) *[]*models.Instrument_change { return &owner.Instrument_change })
		}
	}
}

func __gong__New__Instrument_linkFormCallback(
	_instance *models.Instrument_link,
	probe *Probe,
	formGroup *form.FormGroup,
) (instrument_linkFormCallback *FormCallback[*models.Instrument_link]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInstrument_linkFields,
	)
}

type Instrument_linkFormCallback = FormCallback[*models.Instrument_link]

func saveInstrument_linkFields(
	_instance *models.Instrument_link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Part_link:Instrument_link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Instrument_link", func(owner *models.Part_link) *[]*models.Instrument_link { return &owner.Instrument_link })
		}
	}
}

func __gong__New__InterchangeableFormCallback(
	_instance *models.Interchangeable,
	probe *Probe,
	formGroup *form.FormGroup,
) (interchangeableFormCallback *FormCallback[*models.Interchangeable]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInterchangeableFields,
	)
}

type InterchangeableFormCallback = FormCallback[*models.Interchangeable]

func saveInterchangeableFields(
	_instance *models.Interchangeable,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Symbol":
			FormDivBasicFieldToField(&(_instance.Symbol), formDiv)
		case "Separator":
			FormDivBasicFieldToField(&(_instance.Separator), formDiv)
		case "Time_relation":
			FormDivBasicFieldToField(&(_instance.Time_relation), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "Beat_type":
			FormDivBasicFieldToField(&(_instance.Beat_type), formDiv)
		}
	}
}

func __gong__New__InversionFormCallback(
	_instance *models.Inversion,
	probe *Probe,
	formGroup *form.FormGroup,
) (inversionFormCallback *FormCallback[*models.Inversion]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInversionFields,
	)
}

type InversionFormCallback = FormCallback[*models.Inversion]

func saveInversionFields(
	_instance *models.Inversion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__KeyFormCallback(
	_instance *models.Key,
	probe *Probe,
	formGroup *form.FormGroup,
) (keyFormCallback *FormCallback[*models.Key]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKeyFields,
	)
}

type KeyFormCallback = FormCallback[*models.Key]

func saveKeyFields(
	_instance *models.Key,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Cancel":
			FormDivSelectFieldToField(&(_instance.Cancel), probe.stageOfInterest, formDiv)
		case "Fifths":
			FormDivBasicFieldToField(&(_instance.Fifths), formDiv)
		case "Mode":
			FormDivBasicFieldToField(&(_instance.Mode), formDiv)
		case "Key_step":
			FormDivEnumStringFieldToField(&(_instance.Key_step), formDiv)
		case "Key_alter":
			FormDivBasicFieldToField(&(_instance.Key_alter), formDiv)
		case "Key_accidental":
			FormDivSelectFieldToField(&(_instance.Key_accidental), probe.stageOfInterest, formDiv)
		case "Key_octave":
			FormDivSliceOfPointersToField(_instance, "Key_octave", &(_instance.Key_octave), formDiv, probe)
		case "Attributes:Key":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Key", func(owner *models.Attributes) *[]*models.Key { return &owner.Key })
		}
	}
}

func __gong__New__Key_accidentalFormCallback(
	_instance *models.Key_accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) (key_accidentalFormCallback *FormCallback[*models.Key_accidental]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKey_accidentalFields,
	)
}

type Key_accidentalFormCallback = FormCallback[*models.Key_accidental]

func saveKey_accidentalFields(
	_instance *models.Key_accidental,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Key_octaveFormCallback(
	_instance *models.Key_octave,
	probe *Probe,
	formGroup *form.FormGroup,
) (key_octaveFormCallback *FormCallback[*models.Key_octave]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKey_octaveFields,
	)
}

type Key_octaveFormCallback = FormCallback[*models.Key_octave]

func saveKey_octaveFields(
	_instance *models.Key_octave,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Cancel":
			FormDivEnumStringFieldToField(&(_instance.Cancel), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Key:Key_octave":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Key_octave", func(owner *models.Key) *[]*models.Key_octave { return &owner.Key_octave })
		}
	}
}

func __gong__New__KindFormCallback(
	_instance *models.Kind,
	probe *Probe,
	formGroup *form.FormGroup,
) (kindFormCallback *FormCallback[*models.Kind]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKindFields,
	)
}

type KindFormCallback = FormCallback[*models.Kind]

func saveKindFields(
	_instance *models.Kind,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Use_symbols":
			FormDivEnumStringFieldToField(&(_instance.Use_symbols), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Stack_degrees":
			FormDivEnumStringFieldToField(&(_instance.Stack_degrees), formDiv)
		case "Parentheses_degrees":
			FormDivEnumStringFieldToField(&(_instance.Parentheses_degrees), formDiv)
		case "Bracket_degrees":
			FormDivEnumStringFieldToField(&(_instance.Bracket_degrees), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__LevelFormCallback(
	_instance *models.Level,
	probe *Probe,
	formGroup *form.FormGroup,
) (levelFormCallback *FormCallback[*models.Level]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLevelFields,
	)
}

type LevelFormCallback = FormCallback[*models.Level]

func saveLevelFields(
	_instance *models.Level,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Reference":
			FormDivEnumStringFieldToField(&(_instance.Reference), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Bracket":
			FormDivEnumStringFieldToField(&(_instance.Bracket), formDiv)
		case "Size":
			FormDivEnumStringFieldToField(&(_instance.Size), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Line_detailFormCallback(
	_instance *models.Line_detail,
	probe *Probe,
	formGroup *form.FormGroup,
) (line_detailFormCallback *FormCallback[*models.Line_detail]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLine_detailFields,
	)
}

type Line_detailFormCallback = FormCallback[*models.Line_detail]

func saveLine_detailFields(
	_instance *models.Line_detail,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Line":
			FormDivBasicFieldToField(&(_instance.Line), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Staff_details:Line_detail":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Line_detail", func(owner *models.Staff_details) *[]*models.Line_detail { return &owner.Line_detail })
		}
	}
}

func __gong__New__Line_widthFormCallback(
	_instance *models.Line_width,
	probe *Probe,
	formGroup *form.FormGroup,
) (line_widthFormCallback *FormCallback[*models.Line_width]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLine_widthFields,
	)
}

type Line_widthFormCallback = FormCallback[*models.Line_width]

func saveLine_widthFields(
	_instance *models.Line_width,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Appearance:Line_width":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Line_width", func(owner *models.Appearance) *[]*models.Line_width { return &owner.Line_width })
		}
	}
}

func __gong__New__LinkFormCallback(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *FormCallback[*models.Link]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkFields,
	)
}

type LinkFormCallback = FormCallback[*models.Link]

func saveLinkFields(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Href":
			FormDivBasicFieldToField(&(_instance.Href), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Role":
			FormDivBasicFieldToField(&(_instance.Role), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(_instance.Title), formDiv)
		case "Show":
			FormDivBasicFieldToField(&(_instance.Show), formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&(_instance.Actuate), formDiv)
		case "Element":
			FormDivBasicFieldToField(&(_instance.Element), formDiv)
		case "Position":
			FormDivBasicFieldToField(&(_instance.Position), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "A_measure:Link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Link", func(owner *models.A_measure) *[]*models.Link { return &owner.Link })
		case "A_part_1:Link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Link", func(owner *models.A_part_1) *[]*models.Link { return &owner.Link })
		case "Credit:Link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Link", func(owner *models.Credit) *[]*models.Link { return &owner.Link })
		}
	}
}

func __gong__New__ListenFormCallback(
	_instance *models.Listen,
	probe *Probe,
	formGroup *form.FormGroup,
) (listenFormCallback *FormCallback[*models.Listen]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveListenFields,
	)
}

type ListenFormCallback = FormCallback[*models.Listen]

func saveListenFields(
	_instance *models.Listen,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Assess":
			FormDivSliceOfPointersToField(_instance, "Assess", &(_instance.Assess), formDiv, probe)
		case "Wait":
			FormDivSliceOfPointersToField(_instance, "Wait", &(_instance.Wait), formDiv, probe)
		case "Other_listen":
			FormDivSliceOfPointersToField(_instance, "Other_listen", &(_instance.Other_listen), formDiv, probe)
		}
	}
}

func __gong__New__ListeningFormCallback(
	_instance *models.Listening,
	probe *Probe,
	formGroup *form.FormGroup,
) (listeningFormCallback *FormCallback[*models.Listening]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveListeningFields,
	)
}

type ListeningFormCallback = FormCallback[*models.Listening]

func saveListeningFields(
	_instance *models.Listening,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Sync":
			FormDivSliceOfPointersToField(_instance, "Sync", &(_instance.Sync), formDiv, probe)
		case "Other_listening":
			FormDivSliceOfPointersToField(_instance, "Other_listening", &(_instance.Other_listening), formDiv, probe)
		case "Offset":
			FormDivSelectFieldToField(&(_instance.Offset), probe.stageOfInterest, formDiv)
		case "A_measure:Listening":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Listening", func(owner *models.A_measure) *[]*models.Listening { return &owner.Listening })
		case "A_part_1:Listening":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Listening", func(owner *models.A_part_1) *[]*models.Listening { return &owner.Listening })
		}
	}
}

func __gong__New__LyricFormCallback(
	_instance *models.Lyric,
	probe *Probe,
	formGroup *form.FormGroup,
) (lyricFormCallback *FormCallback[*models.Lyric]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLyricFields,
	)
}

type LyricFormCallback = FormCallback[*models.Lyric]

func saveLyricFields(
	_instance *models.Lyric,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Elision":
			FormDivSliceOfPointersToField(_instance, "Elision", &(_instance.Elision), formDiv, probe)
		case "Syllabic":
			FormDivBasicFieldToField(&(_instance.Syllabic), formDiv)
		case "Text":
			FormDivSliceOfPointersToField(_instance, "Text", &(_instance.Text), formDiv, probe)
		case "Extend":
			FormDivSelectFieldToField(&(_instance.Extend), probe.stageOfInterest, formDiv)
		case "Laughing":
			FormDivBasicFieldToField(&(_instance.Laughing), formDiv)
		case "Humming":
			FormDivBasicFieldToField(&(_instance.Humming), formDiv)
		case "End_line":
			FormDivBasicFieldToField(&(_instance.End_line), formDiv)
		case "End_paragraph":
			FormDivBasicFieldToField(&(_instance.End_paragraph), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Note:Lyric":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Lyric", func(owner *models.Note) *[]*models.Lyric { return &owner.Lyric })
		}
	}
}

func __gong__New__Lyric_fontFormCallback(
	_instance *models.Lyric_font,
	probe *Probe,
	formGroup *form.FormGroup,
) (lyric_fontFormCallback *FormCallback[*models.Lyric_font]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLyric_fontFields,
	)
}

type Lyric_fontFormCallback = FormCallback[*models.Lyric_font]

func saveLyric_fontFields(
	_instance *models.Lyric_font,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Defaults:Lyric_font":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Lyric_font", func(owner *models.Defaults) *[]*models.Lyric_font { return &owner.Lyric_font })
		}
	}
}

func __gong__New__Lyric_languageFormCallback(
	_instance *models.Lyric_language,
	probe *Probe,
	formGroup *form.FormGroup,
) (lyric_languageFormCallback *FormCallback[*models.Lyric_language]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLyric_languageFields,
	)
}

type Lyric_languageFormCallback = FormCallback[*models.Lyric_language]

func saveLyric_languageFields(
	_instance *models.Lyric_language,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Defaults:Lyric_language":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Lyric_language", func(owner *models.Defaults) *[]*models.Lyric_language { return &owner.Lyric_language })
		}
	}
}

func __gong__New__Measure_layoutFormCallback(
	_instance *models.Measure_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) (measure_layoutFormCallback *FormCallback[*models.Measure_layout]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeasure_layoutFields,
	)
}

type Measure_layoutFormCallback = FormCallback[*models.Measure_layout]

func saveMeasure_layoutFields(
	_instance *models.Measure_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Measure_distance":
			FormDivBasicFieldToField(&(_instance.Measure_distance), formDiv)
		}
	}
}

func __gong__New__Measure_numberingFormCallback(
	_instance *models.Measure_numbering,
	probe *Probe,
	formGroup *form.FormGroup,
) (measure_numberingFormCallback *FormCallback[*models.Measure_numbering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeasure_numberingFields,
	)
}

type Measure_numberingFormCallback = FormCallback[*models.Measure_numbering]

func saveMeasure_numberingFields(
	_instance *models.Measure_numbering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "System":
			FormDivBasicFieldToField(&(_instance.System), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&(_instance.Staff), formDiv)
		case "Multiple_rest_always":
			FormDivEnumStringFieldToField(&(_instance.Multiple_rest_always), formDiv)
		case "Multiple_rest_range":
			FormDivEnumStringFieldToField(&(_instance.Multiple_rest_range), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Measure_repeatFormCallback(
	_instance *models.Measure_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) (measure_repeatFormCallback *FormCallback[*models.Measure_repeat]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeasure_repeatFields,
	)
}

type Measure_repeatFormCallback = FormCallback[*models.Measure_repeat]

func saveMeasure_repeatFields(
	_instance *models.Measure_repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&(_instance.Slashes), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Measure_styleFormCallback(
	_instance *models.Measure_style,
	probe *Probe,
	formGroup *form.FormGroup,
) (measure_styleFormCallback *FormCallback[*models.Measure_style]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMeasure_styleFields,
	)
}

type Measure_styleFormCallback = FormCallback[*models.Measure_style]

func saveMeasure_styleFields(
	_instance *models.Measure_style,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Multiple_rest":
			FormDivSelectFieldToField(&(_instance.Multiple_rest), probe.stageOfInterest, formDiv)
		case "Measure_repeat":
			FormDivSelectFieldToField(&(_instance.Measure_repeat), probe.stageOfInterest, formDiv)
		case "Beat_repeat":
			FormDivSelectFieldToField(&(_instance.Beat_repeat), probe.stageOfInterest, formDiv)
		case "Slash":
			FormDivSelectFieldToField(&(_instance.Slash), probe.stageOfInterest, formDiv)
		case "Attributes:Measure_style":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Measure_style", func(owner *models.Attributes) *[]*models.Measure_style { return &owner.Measure_style })
		}
	}
}

func __gong__New__MembraneFormCallback(
	_instance *models.Membrane,
	probe *Probe,
	formGroup *form.FormGroup,
) (membraneFormCallback *FormCallback[*models.Membrane]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMembraneFields,
	)
}

type MembraneFormCallback = FormCallback[*models.Membrane]

func saveMembraneFields(
	_instance *models.Membrane,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__MetalFormCallback(
	_instance *models.Metal,
	probe *Probe,
	formGroup *form.FormGroup,
) (metalFormCallback *FormCallback[*models.Metal]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetalFields,
	)
}

type MetalFormCallback = FormCallback[*models.Metal]

func saveMetalFields(
	_instance *models.Metal,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__MetronomeFormCallback(
	_instance *models.Metronome,
	probe *Probe,
	formGroup *form.FormGroup,
) (metronomeFormCallback *FormCallback[*models.Metronome]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetronomeFields,
	)
}

type MetronomeFormCallback = FormCallback[*models.Metronome]

func saveMetronomeFields(
	_instance *models.Metronome,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Beat_unit":
			FormDivEnumStringFieldToField(&(_instance.Beat_unit), formDiv)
		case "Beat_unit_dot":
			FormDivBasicFieldToField(&(_instance.Beat_unit_dot), formDiv)
		case "Per_minute":
			FormDivSelectFieldToField(&(_instance.Per_minute), probe.stageOfInterest, formDiv)
		case "Beat_unit_tied":
			FormDivSliceOfPointersToField(_instance, "Beat_unit_tied", &(_instance.Beat_unit_tied), formDiv, probe)
		case "Metronome_arrows":
			FormDivBasicFieldToField(&(_instance.Metronome_arrows), formDiv)
		case "Metronome_relation":
			FormDivBasicFieldToField(&(_instance.Metronome_relation), formDiv)
		case "Metronome_note":
			FormDivSliceOfPointersToField(_instance, "Metronome_note", &(_instance.Metronome_note), formDiv, probe)
		}
	}
}

func __gong__New__Metronome_beamFormCallback(
	_instance *models.Metronome_beam,
	probe *Probe,
	formGroup *form.FormGroup,
) (metronome_beamFormCallback *FormCallback[*models.Metronome_beam]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetronome_beamFields,
	)
}

type Metronome_beamFormCallback = FormCallback[*models.Metronome_beam]

func saveMetronome_beamFields(
	_instance *models.Metronome_beam,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		case "Metronome_note:Metronome_beam":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Metronome_beam", func(owner *models.Metronome_note) *[]*models.Metronome_beam { return &owner.Metronome_beam })
		}
	}
}

func __gong__New__Metronome_noteFormCallback(
	_instance *models.Metronome_note,
	probe *Probe,
	formGroup *form.FormGroup,
) (metronome_noteFormCallback *FormCallback[*models.Metronome_note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetronome_noteFields,
	)
}

type Metronome_noteFormCallback = FormCallback[*models.Metronome_note]

func saveMetronome_noteFields(
	_instance *models.Metronome_note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Metronome_type":
			FormDivBasicFieldToField(&(_instance.Metronome_type), formDiv)
		case "Metronome_dot":
			FormDivBasicFieldToField(&(_instance.Metronome_dot), formDiv)
		case "Metronome_beam":
			FormDivSliceOfPointersToField(_instance, "Metronome_beam", &(_instance.Metronome_beam), formDiv, probe)
		case "Metronome_tied":
			FormDivSelectFieldToField(&(_instance.Metronome_tied), probe.stageOfInterest, formDiv)
		case "Metronome_tuplet":
			FormDivSelectFieldToField(&(_instance.Metronome_tuplet), probe.stageOfInterest, formDiv)
		case "Metronome:Metronome_note":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Metronome_note", func(owner *models.Metronome) *[]*models.Metronome_note { return &owner.Metronome_note })
		}
	}
}

func __gong__New__Metronome_tiedFormCallback(
	_instance *models.Metronome_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) (metronome_tiedFormCallback *FormCallback[*models.Metronome_tied]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetronome_tiedFields,
	)
}

type Metronome_tiedFormCallback = FormCallback[*models.Metronome_tied]

func saveMetronome_tiedFields(
	_instance *models.Metronome_tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		}
	}
}

func __gong__New__Metronome_tupletFormCallback(
	_instance *models.Metronome_tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) (metronome_tupletFormCallback *FormCallback[*models.Metronome_tuplet]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetronome_tupletFields,
	)
}

type Metronome_tupletFormCallback = FormCallback[*models.Metronome_tuplet]

func saveMetronome_tupletFields(
	_instance *models.Metronome_tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Midi_deviceFormCallback(
	_instance *models.Midi_device,
	probe *Probe,
	formGroup *form.FormGroup,
) (midi_deviceFormCallback *FormCallback[*models.Midi_device]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMidi_deviceFields,
	)
}

type Midi_deviceFormCallback = FormCallback[*models.Midi_device]

func saveMidi_deviceFields(
	_instance *models.Midi_device,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Port":
			FormDivBasicFieldToField(&(_instance.Port), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Score_part:Midi_device":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Midi_device", func(owner *models.Score_part) *[]*models.Midi_device { return &owner.Midi_device })
		case "Sound:Midi_device":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Midi_device", func(owner *models.Sound) *[]*models.Midi_device { return &owner.Midi_device })
		}
	}
}

func __gong__New__Midi_instrumentFormCallback(
	_instance *models.Midi_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) (midi_instrumentFormCallback *FormCallback[*models.Midi_instrument]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMidi_instrumentFields,
	)
}

type Midi_instrumentFormCallback = FormCallback[*models.Midi_instrument]

func saveMidi_instrumentFields(
	_instance *models.Midi_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Midi_channel":
			FormDivBasicFieldToField(&(_instance.Midi_channel), formDiv)
		case "Midi_name":
			FormDivBasicFieldToField(&(_instance.Midi_name), formDiv)
		case "Midi_bank":
			FormDivBasicFieldToField(&(_instance.Midi_bank), formDiv)
		case "Midi_program":
			FormDivBasicFieldToField(&(_instance.Midi_program), formDiv)
		case "Midi_unpitched":
			FormDivBasicFieldToField(&(_instance.Midi_unpitched), formDiv)
		case "Volume":
			FormDivBasicFieldToField(&(_instance.Volume), formDiv)
		case "Pan":
			FormDivBasicFieldToField(&(_instance.Pan), formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&(_instance.Elevation), formDiv)
		case "Score_part:Midi_instrument":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Midi_instrument", func(owner *models.Score_part) *[]*models.Midi_instrument { return &owner.Midi_instrument })
		case "Sound:Midi_instrument":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Midi_instrument", func(owner *models.Sound) *[]*models.Midi_instrument { return &owner.Midi_instrument })
		}
	}
}

func __gong__New__MiscellaneousFormCallback(
	_instance *models.Miscellaneous,
	probe *Probe,
	formGroup *form.FormGroup,
) (miscellaneousFormCallback *FormCallback[*models.Miscellaneous]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMiscellaneousFields,
	)
}

type MiscellaneousFormCallback = FormCallback[*models.Miscellaneous]

func saveMiscellaneousFields(
	_instance *models.Miscellaneous,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Miscellaneous_field":
			FormDivSliceOfPointersToField(_instance, "Miscellaneous_field", &(_instance.Miscellaneous_field), formDiv, probe)
		}
	}
}

func __gong__New__Miscellaneous_fieldFormCallback(
	_instance *models.Miscellaneous_field,
	probe *Probe,
	formGroup *form.FormGroup,
) (miscellaneous_fieldFormCallback *FormCallback[*models.Miscellaneous_field]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMiscellaneous_fieldFields,
	)
}

type Miscellaneous_fieldFormCallback = FormCallback[*models.Miscellaneous_field]

func saveMiscellaneous_fieldFields(
	_instance *models.Miscellaneous_field,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Miscellaneous:Miscellaneous_field":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Miscellaneous_field", func(owner *models.Miscellaneous) *[]*models.Miscellaneous_field { return &owner.Miscellaneous_field })
		}
	}
}

func __gong__New__MordentFormCallback(
	_instance *models.Mordent,
	probe *Probe,
	formGroup *form.FormGroup,
) (mordentFormCallback *FormCallback[*models.Mordent]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMordentFields,
	)
}

type MordentFormCallback = FormCallback[*models.Mordent]

func saveMordentFields(
	_instance *models.Mordent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Ornaments:Mordent":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Mordent", func(owner *models.Ornaments) *[]*models.Mordent { return &owner.Mordent })
		case "Ornaments:Inverted_mordent":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Inverted_mordent", func(owner *models.Ornaments) *[]*models.Mordent { return &owner.Inverted_mordent })
		}
	}
}

func __gong__New__Multiple_restFormCallback(
	_instance *models.Multiple_rest,
	probe *Probe,
	formGroup *form.FormGroup,
) (multiple_restFormCallback *FormCallback[*models.Multiple_rest]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMultiple_restFields,
	)
}

type Multiple_restFormCallback = FormCallback[*models.Multiple_rest]

func saveMultiple_restFields(
	_instance *models.Multiple_rest,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Use_symbols":
			FormDivEnumStringFieldToField(&(_instance.Use_symbols), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Name_displayFormCallback(
	_instance *models.Name_display,
	probe *Probe,
	formGroup *form.FormGroup,
) (name_displayFormCallback *FormCallback[*models.Name_display]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveName_displayFields,
	)
}

type Name_displayFormCallback = FormCallback[*models.Name_display]

func saveName_displayFields(
	_instance *models.Name_display,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Display_text":
			FormDivSliceOfPointersToField(_instance, "Display_text", &(_instance.Display_text), formDiv, probe)
		case "Accidental_text":
			FormDivSliceOfPointersToField(_instance, "Accidental_text", &(_instance.Accidental_text), formDiv, probe)
		}
	}
}

func __gong__New__Non_arpeggiateFormCallback(
	_instance *models.Non_arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) (non_arpeggiateFormCallback *FormCallback[*models.Non_arpeggiate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNon_arpeggiateFields,
	)
}

type Non_arpeggiateFormCallback = FormCallback[*models.Non_arpeggiate]

func saveNon_arpeggiateFields(
	_instance *models.Non_arpeggiate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Notations:Non_arpeggiate":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Non_arpeggiate", func(owner *models.Notations) *[]*models.Non_arpeggiate { return &owner.Non_arpeggiate })
		}
	}
}

func __gong__New__NotationsFormCallback(
	_instance *models.Notations,
	probe *Probe,
	formGroup *form.FormGroup,
) (notationsFormCallback *FormCallback[*models.Notations]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNotationsFields,
	)
}

type NotationsFormCallback = FormCallback[*models.Notations]

func saveNotationsFields(
	_instance *models.Notations,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Tied":
			FormDivSliceOfPointersToField(_instance, "Tied", &(_instance.Tied), formDiv, probe)
		case "Slur":
			FormDivSliceOfPointersToField(_instance, "Slur", &(_instance.Slur), formDiv, probe)
		case "Tuplet":
			FormDivSliceOfPointersToField(_instance, "Tuplet", &(_instance.Tuplet), formDiv, probe)
		case "Glissando":
			FormDivSliceOfPointersToField(_instance, "Glissando", &(_instance.Glissando), formDiv, probe)
		case "Slide":
			FormDivSliceOfPointersToField(_instance, "Slide", &(_instance.Slide), formDiv, probe)
		case "Ornaments":
			FormDivSliceOfPointersToField(_instance, "Ornaments", &(_instance.Ornaments), formDiv, probe)
		case "Technical":
			FormDivSliceOfPointersToField(_instance, "Technical", &(_instance.Technical), formDiv, probe)
		case "Articulations":
			FormDivSliceOfPointersToField(_instance, "Articulations", &(_instance.Articulations), formDiv, probe)
		case "Dynamics":
			FormDivSliceOfPointersToField(_instance, "Dynamics", &(_instance.Dynamics), formDiv, probe)
		case "Fermata":
			FormDivSliceOfPointersToField(_instance, "Fermata", &(_instance.Fermata), formDiv, probe)
		case "Arpeggiate":
			FormDivSliceOfPointersToField(_instance, "Arpeggiate", &(_instance.Arpeggiate), formDiv, probe)
		case "Non_arpeggiate":
			FormDivSliceOfPointersToField(_instance, "Non_arpeggiate", &(_instance.Non_arpeggiate), formDiv, probe)
		case "Accidental_mark":
			FormDivSliceOfPointersToField(_instance, "Accidental_mark", &(_instance.Accidental_mark), formDiv, probe)
		case "Other_notation":
			FormDivSliceOfPointersToField(_instance, "Other_notation", &(_instance.Other_notation), formDiv, probe)
		case "Note:Notations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Notations", func(owner *models.Note) *[]*models.Notations { return &owner.Notations })
		}
	}
}

func __gong__New__NoteFormCallback(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *FormCallback[*models.Note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteFields,
	)
}

type NoteFormCallback = FormCallback[*models.Note]

func saveNoteFields(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_leger":
			FormDivEnumStringFieldToField(&(_instance.Print_leger), formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&(_instance.Dynamics), formDiv)
		case "End_dynamics":
			FormDivBasicFieldToField(&(_instance.End_dynamics), formDiv)
		case "Attack":
			FormDivBasicFieldToField(&(_instance.Attack), formDiv)
		case "Release":
			FormDivBasicFieldToField(&(_instance.Release), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "Pizzicato":
			FormDivEnumStringFieldToField(&(_instance.Pizzicato), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Print_dot":
			FormDivEnumStringFieldToField(&(_instance.Print_dot), formDiv)
		case "Print_lyric":
			FormDivEnumStringFieldToField(&(_instance.Print_lyric), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Print_spacing":
			FormDivEnumStringFieldToField(&(_instance.Print_spacing), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Grace":
			FormDivSelectFieldToField(&(_instance.Grace), probe.stageOfInterest, formDiv)
		case "Chord":
			FormDivBasicFieldToField(&(_instance.Chord), formDiv)
		case "Pitch":
			FormDivSelectFieldToField(&(_instance.Pitch), probe.stageOfInterest, formDiv)
		case "Unpitched":
			FormDivSelectFieldToField(&(_instance.Unpitched), probe.stageOfInterest, formDiv)
		case "Rest":
			FormDivSelectFieldToField(&(_instance.Rest), probe.stageOfInterest, formDiv)
		case "Tie":
			FormDivSelectFieldToField(&(_instance.Tie), probe.stageOfInterest, formDiv)
		case "Cue":
			FormDivBasicFieldToField(&(_instance.Cue), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(_instance.Duration), formDiv)
		case "Instrument":
			FormDivSliceOfPointersToField(_instance, "Instrument", &(_instance.Instrument), formDiv, probe)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		case "Voice":
			FormDivBasicFieldToField(&(_instance.Voice), formDiv)
		case "Type":
			FormDivSelectFieldToField(&(_instance.Type), probe.stageOfInterest, formDiv)
		case "Dot":
			FormDivSliceOfPointersToField(_instance, "Dot", &(_instance.Dot), formDiv, probe)
		case "Accidental":
			FormDivSelectFieldToField(&(_instance.Accidental), probe.stageOfInterest, formDiv)
		case "Time_modification":
			FormDivSelectFieldToField(&(_instance.Time_modification), probe.stageOfInterest, formDiv)
		case "Stem":
			FormDivSelectFieldToField(&(_instance.Stem), probe.stageOfInterest, formDiv)
		case "Notehead":
			FormDivSelectFieldToField(&(_instance.Notehead), probe.stageOfInterest, formDiv)
		case "Notehead_text":
			FormDivSelectFieldToField(&(_instance.Notehead_text), probe.stageOfInterest, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&(_instance.Staff), formDiv)
		case "Beam":
			FormDivSelectFieldToField(&(_instance.Beam), probe.stageOfInterest, formDiv)
		case "Notations":
			FormDivSliceOfPointersToField(_instance, "Notations", &(_instance.Notations), formDiv, probe)
		case "Lyric":
			FormDivSliceOfPointersToField(_instance, "Lyric", &(_instance.Lyric), formDiv, probe)
		case "Play":
			FormDivSelectFieldToField(&(_instance.Play), probe.stageOfInterest, formDiv)
		case "Listen":
			FormDivSelectFieldToField(&(_instance.Listen), probe.stageOfInterest, formDiv)
		case "A_measure:Note":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note", func(owner *models.A_measure) *[]*models.Note { return &owner.Note })
		case "A_part_1:Note":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note", func(owner *models.A_part_1) *[]*models.Note { return &owner.Note })
		}
	}
}

func __gong__New__Note_sizeFormCallback(
	_instance *models.Note_size,
	probe *Probe,
	formGroup *form.FormGroup,
) (note_sizeFormCallback *FormCallback[*models.Note_size]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNote_sizeFields,
	)
}

type Note_sizeFormCallback = FormCallback[*models.Note_size]

func saveNote_sizeFields(
	_instance *models.Note_size,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Appearance:Note_size":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_size", func(owner *models.Appearance) *[]*models.Note_size { return &owner.Note_size })
		}
	}
}

func __gong__New__Note_typeFormCallback(
	_instance *models.Note_type,
	probe *Probe,
	formGroup *form.FormGroup,
) (note_typeFormCallback *FormCallback[*models.Note_type]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNote_typeFields,
	)
}

type Note_typeFormCallback = FormCallback[*models.Note_type]

func saveNote_typeFields(
	_instance *models.Note_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Size":
			FormDivEnumStringFieldToField(&(_instance.Size), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__NoteheadFormCallback(
	_instance *models.Notehead,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteheadFormCallback *FormCallback[*models.Notehead]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteheadFields,
	)
}

type NoteheadFormCallback = FormCallback[*models.Notehead]

func saveNoteheadFields(
	_instance *models.Notehead,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Filled":
			FormDivEnumStringFieldToField(&(_instance.Filled), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Notehead_textFormCallback(
	_instance *models.Notehead_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (notehead_textFormCallback *FormCallback[*models.Notehead_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNotehead_textFields,
	)
}

type Notehead_textFormCallback = FormCallback[*models.Notehead_text]

func saveNotehead_textFields(
	_instance *models.Notehead_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Display_text":
			FormDivSliceOfPointersToField(_instance, "Display_text", &(_instance.Display_text), formDiv, probe)
		case "Accidental_text":
			FormDivSliceOfPointersToField(_instance, "Accidental_text", &(_instance.Accidental_text), formDiv, probe)
		}
	}
}

func __gong__New__NumeralFormCallback(
	_instance *models.Numeral,
	probe *Probe,
	formGroup *form.FormGroup,
) (numeralFormCallback *FormCallback[*models.Numeral]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNumeralFields,
	)
}

type NumeralFormCallback = FormCallback[*models.Numeral]

func saveNumeralFields(
	_instance *models.Numeral,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Numeral_root":
			FormDivSelectFieldToField(&(_instance.Numeral_root), probe.stageOfInterest, formDiv)
		case "Numeral_alter":
			FormDivSelectFieldToField(&(_instance.Numeral_alter), probe.stageOfInterest, formDiv)
		case "Numeral_key":
			FormDivSelectFieldToField(&(_instance.Numeral_key), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Numeral_keyFormCallback(
	_instance *models.Numeral_key,
	probe *Probe,
	formGroup *form.FormGroup,
) (numeral_keyFormCallback *FormCallback[*models.Numeral_key]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNumeral_keyFields,
	)
}

type Numeral_keyFormCallback = FormCallback[*models.Numeral_key]

func saveNumeral_keyFields(
	_instance *models.Numeral_key,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Numeral_fifths":
			FormDivBasicFieldToField(&(_instance.Numeral_fifths), formDiv)
		case "Numeral_mode":
			FormDivBasicFieldToField(&(_instance.Numeral_mode), formDiv)
		}
	}
}

func __gong__New__Numeral_rootFormCallback(
	_instance *models.Numeral_root,
	probe *Probe,
	formGroup *form.FormGroup,
) (numeral_rootFormCallback *FormCallback[*models.Numeral_root]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNumeral_rootFields,
	)
}

type Numeral_rootFormCallback = FormCallback[*models.Numeral_root]

func saveNumeral_rootFields(
	_instance *models.Numeral_root,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Octave_shiftFormCallback(
	_instance *models.Octave_shift,
	probe *Probe,
	formGroup *form.FormGroup,
) (octave_shiftFormCallback *FormCallback[*models.Octave_shift]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOctave_shiftFields,
	)
}

type Octave_shiftFormCallback = FormCallback[*models.Octave_shift]

func saveOctave_shiftFields(
	_instance *models.Octave_shift,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(_instance.Size), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__OffsetFormCallback(
	_instance *models.Offset,
	probe *Probe,
	formGroup *form.FormGroup,
) (offsetFormCallback *FormCallback[*models.Offset]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOffsetFields,
	)
}

type OffsetFormCallback = FormCallback[*models.Offset]

func saveOffsetFields(
	_instance *models.Offset,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Sound":
			FormDivEnumStringFieldToField(&(_instance.Sound), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__OpusFormCallback(
	_instance *models.Opus,
	probe *Probe,
	formGroup *form.FormGroup,
) (opusFormCallback *FormCallback[*models.Opus]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOpusFields,
	)
}

type OpusFormCallback = FormCallback[*models.Opus]

func saveOpusFields(
	_instance *models.Opus,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Href":
			FormDivBasicFieldToField(&(_instance.Href), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Role":
			FormDivBasicFieldToField(&(_instance.Role), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(_instance.Title), formDiv)
		case "Show":
			FormDivBasicFieldToField(&(_instance.Show), formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&(_instance.Actuate), formDiv)
		}
	}
}

func __gong__New__OrnamentsFormCallback(
	_instance *models.Ornaments,
	probe *Probe,
	formGroup *form.FormGroup,
) (ornamentsFormCallback *FormCallback[*models.Ornaments]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOrnamentsFields,
	)
}

type OrnamentsFormCallback = FormCallback[*models.Ornaments]

func saveOrnamentsFields(
	_instance *models.Ornaments,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Trill_mark":
			FormDivSliceOfPointersToField(_instance, "Trill_mark", &(_instance.Trill_mark), formDiv, probe)
		case "Turn":
			FormDivSliceOfPointersToField(_instance, "Turn", &(_instance.Turn), formDiv, probe)
		case "Delayed_turn":
			FormDivSliceOfPointersToField(_instance, "Delayed_turn", &(_instance.Delayed_turn), formDiv, probe)
		case "Inverted_turn":
			FormDivSliceOfPointersToField(_instance, "Inverted_turn", &(_instance.Inverted_turn), formDiv, probe)
		case "Delayed_inverted_turn":
			FormDivSliceOfPointersToField(_instance, "Delayed_inverted_turn", &(_instance.Delayed_inverted_turn), formDiv, probe)
		case "Vertical_turn":
			FormDivSliceOfPointersToField(_instance, "Vertical_turn", &(_instance.Vertical_turn), formDiv, probe)
		case "Inverted_vertical_turn":
			FormDivSliceOfPointersToField(_instance, "Inverted_vertical_turn", &(_instance.Inverted_vertical_turn), formDiv, probe)
		case "Shake":
			FormDivSliceOfPointersToField(_instance, "Shake", &(_instance.Shake), formDiv, probe)
		case "Wavy_line":
			FormDivSliceOfPointersToField(_instance, "Wavy_line", &(_instance.Wavy_line), formDiv, probe)
		case "Mordent":
			FormDivSliceOfPointersToField(_instance, "Mordent", &(_instance.Mordent), formDiv, probe)
		case "Inverted_mordent":
			FormDivSliceOfPointersToField(_instance, "Inverted_mordent", &(_instance.Inverted_mordent), formDiv, probe)
		case "Schleifer":
			FormDivSliceOfPointersToField(_instance, "Schleifer", &(_instance.Schleifer), formDiv, probe)
		case "Tremolo":
			FormDivSliceOfPointersToField(_instance, "Tremolo", &(_instance.Tremolo), formDiv, probe)
		case "Haydn":
			FormDivSliceOfPointersToField(_instance, "Haydn", &(_instance.Haydn), formDiv, probe)
		case "Other_ornament":
			FormDivSliceOfPointersToField(_instance, "Other_ornament", &(_instance.Other_ornament), formDiv, probe)
		case "Accidental_mark":
			FormDivSliceOfPointersToField(_instance, "Accidental_mark", &(_instance.Accidental_mark), formDiv, probe)
		case "Notations:Ornaments":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Ornaments", func(owner *models.Notations) *[]*models.Ornaments { return &owner.Ornaments })
		}
	}
}

func __gong__New__Other_appearanceFormCallback(
	_instance *models.Other_appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_appearanceFormCallback *FormCallback[*models.Other_appearance]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_appearanceFields,
	)
}

type Other_appearanceFormCallback = FormCallback[*models.Other_appearance]

func saveOther_appearanceFields(
	_instance *models.Other_appearance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Appearance:Other_appearance":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_appearance", func(owner *models.Appearance) *[]*models.Other_appearance { return &owner.Other_appearance })
		}
	}
}

func __gong__New__Other_directionFormCallback(
	_instance *models.Other_direction,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_directionFormCallback *FormCallback[*models.Other_direction]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_directionFields,
	)
}

type Other_directionFormCallback = FormCallback[*models.Other_direction]

func saveOther_directionFields(
	_instance *models.Other_direction,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Other_listeningFormCallback(
	_instance *models.Other_listening,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_listeningFormCallback *FormCallback[*models.Other_listening]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_listeningFields,
	)
}

type Other_listeningFormCallback = FormCallback[*models.Other_listening]

func saveOther_listeningFields(
	_instance *models.Other_listening,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Player":
			FormDivBasicFieldToField(&(_instance.Player), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Listen:Other_listen":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_listen", func(owner *models.Listen) *[]*models.Other_listening { return &owner.Other_listen })
		case "Listening:Other_listening":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_listening", func(owner *models.Listening) *[]*models.Other_listening { return &owner.Other_listening })
		}
	}
}

func __gong__New__Other_notationFormCallback(
	_instance *models.Other_notation,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_notationFormCallback *FormCallback[*models.Other_notation]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_notationFields,
	)
}

type Other_notationFormCallback = FormCallback[*models.Other_notation]

func saveOther_notationFields(
	_instance *models.Other_notation,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Notations:Other_notation":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_notation", func(owner *models.Notations) *[]*models.Other_notation { return &owner.Other_notation })
		}
	}
}

func __gong__New__Other_placement_textFormCallback(
	_instance *models.Other_placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_placement_textFormCallback *FormCallback[*models.Other_placement_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_placement_textFields,
	)
}

type Other_placement_textFormCallback = FormCallback[*models.Other_placement_text]

func saveOther_placement_textFields(
	_instance *models.Other_placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Articulations:Other_articulation":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_articulation", func(owner *models.Articulations) *[]*models.Other_placement_text { return &owner.Other_articulation })
		case "Ornaments:Other_ornament":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_ornament", func(owner *models.Ornaments) *[]*models.Other_placement_text { return &owner.Other_ornament })
		case "Technical:Other_technical":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_technical", func(owner *models.Technical) *[]*models.Other_placement_text { return &owner.Other_technical })
		}
	}
}

func __gong__New__Other_playFormCallback(
	_instance *models.Other_play,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_playFormCallback *FormCallback[*models.Other_play]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_playFields,
	)
}

type Other_playFormCallback = FormCallback[*models.Other_play]

func saveOther_playFields(
	_instance *models.Other_play,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Play:Other_play":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_play", func(owner *models.Play) *[]*models.Other_play { return &owner.Other_play })
		}
	}
}

func __gong__New__Other_textFormCallback(
	_instance *models.Other_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (other_textFormCallback *FormCallback[*models.Other_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOther_textFields,
	)
}

type Other_textFormCallback = FormCallback[*models.Other_text]

func saveOther_textFields(
	_instance *models.Other_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Dynamics:Other_dynamics":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Other_dynamics", func(owner *models.Dynamics) *[]*models.Other_text { return &owner.Other_dynamics })
		}
	}
}

func __gong__New__Page_layoutFormCallback(
	_instance *models.Page_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) (page_layoutFormCallback *FormCallback[*models.Page_layout]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePage_layoutFields,
	)
}

type Page_layoutFormCallback = FormCallback[*models.Page_layout]

func savePage_layoutFields(
	_instance *models.Page_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Page_height":
			FormDivBasicFieldToField(&(_instance.Page_height), formDiv)
		case "Page_width":
			FormDivBasicFieldToField(&(_instance.Page_width), formDiv)
		case "Page_margins":
			FormDivSelectFieldToField(&(_instance.Page_margins), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Page_marginsFormCallback(
	_instance *models.Page_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) (page_marginsFormCallback *FormCallback[*models.Page_margins]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePage_marginsFields,
	)
}

type Page_marginsFormCallback = FormCallback[*models.Page_margins]

func savePage_marginsFields(
	_instance *models.Page_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Left_margin":
			FormDivBasicFieldToField(&(_instance.Left_margin), formDiv)
		case "Right_margin":
			FormDivBasicFieldToField(&(_instance.Right_margin), formDiv)
		case "Top_margin":
			FormDivBasicFieldToField(&(_instance.Top_margin), formDiv)
		case "Bottom_margin":
			FormDivBasicFieldToField(&(_instance.Bottom_margin), formDiv)
		}
	}
}

func __gong__New__Part_clefFormCallback(
	_instance *models.Part_clef,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_clefFormCallback *FormCallback[*models.Part_clef]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_clefFields,
	)
}

type Part_clefFormCallback = FormCallback[*models.Part_clef]

func savePart_clefFields(
	_instance *models.Part_clef,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Sign":
			FormDivBasicFieldToField(&(_instance.Sign), formDiv)
		case "Line":
			FormDivBasicFieldToField(&(_instance.Line), formDiv)
		case "Clef_octave_change":
			FormDivBasicFieldToField(&(_instance.Clef_octave_change), formDiv)
		}
	}
}

func __gong__New__Part_groupFormCallback(
	_instance *models.Part_group,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_groupFormCallback *FormCallback[*models.Part_group]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_groupFields,
	)
}

type Part_groupFormCallback = FormCallback[*models.Part_group]

func savePart_groupFields(
	_instance *models.Part_group,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Group_name":
			FormDivSelectFieldToField(&(_instance.Group_name), probe.stageOfInterest, formDiv)
		case "Group_name_display":
			FormDivSelectFieldToField(&(_instance.Group_name_display), probe.stageOfInterest, formDiv)
		case "Group_abbreviation":
			FormDivSelectFieldToField(&(_instance.Group_abbreviation), probe.stageOfInterest, formDiv)
		case "Group_abbreviation_display":
			FormDivSelectFieldToField(&(_instance.Group_abbreviation_display), probe.stageOfInterest, formDiv)
		case "Group_symbol":
			FormDivSelectFieldToField(&(_instance.Group_symbol), probe.stageOfInterest, formDiv)
		case "Group_barline":
			FormDivSelectFieldToField(&(_instance.Group_barline), probe.stageOfInterest, formDiv)
		case "Group_time":
			FormDivBasicFieldToField(&(_instance.Group_time), formDiv)
		case "Footnote":
			FormDivSelectFieldToField(&(_instance.Footnote), probe.stageOfInterest, formDiv)
		case "Level":
			FormDivSelectFieldToField(&(_instance.Level), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Part_linkFormCallback(
	_instance *models.Part_link,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_linkFormCallback *FormCallback[*models.Part_link]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_linkFields,
	)
}

type Part_linkFormCallback = FormCallback[*models.Part_link]

func savePart_linkFields(
	_instance *models.Part_link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Href":
			FormDivBasicFieldToField(&(_instance.Href), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Role":
			FormDivBasicFieldToField(&(_instance.Role), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(_instance.Title), formDiv)
		case "Show":
			FormDivBasicFieldToField(&(_instance.Show), formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&(_instance.Actuate), formDiv)
		case "Instrument_link":
			FormDivSliceOfPointersToField(_instance, "Instrument_link", &(_instance.Instrument_link), formDiv, probe)
		case "Group_link":
			FormDivBasicFieldToField(&(_instance.Group_link), formDiv)
		case "Score_part:Part_link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Part_link", func(owner *models.Score_part) *[]*models.Part_link { return &owner.Part_link })
		}
	}
}

func __gong__New__Part_listFormCallback(
	_instance *models.Part_list,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_listFormCallback *FormCallback[*models.Part_list]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_listFields,
	)
}

type Part_listFormCallback = FormCallback[*models.Part_list]

func savePart_listFields(
	_instance *models.Part_list,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Part_group":
			FormDivSelectFieldToField(&(_instance.Part_group), probe.stageOfInterest, formDiv)
		case "Score_part":
			FormDivSelectFieldToField(&(_instance.Score_part), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Part_nameFormCallback(
	_instance *models.Part_name,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_nameFormCallback *FormCallback[*models.Part_name]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_nameFields,
	)
}

type Part_nameFormCallback = FormCallback[*models.Part_name]

func savePart_nameFields(
	_instance *models.Part_name,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Justify":
			FormDivEnumStringFieldToField(&(_instance.Justify), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Part_symbolFormCallback(
	_instance *models.Part_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_symbolFormCallback *FormCallback[*models.Part_symbol]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_symbolFields,
	)
}

type Part_symbolFormCallback = FormCallback[*models.Part_symbol]

func savePart_symbolFields(
	_instance *models.Part_symbol,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Top_staff":
			FormDivBasicFieldToField(&(_instance.Top_staff), formDiv)
		case "Bottom_staff":
			FormDivBasicFieldToField(&(_instance.Bottom_staff), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Part_transposeFormCallback(
	_instance *models.Part_transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) (part_transposeFormCallback *FormCallback[*models.Part_transpose]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePart_transposeFields,
	)
}

type Part_transposeFormCallback = FormCallback[*models.Part_transpose]

func savePart_transposeFields(
	_instance *models.Part_transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&(_instance.Diatonic), formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&(_instance.Chromatic), formDiv)
		case "Octave_change":
			FormDivBasicFieldToField(&(_instance.Octave_change), formDiv)
		case "Double":
			FormDivBasicFieldToField(&(_instance.Double), formDiv)
		}
	}
}

func __gong__New__PedalFormCallback(
	_instance *models.Pedal,
	probe *Probe,
	formGroup *form.FormGroup,
) (pedalFormCallback *FormCallback[*models.Pedal]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePedalFields,
	)
}

type PedalFormCallback = FormCallback[*models.Pedal]

func savePedalFields(
	_instance *models.Pedal,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line":
			FormDivEnumStringFieldToField(&(_instance.Line), formDiv)
		case "Sign":
			FormDivEnumStringFieldToField(&(_instance.Sign), formDiv)
		case "Abbreviated":
			FormDivEnumStringFieldToField(&(_instance.Abbreviated), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__Pedal_tuningFormCallback(
	_instance *models.Pedal_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) (pedal_tuningFormCallback *FormCallback[*models.Pedal_tuning]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePedal_tuningFields,
	)
}

type Pedal_tuningFormCallback = FormCallback[*models.Pedal_tuning]

func savePedal_tuningFields(
	_instance *models.Pedal_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Pedal_step":
			FormDivEnumStringFieldToField(&(_instance.Pedal_step), formDiv)
		case "Pedal_alter":
			FormDivBasicFieldToField(&(_instance.Pedal_alter), formDiv)
		case "Harp_pedals:Pedal_tuning":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Pedal_tuning", func(owner *models.Harp_pedals) *[]*models.Pedal_tuning { return &owner.Pedal_tuning })
		}
	}
}

func __gong__New__Per_minuteFormCallback(
	_instance *models.Per_minute,
	probe *Probe,
	formGroup *form.FormGroup,
) (per_minuteFormCallback *FormCallback[*models.Per_minute]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePer_minuteFields,
	)
}

type Per_minuteFormCallback = FormCallback[*models.Per_minute]

func savePer_minuteFields(
	_instance *models.Per_minute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__PercussionFormCallback(
	_instance *models.Percussion,
	probe *Probe,
	formGroup *form.FormGroup,
) (percussionFormCallback *FormCallback[*models.Percussion]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePercussionFields,
	)
}

type PercussionFormCallback = FormCallback[*models.Percussion]

func savePercussionFields(
	_instance *models.Percussion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Enclosure":
			FormDivBasicFieldToField(&(_instance.Enclosure), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Glass":
			FormDivSelectFieldToField(&(_instance.Glass), probe.stageOfInterest, formDiv)
		case "Metal":
			FormDivSelectFieldToField(&(_instance.Metal), probe.stageOfInterest, formDiv)
		case "Wood":
			FormDivSelectFieldToField(&(_instance.Wood), probe.stageOfInterest, formDiv)
		case "Pitched":
			FormDivSelectFieldToField(&(_instance.Pitched), probe.stageOfInterest, formDiv)
		case "Membrane":
			FormDivSelectFieldToField(&(_instance.Membrane), probe.stageOfInterest, formDiv)
		case "Effect":
			FormDivSelectFieldToField(&(_instance.Effect), probe.stageOfInterest, formDiv)
		case "Timpani":
			FormDivSelectFieldToField(&(_instance.Timpani), probe.stageOfInterest, formDiv)
		case "Beater":
			FormDivSelectFieldToField(&(_instance.Beater), probe.stageOfInterest, formDiv)
		case "Stick":
			FormDivSelectFieldToField(&(_instance.Stick), probe.stageOfInterest, formDiv)
		case "Stick_location":
			FormDivBasicFieldToField(&(_instance.Stick_location), formDiv)
		case "Other_percussion":
			FormDivSelectFieldToField(&(_instance.Other_percussion), probe.stageOfInterest, formDiv)
		case "Direction_type:Percussion":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Percussion", func(owner *models.Direction_type) *[]*models.Percussion { return &owner.Percussion })
		}
	}
}

func __gong__New__PitchFormCallback(
	_instance *models.Pitch,
	probe *Probe,
	formGroup *form.FormGroup,
) (pitchFormCallback *FormCallback[*models.Pitch]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePitchFields,
	)
}

type PitchFormCallback = FormCallback[*models.Pitch]

func savePitchFields(
	_instance *models.Pitch,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Step":
			FormDivEnumStringFieldToField(&(_instance.Step), formDiv)
		case "Alter":
			FormDivBasicFieldToField(&(_instance.Alter), formDiv)
		case "Octave":
			FormDivBasicFieldToField(&(_instance.Octave), formDiv)
		}
	}
}

func __gong__New__PitchedFormCallback(
	_instance *models.Pitched,
	probe *Probe,
	formGroup *form.FormGroup,
) (pitchedFormCallback *FormCallback[*models.Pitched]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePitchedFields,
	)
}

type PitchedFormCallback = FormCallback[*models.Pitched]

func savePitchedFields(
	_instance *models.Pitched,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Placement_textFormCallback(
	_instance *models.Placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (placement_textFormCallback *FormCallback[*models.Placement_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlacement_textFields,
	)
}

type Placement_textFormCallback = FormCallback[*models.Placement_text]

func savePlacement_textFields(
	_instance *models.Placement_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Pluck":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Pluck", func(owner *models.Technical) *[]*models.Placement_text { return &owner.Pluck })
		}
	}
}

func __gong__New__PlayFormCallback(
	_instance *models.Play,
	probe *Probe,
	formGroup *form.FormGroup,
) (playFormCallback *FormCallback[*models.Play]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlayFields,
	)
}

type PlayFormCallback = FormCallback[*models.Play]

func savePlayFields(
	_instance *models.Play,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Ipa":
			FormDivBasicFieldToField(&(_instance.Ipa), formDiv)
		case "Mute":
			FormDivBasicFieldToField(&(_instance.Mute), formDiv)
		case "Semi_pitched":
			FormDivBasicFieldToField(&(_instance.Semi_pitched), formDiv)
		case "Other_play":
			FormDivSliceOfPointersToField(_instance, "Other_play", &(_instance.Other_play), formDiv, probe)
		case "Sound:Play":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Play", func(owner *models.Sound) *[]*models.Play { return &owner.Play })
		}
	}
}

func __gong__New__PlayerFormCallback(
	_instance *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) (playerFormCallback *FormCallback[*models.Player]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlayerFields,
	)
}

type PlayerFormCallback = FormCallback[*models.Player]

func savePlayerFields(
	_instance *models.Player,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Player_name":
			FormDivBasicFieldToField(&(_instance.Player_name), formDiv)
		case "Score_part:Player":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Player", func(owner *models.Score_part) *[]*models.Player { return &owner.Player })
		}
	}
}

func __gong__New__Principal_voiceFormCallback(
	_instance *models.Principal_voice,
	probe *Probe,
	formGroup *form.FormGroup,
) (principal_voiceFormCallback *FormCallback[*models.Principal_voice]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePrincipal_voiceFields,
	)
}

type Principal_voiceFormCallback = FormCallback[*models.Principal_voice]

func savePrincipal_voiceFields(
	_instance *models.Principal_voice,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Symbol":
			FormDivBasicFieldToField(&(_instance.Symbol), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__PrintFormCallback(
	_instance *models.Print,
	probe *Probe,
	formGroup *form.FormGroup,
) (printFormCallback *FormCallback[*models.Print]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePrintFields,
	)
}

type PrintFormCallback = FormCallback[*models.Print]

func savePrintFields(
	_instance *models.Print,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Staff_spacing":
			FormDivBasicFieldToField(&(_instance.Staff_spacing), formDiv)
		case "New_system":
			FormDivEnumStringFieldToField(&(_instance.New_system), formDiv)
		case "New_page":
			FormDivEnumStringFieldToField(&(_instance.New_page), formDiv)
		case "Blank_page":
			FormDivBasicFieldToField(&(_instance.Blank_page), formDiv)
		case "Page_number":
			FormDivBasicFieldToField(&(_instance.Page_number), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Page_layout":
			FormDivSelectFieldToField(&(_instance.Page_layout), probe.stageOfInterest, formDiv)
		case "System_layout":
			FormDivSelectFieldToField(&(_instance.System_layout), probe.stageOfInterest, formDiv)
		case "Staff_layout":
			FormDivSliceOfPointersToField(_instance, "Staff_layout", &(_instance.Staff_layout), formDiv, probe)
		case "Measure_layout":
			FormDivSelectFieldToField(&(_instance.Measure_layout), probe.stageOfInterest, formDiv)
		case "Measure_numbering":
			FormDivSelectFieldToField(&(_instance.Measure_numbering), probe.stageOfInterest, formDiv)
		case "Part_name_display":
			FormDivSelectFieldToField(&(_instance.Part_name_display), probe.stageOfInterest, formDiv)
		case "Part_abbreviation_display":
			FormDivSelectFieldToField(&(_instance.Part_abbreviation_display), probe.stageOfInterest, formDiv)
		case "A_measure:Print":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Print", func(owner *models.A_measure) *[]*models.Print { return &owner.Print })
		case "A_part_1:Print":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Print", func(owner *models.A_part_1) *[]*models.Print { return &owner.Print })
		}
	}
}

func __gong__New__ReleaseFormCallback(
	_instance *models.Release,
	probe *Probe,
	formGroup *form.FormGroup,
) (releaseFormCallback *FormCallback[*models.Release]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveReleaseFields,
	)
}

type ReleaseFormCallback = FormCallback[*models.Release]

func saveReleaseFields(
	_instance *models.Release,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__RepeatFormCallback(
	_instance *models.Repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) (repeatFormCallback *FormCallback[*models.Repeat]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRepeatFields,
	)
}

type RepeatFormCallback = FormCallback[*models.Repeat]

func saveRepeatFields(
	_instance *models.Repeat,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Direction":
			FormDivBasicFieldToField(&(_instance.Direction), formDiv)
		case "Times":
			FormDivBasicFieldToField(&(_instance.Times), formDiv)
		case "After_jump":
			FormDivEnumStringFieldToField(&(_instance.After_jump), formDiv)
		case "Winged":
			FormDivBasicFieldToField(&(_instance.Winged), formDiv)
		}
	}
}

func __gong__New__RestFormCallback(
	_instance *models.Rest,
	probe *Probe,
	formGroup *form.FormGroup,
) (restFormCallback *FormCallback[*models.Rest]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRestFields,
	)
}

type RestFormCallback = FormCallback[*models.Rest]

func saveRestFields(
	_instance *models.Rest,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Measure":
			FormDivEnumStringFieldToField(&(_instance.Measure), formDiv)
		case "Display_step":
			FormDivEnumStringFieldToField(&(_instance.Display_step), formDiv)
		case "Display_octave":
			FormDivBasicFieldToField(&(_instance.Display_octave), formDiv)
		}
	}
}

func __gong__New__RootFormCallback(
	_instance *models.Root,
	probe *Probe,
	formGroup *form.FormGroup,
) (rootFormCallback *FormCallback[*models.Root]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRootFields,
	)
}

type RootFormCallback = FormCallback[*models.Root]

func saveRootFields(
	_instance *models.Root,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Root_step":
			FormDivSelectFieldToField(&(_instance.Root_step), probe.stageOfInterest, formDiv)
		case "Root_alter":
			FormDivSelectFieldToField(&(_instance.Root_alter), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Root_stepFormCallback(
	_instance *models.Root_step,
	probe *Probe,
	formGroup *form.FormGroup,
) (root_stepFormCallback *FormCallback[*models.Root_step]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRoot_stepFields,
	)
}

type Root_stepFormCallback = FormCallback[*models.Root_step]

func saveRoot_stepFields(
	_instance *models.Root_step,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__ScalingFormCallback(
	_instance *models.Scaling,
	probe *Probe,
	formGroup *form.FormGroup,
) (scalingFormCallback *FormCallback[*models.Scaling]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScalingFields,
	)
}

type ScalingFormCallback = FormCallback[*models.Scaling]

func saveScalingFields(
	_instance *models.Scaling,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Millimeters":
			FormDivBasicFieldToField(&(_instance.Millimeters), formDiv)
		case "Tenths":
			FormDivBasicFieldToField(&(_instance.Tenths), formDiv)
		}
	}
}

func __gong__New__ScordaturaFormCallback(
	_instance *models.Scordatura,
	probe *Probe,
	formGroup *form.FormGroup,
) (scordaturaFormCallback *FormCallback[*models.Scordatura]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScordaturaFields,
	)
}

type ScordaturaFormCallback = FormCallback[*models.Scordatura]

func saveScordaturaFields(
	_instance *models.Scordatura,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Accord":
			FormDivSliceOfPointersToField(_instance, "Accord", &(_instance.Accord), formDiv, probe)
		}
	}
}

func __gong__New__Score_instrumentFormCallback(
	_instance *models.Score_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) (score_instrumentFormCallback *FormCallback[*models.Score_instrument]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScore_instrumentFields,
	)
}

type Score_instrumentFormCallback = FormCallback[*models.Score_instrument]

func saveScore_instrumentFields(
	_instance *models.Score_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Instrument_name":
			FormDivBasicFieldToField(&(_instance.Instrument_name), formDiv)
		case "Instrument_abbreviation":
			FormDivBasicFieldToField(&(_instance.Instrument_abbreviation), formDiv)
		case "Instrument_sound":
			FormDivBasicFieldToField(&(_instance.Instrument_sound), formDiv)
		case "Solo":
			FormDivBasicFieldToField(&(_instance.Solo), formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&(_instance.Ensemble), formDiv)
		case "Virtual_instrument":
			FormDivSelectFieldToField(&(_instance.Virtual_instrument), probe.stageOfInterest, formDiv)
		case "Score_part:Score_instrument":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Score_instrument", func(owner *models.Score_part) *[]*models.Score_instrument { return &owner.Score_instrument })
		}
	}
}

func __gong__New__Score_partFormCallback(
	_instance *models.Score_part,
	probe *Probe,
	formGroup *form.FormGroup,
) (score_partFormCallback *FormCallback[*models.Score_part]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScore_partFields,
	)
}

type Score_partFormCallback = FormCallback[*models.Score_part]

func saveScore_partFields(
	_instance *models.Score_part,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Identification":
			FormDivSelectFieldToField(&(_instance.Identification), probe.stageOfInterest, formDiv)
		case "Part_link":
			FormDivSliceOfPointersToField(_instance, "Part_link", &(_instance.Part_link), formDiv, probe)
		case "Part_name":
			FormDivSelectFieldToField(&(_instance.Part_name), probe.stageOfInterest, formDiv)
		case "Part_name_display":
			FormDivSelectFieldToField(&(_instance.Part_name_display), probe.stageOfInterest, formDiv)
		case "Part_abbreviation":
			FormDivSelectFieldToField(&(_instance.Part_abbreviation), probe.stageOfInterest, formDiv)
		case "Part_abbreviation_display":
			FormDivSelectFieldToField(&(_instance.Part_abbreviation_display), probe.stageOfInterest, formDiv)
		case "Group":
			FormDivBasicFieldToField(&(_instance.Group), formDiv)
		case "Score_instrument":
			FormDivSliceOfPointersToField(_instance, "Score_instrument", &(_instance.Score_instrument), formDiv, probe)
		case "Player":
			FormDivSliceOfPointersToField(_instance, "Player", &(_instance.Player), formDiv, probe)
		case "Midi_device":
			FormDivSliceOfPointersToField(_instance, "Midi_device", &(_instance.Midi_device), formDiv, probe)
		case "Midi_instrument":
			FormDivSliceOfPointersToField(_instance, "Midi_instrument", &(_instance.Midi_instrument), formDiv, probe)
		}
	}
}

func __gong__New__Score_partwiseFormCallback(
	_instance *models.Score_partwise,
	probe *Probe,
	formGroup *form.FormGroup,
) (score_partwiseFormCallback *FormCallback[*models.Score_partwise]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScore_partwiseFields,
	)
}

type Score_partwiseFormCallback = FormCallback[*models.Score_partwise]

func saveScore_partwiseFields(
	_instance *models.Score_partwise,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Version":
			FormDivBasicFieldToField(&(_instance.Version), formDiv)
		case "Work":
			FormDivSelectFieldToField(&(_instance.Work), probe.stageOfInterest, formDiv)
		case "Movement_number":
			FormDivBasicFieldToField(&(_instance.Movement_number), formDiv)
		case "Movement_title":
			FormDivBasicFieldToField(&(_instance.Movement_title), formDiv)
		case "Identification":
			FormDivSelectFieldToField(&(_instance.Identification), probe.stageOfInterest, formDiv)
		case "Defaults":
			FormDivSelectFieldToField(&(_instance.Defaults), probe.stageOfInterest, formDiv)
		case "Credit":
			FormDivSliceOfPointersToField(_instance, "Credit", &(_instance.Credit), formDiv, probe)
		case "Part_list":
			FormDivSelectFieldToField(&(_instance.Part_list), probe.stageOfInterest, formDiv)
		case "Part":
			FormDivSliceOfPointersToField(_instance, "Part", &(_instance.Part), formDiv, probe)
		}
	}
}

func __gong__New__Score_timewiseFormCallback(
	_instance *models.Score_timewise,
	probe *Probe,
	formGroup *form.FormGroup,
) (score_timewiseFormCallback *FormCallback[*models.Score_timewise]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScore_timewiseFields,
	)
}

type Score_timewiseFormCallback = FormCallback[*models.Score_timewise]

func saveScore_timewiseFields(
	_instance *models.Score_timewise,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Version":
			FormDivBasicFieldToField(&(_instance.Version), formDiv)
		case "Work":
			FormDivSelectFieldToField(&(_instance.Work), probe.stageOfInterest, formDiv)
		case "Movement_number":
			FormDivBasicFieldToField(&(_instance.Movement_number), formDiv)
		case "Movement_title":
			FormDivBasicFieldToField(&(_instance.Movement_title), formDiv)
		case "Identification":
			FormDivSelectFieldToField(&(_instance.Identification), probe.stageOfInterest, formDiv)
		case "Defaults":
			FormDivSelectFieldToField(&(_instance.Defaults), probe.stageOfInterest, formDiv)
		case "Credit":
			FormDivSliceOfPointersToField(_instance, "Credit", &(_instance.Credit), formDiv, probe)
		case "Part_list":
			FormDivSelectFieldToField(&(_instance.Part_list), probe.stageOfInterest, formDiv)
		case "Measure":
			FormDivSliceOfPointersToField(_instance, "Measure", &(_instance.Measure), formDiv, probe)
		}
	}
}

func __gong__New__SegnoFormCallback(
	_instance *models.Segno,
	probe *Probe,
	formGroup *form.FormGroup,
) (segnoFormCallback *FormCallback[*models.Segno]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSegnoFields,
	)
}

type SegnoFormCallback = FormCallback[*models.Segno]

func saveSegnoFields(
	_instance *models.Segno,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Direction_type:Segno":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Segno", func(owner *models.Direction_type) *[]*models.Segno { return &owner.Segno })
		}
	}
}

func __gong__New__SlashFormCallback(
	_instance *models.Slash,
	probe *Probe,
	formGroup *form.FormGroup,
) (slashFormCallback *FormCallback[*models.Slash]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSlashFields,
	)
}

type SlashFormCallback = FormCallback[*models.Slash]

func saveSlashFields(
	_instance *models.Slash,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Use_dots":
			FormDivEnumStringFieldToField(&(_instance.Use_dots), formDiv)
		case "Use_stems":
			FormDivEnumStringFieldToField(&(_instance.Use_stems), formDiv)
		case "Slash_type":
			FormDivEnumStringFieldToField(&(_instance.Slash_type), formDiv)
		case "Slash_dot":
			FormDivBasicFieldToField(&(_instance.Slash_dot), formDiv)
		case "Except_voice":
			FormDivBasicFieldToField(&(_instance.Except_voice), formDiv)
		}
	}
}

func __gong__New__SlideFormCallback(
	_instance *models.Slide,
	probe *Probe,
	formGroup *form.FormGroup,
) (slideFormCallback *FormCallback[*models.Slide]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSlideFields,
	)
}

type SlideFormCallback = FormCallback[*models.Slide]

func saveSlideFields(
	_instance *models.Slide,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Accelerate":
			FormDivEnumStringFieldToField(&(_instance.Accelerate), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "First_beat":
			FormDivBasicFieldToField(&(_instance.First_beat), formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&(_instance.Last_beat), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Notations:Slide":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Slide", func(owner *models.Notations) *[]*models.Slide { return &owner.Slide })
		}
	}
}

func __gong__New__SlurFormCallback(
	_instance *models.Slur,
	probe *Probe,
	formGroup *form.FormGroup,
) (slurFormCallback *FormCallback[*models.Slur]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSlurFields,
	)
}

type SlurFormCallback = FormCallback[*models.Slur]

func saveSlurFields(
	_instance *models.Slur,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Orientation":
			FormDivBasicFieldToField(&(_instance.Orientation), formDiv)
		case "Bezier_x":
			FormDivBasicFieldToField(&(_instance.Bezier_x), formDiv)
		case "Bezier_y":
			FormDivBasicFieldToField(&(_instance.Bezier_y), formDiv)
		case "Bezier_x2":
			FormDivBasicFieldToField(&(_instance.Bezier_x2), formDiv)
		case "Bezier_y2":
			FormDivBasicFieldToField(&(_instance.Bezier_y2), formDiv)
		case "Bezier_offset":
			FormDivBasicFieldToField(&(_instance.Bezier_offset), formDiv)
		case "Bezier_offset2":
			FormDivBasicFieldToField(&(_instance.Bezier_offset2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Notations:Slur":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Slur", func(owner *models.Notations) *[]*models.Slur { return &owner.Slur })
		}
	}
}

func __gong__New__SoundFormCallback(
	_instance *models.Sound,
	probe *Probe,
	formGroup *form.FormGroup,
) (soundFormCallback *FormCallback[*models.Sound]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSoundFields,
	)
}

type SoundFormCallback = FormCallback[*models.Sound]

func saveSoundFields(
	_instance *models.Sound,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tempo":
			FormDivBasicFieldToField(&(_instance.Tempo), formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&(_instance.Dynamics), formDiv)
		case "Dacapo":
			FormDivEnumStringFieldToField(&(_instance.Dacapo), formDiv)
		case "Segno":
			FormDivBasicFieldToField(&(_instance.Segno), formDiv)
		case "Dalsegno":
			FormDivBasicFieldToField(&(_instance.Dalsegno), formDiv)
		case "Coda":
			FormDivBasicFieldToField(&(_instance.Coda), formDiv)
		case "Tocoda":
			FormDivBasicFieldToField(&(_instance.Tocoda), formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&(_instance.Divisions), formDiv)
		case "Forward_repeat":
			FormDivEnumStringFieldToField(&(_instance.Forward_repeat), formDiv)
		case "Fine":
			FormDivBasicFieldToField(&(_instance.Fine), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "Pizzicato":
			FormDivEnumStringFieldToField(&(_instance.Pizzicato), formDiv)
		case "Pan":
			FormDivBasicFieldToField(&(_instance.Pan), formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&(_instance.Elevation), formDiv)
		case "Damper_pedal":
			FormDivBasicFieldToField(&(_instance.Damper_pedal), formDiv)
		case "Soft_pedal":
			FormDivBasicFieldToField(&(_instance.Soft_pedal), formDiv)
		case "Sostenuto_pedal":
			FormDivBasicFieldToField(&(_instance.Sostenuto_pedal), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Instrument_change":
			FormDivSliceOfPointersToField(_instance, "Instrument_change", &(_instance.Instrument_change), formDiv, probe)
		case "Midi_device":
			FormDivSliceOfPointersToField(_instance, "Midi_device", &(_instance.Midi_device), formDiv, probe)
		case "Midi_instrument":
			FormDivSliceOfPointersToField(_instance, "Midi_instrument", &(_instance.Midi_instrument), formDiv, probe)
		case "Play":
			FormDivSliceOfPointersToField(_instance, "Play", &(_instance.Play), formDiv, probe)
		case "Swing":
			FormDivSelectFieldToField(&(_instance.Swing), probe.stageOfInterest, formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(_instance.Offset), probe.stageOfInterest, formDiv)
		case "A_measure:Sound":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sound", func(owner *models.A_measure) *[]*models.Sound { return &owner.Sound })
		case "A_part_1:Sound":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sound", func(owner *models.A_part_1) *[]*models.Sound { return &owner.Sound })
		}
	}
}

func __gong__New__Staff_detailsFormCallback(
	_instance *models.Staff_details,
	probe *Probe,
	formGroup *form.FormGroup,
) (staff_detailsFormCallback *FormCallback[*models.Staff_details]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaff_detailsFields,
	)
}

type Staff_detailsFormCallback = FormCallback[*models.Staff_details]

func saveStaff_detailsFields(
	_instance *models.Staff_details,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Show_frets":
			FormDivBasicFieldToField(&(_instance.Show_frets), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Print_spacing":
			FormDivEnumStringFieldToField(&(_instance.Print_spacing), formDiv)
		case "Staff_type":
			FormDivBasicFieldToField(&(_instance.Staff_type), formDiv)
		case "Staff_lines":
			FormDivBasicFieldToField(&(_instance.Staff_lines), formDiv)
		case "Line_detail":
			FormDivSliceOfPointersToField(_instance, "Line_detail", &(_instance.Line_detail), formDiv, probe)
		case "Staff_tuning":
			FormDivSliceOfPointersToField(_instance, "Staff_tuning", &(_instance.Staff_tuning), formDiv, probe)
		case "Capo":
			FormDivBasicFieldToField(&(_instance.Capo), formDiv)
		case "Staff_size":
			FormDivSelectFieldToField(&(_instance.Staff_size), probe.stageOfInterest, formDiv)
		case "Attributes:Staff_details":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staff_details", func(owner *models.Attributes) *[]*models.Staff_details { return &owner.Staff_details })
		}
	}
}

func __gong__New__Staff_divideFormCallback(
	_instance *models.Staff_divide,
	probe *Probe,
	formGroup *form.FormGroup,
) (staff_divideFormCallback *FormCallback[*models.Staff_divide]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaff_divideFields,
	)
}

type Staff_divideFormCallback = FormCallback[*models.Staff_divide]

func saveStaff_divideFields(
	_instance *models.Staff_divide,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__Staff_layoutFormCallback(
	_instance *models.Staff_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) (staff_layoutFormCallback *FormCallback[*models.Staff_layout]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaff_layoutFields,
	)
}

type Staff_layoutFormCallback = FormCallback[*models.Staff_layout]

func saveStaff_layoutFields(
	_instance *models.Staff_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Staff_distance":
			FormDivBasicFieldToField(&(_instance.Staff_distance), formDiv)
		case "Defaults:Staff_layout":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staff_layout", func(owner *models.Defaults) *[]*models.Staff_layout { return &owner.Staff_layout })
		case "Print:Staff_layout":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staff_layout", func(owner *models.Print) *[]*models.Staff_layout { return &owner.Staff_layout })
		}
	}
}

func __gong__New__Staff_sizeFormCallback(
	_instance *models.Staff_size,
	probe *Probe,
	formGroup *form.FormGroup,
) (staff_sizeFormCallback *FormCallback[*models.Staff_size]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaff_sizeFields,
	)
}

type Staff_sizeFormCallback = FormCallback[*models.Staff_size]

func saveStaff_sizeFields(
	_instance *models.Staff_size,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Scaling":
			FormDivBasicFieldToField(&(_instance.Scaling), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Staff_tuningFormCallback(
	_instance *models.Staff_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) (staff_tuningFormCallback *FormCallback[*models.Staff_tuning]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaff_tuningFields,
	)
}

type Staff_tuningFormCallback = FormCallback[*models.Staff_tuning]

func saveStaff_tuningFields(
	_instance *models.Staff_tuning,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Line":
			FormDivBasicFieldToField(&(_instance.Line), formDiv)
		case "Tuning_step":
			FormDivEnumStringFieldToField(&(_instance.Tuning_step), formDiv)
		case "Tuning_alter":
			FormDivBasicFieldToField(&(_instance.Tuning_alter), formDiv)
		case "Tuning_octave":
			FormDivBasicFieldToField(&(_instance.Tuning_octave), formDiv)
		case "Staff_details:Staff_tuning":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Staff_tuning", func(owner *models.Staff_details) *[]*models.Staff_tuning { return &owner.Staff_tuning })
		}
	}
}

func __gong__New__StemFormCallback(
	_instance *models.Stem,
	probe *Probe,
	formGroup *form.FormGroup,
) (stemFormCallback *FormCallback[*models.Stem]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStemFields,
	)
}

type StemFormCallback = FormCallback[*models.Stem]

func saveStemFields(
	_instance *models.Stem,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__StickFormCallback(
	_instance *models.Stick,
	probe *Probe,
	formGroup *form.FormGroup,
) (stickFormCallback *FormCallback[*models.Stick]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStickFields,
	)
}

type StickFormCallback = FormCallback[*models.Stick]

func saveStickFields(
	_instance *models.Stick,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tip":
			FormDivEnumStringFieldToField(&(_instance.Tip), formDiv)
		case "Parentheses":
			FormDivEnumStringFieldToField(&(_instance.Parentheses), formDiv)
		case "Dashed_circle":
			FormDivEnumStringFieldToField(&(_instance.Dashed_circle), formDiv)
		case "Stick_type":
			FormDivBasicFieldToField(&(_instance.Stick_type), formDiv)
		case "Stick_material":
			FormDivBasicFieldToField(&(_instance.Stick_material), formDiv)
		}
	}
}

func __gong__New__String_muteFormCallback(
	_instance *models.String_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) (string_muteFormCallback *FormCallback[*models.String_mute]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveString_muteFields,
	)
}

type String_muteFormCallback = FormCallback[*models.String_mute]

func saveString_muteFields(
	_instance *models.String_mute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__String_typeFormCallback(
	_instance *models.String_type,
	probe *Probe,
	formGroup *form.FormGroup,
) (string_typeFormCallback *FormCallback[*models.String_type]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveString_typeFields,
	)
}

type String_typeFormCallback = FormCallback[*models.String_type]

func saveString_typeFields(
	_instance *models.String_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:String":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "String", func(owner *models.Technical) *[]*models.String_type { return &owner.String })
		}
	}
}

func __gong__New__Strong_accentFormCallback(
	_instance *models.Strong_accent,
	probe *Probe,
	formGroup *form.FormGroup,
) (strong_accentFormCallback *FormCallback[*models.Strong_accent]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStrong_accentFields,
	)
}

type Strong_accentFormCallback = FormCallback[*models.Strong_accent]

func saveStrong_accentFields(
	_instance *models.Strong_accent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Articulations:Strong_accent":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Strong_accent", func(owner *models.Articulations) *[]*models.Strong_accent { return &owner.Strong_accent })
		}
	}
}

func __gong__New__Style_textFormCallback(
	_instance *models.Style_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (style_textFormCallback *FormCallback[*models.Style_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStyle_textFields,
	)
}

type Style_textFormCallback = FormCallback[*models.Style_text]

func saveStyle_textFields(
	_instance *models.Style_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__SupportsFormCallback(
	_instance *models.Supports,
	probe *Probe,
	formGroup *form.FormGroup,
) (supportsFormCallback *FormCallback[*models.Supports]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSupportsFields,
	)
}

type SupportsFormCallback = FormCallback[*models.Supports]

func saveSupportsFields(
	_instance *models.Supports,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Element":
			FormDivBasicFieldToField(&(_instance.Element), formDiv)
		case "Attribute":
			FormDivBasicFieldToField(&(_instance.Attribute), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "Encoding:Supports":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Supports", func(owner *models.Encoding) *[]*models.Supports { return &owner.Supports })
		}
	}
}

func __gong__New__SwingFormCallback(
	_instance *models.Swing,
	probe *Probe,
	formGroup *form.FormGroup,
) (swingFormCallback *FormCallback[*models.Swing]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSwingFields,
	)
}

type SwingFormCallback = FormCallback[*models.Swing]

func saveSwingFields(
	_instance *models.Swing,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Straight":
			FormDivBasicFieldToField(&(_instance.Straight), formDiv)
		case "First":
			FormDivBasicFieldToField(&(_instance.First), formDiv)
		case "Second":
			FormDivBasicFieldToField(&(_instance.Second), formDiv)
		case "Swing_type":
			FormDivEnumStringFieldToField(&(_instance.Swing_type), formDiv)
		case "Swing_style":
			FormDivBasicFieldToField(&(_instance.Swing_style), formDiv)
		}
	}
}

func __gong__New__SyncFormCallback(
	_instance *models.Sync,
	probe *Probe,
	formGroup *form.FormGroup,
) (syncFormCallback *FormCallback[*models.Sync]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSyncFields,
	)
}

type SyncFormCallback = FormCallback[*models.Sync]

func saveSyncFields(
	_instance *models.Sync,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Latency":
			FormDivBasicFieldToField(&(_instance.Latency), formDiv)
		case "Player":
			FormDivBasicFieldToField(&(_instance.Player), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "Listening:Sync":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sync", func(owner *models.Listening) *[]*models.Sync { return &owner.Sync })
		}
	}
}

func __gong__New__System_dividersFormCallback(
	_instance *models.System_dividers,
	probe *Probe,
	formGroup *form.FormGroup,
) (system_dividersFormCallback *FormCallback[*models.System_dividers]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystem_dividersFields,
	)
}

type System_dividersFormCallback = FormCallback[*models.System_dividers]

func saveSystem_dividersFields(
	_instance *models.System_dividers,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Left_divider":
			FormDivSelectFieldToField(&(_instance.Left_divider), probe.stageOfInterest, formDiv)
		case "Right_divider":
			FormDivSelectFieldToField(&(_instance.Right_divider), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__System_layoutFormCallback(
	_instance *models.System_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) (system_layoutFormCallback *FormCallback[*models.System_layout]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystem_layoutFields,
	)
}

type System_layoutFormCallback = FormCallback[*models.System_layout]

func saveSystem_layoutFields(
	_instance *models.System_layout,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "System_margins":
			FormDivSelectFieldToField(&(_instance.System_margins), probe.stageOfInterest, formDiv)
		case "System_distance":
			FormDivBasicFieldToField(&(_instance.System_distance), formDiv)
		case "Top_system_distance":
			FormDivBasicFieldToField(&(_instance.Top_system_distance), formDiv)
		case "System_dividers":
			FormDivSelectFieldToField(&(_instance.System_dividers), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__System_marginsFormCallback(
	_instance *models.System_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) (system_marginsFormCallback *FormCallback[*models.System_margins]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystem_marginsFields,
	)
}

type System_marginsFormCallback = FormCallback[*models.System_margins]

func saveSystem_marginsFields(
	_instance *models.System_margins,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Left_margin":
			FormDivBasicFieldToField(&(_instance.Left_margin), formDiv)
		case "Right_margin":
			FormDivBasicFieldToField(&(_instance.Right_margin), formDiv)
		}
	}
}

func __gong__New__TapFormCallback(
	_instance *models.Tap,
	probe *Probe,
	formGroup *form.FormGroup,
) (tapFormCallback *FormCallback[*models.Tap]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTapFields,
	)
}

type TapFormCallback = FormCallback[*models.Tap]

func saveTapFields(
	_instance *models.Tap,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Hand":
			FormDivBasicFieldToField(&(_instance.Hand), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Technical:Tap":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tap", func(owner *models.Technical) *[]*models.Tap { return &owner.Tap })
		}
	}
}

func __gong__New__TechnicalFormCallback(
	_instance *models.Technical,
	probe *Probe,
	formGroup *form.FormGroup,
) (technicalFormCallback *FormCallback[*models.Technical]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTechnicalFields,
	)
}

type TechnicalFormCallback = FormCallback[*models.Technical]

func saveTechnicalFields(
	_instance *models.Technical,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Up_bow":
			FormDivSliceOfPointersToField(_instance, "Up_bow", &(_instance.Up_bow), formDiv, probe)
		case "Down_bow":
			FormDivSliceOfPointersToField(_instance, "Down_bow", &(_instance.Down_bow), formDiv, probe)
		case "Harmonic":
			FormDivSliceOfPointersToField(_instance, "Harmonic", &(_instance.Harmonic), formDiv, probe)
		case "Open_string":
			FormDivSliceOfPointersToField(_instance, "Open_string", &(_instance.Open_string), formDiv, probe)
		case "Thumb_position":
			FormDivSliceOfPointersToField(_instance, "Thumb_position", &(_instance.Thumb_position), formDiv, probe)
		case "Fingering":
			FormDivSliceOfPointersToField(_instance, "Fingering", &(_instance.Fingering), formDiv, probe)
		case "Pluck":
			FormDivSliceOfPointersToField(_instance, "Pluck", &(_instance.Pluck), formDiv, probe)
		case "Double_tongue":
			FormDivSliceOfPointersToField(_instance, "Double_tongue", &(_instance.Double_tongue), formDiv, probe)
		case "Triple_tongue":
			FormDivSliceOfPointersToField(_instance, "Triple_tongue", &(_instance.Triple_tongue), formDiv, probe)
		case "Stopped":
			FormDivSliceOfPointersToField(_instance, "Stopped", &(_instance.Stopped), formDiv, probe)
		case "Snap_pizzicato":
			FormDivSliceOfPointersToField(_instance, "Snap_pizzicato", &(_instance.Snap_pizzicato), formDiv, probe)
		case "Fret":
			FormDivSliceOfPointersToField(_instance, "Fret", &(_instance.Fret), formDiv, probe)
		case "String":
			FormDivSliceOfPointersToField(_instance, "String", &(_instance.String), formDiv, probe)
		case "Hammer_on":
			FormDivSliceOfPointersToField(_instance, "Hammer_on", &(_instance.Hammer_on), formDiv, probe)
		case "Pull_off":
			FormDivSliceOfPointersToField(_instance, "Pull_off", &(_instance.Pull_off), formDiv, probe)
		case "Bend":
			FormDivSliceOfPointersToField(_instance, "Bend", &(_instance.Bend), formDiv, probe)
		case "Tap":
			FormDivSliceOfPointersToField(_instance, "Tap", &(_instance.Tap), formDiv, probe)
		case "Heel":
			FormDivSliceOfPointersToField(_instance, "Heel", &(_instance.Heel), formDiv, probe)
		case "Toe":
			FormDivSliceOfPointersToField(_instance, "Toe", &(_instance.Toe), formDiv, probe)
		case "Fingernails":
			FormDivSliceOfPointersToField(_instance, "Fingernails", &(_instance.Fingernails), formDiv, probe)
		case "Hole":
			FormDivSliceOfPointersToField(_instance, "Hole", &(_instance.Hole), formDiv, probe)
		case "Arrow":
			FormDivSliceOfPointersToField(_instance, "Arrow", &(_instance.Arrow), formDiv, probe)
		case "Handbell":
			FormDivSliceOfPointersToField(_instance, "Handbell", &(_instance.Handbell), formDiv, probe)
		case "Brass_bend":
			FormDivSliceOfPointersToField(_instance, "Brass_bend", &(_instance.Brass_bend), formDiv, probe)
		case "Flip":
			FormDivSliceOfPointersToField(_instance, "Flip", &(_instance.Flip), formDiv, probe)
		case "Smear":
			FormDivSliceOfPointersToField(_instance, "Smear", &(_instance.Smear), formDiv, probe)
		case "Open":
			FormDivSliceOfPointersToField(_instance, "Open", &(_instance.Open), formDiv, probe)
		case "Half_muted":
			FormDivSliceOfPointersToField(_instance, "Half_muted", &(_instance.Half_muted), formDiv, probe)
		case "Harmon_mute":
			FormDivSliceOfPointersToField(_instance, "Harmon_mute", &(_instance.Harmon_mute), formDiv, probe)
		case "Golpe":
			FormDivSliceOfPointersToField(_instance, "Golpe", &(_instance.Golpe), formDiv, probe)
		case "Other_technical":
			FormDivSliceOfPointersToField(_instance, "Other_technical", &(_instance.Other_technical), formDiv, probe)
		case "Notations:Technical":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Technical", func(owner *models.Notations) *[]*models.Technical { return &owner.Technical })
		}
	}
}

func __gong__New__Text_element_dataFormCallback(
	_instance *models.Text_element_data,
	probe *Probe,
	formGroup *form.FormGroup,
) (text_element_dataFormCallback *FormCallback[*models.Text_element_data]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveText_element_dataFields,
	)
}

type Text_element_dataFormCallback = FormCallback[*models.Text_element_data]

func saveText_element_dataFields(
	_instance *models.Text_element_data,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Underline":
			FormDivBasicFieldToField(&(_instance.Underline), formDiv)
		case "Overline":
			FormDivBasicFieldToField(&(_instance.Overline), formDiv)
		case "Line_through":
			FormDivBasicFieldToField(&(_instance.Line_through), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "Letter_spacing":
			FormDivBasicFieldToField(&(_instance.Letter_spacing), formDiv)
		case "Dir":
			FormDivBasicFieldToField(&(_instance.Dir), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Lyric:Text":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Text", func(owner *models.Lyric) *[]*models.Text_element_data { return &owner.Text })
		}
	}
}

func __gong__New__TieFormCallback(
	_instance *models.Tie,
	probe *Probe,
	formGroup *form.FormGroup,
) (tieFormCallback *FormCallback[*models.Tie]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTieFields,
	)
}

type TieFormCallback = FormCallback[*models.Tie]

func saveTieFields(
	_instance *models.Tie,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		}
	}
}

func __gong__New__TiedFormCallback(
	_instance *models.Tied,
	probe *Probe,
	formGroup *form.FormGroup,
) (tiedFormCallback *FormCallback[*models.Tied]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTiedFields,
	)
}

type TiedFormCallback = FormCallback[*models.Tied]

func saveTiedFields(
	_instance *models.Tied,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Orientation":
			FormDivBasicFieldToField(&(_instance.Orientation), formDiv)
		case "Bezier_x":
			FormDivBasicFieldToField(&(_instance.Bezier_x), formDiv)
		case "Bezier_y":
			FormDivBasicFieldToField(&(_instance.Bezier_y), formDiv)
		case "Bezier_x2":
			FormDivBasicFieldToField(&(_instance.Bezier_x2), formDiv)
		case "Bezier_y2":
			FormDivBasicFieldToField(&(_instance.Bezier_y2), formDiv)
		case "Bezier_offset":
			FormDivBasicFieldToField(&(_instance.Bezier_offset), formDiv)
		case "Bezier_offset2":
			FormDivBasicFieldToField(&(_instance.Bezier_offset2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Notations:Tied":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tied", func(owner *models.Notations) *[]*models.Tied { return &owner.Tied })
		}
	}
}

func __gong__New__TimeFormCallback(
	_instance *models.Time,
	probe *Probe,
	formGroup *form.FormGroup,
) (timeFormCallback *FormCallback[*models.Time]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTimeFields,
	)
}

type TimeFormCallback = FormCallback[*models.Time]

func saveTimeFields(
	_instance *models.Time,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Symbol":
			FormDivEnumStringFieldToField(&(_instance.Symbol), formDiv)
		case "Separator":
			FormDivEnumStringFieldToField(&(_instance.Separator), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Halign":
			FormDivBasicFieldToField(&(_instance.Halign), formDiv)
		case "Valign":
			FormDivBasicFieldToField(&(_instance.Valign), formDiv)
		case "Print_object":
			FormDivEnumStringFieldToField(&(_instance.Print_object), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "Beat_type":
			FormDivBasicFieldToField(&(_instance.Beat_type), formDiv)
		case "Interchangeable":
			FormDivSelectFieldToField(&(_instance.Interchangeable), probe.stageOfInterest, formDiv)
		case "Senza_misura":
			FormDivBasicFieldToField(&(_instance.Senza_misura), formDiv)
		case "Attributes:Time":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Time", func(owner *models.Attributes) *[]*models.Time { return &owner.Time })
		}
	}
}

func __gong__New__Time_modificationFormCallback(
	_instance *models.Time_modification,
	probe *Probe,
	formGroup *form.FormGroup,
) (time_modificationFormCallback *FormCallback[*models.Time_modification]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTime_modificationFields,
	)
}

type Time_modificationFormCallback = FormCallback[*models.Time_modification]

func saveTime_modificationFields(
	_instance *models.Time_modification,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Actual_notes":
			FormDivBasicFieldToField(&(_instance.Actual_notes), formDiv)
		case "Normal_notes":
			FormDivBasicFieldToField(&(_instance.Normal_notes), formDiv)
		case "Normal_type":
			FormDivEnumStringFieldToField(&(_instance.Normal_type), formDiv)
		case "Normal_dot":
			FormDivBasicFieldToField(&(_instance.Normal_dot), formDiv)
		}
	}
}

func __gong__New__TimpaniFormCallback(
	_instance *models.Timpani,
	probe *Probe,
	formGroup *form.FormGroup,
) (timpaniFormCallback *FormCallback[*models.Timpani]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTimpaniFields,
	)
}

type TimpaniFormCallback = FormCallback[*models.Timpani]

func saveTimpaniFields(
	_instance *models.Timpani,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		}
	}
}

func __gong__New__TransposeFormCallback(
	_instance *models.Transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) (transposeFormCallback *FormCallback[*models.Transpose]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTransposeFields,
	)
}

type TransposeFormCallback = FormCallback[*models.Transpose]

func saveTransposeFields(
	_instance *models.Transpose,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&(_instance.Diatonic), formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&(_instance.Chromatic), formDiv)
		case "Octave_change":
			FormDivBasicFieldToField(&(_instance.Octave_change), formDiv)
		case "Double":
			FormDivBasicFieldToField(&(_instance.Double), formDiv)
		case "Attributes:Transpose":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Transpose", func(owner *models.Attributes) *[]*models.Transpose { return &owner.Transpose })
		}
	}
}

func __gong__New__TremoloFormCallback(
	_instance *models.Tremolo,
	probe *Probe,
	formGroup *form.FormGroup,
) (tremoloFormCallback *FormCallback[*models.Tremolo]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTremoloFields,
	)
}

type TremoloFormCallback = FormCallback[*models.Tremolo]

func saveTremoloFields(
	_instance *models.Tremolo,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Ornaments:Tremolo":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tremolo", func(owner *models.Ornaments) *[]*models.Tremolo { return &owner.Tremolo })
		}
	}
}

func __gong__New__TupletFormCallback(
	_instance *models.Tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) (tupletFormCallback *FormCallback[*models.Tuplet]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTupletFields,
	)
}

type TupletFormCallback = FormCallback[*models.Tuplet]

func saveTupletFields(
	_instance *models.Tuplet,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Bracket":
			FormDivEnumStringFieldToField(&(_instance.Bracket), formDiv)
		case "Show_number":
			FormDivBasicFieldToField(&(_instance.Show_number), formDiv)
		case "Show_type":
			FormDivEnumStringFieldToField(&(_instance.Show_type), formDiv)
		case "Line_shape":
			FormDivBasicFieldToField(&(_instance.Line_shape), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		case "Tuplet_actual":
			FormDivSelectFieldToField(&(_instance.Tuplet_actual), probe.stageOfInterest, formDiv)
		case "Tuplet_normal":
			FormDivSelectFieldToField(&(_instance.Tuplet_normal), probe.stageOfInterest, formDiv)
		case "Notations:Tuplet":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tuplet", func(owner *models.Notations) *[]*models.Tuplet { return &owner.Tuplet })
		}
	}
}

func __gong__New__Tuplet_dotFormCallback(
	_instance *models.Tuplet_dot,
	probe *Probe,
	formGroup *form.FormGroup,
) (tuplet_dotFormCallback *FormCallback[*models.Tuplet_dot]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTuplet_dotFields,
	)
}

type Tuplet_dotFormCallback = FormCallback[*models.Tuplet_dot]

func saveTuplet_dotFields(
	_instance *models.Tuplet_dot,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Tuplet_portion:Tuplet_dot":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tuplet_dot", func(owner *models.Tuplet_portion) *[]*models.Tuplet_dot { return &owner.Tuplet_dot })
		}
	}
}

func __gong__New__Tuplet_numberFormCallback(
	_instance *models.Tuplet_number,
	probe *Probe,
	formGroup *form.FormGroup,
) (tuplet_numberFormCallback *FormCallback[*models.Tuplet_number]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTuplet_numberFields,
	)
}

type Tuplet_numberFormCallback = FormCallback[*models.Tuplet_number]

func saveTuplet_numberFields(
	_instance *models.Tuplet_number,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Tuplet_portionFormCallback(
	_instance *models.Tuplet_portion,
	probe *Probe,
	formGroup *form.FormGroup,
) (tuplet_portionFormCallback *FormCallback[*models.Tuplet_portion]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTuplet_portionFields,
	)
}

type Tuplet_portionFormCallback = FormCallback[*models.Tuplet_portion]

func saveTuplet_portionFields(
	_instance *models.Tuplet_portion,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tuplet_number":
			FormDivSelectFieldToField(&(_instance.Tuplet_number), probe.stageOfInterest, formDiv)
		case "Tuplet_type":
			FormDivSelectFieldToField(&(_instance.Tuplet_type), probe.stageOfInterest, formDiv)
		case "Tuplet_dot":
			FormDivSliceOfPointersToField(_instance, "Tuplet_dot", &(_instance.Tuplet_dot), formDiv, probe)
		}
	}
}

func __gong__New__Tuplet_typeFormCallback(
	_instance *models.Tuplet_type,
	probe *Probe,
	formGroup *form.FormGroup,
) (tuplet_typeFormCallback *FormCallback[*models.Tuplet_type]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTuplet_typeFields,
	)
}

type Tuplet_typeFormCallback = FormCallback[*models.Tuplet_type]

func saveTuplet_typeFields(
	_instance *models.Tuplet_type,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Font_family":
			FormDivBasicFieldToField(&(_instance.Font_family), formDiv)
		case "Font_style":
			FormDivBasicFieldToField(&(_instance.Font_style), formDiv)
		case "Font_size":
			FormDivBasicFieldToField(&(_instance.Font_size), formDiv)
		case "Font_weight":
			FormDivBasicFieldToField(&(_instance.Font_weight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "EnclosedText":
			FormDivEnumStringFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__Typed_textFormCallback(
	_instance *models.Typed_text,
	probe *Probe,
	formGroup *form.FormGroup,
) (typed_textFormCallback *FormCallback[*models.Typed_text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTyped_textFields,
	)
}

type Typed_textFormCallback = FormCallback[*models.Typed_text]

func saveTyped_textFields(
	_instance *models.Typed_text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Encoding:Encoder":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Encoder", func(owner *models.Encoding) *[]*models.Typed_text { return &owner.Encoder })
		case "Identification:Creator":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Creator", func(owner *models.Identification) *[]*models.Typed_text { return &owner.Creator })
		case "Identification:Rights":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Rights", func(owner *models.Identification) *[]*models.Typed_text { return &owner.Rights })
		case "Identification:Relation":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Relation", func(owner *models.Identification) *[]*models.Typed_text { return &owner.Relation })
		}
	}
}

func __gong__New__UnpitchedFormCallback(
	_instance *models.Unpitched,
	probe *Probe,
	formGroup *form.FormGroup,
) (unpitchedFormCallback *FormCallback[*models.Unpitched]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUnpitchedFields,
	)
}

type UnpitchedFormCallback = FormCallback[*models.Unpitched]

func saveUnpitchedFields(
	_instance *models.Unpitched,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Display_step":
			FormDivEnumStringFieldToField(&(_instance.Display_step), formDiv)
		case "Display_octave":
			FormDivBasicFieldToField(&(_instance.Display_octave), formDiv)
		}
	}
}

func __gong__New__Virtual_instrumentFormCallback(
	_instance *models.Virtual_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) (virtual_instrumentFormCallback *FormCallback[*models.Virtual_instrument]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVirtual_instrumentFields,
	)
}

type Virtual_instrumentFormCallback = FormCallback[*models.Virtual_instrument]

func saveVirtual_instrumentFields(
	_instance *models.Virtual_instrument,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Virtual_library":
			FormDivBasicFieldToField(&(_instance.Virtual_library), formDiv)
		case "Virtual_name":
			FormDivBasicFieldToField(&(_instance.Virtual_name), formDiv)
		}
	}
}

func __gong__New__WaitFormCallback(
	_instance *models.Wait,
	probe *Probe,
	formGroup *form.FormGroup,
) (waitFormCallback *FormCallback[*models.Wait]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWaitFields,
	)
}

type WaitFormCallback = FormCallback[*models.Wait]

func saveWaitFields(
	_instance *models.Wait,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Player":
			FormDivBasicFieldToField(&(_instance.Player), formDiv)
		case "Time_only":
			FormDivEnumStringFieldToField(&(_instance.Time_only), formDiv)
		case "Listen:Wait":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Wait", func(owner *models.Listen) *[]*models.Wait { return &owner.Wait })
		}
	}
}

func __gong__New__Wavy_lineFormCallback(
	_instance *models.Wavy_line,
	probe *Probe,
	formGroup *form.FormGroup,
) (wavy_lineFormCallback *FormCallback[*models.Wavy_line]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWavy_lineFields,
	)
}

type Wavy_lineFormCallback = FormCallback[*models.Wavy_line]

func saveWavy_lineFields(
	_instance *models.Wavy_line,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Placement":
			FormDivBasicFieldToField(&(_instance.Placement), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Start_note":
			FormDivBasicFieldToField(&(_instance.Start_note), formDiv)
		case "Trill_step":
			FormDivBasicFieldToField(&(_instance.Trill_step), formDiv)
		case "Two_note_turn":
			FormDivBasicFieldToField(&(_instance.Two_note_turn), formDiv)
		case "Accelerate":
			FormDivEnumStringFieldToField(&(_instance.Accelerate), formDiv)
		case "Beats":
			FormDivBasicFieldToField(&(_instance.Beats), formDiv)
		case "Second_beat":
			FormDivBasicFieldToField(&(_instance.Second_beat), formDiv)
		case "Last_beat":
			FormDivBasicFieldToField(&(_instance.Last_beat), formDiv)
		case "Ornaments:Wavy_line":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Wavy_line", func(owner *models.Ornaments) *[]*models.Wavy_line { return &owner.Wavy_line })
		}
	}
}

func __gong__New__WedgeFormCallback(
	_instance *models.Wedge,
	probe *Probe,
	formGroup *form.FormGroup,
) (wedgeFormCallback *FormCallback[*models.Wedge]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWedgeFields,
	)
}

type WedgeFormCallback = FormCallback[*models.Wedge]

func saveWedgeFields(
	_instance *models.Wedge,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(_instance.Number), formDiv)
		case "Spread":
			FormDivBasicFieldToField(&(_instance.Spread), formDiv)
		case "Niente":
			FormDivEnumStringFieldToField(&(_instance.Niente), formDiv)
		case "Line_type":
			FormDivBasicFieldToField(&(_instance.Line_type), formDiv)
		case "Dash_length":
			FormDivBasicFieldToField(&(_instance.Dash_length), formDiv)
		case "Space_length":
			FormDivBasicFieldToField(&(_instance.Space_length), formDiv)
		case "Default_x":
			FormDivBasicFieldToField(&(_instance.Default_x), formDiv)
		case "Default_y":
			FormDivBasicFieldToField(&(_instance.Default_y), formDiv)
		case "Relative_x":
			FormDivBasicFieldToField(&(_instance.Relative_x), formDiv)
		case "Relative_y":
			FormDivBasicFieldToField(&(_instance.Relative_y), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(_instance.Id), formDiv)
		}
	}
}

func __gong__New__WoodFormCallback(
	_instance *models.Wood,
	probe *Probe,
	formGroup *form.FormGroup,
) (woodFormCallback *FormCallback[*models.Wood]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWoodFields,
	)
}

type WoodFormCallback = FormCallback[*models.Wood]

func saveWoodFields(
	_instance *models.Wood,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&(_instance.Smufl), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		}
	}
}

func __gong__New__WorkFormCallback(
	_instance *models.Work,
	probe *Probe,
	formGroup *form.FormGroup,
) (workFormCallback *FormCallback[*models.Work]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWorkFields,
	)
}

type WorkFormCallback = FormCallback[*models.Work]

func saveWorkFields(
	_instance *models.Work,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Work_number":
			FormDivBasicFieldToField(&(_instance.Work_number), formDiv)
		case "Work_title":
			FormDivBasicFieldToField(&(_instance.Work_title), formDiv)
		case "Opus":
			FormDivSelectFieldToField(&(_instance.Opus), probe.stageOfInterest, formDiv)
		}
	}
}

