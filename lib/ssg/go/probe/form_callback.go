// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
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

	// log.Println("ChapterFormCallback, OnSave")

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
		case "MardownContent":
			FormDivBasicFieldToField(&(chapter_.MardownContent), formDiv)
		case "Content:Chapters":
			// WARNING : this form deals with the N-N association "Content.Chapters []*Chapter" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Chapter). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Content
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Content"
				rf.Fieldname = "Chapters"
				formerAssociationSource := models.GetReverseFieldOwner(
					chapterFormCallback.probe.stageOfInterest,
					chapter_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Content)
					if !ok {
						log.Fatalln("Source of Content.Chapters []*Chapter, is not an Content instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
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
			var newSource *models.Content
			for _content := range *models.GetGongstructInstancesSet[models.Content](chapterFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _content.GetName() == newSourceName.GetName() {
					newSource = _content // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Content.Chapters []*Chapter, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Chapters = append(newSource.Chapters, chapter_)
		}
	}

	// manage the suppress operation
	if chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapter_.Unstage(chapterFormCallback.probe.stageOfInterest)
	}

	chapterFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Chapter](
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

	updateAndCommitTree(chapterFormCallback.probe)
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
		case "MardownContent":
			FormDivBasicFieldToField(&(content_.MardownContent), formDiv)
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
func __gong__New__PageFormCallback(
	page *models.Page,
	probe *Probe,
	formGroup *table.FormGroup,
) (pageFormCallback *PageFormCallback) {
	pageFormCallback = new(PageFormCallback)
	pageFormCallback.probe = probe
	pageFormCallback.page = page
	pageFormCallback.formGroup = formGroup

	pageFormCallback.CreationMode = (page == nil)

	return
}

type PageFormCallback struct {
	page *models.Page

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pageFormCallback *PageFormCallback) OnSave() {

	// log.Println("PageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pageFormCallback.probe.formStage.Checkout()

	if pageFormCallback.page == nil {
		pageFormCallback.page = new(models.Page).Stage(pageFormCallback.probe.stageOfInterest)
	}
	page_ := pageFormCallback.page
	_ = page_

	for _, formDiv := range pageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(page_.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(page_.MardownContent), formDiv)
		case "Chapter:Pages":
			// WARNING : this form deals with the N-N association "Chapter.Pages []*Page" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Page). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Chapter
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Chapter"
				rf.Fieldname = "Pages"
				formerAssociationSource := models.GetReverseFieldOwner(
					pageFormCallback.probe.stageOfInterest,
					page_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Chapter)
					if !ok {
						log.Fatalln("Source of Chapter.Pages []*Page, is not an Chapter instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
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
			var newSource *models.Chapter
			for _chapter := range *models.GetGongstructInstancesSet[models.Chapter](pageFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _chapter.GetName() == newSourceName.GetName() {
					newSource = _chapter // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Chapter.Pages []*Page, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Pages = append(newSource.Pages, page_)
		}
	}

	// manage the suppress operation
	if pageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_.Unstage(pageFormCallback.probe.stageOfInterest)
	}

	pageFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Page](
		pageFormCallback.probe,
	)
	pageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pageFormCallback.CreationMode || pageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(pageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PageFormCallback(
			nil,
			pageFormCallback.probe,
			newFormGroup,
		)
		page := new(models.Page)
		FillUpForm(page, newFormGroup, pageFormCallback.probe)
		pageFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(pageFormCallback.probe)
}
