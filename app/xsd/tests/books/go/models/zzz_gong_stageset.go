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
		booktypeOrdered := []*BookType{}
		for booktype := range stageSet.Stage.BookTypes {
			booktypeOrdered = append(booktypeOrdered, booktype)
		}
		sort.Slice(booktypeOrdered, func(i, j int) bool {
			return stageSet.Stage.BookType_stagedOrder[booktypeOrdered[i]] < stageSet.Stage.BookType_stagedOrder[booktypeOrdered[j]]
		})
		for _, booktype := range booktypeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			booktypeIdent := "__models" + booktype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BookType{Name: %s}).Stage(stageSet.Stage)", booktypeIdent, __gong__toRawStringLiteral(booktype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Edition = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Edition)))
			values.WriteString(fmt.Sprintf("\n\t%s.Isbn = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Isbn)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bestseller = %t", booktypeIdent, booktype.Bestseller))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Author = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Author)))
			values.WriteString(fmt.Sprintf("\n\t%s.Year = %d", booktypeIdent, booktype.Year))
			values.WriteString(fmt.Sprintf("\n\t%s.Format = %s", booktypeIdent, __gong__toRawStringLiteral(booktype.Format)))
			for _, elem := range booktype.Credit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", booktypeIdent, booktypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		booksOrdered := []*Books{}
		for books := range stageSet.Stage.Bookss {
			booksOrdered = append(booksOrdered, books)
		}
		sort.Slice(booksOrdered, func(i, j int) bool {
			return stageSet.Stage.Books_stagedOrder[booksOrdered[i]] < stageSet.Stage.Books_stagedOrder[booksOrdered[j]]
		})
		for _, books := range booksOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			booksIdent := "__models" + books.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Books{Name: %s}).Stage(stageSet.Stage)", booksIdent, __gong__toRawStringLiteral(books.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", booksIdent, __gong__toRawStringLiteral(books.Name)))
			for _, elem := range books.Book {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Book = append(%s.Book, %s)", booksIdent, booksIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		creditOrdered := []*Credit{}
		for credit := range stageSet.Stage.Credits {
			creditOrdered = append(creditOrdered, credit)
		}
		sort.Slice(creditOrdered, func(i, j int) bool {
			return stageSet.Stage.Credit_stagedOrder[creditOrdered[i]] < stageSet.Stage.Credit_stagedOrder[creditOrdered[j]]
		})
		for _, credit := range creditOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			creditIdent := "__models" + credit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Credit{Name: %s}).Stage(stageSet.Stage)", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page = %d", creditIdent, credit.Page))
			values.WriteString(fmt.Sprintf("\n\t%s.Credit_type = %s", creditIdent, __gong__toRawStringLiteral(credit.Credit_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Credit_words = %s", creditIdent, __gong__toRawStringLiteral(credit.Credit_words)))
			values.WriteString(fmt.Sprintf("\n\t%s.Credit_symbol = %s", creditIdent, __gong__toRawStringLiteral(credit.Credit_symbol)))
			for _, elem := range credit.Link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", creditIdent, creditIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		linkOrdered := []*Link{}
		for link := range stageSet.Stage.Links {
			linkOrdered = append(linkOrdered, link)
		}
		sort.Slice(linkOrdered, func(i, j int) bool {
			return stageSet.Stage.Link_stagedOrder[linkOrdered[i]] < stageSet.Stage.Link_stagedOrder[linkOrdered[j]]
		})
		for _, link := range linkOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkIdent := "__models" + link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Link{Name: %s}).Stage(stageSet.Stage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkIdent, __gong__toRawStringLiteral(link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", linkIdent, __gong__toRawStringLiteral(link.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", linkIdent, __gong__toRawStringLiteral(link.EnclosedText)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/books/go/models"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/books/go/models":
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
				case "BookType":
					if !preserveOrder {
						inst := (&BookType{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(BookType)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Books":
					if !preserveOrder {
						inst := (&Books{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Books)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Credit":
					if !preserveOrder {
						inst := (&Credit{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Credit)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Link":
					if !preserveOrder {
						inst := (&Link{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Link)
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
				case *BookType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Edition":
						inst.Edition = GongExtractString(rhs)
					case "Isbn":
						inst.Isbn = GongExtractString(rhs)
					case "Bestseller":
						inst.Bestseller = GongExtractBool(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Author":
						inst.Author = GongExtractString(rhs)
					case "Year":
						inst.Year = GongExtractInt(rhs)
					case "Format":
						inst.Format = GongExtractString(rhs)
					case "Credit":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Credit); ok {
										inst.Credit = append(inst.Credit, typedTarget)
									}
								}
							}
						}
					}
				case *Books:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Book":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*BookType); ok {
										inst.Book = append(inst.Book, typedTarget)
									}
								}
							}
						}
					}
				case *Credit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page":
						inst.Page = GongExtractInt(rhs)
					case "Credit_type":
						inst.Credit_type = GongExtractString(rhs)
					case "Link":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Link); ok {
										inst.Link = append(inst.Link, typedTarget)
									}
								}
							}
						}
					case "Credit_words":
						inst.Credit_words = GongExtractString(rhs)
					case "Credit_symbol":
						inst.Credit_symbol = GongExtractString(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
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
