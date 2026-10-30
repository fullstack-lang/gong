// generated code - do not edit
package models

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
	if stage.BasictypesStage != nil {
		stage.BasictypesStage.ComputeReverseMaps()
	}
	if stage.XmlxsdStage != nil {
		stage.XmlxsdStage.ComputeReverseMaps()
	}
	if stage.XlinkStage != nil {
		stage.XlinkStage.ComputeReverseMaps()
	}
	if stage.GcobaseStage != nil {
		stage.GcobaseStage.ComputeReverseMaps()
	}
	if stage.MeasuresStage != nil {
		stage.MeasuresStage.ComputeReverseMaps()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.ComputeReverseMaps()
	}
	if stage.ApplicationschemaStage != nil {
		stage.ApplicationschemaStage.ComputeReverseMaps()
	}
	if stage.ContentStage != nil {
		stage.ContentStage.ComputeReverseMaps()
	}
	if stage.MetadataextensionStage != nil {
		stage.MetadataextensionStage.ComputeReverseMaps()
	}
	if stage.PortrayalcatalogueStage != nil {
		stage.PortrayalcatalogueStage.ComputeReverseMaps()
	}
	if stage.SpatialrepresentationStage != nil {
		stage.SpatialrepresentationStage.ComputeReverseMaps()
	}
	if stage.GmdStage != nil {
		stage.GmdStage.ComputeReverseMaps()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.ComputeReverseMaps()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.ComputeReverseMaps()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.ComputeReverseMaps()
	}
	// insertion point per named struct
	// Compute reverse map for named struct AIXMBasicMessage
	// insertion point per field
	stage.AIXMBasicMessage_HasMember_reverseMap = make(map[*BasicMessageMemberAIXMPropertyType]*AIXMBasicMessage)
	for aixmbasicmessage := range stage.AIXMBasicMessages {
		_ = aixmbasicmessage
		for _, _basicmessagememberaixmpropertytype := range aixmbasicmessage.HasMember {
			stage.AIXMBasicMessage_HasMember_reverseMap[_basicmessagememberaixmpropertytype] = aixmbasicmessage
		}
	}

	// Compute reverse map for named struct AIXMBasicMessageType
	// insertion point per field
	stage.AIXMBasicMessageType_HasMember_reverseMap = make(map[*BasicMessageMemberAIXMPropertyType]*AIXMBasicMessageType)
	for aixmbasicmessagetype := range stage.AIXMBasicMessageTypes {
		_ = aixmbasicmessagetype
		for _, _basicmessagememberaixmpropertytype := range aixmbasicmessagetype.HasMember {
			stage.AIXMBasicMessageType_HasMember_reverseMap[_basicmessagememberaixmpropertytype] = aixmbasicmessagetype
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AIXMBasicMessages)

	res = __gong__appendInstances(res, stage.AIXMBasicMessageTypes)

	res = __gong__appendInstances(res, stage.BasicMessageMemberAIXMPropertyTypes)

	return
}

// insertion point per named struct
func (aixmbasicmessage *AIXMBasicMessage) GongCopy() GongstructIF {
	newInstance := new(AIXMBasicMessage)
	aixmbasicmessage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongCopy() GongstructIF {
	newInstance := new(AIXMBasicMessageType)
	aixmbasicmessagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongCopy() GongstructIF {
	newInstance := new(BasicMessageMemberAIXMPropertyType)
	basicmessagememberaixmpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (aixmbasicmessage *AIXMBasicMessage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aixmbasicmessage)
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aixmbasicmessagetype)
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, basicmessagememberaixmpropertytype)
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
		stage.AIXMBasicMessages,
		stage.AIXMBasicMessage_stagedOrder,
		stage.AIXMBasicMessages_reference,
		&stage.AIXMBasicMessages_referenceOrder,
		stage.AIXMBasicMessages_instance,
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
		stage.AIXMBasicMessageTypes,
		stage.AIXMBasicMessageType_stagedOrder,
		stage.AIXMBasicMessageTypes_reference,
		&stage.AIXMBasicMessageTypes_referenceOrder,
		stage.AIXMBasicMessageTypes_instance,
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
		stage.BasicMessageMemberAIXMPropertyTypes,
		stage.BasicMessageMemberAIXMPropertyType_stagedOrder,
		stage.BasicMessageMemberAIXMPropertyTypes_reference,
		&stage.BasicMessageMemberAIXMPropertyTypes_referenceOrder,
		stage.BasicMessageMemberAIXMPropertyTypes_instance,
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
	if stage.BasictypesStage != nil {
		stage.BasictypesStage.ComputeReferenceAndOrders()
	}
	if stage.XmlxsdStage != nil {
		stage.XmlxsdStage.ComputeReferenceAndOrders()
	}
	if stage.XlinkStage != nil {
		stage.XlinkStage.ComputeReferenceAndOrders()
	}
	if stage.GcobaseStage != nil {
		stage.GcobaseStage.ComputeReferenceAndOrders()
	}
	if stage.MeasuresStage != nil {
		stage.MeasuresStage.ComputeReferenceAndOrders()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.ComputeReferenceAndOrders()
	}
	if stage.ApplicationschemaStage != nil {
		stage.ApplicationschemaStage.ComputeReferenceAndOrders()
	}
	if stage.ContentStage != nil {
		stage.ContentStage.ComputeReferenceAndOrders()
	}
	if stage.MetadataextensionStage != nil {
		stage.MetadataextensionStage.ComputeReferenceAndOrders()
	}
	if stage.PortrayalcatalogueStage != nil {
		stage.PortrayalcatalogueStage.ComputeReferenceAndOrders()
	}
	if stage.SpatialrepresentationStage != nil {
		stage.SpatialrepresentationStage.ComputeReferenceAndOrders()
	}
	if stage.GmdStage != nil {
		stage.GmdStage.ComputeReferenceAndOrders()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.ComputeReferenceAndOrders()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.ComputeReferenceAndOrders()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.ComputeReferenceAndOrders()
	}
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.AIXMBasicMessages, &stage.AIXMBasicMessages_reference, &stage.AIXMBasicMessages_referenceOrder, &stage.AIXMBasicMessages_instance)

	__gong__computeReferencePass1(stage, stage.AIXMBasicMessageTypes, &stage.AIXMBasicMessageTypes_reference, &stage.AIXMBasicMessageTypes_referenceOrder, &stage.AIXMBasicMessageTypes_instance)

	__gong__computeReferencePass1(stage, stage.BasicMessageMemberAIXMPropertyTypes, &stage.BasicMessageMemberAIXMPropertyTypes_reference, &stage.BasicMessageMemberAIXMPropertyTypes_referenceOrder, &stage.BasicMessageMemberAIXMPropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AIXMBasicMessages, stage.AIXMBasicMessages_reference, stage)

	__gong__computeReferencePass2(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageTypes_reference, stage)

	__gong__computeReferencePass2(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (aixmbasicmessage *AIXMBasicMessage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AIXMBasicMessage_stagedOrder, stage.AIXMBasicMessages_referenceOrder, aixmbasicmessage, "AIXMBasicMessage")
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AIXMBasicMessageType_stagedOrder, stage.AIXMBasicMessageTypes_referenceOrder, aixmbasicmessagetype, "AIXMBasicMessageType")
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BasicMessageMemberAIXMPropertyType_stagedOrder, stage.BasicMessageMemberAIXMPropertyTypes_referenceOrder, basicmessagememberaixmpropertytype, "BasicMessageMemberAIXMPropertyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (aixmbasicmessage *AIXMBasicMessage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aixmbasicmessage, aixmbasicmessage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aixmbasicmessage *AIXMBasicMessage) GongGetReferenceIdentifier(stage *Stage) string {
	return aixmbasicmessage.GongGetIdentifier(stage)
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aixmbasicmessagetype, aixmbasicmessagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetReferenceIdentifier(stage *Stage) string {
	return aixmbasicmessagetype.GongGetIdentifier(stage)
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(basicmessagememberaixmpropertytype, basicmessagememberaixmpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return basicmessagememberaixmpropertytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (aixmbasicmessage *AIXMBasicMessage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aixmbasicmessage.GongGetIdentifier(stage), "AIXMBasicMessage", aixmbasicmessage.Name)
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aixmbasicmessagetype.GongGetIdentifier(stage), "AIXMBasicMessageType", aixmbasicmessagetype.Name)
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(basicmessagememberaixmpropertytype.GongGetIdentifier(stage), "BasicMessageMemberAIXMPropertyType", basicmessagememberaixmpropertytype.Name)
}

// insertion point for unstaging
func (aixmbasicmessage *AIXMBasicMessage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aixmbasicmessage.GongGetReferenceIdentifier(stage))
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aixmbasicmessagetype.GongGetReferenceIdentifier(stage))
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(basicmessagememberaixmpropertytype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__models__%s__%08d_", s.GongGetGongstructName(), order)
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
