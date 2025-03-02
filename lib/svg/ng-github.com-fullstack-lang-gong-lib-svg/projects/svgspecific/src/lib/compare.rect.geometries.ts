
import * as svg from '../../../svg/src/public-api'


export function compareRectGeometries(instance1: svg.Rect, instance2: svg.Rect): boolean {

    // during initialization phases, some arguments might be undefined
    if (instance1 == undefined || instance2 == undefined) {
        console.log("comparing undefined rect")
        return true
    }

    if (
        instance1.X != instance2.X ||
        instance1.Y != instance2.Y ||
        instance1.Width != instance2.Width ||
        instance1.Height != instance2.Height
    ) {
        return false
    }
    return true
}