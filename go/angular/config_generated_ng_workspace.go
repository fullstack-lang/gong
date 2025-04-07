package angular

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func configGeneratedNgWorkspace(modelPkg *gong_models.ModelPkg) {

	if true {
		angularJsonFile := filepath.Join(modelPkg.NgWorkspacePath, "angular.json")

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
			modelPkg.NgWorkspaceName,
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
	}

	// check existance of generated angular library. If absent, use "ng generate libray <library>"
	// and generate default app application
	{
		_, errStat := os.Stat(modelPkg.NgDataLibrarySourceCodeDirectory)
		log.Println("module target abs path " + modelPkg.NgDataLibrarySourceCodeDirectory)

		if os.IsNotExist(errStat) {
			log.Printf("library directory %s does not exist, hence gong is generating it with ng generate library command",
				modelPkg.NgDataLibrarySourceCodeDirectory)

			// generate library project
			start := time.Now()
			cmd := exec.Command("ng", "generate", "library", modelPkg.Name, "--defaults=true", "--skip-install=true")
			cmd.Dir = modelPkg.NgWorkspacePath
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
				// patch tsconfig file in order to have the @vendored_components
				filename := filepath.Join(modelPkg.NgWorkspacePath, "tsconfig.json")

				gong_models.InsertStringToFile(filename, TsConfigInsertForPaths, "\"paths\": {")
			}
			{
				// patch styles.css file in order have imports of css stuff and work offline
				filename := filepath.Join(modelPkg.NgWorkspacePath, "src", "styles.css")
				gong_models.InsertStringToFile(filename, StylesCssInsert, "/* You can add global styles to this file, and also import other style files */")
			}
		}

		// generate the "specific" library that will contains stack specific stuff, that can be reused
		if os.IsNotExist(errStat) {
			log.Printf("library directory %sspecific does not exist, hence gong is generating it with ng generate library command",
				modelPkg.NgDataLibrarySourceCodeDirectory)

			// generate library project
			start := time.Now()
			cmd := exec.Command("ng", "generate", "library", modelPkg.Name+"specific", "--defaults=true", "--skip-install=true")
			cmd.Dir = modelPkg.NgWorkspacePath
			log.Printf("Creating a specific library %s in the angular workspace\n", modelPkg.Name+"specific")

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

			// generate c
			// ng generate component my-component --inlineTemplate=false --inlineStyle=false
			{
				componentName := modelPkg.Name + "-" + "specific"
				start := time.Now()
				cmd := exec.Command("ng", "generate", "component", componentName,
					"--inline-template=false", "--inline-style=false", "--project="+modelPkg.Name+"specific")
				cmd.Dir = modelPkg.NgWorkspacePath
				log.Printf("Creating a component %s in the angular workspace\n", componentName)

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
		}

		// npm install
		if true {
			start := time.Now()
			cmd := exec.Command("npm", "i")
			cmd.Dir = modelPkg.NgWorkspacePath
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
		// if true {

		// 	// needed till fix of https://github.com/clientIO/joint/issues/2018
		// 	tsConfigFile := filepath.Join(modelPkg.NgWorkspacePath, "tsconfig.json")
		// 	// Read the contents of the tsconfig.json file
		// 	file, err := ioutil.ReadFile(tsConfigFile)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 		return
		// 	}

		// 	// Modify the skipLibCheck field
		// 	file, err = sjson.SetBytes(file, "compilerOptions.skipLibCheck", true)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 		return
		// 	}

		// 	// Write the updated JSON back to the tsconfig.json file
		// 	err = ioutil.WriteFile(tsConfigFile, file, 0644)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 		return
		// 	}
		// }
	}

}
