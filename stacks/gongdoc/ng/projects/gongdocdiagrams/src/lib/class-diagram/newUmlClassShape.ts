import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'

export function newUmlClassShape(classshape: gongdoc.ClassshapeDB): joint.shapes.uml.Class {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the classshape type
    var attributes = new Array<string>()
    classshape.Fields?.forEach(
        field => {
            console.log(field.Fieldname + " " + field.Structname + " " + field.Fieldtypename)
            attributes.push(field.Fieldname + " : " + field.Fieldtypename)
        }
    )

    return new joint.shapes.uml.Class(
        {
            position: {
                x: classshape.Position.X,
                y: classshape.Position.Y
            },
            size: { width: classshape.Width, height: classshape.Heigth },
            name: [classshape.Structname],
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
            }
        }
    )

}