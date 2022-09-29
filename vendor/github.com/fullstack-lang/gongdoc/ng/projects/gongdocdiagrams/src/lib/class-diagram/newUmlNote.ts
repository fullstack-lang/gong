import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { GongdocCommandService } from 'gongdoc';
import { MatLabel } from '@angular/material/form-field';

export function newUmlNote(note: gongdoc.NoteDB,
    noteService: gongdoc.NoteService,
    gongdocCommandSingloton: gongdoc.GongdocCommandDB,
    gongdocCommandService: gongdoc.GongdocCommandService): joint.shapes.basic.Rect {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the note type
    var attributes = new Array<string>()


    let noteBody = note.Body

    var rect = new joint.shapes.standard.Rectangle(
        {
            position: {
                x: note.X,
                y: note.Y
            },
            size: { width: note.Width, height: note.Heigth },
            name: [noteBody],
            methods: [],
            // store relevant attributes for working when callback are invoked
            note: note,
            noteService: noteService,
            gongdocCommandSingloton: gongdocCommandSingloton,
            gongdocCommandService: gongdocCommandService
        }
    )
    let width = noteBody.length * 12
    let lines = noteBody.split(/\r\n|\r|\n/)
    let maxLength = 0
    for (let lineNb = 0; lineNb < lines.length; lineNb++) {
        if (lines[lineNb].length > maxLength)  {
            maxLength = lines[lineNb].length
        }
    }
    // console.log("maxLength ", maxLength)

    rect.resize(
        maxLength*7, noteBody.split(/\r\n|\r|\n/).length * 17
    )
    rect.attr({
        body: {
            // rx: 10, // add a corner radius
            // ry: 10,
            fill: '#fe976a',
            "stroke-width": 0.5,
        },
        text: {
            text: noteBody
        }
        
    })

    return rect

}