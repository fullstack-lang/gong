// generated code - do not edit
package models

import (
	"fmt"
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += filetodownload.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := filetodownload.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := filetodownload.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", filetodownload.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for filetodownload := range stage.FileToDownloads_reference {
		if _, ok := stage.FileToDownloads[filetodownload]; !ok {
			filetodownloads_deletedInstances = append(filetodownloads_deletedInstances, filetodownload)
			deletedInstancesStmt += filetodownload.GongMarshallUnstaging(stage)
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
			newInstancesStmt += filetoupload.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := filetoupload.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := filetoupload.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", filetoupload.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for filetoupload := range stage.FileToUploads_reference {
		if _, ok := stage.FileToUploads[filetoupload]; !ok {
			filetouploads_deletedInstances = append(filetouploads_deletedInstances, filetoupload)
			deletedInstancesStmt += filetoupload.GongMarshallUnstaging(stage)
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
			newInstancesStmt += message.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := message.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := message.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", message.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for message := range stage.Messages_reference {
		if _, ok := stage.Messages[message]; !ok {
			messages_deletedInstances = append(messages_deletedInstances, message)
			deletedInstancesStmt += message.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(messages_newInstances)
	lenDeletedInstances += len(messages_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.FileToDownloads_reference = make(map[*FileToDownload]*FileToDownload)
	for instance := range stage.FileToDownloads {
		stage.FileToDownloads_reference[instance] = instance.GongCopy().(*FileToDownload)
	}

	stage.FileToUploads_reference = make(map[*FileToUpload]*FileToUpload)
	for instance := range stage.FileToUploads {
		stage.FileToUploads_reference[instance] = instance.GongCopy().(*FileToUpload)
	}

	stage.Messages_reference = make(map[*Message]*Message)
	for instance := range stage.Messages {
		stage.Messages_reference[instance] = instance.GongCopy().(*Message)
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
	return stage.FileToDownloadMap_Staged_Order[filetodownload]
}

func (filetoupload *FileToUpload) GongGetOrder(stage *Stage) uint {
	return stage.FileToUploadMap_Staged_Order[filetoupload]
}

func (message *Message) GongGetOrder(stage *Stage) uint {
	return stage.MessageMap_Staged_Order[message]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (filetodownload *FileToDownload) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetodownload.GongGetGongstructName(), filetodownload.GongGetOrder(stage))
}

func (filetoupload *FileToUpload) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", filetoupload.GongGetGongstructName(), filetoupload.GongGetOrder(stage))
}

func (message *Message) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
	return
}
func (filetoupload *FileToUpload) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetoupload.GongGetIdentifier(stage))
	return
}
func (message *Message) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	return
}
