package models

import (
	"go/parser"
	"go/token"
	"strings"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(
	treeInstance *tree.Tree,
	library *Library,
	parentNodes *[]*tree.Node,
	is3DView bool,
) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.getRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}
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

	confPlants := ItemButtonConfiguration[
		PlantAbstract, *PlantAbstract, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Plants,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confPlants)

	if library == stager.getRootLibrary() && DataFS != nil {
		entries, err := DataFS.ReadDir("data")
		if err == nil {
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				entryName := entry.Name()
				if !strings.HasSuffix(entryName, ".go") {
					continue
				}
				if libraryNode.Menu == nil {
					libraryNode.Menu = &tree.Menu{Name: "Menu"}
				}
				libraryNode.Menu.Buttons = append(libraryNode.Menu.Buttons, &tree.Button{
					Name: entryName,
					Icon: string(buttons.BUTTON_file_open),
					OnClick: func() {
						content, err := DataFS.ReadFile("data/" + entryName)
						if err == nil {
							// if the user loads a data file, we don't want the file to be automatically overwritten
							stager.stage.OnInitCommitCallback = nil

							stager.stage.Reset()

							fset := token.NewFileSet()
							file, err := parser.ParseFile(fset, "", string(content), parser.ParseComments)
							if err == nil {
								ParseAstFileFromAst(stager.stage, file, fset, true)
								stager.stage.ComputeReverseMaps()
								stager.stage.ComputeInstancesNb()
								stager.stage.ComputeReferenceAndOrders()
								stager.stage.Commit()
								stager.probeForm.Refresh()
							}
						}
					},
				})
			}
		}
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children, is3DView)
	}

	for _, plant := range library.Plants {
		stager.treePlant(plant, &libraryNode.Children, is3DView)
	}
}
