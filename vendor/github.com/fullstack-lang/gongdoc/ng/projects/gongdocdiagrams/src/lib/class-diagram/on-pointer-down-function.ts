import * as gongdoc from 'gongdoc'

export function informBackEndOfSelection(cellView: joint.dia.CellView) {

    let umlClassShape = cellView.model

    let gongStructShape = umlClassShape.attributes['classshape'] as gongdoc.GongStructShapeDB
    let gongStructShapeService = umlClassShape.attributes['gongStructShapeService'] as gongdoc.GongStructShapeService

    // if selected object is not a classshape, move on
    if (gongStructShape == undefined) {
        return
    }

    gongStructShape.IsSelected = true
    gongStructShapeService.updateGongStructShape(gongStructShape, "").subscribe(
        classhape => {
            console.log("gongStructShape updated")
        }
    )
}