package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/go/angular"
	"github.com/fullstack-lang/gong/go/golang"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/tidwall/sjson"
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
			{
				{
					start := time.Now()
					//  --skip-confirmation is necessary when executing angular without user interaction
					// otherwise, one meet the error
					// "No terminal detected. '--skip-confirmation' can be used to bypass installation confirmation.
					// Ensure package name is correct prior to '--skip-confirmation' option usage."
					cmd := exec.Command("ng", "add", "@angular/material@15", "--skip-confirmation")
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
						"angular-split@14.1",
						"material-icons",
						"@fontsource/open-sans",
						"@fontsource/roboto",
						"@angular-material-components/datetime-picker@9")
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
						"@types/backbone", "@types/jquery", "@types/lodash", "@types/node", "@types/leaflet", "codelyzer", "install", "jointjs")
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
					angular.NgFileModule)

				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/index.html"),
					angular.NgFileIndex)

				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.component.ts"),
					angular.NgFileAppComponentTs)

				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					modelPkg.Name,
					modelPkg.PkgPath,
					filepath.Join(gong_models.NgWorkspacePath, "src/app/app.component.html"),
					angular.NgFileAppComponentHtml)

			}
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

					gong_models.InsertStringToFile(filename, angular.TsConfigInsertForPaths, "\"paths\": {")

				}
				{
					// patch styles.css file in order have imports of css stuff and work offline
					filename := filepath.Join(gong_models.NgWorkspacePath, "src", "styles.css")
					gong_models.InsertStringToFile(filename, angular.StylesCssInsert, "/* You can add global styles to this file, and also import other style files */")
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

				filename := filepath.Join(gong_models.NgWorkspacePath, "tsconfig.json")
				gong_models.InsertStringToFile(filename, "        \"projects/"+modelPkg.Name+"specific/src/public-api.ts\",", modelPkg.Name+"specific\": [")

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
			if true {

				// needed till fix of https://github.com/clientIO/joint/issues/2018
				tsConfigFile := filepath.Join(gong_models.NgWorkspacePath, "tsconfig.json")
				// Read the contents of the tsconfig.json file
				file, err := ioutil.ReadFile(tsConfigFile)
				if err != nil {
					log.Fatalln(err)
					return
				}

				// Modify the skipLibCheck field
				file, err = sjson.SetBytes(file, "compilerOptions.skipLibCheck", true)
				if err != nil {
					log.Fatalln(err)
					return
				}

				// Write the updated JSON back to the tsconfig.json file
				err = ioutil.WriteFile(tsConfigFile, file, 0644)
				if err != nil {
					log.Fatalln(err)
					return
				}
			}
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

	angular.MultiCodeGeneratorNgDetail(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgClass(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgService(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgCommitNbFromBack(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgGongselectionServiceTs(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgNullInt64(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	angular.CodeGeneratorNgPushFromFrontNb(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath,
		*addr)

	angular.MultiCodeGeneratorNgTable(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.MultiCodeGeneratorNgSorting(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgFrontRepo(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgSidebar(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgEnum(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgPublicApi(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	angular.CodeGeneratorNgSplitter(
		modelPkg,
		modelPkg.Name,
		gong_models.MatTargetPath,
		modelPkg.PkgPath)

	caserEnglish := cases.Title(language.English)
	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, modelPkg.Name+".module.ts"),
		angular.NgLibModuleTemplate, angular.NgLibModuleSubTemplateCode)

	{
		directory, _ := filepath.Abs(
			filepath.Join(gong_models.MaterialLibSpecificTargetPath, "data-model-panel"))
		_, err := os.Stat(directory)
		if os.IsNotExist(err) {
			err := os.MkdirAll(directory, os.ModePerm)

			if err != nil {
				log.Fatalln("Unable to create dir", directory, err.Error())
			}

			log.Println("creating directory : " + gong_models.OrmPkgGenPath)

			// if data-model-panel does not exists, it is an indication that
			// the module file has to be overwritten
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(gong_models.MaterialLibSpecificTargetPath,
					fmt.Sprintf("%sspecific.module.ts", modelPkg.Name)),
				angular.NgFileModuleSpecific)

			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(gong_models.MaterialLibSpecificTargetPath,
					"data-model-panel", "data-model-panel.component.ts"),
				angular.NgFileDataModelPanelTemplateTs)

			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				filepath.Join(gong_models.MaterialLibSpecificTargetPath,
					"data-model-panel", "data-model-panel.component.html"),
				angular.NgFileDataModelPanelTemplateHtml)
		}
		if os.IsExist(err) {
			log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")
		}
	}

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, "app-routing.module.ts"),
		angular.NgRoutingTemplate, angular.NgRoutingSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, "route-service.ts"),
		angular.NgRouteServiceTemplate, angular.NgRouteServiceSubTemplateCode)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath, filepath.Join(gong_models.MatTargetPath, "map-components.ts"),
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
