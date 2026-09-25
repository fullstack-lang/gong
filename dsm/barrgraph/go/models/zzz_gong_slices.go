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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.ArtefactTypes)

	res = __gong__appendInstances(res, stage.ArtefactTypeShapes)

	res = __gong__appendInstances(res, stage.Artists)

	res = __gong__appendInstances(res, stage.ArtistShapes)

	res = __gong__appendInstances(res, stage.ControlPointShapes)

	res = __gong__appendInstances(res, stage.Desks)

	res = __gong__appendInstances(res, stage.Diagrams)

	res = __gong__appendInstances(res, stage.Influences)

	res = __gong__appendInstances(res, stage.InfluenceShapes)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Movements)

	res = __gong__appendInstances(res, stage.MovementShapes)

	res = __gong__appendInstances(res, stage.Places)

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
func (artefacttype *ArtefactType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, artefacttype)
}

func (artefacttypeshape *ArtefactTypeShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, artefacttypeshape)
}

func (artist *Artist) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, artist)
}

func (artistshape *ArtistShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, artistshape)
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlpointshape)
}

func (desk *Desk) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, desk)
}

func (diagram *Diagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagram)
}

func (influence *Influence) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, influence)
}

func (influenceshape *InfluenceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, influenceshape)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (movement *Movement) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, movement)
}

func (movementshape *MovementShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, movementshape)
}

func (place *Place) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, place)
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
	__gong__computeReferencePass1(stage, stage.ArtefactTypes, &stage.ArtefactTypes_reference, &stage.ArtefactTypes_referenceOrder, &stage.ArtefactTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArtefactTypeShapes, &stage.ArtefactTypeShapes_reference, &stage.ArtefactTypeShapes_referenceOrder, &stage.ArtefactTypeShapes_instance)

	__gong__computeReferencePass1(stage, stage.Artists, &stage.Artists_reference, &stage.Artists_referenceOrder, &stage.Artists_instance)

	__gong__computeReferencePass1(stage, stage.ArtistShapes, &stage.ArtistShapes_reference, &stage.ArtistShapes_referenceOrder, &stage.ArtistShapes_instance)

	__gong__computeReferencePass1(stage, stage.ControlPointShapes, &stage.ControlPointShapes_reference, &stage.ControlPointShapes_referenceOrder, &stage.ControlPointShapes_instance)

	__gong__computeReferencePass1(stage, stage.Desks, &stage.Desks_reference, &stage.Desks_referenceOrder, &stage.Desks_instance)

	__gong__computeReferencePass1(stage, stage.Diagrams, &stage.Diagrams_reference, &stage.Diagrams_referenceOrder, &stage.Diagrams_instance)

	__gong__computeReferencePass1(stage, stage.Influences, &stage.Influences_reference, &stage.Influences_referenceOrder, &stage.Influences_instance)

	__gong__computeReferencePass1(stage, stage.InfluenceShapes, &stage.InfluenceShapes_reference, &stage.InfluenceShapes_referenceOrder, &stage.InfluenceShapes_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Movements, &stage.Movements_reference, &stage.Movements_referenceOrder, &stage.Movements_instance)

	__gong__computeReferencePass1(stage, stage.MovementShapes, &stage.MovementShapes_reference, &stage.MovementShapes_referenceOrder, &stage.MovementShapes_instance)

	__gong__computeReferencePass1(stage, stage.Places, &stage.Places_reference, &stage.Places_referenceOrder, &stage.Places_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.ArtefactTypes, stage.ArtefactTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArtefactTypeShapes, stage.ArtefactTypeShapes_reference, stage)

	__gong__computeReferencePass2(stage.Artists, stage.Artists_reference, stage)

	__gong__computeReferencePass2(stage.ArtistShapes, stage.ArtistShapes_reference, stage)

	__gong__computeReferencePass2(stage.ControlPointShapes, stage.ControlPointShapes_reference, stage)

	__gong__computeReferencePass2(stage.Desks, stage.Desks_reference, stage)

	__gong__computeReferencePass2(stage.Diagrams, stage.Diagrams_reference, stage)

	__gong__computeReferencePass2(stage.Influences, stage.Influences_reference, stage)

	__gong__computeReferencePass2(stage.InfluenceShapes, stage.InfluenceShapes_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Movements, stage.Movements_reference, stage)

	__gong__computeReferencePass2(stage.MovementShapes, stage.MovementShapes_reference, stage)

	__gong__computeReferencePass2(stage.Places, stage.Places_reference, stage)

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
	return __gong__getOrder(stage.ArtefactType_stagedOrder, stage.ArtefactTypes_referenceOrder, artefacttype, "ArtefactType")
}

func (artefacttypeshape *ArtefactTypeShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArtefactTypeShape_stagedOrder, stage.ArtefactTypeShapes_referenceOrder, artefacttypeshape, "ArtefactTypeShape")
}

func (artist *Artist) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Artist_stagedOrder, stage.Artists_referenceOrder, artist, "Artist")
}

func (artistshape *ArtistShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArtistShape_stagedOrder, stage.ArtistShapes_referenceOrder, artistshape, "ArtistShape")
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlPointShape_stagedOrder, stage.ControlPointShapes_referenceOrder, controlpointshape, "ControlPointShape")
}

func (desk *Desk) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Desk_stagedOrder, stage.Desks_referenceOrder, desk, "Desk")
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Diagram_stagedOrder, stage.Diagrams_referenceOrder, diagram, "Diagram")
}

func (influence *Influence) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Influence_stagedOrder, stage.Influences_referenceOrder, influence, "Influence")
}

func (influenceshape *InfluenceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InfluenceShape_stagedOrder, stage.InfluenceShapes_referenceOrder, influenceshape, "InfluenceShape")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (movement *Movement) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Movement_stagedOrder, stage.Movements_referenceOrder, movement, "Movement")
}

func (movementshape *MovementShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MovementShape_stagedOrder, stage.MovementShapes_referenceOrder, movementshape, "MovementShape")
}

func (place *Place) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Place_stagedOrder, stage.Places_referenceOrder, place, "Place")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (artefacttype *ArtefactType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(artefacttype, artefacttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttype *ArtefactType) GongGetReferenceIdentifier(stage *Stage) string {
	return artefacttype.GongGetIdentifier(stage)
}

func (artefacttypeshape *ArtefactTypeShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(artefacttypeshape, artefacttypeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttypeshape *ArtefactTypeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return artefacttypeshape.GongGetIdentifier(stage)
}

func (artist *Artist) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(artist, artist.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artist *Artist) GongGetReferenceIdentifier(stage *Stage) string {
	return artist.GongGetIdentifier(stage)
}

func (artistshape *ArtistShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(artistshape, artistshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artistshape *ArtistShape) GongGetReferenceIdentifier(stage *Stage) string {
	return artistshape.GongGetIdentifier(stage)
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlpointshape, controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return controlpointshape.GongGetIdentifier(stage)
}

func (desk *Desk) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(desk, desk.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (desk *Desk) GongGetReferenceIdentifier(stage *Stage) string {
	return desk.GongGetIdentifier(stage)
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagram, diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return diagram.GongGetIdentifier(stage)
}

func (influence *Influence) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(influence, influence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influence *Influence) GongGetReferenceIdentifier(stage *Stage) string {
	return influence.GongGetIdentifier(stage)
}

func (influenceshape *InfluenceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(influenceshape, influenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influenceshape *InfluenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return influenceshape.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (movement *Movement) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(movement, movement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movement *Movement) GongGetReferenceIdentifier(stage *Stage) string {
	return movement.GongGetIdentifier(stage)
}

func (movementshape *MovementShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(movementshape, movementshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movementshape *MovementShape) GongGetReferenceIdentifier(stage *Stage) string {
	return movementshape.GongGetIdentifier(stage)
}

func (place *Place) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(place, place.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (place *Place) GongGetReferenceIdentifier(stage *Stage) string {
	return place.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (artefacttype *ArtefactType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(artefacttype.GongGetIdentifier(stage), "ArtefactType", artefacttype.Name)
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(artefacttypeshape.GongGetIdentifier(stage), "ArtefactTypeShape", artefacttypeshape.Name)
}

func (artist *Artist) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(artist.GongGetIdentifier(stage), "Artist", artist.Name)
}

func (artistshape *ArtistShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(artistshape.GongGetIdentifier(stage), "ArtistShape", artistshape.Name)
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlpointshape.GongGetIdentifier(stage), "ControlPointShape", controlpointshape.Name)
}

func (desk *Desk) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(desk.GongGetIdentifier(stage), "Desk", desk.Name)
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagram.GongGetIdentifier(stage), "Diagram", diagram.Name)
}

func (influence *Influence) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(influence.GongGetIdentifier(stage), "Influence", influence.Name)
}

func (influenceshape *InfluenceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(influenceshape.GongGetIdentifier(stage), "InfluenceShape", influenceshape.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (movement *Movement) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(movement.GongGetIdentifier(stage), "Movement", movement.Name)
}

func (movementshape *MovementShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(movementshape.GongGetIdentifier(stage), "MovementShape", movementshape.Name)
}

func (place *Place) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(place.GongGetIdentifier(stage), "Place", place.Name)
}

// insertion point for unstaging
func (artefacttype *ArtefactType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(artefacttype.GongGetReferenceIdentifier(stage))
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(artefacttypeshape.GongGetReferenceIdentifier(stage))
}

func (artist *Artist) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(artist.GongGetReferenceIdentifier(stage))
}

func (artistshape *ArtistShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(artistshape.GongGetReferenceIdentifier(stage))
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlpointshape.GongGetReferenceIdentifier(stage))
}

func (desk *Desk) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(desk.GongGetReferenceIdentifier(stage))
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagram.GongGetReferenceIdentifier(stage))
}

func (influence *Influence) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(influence.GongGetReferenceIdentifier(stage))
}

func (influenceshape *InfluenceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(influenceshape.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (movement *Movement) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(movement.GongGetReferenceIdentifier(stage))
}

func (movementshape *MovementShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(movementshape.GongGetReferenceIdentifier(stage))
}

func (place *Place) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(place.GongGetReferenceIdentifier(stage))
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
