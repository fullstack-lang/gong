import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { shapeIdentifierToShapeName } from './shape-identifier-to-shape-name';

export function newUmlClassShapeFromGongEnumShape(
    gongEnumShape: gongdoc.GongEnumShapeDB,
    positionService: gongdoc.PositionService,
    gongEnumShapeService: gongdoc.GongEnumShapeService,
    GONG__StackPath: string): joint.shapes.uml.Class {

    var attributes = new Array<string>()
    if (gongEnumShape.GongEnumValueEntrys) {

        // sort fields according to the index
        gongEnumShape.GongEnumValueEntrys.sort(
            (gongEnumValueEntryA: gongdoc.GongEnumValueEntryDB, gongEnumValueEntryB: gongdoc.GongEnumValueEntryDB) => {
                return gongEnumValueEntryA.GongEnumShape_GongEnumValueEntrysDBID_Index.Int64 -
                    gongEnumValueEntryB.GongEnumShape_GongEnumValueEntrysDBID_Index.Int64
            })

        for (let idx = 0; idx < gongEnumShape.GongEnumValueEntrys!.length; idx++) {
            let field = gongEnumShape.GongEnumValueEntrys![idx]
            // console.log("Adding " + field.Fieldname + " " + field.Structname + " " + field.Fieldtypename)

            let parts = field.Identifier.split(".");
            let fieldName = parts[parts.length - 1];

            attributes.push(fieldName)
        }
    }

    let classShapeTitle = shapeIdentifierToShapeName(gongEnumShape.Identifier)

    return new joint.shapes.uml.Class(
        {
            position: {
                x: gongEnumShape.Position!.X,
                y: gongEnumShape.Position!.Y
            },
            size: { width: gongEnumShape.Width, height: gongEnumShape.Heigth },
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
            classshape: gongEnumShape,
            positionService: positionService,
            gongEnumShapeService: gongEnumShapeService,
            GONG__StackPath: GONG__StackPath
        }
    )

}