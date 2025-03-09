// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/tone/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Freqency":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Freqency Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FreqencyFormCallback(
			nil,
			probe,
			formGroup,
		)
		freqency := new(models.Freqency)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(freqency, formGroup, probe)
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "Player":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Player Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayerFormCallback(
			nil,
			probe,
			formGroup,
		)
		player := new(models.Player)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(player, formGroup, probe)
	}
	formStage.Commit()
}
