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
		for _, chapter := range __gong__sortStageSetInstances(stageSet.Stage.Chapters, stageSet.Stage.Chapter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			chapterIdent := "__models" + chapter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Chapter{Name: %s}).Stage(stageSet.Stage)", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", chapterIdent, __gong__toRawStringLiteral(chapter.MardownContent)))
			for _, elem := range chapter.Sections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.Pages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pages = append(%s.Pages, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.SubChapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubChapters = append(%s.SubChapters, %s)", chapterIdent, chapterIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, content := range __gong__sortStageSetInstances(stageSet.Stage.Contents, stageSet.Stage.Content_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			contentIdent := "__models" + content.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Content{Name: %s}).Stage(stageSet.Stage)", contentIdent, __gong__toRawStringLiteral(content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contentIdent, __gong__toRawStringLiteral(content.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", contentIdent, __gong__toRawStringLiteral(content.MardownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.ContentPath = %s", contentIdent, __gong__toRawStringLiteral(content.ContentPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.OutputPath = %s", contentIdent, __gong__toRawStringLiteral(content.OutputPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaticPath = %s", contentIdent, __gong__toRawStringLiteral(content.StaticPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", contentIdent, __gong__toRawStringLiteral(content.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeLogoFileName = %t", contentIdent, content.IsBespokeLogoFileName))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeLogoFileName = %s", contentIdent, __gong__toRawStringLiteral(content.BespokeLogoFileName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokePageTileLogoFileName = %t", contentIdent, content.IsBespokePageTileLogoFileName))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokePageTileLogoFileName = %s", contentIdent, __gong__toRawStringLiteral(content.BespokePageTileLogoFileName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Target = %s", contentIdent, __gong__toRawStringLiteral(string(content.Target))))
			values.WriteString(fmt.Sprintf("\n\t%s.VersionInfo = %s", contentIdent, __gong__toRawStringLiteral(content.VersionInfo)))
			for _, elem := range content.Chapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chapters = append(%s.Chapters, %s)", contentIdent, contentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, downloadablefile := range __gong__sortStageSetInstances(stageSet.Stage.DownloadableFiles, stageSet.Stage.DownloadableFile_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			downloadablefileIdent := "__models" + downloadablefile.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DownloadableFile{Name: %s}).Stage(stageSet.Stage)", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		for _, jpgimage := range __gong__sortStageSetInstances(stageSet.Stage.JpgImages, stageSet.Stage.JpgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			jpgimageIdent := "__models" + jpgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.JpgImage{Name: %s}).Stage(stageSet.Stage)", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		for _, page := range __gong__sortStageSetInstances(stageSet.Stage.Pages, stageSet.Stage.Page_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pageIdent := "__models" + page.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Page{Name: %s}).Stage(stageSet.Stage)", pageIdent, __gong__toRawStringLiteral(page.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pageIdent, __gong__toRawStringLiteral(page.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", pageIdent, __gong__toRawStringLiteral(page.MardownContent)))
			for _, elem := range page.Sections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", pageIdent, pageIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, pngimage := range __gong__sortStageSetInstances(stageSet.Stage.PngImages, stageSet.Stage.PngImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pngimageIdent := "__models" + pngimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PngImage{Name: %s}).Stage(stageSet.Stage)", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		for _, section := range __gong__sortStageSetInstances(stageSet.Stage.Sections, stageSet.Stage.Section_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sectionIdent := "__models" + section.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Section{Name: %s}).Stage(stageSet.Stage)", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", sectionIdent, __gong__toRawStringLiteral(section.MardownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImage = %t", sectionIdent, section.IsImage))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDownloadableFile = %t", sectionIdent, section.IsDownloadableFile))
			if section.SvgImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + section.SvgImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SvgImage = %s", sectionIdent, targetIdent))
			}
			if section.PngImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + section.PngImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PngImage = %s", sectionIdent, targetIdent))
			}
			if section.JpgImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + section.JpgImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.JpgImage = %s", sectionIdent, targetIdent))
			}
			if section.DownloadableFile != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + section.DownloadableFile.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DownloadableFile = %s", sectionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, svgimage := range __gong__sortStageSetInstances(stageSet.Stage.SvgImages, stageSet.Stage.SvgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgimageIdent := "__models" + svgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SvgImage{Name: %s}).Stage(stageSet.Stage)", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Content)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
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
		case "github.com/fullstack-lang/gong/lib/ssg/go/models":
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
				case "Chapter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Chapter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Content":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Content), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DownloadableFile":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DownloadableFile), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "JpgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(JpgImage), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Page":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PngImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PngImage), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Section":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Section), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SvgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SvgImage), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Chapter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "Sections":
						__gong__assignSliceOfPointers(&inst.Sections, rhs, identifierMap)
					case "Pages":
						__gong__assignSliceOfPointers(&inst.Pages, rhs, identifierMap)
					case "SubChapters":
						__gong__assignSliceOfPointers(&inst.SubChapters, rhs, identifierMap)
					}
				case *Content:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "ContentPath":
						inst.ContentPath = GongExtractString(rhs)
					case "OutputPath":
						inst.OutputPath = GongExtractString(rhs)
					case "StaticPath":
						inst.StaticPath = GongExtractString(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "IsBespokeLogoFileName":
						inst.IsBespokeLogoFileName = GongExtractBool(rhs)
					case "BespokeLogoFileName":
						inst.BespokeLogoFileName = GongExtractString(rhs)
					case "IsBespokePageTileLogoFileName":
						inst.IsBespokePageTileLogoFileName = GongExtractBool(rhs)
					case "BespokePageTileLogoFileName":
						inst.BespokePageTileLogoFileName = GongExtractString(rhs)
					case "Target":
						inst.Target = Target(GongExtractString(rhs))
					case "Chapters":
						__gong__assignSliceOfPointers(&inst.Chapters, rhs, identifierMap)
					case "VersionInfo":
						inst.VersionInfo = GongExtractString(rhs)
					}
				case *DownloadableFile:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *JpgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *Page:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "Sections":
						__gong__assignSliceOfPointers(&inst.Sections, rhs, identifierMap)
					}
				case *PngImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *Section:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "IsImage":
						inst.IsImage = GongExtractBool(rhs)
					case "SvgImage":
						__gong__assignPointer(&inst.SvgImage, rhs, identifierMap)
					case "PngImage":
						__gong__assignPointer(&inst.PngImage, rhs, identifierMap)
					case "JpgImage":
						__gong__assignPointer(&inst.JpgImage, rhs, identifierMap)
					case "IsDownloadableFile":
						inst.IsDownloadableFile = GongExtractBool(rhs)
					case "DownloadableFile":
						__gong__assignPointer(&inst.DownloadableFile, rhs, identifierMap)
					}
				case *SvgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
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
