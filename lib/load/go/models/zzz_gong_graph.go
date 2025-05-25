// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *FileToDownload:
		ok = stage.IsStagedFileToDownload(target)

	case *FileToUpload:
		ok = stage.IsStagedFileToUpload(target)

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
