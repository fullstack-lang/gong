package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fullstack-lang/gong/go/golang/models"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// RemoveGeneratedGongFilesButDocs generates the setup file for the gorm
func RemoveGeneratedGongFilesButDocs(
	RelativePkgPath string) {

	{
		filename := filepath.Join(RelativePkgPath, string(gong_models.DocsGoFilePath))
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			absPath, _ := filepath.Abs(RelativePkgPath)
			pkgName := filepath.Base(absPath)
			file, err := os.Create(filename)
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprintf(file, models.DefaultModelDocsTemplate, pkgName)
		}
	}

	for _, file := range gong_models.GeneratedModelFiles {
		removeFile(filepath.Join(RelativePkgPath, file))
	}
}

// RemoveGeneratedSubModelPackageGongFiles removes only generated files within the sub-model directory
func RemoveGeneratedSubModelPackageGongFiles(
	RelativePkgPath string) {

	{
		filename := filepath.Join(RelativePkgPath, string(gong_models.DocsGoFilePath))
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			absPath, _ := filepath.Abs(RelativePkgPath)
			pkgName := filepath.Base(absPath)
			file, err := os.Create(filename)
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprintf(file, models.DefaultModelDocsTemplate, pkgName)
		}
	}

	for _, file := range gong_models.GeneratedModelFiles {
		// Do not delete files in parent directories (e.g. ../embed.go)
		if strings.HasPrefix(file, "..") {
			continue
		}
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
