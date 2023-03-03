import * as gongdoc from 'gongdoc'


// onMove is called each time the shape is moved
export function onClassshapeMove(umlClassShape: joint.shapes.uml.Class) {
    // console.log(umlClassShape.id, ':', umlClassShape.get('position'));

    let classhape = umlClassShape.attributes['classshape'] as gongdoc.GongStructShapeDB
    let positionService = umlClassShape.attributes['positionService'] as gongdoc.PositionService
    let position = classhape.Position
    position!.X = umlClassShape.get('position')!.x
    position!.Y = umlClassShape.get('position')!.y

    positionService.updatePosition(position!, "").subscribe(
        position => {
            // console.log("position updated")
        }
    )
}

// onMove is called each time the shape is moved
export function onNoteMove(umlNote: joint.shapes.standard.Rectangle) {


    let note = umlNote.attributes['note'] as gongdoc.NoteShapeDB
    let noteService = umlNote.attributes['noteService'] as gongdoc.NoteShapeService
    note.X = umlNote.get('position')!.x
    note.Y = umlNote.get('position')!.y
    noteService.updateNoteShape(note!, "").subscribe(
        note => {

        }
    )
    // console.log(note.Name, ':', umlNote.get('position'));
}

// onMove is called each time the shape is moved
export function onLinkMove(standardLink: joint.shapes.standard.Link) {
    // console.log(standardLink.id, ':', standardLink.get('vertices'));

    let middleVertice = standardLink.attributes['middleVertice'] as gongdoc.VerticeDB
    let verticeService = standardLink.attributes['verticeService'] as gongdoc.VerticeService

    middleVertice!.X = standardLink.get('vertices')![0].x
    middleVertice!.Y = standardLink.get('vertices')![0].y

    verticeService.updateVertice(middleVertice!, "").subscribe(
        middleVertice => {
            // console.log("middleVertice updated")
        }
    )
}