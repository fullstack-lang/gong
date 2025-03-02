

import { SvgSpecificComponent } from "./svg-specific/svg-specific.component"

import * as svg from '../../../svg/src/public-api'

// informBackEndOfEndOfLinkDrawing
//
// informs the back end with 2 updates
// on the first update, the svg is updated with the state DRAWING_LINK
// on the second, the svg is updated with the state NOT_DRAWING_LINK
//
// the back ends shall interpret those calls in order to create the link between

// start end end rects
export function informBackEndOfEndOfLinkDrawing(svgSpecificComponent: SvgSpecificComponent) {
    svgSpecificComponent.svg.DrawingState = svg.DrawingState.DRAWING_LINK
    svgSpecificComponent.svgService.updateFront(svgSpecificComponent.svg, svgSpecificComponent.GONG__StackPath).subscribe(
        () => {
            // back to normal state
            svgSpecificComponent.svg.DrawingState = svg.DrawingState.NOT_DRAWING_LINK;
            svgSpecificComponent.svgService.updateFront(svgSpecificComponent.svg, svgSpecificComponent.GONG__StackPath).subscribe();

            // set the isEditable
            svgSpecificComponent.isEditableService.setIsEditable(svgSpecificComponent.svg!.IsEditable);
        }
    );
}