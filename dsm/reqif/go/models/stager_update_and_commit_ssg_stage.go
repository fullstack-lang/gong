package models

import (
	"fmt"
	"log"
	"path/filepath"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

const figureTemplate = `
<figure style="
  width: %dpx;
  height: %dpx;
">
  <img src="%s" alt="%s" style="
	width: 100%%;
	height: 100%%;
	object-fit: cover;
	object-position: center top;
  ">
  <figcaption>
	</figcaption>
</figure>
`

func (stager *Stager) UpdateAndCommitSsgStage() {

	stager.ssgStage.Reset()

	var siteWeb *StaticWebSite
	{
		siteWebs := GetGongstructInstancesSet[StaticWebSite](stager.stage)
		if len(*siteWebs) != 1 {
			log.Fatalln("There should be one siteWeb")
		}
		for k := range *siteWebs {
			siteWeb = k
		}
		_ = siteWeb
	}

	content := &ssg.Content{
		Name:           siteWeb.Name,
		MardownContent: siteWeb.MarkdownContent,

		LayoutPath: filepath.Join(stager.pathToExtractedGongSsgDefaultLayoutFiles, "layouts"),
		StaticPath: filepath.Join(stager.pathToFilesUsedForSsgGeneration, "static"),

		ContentPath: stager.pathToFilesForSsgGeneratedMarkdownFiles,
		OutputPath:  filepath.Join(stager.rootPathToStaticOutputs, siteWeb.OutputStaticWebDir),

		Target: ssg.FILE,

		VersionInfo: siteWeb.VersionInfo,
	}

	for _, chapter := range siteWeb.Chapters {

		ssgChapter := &ssg.Chapter{
			Name:           chapter.Name,
			MardownContent: chapter.MarkdownContent,
		}

		content.Chapters = append(content.Chapters,
			ssgChapter,
		)

		for _, paragraph := range chapter.Paragraphs {
			if paragraph.Image != nil {

				// it is important that the image starts at the first column
				ssgChapter.MardownContent += fmt.Sprintf(figureTemplate,
					paragraph.Image.Width,
					paragraph.Image.Height,

					filepath.Join("../images/", paragraph.Image.Name),
					paragraph.LegendMarkdownContent,
				)
			}
		}

	}

	ssg.StageBranch(stager.ssgStage, content)

	stager.ssgStage.Commit()

	// for image := range *GetGongstructInstancesSet[Image](stager.stage) {
	// 	image.Unstage(stager.stage)
	// }
}
