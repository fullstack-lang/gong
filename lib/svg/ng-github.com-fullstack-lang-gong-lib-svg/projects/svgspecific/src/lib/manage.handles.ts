import * as svg from '../../../svg/src/public-api'

// display or not handles if selected or not
export function manageHandles(rect: svg.Rect) {

    if (rect.IsSelected && rect.CanHaveLeftHandle) {
        rect.HasLeftHandle = true
    } else {
        rect.HasLeftHandle = false
    }
    if (rect.IsSelected && rect.CanHaveRightHandle) {
        rect.HasRightHandle = true
    } else {
        rect.HasRightHandle = false
    }
    if (rect.IsSelected && rect.CanHaveTopHandle) {
        rect.HasTopHandle = true
    } else {
        rect.HasTopHandle = false
    }
    if (rect.IsSelected && rect.CanHaveBottomHandle) {
        rect.HasBottomHandle = true
    } else {
        rect.HasBottomHandle = false
    }
}