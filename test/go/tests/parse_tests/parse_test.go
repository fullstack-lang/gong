package parsetests_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/test/go/fullstack"
	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/gin-gonic/gin"
)

func TestParseTest(t *testing.T) {

	// setup GORM
	r := gin.Default()
	fullstack.Init(r)

	directory, err := filepath.Abs("./input")
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset, directory, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ", errParser.Error())
	}
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}

	pkg, ok := pkgsParser["main"]
	if !ok {
		log.Fatalln("wrong package name")
	}

	for fileName, file := range pkg.Files {

		isOfInterest := strings.Contains(fileName, "stage")
		if !isOfInterest {
			continue
		}
		log.Println("AST File: ", fileName)
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				funcDecl := decl
				if name := funcDecl.Name; name != nil {
					isOfInterest := strings.Contains(funcDecl.Name.Name, "stage")
					if !isOfInterest {
						continue
					}
					log.Println("\tAST FuncDecl: ", name.Name)
				}
				if doc := funcDecl.Doc; doc != nil {
					for _, comment := range doc.List {
						log.Println("\t\tAST Comment: ", comment.Text)
					}
				}
				if body := funcDecl.Body; body != nil {
					for _, stmt := range body.List {
						switch stmt := stmt.(type) {
						case *ast.ExprStmt:
							exprStmt := stmt
							switch expr := exprStmt.X.(type) {
							case *ast.CallExpr:
								callExpr := expr
								switch fun := callExpr.Fun.(type) {
								case *ast.Ident:
									ident := fun
									log.Println("\t\t\tAST CallExpr: ", ident.Name)
								}
							}
						case *ast.AssignStmt:
							log.Println("\t\t\tAST Assigment: ")
							assignStmt := stmt
							instance, id := models.UnmarshallGongstructStaging[models.Astruct](assignStmt)
							_ = instance
							_ = id

						}
					}
				}
			case *ast.GenDecl:
				genDecl := decl
				log.Println("\tAST GenDecl: ")
				if doc := genDecl.Doc; doc != nil {
					for _, comment := range doc.List {
						log.Println("\tAST Comment: ", comment.Text)
					}
				}
				for _, spec := range genDecl.Specs {
					switch spec := spec.(type) {
					case *ast.ImportSpec:
						importSpec := spec
						if path := importSpec.Path; path != nil {
							log.Println("\t\tAST Path: ", path.Value)
						}
					}
				}
			}

		}
	}

	log.Println()
}
