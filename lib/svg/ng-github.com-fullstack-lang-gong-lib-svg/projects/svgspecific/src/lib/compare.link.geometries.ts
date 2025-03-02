
import * as svg from '../../../svg/src/public-api'


export function compareLinkGeometries(instance1: svg.Link, instance2: svg.Link): boolean {

    if (instance1 == undefined || instance2 == undefined) {
        console.log("comparing undefined links")
    }

    if (
        instance1.StartRatio != instance2.StartRatio ||
        instance1.StartOrientation != instance2.StartOrientation ||
        instance1.EndRatio != instance2.EndRatio ||
        instance1.EndOrientation != instance2.EndOrientation ||
        instance1.CornerOffsetRatio != instance2.CornerOffsetRatio
    ) {
        return false
    }
    return true;
}