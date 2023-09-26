package main

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

	"github.com/fullstack-lang/gong/go/angular"
	gong_models "github.com/fullstack-lang/gong/go/models"
)

func configGeneratedNgWorkspace(modelPkg *gong_models.ModelPkg) {

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
		_, errStat := os.Stat(gong_models.NgDataLibrarySourceCodeDirectory)
		log.Println("module target abs path " + gong_models.NgDataLibrarySourceCodeDirectory)

		if os.IsNotExist(errStat) {
			log.Printf("library directory %s does not exist, hence gong is generating it with ng generate library command",
				gong_models.NgDataLibrarySourceCodeDirectory)

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
				gong_models.NgDataLibrarySourceCodeDirectory)

			// generate library project
			start := time.Now()
			cmd := exec.Command("ng", "generate", "library", modelPkg.Name+"specific", "--defaults=true", "--skip-install=true")
			cmd.Dir = gong_models.NgWorkspacePath
			log.Printf("Creating a specific library %s in the angular workspace\n", modelPkg.Name)

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

			{

			}
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
		// if true {

		// 	// needed till fix of https://github.com/clientIO/joint/issues/2018
		// 	tsConfigFile := filepath.Join(gong_models.NgWorkspacePath, "tsconfig.json")
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
