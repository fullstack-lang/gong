// generated code - do not edit
package measures

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
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AngleTypes)

	res = __gong__appendInstances(res, stage.AreaTypes)

	res = __gong__appendInstances(res, stage.GridLengthTypes)

	res = __gong__appendInstances(res, stage.LengthTypes)

	res = __gong__appendInstances(res, stage.ScaleTypes)

	res = __gong__appendInstances(res, stage.SpeedTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.VolumeTypes)

	return
}

// insertion point per named struct
func (angletype *AngleType) GongCopy() GongstructIF {
	newInstance := new(AngleType)
	angletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (areatype *AreaType) GongCopy() GongstructIF {
	newInstance := new(AreaType)
	areatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gridlengthtype *GridLengthType) GongCopy() GongstructIF {
	newInstance := new(GridLengthType)
	gridlengthtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lengthtype *LengthType) GongCopy() GongstructIF {
	newInstance := new(LengthType)
	lengthtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scaletype *ScaleType) GongCopy() GongstructIF {
	newInstance := new(ScaleType)
	scaletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (speedtype *SpeedType) GongCopy() GongstructIF {
	newInstance := new(SpeedType)
	speedtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timetype *TimeType) GongCopy() GongstructIF {
	newInstance := new(TimeType)
	timetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (volumetype *VolumeType) GongCopy() GongstructIF {
	newInstance := new(VolumeType)
	volumetype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (angletype *AngleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angletype)
}

func (areatype *AreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, areatype)
}

func (gridlengthtype *GridLengthType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gridlengthtype)
}

func (lengthtype *LengthType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lengthtype)
}

func (scaletype *ScaleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scaletype)
}

func (speedtype *SpeedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, speedtype)
}

func (timetype *TimeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timetype)
}

func (volumetype *VolumeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, volumetype)
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
		stage.AngleTypes,
		stage.AngleType_stagedOrder,
		stage.AngleTypes_reference,
		&stage.AngleTypes_referenceOrder,
		stage.AngleTypes_instance,
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
		stage.AreaTypes,
		stage.AreaType_stagedOrder,
		stage.AreaTypes_reference,
		&stage.AreaTypes_referenceOrder,
		stage.AreaTypes_instance,
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
		stage.GridLengthTypes,
		stage.GridLengthType_stagedOrder,
		stage.GridLengthTypes_reference,
		&stage.GridLengthTypes_referenceOrder,
		stage.GridLengthTypes_instance,
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
		stage.LengthTypes,
		stage.LengthType_stagedOrder,
		stage.LengthTypes_reference,
		&stage.LengthTypes_referenceOrder,
		stage.LengthTypes_instance,
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
		stage.ScaleTypes,
		stage.ScaleType_stagedOrder,
		stage.ScaleTypes_reference,
		&stage.ScaleTypes_referenceOrder,
		stage.ScaleTypes_instance,
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
		stage.SpeedTypes,
		stage.SpeedType_stagedOrder,
		stage.SpeedTypes_reference,
		&stage.SpeedTypes_referenceOrder,
		stage.SpeedTypes_instance,
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
		stage.TimeTypes,
		stage.TimeType_stagedOrder,
		stage.TimeTypes_reference,
		&stage.TimeTypes_referenceOrder,
		stage.TimeTypes_instance,
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
		stage.VolumeTypes,
		stage.VolumeType_stagedOrder,
		stage.VolumeTypes_reference,
		&stage.VolumeTypes_referenceOrder,
		stage.VolumeTypes_instance,
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
	__gong__computeReferencePass1(stage, stage.AngleTypes, &stage.AngleTypes_reference, &stage.AngleTypes_referenceOrder, &stage.AngleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AreaTypes, &stage.AreaTypes_reference, &stage.AreaTypes_referenceOrder, &stage.AreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.GridLengthTypes, &stage.GridLengthTypes_reference, &stage.GridLengthTypes_referenceOrder, &stage.GridLengthTypes_instance)

	__gong__computeReferencePass1(stage, stage.LengthTypes, &stage.LengthTypes_reference, &stage.LengthTypes_referenceOrder, &stage.LengthTypes_instance)

	__gong__computeReferencePass1(stage, stage.ScaleTypes, &stage.ScaleTypes_reference, &stage.ScaleTypes_referenceOrder, &stage.ScaleTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpeedTypes, &stage.SpeedTypes_reference, &stage.SpeedTypes_referenceOrder, &stage.SpeedTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.VolumeTypes, &stage.VolumeTypes_reference, &stage.VolumeTypes_referenceOrder, &stage.VolumeTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AngleTypes, stage.AngleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AreaTypes, stage.AreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.GridLengthTypes, stage.GridLengthTypes_reference, stage)

	__gong__computeReferencePass2(stage.LengthTypes, stage.LengthTypes_reference, stage)

	__gong__computeReferencePass2(stage.ScaleTypes, stage.ScaleTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpeedTypes, stage.SpeedTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.VolumeTypes, stage.VolumeTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (angletype *AngleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleType_stagedOrder, stage.AngleTypes_referenceOrder, angletype, "AngleType")
}

func (areatype *AreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AreaType_stagedOrder, stage.AreaTypes_referenceOrder, areatype, "AreaType")
}

func (gridlengthtype *GridLengthType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GridLengthType_stagedOrder, stage.GridLengthTypes_referenceOrder, gridlengthtype, "GridLengthType")
}

func (lengthtype *LengthType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LengthType_stagedOrder, stage.LengthTypes_referenceOrder, lengthtype, "LengthType")
}

func (scaletype *ScaleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScaleType_stagedOrder, stage.ScaleTypes_referenceOrder, scaletype, "ScaleType")
}

func (speedtype *SpeedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpeedType_stagedOrder, stage.SpeedTypes_referenceOrder, speedtype, "SpeedType")
}

func (timetype *TimeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeType_stagedOrder, stage.TimeTypes_referenceOrder, timetype, "TimeType")
}

func (volumetype *VolumeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VolumeType_stagedOrder, stage.VolumeTypes_referenceOrder, volumetype, "VolumeType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (angletype *AngleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angletype, angletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angletype *AngleType) GongGetReferenceIdentifier(stage *Stage) string {
	return angletype.GongGetIdentifier(stage)
}

func (areatype *AreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(areatype, areatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (areatype *AreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return areatype.GongGetIdentifier(stage)
}

func (gridlengthtype *GridLengthType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gridlengthtype, gridlengthtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gridlengthtype *GridLengthType) GongGetReferenceIdentifier(stage *Stage) string {
	return gridlengthtype.GongGetIdentifier(stage)
}

func (lengthtype *LengthType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lengthtype, lengthtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lengthtype *LengthType) GongGetReferenceIdentifier(stage *Stage) string {
	return lengthtype.GongGetIdentifier(stage)
}

func (scaletype *ScaleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scaletype, scaletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scaletype *ScaleType) GongGetReferenceIdentifier(stage *Stage) string {
	return scaletype.GongGetIdentifier(stage)
}

func (speedtype *SpeedType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(speedtype, speedtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (speedtype *SpeedType) GongGetReferenceIdentifier(stage *Stage) string {
	return speedtype.GongGetIdentifier(stage)
}

func (timetype *TimeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timetype, timetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timetype *TimeType) GongGetReferenceIdentifier(stage *Stage) string {
	return timetype.GongGetIdentifier(stage)
}

func (volumetype *VolumeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(volumetype, volumetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (volumetype *VolumeType) GongGetReferenceIdentifier(stage *Stage) string {
	return volumetype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (angletype *AngleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angletype.GongGetIdentifier(stage), "AngleType", angletype.Name)
}

func (areatype *AreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(areatype.GongGetIdentifier(stage), "AreaType", areatype.Name)
}

func (gridlengthtype *GridLengthType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gridlengthtype.GongGetIdentifier(stage), "GridLengthType", gridlengthtype.Name)
}

func (lengthtype *LengthType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lengthtype.GongGetIdentifier(stage), "LengthType", lengthtype.Name)
}

func (scaletype *ScaleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scaletype.GongGetIdentifier(stage), "ScaleType", scaletype.Name)
}

func (speedtype *SpeedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(speedtype.GongGetIdentifier(stage), "SpeedType", speedtype.Name)
}

func (timetype *TimeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timetype.GongGetIdentifier(stage), "TimeType", timetype.Name)
}

func (volumetype *VolumeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(volumetype.GongGetIdentifier(stage), "VolumeType", volumetype.Name)
}

// insertion point for unstaging
func (angletype *AngleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angletype.GongGetReferenceIdentifier(stage))
}

func (areatype *AreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(areatype.GongGetReferenceIdentifier(stage))
}

func (gridlengthtype *GridLengthType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gridlengthtype.GongGetReferenceIdentifier(stage))
}

func (lengthtype *LengthType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lengthtype.GongGetReferenceIdentifier(stage))
}

func (scaletype *ScaleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scaletype.GongGetReferenceIdentifier(stage))
}

func (speedtype *SpeedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(speedtype.GongGetReferenceIdentifier(stage))
}

func (timetype *TimeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timetype.GongGetReferenceIdentifier(stage))
}

func (volumetype *VolumeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(volumetype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__measures__%s__%08d_", s.GongGetGongstructName(), order)
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
