package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/go/golang/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
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

	for _, file := range gong_models.GeneratedModelFiles {
		removeFile(filepath.Join(RelativePkgPath, file))
	}
}

func removeFile(filename string) {
	// log.Println("removing file : " + filename)

	if err := os.Remove(filename); err != nil {
		if os.IsExist(err) {
			log.Fatalf("Unable to remove %s", filename)
		}
	}
}
