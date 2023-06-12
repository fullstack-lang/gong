
import * as gongsvg from 'gongsvg'


export function compareLinkGeometries(instance1: gongsvg.LinkDB, instance2: gongsvg.LinkDB): boolean {

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