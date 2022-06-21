package models

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"time"

	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

// Walk parses go files in the `models` directory pointed by relativePathToModel and
// fills up modelPkg with its Gongstruct & Gongenum
//
// if useParse is used, Walk uses go/parser.Parser
//
// Walk leverages go Parser capabilities to fetch identifiers in go files
//
// The algo is in several steps:
// - First pass gather Gongstruct & Gongenums identifiers
// - Second pass parses fields and link them to other Gongstructs
func Walk(relativePathToModel string, modelPkg *ModelPkg, useParse ...bool) {

	var useParser bool
	if len(useParse) > 1 {
		log.Fatal("Too many args to useParse")
	}
	if len(useParse) == 1 {
		useParser = useParse[0]
	}

	directory, err := filepath.Abs(relativePathToModel)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	// log.Println("Loading package " + directory)

	if useParser {
		fset := token.NewFileSet()
		startParser := time.Now()
		pkgsParser, errParser := parser.ParseDir(fset, directory, nil, parser.ParseComments)
		log.Printf("Parser took %s", time.Since(startParser))

		if errParser != nil {
			log.Panic("Unable to parser ")
		}
		if len(pkgsParser) != 1 {
			log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
		}

		WalkParser(pkgsParser, modelPkg)
	}

	var refModelPkg ModelPkg
	{
		//
		// prepare package load
		//
		cfg := &packages.Config{
			Dir:   directory,
			Mode:  pkgLoadMode,
			Tests: false,
		}
		start := time.Now()
		pkgs, err := packages.Load(cfg, "./...")
		log.Printf("package Load took %s", time.Since(start))

		if err != nil {
			s := fmt.Sprintf("cannot process package at path %s, err %s", relativePathToModel, err.Error())
			log.Panic(s)
		}

		if len(pkgs) != 1 {
			log.Panicf("Expected 1 package to scope, found %d", len(pkgs))
		}
		pkg := pkgs[0]

		WalkLoader(pkg, &refModelPkg)
	}

	log.Println("Walk is over, nb of gongstruct ", len(modelPkg.GongStructs))
}
