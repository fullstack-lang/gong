// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/load/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

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
func __gong__New__FileToUploadFormCallback(
	filetoupload *models.FileToUpload,
	probe *Probe,
	formGroup *table.FormGroup,
) (filetouploadFormCallback *FileToUploadFormCallback) {
	filetouploadFormCallback = new(FileToUploadFormCallback)
	filetouploadFormCallback.probe = probe
	filetouploadFormCallback.filetoupload = filetoupload
	filetouploadFormCallback.formGroup = formGroup

	filetouploadFormCallback.CreationMode = (filetoupload == nil)

	return
}

type FileToUploadFormCallback struct {
	filetoupload *models.FileToUpload

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (filetouploadFormCallback *FileToUploadFormCallback) OnSave() {

	// log.Println("FileToUploadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	filetouploadFormCallback.probe.formStage.Checkout()

	if filetouploadFormCallback.filetoupload == nil {
		filetouploadFormCallback.filetoupload = new(models.FileToUpload).Stage(filetouploadFormCallback.probe.stageOfInterest)
	}
	filetoupload_ := filetouploadFormCallback.filetoupload
	_ = filetoupload_

	for _, formDiv := range filetouploadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(filetoupload_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(filetoupload_.Content), formDiv)
		case "InvitationToUpload":
			FormDivBasicFieldToField(&(filetoupload_.InvitationToUpload), formDiv)
		case "AfterProcessingMessage":
			FormDivBasicFieldToField(&(filetoupload_.AfterProcessingMessage), formDiv)
		}
	}

	// manage the suppress operation
	if filetouploadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetoupload_.Unstage(filetouploadFormCallback.probe.stageOfInterest)
	}

	filetouploadFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FileToUpload](
		filetouploadFormCallback.probe,
	)
	filetouploadFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if filetouploadFormCallback.CreationMode || filetouploadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		filetouploadFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(filetouploadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FileToUploadFormCallback(
			nil,
			filetouploadFormCallback.probe,
			newFormGroup,
		)
		filetoupload := new(models.FileToUpload)
		FillUpForm(filetoupload, newFormGroup, filetouploadFormCallback.probe)
		filetouploadFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(filetouploadFormCallback.probe)
}
