package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	diagram := stager.desk.SelectedDiagram
	svgObject := stager.svgObject

	if svgObject.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connexion
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	var sourceArtElement, targetArtElement ArtElement
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	// determine the association type from the source and target rects
	if movementShape, ok := diagram.map_SvgRect_MovementShape[startRect]; ok {
		sourceArtElement = movementShape.Movement
	} else if artefactTypeShape, ok := diagram.map_SvgRect_ArtefactTypeShape[startRect]; ok {
		sourceArtElement = artefactTypeShape.ArtefactType
	} else if artistShape, ok := diagram.map_SvgRect_ArtistShape[startRect]; ok {
		sourceArtElement = artistShape.Artist
	}

	if movementShape, ok := diagram.map_SvgRect_MovementShape[endRect]; ok {
		targetArtElement = movementShape.Movement
	} else if artefactTypeShape, ok := diagram.map_SvgRect_ArtefactTypeShape[endRect]; ok {
		targetArtElement = artefactTypeShape.ArtefactType
	} else if artistShape, ok := diagram.map_SvgRect_ArtistShape[endRect]; ok {
		targetArtElement = artistShape.Artist
	}

	if sourceArtElement == nil || targetArtElement == nil {
		log.Println("unsupported association: missing source or target")
		return
	}

	// create the influence
	influence := new(Influence).Stage(stager.stage)
	influence.Name = sourceArtElement.GetName() + " to " + targetArtElement.GetName()
	influence.source = sourceArtElement
	influence.target = targetArtElement

	if movement, ok := sourceArtElement.(*Movement); ok {
		influence.SourceMovement = movement
	} else if artefactType, ok := sourceArtElement.(*ArtefactType); ok {
		influence.SourceArtefactType = artefactType
	} else if artist, ok := sourceArtElement.(*Artist); ok {
		influence.SourceArtist = artist
	}

	if movement, ok := targetArtElement.(*Movement); ok {
		influence.TargetMovement = movement
	} else if artefactType, ok := targetArtElement.(*ArtefactType); ok {
		influence.TargetArtefactType = artefactType
	} else if artist, ok := targetArtElement.(*Artist); ok {
		influence.TargetArtist = artist
	}

	influenceShape := new(InfluenceShape).Stage(stager.stage)
	influenceShape.Name = influence.Name
	influenceShape.Influence = influence

	diagram.InfluenceShapes = append(diagram.InfluenceShapes, influenceShape)

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
