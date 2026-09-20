package models

import (
	"log"
	"strings"
)

const RefPrefixReferencedPackage = "ref_"
const RefPackagePlusPeriod = "models."

// IdentifierToFieldName take an ident in the forms
// "ref_models.Foo.Name" or "ref_x.Foo.Name" and returns "Name"
func IdentifierToFieldName(fieldIdentifier string) (fieldName string) {

	clean := strings.TrimPrefix(fieldIdentifier, RefPrefixReferencedPackage)

	subStrings := strings.Split(clean, ".")
	if len(subStrings) < 2 {
		log.Fatalln("IdentifierToFieldName: wrong number of substrings in ", fieldIdentifier)
	}

	fieldName = subStrings[len(subStrings)-1]

	return
}

// IdentifierMetaToFieldName take an ident in the forms
// "ref_models.Foo{}.Name" or "ref_x.Foo{}.Name" and returns "Name"
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

// IdentifierMetaToStructAndFieldName take an ident in the forms
// "ref_models.Foo{}.Name" and returns "Foo", "Name"
func IdentifierMetaToStructAndFieldName(fieldMetaIdentifier any) (structname, fieldName string) {

	var fieldMetaIdentifierString string
	var ok bool
	if fieldMetaIdentifierString, ok = fieldMetaIdentifier.(string); !ok {
		return "", ""
	}

	fielddentifier := strings.ReplaceAll(fieldMetaIdentifierString, "{}", "")

	structname, fieldName = IdentifierToReceiverAndFieldName(fielddentifier)
	return
}

// GongEnumValueShapeIdentifierMetaToValueName take an ident in the forms
// ref_models.Foo{}.Name and returns "Name"
func GongEnumValueShapeIdentifierMetaToValueName(identifierMeta any) (valueName string) {

	var identifierMetaString string
	var ok bool
	if identifierMetaString, ok = identifierMeta.(string); !ok {
		return ""
	}

	clean := strings.TrimPrefix(identifierMetaString, RefPrefixReferencedPackage)
	if idx := strings.Index(clean, "."); idx != -1 {
		valueName = clean[idx+1:]
	} else {
		valueName = clean
	}

	return
}

// IdentifierToReceiverAndFieldName take an ident in the forms
// "ref_models.Foo.Name" or "ref_x.Foo.Name" and returns "Foo", "Name"
func IdentifierToReceiverAndFieldName(fieldIdentifier string) (receiver, fieldName string) {

	clean := strings.TrimPrefix(fieldIdentifier, RefPrefixReferencedPackage)

	subStrings := strings.Split(clean, ".")
	if len(subStrings) < 2 {
		log.Fatalln("IdentifierToReceiverAndFieldName: wrong number of substrings in ", fieldIdentifier)
	}

	fieldName = subStrings[len(subStrings)-1]
	receiver = subStrings[len(subStrings)-2]

	return
}

// GongstructAndFieldnameToFieldIdentifier takes "Foo" "Name" and returns "ref_models.Foo.Name"
func GongstructAndFieldnameToFieldIdentifier(structName string, fieldName string) (fieldIdentifier string) {
	return GongstructAndFieldnameToFieldIdentifierWithPackage("models", structName, fieldName)
}

func GongstructAndFieldnameToFieldIdentifierWithPackage(pkgName string, structName string, fieldName string) (fieldIdentifier string) {
	if pkgName == "" {
		pkgName = "models"
	}
	return RefPrefixReferencedPackage + pkgName + "." + structName + "." + fieldName
}

func GongEnumValueToIdentifierMeta(gongEnumValue string) (res string) {
	return GongEnumValueToIdentifierMetaWithPackage("models", gongEnumValue)
}

func GongEnumValueToIdentifierMetaWithPackage(pkgName string, gongEnumValue string) (res string) {
	if pkgName == "" {
		pkgName = "models"
	}
	return RefPrefixReferencedPackage + pkgName + "." + gongEnumValue
}

// IdentifierToGongStructName take an ident in the forms
// "ref_models.Foo", "ref_models.Foo{}", or "ref_x.Foo{}" and returns "Foo"
func IdentifierToGongStructName(structIdentifier string) (structName string) {

	clean := strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage)
	clean = strings.TrimSuffix(clean, "{}")

	subStrings := strings.Split(clean, ".")
	if len(subStrings) >= 2 {
		return subStrings[len(subStrings)-1]
	}

	return clean
}

// IdentifierMetaToGongStructName take an ident in the forms
// "ref_models.Foo", "ref_models.Foo{}", or "ref_x.Foo{}" and returns "Foo"
func IdentifierMetaToGongStructName(structIdentifierMeta any) (structName string) {

	var structIdentifier string
	var ok bool
	if structIdentifier, ok = structIdentifierMeta.(string); !ok {
		return ""
	}

	return IdentifierToGongStructName(structIdentifier)
}

// IdentifierMetaToPackageAndGongStructName returns package name and struct name
// e.g. "ref_x.Foo{}" -> ("x", "Foo")
func IdentifierMetaToPackageAndGongStructName(structIdentifierMeta any) (pkgName, structName string) {

	var structIdentifier string
	var ok bool
	if structIdentifier, ok = structIdentifierMeta.(string); !ok {
		return "", ""
	}

	clean := strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage)
	clean = strings.TrimSuffix(clean, "{}")

	subStrings := strings.Split(clean, ".")
	if len(subStrings) >= 2 {
		return subStrings[0], subStrings[len(subStrings)-1]
	}

	return "models", clean
}

func GongStructNameToIdentifier(structName string) (identifier string) {
	return GongStructNameToIdentifierWithPackage("models", structName)
}

func GongStructNameToIdentifierWithPackage(pkgName string, structName string) (identifier string) {
	if pkgName == "" {
		pkgName = "models"
	}
	return RefPrefixReferencedPackage + pkgName + "." + structName
}

func GongNoteNameToIdentifier(gongNoteName string) (identifier string) {
	return GongStructNameToIdentifier(gongNoteName)
}

func GongNoteNameToIdentifierWithPackage(pkgName string, gongNoteName string) (identifier string) {
	return GongStructNameToIdentifierWithPackage(pkgName, gongNoteName)
}

// turns new(ref_models.AEnumType2) or new(ref_x.AEnumType2) into AEnumType2
func GongEnumIdentifierMetaToGongEnumName(gongEnumIdentifierMeta any) (gongStructName string) {

	var gongEnumIdentifier string
	var ok bool
	if gongEnumIdentifier, ok = gongEnumIdentifierMeta.(string); !ok {
		return ""
	}

	clean := strings.TrimPrefix(gongEnumIdentifier, "new(")
	clean = strings.TrimSuffix(clean, ")")
	clean = strings.TrimPrefix(clean, RefPrefixReferencedPackage)

	subStrings := strings.Split(clean, ".")
	if len(subStrings) >= 2 {
		return subStrings[len(subStrings)-1]
	}

	return clean
}
