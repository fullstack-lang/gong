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
	// Compute reverse map for named struct ArtefactType
	// insertion point per field

	// Compute reverse map for named struct ArtefactTypeShape
	// insertion point per field

	// Compute reverse map for named struct Artist
	// insertion point per field

	// Compute reverse map for named struct ArtistShape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

	// Compute reverse map for named struct Desk
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_MovementShapes_reverseMap = make(map[*MovementShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _movementshape := range diagram.MovementShapes {
			stage.Diagram_MovementShapes_reverseMap[_movementshape] = diagram
		}
	}
	stage.Diagram_ArtefactTypeShapes_reverseMap = make(map[*ArtefactTypeShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			stage.Diagram_ArtefactTypeShapes_reverseMap[_artefacttypeshape] = diagram
		}
	}
	stage.Diagram_ArtistShapes_reverseMap = make(map[*ArtistShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artistshape := range diagram.ArtistShapes {
			stage.Diagram_ArtistShapes_reverseMap[_artistshape] = diagram
		}
	}
	stage.Diagram_InfluenceShapes_reverseMap = make(map[*InfluenceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _influenceshape := range diagram.InfluenceShapes {
			stage.Diagram_InfluenceShapes_reverseMap[_influenceshape] = diagram
		}
	}

	// Compute reverse map for named struct Influence
	// insertion point per field

	// Compute reverse map for named struct InfluenceShape
	// insertion point per field
	stage.InfluenceShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*InfluenceShape)
	for influenceshape := range stage.InfluenceShapes {
		_ = influenceshape
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			stage.InfluenceShape_ControlPointShapes_reverseMap[_controlpointshape] = influenceshape
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}

	// Compute reverse map for named struct Movement
	// insertion point per field
	stage.Movement_Places_reverseMap = make(map[*Place]*Movement)
	for movement := range stage.Movements {
		_ = movement
		for _, _place := range movement.Places {
			stage.Movement_Places_reverseMap[_place] = movement
		}
	}

	// Compute reverse map for named struct MovementShape
	// insertion point per field

	// Compute reverse map for named struct Place
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ArtefactTypes {
		res = append(res, instance)
	}

	for instance := range stage.ArtefactTypeShapes {
		res = append(res, instance)
	}

	for instance := range stage.Artists {
		res = append(res, instance)
	}

	for instance := range stage.ArtistShapes {
		res = append(res, instance)
	}

	for instance := range stage.ControlPointShapes {
		res = append(res, instance)
	}

	for instance := range stage.Desks {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Influences {
		res = append(res, instance)
	}

	for instance := range stage.InfluenceShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Movements {
		res = append(res, instance)
	}

	for instance := range stage.MovementShapes {
		res = append(res, instance)
	}

	for instance := range stage.Places {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (artefacttype *ArtefactType) GongCopy() GongstructIF {
	newInstance := new(ArtefactType)
	artefacttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (artefacttypeshape *ArtefactTypeShape) GongCopy() GongstructIF {
	newInstance := new(ArtefactTypeShape)
	artefacttypeshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (artist *Artist) GongCopy() GongstructIF {
	newInstance := new(Artist)
	artist.GongCopyBasicFields(newInstance)
	return newInstance
}

func (artistshape *ArtistShape) GongCopy() GongstructIF {
	newInstance := new(ArtistShape)
	artistshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := new(ControlPointShape)
	controlpointshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (desk *Desk) GongCopy() GongstructIF {
	newInstance := new(Desk)
	desk.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (influence *Influence) GongCopy() GongstructIF {
	newInstance := new(Influence)
	influence.GongCopyBasicFields(newInstance)
	return newInstance
}

func (influenceshape *InfluenceShape) GongCopy() GongstructIF {
	newInstance := new(InfluenceShape)
	influenceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (movement *Movement) GongCopy() GongstructIF {
	newInstance := new(Movement)
	movement.GongCopyBasicFields(newInstance)
	return newInstance
}

func (movementshape *MovementShape) GongCopy() GongstructIF {
	newInstance := new(MovementShape)
	movementshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (place *Place) GongCopy() GongstructIF {
	newInstance := new(Place)
	place.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (artefacttype *ArtefactType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(artefacttype).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(artefacttype), uint64(stage.GetOrder(artefacttype)))
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(artefacttypeshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(artefacttypeshape), uint64(stage.GetOrder(artefacttypeshape)))
	return
}

func (artist *Artist) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(artist).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(artist), uint64(stage.GetOrder(artist)))
	return
}

func (artistshape *ArtistShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(artistshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(artistshape), uint64(stage.GetOrder(artistshape)))
	return
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlpointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlpointshape), uint64(stage.GetOrder(controlpointshape)))
	return
}

func (desk *Desk) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(desk).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(desk), uint64(stage.GetOrder(desk)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagram), uint64(stage.GetOrder(diagram)))
	return
}

func (influence *Influence) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(influence).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(influence), uint64(stage.GetOrder(influence)))
	return
}

func (influenceshape *InfluenceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(influenceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(influenceshape), uint64(stage.GetOrder(influenceshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (movement *Movement) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(movement).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(movement), uint64(stage.GetOrder(movement)))
	return
}

func (movementshape *MovementShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(movementshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(movementshape), uint64(stage.GetOrder(movementshape)))
	return
}

func (place *Place) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(place).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(place), uint64(stage.GetOrder(place)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
		stage.ArtefactTypes,
		stage.ArtefactType_stagedOrder,
		stage.ArtefactTypes_reference,
		&stage.ArtefactTypes_referenceOrder,
		stage.ArtefactTypes_instance,
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
		stage.ArtefactTypeShapes,
		stage.ArtefactTypeShape_stagedOrder,
		stage.ArtefactTypeShapes_reference,
		&stage.ArtefactTypeShapes_referenceOrder,
		stage.ArtefactTypeShapes_instance,
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
		stage.Artists,
		stage.Artist_stagedOrder,
		stage.Artists_reference,
		&stage.Artists_referenceOrder,
		stage.Artists_instance,
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
		stage.ArtistShapes,
		stage.ArtistShape_stagedOrder,
		stage.ArtistShapes_reference,
		&stage.ArtistShapes_referenceOrder,
		stage.ArtistShapes_instance,
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
		stage.ControlPointShapes,
		stage.ControlPointShape_stagedOrder,
		stage.ControlPointShapes_reference,
		&stage.ControlPointShapes_referenceOrder,
		stage.ControlPointShapes_instance,
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
		stage.Desks,
		stage.Desk_stagedOrder,
		stage.Desks_reference,
		&stage.Desks_referenceOrder,
		stage.Desks_instance,
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
		stage.Diagrams,
		stage.Diagram_stagedOrder,
		stage.Diagrams_reference,
		&stage.Diagrams_referenceOrder,
		stage.Diagrams_instance,
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
		stage.Influences,
		stage.Influence_stagedOrder,
		stage.Influences_reference,
		&stage.Influences_referenceOrder,
		stage.Influences_instance,
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
		stage.InfluenceShapes,
		stage.InfluenceShape_stagedOrder,
		stage.InfluenceShapes_reference,
		&stage.InfluenceShapes_referenceOrder,
		stage.InfluenceShapes_instance,
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
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
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
		stage.Movements,
		stage.Movement_stagedOrder,
		stage.Movements_reference,
		&stage.Movements_referenceOrder,
		stage.Movements_instance,
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
		stage.MovementShapes,
		stage.MovementShape_stagedOrder,
		stage.MovementShapes_reference,
		&stage.MovementShapes_referenceOrder,
		stage.MovementShapes_instance,
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
		stage.Places,
		stage.Place_stagedOrder,
		stage.Places_reference,
		&stage.Places_referenceOrder,
		stage.Places_instance,
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
	stage.ArtefactTypes_reference = make(map[*ArtefactType]*ArtefactType)
	stage.ArtefactTypes_referenceOrder = make(map[*ArtefactType]uint) // diff Unstage needs the reference order
	stage.ArtefactTypes_instance = make(map[*ArtefactType]*ArtefactType)
	for instance := range stage.ArtefactTypes {
		_copy := instance.GongCopy().(*ArtefactType)
		stage.ArtefactTypes_reference[instance] = _copy
		stage.ArtefactTypes_instance[_copy] = instance
		stage.ArtefactTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ArtefactTypeShapes_reference = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	stage.ArtefactTypeShapes_referenceOrder = make(map[*ArtefactTypeShape]uint) // diff Unstage needs the reference order
	stage.ArtefactTypeShapes_instance = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	for instance := range stage.ArtefactTypeShapes {
		_copy := instance.GongCopy().(*ArtefactTypeShape)
		stage.ArtefactTypeShapes_reference[instance] = _copy
		stage.ArtefactTypeShapes_instance[_copy] = instance
		stage.ArtefactTypeShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Artists_reference = make(map[*Artist]*Artist)
	stage.Artists_referenceOrder = make(map[*Artist]uint) // diff Unstage needs the reference order
	stage.Artists_instance = make(map[*Artist]*Artist)
	for instance := range stage.Artists {
		_copy := instance.GongCopy().(*Artist)
		stage.Artists_reference[instance] = _copy
		stage.Artists_instance[_copy] = instance
		stage.Artists_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ArtistShapes_reference = make(map[*ArtistShape]*ArtistShape)
	stage.ArtistShapes_referenceOrder = make(map[*ArtistShape]uint) // diff Unstage needs the reference order
	stage.ArtistShapes_instance = make(map[*ArtistShape]*ArtistShape)
	for instance := range stage.ArtistShapes {
		_copy := instance.GongCopy().(*ArtistShape)
		stage.ArtistShapes_reference[instance] = _copy
		stage.ArtistShapes_instance[_copy] = instance
		stage.ArtistShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint) // diff Unstage needs the reference order
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	for instance := range stage.ControlPointShapes {
		_copy := instance.GongCopy().(*ControlPointShape)
		stage.ControlPointShapes_reference[instance] = _copy
		stage.ControlPointShapes_instance[_copy] = instance
		stage.ControlPointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Desks_reference = make(map[*Desk]*Desk)
	stage.Desks_referenceOrder = make(map[*Desk]uint) // diff Unstage needs the reference order
	stage.Desks_instance = make(map[*Desk]*Desk)
	for instance := range stage.Desks {
		_copy := instance.GongCopy().(*Desk)
		stage.Desks_reference[instance] = _copy
		stage.Desks_instance[_copy] = instance
		stage.Desks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Influences_reference = make(map[*Influence]*Influence)
	stage.Influences_referenceOrder = make(map[*Influence]uint) // diff Unstage needs the reference order
	stage.Influences_instance = make(map[*Influence]*Influence)
	for instance := range stage.Influences {
		_copy := instance.GongCopy().(*Influence)
		stage.Influences_reference[instance] = _copy
		stage.Influences_instance[_copy] = instance
		stage.Influences_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.InfluenceShapes_reference = make(map[*InfluenceShape]*InfluenceShape)
	stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint) // diff Unstage needs the reference order
	stage.InfluenceShapes_instance = make(map[*InfluenceShape]*InfluenceShape)
	for instance := range stage.InfluenceShapes {
		_copy := instance.GongCopy().(*InfluenceShape)
		stage.InfluenceShapes_reference[instance] = _copy
		stage.InfluenceShapes_instance[_copy] = instance
		stage.InfluenceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Movements_reference = make(map[*Movement]*Movement)
	stage.Movements_referenceOrder = make(map[*Movement]uint) // diff Unstage needs the reference order
	stage.Movements_instance = make(map[*Movement]*Movement)
	for instance := range stage.Movements {
		_copy := instance.GongCopy().(*Movement)
		stage.Movements_reference[instance] = _copy
		stage.Movements_instance[_copy] = instance
		stage.Movements_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MovementShapes_reference = make(map[*MovementShape]*MovementShape)
	stage.MovementShapes_referenceOrder = make(map[*MovementShape]uint) // diff Unstage needs the reference order
	stage.MovementShapes_instance = make(map[*MovementShape]*MovementShape)
	for instance := range stage.MovementShapes {
		_copy := instance.GongCopy().(*MovementShape)
		stage.MovementShapes_reference[instance] = _copy
		stage.MovementShapes_instance[_copy] = instance
		stage.MovementShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Places_reference = make(map[*Place]*Place)
	stage.Places_referenceOrder = make(map[*Place]uint) // diff Unstage needs the reference order
	stage.Places_instance = make(map[*Place]*Place)
	for instance := range stage.Places {
		_copy := instance.GongCopy().(*Place)
		stage.Places_reference[instance] = _copy
		stage.Places_instance[_copy] = instance
		stage.Places_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.ArtefactTypes {
		reference := stage.ArtefactTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ArtefactTypeShapes {
		reference := stage.ArtefactTypeShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Artists {
		reference := stage.Artists_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ArtistShapes {
		reference := stage.ArtistShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlPointShapes {
		reference := stage.ControlPointShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Desks {
		reference := stage.Desks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Influences {
		reference := stage.Influences_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.InfluenceShapes {
		reference := stage.InfluenceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Movements {
		reference := stage.Movements_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MovementShapes {
		reference := stage.MovementShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Places {
		reference := stage.Places_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (artefacttype *ArtefactType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtefactType_stagedOrder[artefacttype]; ok {
		return order
	}
	if order, ok := stage.ArtefactTypes_referenceOrder[artefacttype]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtefactType was not staged and does not have a reference order", artefacttype)
		return 0
	}
}

func (artefacttypeshape *ArtefactTypeShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtefactTypeShape_stagedOrder[artefacttypeshape]; ok {
		return order
	}
	if order, ok := stage.ArtefactTypeShapes_referenceOrder[artefacttypeshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtefactTypeShape was not staged and does not have a reference order", artefacttypeshape)
		return 0
	}
}

func (artist *Artist) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Artist_stagedOrder[artist]; ok {
		return order
	}
	if order, ok := stage.Artists_referenceOrder[artist]; ok {
		return order
	} else {
		log.Printf("instance %p of type Artist was not staged and does not have a reference order", artist)
		return 0
	}
}

func (artistshape *ArtistShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtistShape_stagedOrder[artistshape]; ok {
		return order
	}
	if order, ok := stage.ArtistShapes_referenceOrder[artistshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtistShape was not staged and does not have a reference order", artistshape)
		return 0
	}
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointShape_stagedOrder[controlpointshape]; ok {
		return order
	}
	if order, ok := stage.ControlPointShapes_referenceOrder[controlpointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlPointShape was not staged and does not have a reference order", controlpointshape)
		return 0
	}
}

func (desk *Desk) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Desk_stagedOrder[desk]; ok {
		return order
	}
	if order, ok := stage.Desks_referenceOrder[desk]; ok {
		return order
	} else {
		log.Printf("instance %p of type Desk was not staged and does not have a reference order", desk)
		return 0
	}
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
		return 0
	}
}

func (influence *Influence) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Influence_stagedOrder[influence]; ok {
		return order
	}
	if order, ok := stage.Influences_referenceOrder[influence]; ok {
		return order
	} else {
		log.Printf("instance %p of type Influence was not staged and does not have a reference order", influence)
		return 0
	}
}

func (influenceshape *InfluenceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InfluenceShape_stagedOrder[influenceshape]; ok {
		return order
	}
	if order, ok := stage.InfluenceShapes_referenceOrder[influenceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type InfluenceShape was not staged and does not have a reference order", influenceshape)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (movement *Movement) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Movement_stagedOrder[movement]; ok {
		return order
	}
	if order, ok := stage.Movements_referenceOrder[movement]; ok {
		return order
	} else {
		log.Printf("instance %p of type Movement was not staged and does not have a reference order", movement)
		return 0
	}
}

func (movementshape *MovementShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MovementShape_stagedOrder[movementshape]; ok {
		return order
	}
	if order, ok := stage.MovementShapes_referenceOrder[movementshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type MovementShape was not staged and does not have a reference order", movementshape)
		return 0
	}
}

func (place *Place) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Place_stagedOrder[place]; ok {
		return order
	}
	if order, ok := stage.Places_referenceOrder[place]; ok {
		return order
	} else {
		log.Printf("instance %p of type Place was not staged and does not have a reference order", place)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (artefacttype *ArtefactType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttype.GongGetGongstructName(), artefacttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttype *ArtefactType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttype.GongGetGongstructName(), artefacttype.GongGetOrder(stage))
}

func (artefacttypeshape *ArtefactTypeShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttypeshape.GongGetGongstructName(), artefacttypeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttypeshape *ArtefactTypeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttypeshape.GongGetGongstructName(), artefacttypeshape.GongGetOrder(stage))
}

func (artist *Artist) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artist.GongGetGongstructName(), artist.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artist *Artist) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artist.GongGetGongstructName(), artist.GongGetOrder(stage))
}

func (artistshape *ArtistShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artistshape.GongGetGongstructName(), artistshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artistshape *ArtistShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artistshape.GongGetGongstructName(), artistshape.GongGetOrder(stage))
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

func (desk *Desk) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (desk *Desk) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (influence *Influence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influence *Influence) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

func (influenceshape *InfluenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influenceshape *InfluenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (movement *Movement) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movement.GongGetGongstructName(), movement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movement *Movement) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movement.GongGetGongstructName(), movement.GongGetOrder(stage))
}

func (movementshape *MovementShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movementshape.GongGetGongstructName(), movementshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movementshape *MovementShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movementshape.GongGetGongstructName(), movementshape.GongGetOrder(stage))
}

func (place *Place) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", place.GongGetGongstructName(), place.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (place *Place) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", place.GongGetGongstructName(), place.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (artefacttype *ArtefactType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(artefacttype.Name))
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactTypeShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(artefacttypeshape.Name))
	return
}

func (artist *Artist) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artist.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Artist")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(artist.Name))
	return
}

func (artistshape *ArtistShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtistShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(artistshape.Name))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlpointshape.Name))
	return
}

func (desk *Desk) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Desk")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(desk.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagram.Name))
	return
}

func (influence *Influence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Influence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(influence.Name))
	return
}

func (influenceshape *InfluenceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InfluenceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(influenceshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (movement *Movement) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movement.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Movement")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(movement.Name))
	return
}

func (movementshape *MovementShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MovementShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(movementshape.Name))
	return
}

func (place *Place) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", place.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Place")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(place.Name))
	return
}

// insertion point for unstaging
func (artefacttype *ArtefactType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttype.GongGetReferenceIdentifier(stage))
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttypeshape.GongGetReferenceIdentifier(stage))
	return
}

func (artist *Artist) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artist.GongGetReferenceIdentifier(stage))
	return
}

func (artistshape *ArtistShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artistshape.GongGetReferenceIdentifier(stage))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetReferenceIdentifier(stage))
	return
}

func (desk *Desk) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (influence *Influence) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetReferenceIdentifier(stage))
	return
}

func (influenceshape *InfluenceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (movement *Movement) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movement.GongGetReferenceIdentifier(stage))
	return
}

func (movementshape *MovementShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movementshape.GongGetReferenceIdentifier(stage))
	return
}

func (place *Place) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", place.GongGetReferenceIdentifier(stage))
	return
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

// end of template
