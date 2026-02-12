// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__BodyFormCallback(
	body *models.Body,
	probe *Probe,
	formGroup *table.FormGroup,
) (bodyFormCallback *BodyFormCallback) {
	bodyFormCallback = new(BodyFormCallback)
	bodyFormCallback.probe = probe
	bodyFormCallback.body = body
	bodyFormCallback.formGroup = formGroup

	bodyFormCallback.CreationMode = (body == nil)

	return
}

type BodyFormCallback struct {
	body *models.Body

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bodyFormCallback *BodyFormCallback) OnSave() {
	bodyFormCallback.probe.stageOfInterest.Lock()
	defer bodyFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BodyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bodyFormCallback.probe.formStage.Checkout()

	if bodyFormCallback.body == nil {
		bodyFormCallback.body = new(models.Body).Stage(bodyFormCallback.probe.stageOfInterest)
	}
	body_ := bodyFormCallback.body
	_ = body_

	for _, formDiv := range bodyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(body_.Name), formDiv)
		case "Paragraphs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Paragraph](bodyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Paragraph, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Paragraph)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bodyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Paragraph](bodyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			body_.Paragraphs = instanceSlice

		case "Tables":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Table](bodyFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Table, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Table)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bodyFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Table](bodyFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			body_.Tables = instanceSlice

		case "LastParagraph":
			FormDivSelectFieldToField(&(body_.LastParagraph), bodyFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if bodyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		body_.Unstage(bodyFormCallback.probe.stageOfInterest)
	}

	bodyFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Body](
		bodyFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bodyFormCallback.CreationMode || bodyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bodyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(bodyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BodyFormCallback(
			nil,
			bodyFormCallback.probe,
			newFormGroup,
		)
		body := new(models.Body)
		FillUpForm(body, newFormGroup, bodyFormCallback.probe)
		bodyFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(bodyFormCallback.probe)
}
func __gong__New__DocumentFormCallback(
	document *models.Document,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentFormCallback *DocumentFormCallback) {
	documentFormCallback = new(DocumentFormCallback)
	documentFormCallback.probe = probe
	documentFormCallback.document = document
	documentFormCallback.formGroup = formGroup

	documentFormCallback.CreationMode = (document == nil)

	return
}

type DocumentFormCallback struct {
	document *models.Document

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentFormCallback *DocumentFormCallback) OnSave() {
	documentFormCallback.probe.stageOfInterest.Lock()
	defer documentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DocumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentFormCallback.probe.formStage.Checkout()

	if documentFormCallback.document == nil {
		documentFormCallback.document = new(models.Document).Stage(documentFormCallback.probe.stageOfInterest)
	}
	document_ := documentFormCallback.document
	_ = document_

	for _, formDiv := range documentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(document_.Name), formDiv)
		case "File":
			FormDivSelectFieldToField(&(document_.File), documentFormCallback.probe.stageOfInterest, formDiv)
		case "Root":
			FormDivSelectFieldToField(&(document_.Root), documentFormCallback.probe.stageOfInterest, formDiv)
		case "Body":
			FormDivSelectFieldToField(&(document_.Body), documentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		document_.Unstage(documentFormCallback.probe.stageOfInterest)
	}

	documentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Document](
		documentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if documentFormCallback.CreationMode || documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(documentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			documentFormCallback.probe,
			newFormGroup,
		)
		document := new(models.Document)
		FillUpForm(document, newFormGroup, documentFormCallback.probe)
		documentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(documentFormCallback.probe)
}
func __gong__New__DocxFormCallback(
	docx *models.Docx,
	probe *Probe,
	formGroup *table.FormGroup,
) (docxFormCallback *DocxFormCallback) {
	docxFormCallback = new(DocxFormCallback)
	docxFormCallback.probe = probe
	docxFormCallback.docx = docx
	docxFormCallback.formGroup = formGroup

	docxFormCallback.CreationMode = (docx == nil)

	return
}

type DocxFormCallback struct {
	docx *models.Docx

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (docxFormCallback *DocxFormCallback) OnSave() {
	docxFormCallback.probe.stageOfInterest.Lock()
	defer docxFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DocxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	docxFormCallback.probe.formStage.Checkout()

	if docxFormCallback.docx == nil {
		docxFormCallback.docx = new(models.Docx).Stage(docxFormCallback.probe.stageOfInterest)
	}
	docx_ := docxFormCallback.docx
	_ = docx_

	for _, formDiv := range docxFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(docx_.Name), formDiv)
		case "Files":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.File](docxFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.File, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.File)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					docxFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.File](docxFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			docx_.Files = instanceSlice

		case "Document":
			FormDivSelectFieldToField(&(docx_.Document), docxFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if docxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		docx_.Unstage(docxFormCallback.probe.stageOfInterest)
	}

	docxFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Docx](
		docxFormCallback.probe,
	)

	// display a new form by reset the form stage
	if docxFormCallback.CreationMode || docxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		docxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(docxFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocxFormCallback(
			nil,
			docxFormCallback.probe,
			newFormGroup,
		)
		docx := new(models.Docx)
		FillUpForm(docx, newFormGroup, docxFormCallback.probe)
		docxFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(docxFormCallback.probe)
}
func __gong__New__FileFormCallback(
	file *models.File,
	probe *Probe,
	formGroup *table.FormGroup,
) (fileFormCallback *FileFormCallback) {
	fileFormCallback = new(FileFormCallback)
	fileFormCallback.probe = probe
	fileFormCallback.file = file
	fileFormCallback.formGroup = formGroup

	fileFormCallback.CreationMode = (file == nil)

	return
}

type FileFormCallback struct {
	file *models.File

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fileFormCallback *FileFormCallback) OnSave() {
	fileFormCallback.probe.stageOfInterest.Lock()
	defer fileFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fileFormCallback.probe.formStage.Checkout()

	if fileFormCallback.file == nil {
		fileFormCallback.file = new(models.File).Stage(fileFormCallback.probe.stageOfInterest)
	}
	file_ := fileFormCallback.file
	_ = file_

	for _, formDiv := range fileFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(file_.Name), formDiv)
		case "Docx:Files":
			// WARNING : this form deals with the N-N association "Docx.Files []*File" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of File). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Docx
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Docx"
				rf.Fieldname = "Files"
				formerAssociationSource := file_.GongGetReverseFieldOwner(
					fileFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Docx)
					if !ok {
						log.Fatalln("Source of Docx.Files []*File, is not an Docx instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Files, file_)
					formerSource.Files = slices.Delete(formerSource.Files, idx, idx+1)
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
			var newSource *models.Docx
			for _docx := range *models.GetGongstructInstancesSet[models.Docx](fileFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _docx.GetName() == newSourceName.GetName() {
					newSource = _docx // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Docx.Files []*File, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Files = append(newSource.Files, file_)
		}
	}

	// manage the suppress operation
	if fileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		file_.Unstage(fileFormCallback.probe.stageOfInterest)
	}

	fileFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.File](
		fileFormCallback.probe,
	)

	// display a new form by reset the form stage
	if fileFormCallback.CreationMode || fileFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fileFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(fileFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FileFormCallback(
			nil,
			fileFormCallback.probe,
			newFormGroup,
		)
		file := new(models.File)
		FillUpForm(file, newFormGroup, fileFormCallback.probe)
		fileFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(fileFormCallback.probe)
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	probe *Probe,
	formGroup *table.FormGroup,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.probe = probe
	nodeFormCallback.node = node
	nodeFormCallback.formGroup = formGroup

	nodeFormCallback.CreationMode = (node == nil)

	return
}

type NodeFormCallback struct {
	node *models.Node

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (nodeFormCallback *NodeFormCallback) OnSave() {
	nodeFormCallback.probe.stageOfInterest.Lock()
	defer nodeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.probe.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.probe.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	for _, formDiv := range nodeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(node_.Name), formDiv)
		case "Nodes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Node](nodeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Node, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Node)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					nodeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Node](nodeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			node_.Nodes = instanceSlice

		case "Node:Nodes":
			// WARNING : this form deals with the N-N association "Node.Nodes []*Node" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Node). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Node
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Node"
				rf.Fieldname = "Nodes"
				formerAssociationSource := node_.GongGetReverseFieldOwner(
					nodeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Node)
					if !ok {
						log.Fatalln("Source of Node.Nodes []*Node, is not an Node instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Nodes, node_)
					formerSource.Nodes = slices.Delete(formerSource.Nodes, idx, idx+1)
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
			var newSource *models.Node
			for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _node.GetName() == newSourceName.GetName() {
					newSource = _node // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Node.Nodes []*Node, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Nodes = append(newSource.Nodes, node_)
		}
	}

	// manage the suppress operation
	if nodeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		node_.Unstage(nodeFormCallback.probe.stageOfInterest)
	}

	nodeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Node](
		nodeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if nodeFormCallback.CreationMode || nodeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		nodeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(nodeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NodeFormCallback(
			nil,
			nodeFormCallback.probe,
			newFormGroup,
		)
		node := new(models.Node)
		FillUpForm(node, newFormGroup, nodeFormCallback.probe)
		nodeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(nodeFormCallback.probe)
}
func __gong__New__ParagraphFormCallback(
	paragraph *models.Paragraph,
	probe *Probe,
	formGroup *table.FormGroup,
) (paragraphFormCallback *ParagraphFormCallback) {
	paragraphFormCallback = new(ParagraphFormCallback)
	paragraphFormCallback.probe = probe
	paragraphFormCallback.paragraph = paragraph
	paragraphFormCallback.formGroup = formGroup

	paragraphFormCallback.CreationMode = (paragraph == nil)

	return
}

type ParagraphFormCallback struct {
	paragraph *models.Paragraph

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (paragraphFormCallback *ParagraphFormCallback) OnSave() {
	paragraphFormCallback.probe.stageOfInterest.Lock()
	defer paragraphFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParagraphFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphFormCallback.probe.formStage.Checkout()

	if paragraphFormCallback.paragraph == nil {
		paragraphFormCallback.paragraph = new(models.Paragraph).Stage(paragraphFormCallback.probe.stageOfInterest)
	}
	paragraph_ := paragraphFormCallback.paragraph
	_ = paragraph_

	for _, formDiv := range paragraphFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraph_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraph_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraph_.Node), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "ParagraphProperties":
			FormDivSelectFieldToField(&(paragraph_.ParagraphProperties), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Runes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Rune](paragraphFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Rune, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Rune)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					paragraphFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Rune](paragraphFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			paragraph_.Runes = instanceSlice

		case "CollatedText":
			FormDivBasicFieldToField(&(paragraph_.CollatedText), formDiv)
		case "Next":
			FormDivSelectFieldToField(&(paragraph_.Next), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Previous":
			FormDivSelectFieldToField(&(paragraph_.Previous), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingBody":
			FormDivSelectFieldToField(&(paragraph_.EnclosingBody), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingTableColumn":
			FormDivSelectFieldToField(&(paragraph_.EnclosingTableColumn), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Body:Paragraphs":
			// WARNING : this form deals with the N-N association "Body.Paragraphs []*Paragraph" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Paragraph). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Body
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Body"
				rf.Fieldname = "Paragraphs"
				formerAssociationSource := paragraph_.GongGetReverseFieldOwner(
					paragraphFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Body)
					if !ok {
						log.Fatalln("Source of Body.Paragraphs []*Paragraph, is not an Body instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Paragraphs, paragraph_)
					formerSource.Paragraphs = slices.Delete(formerSource.Paragraphs, idx, idx+1)
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
			var newSource *models.Body
			for _body := range *models.GetGongstructInstancesSet[models.Body](paragraphFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _body.GetName() == newSourceName.GetName() {
					newSource = _body // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Body.Paragraphs []*Paragraph, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Paragraphs = append(newSource.Paragraphs, paragraph_)
		case "TableColumn:Paragraphs":
			// WARNING : this form deals with the N-N association "TableColumn.Paragraphs []*Paragraph" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Paragraph). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.TableColumn
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "TableColumn"
				rf.Fieldname = "Paragraphs"
				formerAssociationSource := paragraph_.GongGetReverseFieldOwner(
					paragraphFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.TableColumn)
					if !ok {
						log.Fatalln("Source of TableColumn.Paragraphs []*Paragraph, is not an TableColumn instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Paragraphs, paragraph_)
					formerSource.Paragraphs = slices.Delete(formerSource.Paragraphs, idx, idx+1)
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
			var newSource *models.TableColumn
			for _tablecolumn := range *models.GetGongstructInstancesSet[models.TableColumn](paragraphFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _tablecolumn.GetName() == newSourceName.GetName() {
					newSource = _tablecolumn // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of TableColumn.Paragraphs []*Paragraph, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Paragraphs = append(newSource.Paragraphs, paragraph_)
		}
	}

	// manage the suppress operation
	if paragraphFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraph_.Unstage(paragraphFormCallback.probe.stageOfInterest)
	}

	paragraphFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Paragraph](
		paragraphFormCallback.probe,
	)

	// display a new form by reset the form stage
	if paragraphFormCallback.CreationMode || paragraphFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraphFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(paragraphFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParagraphFormCallback(
			nil,
			paragraphFormCallback.probe,
			newFormGroup,
		)
		paragraph := new(models.Paragraph)
		FillUpForm(paragraph, newFormGroup, paragraphFormCallback.probe)
		paragraphFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(paragraphFormCallback.probe)
}
func __gong__New__ParagraphPropertiesFormCallback(
	paragraphproperties *models.ParagraphProperties,
	probe *Probe,
	formGroup *table.FormGroup,
) (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) {
	paragraphpropertiesFormCallback = new(ParagraphPropertiesFormCallback)
	paragraphpropertiesFormCallback.probe = probe
	paragraphpropertiesFormCallback.paragraphproperties = paragraphproperties
	paragraphpropertiesFormCallback.formGroup = formGroup

	paragraphpropertiesFormCallback.CreationMode = (paragraphproperties == nil)

	return
}

type ParagraphPropertiesFormCallback struct {
	paragraphproperties *models.ParagraphProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) OnSave() {
	paragraphpropertiesFormCallback.probe.stageOfInterest.Lock()
	defer paragraphpropertiesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParagraphPropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphpropertiesFormCallback.probe.formStage.Checkout()

	if paragraphpropertiesFormCallback.paragraphproperties == nil {
		paragraphpropertiesFormCallback.paragraphproperties = new(models.ParagraphProperties).Stage(paragraphpropertiesFormCallback.probe.stageOfInterest)
	}
	paragraphproperties_ := paragraphpropertiesFormCallback.paragraphproperties
	_ = paragraphproperties_

	for _, formDiv := range paragraphpropertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphproperties_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphproperties_.Content), formDiv)
		case "ParagraphStyle":
			FormDivSelectFieldToField(&(paragraphproperties_.ParagraphStyle), paragraphpropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphproperties_.Node), paragraphpropertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if paragraphpropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraphproperties_.Unstage(paragraphpropertiesFormCallback.probe.stageOfInterest)
	}

	paragraphpropertiesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParagraphProperties](
		paragraphpropertiesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if paragraphpropertiesFormCallback.CreationMode || paragraphpropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraphpropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(paragraphpropertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParagraphPropertiesFormCallback(
			nil,
			paragraphpropertiesFormCallback.probe,
			newFormGroup,
		)
		paragraphproperties := new(models.ParagraphProperties)
		FillUpForm(paragraphproperties, newFormGroup, paragraphpropertiesFormCallback.probe)
		paragraphpropertiesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(paragraphpropertiesFormCallback.probe)
}
func __gong__New__ParagraphStyleFormCallback(
	paragraphstyle *models.ParagraphStyle,
	probe *Probe,
	formGroup *table.FormGroup,
) (paragraphstyleFormCallback *ParagraphStyleFormCallback) {
	paragraphstyleFormCallback = new(ParagraphStyleFormCallback)
	paragraphstyleFormCallback.probe = probe
	paragraphstyleFormCallback.paragraphstyle = paragraphstyle
	paragraphstyleFormCallback.formGroup = formGroup

	paragraphstyleFormCallback.CreationMode = (paragraphstyle == nil)

	return
}

type ParagraphStyleFormCallback struct {
	paragraphstyle *models.ParagraphStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (paragraphstyleFormCallback *ParagraphStyleFormCallback) OnSave() {
	paragraphstyleFormCallback.probe.stageOfInterest.Lock()
	defer paragraphstyleFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParagraphStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphstyleFormCallback.probe.formStage.Checkout()

	if paragraphstyleFormCallback.paragraphstyle == nil {
		paragraphstyleFormCallback.paragraphstyle = new(models.ParagraphStyle).Stage(paragraphstyleFormCallback.probe.stageOfInterest)
	}
	paragraphstyle_ := paragraphstyleFormCallback.paragraphstyle
	_ = paragraphstyle_

	for _, formDiv := range paragraphstyleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphstyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphstyle_.Node), paragraphstyleFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphstyle_.Content), formDiv)
		case "ValAttr":
			FormDivBasicFieldToField(&(paragraphstyle_.ValAttr), formDiv)
		}
	}

	// manage the suppress operation
	if paragraphstyleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraphstyle_.Unstage(paragraphstyleFormCallback.probe.stageOfInterest)
	}

	paragraphstyleFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParagraphStyle](
		paragraphstyleFormCallback.probe,
	)

	// display a new form by reset the form stage
	if paragraphstyleFormCallback.CreationMode || paragraphstyleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		paragraphstyleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(paragraphstyleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParagraphStyleFormCallback(
			nil,
			paragraphstyleFormCallback.probe,
			newFormGroup,
		)
		paragraphstyle := new(models.ParagraphStyle)
		FillUpForm(paragraphstyle, newFormGroup, paragraphstyleFormCallback.probe)
		paragraphstyleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(paragraphstyleFormCallback.probe)
}
func __gong__New__RuneFormCallback(
	rune *models.Rune,
	probe *Probe,
	formGroup *table.FormGroup,
) (runeFormCallback *RuneFormCallback) {
	runeFormCallback = new(RuneFormCallback)
	runeFormCallback.probe = probe
	runeFormCallback.rune = rune
	runeFormCallback.formGroup = formGroup

	runeFormCallback.CreationMode = (rune == nil)

	return
}

type RuneFormCallback struct {
	rune *models.Rune

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (runeFormCallback *RuneFormCallback) OnSave() {
	runeFormCallback.probe.stageOfInterest.Lock()
	defer runeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RuneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runeFormCallback.probe.formStage.Checkout()

	if runeFormCallback.rune == nil {
		runeFormCallback.rune = new(models.Rune).Stage(runeFormCallback.probe.stageOfInterest)
	}
	rune_ := runeFormCallback.rune
	_ = rune_

	for _, formDiv := range runeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rune_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rune_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(rune_.Node), runeFormCallback.probe.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(rune_.Text), runeFormCallback.probe.stageOfInterest, formDiv)
		case "RuneProperties":
			FormDivSelectFieldToField(&(rune_.RuneProperties), runeFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingParagraph":
			FormDivSelectFieldToField(&(rune_.EnclosingParagraph), runeFormCallback.probe.stageOfInterest, formDiv)
		case "Paragraph:Runes":
			// WARNING : this form deals with the N-N association "Paragraph.Runes []*Rune" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Rune). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Paragraph
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Paragraph"
				rf.Fieldname = "Runes"
				formerAssociationSource := rune_.GongGetReverseFieldOwner(
					runeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Paragraph)
					if !ok {
						log.Fatalln("Source of Paragraph.Runes []*Rune, is not an Paragraph instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Runes, rune_)
					formerSource.Runes = slices.Delete(formerSource.Runes, idx, idx+1)
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
			var newSource *models.Paragraph
			for _paragraph := range *models.GetGongstructInstancesSet[models.Paragraph](runeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _paragraph.GetName() == newSourceName.GetName() {
					newSource = _paragraph // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Paragraph.Runes []*Rune, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Runes = append(newSource.Runes, rune_)
		}
	}

	// manage the suppress operation
	if runeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rune_.Unstage(runeFormCallback.probe.stageOfInterest)
	}

	runeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Rune](
		runeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if runeFormCallback.CreationMode || runeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		runeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(runeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RuneFormCallback(
			nil,
			runeFormCallback.probe,
			newFormGroup,
		)
		rune := new(models.Rune)
		FillUpForm(rune, newFormGroup, runeFormCallback.probe)
		runeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(runeFormCallback.probe)
}
func __gong__New__RunePropertiesFormCallback(
	runeproperties *models.RuneProperties,
	probe *Probe,
	formGroup *table.FormGroup,
) (runepropertiesFormCallback *RunePropertiesFormCallback) {
	runepropertiesFormCallback = new(RunePropertiesFormCallback)
	runepropertiesFormCallback.probe = probe
	runepropertiesFormCallback.runeproperties = runeproperties
	runepropertiesFormCallback.formGroup = formGroup

	runepropertiesFormCallback.CreationMode = (runeproperties == nil)

	return
}

type RunePropertiesFormCallback struct {
	runeproperties *models.RuneProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (runepropertiesFormCallback *RunePropertiesFormCallback) OnSave() {
	runepropertiesFormCallback.probe.stageOfInterest.Lock()
	defer runepropertiesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RunePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runepropertiesFormCallback.probe.formStage.Checkout()

	if runepropertiesFormCallback.runeproperties == nil {
		runepropertiesFormCallback.runeproperties = new(models.RuneProperties).Stage(runepropertiesFormCallback.probe.stageOfInterest)
	}
	runeproperties_ := runepropertiesFormCallback.runeproperties
	_ = runeproperties_

	for _, formDiv := range runepropertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(runeproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(runeproperties_.Node), runepropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "IsBold":
			FormDivBasicFieldToField(&(runeproperties_.IsBold), formDiv)
		case "IsStrike":
			FormDivBasicFieldToField(&(runeproperties_.IsStrike), formDiv)
		case "IsItalic":
			FormDivBasicFieldToField(&(runeproperties_.IsItalic), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(runeproperties_.Content), formDiv)
		}
	}

	// manage the suppress operation
	if runepropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		runeproperties_.Unstage(runepropertiesFormCallback.probe.stageOfInterest)
	}

	runepropertiesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RuneProperties](
		runepropertiesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if runepropertiesFormCallback.CreationMode || runepropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		runepropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(runepropertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RunePropertiesFormCallback(
			nil,
			runepropertiesFormCallback.probe,
			newFormGroup,
		)
		runeproperties := new(models.RuneProperties)
		FillUpForm(runeproperties, newFormGroup, runepropertiesFormCallback.probe)
		runepropertiesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(runepropertiesFormCallback.probe)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
	formGroup *table.FormGroup,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.probe = probe
	tableFormCallback.table = table
	tableFormCallback.formGroup = formGroup

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tableFormCallback *TableFormCallback) OnSave() {
	tableFormCallback.probe.stageOfInterest.Lock()
	defer tableFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.probe.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.probe.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	for _, formDiv := range tableFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(table_.Node), tableFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(table_.Content), formDiv)
		case "TableProperties":
			FormDivSelectFieldToField(&(table_.TableProperties), tableFormCallback.probe.stageOfInterest, formDiv)
		case "TableRows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TableRow](tableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TableRow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TableRow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TableRow](tableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			table_.TableRows = instanceSlice

		case "Body:Tables":
			// WARNING : this form deals with the N-N association "Body.Tables []*Table" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Table). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Body
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Body"
				rf.Fieldname = "Tables"
				formerAssociationSource := table_.GongGetReverseFieldOwner(
					tableFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Body)
					if !ok {
						log.Fatalln("Source of Body.Tables []*Table, is not an Body instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Tables, table_)
					formerSource.Tables = slices.Delete(formerSource.Tables, idx, idx+1)
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
			var newSource *models.Body
			for _body := range *models.GetGongstructInstancesSet[models.Body](tableFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _body.GetName() == newSourceName.GetName() {
					newSource = _body // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Body.Tables []*Table, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Tables = append(newSource.Tables, table_)
		}
	}

	// manage the suppress operation
	if tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		table_.Unstage(tableFormCallback.probe.stageOfInterest)
	}

	tableFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Table](
		tableFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode || tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tableFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tableFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			tableFormCallback.probe,
			newFormGroup,
		)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.probe)
		tableFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tableFormCallback.probe)
}
func __gong__New__TableColumnFormCallback(
	tablecolumn *models.TableColumn,
	probe *Probe,
	formGroup *table.FormGroup,
) (tablecolumnFormCallback *TableColumnFormCallback) {
	tablecolumnFormCallback = new(TableColumnFormCallback)
	tablecolumnFormCallback.probe = probe
	tablecolumnFormCallback.tablecolumn = tablecolumn
	tablecolumnFormCallback.formGroup = formGroup

	tablecolumnFormCallback.CreationMode = (tablecolumn == nil)

	return
}

type TableColumnFormCallback struct {
	tablecolumn *models.TableColumn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tablecolumnFormCallback *TableColumnFormCallback) OnSave() {
	tablecolumnFormCallback.probe.stageOfInterest.Lock()
	defer tablecolumnFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TableColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablecolumnFormCallback.probe.formStage.Checkout()

	if tablecolumnFormCallback.tablecolumn == nil {
		tablecolumnFormCallback.tablecolumn = new(models.TableColumn).Stage(tablecolumnFormCallback.probe.stageOfInterest)
	}
	tablecolumn_ := tablecolumnFormCallback.tablecolumn
	_ = tablecolumn_

	for _, formDiv := range tablecolumnFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablecolumn_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablecolumn_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablecolumn_.Node), tablecolumnFormCallback.probe.stageOfInterest, formDiv)
		case "Paragraphs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Paragraph](tablecolumnFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Paragraph, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Paragraph)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tablecolumnFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Paragraph](tablecolumnFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			tablecolumn_.Paragraphs = instanceSlice

		case "TableRow:TableColumns":
			// WARNING : this form deals with the N-N association "TableRow.TableColumns []*TableColumn" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TableColumn). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.TableRow
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "TableRow"
				rf.Fieldname = "TableColumns"
				formerAssociationSource := tablecolumn_.GongGetReverseFieldOwner(
					tablecolumnFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.TableRow)
					if !ok {
						log.Fatalln("Source of TableRow.TableColumns []*TableColumn, is not an TableRow instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TableColumns, tablecolumn_)
					formerSource.TableColumns = slices.Delete(formerSource.TableColumns, idx, idx+1)
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
			var newSource *models.TableRow
			for _tablerow := range *models.GetGongstructInstancesSet[models.TableRow](tablecolumnFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _tablerow.GetName() == newSourceName.GetName() {
					newSource = _tablerow // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of TableRow.TableColumns []*TableColumn, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TableColumns = append(newSource.TableColumns, tablecolumn_)
		}
	}

	// manage the suppress operation
	if tablecolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablecolumn_.Unstage(tablecolumnFormCallback.probe.stageOfInterest)
	}

	tablecolumnFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TableColumn](
		tablecolumnFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tablecolumnFormCallback.CreationMode || tablecolumnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablecolumnFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tablecolumnFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableColumnFormCallback(
			nil,
			tablecolumnFormCallback.probe,
			newFormGroup,
		)
		tablecolumn := new(models.TableColumn)
		FillUpForm(tablecolumn, newFormGroup, tablecolumnFormCallback.probe)
		tablecolumnFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tablecolumnFormCallback.probe)
}
func __gong__New__TablePropertiesFormCallback(
	tableproperties *models.TableProperties,
	probe *Probe,
	formGroup *table.FormGroup,
) (tablepropertiesFormCallback *TablePropertiesFormCallback) {
	tablepropertiesFormCallback = new(TablePropertiesFormCallback)
	tablepropertiesFormCallback.probe = probe
	tablepropertiesFormCallback.tableproperties = tableproperties
	tablepropertiesFormCallback.formGroup = formGroup

	tablepropertiesFormCallback.CreationMode = (tableproperties == nil)

	return
}

type TablePropertiesFormCallback struct {
	tableproperties *models.TableProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tablepropertiesFormCallback *TablePropertiesFormCallback) OnSave() {
	tablepropertiesFormCallback.probe.stageOfInterest.Lock()
	defer tablepropertiesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TablePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablepropertiesFormCallback.probe.formStage.Checkout()

	if tablepropertiesFormCallback.tableproperties == nil {
		tablepropertiesFormCallback.tableproperties = new(models.TableProperties).Stage(tablepropertiesFormCallback.probe.stageOfInterest)
	}
	tableproperties_ := tablepropertiesFormCallback.tableproperties
	_ = tableproperties_

	for _, formDiv := range tablepropertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tableproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tableproperties_.Node), tablepropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tableproperties_.Content), formDiv)
		case "TableStyle":
			FormDivSelectFieldToField(&(tableproperties_.TableStyle), tablepropertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if tablepropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tableproperties_.Unstage(tablepropertiesFormCallback.probe.stageOfInterest)
	}

	tablepropertiesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TableProperties](
		tablepropertiesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tablepropertiesFormCallback.CreationMode || tablepropertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablepropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tablepropertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TablePropertiesFormCallback(
			nil,
			tablepropertiesFormCallback.probe,
			newFormGroup,
		)
		tableproperties := new(models.TableProperties)
		FillUpForm(tableproperties, newFormGroup, tablepropertiesFormCallback.probe)
		tablepropertiesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tablepropertiesFormCallback.probe)
}
func __gong__New__TableRowFormCallback(
	tablerow *models.TableRow,
	probe *Probe,
	formGroup *table.FormGroup,
) (tablerowFormCallback *TableRowFormCallback) {
	tablerowFormCallback = new(TableRowFormCallback)
	tablerowFormCallback.probe = probe
	tablerowFormCallback.tablerow = tablerow
	tablerowFormCallback.formGroup = formGroup

	tablerowFormCallback.CreationMode = (tablerow == nil)

	return
}

type TableRowFormCallback struct {
	tablerow *models.TableRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tablerowFormCallback *TableRowFormCallback) OnSave() {
	tablerowFormCallback.probe.stageOfInterest.Lock()
	defer tablerowFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TableRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablerowFormCallback.probe.formStage.Checkout()

	if tablerowFormCallback.tablerow == nil {
		tablerowFormCallback.tablerow = new(models.TableRow).Stage(tablerowFormCallback.probe.stageOfInterest)
	}
	tablerow_ := tablerowFormCallback.tablerow
	_ = tablerow_

	for _, formDiv := range tablerowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablerow_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablerow_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablerow_.Node), tablerowFormCallback.probe.stageOfInterest, formDiv)
		case "TableColumns":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TableColumn](tablerowFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TableColumn, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TableColumn)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					tablerowFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TableColumn](tablerowFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			tablerow_.TableColumns = instanceSlice

		case "Table:TableRows":
			// WARNING : this form deals with the N-N association "Table.TableRows []*TableRow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TableRow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Table
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Table"
				rf.Fieldname = "TableRows"
				formerAssociationSource := tablerow_.GongGetReverseFieldOwner(
					tablerowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Table)
					if !ok {
						log.Fatalln("Source of Table.TableRows []*TableRow, is not an Table instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TableRows, tablerow_)
					formerSource.TableRows = slices.Delete(formerSource.TableRows, idx, idx+1)
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
			var newSource *models.Table
			for _table := range *models.GetGongstructInstancesSet[models.Table](tablerowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _table.GetName() == newSourceName.GetName() {
					newSource = _table // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Table.TableRows []*TableRow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TableRows = append(newSource.TableRows, tablerow_)
		}
	}

	// manage the suppress operation
	if tablerowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablerow_.Unstage(tablerowFormCallback.probe.stageOfInterest)
	}

	tablerowFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TableRow](
		tablerowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tablerowFormCallback.CreationMode || tablerowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablerowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tablerowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableRowFormCallback(
			nil,
			tablerowFormCallback.probe,
			newFormGroup,
		)
		tablerow := new(models.TableRow)
		FillUpForm(tablerow, newFormGroup, tablerowFormCallback.probe)
		tablerowFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tablerowFormCallback.probe)
}
func __gong__New__TableStyleFormCallback(
	tablestyle *models.TableStyle,
	probe *Probe,
	formGroup *table.FormGroup,
) (tablestyleFormCallback *TableStyleFormCallback) {
	tablestyleFormCallback = new(TableStyleFormCallback)
	tablestyleFormCallback.probe = probe
	tablestyleFormCallback.tablestyle = tablestyle
	tablestyleFormCallback.formGroup = formGroup

	tablestyleFormCallback.CreationMode = (tablestyle == nil)

	return
}

type TableStyleFormCallback struct {
	tablestyle *models.TableStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tablestyleFormCallback *TableStyleFormCallback) OnSave() {
	tablestyleFormCallback.probe.stageOfInterest.Lock()
	defer tablestyleFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TableStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablestyleFormCallback.probe.formStage.Checkout()

	if tablestyleFormCallback.tablestyle == nil {
		tablestyleFormCallback.tablestyle = new(models.TableStyle).Stage(tablestyleFormCallback.probe.stageOfInterest)
	}
	tablestyle_ := tablestyleFormCallback.tablestyle
	_ = tablestyle_

	for _, formDiv := range tablestyleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablestyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablestyle_.Node), tablestyleFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablestyle_.Content), formDiv)
		case "Val":
			FormDivBasicFieldToField(&(tablestyle_.Val), formDiv)
		}
	}

	// manage the suppress operation
	if tablestyleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablestyle_.Unstage(tablestyleFormCallback.probe.stageOfInterest)
	}

	tablestyleFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TableStyle](
		tablestyleFormCallback.probe,
	)

	// display a new form by reset the form stage
	if tablestyleFormCallback.CreationMode || tablestyleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tablestyleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tablestyleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableStyleFormCallback(
			nil,
			tablestyleFormCallback.probe,
			newFormGroup,
		)
		tablestyle := new(models.TableStyle)
		FillUpForm(tablestyle, newFormGroup, tablestyleFormCallback.probe)
		tablestyleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tablestyleFormCallback.probe)
}
func __gong__New__TextFormCallback(
	text *models.Text,
	probe *Probe,
	formGroup *table.FormGroup,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.probe = probe
	textFormCallback.text = text
	textFormCallback.formGroup = formGroup

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (textFormCallback *TextFormCallback) OnSave() {
	textFormCallback.probe.stageOfInterest.Lock()
	defer textFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.probe.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.probe.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	for _, formDiv := range textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(text_.Node), textFormCallback.probe.stageOfInterest, formDiv)
		case "PreserveWhiteSpace":
			FormDivBasicFieldToField(&(text_.PreserveWhiteSpace), formDiv)
		case "EnclosingRune":
			FormDivSelectFieldToField(&(text_.EnclosingRune), textFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		text_.Unstage(textFormCallback.probe.stageOfInterest)
	}

	textFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Text](
		textFormCallback.probe,
	)

	// display a new form by reset the form stage
	if textFormCallback.CreationMode || textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TextFormCallback(
			nil,
			textFormCallback.probe,
			newFormGroup,
		)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.probe)
		textFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(textFormCallback.probe)
}
