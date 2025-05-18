// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/load/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

// insertion point
func __gong__New__FileToDownloadFormCallback(
	filetodownload *models.FileToDownload,
	probe *Probe,
	formGroup *table.FormGroup,
) (filetodownloadFormCallback *FileToDownloadFormCallback) {
	filetodownloadFormCallback = new(FileToDownloadFormCallback)
	filetodownloadFormCallback.probe = probe
	filetodownloadFormCallback.filetodownload = filetodownload
	filetodownloadFormCallback.formGroup = formGroup

	filetodownloadFormCallback.CreationMode = (filetodownload == nil)

	return
}

type FileToDownloadFormCallback struct {
	filetodownload *models.FileToDownload

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (filetodownloadFormCallback *FileToDownloadFormCallback) OnSave() {

	// log.Println("FileToDownloadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	filetodownloadFormCallback.probe.formStage.Checkout()

	if filetodownloadFormCallback.filetodownload == nil {
		filetodownloadFormCallback.filetodownload = new(models.FileToDownload).Stage(filetodownloadFormCallback.probe.stageOfInterest)
	}
	filetodownload_ := filetodownloadFormCallback.filetodownload
	_ = filetodownload_

	for _, formDiv := range filetodownloadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(filetodownload_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(filetodownload_.Content), formDiv)
		}
	}

	// manage the suppress operation
	if filetodownloadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetodownload_.Unstage(filetodownloadFormCallback.probe.stageOfInterest)
	}

	filetodownloadFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FileToDownload](
		filetodownloadFormCallback.probe,
	)
	filetodownloadFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if filetodownloadFormCallback.CreationMode || filetodownloadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetodownloadFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(filetodownloadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FileToDownloadFormCallback(
			nil,
			filetodownloadFormCallback.probe,
			newFormGroup,
		)
		filetodownload := new(models.FileToDownload)
		FillUpForm(filetodownload, newFormGroup, filetodownloadFormCallback.probe)
		filetodownloadFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(filetodownloadFormCallback.probe)
}
