
import { SelectAreaConfig, SweepDirection } from "./svg-event.service";
import { SvgSpecificComponent } from "./svg-specific/svg-specific.component"

import * as svg from '../../../svg/src/public-api'
import { getFunctionName } from "./get.function.name";

export function selectRectsByArea(svgSpecificComponent: SvgSpecificComponent) {
    let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

    if (svgSpecificComponent.PointAtMouseUp.X > svgSpecificComponent.PointAtMouseDown.X) {
        selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
    } else {
        selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
    }

    selectAreaConfig.TopLeft = [
        Math.min(svgSpecificComponent.PointAtMouseDown.X, svgSpecificComponent.PointAtMouseUp.X),
        Math.min(svgSpecificComponent.PointAtMouseDown.Y, svgSpecificComponent.PointAtMouseUp.Y)]
    selectAreaConfig.BottomRigth = [
        Math.max(svgSpecificComponent.PointAtMouseDown.X, svgSpecificComponent.PointAtMouseUp.X),
        Math.max(svgSpecificComponent.PointAtMouseDown.Y, svgSpecificComponent.PointAtMouseUp.Y)]

    for (let layer of svgSpecificComponent.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
        for (let rect of layer.Rects) {
            switch (selectAreaConfig.SweepDirection) {
                case SweepDirection.LEFT_TO_RIGHT:
                    if (rect.X > selectAreaConfig.TopLeft[0] &&
                        rect.X + rect.Width < selectAreaConfig.BottomRigth[0] &&
                        rect.Y > selectAreaConfig.TopLeft[1] &&
                        rect.Y + rect.Height < selectAreaConfig.BottomRigth[1]) {
                        if (!rect.IsSelected) {
                            svgSpecificComponent.selectRect(rect);
                        }
                    }
                    break;
                case SweepDirection.RIGHT_TO_LEFT:
                    // rectangle has to be partialy boxed in the rect
                    if (rect.X < selectAreaConfig.BottomRigth[0] &&
                        rect.X + rect.Width > selectAreaConfig.TopLeft[0] &&
                        rect.Y < selectAreaConfig.BottomRigth[1] &&
                        rect.Y + rect.Height > selectAreaConfig.TopLeft[1]) {
                        console.log(getFunctionName(), "selecting rect", rect.Name);
                        if (!rect.IsSelected) {
                            console.log(getFunctionName(), "selecting rect", rect.Name);
                            rect.IsSelected = true;
                            svgSpecificComponent.manageHandles(rect)
                            svgSpecificComponent.rectService.updateFront(rect, svgSpecificComponent.GONG__StackPath).subscribe(
                                _ => {
                                }
                            );
                        }
                    }
                    break;
            }
        }
    }
}