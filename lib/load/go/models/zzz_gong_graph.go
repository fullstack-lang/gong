// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (filetodownload *FileToDownload) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FileToDownloads[filetodownload]

	return
}

func (stage *Stage) IsStagedFileToDownload(filetodownload *FileToDownload) (ok bool) {

	return filetodownload.GongIsStaged(stage)
}

func (filetoupload *FileToUpload) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FileToUploads[filetoupload]

	return
}

func (stage *Stage) IsStagedFileToUpload(filetoupload *FileToUpload) (ok bool) {

	return filetoupload.GongIsStaged(stage)
}

func (message *Message) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

func (stage *Stage) IsStagedMessage(message *Message) (ok bool) {

	return message.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (filetodownload *FileToDownload) GongStageBranch(stage *Stage) {
	stage.StageBranchFileToDownload(filetodownload)
}

func (stage *Stage) StageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (filetoupload *FileToUpload) GongStageBranch(stage *Stage) {
	stage.StageBranchFileToUpload(filetoupload)
}

func (stage *Stage) StageBranchFileToUpload(filetoupload *FileToUpload) {

	// check if instance is already staged
	if stage.IsStaged(filetoupload) {
		return
	}

	filetoupload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (message *Message) GongStageBranch(stage *Stage) {
	stage.StageBranchMessage(message)
}

func (stage *Stage) StageBranchMessage(message *Message) {

	// check if instance is already staged
	if stage.IsStaged(message) {
		return
	}

	message.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *FileToDownload:
		toT := GongCopyBranchFileToDownload(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FileToUpload:
		toT := GongCopyBranchFileToUpload(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Message:
		toT := GongCopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchFileToDownload(mapOrigCopy map[any]any, filetodownloadFrom *FileToDownload) (filetodownloadTo *FileToDownload) {

	// filetodownloadFrom has already been copied
	if _filetodownloadTo, ok := mapOrigCopy[filetodownloadFrom]; ok {
		filetodownloadTo = _filetodownloadTo.(*FileToDownload)
		return
	}

	filetodownloadTo = new(FileToDownload)
	mapOrigCopy[filetodownloadFrom] = filetodownloadTo
	filetodownloadFrom.GongCopyBasicFields(filetodownloadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFileToUpload(mapOrigCopy map[any]any, filetouploadFrom *FileToUpload) (filetouploadTo *FileToUpload) {

	// filetouploadFrom has already been copied
	if _filetouploadTo, ok := mapOrigCopy[filetouploadFrom]; ok {
		filetouploadTo = _filetouploadTo.(*FileToUpload)
		return
	}

	filetouploadTo = new(FileToUpload)
	mapOrigCopy[filetouploadFrom] = filetouploadTo
	filetouploadFrom.GongCopyBasicFields(filetouploadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {

	// messageFrom has already been copied
	if _messageTo, ok := mapOrigCopy[messageFrom]; ok {
		messageTo = _messageTo.(*Message)
		return
	}

	messageTo = new(Message)
	mapOrigCopy[messageFrom] = messageTo
	messageFrom.GongCopyBasicFields(messageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (filetodownload *FileToDownload) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFileToDownload(filetodownload)
}

func (stage *Stage) UnstageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if !stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (filetoupload *FileToUpload) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFileToUpload(filetoupload)
}

func (stage *Stage) UnstageBranchFileToUpload(filetoupload *FileToUpload) {

	// check if instance is already staged
	if !stage.IsStaged(filetoupload) {
		return
	}

	filetoupload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (message *Message) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMessage(message)
}

func (stage *Stage) UnstageBranchMessage(message *Message) {

	// check if instance is already staged
	if !stage.IsStaged(message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *FileToDownload) GongReconstructPointersFromReferences(stage *Stage, instance *FileToDownload) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FileToUpload) GongReconstructPointersFromReferences(stage *Stage, instance *FileToUpload) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Message) GongReconstructPointersFromReferences(stage *Stage, instance *Message) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *FileToDownload) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FileToUpload) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Message) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (filetodownload *FileToDownload) GongDiff(stage *Stage, filetodownloadOther *FileToDownload) (diffs []string) {
	// insertion point for field diffs
	if filetodownload.Name != filetodownloadOther.Name {
		diffs = append(diffs, filetodownload.GongMarshallField(stage, "Name"))
	}
	if filetodownload.Base64EncodedContent != filetodownloadOther.Base64EncodedContent {
		diffs = append(diffs, filetodownload.GongMarshallField(stage, "Base64EncodedContent"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (filetoupload *FileToUpload) GongDiff(stage *Stage, filetouploadOther *FileToUpload) (diffs []string) {
	// insertion point for field diffs
	if filetoupload.Name != filetouploadOther.Name {
		diffs = append(diffs, filetoupload.GongMarshallField(stage, "Name"))
	}
	if filetoupload.Base64EncodedContent != filetouploadOther.Base64EncodedContent {
		diffs = append(diffs, filetoupload.GongMarshallField(stage, "Base64EncodedContent"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (message *Message) GongDiff(stage *Stage, messageOther *Message) (diffs []string) {
	// insertion point for field diffs
	if message.Name != messageOther.Name {
		diffs = append(diffs, message.GongMarshallField(stage, "Name"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff[T1, T2 PointerToGongstruct](a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
