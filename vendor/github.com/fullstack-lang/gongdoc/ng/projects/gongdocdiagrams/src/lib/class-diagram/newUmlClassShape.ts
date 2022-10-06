import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { GongdocCommandService } from 'gongdoc';

export function newUmlClassShape(classshape: gongdoc.ClassshapeDB,
    positionService: gongdoc.PositionService,
    gongdocCommandSingloton: gongdoc.GongdocCommandDB,
    gongdocCommandService: gongdoc.GongdocCommandService): joint.shapes.uml.Class {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the classshape type
    var attributes = new Array<string>()
    if (classshape.Fields) {

        // sort fields according to the index
        classshape.Fields.sort((fieldA: gongdoc.FieldDB, fieldB: gongdoc.FieldDB) => {
            return fieldA.Classshape_FieldsDBID_Index.Int64 - fieldB.Classshape_FieldsDBID_Index.Int64
        })

        for (let idx = 0; idx < classshape.Fields!.length; idx++) {
            let field = classshape.Fields![idx]
            // console.log("Adding " + field.Fieldname + " " + field.Structname + " " + field.Fieldtypename)
            attributes.push(field.Fieldname + " : " + field.Fieldtypename)
        }
    }

    let classShapeTitle = classshape.Structname

    // show nb of instances if necessary
    if (classshape.ShowNbInstances) {
        classShapeTitle += " ( " + classshape.NbInstances + " )"
    }

    return new joint.shapes.uml.Class(
        {
            position: {
                x: classshape.Position!.X,
                y: classshape.Position!.Y
            },
            size: { width: classshape.Width, height: classshape.Heigth },
            name: [classShapeTitle],
            attributes: attributes,
            methods: [],
            attrs: {
                '.uml-class-name-rect': {
                    fill: '#ff8450',
                    stroke: '#fff',
                    'stroke-width': 0.5,
                },
                '.uml-class-name-text': {
                    'font-family': 'Roboto'
                },
                '.uml-class-attrs-rect': {
                    fill: '#fe976a',
                    stroke: '#fff',
                    height: 10,
                    'stroke-width': 0.5,
                    'font-family': 'Roboto'
                },
                '.uml-class-methods-rect': {
                    fill: '#fe976a',
                    stroke: '#fff',
                    height: 0,
                    'stroke-width': 0
                },
                '.uml-class-attrs-text': {
                    'ref-y': 0,
                    'y-alignment': 'top',
                    'font-family': 'Roboto'
                }
            },

            // store relevant attributes for working when callback are invoked
            classshape: classshape,
            positionService: positionService,
            gongdocCommandSingloton: gongdocCommandSingloton,
            gongdocCommandService: gongdocCommandService
        }
    )

}