// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder

	if stageSet.Stage != nil {
		assplitOrdered := []*AsSplit{}
		for assplit := range stageSet.Stage.AsSplits {
			assplitOrdered = append(assplitOrdered, assplit)
		}
		sort.Slice(assplitOrdered, func(i, j int) bool {
			return stageSet.Stage.AsSplit_stagedOrder[assplitOrdered[i]] < stageSet.Stage.AsSplit_stagedOrder[assplitOrdered[j]]
		})
		for _, assplit := range assplitOrdered {
			assplitIdent := "__stage_0" + assplit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AsSplit{Name: %s}).Stage(stageSet.Stage)", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", assplitIdent, __gong__toRawStringLiteral(string(assplit.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", assplitIdent, assplit.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", assplitIdent, assplit.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", assplitIdent, assplit.GutterSize))
			for _, elem := range assplit.AsSplitAreas {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplitAreas = append(%s.AsSplitAreas, %s)", assplitIdent, assplitIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		assplitareaOrdered := []*AsSplitArea{}
		for assplitarea := range stageSet.Stage.AsSplitAreas {
			assplitareaOrdered = append(assplitareaOrdered, assplitarea)
		}
		sort.Slice(assplitareaOrdered, func(i, j int) bool {
			return stageSet.Stage.AsSplitArea_stagedOrder[assplitareaOrdered[i]] < stageSet.Stage.AsSplitArea_stagedOrder[assplitareaOrdered[j]]
		})
		for _, assplitarea := range assplitareaOrdered {
			assplitareaIdent := "__stage_0" + assplitarea.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AsSplitArea{Name: %s}).Stage(stageSet.Stage)", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNameInHeader = %t", assplitareaIdent, assplitarea.ShowNameInHeader))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %f", assplitareaIdent, assplitarea.Size))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAny = %t", assplitareaIdent, assplitarea.IsAny))
			values.WriteString(fmt.Sprintf("\n\t%s.HasDiv = %t", assplitareaIdent, assplitarea.HasDiv))
			values.WriteString(fmt.Sprintf("\n\t%s.DivStyle = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.DivStyle)))
			if assplitarea.AsSplit != nil {
				targetIdent := "__stage_0" + assplitarea.AsSplit.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplit = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Button != nil {
				targetIdent := "__stage_0" + assplitarea.Button.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Button = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Form != nil {
				targetIdent := "__stage_0" + assplitarea.Form.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Form = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Load != nil {
				targetIdent := "__stage_0" + assplitarea.Load.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Load = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Split != nil {
				targetIdent := "__stage_0" + assplitarea.Split.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Split = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Svg != nil {
				targetIdent := "__stage_0" + assplitarea.Svg.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Svg = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Table != nil {
				targetIdent := "__stage_0" + assplitarea.Table.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Table = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Tree != nil {
				targetIdent := "__stage_0" + assplitarea.Tree.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tree = %s", assplitareaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		buttonOrdered := []*Button{}
		for button := range stageSet.Stage.Buttons {
			buttonOrdered = append(buttonOrdered, button)
		}
		sort.Slice(buttonOrdered, func(i, j int) bool {
			return stageSet.Stage.Button_stagedOrder[buttonOrdered[i]] < stageSet.Stage.Button_stagedOrder[buttonOrdered[j]]
		})
		for _, button := range buttonOrdered {
			buttonIdent := "__stage_0" + button.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Button{Name: %s}).Stage(stageSet.Stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", buttonIdent, __gong__toRawStringLiteral(button.StackName)))
		}
	}
	if stageSet.Stage != nil {
		faviconOrdered := []*FavIcon{}
		for favicon := range stageSet.Stage.FavIcons {
			faviconOrdered = append(faviconOrdered, favicon)
		}
		sort.Slice(faviconOrdered, func(i, j int) bool {
			return stageSet.Stage.FavIcon_stagedOrder[faviconOrdered[i]] < stageSet.Stage.FavIcon_stagedOrder[faviconOrdered[j]]
		})
		for _, favicon := range faviconOrdered {
			faviconIdent := "__stage_0" + favicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FavIcon{Name: %s}).Stage(stageSet.Stage)", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", faviconIdent, __gong__toRawStringLiteral(favicon.SVG)))
		}
	}
	if stageSet.Stage != nil {
		formOrdered := []*Form{}
		for form := range stageSet.Stage.Forms {
			formOrdered = append(formOrdered, form)
		}
		sort.Slice(formOrdered, func(i, j int) bool {
			return stageSet.Stage.Form_stagedOrder[formOrdered[i]] < stageSet.Stage.Form_stagedOrder[formOrdered[j]]
		})
		for _, form := range formOrdered {
			formIdent := "__stage_0" + form.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Form{Name: %s}).Stage(stageSet.Stage)", formIdent, __gong__toRawStringLiteral(form.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formIdent, __gong__toRawStringLiteral(form.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", formIdent, __gong__toRawStringLiteral(form.StackName)))
		}
	}
	if stageSet.Stage != nil {
		loadOrdered := []*Load{}
		for load := range stageSet.Stage.Loads {
			loadOrdered = append(loadOrdered, load)
		}
		sort.Slice(loadOrdered, func(i, j int) bool {
			return stageSet.Stage.Load_stagedOrder[loadOrdered[i]] < stageSet.Stage.Load_stagedOrder[loadOrdered[j]]
		})
		for _, load := range loadOrdered {
			loadIdent := "__stage_0" + load.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Load{Name: %s}).Stage(stageSet.Stage)", loadIdent, __gong__toRawStringLiteral(load.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", loadIdent, __gong__toRawStringLiteral(load.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", loadIdent, __gong__toRawStringLiteral(load.StackName)))
		}
	}
	if stageSet.Stage != nil {
		logoontheleftOrdered := []*LogoOnTheLeft{}
		for logoontheleft := range stageSet.Stage.LogoOnTheLefts {
			logoontheleftOrdered = append(logoontheleftOrdered, logoontheleft)
		}
		sort.Slice(logoontheleftOrdered, func(i, j int) bool {
			return stageSet.Stage.LogoOnTheLeft_stagedOrder[logoontheleftOrdered[i]] < stageSet.Stage.LogoOnTheLeft_stagedOrder[logoontheleftOrdered[j]]
		})
		for _, logoontheleft := range logoontheleftOrdered {
			logoontheleftIdent := "__stage_0" + logoontheleft.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.LogoOnTheLeft{Name: %s}).Stage(stageSet.Stage)", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontheleftIdent, logoontheleft.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontheleftIdent, logoontheleft.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.SVG)))
		}
	}
	if stageSet.Stage != nil {
		logoontherightOrdered := []*LogoOnTheRight{}
		for logoontheright := range stageSet.Stage.LogoOnTheRights {
			logoontherightOrdered = append(logoontherightOrdered, logoontheright)
		}
		sort.Slice(logoontherightOrdered, func(i, j int) bool {
			return stageSet.Stage.LogoOnTheRight_stagedOrder[logoontherightOrdered[i]] < stageSet.Stage.LogoOnTheRight_stagedOrder[logoontherightOrdered[j]]
		})
		for _, logoontheright := range logoontherightOrdered {
			logoontherightIdent := "__stage_0" + logoontheright.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.LogoOnTheRight{Name: %s}).Stage(stageSet.Stage)", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontherightIdent, logoontheright.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontherightIdent, logoontheright.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.SVG)))
		}
	}
	if stageSet.Stage != nil {
		splitOrdered := []*Split{}
		for split := range stageSet.Stage.Splits {
			splitOrdered = append(splitOrdered, split)
		}
		sort.Slice(splitOrdered, func(i, j int) bool {
			return stageSet.Stage.Split_stagedOrder[splitOrdered[i]] < stageSet.Stage.Split_stagedOrder[splitOrdered[j]]
		})
		for _, split := range splitOrdered {
			splitIdent := "__stage_0" + split.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Split{Name: %s}).Stage(stageSet.Stage)", splitIdent, __gong__toRawStringLiteral(split.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", splitIdent, __gong__toRawStringLiteral(split.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", splitIdent, __gong__toRawStringLiteral(split.StackName)))
		}
	}
	if stageSet.Stage != nil {
		svgOrdered := []*Svg{}
		for svg := range stageSet.Stage.Svgs {
			svgOrdered = append(svgOrdered, svg)
		}
		sort.Slice(svgOrdered, func(i, j int) bool {
			return stageSet.Stage.Svg_stagedOrder[svgOrdered[i]] < stageSet.Stage.Svg_stagedOrder[svgOrdered[j]]
		})
		for _, svg := range svgOrdered {
			svgIdent := "__stage_0" + svg.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Svg{Name: %s}).Stage(stageSet.Stage)", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", svgIdent, __gong__toRawStringLiteral(svg.StackName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Style = %s", svgIdent, __gong__toRawStringLiteral(svg.Style)))
		}
	}
	if stageSet.Stage != nil {
		tableOrdered := []*Table{}
		for table := range stageSet.Stage.Tables {
			tableOrdered = append(tableOrdered, table)
		}
		sort.Slice(tableOrdered, func(i, j int) bool {
			return stageSet.Stage.Table_stagedOrder[tableOrdered[i]] < stageSet.Stage.Table_stagedOrder[tableOrdered[j]]
		})
		for _, table := range tableOrdered {
			tableIdent := "__stage_0" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", tableIdent, __gong__toRawStringLiteral(table.StackName)))
		}
	}
	if stageSet.Stage != nil {
		titleOrdered := []*Title{}
		for title := range stageSet.Stage.Titles {
			titleOrdered = append(titleOrdered, title)
		}
		sort.Slice(titleOrdered, func(i, j int) bool {
			return stageSet.Stage.Title_stagedOrder[titleOrdered[i]] < stageSet.Stage.Title_stagedOrder[titleOrdered[j]]
		})
		for _, title := range titleOrdered {
			titleIdent := "__stage_0" + title.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Title{Name: %s}).Stage(stageSet.Stage)", titleIdent, __gong__toRawStringLiteral(title.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleIdent, __gong__toRawStringLiteral(title.Name)))
		}
	}
	if stageSet.Stage != nil {
		treeOrdered := []*Tree{}
		for tree := range stageSet.Stage.Trees {
			treeOrdered = append(treeOrdered, tree)
		}
		sort.Slice(treeOrdered, func(i, j int) bool {
			return stageSet.Stage.Tree_stagedOrder[treeOrdered[i]] < stageSet.Stage.Tree_stagedOrder[treeOrdered[j]]
		})
		for _, tree := range treeOrdered {
			treeIdent := "__stage_0" + tree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tree{Name: %s}).Stage(stageSet.Stage)", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", treeIdent, __gong__toRawStringLiteral(tree.StackName)))
		}
	}
	if stageSet.Stage != nil {
		viewOrdered := []*View{}
		for view := range stageSet.Stage.Views {
			viewOrdered = append(viewOrdered, view)
		}
		sort.Slice(viewOrdered, func(i, j int) bool {
			return stageSet.Stage.View_stagedOrder[viewOrdered[i]] < stageSet.Stage.View_stagedOrder[viewOrdered[j]]
		})
		for _, view := range viewOrdered {
			viewIdent := "__stage_0" + view.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.View{Name: %s}).Stage(stageSet.Stage)", viewIdent, __gong__toRawStringLiteral(view.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewIdent, __gong__toRawStringLiteral(view.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowViewName = %t", viewIdent, view.ShowViewName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelectedView = %t", viewIdent, view.IsSelectedView))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", viewIdent, __gong__toRawStringLiteral(string(view.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondaryView = %t", viewIdent, view.IsSecondaryView))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", viewIdent, view.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", viewIdent, view.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", viewIdent, view.GutterSize))
			for _, elem := range view.RootAsSplitAreas {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootAsSplitAreas = append(%s.RootAsSplitAreas, %s)", viewIdent, viewIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/splitlite/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "AsSplit":
					if !preserveOrder {
						inst := (&AsSplit{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AsSplit)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "AsSplitArea":
					if !preserveOrder {
						inst := (&AsSplitArea{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AsSplitArea)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Button":
					if !preserveOrder {
						inst := (&Button{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Button)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FavIcon":
					if !preserveOrder {
						inst := (&FavIcon{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FavIcon)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Form":
					if !preserveOrder {
						inst := (&Form{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Form)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Load":
					if !preserveOrder {
						inst := (&Load{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Load)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "LogoOnTheLeft":
					if !preserveOrder {
						inst := (&LogoOnTheLeft{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(LogoOnTheLeft)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "LogoOnTheRight":
					if !preserveOrder {
						inst := (&LogoOnTheRight{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(LogoOnTheRight)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Split":
					if !preserveOrder {
						inst := (&Split{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Split)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Svg":
					if !preserveOrder {
						inst := (&Svg{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Svg)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Table":
					if !preserveOrder {
						inst := (&Table{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Table)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Title":
					if !preserveOrder {
						inst := (&Title{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Title)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tree":
					if !preserveOrder {
						inst := (&Tree{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tree)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "View":
					if !preserveOrder {
						inst := (&View{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(View)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *AsSplit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						inst.Direction = Direction(GongExtractString(rhs))
					case "AsSplitAreas":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AsSplitArea); ok {
										inst.AsSplitAreas = append(inst.AsSplitAreas, typedTarget)
									}
								}
							}
						}
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *AsSplitArea:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowNameInHeader":
						inst.ShowNameInHeader = GongExtractBool(rhs)
					case "Size":
						inst.Size = GongExtractFloat(rhs)
					case "IsAny":
						inst.IsAny = GongExtractBool(rhs)
					case "AsSplit":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AsSplit); ok {
									inst.AsSplit = typedTarget
								}
							}
						}
					case "Button":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Button); ok {
									inst.Button = typedTarget
								}
							}
						}
					case "Form":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Form); ok {
									inst.Form = typedTarget
								}
							}
						}
					case "Load":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Load); ok {
									inst.Load = typedTarget
								}
							}
						}
					case "Split":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Split); ok {
									inst.Split = typedTarget
								}
							}
						}
					case "Svg":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Svg); ok {
									inst.Svg = typedTarget
								}
							}
						}
					case "Table":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Table); ok {
									inst.Table = typedTarget
								}
							}
						}
					case "Tree":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tree); ok {
									inst.Tree = typedTarget
								}
							}
						}
					case "HasDiv":
						inst.HasDiv = GongExtractBool(rhs)
					case "DivStyle":
						inst.DivStyle = GongExtractString(rhs)
					}
				case *Button:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *FavIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Form:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Load:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *LogoOnTheLeft:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *LogoOnTheRight:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Split:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Svg:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					case "Style":
						inst.Style = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Title:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Tree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *View:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowViewName":
						inst.ShowViewName = GongExtractBool(rhs)
					case "RootAsSplitAreas":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AsSplitArea); ok {
										inst.RootAsSplitAreas = append(inst.RootAsSplitAreas, typedTarget)
									}
								}
							}
						}
					case "IsSelectedView":
						inst.IsSelectedView = GongExtractBool(rhs)
					case "Direction":
						inst.Direction = Direction(GongExtractString(rhs))
					case "IsSecondaryView":
						inst.IsSecondaryView = GongExtractBool(rhs)
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}
