package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ProductAddButtonProxy struct {
	stager   *Stager
	products []*Product
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *ProductAddButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	product := (&Product{
		Name: "New Product",
	}).Stage(p.stager.stage)

	p.products = append(p.products, product)

	p.stager.stage.Commit()
}

var productNumber int
