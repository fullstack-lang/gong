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
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, arrow := range __gong__sortStageSetInstances(stageSet.Stage.Arrows, stageSet.Stage.Arrow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrowIdent := "__models" + arrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Arrow{Name: %s}).Stage(stageSet.Stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalStroke)))
			if arrow.From != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + arrow.From.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.From = %s", arrowIdent, targetIdent))
			}
			if arrow.To != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + arrow.To.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.To = %s", arrowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bar := range __gong__sortStageSetInstances(stageSet.Stage.Bars, stageSet.Stage.Bar_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			barIdent := "__models" + bar.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bar{Name: %s}).Stage(stageSet.Stage)", barIdent, __gong__toRawStringLiteral(bar.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barIdent, __gong__toRawStringLiteral(bar.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", barIdent, int64(bar.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalStroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", barIdent, bar.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", barIdent, bar.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", barIdent, __gong__toRawStringLiteral(bar.StrokeDashArray)))
		}
	}
	if stageSet.Stage != nil {
		for _, gantt := range __gong__sortStageSetInstances(stageSet.Stage.Gantts, stageSet.Stage.Gantt_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ganttIdent := "__models" + gantt.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Gantt{Name: %s}).Stage(stageSet.Stage)", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", ganttIdent, int64(gantt.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", ganttIdent, gantt.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.LaneHeight = %f", ganttIdent, gantt.LaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RatioBarToLaneHeight = %f", ganttIdent, gantt.RatioBarToLaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.YTopMargin = %f", ganttIdent, gantt.YTopMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftText = %f", ganttIdent, gantt.XLeftText))
			values.WriteString(fmt.Sprintf("\n\t%s.TextHeight = %f", ganttIdent, gantt.TextHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftLanes = %f", ganttIdent, gantt.XLeftLanes))
			values.WriteString(fmt.Sprintf("\n\t%s.XRightMargin = %f", ganttIdent, gantt.XRightMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowLengthToTheRightOfStartBar = %f", ganttIdent, gantt.ArrowLengthToTheRightOfStartBar))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowTipLenght = %f", ganttIdent, gantt.ArrowTipLenght))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Color = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_FillOpacity = %f", ganttIdent, gantt.TimeLine_FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_StrokeWidth = %f", ganttIdent, gantt.TimeLine_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", ganttIdent, gantt.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", ganttIdent, gantt.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnStartEndOnYearStart = %t", ganttIdent, gantt.AlignOnStartEndOnYearStart))
			for _, elem := range gantt.Lanes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lanes = append(%s.Lanes, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Milestones {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Milestones = append(%s.Milestones, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Arrows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrows = append(%s.Arrows, %s)", ganttIdent, ganttIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, group := range __gong__sortStageSetInstances(stageSet.Stage.Groups, stageSet.Stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := "__models" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			for _, elem := range group.GroupLanes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupLanes = append(%s.GroupLanes, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, lane := range __gong__sortStageSetInstances(stageSet.Stage.Lanes, stageSet.Stage.Lane_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			laneIdent := "__models" + lane.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Lane{Name: %s}).Stage(stageSet.Stage)", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", laneIdent, lane.Order))
			for _, elem := range lane.Bars {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bars = append(%s.Bars, %s)", laneIdent, laneIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, laneuse := range __gong__sortStageSetInstances(stageSet.Stage.LaneUses, stageSet.Stage.LaneUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			laneuseIdent := "__models" + laneuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LaneUse{Name: %s}).Stage(stageSet.Stage)", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			if laneuse.Lane != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + laneuse.Lane.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lane = %s", laneuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, milestone := range __gong__sortStageSetInstances(stageSet.Stage.Milestones, stageSet.Stage.Milestone_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			milestoneIdent := "__models" + milestone.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Milestone{Name: %s}).Stage(stageSet.Stage)", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", milestoneIdent, milestone.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", milestoneIdent, milestone.DisplayVerticalBar))
			for _, elem := range milestone.LanesToDisplay {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LanesToDisplay = append(%s.LanesToDisplay, %s)", milestoneIdent, milestoneIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

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

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/lib/gantt/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

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

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arrow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bar":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bar), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Gantt":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Gantt), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Lane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lane), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LaneUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LaneUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Milestone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Milestone), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "From":
						__gong__assignPointer(&inst.From, rhs, identifierMap)
					case "To":
						__gong__assignPointer(&inst.To, rhs, identifierMap)
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					}
				case *Bar:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					}
				case *Gantt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedStart":
						inst.ComputedStart = GongExtractDate(rhs)
					case "ComputedEnd":
						inst.ComputedEnd = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						inst.ManualStart = GongExtractDate(rhs)
					case "ManualEnd":
						inst.ManualEnd = GongExtractDate(rhs)
					case "LaneHeight":
						inst.LaneHeight = GongExtractFloat(rhs)
					case "RatioBarToLaneHeight":
						inst.RatioBarToLaneHeight = GongExtractFloat(rhs)
					case "YTopMargin":
						inst.YTopMargin = GongExtractFloat(rhs)
					case "XLeftText":
						inst.XLeftText = GongExtractFloat(rhs)
					case "TextHeight":
						inst.TextHeight = GongExtractFloat(rhs)
					case "XLeftLanes":
						inst.XLeftLanes = GongExtractFloat(rhs)
					case "XRightMargin":
						inst.XRightMargin = GongExtractFloat(rhs)
					case "ArrowLengthToTheRightOfStartBar":
						inst.ArrowLengthToTheRightOfStartBar = GongExtractFloat(rhs)
					case "ArrowTipLenght":
						inst.ArrowTipLenght = GongExtractFloat(rhs)
					case "TimeLine_Color":
						inst.TimeLine_Color = GongExtractString(rhs)
					case "TimeLine_FillOpacity":
						inst.TimeLine_FillOpacity = GongExtractFloat(rhs)
					case "TimeLine_Stroke":
						inst.TimeLine_Stroke = GongExtractString(rhs)
					case "TimeLine_StrokeWidth":
						inst.TimeLine_StrokeWidth = GongExtractFloat(rhs)
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnStartEndOnYearStart":
						inst.AlignOnStartEndOnYearStart = GongExtractBool(rhs)
					case "Lanes":
						__gong__assignSliceOfPointers(&inst.Lanes, rhs, identifierMap)
					case "Milestones":
						__gong__assignSliceOfPointers(&inst.Milestones, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Arrows":
						__gong__assignSliceOfPointers(&inst.Arrows, rhs, identifierMap)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GroupLanes":
						__gong__assignSliceOfPointers(&inst.GroupLanes, rhs, identifierMap)
					}
				case *Lane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Bars":
						__gong__assignSliceOfPointers(&inst.Bars, rhs, identifierMap)
					}
				case *LaneUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lane":
						__gong__assignPointer(&inst.Lane, rhs, identifierMap)
					}
				case *Milestone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
					case "LanesToDisplay":
						__gong__assignSliceOfPointers(&inst.LanesToDisplay, rhs, identifierMap)
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

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
