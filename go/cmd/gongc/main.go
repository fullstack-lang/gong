package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/fullstack-lang/gong/go/vscode"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const COMPUTED_FROM_PKG_PATH string = "computed from pkgPath (path to package for analysis)"

var (
	pkgPath     = flag.String("pkgPath", ".", "path to the models package to be compiled by gongc compilation")
	skipSwagger = flag.Bool("skipSwagger", true, "skip swagger")
	skipNg      = flag.Bool("skipNg", false, "generates skipNg")
	skipFlutter = flag.Bool("skipFlutter", true, "do not generate flutter front")
	skipCoder   = flag.Bool("skipCoder", true, "do not generate coder file")

	addr = flag.String("addr", "localhost:8080/api",
		"network address addr where the angular generated service will lookup the server")
	run               = flag.Bool("run", false, "run 'go run main.go' after compilation")
	skipGoModCommands = flag.Bool("skipGoModCommands", false, "avoid calls to go mod init, tidy and vendor")

	skipNpmInstall = flag.Bool("skipNpmInstall", false, "skip the npm install command")

	compileForDebug = flag.Bool("compileForDebug", false, "Delve can be slow to start (more than 60'). A workaround is to generate a go build with with '-N -l' options")
)

func main() {

	log.SetPrefix("gongc: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatal("surplus arguments")
	}
	if len(flag.Args()) == 1 {
		*pkgPath = flag.Arg(0)
	}

	// remove gong generated files
	golang.RemoveGeneratedGongFilesButDocs(*pkgPath)

	// initiate model package
	modelPkg, _ := gong_models.LoadSource(*pkgPath)

	// check wether the package name follows gong naming convention
	if strings.ContainsAny(modelPkg.Name, "-") {
		log.Panicln(modelPkg.Name + " is not OK for a gong package name because it contains a - (dash) " +
			"and it cannot be used for naming a typescript class (the generated module in this cause)")
	}

	//
	// compute paths to generation directory
	//
	{

		directory, err := filepath.Abs(filepath.Join(*pkgPath, ".."))
		if err != nil {
			log.Fatal("Problem with backend target path " + err.Error())
		}

		// check existance of path
		fileInfo, err := os.Stat(directory)
		if os.IsNotExist(err) {
			log.Panicf("Folder %s does not exist.", directory)
		}
		if fileInfo.Mode().Perm()&(1<<(uint(7))) == 0 {
			log.Panicf("Folder %s is not writtable", directory)
		}
		gong_models.PathToGoSubDirectory = directory
		log.Println("backend target path " + gong_models.PathToGoSubDirectory)
	}

	// check wether one is in a git managed directory. If absent, use "git init"
	{
		gitStatusCommand := exec.Command("git", "status")

		var isGitActive bool

		// Execute the command
		if err := gitStatusCommand.Run(); err != nil {
			isGitActive = false
			log.Println("not a git directory")
		} else {
			isGitActive = true
			log.Println("a git directory")
		}

		if !isGitActive {
			log.Printf("git is not active, hence gong is generating a git directory with git init command")

			// git init
			cmd := exec.Command("git", "init")
			cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, "../.."))

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
		}
	}

	// generate diagrams/docs.go is absent
	{
		diagramsDocFilePath := filepath.Join(*pkgPath, "../diagrams/docs.go")
		_, errd := os.Stat(diagramsDocFilePath)
		if os.IsNotExist(errd) {
			log.Printf("../diagrams/docs.go does not exist, gongc creates a default one")

			diagramsDocFileDirPath := filepath.Dir(diagramsDocFilePath)
			diagramsDocFileDirAbsPath, _ := filepath.Abs(diagramsDocFileDirPath)

			errd := os.MkdirAll(diagramsDocFileDirAbsPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + diagramsDocFileDirAbsPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + diagramsDocFileDirAbsPath + " allready exists")
			}
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				diagramsDocFilePath,
				golang.DiagramsDocFile)
		}
	}

	// generate main.go if absent
	{
		// check existance of "main.go" file and generate a default "main.go" if absent
		mainFilePath := filepath.Join(*pkgPath, fmt.Sprintf("../cmd/%s/main.go", gong_models.ComputePkgNameFromPkgPath(*pkgPath)))

		_, errd := os.Stat(mainFilePath)
		if os.IsNotExist(errd) {
			log.Printf("maing.go does not exist, gongc creates a default main.go")

			mainFileDirPath := filepath.Dir(mainFilePath)
			mainFileDirAbsPath, _ := filepath.Abs(mainFileDirPath)

			errd := os.MkdirAll(mainFileDirAbsPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + mainFileDirAbsPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + mainFileDirAbsPath + " allready exists")
			}

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			golang.CodeGeneratorPackageMain(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				mainFilePath,
				*skipNg)
		}
	}

	// check existance of .vscode directory. If absent, generates default vscode configurations
	// that are usefull for development
	{
		vscodeDirFilePath := filepath.Join(*pkgPath, "../../.vscode")

		_, errd := os.Stat(vscodeDirFilePath)
		if os.IsNotExist(errd) {
			log.Printf(".vscode directory does not exist, gongc creates a default .vscode directory and the debug and launch configs")

			errd := os.MkdirAll(vscodeDirFilePath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + vscodeDirFilePath)
			}
			if os.IsExist(errd) {
				log.Fatal("directory " + vscodeDirFilePath + " should not allready exists")
			}

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(vscodeDirFilePath, "launch.json"),
				vscode.VsCodeLaunchConfig)

			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(vscodeDirFilePath, "tasks.json"),
				vscode.VsCodeTasksConfig)
		}

	}

	// generate directory for orm package
	gong_models.OrmPkgGenPath = filepath.Join(gong_models.PathToGoSubDirectory, "orm")

	os.RemoveAll(gong_models.OrmPkgGenPath)
	errd := os.MkdirAll(gong_models.OrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + gong_models.OrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")

		// supppress all files in it
	}

	// generate directory for controllers package
	gong_models.ControllersPkgGenPath = filepath.Join(gong_models.PathToGoSubDirectory, "controllers")

	os.RemoveAll(gong_models.ControllersPkgGenPath)
	errd = os.MkdirAll(gong_models.ControllersPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + gong_models.ControllersPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + gong_models.ControllersPkgGenPath + " allready exists")
	}

	// generate directory for fullstack package
	gong_models.FullstackPkgGenPath = filepath.Join(gong_models.PathToGoSubDirectory, "fullstack")

	os.RemoveAll(gong_models.FullstackPkgGenPath)
	errd = os.MkdirAll(gong_models.FullstackPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + gong_models.FullstackPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + gong_models.FullstackPkgGenPath + " allready exists")
	}

	// compute source path
	sourcePath, errd2 := filepath.Abs(*pkgPath)
	if errd2 != nil {
		log.Panic("Problem with source path " + errd2.Error())
	}

	caserEnglish := cases.Title(language.English)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../../embed.go"),
		golang.EmebedGoDirTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../fullstack/init.go"),
		golang.FullstackInitTemplate)

	golang.CodeGeneratorModelGong(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	golang.CodeGeneratorModelGongEnum(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	golang.CodeGeneratorModelGongMarshall(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	golang.CodeGeneratorModelGongGraph(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	if !*skipCoder {
		golang.CodeGeneratorModelGongCoder(
			modelPkg,
			modelPkg.Name,
			*pkgPath)
	}

	golang.GongAstGenerator(modelPkg, *pkgPath)

	// generate files
	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/setup.go"),
		golang.OrmFileSetupTemplate, golang.OrmSetupCumulSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/back_repo.go"),
		golang.BackRepoTemplateCode, golang.BackRepoSubTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../controllers/register_controllers.go"),
		golang.ControllersRegisterTemplate, golang.ControllersRegistrationsSubTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../models/gong_callbacks.go"),
		golang.ModelGongCallbacksFileTemplate, golang.ModelGongCallbacksStructSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../models/gong_serialize.go"),
		golang.ModelGongSerializeFileTemplate, golang.ModelGongSerializeStructSubTemplateCode)

	golang.MultiCodeGeneratorBackRepo(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		gong_models.OrmPkgGenPath)

	golang.MultiCodeGeneratorControllers(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		gong_models.ControllersPkgGenPath)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../docs.go"),
		golang.RootFileDocsTemplate)

	// go mod vendor to get the ng code of dependant gong stacks
	if !*skipGoModCommands {
		start := time.Now()
		cmd := exec.Command("go", "get", "github.com/xuri/excelize/v2@v2.6.1")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, fmt.Sprintf("../cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
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
		log.Printf("go get github.com/xuri/excelize/v2@v2.6.1 is over and took %s", time.Since(start))
	}

	// since go mod vendor brings angular dependencies into the vendor directory
	// the go mod vendor command has to be issued before the ng build command
	if !*skipNg {
		genAngular(modelPkg, *skipNpmInstall, *skipGoModCommands)
	}

	if !*skipFlutter {
		genFlutter(modelPkg)
	}

	apiYamlFilePath := fmt.Sprintf("%s/%sapi.yml", gong_models.ControllersPkgGenPath, modelPkg.Name)
	if !*skipSwagger {

		// generate open api specification with swagger
		cmd := exec.Command("swagger",
			"generate", "spec",
			"-w", filepath.Dir(sourcePath), // swagger is interested in the "docs.go" in the package
			// by convention, this file is located on the parent dir
			"-o", apiYamlFilePath)
		log.Printf("Running swagger command and waiting for it to finish...\n")

		// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Printf("problem with swagger : %s", err.Error())
		}
	}

	// if *skipNg {
	// 	return
	// }

	// go build
	if true {
		start := time.Now()
		var cmd *exec.Cmd
		if !*compileForDebug {
			cmd = exec.Command("go", "build")
		} else {
			// gcFlags allows for speedup of delve
			cmd = exec.Command("go", "build", "-gcflags", "-N -l")
		}
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, fmt.Sprintf("../cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
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
		log.Printf("go build over and took %s", time.Since(start))
	}

	// run application
	if *run {
		cmd := exec.Command("go", "run", "main.go")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath,
			fmt.Sprintf("../../go/cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
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
	}
}
