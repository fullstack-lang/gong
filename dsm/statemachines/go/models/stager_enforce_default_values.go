package models



func (stager *Stager) enforceDefaultValues() (needCommit bool) {

	root := stager.getRootLibrary()
	if root != nil && root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true
	}

	return
}
