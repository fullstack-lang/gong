package models

import "fmt"

func GetStagerTemplate(useSplitlite bool, dsm bool) string {
	splitImport := `split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"`
	if useSplitlite {
		splitImport = `split "github.com/fullstack-lang/gong/lib/splitlite/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/splitlite/go/stack"`
	}

	if !dsm {
		return fmt.Sprintf(`// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"net/http"

	%s
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF
}

func NewStager(
	r *http.ServeMux,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager.splitStage.StageBranch(&split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()

	callbacks := &BeforeCommitImplementation{
		stager: stager,
	}
	stager.stage.OnInitCommitFromBackCallback = callbacks
	callbacks.BeforeCommit(stage)

	return
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}
`, splitImport)
	}

	return fmt.Sprintf(`// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"net/http"

	%s

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	treeStage           *tree.Stage
	svgStage            *svg.Stage
	loadStage           *load.Stage
	loadStageMultistage *load.Stage
	buttonStage         *button.Stage

	svgObject *svg.SVG
	fileName  string

	map_Element_Diagrams map[AbstractType][]DiagramIF
}

func (stager *Stager) GetSvgObject() *svg.SVG {
	return stager.svgObject
}

func (stager *Stager) GetTreeStage() *tree.Stage {
	return stager.treeStage
}

func (stager *Stager) GetSvgStage() *svg.Stage {
	return stager.svgStage
}

func (stager *Stager) exportWebsite() {}

func (stager *Stager) enforceSemantic() (needCommit bool) {
	needCommit = stager.enforceThereIsARootLibrary() || needCommit
	return
}

func (stager *Stager) svg() {
	stager.svgStage.Reset()
	stager.svgStage.Commit()
}

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.getRootLibrary()
	if rootLibrary == nil {
		stager.treeStage.Commit()
		return
	}

	treeInstance := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}

	if stager.probeForm != nil {
		stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
			treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
		})
	}

	stager.treeLibrary(treeInstance, rootLibrary, &treeInstance.RootNodes)

	EnsureRenameIsFirstOnTree(treeInstance)

	stager.treeStage.StageBranch(treeInstance)

	stager.treeStage.Commit()
}

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	if library == nil {
		return
	}
	name := library.Name
	if name == "" {
		name = "Root Library"
	}
	libraryNode := &tree.Node{
		Name:                 name,
		IsExpanded:           library.IsExpanded,
		IsNodeClickable:      true,
		IsInEditMode:         library.GetIsInRenameMode(),
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(tree_buttons.BUTTON_local_library),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	addRenameButton(library, libraryNode, stager)
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpanded)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confSubLibraries)

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}

	EnsureRenameIsFirst(libraryNode)
}

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	stager.splitStage.StageBranch(&split.View{
		Name:           "Working View",
		Direction:      split.Horizontal,
		IsSelectedView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with tree and SVG",
				ShowNameInHeader: false,
				Size:             80,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 30,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Tree",
										Size:             80,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
									{
										Size: 10,
										Load: &split.Load{
											StackName: stager.loadStage.GetName(),
										},
									},
									{
										Size: 10,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 70,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 20,
				Form: &split.Form{
					StackName: stager.probeForm.GetFormStage().GetName(),
				},
			},
		},
	})

	stager.splitStage.StageBranch(&split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.stage.GetProbeSplitStageName(),
				},
			},
		},
	})
}

func NewStager(
	r *http.ServeMux,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	stage.SetDeltaMode(true)

	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.treeStage.RegisterBeforeCommit(func(treeStage *tree.Stage) {
		for node := range *treeStage.GetInstancesSet[*tree.Node]() {
			EnsureRenameIsFirst(node)
		}
	})
	stager.svgStage = svg_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")
	stager.loadStageMultistage, _ = load_fullstack.NewStackInstance(r, "multistage")
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.createViews()

	stager.splitStage.Commit()

	// Setup your before commit sequence

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree() // DSM mandatory name, to be changed
		stager.svg()
		stager.button()
		stager.load()
		if stager.probeForm != nil {
			stager.probeForm.Refresh()
		}
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	return
}
`, splitImport)
}

var StagerFileTemplate = GetStagerTemplate(false, false)
