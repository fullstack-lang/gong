package exporter

import (
	"encoding/base64"
	"log"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
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

	contentStr, err := stageForRenderingConf.MarshallToString("github.com/fullstack-lang/gong/app/reqif/go/models", "main")
	if err != nil {
		log.Fatalf("Failed to marshall to string: %v", err)
	}

	content := []byte(contentStr)

	// 5. The content is now a string in memory, and the file will be deleted.
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(content)

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the rendering configuration file", fileToDownload.Name)

}
