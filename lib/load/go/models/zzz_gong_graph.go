// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *FileToDownload:
		ok = stage.IsStagedFileToDownload(target)

	case *FileToUpload:
		ok = stage.IsStagedFileToUpload(target)

	case *Message:
		ok = stage.IsStagedMessage(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *FileToDownload:
		ok = stage.IsStagedFileToDownload(target)

	case *FileToUpload:
		ok = stage.IsStagedFileToUpload(target)

	case *Message:
		ok = stage.IsStagedMessage(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedFileToDownload(filetodownload *FileToDownload) (ok bool) {

	_, ok = stage.FileToDownloads[filetodownload]

	return
}

func (stage *Stage) IsStagedFileToUpload(filetoupload *FileToUpload) (ok bool) {

	_, ok = stage.FileToUploads[filetoupload]

	return
}

func (stage *Stage) IsStagedMessage(message *Message) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *FileToDownload:
		stage.StageBranchFileToDownload(target)

	case *FileToUpload:
		stage.StageBranchFileToUpload(target)

	case *Message:
		stage.StageBranchMessage(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if IsStaged(stage, filetodownload) {
		return
	}

	filetodownload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFileToUpload(filetoupload *FileToUpload) {

	// check if instance is already staged
	if IsStaged(stage, filetoupload) {
		return
	}

	filetoupload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMessage(message *Message) {

	// check if instance is already staged
	if IsStaged(stage, message) {
		return
	}

	message.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *FileToDownload:
		toT := CopyBranchFileToDownload(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FileToUpload:
		toT := CopyBranchFileToUpload(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Message:
		toT := CopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchFileToDownload(mapOrigCopy map[any]any, filetodownloadFrom *FileToDownload) (filetodownloadTo *FileToDownload) {

	// filetodownloadFrom has already been copied
	if _filetodownloadTo, ok := mapOrigCopy[filetodownloadFrom]; ok {
		filetodownloadTo = _filetodownloadTo.(*FileToDownload)
		return
	}

	filetodownloadTo = new(FileToDownload)
	mapOrigCopy[filetodownloadFrom] = filetodownloadTo
	filetodownloadFrom.CopyBasicFields(filetodownloadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFileToUpload(mapOrigCopy map[any]any, filetouploadFrom *FileToUpload) (filetouploadTo *FileToUpload) {

	// filetouploadFrom has already been copied
	if _filetouploadTo, ok := mapOrigCopy[filetouploadFrom]; ok {
		filetouploadTo = _filetouploadTo.(*FileToUpload)
		return
	}

	filetouploadTo = new(FileToUpload)
	mapOrigCopy[filetouploadFrom] = filetouploadTo
	filetouploadFrom.CopyBasicFields(filetouploadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {

	// messageFrom has already been copied
	if _messageTo, ok := mapOrigCopy[messageFrom]; ok {
		messageTo = _messageTo.(*Message)
		return
	}

	messageTo = new(Message)
	mapOrigCopy[messageFrom] = messageTo
	messageFrom.CopyBasicFields(messageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *FileToDownload:
		stage.UnstageBranchFileToDownload(target)

	case *FileToUpload:
		stage.UnstageBranchFileToUpload(target)

	case *Message:
		stage.UnstageBranchMessage(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if !IsStaged(stage, filetodownload) {
		return
	}

	filetodownload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFileToUpload(filetoupload *FileToUpload) {

	// check if instance is already staged
	if !IsStaged(stage, filetoupload) {
		return
	}

	filetoupload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMessage(message *Message) {

	// check if instance is already staged
	if !IsStaged(stage, message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
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
			ops += fmt.Sprintf("\t%s.%s = slices.Delete( %s.%s, %d, %d)\n", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
			ops += fmt.Sprintf("\t%s.%s = slices.Insert( %s.%s, %d, %s)\n",  a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
