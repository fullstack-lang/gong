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

// IdentifierMetaToFieldName take an ident in the forms
// "ref_models.Foo{}.Name" and returns "Name"
func IdentifierMetaToFieldName(fieldMetaIdentifier any) (fieldName string) {

	var fieldMetaIdentifierString string
	var ok bool
	if fieldMetaIdentifierString, ok = fieldMetaIdentifier.(string); !ok {
		return ""
	}

	fielddentifier := strings.ReplaceAll(fieldMetaIdentifierString, "{}", "")

	fieldName = IdentifierToFieldName(fielddentifier)
	return
}

// IdentifierToFieldName take an ident in the forms
// ref_models.Foo{}.Name and returns "Name"
func GongEnumValueShapeIdentifierMetaToValueName(identifierMeta any) (valueName string) {

	var identifierMetaString string
	var ok bool
	if identifierMetaString, ok = identifierMeta.(string); !ok {
		return ""
	}

	valueName = strings.TrimPrefix(identifierMetaString, RefPrefixReferencedPackage+RefPackagePlusPeriod)

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

func GongEnumValueToIdentifierMeta(gongEnumValue string) (res string) {

	res = RefPrefixReferencedPackage + RefPackagePlusPeriod + gongEnumValue
	return
}

// IdentifierToGongStructName take an ident in the forms
// "ref_models.Foo" and returns "Foo"
func IdentifierToGongStructName(structIdentifier string) (structName string) {

	structName = strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage+"models.")

	// in case we uses the identifier meta
	structName = strings.TrimSuffix(structName, "{}")

	return
}

// IdentifierMetaToGongStructName take an ident in the forms
// "ref_models.Foo" and returns "Foo"
func IdentifierMetaToGongStructName(structIdentifierMeta any) (structName string) {

	var structIdentifier string
	var ok bool
	if structIdentifier, ok = structIdentifierMeta.(string); !ok {
		return ""
	}

	structName = strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage+"models.")

	// in case we uses the identifier meta
	structName = strings.TrimSuffix(structName, "{}")

	return
}

func GongStructNameToIdentifier(structName string) (identifier string) {

	identifier = RefPrefixReferencedPackage + RefPackagePlusPeriod + structName
	return
}

func GongNoteNameToIdentifier(gongNoteName string) (identifier string) {

	return GongStructNameToIdentifier(gongNoteName)
}

// turns new(ref_models.AEnumType2) into ref_models.AEnumType2
func GongEnumIdentifierMetaToGongEnumName(gongEnumIdentifierMeta any) (gongStructName string) {

	var gongEnumIdentifier string
	var ok bool
	if gongEnumIdentifier, ok = gongEnumIdentifierMeta.(string); !ok {
		return ""
	}

	gongStructName = strings.TrimPrefix(gongEnumIdentifier, "new("+RefPrefixReferencedPackage+RefPackagePlusPeriod)

	gongStructName = strings.TrimSuffix(gongStructName, ")")

	return
}
