package models

import markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"

func (stager *Stager) ux_markdown() {

	stage := stager.markdownStage

	stage.Reset()

	content := markdown.Content{
		Content: `# BarrGraph

**BarrGraph** is a specialized visualization tool designed for art curators, historians, and researchers. It enables the creation of chronological flowcharts to map the genealogy and evolution of art movements, inspired by the methodology of **Alfred Barr**.

## Inspiration

Alfred H. Barr Jr., the first director of the Museum of Modern Art (MoMA), famously used flowcharts to explain the origins and development of abstract art. His diagram for the 1936 exhibition *Cubism and Abstract Art* remains a canonical example of data visualization in art history.

**BarrGraph** digitizes this approach, allowing users to construct directed graphs where nodes represent art movements and edges represent influence, all plotted against a chronological axis.

## Features

- **Directed Influence Maps:** Model the flow of influence between movements (e.g., Cézanne → Cubism).
- **Chronological Alignment:** Automatically align movements by their active decades.
- **Curatorial Metadata:** Attach descriptions, key artists, and representative works to each node.
- **Visual Styling:** Export charts that pay homage to the classic mid-century aesthetic of Barr's original posters.

## Abstract Syntax

The application is built around a domain-specific model that captures the relationships between art movements.

![Abstract Syntax](svg:abstractSyntax?width=1000px)

*The abstract syntax defining ArtMovements and their Influences.*

The model consists of:
- **ArtMovement**: Represents a distinct artistic movement or period.
- **ArtefactType**: A categorization of art (e.g., Painting, Sculpture, Architecture).
- **Artist**: Individual artists associated with movements.
- **Influence**: Represents a directed relationship where one movement influences another. 
An "Influence" links a **Source** "ArtMovement/ArtefactType/Artist" 
to a **Target** "ArtMovement/ArtefactType/Artist".

## Gallery

Below are examples of the tool in use, including the original inspiration and the digital recreation.

![Alfred Barr's Original Poster, redrawn](svg:barrgraph?width=800px)

![Renaissance Art Movements](svg:renaissance?width=800px)

## Installation

go run github.com/fullstack-lang/gong/dsm/barrgraph/go/cmd/barrgraph@latest

then open a browser on http://localhost:8080/.

## License

Distributed under the MIT License.
		
		
`,
	}
	content.Stage(stage)

	{
		svgImage := markdown.SvgImage{
			Name:    "barrgraph",
			Content: svgCubism,
		}
		svgImage.Stage(stage)
	}

	{
		svgImage := markdown.SvgImage{
			Name:    "renaissance",
			Content: renaissance,
		}
		svgImage.Stage(stage)
	}

	{
		svgImage := markdown.SvgImage{
			Name:    "abstractSyntax",
			Content: abstractSyntax,
		}
		svgImage.Stage(stage)
	}

	stage.Commit()
}
