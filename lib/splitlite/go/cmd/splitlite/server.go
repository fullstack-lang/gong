//go:build !js

package main

import (
	"log"
	"strconv"

	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
	splitlite_stack "github.com/fullstack-lang/gong/lib/splitlite/go/stack"
	splitlite_static "github.com/fullstack-lang/gong/lib/splitlite/go/static"

	button_models "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

func executeServer() {
	r := splitlite_static.ServeStaticFiles(logGINFlag)

	stack := splitlite_stack.NewStack(r, "", "", "", "", embeddedDiagrams, true)
	splitliteStage := stack.Stage
	stack.Probe.Refresh()

	(&splitlite.Title{Name: "Splitlite Test"}).Stage(splitliteStage)
	(&splitlite.FavIcon{
		Name: "Test",
		SVG: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none">
			<circle cx="12" cy="12" r="10" fill="#007bff"/>
		</svg>`,
	}).Stage(splitliteStage)

	buttonStackName := "button"
	stackbutton := button_stack.NewStack(r, buttonStackName, "", "", "", true, true)
	buttonStage := stackbutton.Stage

	{
		layout := new(button_models.Layout).Stage(buttonStage)
		group := new(button_models.Group).Stage(buttonStage)
		group.Percentage = 100
		layout.Groups = append(layout.Groups, group)

		btn := new(button_models.Button).Stage(buttonStage)
		btn.Name = "example"
		btn.Icon = "draw"
		btn.Label = "Example Button"
		group.Buttons = append(group.Buttons, btn)

		buttonStage.Commit()
	}

	svgStackName := "svg"
	stacksvg := svg_stack.NewStack(r, svgStackName, "", "", "", true, true)
	svgStage := stacksvg.Stage

	{
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

	(&splitlite.View{
		Name: "Main View",
		RootAsSplitAreas: []*splitlite.AsSplitArea{
			{
				Button: &splitlite.Button{
					StackName: buttonStage.GetName(),
				},
			},
			{
				Svg: &splitlite.Svg{
					StackName: svgStage.GetName(),
				},
			},
		},
		IsSelectedView: true,
	}).Stage(splitliteStage)

	splitliteStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := splitlite_static.RunServer(r, ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
