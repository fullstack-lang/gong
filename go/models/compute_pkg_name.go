package models

import (
	"log"
	"path/filepath"
	"strings"
)

// ComputePkgNameFromPkgPath computes the name of the package from the current working directory
func ComputePkgNameFromPkgPath(pkgPathArg string) (pkgName string) {
	// compute name of package
	abs, _ := filepath.Abs(filepath.Join(pkgPathArg, "../.."))
	log.Println("Abs is " + abs)

	// to slash to have standardized separators between unix and windows
	abs = filepath.ToSlash(abs)

	dirs := strings.Split(abs, "/")
	pkgName = dirs[len(dirs)-1]
	log.Println("PkgName is " + pkgName)

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
