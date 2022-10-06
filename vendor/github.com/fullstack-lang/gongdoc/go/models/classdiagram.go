package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
	"github.com/tdewolff/canvas/renderers/svg"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/static"
	"github.com/fullstack-lang/gongdoc/go/walk"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// list of classshapes in the diagram
	Classshapes []*Classshape

	// list of notes in the diagram
	Notes []*Note

	// IsEditable indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsEditable bool
}

const DiagramMarginX = 10.0
const DiagramMarginY = 10.0

// Extent compute max X and max Y
func (classdiagram *Classdiagram) Extent() (x, y, maxClassshapeHeigth float64) {

	for _, classshape := range classdiagram.Classshapes {

		if maxClassshapeHeigth < classshape.Heigth {
			maxClassshapeHeigth = classshape.Heigth
		}

		if classshape.Position.X+classshape.Width > x {
			x = classshape.Position.X + classshape.Width
		}
		if classshape.Position.Y+classshape.Heigth > y {
			y = classshape.Position.Y + classshape.Heigth
		}

		for _, link := range classshape.Links {
			if link.Middlevertice.X > x {
				x = link.Middlevertice.X
			}
			if link.Middlevertice.Y > y {
				y = link.Middlevertice.Y
			}
		}
	}
	// margin
	x += DiagramMarginX
	y += DiagramMarginY
	return
}

// MarshallAsVariable transform the class diagram into a var
func (classdiagram *Classdiagram) MarshallAsVariable(file *os.File) error {

	fmt.Fprintf(file, "var %s uml.Classdiagram = uml.Classdiagram{\n", classdiagram.Name)

	fmt.Fprintf(file, "\tClassshapes: []*uml.Classshape{\n")

	if len(classdiagram.Classshapes) > 0 {
		// sort Classshapes
		sort.Slice(classdiagram.Classshapes[:], func(i, j int) bool {
			return classdiagram.Classshapes[i].Structname < classdiagram.Classshapes[j].Structname
		})
		for _, classshape := range classdiagram.Classshapes {
			classshape.Marshall(file, 2)
			fmt.Fprintf(file, ",\n")
		}
	}
	fmt.Fprintf(file, "\t},\n")

	fmt.Fprintf(file, "\tNotes: []*uml.Note{\n")

	if len(classdiagram.Notes) > 0 {
		// sort Notes
		sort.Slice(classdiagram.Notes[:], func(i, j int) bool {
			return classdiagram.Notes[i].Body < classdiagram.Notes[j].Body
		})
		for _, note := range classdiagram.Notes {
			note.Marshall(file, 2)
			fmt.Fprintf(file, ",\n")
		}
	}
	fmt.Fprintf(file, "\t},\n")

	fmt.Fprintf(file, "}\n")
	return nil
}

// ClassdiagramMap is a Map of all Classdiagram via their Name
type ClassdiagramMap map[string]*Classdiagram

// ClassdiagramStore is a handy ClassdiagramMap
var ClassdiagramStore ClassdiagramMap = make(map[string]*Classdiagram, 0)

// Unmarshall updates a classdiagram values from an ast.Epr
// and appends it to the ClassdiagramStore
func (classdiagram *Classdiagram) Unmarshall(modelPkg *gong_models.ModelPkg, expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// "uml.Classdiagram{
	//   	Classshapes: []*uml.Classshape{ ...
	//		Links: []*uml.Links{ ...
	if cl, ok := expr.(*ast.CompositeLit); ok {

		// parse all KeyValues of the Classdiagram
		for _, elt := range cl.Elts {
			if structvaluekeyexpr, ok := elt.(*ast.KeyValueExpr); !ok {
				log.Panic("Expression should be a struct key" +
					fset.Position(structvaluekeyexpr.Pos()).String())
			} else {

				if key, ok := structvaluekeyexpr.Key.(*ast.Ident); !ok {
					log.Panic("Key shoud be an ident" +
						fset.Position(key.Pos()).String())
				} else {
					switch key.Name {
					case "Classshapes":
						var cl *ast.CompositeLit
						var ok bool
						if cl, ok = structvaluekeyexpr.Value.(*ast.CompositeLit); !ok {
							log.Panic("Value shoud be a composite lit" +
								fset.Position(structvaluekeyexpr.Pos()).String())
						}
						// get the array of shapes (either as definition or as reference)
						for _, expr := range cl.Elts {

							var classshape *Classshape
							switch exp := expr.(type) {
							case *ast.UnaryExpr: // this is a reference to a variable
								if ident, ok := exp.X.(*ast.Ident); !ok {
									log.Panic("" + fset.Position(exp.Pos()).String())
								} else {
									log.Printf("found %s", ident.Name)

									for _classshape := range Stage.Classshapes {
										if _classshape.Name == ident.Name {
											classshape = _classshape
											continue
										}
									}
									if classshape == nil {
										log.Panic("Unable to find the shape " + ident.Name + " " +
											fset.Position(cl.Pos()).String())
									}
								}
							case *ast.CompositeLit: // this is a definition
								classshape = new(Classshape)
								classshape.Unmarshall(modelPkg, exp, fset)
							default:
								log.Panic("Value shoud be a composite lit or a unary" +
									fset.Position(structvaluekeyexpr.Pos()).String())
							}

							classdiagram.Classshapes = append(classdiagram.Classshapes, classshape)
						}
					case "Notes":
						var cl *ast.CompositeLit
						var ok bool
						if cl, ok = structvaluekeyexpr.Value.(*ast.CompositeLit); !ok {
							log.Panic("Value shoud be a composite lit" +
								fset.Position(structvaluekeyexpr.Pos()).String())
						}
						for _, expr := range cl.Elts {

							var note *Note
							switch exp := expr.(type) {
							case *ast.UnaryExpr: // this is a reference to a variable
								if ident, ok := exp.X.(*ast.Ident); !ok {
									log.Panic("" + fset.Position(exp.Pos()).String())
								} else {
									log.Printf("found %s", ident.Name)
								}
							case *ast.CompositeLit: // this is a definition
								note = new(Note)
								note.Unmarshall(modelPkg, exp, fset)
							default:
								log.Panic("Value shoud be a composite lit or a unary" +
									fset.Position(structvaluekeyexpr.Pos()).String())
							}

							// of the note is matched, add it to the diagram
							if note.Matched {
								classdiagram.Notes = append(classdiagram.Notes, note)
							}
						}
					case "Name":
						// already initialized
					default:
						log.Panic("Key shoud be Classshapes, Field or Link" +
							fset.Position(key.Pos()).String())
					}
				}

			}

		}
	}
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (classdiagram *Classdiagram) SerializeToStage() {

	classdiagram.Stage()

	for _, classshape := range classdiagram.Classshapes {
		classshape.SerializeToStage()
	}
	for _, note := range classdiagram.Notes {
		note.Stage()
	}
}

// ModelSpaceToSVGPaperYCoord convert the Y coordinates from the model space
// to the SVG space
// in model space, Y coordinates are descending. In SVG, Y coordinates are
// ascending.
// You need to have the paper heigth mind
var SVGPaperHeight float64

// MaxClassshapeHeigth because you also need the max heigth of all classshapes
var MaxClassshapeHeigth float64

// ModelToSVGYCoord converts a Y coordinate in the model to a Y coordinate in the SVG canvas
func ModelToSVGYCoord(yModel float64) (ySVG float64) {
	return SVGPaperHeight - yModel
}

// ModelToSVGRectangleYOrigin converts a top left Y coordinate of a rectangle in the model to a
// bottom left Y coordinate in the SVG canvas
func ModelToSVGRectangleYOrigin(yModel, classshapeHeigth float64) (ySVG float64) {
	return ModelToSVGYCoord(yModel + classshapeHeigth)
}

func (classdiagram *Classdiagram) OutputSVG(path string) {

	var maxx, maxy float64
	maxx, maxy, MaxClassshapeHeigth = classdiagram.Extent()
	var width float64
	width, SVGPaperHeight = maxx, maxy

	c := canvas.New(width, SVGPaperHeight)
	ctx := canvas.NewContext(c)
	ctx.SetStrokeColor(color.Black)

	mapStringClassshape := make(map[string]*Classshape)
	for _, classshape := range classdiagram.Classshapes {
		mapStringClassshape[classshape.Structname] = classshape
	}

	dejaVuSerif := canvas.NewFontFamily("dejavu-serif")
	staticFontFile := []byte(static.Files["../gongdoc/DejaVuSerif.ttf"])
	if err := dejaVuSerif.LoadFont(
		staticFontFile,
		0,
		canvas.FontRegular); err != nil { // TTF, OTF, WOFF, or WOFF2
		log.Panic("cannot load font", err.Error())
	}
	ff := dejaVuSerif.Face(24.0, color.Black, canvas.FontBlack, canvas.FontNormal)

	// First, draw all links with full transparency
	ctx.SetFillColor(color.Transparent)
	for _, classshape := range classdiagram.Classshapes {

		bottomLeftX := classshape.Position.X
		bottomLeftY := ModelToSVGRectangleYOrigin(classshape.Position.Y, classshape.Heigth)

		for _, link := range classshape.Links {
			// tech target point of link

			targetClassshape := mapStringClassshape[link.Fieldtypename]

			// sometimes, links to not refer to classhape that are existing on the diagram (for instance Interfaces)
			if targetClassshape == nil {
				continue
			}
			targetPosition := targetClassshape.Position

			// from center to center
			sourceClassshapeToMiddleVerticePath := &canvas.Path{}

			var verticeSVGX = link.Middlevertice.X
			var verticeSVGY = ModelToSVGYCoord(link.Middlevertice.Y)

			var centerOfClassshapeX = bottomLeftX + classshape.Width/2.0
			var centerOfClassshapeY = bottomLeftY + classshape.Heigth/2.0

			// draw with middle vertice
			linkToMiddleVerticePathInDiagram := Position{
				X: verticeSVGX - centerOfClassshapeX,
				Y: verticeSVGY - centerOfClassshapeY,
			}
			sourceClassshapeToMiddleVerticePath.LineTo(linkToMiddleVerticePathInDiagram.X, linkToMiddleVerticePathInDiagram.Y)

			ctx.DrawPath(centerOfClassshapeX, centerOfClassshapeY, sourceClassshapeToMiddleVerticePath)

			targetClassshapeToMiddleVerticePath := &canvas.Path{}

			targetBottomLeftX := targetPosition.X
			targetBottomLeftY := ModelToSVGRectangleYOrigin(targetPosition.Y, classshape.Heigth)

			centerOfClassshapeX = targetBottomLeftX + classshape.Width/2.0
			centerOfClassshapeY = targetBottomLeftY + classshape.Heigth/2.0

			// draw with middle vertice
			linkToMiddleVerticePathInDiagram = Position{
				X: verticeSVGX - centerOfClassshapeX,
				Y: verticeSVGY - centerOfClassshapeY,
			}
			targetClassshapeToMiddleVerticePath.LineTo(linkToMiddleVerticePathInDiagram.X, linkToMiddleVerticePathInDiagram.Y)

			ctx.DrawPath(centerOfClassshapeX, centerOfClassshapeY, targetClassshapeToMiddleVerticePath)

			// draw text between middle vertice & target position
			text := canvas.NewTextLine(ff, link.Fieldname, canvas.TextAlign(canvas.Left)) // simple text line
			linkFieldPosition := Position{
				X: (verticeSVGX + centerOfClassshapeX) / 2.0,
				Y: (verticeSVGY + centerOfClassshapeY) / 2.0,
			}
			ctx.DrawText(linkFieldPosition.X, linkFieldPosition.Y, text)
		}
	}

	// draw the classshape AFTER the links
	ctx.SetFillColor(color.RGBA{204, 224, 218, 255})

	for _, classshape := range classdiagram.Classshapes {

		heigthBetweenLines := 16.0

		bottomLeftX := classshape.Position.X
		bottomLeftY := ModelToSVGRectangleYOrigin(classshape.Position.Y, classshape.Heigth)

		text := canvas.NewTextLine(ff, classshape.Structname, canvas.TextAlign(canvas.Left)) // simple text line
		p := canvas.Rectangle(classshape.Width, classshape.Heigth)

		// draw the line below the name of the class
		p.MoveTo(0, classshape.Heigth-heigthBetweenLines)
		p.LineTo(classshape.Width, classshape.Heigth-heigthBetweenLines)
		ctx.DrawPath(bottomLeftX, bottomLeftY, p)
		// draw the name of the class
		ctx.DrawText(bottomLeftX+10, bottomLeftY+classshape.Heigth-heigthBetweenLines+6, text)

		// draw fields name
		for i, field := range classshape.Fields {
			text := canvas.NewTextLine(ff,
				field.Fieldname+":"+field.Fieldtypename,
				canvas.TextAlign(canvas.Left)) // simple text line
			ctx.DrawText(bottomLeftX+10.0, bottomLeftY+classshape.Heigth-12.0-float64(i+1)*heigthBetweenLines, text)
		}
	}

	c.WriteFile(fmt.Sprintf(filepath.Join(path, "%s.png"), classdiagram.Name), rasterizer.PNGWriter(5.0))
	c.WriteFile(fmt.Sprintf(filepath.Join(path, "%s.svg"), classdiagram.Name), svg.Writer)
}

func (classDiagram *Classdiagram) Marshall(pkgelt *Pkgelt, pkgPath string) error {

	// open file

	log.SetFlags(log.Lshortfile)
	filename := walk.CaptureOutput(func() { log.Printf("") })
	log.SetFlags(log.LstdFlags)
	prelude := strings.ReplaceAll(preludeRef, "{{filename}}", filename)
	prelude = strings.ReplaceAll(prelude, "{{ClassdiagramName}}", classDiagram.Name)
	if len(classDiagram.Classshapes) > 0 {
		prelude = strings.ReplaceAll(prelude, "{{Imports}}", "\n\t\""+pkgelt.GongModelPath+"\"")
	} else {
		prelude = strings.ReplaceAll(prelude, "{{Imports}}", "")
	}
	prelude = strings.ReplaceAll(prelude, "docs", "models")

	filepath := filepath.Join(pkgPath, classDiagram.Name) + ".go"
	file, err := os.Create(filepath)
	defer closeFile(file)
	if err == nil {
		fmt.Fprintf(file, prelude)
	} else {
		cwd, _ := os.Getwd()
		log.Fatal("Cannot open file ", filepath, " from cwd ", cwd, ", Error is ", err)
	}
	if err2 := classDiagram.MarshallAsVariable(file); err != nil {
		return err2
	}

	return err
}
