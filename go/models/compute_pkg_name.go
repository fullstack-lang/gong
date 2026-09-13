package models

import (
	"log"
	"path/filepath"
	"strings"
)

// ComputePkgNameFromPkgPath computes the name of the package from the current working directory
func ComputePkgNameFromPkgPath(pkgPathArg string) (pkgName string) {
	abs, _ := filepath.Abs(pkgPathArg)
	abs = filepath.ToSlash(abs)

	// If the path is in or under /go/models, the stack directory is before /go/models
	if idx := strings.Index(abs, "/go/models"); idx != -1 {
		stackDir := abs[:idx]
		dirs := strings.Split(stackDir, "/")
		pkgName = dirs[len(dirs)-1]
	} else {
		absTwoLevelsUp, _ := filepath.Abs(filepath.Join(pkgPathArg, "../.."))
		absTwoLevelsUp = filepath.ToSlash(absTwoLevelsUp)
		dirs := strings.Split(absTwoLevelsUp, "/")
		pkgName = dirs[len(dirs)-1]
	}
	// log.Println("PkgName is " + pkgName)

	// check name

	// no upper cases allowed
	if strings.ToLower(pkgName) != pkgName {
		log.Fatalln("only lower cases are allowed in directory/package name, because angular workspace does not support it", pkgName)
	}

	if strings.Contains(pkgName, "_") {
		log.Fatalln("underscores '_' are not allowed in directory/package name, because angular workspace does not support it", pkgName)
	}

	return
}
