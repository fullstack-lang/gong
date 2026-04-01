package specifications

import (

	// Corrected path
	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// UpdateAndCommitSpecificationsMarkdownStage implements models.SpecificationsTreeUpdaterInterface.
func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsMarkdownStage(stager *m.Stager) {
	stage := stager.GetStage()

	markdownStage := stager.GetMarkdownStage()
	markdownStage.Reset()

	if stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS == nil {
		markdownStage.Commit()
		return
	}

	selectedSpecification := GetSelectedSpecification(stage)
	if selectedSpecification == nil {
		markdownStage.Commit()
		return
	}

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS.SPECIFICATION
	for _, specification := range specifications {

		if selectedSpecification != specification {
			continue
		}

		// --- updated logic to generate and assign markdown content ---

		// 1. Initialize markdown content string
		// markDownContent := "**Specification \"" + specification.Name + "\"**\n\n"
		markDownContent := ""

		// 2. A dummy parent node is created because processSpecHierarchy expects a parent
		// to append children to. This node is temporary and will be discarded.
		hierarchyParentNode := &tree.Node{}
		depth := 1 // initial depth for chapters

		// 3. Recursively process spec hierarchies to build the markdown string
		for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchy,
				hierarchyParentNode,
				depth,
				&markDownContent)
		}
		// --- end of update ---

		content := &markdown.Content{
			Name:    "*" + specification.GetName() + "*",
			Content: markDownContent, // Assign the generated markdown
		}

		for _, svgImage := range models.GetGongstrucsSorted[*models.EmbeddedSvgImage](stager.GetStage()) {
			markdownSvgImage := new(markdown.SvgImage).Stage(markdownStage)
			markdownSvgImage.Content = svgImage.Content
			markdownSvgImage.Name = svgImage.Name
		}

		for _, jpgImage := range models.GetGongstrucsSorted[*models.EmbeddedJpgImage](stager.GetStage()) {
			markdownSvgImage := new(markdown.JpgImage).Stage(markdownStage)
			markdownSvgImage.Base64Content = jpgImage.Base64Content
			markdownSvgImage.Name = jpgImage.Name
		}

		for _, pngImage := range models.GetGongstrucsSorted[*models.EmbeddedPngImage](stager.GetStage()) {
			markdownSvgImage := new(markdown.PngImage).Stage(markdownStage)
			markdownSvgImage.Base64Content = pngImage.Base64Content
			markdownSvgImage.Name = pngImage.Name
		}

		markdown.StageBranch(markdownStage, content)

		if GetSpecificationRendering(stage, selectedSpecification).IsWithHeadingNumbering {
			content.Content = AddHeaderNumbering(content.Content)
		}
	}

	markdownStage.Commit()
}
