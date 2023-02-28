package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// RemoveGeneratedGongFilesButDocs generates the setup file for the gorm
func RemoveGeneratedGongFilesButDocs(
	RelativePkgPath string) {

	{
		filename := filepath.Join(RelativePkgPath, "docs.go")
		file, err := os.Create(filename)

		if err != nil {
			file, err = os.Create(filename)
		}

		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, DefaultModelDocsTemplate)

	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_coder.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_ast.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_serialize.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_marshall.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}

	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_graph.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}
	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_enum.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}
	{
		// relative to the models package, swith to ./controlers package
		filename := filepath.Join(RelativePkgPath, "gong_callbacks.go")

		// we should use go generate
		log.Println("removing file : " + filename)

		if err := os.Remove(filename); err != nil {
			if os.IsExist(err) {
				log.Fatalf("Unable to remove %s", filename)
			}
		}
	}
}
