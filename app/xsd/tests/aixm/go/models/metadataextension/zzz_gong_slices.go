// generated code - do not edit
package metadataextension

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct MD_MetadataExtensionInformation_Type
	// insertion point per field
	stage.MD_MetadataExtensionInformation_Type_ExtendedElementInformation_reverseMap = make(map[*MD_ExtendedElementInformation_PropertyType]*MD_MetadataExtensionInformation_Type)
	for md_metadataextensioninformation_type := range stage.MD_MetadataExtensionInformation_Types {
		_ = md_metadataextensioninformation_type
		for _, _md_extendedelementinformation_propertytype := range md_metadataextensioninformation_type.ExtendedElementInformation {
			stage.MD_MetadataExtensionInformation_Type_ExtendedElementInformation_reverseMap[_md_extendedelementinformation_propertytype] = md_metadataextensioninformation_type
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.MD_DatatypeCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_ExtendedElementInformation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_ExtendedElementInformation_Types)

	res = __gong__appendInstances(res, stage.MD_MetadataExtensionInformation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_MetadataExtensionInformation_Types)

	res = __gong__appendInstances(res, stage.MD_ObligationCode_PropertyTypes)

	return
}

// insertion point per named struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_DatatypeCode_PropertyType)
	md_datatypecode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_ExtendedElementInformation_PropertyType)
	md_extendedelementinformation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongCopy() GongstructIF {
	newInstance := new(MD_ExtendedElementInformation_Type)
	md_extendedelementinformation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_MetadataExtensionInformation_PropertyType)
	md_metadataextensioninformation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongCopy() GongstructIF {
	newInstance := new(MD_MetadataExtensionInformation_Type)
	md_metadataextensioninformation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_ObligationCode_PropertyType)
	md_obligationcode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_datatypecode_propertytype)
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_extendedelementinformation_propertytype)
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_extendedelementinformation_type)
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_metadataextensioninformation_propertytype)
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_metadataextensioninformation_type)
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_obligationcode_propertytype)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.MD_DatatypeCode_PropertyTypes,
		stage.MD_DatatypeCode_PropertyType_stagedOrder,
		stage.MD_DatatypeCode_PropertyTypes_reference,
		&stage.MD_DatatypeCode_PropertyTypes_referenceOrder,
		stage.MD_DatatypeCode_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_ExtendedElementInformation_PropertyTypes,
		stage.MD_ExtendedElementInformation_PropertyType_stagedOrder,
		stage.MD_ExtendedElementInformation_PropertyTypes_reference,
		&stage.MD_ExtendedElementInformation_PropertyTypes_referenceOrder,
		stage.MD_ExtendedElementInformation_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_ExtendedElementInformation_Types,
		stage.MD_ExtendedElementInformation_Type_stagedOrder,
		stage.MD_ExtendedElementInformation_Types_reference,
		&stage.MD_ExtendedElementInformation_Types_referenceOrder,
		stage.MD_ExtendedElementInformation_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_MetadataExtensionInformation_PropertyTypes,
		stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder,
		stage.MD_MetadataExtensionInformation_PropertyTypes_reference,
		&stage.MD_MetadataExtensionInformation_PropertyTypes_referenceOrder,
		stage.MD_MetadataExtensionInformation_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_MetadataExtensionInformation_Types,
		stage.MD_MetadataExtensionInformation_Type_stagedOrder,
		stage.MD_MetadataExtensionInformation_Types_reference,
		&stage.MD_MetadataExtensionInformation_Types_referenceOrder,
		stage.MD_MetadataExtensionInformation_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_ObligationCode_PropertyTypes,
		stage.MD_ObligationCode_PropertyType_stagedOrder,
		stage.MD_ObligationCode_PropertyTypes_reference,
		&stage.MD_ObligationCode_PropertyTypes_referenceOrder,
		stage.MD_ObligationCode_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.MD_DatatypeCode_PropertyTypes, &stage.MD_DatatypeCode_PropertyTypes_reference, &stage.MD_DatatypeCode_PropertyTypes_referenceOrder, &stage.MD_DatatypeCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_ExtendedElementInformation_PropertyTypes, &stage.MD_ExtendedElementInformation_PropertyTypes_reference, &stage.MD_ExtendedElementInformation_PropertyTypes_referenceOrder, &stage.MD_ExtendedElementInformation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_ExtendedElementInformation_Types, &stage.MD_ExtendedElementInformation_Types_reference, &stage.MD_ExtendedElementInformation_Types_referenceOrder, &stage.MD_ExtendedElementInformation_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_MetadataExtensionInformation_PropertyTypes, &stage.MD_MetadataExtensionInformation_PropertyTypes_reference, &stage.MD_MetadataExtensionInformation_PropertyTypes_referenceOrder, &stage.MD_MetadataExtensionInformation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_MetadataExtensionInformation_Types, &stage.MD_MetadataExtensionInformation_Types_reference, &stage.MD_MetadataExtensionInformation_Types_referenceOrder, &stage.MD_MetadataExtensionInformation_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_ObligationCode_PropertyTypes, &stage.MD_ObligationCode_PropertyTypes_reference, &stage.MD_ObligationCode_PropertyTypes_referenceOrder, &stage.MD_ObligationCode_PropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_DatatypeCode_PropertyType_stagedOrder, stage.MD_DatatypeCode_PropertyTypes_referenceOrder, md_datatypecode_propertytype, "MD_DatatypeCode_PropertyType")
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ExtendedElementInformation_PropertyType_stagedOrder, stage.MD_ExtendedElementInformation_PropertyTypes_referenceOrder, md_extendedelementinformation_propertytype, "MD_ExtendedElementInformation_PropertyType")
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ExtendedElementInformation_Type_stagedOrder, stage.MD_ExtendedElementInformation_Types_referenceOrder, md_extendedelementinformation_type, "MD_ExtendedElementInformation_Type")
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder, stage.MD_MetadataExtensionInformation_PropertyTypes_referenceOrder, md_metadataextensioninformation_propertytype, "MD_MetadataExtensionInformation_PropertyType")
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_MetadataExtensionInformation_Type_stagedOrder, stage.MD_MetadataExtensionInformation_Types_referenceOrder, md_metadataextensioninformation_type, "MD_MetadataExtensionInformation_Type")
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ObligationCode_PropertyType_stagedOrder, stage.MD_ObligationCode_PropertyTypes_referenceOrder, md_obligationcode_propertytype, "MD_ObligationCode_PropertyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_datatypecode_propertytype, md_datatypecode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_datatypecode_propertytype.GongGetIdentifier(stage)
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_extendedelementinformation_propertytype, md_extendedelementinformation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_extendedelementinformation_propertytype.GongGetIdentifier(stage)
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_extendedelementinformation_type, md_extendedelementinformation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_extendedelementinformation_type.GongGetIdentifier(stage)
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_metadataextensioninformation_propertytype, md_metadataextensioninformation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_metadataextensioninformation_propertytype.GongGetIdentifier(stage)
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_metadataextensioninformation_type, md_metadataextensioninformation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_metadataextensioninformation_type.GongGetIdentifier(stage)
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_obligationcode_propertytype, md_obligationcode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_obligationcode_propertytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_datatypecode_propertytype.GongGetIdentifier(stage), "MD_DatatypeCode_PropertyType", md_datatypecode_propertytype.Name)
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_extendedelementinformation_propertytype.GongGetIdentifier(stage), "MD_ExtendedElementInformation_PropertyType", md_extendedelementinformation_propertytype.Name)
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_extendedelementinformation_type.GongGetIdentifier(stage), "MD_ExtendedElementInformation_Type", md_extendedelementinformation_type.Name)
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_metadataextensioninformation_propertytype.GongGetIdentifier(stage), "MD_MetadataExtensionInformation_PropertyType", md_metadataextensioninformation_propertytype.Name)
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_metadataextensioninformation_type.GongGetIdentifier(stage), "MD_MetadataExtensionInformation_Type", md_metadataextensioninformation_type.Name)
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_obligationcode_propertytype.GongGetIdentifier(stage), "MD_ObligationCode_PropertyType", md_obligationcode_propertytype.Name)
}

// insertion point for unstaging
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_datatypecode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_extendedelementinformation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_extendedelementinformation_type.GongGetReferenceIdentifier(stage))
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_metadataextensioninformation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_metadataextensioninformation_type.GongGetReferenceIdentifier(stage))
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_obligationcode_propertytype.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__metadataextension__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
