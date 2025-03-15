package main

import (
	"flag"
	"log"
	"strconv"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	split_static "github.com/fullstack-lang/gong/lib/split/go/static"

	slider_models "github.com/fullstack-lang/gong/lib/slider/go/models"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("split: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	log.Println("marshallOnCommit", *marshallOnCommit)
	log.Println("unmarshallFromCode", *unmarshallFromCode)

	// setup the static file server and get the controller
	r := split_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := split_stack.NewStack(r, "split", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	stackSlider := slider_stack.NewStack(r, "slider", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	sliderStage := stackSlider.Stage

	layout := new(slider_models.Layout).Stage(sliderStage)
	group := new(slider_models.Group).Stage(sliderStage)
	group.Percentage = 100
	layout.Groups = append(layout.Groups, group)

	slider := new(slider_models.Slider).Stage(sliderStage)
	slider.IsInt = true
	slider.MinInt = 10
	slider.MaxInt = 100
	slider.StepInt = 5
	slider.Name = "example"
	group.Sliders = append(group.Sliders, slider)

	sliderStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
