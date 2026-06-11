package cmd

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the module to the latest gong commit on the main branch",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gong: ")
		log.SetFlags(0)

		log.Println("Updating gong to the latest commit on main branch")

		// go get github.com/fullstack-lang/gong@main
		goGetCmd := exec.Command("go", "get", "github.com/fullstack-lang/gong@main")

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		goGetCmd.Stdout = mw
		goGetCmd.Stderr = mw

		log.Println(goGetCmd.String())

		if err := goGetCmd.Run(); err != nil {
			log.Panic(err)
		}

		// go mod tidy
		goModTidyCmd := exec.Command("go", "mod", "tidy")
		goModTidyCmd.Stdout = mw
		goModTidyCmd.Stderr = mw

		log.Println(goModTidyCmd.String())

		if err := goModTidyCmd.Run(); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
