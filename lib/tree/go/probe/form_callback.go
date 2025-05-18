// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__ButtonFormCallback(
	button *models.Button,
	probe *Probe,
	formGroup *table.FormGroup,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.probe = probe
	buttonFormCallback.button = button
	buttonFormCallback.formGroup = formGroup

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	// log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.probe.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.probe.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	for _, formDiv := range buttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(button_.Icon), formDiv)
		case "SVGIcon":
			FormDivSelectFieldToField(&(button_.SVGIcon), buttonFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(button_.IsDisabled), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(button_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(button_.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(button_.ToolTipPosition), formDiv)
		case "Node:Buttons":
			// WARNING : this form deals with the N-N association "Node.Buttons []*Button" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Button). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Node
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Node"
				rf.Fieldname = "Buttons"
				formerAssociationSource := models.GetReverseFieldOwner(
					buttonFormCallback.probe.stageOfInterest,
					button_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Node)
					if !ok {
						log.Fatalln("Source of Node.Buttons []*Button, is not an Node instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Buttons, button_)
					formerSource.Buttons = slices.Delete(formerSource.Buttons, idx, idx+1)
				}
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
			var newSource *models.Node
			for _node := range *models.GetGongstructInstancesSet[models.Node](buttonFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _node.GetName() == fieldValue.GetName() {
							newNodeOwner := _node // we have a match
							
							// we remove the button_ instance from the pastNodeOwner field
							if pastNodeOwner != nil {
								if newNodeOwner != pastNodeOwner {
									idx := slices.Index(pastNodeOwner.Buttons, button_)
									pastNodeOwner.Buttons = slices.Delete(pastNodeOwner.Buttons, idx, idx+1)
									newNodeOwner.Buttons = append(newNodeOwner.Buttons, button_)
								}
							} else {
								newNodeOwner.Buttons = append(newNodeOwner.Buttons, button_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		button_.Unstage(buttonFormCallback.probe.stageOfInterest)
	}

	buttonFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Button](
		buttonFormCallback.probe,
	)
	buttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(buttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			buttonFormCallback.probe,
			newFormGroup,
		)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.probe)
		buttonFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(buttonFormCallback.probe)
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
		case "FontStyle":
			FormDivEnumStringFieldToField(&(node_.FontStyle), formDiv)
		case "BackgroundColor":
			FormDivBasicFieldToField(&(node_.BackgroundColor), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(node_.IsExpanded), formDiv)
		case "HasCheckboxButton":
			FormDivBasicFieldToField(&(node_.HasCheckboxButton), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(node_.IsChecked), formDiv)
		case "IsCheckboxDisabled":
			FormDivBasicFieldToField(&(node_.IsCheckboxDisabled), formDiv)
		case "CheckboxHasToolTip":
			FormDivBasicFieldToField(&(node_.CheckboxHasToolTip), formDiv)
		case "CheckboxToolTipText":
			FormDivBasicFieldToField(&(node_.CheckboxToolTipText), formDiv)
		case "CheckboxToolTipPosition":
			FormDivEnumStringFieldToField(&(node_.CheckboxToolTipPosition), formDiv)
		case "HasSecondCheckboxButton":
			FormDivBasicFieldToField(&(node_.HasSecondCheckboxButton), formDiv)
		case "IsSecondCheckboxChecked":
			FormDivBasicFieldToField(&(node_.IsSecondCheckboxChecked), formDiv)
		case "IsSecondCheckboxDisabled":
			FormDivBasicFieldToField(&(node_.IsSecondCheckboxDisabled), formDiv)
		case "TextAfterSecondCheckbox":
			FormDivBasicFieldToField(&(node_.TextAfterSecondCheckbox), formDiv)
		case "SecondCheckboxHasToolTip":
			FormDivBasicFieldToField(&(node_.SecondCheckboxHasToolTip), formDiv)
		case "SecondCheckboxToolTipText":
			FormDivBasicFieldToField(&(node_.SecondCheckboxToolTipText), formDiv)
		case "SecondCheckboxToolTipPosition":
			FormDivEnumStringFieldToField(&(node_.SecondCheckboxToolTipPosition), formDiv)
		case "IsInEditMode":
			FormDivBasicFieldToField(&(node_.IsInEditMode), formDiv)
		case "IsNodeClickable":
			FormDivBasicFieldToField(&(node_.IsNodeClickable), formDiv)
		case "IsWithPreceedingIcon":
			FormDivBasicFieldToField(&(node_.IsWithPreceedingIcon), formDiv)
		case "PreceedingIcon":
			FormDivBasicFieldToField(&(node_.PreceedingIcon), formDiv)
		case "PreceedingSVGIcon":
			FormDivSelectFieldToField(&(node_.PreceedingSVGIcon), nodeFormCallback.probe.stageOfInterest, formDiv)
	case "Children":
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			node_.Children = instanceSlice

	case "Buttons":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Button](nodeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Button, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Button)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					nodeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			node_.Buttons = instanceSlice

		case "Node:Children":
			// WARNING : this form deals with the N-N association "Node.Children []*Node" but
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
				rf.Fieldname = "Children"
				formerAssociationSource := models.GetReverseFieldOwner(
					nodeFormCallback.probe.stageOfInterest,
					node_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Node)
					if !ok {
						log.Fatalln("Source of Node.Children []*Node, is not an Node instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Children, node_)
					formerSource.Children = slices.Delete(formerSource.Children, idx, idx+1)
				}
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
			var newSource *models.Node
			for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _node.GetName() == fieldValue.GetName() {
							newNodeOwner := _node // we have a match
							
							// we remove the node_ instance from the pastNodeOwner field
							if pastNodeOwner != nil {
								if newNodeOwner != pastNodeOwner {
									idx := slices.Index(pastNodeOwner.Children, node_)
									pastNodeOwner.Children = slices.Delete(pastNodeOwner.Children, idx, idx+1)
									newNodeOwner.Children = append(newNodeOwner.Children, node_)
								}
							} else {
								newNodeOwner.Children = append(newNodeOwner.Children, node_)
							}
						}
					}
				}
			}
		case "Tree:RootNodes":
			// WARNING : this form deals with the N-N association "Tree.RootNodes []*Node" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Node). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Tree
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Tree"
				rf.Fieldname = "RootNodes"
				formerAssociationSource := models.GetReverseFieldOwner(
					nodeFormCallback.probe.stageOfInterest,
					node_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Tree)
					if !ok {
						log.Fatalln("Source of Tree.RootNodes []*Node, is not an Tree instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.RootNodes, node_)
					formerSource.RootNodes = slices.Delete(formerSource.RootNodes, idx, idx+1)
				}
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
			var newSource *models.Tree
			for _tree := range *models.GetGongstructInstancesSet[models.Tree](nodeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _tree.GetName() == fieldValue.GetName() {
							newTreeOwner := _tree // we have a match
							
							// we remove the node_ instance from the pastTreeOwner field
							if pastTreeOwner != nil {
								if newTreeOwner != pastTreeOwner {
									idx := slices.Index(pastTreeOwner.RootNodes, node_)
									pastTreeOwner.RootNodes = slices.Delete(pastTreeOwner.RootNodes, idx, idx+1)
									newTreeOwner.RootNodes = append(newTreeOwner.RootNodes, node_)
								}
							} else {
								newTreeOwner.RootNodes = append(newTreeOwner.RootNodes, node_)
							}
						}
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
	updateAndCommitTable[models.Node](
		nodeFormCallback.probe,
	)
	nodeFormCallback.probe.tableStage.Commit()

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
func __gong__New__SVGIconFormCallback(
	svgicon *models.SVGIcon,
	probe *Probe,
	formGroup *table.FormGroup,
) (svgiconFormCallback *SVGIconFormCallback) {
	svgiconFormCallback = new(SVGIconFormCallback)
	svgiconFormCallback.probe = probe
	svgiconFormCallback.svgicon = svgicon
	svgiconFormCallback.formGroup = formGroup

	svgiconFormCallback.CreationMode = (svgicon == nil)

	return
}

type SVGIconFormCallback struct {
	svgicon *models.SVGIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (svgiconFormCallback *SVGIconFormCallback) OnSave() {

	// log.Println("SVGIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgiconFormCallback.probe.formStage.Checkout()

	if svgiconFormCallback.svgicon == nil {
		svgiconFormCallback.svgicon = new(models.SVGIcon).Stage(svgiconFormCallback.probe.stageOfInterest)
	}
	svgicon_ := svgiconFormCallback.svgicon
	_ = svgicon_

	for _, formDiv := range svgiconFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(svgicon_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if svgiconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgicon_.Unstage(svgiconFormCallback.probe.stageOfInterest)
	}

	svgiconFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SVGIcon](
		svgiconFormCallback.probe,
	)
	svgiconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgiconFormCallback.CreationMode || svgiconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgiconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(svgiconFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SVGIconFormCallback(
			nil,
			svgiconFormCallback.probe,
			newFormGroup,
		)
		svgicon := new(models.SVGIcon)
		FillUpForm(svgicon, newFormGroup, svgiconFormCallback.probe)
		svgiconFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(svgiconFormCallback.probe)
}
func __gong__New__TreeFormCallback(
	tree *models.Tree,
	probe *Probe,
	formGroup *table.FormGroup,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.probe = probe
	treeFormCallback.tree = tree
	treeFormCallback.formGroup = formGroup

	treeFormCallback.CreationMode = (tree == nil)

	return
}

type TreeFormCallback struct {
	tree *models.Tree

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (treeFormCallback *TreeFormCallback) OnSave() {

	// log.Println("TreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	treeFormCallback.probe.formStage.Checkout()

	if treeFormCallback.tree == nil {
		treeFormCallback.tree = new(models.Tree).Stage(treeFormCallback.probe.stageOfInterest)
	}
	tree_ := treeFormCallback.tree
	_ = tree_

	for _, formDiv := range treeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tree_.Name), formDiv)
	case "RootNodes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Node](treeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Node, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Node)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					treeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			tree_.RootNodes = instanceSlice

		}
	}

	// manage the suppress operation
	if treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tree_.Unstage(treeFormCallback.probe.stageOfInterest)
	}

	treeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Tree](
		treeFormCallback.probe,
	)
	treeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if treeFormCallback.CreationMode || treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		treeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(treeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TreeFormCallback(
			nil,
			treeFormCallback.probe,
			newFormGroup,
		)
		tree := new(models.Tree)
		FillUpForm(tree, newFormGroup, treeFormCallback.probe)
		treeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(treeFormCallback.probe)
}
