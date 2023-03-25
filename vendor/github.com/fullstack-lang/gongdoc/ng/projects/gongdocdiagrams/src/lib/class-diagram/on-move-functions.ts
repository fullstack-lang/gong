import * as gongdoc from 'gongdoc'


// onMove is called each time the shape is moved
export function onClassshapeMove(umlClassShape: joint.shapes.uml.Class) {
    // console.log(umlClassShape.id, ':', umlClassShape.get('position'));

    let classhape = umlClassShape.attributes['classshape'] as gongdoc.GongStructShapeDB
    let positionService = umlClassShape.attributes['positionService'] as gongdoc.PositionService
    let GONG__StackPath = umlClassShape.attributes['GONG__StackPath'] 

    let position = classhape.Position

    position!.X = umlClassShape.get('position')!.x
    position!.Y = umlClassShape.get('position')!.y

    positionService.updatePosition(position!, GONG__StackPath).subscribe(
        position => {
            // console.log("position updated")
        }
    )
}

// onMove is called each time the shape is moved
export function onNoteMove(umlNote: joint.shapes.standard.Rectangle) {


    let note = umlNote.attributes['note'] as gongdoc.NoteShapeDB
    let noteService = umlNote.attributes['noteService'] as gongdoc.NoteShapeService
    let GONG__StackPath = umlNote.attributes['GONG__StackPath'] 

    note.X = umlNote.get('position')!.x
    note.Y = umlNote.get('position')!.y
    noteService.updateNoteShape(note!, GONG__StackPath).subscribe(
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
    let GONG__StackPath = standardLink.attributes['GONG__StackPath'] 

    middleVertice!.X = standardLink.get('vertices')![0].x
    middleVertice!.Y = standardLink.get('vertices')![0].y

    verticeService.updateVertice(middleVertice!, GONG__StackPath).subscribe(
        middleVertice => {
            // console.log("middleVertice updated")
        }
    )
}