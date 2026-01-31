// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct FileToDownload
	// insertion point per field

	// Compute reverse map for named struct FileToUpload
	// insertion point per field

	// Compute reverse map for named struct Message
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.FileToDownloads {
		res = append(res, instance)
	}

	for instance := range stage.FileToUploads {
		res = append(res, instance)
	}

	for instance := range stage.Messages {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (filetodownload *FileToDownload) GongCopy() GongstructIF {
	newInstance := *filetodownload
	return &newInstance
}

func (filetoupload *FileToUpload) GongCopy() GongstructIF {
	newInstance := *filetoupload
	return &newInstance
}

func (message *Message) GongCopy() GongstructIF {
	newInstance := *message
	return &newInstance
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
	var filetodownloads_newInstances []*FileToDownload
	var filetodownloads_deletedInstances []*FileToDownload

	// parse all staged instances and check if they have a reference
	for filetodownload := range stage.FileToDownloads {
		if ref, ok := stage.FileToDownloads_reference[filetodownload]; !ok {
			filetodownloads_newInstances = append(filetodownloads_newInstances, filetodownload)
			newInstancesSlice = append(newInstancesSlice, filetodownload.GongMarshallIdentifier(stage))
			if stage.FileToDownloads_referenceOrder == nil {
				stage.FileToDownloads_referenceOrder = make(map[*FileToDownload]uint)
			}
			stage.FileToDownloads_referenceOrder[filetodownload] = stage.FileToDownloadMap_Staged_Order[filetodownload]
			newInstancesReverseSlice = append(newInstancesReverseSlice, filetodownload.GongMarshallUnstaging(stage))
			delete(stage.FileToDownloads_referenceOrder, filetodownload)
			fieldInitializers, pointersInitializations := filetodownload.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FileToDownloadMap_Staged_Order[ref] = stage.FileToDownloadMap_Staged_Order[filetodownload]
			diffs := filetodownload.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, filetodownload)
			delete(stage.FileToDownloadMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", filetodownload.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FileToDownloads_reference {
		if _, ok := stage.FileToDownloads[ref]; !ok {
			filetodownloads_deletedInstances = append(filetodownloads_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(filetodownloads_newInstances)
	lenDeletedInstances += len(filetodownloads_deletedInstances)
	var filetouploads_newInstances []*FileToUpload
	var filetouploads_deletedInstances []*FileToUpload

	// parse all staged instances and check if they have a reference
	for filetoupload := range stage.FileToUploads {
		if ref, ok := stage.FileToUploads_reference[filetoupload]; !ok {
			filetouploads_newInstances = append(filetouploads_newInstances, filetoupload)
			newInstancesSlice = append(newInstancesSlice, filetoupload.GongMarshallIdentifier(stage))
			if stage.FileToUploads_referenceOrder == nil {
				stage.FileToUploads_referenceOrder = make(map[*FileToUpload]uint)
			}
			stage.FileToUploads_referenceOrder[filetoupload] = stage.FileToUploadMap_Staged_Order[filetoupload]
			newInstancesReverseSlice = append(newInstancesReverseSlice, filetoupload.GongMarshallUnstaging(stage))
			delete(stage.FileToUploads_referenceOrder, filetoupload)
			fieldInitializers, pointersInitializations := filetoupload.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FileToUploadMap_Staged_Order[ref] = stage.FileToUploadMap_Staged_Order[filetoupload]
			diffs := filetoupload.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, filetoupload)
			delete(stage.FileToUploadMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", filetoupload.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FileToUploads_reference {
		if _, ok := stage.FileToUploads[ref]; !ok {
			filetouploads_deletedInstances = append(filetouploads_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(filetouploads_newInstances)
	lenDeletedInstances += len(filetouploads_deletedInstances)
	var messages_newInstances []*Message
	var messages_deletedInstances []*Message

	// parse all staged instances and check if they have a reference
	for message := range stage.Messages {
		if ref, ok := stage.Messages_reference[message]; !ok {
			messages_newInstances = append(messages_newInstances, message)
			newInstancesSlice = append(newInstancesSlice, message.GongMarshallIdentifier(stage))
			if stage.Messages_referenceOrder == nil {
				stage.Messages_referenceOrder = make(map[*Message]uint)
			}
			stage.Messages_referenceOrder[message] = stage.MessageMap_Staged_Order[message]
			newInstancesReverseSlice = append(newInstancesReverseSlice, message.GongMarshallUnstaging(stage))
			delete(stage.Messages_referenceOrder, message)
			fieldInitializers, pointersInitializations := message.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MessageMap_Staged_Order[ref] = stage.MessageMap_Staged_Order[message]
			diffs := message.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, message)
			delete(stage.MessageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", message.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Messages_reference {
		if _, ok := stage.Messages[ref]; !ok {
			messages_deletedInstances = append(messages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(messages_newInstances)
	lenDeletedInstances += len(messages_deletedInstances)

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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.FileToDownloads_reference = make(map[*FileToDownload]*FileToDownload)
	stage.FileToDownloads_referenceOrder = make(map[*FileToDownload]uint) // diff Unstage needs the reference order
	for instance := range stage.FileToDownloads {
		stage.FileToDownloads_reference[instance] = instance.GongCopy().(*FileToDownload)
		stage.FileToDownloads_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FileToUploads_reference = make(map[*FileToUpload]*FileToUpload)
	stage.FileToUploads_referenceOrder = make(map[*FileToUpload]uint) // diff Unstage needs the reference order
	for instance := range stage.FileToUploads {
		stage.FileToUploads_reference[instance] = instance.GongCopy().(*FileToUpload)
		stage.FileToUploads_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Messages_reference = make(map[*Message]*Message)
	stage.Messages_referenceOrder = make(map[*Message]uint) // diff Unstage needs the reference order
	for instance := range stage.Messages {
		stage.Messages_reference[instance] = instance.GongCopy().(*Message)
		stage.Messages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (filetodownload *FileToDownload) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FileToDownloadMap_Staged_Order[filetodownload]; ok {
		return order
	}
	return stage.FileToDownloads_referenceOrder[filetodownload]
}

func (filetodownload *FileToDownload) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FileToDownloads_referenceOrder[filetodownload]
}

func (filetoupload *FileToUpload) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FileToUploadMap_Staged_Order[filetoupload]; ok {
		return order
	}
	return stage.FileToUploads_referenceOrder[filetoupload]
}

func (filetoupload *FileToUpload) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FileToUploads_referenceOrder[filetoupload]
}

func (message *Message) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MessageMap_Staged_Order[message]; ok {
		return order
	}
	return stage.Messages_referenceOrder[message]
}

func (message *Message) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Messages_referenceOrder[message]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (filetodownload *FileToDownload) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetodownload.GongGetGongstructName(), filetodownload.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (filetodownload *FileToDownload) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetodownload.GongGetGongstructName(), filetodownload.GongGetReferenceOrder(stage))
}

func (filetoupload *FileToUpload) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetoupload.GongGetGongstructName(), filetoupload.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (filetoupload *FileToUpload) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetoupload.GongGetGongstructName(), filetoupload.GongGetReferenceOrder(stage))
}

func (message *Message) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (message *Message) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (filetodownload *FileToDownload) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FileToDownload")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", filetodownload.Name)
	return
}
func (filetoupload *FileToUpload) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetoupload.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FileToUpload")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", filetoupload.Name)
	return
}
func (message *Message) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", message.Name)
	return
}

// insertion point for unstaging
func (filetodownload *FileToDownload) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetReferenceIdentifier(stage))
	return
}
func (filetoupload *FileToUpload) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetoupload.GongGetReferenceIdentifier(stage))
	return
}
func (message *Message) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetReferenceIdentifier(stage))
	return
}
