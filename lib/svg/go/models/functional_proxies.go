package models

// Define a function signature that matches the interface method logic
type SvgURectpdatedFunc func(frontRect *Rect)

// Generic Proxy that implements models.ButtonImplInterface
type FunctionalSvgRectProxy struct {
	OnUpdated SvgURectpdatedFunc
}

// Implement the interface method
func (p *FunctionalSvgRectProxy) RectUpdated(frontRect *Rect) {
	if p.OnUpdated != nil {
		p.OnUpdated(frontRect)
	}
}

// Define a function signature that matches the interface method logic
type SvgULinkpdatedFunc func(frontLink *Link)

// Generic Proxy that implements models.ButtonImplInterface
type FunctionalSvgLinkProxy struct {
	OnUpdated SvgULinkpdatedFunc
}

// Implement the interface method
func (p *FunctionalSvgLinkProxy) LinkUpdated(frontLink *Link) {
	if p.OnUpdated != nil {
		p.OnUpdated(frontLink)
	}
}
