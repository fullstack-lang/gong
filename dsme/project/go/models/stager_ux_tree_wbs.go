package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeWBSinDiagram(diagram *Diagram, task *Task, parentNode *tree.Node) {
	stage := stager.stage

	svgIconLinkVisibilityOn := &tree.SVGIcon{
		Name: "Show Visibility On",
		SVG:  `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24"><path d="M 40,-480 C 160,-220 800,-220 920,-480 C 800,-740 160,-740 40,-480 Z M 640,-480 C 640,-568.4 568.4,-640 480,-640 C 391.6,-640 320,-568.4 320,-480 C 320,-391.6 391.6,-320 480,-320 C 568.4,-320 640,-391.6 640,-480 Z M 354,-491.6 L 427.6,-430 A 40 40 0 0 0 479.2,-491.2 L 405.6,-552.8 A 40 40 0 0 0 354,-491.6 Z M 480.8,-468.8 L 554.4,-407.2 A 40 40 0 0 0 606,-468.4 L 532.4,-530 A 40 40 0 0 0 480.8,-468.8 Z" fill-rule="evenodd"/></svg>`,
	}

	svgIconLinkVisibilityOff := &tree.SVGIcon{
		Name: "Show Visibility Off",
		SVG: `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!-- Created with Inkscape (http://www.inkscape.org/) -->

<svg
   width="24"
   height="24"
   viewBox="0 0 6.3499999 6.35"
   version="1.1"
   id="svg1"
   xml:space="preserve"
   inkscape:version="1.3.2 (091e20e, 2023-11-25, custom)"
   sodipodi:docname="link visibility off.svg"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg"><sodipodi:namedview
     id="namedview1"
     pagecolor="#ffffff"
     bordercolor="#000000"
     borderopacity="0.25"
     inkscape:showpageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="0"
     inkscape:deskcolor="#d1d1d1"
     inkscape:document-units="mm"
     showguides="false"
     showgrid="true"
     inkscape:zoom="9.3102393"
     inkscape:cx="-19.54837"
     inkscape:cy="2.3092854"
     inkscape:window-width="1920"
     inkscape:window-height="1009"
     inkscape:window-x="-8"
     inkscape:window-y="123"
     inkscape:window-maximized="1"
     inkscape:current-layer="layer1"><inkscape:grid
       id="grid2"
       units="mm"
       originx="0"
       originy="0"
       spacingx="0.99999998"
       spacingy="1"
       empcolor="#0099e5"
       empopacity="0.30196078"
       color="#0099e5"
       opacity="0.14901961"
       empspacing="5"
       dotted="false"
       gridanglex="30"
       gridanglez="30"
       visible="true" /></sodipodi:namedview><defs
     id="defs1" /><g
     inkscape:label="Calque 1"
     inkscape:groupmode="layer"
     id="layer1"><g
       style="fill:#e3e3e3;stroke-width:1.14644"
       id="g1"
       transform="matrix(0.00574265,0,0,0.00579678,0.75734243,5.2870158)"><path
         d="m 423,-716 c -429.1108007,-232.00036 -429.1108007,-232.00036 0,0 z m 349,414 -58,-56 c 25.33333,-19.33333 47.83333,-40.5 67.5,-63.5 19.66667,-23 36.5,-49.16667 50.5,-78.5 C 798.66667,-567.33333 750.83333,-620.83333 688.5,-660.5 626.16667,-700.16667 556.66667,-720 480,-720 c -19.33333,0 -38.33333,1.33333 -57,4 -18.66667,2.66667 -37,6.66667 -55,12 l -62,-62 c 27.33333,-11.33333 55.33333,-19.83333 84,-25.5 28.66667,-5.66667 58.66667,-8.5 90,-8.5 100.66667,0 190.33333,27.83333 269,83.5 78.66667,55.66667 135.66667,127.83333 171,216.5 -15.33333,39.33333 -35.5,75.83333 -60.5,109.5 -25,33.66667 -54.16667,63.16667 -87.5,88.5 z m 20,246 -168,-166 c -23.33333,7.33333 -46.83333,12.83333 -70.5,16.5 -23.66667,3.66667 -48.16667,5.5 -73.5,5.5 -100.66667,0 -190.33333,-27.83333 -269,-83.5 C 132.33333,-339.16667 75.333333,-411.33333 40,-500 c 14,-35.33333 31.666667,-68.16667 53,-98.5 21.33333,-30.33333 45.66667,-57.5 73,-81.5 l -110,-112 56,-56 736,736 z M 222,-624 c -19.33333,17.33333 -37,36.33333 -53,57 -16,20.66667 -29.66667,43 -41,67 33.33333,67.33333 81.16667,120.83333 143.5,160.5 62.33333,39.66667 131.83333,59.5 208.5,59.5 13.33333,0 26.33333,-0.83333 39,-2.5 12.66667,-1.66667 25.66667,-3.5 39,-5.5 l -36,-38 C 453.14989,-394.85011 379.04113,-468.00011 306,-542 Z"
         id="path1"
         sodipodi:nodetypes="ssccscsssccssscscccssscsccccccscssscccc"
         style="stroke-width:1.14644" /></g><g
       style="fill:#e3e3e3;stroke-width:1.9377"
       id="g2"
       transform="matrix(0,0.00329682,-0.00353458,0,-0.09052875,3.3237783)"><path
         d="m 318,-120 q -82,0 -140,-58 -58,-58 -58,-140 0,-40 15,-76 15,-36 43,-64 l 134,-133 56,56 -134,134 q -17,17 -25.5,38.5 -8.5,21.5 -8.5,44.5 0,49 34.5,83.5 34.5,34.5 83.5,34.5 23,0 45,-8.5 22,-8.5 39,-25.5 l 133,-134 57,57 -134,133 q -28,28 -64,43 -36,15 -76,15 z m 79,-220 -57,-57 223,-223 57,57 z m 251,-28 -56,-57 134,-133 q 17,-17 25,-38 8,-21 8,-44 0,-50 -34,-85 -34,-35 -84,-35 -23,0 -44.5,8.5 Q 575,-743 558,-726 l -133,134 -57,-56 134,-134 q 28,-28 64,-43 36,-15 76,-15 82,0 139.5,58 57.5,58 57.5,141 0,39 -14.5,75 -14.5,36 -42.5,64 z"
         id="path1-3"
         style="stroke-width:1.9377" /></g></g></svg>
`,
	}

	taskNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		task,
		&diagram.TasksWhoseNodeIsExpanded,
		&diagram.Task_Shapes,
		&diagram.map_Task_TaskShape,
	)

	addAddItemButton(stager, &diagram.TasksWhoseNodeIsExpanded, task, nil, taskNode, &task.SubTasks, diagram, &diagram.Task_Shapes, &diagram.TaskComposition_Shapes)

	// if task has a parent task, add a button to show/hide the link to the parent
	addShowHideCompositionButton(
		stager,
		task,
		task.parentTask,
		taskNode,
		diagram.map_Task_TaskShape,
		diagram.map_Task_TaskCompositionShape,
		&diagram.TaskComposition_Shapes,
	)

	for _, task := range task.SubTasks {
		stager.treeWBSinDiagram(diagram, task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("In (%d)", len(task.Inputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseInputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)

		inputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagram.TasksWhoseInputNodeIsExpanded)

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			// if input task is present in diagram as well as the input product
			// display the show/hide input relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Task_TaskShape[task]; ok {

					inputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskInputShape, ok := diagram.map_Task_TaskInputShape[taskProductKey]
					inputProductNode.IsChecked = ok

					if ok {
						inputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						inputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					inputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskInputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskInputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					inputProductNode.Buttons = []*tree.Button{
						{
							Name: diagram.GetName(),
							// Icon:            string(buttons.BUTTON_visibility_off),
							SVGIcon:         svgIconLinkVisibilityOff,
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								taskInputShape.SetIsHidden(!taskInputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskInputShape.GetIsHidden() {
							// inputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							inputProductNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
							inputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						inputProductNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}

	}
	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("Out (%d)", len(task.Outputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseOutputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagram.TasksWhoseOutputNodeIsExpanded)

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			// if output task is present in diagram as well as the output product
			// display the show/hide output relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Task_TaskShape[task]; ok {

					outputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskOutputShape, ok := diagram.map_Task_TaskOutputShape[taskProductKey]
					outputProductNode.IsChecked = ok

					if ok {
						outputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						outputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					outputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskOutputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskOutputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					outputProductNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								taskOutputShape.SetIsHidden(!taskOutputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskOutputShape.GetIsHidden() {
							outputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							outputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						outputProductNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
