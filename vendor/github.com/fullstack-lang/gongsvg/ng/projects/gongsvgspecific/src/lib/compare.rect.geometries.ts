
import * as gongsvg from 'gongsvg'


export function compareRectGeometries(instance1: gongsvg.RectDB, instance2: gongsvg.RectDB): boolean {

    if (instance1 == undefined || instance2 == undefined) {
        console.log("comparing undefined rect")
    }

    if (
        instance1.X != instance2.X ||
        instance1.Y != instance2.Y ||
        instance1.Width != instance2.Width ||
        instance1.Height != instance2.Height
    ) {
        return false
    }
    return true;
}