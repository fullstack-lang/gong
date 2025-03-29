package main

import (
	"flag"
	"log"
	"strconv"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	split_static "github.com/fullstack-lang/gong/lib/split/go/static"

	button_models "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	cursor_models "github.com/fullstack-lang/gong/lib/cursor/go/models"
	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"

	slider_models "github.com/fullstack-lang/gong/lib/slider/go/models"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	xlsx_models "github.com/fullstack-lang/gong/lib/xlsx/go/models"
	xlsx_stack "github.com/fullstack-lang/gong/lib/xlsx/go/stack"

	tone_stack "github.com/fullstack-lang/gong/lib/tone/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	issue564         = flag.Bool("issue564", false, "generating issue564")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("split: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := split_static.ServeStaticFiles(*logGINFlag)

	// setup root stack
	stack := split_stack.NewStack(r, "", "", "", "", *embeddedDiagrams, true)
	splitStage := stack.Stage
	stack.Probe.Refresh()

	sliderStage1 := slider_stack.NewStack(r, "slider 1", "", "", "", true, true).Stage

	sliderStage2 := slider_stack.NewStack(r, "slider 2", "", "", "", true, true).Stage

	buttonStackName := "button"
	stackbutton := button_stack.NewStack(r, buttonStackName, "", "", "", true, true)
	buttonStage := stackbutton.Stage

	toneStackName := "tone"
	stacktone := tone_stack.NewStack(r, toneStackName, "", "", "", true, true)
	toneStage := stacktone.Stage

	xlsxStackName := "xlsx"
	stackxlsx := xlsx_stack.NewStack(r, xlsxStackName, "", "", "", true, true)
	xlsxStage := stackxlsx.Stage

	{

		layout := new(slider_models.Layout).Stage(sliderStage1)
		group := new(slider_models.Group).Stage(sliderStage1)
		group.Percentage = 100
		layout.Groups = append(layout.Groups, group)

		slider := new(slider_models.Slider).Stage(sliderStage1)
		slider.IsInt = true
		slider.MinInt = 10
		slider.MaxInt = 100
		slider.StepInt = 5
		slider.Name = "example"
		group.Sliders = append(group.Sliders, slider)

		sliderStage1.Commit()
	}

	{

		layout := new(slider_models.Layout).Stage(sliderStage2)
		group := new(slider_models.Group).Stage(sliderStage2)
		group.Percentage = 100
		layout.Groups = append(layout.Groups, group)

		slider := new(slider_models.Slider).Stage(sliderStage2)
		slider.IsInt = true
		slider.MinInt = 10
		slider.MaxInt = 100
		slider.StepInt = 5
		slider.Name = "example"
		group.Sliders = append(group.Sliders, slider)

		sliderStage2.Commit()
	}

	{

		layout := new(button_models.Layout).Stage(buttonStage)
		group := new(button_models.Group).Stage(buttonStage)
		group.Percentage = 100
		layout.Groups = append(layout.Groups, group)

		button := new(button_models.Button).Stage(buttonStage)
		button.Name = "example"
		button.Icon = "draw"
		button.Label = "Draw a Phyllotaxy Growth Curve"
		group.Buttons = append(group.Buttons, button)

		buttonStage.Commit()
	}

	{

		toneStage.Commit()
	}

	{
		cursorStackName := "cursor"
		stackcursor := cursor_stack.NewStack(r, cursorStackName, "", "", "", *embeddedDiagrams, true)
		cursorStage := stackcursor.Stage

		cursor := (&cursor_models.Cursor{Name: "cursor"}).Stage(cursorStage)

		cursor.Y1 = 100
		cursor.Y2 = 200

		cursor.StartX,
			cursor.EndX = 20, 20
		cursor.DurationSeconds = 7
		cursor.Stroke = "red"
		cursor.StrokeOpacity = 1.0
		cursor.StrokeWidth = 2.0
		cursor.IsPlaying = true

		cursorStage.Commit()

		svgStackName := "svg"
		stacksvg := svg_stack.NewStack(r, svgStackName, "", "", "", *embeddedDiagrams, true)
		svgStage := stacksvg.Stage

		svg := (&svg_models.SVG{Name: "svg"}).Stage(svgStage)
		layer := (&svg_models.Layer{Name: "layer"}).Stage(svgStage)
		svg.Layers = append(svg.Layers, layer)

		rect := (&svg_models.Rect{Name: "rect"}).Stage(svgStage)
		layer.Rects = append(layer.Rects, rect)
		rect.X = 10
		rect.Y = 10
		rect.Width = 200
		rect.Height = 100
		rect.Stroke = "black"
		rect.StrokeOpacity = 1.0
		rect.StrokeWidth = 2.0

		svgStage.Commit()
	}

	if false {
		stack.Stage.Checkout()

		// fetch the view tone probe view
		key := "view of tone probe"
		mapView := *split.GetGongstructInstancesMap[split.View](stack.Stage)

		viewToneProbe, ok := mapView[key]
		if !ok {
			log.Fatalln("view not found")
		}
		// remove the tone probe area from the probe view and commit
		asSplitArea := viewToneProbe.RootAsSplitAreas[0]
		if asSplitArea == nil {
			log.Fatalln("split area not found")
		}
		viewToneProbe.RootAsSplitAreas = viewToneProbe.RootAsSplitAreas[:]
		stack.Stage.Commit()

		// add the tone probe area from the probe view and commit
		viewToneProbe.RootAsSplitAreas = append(viewToneProbe.RootAsSplitAreas, asSplitArea)
		stack.Stage.Commit()
	}

	{

		sampleXLFile := new(xlsx_models.XLFile).Stage(xlsxStage)
		sampleXLFile.Open(xlsxStage, "data/samplefile.xlsx")

		xlsxStage.Commit()
		stackxlsx.Probe.Refresh()
	}

	(&split.View{
		Name: "Probe Split",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: splitStage.GetProbeSplitStageName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Slider 1",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Slider: (&split.Slider{
					StackName: sliderStage1.GetName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Probe Slider 1",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: sliderStage1.GetProbeSplitStageName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Slider 2",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Slider: (&split.Slider{
					StackName: sliderStage2.GetName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Probe Slider 2",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: sliderStage1.GetProbeSplitStageName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Button",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Button: (&split.Button{
					StackName: buttonStage.GetName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Probe Button",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: buttonStage.GetProbeSplitStageName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Xlsx",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Xlsx: (&split.Xlsx{
					StackName: xlsxStage.GetName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Probe Xlsx",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: xlsxStage.GetProbeSplitStageName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Spliting views",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				AsSplit: (&split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name:             "Top",
							Size:             50,
							ShowNameInHeader: true,
							AsSplit: (&split.AsSplit{
								Name:      "as split",
								Direction: split.Horizontal,
								AsSplitAreas: []*split.AsSplitArea{
									(&split.AsSplitArea{
										Name:             "Top",
										Size:             50,
										ShowNameInHeader: true,
									}).Stage(splitStage),
									(&split.AsSplitArea{
										Name:             "Bottom",
										Size:             50,
										ShowNameInHeader: true,
									}).Stage(splitStage),
								},
							}).Stage(splitStage),
						}).Stage(splitStage),
						(&split.AsSplitArea{
							Name:             "Bottom",
							Size:             50,
							ShowNameInHeader: true,
							AsSplit: (&split.AsSplit{
								Name:      "as split",
								Direction: split.Horizontal,
								AsSplitAreas: []*split.AsSplitArea{
									(&split.AsSplitArea{
										Name:             "Top",
										Size:             50,
										ShowNameInHeader: true,
									}).Stage(splitStage),
									(&split.AsSplitArea{
										Name:             "Bottom",
										Size:             50,
										ShowNameInHeader: true,
									}).Stage(splitStage),
								},
							}).Stage(splitStage),
						}).Stage(splitStage),
					},
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
