package main

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

func genFlutter(modelPkg *gong_models.ModelPkg) {

	// check if "flutter" directory exist
	// if not, call "flutter create flutter" to create a flutter project
	var flutterCreateHasBeenPerformed bool
	{

		flutterProjectDir, err := filepath.Abs(filepath.Join(*pkgPath, "../../flutter_project"))
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
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

	if flutterCreateHasBeenPerformed {

	}
}
