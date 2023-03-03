package models

import (
	"log"
	"strings"
)

const RefPrefixReferencedPackage = "ref_"
const RefPackagePlusPeriod = "models."

// IdentifierToFieldName take an ident in the forms
// "ref_models.Foo.Name" and returns "Name"
func IdentifierToFieldName(fieldIdentifier string) (fieldName string) {

	if !strings.Contains(fieldIdentifier, RefPackagePlusPeriod) {
		log.Fatalln("IdentifierToFieldName: missing", RefPackagePlusPeriod, "in", fieldIdentifier)
	}

	structNameWithFieldName := strings.TrimPrefix(fieldIdentifier, RefPrefixReferencedPackage+RefPackagePlusPeriod)

	subStrings := strings.Split(structNameWithFieldName, ".")
	if len(subStrings) != 2 {
		log.Fatalln("IdentifierToFieldName: wrong number of substrings in ", structNameWithFieldName)
	}

	fieldName = subStrings[1]

	return
}

func IdentifierToReceiverAndFieldName(fieldIdentifier string) (receiver, fieldName string) {

	if !strings.Contains(fieldIdentifier, RefPackagePlusPeriod) {
		log.Fatalln("IdentifierToReceiverAndFieldName: missing", RefPackagePlusPeriod, "in", fieldIdentifier)
	}

	structNameWithFieldName := strings.TrimPrefix(fieldIdentifier, RefPrefixReferencedPackage+RefPackagePlusPeriod)

	subStrings := strings.Split(structNameWithFieldName, ".")
	if len(subStrings) != 2 {
		log.Fatalln("IdentifierToReceiverAndFieldName: wrong number of substrings in ", structNameWithFieldName)
	}

	fieldName = subStrings[1]
	receiver = subStrings[0]

	return
}

// GongstructAndFieldnameToFieldIdentifier takes "Foo" "Name" and returns "ref_models.Foo.Name"
func GongstructAndFieldnameToFieldIdentifier(structName string, fieldName string) (fieldIdentifier string) {

	fieldIdentifier = RefPrefixReferencedPackage + RefPackagePlusPeriod +
		structName + "." + fieldName

	return
}

// IdentifierToGongObjectName take an ident in the forms
// "ref_models.Foo" and returns "Foo"
func IdentifierToGongObjectName(structIdentifier string) (structName string) {

	structName = strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage+"models.")

	return
}

func GongStructNameToIdentifier(structName string) (identifier string) {

	identifier = RefPrefixReferencedPackage + "models." + structName
	return
}
