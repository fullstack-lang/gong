

import { GongsvgDiagrammingComponent } from "./gongsvg-diagramming"

import * as gongsvg from 'gongsvg'

// informBackEndOfEndOfLinkDrawing
//
// informs the back end with 2 updates
// on the first update, the svg is updated with the state DRAWING_LINK
// on the second, the svg is updated with the state NOT_DRAWING_LINK
//
// the back ends shall interpret those calls in order to create the link between

// start end end rects
export function informBackEndOfEndOfLinkDrawing(gongsvgDiagrammingComponent: GongsvgDiagrammingComponent) {
    gongsvgDiagrammingComponent.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINK
    gongsvgDiagrammingComponent.svgService.updateSVG(gongsvgDiagrammingComponent.svg, gongsvgDiagrammingComponent.GONG__StackPath, gongsvgDiagrammingComponent.gongsvgFrontRepoService.frontRepo).subscribe(
        () => {

            gongsvgDiagrammingComponent.gongsvgFrontRepoService.pull(gongsvgDiagrammingComponent.GONG__StackPath).subscribe(
                gongsvgsFrontRepo => {
                    gongsvgDiagrammingComponent.gongsvgFrontRepo = gongsvgsFrontRepo;

                    if (gongsvgDiagrammingComponent.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
                        gongsvgDiagrammingComponent.svg = gongsvgDiagrammingComponent.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0];

                        // back to normal state
                        gongsvgDiagrammingComponent.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINK;
                        gongsvgDiagrammingComponent.svgService.updateSVG(gongsvgDiagrammingComponent.svg, gongsvgDiagrammingComponent.GONG__StackPath, gongsvgDiagrammingComponent.gongsvgFrontRepoService.frontRepo).subscribe();

                        // set the isEditable
                        gongsvgDiagrammingComponent.isEditableService.setIsEditable(gongsvgDiagrammingComponent.svg!.IsEditable);
                    } else {
                        return;
                    }
                }
            );
        }
    );
}