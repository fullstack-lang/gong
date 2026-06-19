package models

import (
	"embed"
	"go/parser"
	"go/token"
	"strings"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var DataFS *embed.FS

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.GetRootLibrary()
	_ = rootLibrary

	treeInstance := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(treeInstance, rootLibrary, &treeInstance.RootNodes)

	examplesNode := &tree.Node{
		Name:            "Examples",
		IsNodeClickable: true,
		IsExpanded:      true,
	}

	examplesNode.Menu = &tree.Menu{
		Name: "Examples Menu",
	}

	if DataFS != nil {
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
				examplesNode.Menu.Buttons = append(examplesNode.Menu.Buttons, &tree.Button{
					Name: entryName,
					Icon: string(buttons.BUTTON_file_open),
					OnClick: func() {
						content, err := DataFS.ReadFile("data/" + entryName)
						if err == nil {
							// if the user loads an example file, we don't want the file to be automatically overwritten
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
	treeInstance.RootNodes = append(treeInstance.RootNodes, examplesNode)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}
