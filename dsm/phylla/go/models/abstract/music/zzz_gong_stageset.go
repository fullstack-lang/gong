// generated code - do not edit
package music

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
		musicabstractOrdered := []*MusicAbstract{}
		for musicabstract := range stageSet.Stage.MusicAbstracts {
			musicabstractOrdered = append(musicabstractOrdered, musicabstract)
		}
		sort.Slice(musicabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.MusicAbstract_stagedOrder[musicabstractOrdered[i]] < stageSet.Stage.MusicAbstract_stagedOrder[musicabstractOrdered[j]]
		})
		for _, musicabstract := range musicabstractOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			musicabstractIdent := "__music" + musicabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&music.MusicAbstract{Name: %s}).Stage(stageSet.Stage)", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", musicabstractIdent, musicabstract.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchHeight = %f", musicabstractIdent, musicabstract.PitchHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.NbOfBeatsInTheme = %d", musicabstractIdent, musicabstract.NbOfBeatsInTheme))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatsPerSecond = %f", musicabstractIdent, musicabstract.BeatsPerSecond))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftX = %f", musicabstractIdent, musicabstract.FirstVoiceShiftX))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftY = %f", musicabstractIdent, musicabstract.FirstVoiceShiftY))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchDifference = %d", musicabstractIdent, musicabstract.PitchDifference))
			values.WriteString(fmt.Sprintf("\n\t%s.Level = %f", musicabstractIdent, musicabstract.Level))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualBeatsTemporalShift = %d", musicabstractIdent, musicabstract.ActualBeatsTemporalShift))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMinor = %t", musicabstractIdent, musicabstract.IsMinor))
			values.WriteString(fmt.Sprintf("\n\t%s.ThemeBinaryEncoding = %d", musicabstractIdent, musicabstract.ThemeBinaryEncoding))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierControlLengthRatio = %f", musicabstractIdent, musicabstract.BezierControlLengthRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPitchLines = %d", musicabstractIdent, musicabstract.NbPitchLines))
			values.WriteString(fmt.Sprintf("\n\t%s.NbBeatLines = %d", musicabstractIdent, musicabstract.NbBeatLines))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginX = %f", musicabstractIdent, musicabstract.OriginX))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginY = %f", musicabstractIdent, musicabstract.OriginY))
			values.WriteString(fmt.Sprintf("\n\t%s.ScoreScale = %f", musicabstractIdent, musicabstract.ScoreScale))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoice = %t", musicabstractIdent, musicabstract.ShowFirstVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoice = %t", musicabstractIdent, musicabstract.ShowSecondVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotes = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotes = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComposerNodeExpanded = %t", musicabstractIdent, musicabstract.IsComposerNodeExpanded))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *music.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *music.StageSet) {

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
		case "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music":
			aliasToCanonical[alias] = "music"
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
			case "music":
				switch typeName {
				case "MusicAbstract":
					if !preserveOrder {
						inst := (&MusicAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MusicAbstract)
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
				case *MusicAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "PitchHeight":
						inst.PitchHeight = GongExtractFloat(rhs)
					case "NbOfBeatsInTheme":
						inst.NbOfBeatsInTheme = GongExtractInt(rhs)
					case "BeatsPerSecond":
						inst.BeatsPerSecond = GongExtractFloat(rhs)
					case "FirstVoiceShiftX":
						inst.FirstVoiceShiftX = GongExtractFloat(rhs)
					case "FirstVoiceShiftY":
						inst.FirstVoiceShiftY = GongExtractFloat(rhs)
					case "PitchDifference":
						inst.PitchDifference = GongExtractInt(rhs)
					case "Level":
						inst.Level = GongExtractFloat(rhs)
					case "ActualBeatsTemporalShift":
						inst.ActualBeatsTemporalShift = GongExtractInt(rhs)
					case "IsMinor":
						inst.IsMinor = GongExtractBool(rhs)
					case "ThemeBinaryEncoding":
						inst.ThemeBinaryEncoding = GongExtractInt(rhs)
					case "BezierControlLengthRatio":
						inst.BezierControlLengthRatio = GongExtractFloat(rhs)
					case "NbPitchLines":
						inst.NbPitchLines = GongExtractInt(rhs)
					case "NbBeatLines":
						inst.NbBeatLines = GongExtractInt(rhs)
					case "OriginX":
						inst.OriginX = GongExtractFloat(rhs)
					case "OriginY":
						inst.OriginY = GongExtractFloat(rhs)
					case "ScoreScale":
						inst.ScoreScale = GongExtractFloat(rhs)
					case "ShowFirstVoice":
						inst.ShowFirstVoice = GongExtractBool(rhs)
					case "ShowFirstVoiceShiftRight":
						inst.ShowFirstVoiceShiftRight = GongExtractBool(rhs)
					case "ShowSecondVoice":
						inst.ShowSecondVoice = GongExtractBool(rhs)
					case "ShowSecondVoiceShiftRight":
						inst.ShowSecondVoiceShiftRight = GongExtractBool(rhs)
					case "ShowFirstVoiceNotes":
						inst.ShowFirstVoiceNotes = GongExtractBool(rhs)
					case "ShowFirstVoiceNotesShiftRight":
						inst.ShowFirstVoiceNotesShiftRight = GongExtractBool(rhs)
					case "ShowSecondVoiceNotes":
						inst.ShowSecondVoiceNotes = GongExtractBool(rhs)
					case "ShowSecondVoiceNotesShiftRight":
						inst.ShowSecondVoiceNotesShiftRight = GongExtractBool(rhs)
					case "IsComposerNodeExpanded":
						inst.IsComposerNodeExpanded = GongExtractBool(rhs)
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
