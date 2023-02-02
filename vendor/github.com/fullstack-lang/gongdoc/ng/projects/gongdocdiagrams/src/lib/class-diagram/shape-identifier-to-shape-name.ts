export function shapeIdentifierToShapeName(identifier: string) {
    return identifier.replace("ref_models.", "")
}