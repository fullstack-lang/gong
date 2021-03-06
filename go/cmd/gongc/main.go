package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const COMPUTED_FROM_PKG_PATH string = "computed from pkgPath (path to package for analysis)"

var (
	pkgPath     = flag.String("pkgPath", ".", "path to the models package to be compiled by gongc compilation")
	skipSwagger = flag.Bool("skipSwagger", true, "skip swagger")
	backendOnly = flag.Bool("backendOnly", false, "generates backendOnly")
	addr        = flag.String("addr", "localhost:8080/api",
		"network address addr where the angular generated service will lookup the server")
	run               = flag.Bool("run", false, "run 'go run main.go' after compilation")
	skipGoModCommands = flag.Bool("skipGoModCommands", false, "avoid calls to go mod init, tidy and vendor")

	useParser = flag.Bool("useParser", true, "use go/parser.Parse instead of packages.Load (which requires go installled)")
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

	// check existance of go.mod file in the path to the 'models' package
	//
	// if no go.mod file is found above the 'models' package, gongc fails
	//
	// if go.mod exists, it means the package path has been defined
	// for instance "github.com/fullstack-lang/gongsvg"
	//
	// if go.mod does not exist, gongc can only infer the package name
	// from the name of directory that is two levels above "go/models"
	// it is up to the developper to change the module name after the first gong generation
	//
	//
	pkgName, fullPkgPath := gong_models.ComputePkgPathFromGoModFile(*pkgPath)

	// initiate model package
	modelPkg := (&gong_models.ModelPkg{
		Name:    pkgName,
		PkgPath: fullPkgPath,
	})

	gong_models.Walk(*pkgPath, modelPkg, *useParser)

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
	{

		directory, err :=
			filepath.Abs(
				filepath.Join(*pkgPath,
					fmt.Sprintf("../../ng/projects/%s/src/lib", modelPkg.Name)))
		gong_models.MatTargetPath = directory
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
		}
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
				gong_models.DiagramsDocFile)
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
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				mainFilePath,
				gong_models.PackageMain)
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
				gong_models.VsCodeLaunchConfig)

			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(vscodeDirFilePath, "tasks.json"),
				gong_models.VsCodeTasksConfig)
		}

	}

	// generate things in ng  lib directory
	var ngNewWsPerformed bool
	{

		directory, err := filepath.Abs(
			filepath.Join(*pkgPath, "../../ng"))
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
		}
		gong_models.NgWorkspacePath = directory
		log.Println("module target abs path " + gong_models.NgWorkspacePath)

		// check existance of angular directory. If absent, use "ng new ng", then generate default app.component.html
		_, errd := os.Stat(gong_models.NgWorkspacePath)
		if os.IsNotExist(errd) {
			log.Printf("ng directory %s does not exist, hence gong is generating it with ng new ng command", gong_models.NgWorkspacePath)

			// generate ng workspace

			cmd := exec.Command("ng", "new", "ng", "--defaults=true", "--minimal=true")
			cmd.Dir = filepath.Dir(gong_models.NgWorkspacePath)
			log.Printf("Creating angular workspace\n")

			// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			cmd.Stdout = mw
			cmd.Stderr = mw

			log.Println(cmd.String())
			log.Println(stdBuffer.String())

			// Execute the command
			start := time.Now()
			if err := cmd.Run(); err != nil {
				log.Panic(err)
			}
			log.Printf("ng new ng is over and took %s", time.Since(start))
			ngNewWsPerformed = true

			// add the necessary libraries to gong applications
			{
				{
					start := time.Now()
					//  --skip-confirmation is necessary when executing angular without user interaction
					// otherwise, one meet the error
					// "No terminal detected. '--skip-confirmation' can be used to bypass installation confirmation.
					// Ensure package name is correct prior to '--skip-confirmation' option usage."
					cmd := exec.Command("ng", "add", "@angular/material", "--skip-confirmation")
					cmd.Dir = gong_models.NgWorkspacePath
					log.Printf("Adding angular material\n")

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
					log.Printf("ng add is over and took %s", time.Since(start))
				}

				{
					start := time.Now()
					cmd := exec.Command("npm", "install", "--save",
						"angular-split@13.2.0", "material-design-icons", "typeface-open-sans", "typeface-roboto", "@angular-material-components/datetime-picker@7")
					cmd.Dir = gong_models.NgWorkspacePath
					log.Printf("Installing some packages\n")

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
					log.Printf("npm install is over and took %s", time.Since(start))
				}
				{
					start := time.Now()
					cmd := exec.Command("npm", "install", "--save",
						"@types/backbone", "@types/jointjs", "@types/jquery", "@types/lodash", "@types/node", "@types/leaflet", "backbone", "codelyzer", "install", "jointjs", "jquery", "lodash")
					cmd.Dir = gong_models.NgWorkspacePath
					log.Printf("Installing some packages\n")

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
					log.Printf("npm install is over and took %s", time.Since(start))
				}
			}

			// generate default app.component.ts, app.component.html and app.module.ts
			{
				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.module.ts"),
					gong_models.NgFileModule)

				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.component.ts"),
					gong_models.NgFileAppComponentTs)

				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.component.html"),
					gong_models.NgFileAppComponentHtml)

			}
		}
	}

	// if ng new was performed, modify the angular json file
	// for some configuration stuff
	// mostly
	if ngNewWsPerformed {
		angularJsonFile := filepath.Join(gong_models.NgWorkspacePath, "angular.json")

		// Read the file
		input, err := os.ReadFile(angularJsonFile)
		if err != nil {
			log.Fatalln(err)
		}

		var root interface{}
		if err := json.Unmarshal(input, &root); err != nil {
			log.Fatal(err)
		}

		// Walk down path to target object.
		v := root
		var path = []string{"projects",
			"ng",
			"architect",
			"build",
			"configurations",
			"production",
			"budgets"}
		for i, k := range path {
			m, ok := v.(map[string]interface{})
			if !ok {
				log.Fatalf("map not found at %s", strings.Join(path[:i+1], ", "))
			}
			v, ok = m[k]
			if !ok {
				log.Fatalf("value not found at %s", strings.Join(path[:i+1], ", "))
			}
		}

		// Set value in the target object.
		// budgets is a slice and it is the first one
		slice, ok := v.([]interface{})
		if !ok {
			log.Fatalf("not slice found at %s", strings.Join(path, ", "))
		}
		v = slice[0]
		m, ok := v.(map[string]interface{})
		if !ok {
			log.Fatalf("map not found at %s", strings.Join(path, ", "))
		}
		m["maximumError"] = "10mb"

		// Marshal back to JSON. Variable d is []byte with the JSON
		d, err := json.Marshal(root)
		if err != nil {
			log.Fatal(err)
		}

		errWrite := os.WriteFile(angularJsonFile, d, os.ModePerm)
		if errWrite != nil {
			log.Fatal(errWrite)
		}

		// check existance of generated angular library. If absent, use "ng generate libray <library>"
		// and generate default app application
		{
			_, errStat := os.Stat(gong_models.MatTargetPath)
			log.Println("module target abs path " + gong_models.MatTargetPath)

			if os.IsNotExist(errStat) {
				log.Printf("library directory %s does not exist, hence gong is generating it with ng generate library command",
					gong_models.MatTargetPath)

				// generate library project
				start := time.Now()
				cmd := exec.Command("ng", "generate", "library", modelPkg.Name, "--defaults=true", "--skip-install=true")
				cmd.Dir = gong_models.NgWorkspacePath
				log.Printf("Creating a library %s in the angular workspace\n", modelPkg.Name)

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
				log.Printf("ng generate library is over and took %s", time.Since(start))
				{
					// patch tsconfig file in order to have the path to the public-api of the
					// generated library (instead of the path to "dist")
					filename := filepath.Join(gong_models.NgWorkspacePath, "tsconfig.json")
					gong_models.InsertStringToFile(filename, "        \"projects/"+modelPkg.Name+"/src/public-api.ts\",", modelPkg.Name+"\": [")

					gong_models.InsertStringToFile(filename, gong_models.TsConfigInsertForPaths, "\"paths\": {")

				}
				{
					// patch styles.css file in order have imports of css stuff and work offline
					filename := filepath.Join(gong_models.NgWorkspacePath, "src", "styles.css")
					gong_models.InsertStringToFile(filename, gong_models.StylesCssInsert, "/* You can add global styles to this file, and also import other style files */")
				}
			}

			// generate the "specific" library that will contains stack specific stuff, that can be reused
			if os.IsNotExist(errStat) {
				log.Printf("library directory %sspecific does not exist, hence gong is generating it with ng generate library command",
					gong_models.MatTargetPath)

				// generate library project
				start := time.Now()
				cmd := exec.Command("ng", "generate", "library", modelPkg.Name+"specific", "--defaults=true", "--skip-install=true")
				cmd.Dir = gong_models.NgWorkspacePath
				log.Printf("Creating a library %s in the angular workspace\n", modelPkg.Name)

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
				log.Printf("ng generate library is over and took %s", time.Since(start))
			}

			// npm install
			if true {
				start := time.Now()
				cmd := exec.Command("npm", "i")
				cmd.Dir = gong_models.NgWorkspacePath
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
				log.Printf("npm i is over and took %s", time.Since(start))
			}
		}
	}

	// now replace the generated content in the library
	{
		if !*backendOnly {
			log.Println("Removing all content of " + gong_models.MatTargetPath)
			gong_models.RemoveContents(gong_models.MatTargetPath)
		}
	}

	// generates styles
	{
		directory, _ := filepath.Abs(gong_models.MatTargetPath)
		errForCreationOfStylesDir := os.MkdirAll(filepath.Join(directory, "styles"), os.ModePerm)
		if os.IsNotExist(errForCreationOfStylesDir) {
			log.Println("creating directory : " + gong_models.OrmPkgGenPath)
		}
		if os.IsExist(errForCreationOfStylesDir) {
			log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")
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

	// compute source path
	sourcePath, errd2 := filepath.Abs(*pkgPath)
	if errd2 != nil {
		log.Panic("Problem with source path " + errd2.Error())
	}

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		strings.Title(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgWorkspacePath, "embed.go"),
		gong_models.GoProjectsGo)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		strings.Title(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgWorkspacePath, "../embed.go"),
		gong_models.EmebedNgDistNg)

	// remove "gong.go" file
	gong_models.RemoveGeneratedGongFile(*pkgPath)
	gong_models.CodeGeneratorModelGong(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	// generate files
	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/setup.go"),
		gong_models.OrmFileSetupTemplate, gong_models.OrmSetupCumulSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/back_repo.go"),
		gong_models.BackRepoTemplateCode, gong_models.BackRepoSubTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../controllers/register_controllers.go"),
		gong_models.ControllersRegisterTemplate, gong_models.ControllersRegistrationsSubTemplate)

	gong_models.MultiCodeGeneratorBackRepo(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		gong_models.OrmPkgGenPath)

	gong_models.MultiCodeGeneratorControllers(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		gong_models.ControllersPkgGenPath)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../docs.go"),
		gong_models.RootFileDocsTemplate)

	// go mod vendor to get the ng code of dependant gong stacks
	if !*skipGoModCommands {
		start := time.Now()
		cmd := exec.Command("go", "mod", "tidy")
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
		log.Printf("go mod tidy is over and took %s", time.Since(start))
	}

	// go mod vendor to get the ng code of dependant gong stacks
	if !*skipGoModCommands {
		start := time.Now()
		cmd := exec.Command("go", "mod", "vendor")
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
		log.Printf("go mod vendor is over and took %s", time.Since(start))
	}

	gong_models.MultiCodeGeneratorNgDetail(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.MultiCodeGeneratorNgPresentation(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.MultiCodeGeneratorNgClass(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.MultiCodeGeneratorNgService(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	gong_models.CodeGeneratorNgCommitNb(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	gong_models.CodeGeneratorNgGongselectionServiceTs(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	gong_models.CodeGeneratorNgNullInt64(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	gong_models.CodeGeneratorNgPushFromFrontNb(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	gong_models.MultiCodeGeneratorNgTable(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.MultiCodeGeneratorNgSorting(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.CodeGeneratorNgFrontRepo(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.CodeGeneratorNgSidebar(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.CodeGeneratorNgEnum(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.CodeGeneratorNgPublicApi(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.CodeGeneratorNgSplitter(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		strings.Title(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, modelPkg.Name+".module.ts"),
		gong_models.NgLibModuleTemplate, gong_models.NgLibModuleSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		strings.Title(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, "app-routing.module.ts"),
		gong_models.NgRoutingTemplate, gong_models.NgRoutingSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, "map-components.ts"),
		gong_models.NgLibMapComponentsServiceTemplate, gong_models.NgLibMapComponentsSubTemplateCode)

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

	if *backendOnly {
		return
	}

	// ng build
	if true {
		start := time.Now()
		cmd := exec.Command("ng", "build")
		cmd.Dir = gong_models.NgWorkspacePath
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
		log.Printf("ng build is over and took %s", time.Since(start))
	}

	// go get
	if !*skipGoModCommands {
		start := time.Now()
		cmd := exec.Command("go", "get")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, fmt.Sprintf("../../go/cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
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
		log.Printf("go get over and took %s", time.Since(start))
	}

	// go get isatty
	if !*skipGoModCommands {
		start := time.Now()
		// path gin since isatty fails if v0.0.12 (patch version 0.0.14 is OK)
		cmd := exec.Command("go", "get", "-d", "github.com/mattn/go-isatty")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, "../.."))
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
		log.Printf("go get -d github.com/mattn/go-isatty over and took %s", time.Since(start))
	}

	// go build
	if true {
		start := time.Now()
		cmd := exec.Command("go", "build")
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
		cmd.Dir, _ = filepath.Abs(filepath.Join(gong_models.NgWorkspacePath, fmt.Sprintf("../go/cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
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
