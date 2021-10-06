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

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const COMPUTED_FROM_PKG_PATH string = "computed from pkgPath (path to package for analysis)"

var (
	pkgPath = flag.String("pkgPath", ".", "path to package for gongc compilation")

	backendTargetPath = flag.String("backendTargetPath", COMPUTED_FROM_PKG_PATH,
		"relative path to the directory where orm & controllers packages are generated"+
			" (by convention, one level above path to package for analysis)."+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	matTargetPath = flag.String("matTargetPath", COMPUTED_FROM_PKG_PATH, //
		"path to the ng directory where material components are generated"+
			"(by convention, relative to pkgPath, "+
			"into ../../ng/projects/<pkgName>/src/lib)"+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	ngWorkspacePath = flag.String("ngWorkspacePath", COMPUTED_FROM_PKG_PATH, //
		"path to the ng workspace directory (for performing npm isntall commands"+
			"(by convention, ../../ng relative to path to package for analysis)."+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	skipSwagger = flag.Bool("skipSwagger", false, "skip swagger")

	backendOnly = flag.Bool("backendOnly", false, "generates backendOnly")

	addr = flag.String("addr", "localhost:8080/api",
		"network address addr where the angular generated service will lookup the server")

	run = flag.Bool("run", false, "run 'go run main.go' after compilation")
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

	gong_models.ADDR = *addr

	// parse package and generate code if flag set
	gong_models.RemoveGoAllGongStruct(*pkgPath)

	// load package into database
	var modelPkg gong_models.ModelPkg
	gong_models.Walk(*pkgPath, &modelPkg)

	// check wether the package name follows gong naming convention
	if strings.ContainsAny(modelPkg.Name, "-") {
		log.Panicln(modelPkg.Name + " is not OK for a gong package name because it contains a - (dash) " +
			"and it cannot be used for naming a typescript class (the generated module in this cause)")
	}

	// compute path to generation directory
	{
		if *backendTargetPath == COMPUTED_FROM_PKG_PATH {
			*backendTargetPath = filepath.Join(*pkgPath, "..")
		}

		directory, err := filepath.Abs(*backendTargetPath)
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
		gong_models.BackendTargetPath = directory
		log.Println("backend target path " + gong_models.BackendTargetPath)
	}

	{
		// check existance of "main.go" file and generate a default "main.go" if absent
		mainFilePath := filepath.Join(*pkgPath, "../../main.go")

		_, errd := os.Stat(mainFilePath)
		if os.IsNotExist(errd) {
			log.Printf("maing.go does not exist, gongc creates a default main.go")

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			gong_models.VerySimpleCodeGenerator(
				&modelPkg,
				gong_models.PkgName,
				gong_models.PkgGoPath,
				mainFilePath,
				gong_models.PackageMain)
		}
	}
	{
		if *ngWorkspacePath == COMPUTED_FROM_PKG_PATH {
			*ngWorkspacePath = filepath.Join(*pkgPath, "../../ng")
		}

		directory, err := filepath.Abs(*ngWorkspacePath)
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
		}
		gong_models.NgWorkspacePath = directory
		log.Println("module target abs path " + gong_models.NgWorkspacePath)
	}

	// check existance of angular directory
	{

		_, errd := os.Stat(gong_models.NgWorkspacePath)
		if os.IsNotExist(errd) {
			log.Printf("ng directory %s does not exist", gong_models.NgWorkspacePath)

			// generate ng workspace
			{
				cmd := exec.Command("ng", "new", "ng", "--defaults=true")
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
				if err := cmd.Run(); err != nil {
					log.Panic(err)
				}
			}
			// generate library project
			{
				cmd := exec.Command("ng", "generate", "library", gong_models.PkgName, "--defaults=true")
				cmd.Dir = gong_models.NgWorkspacePath
				log.Printf("Creating a library %s in the angular workspace\n", gong_models.PkgName)

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
			{
				cmd := exec.Command("ng", "add", "@angular/material", "--defaults=true", "--skip-confirmation")
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
			}
			{
				cmd := exec.Command("npm", "install", "--save",
					"angular-split", "material-design-icons", "typeface-open-sans", "typeface-roboto")
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
			}
			// generate default app.component.ts, app.component.html and app.module.ts
			{

				gong_models.VerySimpleCodeGenerator(
					&modelPkg,
					gong_models.PkgName,
					gong_models.PkgGoPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.module.ts"),
					gong_models.NgFileModule)
				filename := filepath.Join(gong_models.NgWorkspacePath, "src/app/app.component.html")

				// we should use go generate
				log.Println("generating app component file: " + filename)

				f, err := os.Create(filename)
				if err != nil {
					log.Panic(err)
				}
				defer f.Close()
				res := fmt.Sprintf("<app-%s-splitter></app-%s-splitter>", gong_models.PkgName, gong_models.PkgName)

				fmt.Fprintf(f, "%s", res)
			}

			{
				// patch tsconfig file in order to have the path to the public-api of the
				// generated library (instead of the path to "dist")
				filename := filepath.Join(gong_models.NgWorkspacePath, "tsconfig.json")
				InsertStringToFile(filename, "        \"projects/"+modelPkg.Name+"/src/public-api.ts\",", modelPkg.Name+"\": [")
			}

		} else {
			log.Printf("ng directory %s does exist", gong_models.NgWorkspacePath)
		}

	}
	{
		if *matTargetPath == COMPUTED_FROM_PKG_PATH {
			*matTargetPath = filepath.Join(*pkgPath, fmt.Sprintf("../../ng/projects/%s/src/lib", gong_models.PkgName))
		}

		directory, err := filepath.Abs(*matTargetPath)
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
		}
		// check existance of path
		fileInfo, err := os.Stat(directory)
		if os.IsNotExist(err) {
			log.Panicf("Folder %s does not exist.", directory)
		}
		if fileInfo.Mode().Perm()&(1<<(uint(7))) == 0 {
			log.Panicf("Folder %s is not writtable", directory)
		}
		if !*backendOnly {
			RemoveContents(*matTargetPath)
		}
		gong_models.MatTargetPath = directory
		log.Println("module target abs path " + gong_models.MatTargetPath)
	}

	{
		directory, errForCreationOfStylesDir := filepath.Abs(*matTargetPath)
		errForCreationOfStylesDir = os.Mkdir(filepath.Join(directory, "styles"), os.ModePerm)
		if os.IsNotExist(errForCreationOfStylesDir) {
			log.Println("creating directory : " + gong_models.OrmPkgGenPath)
		}
		if os.IsExist(errForCreationOfStylesDir) {
			log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")
		}
	}

	// generate directory for orm package
	gong_models.OrmPkgGenPath = filepath.Join(gong_models.BackendTargetPath, "orm")

	os.RemoveAll(gong_models.OrmPkgGenPath)
	errd := os.Mkdir(gong_models.OrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + gong_models.OrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")

		// supppress all files in it
	}

	// generate directory for controllers package
	gong_models.ControllersPkgGenPath = filepath.Join(gong_models.BackendTargetPath, "controllers")

	os.RemoveAll(gong_models.ControllersPkgGenPath)
	errd = os.Mkdir(gong_models.ControllersPkgGenPath, os.ModePerm)
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

	gong_models.MultiCodeGeneratorNgDetail(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.MultiCodeGeneratorNgPresentation(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.MultiCodeGeneratorNgClass(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.MultiCodeGeneratorNgService(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath,
		*addr)

	gong_models.CodeGeneratorNgCommitNb(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath,
		*addr)

	gong_models.CodeGeneratorNgPushFromFrontNb(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath,
		*addr)

	gong_models.MultiCodeGeneratorNgTable(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.MultiCodeGeneratorNgSorting(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.CodeGeneratorModelGong(
		&modelPkg,
		gong_models.PkgName,
		*pkgPath)

	// generate files
	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		filepath.Join(*pkgPath, "../orm/setup.go"),
		gong_models.OrmFileSetupTemplate, gong_models.OrmSetupCumulSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		filepath.Join(*pkgPath, "../orm/back_repo.go"),
		gong_models.BackRepoTemplateCode, gong_models.BackRepoSubTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		filepath.Join(*pkgPath, "../controllers/register_controllers.go"),
		gong_models.ControllersRegisterTemplate, gong_models.ControllersRegistrationsSubTemplate)

	gong_models.MultiCodeGeneratorBackRepo(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		gong_models.OrmPkgGenPath)

	gong_models.MultiCodeGeneratorControllers(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		gong_models.ControllersPkgGenPath)

	gong_models.VerySimpleCodeGenerator(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath,
		filepath.Join(*pkgPath, "../docs.go"),
		gong_models.RootFileDocsTemplate)

	gong_models.CodeGeneratorNgFrontRepo(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.CodeGeneratorNgSidebar(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.CodeGeneratorNgEnum(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.CodeGeneratorNgPublicApi(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.CodeGeneratorNgSplitter(
		&modelPkg,
		gong_models.PkgName,
		gong_models.MatTargetPath,
		gong_models.PkgGoPath)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		strings.Title(gong_models.PkgName),
		gong_models.PkgGoPath, filepath.Join(gong_models.MatTargetPath, gong_models.PkgName+".module.ts"),
		gong_models.NgLibModuleTemplate, gong_models.NgLibModuleSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		strings.Title(gong_models.PkgName),
		gong_models.PkgGoPath, filepath.Join(gong_models.MatTargetPath, "app-routing.module.ts"),
		gong_models.NgRoutingTemplate, gong_models.NgRoutingSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		&modelPkg,
		gong_models.PkgName,
		gong_models.PkgGoPath, filepath.Join(gong_models.MatTargetPath, "map-components.ts"),
		gong_models.NgLibMapComponentsServiceTemplate, gong_models.NgLibMapComponentsSubTemplateCode)

	apiYamlFilePath := fmt.Sprintf("%s/%sapi.yml", gong_models.ControllersPkgGenPath, gong_models.PkgName)
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

	if true {

		cmd := exec.Command("npm", "i")
		cmd.Dir, _ = filepath.Abs(*ngWorkspacePath)
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

	if false { // this step is unnecessary since compilation is done with tsconfig dependencies
		cmd := exec.Command("ng", "build", gong_models.PkgName)
		cmd.Dir, _ = filepath.Abs(*ngWorkspacePath)
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

	if true {

		cmd := exec.Command("ng", "build")
		cmd.Dir, _ = filepath.Abs(*ngWorkspacePath)
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

	log.Printf("compilation over")

	if *run {

		cmd := exec.Command("go", "run", "main.go")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*ngWorkspacePath, ".."))
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

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
