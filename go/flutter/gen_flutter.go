package flutter

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func GenFlutter(modelPkg *gong_models.ModelPkg, pkgPath string) {

	// check if "flutter" directory exist
	// if not, call "flutter create flutter" to create a flutter project
	var flutterProjectDir string
	var flutterCreateHasBeenPerformed bool
	{

		var err error
		flutterProjectDir, err = filepath.Abs(filepath.Join(pkgPath, "../../flutter_project"))
		if err != nil {
			log.Panic("Problem with flutter path " + err.Error())
		}

		_, errd := os.Stat(flutterProjectDir)
		if os.IsNotExist(errd) {
			log.Printf("flutter directory %s does not exist, hence flutter create flutter_project is called", flutterProjectDir)

			cmd := exec.Command("flutter", "create", "flutter_project", "--template=app", "--platforms=web")
			cmd.Dir = filepath.Dir(flutterProjectDir)
			log.Printf("Creating flutter project flutter_project\n")

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
			log.Printf("%s is over and took %s", cmd.String(), time.Since(start))

			flutterCreateHasBeenPerformed = true
		}
	}

	// for the first generation, overides the main.dart code
	if flutterCreateHasBeenPerformed {
		pathToMain := filepath.Join(flutterProjectDir, "lib", "main.dart")
		GenerateDefaultMainFile(modelPkg, pathToMain)

		{
			// delete the test directory (temporary)
			testDirectory := filepath.Join(flutterProjectDir, "test")
			errd := os.RemoveAll(testDirectory)
			if errd != nil {
				log.Panic("Problem with flutter path " + errd.Error())
			}
		}

		{
			// delete the iml file  (for intelli J)
			inteliJfile := filepath.Join(flutterProjectDir, "flutter_project.iml")
			errd := os.Remove(inteliJfile)
			if errd != nil {
				log.Panic("Problem with flutter path " + errd.Error())
			}
		}
	}
}
