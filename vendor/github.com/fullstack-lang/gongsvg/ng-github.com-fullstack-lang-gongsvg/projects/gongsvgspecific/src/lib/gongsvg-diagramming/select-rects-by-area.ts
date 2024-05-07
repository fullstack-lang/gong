
import { SelectAreaConfig, SweepDirection } from "../svg-event.service";
import { GongsvgDiagrammingComponent } from "./gongsvg-diagramming"

import * as gongsvg from '../../../../gongsvg/src/public-api'
import { getFunctionName } from "./get.function.name";

export function selectRectsByArea(gongsvgDiagrammingComponent: GongsvgDiagrammingComponent) {
    let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

    if (gongsvgDiagrammingComponent.PointAtMouseUp.X > gongsvgDiagrammingComponent.PointAtMouseDown.X) {
        selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
    } else {
        selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
    }

    selectAreaConfig.TopLeft = [
        Math.min(gongsvgDiagrammingComponent.PointAtMouseDown.X, gongsvgDiagrammingComponent.PointAtMouseUp.X),
        Math.min(gongsvgDiagrammingComponent.PointAtMouseDown.Y, gongsvgDiagrammingComponent.PointAtMouseUp.Y)]
    selectAreaConfig.BottomRigth = [
        Math.max(gongsvgDiagrammingComponent.PointAtMouseDown.X, gongsvgDiagrammingComponent.PointAtMouseUp.X),
        Math.max(gongsvgDiagrammingComponent.PointAtMouseDown.Y, gongsvgDiagrammingComponent.PointAtMouseUp.Y)]

    for (let layer of gongsvgDiagrammingComponent.gongsvgFrontRepo?.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)!) {
        for (let rect of layer.Rects) {
            switch (selectAreaConfig.SweepDirection) {
                case SweepDirection.LEFT_TO_RIGHT:
                    if (rect.X > selectAreaConfig.TopLeft[0] &&
                        rect.X + rect.Width < selectAreaConfig.BottomRigth[0] &&
                        rect.Y > selectAreaConfig.TopLeft[1] &&
                        rect.Y + rect.Height < selectAreaConfig.BottomRigth[1]) {
                        if (!rect.IsSelected) {
                            gongsvgDiagrammingComponent.selectRect(rect);
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
                            gongsvgDiagrammingComponent.manageHandles(rect)
                            gongsvgDiagrammingComponent.rectService.updateFront(rect, gongsvgDiagrammingComponent.GONG__StackPath).subscribe(
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