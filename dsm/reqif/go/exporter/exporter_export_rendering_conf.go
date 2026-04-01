package exporter

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
)

func (exporter *Exporter) ExportRenderingConf(stager *models.Stager) {

	log.Println("Exporting the rendering configuration")

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	fileToDownload.Name = strings.TrimSuffix(stager.PathToReqifFile, ".reqif") + "-renderingConf.go"

	// 0. we create a new stage for just the marshall of the rendering configuration
	stageForRenderingConf := models.NewStage("renderingConf")

	stage := stager.GetStage()
	models.StageAllOfTypeToAnotherStage[*models.SPECIFICATION_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.SPEC_OBJECT_TYPE_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_DATE_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_REAL_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_STRING_Rendering](stage, stageForRenderingConf)
	models.StageAllOfTypeToAnotherStage[*models.ATTRIBUTE_DEFINITION_XHTML_Rendering](stage, stageForRenderingConf)

	fileName := filepath.Base(fileToDownload.Name)

	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}

	// 2. Defer the removal of the file to ensure it's cleaned up.
	defer os.Remove(tempFile.Name())

	// 3. Marshall the data into the temporary file.
	stageForRenderingConf.Marshall(tempFile, "github.com/fullstack-lang/gong/dsm/reqif/go/models", "main")

	// 4. Read the content back from the file.
	// os.ReadFile needs a file path, so we use the name.
	// First, ensure the file is closed so all data is flushed to disk.
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Failed to close temp file: %v", err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		log.Fatalf("Failed to read back temp file: %v", err)
	}

	// 5. The content is now a string in memory, and the file will be deleted.
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(content)

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the rendering configuration file", tempFile.Name())

}
