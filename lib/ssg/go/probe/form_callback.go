// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__ChapterFormCallback(
	chapter *models.Chapter,
	probe *Probe,
	formGroup *table.FormGroup,
) (chapterFormCallback *ChapterFormCallback) {
	chapterFormCallback = new(ChapterFormCallback)
	chapterFormCallback.probe = probe
	chapterFormCallback.chapter = chapter
	chapterFormCallback.formGroup = formGroup

	chapterFormCallback.CreationMode = (chapter == nil)

	return
}

type ChapterFormCallback struct {
	chapter *models.Chapter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (chapterFormCallback *ChapterFormCallback) OnSave() {

	log.Println("ChapterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	chapterFormCallback.probe.formStage.Checkout()

	if chapterFormCallback.chapter == nil {
		chapterFormCallback.chapter = new(models.Chapter).Stage(chapterFormCallback.probe.stageOfInterest)
	}
	chapter_ := chapterFormCallback.chapter
	_ = chapter_

	for _, formDiv := range chapterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(chapter_.Name), formDiv)
		case "Weigth":
			FormDivBasicFieldToField(&(chapter_.Weigth), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(chapter_.Description), formDiv)
		case "Content:Chapters":
			// we need to retrieve the field owner before the change
			var pastContentOwner *models.Content
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Content"
			rf.Fieldname = "Chapters"
			reverseFieldOwner := models.GetReverseFieldOwner(
				chapterFormCallback.probe.stageOfInterest,
				chapter_,
				&rf)

			if reverseFieldOwner != nil {
				pastContentOwner = reverseFieldOwner.(*models.Content)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastContentOwner != nil {
					idx := slices.Index(pastContentOwner.Chapters, chapter_)
					pastContentOwner.Chapters = slices.Delete(pastContentOwner.Chapters, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _content := range *models.GetGongstructInstancesSet[models.Content](chapterFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newContentOwner := _content // we have a match
						if pastContentOwner != nil {
							if newContentOwner != pastContentOwner {
								idx := slices.Index(pastContentOwner.Chapters, chapter_)
								pastContentOwner.Chapters = slices.Delete(pastContentOwner.Chapters, idx, idx+1)
								newContentOwner.Chapters = append(newContentOwner.Chapters, chapter_)
							}
						} else {
							newContentOwner.Chapters = append(newContentOwner.Chapters, chapter_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapter_.Unstage(chapterFormCallback.probe.stageOfInterest)
	}

	chapterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Chapter](
		chapterFormCallback.probe,
	)
	chapterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if chapterFormCallback.CreationMode || chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(chapterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ChapterFormCallback(
			nil,
			chapterFormCallback.probe,
			newFormGroup,
		)
		chapter := new(models.Chapter)
		FillUpForm(chapter, newFormGroup, chapterFormCallback.probe)
		chapterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(chapterFormCallback.probe)
}
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

	log.Println("ContentFormCallback, OnSave")

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
		case "Text":
			FormDivBasicFieldToField(&(content_.Text), formDiv)
		case "ContentPath":
			FormDivBasicFieldToField(&(content_.ContentPath), formDiv)
		case "OutputPath":
			FormDivBasicFieldToField(&(content_.OutputPath), formDiv)
		case "LayoutPath":
			FormDivBasicFieldToField(&(content_.LayoutPath), formDiv)
		case "StaticPath":
			FormDivBasicFieldToField(&(content_.StaticPath), formDiv)
		case "Target":
			FormDivEnumStringFieldToField(&(content_.Target), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Content](
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

	fillUpTree(contentFormCallback.probe)
}
