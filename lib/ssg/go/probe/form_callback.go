// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ChapterFormCallback(
	chapter *models.Chapter,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (chapterFormCallback *ChapterFormCallback) OnSave() {
	chapterFormCallback.probe.stageOfInterest.Lock()
	defer chapterFormCallback.probe.stageOfInterest.Unlock()

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
		case "Sections":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Section](chapterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Section, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Section)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					chapterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Section](chapterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			chapter_.Sections = instanceSlice
			chapterFormCallback.probe.UpdateSliceOfPointersCallback(chapter_, "Sections", &chapter_.Sections)

		case "Pages":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Page](chapterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Page, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Page)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					chapterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Page](chapterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			chapter_.Pages = instanceSlice
			chapterFormCallback.probe.UpdateSliceOfPointersCallback(chapter_, "Pages", &chapter_.Pages)

		case "SubChapters":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Chapter](chapterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Chapter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Chapter)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					chapterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Chapter](chapterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			chapter_.SubChapters = instanceSlice
			chapterFormCallback.probe.UpdateSliceOfPointersCallback(chapter_, "SubChapters", &chapter_.SubChapters)

		case "Chapter:SubChapters":
			// WARNING : this form deals with the N-N association "Chapter.SubChapters []*Chapter" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Chapter). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Chapter
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Chapter"
				rf.Fieldname = "SubChapters"
				formerAssociationSource := chapter_.GongGetReverseFieldOwner(
					chapterFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Chapter)
					if !ok {
						log.Fatalln("Source of Chapter.SubChapters []*Chapter, is not an Chapter instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubChapters, chapter_)
					formerSource.SubChapters = slices.Delete(formerSource.SubChapters, idx, idx+1)
				}
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
			for _chapter := range *models.GetGongstructInstancesSet[models.Chapter](chapterFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _chapter.GetName() == newSourceName.GetName() {
					newSource = _chapter // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Chapter.SubChapters []*Chapter, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubChapters = append(newSource.SubChapters, chapter_)
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
				formerAssociationSource := chapter_.GongGetReverseFieldOwner(
					chapterFormCallback.probe.stageOfInterest,
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
				if formerSource != nil {
					idx := slices.Index(formerSource.Chapters, chapter_)
					formerSource.Chapters = slices.Delete(formerSource.Chapters, idx, idx+1)
				}
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

			// (3) append the new value to the new source field
			newSource.Chapters = append(newSource.Chapters, chapter_)
		}
	}

	// manage the suppress operation
	if chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapter_.Unstage(chapterFormCallback.probe.stageOfInterest)
	}

	chapterFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Chapter](
		chapterFormCallback.probe,
	)

	// display a new form by reset the form stage
	if chapterFormCallback.CreationMode || chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapterFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	chapterFormCallback.probe.ux_tree()
}
func __gong__New__ContentFormCallback(
	content *models.Content,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (contentFormCallback *ContentFormCallback) OnSave() {
	contentFormCallback.probe.stageOfInterest.Lock()
	defer contentFormCallback.probe.stageOfInterest.Unlock()

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
		case "StaticPath":
			FormDivBasicFieldToField(&(content_.StaticPath), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(content_.LogoSVGFile), formDiv)
		case "IsBespokeLogoFileName":
			FormDivBasicFieldToField(&(content_.IsBespokeLogoFileName), formDiv)
		case "BespokeLogoFileName":
			FormDivBasicFieldToField(&(content_.BespokeLogoFileName), formDiv)
		case "IsBespokePageTileLogoFileName":
			FormDivBasicFieldToField(&(content_.IsBespokePageTileLogoFileName), formDiv)
		case "BespokePageTileLogoFileName":
			FormDivBasicFieldToField(&(content_.BespokePageTileLogoFileName), formDiv)
		case "Target":
			FormDivEnumStringFieldToField(&(content_.Target), formDiv)
		case "Chapters":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Chapter](contentFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Chapter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Chapter)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					contentFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Chapter](contentFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			content_.Chapters = instanceSlice
			contentFormCallback.probe.UpdateSliceOfPointersCallback(content_, "Chapters", &content_.Chapters)

		case "VersionInfo":
			FormDivBasicFieldToField(&(content_.VersionInfo), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Content](
		contentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if contentFormCallback.CreationMode || contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	contentFormCallback.probe.ux_tree()
}
func __gong__New__DownloadableFileFormCallback(
	downloadablefile *models.DownloadableFile,
	probe *Probe,
	formGroup *form.FormGroup,
) (downloadablefileFormCallback *DownloadableFileFormCallback) {
	downloadablefileFormCallback = new(DownloadableFileFormCallback)
	downloadablefileFormCallback.probe = probe
	downloadablefileFormCallback.downloadablefile = downloadablefile
	downloadablefileFormCallback.formGroup = formGroup

	downloadablefileFormCallback.CreationMode = (downloadablefile == nil)

	return
}

type DownloadableFileFormCallback struct {
	downloadablefile *models.DownloadableFile

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (downloadablefileFormCallback *DownloadableFileFormCallback) OnSave() {
	downloadablefileFormCallback.probe.stageOfInterest.Lock()
	defer downloadablefileFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DownloadableFileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	downloadablefileFormCallback.probe.formStage.Checkout()

	if downloadablefileFormCallback.downloadablefile == nil {
		downloadablefileFormCallback.downloadablefile = new(models.DownloadableFile).Stage(downloadablefileFormCallback.probe.stageOfInterest)
	}
	downloadablefile_ := downloadablefileFormCallback.downloadablefile
	_ = downloadablefile_

	for _, formDiv := range downloadablefileFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(downloadablefile_.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(downloadablefile_.Base64Content), formDiv)
		}
	}

	// manage the suppress operation
	if downloadablefileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		downloadablefile_.Unstage(downloadablefileFormCallback.probe.stageOfInterest)
	}

	downloadablefileFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DownloadableFile](
		downloadablefileFormCallback.probe,
	)

	// display a new form by reset the form stage
	if downloadablefileFormCallback.CreationMode || downloadablefileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		downloadablefileFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(downloadablefileFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DownloadableFileFormCallback(
			nil,
			downloadablefileFormCallback.probe,
			newFormGroup,
		)
		downloadablefile := new(models.DownloadableFile)
		FillUpForm(downloadablefile, newFormGroup, downloadablefileFormCallback.probe)
		downloadablefileFormCallback.probe.formStage.Commit()
	}

	downloadablefileFormCallback.probe.ux_tree()
}
func __gong__New__JpgImageFormCallback(
	jpgimage *models.JpgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (jpgimageFormCallback *JpgImageFormCallback) {
	jpgimageFormCallback = new(JpgImageFormCallback)
	jpgimageFormCallback.probe = probe
	jpgimageFormCallback.jpgimage = jpgimage
	jpgimageFormCallback.formGroup = formGroup

	jpgimageFormCallback.CreationMode = (jpgimage == nil)

	return
}

type JpgImageFormCallback struct {
	jpgimage *models.JpgImage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (jpgimageFormCallback *JpgImageFormCallback) OnSave() {
	jpgimageFormCallback.probe.stageOfInterest.Lock()
	defer jpgimageFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("JpgImageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	jpgimageFormCallback.probe.formStage.Checkout()

	if jpgimageFormCallback.jpgimage == nil {
		jpgimageFormCallback.jpgimage = new(models.JpgImage).Stage(jpgimageFormCallback.probe.stageOfInterest)
	}
	jpgimage_ := jpgimageFormCallback.jpgimage
	_ = jpgimage_

	for _, formDiv := range jpgimageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(jpgimage_.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(jpgimage_.Base64Content), formDiv)
		}
	}

	// manage the suppress operation
	if jpgimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		jpgimage_.Unstage(jpgimageFormCallback.probe.stageOfInterest)
	}

	jpgimageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.JpgImage](
		jpgimageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if jpgimageFormCallback.CreationMode || jpgimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		jpgimageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(jpgimageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__JpgImageFormCallback(
			nil,
			jpgimageFormCallback.probe,
			newFormGroup,
		)
		jpgimage := new(models.JpgImage)
		FillUpForm(jpgimage, newFormGroup, jpgimageFormCallback.probe)
		jpgimageFormCallback.probe.formStage.Commit()
	}

	jpgimageFormCallback.probe.ux_tree()
}
func __gong__New__PageFormCallback(
	page *models.Page,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (pageFormCallback *PageFormCallback) OnSave() {
	pageFormCallback.probe.stageOfInterest.Lock()
	defer pageFormCallback.probe.stageOfInterest.Unlock()

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
		case "Sections":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Section](pageFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Section, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Section)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					pageFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Section](pageFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			page_.Sections = instanceSlice
			pageFormCallback.probe.UpdateSliceOfPointersCallback(page_, "Sections", &page_.Sections)

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
				formerAssociationSource := page_.GongGetReverseFieldOwner(
					pageFormCallback.probe.stageOfInterest,
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
				if formerSource != nil {
					idx := slices.Index(formerSource.Pages, page_)
					formerSource.Pages = slices.Delete(formerSource.Pages, idx, idx+1)
				}
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

			// (3) append the new value to the new source field
			newSource.Pages = append(newSource.Pages, page_)
		}
	}

	// manage the suppress operation
	if pageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_.Unstage(pageFormCallback.probe.stageOfInterest)
	}

	pageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Page](
		pageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if pageFormCallback.CreationMode || pageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	pageFormCallback.probe.ux_tree()
}
func __gong__New__PngImageFormCallback(
	pngimage *models.PngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (pngimageFormCallback *PngImageFormCallback) {
	pngimageFormCallback = new(PngImageFormCallback)
	pngimageFormCallback.probe = probe
	pngimageFormCallback.pngimage = pngimage
	pngimageFormCallback.formGroup = formGroup

	pngimageFormCallback.CreationMode = (pngimage == nil)

	return
}

type PngImageFormCallback struct {
	pngimage *models.PngImage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (pngimageFormCallback *PngImageFormCallback) OnSave() {
	pngimageFormCallback.probe.stageOfInterest.Lock()
	defer pngimageFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PngImageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pngimageFormCallback.probe.formStage.Checkout()

	if pngimageFormCallback.pngimage == nil {
		pngimageFormCallback.pngimage = new(models.PngImage).Stage(pngimageFormCallback.probe.stageOfInterest)
	}
	pngimage_ := pngimageFormCallback.pngimage
	_ = pngimage_

	for _, formDiv := range pngimageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pngimage_.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(pngimage_.Base64Content), formDiv)
		}
	}

	// manage the suppress operation
	if pngimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pngimage_.Unstage(pngimageFormCallback.probe.stageOfInterest)
	}

	pngimageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PngImage](
		pngimageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if pngimageFormCallback.CreationMode || pngimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pngimageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(pngimageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PngImageFormCallback(
			nil,
			pngimageFormCallback.probe,
			newFormGroup,
		)
		pngimage := new(models.PngImage)
		FillUpForm(pngimage, newFormGroup, pngimageFormCallback.probe)
		pngimageFormCallback.probe.formStage.Commit()
	}

	pngimageFormCallback.probe.ux_tree()
}
func __gong__New__SectionFormCallback(
	section *models.Section,
	probe *Probe,
	formGroup *form.FormGroup,
) (sectionFormCallback *SectionFormCallback) {
	sectionFormCallback = new(SectionFormCallback)
	sectionFormCallback.probe = probe
	sectionFormCallback.section = section
	sectionFormCallback.formGroup = formGroup

	sectionFormCallback.CreationMode = (section == nil)

	return
}

type SectionFormCallback struct {
	section *models.Section

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (sectionFormCallback *SectionFormCallback) OnSave() {
	sectionFormCallback.probe.stageOfInterest.Lock()
	defer sectionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sectionFormCallback.probe.formStage.Checkout()

	if sectionFormCallback.section == nil {
		sectionFormCallback.section = new(models.Section).Stage(sectionFormCallback.probe.stageOfInterest)
	}
	section_ := sectionFormCallback.section
	_ = section_

	for _, formDiv := range sectionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(section_.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(section_.MardownContent), formDiv)
		case "IsImage":
			FormDivBasicFieldToField(&(section_.IsImage), formDiv)
		case "SvgImage":
			FormDivSelectFieldToField(&(section_.SvgImage), sectionFormCallback.probe.stageOfInterest, formDiv)
		case "PngImage":
			FormDivSelectFieldToField(&(section_.PngImage), sectionFormCallback.probe.stageOfInterest, formDiv)
		case "JpgImage":
			FormDivSelectFieldToField(&(section_.JpgImage), sectionFormCallback.probe.stageOfInterest, formDiv)
		case "IsDownloadableFile":
			FormDivBasicFieldToField(&(section_.IsDownloadableFile), formDiv)
		case "DownloadableFile":
			FormDivSelectFieldToField(&(section_.DownloadableFile), sectionFormCallback.probe.stageOfInterest, formDiv)
		case "Chapter:Sections":
			// WARNING : this form deals with the N-N association "Chapter.Sections []*Section" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Section). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Chapter
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Chapter"
				rf.Fieldname = "Sections"
				formerAssociationSource := section_.GongGetReverseFieldOwner(
					sectionFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Chapter)
					if !ok {
						log.Fatalln("Source of Chapter.Sections []*Section, is not an Chapter instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sections, section_)
					formerSource.Sections = slices.Delete(formerSource.Sections, idx, idx+1)
				}
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
			for _chapter := range *models.GetGongstructInstancesSet[models.Chapter](sectionFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _chapter.GetName() == newSourceName.GetName() {
					newSource = _chapter // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Chapter.Sections []*Section, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sections = append(newSource.Sections, section_)
		case "Page:Sections":
			// WARNING : this form deals with the N-N association "Page.Sections []*Section" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Section). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Page
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Page"
				rf.Fieldname = "Sections"
				formerAssociationSource := section_.GongGetReverseFieldOwner(
					sectionFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Page)
					if !ok {
						log.Fatalln("Source of Page.Sections []*Section, is not an Page instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Sections, section_)
					formerSource.Sections = slices.Delete(formerSource.Sections, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Page
			for _page := range *models.GetGongstructInstancesSet[models.Page](sectionFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _page.GetName() == newSourceName.GetName() {
					newSource = _page // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Page.Sections []*Section, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sections = append(newSource.Sections, section_)
		}
	}

	// manage the suppress operation
	if sectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		section_.Unstage(sectionFormCallback.probe.stageOfInterest)
	}

	sectionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Section](
		sectionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if sectionFormCallback.CreationMode || sectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sectionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(sectionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SectionFormCallback(
			nil,
			sectionFormCallback.probe,
			newFormGroup,
		)
		section := new(models.Section)
		FillUpForm(section, newFormGroup, sectionFormCallback.probe)
		sectionFormCallback.probe.formStage.Commit()
	}

	sectionFormCallback.probe.ux_tree()
}
func __gong__New__SvgImageFormCallback(
	svgimage *models.SvgImage,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (svgimageFormCallback *SvgImageFormCallback) OnSave() {
	svgimageFormCallback.probe.stageOfInterest.Lock()
	defer svgimageFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.SvgImage](
		svgimageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if svgimageFormCallback.CreationMode || svgimageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgimageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	svgimageFormCallback.probe.ux_tree()
}
