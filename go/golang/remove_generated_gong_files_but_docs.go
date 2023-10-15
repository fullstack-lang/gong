package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/go/golang/models"
)

// RemoveGeneratedGongFilesButDocs generates the setup file for the gorm
func RemoveGeneratedGongFilesButDocs(
	RelativePkgPath string) {

	{
		filename := filepath.Join(RelativePkgPath, "docs.go")
		file, err := os.Create(filename)

		if err != nil {
			file, err = os.Create(filename)
		}

		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, models.DefaultModelDocsTemplate)

	}

	filesToRemove := []string{
		"gong.go",
		"../..embed.go",
		"../embed.go",
		"gong_coder.go",
		"gong_ast.go",
		"gong_serialize.go",
		"gong_slices.go",
		"gong_marshall.go",
		"gong_graph.go",
		"gong_enum.go",
		"gong_callbacks.go",
		"gong_orchestrator.go",
		"gong_wop.go",
	}

	for _, file := range filesToRemove {
		removeFile(filepath.Join(RelativePkgPath, file))
	}
}

func removeFile(filename string) {
	log.Println("removing file : " + filename)

	if err := os.Remove(filename); err != nil {
		if os.IsExist(err) {
			log.Fatalf("Unable to remove %s", filename)
		}
	}
}
