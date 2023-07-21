package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fullstack-lang/gong/go/angular"
	"github.com/fullstack-lang/gong/go/golang"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func genAngular(modelPkg *gong_models.ModelPkg, skipNpmInstall bool, skipGoModCommands bool) {

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
			var cmd *exec.Cmd
			if !skipNpmInstall {
				cmd = exec.Command("ng", "new", "ng", "--defaults=true", "--minimal=true")
			} else {
				cmd = exec.Command("ng", "new", "ng", "--defaults=true", "--minimal=true", "--skip-install")
			}
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
			installMaterialLibs(modelPkg)
		} else {
			// if ng is already present, one might have to perform an installation
			// for example, if the repo has been checkout and "node_modules" is empty
			if !skipNpmInstall {
				start := time.Now()
				cmd := exec.Command("npm", "install")
				cmd.Dir = gong_models.NgWorkspacePath
				log.Printf("Performing default npm install\n")

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
	}

	// if ng new was performed, modify the angular json file
	// for some configuration stuff
	// mostly
	if ngNewWsPerformed {
		configGeneratedNgWorkspace(modelPkg)
	}

	// if data-model-panel does not exists, it is an indication that
	// the module file has to be overwritten
	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			fmt.Sprintf("%sdatamodel.module.ts", modelPkg.Name)),
		angular.NgFileModuleDatamodel)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			"data-model-panel", "data-model-panel.component.ts"),
		angular.NgFileDataModelPanelTemplateTs)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			"data-model-panel", "data-model-panel.component.html"),
		angular.NgFileDataModelPanelTemplateHtml)

	// generates styles
	{
		directory, _ := filepath.Abs(gong_models.NgDataLibrarySourceCodeDirectory)
		errForCreationOfStylesDir := os.MkdirAll(filepath.Join(directory, "styles"), os.ModePerm)
		if os.IsNotExist(errForCreationOfStylesDir) {
			log.Println("creating directory : " + gong_models.OrmPkgGenPath)
		}
		if os.IsExist(errForCreationOfStylesDir) {
			log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")
		}
	}

	angular.MultiCodeGeneratorNgDetail(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgClass(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgService(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgCommitNbFromBack(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgGongselectionServiceTs(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgNullInt64(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgPushFromFrontNb(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath,
		*addr)

	angular.MultiCodeGeneratorNgTable(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgSorting(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgFrontRepo(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgSidebar(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgEnum(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgPublicApi(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgSplitter(
		modelPkg,
		modelPkg.Name,
		gong_models.NgDataLibrarySourceCodeDirectory,
		modelPkg.PkgPath)

	caserEnglish := cases.Title(language.English)
	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgDataLibrarySourceCodeDirectory, modelPkg.Name+".module.ts"),
		angular.NgLibModuleTemplate, angular.NgLibModuleSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgDataLibrarySourceCodeDirectory, "app-routing.module.ts"),
		angular.NgRoutingTemplate, angular.NgRoutingSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgDataLibrarySourceCodeDirectory, "route-service.ts"),
		angular.NgRouteServiceTemplate, angular.NgRouteServiceSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath, filepath.Join(gong_models.NgDataLibrarySourceCodeDirectory, "map-components.ts"),
		angular.NgLibMapComponentsServiceTemplate, angular.NgLibMapComponentsSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgWorkspacePath, "projects/embed.go"),
		golang.GoProjectsGo)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.NgWorkspacePath, "../embed_ng_dist_ng.go"),
		angular.EmebedNgDistNg)

	// go mod tidy to get the new dependencies
	if !skipGoModCommands {
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
	// it has to be done after the go mod tidy and ust before the ng build
	if !skipGoModCommands {
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

	// ng build
	{
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
}
