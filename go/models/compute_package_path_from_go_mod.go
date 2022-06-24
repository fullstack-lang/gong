package models

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func ComputePkgPathFromGoModFile(pkgPathArg string) (pkgName, fullPkgPath string) {
	// compute the number of steps to reach the
	pathToTheModel, _ := filepath.Abs(pkgPathArg)
	_ = pathToTheModel
	dirs := strings.Split(pathToTheModel, string(os.PathSeparator))
	_ = dirs

	dirs = dirs[:len(dirs)-2]
	log.Println(filepath.Join(dirs...))

	// parse the above directories for find a go.mod file
	var nbOfLevelBetweenPackageAndModule int

	// if the package starts at the module level, then this path is empty
	// for instance "gongsvg" and "github.com/fullstack-lang/gongsvg"
	//
	// if the package starts below the module level, then this path is not empty
	// "test" and "github.com/fullstack-lang/gong" where "test" is in "github.com/fullstack-lang/gong/test"
	var goModFileMissing = true
	var goModFilePath string

	// loop that search up for the "go.mod" file
	for nbOfLevelBetweenPackageAndModule = 2; // start at 2 levels above the "models" package since the "go.mod" file closest place is 2 levels
	nbOfLevelBetweenPackageAndModule < 7 &&   // cannot go above 7
		goModFileMissing; // stop is go.mod file found
	nbOfLevelBetweenPackageAndModule++ {
		var pathBetweenPackageAndModule = "."
		for i := 0; i < nbOfLevelBetweenPackageAndModule; i++ {
			pathBetweenPackageAndModule = filepath.Join(pathBetweenPackageAndModule, "..")
		}
		goModFilePath = filepath.Join(pkgPathArg, pathBetweenPackageAndModule, "go.mod")
		_, errGoModFile := os.Stat(goModFilePath)
		goModFileMissing = os.IsNotExist(errGoModFile)
	}

	pkgName = ComputePkgNameFromPkgPath(pkgPathArg)
	if goModFileMissing {
		cmd := exec.Command("go", "mod", "init", pkgName)
		cmd.Dir, _ = filepath.Abs(filepath.Join(pkgPathArg, "../.."))
		log.Printf("Running %s command in directory %s and waiting for it to finish...\n", cmd.Args, cmd.Dir)

		// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		nbOfLevelBetweenPackageAndModule = 2
		var pathBetweenPackageAndModule = "."
		for i := 0; i < nbOfLevelBetweenPackageAndModule; i++ {
			pathBetweenPackageAndModule = filepath.Join(pathBetweenPackageAndModule, "..")
		}
		goModFilePath = filepath.Join(pkgPathArg, pathBetweenPackageAndModule, "go.mod")
	}

	// read the go.mod file
	buf, err := ioutil.ReadFile(goModFilePath)
	if err != nil {
		log.Fatalln("Cannot read go.mod file at ", goModFilePath)
	}
	modFile, err := modfile.Parse(goModFilePath, buf, nil)
	if err != nil {
		panic(err)
	}
	log.Println("module path is ", modFile.Module.Mod.Path)
	_ = modFile // now we have mod file

	// load package into database
	// module path + (if necessary) relative path to the package
	// for instance "github.com/fullstack-lang/gongsvg" or
	// for instance "github.com/fullstack-lang/gong/test" or
	pikPlace := len(dirs) - nbOfLevelBetweenPackageAndModule + 3
	if pikPlace > len(dirs) {
		pikPlace = len(dirs)
	}
	relPath := dirs[pikPlace:]
	joinedPath := append([]string{modFile.Module.Mod.Path}, relPath...)

	fullPkgPath = filepath.Join(joinedPath...)
	fullPkgPath = filepath.Join(fullPkgPath, "go", "models")
	// case for windows
	fullPkgPath = strings.ReplaceAll(fullPkgPath, "\\", "/")

	return
}
