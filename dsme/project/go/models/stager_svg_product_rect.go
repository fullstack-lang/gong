package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

func (stager *Stager) svgProductRect(
	diagram *Diagram, // might be ni
	ProductShape *ProductShape,
	layer *svg.Layer,
) *svg.Rect {
	root := stager.root

	Product := ProductShape.Product

	rect := new(svg.Rect)
	rect.Name = Product.Name
	rect.Stroke = svg.Black.ToString()
	rect.StrokeWidth = 2
	rect.StrokeOpacity = 1
	rect.X = ProductShape.X
	rect.Y = ProductShape.Y
	rect.Width = ProductShape.Width
	rect.Height = ProductShape.Height
	rect.RX = 3

	// rect is editable if diagram is not null

	if diagram.IsEditable() {
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		rect.Impl = NewRectShape_ProductProxy(
			&ProductShape.RectShape,
			Product,
			stager,
		)
		// for allowing later Stage() on the rect shape
		ProductShape.receiver = ProductShape
	}

	ProductTitleText := new(svg.RectAnchoredText)
	{
		ProductTitleText.Name = Product.Name
		ProductTitleText.Content = Product.Name

		if rect.Width > 0 {
			ProductTitleText.Content = strutils.WrapString(ProductShape.Product.Name, int(rect.Width/root.NbPixPerCharacter))
		}
		ProductTitleText.Stroke = svg.Black.ToString()
		ProductTitleText.StrokeWidth = 1
		ProductTitleText.StrokeOpacity = 1
		ProductTitleText.Color = svg.Black.ToString()
		ProductTitleText.FillOpacity = 1

		ProductTitleText.FontSize = "16px"
		ProductTitleText.X_Offset = 0
		ProductTitleText.Y_Offset = 30
		ProductTitleText.RectAnchorType = svg.RECT_TOP
		ProductTitleText.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, ProductTitleText)
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}

func NewRectShape_ProductProxy(
	ProductShapeInterface RectShapeInterface,
	Product *Product,
	stager *Stager,
) (proxy *ShapeRect_ProductProxy) {
	proxy = new(ShapeRect_ProductProxy)
	proxy.rectShapeInterface = ProductShapeInterface
	proxy.Product = Product
	proxy.stager = stager
	return
}

type ShapeRect_ProductProxy struct {
	rectShapeInterface RectShapeInterface
	Product            *Product
	stager             *Stager
}

// RectUpdated implements models.RectImplInterface.
func (p *ShapeRect_ProductProxy) RectUpdated(updatedRect *svg.Rect) {
	diffSize := p.rectShapeInterface.GetWidth() != updatedRect.Width ||
		p.rectShapeInterface.GetHeight() != updatedRect.Height

	diffPosition := p.rectShapeInterface.GetX() != updatedRect.X ||
		p.rectShapeInterface.GetY() != updatedRect.Y

	p.rectShapeInterface.SetX(updatedRect.X)
	p.rectShapeInterface.SetY(updatedRect.Y)
	p.rectShapeInterface.SetWidth(updatedRect.Width)
	p.rectShapeInterface.SetHeight(updatedRect.Height)

	if !diffSize && !diffPosition {
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.probeForm.FillUpFormFromGongstruct(p.Product, "Product")
	}

	if diffPosition {
		// Issue #7, this will allow multiple rect to be moved together
		p.stager.stage.CommitWithSuspendedCallbacks()
	}

	if diffSize {
		p.stager.stage.Commit()
	}
}
