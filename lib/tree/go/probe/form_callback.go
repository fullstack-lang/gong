// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ButtonFormCallback(
	button *models.Button,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {
	buttonFormCallback.probe.stageOfInterest.Lock()
	defer buttonFormCallback.probe.stageOfInterest.Unlock()

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
		case "Menu:Buttons":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Menu instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Menu instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Menu](buttonFormCallback.probe.stageOfInterest)
			targetMenuIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetMenuIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Menu instances and update their Buttons slice
			for _menu := range *models.GetGongstructInstancesSetFromPointerType[*models.Menu](buttonFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(buttonFormCallback.probe.stageOfInterest, _menu)
				
				// if Menu is selected
				if targetMenuIDs[id] {
					// ensure button_ is in _menu.Buttons
					found := false
					for _, _b := range _menu.Buttons {
						if _b == button_ {
							found = true
							break
						}
					}
					if !found {
						_menu.Buttons = append(_menu.Buttons, button_)
						buttonFormCallback.probe.UpdateSliceOfPointersCallback(_menu, "Buttons", &_menu.Buttons)
					}
				} else {
					// ensure button_ is NOT in _menu.Buttons
					idx := slices.Index(_menu.Buttons, button_)
					if idx != -1 {
						_menu.Buttons = slices.Delete(_menu.Buttons, idx, idx+1)
						buttonFormCallback.probe.UpdateSliceOfPointersCallback(_menu, "Buttons", &_menu.Buttons)
					}
				}
			}
		case "Node:Buttons":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Node instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Node instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Node](buttonFormCallback.probe.stageOfInterest)
			targetNodeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNodeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Node instances and update their Buttons slice
			for _node := range *models.GetGongstructInstancesSetFromPointerType[*models.Node](buttonFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(buttonFormCallback.probe.stageOfInterest, _node)
				
				// if Node is selected
				if targetNodeIDs[id] {
					// ensure button_ is in _node.Buttons
					found := false
					for _, _b := range _node.Buttons {
						if _b == button_ {
							found = true
							break
						}
					}
					if !found {
						_node.Buttons = append(_node.Buttons, button_)
						buttonFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Buttons", &_node.Buttons)
					}
				} else {
					// ensure button_ is NOT in _node.Buttons
					idx := slices.Index(_node.Buttons, button_)
					if idx != -1 {
						_node.Buttons = slices.Delete(_node.Buttons, idx, idx+1)
						buttonFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Buttons", &_node.Buttons)
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
	updateProbeTable[*models.Button](
		buttonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	buttonFormCallback.probe.ux_tree()
}
func __gong__New__MenuFormCallback(
	menu *models.Menu,
	probe *Probe,
	formGroup *form.FormGroup,
) (menuFormCallback *MenuFormCallback) {
	menuFormCallback = new(MenuFormCallback)
	menuFormCallback.probe = probe
	menuFormCallback.menu = menu
	menuFormCallback.formGroup = formGroup

	menuFormCallback.CreationMode = (menu == nil)

	return
}

type MenuFormCallback struct {
	menu *models.Menu

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (menuFormCallback *MenuFormCallback) OnSave() {
	menuFormCallback.probe.stageOfInterest.Lock()
	defer menuFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MenuFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	menuFormCallback.probe.formStage.Checkout()

	if menuFormCallback.menu == nil {
		menuFormCallback.menu = new(models.Menu).Stage(menuFormCallback.probe.stageOfInterest)
	}
	menu_ := menuFormCallback.menu
	_ = menu_

	for _, formDiv := range menuFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(menu_.Name), formDiv)
		case "Buttons":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Button](menuFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Button, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Button)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					menuFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Button](menuFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			menu_.Buttons = instanceSlice
			menuFormCallback.probe.UpdateSliceOfPointersCallback(menu_, "Buttons", &menu_.Buttons)

		}
	}

	// manage the suppress operation
	if menuFormCallback.formGroup.HasSuppressButtonBeenPressed {
		menu_.Unstage(menuFormCallback.probe.stageOfInterest)
	}

	menuFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Menu](
		menuFormCallback.probe,
	)

	// display a new form by reset the form stage
	if menuFormCallback.CreationMode || menuFormCallback.formGroup.HasSuppressButtonBeenPressed {
		menuFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(menuFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MenuFormCallback(
			nil,
			menuFormCallback.probe,
			newFormGroup,
		)
		menu := new(models.Menu)
		FillUpForm(menu, newFormGroup, menuFormCallback.probe)
		menuFormCallback.probe.formStage.Commit()
	}

	menuFormCallback.probe.ux_tree()
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
		case "IsWithPrefix":
			FormDivBasicFieldToField(&(node_.IsWithPrefix), formDiv)
		case "Prefix":
			FormDivBasicFieldToField(&(node_.Prefix), formDiv)
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
		case "SecondCheckboxHasToolTip":
			FormDivBasicFieldToField(&(node_.SecondCheckboxHasToolTip), formDiv)
		case "SecondCheckboxToolTipText":
			FormDivBasicFieldToField(&(node_.SecondCheckboxToolTipText), formDiv)
		case "SecondCheckboxToolTipPosition":
			FormDivEnumStringFieldToField(&(node_.SecondCheckboxToolTipPosition), formDiv)
		case "TextAfterSecondCheckbox":
			FormDivBasicFieldToField(&(node_.TextAfterSecondCheckbox), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(node_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(node_.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(node_.ToolTipPosition), formDiv)
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
			node_.Children = instanceSlice
			nodeFormCallback.probe.UpdateSliceOfPointersCallback(node_, "Children", &node_.Children)

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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Button](nodeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			node_.Buttons = instanceSlice
			nodeFormCallback.probe.UpdateSliceOfPointersCallback(node_, "Buttons", &node_.Buttons)

		case "Menu":
			FormDivSelectFieldToField(&(node_.Menu), nodeFormCallback.probe.stageOfInterest, formDiv)
		case "Node:Children":
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

			// 3. Iterate over all Node instances and update their Children slice
			for _node := range *models.GetGongstructInstancesSetFromPointerType[*models.Node](nodeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(nodeFormCallback.probe.stageOfInterest, _node)
				
				// if Node is selected
				if targetNodeIDs[id] {
					// ensure node_ is in _node.Children
					found := false
					for _, _b := range _node.Children {
						if _b == node_ {
							found = true
							break
						}
					}
					if !found {
						_node.Children = append(_node.Children, node_)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Children", &_node.Children)
					}
				} else {
					// ensure node_ is NOT in _node.Children
					idx := slices.Index(_node.Children, node_)
					if idx != -1 {
						_node.Children = slices.Delete(_node.Children, idx, idx+1)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_node, "Children", &_node.Children)
					}
				}
			}
		case "Tree:RootNodes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Tree instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Tree instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Tree](nodeFormCallback.probe.stageOfInterest)
			targetTreeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTreeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Tree instances and update their RootNodes slice
			for _tree := range *models.GetGongstructInstancesSetFromPointerType[*models.Tree](nodeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(nodeFormCallback.probe.stageOfInterest, _tree)
				
				// if Tree is selected
				if targetTreeIDs[id] {
					// ensure node_ is in _tree.RootNodes
					found := false
					for _, _b := range _tree.RootNodes {
						if _b == node_ {
							found = true
							break
						}
					}
					if !found {
						_tree.RootNodes = append(_tree.RootNodes, node_)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_tree, "RootNodes", &_tree.RootNodes)
					}
				} else {
					// ensure node_ is NOT in _tree.RootNodes
					idx := slices.Index(_tree.RootNodes, node_)
					if idx != -1 {
						_tree.RootNodes = slices.Delete(_tree.RootNodes, idx, idx+1)
						nodeFormCallback.probe.UpdateSliceOfPointersCallback(_tree, "RootNodes", &_tree.RootNodes)
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
func __gong__New__SVGIconFormCallback(
	svgicon *models.SVGIcon,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (svgiconFormCallback *SVGIconFormCallback) OnSave() {
	svgiconFormCallback.probe.stageOfInterest.Lock()
	defer svgiconFormCallback.probe.stageOfInterest.Unlock()

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
	updateProbeTable[*models.SVGIcon](
		svgiconFormCallback.probe,
	)

	// display a new form by reset the form stage
	if svgiconFormCallback.CreationMode || svgiconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgiconFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	svgiconFormCallback.probe.ux_tree()
}
func __gong__New__TreeFormCallback(
	tree *models.Tree,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (treeFormCallback *TreeFormCallback) OnSave() {
	treeFormCallback.probe.stageOfInterest.Lock()
	defer treeFormCallback.probe.stageOfInterest.Unlock()

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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Node](treeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			tree_.RootNodes = instanceSlice
			treeFormCallback.probe.UpdateSliceOfPointersCallback(tree_, "RootNodes", &tree_.RootNodes)

		}
	}

	// manage the suppress operation
	if treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tree_.Unstage(treeFormCallback.probe.stageOfInterest)
	}

	treeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Tree](
		treeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if treeFormCallback.CreationMode || treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		treeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	treeFormCallback.probe.ux_tree()
}
