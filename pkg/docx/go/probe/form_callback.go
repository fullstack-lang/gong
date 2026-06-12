// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

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
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			bodyFormCallback.probe.UpdateSliceOfPointersCallback(body_, "Paragraphs", &body_.Paragraphs)

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
			bodyFormCallback.probe.UpdateSliceOfPointersCallback(body_, "Tables", &body_.Tables)

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
		newFormGroup := (&form.FormGroup{
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

	bodyFormCallback.probe.ux_tree()
}
func __gong__New__DocumentFormCallback(
	document *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	documentFormCallback.probe.ux_tree()
}
func __gong__New__DocxFormCallback(
	docx *models.Docx,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			docxFormCallback.probe.UpdateSliceOfPointersCallback(docx_, "Files", &docx_.Files)

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
		newFormGroup := (&form.FormGroup{
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

	docxFormCallback.probe.ux_tree()
}
func __gong__New__FileFormCallback(
	file *models.File,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Docx instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Docx instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Docx](fileFormCallback.probe.stageOfInterest)
			targetDocxIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDocxIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Docx instances and update their Files slice
			for _docx := range *models.GetGongstructInstancesSetFromPointerType[*models.Docx](fileFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(fileFormCallback.probe.stageOfInterest, _docx)
				
				// if Docx is selected
				if targetDocxIDs[id] {
					// ensure file_ is in _docx.Files
					found := false
					for _, _b := range _docx.Files {
						if _b == file_ {
							found = true
							break
						}
					}
					if !found {
						_docx.Files = append(_docx.Files, file_)
						fileFormCallback.probe.UpdateSliceOfPointersCallback(_docx, "Files", &_docx.Files)
					}
				} else {
					// ensure file_ is NOT in _docx.Files
					idx := slices.Index(_docx.Files, file_)
					if idx != -1 {
						_docx.Files = slices.Delete(_docx.Files, idx, idx+1)
						fileFormCallback.probe.UpdateSliceOfPointersCallback(_docx, "Files", &_docx.Files)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	fileFormCallback.probe.ux_tree()
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			nodeFormCallback.probe.UpdateSliceOfPointersCallback(node_, "Nodes", &node_.Nodes)

		case "Node:Nodes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Node instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Node instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Node](nodeFormCallback.probe.stageOfInterest)
			targetNodeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNodeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Node instances and update their Nodes slice
			for _node := range *models.GetGongstructInstancesSetFromPointerType[*models.Node](nodeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(nodeFormCallback.probe.stageOfInterest, _node)
				
				// if Node is selected
				if targetNodeIDs[id] {
					// ensure node_ is in _node.Nodes
					found := false
					for _, _b := range _node.Nodes {
						if _b == node_ {
							found = true
							break
						}
					}
					if !found {
						_node.Nodes = append(_node.Nodes, node_)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Nodes", &_node.Nodes)
					}
				} else {
					// ensure node_ is NOT in _node.Nodes
					idx := slices.Index(_node.Nodes, node_)
					if idx != -1 {
						_node.Nodes = slices.Delete(_node.Nodes, idx, idx+1)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Nodes", &_node.Nodes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	nodeFormCallback.probe.ux_tree()
}
func __gong__New__ParagraphFormCallback(
	paragraph *models.Paragraph,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			paragraphFormCallback.probe.UpdateSliceOfPointersCallback(paragraph_, "Runes", &paragraph_.Runes)

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Body instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Body instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Body](paragraphFormCallback.probe.stageOfInterest)
			targetBodyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBodyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Body instances and update their Paragraphs slice
			for _body := range *models.GetGongstructInstancesSetFromPointerType[*models.Body](paragraphFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(paragraphFormCallback.probe.stageOfInterest, _body)
				
				// if Body is selected
				if targetBodyIDs[id] {
					// ensure paragraph_ is in _body.Paragraphs
					found := false
					for _, _b := range _body.Paragraphs {
						if _b == paragraph_ {
							found = true
							break
						}
					}
					if !found {
						_body.Paragraphs = append(_body.Paragraphs, paragraph_)
						paragraphFormCallback.probe.UpdateSliceOfPointersCallback(_body, "Paragraphs", &_body.Paragraphs)
					}
				} else {
					// ensure paragraph_ is NOT in _body.Paragraphs
					idx := slices.Index(_body.Paragraphs, paragraph_)
					if idx != -1 {
						_body.Paragraphs = slices.Delete(_body.Paragraphs, idx, idx+1)
						paragraphFormCallback.probe.UpdateSliceOfPointersCallback(_body, "Paragraphs", &_body.Paragraphs)
					}
				}
			}
		case "TableColumn:Paragraphs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TableColumn instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TableColumn instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TableColumn](paragraphFormCallback.probe.stageOfInterest)
			targetTableColumnIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTableColumnIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TableColumn instances and update their Paragraphs slice
			for _tablecolumn := range *models.GetGongstructInstancesSetFromPointerType[*models.TableColumn](paragraphFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(paragraphFormCallback.probe.stageOfInterest, _tablecolumn)
				
				// if TableColumn is selected
				if targetTableColumnIDs[id] {
					// ensure paragraph_ is in _tablecolumn.Paragraphs
					found := false
					for _, _b := range _tablecolumn.Paragraphs {
						if _b == paragraph_ {
							found = true
							break
						}
					}
					if !found {
						_tablecolumn.Paragraphs = append(_tablecolumn.Paragraphs, paragraph_)
						paragraphFormCallback.probe.UpdateSliceOfPointersCallback(_tablecolumn, "Paragraphs", &_tablecolumn.Paragraphs)
					}
				} else {
					// ensure paragraph_ is NOT in _tablecolumn.Paragraphs
					idx := slices.Index(_tablecolumn.Paragraphs, paragraph_)
					if idx != -1 {
						_tablecolumn.Paragraphs = slices.Delete(_tablecolumn.Paragraphs, idx, idx+1)
						paragraphFormCallback.probe.UpdateSliceOfPointersCallback(_tablecolumn, "Paragraphs", &_tablecolumn.Paragraphs)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	paragraphFormCallback.probe.ux_tree()
}
func __gong__New__ParagraphPropertiesFormCallback(
	paragraphproperties *models.ParagraphProperties,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	paragraphpropertiesFormCallback.probe.ux_tree()
}
func __gong__New__ParagraphStyleFormCallback(
	paragraphstyle *models.ParagraphStyle,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	paragraphstyleFormCallback.probe.ux_tree()
}
func __gong__New__RuneFormCallback(
	rune *models.Rune,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Paragraph instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Paragraph instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Paragraph](runeFormCallback.probe.stageOfInterest)
			targetParagraphIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParagraphIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Paragraph instances and update their Runes slice
			for _paragraph := range *models.GetGongstructInstancesSetFromPointerType[*models.Paragraph](runeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(runeFormCallback.probe.stageOfInterest, _paragraph)
				
				// if Paragraph is selected
				if targetParagraphIDs[id] {
					// ensure rune_ is in _paragraph.Runes
					found := false
					for _, _b := range _paragraph.Runes {
						if _b == rune_ {
							found = true
							break
						}
					}
					if !found {
						_paragraph.Runes = append(_paragraph.Runes, rune_)
						runeFormCallback.probe.UpdateSliceOfPointersCallback(_paragraph, "Runes", &_paragraph.Runes)
					}
				} else {
					// ensure rune_ is NOT in _paragraph.Runes
					idx := slices.Index(_paragraph.Runes, rune_)
					if idx != -1 {
						_paragraph.Runes = slices.Delete(_paragraph.Runes, idx, idx+1)
						runeFormCallback.probe.UpdateSliceOfPointersCallback(_paragraph, "Runes", &_paragraph.Runes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	runeFormCallback.probe.ux_tree()
}
func __gong__New__RunePropertiesFormCallback(
	runeproperties *models.RuneProperties,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	runepropertiesFormCallback.probe.ux_tree()
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			tableFormCallback.probe.UpdateSliceOfPointersCallback(table_, "TableRows", &table_.TableRows)

		case "Body:Tables":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Body instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Body instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Body](tableFormCallback.probe.stageOfInterest)
			targetBodyIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBodyIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Body instances and update their Tables slice
			for _body := range *models.GetGongstructInstancesSetFromPointerType[*models.Body](tableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(tableFormCallback.probe.stageOfInterest, _body)
				
				// if Body is selected
				if targetBodyIDs[id] {
					// ensure table_ is in _body.Tables
					found := false
					for _, _b := range _body.Tables {
						if _b == table_ {
							found = true
							break
						}
					}
					if !found {
						_body.Tables = append(_body.Tables, table_)
						tableFormCallback.probe.UpdateSliceOfPointersCallback(_body, "Tables", &_body.Tables)
					}
				} else {
					// ensure table_ is NOT in _body.Tables
					idx := slices.Index(_body.Tables, table_)
					if idx != -1 {
						_body.Tables = slices.Delete(_body.Tables, idx, idx+1)
						tableFormCallback.probe.UpdateSliceOfPointersCallback(_body, "Tables", &_body.Tables)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	tableFormCallback.probe.ux_tree()
}
func __gong__New__TableColumnFormCallback(
	tablecolumn *models.TableColumn,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			tablecolumnFormCallback.probe.UpdateSliceOfPointersCallback(tablecolumn_, "Paragraphs", &tablecolumn_.Paragraphs)

		case "TableRow:TableColumns":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TableRow instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TableRow instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TableRow](tablecolumnFormCallback.probe.stageOfInterest)
			targetTableRowIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTableRowIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TableRow instances and update their TableColumns slice
			for _tablerow := range *models.GetGongstructInstancesSetFromPointerType[*models.TableRow](tablecolumnFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(tablecolumnFormCallback.probe.stageOfInterest, _tablerow)
				
				// if TableRow is selected
				if targetTableRowIDs[id] {
					// ensure tablecolumn_ is in _tablerow.TableColumns
					found := false
					for _, _b := range _tablerow.TableColumns {
						if _b == tablecolumn_ {
							found = true
							break
						}
					}
					if !found {
						_tablerow.TableColumns = append(_tablerow.TableColumns, tablecolumn_)
						tablecolumnFormCallback.probe.UpdateSliceOfPointersCallback(_tablerow, "TableColumns", &_tablerow.TableColumns)
					}
				} else {
					// ensure tablecolumn_ is NOT in _tablerow.TableColumns
					idx := slices.Index(_tablerow.TableColumns, tablecolumn_)
					if idx != -1 {
						_tablerow.TableColumns = slices.Delete(_tablerow.TableColumns, idx, idx+1)
						tablecolumnFormCallback.probe.UpdateSliceOfPointersCallback(_tablerow, "TableColumns", &_tablerow.TableColumns)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	tablecolumnFormCallback.probe.ux_tree()
}
func __gong__New__TablePropertiesFormCallback(
	tableproperties *models.TableProperties,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	tablepropertiesFormCallback.probe.ux_tree()
}
func __gong__New__TableRowFormCallback(
	tablerow *models.TableRow,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
			tablerowFormCallback.probe.UpdateSliceOfPointersCallback(tablerow_, "TableColumns", &tablerow_.TableColumns)

		case "Table:TableRows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Table instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Table instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Table](tablerowFormCallback.probe.stageOfInterest)
			targetTableIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTableIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Table instances and update their TableRows slice
			for _table := range *models.GetGongstructInstancesSetFromPointerType[*models.Table](tablerowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(tablerowFormCallback.probe.stageOfInterest, _table)
				
				// if Table is selected
				if targetTableIDs[id] {
					// ensure tablerow_ is in _table.TableRows
					found := false
					for _, _b := range _table.TableRows {
						if _b == tablerow_ {
							found = true
							break
						}
					}
					if !found {
						_table.TableRows = append(_table.TableRows, tablerow_)
						tablerowFormCallback.probe.UpdateSliceOfPointersCallback(_table, "TableRows", &_table.TableRows)
					}
				} else {
					// ensure tablerow_ is NOT in _table.TableRows
					idx := slices.Index(_table.TableRows, tablerow_)
					if idx != -1 {
						_table.TableRows = slices.Delete(_table.TableRows, idx, idx+1)
						tablerowFormCallback.probe.UpdateSliceOfPointersCallback(_table, "TableRows", &_table.TableRows)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	tablerowFormCallback.probe.ux_tree()
}
func __gong__New__TableStyleFormCallback(
	tablestyle *models.TableStyle,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	tablestyleFormCallback.probe.ux_tree()
}
func __gong__New__TextFormCallback(
	text *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
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
		newFormGroup := (&form.FormGroup{
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

	textFormCallback.probe.ux_tree()
}
