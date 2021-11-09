package main

import (
	"log"
	"path/filepath"
	"strings"
)

func computePkgName() (pkgName string) {
	// compute name of package
	abs, _ := filepath.Abs(filepath.Join(*pkgPath, "../.."))
	log.Println("Abs is " + abs)

	// to slash to have standardized separators between unix and windows
	abs = filepath.ToSlash(abs)

	dirs := strings.Split(abs, "/")
	pkgName = dirs[len(dirs)-1]
	log.Println("PkgName is " + pkgName)

	return
}
