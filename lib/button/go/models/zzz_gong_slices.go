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
	// insertion point per named struct
	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Buttons_reverseMap = make(map[*Button]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _button := range group.Buttons {
			stage.Group_Buttons_reverseMap[_button] = group
		}
	}

	// Compute reverse map for named struct GroupToogle
	// insertion point per field
	stage.GroupToogle_ButtonToggles_reverseMap = make(map[*ButtonToggle]*GroupToogle)
	for grouptoogle := range stage.GroupToogles {
		_ = grouptoogle
		for _, _buttontoggle := range grouptoogle.ButtonToggles {
			stage.GroupToogle_ButtonToggles_reverseMap[_buttontoggle] = grouptoogle
		}
	}

	// Compute reverse map for named struct Layout
	// insertion point per field
	stage.Layout_Groups_reverseMap = make(map[*Group]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _group := range layout.Groups {
			stage.Layout_Groups_reverseMap[_group] = layout
		}
	}
	stage.Layout_GroupToogles_reverseMap = make(map[*GroupToogle]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _grouptoogle := range layout.GroupToogles {
			stage.Layout_GroupToogles_reverseMap[_grouptoogle] = layout
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Buttons)

	res = __gong__appendInstances(res, stage.ButtonToggles)

	res = __gong__appendInstances(res, stage.Groups)

	res = __gong__appendInstances(res, stage.GroupToogles)

	res = __gong__appendInstances(res, stage.Layouts)

	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := new(Button)
	button.GongCopyBasicFields(newInstance)
	return newInstance
}

func (buttontoggle *ButtonToggle) GongCopy() GongstructIF {
	newInstance := new(ButtonToggle)
	buttontoggle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (grouptoogle *GroupToogle) GongCopy() GongstructIF {
	newInstance := new(GroupToogle)
	grouptoogle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (layout *Layout) GongCopy() GongstructIF {
	newInstance := new(Layout)
	layout.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (button *Button) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, button)
}

func (buttontoggle *ButtonToggle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, buttontoggle)
}

func (group *Group) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group)
}

func (grouptoogle *GroupToogle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, grouptoogle)
}

func (layout *Layout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, layout)
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
		stage.Buttons,
		stage.Button_stagedOrder,
		stage.Buttons_reference,
		&stage.Buttons_referenceOrder,
		stage.Buttons_instance,
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
		stage.ButtonToggles,
		stage.ButtonToggle_stagedOrder,
		stage.ButtonToggles_reference,
		&stage.ButtonToggles_referenceOrder,
		stage.ButtonToggles_instance,
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
		stage.Groups,
		stage.Group_stagedOrder,
		stage.Groups_reference,
		&stage.Groups_referenceOrder,
		stage.Groups_instance,
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
		stage.GroupToogles,
		stage.GroupToogle_stagedOrder,
		stage.GroupToogles_reference,
		&stage.GroupToogles_referenceOrder,
		stage.GroupToogles_instance,
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
		stage.Layouts,
		stage.Layout_stagedOrder,
		stage.Layouts_reference,
		&stage.Layouts_referenceOrder,
		stage.Layouts_instance,
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
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Buttons, &stage.Buttons_reference, &stage.Buttons_referenceOrder, &stage.Buttons_instance)

	__gong__computeReferencePass1(stage, stage.ButtonToggles, &stage.ButtonToggles_reference, &stage.ButtonToggles_referenceOrder, &stage.ButtonToggles_instance)

	__gong__computeReferencePass1(stage, stage.Groups, &stage.Groups_reference, &stage.Groups_referenceOrder, &stage.Groups_instance)

	__gong__computeReferencePass1(stage, stage.GroupToogles, &stage.GroupToogles_reference, &stage.GroupToogles_referenceOrder, &stage.GroupToogles_instance)

	__gong__computeReferencePass1(stage, stage.Layouts, &stage.Layouts_reference, &stage.Layouts_referenceOrder, &stage.Layouts_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Buttons, stage.Buttons_reference, stage)

	__gong__computeReferencePass2(stage.ButtonToggles, stage.ButtonToggles_reference, stage)

	__gong__computeReferencePass2(stage.Groups, stage.Groups_reference, stage)

	__gong__computeReferencePass2(stage.GroupToogles, stage.GroupToogles_reference, stage)

	__gong__computeReferencePass2(stage.Layouts, stage.Layouts_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (button *Button) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Button_stagedOrder, stage.Buttons_referenceOrder, button, "Button")
}

func (buttontoggle *ButtonToggle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ButtonToggle_stagedOrder, stage.ButtonToggles_referenceOrder, buttontoggle, "ButtonToggle")
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_stagedOrder, stage.Groups_referenceOrder, group, "Group")
}

func (grouptoogle *GroupToogle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroupToogle_stagedOrder, stage.GroupToogles_referenceOrder, grouptoogle, "GroupToogle")
}

func (layout *Layout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Layout_stagedOrder, stage.Layouts_referenceOrder, layout, "Layout")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(button, button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return button.GongGetIdentifier(stage)
}

func (buttontoggle *ButtonToggle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(buttontoggle, buttontoggle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (buttontoggle *ButtonToggle) GongGetReferenceIdentifier(stage *Stage) string {
	return buttontoggle.GongGetIdentifier(stage)
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group, group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return group.GongGetIdentifier(stage)
}

func (grouptoogle *GroupToogle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(grouptoogle, grouptoogle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grouptoogle *GroupToogle) GongGetReferenceIdentifier(stage *Stage) string {
	return grouptoogle.GongGetIdentifier(stage)
}

func (layout *Layout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(layout, layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layout *Layout) GongGetReferenceIdentifier(stage *Stage) string {
	return layout.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (button *Button) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(button.GongGetIdentifier(stage), "Button", button.Name)
}

func (buttontoggle *ButtonToggle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(buttontoggle.GongGetIdentifier(stage), "ButtonToggle", buttontoggle.Name)
}

func (group *Group) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group.GongGetIdentifier(stage), "Group", group.Name)
}

func (grouptoogle *GroupToogle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(grouptoogle.GongGetIdentifier(stage), "GroupToogle", grouptoogle.Name)
}

func (layout *Layout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(layout.GongGetIdentifier(stage), "Layout", layout.Name)
}

// insertion point for unstaging
func (button *Button) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(button.GongGetReferenceIdentifier(stage))
}

func (buttontoggle *ButtonToggle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(buttontoggle.GongGetReferenceIdentifier(stage))
}

func (group *Group) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group.GongGetReferenceIdentifier(stage))
}

func (grouptoogle *GroupToogle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(grouptoogle.GongGetReferenceIdentifier(stage))
}

func (layout *Layout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(layout.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
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
