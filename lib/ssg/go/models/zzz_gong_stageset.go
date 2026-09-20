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
		chapterOrdered := []*Chapter{}
		for chapter := range stageSet.Stage.Chapters {
			chapterOrdered = append(chapterOrdered, chapter)
		}
		sort.Slice(chapterOrdered, func(i, j int) bool {
			return stageSet.Stage.Chapter_stagedOrder[chapterOrdered[i]] < stageSet.Stage.Chapter_stagedOrder[chapterOrdered[j]]
		})
		for _, chapter := range chapterOrdered {
			chapterIdent := "__stage_0" + chapter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Chapter{Name: %s}).Stage(stageSet.Stage)", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", chapterIdent, __gong__toRawStringLiteral(chapter.MardownContent)))
			for _, elem := range chapter.Sections {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.Pages {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pages = append(%s.Pages, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.SubChapters {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubChapters = append(%s.SubChapters, %s)", chapterIdent, chapterIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		contentOrdered := []*Content{}
		for content := range stageSet.Stage.Contents {
			contentOrdered = append(contentOrdered, content)
		}
		sort.Slice(contentOrdered, func(i, j int) bool {
			return stageSet.Stage.Content_stagedOrder[contentOrdered[i]] < stageSet.Stage.Content_stagedOrder[contentOrdered[j]]
		})
		for _, content := range contentOrdered {
			contentIdent := "__stage_0" + content.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Content{Name: %s}).Stage(stageSet.Stage)", contentIdent, __gong__toRawStringLiteral(content.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chapters = append(%s.Chapters, %s)", contentIdent, contentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		downloadablefileOrdered := []*DownloadableFile{}
		for downloadablefile := range stageSet.Stage.DownloadableFiles {
			downloadablefileOrdered = append(downloadablefileOrdered, downloadablefile)
		}
		sort.Slice(downloadablefileOrdered, func(i, j int) bool {
			return stageSet.Stage.DownloadableFile_stagedOrder[downloadablefileOrdered[i]] < stageSet.Stage.DownloadableFile_stagedOrder[downloadablefileOrdered[j]]
		})
		for _, downloadablefile := range downloadablefileOrdered {
			downloadablefileIdent := "__stage_0" + downloadablefile.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DownloadableFile{Name: %s}).Stage(stageSet.Stage)", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		jpgimageOrdered := []*JpgImage{}
		for jpgimage := range stageSet.Stage.JpgImages {
			jpgimageOrdered = append(jpgimageOrdered, jpgimage)
		}
		sort.Slice(jpgimageOrdered, func(i, j int) bool {
			return stageSet.Stage.JpgImage_stagedOrder[jpgimageOrdered[i]] < stageSet.Stage.JpgImage_stagedOrder[jpgimageOrdered[j]]
		})
		for _, jpgimage := range jpgimageOrdered {
			jpgimageIdent := "__stage_0" + jpgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.JpgImage{Name: %s}).Stage(stageSet.Stage)", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		pageOrdered := []*Page{}
		for page := range stageSet.Stage.Pages {
			pageOrdered = append(pageOrdered, page)
		}
		sort.Slice(pageOrdered, func(i, j int) bool {
			return stageSet.Stage.Page_stagedOrder[pageOrdered[i]] < stageSet.Stage.Page_stagedOrder[pageOrdered[j]]
		})
		for _, page := range pageOrdered {
			pageIdent := "__stage_0" + page.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Page{Name: %s}).Stage(stageSet.Stage)", pageIdent, __gong__toRawStringLiteral(page.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pageIdent, __gong__toRawStringLiteral(page.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", pageIdent, __gong__toRawStringLiteral(page.MardownContent)))
			for _, elem := range page.Sections {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", pageIdent, pageIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		pngimageOrdered := []*PngImage{}
		for pngimage := range stageSet.Stage.PngImages {
			pngimageOrdered = append(pngimageOrdered, pngimage)
		}
		sort.Slice(pngimageOrdered, func(i, j int) bool {
			return stageSet.Stage.PngImage_stagedOrder[pngimageOrdered[i]] < stageSet.Stage.PngImage_stagedOrder[pngimageOrdered[j]]
		})
		for _, pngimage := range pngimageOrdered {
			pngimageIdent := "__stage_0" + pngimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.PngImage{Name: %s}).Stage(stageSet.Stage)", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		sectionOrdered := []*Section{}
		for section := range stageSet.Stage.Sections {
			sectionOrdered = append(sectionOrdered, section)
		}
		sort.Slice(sectionOrdered, func(i, j int) bool {
			return stageSet.Stage.Section_stagedOrder[sectionOrdered[i]] < stageSet.Stage.Section_stagedOrder[sectionOrdered[j]]
		})
		for _, section := range sectionOrdered {
			sectionIdent := "__stage_0" + section.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Section{Name: %s}).Stage(stageSet.Stage)", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", sectionIdent, __gong__toRawStringLiteral(section.MardownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImage = %t", sectionIdent, section.IsImage))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDownloadableFile = %t", sectionIdent, section.IsDownloadableFile))
			if section.SvgImage != nil {
				targetIdent := "__stage_0" + section.SvgImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SvgImage = %s", sectionIdent, targetIdent))
			}
			if section.PngImage != nil {
				targetIdent := "__stage_0" + section.PngImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PngImage = %s", sectionIdent, targetIdent))
			}
			if section.JpgImage != nil {
				targetIdent := "__stage_0" + section.JpgImage.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.JpgImage = %s", sectionIdent, targetIdent))
			}
			if section.DownloadableFile != nil {
				targetIdent := "__stage_0" + section.DownloadableFile.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DownloadableFile = %s", sectionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		svgimageOrdered := []*SvgImage{}
		for svgimage := range stageSet.Stage.SvgImages {
			svgimageOrdered = append(svgimageOrdered, svgimage)
		}
		sort.Slice(svgimageOrdered, func(i, j int) bool {
			return stageSet.Stage.SvgImage_stagedOrder[svgimageOrdered[i]] < stageSet.Stage.SvgImage_stagedOrder[svgimageOrdered[j]]
		})
		for _, svgimage := range svgimageOrdered {
			svgimageIdent := "__stage_0" + svgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SvgImage{Name: %s}).Stage(stageSet.Stage)", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Content)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/ssg/go/models"
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
				case "Chapter":
					if !preserveOrder {
						inst := (&Chapter{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Chapter)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Content":
					if !preserveOrder {
						inst := (&Content{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Content)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DownloadableFile":
					if !preserveOrder {
						inst := (&DownloadableFile{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DownloadableFile)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "JpgImage":
					if !preserveOrder {
						inst := (&JpgImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(JpgImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Page":
					if !preserveOrder {
						inst := (&Page{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Page)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PngImage":
					if !preserveOrder {
						inst := (&PngImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PngImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Section":
					if !preserveOrder {
						inst := (&Section{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Section)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SvgImage":
					if !preserveOrder {
						inst := (&SvgImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SvgImage)
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
				case *Chapter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "Sections":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Section); ok {
										inst.Sections = append(inst.Sections, typedTarget)
									}
								}
							}
						}
					case "Pages":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Page); ok {
										inst.Pages = append(inst.Pages, typedTarget)
									}
								}
							}
						}
					case "SubChapters":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Chapter); ok {
										inst.SubChapters = append(inst.SubChapters, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Chapter); ok {
										inst.Chapters = append(inst.Chapters, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Section); ok {
										inst.Sections = append(inst.Sections, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SvgImage); ok {
									inst.SvgImage = typedTarget
								}
							}
						}
					case "PngImage":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PngImage); ok {
									inst.PngImage = typedTarget
								}
							}
						}
					case "JpgImage":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*JpgImage); ok {
									inst.JpgImage = typedTarget
								}
							}
						}
					case "IsDownloadableFile":
						inst.IsDownloadableFile = GongExtractBool(rhs)
					case "DownloadableFile":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DownloadableFile); ok {
									inst.DownloadableFile = typedTarget
								}
							}
						}
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
