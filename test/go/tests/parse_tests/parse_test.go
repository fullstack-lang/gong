package parsetests_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fullstack-lang/gong/test/go/fullstack"
	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/gin-gonic/gin"
)

func TestParseTest(t *testing.T) {

	// setup GORM
	r := gin.Default()
	fullstack.Init(r)

	models.ParseAstFile(&models.Stage, "./input/stage.go")

	// serialize into another file
	out_file, err := os.Create(fmt.Sprintf("./out/%s.go", "out"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer out_file.Close()
	models.Stage.Marshall(out_file, "github.com/fullstack-lang/gong/test/go/models", "out")
}
