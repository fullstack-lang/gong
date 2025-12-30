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

	// insertion point per named struct
	var filetodownloads_newInstances []*FileToDownload
	var filetodownloads_deletedInstances []*FileToDownload

	// parse all staged instances and check if they have a reference
	for filetodownload := range stage.FileToDownloads {
		if ref, ok := stage.FileToDownloads_reference[filetodownload]; !ok {
			filetodownloads_newInstances = append(filetodownloads_newInstances, filetodownload)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FileToDownload "+filetodownload.Name,
				)
			}
		} else {
			diffs := filetodownload.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FileToDownload \""+filetodownload.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for filetodownload := range stage.FileToDownloads_reference {
		if _, ok := stage.FileToDownloads[filetodownload]; !ok {
			filetodownloads_deletedInstances = append(filetodownloads_deletedInstances, filetodownload)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FileToDownload "+filetodownload.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FileToUpload "+filetoupload.Name,
				)
			}
		} else {
			diffs := filetoupload.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FileToUpload \""+filetoupload.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for filetoupload := range stage.FileToUploads_reference {
		if _, ok := stage.FileToUploads[filetoupload]; !ok {
			filetouploads_deletedInstances = append(filetouploads_deletedInstances, filetoupload)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FileToUpload "+filetoupload.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Message "+message.Name,
				)
			}
		} else {
			diffs := message.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Message \""+message.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for message := range stage.Messages_reference {
		if _, ok := stage.Messages[message]; !ok {
			messages_deletedInstances = append(messages_deletedInstances, message)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Message "+message.Name,
				)
			}
		}
	}

	lenNewInstances += len(messages_newInstances)
	lenDeletedInstances += len(messages_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
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
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FileToDownload")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", filetodownload.Name)
	return
}
func (filetoupload *FileToUpload) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", filetoupload.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FileToUpload")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", filetoupload.Name)
	return
}
func (message *Message) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", message.Name)
	return
}
