// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/load/go/models"
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
	case "FileToDownload":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FileToDownload Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FileToDownloadFormCallback(
			nil,
			probe,
			formGroup,
		)
		filetodownload := new(models.FileToDownload)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(filetodownload, formGroup, probe)
	}
	formStage.Commit()
}
