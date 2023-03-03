
const RefPrefixReferencedPackage: string = "ref_"
const RefPackagePlusPeriod: string = "models."

export function IdentifierToReceiverAndFieldName(fieldIdentifier: string): { receiver: string, fieldName: string } {

    let prefix = RefPrefixReferencedPackage + RefPackagePlusPeriod

    let structNameWithFieldName = fieldIdentifier.replace(prefix, "")

    let subStrings = structNameWithFieldName.split(".")

    let fieldName = subStrings[1]
    let receiver = subStrings[0]

    return { receiver: receiver, fieldName: fieldName }
}

export function IdentifierToStructname(fieldIdentifier: string): string {

    let prefix = RefPrefixReferencedPackage + RefPackagePlusPeriod

    let structName = fieldIdentifier.replace(prefix, "")

    return structName
}