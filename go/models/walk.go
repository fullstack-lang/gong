package models

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	// to parse the .frontignore file
	"github.com/fullstack-lang/gong/go/ignore"
)

// Walk parses go files in the `models` directory pointed by relativePathToModel and
// fills up modelPkg with its Gongstruct & Gongenum
//
// if useParse is used, Walk uses go/parser.Parser
//
// # Walk leverages go Parser capabilities to fetch identifiers in go files
//
// The algo is in several steps:
// - First pass gather Gongstruct & Gongenums identifiers
// - Second pass parses fields and link them to other Gongstructs
func Walk(relativePathToModel string, modelPkg *ModelPkg) {

	directory, err := filepath.Abs(relativePathToModel)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset, directory, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ", errParser.Error())
	}
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}

	filepath := filepath.Join(relativePathToModel, ".frontignore")

	// Parse the gitignore file
	goGitignoreEntries, err := ignore.ParseGoGitignore(filepath)
	_ = goGitignoreEntries

	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			// log.Println("No .frontignore file present")
		} else {
			log.Fatalf("Failed to compile .frontignore: %v", err)
		}
	}

	WalkParser(pkgsParser, modelPkg, &goGitignoreEntries)

	// fetch meta information
	inspectMeta(modelPkg.GetStage(), pkgsParser["models"])
}
