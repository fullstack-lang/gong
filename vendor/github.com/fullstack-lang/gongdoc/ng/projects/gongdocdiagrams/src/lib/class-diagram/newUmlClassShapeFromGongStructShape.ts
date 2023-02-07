import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { shapeIdentifierToShapeName } from './shape-identifier-to-shape-name';

export function newUmlClassShapeFromGongStructShape(gongStructShape: gongdoc.GongStructShapeDB,
    positionService: gongdoc.PositionService,
    gongStructShapeService: gongdoc.GongStructShapeService): joint.shapes.uml.Class {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the classshape type
    var attributes = new Array<string>()
    if (gongStructShape.Fields) {

        // sort fields according to the index
        gongStructShape.Fields.sort((fieldA: gongdoc.FieldDB, fieldB: gongdoc.FieldDB) => {
            return fieldA.GongStructShape_FieldsDBID_Index.Int64 - fieldB.GongStructShape_FieldsDBID_Index.Int64
        })

        for (let idx = 0; idx < gongStructShape.Fields!.length; idx++) {
            let field = gongStructShape.Fields![idx]
            // console.log("Adding " + field.Fieldname + " " + field.Structname + " " + field.Fieldtypename)

            let parts = field.Identifier.split(".");
            let fieldName = parts[parts.length - 1];

            attributes.push(fieldName + " : " + field.Fieldtypename)
        }
    }

    let classShapeTitle = shapeIdentifierToShapeName(gongStructShape.Identifier)

    // show nb of instances if necessary
    if (gongStructShape.ShowNbInstances) {
        classShapeTitle += " ( " + gongStructShape.NbInstances + " )"
    }

    return new joint.shapes.uml.Class(
        {
            position: {
                x: gongStructShape.Position!.X,
                y: gongStructShape.Position!.Y
            },
            size: { width: gongStructShape.Width, height: gongStructShape.Heigth },
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
            classshape: gongStructShape,
            positionService: positionService,
            gongStructShapeService: gongStructShapeService
        }
    )

}