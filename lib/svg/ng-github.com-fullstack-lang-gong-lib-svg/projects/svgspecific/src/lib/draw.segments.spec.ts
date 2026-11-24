import * as svg from '../../../svg/src/public-api';
import { drawSegmentsFromLink } from './draw.segments';

describe('drawSegmentsFromLink with control points', () => {
    it('attaches ANCHOR_CENTER start and end ports toward adjacent control points instead of opposite rect', () => {
        // Reproduce port geometry: Start port ("ASD Track") at top-right, End port ("Surv In") at bottom-left.
        // Without control points, Start -> End exits the left edge of Start and enters the right edge of End.
        // With CP0 below Start and CP1 directly above End, Start must exit its bottom edge and End must enter its top edge.
        const startRect = new svg.Rect();
        startRect.ID = 1;
        startRect.X = 800;
        startRect.Y = 80;
        startRect.Width = 24;
        startRect.Height = 24;

        const endRect = new svg.Rect();
        endRect.ID = 2;
        endRect.X = 40;
        endRect.Y = 420;
        endRect.Width = 24;
        endRect.Height = 24;

        // CP0 at (820, 140) -> below and slightly right of startRect center (812, 92)
        const cp0 = new svg.ControlPoint();
        cp0.ClosestRect = startRect;
        cp0.X_Relative = (820 - startRect.X) / startRect.Width;
        cp0.Y_Relative = (140 - startRect.Y) / startRect.Height;

        // CP1 at (56, 370) -> directly above endRect center (52, 432)
        const cp1 = new svg.ControlPoint();
        cp1.ClosestRect = endRect;
        cp1.X_Relative = (56 - endRect.X) / endRect.Width;
        cp1.Y_Relative = (370 - endRect.Y) / endRect.Height;

        const link = new svg.Link();
        link.Type = svg.LinkType.LINK_TYPE_LINE_WITH_CONTROL_POINTS;
        link.Start = startRect;
        link.End = endRect;
        link.StartAnchorType = svg.AnchorType.ANCHOR_CENTER;
        link.EndAnchorType = svg.AnchorType.ANCHOR_CENTER;
        link.ControlPoints = [cp0, cp1];
        link.StartArrowOffset = 0;
        link.EndArrowOffset = 0;
        link.CornerRadius = 0;

        const segments = drawSegmentsFromLink(link);
        expect(segments.length).toBe(3);

        // First segment must exit bottom edge of startRect (Y == 104), not left edge (X == 800)
        expect(segments[0].StartPoint.Y).toBeCloseTo(startRect.Y + startRect.Height, 5);
        expect(segments[0].StartPoint.X).toBeGreaterThan(startRect.X);
        expect(segments[0].StartPoint.X).toBeLessThan(startRect.X + startRect.Width);

        // Last segment must enter top edge of endRect (Y == 420), not right edge (X == 64)
        expect(segments[2].EndPoint.Y).toBeCloseTo(endRect.Y, 5);
        expect(segments[2].EndPoint.X).toBeGreaterThan(endRect.X);
        expect(segments[2].EndPoint.X).toBeLessThan(endRect.X + endRect.Width);
    });
});
