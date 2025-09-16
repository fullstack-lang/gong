// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ContentFormCallback(
	content *models.Content,
	probe *Probe,
	formGroup *table.FormGroup,
) (contentFormCallback *ContentFormCallback) {
	contentFormCallback = new(ContentFormCallback)
	contentFormCallback.probe = probe
	contentFormCallback.content = content
	contentFormCallback.formGroup = formGroup

	contentFormCallback.CreationMode = (content == nil)

	return
}

type ContentFormCallback struct {
	content *models.Content

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (contentFormCallback *ContentFormCallback) OnSave() {

	// log.Println("ContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	contentFormCallback.probe.formStage.Checkout()

	if contentFormCallback.content == nil {
		contentFormCallback.content = new(models.Content).Stage(contentFormCallback.probe.stageOfInterest)
	}
	content_ := contentFormCallback.content
	_ = content_

	for _, formDiv := range contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(content_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(content_.Content), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Content](
		contentFormCallback.probe,
	)
	contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if contentFormCallback.CreationMode || contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			contentFormCallback.probe,
			newFormGroup,
		)
		content := new(models.Content)
		FillUpForm(content, newFormGroup, contentFormCallback.probe)
		contentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(contentFormCallback.probe)
}
func __gong__New__SvgImageFormCallback(
	svgimage *models.SvgImage,
	probe *Probe,
	formGroup *table.FormGroup,
) (svgimageFormCallback *SvgImageFormCallback) {
	svgimageFormCallback = new(SvgImageFormCallback)
	svgimageFormCallback.probe = probe
	svgimageFormCallback.svgimage = svgimage
	svgimageFormCallback.formGroup = formGroup

	svgimageFormCallback.CreationMode = (svgimage == nil)

	return
}

type SvgImageFormCallback struct {
	svgimage *models.SvgImage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (svgimageFormCallback *SvgImageFormCallback) OnSave() {

	// log.Println("SvgImageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgimageFormCallback.probe.formStage.Checkout()

	if svgimageFormCallback.svgimage == nil {
		svgimageFormCallback.svgimage = new(models.SvgImage).Stage(svgimageFormCallback.probe.stageOfInterest)
	}
	svgimage_ := svgimageFormCallback.svgimage
	_ = svgimage_

	for _, formDiv := range svgimageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgimage_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(svgimage_.Content), formDiv)
		}
	}

	// manage the suppress operation
	if svgimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgimage_.Unstage(svgimageFormCallback.probe.stageOfInterest)
	}

	svgimageFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SvgImage](
		svgimageFormCallback.probe,
	)
	svgimageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgimageFormCallback.CreationMode || svgimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgimageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(svgimageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SvgImageFormCallback(
			nil,
			svgimageFormCallback.probe,
			newFormGroup,
		)
		svgimage := new(models.SvgImage)
		FillUpForm(svgimage, newFormGroup, svgimageFormCallback.probe)
		svgimageFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(svgimageFormCallback.probe)
}
